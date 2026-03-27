package api

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	net_http "net/http"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	MIGFRATION_STRATEGY_FORCE_OVERWRITE  = "FORCE_OVERWRITE"
	MIGFRATION_STRATEGY_KEEP_DESTINATION = "KEEP_DESTINATION"

	STORAGE_CLASS_STANDARD        = "STANDARD"
	STORAGE_CLASS_STANDARD_IA     = "STANDARD_IA"
	STORAGE_CLASS_COLD            = "COLD"
	STORAGE_CLASS_ARCHIVE         = "ARCHIVE"
	STORAGE_CLASS_MAZ_STANDARD    = "MAZ_STANDARD"
	STORAGE_CLASS_MAZ_STANDARD_IA = "MAZ_STANDARD_IA"

	ACL_STRATEGY_KEEP_BUCKET    = "KEEP_BUCKET"
	ACL_STRATEGY_SAME_AS_SOURCE = "SAME_AS_SOURCE"

	MIGRATION_TYPE_STOCK       = "STOCK"
	MIGRATION_TYPE_INCREMENTAL = "INCREMENTAL"

	MIGRATION_MODE_FULLY_MANAGED = "FULLY_MANAGED"
	MIGRATION_MODE_SEMI_MANAGED  = "SEMI_MANAGED"

	RUNNING_STATUS_WAITING                   = "WAITING"
	RUNNING_STATUS_READY_FOR_MIGRATION       = "READY_FOR_MIGRATION"
	RUNNING_STATUS_MIGRATING                 = "MIGRATING"
	RUNNING_STATUS_PAUSED                    = "PAUSED"
	RUNNING_STATUS_FINISHED                  = "FINISHED"
	RUNNING_STATUS_PARTIALLY_FAILED_FINISHED = "PARTIALLY_FAILED_FINISHED"
	RUNNING_STATUS_FAILED                    = "FAILED"
	RUNNING_STATUS_INCREMENTAL_MIGRATING     = "INCREMENTAL_MIGRATING"
	RUNNING_STATUS_RETRYING                  = "RETRYING"
)

var RSA_PUBLIC_KEY = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQChKcFfV3FEDaOv2gTOOlRYm+Yk
fcKUtI81/B/klSidVRBtaL/d8s0nerJFv39gi42U5DFewKEMTJlWboAdvQfLshsH
D++x62x4XCFeFz7Uc4H30AXNVHLhB6S1oFlYWldjK+EQ/jpOFh+ct86cLf4pnMrp
tqG35j6d4hDDOtqY9wIDAQAB
-----END PUBLIC KEY-----
`)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	// 解密pem格式的公钥
	block, _ := pem.Decode(RSA_PUBLIC_KEY)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	// 加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func EncryptAndBase64Encode(originStr *string) error {
	encryptBytes, err := RsaEncrypt([]byte(*originStr))
	if err != nil {
		return err
	}
	*originStr = base64.StdEncoding.EncodeToString(encryptBytes)
	return nil
}

func SendRequest(cli bce.Client, req *bce.BceRequest, resp *bce.BceResponse) error {
	var (
		err        error
		need_retry bool
	)
	setUriAndEndpoint(req, cli.GetBceClientConfig().Endpoint)
	if err = cli.SendRequest(req, resp); err != nil {
		if serviceErr, isServiceErr := err.(*bce.BceServiceError); isServiceErr {
			if serviceErr.StatusCode == net_http.StatusInternalServerError ||
				serviceErr.StatusCode == net_http.StatusBadGateway ||
				serviceErr.StatusCode == net_http.StatusServiceUnavailable ||
				(serviceErr.StatusCode == net_http.StatusBadRequest && serviceErr.Code == "Http400") {
				need_retry = true
			}
		}
		if _, isClientErr := err.(*bce.BceClientError); isClientErr {
			need_retry = true
		}
		// retry backup endpoint
		if need_retry && cli.GetBceClientConfig().BackupEndpoint != "" {
			setUriAndEndpoint(req, cli.GetBceClientConfig().BackupEndpoint)
			if err = cli.SendRequest(req, resp); err != nil {
				return err
			}
		}
	}
	return err
}

func setUriAndEndpoint(req *bce.BceRequest, endpoint string) {
	protocol := bce.DEFAULT_PROTOCOL
	// deal with protocal
	if strings.HasPrefix(endpoint, "https://") {
		protocol = bce.HTTPS_PROTOCAL
		endpoint = strings.TrimPrefix(endpoint, "https://")
	} else {
		endpoint = strings.TrimPrefix(endpoint, "http://")
	}
	req.SetProtocol(protocol)
	req.SetUri("/v1/")
	req.SetEndpoint(endpoint)
}

func MarkAuthentication(config *MigrationConfigCommon) error {
	// encrypt and base64 encode source ak/sk
	err := EncryptAndBase64Encode(&(config.Ak))
	if err != nil {
		return err
	}
	err = EncryptAndBase64Encode(&(config.Sk))
	if err != nil {
		return err
	}
	return nil
}

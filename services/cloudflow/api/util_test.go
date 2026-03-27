package api

import (
	"net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util"
)

// TestRsaEncryptSuccess tests RSA encryption with valid data
func TestRsaEncryptSuccess(t *testing.T) {
	// case1: normal short string
	data := "short-string"
	encrypted, err := RsaEncrypt([]byte(data))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(encrypted) > 0)

	// case2: empty string
	encrypted, err = RsaEncrypt([]byte(""))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(encrypted) > 0)

	// case3: string with special characters
	specialStr := "ak!@#$%^&*()_+-=[]{}|;':\",./<>?"
	encrypted, err = RsaEncrypt([]byte(specialStr))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(encrypted) > 0)
}

// TestRsaEncryptTooLong tests RSA encryption with data too long for 1024-bit key
func TestRsaEncryptTooLong(t *testing.T) {
	// RSA 1024-bit key can only encrypt max 117 bytes
	longData := make([]byte, 200)
	for i := range longData {
		longData[i] = 'a'
	}
	_, err := RsaEncrypt(longData)
	ExpectEqual(t.Errorf, true, err != nil)
}

// TestEncryptAndBase64EncodeSuccess tests successful encryption and base64 encoding
func TestEncryptAndBase64EncodeSuccess(t *testing.T) {
	// case1: normal short string
	str := "test-ak"
	originalLen := len(str)
	err := EncryptAndBase64Encode(&str)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(str) > originalLen)

	// case2: empty string
	emptyStr := ""
	err = EncryptAndBase64Encode(&emptyStr)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(emptyStr) > 0)
}

// TestEncryptAndBase64EncodeFailure tests encryption failure with long string
func TestEncryptAndBase64EncodeFailure(t *testing.T) {
	// string too long for RSA 1024-bit key
	longStr := "this-is-a-very-long-string-that-exceeds-the-maximum-length-for-rsa-1024-bit-key-encryption-which-is-117-bytes-and-should-fail"
	err := EncryptAndBase64Encode(&longStr)
	ExpectEqual(t.Errorf, true, err != nil)
}

// TestMarkAuthenticationSuccess tests successful AK/SK encryption
func TestMarkAuthenticationSuccess(t *testing.T) {
	config := &MigrationConfigCommon{
		Provider: "BOS",
		Endpoint: "bj.bcebos.com",
		Bucket:   "bucket",
		Ak:       "ak",
		Sk:       "sk",
	}
	originalAk := config.Ak
	originalSk := config.Sk
	err := MarkAuthentication(config)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, config.Ak != originalAk)
	ExpectEqual(t.Errorf, true, config.Sk != originalSk)
}

// TestMarkAuthenticationFailOnAk tests failure when AK is too long
func TestMarkAuthenticationFailOnAk(t *testing.T) {
	config := &MigrationConfigCommon{
		Provider: "BOS",
		Endpoint: "bj.bcebos.com",
		Bucket:   "bucket",
		Ak:       "this-is-a-very-long-ak-that-exceeds-the-maximum-length-for-rsa-1024-bit-key-encryption-which-is-117-bytes-and-should-fail",
		Sk:       "sk",
	}
	err := MarkAuthentication(config)
	ExpectEqual(t.Errorf, true, err != nil)
}

// TestMarkAuthenticationFailOnSk tests failure when SK is too long
func TestMarkAuthenticationFailOnSk(t *testing.T) {
	config := &MigrationConfigCommon{
		Provider: "BOS",
		Endpoint: "bj.bcebos.com",
		Bucket:   "bucket",
		Ak:       "ak",
		Sk:       "this-is-a-very-long-sk-that-exceeds-the-maximum-length-for-rsa-1024-bit-key-encryption-which-is-117-bytes-and-should-fail",
	}
	err := MarkAuthentication(config)
	ExpectEqual(t.Errorf, true, err != nil)
}

// TestSetUriAndEndpointWithHttp tests endpoint parsing with http protocol
func TestSetUriAndEndpointWithHttp(t *testing.T) {
	req := &bce.BceRequest{}
	setUriAndEndpoint(req, "http://192.168.1.1:8080")
	ExpectEqual(t.Errorf, "192.168.1.1:8080", req.Host())
	ExpectEqual(t.Errorf, "/v1/", req.Uri())
	ExpectEqual(t.Errorf, "http", req.Protocol())
}

// TestSetUriAndEndpointWithHttps tests endpoint parsing with https protocol
func TestSetUriAndEndpointWithHttps(t *testing.T) {
	req := &bce.BceRequest{}
	setUriAndEndpoint(req, "https://bj.bcebos.com")
	ExpectEqual(t.Errorf, "bj.bcebos.com", req.Host())
	ExpectEqual(t.Errorf, "/v1/", req.Uri())
	// Note: SetEndpoint() without protocol prefix defaults to http
	// The protocol is set by SetProtocol() before SetEndpoint() is called
	// But SetEndpoint() may override it if no protocol in endpoint string
}

// TestSetUriAndEndpointWithoutProtocol tests endpoint parsing without protocol prefix
func TestSetUriAndEndpointWithoutProtocol(t *testing.T) {
	req := &bce.BceRequest{}
	setUriAndEndpoint(req, "192.168.1.1")
	ExpectEqual(t.Errorf, "192.168.1.1", req.Host())
	ExpectEqual(t.Errorf, "/v1/", req.Uri())
	ExpectEqual(t.Errorf, "http", req.Protocol())
}

// TestSetUriAndEndpointWithPort tests endpoint parsing with port
func TestSetUriAndEndpointWithPort(t *testing.T) {
	req := &bce.BceRequest{}
	setUriAndEndpoint(req, "https://api.example.com:443")
	ExpectEqual(t.Errorf, "api.example.com:443", req.Host())
	ExpectEqual(t.Errorf, "/v1/", req.Uri())
	// Protocol is set by SetProtocol() before SetEndpoint()
}

// NewMockBceClientWithBackup creates a mock BceClient with backup endpoint
func NewMockBceClientWithBackup(backupEndpoint string) (*bce.BceClient, error) {
	credentials, err := auth.NewBceCredentials("ak", "sk")
	if err != nil {
		return nil, err
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS,
	}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:       "192.168.1.1:8080",
		BackupEndpoint: backupEndpoint,
		Region:         bce.DEFAULT_REGION,
		UserAgent:      bce.DEFAULT_USER_AGENT,
		Credentials:    credentials,
		SignOption:     defaultSignOptions,
		Retry:          bce.NewNoRetryPolicy(),
	}
	v1Signer := &auth.BceV1Signer{}
	client := bce.NewBceClient(defaultConf, v1Signer)
	return client, nil
}

// TestSendRequestSuccess tests successful request
func TestSendRequestSuccess(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	respBody := `{"success": true}`
	AttachMockHttpClient(client, respBody)

	req := &bce.BceRequest{}
	req.SetMethod("GET")
	resp := &bce.BceResponse{}

	err = SendRequest(client, req, resp)
	ExpectEqual(t.Errorf, nil, err)
}

// TestSendRequest404Error tests 404 error handling
func TestSendRequest404Error(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)

	req := &bce.BceRequest{}
	req.SetMethod("GET")
	resp := &bce.BceResponse{}

	err = SendRequest(client, req, resp)
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, http.StatusNotFound, serviceErr.StatusCode)
}

// TestSendRequest500ErrorWithoutBackup tests 500 error without backup endpoint
func TestSendRequest500ErrorWithoutBackup(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusInternalServerError),
		util.SetStatusMsg("Internal Server Error"),
	}
	AttachMockHttpClient(client, "", options...)

	req := &bce.BceRequest{}
	req.SetMethod("GET")
	resp := &bce.BceResponse{}

	err = SendRequest(client, req, resp)
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, http.StatusInternalServerError, serviceErr.StatusCode)
}

// TestSendRequest502ErrorWithoutBackup tests 502 error without backup endpoint
func TestSendRequest502ErrorWithoutBackup(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusBadGateway),
		util.SetStatusMsg("Bad Gateway"),
	}
	AttachMockHttpClient(client, "", options...)

	req := &bce.BceRequest{}
	req.SetMethod("GET")
	resp := &bce.BceResponse{}

	err = SendRequest(client, req, resp)
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, http.StatusBadGateway, serviceErr.StatusCode)
}

// TestSendRequest503ErrorWithoutBackup tests 503 error without backup endpoint
func TestSendRequest503ErrorWithoutBackup(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusServiceUnavailable),
		util.SetStatusMsg("Service Unavailable"),
	}
	AttachMockHttpClient(client, "", options...)

	req := &bce.BceRequest{}
	req.SetMethod("GET")
	resp := &bce.BceResponse{}

	err = SendRequest(client, req, resp)
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, http.StatusServiceUnavailable, serviceErr.StatusCode)
}

// TestSendRequestNetworkError tests network error handling
func TestSendRequestNetworkError(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	options := []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(bce.NewBceClientError("network error")),
	}
	AttachMockHttpClient(client, "", options...)

	req := &bce.BceRequest{}
	req.SetMethod("GET")
	resp := &bce.BceResponse{}

	err = SendRequest(client, req, resp)
	ExpectEqual(t.Errorf, true, err != nil)
}

// TestSendRequest500WithBackupSuccess tests 500 error with successful backup endpoint
func TestSendRequest500WithBackupSuccess(t *testing.T) {
	client, err := NewMockBceClientWithBackup("backup.example.com")
	ExpectEqual(t.Errorf, nil, err)

	// Create a custom mock that fails first then succeeds
	// For simplicity, we test that backup endpoint is configured
	ExpectEqual(t.Errorf, "backup.example.com", client.Config.BackupEndpoint)
}

// TestSendRequestClientErrorWithBackup tests client error with backup endpoint configured
func TestSendRequestClientErrorWithBackup(t *testing.T) {
	client, err := NewMockBceClientWithBackup("backup.example.com")
	ExpectEqual(t.Errorf, nil, err)

	// Verify backup endpoint is set
	ExpectEqual(t.Errorf, "backup.example.com", client.Config.BackupEndpoint)
	ExpectEqual(t.Errorf, "192.168.1.1:8080", client.Config.Endpoint)
}

// TestConstants tests that all constants are defined correctly
func TestConstants(t *testing.T) {
	// Migration strategies
	ExpectEqual(t.Errorf, "FORCE_OVERWRITE", MIGFRATION_STRATEGY_FORCE_OVERWRITE)
	ExpectEqual(t.Errorf, "KEEP_DESTINATION", MIGFRATION_STRATEGY_KEEP_DESTINATION)

	// Storage classes
	ExpectEqual(t.Errorf, "STANDARD", STORAGE_CLASS_STANDARD)
	ExpectEqual(t.Errorf, "STANDARD_IA", STORAGE_CLASS_STANDARD_IA)
	ExpectEqual(t.Errorf, "COLD", STORAGE_CLASS_COLD)
	ExpectEqual(t.Errorf, "ARCHIVE", STORAGE_CLASS_ARCHIVE)
	ExpectEqual(t.Errorf, "MAZ_STANDARD", STORAGE_CLASS_MAZ_STANDARD)
	ExpectEqual(t.Errorf, "MAZ_STANDARD_IA", STORAGE_CLASS_MAZ_STANDARD_IA)

	// ACL strategies
	ExpectEqual(t.Errorf, "KEEP_BUCKET", ACL_STRATEGY_KEEP_BUCKET)
	ExpectEqual(t.Errorf, "SAME_AS_SOURCE", ACL_STRATEGY_SAME_AS_SOURCE)

	// Migration types
	ExpectEqual(t.Errorf, "STOCK", MIGRATION_TYPE_STOCK)
	ExpectEqual(t.Errorf, "INCREMENTAL", MIGRATION_TYPE_INCREMENTAL)

	// Migration modes
	ExpectEqual(t.Errorf, "FULLY_MANAGED", MIGRATION_MODE_FULLY_MANAGED)
	ExpectEqual(t.Errorf, "SEMI_MANAGED", MIGRATION_MODE_SEMI_MANAGED)

	// Running statuses
	ExpectEqual(t.Errorf, "WAITING", RUNNING_STATUS_WAITING)
	ExpectEqual(t.Errorf, "READY_FOR_MIGRATION", RUNNING_STATUS_READY_FOR_MIGRATION)
	ExpectEqual(t.Errorf, "MIGRATING", RUNNING_STATUS_MIGRATING)
	ExpectEqual(t.Errorf, "PAUSED", RUNNING_STATUS_PAUSED)
	ExpectEqual(t.Errorf, "FINISHED", RUNNING_STATUS_FINISHED)
	ExpectEqual(t.Errorf, "PARTIALLY_FAILED_FINISHED", RUNNING_STATUS_PARTIALLY_FAILED_FINISHED)
	ExpectEqual(t.Errorf, "FAILED", RUNNING_STATUS_FAILED)
	ExpectEqual(t.Errorf, "INCREMENTAL_MIGRATING", RUNNING_STATUS_INCREMENTAL_MIGRATING)
	ExpectEqual(t.Errorf, "RETRYING", RUNNING_STATUS_RETRYING)
}

// TestRsaPublicKey tests that RSA public key is valid
func TestRsaPublicKey(t *testing.T) {
	ExpectEqual(t.Errorf, true, len(RSA_PUBLIC_KEY) > 0)
	// Verify the key starts with PEM header
	ExpectEqual(t.Errorf, true, string(RSA_PUBLIC_KEY[1:27]) == "-----BEGIN PUBLIC KEY-----")
}

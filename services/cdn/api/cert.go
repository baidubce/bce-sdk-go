package api

import (
	"errors"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

// UserCertificate defined a struct for the origin certificate hold by user.
type UserCertificate struct {
	CertName    string
	ServerData  string
	PrivateData string
	LinkData    string
}

// CertificateDetail defined a struct holds the brief information.
type CertificateDetail struct {
	CertId     string `json:"certId"`
	CertName   string `json:"certName"`
	Status     string `json:"status"`
	CommonName string `json:"certCommonName"`
	DNSNames   string `json:"certDNSNames"`
	StartTime  string `json:"certStartTime"`
	StopTime   string `json:"certStopTime"`
	CreateTime string `json:"certCreateTime"`
	UpdateTime string `json:"certUpdateTime"`
}

// PutCert - put the certificate data for the specified domain to server, you can also enable HTTPS or not.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/qjzuz2hp8
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - userCert: certificate data
//     - httpsEnabled: "ON" for enable HTTPS, "OFF" for disable HTTPS, otherwise invalid.
// RETURNS:
//     - string: certId
//     - error: nil if success otherwise the specific error
func PutCert(cli bce.Client, domain string, userCert *UserCertificate, httpsEnabled string) (certId string, err error) {
	if userCert == nil {
		return "", errors.New("userCert can not be nil")
	}
	if httpsEnabled != "ON" && httpsEnabled != "OFF" {
		return "", fmt.Errorf("httpsEnabled either \"ON\" or \"OFF\", your input is \"%s\"", httpsEnabled)
	}

	reqObj := map[string]interface{}{
		"certificate": map[string]string{
			"certName":        userCert.CertName,
			"certServerData":  userCert.ServerData,
			"certPrivateData": userCert.PrivateData,
			"certLinkData":    userCert.LinkData,
		},
		"httpsEnable": httpsEnabled,
	}

	urlPath := fmt.Sprintf("/v2/%s/certificates", domain)
	respObj := struct {
		CertId string `json:"certId"`
	}{}
	if err := httpRequest(cli, "PUT", urlPath, nil, reqObj, &respObj); err != nil {
		return "", err
	}

	return respObj.CertId, nil
}

// GetCert - query the certificate data for the specified domain.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/kjzuvz70t
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *CertificateDetail: certificate details
//     - error: nil if success otherwise the specific error
func GetCert(cli bce.Client, domain string) (certDetail *CertificateDetail, err error) {
	urlPath := fmt.Sprintf("/v2/%s/certificates", domain)
	respObj := &CertificateDetail{}
	if err := httpRequest(cli, "GET", urlPath, nil, nil, respObj); err != nil {
		return nil, err
	}
	return respObj, err
}

// DeleteCert - delete the certificate data for the specified domain.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ljzuylmee
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *CertificateDetail: certificate details
//     - error: nil if success otherwise the specific error
func DeleteCert(cli bce.Client, domain string) error {
	urlPath := fmt.Sprintf("/v2/%s/certificates", domain)
	var respObj interface{}
	if err := httpRequest(cli, "DELETE", urlPath, nil, nil, respObj); err != nil {
		return err
	}
	return nil
}

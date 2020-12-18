package iot

const (
	SIGNATURE   = "SIGNATURE"
	CERTIFICATE = "CERT"
)

type Device struct {
	Name            string `json:"name"`
	Description     string `json:"desc"`
	CoreId          string `json:"iotCoreId"`
	TemplateId      string `json:"templateId"`
	AuthType        string `json:"authType"`
	SecretKey       string `json:"secretKey,omitempty"`
	PrivateKey      string `json:"privateKey,omitempty"`
	ClientCert      string `json:"clientCert,omitempty"`
	CreateTimestamp int64  `json:"createTs,omitempty"`
}

package bcm

import (
	"strings"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	ProductName        = "bcm"
	DefaultBcmEndpoint = ProductName + "." + bce.DEFAULT_REGION + "." + bce.DEFAULT_DOMAIN
)

// Client of BCM service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the bcm service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endpoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}

	if len(endpoint) == 0 {
		endpoint = DefaultBcmEndpoint
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: map[string]struct{}{
			strings.ToLower(http.HOST):     {},
			strings.ToLower(http.BCE_DATE): {},
		},
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}
	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

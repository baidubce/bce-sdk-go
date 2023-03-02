package billing

// client.go - define the client for billing service

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/billing/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "billing.baidubce.com"
)

// Client of BCC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, region string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	endPoint := DEFAULT_SERVICE_DOMAIN

	if err != nil {
		return nil, err
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endPoint,
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

func (c *Client) GetBalance() (*api.BalanceResponse, error) {
	return api.GetBalance(c)
}

func (c *Client) GetBilling(queryArgs *api.BillingParams) (*api.BillingResponse, error) {
	return api.GetBilling(c, queryArgs)
}

package client

import (
	"strings"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_GLOBAL_ENDPOINT = "aihc.baidubce.com"
	DEFAULT_REGION_ENDPOINT = "aihc." + bce.DEFAULT_REGION + ".baidubce.com"
)

// Client of BBC service is a kind of BceClient, so derived from BceClient
type Client struct {
	DefaultClient *bce.BceClient
	RegionClient  *bce.BceClient
	GlobalClient  *bce.BceClient
}

// NewClient make the BBC service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endpoint string) (*Client, error) {
	regionClient, err := NewBceClient(ak, sk, "", DEFAULT_GLOBAL_ENDPOINT)
	if err != nil {
		return nil, err
	}

	gobalClient, err := NewBceClient(ak, sk, "", DEFAULT_REGION_ENDPOINT)
	if err != nil {
		return nil, err
	}

	client := &Client{
		RegionClient: regionClient,
		GlobalClient: gobalClient}
	if len(endpoint) == 0 {
		client.DefaultClient = regionClient
	} else {
		client.DefaultClient, err = NewBceClient(ak, sk, "", endpoint)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

// NewClientWithSTS make the aihc inference service client with STS configuration.
func NewClientWithSTS(ak, sk, sessionToken, endpoint string) (*Client, error) {
	regionClient, err := NewBceClient(ak, sk, sessionToken, DEFAULT_GLOBAL_ENDPOINT)
	if err != nil {
		return nil, err
	}

	gobalClient, err := NewBceClient(ak, sk, sessionToken, DEFAULT_REGION_ENDPOINT)
	if err != nil {
		return nil, err
	}

	client := &Client{
		RegionClient: regionClient,
		GlobalClient: gobalClient}
	if len(endpoint) == 0 {
		client.DefaultClient = regionClient
	} else {
		client.DefaultClient, err = NewBceClient(ak, sk, sessionToken, endpoint)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

func NewBceClient(ak, sk, sessionToken, endpoint string) (*bce.BceClient, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(sessionToken) > 0 {
		credentials, err = auth.NewSessionBceCredentials(ak, sk, sessionToken)
	} else {
		credentials, err = auth.NewBceCredentials(ak, sk)
	}
	if err != nil {
		return nil, err
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_REGION_ENDPOINT
	}

	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
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
	return bce.NewBceClient(defaultConf, v1Signer), nil
}

func IsRegionedEndpoint(endpoint string) bool {
	endpointSlice := strings.Split(endpoint, ".")
	return len(endpointSlice) <= 3
}

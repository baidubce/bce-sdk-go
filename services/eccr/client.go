package eccr

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_ENDPOINT = "ccr." + bce.DEFAULT_REGION + ".baidubce.com"

	URI_PREFIX = bce.URI_PREFIX + "api/ccr/esvc/v1"

	REQUEST_INSTANCE_URL = "/instances"

	REQUEST_PRIVATELINK_URL = "/privatelinks"
)

// Client ccr enterprise interface.Interface
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getInstanceListURI() string {
	return URI_PREFIX + REQUEST_INSTANCE_URL
}

func getInstanceURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID
}

func getPrivateNetworkListResponseURI(instanceID string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URL + "/" + instanceID + REQUEST_PRIVATELINK_URL
}

package bbc

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_ENDPOINT = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"
)

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

func getURL(version int) (ret string) {
	return fmt.Sprintf("%sv%d/instance", bce.URI_PREFIX, version)
}

func getURLwithID(version int, instanceID string) (ret string) {
	return fmt.Sprintf("%s%d/instance/%s", bce.URI_PREFIX, version, instanceID)
}

func getURLforVPC(version int) string {
	return fmt.Sprintf("%s%d/vpcSubnet", bce.URI_PREFIX, version)
}

package bill

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	// DefaultEndpoint -- xx
	DefaultEndpoint = "billing.baidubce.com"
)

// Client used for client
type Client struct {
	*bce.BceClient
}

// NewClient return a client
func NewClient(ak, sk, endPoint string) (ret *Client, err error) {
	if len(endPoint) == 0 {
		endPoint = DefaultEndpoint
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

package bbc

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	// DefaultEndpoint -- xx
	DefaultEndpoint = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"
)

// Client -- xx
type Client struct {
	*bce.BceClient
}

// NewClient -- xx
func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DefaultEndpoint
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
	return fmt.Sprintf("%sv%d/instance/%s", bce.URI_PREFIX, version, instanceID)
}

func getURLforVPC(version int) string {
	return fmt.Sprintf("%sv%d/vpcSubnet", bce.URI_PREFIX, version)
}

func getURLforTagwithID(version int, instanceID string) string {
	return fmt.Sprintf("%sv%d/instance/%s/tag", bce.URI_PREFIX, version, instanceID)
}

func getURLforFlavor(version int) string {
	return fmt.Sprintf("%sv%d/flavor", bce.URI_PREFIX, version)
}

func getURLforFlavorwithID(version int, instanceID string) string {
	return fmt.Sprintf("%sv%d/floavor/%s", bce.URI_PREFIX, version, instanceID)
}

func getURLforFlavorRaid(version int, instanceID string) string {
	return fmt.Sprintf("%sv%d/flavorRaid/%s", bce.URI_PREFIX, version, instanceID)
}

func getURLforImage(version int) string {
	return fmt.Sprintf("%sv%d/image", bce.URI_PREFIX, version)
}

func getURLforImagewithID(version int, imageID string) string {
	return fmt.Sprintf("%sv%d/image/%s", bce.URI_PREFIX, version, imageID)
}

func getURLforOperationLog(version int) string {
	return fmt.Sprintf("%sv%d/operationLog", bce.URI_PREFIX, version)
}

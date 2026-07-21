package iot

import (
	"fmt"
	"os"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	defaultServiceDomain = "iot." + bce.DEFAULT_DOMAIN
)

type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk string) (*Client, error) {
	client, err := bce.NewBceClientWithAkSk(ak, sk, defaultServiceDomain)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func (c *Client) CreateDevice(coreID, templateId, deviceName, authType, desc string) (*Device, error) {
	if coreID == "" || templateId == "" || deviceName == "" || authType == "" {
		return nil, os.ErrInvalid
	}
	result := &Device{
		CoreId:      coreID,
		TemplateId:  templateId,
		Name:        deviceName,
		AuthType:    authType,
		Description: desc,
	}
	err := bce.NewRequestBuilder(c).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithURL(getDeviceURL(coreID) + "new").
		WithMethod(http.POST).
		WithBody(result).
		WithResult(result).
		Do()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) DeleteDevice(coreID, deviceName string) error {
	if coreID == "" || deviceName == "" {
		return os.ErrInvalid
	}
	return bce.NewRequestBuilder(c).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithURL(getDeviceURL(coreID) + deviceName).
		WithMethod(http.DELETE).
		Do()
}

func getDeviceURL(id string) string {
	return fmt.Sprintf("/v1/iotcore/%s/device/", id)
}

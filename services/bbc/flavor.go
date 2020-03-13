package bbc

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListFlavor -- xx
func (c *Client) ListFlavor() (list *ListFlavorsResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithURL(getURLforFlavor(Version1)).
		WithMethod(http.GET).
		WithResult(&list).
		Do()
	return
}

// GetFlavorDetail -- xx
func (c *Client) GetFlavorDetail(instanceID string) (ret *FlavorModel, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getURLforFlavorwithID(Version1, instanceID)).
		WithResult(&ret).
		Do()
	return
}

// GetRaidofFlavor -- xx
func (c *Client) GetRaidofFlavor(instanceID string) (ret *GetRaidofFlavorResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getURLforFlavorRaid(Version1, instanceID)).
		WithResult(&ret).
		Do()
	return
}

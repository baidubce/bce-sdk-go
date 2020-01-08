package dcc

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

//ListDedicatedHosts -- xx
func (c *Client) ListDedicatedHosts(args *ListDedicatedHostArgs) (list *ListDedicatedHostResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithBody(args).
		WithURL(bce.URI_PREFIX+"v1/dedicatedHost").
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", fmt.Sprintf("%d", args.MaxKeys)).
		WithQueryParamFilter("zoneName", args.ZoneName).
		WithResult(&list).
		Do()
	return
}

//GetDedicatedHostDetail -- xx
func (c *Client) GetDedicatedHostDetail(hostID string) (ret *GetDedicatedHostDetailResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(bce.URI_PREFIX + "v1/dedicatedHost/" + hostID).
		WithBody(&ret).
		WithResult(&ret).
		Do()
	return

}

//PurchaseReserved -- xx
func (c *Client) PurchaseReserved(hostID string, args *PurchaseReservedArgs) (err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(bce.URI_PREFIX+"v1/dedicatedHost/"+hostID).
		WithQueryParamFilter("purchaseReserved", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(&args).
		Do()
	return
}

//Create -- xx
func (c *Client) Create(args *CreateArgs) (ret *CreateResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(bce.URI_PREFIX+"v1/dedicatedHost").
		WithQueryParamFilter("clientToken", args.clientToken).
		WithBody(&args).
		WithResult(&ret).
		Do()
	return

}
func (c *Client) tag(action, uri string, args interface{}) (err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(uri).
		WithQueryParamFilter(action, "").
		WithBody(&args).
		Do()
	return

}

// BindTag -- xx
func (c *Client) BindTag(dccID string, args *BindTagArgs) (err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(bce.URI_PREFIX+"v1/dedicatedHost/"+dccID+"/tag").
		WithQueryParamFilter("bind", "").
		WithBody(&args).
		Do()
	return

}

// UnbindTag -- xx
func (c *Client) UnbindTag(dccID string, args *BindTagArgs) (err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(bce.URI_PREFIX+"v1/dedicatedHost/"+dccID+"/tag").
		WithQueryParamFilter("unbind", "").
		WithBody(&args).
		Do()
	return
}

// CreateInstance -- xx
func (c *Client) CreateInstance(args *CreateInstanceArgs) (ret *CreateInstanceResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(bce.URI_PREFIX+"v1/dedicatedHost").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(&args).
		WithResult(&ret).
		Do()
	return
}

// ModityInstance -- xx
func (c *Client) ModityInstance(instanceID string, args *ModityInstanceArgs) (err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(bce.URI_PREFIX+"v1/dedicatedHost/instance/"+instanceID).
		WithQueryParamFilter("modifyAttribute", "").
		WithBody(&args).
		Do()
	return
}
func (c *Client) tagforInstance(action, instanceID string, args *BindTagArgs) (err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(bce.URI_PREFIX+"v1/dedicatedHost/instance/"+instanceID+"/tag").
		WithQueryParamFilter(action, "").
		WithBody(&args).
		Do()
	return
}

//BindTagforInstance -- xx
func (c *Client) BindTagforInstance(instanceID string, args *BindTagArgs) error {
	return c.tagforInstance("bind", instanceID, args)
}

// UnbindTagforInstance -- xx
func (c *Client) UnbindTagforInstance(instanceID string, args *BindTagArgs) error {
	return c.tagforInstance("unbind", instanceID, args)
}

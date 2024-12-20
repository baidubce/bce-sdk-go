package vpc

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

func (c *Client) CreateIpSet(args *CreateIpSetArgs) (*CreateIpSetResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The CreateIpSetArgs cannot be nil.")
	}
	result := &CreateIpSetResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpSet()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) AddIpAddress2IpSet(ipSetId string, args *AddIpAddress2IpSetArgs) error {
	if ipSetId == "" {
		return fmt.Errorf("The ipSetId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The AddIpAddress2IpSetArgs cannot be nil.")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpSetId(ipSetId)+"/ipAddress").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
	return err
}

func (c *Client) DeleteIpAddress(ipSetId string, args *DeleteIpAddressArgs) error {
	if ipSetId == "" {
		return fmt.Errorf("The ipSetId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The DeleteIpAddressArgs cannot be nil.")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpSetId(ipSetId)+"/deleteIpAddress").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
	return err
}

func (c *Client) UpdateIpSet(ipSetId string, args *UpdateIpSetArgs) error {
	if ipSetId == "" {
		return fmt.Errorf("The ipSetId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The UpdateIpSet cannot be nil.")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpSetId(ipSetId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("modifyAttribute", "").
		Do()
	return err
}

func (c *Client) DeleteIpSet(ipSetId string, args *DeleteIpSetArgs) error {
	if ipSetId == "" {
		return fmt.Errorf("The ipSetId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The DeleteIpSetArgs cannot be nil.")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpSetId(ipSetId)).
		WithMethod(http.DELETE).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
	return err
}

func (c *Client) ListIpSet(args *ListIpSetArgs) (*ListIpSetResult, error) {
	if args == nil {
		args = &ListIpSetArgs{}
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}
	result := &ListIpSetResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpSet()).
		WithMethod(http.GET).
		WithQueryParamFilter("ipVersion", args.IpVersion).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetIpSetDetail(ipSetId string) (*GetIpSetDetailResult, error) {
	if ipSetId == "" {
		return nil, fmt.Errorf("The ipSetId cannot be blank.")
	}
	result := &GetIpSetDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpSetId(ipSetId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) CreateIpGroup(args *CreateIpGroupArgs) (*CreateIpGroupResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The CreateIpGroupArgs cannot be nil.")
	}
	result := &CreateIpGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpGroup()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) AddIpSet2IpGroup(ipGroupId string, args *AddIpSet2IpGroupArgs) error {
	if ipGroupId == "" {
		return fmt.Errorf("The ipGroupId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The AddIpSet2IpGroupArgs cannot be nil.")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpGroupId(ipGroupId)+"/bindIpSet").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
	return err
}

func (c *Client) UnbindIpSet(ipGroupId string, args *UnbindIpSetArgs) error {
	if ipGroupId == "" {
		return fmt.Errorf("The ipGroupId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The UnbindIpSetArgs cannot be nil.")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpGroupId(ipGroupId)+"/unbindIpSet").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
	return err
}

func (c *Client) UpdateIpGroup(ipGroupId string, args *UpdateIpGroupArgs) error {
	if ipGroupId == "" {
		return fmt.Errorf("The ipGroupId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The UpdateIpGroupArgs cannot be nil.")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpGroupId(ipGroupId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("modifyAttribute", "").
		Do()
	return err
}

func (c *Client) DeleteIpGroup(ipGroupId string, args *DeleteIpGroupArgs) error {
	if ipGroupId == "" {
		return fmt.Errorf("The ipGroupId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The DeleteIpGroupArgs cannot be nil.")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpGroupId(ipGroupId)).
		WithMethod(http.DELETE).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
	return err
}

func (c *Client) ListIpGroup(args *ListIpGroupArgs) (*ListIpGroupResult, error) {
	if args == nil {
		args = &ListIpGroupArgs{}
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}
	result := &ListIpGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpGroup()).
		WithMethod(http.GET).
		WithQueryParamFilter("ipVersion", args.IpVersion).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetIpGroupDetail(ipGroupId string) (*GetIpGroupDetailResult, error) {
	if ipGroupId == "" {
		return nil, fmt.Errorf("The ipGroupId cannot be blank.")
	}
	result := &GetIpGroupDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpGroupId(ipGroupId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

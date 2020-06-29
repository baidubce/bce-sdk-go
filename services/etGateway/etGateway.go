package etGateway

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

// CreateEtGateway - create a new Et gateway
//
// PARAMS:
//    - args: the arguments to create Et gateway
// RETURNS:
//    - *CreateVpnGatewayResult: the id of the et gateway newly created
//    - error: nil if success otherwise the specific error

func (c *Client) CreateEtGateway(args *CreateEtGatewayArgs) (*CreateEtGatewayResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The CreateEtGatewayArgs cannot be nil.")
	}

	result := &CreateEtGatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEtGateway()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

//
// ListEtGateway - list all Et gateways with the specific parameters
// PARAMS:
//    - args: the arguments to list et gateways
// RETURNS:
//    - *ListEtGatewayResult: the result of Et gateway list
//    - error: nil if success otherwise the specific error
func (c *Client) ListEtGateway(args *ListEtGatewayArgs) (*ListEtGatewayResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The ListEtGatewayArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}
	result := &ListEtGatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEtGateway()).
		WithMethod(http.GET).
		WithQueryParam("vpcId", args.VpcId).
		WithQueryParamFilter("etGatewayId", args.EtGatewayId).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("status", args.Status).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()
	return result, err
}

// GetEtGatewayDetail - Get the Et gateways with the specific parameters
// PARAMS:
//    - etGatewayId: the id of  the EtGateway's
// RETURNS:
//    - *EtGatewayDetail: the result of EtGgateway detail
//    - error: nil if success otherwise the specific error
func (c *Client) GetEtGatewayDetail(etGatewayId string) (*EtGatewayDetail, error) {
	result := &EtGatewayDetail{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEtGatewayId(etGatewayId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// UpdateEtGateway - update the Et gateways with the specific parameters
// PARAMS:
//    - args: the arguments to update the EtGateway
// RETURNS:
//    - error: nil if success otherwise the specific error
func (c *Client) UpdateEtGateway(updateEtGatewayArgs *UpdateEtGatewayArgs) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForEtGatewayId(updateEtGatewayArgs.EtGatewayId)).
		WithQueryParamFilter("clientToken", updateEtGatewayArgs.ClientToken).
		WithMethod(http.PUT).
		WithBody(updateEtGatewayArgs).
		Do()
}

// DeleteEtGateway - delete the Et gateways with the specific parameters
// PARAMS:
//    - etGatewayId: the id to delete the EtGateway
//    - clientToken: the idempotent string
// RETURNS:
//    - error: nil if success otherwise the specific error
func (c *Client) DeleteEtGateway(etGatewayId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForEtGatewayId(etGatewayId)).
		WithQueryParam("clientToken", clientToken).
		WithMethod(http.DELETE).
		Do()
}

// UnBindEt -  bind the Et
// PARAMS:
//    - args: the arguments to bind the Et
// RETURNS:
//    - error: nil if success otherwise the specific error
func (c *Client) BindEt(args *BindEtArgs) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForEtGatewayId(args.EtGatewayId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("bind", "").
		WithBody(args).
		WithMethod(http.PUT).
		Do()
}

// UnBindEt -  unbind the Et
// PARAMS:
//    - args: the arguments to unbind the Et
// RETURNS:
//    - error: nil if success otherwise the specific error
func (c *Client) UnBindEt(EtGatewayId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForEtGatewayId(EtGatewayId)).
		WithQueryParamFilter("clientToken", clientToken).
		WithQueryParam("unbind", "").
		WithMethod(http.PUT).
		Do()
}

// CreateHealthCheck - create the Et gateway's healthcheck with the specific parameters
// PARAMS:
//    - args: the arguments to create the EtGateway's healthcheck
// RETURNS:
//    - error: nil if success otherwise the specific error
func (c *Client) CreateHealthCheck(args *CreateHealthCheckArgs) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForEtGatewayId(args.EtGatewayId)+"/healthCheck").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithMethod(http.POST).
		Do()
}

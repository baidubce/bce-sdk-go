/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// nat.go - the nat gateway APIs definition supported by the VPC service

package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateNatGateway - create a new nat gateway
//
// PARAMS:
//     - args: the arguments to create nat gateway
// RETURNS:
//     - *CreateNatGatewayResult: the id of the nat gateway newly created
//     - error: nil if success otherwise the specific error
func (c *Client) CreateNatGateway(args *CreateNatGatewayArgs) (*CreateNatGatewayResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createNatGatewayArgs cannot be nil.")
	}

	result := &CreateNatGatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNat()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListNatGateway - list all nat gateways with the specific parameters
//
// PARAMS:
//     - args: the arguments to list nat gateways
// RETURNS:
//     - *ListNatGatewayResult: the result of nat gateway list
//     - error: nil if success otherwise the specific error
func (c *Client) ListNatGateway(args *ListNatGatewayArgs) (*ListNatGatewayResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The listNatGatewayArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListNatGatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNat()).
		WithMethod(http.GET).
		WithQueryParam("vpcId", args.VpcId).
		WithQueryParamFilter("natId", args.NatId).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("ip", args.Ip).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetNatGatewayDetail - get details of the specific nat gateway
//
// PARAMS:
//     - natId: the id of the specified nat
// RETURNS:
//     - *NAT: the result of the specific nat gateway details
//     - error: nil if success otherwise the specific error
func (c *Client) GetNatGatewayDetail(natId string) (*NAT, error) {
	result := &NAT{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

// UpdateNatGateway - update the specified nat gateway
//
// PARAMS:
//     - natId: the id of the specific nat gateway
//     - args: the arguments to update nat gateway
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateNatGateway(natId string, args *UpdateNatGatewayArgs) error {
	if args == nil {
		return fmt.Errorf("The updateNatGatewayArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// BindEips - bind eips for the specific nat gateway
//
// PARAMS:
//     - natId: the id of the specific nat gateway
//     - args: the arguments to bind eips
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindEips(natId string, args *BindEipsArgs) error {
	if args == nil {
		return fmt.Errorf("The bindEipArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("bind", "").
		Do()
}

// UnBindEips - unbind eips for the specific nat gateway
//
// PARAMS:
//     - natId: the id of the specific nat gateway
//     - args: the arguments to unbind eips
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnBindEips(natId string, args *UnBindEipsArgs) error {
	if args == nil {
		return fmt.Errorf("The unBindEipArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("unbind", "").
		Do()
}

// BindDnatEips - bind dnatEips for the specific nat gateway
//
// PARAMS:
//     - natId: the id of the specific nat gateway
//     - args: the arguments to bind dnatEips
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindDnatEips(natId string, args *BindDnatEipsArgs) error {
	if args == nil {
		return fmt.Errorf("The bindDnatEipsArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("bind", "").
		Do()
}

// UnBindDnatEips - unbind dnatEips for the specific nat gateway
//
// PARAMS:
//     - natId: the id of the specific nat gateway
//     - args: the arguments to unbind dnatEips
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnBindDnatEips(natId string, args *UnBindDnatEipsArgs) error {
	if args == nil {
		return fmt.Errorf("the unBindDnatEipArgs cannot be nil")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("unbind", "").
		Do()
}

// EnhanceNatBindEips - bind eips for the specific enhanced nat gateway
//
// PARAMS:
//     - natId: the id of the specific enhanced nat gateway
//     - args: the arguments to bind eips
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) EnhanceNatBindEips(natId string, args *EnhanceNatBindEipsArgs) error {
	if args == nil {
		return fmt.Errorf("The bindDnatEipsArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("bind", "").
		Do()
}

// EnhanceNatUnBindEips - unbind eips for the specific enhanced nat gateway
//
// PARAMS:
//     - natId: the id of the specific enhanced nat gateway
//     - args: the arguments to unbind eips
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) EnhanceNatUnBindEips(natId string, args *EnhanceNatUnBindEipsArgs) error {
	if args == nil {
		return fmt.Errorf("the unBindDnatEipArgs cannot be nil")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("unbind", "").
		Do()
}

// DeleteNatGateway - delete the specific nat gateway
//
// PARAMS:
//     - natId: the id of the specific nat gateway
//     - clientToken: the idempotent token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteNatGateway(natId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// RenewNatGateway - renew nat gateway with the specific parameters
//
// PARAMS:
//     - natId: the id of the specific nat gateway
//     - args: the arguments to renew nat gateway
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RenewNatGateway(natId string, args *RenewNatGatewayArgs) error {
	if args == nil {
		return fmt.Errorf("The renewNatGatewayArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("purchaseReserved", "").
		Do()
}

func (c *Client) CreateNatGatewaySnatRule(natId string, args *CreateNatGatewaySnatRuleArgs) (*CreateNatGatewaySnatRuleResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The CreateNatGatewaySnatRuleArgs cannot be nil.")
	}

	result := &CreateNatGatewaySnatRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)+"/snatRule").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) BatchCreateNatGatewaySnatRule(args *BatchCreateNatGatewaySnatRuleArgs) (*BatchCreateNatGatewaySnatRuleResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The BatchCreateNatGatewaySnatRuleArgs cannot be nil.")
	}

	result := &BatchCreateNatGatewaySnatRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNat()+"/snatRule/batchCreate").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteNatGatewaySnatRule(natId string, snatRuleId string, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)+"/snatRule/"+snatRuleId).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

func (c *Client) UpdateNatGatewaySnatRule(natId string, snatRuleId string, args *UpdateNatGatewaySnatRuleArgs) error {
	if args == nil {
		return fmt.Errorf("The UpdateNatGatewaySnatRuleArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)+"/snatRule/"+snatRuleId).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

func (c *Client) ListNatGatewaySnatRules(args *ListNatGatewaySnatRuleArgs) (*ListNatGatewaySnatRulesResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The ListNatGatewaySnatRuleArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListNatGatewaySnatRulesResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(args.NatId)+"/snatRule").
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) CreateNatGatewayDnatRule(natId string, args *CreateNatGatewayDnatRuleArgs) (*CreateNatGatewayDnatRuleResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The CreateNatGatewayDnatRuleArgs cannot be nil.")
	}

	result := &CreateNatGatewayDnatRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)+"/dnatRule").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) BatchCreateNatGatewayDnatRule(natId string, args *BatchCreateNatGatewayDnatRuleArgs) (*BatchCreateNatGatewayDnatRuleResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The BatchCreateNatGatewayDnatRuleArgs cannot be nil.")
	}

	result := &BatchCreateNatGatewayDnatRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)+"/dnatRule/batchCreate").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) DeleteNatGatewayDnatRule(natId string, dnatRuleId string, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)+"/dnatRule/"+dnatRuleId).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

func (c *Client) UpdateNatGatewayDnatRule(natId string, dnatRuleId string, args *UpdateNatGatewayDnatRuleArgs) error {
	if args == nil {
		return fmt.Errorf("The UpdateNatGatewayDnatRuleArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)+"/dnatRule/"+dnatRuleId).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

func (c *Client) ListNatGatewayDnatRules(natId string, args *ListNatGatewaDnatRuleArgs) (*ListNatGatewayDnatRulesResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The ListNatGatewaDnatRuleArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListNatGatewayDnatRulesResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)+"/dnatRule").
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ResizeNatGateway(natId string, args *ResizeNatGatewayArgs) error {
	if args == nil {
		return fmt.Errorf("The ResizeNatGatewayArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForNatId(natId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("resize", " ").
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

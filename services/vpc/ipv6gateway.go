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

// ipv6gateway.go - the ipv6 gateway APIs definition supported by the VPC service

package vpc

import (
	"errors"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateIPv6Gateway - create a new ipv6 gateway
//
// PARAMS:
//   - args: the arguments to create ipv6 gateway
//
// RETURNS:
//   - *CreateIPv6GatewayResult: the id of the ipv6 gateway newly created
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIPv6Gateway(args *CreateIPv6GatewayArgs) (*CreateIPv6GatewayResult, error) {
	if args == nil {
		return nil, errors.New("The CreateIPv6GatewayArgs cannot be nil.")
	}

	result := &CreateIPv6GatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6Gateway()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListIPv6Gateway - list all ipv6 gateways with the specific parameters
//
// PARAMS:
//   - args: the arguments to list ipv6 gateways
//
// RETURNS:
//   - *ListIPv6GatewayResult: the result of ipv6 gateway list
//   - error: nil if success otherwise the specific error
func (c *Client) ListIPv6Gateway(args *ListIPv6GatewayArgs) (*ListIPv6GatewayResult, error) {
	if args == nil {
		return nil, errors.New("The ListIPv6GatewayArgs cannot be nil.")
	}

	result := &ListIPv6GatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6Gateway()).
		WithMethod(http.GET).
		WithQueryParam("vpcId", args.VpcId).
		WithResult(result).
		Do()

	return result, err
}

// DeleteIPv6Gateway - delete the specified ipv6 gateway
//
// PARAMS:
//   - gatewayId: the id of the specific ipv6 gateway
//   - args: the arguments to delete ipv6 gateway
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteIPv6Gateway(gatewayId string, args *DeleteIPv6GatewayArgs) error {
	if args == nil {
		return errors.New("The DeleteIPv6GatewayArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6GatewayId(gatewayId)).
		WithMethod(http.DELETE).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// ResizeIPv6Gateway - resize the specified ipv6 gateway
//
// PARAMS:
//   - gatewayId: the id of the specific ipv6 gateway
//   - args: the arguments to resize ipv6 gateway
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ResizeIPv6Gateway(gatewayId string, args *ResizeIPv6GatewayArgs) error {
	if args == nil {
		return errors.New("The ResizeIPv6GatewayArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6GatewayId(gatewayId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// CreateIPv6GatewayEgressOnlyRule - create a new ipv6 gateway egress only rule
//
// PARAMS:
//   - gatewayId: the id of the specific ipv6 gateway
//   - args: the arguments to create ipv6 gateway egress only rule
//
// RETURNS:
//   - *CreateIPv6GatewayEgressOnlyRuleResult: the id of the ipv6 gateway egress only rule newly created
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIPv6GatewayEgressOnlyRule(gatewayId string, args *CreateIPv6GatewayEgressOnlyRuleArgs) (
	*CreateIPv6GatewayEgressOnlyRuleResult, error) {
	if args == nil {
		return nil, errors.New("The CreateIPv6GatewayEgressOnlyRuleArgs cannot be nil.")
	}

	result := &CreateIPv6GatewayEgressOnlyRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6GatewayId(gatewayId)+"/egressOnlyRule").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListIPv6GatewayEgressOnlyRule - list all ipv6 gateway egress only rules with the specific parameters
//
// PARAMS:
//   - gatewayId: the id of the specific ipv6 gateway
//   - args: the arguments to list ipv6 gateway egress only rules
//
// RETURNS:
//   - *ListIPv6GatewayEgressOnlyRuleResult: the result of ipv6 gateway egress only rule list
//   - error: nil if success otherwise the specific error
func (c *Client) ListIPv6GatewayEgressOnlyRule(gatewayId string, args *ListIPv6GatewayEgressOnlyRuleArgs) (*ListIPv6GatewayEgressOnlyRuleResult, error) {
	if args == nil {
		return nil, errors.New("The ListIPv6GatewayEgressOnlyRuleArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListIPv6GatewayEgressOnlyRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6GatewayId(gatewayId)+"/egressOnlyRule").
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// DeleteIPv6GatewayEgressOnlyRule - delete the specified ipv6 gateway egress only rule
//
// PARAMS:
//   - gatewayId: the id of the specific ipv6 gateway
//   - egressOnlyRuleId: the id of the specific ipv6 gateway egress only rule
//   - args: the arguments to delete ipv6 gateway egress only rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteIPv6GatewayEgressOnlyRule(gatewayId, egressOnlyRuleId string, args *DeleteIPv6GatewayEgressOnlyRuleArgs) error {
	if args == nil {
		return errors.New("The DeleteIPv6GatewayEgressOnlyRuleArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6GatewayId(gatewayId)+"/egressOnlyRule/"+egressOnlyRuleId).
		WithMethod(http.DELETE).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// CreateIPv6GatewayRateLimitRule - create a new ipv6 gateway rate limit rule
//
// PARAMS:
//   - gatewayId: the id of the specific ipv6 gateway
//   - args: the arguments to create ipv6 gateway rate limit rule
//
// RETURNS:
//   - *CreateIPv6GatewayRateLimitRuleResult: the id of the ipv6 gateway rate limit rule newly created
//   - error: nil if success otherwise the specific error
func (c *Client) CreateIPv6GatewayRateLimitRule(gatewayId string, args *CreateIPv6GatewayRateLimitRuleArgs) (
	*CreateIPv6GatewayRateLimitRuleResult, error) {
	if args == nil {
		return nil, errors.New("The CreateIPv6GatewayRateLimitRuleArgs cannot be nil.")
	}

	result := &CreateIPv6GatewayRateLimitRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6GatewayId(gatewayId)+"/rateLimitRule").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListIPv6GatewayRateLimitRule - list all ipv6 gateway rate limit rules with the specific parameters
//
// PARAMS:
//   - gatewayId: the id of the specific ipv6 gateway
//   - args: the arguments to list ipv6 gateway rate limit rules
//
// RETURNS:
//   - *ListIPv6GatewayRateLimitRuleResult: the result of ipv6 gateway rate limit rule list
//   - error: nil if success otherwise the specific error
func (c *Client) ListIPv6GatewayRateLimitRule(gatewayId string, args *ListIPv6GatewayRateLimitRuleArgs) (*ListIPv6GatewayRateLimitRuleResult, error) {
	if args == nil {
		return nil, errors.New("The ListIPv6GatewayRateLimitRuleArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListIPv6GatewayRateLimitRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6GatewayId(gatewayId)+"/rateLimitRule").
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// DeleteIPv6GatewayRateLimitRule - delete the specified ipv6 gateway rate limit rule
//
// PARAMS:
//   - gatewayId: the id of the specific ipv6 gateway
//   - rateLimitRuleId: the id of the specific ipv6 gateway rate limit rule
//   - args: the arguments to delete ipv6 gateway rate limit rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteIPv6GatewayRateLimitRule(gatewayId, rateLimitRuleId string, args *DeleteIPv6GatewayRateLimitRuleArgs) error {
	if args == nil {
		return errors.New("The DeleteIPv6GatewayRateLimitRuleArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6GatewayId(gatewayId)+"/rateLimitRule/"+rateLimitRuleId).
		WithMethod(http.DELETE).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// UpdateIPv6GatewayRateLimitRule - update the specified ipv6 gateway rate limit rule
//
// PARAMS:
//   - gatewayId: the id of the specific ipv6 gateway
//   - rateLimitRuleId: the id of the specific ipv6 gateway rate limit rule
//   - args: the arguments to update ipv6 gateway rate limit rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateIPv6GatewayRateLimitRule(gatewayId, rateLimitRuleId string, args *UpdateIPv6GatewayRateLimitRuleArgs) error {
	if args == nil {
		return errors.New("The UpdateIPv6GatewayRateLimitRuleArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForIpv6GatewayId(gatewayId)+"/rateLimitRule/"+rateLimitRuleId).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

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

// route.go - the route APIs definition supported by the VPC service

package vpc

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// GetRouteTableDetail - get details of the given routeTableId or vpcId
//
// PARAMS:
//     - routeTableId: the id of the specific route table
//     - vpcId: the id of the specific VPC
// RETURNS:
//     - *GetRouteTableResult: the result of route table details
//     - error: nil if success otherwise the specific error
func (c *Client) GetRouteTableDetail(routeTableId, vpcId string) (*GetRouteTableResult, error) {
	if routeTableId == "" && vpcId == "" {
		return nil, fmt.Errorf("The routeTableId and vpcId cannot be blank at the same time.")
	}

	result := &GetRouteTableResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForRouteTable()).
		WithMethod(http.GET).
		WithQueryParamFilter("routeTableId", routeTableId).
		WithQueryParamFilter("vpcId", vpcId).
		WithResult(result).
		Do()

	return result, err
}

// CreateRouteRule - create a new route rule with the given parameters
//
// PARAMS:
//     - args: the arguments to create route rule
// RETURNS:
//     - *CreateRouteRuleResult: the id of the route rule newly created
//     - error: nil if success otherwise the specific error
func (c *Client) CreateRouteRule(args *CreateRouteRuleArgs) (*CreateRouteRuleResult, error) {
	if args == nil {
		return nil, fmt.Errorf("CreateRouteRuleArgs cannot be nil.")
	}

	result := &CreateRouteRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForRouteRule()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// DeleteRouteRule - delete the given routing rule
//
// PARAMS:
//     - routeRuleId: the id of the specific routing rule
//     - clientToken: the idempotent token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteRouteRule(routeRuleId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForRouteRuleId(routeRuleId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

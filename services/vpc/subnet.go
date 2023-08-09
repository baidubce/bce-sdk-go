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

// subnet.go - the subnet APIs definition supported by the VPC service

package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateSubnet - create a new subnet with the specified parameters
//
// PARAMS:
//   - args: the arguments to create subnet
//
// RETURNS:
//   - *CreateSubnetResult: the ID of the subnet newly created
//   - error: nil if success otherwise the specific error
func (c *Client) CreateSubnet(args *CreateSubnetArgs) (*CreateSubnetResult, error) {
	if args == nil {
		return nil, fmt.Errorf("CreateSubnetArgs cannot be nil.")
	}

	result := &CreateSubnetResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForSubnet()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListSubnets - list all subnets with the specified parameters
//
// PARAMS:
//   - args: the arguments to list subnets
//   - :
//
// RETURNS:
//   - *ListSubnetResult: the result of all subnets
//   - error: nil if success otherwise the specific error
func (c *Client) ListSubnets(args *ListSubnetArgs) (*ListSubnetResult, error) {
	if args == nil {
		args = &ListSubnetArgs{}
	}
	if args.MaxKeys < 0 || args.MaxKeys > 1000 {
		return nil, fmt.Errorf("The field maxKeys is out of range [0, 1000]")
	} else if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListSubnetResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForSubnet()).
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("vpcId", args.VpcId).
		WithQueryParamFilter("zoneName", args.ZoneName).
		WithQueryParamFilter("subnetType", string(args.SubnetType)).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetSubnetDetail - get details of the given subnet
//
// PARAMS:
//   - subnetId: the id of the specified subnet
//
// RETURNS:
//   - *GetSubnetDetailResult: the result of the given subnet details
//   - error: nil if success otherwise the specific error
func (c *Client) GetSubnetDetail(subnetId string) (*GetSubnetDetailResult, error) {
	if subnetId == "" {
		return nil, fmt.Errorf("The subnetId cannot be blank.")
	}

	result := &GetSubnetDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForSubnetId(subnetId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

// UpdateSubnet - update the given subnet
//
// PARAMS:
//   - subnetId: the id of the given subnet
//   - args: the arguments to update subnet
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateSubnet(subnetId string, args *UpdateSubnetArgs) error {
	if subnetId == "" {
		return fmt.Errorf("The subnetId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("The UpdateSubnetArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForSubnetId(subnetId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("modifyAttribute", "").
		Do()
}

// DeleteSubnet - delete the given subnet
//
// PARAMS:
//   - subnetId: the id of the specified subnet
//   - clientToken: the idempotent token
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteSubnet(subnetId string, clientToken string) error {
	if subnetId == "" {
		return fmt.Errorf("The subnetId cannot be blank.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForSubnetId(subnetId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

/*
 * Copyright 2021 Baidu, Inc.
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

// appipgroup.go - the Application BLB Ip Group APIs definition supported by the APPBLB service

package appblb

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateAppIpGroup - create an ip group
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create IpGroup
// RETURNS:
//     - *CreateAppIpGroupResult: the result of create IpGroup, contains new IpGroup's ID
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppIpGroup(blbId string, args *CreateAppIpGroupArgs) (*CreateAppIpGroupResult, error) {
	if args == nil {
		args = &CreateAppIpGroupArgs{}
	}

	result := &CreateAppIpGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppIpGroupUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpdateAppIpGroup - update an ip group
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update an ip group
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppIpGroup(blbId string, args *UpdateAppIpGroupArgs) error {
	if args == nil || len(args.IpGroupId) == 0 {
		return fmt.Errorf("unset ip group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppIpGroupUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribeAppIpGroup - describe all ip groups
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe all ip groups
// RETURNS:
//     - *DescribeAppIpGroupResult: the result of describe all ip groups
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeAppIpGroup(blbId string, args *DescribeAppIpGroupArgs) (*DescribeAppIpGroupResult, error) {
	if args == nil {
		args = &DescribeAppIpGroupArgs{}
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppIpGroupResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppIpGroupUri(blbId)).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ExactlyMatch {
		request.WithQueryParam("exactlyMatch", "true")
	}

	err := request.Do()
	return result, err
}

// DeleteAppIpGroup - delete an ip group
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to delete an ip group
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteAppIpGroup(blbId string, args *DeleteAppIpGroupArgs) error {
	if args == nil || len(args.IpGroupId) == 0 {
		return fmt.Errorf("unset ip group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppIpGroupUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("delete", "").
		WithBody(args).
		Do()
}

// CreateAppIpGroupBackendPolicy - create an ip group backend policy
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create an ip group backend policy
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppIpGroupBackendPolicy(blbId string, args *CreateAppIpGroupBackendPolicyArgs) error {
	if args == nil || len(args.IpGroupId) == 0 {
		return fmt.Errorf("unset ip group id")
	}

	if len(args.Type) == 0 {
		return fmt.Errorf("unset type")
	}

	if args.Type == "UDP" && len(args.UdpHealthCheckString) == 0 {
		return fmt.Errorf("unset udpHealthCheckString")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppIpGroupBackendPolicyUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateAppIpGroupBackendPolicy - update ip group backend policy
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update ip group backend policy
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppIpGroupBackendPolicy (blbId string, args *UpdateAppIpGroupBackendPolicyArgs) error {
	if args == nil || len(args.IpGroupId) == 0 {
		return fmt.Errorf("unset ip group id")
	}

	if len(args.Id) == 0 {
		return fmt.Errorf("unset ip group backend policy id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppIpGroupBackendPolicyUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DeleteAppIpGroupBackendPolicy - delete an ip group backend policy
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to delete ip group backend policies
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteAppIpGroupBackendPolicy(blbId string, args *DeleteAppIpGroupBackendPolicyArgs) error {
	if args == nil || len(args.IpGroupId) == 0 {
		return fmt.Errorf("unset ip group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppIpGroupBackendPolicyUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("delete", "").
		WithBody(args).
		Do()
}

// CreateAppIpGroupMember - create ip group members
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create ip group members
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppIpGroupMember(blbId string, args *CreateAppIpGroupMemberArgs) error {
	if args == nil || len(args.IpGroupId) == 0 {
		return fmt.Errorf("unset ip group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppIpGroupMemberUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateAppIpGroupMember - update ip group members
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update ip group members
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppIpGroupMember(blbId string, args *UpdateAppIpGroupMemberArgs) error {
	if args == nil || len(args.IpGroupId) == 0 {
		return fmt.Errorf("unset ip group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppIpGroupMemberUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribeAppIpGroupMember - describe ip group members
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe ip group members
// RETURNS:
//     - *DescribeAppIpGroupMemberResult: the result of describe ip group members
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeAppIpGroupMember(blbId string, args *DescribeAppIpGroupMemberArgs) (*DescribeAppIpGroupMemberResult, error) {
	if args == nil || len(args.IpGroupId) == 0 {
		return nil, fmt.Errorf("unset ip group id")
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppIpGroupMemberResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppIpGroupMemberUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParam("ipGroupId", args.IpGroupId).
		WithResult(result).
		Do()

	return result, err
}

// DeleteAppIpGroupMember - delete ip group members
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to delete ip group members
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteAppIpGroupMember(blbId string, args *DeleteAppIpGroupMemberArgs) error {
	if args == nil || len(args.IpGroupId) == 0 {
		return fmt.Errorf("unset ip group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppIpGroupMemberUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("delete", "").
		WithBody(args).
		Do()
}



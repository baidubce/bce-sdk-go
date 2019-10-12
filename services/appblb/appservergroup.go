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

// appservergroup.go - the Application BLB Server Group APIs definition supported by the APPBLB service

package appblb

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateAppServerGroup - create a LoadBalancer
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create ServerGroup
// RETURNS:
//     - *CreateAppServerGroupResult: the result of create ServerGroup, contains new ServerGroup's ID
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppServerGroup(blbId string, args *CreateAppServerGroupArgs) (*CreateAppServerGroupResult, error) {
	if args == nil {
		args = &CreateAppServerGroupArgs{}
	}

	result := &CreateAppServerGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppServerGroupUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpdateAppServerGroup - update a server group
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update a server group
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppServerGroup(blbId string, args *UpdateAppServerGroupArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppServerGroupUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribeAppServerGroup - describe all server groups
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe all server groups
// RETURNS:
//     - *DescribeAppServerGroupResult: the result of describe all server groups
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeAppServerGroup(blbId string, args *DescribeAppServerGroupArgs) (*DescribeAppServerGroupResult, error) {
	if args == nil {
		args = &DescribeAppServerGroupArgs{}
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppServerGroupResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppServerGroupUri(blbId)).
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

// DeleteAppServerGroup - delete a server group
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to delete a server group
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteAppServerGroup(blbId string, args *DeleteAppServerGroupArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppServerGroupUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("delete", "").
		WithBody(args).
		Do()
}

// CreateAppServerGroupPort - create a server group port
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create a server group port
// RETURNS:
//     - *CreateAppServerGroupPortResult: the result of create a server group port
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppServerGroupPort(blbId string, args *CreateAppServerGroupPortArgs) (*CreateAppServerGroupPortResult, error) {
	if args == nil || len(args.SgId) == 0 {
		return nil, fmt.Errorf("unset server group id")
	}

	if len(args.Type) == 0 {
		return nil, fmt.Errorf("unset type")
	}

	if args.Type == "UDP" && len(args.UdpHealthCheckString) == 0 {
		return nil, fmt.Errorf("unset udpHealthCheckString")
	}

	result := &CreateAppServerGroupPortResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppServerGroupPortUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpdateAppServerGroupPort - update server group port
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update server group port
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppServerGroupPort(blbId string, args *UpdateAppServerGroupPortArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppServerGroupPortUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DeleteAppServerGroupPort - delete server group ports
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to delete server group ports
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteAppServerGroupPort(blbId string, args *DeleteAppServerGroupPortArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppServerGroupPortUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("batchdelete", "").
		WithBody(args).
		Do()
}

// CreateBlbRs - add backend servers
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to add backend servers
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateBlbRs(blbId string, args *CreateBlbRsArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBlbRsUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateBlbRs - update backend servers
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update backend servers
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateBlbRs(blbId string, args *UpdateBlbRsArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBlbRsUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribeBlbRs - describe backend servers
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe backend servers
// RETURNS:
//     - *DescribeBlbRsResult: the result of describe backend servers
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeBlbRs(blbId string, args *DescribeBlbRsArgs) (*DescribeBlbRsResult, error) {
	if args == nil || len(args.SgId) == 0 {
		return nil, fmt.Errorf("unset server group id")
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeBlbRsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbRsUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParam("sgId", args.SgId).
		WithResult(result).
		Do()

	return result, err
}

// DeleteBlbRs - delete backend servers
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to delete backend servers
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteBlbRs(blbId string, args *DeleteBlbRsArgs) error {
	if args == nil || len(args.SgId) == 0 {
		return fmt.Errorf("unset server group id")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBlbRsUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("batchdelete", "").
		WithBody(args).
		Do()
}

// DescribeRsMount - get all mount backend server list
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - sgId: ServerGroup's ID
// RETURNS:
//     - *DescribeRsMountResult: the mount backend server list
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeRsMount(blbId, sgId string) (*DescribeRsMountResult, error) {
	if len(sgId) == 0 {
		return nil, fmt.Errorf("unset server group id")
	}

	result := &DescribeRsMountResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbRsMountUri(blbId)).
		WithQueryParam("sgId", sgId).
		WithResult(result).
		Do()

	return result, err
}

// DescribeRsUnMount - get all unmount backend server list
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - sgId: ServerGroup's ID
// RETURNS:
//     - *DescribeRsMountResult: the unMount backend server list
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeRsUnMount(blbId, sgId string) (*DescribeRsMountResult, error) {
	if len(sgId) == 0 {
		return nil, fmt.Errorf("unset server group id")
	}

	result := &DescribeRsMountResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbRsUnMountUri(blbId)).
		WithQueryParam("sgId", sgId).
		WithResult(result).
		Do()

	return result, err
}

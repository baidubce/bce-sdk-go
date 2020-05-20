/*
 * Copyright 2020 Baidu, Inc.
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

// backendserver.go - the backendserver APIs definition supported by the BLB service


package blb

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)


// AddBackendServers - add backend servers
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to add backend servers
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) AddBackendServers(blbId string, args *AddBackendServersArgs) error {

	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.BackendServerList) == 0 {
		return fmt.Errorf("unset backendServer list")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBackendServerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateBackendServers - update backend servers
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update backend servers
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateBackendServers(blbId string, args *UpdateBackendServersArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.BackendServerList) == 0 {
		return fmt.Errorf("unset backendServer list")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBackendServerUri(blbId)).
		WithQueryParam("update", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}


// DescribeBackendServers - describe all backend servers
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe all backend servers
// RETURNS:
//     - *DescribeBackendServersResult: the result of describe all backend servers
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeBackendServers(blbId string, args *DescribeBackendServersArgs) (*DescribeBackendServersResult, error) {
	if args == nil {
		args = &DescribeBackendServersArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeBackendServersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBackendServerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	err := request.Do()
	return result, err
}

// DescribeHealthStatus - describe all backend servers health status
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe all backend servers health status
// RETURNS:
//     - *DescribeHealthStatusResult: the result of describe all backend servers health status
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeHealthStatus(blbId string, args *DescribeHealthStatusArgs) (*DescribeHealthStatusResult, error) {
	if args == nil {
		args = &DescribeHealthStatusArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeHealthStatusResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBackendServerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	err := request.Do()
	return result, err
}

// RemoveBackendServers - remove backend servers
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to remove backend servers, a backend server list
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) RemoveBackendServers(blbId string, args *RemoveBackendServersArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.BackendServerList) == 0 {
		return fmt.Errorf("unset backend server list")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBackendServerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}








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

// endpoint.go - the endpoint APIs definition supported by the endpoint service
package endpoint

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// GetServices - get the public services
// RETURNS:
//   - *ListServiceResult: the result of public service
//   - error: nil if success otherwise the specific error
func (c *Client) GetServices() (*ListServiceResult, error) {
	result := &ListServiceResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEndpoint() + "/publicService").
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

// CreateEndpoint - create an endpoint with the specific parameters
//
// PARAMS:
//   - args: the arguments to create an endpoint
//
// RETURNS:
//   - *CreateEndpointResult: the result of create endpoint
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEndpoint(args *CreateEndpointArgs) (*CreateEndpointResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createEndpointArgs cannot be nil.")
	}

	result := &CreateEndpointResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEndpoint()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// DeleteEndpoint - delete an endpoint
//
// PARAMS:
//   - endpointId: the specific endpointId
//   - clientToken: optional parameter, an Idempotent Token
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteEndpoint(endpointId string, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForEndpointId(endpointId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// UpdateEndpoint - update an endpoint
//
// PARAMS:
//   - endpointId: the specific endpointId
//   - UpdateEndpointArgs: the arguments to update an endpoint
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateEndpoint(endpointId string, args *UpdateEndpointArgs) error {
	if args == nil {
		return fmt.Errorf("The updateEndpointArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForEndpointId(endpointId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// ListEndpoints - list all endpoint with the specific parameters
//
// PARAMS:
//   - args: the arguments to list all endpoint
//
// RETURNS:
//   - *ListEndpointResult: the result of list all endpoint
//   - error: nil if success otherwise the specific error
func (c *Client) ListEndpoints(args *ListEndpointArgs) (*ListEndpointResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The ListEndpointArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListEndpointResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEndpoint()).
		WithMethod(http.GET).
		WithQueryParam("vpcId", args.VpcId).
		WithQueryParam("name", args.Name).
		WithQueryParam("ipAddress", args.IpAddress).
		WithQueryParam("status", args.Status).
		WithQueryParam("subnetId", args.SubnetId).
		WithQueryParam("service", args.Service).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetEndpointDetail - get the endpoint detail
//
// PARAMS:
//   - endpointId: the specific endpointId
//
// RETURNS:
//   - *Endpoint: the endpoint
//   - error: nil if success otherwise the specific error
func (c *Client) GetEndpointDetail(endpointId string) (*Endpoint, error) {
	result := &Endpoint{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEndpointId(endpointId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

// UpdateEndpointNormalSecurityGroup - update normal security group bound to the endpoint
//
// PARAMS:
//   - endpointId: the specific endpointId
//   - args: the arguments to update a normal security group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateEndpointNormalSecurityGroup(endpointId string, args *UpdateEndpointNSGArgs) error {
	if args == nil {
		return fmt.Errorf("The UpdateEndpointNSGArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForEndpointId(endpointId)).
		WithMethod(http.PUT).
		WithQueryParam("bindSg", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateEndpointEnterpriseSecurityGroup - update enterprise security group bound to the endpoint
//
// PARAMS:
//   - endpointId: the specific endpointId
//   - args: the arguments to update an enterprise security group
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateEndpointEnterpriseSecurityGroup(endpointId string, args *UpdateEndpointESGArgs) error {
	if args == nil {
		return fmt.Errorf("The UpdateEndpointESGArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForEndpointId(endpointId)).
		WithMethod(http.PUT).
		WithQueryParam("bindEsg", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

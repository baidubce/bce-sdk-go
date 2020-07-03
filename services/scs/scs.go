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

// scs.go - the SCS for Redis APIs definition supported by the redis service
package scs

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	INSTANCE_URL_V1 = bce.URI_PREFIX + "v1" + "/instance"
	INSTANCE_URL_V2 = bce.URI_PREFIX + "v2" + "/instance"
)

// CreateInstance - create an instance with specified parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to create instance
// RETURNS:
//     - *CreateInstanceResult: result of the instance ids newly created
//     - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(args *CreateInstanceArgs) (*CreateInstanceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set create scs argments")
	}

	result := &CreateInstanceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(INSTANCE_URL_V2).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListInstances - list all instances with the specified parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list instances
// RETURNS:
//     - *ListInstanceResult: result of the instance list
//     - error: nil if success otherwise the specific error
func (c *Client) ListInstances(args *ListInstancesArgs) (*ListInstancesResult, error) {
	if args == nil {
		args = &ListInstancesArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListInstancesResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(INSTANCE_URL_V2).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetInstanceDetail - get details of the specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
// RETURNS:
//     - *GetInstanceDetailResult: result of the instance details
//     - error: nil if success otherwise the specific error
func (c *Client) GetInstanceDetail(instanceId string) (*GetInstanceDetailResult, error) {
	result := &GetInstanceDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(INSTANCE_URL_V2 +  "/" + instanceId).
		WithResult(result).
		Do()

	return result, err
}


// ResizeInstance - resize a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be resized
//     - reqBody: the request body to resize instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ResizeInstance(instanceId string, args *ResizeInstanceArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1 +  "/" + instanceId  + "/resize").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DeleteInstance - delete a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be deleted
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(instanceId string, clientToken string) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(INSTANCE_URL_V1 +  "/" + instanceId).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}


// UpdateInstanceName - update name of a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be deleted
//     - args: the arguments to Update instanceName
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceName(instanceId string, args *UpdateInstanceNameArgs) error {

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(INSTANCE_URL_V1 +  "/" + instanceId + "/rename").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// GetNodeTypeList - list all nodetype
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be deleted
//     - args: the arguments to Update instanceName
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) GetNodeTypeList() (*GetNodeTypeListResult, error){
	getNodeTypeListResult := &GetNodeTypeListResult{}
	err :=  bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL("/v2/nodetypes").
		WithResult(getNodeTypeListResult).
		Do()

	return getNodeTypeListResult, err
}

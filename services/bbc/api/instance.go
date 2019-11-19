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

// instance.go - the BBC APIs definition supported by the BBC service

package api

import (
	//"encoding/json"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateInstance - create an instance with specified parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to create instance
// RETURNS:
//     - *CreateInstanceResult: result of the instance ids newly created
//     - error: nil if success otherwise the specific error
func CreateInstance(cli bce.Client, clientToken string, reqBody *bce.Body) (*CreateInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	if clientToken != "" {
		req.SetParam("clientToken", clientToken)
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ListInstances - list all instances with the specified parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list instances
// RETURNS:
//     - *ListInstanceResult: result of the instance list
//     - error: nil if success otherwise the specific error
func ListInstances(cli bce.Client, args *ListInstanceArgs) (*ListInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUri())
	req.SetMethod(http.GET)

	// Optional arguments settings
	if args != nil {
		if len(args.Marker) != 0 {
			req.SetParam("marker", args.Marker)
		}
		if args.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(args.MaxKeys))
		}
		if len(args.InternalIp) != 0 {
			req.SetParam("internalIp", args.InternalIp)
		}
	}
	if args == nil || args.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetInstanceDetail - get details of the specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
// RETURNS:
//     - *GetInstanceDetailResult: result of the instance details
//     - error: nil if success otherwise the specific error
func GetInstanceDetail(cli bce.Client, instanceId string) (*GetInstanceDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetInstanceDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

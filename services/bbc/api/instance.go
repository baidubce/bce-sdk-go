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

// instance.go - the instance APIs definition supported by the BBC service

// Package api defines all APIs supported by the BBC service of BCE.
package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateInstance - create a bbc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - clientToken: idempotent token, an ASCII string no longer than 64 bits
//     - reqBody: http request body
// RETURNS:
//     - *CreateInstanceResult: results of creating a bbc instance
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

// ListInstances - list all bbc instances
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list bbc instances
// RETURNS:
//     - *ListInstanceResult: results of list bbc instances
//     - error: nil if success otherwise the specific error
func ListInstances(cli bce.Client, args *ListInstancesArgs) (*ListInstancesResult, error) {
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

	jsonBody := &ListInstancesResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetInstanceDetail - get a bbc instance detail msg
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
// RETURNS:
//     - *InstanceModel: instance detail msg
//     - error: nil if success otherwise the specific error
func GetInstanceDetail(cli bce.Client, instanceId string) (*InstanceModel, error) {
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

	jsonBody := &InstanceModel{}
	print(jsonBody)
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// StartInstance - start a bbc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func StartInstance(cli bce.Client, instanceId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("start", "")

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// StopInstance - stop a bbc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func StopInstance(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("stop", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// RebootInstance - reboot a bbc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func RebootInstance(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("reboot", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// ModifyInstanceName - modify a bbc instance name
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyInstanceName(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("rename", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// ModifyInstanceDesc - modify a bbc instance desc
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyInstanceDesc(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("updateDesc", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// RebuildInstance - rebuild a bbc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func RebuildInstance(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("rebuild", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// DeleteInstance - delete a bbc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteInstance(cli bce.Client, instanceId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithIdV2(instanceId))
	req.SetMethod(http.DELETE)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// GetVpcSubnet - get multi instances vpc and subnet
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
// 	   - *GetVpcSubnetResult: result of vpc and subnet
//     - error: nil if success otherwise the specific error
func GetVpcSubnet(cli bce.Client, reqBody *bce.Body) (*GetVpcSubnetResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSubnetUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetVpcSubnetResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ModifyInstancePassword - modify a bbc instance password
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyInstancePassword(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("changePass", "")
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func getInstanceUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI
}

func getInstanceUriWithId(id string) string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/" + id
}

func getInstanceUriWithIdV2(id string) string {
	return URI_PREFIX_V2 + REQUEST_INSTANCE_URI + "/" + id
}

func getSubnetUri() string {
	return URI_PREFIX_V1 + REQUEST_SUBNET_URI
}

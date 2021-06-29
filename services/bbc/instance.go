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

// Package bbc defines all APIs supported by the BBC service of BCE.
package bbc

import (
	"encoding/json"
	"fmt"
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
func CreateInstance(cli bce.Client, args *CreateInstanceArgs, reqBody *bce.Body) (*CreateInstanceResult,
	error) {
	clientToken := args.ClientToken
	requestToken := args.RequestToken
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	req.SetHeader("x-request-token", requestToken)

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
func GetInstanceDetailWithDeploySet(cli bce.Client, instanceId string, isDeploySet bool) (*InstanceModel, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.GET)
	if isDeploySet == true {
		req.SetParam("isDeploySet", "true")
	}
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

// GetInstanceDetail - get a bbc instance detail msg
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
// RETURNS:
//     - *InstanceModel: instance detail msg
//     - error: nil if success otherwise the specific error
func GetInstanceDetailWithDeploySetAndFailed(cli bce.Client, instanceId string,
	isDeploySet bool, containsFailed bool) (*InstanceModel, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.GET)
	if isDeploySet == true {
		req.SetParam("isDeploySet", "true")
	}
	if containsFailed == true {
		req.SetParam("containsFailed", "true")
	} else {
		req.SetParam("containsFailed", "false")
	}
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

func GetInstanceDetail(cli bce.Client, instanceId string) (*InstanceModel, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.GET)
	req.SetParam("isDeploySet", "false")

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

// ListInstances - list all bbc instances
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list bbc instances
// RETURNS:
//     - *ListInstanceResult: results of list bbc instances
//     - error: nil if success otherwise the specific error
func ListRecycledInstances(cli bce.Client, reqBody *bce.Body) (*ListRecycledInstancesResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getRecycledInstanceUri())
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

	jsonBody := &ListRecycledInstancesResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func RecoveryInstances(cli bce.Client, reqBody *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getRecoveryInstancesUri())
	req.SetMethod(http.POST)
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

// DeleteInstance - delete a specified instance,contains prepay or postpay instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be deleted
// RETURNS:
//     - error: nil if success otherwise the specific error

func DeleteBbcIngorePayment(cli bce.Client, args *DeleteInstanceIngorePaymentArgs) (*DeleteInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeleteBbcDeleteIngorePaymentUri())
	req.SetMethod(http.POST)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DeleteInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func getDeleteBbcDeleteIngorePaymentUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/delete"
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
func ModifyInstanceDesc(cli bce.Client, instanceId string, clientToken string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("updateDesc", "")
	req.SetBody(reqBody)
	if clientToken != "" {
		req.SetParam("clientToken", clientToken)
	}

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

// BatchRebuildInstances - batch rebuild instances
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to rebuild instance
// RETURNS:
//     - *BatchRebuildResponse: result of batch rebuild instances
//     - error: nil if success otherwise the specific error
func BatchRebuildInstances(cli bce.Client, reqBody *bce.Body) (*BatchRebuildResponse, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getRebuildBatchInstanceUri())
	req.SetMethod(http.PUT)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	fmt.Println(resp)
	jsonBody := &BatchRebuildResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
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

// DeleteInstance - delete a bbc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteInstances(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBatchDeleteInstanceUri())
	req.SetMethod(http.POST)
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

// BatchAddIp - Add ips to instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func BatchAddIp(cli bce.Client, args *BatchAddIpArgs, reqBody *bce.Body) (*BatchAddIpResponse, error) {
	// Build the request
	clientToken := args.ClientToken
	req := &bce.BceRequest{}
	req.SetUri(getBatchAddIpUri())
	req.SetMethod(http.PUT)
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

	jsonBody := &BatchAddIpResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// BatchAddIp - Add ips to instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func BatchAddIpCrossSubnet(cli bce.Client, args *BatchAddIpCrossSubnetArgs, reqBody *bce.Body) (*BatchAddIpResponse,
	error) {
	// Build the request
	clientToken := args.ClientToken
	req := &bce.BceRequest{}
	req.SetUri(getBatchAddIpCrossSubnetUri())
	req.SetMethod(http.PUT)
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

	jsonBody := &BatchAddIpResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// BatchDelIp - Delete ips of instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func BatchDelIp(cli bce.Client, args *BatchDelIpArgs, reqBody *bce.Body) error {
	// Build the request
	clientToken := args.ClientToken
	req := &bce.BceRequest{}
	req.SetUri(getBatchDelIpUri())
	req.SetMethod(http.PUT)
	req.SetBody(reqBody)

	if clientToken != "" {
		req.SetParam("clientToken", clientToken)
	}

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

func GetInstanceCreateStock(cli bce.Client, args *CreateInstanceStockArgs) (*InstanceStockResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getCreateInstanceStock())
	req.SetMethod(http.POST)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &InstanceStockResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetSimpleFlavor(cli bce.Client, args *GetSimpleFlavorArgs) (*SimpleFlavorResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getsimpleFlavor())
	req.SetMethod(http.POST)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &SimpleFlavorResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetInstancePirce(cli bce.Client, args *InstancePirceArgs) (*InstancePirceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstancePirce())
	req.SetMethod(http.POST)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &InstancePirceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetInstanceEni - get the eni of the bbc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the bbc instance id
// RETURNS:
//     - error: nil if success otherwise the specific error
func GetInstanceEni(cli bce.Client, instanceId string) (*GetInstanceEniResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceEniUri(instanceId))
	req.SetMethod(http.GET)
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetInstanceEniResult{}
	print(jsonBody)
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetInstanceVNC - get VNC address of the specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
// RETURNS:
//     - *GetInstanceVNCResult: result of the VNC address of the instance
//     - error: nil if success otherwise the specific error
func GetInstanceVNC(cli bce.Client, instanceId string) (*GetInstanceVNCResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceVNCUri(instanceId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetInstanceVNCResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// BatchCreateAutoRenewRules - Batch Create AutoRenew Rules
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func BatchCreateAutoRenewRules(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBatchCreateAutoRenewRulesUri())
	req.SetMethod(http.POST)
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

// BatchDeleteAutoRenewRules - Batch Delete AutoRenew Rules
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func BatchDeleteAutoRenewRules(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBatchDeleteAutoRenewRulesUri())
	req.SetMethod(http.POST)
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

// InstanceChangeVpc - change the subnet to which the instance belongs
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: request body to change subnet of instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func InstanceChangeSubnet(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getChangeSubnetUri())
	req.SetMethod(http.PUT)
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

// InstanceChangeVpc - change the vpc to which the instance belongs
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: request body to change vpc of instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func InstanceChangeVpc(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getChangeVpcUri())
	req.SetMethod(http.PUT)
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

func getInstanceVNCUri(instanceId string) string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/" + instanceId + "/vnc"
}

func getInstanceEniUri(instanceId string) string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_PORT_URI + "/" + instanceId
}

func getInstanceUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI
}

func getRecycledInstanceUri() string {
	return URI_PREFIX_V1 + REQUEST_RECYCLE_URI + REQUEST_INSTANCE_URI
}

func getRecoveryInstancesUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + REQUEST_RECOVERY_URI
}

func getInstanceUriWithId(id string) string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/" + id
}

func getBatchAddIpUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + REQUEST_BATCHADDIP_URI
}

func getBatchAddIpCrossSubnetUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + REQUEST_BATCHADDIPCROSSSUBNET_URI
}

func getBatchDelIpUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + REQUEST_BATCHDELIP_URI
}

func getBatchDeleteInstanceUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + REQUEST_BATCH_DELETE_URI
}

func getInstanceUriWithIdV2(id string) string {
	return URI_PREFIX_V2 + REQUEST_INSTANCE_URI + "/" + id
}

func getSubnetUri() string {
	return URI_PREFIX_V1 + REQUEST_SUBNET_URI
}

func getCreateInstanceStock() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/stock/createInstance"
}

func getsimpleFlavor() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/simpleFlavor"
}

func getInstancePirce() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/price"
}

func getBatchCreateAutoRenewRulesUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + REQUEST_BATCH_CREATE_AUTORENEW_RULES_URI
}

func getBatchDeleteAutoRenewRulesUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + REQUEST_BATCH_Delete_AUTORENEW_RULES_URI
}

func getRebuildBatchInstanceUri() string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + REQUEST_BATCH_REBUILD_INSTANCE_URI
}

func getChangeSubnetUri() string {
	return URI_PREFIX_V1 + "/subnet" + "/changeSubnet"
}

func getChangeVpcUri() string {
	return URI_PREFIX_V1 + REQUEST_VPC_URI + "/changeVpc"
}

// GetStockWithDeploySet - get the bbc's stock with deploySet
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to get the bbc's stock with deploySet
// RETURNS:
//     - *GetBbcStocksResult: the result of the bbc's stock
//     - error: nil if success otherwise the specific error
func GetStockWithDeploySet(cli bce.Client, args *GetBbcStockArgs) (*GetBbcStocksResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(geBbcStockWithDeploySetUri())
	req.SetMethod(http.POST)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetBbcStocksResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}
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

// instance.go - the instance APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"encoding/json"
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
func CreateInstance(cli bce.Client, args *CreateInstanceArgs, reqBody *bce.Body) (*CreateInstanceResult,
	error) {
	// Build the request
	clientToken := args.ClientToken
	requestToken := args.RequestToken
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

// CreateInstanceBySpec - create an instance with specified spec.
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to create instance
// RETURNS:
//     - *CreateInstanceBySpecResult: result of the instance ids newly created
//     - error: nil if success otherwise the specific error
func CreateInstanceBySpec(cli bce.Client, args *CreateInstanceBySpecArgs, reqBody *bce.Body) (
	*CreateInstanceBySpecResult, error) {
	// Build the request
	clientToken := args.ClientToken
	requestToken := args.RequestToken
	req := &bce.BceRequest{}
	req.SetUri(getInstanceBySpecUri())
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

	jsonBody := &CreateInstanceBySpecResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// CreateInstanceV3 - create an instance with specified spec.
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to create instance
// RETURNS:
//     - *CreateInstanceV3Result: result of the instance ids newly created
//     - error: nil if success otherwise the specific error
func CreateInstanceV3(cli bce.Client, args *CreateInstanceV3Args, reqBody *bce.Body) (
	*CreateInstanceV3Result, error) {
	// Build the request
	clientToken := args.ClientToken
	requestToken := args.RequestToken
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriV3())
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

	jsonBody := &CreateInstanceV3Result{}
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
		if len(args.DedicatedHostId) != 0 {
			req.SetParam("dedicateHostId", args.DedicatedHostId)
		}
		if len(args.InternalIp) != 0 {
			req.SetParam("internalIp", args.InternalIp)
		}
		if len(args.ZoneName) != 0 {
			req.SetParam("zoneName", args.ZoneName)
		}
		if len(args.KeypairId) != 0 {
			req.SetParam("keypairId", args.KeypairId)
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

// ListRecycleInstances - list all instances in the recycle bin with the specified parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list instances in the recycle bin
// RETURNS:
//     - *ListRecycleInstanceResult: result of the instance in the recycle bin list
//     - error: nil if success otherwise the specific error
func ListRecycleInstances(cli bce.Client, args *ListRecycleInstanceArgs) (*ListRecycleInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getRecycleInstanceListUri())
	req.SetMethod(http.POST)

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
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

	jsonBody := &ListRecycleInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// listServersByMarkerV3 - list all instances  with the specified parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list instances
// RETURNS:
//     - *LogicMarkerResultResponseV3: result of the instance
//     - error: nil if success otherwise the specific error
func ListServersByMarkerV3(cli bce.Client, args *ListServerRequestV3Args) (*LogicMarkerResultResponseV3, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getServersByMarkerV3Uri())
	req.SetMethod(http.POST)

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
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

	jsonBody := &LogicMarkerResultResponseV3{}
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
	req.SetParam("isDeploySet", "false")

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

func GetInstanceDetailWithDeploySet(cli bce.Client, instanceId string, isDeploySet bool) (*GetInstanceDetailResult, error) {
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

	jsonBody := &GetInstanceDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func GetInstanceDetailWithDeploySetAndFailed(cli bce.Client, instanceId string,
	isDeploySet bool, containsFailed bool) (*GetInstanceDetailResult, error) {
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

	jsonBody := &GetInstanceDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// DeleteInstance - delete a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be deleted
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteInstance(cli bce.Client, instanceId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
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

// DeleteInstance - delete a specified instance,contains prepay or postpay instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be deleted
// RETURNS:
//     - error: nil if success otherwise the specific error

func DeleteInstanceIngorePayment(cli bce.Client, args *DeleteInstanceIngorePaymentArgs) (*DeleteInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeleteInstanceDeleteIngorePaymentUri())
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

// DeleteRecycledInstance - delete a recycled bcc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteRecycledInstance(cli bce.Client, instanceId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeleteRecycledInstanceUri(instanceId))
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

// AutoReleaseInstance - set releaseTime of a postpay instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the specific instance ID
//     - args: the arguments to auto release instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func AutoReleaseInstance(cli bce.Client, instanceId string, args *AutoReleaseArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("autorelease", "")
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)
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

// ResizeInstance - resize a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be resized
//     - reqBody: the request body to resize instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func ResizeInstance(cli bce.Client, instanceId, clientToken string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("resize", "")
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

// RebuildInstance - rebuild a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be rebuilded
//     - reqBody: the request body to rebuild instance
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

// StartInstance - start a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be started
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

// StopInstance - stop a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be stoped
//	   - reqBody: the request body to stop instance
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

// RebootInstance - reboot a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be rebooted
//	   - reqBody: the request body to reboot instance
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

func RecoveryInstance(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getRecoveryInstanceUri())
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

// ChangeInstancePass - change password of specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
//	   - reqBody: the request body to change paasword of instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func ChangeInstancePass(cli bce.Client, instanceId string, reqBody *bce.Body) error {
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

// ModifyDeletionProtection - Modify deletion protection of specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
//	   - reqBody: the request body to set an instance, default 0 for deletable and 1 for deletion protection
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyDeletionProtection(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceDeletionProtectionUri(instanceId))
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

// ModifyInstanceAttribute - modify attribute of a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be modified
//	   - reqBody: the request body to modify instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyInstanceAttribute(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("modifyAttribute", "")
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

// ModifyInstanceDesc - modify desc of a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be modified
//	   - reqBody: the request body to modify instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyInstanceDesc(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("modifyDesc", "")
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

// ModifyInstanceHostname - modify hostname of a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be modified
//	   - reqBody: the request body to modify instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyInstanceHostname(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("changeHostname", "")
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

// BindSecurityGroup - bind security group for a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
//	   - reqBody: the request body to bind security group associate to the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func BindSecurityGroup(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("bind", "")
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

// UnBindSecurityGroup - unbind security group for a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance
//	   - reqBody: the request body to unbind security group associate to the instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func UnBindSecurityGroup(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("unbind", "")
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

// InstancePurchaseReserved - renew a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be renewed
//     - reqBody: the request body to renew instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func InstancePurchaseReserved(cli bce.Client, instanceId, relatedRenewFlag, clientToken string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("purchaseReserved", "")
	req.SetParam("relatedRenewFlag", relatedRenewFlag)
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

// DeleteInstanceWithRelatedResource - delete an instance with related resources
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be deleted
//     - reqBody: request body to delete instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteInstanceWithRelatedResource(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
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

// InstanceChangeSubnet - change the subnet to which the instance belongs
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

// ResizeInstanceBySpec - resize a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: id of the instance to be resized
//     - reqBody: the request body to resize instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func ResizeInstanceBySpec(cli bce.Client, instanceId, clientToken string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getResizeInstanceBySpec(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("resize", "")
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

// BatchRebuildInstances - batch rebuild instances
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to rebuild instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func BatchRebuildInstances(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getRebuildBatchInstanceUri())
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

// ChangeToPrepaid - to prepaid
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to ChangeToPrepaid
// RETURNS:
//     - error: nil if success otherwise the specific error
func ChangeToPrepaid(cli bce.Client, instanceId string, reqBody *bce.Body) (*ChangeToPrepaidResponse, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getChangeToPrepaidUri(instanceId))
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	req.SetParam("toPrepay", "")

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ChangeToPrepaidResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// bindInstanceToTags - bind instance to tags
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to bindInstanceToTags
// RETURNS:
//     - error: nil if success otherwise the specific error
func BindInstanceToTags(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getbindInstanceToTagsUri(instanceId))
	req.SetMethod(http.PUT)
	req.SetBody(reqBody)
	req.SetParam("bind", "")

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

// UnBindInstanceToTags - unbind instance to tags
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to unbindInstanceToTags
// RETURNS:
//     - error: nil if success otherwise the specific error
func UnBindInstanceToTags(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getbindInstanceToTagsUri(instanceId))
	req.SetMethod(http.PUT)
	req.SetBody(reqBody)
	req.SetParam("unbind", "")

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

// GetInstanceNoChargeList - get instance with nocharge list
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list instances
// RETURNS:
//     - *ListInstanceResult: result of the instance list
//     - error: nil if success otherwise the specific error
func GetInstanceNoChargeList(cli bce.Client, args *ListInstanceArgs) (*ListInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(GetInstanceNoChargeListUri())
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
		if len(args.ZoneName) != 0 {
			req.SetParam("zoneName", args.ZoneName)
		}
		if len(args.KeypairId) != 0 {
			req.SetParam("keypairId", args.KeypairId)
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

// createBidInstance - create an instance with specified parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to create instance
// RETURNS:
//     - *CreateInstanceResult: result of the instance ids newly created
//     - error: nil if success otherwise the specific error
func CreateBidInstance(cli bce.Client, clientToken string, reqBody *bce.Body) (*CreateInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(GetCreateBidInstanceUri())
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

// CancelBidOrder - Cancel the bidding instance order.
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to cancel bid order
// RETURNS:
//     - error: nil if success otherwise the specific error
func CancelBidOrder(cli bce.Client, clientToken string, reqBody *bce.Body) (*CreateBidInstanceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(GetCancelBidOrderUri())
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

	jsonBody := &CreateBidInstanceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetBidInstancePrice - get the market price of the specified bidding instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - *GetBidInstancePriceResult: result of the market price of the specified bidding instance
//     - error: nil if success otherwise the specific error
func GetBidInstancePrice(cli bce.Client, clientToken string, reqBody *bce.Body) (*GetBidInstancePriceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBidInstancePriceUri())
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

	jsonBody := &GetBidInstancePriceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// ListBidFlavor - list all flavors of the bidding instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - *ListBidFlavorResult: result of the flavor list
//     - error: nil if success otherwise the specific error
func ListBidFlavor(cli bce.Client) (*ListBidFlavorResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(listBidFlavorUri())
	req.SetMethod(http.POST)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListBidFlavorResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetInstanceResizeStock(cli bce.Client, args *ResizeInstanceStockArgs) (*InstanceStockResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getResizeInstanceStock())
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

// GetAllStocks - get the bcc and bbc's stock
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - *GetAllStocksResult: the result of the bcc and bbc's stock
//     - error: nil if success otherwise the specific error
func GetAllStocks(cli bce.Client) (*GetAllStocksResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getAllStocks())
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetAllStocksResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetStockWithDeploySet - get the bcc's stock with deploySet
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to get the bcc's stock with deploySet
// RETURNS:
//     - *GetAllStocksResult: the result of the bcc's stock
//     - error: nil if success otherwise the specific error
func GetStockWithDeploySet(cli bce.Client, args *GetStockWithDeploySetArgs) (*GetStockWithDeploySetResults, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getStockWithDeploySet())
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

	jsonBody := &GetStockWithDeploySetResults{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
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
	//print(resp)

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
	//print(resp)

	defer func() { resp.Body().Close() }()
	return nil
}

// ListInstanceByInstanceIds - list instance by instanceId
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - *ListInstancesResult: result of the instance list
//     - error: nil if success otherwise the specific error
func ListInstanceByInstanceIds(cli bce.Client, args *ListInstanceByInstanceIdArgs) (*ListInstancesResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getListInstancesByIdsUrl())
	req.SetMethod(http.POST)

	if args != nil {
		if len(args.Marker) != 0 {
			req.SetParam("marker", args.Marker)
		}
		if args.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(args.MaxKeys))
		}
	}
	if args == nil || args.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}
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

	jsonBody := &ListInstancesResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

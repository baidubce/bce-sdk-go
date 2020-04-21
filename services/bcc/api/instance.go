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

// CreateInstanceBySpec - create an instance with specified spec.
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: the request body to create instance
// RETURNS:
//     - *CreateInstanceBySpecResult: result of the instance ids newly created
//     - error: nil if success otherwise the specific error
func CreateInstanceBySpec(cli bce.Client, clientToken string, reqBody *bce.Body) (*CreateInstanceBySpecResult, error)  {
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

	jsonBody := &CreateInstanceBySpecResult{}
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
func InstancePurchaseReserved(cli bce.Client, instanceId, clientToken string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceUriWithId(instanceId))
	req.SetMethod(http.PUT)
	req.SetParam("purchaseReserved", "")
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

// BatchAddIp - Add ips to instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func BatchAddIp(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBatchAddIpUri())
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

// BatchDelIp - Delete ips of instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func BatchDelIp(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBatchDelIpUri())
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


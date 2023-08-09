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

// autoSnapshotPolicy.go - the autoSnapshotPolicy APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"encoding/json"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateAutoSnapshotPolicy - create an automatic snapshot policy
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to create automatic snapshot policy
//
// RETURNS:
//   - *CreateASPResult: the ID of the automatic snapshot policy newly created
//   - error: nil if success otherwise the specific error
func CreateAutoSnapshotPolicy(cli bce.Client, args *CreateASPArgs) (*CreateASPResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUri())
	req.SetMethod(http.POST)

	if args != nil && len(args.ClientToken) != 0 {
		req.SetParam("clientToken", args.ClientToken)
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

	jsonBody := &CreateASPResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// AttachAutoSnapshotPolicy - attach an automatic snapshot policy to specified volumes
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - aspId: the id of the automatic snapshot policy
//   - args: the arguments to attach automatic snapshot policy
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func AttachAutoSnapshotPolicy(cli bce.Client, aspId string, args *AttachASPArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUriWithId(aspId))
	req.SetMethod(http.PUT)

	req.SetParam("attach", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
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

// DetachAutoSnapshotPolicy - detach an automatic snapshot policy for specified volumes
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - aspId: the id of the automatic snapshot policy
//   - args: the arguments to detach automatic snapshot policy
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DetachAutoSnapshotPolicy(cli bce.Client, aspId string, args *DetachASPArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUriWithId(aspId))
	req.SetMethod(http.PUT)

	req.SetParam("detach", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
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

// DeleteAutoSnapshotPolicy - delete an automatic snapshot policy
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - aspId: the id of the automatic snapshot policy
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DeleteAutoSnapshotPolicy(cli bce.Client, aspId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUriWithId(aspId))
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

// ListAutoSnapshotPolicy - list all automatic snapshot policies with the specified parameters
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryArgs: the arguments to list automatic snapshot policies
//   - :
//
// RETURNS:
//   - *ListASPResult: the result of the automatic snapshot policies
//   - error: nil if success otherwise the specific error
func ListAutoSnapshotPolicy(cli bce.Client, queryArgs *ListASPArgs) (*ListASPResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.Marker) != 0 {
			req.SetParam("marker", queryArgs.Marker)
		}
		if queryArgs.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(queryArgs.MaxKeys))
		}
		if len(queryArgs.AspName) != 0 {
			req.SetParam("aspName", queryArgs.AspName)
		}
		if len(queryArgs.VolumeName) != 0 {
			req.SetParam("volumeName", queryArgs.VolumeName)
		}
	}

	if queryArgs == nil || queryArgs.MaxKeys == 0 {
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

	jsonBody := &ListASPResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetAutoSnapshotPolicyDetail - get details of the specified automatic snapshot policy
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - aspId: the id of the automatic snapshot policy
//
// RETURNS:
//   - *GetASPDetailResult: the result of the given automatic snapshot policy
//   - error: nil if success otherwise the specific error
func GetAutoSnapshotPolicyDetail(cli bce.Client, aspId string) (*GetASPDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUriWithId(aspId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetASPDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// UpdateAutoSnapshotPolicy - update an automatic snapshot policy
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to update automatic snapshot policy
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func UpdateAutoSnapshotPolicy(cli bce.Client, args *UpdateASPArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getASPUri() + "/update")
	req.SetMethod(http.PUT)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
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

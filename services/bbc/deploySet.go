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

// deploySet.go - the deploy set APIs definition supported by the BBC service

// Package bbc defines all APIs supported by the BBC service of BCE.
package bbc

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

// CreateDeploySet - create a deploy set
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - clientToken: idempotent token,  an ASCII string no longer than 64 bits
//     - reqBody: http request body
// RETURNS:
//     - *CreateDeploySetResult: results of creating a deploy set
//     - error: nil if success otherwise the specific error
func CreateDeploySet(cli bce.Client, clientToken string, reqBody *bce.Body) (*CreateDeploySetResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUri())
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

	jsonBody := &CreateDeploySetResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ListDeploySets - list all deploy sets
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - *ListDeploySetsResult: the result of list all deploy sets
//     - error: nil if success otherwise the specific error
func ListDeploySets(cli bce.Client) (*ListDeploySetsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUri())
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListDeploySetsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ListDeploySets - list all deploy sets
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the filter of deployset
// RETURNS:
//     - *ListDeploySetsResult: the result of list all deploy sets
//     - error: nil if success otherwise the specific error
func ListDeploySetsPage(cli bce.Client, args *ListDeploySetsArgs) (*ListDeploySetsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUri())
	req.SetMethod(http.GET)

	// Optional arguments settings
	if args != nil {
		if len(args.Marker) != 0 {
			req.SetParam("marker", args.Marker)
		}
		if args.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(args.MaxKeys))
		}
		if len(args.Strategy) != 0 {
			req.SetParam("strategy", args.Strategy)
		}
	}
	if args == nil || args.MaxKeys == 0 {
		req.SetParam("maxKeys", "500")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListDeploySetsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetDeploySet - get details of the deploy set
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - *GetDeploySetResult: the detail of the deploy set
//     - error: nil if success otherwise the specific error
func GetDeploySet(cli bce.Client, deploySetId string) (*DeploySetResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUriWithId(deploySetId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &DeploySetResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// DeleteDeploySet - delete a deploy set
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteDeploySet(cli bce.Client, deploySetId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeploySetUriWithId(deploySetId))
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

func getDeploySetUri() string {
	return URI_PREFIX_V1 + REQUEST_DEPLOY_SET_URI
}

func getDeploySetUriWithId(id string) string {
	return URI_PREFIX_V1 + REQUEST_DEPLOY_SET_URI + "/" + id
}

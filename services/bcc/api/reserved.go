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
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func BindReservedInstanceToTags(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(GetBccReservedToTagsUri())
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

func UnbindReservedInstanceFromTags(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(GetBccReservedToTagsUri())
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

func CreateReservedInstance(cli bce.Client, clientToken string, reqBody *bce.Body) (*CreateReservedInstanceResponse, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getCreateReservedInstanceUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)
	if len(clientToken) > 0 {
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

	jsonBody := &CreateReservedInstanceResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ModifyReservedInstances(cli bce.Client, clientToken string, reqBody *bce.Body) (*ModifyReservedInstancesResponse, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getModifyReservedInstancesUri())
	req.SetMethod(http.PUT)
	req.SetBody(reqBody)
	if len(clientToken) > 0 {
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

	jsonBody := &ModifyReservedInstancesResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}


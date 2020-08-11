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

// tag.go - the tag APIs definition supported by the BBC service

// Package bbc defines all APIs supported by the BBC service of BCE.
package bbc

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// bindTags - bind a bbc instance tags
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func BindTags(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBindTagsUriWithId(instanceId))
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

// UnbindTags - unbind a bbc instance tags
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the id of the instance
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func UnbindTags(cli bce.Client, instanceId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getUnbindTagsUriWithId(instanceId))
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

func getBindTagsUriWithId(id string) string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/" + id + "/tag"
}

func getUnbindTagsUriWithId(id string) string {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + "/" + id + "/tag"
}

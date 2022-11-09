/*
 * Copyright 2021 Baidu, Inc.
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

package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateRole(cli bce.Client, body *bce.Body) (*CreateRoleResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_ROLE)
	req.SetMethod(http.POST)
	req.SetBody(body)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateRoleResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func GetRole(cli bce.Client, roleName string) (*GetRoleResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getRoleUri(roleName))
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetRoleResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func UpdateRole(cli bce.Client, roleName string, body *bce.Body) (*UpdateRoleResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getRoleUri(roleName))
	req.SetMethod(http.PUT)
	req.SetBody(body)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &UpdateRoleResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func DeleteRole(cli bce.Client, roleName string) error {
	req := &bce.BceRequest{}
	req.SetUri(getRoleUri(roleName))
	req.SetMethod(http.DELETE)

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

func ListRole(cli bce.Client) (*ListRoleResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_ROLE)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListRoleResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func getRoleUri(roleName string) string {
	return URI_PREFIX + URI_ROLE + "/" + roleName
}

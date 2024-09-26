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

func CreateUser(cli bce.Client, body *bce.Body) (*CreateUserResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_USER)
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
	jsonBody := &CreateUserResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func GetUser(cli bce.Client, name string) (*GetUserResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(name))
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetUserResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func DeleteUser(cli bce.Client, name string) error {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(name))
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

func UpdateUser(cli bce.Client, name string, body *bce.Body) (*UpdateUserResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(name))
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
	jsonBody := &UpdateUserResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func ListUser(cli bce.Client) (*ListUserResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_USER)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListUserResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func UpdateUserLoginProfile(cli bce.Client, name string, body *bce.Body) (*UpdateUserLoginProfileResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(name) + "/loginProfile")
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
	jsonBody := &UpdateUserLoginProfileResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func GetUserLoginProfile(cli bce.Client, name string) (*GetUserLoginProfileResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(name) + "/loginProfile")
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetUserLoginProfileResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func DeleteUserLoginProfile(cli bce.Client, name string) error {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(name) + "/loginProfile")
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

func UserOperationMfaSwitch(cli bce.Client, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_USER + "/operation/mfa/switch")
	req.SetMethod(http.POST)
	req.SetBody(body)

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

func SubUserUpdate(cli bce.Client, body *bce.Body, name string) (*UpdateUserResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getSubUserUri(name) + "/update")
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
	jsonBody := &UpdateUserResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func GetSubUserIdpConfig(cli bce.Client) (*IdpWithStatus, error) {
	req := &bce.BceRequest{}
	req.SetUri("/v1/subUser/idp/query")
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &IdpWithStatus{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func UpdateSubUserIdpStatus(cli bce.Client, status string) (*IdpWithStatus, error) {
	req := &bce.BceRequest{}
	req.SetUri("/v1/subUser/idp/updateStatus")
	req.SetMethod(http.POST)
	req.SetParam("status", status)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &IdpWithStatus{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil

}

func UpdateSubUserIdp(cli bce.Client, body *bce.Body) (*IdpWithStatus, error) {
	req := &bce.BceRequest{}
	req.SetUri("/v1/subUser/idp/update")
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
	jsonBody := &IdpWithStatus{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil

}

func DeleteSubUserIdp(cli bce.Client) error {
	req := &bce.BceRequest{}
	req.SetUri("/v1/subUser/idp/delete")
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

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

func getUserUri(name string) string {
	return URI_PREFIX + URI_USER + "/" + name
}

func getSubUserUri(name string) string {
	return URI_PREFIX + SUB_USER + "/" + name
}

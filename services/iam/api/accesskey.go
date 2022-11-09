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

func CreateAccessKey(cli bce.Client, userName string) (*CreateAccessKeyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getAccessKeyUri(userName))
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateAccessKeyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func DisableAccessKey(cli bce.Client, userName, accessKeyId string) (*UpdateAccessKeyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getAccessKeyWithIdUri(userName, accessKeyId))
	req.SetParam(ACCESSKEY_STATUS_DISABLE, "")
	req.SetMethod(http.PUT)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &UpdateAccessKeyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func EnableAccessKey(cli bce.Client, userName, accessKeyId string) (*UpdateAccessKeyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getAccessKeyWithIdUri(userName, accessKeyId))
	req.SetParam(ACCESSKEY_STATUS_ENABLE, "")
	req.SetMethod(http.PUT)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &UpdateAccessKeyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func DeleteAccessKey(cli bce.Client, userName, accessKeyId string) error {
	req := &bce.BceRequest{}
	req.SetUri(getAccessKeyWithIdUri(userName, accessKeyId))
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

func ListAccessKey(cli bce.Client, userName string) (*ListAccessKeyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getAccessKeyUri(userName))
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListAccessKeyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func getAccessKeyUri(userName string) string {
	return URI_PREFIX + URI_USER + "/" + userName + URI_ACCESSKEY
}

func getAccessKeyWithIdUri(userName, accessKeyId string) string {
	return getAccessKeyUri(userName) + "/" + accessKeyId
}

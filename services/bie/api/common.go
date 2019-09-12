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

// commond.go - commmon and shared logic

// Package api defines all APIs supported by the BIE service of BCE.
package api

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

type PostHttpReq struct {
	Url    string
	Body   interface{}
	Result interface{}
	Params map[string]string
}

type GetHttpReq struct {
	Url    string
	Result interface{}
	Params map[string]string
}

func Post(cli bce.Client, phr *PostHttpReq) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)

	return PostOrPut(cli, phr, req)
}

func Put(cli bce.Client, phr *PostHttpReq) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)

	return PostOrPut(cli, phr, req)
}

func PostOrPut(cli bce.Client, phr *PostHttpReq, req *bce.BceRequest) error {
	req.SetUri(phr.Url)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	if phr.Body != nil {
		jsonBytes, jsonErr := json.Marshal(phr.Body)
		if jsonErr != nil {
			return jsonErr
		}
		// fmt.Println(string(jsonBytes))
		bodyObj, err := bce.NewBodyFromBytes(jsonBytes)
		if err != nil {
			return err
		}

		req.SetBody(bodyObj)
	}

	if phr.Params != nil {
		req.SetParams(phr.Params)
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	if phr.Result != nil {
		if err := resp.ParseJsonBody(phr.Result); err != nil {
			return err
		}
	}
	return nil
}

func Get(cli bce.Client, ghr *GetHttpReq) error {
	req := &bce.BceRequest{}
	req.SetUri(ghr.Url)
	req.SetMethod(http.GET)

	if ghr.Params != nil {
		req.SetParams(ghr.Params)
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	if err := resp.ParseJsonBody(ghr.Result); err != nil {
		return err
	}
	return nil
}

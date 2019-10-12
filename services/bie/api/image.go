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

// iamge.go - the docker image related APIs definition supported by the BIE service

// Package api defines all APIs supported by the BIE service of BCE.
package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	PREFIX_V3MODSYS = "/v3/module/system"
	PREFIX_V3MODUSR = "/v3/module/user"
)

// ListImageSys - list all system docker images
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - ListImageReq: list request parameters
// RETURNS:
//     - *ListImageResult: the result iamge list
//     - error: nil if ok otherwise the specific error
func ListImageSys(cli bce.Client, lir *ListImageReq) (*ListImageResult, error) {
	url := PREFIX_V3MODSYS
	result := &ListImageResult{}
	params := map[string]string{}
	params["pageNo"] = strconv.Itoa(lir.PageNo)
	params["pageSize"] = strconv.Itoa(lir.PageSize)
	if lir.Tag != "" {
		params["tag"] = lir.Tag
	}
	req := &GetHttpReq{Url: url, Result: result, Params: params}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetImageSys - get a system docker images
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - uuid: the image uuid
// RETURNS:
//     - *Image: the result iamge
//     - error: nil if ok otherwise the specific error
func GetImageSys(cli bce.Client, uuid string) (*Image, error) {
	url := PREFIX_V3MODSYS + "/" + uuid
	result := &Image{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListImageUser - list all user docker images
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - ListImageReq: list request parameters
// RETURNS:
//     - *ListImageResult: the result iamge list
//     - error: nil if ok otherwise the specific error
func ListImageUser(cli bce.Client, lir *ListImageReq) (*ListImageResult, error) {
	url := PREFIX_V3MODUSR
	result := &ListImageResult{}
	params := map[string]string{}
	params["pageNo"] = strconv.Itoa(lir.PageNo)
	params["pageSize"] = strconv.Itoa(lir.PageSize)
	if lir.Tag != "" {
		params["tag"] = lir.Tag
	}
	req := &GetHttpReq{Url: url, Result: result, Params: params}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetImageUser - get a user docker image
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - uuid: the image uuid
// RETURNS:
//     - *Image: the result iamge
//     - error: nil if ok otherwise the specific error
func GetImageUser(cli bce.Client, uuid string) (*Image, error) {
	url := PREFIX_V3MODUSR + "/" + uuid
	result := &Image{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateImageUser - create a user docker image
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - CreateImageReq: request parameters, name, image url, description
// RETURNS:
//     - *Image: the result iamge
//     - error: nil if ok otherwise the specific error
func CreateImageUser(cli bce.Client, cir *CreateImageReq) (*Image, error) {
	url := PREFIX_V3MODUSR
	result := &Image{}
	req := &PostHttpReq{Url: url, Result: result, Body: cir}
	err := Post(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EditImageUser - edit a user docker image information
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - uuid: the image uuid
//     - EditImageReq: request parameter: description
// RETURNS:
//     - *Image: the result iamge
//     - error: nil if ok otherwise the specific error
func EditImageUser(cli bce.Client, uuid string, eir *EditImageReq) (*Image, error) {
	url := PREFIX_V3MODUSR + "/" + uuid
	result := &Image{}
	req := &PostHttpReq{Url: url, Result: result, Body: eir}
	err := Put(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteImageUser - delete a user docker image
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - uuid: the image uuid
// RETURNS:
//     - error: nil if ok otherwise the specific error
func DeleteImageUser(cli bce.Client, uuid string) error {
	url := PREFIX_V3MODUSR + "/" + uuid
	req := &bce.BceRequest{}
	req.SetUri(url)
	req.SetMethod(http.DELETE)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	return nil
}

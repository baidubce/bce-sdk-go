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

// image.go - the image APIs definition supported by the BBC service

// Package bbc defines all APIs supported by the BBC service of BCE.
package bbc

import (
	"encoding/json"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateImageFromInstanceId - create image from specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - clientToken: idempotent token,  an ASCII string no longer than 64 bits
//     - reqBody: http request body
// RETURNS:
//     - *CreateImageResult: the result of create Image
//     - error: nil if success otherwise the specific error
func CreateImageFromInstanceId(cli bce.Client, clientToken string, reqBody *bce.Body) (*CreateImageResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUri())
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

	jsonBody := &CreateImageResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

//ListImage - list all images
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list all images
// RETURNS:
//     - *ListImageResult: the result of list all images
//     - error: nil if success otherwise the specific error
func ListImage(cli bce.Client, queryArgs *ListImageArgs) (*ListImageResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.Marker) != 0 {
			req.SetParam("marker", queryArgs.Marker)
		}
		if queryArgs.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(queryArgs.MaxKeys))
		}
		if len(queryArgs.ImageType) != 0 {
			req.SetParam("imageType", queryArgs.ImageType)
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

	jsonBody := &ListImageResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetImageDetail - get an image's detail info
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: the specific image ID
// RETURNS:
//     - *GetImageDetailResult: the result of get image's detail
//     - error: nil if success otherwise the specific error
func GetImageDetail(cli bce.Client, imageId string) (*GetImageDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUriWithId(imageId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetImageDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// DeleteImage - delete an image
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: the specific image ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteImage(cli bce.Client, imageId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUriWithId(imageId))
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

// GetCommonImage - get common flavor image list
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - flavorIds: the specific flavorIds, can be nil
// RETURNS:
//     - *GetImageDetailResult: the result of get image's detail
//     - error: nil if success otherwise the specific error
func GetCommonImage(cli bce.Client, reqBody *bce.Body) (*GetImagesResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getCommonImageUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetImagesResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetCustomImage - get user onwer flavor image list
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - flavorIds: the specific flavorIds, can be nil
// RETURNS:
//     - *GetImageDetailResult: the result of get image's detail
//     - error: nil if success otherwise the specific error
func GetCustomImage(cli bce.Client, reqBody *bce.Body) (*GetImagesResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getCustomImageUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetImagesResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// ShareImage - share a specified custom image
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: id of the image to be shared
//     - args: the arguments to share image
// RETURNS:
//     - error: nil if success otherwise the specific error
func ShareImage(cli bce.Client, imageId string, args *SharedUser) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUriWithId(imageId))
	req.SetMethod(http.POST)

	req.SetParam("share", "")

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

// UnShareImage - unshare a specified image
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: id of the image to be unshared
//     - args: the arguments to unshare image
// RETURNS:
//     - error: nil if success otherwise the specific error
func UnShareImage(cli bce.Client, imageId string, args *SharedUser) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUriWithId(imageId))
	req.SetMethod(http.POST)

	req.SetParam("unshare", "")

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

func getImageUri() string {
	return URI_PREFIX_V1 + REQUEST_IMAGE_URI
}

func getImageUriWithId(id string) string {
	return URI_PREFIX_V1 + REQUEST_IMAGE_URI + "/" + id
}

func getCommonImageUri() string {
	return URI_PREFIX_V1 + REQUEST_COMMON_IMAGE_URI
}

func getCustomImageUri() string {
	return URI_PREFIX_V1 + REQUEST_CUSTOM_IMAGE_URI
}

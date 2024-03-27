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

// image.go - the image APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
	"strings"
)

// CreateImage - create an image
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to create image
// RETURNS:
//     - *CreateImageResult: the result of the image newly created
//     - error: nil if success otherwise the specific error
func CreateImage(cli bce.Client, args *CreateImageArgs) (*CreateImageResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUri())
	req.SetMethod(http.POST)

	if args.ClientToken != "" {
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

	jsonBody := &CreateImageResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// ListImage - list all images with the specified parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryArgs: the arguments to list images
// RETURNS:
//     - *ListImageResult: result of the image list
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
		if len(queryArgs.ImageName) != 0 {
			if len(queryArgs.ImageType) != 0 && strings.EqualFold("custom", queryArgs.ImageType) {
				req.SetParam("imageName", queryArgs.ImageName)
			} else {
				return nil, errors.New("only the custom image type could filter by name")
			}
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

// GetImageDetail - get details of the specified image
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: id of the image
// RETURNS:
//     - *GetImageDetailResult: result of image details
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

// DeleteImage - delete a specified image
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: id of image to be deleted
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

// RemoteCopyImage - copy custom images across regions, only custom images supported, the system \
// and service integration images cannot be copied.
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: id of the image to be copied
//     - args: the arguments to copy image
// RETURNS:
//     - error: nil if success otherwise the specific error
func RemoteCopyImage(cli bce.Client, imageId string, args *RemoteCopyImageArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUriWithId(imageId))
	req.SetMethod(http.POST)

	req.SetParam("remoteCopy", "")

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

// RemoteCopyImageReturnImageIds - copy custom images across regions, only custom images supported, the system \
// and service integration images cannot be copied.
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: id of the image to be copied
//     - args: the arguments to copy image
// RETURNS:
//     - imageIds of destination region if success otherwise the specific error
func RemoteCopyImageReturnImageIds(cli bce.Client, imageId string, args *RemoteCopyImageArgs) (*RemoteCopyImageResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUriWithId(imageId))
	req.SetMethod(http.POST)

	req.SetParam("remoteCopy", "")

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

	jsonBody := &RemoteCopyImageResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// CancelRemoteCopyImage - cancel the image copy across regions
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: id of the image
// RETURNS:
//     - error: nil if success otherwise the specific error
func CancelRemoteCopyImage(cli bce.Client, imageId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageUriWithId(imageId))
	req.SetMethod(http.POST)

	req.SetParam("cancelRemoteCopy", "")

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

// GetImageSharedUser - get the list of users that the image has been shared with
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - imageId: id of the image
// RETURNS:
//     - *GetImageSharedUserResult: result of the shared users
//     - error: nil if success otherwise the specific error
func GetImageSharedUser(cli bce.Client, imageId string) (*GetImageSharedUserResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageSharedUserUri(imageId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetImageSharedUserResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetImageOS - get the operating system information of the instance in batches according to the instance ids
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments of instance ids
// RETURNS:
//     - *GetImageOsResult: result of the operating system information
//     - error: nil if success otherwise the specific error
func GetImageOS(cli bce.Client, args *GetImageOsArgs) (*GetImageOsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageOsUri())
	req.SetMethod(http.POST)

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

	jsonBody := &GetImageOsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}


func BindImageToTags(cli bce.Client, imageId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageToTagsUri(imageId))
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

	return nil
}

func UnBindImageToTags(cli bce.Client, imageId string, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImageToTagsUri(imageId))
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

	return nil
}

func ImportCustomImage(cli bce.Client, args *ImportCustomImageArgs) (*ImportCustomImageResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getImportCustomImageUri())
	req.SetMethod(http.POST)
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

	jsonBody := &ImportCustomImageResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func BatchRefundResource(cli bce.Client, args *BatchRefundResourceArg) (*BatchRefundResourceResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBatchRefundResourceUri())
	req.SetMethod(http.POST)
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

	jsonBody := &BatchRefundResourceResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetAvailableImagesBySpec(cli bce.Client, args *GetAvailableImagesBySpecArg) (*GetAvailableImagesBySpecResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getAvailableImagesBySpecUri())
	req.SetMethod(http.GET)

	if len(args.Spec) > 0 {
		req.SetParam("spec", args.Spec)
	}
	if len(args.OsName) > 0 {
		req.SetParam("osName", args.OsName)
	}
	if len(args.Marker) > 0 {
		req.SetParam("marker", args.Marker)
	}
	if args.MaxKeys > 0{
		req.SetParam("maxKeys", fmt.Sprint(args.MaxKeys))
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetAvailableImagesBySpecResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}
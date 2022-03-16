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

package mms

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/mms/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "mms." + bce.DEFAULT_REGION + "." + bce.DEFAULT_DOMAIN
)

// Client of MMS service
type Client struct {
	*bce.BceClient
}

// NewClient make the MMS service client
func NewClient(ak, sk, endpoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

// InsertVideo - insert a video, and create a video recognize task
// PARAMS:
// 		- lib: vedio library name
// 		- args: vedio source and description
// RETURN:
// 		- BaseResponse: the result of insert a video
// 		- error: nil if success otherwise the specific error
func (c *Client) InsertVideo(lib string, args *api.BaseRequest) (*api.BaseResponse, error) {
	return api.InsertVideo(c, lib, args)
}

// GetInsertVideoResult - get video insert result and video recognize task result
// PARAMS:
// 		- lib: vedio library name
// 		- source: vedio source
// RETURN:
// 		- BaseResponse: the result of get result
// 		- error: nil if success otherwise the specific error
func (c *Client) GetInsertVideoResult(lib, source string) (*api.BaseResponse, error) {
	return api.GetInsertVideoResult(c, lib, source)
}

// GetInsertVideoResultById - get video insert result and video recognize task result by id
// PARAMS:
// 		- libId: vedio library id
// 		- mediaId: vedio id
// RETURN:
// 		- BaseResponse: the result of get result
// 		- error: nil if success otherwise the specific error
func (c *Client) GetInsertVideoResultById(libId, mediaId string) (*api.BaseResponse, error) {
	return api.GetInsertVideoResultById(c, libId, mediaId)
}

// DeleteVideo - delete a video
// PARAMS:
// 		- lib: vedio library name
// 		- source: vedio source
// RETURN:
// 		- BaseResponse: the result of delete video
// 		- error: nil if success otherwise the specific error
func (c *Client) DeleteVideo(lib, source string) (*api.BaseResponse, error) {
	return api.DeleteVideo(c, lib, source)
}

// DeleteVideoById - delete a video
// PARAMS:
// 		- libId: video library id
// 		- mediaId: video id
// RETURN:
// 		- BaseResponse: the result of delete video
// 		- error: nil if success otherwise the specific error
func (c *Client) DeleteVideoById(libId, mediaId string) (*api.BaseResponse, error) {
	return api.DeleteVideoById(c, libId, mediaId)
}

// InsertImage - insert a image
// PARAMS:
// 		- lib: image library name
// 		- args: image source and description
// RETURN:
// 		- BaseResponse: the result of insert a image
// 		- error: nil if success otherwise the specific error
func (c *Client) InsertImage(lib string, args *api.BaseRequest) (*api.BaseResponse, error) {
	return api.InsertImage(c, lib, args)
}

// DeleteImage - delete a image
// PARAMS:
// 		- lib: image library name
// 		- source: image source
// RETURN:
// 		- BaseResponse: the result of delete a image
// 		- error: nil if success otherwise the specific error
func (c *Client) DeleteImage(lib, source string) (*api.BaseResponse, error) {
	return api.DeleteImage(c, lib, source)
}

// DeleteImageById - delete a image
// PARAMS:
// 		- libId: image library id
// 		- mediaId: image id
// RETURN:
// 		- BaseResponse: the result of delete a image
// 		- error: nil if success otherwise the specific error
func (c *Client) DeleteImageById(libId, mediaId string) (*api.BaseResponse, error) {
	return api.DeleteImageById(c, libId, mediaId)
}

// SearchImageByImage - search images by a image
// PARAMS:
// 		- lib: image library name
// 		- args: image source and description
// RETURN:
// 		- BaseResponse: the result of search
// 		- error: nil if success otherwise the specific error
func (c *Client) SearchImageByImage(lib string, args *api.BaseRequest) (*api.SearchTaskResultResponse, error) {
	return api.SearchImageByImage(c, lib, args)
}

// SearchVideoByImage - search videos by a image
// PARAMS:
// 		- lib: video library name
// 		- args: image source and description
// RETURN:
// 		- BaseResponse: the result of search
// 		- error: nil if success otherwise the specific error
func (c *Client) SearchVideoByImage(lib string, args *api.BaseRequest) (*api.SearchTaskResultResponse, error) {
	return api.SearchVideoByImage(c, lib, args)
}

// SearchVideoByVideo - create a search videos by a video task
// PARAMS:
// 		- lib: video library name
// 		- args: video source and description
// RETURN:
// 		- BaseResponse: search task info
// 		- error: nil if success otherwise the specific error
func (c *Client) SearchVideoByVideo(lib string, args *api.BaseRequest) (*api.SearchTaskResultResponse, error) {
	return api.SearchVideoByVideo(c, lib, args)
}

// GetSearchVideoByVideoResult - get result of searching videos by video
// PARAMS:
// 		- lib: video library name
// 		- source: video source
// RETURN:
// 		- BaseResponse: the result of searching videos by video
// 		- error: nil if success otherwise the specific error
func (c *Client) GetSearchVideoByVideoResult(lib, source string) (*api.SearchTaskResultResponse, error) {
	return api.GetSearchVideoByVideoResult(c, lib, source)
}

// GetSearchVideoByVideoResultById - get result of searching videos by taskId
// PARAMS:
// 		- lib: video library name
// 		- taskId: search task id
// RETURN:
// 		- BaseResponse: the result of searching videos by video
// 		- error: nil if success otherwise the specific error
func (c *Client) GetSearchVideoByVideoResultById(lib, taskId string) (*api.SearchTaskResultResponse, error) {
	return api.GetSearchVideoByVideoResultById(c, lib, taskId)
}

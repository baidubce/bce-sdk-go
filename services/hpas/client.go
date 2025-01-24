/*
 * Copyright 2025 Baidu, Inc.
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
package hpas

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/hpas/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = ""
	DEFAULT_MAX_PARALLEL   = 10
	MULTIPART_ALIGN        = 1 << 20        // 1MB
	MIN_MULTIPART_SIZE     = 1 << 20        // 1MB
	DEFAULT_MULTIPART_SIZE = 12 * (1 << 20) // 12MB
	MAX_PART_NUMBER        = 10000
)

// Client of bcd service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient

	// Fileds that used in parallel operation for BOS service
	MaxParallel   int64
	MultipartSize int64
}

// NewClient make the bcd service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(ak) == 0 && len(sk) == 0 { // to support public-read-write request
		credentials, err = nil, nil
	} else {
		credentials, err = auth.NewBceCredentials(ak, sk)
		if err != nil {
			return nil, err
		}
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

	client := &Client{bce.NewBceClient(defaultConf, v1Signer),
		DEFAULT_MAX_PARALLEL, DEFAULT_MULTIPART_SIZE}
	return client, nil
}

// CreateHpas -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - *api.CreateHpasResp:
//     - error: the return error if any occurs
func (c *Client) CreateHpas(body *api.CreateHpasReq) (*api.CreateHpasResp, error) {
	return api.CreateHpas(c, body)
}

// DeleteHpas -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) DeleteHpas(body *api.DeleteHpasReq) error {
	return api.DeleteHpas(c, body)
}

// StopHpas -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) StopHpas(body *api.StopHpasReq) error {
	return api.StopHpas(c, body)
}

// StartHpas -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) StartHpas(body *api.StartHpasReq) error {
	return api.StartHpas(c, body)
}

// RebootHpas -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) RebootHpas(body *api.RebootHpasReq) error {
	return api.RebootHpas(c, body)
}

// ResetHpas -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) ResetHpas(body *api.ResetHpasReq) error {
	return api.ResetHpas(c, body)
}

// ModifyPasswordHpas -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) ModifyPasswordHpas(body *api.ModifyPasswordHpasReq) error {
	return api.ModifyPasswordHpas(c, body)
}

// CreateHpasCoupon -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - *api.CreateHpasCouponResp:
//     - error: the return error if any occurs
func (c *Client) CreateHpasCoupon(body *api.CreateHpasCouponReq) (
	*api.CreateHpasCouponResp, error) {
	return api.CreateHpasCoupon(c, body)
}

// DescribeCouponHpas -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - *api.ListHpasCouponByPageResp:
//     - error: the return error if any occurs
func (c *Client) DescribeCouponHpas(body *api.ListCouponHpasPageReq) (
	*api.ListHpasCouponByPageResp, error) {
	return api.DescribeCouponHpas(c, body)
}

// ListHpas -
//
// PARAMS:
//     - showRdmaTopo:
//     - body: body参数
// RETURNS:
//     - *api.ListHpasByPageResp:
//     - error: the return error if any occurs
func (c *Client) ListHpas(body *api.ListHpasPageReq) (
	*api.ListHpasByPageResp, error) {
	return api.ListHpas(c, body)
}

// ImageList - 查询镜像接口
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - *api.DescribeHpasImageResp:
//     - error: the return error if any occurs
func (c *Client) ImageList(body *api.BaseMarkerV3Req) (*api.DescribeHpasImageResp, error) {
	return api.ImageList(c, body)
}

// AttachTags -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) AttachTags(body *api.TagsOperationRequest) error {
	return api.AttachTags(c, body)
}

func (c *Client) DetachTags(body *api.TagsOperationRequest) error {
	return api.DetachTags(c, body)
}

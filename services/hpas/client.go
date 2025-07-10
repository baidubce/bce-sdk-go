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
	MaxParallel int64

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

	client := &Client{BceClient: bce.NewBceClient(defaultConf, v1Signer),
		MaxParallel: DEFAULT_MAX_PARALLEL, MultipartSize: DEFAULT_MULTIPART_SIZE}
	return client, nil
}

// CreateHpas -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.CreateHpasResp:
//   - error: the return error if any occurs
func (c *Client) CreateHpas(body *api.CreateHpasReq) (*api.CreateHpasResp, error) {
	return api.CreateHpas(c, body)
}

// DeleteHpas -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteHpas(body *api.DeleteHpasReq) error {
	return api.DeleteHpas(c, body)
}

// StopHpas -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) StopHpas(body *api.StopHpasReq) error {
	return api.StopHpas(c, body)
}

// StartHpas -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) StartHpas(body *api.StartHpasReq) error {
	return api.StartHpas(c, body)
}

// RebootHpas -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) RebootHpas(body *api.RebootHpasReq) error {
	return api.RebootHpas(c, body)
}

// ResetHpas -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) ResetHpas(body *api.ResetHpasReq) error {
	return api.ResetHpas(c, body)
}

// ModifyPasswordHpas -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) ModifyPasswordHpas(body *api.ModifyPasswordHpasReq) error {
	return api.ModifyPasswordHpas(c, body)
}

// ModifyInstancesAttribute:
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) ModifyInstancesAttribute(body *api.ModifyInstancesAttributeReq) error {
	return api.ModifyInstancesAttribute(c, body)
}

// ModifyInstancesSubnet - 修改实例的子网
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func (c *Client) ModifyInstancesSubnet(body *api.ModifyInstancesSubnetRequest) (*api.BaseV3Resp, error) {
	return api.ModifyInstancesSubnet(c, body)
}

// ModifyInstanceVpc - 修改实例的vpc
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func (c *Client) ModifyInstanceVpc(body *api.ModifyInstanceVpcRequest) (*api.BaseV3Resp, error) {
	return api.ModifyInstanceVpc(c, body)
}

// CreateReservedHpas -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.CreateReservedHpasResp:
//   - error: the return error if any occurs
func (c *Client) CreateReservedHpas(body *api.CreateReservedHpasReq) (
	*api.CreateReservedHpasResp, error) {
	return api.CreateReservedHpas(c, body)
}

// DescribeReservedHpas -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.ListReservedHpasByPageResp:
//   - error: the return error if any occurs
func (c *Client) DescribeReservedHpas(body *api.ListReservedHpasPageReq) (
	*api.ListReservedHpasByPageResp, error) {
	return api.DescribeReservedHpas(c, body)
}

// ListHpas -
//
// PARAMS:
//   - showRdmaTopo:
//   - body: body参数
//
// RETURNS:
//   - *api.ListHpasByPageResp:
//   - error: the return error if any occurs
func (c *Client) ListHpas(body *api.ListHpasPageReq) (
	*api.ListHpasByPageResp, error) {
	return api.ListHpas(c, body)
}

// ImageList - 查询镜像接口
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.DescribeHpasImageResp:
//   - error: the return error if any occurs
func (c *Client) ImageList(body *api.DescribeHpasImageReq) (*api.DescribeHpasImageResp, error) {
	return api.ImageList(c, body)
}

// ImageList - 查询镜像接口
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.DescribeHpasImageResp:
//   - error: the return error if any occurs
func (c *Client) CreateImage(body *api.CreateImageReq) (*api.CreateImageResp, error) {
	return api.CreateImage(c, body)
}

// ModifyImageAttribute - 修改自定义镜像
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func (c *Client) ModifyImageAttribute(body *api.ModifyImageAttributeReq) (*api.BaseV3Resp, error) {
	return api.ModifyImageAttribute(c, body)
}

// DeleteImages - 删除自定义镜像
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func (c *Client) DeleteImages(body *api.DeleteImagesReq) (*api.BaseV3Resp, error) {
	return api.DeleteImages(c, body)
}

// AttachTags -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) AttachTags(body *api.TagsOperationRequest) error {
	return api.AttachTags(c, body)
}

func (c *Client) DetachTags(body *api.TagsOperationRequest) error {
	return api.DetachTags(c, body)
}

// AssignPrivateIpAddresses -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.AssignIpv4Resp:
//   - error: the return error if any occurs
func (c *Client) AssignPrivateIpAddresses(body *api.AssignIpv4Req) (
	*api.AssignIpv4Resp, error) {
	return api.AssignPrivateIpAddresses(c, body)
}

// UnAssignPrivateIpAddresses -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func (c *Client) UnAssignPrivateIpAddresses(body *api.UnAssignIpv4Req) (
	*api.BaseV3Resp, error) {
	return api.UnAssignPrivateIpAddresses(c, body)
}

// AssignIpv6Addresses -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.AssignIpv6Resp:
//   - error: the return error if any occurs
func (c *Client) AssignIpv6Addresses(body *api.AssignIpv6Req) (
	*api.AssignIpv6Resp, error) {
	return api.AssignIpv6Addresses(c, body)
}

// UnAssignIpv6Addresses -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func (c *Client) UnAssignIpv6Addresses(body *api.UnAssignIpv6Req) (*api.BaseV3Resp, error) {
	return api.UnAssignIpv6Addresses(c, body)
}

// DescribeReservedHpasByMaker -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.ListReservedHpasByMakerResp:
//   - error: the return error if any occurs
func (c *Client) DescribeReservedHpasByMaker(body *api.ListReservedHpasByMakerReq) (
	*api.ListReservedHpasByMakerResp, error) {
	return api.ListReservedHpasByMaker(c, body)
}

// DescribeHPASInstancesByMaker:
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.ListHpasByMakerResp:
//   - error: the return error if any occurs
func (c *Client) DescribeHPASInstancesByMaker(body *api.ListHpasByMakerReq) (
	*api.ListHpasByMakerResp, error) {
	return api.ListHpasByMaker(c, body)
}

// DescribeHpasVncUrl -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.DescribeHpasVncUrlResp:
//   - error: the return error if any occurs
func (c *Client) DescribeHpasVncUrl(body *api.DescribeHpasVncUrlReq) (
	*api.DescribeHpasVncUrlResp, error) {
	return api.DescribeHpasVncUrl(c, body)
}

// AttachSecurityGroups -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func (c *Client) AttachSecurityGroups(body *api.SecurityGroupsReq) (*api.BaseV3Resp, error) {
	return api.AttachSecurityGroups(c, body)
}

// DescribeHpasVncUrl -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func (c *Client) ReplaceSecurityGroups(body *api.SecurityGroupsReq) (*api.BaseV3Resp, error) {
	return api.ReplaceSecurityGroups(c, body)
}

// DescribeHpasVncUrl -
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func (c *Client) DetachSecurityGroups(body *api.SecurityGroupsReq) (*api.BaseV3Resp, error) {
	return api.DetachSecurityGroups(c, body)
}

func (c *Client) DescribeInstanceInventoryQuantity(body *api.DescribeInstanceInventoryQuantityReq) (*api.DescribeInstanceInventoryQuantityResp, error) {
	return api.DescribeInstanceInventoryQuantity(c, body)
}
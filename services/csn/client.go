/*
 * Copyright 2022 Baidu, Inc.
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
package csn

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_SERVICE_DOMAIN = "http://csn.baidubce.com"
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

// AttachInstance - 将网络实例加载进云智能网。
//
// PARAMS:
//   - csnId: 云智能网的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) AttachInstance(csnId string, body *AttachInstanceRequest, clientToken string) error {
	return AttachInstance(c, csnId, body, clientToken)
}

// BindCsnBp - 带宽包绑定云智能网。
//
// PARAMS:
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) BindCsnBp(csnBpId string, body *BindCsnBpRequest, clientToken string) error {
	return BindCsnBp(c, csnBpId, body, clientToken)
}

// CreateAssociation - 创建路由表的关联关系。
//
// PARAMS:
//   - csnRtId: 云智能网路由表的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) CreateAssociation(csnRtId string, body *CreateAssociationRequest,
	clientToken string) error {
	return CreateAssociation(c, csnRtId, body, clientToken)
}

// CreateCsn - 创建云智能网。
//
// PARAMS:
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body: body参数
//
// RETURNS:
//   - *CreateCsnResponse:
//   - error: the return error if any occurs
func (c *Client) CreateCsn(body *CreateCsnRequest, clientToken string) (
	*CreateCsnResponse, error) {
	return CreateCsn(c, body, clientToken)
}

// CreateCsnBp - 创建云智能网共享带宽包。
//
// PARAMS:
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body: body参数
//
// RETURNS:
//   - *CreateCsnBpResponse:
//   - error: the return error if any occurs
func (c *Client) CreateCsnBp(body *CreateCsnBpRequest, clientToken string) (
	*CreateCsnBpResponse, error) {
	return CreateCsnBp(c, body, clientToken)
}

// CreateCsnBpLimit - 创建带宽包中两个地域间的地域带宽。
//
// PARAMS:
//   - csnBpId: 带宽包的ID
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) CreateCsnBpLimit(csnBpId string, body *CreateCsnBpLimitRequest, clientToken string) error {
	return CreateCsnBpLimit(c, csnBpId, body, clientToken)
}

// CreatePropagation - 创建路由表的学习关系。
//
// PARAMS:
//   - csnRtId: 云智能网路由表的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) CreatePropagation(csnRtId string, body *CreatePropagationRequest,
	clientToken string) error {
	return CreatePropagation(c, csnRtId, body, clientToken)
}

// CreateRouteRule - 添加云智能网路由表的路由条目。
//
// PARAMS:
//   - csnRtId: 云智能网路由表的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) CreateRouteRule(csnRtId string, body *CreateRouteRuleRequest,
	clientToken string) error {
	return CreateRouteRule(c, csnRtId, body, clientToken)
}

// DeleteAssociation - 删除云智能网路由表的关联关系。
//
// PARAMS:
//   - csnRtId: 路由表的ID
//   - attachId: 网络实例在云智能网中的身份ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteAssociation(csnRtId string, attachId string, clientToken string) error {
	return DeleteAssociation(c, csnRtId, attachId, clientToken)
}

// DeleteCsn - 删除云智能网。  已经加载了网络实例的云智能网不能直接删除，必须先卸载实例。
//
// PARAMS:
//   - csnId: 云智能网的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteCsn(csnId string, clientToken string) error {
	return DeleteCsn(c, csnId, clientToken)
}

// DeleteCsnBp - 删除带宽包。
//
// PARAMS:
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteCsnBp(csnBpId string, clientToken string) error {
	return DeleteCsnBp(c, csnBpId, clientToken)
}

// DeleteCsnBpLimit - 删除带宽包中两个地域间的地域带宽。
//
// PARAMS:
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteCsnBpLimit(csnBpId string, body *DeleteCsnBpLimitRequest,
	clientToken string) error {
	return DeleteCsnBpLimit(c, csnBpId, body, clientToken)
}

// DeletePropagation - 删除云智能网路由表的学习关系。
//
// PARAMS:
//   - csnRtId: 路由表的ID
//   - attachId: 网络实例在云智能网中的身份ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeletePropagation(csnRtId string, attachId string, clientToken string) error {
	return DeletePropagation(c, csnRtId, attachId, clientToken)
}

// DeleteRouteRule - 删除云智能网路由表的指定路由条目。
//
// PARAMS:
//   - csnRtId: 路由表的ID
//   - csnRtRuleId: 路由条目的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteRouteRule(csnRtId string, csnRtRuleId string, clientToken string) error {
	return DeleteRouteRule(c, csnRtId, csnRtRuleId, clientToken)
}

// DetachInstance - 从云智能网中移出指定的网络实例。
//
// PARAMS:
//   - csnId: 云智能网的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DetachInstance(csnId string, body *DetachInstanceRequest, clientToken string) error {
	return DetachInstance(c, csnId, body, clientToken)
}

// GetCsn - 查询云智能网详情。
//
// PARAMS:
//   - csnId: csnId
//
// RETURNS:
//   - *api.GetCsnResponse:
//   - error: the return error if any occurs
func (c *Client) GetCsn(csnId string) (*GetCsnResponse, error) {
	return GetCsn(c, csnId)
}

// GetCsnBp - 查询指定云智能网带宽包详情。
//
// PARAMS:
//   - csnBpId: 带宽包的ID
//
// RETURNS:
//   - *GetCsnBpResponse:
//   - error: the return error if any occurs
func (c *Client) GetCsnBp(csnBpId string) (*GetCsnBpResponse, error) {
	return GetCsnBp(c, csnBpId)
}

// GetCsnBpPrice - 云智能网共享带宽包查询价格
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *GetCsnBpPriceResponse:
//   - error: the return error if any occurs
func (c *Client) GetCsnBpPrice(body *GetCsnBpPriceRequest) (*GetCsnBpPriceResponse, error) {
	return GetCsnBpPrice(c, body)
}

// ListAssociation - 查询指定云智能网路由表的关联关系。
//
// PARAMS:
//   - csnRtId: 云智能网路由表的ID
//
// RETURNS:
//   - *ListAssociationResponse:
//   - error: the return error if any occurs
func (c *Client) ListAssociation(csnRtId string) (*ListAssociationResponse, error) {
	return ListAssociation(c, csnRtId)
}

// ListCsn - 查询云智能网列表。
//
// PARAMS:
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *ListCsnResponse:
//   - error: the return error if any occurs
func (c *Client) ListCsn(listCsnArgs *ListCsnArgs) (*ListCsnResponse, error) {
	return ListCsn(c, listCsnArgs)
}

// ListCsnBp - 查询云智能网带宽包列表。
//
// PARAMS:
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *ListCsnBpResponse:
//   - error: the return error if any occurs
func (c *Client) ListCsnBp(listCsnBpArgs *ListCsnBpArgs) (*ListCsnBpResponse, error) {
	return ListCsnBp(c, listCsnBpArgs)
}

// ListCsnBpLimit - 查询带宽包的地域带宽列表。
//
// PARAMS:
//   - csnBpId:
//
// RETURNS:
//   - *ListCsnBpLimitResponse:
//   - error: the return error if any occurs
func (c *Client) ListCsnBpLimit(csnBpId string) (*ListCsnBpLimitResponse, error) {
	return ListCsnBpLimit(c, csnBpId)
}

// ListCsnBpLimitByCsnId - 查询云智能网的地域带宽列表。
//
// PARAMS:
//   - csnId: 云智能网的ID
//   - body: body参数
//
// RETURNS:
//   - *ListCsnBpLimitByCsnIdResponse:
//   - error: the return error if any occurs
func (c *Client) ListCsnBpLimitByCsnId(csnId string) (
	*ListCsnBpLimitByCsnIdResponse, error) {
	return ListCsnBpLimitByCsnId(c, csnId)
}

// ListInstance - 查询指定云智能网下加载的网络实例信息。
//
// PARAMS:
//   - csnId: 云智能网的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *ListInstanceResponse:
//   - error: the return error if any occurs
func (c *Client) ListInstance(csnId string, listInstanceArgs *ListInstanceArgs) (
	*ListInstanceResponse, error) {
	return ListInstance(c, csnId, listInstanceArgs)
}

// ListPropagation - 查询指定云智能网路由表的学习关系。
//
// PARAMS:
//   - csnRtId: 云智能网路由表的ID
//
// RETURNS:
//   - *ListPropagationResponse:
//   - error: the return error if any occurs
func (c *Client) ListPropagation(csnRtId string) (*ListPropagationResponse, error) {
	return ListPropagation(c, csnRtId)
}

// ListRouteRule - 查询指定云智能网路由表的路由条目。
//
// PARAMS:
//   - csnRtId: 云智能网路由表的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000。缺省值为1000
//
// RETURNS:
//   - *ListRouteRuleResponse:
//   - error: the return error if any occurs
func (c *Client) ListRouteRule(csnRtId string, listRouteRuleArgs *ListRouteRuleArgs) (
	*ListRouteRuleResponse, error) {
	return ListRouteRule(c, csnRtId, listRouteRuleArgs)
}

// ListRouteTable - 查询云智能网的路由表列表。
//
// PARAMS:
//   - csnId: 云智能网的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *ListRouteTableResponse:
//   - error: the return error if any occurs
func (c *Client) ListRouteTable(csnId string, listRouteTableArgs *ListRouteTableArgs) (
	*ListRouteTableResponse, error) {
	return ListRouteTable(c, csnId, listRouteTableArgs)
}

// ListTgw - 查询云智能网TGW列表。
//
// PARAMS:
//   - csnId: 云智能网的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *ListTgwResponse:
//   - error: the return error if any occurs
func (c *Client) ListTgw(csnId string, listTgwArgs *ListTgwArgs) (
	*ListTgwResponse, error) {
	return ListTgw(c, csnId, listTgwArgs)
}

// ListTgwRule - 查询指定TGW的路由条目。
//
// PARAMS:
//   - csnId: 云智能网的ID
//   - tgwId: TGW的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *ListTgwRuleResponse:
//   - error: the return error if any occurs
func (c *Client) ListTgwRule(csnId string, tgwId string, listTgwRuleArgs *ListTgwRuleArgs,
) (*ListTgwRuleResponse, error) {
	return ListTgwRule(c, csnId, tgwId, listTgwRuleArgs)
}

// ResizeCsnBp - 带宽包的带宽升降级。
//
// PARAMS:
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) ResizeCsnBp(csnBpId string, body *ResizeCsnBpRequest, clientToken string) error {
	return ResizeCsnBp(c, csnBpId, body, clientToken)
}

// UnbindCsnBp - 带宽包解绑云智能网。
//
// PARAMS:
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UnbindCsnBp(csnBpId string, body *UnbindCsnBpRequest, clientToken string) error {
	return UnbindCsnBp(c, csnBpId, body, clientToken)
}

// UpdateCsn - 更新云智能网。  更新云智能网的名称和描述。
//
// PARAMS:
//   - csnId: 云智能网ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateCsn(csnId string, body *UpdateCsnRequest, clientToken string) error {
	return UpdateCsn(c, csnId, body, clientToken)
}

// UpdateCsnBp - 更新带宽包的名称信息。
//
// PARAMS:
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateCsnBp(csnBpId string, body *UpdateCsnBpRequest, clientToken string) error {
	return UpdateCsnBp(c, csnBpId, body, clientToken)
}

// UpdateCsnBpLimit - 更新带宽包中两个地域间的地域带宽。
//
// PARAMS:
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateCsnBpLimit(csnBpId string, body *UpdateCsnBpLimitRequest,
	clientToken string) error {
	return UpdateCsnBpLimit(c, csnBpId, body, clientToken)
}

// UpdateTgw - 更新TGW的名称、描述。
//
// PARAMS:
//   - csnId: 云智能网的ID
//   - tgwId: TGW实例的ID
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateTgw(csnId string, tgwId string, body *UpdateTgwRequest,
	clientToken string) error {
	return UpdateTgw(c, csnId, tgwId, body, clientToken)
}

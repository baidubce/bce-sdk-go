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
package cfw

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_SERVICE_DOMAIN = "http://cfw.baidubce.com"
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

// BindCfw - 批量实例绑定CFW策略。 - 没有规则的CFW不能绑定到实例
//
// PARAMS:
//   - cfwId: CFW的id
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) BindCfw(cfwId string, body *BindCfwRequest) error {
	return BindCfw(c, cfwId, body)
}

// CreateCfw - 创建CFW策略。
//
// PARAMS:
//   - body: body参数
//
// RETURNS:
//   - *CreateCfwResponse:
//   - error: the return error if any occurs
func (c *Client) CreateCfw(body *CreateCfwRequest) (
	*CreateCfwResponse, error) {
	return CreateCfw(c, body)
}

// CreateCfwRule - 批量创建CFW中防护规则。 - 五元组(protocol/sourceAddress/destAddress/sourcePort/destPort) + 方向(direction)不能全部相同。 - 一次最多创建100条规则。
//
// PARAMS:
//   - cfwId: CFW的id
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) CreateCfwRule(cfwId string, body *CreateCfwRuleRequest) error {
	return CreateCfwRule(c, cfwId, body)
}

// DeleteCfw - 删除指定CFW策略。 - CFW存在绑定关系时不允许删除
//
// PARAMS:
//   - cfwId: CFW的id
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteCfw(cfwId string) error {
	return DeleteCfw(c, cfwId)
}

// DeleteCfwRule - 批量删除指定CFW中某些规则。 - CFW已绑定到实例时，至少保留一条规则。
//
// PARAMS:
//   - cfwId: CFW的id
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteCfwRule(cfwId string, body *DeleteCfwRuleRequest) error {
	return DeleteCfwRule(c, cfwId, body)
}

// DisableCfw - 已绑定CFW的实例，使用该接口临时关闭CFW的防护功能。
//
// PARAMS:
//   - cfwId: CFW的id
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DisableCfw(cfwId string, body *DisableCfwRequest) error {
	return DisableCfw(c, cfwId, body)
}

// EnableCfw - 已绑定CFW并且临时关闭了防护功能的实例，使用该接口恢复CFW的防护功能。
//
// PARAMS:
//   - cfwId: CFW的id
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) EnableCfw(cfwId string, body *EnableCfwRequest) error {
	return EnableCfw(c, cfwId, body)
}

// GetCfw - 查询指定CFW策略的详情信息。
//
// PARAMS:
//   - cfwId: CFW的id
//
// RETURNS:
//   - *GetCfwResponse:
//   - error: the return error if any occurs
func (c *Client) GetCfw(cfwId string) (*GetCfwResponse, error) {
	return GetCfw(c, cfwId)
}

// ListCfw - 查询CFW策略列表信息。
//
// PARAMS:
//   - marker: 批量获取列表查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量通常不超过1000，缺省值为1000
//
// RETURNS:
//   - *ListCfwResponse:
//   - error: the return error if any occurs
func (c *Client) ListCfw(listCfwArgs *ListCfwArgs) (
	*ListCfwResponse, error) {
	return ListCfw(c, listCfwArgs)
}

// ListInstance - 查询防护边界实例的列表。
//
// PARAMS:
//   - instanceType: 实例类型，取值[ eip | nat | etGateway | peerconn | csn | ipv6Gateway ]
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量通常不超过1000，缺省值为1000
//   - status: 防护状态，取值 [ unbound | protected | unprotected ]
//   - region: 地域信息，取值 [ bj | gz | su | fsh | hkg | bd | fwh | sin ]
//   - body: body参数
//
// RETURNS:
//   - *ListInstanceResponse:
//   - error: the return error if any occurs
func (c *Client) ListInstance(body *ListInstanceRequest) (*ListInstanceResponse, error) {
	return ListInstance(c, body)
}

// UnbindCfw - 实例批量解绑CFW。
//
// PARAMS:
//   - cfwId: CFW的id
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UnbindCfw(cfwId string, body *UnbindCfwRequest) error {
	return UnbindCfw(c, cfwId, body)
}

// UpdateCfw - 更新CFW策略的基本信息。
//
// PARAMS:
//   - cfwId: CFW的id
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateCfw(cfwId string, body *UpdateCfwRequest) error {
	return UpdateCfw(c, cfwId, body)
}

// UpdateCfwRule - 修改指定CFW规则。 - 五元组(protocol/sourceAddress/destAddress/sourcePort/destPort) + 方向(direction)不能全部相同。
//
// PARAMS:
//   - cfwId: CFW策略的id
//   - cfwRuleId: CFW规则的id
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateCfwRule(cfwId string, cfwRuleId string,
	body *UpdateCfwRuleRequest) error {
	return UpdateCfwRule(c, cfwId, cfwRuleId, body)
}

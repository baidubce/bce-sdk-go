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
package dns

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_SERVICE_DOMAIN = "http://dns.baidubce.com"
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

// AddLineGroup -
//
// PARAMS:
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) AddLineGroup(body *AddLineGroupRequest, clientToken string) error {
	return AddLineGroup(c, body, clientToken)
}

// CreatePaidZone -
//
// PARAMS:
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) CreatePaidZone(body *CreatePaidZoneRequest, clientToken string) error {
	return CreatePaidZone(c, body, clientToken)
}

// CreateRecord -
//
// PARAMS:
//   - zoneName: 域名名称。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) CreateRecord(zoneName string, body *CreateRecordRequest, clientToken string) error {
	return CreateRecord(c, zoneName, body, clientToken)
}

// CreateZone -
//
// PARAMS:
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) CreateZone(body *CreateZoneRequest, clientToken string) error {
	return CreateZone(c, body, clientToken)
}

// DeleteLineGroup -
//
// PARAMS:
//   - lineId: 线路组id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteLineGroup(lineId string, clientToken string) error {
	return DeleteLineGroup(c, lineId, clientToken)
}

// DeleteRecord -
//
// PARAMS:
//   - zoneName: 域名名称。
//   - recordId: 解析记录id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteRecord(zoneName string, recordId string, clientToken string) error {
	return DeleteRecord(c, zoneName, recordId, clientToken)
}

// DeleteZone -
//
// PARAMS:
//   - zoneName: 域名的名称。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) DeleteZone(zoneName string, clientToken string) error {
	return DeleteZone(c, zoneName, clientToken)
}

// ListLineGroup -
//
// PARAMS:
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串。
//   - maxKeys: 每页包含的最大数量，最大数量通常不超过1000，缺省值为1000。
//   - body: body参数
//
// RETURNS:
//   - *ListLineGroupResponse:
//   - error: the return error if any occurs
func (c *Client) ListLineGroup(body *ListLineGroupRequest) (*ListLineGroupResponse, error) {
	return ListLineGroup(c, body.Marker, body.MaxKeys)
}

// ListRecord -
//
// PARAMS:
//   - zoneName: 域名的名称。
//   - rr: 主机记录，例如“www”。
//   - id: 解析记录id。
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串。
//   - maxKeys: 每页包含的最大数量，最大数量通常不超过1000。缺省值为1000。
//   - body: body参数
//
// RETURNS:
//   - *ListRecordResponse:
//   - error: the return error if any occurs
func (c *Client) ListRecord(zoneName string, request *ListRecordRequest) (*ListRecordResponse, error) {
	return ListRecord(c, zoneName, request.Rr, request.Id, request.Marker, request.MaxKeys)
}

// ListZone -
//
// PARAMS:
//   - name: 域名的名称，支持模糊搜索。
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
//   - body: body参数
//
// RETURNS:
//   - *ListZoneResponse:
//   - error: the return error if any occurs
func (c *Client) ListZone(body *ListZoneRequest) (
	*ListZoneResponse, error) {
	return ListZone(c, body, body.Name, body.Marker, body.MaxKeys)
}

// RenewZone -
//
// PARAMS:
//   - name: 续费的域名。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) RenewZone(name string, body *RenewZoneRequest, clientToken string) error {
	return RenewZone(c, name, body, clientToken)
}

// UpdateLineGroup -
//
// PARAMS:
//   - lineId: 线路组id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateLineGroup(lineId string, body *UpdateLineGroupRequest,
	clientToken string) error {
	return UpdateLineGroup(c, lineId, body, clientToken)
}

// UpdateRecord -
//
// PARAMS:
//   - zoneName: 域名名称。
//   - recordId: 解析记录id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateRecord(zoneName string, recordId string, body *UpdateRecordRequest,
	clientToken string) error {
	return UpdateRecord(c, zoneName, recordId, body, clientToken)
}

// UpdateRecordDisable -
//
// PARAMS:
//   - zoneName: 域名名称。
//   - recordId: 解析记录id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateRecordDisable(zoneName string, recordId string, clientToken string) error {
	return UpdateRecordDisable(c, zoneName, recordId, clientToken)
}

// UpdateRecordEnable -
//
// PARAMS:
//   - zoneName: 域名名称。
//   - recordId: 解析记录id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpdateRecordEnable(zoneName string, recordId string, clientToken string) error {
	return UpdateRecordEnable(c, zoneName, recordId, clientToken)
}

// UpgradeZone -
//
// PARAMS:
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body: body参数
//
// RETURNS:
//   - error: the return error if any occurs
func (c *Client) UpgradeZone(body *UpgradeZoneRequest, clientToken string) error {
	return UpgradeZone(c, body, clientToken)
}

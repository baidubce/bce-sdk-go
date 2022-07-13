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
package localDns

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_SERVICE_DOMAIN = "http://privatezone.baidubce.com"
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

// AddRecord -
//
// PARAMS:
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body: body参数
// RETURNS:
//     - *api.AddRecordResponse:
//     - error: the return error if any occurs
func (c *Client) AddRecord(zoneId string, body *AddRecordRequest) (
	*AddRecordResponse, error) {
	return AddRecord(c, zoneId, body, body.ClientToken)
}

// DeletePrivateZone -
//
// PARAMS:
//     - zoneId: zone的id
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) DeletePrivateZone(zoneId string, clientToken string) error {
	return DeletePrivateZone(c, zoneId, clientToken)
}

// CreatePrivateZone -
//
// PARAMS:
//     - body: body参数
// RETURNS:
//     - *api.CreatePrivateZoneResponse:
//     - error: the return error if any occurs
func (c *Client) CreatePrivateZone(body *CreatePrivateZoneRequest) (
	*CreatePrivateZoneResponse, error) {
	return CreatePrivateZone(c, body, body.ClientToken)
}

// BindVpc -
//
// PARAMS:
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) BindVpc(zoneId string, body *BindVpcRequest) error {
	return BindVpc(c, zoneId, body, body.ClientToken)
}

// DeleteRecord -
//
// PARAMS:
//     - recordId: 解析记录ID
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) DeleteRecord(recordId string, clientToken string) error {
	return DeleteRecord(c, recordId, clientToken)
}

// DisableRecord -
//
// PARAMS:
//     - recordId: 解析记录ID
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) DisableRecord(recordId string, clientToken string) error {
	return DisableRecord(c, recordId, clientToken)
}

// EnableRecord -
//
// PARAMS:
//     - recordId: 解析记录ID
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) EnableRecord(recordId string, clientToken string) error {
	return EnableRecord(c, recordId, clientToken)
}

// GetPrivateZone -
//
// PARAMS:
//     - zoneId: zone的ID
// RETURNS:
//     - *api.GetPrivateZoneResponse:
//     - error: the return error if any occurs
func (c *Client) GetPrivateZone(zoneId string) (*GetPrivateZoneResponse, error) {
	return GetPrivateZone(c, zoneId)
}

// ListPrivateZone -
//
// PARAMS:
//     - request: 获取privateZone列表的入参
// RETURNS:
//     - *api.ListPrivateZoneResponse:
//     - error: the return error if any occurs
func (c *Client) ListPrivateZone(request *ListPrivateZoneRequest) (
	*ListPrivateZoneResponse, error) {
	return ListPrivateZone(c, request.Marker, request.MaxKeys)
}

// ListRecord -
//
// PARAMS:
//     - zoneId: Zone的ID
// RETURNS:
//     - *api.ListRecordResponse:
//     - error: the return error if any occurs
func (c *Client) ListRecord(zoneId string) (*ListRecordResponse, error) {
	return ListRecord(c, zoneId)
}

// UnbindVpc -
//
// PARAMS:
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) UnbindVpc(zoneId string, body *UnbindVpcRequest) error {
	return UnbindVpc(c, zoneId, body, body.ClientToken)
}

// UpdateRecord -
//
// PARAMS:
//     - recordId: 解析记录的ID
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs
func (c *Client) UpdateRecord(recordId string, body *UpdateRecordRequest) error {
	return UpdateRecord(c, recordId, body, body.ClientToken)
}

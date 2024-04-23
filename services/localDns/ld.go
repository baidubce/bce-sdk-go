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
	"encoding/json"
	"strconv"
	"strings"
)
import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// AddRecord -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body:
// RETURNS:
//     - *api.AddRecordResponse:
//     - error: the return error if any occurs
func AddRecord(cli bce.Client, zoneId string, body *AddRecordRequest, clientToken string) (
	*AddRecordResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/privatezone/[zoneId]/record"
	path = strings.Replace(path, "[zoneId]", zoneId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	jsonBody, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(jsonBody)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &AddRecordResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// BindVpc -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func BindVpc(cli bce.Client, zoneId string, body *BindVpcRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/privatezone/[zoneId]"
	path = strings.Replace(path, "[zoneId]", zoneId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetParam("bind", "")

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	jsonBody, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(jsonBody)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// CreatePrivateZone -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body:
// RETURNS:
//     - *api.CreatePrivateZoneResponse:
//     - error: the return error if any occurs
func CreatePrivateZone(cli bce.Client, body *CreatePrivateZoneRequest, clientToken string) (
	*CreatePrivateZoneResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/privatezone"
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	jsonBody, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(jsonBody)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &CreatePrivateZoneResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// DeletePrivateZone -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - zoneId: zone的id
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func DeletePrivateZone(cli bce.Client, zoneId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/privatezone/[zoneId]"
	path = strings.Replace(path, "[zoneId]", zoneId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// DeleteRecord -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - recordId: 解析记录ID
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func DeleteRecord(cli bce.Client, recordId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/privatezone/record/[recordId]"
	path = strings.Replace(path, "[recordId]", recordId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// DisableRecord -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - recordId: 解析记录ID
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
// RETURNS:
//     - error: the return error if any occurs
func DisableRecord(cli bce.Client, recordId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/privatezone/record/[recordId]"
	path = strings.Replace(path, "[recordId]", recordId, -1)
	req.SetUri(path)
	req.SetParam("disable", "")
	req.SetParam("clientToken", clientToken)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// EnableRecord -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - recordId: 解析记录ID
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
// RETURNS:
//     - error: the return error if any occurs
func EnableRecord(cli bce.Client, recordId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/privatezone/record/[recordId]"
	path = strings.Replace(path, "[recordId]", recordId, -1)
	req.SetUri(path)
	req.SetParam("enable", "")
	req.SetParam("clientToken", clientToken)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// GetPrivateZone -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - zoneId: zone的ID
// RETURNS:
//     - *api.GetPrivateZoneResponse:
//     - error: the return error if any occurs
func GetPrivateZone(cli bce.Client, zoneId string) (*GetPrivateZoneResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/privatezone/[zoneId]"
	path = strings.Replace(path, "[zoneId]", zoneId, -1)
	req.SetUri(path)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &GetPrivateZoneResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListPrivateZone -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//     - maxKeys: 每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
// RETURNS:
//     - *api.ListPrivateZoneResponse:
//     - error: the return error if any occurs
func ListPrivateZone(cli bce.Client, marker string, maxKeys int) (
	*ListPrivateZoneResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/privatezone"
	req.SetUri(path)
	if "" != marker {
		req.SetParam("marker", marker)
	}
	if 0 != maxKeys {
		req.SetParam("maxKeys", strconv.Itoa(maxKeys))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListPrivateZoneResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListRecord -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - zoneId: Zone的ID
//     - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//     - maxKeys: 每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
//     - sourceType: 记录类型，可选值
// RETURNS:
//     - *api.ListRecordResponse:
//     - error: the return error if any occurs
func ListRecord(cli bce.Client, zoneId string, marker string, maxKeys int,
	sourceType string) (*ListRecordResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/privatezone/[zoneId]/record"
	path = strings.Replace(path, "[zoneId]", zoneId, -1)
	req.SetUri(path)

	if "" != sourceType {
		req.SetParam("sourceType", sourceType)
	}
	if "" != marker {
		req.SetParam("marker", marker)
	}
	if 0 != maxKeys {
		req.SetParam("maxKeys", strconv.Itoa(maxKeys))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListRecordResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// UnbindVpc -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func UnbindVpc(cli bce.Client, zoneId string, body *UnbindVpcRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/privatezone/[zoneId]"
	path = strings.Replace(path, "[zoneId]", zoneId, -1)
	req.SetUri(path)
	req.SetParam("unbind", "")
	req.SetParam("clientToken", clientToken)

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	jsonBody, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(jsonBody)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// UpdateRecord -
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - recordId: 解析记录的ID
//     - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func UpdateRecord(cli bce.Client, recordId string, body *UpdateRecordRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/privatezone/record/[recordId]"
	path = strings.Replace(path, "[recordId]", recordId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return err
	}
	jsonBody, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(jsonBody)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

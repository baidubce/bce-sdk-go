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
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
	"strings"
)

// AddLineGroup -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func AddLineGroup(cli bce.Client, body *AddLineGroupRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/dns/customline"
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

// CreatePaidZone -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func CreatePaidZone(cli bce.Client, body *CreatePaidZoneRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/dns/zone/order"
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

// CreateRecord -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - zoneName: 域名名称。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func CreateRecord(cli bce.Client, zoneName string, body *CreateRecordRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/dns/zone/[zoneName]/record"
	path = strings.Replace(path, "[zoneName]", zoneName, -1)
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

// CreateZone -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func CreateZone(cli bce.Client, body *CreateZoneRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/dns/zone"
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

// DeleteLineGroup -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - lineId: 线路组id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//
// RETURNS:
//   - error: the return error if any occurs
func DeleteLineGroup(cli bce.Client, lineId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/dns/customline/[lineId]"
	path = strings.Replace(path, "[lineId]", lineId, -1)
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
//   - cli: the client agent which can perform sending request
//   - zoneName: 域名名称。
//   - recordId: 解析记录id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func DeleteRecord(cli bce.Client, zoneName string, recordId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/dns/zone/[zoneName]/record/[recordId]"
	path = strings.Replace(path, "[zoneName]", zoneName, -1)
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

// DeleteZone -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - zoneName: 域名的名称。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func DeleteZone(cli bce.Client, zoneName string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/dns/zone/[zoneName]"
	path = strings.Replace(path, "[zoneName]", zoneName, -1)
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

// ListLineGroup -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串。
//   - maxKeys: 每页包含的最大数量，最大数量通常不超过1000，缺省值为1000。
//   - body:
//
// RETURNS:
//   - *api.ListLineGroupResponse:
//   - error: the return error if any occurs
func ListLineGroup(cli bce.Client, marker string, maxKeys int) (*ListLineGroupResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/dns/customline"
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
	res := &ListLineGroupResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListRecord -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - zoneName: 域名的名称。
//   - rr: 主机记录，例如“www”。
//   - id: 解析记录id。
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串。
//   - maxKeys: 每页包含的最大数量，最大数量通常不超过1000。缺省值为1000。
//   - body:
//
// RETURNS:
//   - *api.ListRecordResponse:
//   - error: the return error if any occurs
func ListRecord(cli bce.Client, zoneName string, rr string, id string,
	marker string, maxKeys int) (*ListRecordResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/dns/zone/[zoneName]/record"
	path = strings.Replace(path, "[zoneName]", zoneName, -1)
	req.SetUri(path)
	if "" != rr {
		req.SetParam("rr", rr)
	}
	if "" != id {
		req.SetParam("id", id)
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

// ListZone -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - name: 域名的名称，支持模糊搜索。
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
//   - body:
//
// RETURNS:
//   - *api.ListZoneResponse:
//   - error: the return error if any occurs
func ListZone(cli bce.Client, body *ListZoneRequest, name string, marker string, maxKeys int) (
	*ListZoneResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/dns/zone"
	req.SetUri(path)
	if "" != name {
		req.SetParam("name", name)
	}
	if "" != marker {
		req.SetParam("marker", marker)
	}
	if 0 != maxKeys {
		req.SetParam("maxKeys", strconv.Itoa(maxKeys))
	}

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
	res := &ListZoneResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// RenewZone -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - name: 续费的域名。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func RenewZone(cli bce.Client, name string, body *RenewZoneRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/dns/zone/order/[name]"
	path = strings.Replace(path, "[name]", name, -1)
	req.SetUri(path)
	req.SetParam("purchaseReserved", "")
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

// UpdateLineGroup -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - lineId: 线路组id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UpdateLineGroup(cli bce.Client, lineId string, body *UpdateLineGroupRequest,
	clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/dns/customline/[lineId]"
	path = strings.Replace(path, "[lineId]", lineId, -1)
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

// UpdateRecord -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - zoneName: 域名名称。
//   - recordId: 解析记录id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UpdateRecord(cli bce.Client, zoneName string, recordId string, body *UpdateRecordRequest,
	clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/dns/zone/[zoneName]/record/[recordId]"
	path = strings.Replace(path, "[zoneName]", zoneName, -1)
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

// UpdateRecordDisable -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - zoneName: 域名名称。
//   - recordId: 解析记录id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UpdateRecordDisable(cli bce.Client, zoneName string, recordId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/dns/zone/[zoneName]/record/[recordId]"
	path = strings.Replace(path, "[zoneName]", zoneName, -1)
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

// UpdateRecordEnable -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - zoneName: 域名名称。
//   - recordId: 解析记录id。
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UpdateRecordEnable(cli bce.Client, zoneName string, recordId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/dns/zone/[zoneName]/record/[recordId]"
	path = strings.Replace(path, "[zoneName]", zoneName, -1)
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

// UpgradeZone -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串。
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UpgradeZone(cli bce.Client, body *UpgradeZoneRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/dns/zone/order"
	req.SetUri(path)
	req.SetParam("upgradeToDiscount", "")
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

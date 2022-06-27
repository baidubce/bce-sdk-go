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
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
	"strings"
)

// BindCfw - 批量实例绑定CFW策略。 - 没有规则的CFW不能绑定到实例
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW的id
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func BindCfw(cli *Client, cfwId string, body *BindCfwRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/cfw/[cfwId]"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	req.SetUri(path)
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

// CreateCfw - 创建CFW策略。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - body:
// RETURNS:
//     - *api.CreateCfwResponse:
//     - error: the return error if any occurs
func CreateCfw(cli *Client, body *CreateCfwRequest) (*CreateCfwResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/cfw"
	req.SetUri(path)

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
	res := &CreateCfwResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// CreateCfwRule - 批量创建CFW中防护规则。 - 五元组(protocol/sourceAddress/destAddress/sourcePort/destPort) + 方向(direction)不能全部相同。 - 一次最多创建100条规则。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW的id
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func CreateCfwRule(cli *Client, cfwId string, body *CreateCfwRuleRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/cfw/[cfwId]/rule"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	req.SetUri(path)

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

// DeleteCfw - 删除指定CFW策略。 - CFW存在绑定关系时不允许删除
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW的id
// RETURNS:
//     - error: the return error if any occurs
func DeleteCfw(cli *Client, cfwId string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/cfw/[cfwId]"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	req.SetUri(path)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// DeleteCfwRule - 批量删除指定CFW中某些规则。 - CFW已绑定到实例时，至少保留一条规则。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW的id
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func DeleteCfwRule(cli *Client, cfwId string, body *DeleteCfwRuleRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/cfw/[cfwId]/delete/rule"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	req.SetUri(path)

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

// DisableCfw - 已绑定CFW的实例，使用该接口临时关闭CFW的防护功能。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW的id
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func DisableCfw(cli *Client, cfwId string, body *DisableCfwRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/cfw/[cfwId]"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	req.SetUri(path)
	req.SetParam("off", "")

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

// EnableCfw - 已绑定CFW并且临时关闭了防护功能的实例，使用该接口恢复CFW的防护功能。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW的id
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func EnableCfw(cli *Client, cfwId string, body *EnableCfwRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/cfw/[cfwId]"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	req.SetUri(path)
	req.SetParam("on", "")

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

// GetCfw - 查询指定CFW策略的详情信息。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW的id
// RETURNS:
//     - *api.GetCfwResponse:
//     - error: the return error if any occurs
func GetCfw(cli *Client, cfwId string) (*GetCfwResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/cfw/[cfwId]"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	req.SetUri(path)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &GetCfwResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListCfw - 查询CFW策略列表信息。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - listCfwArgs:
// RETURNS:
//     - *api.ListCfwResponse:
//     - error: the return error if any occurs
func ListCfw(cli *Client, listCfwArgs *ListCfwArgs) (
	*ListCfwResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/cfw"
	req.SetUri(path)
	if "" != listCfwArgs.Marker {
		req.SetParam("marker", listCfwArgs.Marker)
	}
	if 0 != listCfwArgs.MaxKeys {
		req.SetParam("maxKeys", strconv.Itoa(listCfwArgs.MaxKeys))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListCfwResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListInstance - 查询防护边界实例的列表。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - listInstanceRequest:
// RETURNS:
//     - *api.ListInstanceResponse:
//     - error: the return error if any occurs
func ListInstance(cli *Client, listInstanceRequest *ListInstanceRequest) (*ListInstanceResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/cfw/instance"
	req.SetUri(path)
	req.SetParam("instanceType", listInstanceRequest.InstanceType)
	if "" != listInstanceRequest.Marker {
		req.SetParam("marker", listInstanceRequest.Marker)
	}
	if 0 != listInstanceRequest.MaxKeys {
		req.SetParam("maxKeys", strconv.Itoa(listInstanceRequest.MaxKeys))
	}
	if "" != listInstanceRequest.Status {
		req.SetParam("status", listInstanceRequest.Status)
	}
	if "" != listInstanceRequest.Region {
		req.SetParam("region", listInstanceRequest.Region)
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListInstanceResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// UnbindCfw - 实例批量解绑CFW。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW的id
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func UnbindCfw(cli *Client, cfwId string, body *UnbindCfwRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/cfw/[cfwId]"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	req.SetUri(path)
	req.SetParam("unbind", "")

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

// UpdateCfw - 更新CFW策略的基本信息。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW的id
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func UpdateCfw(cli *Client, cfwId string, body *UpdateCfwRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/cfw/[cfwId]"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	req.SetUri(path)

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

// UpdateCfwRule - 修改指定CFW规则。 - 五元组(protocol/sourceAddress/destAddress/sourcePort/destPort) + 方向(direction)不能全部相同。
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cfwId: CFW策略的id
//     - cfwRuleId: CFW规则的id
//     - body:
// RETURNS:
//     - error: the return error if any occurs
func UpdateCfwRule(cli *Client, cfwId string, cfwRuleId string,
	body *UpdateCfwRuleRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/cfw/[cfwId]/rule/[cfwRuleId]"
	path = strings.Replace(path, "[cfwId]", cfwId, -1)
	path = strings.Replace(path, "[cfwRuleId]", cfwRuleId, -1)
	req.SetUri(path)

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

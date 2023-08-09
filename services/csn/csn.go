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
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
	"strings"
)

// AttachInstance - 将网络实例加载进云智能网。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func AttachInstance(cli bce.Client, csnId string, body *AttachInstanceRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/csn/[csnId]"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	req.SetUri(path)
	req.SetParam("attach", "")
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// BindCsnBp - 带宽包绑定云智能网。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func BindCsnBp(cli bce.Client, csnBpId string, body *BindCsnBpRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/csn/bp/[csnBpId]"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetParam("bind", "")
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// CreateAssociation - 创建路由表的关联关系。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnRtId: 云智能网路由表的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func CreateAssociation(cli bce.Client, csnRtId string, body *CreateAssociationRequest,
	clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/csn/routeTable/[csnRtId]/association"
	path = strings.Replace(path, "[csnRtId]", csnRtId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// CreateCsn - 创建云智能网。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body:
//
// RETURNS:
//   - *api.CreateCsnResponse:
//   - error: the return error if any occurs
func CreateCsn(cli bce.Client, body *CreateCsnRequest, clientToken string) (
	*CreateCsnResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/csn"
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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
	res := &CreateCsnResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// CreateCsnBp - 创建云智能网共享带宽包。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body:
//
// RETURNS:
//   - *api.CreateCsnBpResponse:
//   - error: the return error if any occurs
func CreateCsnBp(cli bce.Client, body *CreateCsnBpRequest, clientToken string) (
	*CreateCsnBpResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/csn/bp"
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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
	res := &CreateCsnBpResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// CreateCsnBpLimit - 创建带宽包中两个地域间的地域带宽。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId: 带宽包的ID
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func CreateCsnBpLimit(cli bce.Client, csnBpId string, body *CreateCsnBpLimitRequest,
	clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/csn/bp/[csnBpId]/limit"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// CreatePropagation - 创建路由表的学习关系。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnRtId: 云智能网路由表的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func CreatePropagation(cli bce.Client, csnRtId string, body *CreatePropagationRequest,
	clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/csn/routeTable/[csnRtId]/propagation"
	path = strings.Replace(path, "[csnRtId]", csnRtId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// CreateRouteRule - 添加云智能网路由表的路由条目。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnRtId: 云智能网路由表的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func CreateRouteRule(cli bce.Client, csnRtId string, body *CreateRouteRuleRequest,
	clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/csn/routeTable/[csnRtId]/rule"
	path = strings.Replace(path, "[csnRtId]", csnRtId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// DeleteAssociation - 删除云智能网路由表的关联关系。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnRtId: 路由表的ID
//   - attachId: 网络实例在云智能网中的身份ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//
// RETURNS:
//   - error: the return error if any occurs
func DeleteAssociation(cli bce.Client, csnRtId string, attachId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/csn/routeTable/[csnRtId]/association/[attachId]"
	path = strings.Replace(path, "[csnRtId]", csnRtId, -1)
	path = strings.Replace(path, "[attachId]", attachId, -1)
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

// DeleteCsn - 删除云智能网。  已经加载了网络实例的云智能网不能直接删除，必须先卸载实例。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//
// RETURNS:
//   - error: the return error if any occurs
func DeleteCsn(cli bce.Client, csnId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/csn/[csnId]"
	path = strings.Replace(path, "[csnId]", csnId, -1)
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

// DeleteCsnBp - 删除带宽包。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//
// RETURNS:
//   - error: the return error if any occurs
func DeleteCsnBp(cli bce.Client, csnBpId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/csn/bp/[csnBpId]"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
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

// DeleteCsnBpLimit - 删除带宽包中两个地域间的地域带宽。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func DeleteCsnBpLimit(cli bce.Client, csnBpId string, body *DeleteCsnBpLimitRequest,
	clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/v1/csn/bp/[csnBpId]/limit/delete"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// DeletePropagation - 删除云智能网路由表的学习关系。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnRtId: 路由表的ID
//   - attachId: 网络实例在云智能网中的身份ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//
// RETURNS:
//   - error: the return error if any occurs
func DeletePropagation(cli bce.Client, csnRtId string, attachId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/csn/routeTable/[csnRtId]/propagation/[attachId]"
	path = strings.Replace(path, "[csnRtId]", csnRtId, -1)
	path = strings.Replace(path, "[attachId]", attachId, -1)
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

// DeleteRouteRule - 删除云智能网路由表的指定路由条目。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnRtId: 路由表的ID
//   - csnRtRuleId: 路由条目的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//
// RETURNS:
//   - error: the return error if any occurs
func DeleteRouteRule(cli bce.Client, csnRtId string, csnRtRuleId string, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	path := "/v1/csn/routeTable/[csnRtId]/rule/[csnRtRuleId]"
	path = strings.Replace(path, "[csnRtId]", csnRtId, -1)
	path = strings.Replace(path, "[csnRtRuleId]", csnRtRuleId, -1)
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

// DetachInstance - 从云智能网中移出指定的网络实例。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func DetachInstance(cli bce.Client, csnId string, body *DetachInstanceRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/csn/[csnId]"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	req.SetUri(path)
	req.SetParam("detach", "")
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// GetCsn - 查询云智能网详情。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: csnId
//
// RETURNS:
//   - *api.GetCsnResponse:
//   - error: the return error if any occurs
func GetCsn(cli bce.Client, csnId string) (*GetCsnResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/[csnId]"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	req.SetUri(path)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &GetCsnResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// GetCsnBp - 查询指定云智能网带宽包详情。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId: 带宽包的ID
//
// RETURNS:
//   - *api.GetCsnBpResponse:
//   - error: the return error if any occurs
func GetCsnBp(cli bce.Client, csnBpId string) (*GetCsnBpResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/bp/[csnBpId]"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
	req.SetUri(path)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &GetCsnBpResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListAssociation - 查询指定云智能网路由表的关联关系。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnRtId: 云智能网路由表的ID
//
// RETURNS:
//   - *api.ListAssociationResponse:
//   - error: the return error if any occurs
func ListAssociation(cli bce.Client, csnRtId string) (*ListAssociationResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/routeTable/[csnRtId]/association"
	path = strings.Replace(path, "[csnRtId]", csnRtId, -1)
	req.SetUri(path)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListAssociationResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListCsn - 查询云智能网列表。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *api.ListCsnResponse:
//   - error: the return error if any occurs
func ListCsn(cli bce.Client, listCsnArgs *ListCsnArgs) (*ListCsnResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn"
	req.SetUri(path)
	if "" != listCsnArgs.Marker {
		req.SetParam("marker", listCsnArgs.Marker)
	}
	if 0 != listCsnArgs.MaxKeys {
		req.SetParam("maxKeys", strconv.Itoa(listCsnArgs.MaxKeys))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListCsnResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListCsnBp - 查询云智能网带宽包列表。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *api.ListCsnBpResponse:
//   - error: the return error if any occurs
func ListCsnBp(cli bce.Client, listCsnBpArgs *ListCsnBpArgs) (*ListCsnBpResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/bp"
	req.SetUri(path)
	if "" != listCsnBpArgs.Marker {
		req.SetParam("marker", listCsnBpArgs.Marker)
	}
	if 0 != listCsnBpArgs.MaxKeys {
		req.SetParam("maxKeys", strconv.Itoa(listCsnBpArgs.MaxKeys))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListCsnBpResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListCsnBpLimit - 查询带宽包的地域带宽列表。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId:
//
// RETURNS:
//   - *api.ListCsnBpLimitResponse:
//   - error: the return error if any occurs
func ListCsnBpLimit(cli bce.Client, csnBpId string) (*ListCsnBpLimitResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/bp/[csnBpId]/limit"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
	req.SetUri(path)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListCsnBpLimitResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListCsnBpLimitByCsnId - 查询云智能网的地域带宽列表。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网的ID
//   - body:
//
// RETURNS:
//   - *api.ListCsnBpLimitByCsnIdResponse:
//   - error: the return error if any occurs
func ListCsnBpLimitByCsnId(cli bce.Client, csnId string) (
	*ListCsnBpLimitByCsnIdResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/[csnId]/bp/limit"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	req.SetUri(path)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListCsnBpLimitByCsnIdResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListInstance - 查询指定云智能网下加载的网络实例信息。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *api.ListInstanceResponse:
//   - error: the return error if any occurs
func ListInstance(cli bce.Client, csnId string, listInstanceArgs *ListInstanceArgs) (
	*ListInstanceResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/[csnId]/instance"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	req.SetUri(path)
	if "" != listInstanceArgs.Marker {
		req.SetParam("marker", listInstanceArgs.Marker)
	}
	if 0 != listInstanceArgs.MaxKeys {
		req.SetParam("maxKeys", strconv.Itoa(listInstanceArgs.MaxKeys))
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

// ListPropagation - 查询指定云智能网路由表的学习关系。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnRtId: 云智能网路由表的ID
//
// RETURNS:
//   - *api.ListPropagationResponse:
//   - error: the return error if any occurs
func ListPropagation(cli bce.Client, csnRtId string) (*ListPropagationResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/routeTable/[csnRtId]/propagation"
	path = strings.Replace(path, "[csnRtId]", csnRtId, -1)
	req.SetUri(path)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListPropagationResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListRouteRule - 查询指定云智能网路由表的路由条目。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnRtId: 云智能网路由表的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000。缺省值为1000
//
// RETURNS:
//   - *api.ListRouteRuleResponse:
//   - error: the return error if any occurs
func ListRouteRule(cli bce.Client, csnRtId string, listRouteRuleArgs *ListRouteRuleArgs) (
	*ListRouteRuleResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/routeTable/[csnRtId]/rule"
	path = strings.Replace(path, "[csnRtId]", csnRtId, -1)
	req.SetUri(path)
	if "" != listRouteRuleArgs.Marker {
		req.SetParam("marker", listRouteRuleArgs.Marker)
	}
	if 0 != listRouteRuleArgs.MaxKeys {
		req.SetParam("maxKeys", strconv.Itoa(listRouteRuleArgs.MaxKeys))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListRouteRuleResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListRouteTable - 查询云智能网的路由表列表。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *api.ListRouteTableResponse:
//   - error: the return error if any occurs
func ListRouteTable(cli bce.Client, csnId string, listRouteTableArgs *ListRouteTableArgs) (
	*ListRouteTableResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/[csnId]/routeTable"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	req.SetUri(path)
	if "" != listRouteTableArgs.Marker {
		req.SetParam("marker", listRouteTableArgs.Marker)
	}
	if 0 != listRouteTableArgs.MaxKeys {
		req.SetParam("maxKeys", strconv.Itoa(listRouteTableArgs.MaxKeys))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListRouteTableResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListTgw - 查询云智能网TGW列表。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//
// RETURNS:
//   - *api.ListTgwResponse:
//   - error: the return error if any occurs
func ListTgw(cli bce.Client, csnId string, listTgwArgs *ListTgwArgs) (*ListTgwResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/[csnId]/tgw"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	req.SetUri(path)
	if "" != listTgwArgs.Marker {
		req.SetParam("marker", listTgwArgs.Marker)
	}
	if 0 != listTgwArgs.MaxKeys {
		req.SetParam("maxKeys", strconv.Itoa(listTgwArgs.MaxKeys))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListTgwResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListTgwRule - 查询指定TGW的路由条目。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网的ID
//   - tgwId: TGW的ID
//   - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//   - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
//   - body:
//
// RETURNS:
//   - *api.ListTgwRuleResponse:
//   - error: the return error if any occurs
func ListTgwRule(cli bce.Client, csnId string, tgwId string, listTgwRuleArgs *ListTgwRuleArgs,
) (*ListTgwRuleResponse, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	path := "/v1/csn/[csnId]/tgw/[tgwId]/rule"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	path = strings.Replace(path, "[tgwId]", tgwId, -1)
	req.SetUri(path)
	if "" != listTgwRuleArgs.Marker {
		req.SetParam("marker", listTgwRuleArgs.Marker)
	}
	if 0 != listTgwRuleArgs.MaxKeys {
		req.SetParam("maxKeys", strconv.Itoa(listTgwRuleArgs.MaxKeys))
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	res := &ListTgwRuleResponse{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ResizeCsnBp - 带宽包的带宽升降级。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func ResizeCsnBp(cli bce.Client, csnBpId string, body *ResizeCsnBpRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/csn/bp/[csnBpId]"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetParam("resize", "")
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// UnbindCsnBp - 带宽包解绑云智能网。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UnbindCsnBp(cli bce.Client, csnBpId string, body *UnbindCsnBpRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/csn/bp/[csnBpId]"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetParam("unbind", "")
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// UpdateCsn - 更新云智能网。  更新云智能网的名称和描述。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UpdateCsn(cli bce.Client, csnId string, body *UpdateCsnRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/csn/[csnId]"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// UpdateCsnBp - 更新带宽包的名称信息。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UpdateCsnBp(cli bce.Client, csnBpId string, body *UpdateCsnBpRequest, clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/csn/bp/[csnBpId]"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// UpdateCsnBpLimit - 更新带宽包中两个地域间的地域带宽。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnBpId: 带宽包的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UpdateCsnBpLimit(cli bce.Client, csnBpId string, body *UpdateCsnBpLimitRequest,
	clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/csn/bp/[csnBpId]/limit"
	path = strings.Replace(path, "[csnBpId]", csnBpId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

// UpdateTgw - 更新TGW的名称、描述。
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - csnId: 云智能网的ID
//   - tgwId: TGW实例的ID
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func UpdateTgw(cli bce.Client, csnId string, tgwId string, body *UpdateTgwRequest,
	clientToken string) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.PUT)
	path := "/v1/csn/[csnId]/tgw/[tgwId]"
	path = strings.Replace(path, "[csnId]", csnId, -1)
	path = strings.Replace(path, "[tgwId]", tgwId, -1)
	req.SetUri(path)
	req.SetParam("clientToken", clientToken)
	req.SetHeader("Content-Type", "application/json;charset=UTF-8")

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

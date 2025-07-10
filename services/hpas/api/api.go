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
package api

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.CreateHpasResp:
//   - error: the return error if any occurs
func CreateHpas(cli bce.Client, body *CreateHpasReq) (*CreateHpasResp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "CreateInstances")

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
	res := &CreateHpasResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func DeleteHpas(cli bce.Client, body *DeleteHpasReq) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DeleteInstances")

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

// StopHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func StopHpas(cli bce.Client, body *StopHpasReq) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "StopInstances")

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

// StartHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func StartHpas(cli bce.Client, body *StartHpasReq) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "StartInstances")

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

// RebootHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func RebootHpas(cli bce.Client, body *RebootHpasReq) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "RebootInstances")

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

// ResetHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func ResetHpas(cli bce.Client, body *ResetHpasReq) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "ResetInstances")

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

// ModifyPasswordHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func ModifyPasswordHpas(cli bce.Client, body *ModifyPasswordHpasReq) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "ModifyInstanceApplicationPassword")

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

// ModifyInstancesAttribute -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func ModifyInstancesAttribute(cli bce.Client, body *ModifyInstancesAttributeReq) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "ModifyInstancesAttribute")

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

// CreateReservedHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.CreateReservedHpasResp:
//   - error: the return error if any occurs
func CreateReservedHpas(cli bce.Client, body *CreateReservedHpasReq) (*CreateReservedHpasResp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "CreateReservedInstances")

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
	res := &CreateReservedHpasResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// DescribeReservedHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.ListReservedHpasByPageResp:
//   - error: the return error if any occurs
func DescribeReservedHpas(cli bce.Client, body *ListReservedHpasPageReq) (
	*ListReservedHpasByPageResp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DescribeReservedInstances")

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
	res := &ListReservedHpasByPageResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListHpas -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - showRdmaTopo:
//   - body:
//
// RETURNS:
//   - *api.ListHpasByPageResp:
//   - error: the return error if any occurs
func ListHpas(cli bce.Client, body *ListHpasPageReq) (
	*ListHpasByPageResp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DescribeInstances")
	if body != nil && body.ShowRdmaTopo {
		req.SetParam("showRdmaTopo", "true")
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
	res := &ListHpasByPageResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListHpasByMaker -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.ListHpasByPageResp:
//   - error: the return error if any occurs
func ListHpasByMaker(cli bce.Client, body *ListHpasByMakerReq) (
	*ListHpasByMakerResp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DescribeHPASInstances")
	if body != nil && body.ShowRdmaTopo {
		req.SetParam("showRdmaTopo", "true")
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
	res := &ListHpasByMakerResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ListReservedHpasByMaker -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.ListHpasByPageResp:
//   - error: the return error if any occurs
func ListReservedHpasByMaker(cli bce.Client, body *ListReservedHpasByMakerReq) (
	*ListReservedHpasByMakerResp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DescribeHPASReservedInstances")

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
	res := &ListReservedHpasByMakerResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ImageList - 查询镜像接口
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.DescribeHpasImageResp:
//   - error: the return error if any occurs
func ImageList(cli bce.Client, body *DescribeHpasImageReq) (*DescribeHpasImageResp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DescribeImages")

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
	res := &DescribeHpasImageResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// CreateImage - 创建自定义镜像接口
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.DescribeHpasImageResp:
//   - error: the return error if any occurs
func CreateImage(cli bce.Client, body *CreateImageReq) (*CreateImageResp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "CreateImage")
	req.SetMethod(http.POST)

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
	res := &CreateImageResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ModifyImageAttribute - 修改自定义镜像
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func ModifyImageAttribute(cli bce.Client, body *ModifyImageAttributeReq) (*BaseV3Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "ModifyImageAttribute")
	req.SetMethod(http.POST)

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
	res := &BaseV3Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// DeleteImages - 删除自定义镜像
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.BaseV3Resp:
//   - error: the return error if any occurs
func DeleteImages(cli bce.Client, body *DeleteImagesReq) (*BaseV3Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DeleteImages")
	req.SetMethod(http.POST)

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
	res := &BaseV3Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// AttachTags -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func AttachTags(cli bce.Client, body *TagsOperationRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "AttachTags")

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

// DetachTags -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func DetachTags(cli bce.Client, body *TagsOperationRequest) error {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DetachTags")

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

// AssignPrivateIpAddresses -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.EdpResultRespAssignIpv4Resp:
//   - error: the return error if any occurs
func AssignPrivateIpAddresses(cli bce.Client, body *AssignIpv4Req) (
	*AssignIpv4Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "AssignPrivateIpAddresses")

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
	res := &AssignIpv4Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// AssignIpv6Addresses -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.EdpResultRespAssignIpv6Resp:
//   - error: the return error if any occurs
func AssignIpv6Addresses(cli bce.Client, body *AssignIpv6Req) (
	*AssignIpv6Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "AssignIpv6Addresses")

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
	res := &AssignIpv6Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// UnAssignPrivateIpAddresses -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.EdpResultRespVoid:
//   - error: the return error if any occurs
func UnAssignPrivateIpAddresses(cli bce.Client, body *UnAssignIpv4Req) (*BaseV3Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "UnassignPrivateIpAddresses")

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
	res := &BaseV3Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ModifyInstancesSubnet -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func ModifyInstancesSubnet(cli bce.Client, body *ModifyInstancesSubnetRequest) (*BaseV3Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "ModifyInstancesSubnet")

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
	res := &BaseV3Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ModifyInstanceVpc -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func ModifyInstanceVpc(cli bce.Client, body *ModifyInstanceVpcRequest) (*BaseV3Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "ModifyInstanceVPC")

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
	res := &BaseV3Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// UnAssignIpv6Addresses -
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - *api.EdpResultRespVoid:
//   - error: the return error if any occurs
func UnAssignIpv6Addresses(cli bce.Client, body *UnAssignIpv6Req) (*BaseV3Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "UnAssignIpv6Addresses")

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
	res := &BaseV3Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func DescribeHpasVncUrl(cli bce.Client, body *DescribeHpasVncUrlReq) (*DescribeHpasVncUrlResp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DescribeHpasVncUrl")

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
	res := &DescribeHpasVncUrlResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// AttachSecurityGroups - 加入安全组
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func AttachSecurityGroups(cli bce.Client, body *SecurityGroupsReq) (*BaseV3Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "AttachSecurityGroups")

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
	res := &BaseV3Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// ReplaceSecurityGroups - 变更安全组
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func ReplaceSecurityGroups(cli bce.Client, body *SecurityGroupsReq) (*BaseV3Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "ModifySecurityGroups")

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
	res := &BaseV3Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

// DetachSecurityGroups - 移除安全组
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body:
//
// RETURNS:
//   - error: the return error if any occurs
func DetachSecurityGroups(cli bce.Client, body *SecurityGroupsReq) (*BaseV3Resp, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DetachSecurityGroups")

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
	res := &BaseV3Resp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

func DescribeInstanceInventoryQuantity(cli bce.Client, body *DescribeInstanceInventoryQuantityReq) (*DescribeInstanceInventoryQuantityResp, error){
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	path := "/"
	req.SetUri(path)
	req.SetParam("action", "DescribeInstanceInventoryQuantity")

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
	res := &DescribeInstanceInventoryQuantityResp{}
	if err := resp.ParseJsonBody(res); err != nil {
		return nil, err
	}
	return res, nil
}

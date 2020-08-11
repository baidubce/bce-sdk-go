/*
 * Copyright 2020 Baidu, Inc.
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

// repairPlat.go - the repair plat APIs definition supported by the BBC service

// Package bbc defines all APIs supported by the BBC service of BCE.
package bbc

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

func ListRepairTasks(cli bce.Client, args *ListRepairTaskArgs) (*ListRepairTaskResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getRepairTaskUri())
	req.SetMethod(http.GET)

	// Optional arguments settings
	if args != nil {
		if len(args.Marker) != 0 {
			req.SetParam("marker", args.Marker)
		}
		if args.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(args.MaxKeys))
		}
		if len(args.InstanceId) != 0 {
			req.SetParam("instanceId", args.InstanceId)
		}
		if len(args.ErrResult) != 0 {
			req.SetParam("errResult", args.ErrResult)
		}
	}
	if args == nil || args.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListRepairTaskResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func ListClosedRepairTasks(cli bce.Client, args *ListClosedRepairTaskArgs) (*ListClosedRepairTaskResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getClosedRepairTaskUri())
	req.SetMethod(http.GET)

	// Optional arguments settings
	if args != nil {
		if len(args.Marker) != 0 {
			req.SetParam("marker", args.Marker)
		}
		if args.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(args.MaxKeys))
		}
		if len(args.InstanceId) != 0 {
			req.SetParam("instanceId", args.InstanceId)
		}
		if len(args.TaskId) != 0 {
			req.SetParam("taskId", args.TaskId)
		}
		if len(args.ErrResult) != 0 {
			req.SetParam("errResult", args.ErrResult)
		}
		if args.StartTime != "" {
			req.SetParam("startTime", args.StartTime)
		}
		if args.EndTime != "" {
			req.SetParam("endTime", args.EndTime)
		}
	}
	if args == nil || args.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListClosedRepairTaskResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func GetTaskDetail(cli bce.Client, instanceId string) (*GetRepairTaskResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getTaskUriWithId(instanceId))
	req.SetMethod(http.GET)
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetRepairTaskResult{}
	print(jsonBody)
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func AuthorizeRepairTask(cli bce.Client, reqBody *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getAuthorizeTaskUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func UnAuthorizeRepairTask(cli bce.Client, reqBody *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getUnAuthorizeTaskUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func ConfirmRepairTask(cli bce.Client, reqBody *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getConfirmTaskUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func DisConfirmRepairTask(cli bce.Client, reqBody *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getDisConfirmTaskUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func GetRepairTaskReocrd(cli bce.Client, reqBody *bce.Body) (*GetRepairRecords, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getTaskRecordUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetRepairRecords{}
	print(jsonBody)
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ListRule - list the repair plat rules
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - *ListRuleResult: results of listing the repair plat rules
//     - error: nil if success otherwise the specific error
func ListRule(cli bce.Client, reqBody *bce.Body) (*ListRuleResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getListRuleUri())
	req.SetMethod(http.GET)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListRuleResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetRuleDetail - get the repair plat rule detail
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - ruleId: the specified rule id
// RETURNS:
//     - *Rule: results of listing the repair plat rules
//     - error: nil if success otherwise the specific error
func GetRuleDetail(cli bce.Client, ruleId string) (*Rule, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getRuleDetailUri(ruleId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &Rule{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// CreateRule - create the repair plat rule
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - *CreateRuleResult: results of the id of the repair plat rule which is created
//     - error: nil if success otherwise the specific error
func CreateRule(cli bce.Client, reqBody *bce.Body) (*CreateRuleResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getCreateRuleUri())
	req.SetMethod(http.POST)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}

	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateRuleResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// DeleteRule - delete the repair plat rule
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteRule(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeleteRuleUri())
	req.SetMethod(http.PUT)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// DisableRule - disable the repair plat rule
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func DisableRule(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDisableRuleUri())
	req.SetMethod(http.PUT)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// EnableRule - enable the repair plat rule
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func EnableRule(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getEnableRuleUri())
	req.SetMethod(http.PUT)
	req.SetBody(reqBody)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

func getRepairTaskUri() string {
	return URI_PREFIX_V2 + REQUEST_REPAIR_TASK_URI
}

func getClosedRepairTaskUri() string {
	return URI_PREFIX_V2 + REQUEST_REPAIR_CLOSED_TASK_URI
}

func getTaskUriWithId(id string) string {
	return URI_PREFIX_V2 + REQUEST_REPAIR_TASK_URI + "/" + id
}

func getAuthorizeTaskUri() string {
	return URI_PREFIX_V2 + REQUEST_REPAIR_TASK_URI + "/authorize"
}

func getUnAuthorizeTaskUri() string {
	return URI_PREFIX_V2 + REQUEST_REPAIR_TASK_URI + "/unauthorize"
}

func getConfirmTaskUri() string {
	return URI_PREFIX_V2 + REQUEST_REPAIR_TASK_URI + "/confirm"
}

func getDisConfirmTaskUri() string {
	return URI_PREFIX_V2 + REQUEST_REPAIR_TASK_URI + "/disconfirm"
}

func getTaskRecordUri() string {
	return URI_PREFIX_V2 + REQUEST_REPAIR_TASK_URI + "/record"
}

func getListRuleUri() string {
	return URI_PREFIX_V2 + REQUEST_RULE_URI
}

func getRuleDetailUri(ruleId string) string {
	return URI_PREFIX_V2 + REQUEST_RULE_URI + "/" + ruleId
}

func getCreateRuleUri() string {
	return URI_PREFIX_V2 + REQUEST_RULE_URI + REQUEST_CREATE_URI
}

func getDeleteRuleUri() string {
	return URI_PREFIX_V2 + REQUEST_RULE_URI + REQUEST_DELETE_URI
}

func getEnableRuleUri() string {
	return URI_PREFIX_V2 + REQUEST_RULE_URI + REQUEST_ENABLE_URI
}

func getDisableRuleUri() string {
	return URI_PREFIX_V2 + REQUEST_RULE_URI + REQUEST_DISABLE_URI
}

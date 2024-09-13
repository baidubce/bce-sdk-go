/*
 * Copyright 2021 Baidu, Inc.
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

// logstore.go - the logStore APIs definition supported by the BLS service

package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateProject - create project
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body: project parameters body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func CreateProject(cli bce.Client, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(PROJECT_PREFIX)
	req.SetMethod(http.POST)
	if body != nil {
		req.SetBody(body)
	}
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

// UpdateProject - update project
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body: logStore parameters body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func UpdateProject(cli bce.Client, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(PROJECT_PREFIX)
	req.SetMethod(http.PUT)
	req.SetBody(body)
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

// DescribeProject - get logStore info
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - uuid: project uuid
//
// RETURNS:
//   - *Project: project info
//   - error: nil if success otherwise the specific error
func DescribeProject(cli bce.Client, UUID string) (*Project, error) {
	req := &bce.BceRequest{}
	req.SetUri(getProjectUri(UUID))
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &DescribeProjectResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result.Result.Project, nil
}

// DeleteProject - delete project
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - uuid: project uuid
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DeleteProject(cli bce.Client, UUID string) error {
	req := &bce.BceRequest{}
	req.SetUri(getProjectUri(UUID))
	req.SetMethod(http.DELETE)
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

// ListProject - get all pattern-match project info
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - project: logstore project
//   - body: conditions project should match
//
// RETURNS:
//   - *ListProjectResult: project result set
//   - error: nil if success otherwise the specific error
func ListProject(cli bce.Client, body *bce.Body) (*ListProjectResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(LIST_PROJECT_PREFIX)
	req.SetMethod(http.POST)
	if body != nil {
		req.SetBody(body)
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &ListProjectResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result.Result, nil
}

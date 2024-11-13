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

// CreateDownloadTask - create download task
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body: download task parameters body
//
// RETURNS:
//   - string: download task uuid
//   - error: nil if success otherwise the specific error
func CreateDownloadTask(cli bce.Client, body *bce.Body) (string, error) {
	req := &bce.BceRequest{}
	req.SetUri(DOWNLOAD_TASK_PREFIX)
	req.SetMethod(http.POST)
	if body != nil {
		req.SetBody(body)
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return "", err
	}
	if resp.IsFail() {
		return "nil", resp.ServiceError()
	}
	result := &CreateDownloadResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return "", err
	}
	return result.Result.UUID, nil
}

// DescribeDownloadTask - get download task
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - uuid: download task uuid
//
// RETURNS:
//   - *DownloadTask: download task
//   - error: nil if success otherwise the specific error
func DescribeDownloadTask(cli bce.Client, UUID string) (*DownloadTask, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDownloadTaskUri(UUID))
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &DescribeDownloadTaskResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result.Result.Task, nil
}

// GetDownloadTaskLink - get download link
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - uuid: download task uuid
//
// RETURNS:
//   - *GetDownloadTaskLinkResult: download link info
//   - error: nil if success otherwise the specific error
func GetDownloadTaskLink(cli bce.Client, UUID string) (*GetDownloadTaskLinkResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getDownloadTaskLinkUri(UUID))
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &GetDownloadTaskLinkResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result.Result, nil
}

// DeleteDownloadTask - delete download task
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - uuid: download task uuid
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DeleteDownloadTask(cli bce.Client, UUID string) error {
	req := &bce.BceRequest{}
	req.SetUri(getDownloadTaskUri(UUID))
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

// ListDownloadTask - get all pattern-match download tasks
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body: conditions download task should match
//
// RETURNS:
//   - *ListDownloadTaskResult: download task result set
//   - error: nil if success otherwise the specific error
func ListDownloadTask(cli bce.Client, body *bce.Body) (*ListDownloadTaskResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(LIST_DOWNLOAD_TASK_PREFIX)
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
	result := &ListDownloadTaskResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result.Result, nil
}

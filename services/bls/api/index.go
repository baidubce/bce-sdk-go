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

// index.go - the Index APIs definition supported by the BLS service

package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateIndex - create index for logStore
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - logStoreName: logStore needs to be indexed
//     - body: index mappings body
// RETURNS:
//     - error: nil if success otherwise the specific error
func CreateIndex(cli bce.Client, logStoreName string, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getIndexUri(logStoreName))
	req.SetMethod(http.POST)
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

// UpdateIndex - update index info
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - logStoreName: logStore needs to be updated
//     - body: index mappings body
// RETURNS:
//     - error: nil if success otherwise the specific error
func UpdateIndex(cli bce.Client, logStoreName string, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getIndexUri(logStoreName))
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

// DeleteIndex - delete index for logStore
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - logStoreName: logStore to be deleted
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteIndex(cli bce.Client, logStoreName string) error {
	req := &bce.BceRequest{}
	req.SetUri(getIndexUri(logStoreName))
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

// DescribeIndex - get specific logStore index info
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - logStoreName: logStore needs to be get
// RETURNS:
//     - *IndexFields: index mappings info
//     - error: nil if success otherwise the specific error
func DescribeIndex(cli bce.Client, logStoreName string) (*IndexFields, error) {
	req := &bce.BceRequest{}
	req.SetUri(getIndexUri(logStoreName))
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &IndexFields{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

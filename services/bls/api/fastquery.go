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

// fastquery.go - the fastQuery APIs definition supported by the BLS service

package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateFastQuery - create a fastQuery
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body: the fastQuery body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func CreateFastQuery(cli bce.Client, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(FASTQUERY_PREFIX)
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

// DescribeFastQuery - get specific fastQuery info
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - fastQueryName: fastQuery name to get
//
// RETURNS:
//   - *FastQuery: target fastQuery info
//   - error: nil if success otherwise the specific error
func DescribeFastQuery(cli bce.Client, fastQueryName string) (*FastQuery, error) {
	req := &bce.BceRequest{}
	req.SetUri(getFastQueryUri(fastQueryName))
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &FastQuery{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateFastQuery - update specific fastQuery info
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body: update fastQuery body
//   - fastQueryName: fastQuery to update
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func UpdateFastQuery(cli bce.Client, body *bce.Body, fastQueryName string) error {
	req := &bce.BceRequest{}
	req.SetUri(getFastQueryUri(fastQueryName))
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

// DeleteFastQuery - delete specific fastQuery
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - fastQueryName: fastQuery name to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DeleteFastQuery(cli bce.Client, fastQueryName string) error {
	req := &bce.BceRequest{}
	req.SetUri(getFastQueryUri(fastQueryName))
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

// ListFastQuery - get all fastQuery info
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - project: logstore project
//   - logStore: logStore to parse
//   - args: query args to get pattern-match fastQuery
//
// RETURNS:
//   - *ListFastQueryResult: pattern-match fastQuery result
//   - error: nil if success otherwise the specific error
func ListFastQuery(cli bce.Client, project string, logStore string, args *QueryConditions) (*ListFastQueryResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(FASTQUERY_PREFIX)
	req.SetMethod(http.GET)
	if project != "" {
		req.SetParam("project", project)
	}
	if logStore != "" {
		req.SetParam("logStoreName", logStore)
	}
	// Set optional args
	if args != nil {
		if args.NamePattern != "" {
			req.SetParam("namePattern", args.NamePattern)
		}
		if args.Order != "" {
			req.SetParam("order", args.Order)
		}
		if args.OrderBy != "" {
			req.SetParam("orderBy", args.OrderBy)
		}
		if args.PageNo > 0 {
			req.SetParam("pageNo", strconv.Itoa(args.PageNo))
		}
		if args.PageSize > 0 {
			req.SetParam("pageSize", strconv.Itoa(args.PageSize))
		}
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &ListFastQueryResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

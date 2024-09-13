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

// logshipper.go - the logShipper APIs definition supported by the BLS service

package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateLogShipper - create logShipper
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body: logShipper parameters body
//
// RETURNS:
//   - string: snowflake base64 id for logShipper, empty string if creation fail
//   - error: nil if success otherwise the specific error
func CreateLogShipper(cli bce.Client, body *bce.Body) (string, error) {
	req := &bce.BceRequest{}
	req.SetUri(LOGSHIPPER_PREFIX)
	req.SetMethod(http.POST)
	if body != nil {
		req.SetBody(body)
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return "", err
	}
	if resp.IsFail() {
		return "", resp.ServiceError()
	}
	result := &CreateLogShipperResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return "", err
	}
	return result.LogShipperID, nil
}

// UpdateLogShipper - update logShipper name and destConfig
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - logShipperID: logShipperID to update
//   - body: logShipper parameters body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func UpdateLogShipper(cli bce.Client, body *bce.Body, logShipperID string) error {
	req := &bce.BceRequest{}
	req.SetUri(getLogShipperUri(logShipperID))
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

// GetLogShipper - get logShipper info
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - logShipperID: logShipper to get
//
// RETURNS:
//   - *LogShipper: logShipper info
//   - error: nil if success otherwise the specific error
func GetLogShipper(cli bce.Client, logShipperID string) (*LogShipper, error) {
	req := &bce.BceRequest{}
	req.SetUri(getLogShipperUri(logShipperID))
	req.SetMethod(http.GET)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &LogShipper{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListLogShipper - get all pattern-match logShipper info
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: conditions logShipper should match
//
// RETURNS:
//   - *ListShipperResult: logShipper result set
//   - error: nil if success otherwise the specific error
func ListLogShipper(cli bce.Client, args *ListLogShipperCondition) (*ListShipperResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(LOGSHIPPER_PREFIX)
	req.SetMethod(http.GET)
	if args != nil {
		if args.LogShipperID != "" {
			req.SetParam("logShipperID", args.LogShipperID)
		}
		if args.LogShipperName != "" {
			req.SetParam("logShipperName", args.LogShipperName)
		}
		if args.Project != "" {
			req.SetParam("project", args.Project)
		}
		if args.LogStoreName != "" {
			req.SetParam("logStoreName", args.LogStoreName)
		}
		if args.DestType != "" {
			req.SetParam("destType", args.DestType)
		}
		if args.Status != "" {
			req.SetParam("status", args.Status)
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
	result := &ListShipperResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListLogShipperRecord - get logShipper's execution records
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - logShipperID: logShipper to get
//   - args: conditions records should match
//
// RETURNS:
//   - *ListShipperRecordResult: logShipper records result set
//   - error: nil if success otherwise the specific error
func ListLogShipperRecord(cli bce.Client, logShipperID string, args *ListShipperRecordCondition) (*ListShipperRecordResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getLogShipperRecordUri(logShipperID))
	req.SetMethod(http.GET)
	if args != nil {
		if args.SinceHours > 0 {
			req.SetParam("sinceHours", strconv.Itoa(args.SinceHours))
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
	result := &ListShipperRecordResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSingleLogShipper - delete logShipper by id
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - logShipperID: logShipper to delete
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DeleteSingleLogShipper(cli bce.Client, logShipperID string) error {
	req := &bce.BceRequest{}
	req.SetUri(getLogShipperUri(logShipperID))
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

// BulkDeleteLogShipper - bulk delete logShipper by id
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body: bulkDeleteLogShipper parameters body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func BulkDeleteLogShipper(cli bce.Client, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(LOGSHIPPER_PREFIX)
	req.SetMethod(http.DELETE)
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

// SetSingleLogShipperStatus - set logShipper status by id
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - logShipperID: logShipper to set
//   - body: setSingleLogShipperStatus parameters body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetSingleLogShipperStatus(cli bce.Client, logShipperID string, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getLogShipperStatusUri(logShipperID))
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

// BulkSetLogShipperStatus - bulk set logShipper status by id
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - body: bulkSetLogShipperStatus parameters body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func BulkSetLogShipperStatus(cli bce.Client, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getBulkSetLogShipperStatusUri())
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

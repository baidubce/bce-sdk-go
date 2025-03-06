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

// logrecord.go - the logRecord APIs definition supported by the BLS service

package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// PushLogRecord - push logRecords into logStore
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - project: logstore project
//   - logStore: target logStore to store logRecords
//   - body: logRecord body
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func PushLogRecord(cli bce.Client, project string, logStore string, body *bce.Body) error {
	req := &bce.BceRequest{}
	req.SetUri(getLogRecordUri(logStore))
	req.SetParam("project", project)
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

// PullLogRecord - get logRecords from logStore
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - project: logstore project
//   - logStore: target logStore to get logRecords
//   - args: pull logRecords limitation args
//
// RETURNS:
//   - *PullLogRecordResult: pull logRecord result set
//   - error: nil if success otherwise the specific error
func PullLogRecord(cli bce.Client, project string, logStore string, args *PullLogRecordArgs) (*PullLogRecordResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getLogRecordUri(logStore))
	req.SetParam("project", project)
	req.SetMethod(http.GET)
	if args != nil {
		if args.LogStreamName != "" {
			req.SetParam("logStreamName", args.LogStreamName)
		}
		if len(args.StartDateTime) != 0 {
			req.SetParam("startDateTime", string(args.StartDateTime))
		}
		if len(args.EndDateTime) != 0 {
			req.SetParam("endDateTime", string(args.EndDateTime))
		}
		if args.Limit > 0 && args.Limit <= 1000 {
			req.SetParam("limit", strconv.Itoa(args.Limit))
		}
		if len(args.Marker) != 0 {
			req.SetParam("marker", args.Marker)
		}
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &PullLogRecordResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

// QueryLogRecord - retrieve logRecords from logStore
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - project: logstore project
//   - logStore: target logStore to retrieve logRecords
//   - args: query logRecords conditions args
//
// RETURNS:
//   - *QueryLogResult: query logRecord result set
//   - error: nil if success otherwise the specific error
func QueryLogRecord(cli bce.Client, project string, logStore string, args *QueryLogRecordArgs) (*QueryLogResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getLogRecordUri(logStore))
	req.SetParam("project", project)
	req.SetMethod(http.GET)
	if args != nil {
		req.SetParam("logStreamName", args.LogStreamName)
		req.SetParam("query", args.Query)
		req.SetParam("startDateTime", string(args.StartDateTime))
		req.SetParam("endDateTime", string(args.EndDateTime))
		if args.Limit > 0 {
			req.SetParam("limit", strconv.Itoa(args.Limit))
		}
		if len(args.Sort) > 0 {
			req.SetParam("sort", args.Sort)
		}
		if len(args.Marker) > 0 {
			req.SetParam("marker", args.Marker)
		}
		if args.SamplePercentage > 0 {
			req.SetParam("samplePercentage", strconv.FormatFloat(args.SamplePercentage, 'f', -1, 64))
		}
		if args.SampleSeed > 0 {
			req.SetParam("sampleSeed", strconv.Itoa(args.SampleSeed))
		}
	}
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &QueryLogResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

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

// logstream.go - the logStream APIs definition supported by the BLS service

package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListLogStream - get all pattern-match logStream in logStore
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - logStore: logStore to parse
//   - args: conditions logStream should match
//
// RETURNS:
//   - *ListLogStreamResult: pattern-match logStream result set
//   - error: nil if success otherwise the specific error
func ListLogStream(cli bce.Client, logStore string, args *QueryConditions) (*ListLogStreamResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getLogStreamName(logStore))
	req.SetMethod(http.GET)
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
	result := &ListLogStreamResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

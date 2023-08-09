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

// operationLog.go - the operationLog APIs definition supported by the BBC service

// Package bbc defines all APIs supported by the BBC service of BCE.
package bbc

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// GetOperationLog - get operation log
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to get operation log
//
// RETURNS:
//   - *GetOperationLogResult: results of getting operation log
//   - error: nil if success otherwise the specific error
func GetOperationLog(cli bce.Client, args *GetOperationLogArgs) (*GetOperationLogResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getOperationLogUri())
	req.SetMethod(http.GET)

	if args.Marker != "" {
		req.SetParam("marker", args.Marker)
	}

	if args.MaxKeys != 0 {
		req.SetParam("maxKeys", fmt.Sprintf("%d", args.MaxKeys))
	}

	if args.StartTime != "" {
		req.SetParam("startTime", args.StartTime)
	}

	if args.EndTime != "" {
		req.SetParam("endTime", args.EndTime)
	}
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetOperationLogResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func getOperationLogUri() string {
	return URI_PREFIX_V1 + REQUEST_OPERATION_LOG_URI
}

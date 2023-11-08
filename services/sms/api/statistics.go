/*
 * Copyright 2023 Baidu, Inc.
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

// mobile_black.go - the sms MobileBlack APIs definition supported by the SMS service

package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListStatistics - get sms statistics data
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to get sms statistics data
//
// RETURNS:
//   - error: the return error if any occurs
//   - *api.ListStatisticsResponse: the result of get sms MobileBlackList
func ListStatistics(cli bce.Client, args *ListStatisticsArgs) (*ListStatisticsResponse, error) {
	if err := CheckError(args != nil, "ListStatisticsArgs can not be nil"); err != nil {
		return nil, err
	}

	paramsMap := make(map[string]string)

	if err := CheckError(len(args.StartTime) > 0,
		"ListStatistics query start time can not be nil"); err != nil {
		return nil, err
	}

	paramsMap["startTime"] = args.StartTime + " 00:00:00"

	if err := CheckError(len(args.EndTime) > 0,
		"ListStatistics query end time can not be nil"); err != nil {
		return nil, err
	}

	paramsMap["endTime"] = args.EndTime + " 23:59:59"

	// default value
	paramsMap["smsType"] = "all"
	paramsMap["dimension"] = "day"

	if len(args.SmsType) > 0 {
		paramsMap["smsType"] = args.SmsType
	}
	if len(args.CountryType) > 0 {
		paramsMap["countryType"] = args.CountryType
	}
	if len(args.SignatureId) > 0 {
		paramsMap["signatureId"] = args.SignatureId
	}
	if len(args.TemplateCode) > 0 {
		paramsMap["templateCode"] = args.TemplateCode
	}

	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_STATISTICS)
	req.SetMethod(http.GET)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
	req.SetParams(paramsMap)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	result := &ListStatisticsResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

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

// GetPrepaidPackages - get sms prepaidpackages data
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to get sms prepaidpackages data
//
// RETURNS:
//   - error: the return error if any occurs
//   - *api.GetPrepaidPackageResponse: the result of get sms PrepaidPackages
func GetPrepaidPackages(cli bce.Client, args *GetPrepaidPackageArgs) (*GetPrepaidPackageResponse, error) {
	if err := CheckError(args != nil, "GetPrepaidPackageArgs can not be nil"); err != nil {
		return nil, err
	}

	paramsMap := make(map[string]string)

	if err := CheckError(len(args.UserID) > 0,
		"PrepaidPackages query user id can not be nil"); err != nil {
		return nil, err
	}
	if len(args.CountryType) > 0 {
		paramsMap["countryType"] = args.CountryType
	}
	if len(args.PackageStatus) > 0 {
		paramsMap["packageStatus"] = args.PackageStatus
	}
	if len(args.PageNo) > 0 {
		paramsMap["pageNo"] = args.PageNo
	}
	if len(args.PageSize) > 0 {
		paramsMap["pageSize"] = args.PageSize
	}

	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_PREPAY + args.UserID)
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

	result := &GetPrepaidPackageResponse{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

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
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateMobileBlack - create an sms MobileBlack
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to create an sms mobileBlack
// RETURNS:
//     - error: the return error if any occurs
func CreateMobileBlack(cli bce.Client, args *CreateMobileBlackArgs) error {
	if err := CheckError(args != nil, "CreateMobileBlackArgs can not be nil"); err != nil {
		return err
	}
	if err := CheckError(len(args.Type) > 0, "type can not be blank"); err != nil {
		return err
	}
	if err := CheckError(len(args.Phone) > 0, "phone can not be blank"); err != nil {
		return err
	}
	if args.Type == "SignatureBlack" {
		if err := CheckError(len(args.SmsType) > 0,
			"smsType can not be blank, when 'type' is 'SignatureBlack'."); err != nil {
			return err
		}
		if err := CheckError(len(args.SignatureIdStr) > 0,
			"signatureIdStr can not be blank, when 'type' is 'SignatureBlack'."); err != nil {
			return err
		}
	}

	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_BLACK)
	req.SetMethod(http.POST)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// DeleteMobileBlack - delete sms mobileBlack by phones
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to delete sms mobileBlack
// RETURNS:
//     - error: the return error if any occurs
func DeleteMobileBlack(cli bce.Client, args *DeleteMobileBlackArgs) error {
	if err := CheckError(args != nil, "DeleteMobileBlackArgs can not be nil"); err != nil {
		return err
	}
	if err := CheckError(len(args.Phones) > 0, "Phones can not be blank"); err != nil {
		return err
	}
	return bce.NewRequestBuilder(cli).
		WithMethod(http.DELETE).
		WithURL(REQUEST_URI_BLACK+"/delete").
		WithQueryParam("phones", args.Phones).
		Do()
}

// GetMobileBlack - get sms mobileBlackList
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to get sms mobileBlackList
// RETURNS:
//     - error: the return error if any occurs
//     - *api.GetMobileBlackResult: the result of get sms MobileBlackList
func GetMobileBlack(cli bce.Client, args *GetMobileBlackArgs) (*GetMobileBlackResult, error) {
	if err := CheckError(args != nil, "GetMobileBlackArgs can not be nil"); err != nil {
		return nil, err
	}

	paramsMap := make(map[string]string)
	if len(args.Phone) > 0 {
		paramsMap["phone"] = args.Phone
	}
	if len(args.SmsType) > 0 {
		paramsMap["smsType"] = args.SmsType
	}
	if len(args.SignatureIdStr) > 0 {
		paramsMap["signatureIdStr"] = args.SignatureIdStr
	}
	if len(args.StartTime) > 0 {
		paramsMap["startTime"] = args.StartTime
	}
	if len(args.EndTime) > 0 {
		paramsMap["endTime"] = args.EndTime
	}
	if len(args.PageNo) > 0 {
		paramsMap["pageNo"] = args.PageNo
	}
	if len(args.PageSize) > 0 {
		paramsMap["pageSize"] = args.PageSize
	}

	req := &bce.BceRequest{}
	req.SetUri(REQUEST_URI_BLACK)
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

	result := &GetMobileBlackResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	return result, nil
}

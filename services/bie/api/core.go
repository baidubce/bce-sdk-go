/*
 * Copyright 2017 Baidu, Inc.
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

// core.go - the core APIs definition supported by the BIE service

// Package api defines all APIs supported by the BIE service of BCE.
package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

// ListCore - list all core of a group
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - groupUuid: id of the group
// RETURNS:
//     - *ListCoreResult: the result core list
//     - error: nil if ok otherwise the specific error
func ListCore(cli bce.Client, groupUuid string) (*ListCoreResult, error) {
	url := PREFIX + "/" + groupUuid + "/core"

	result := &ListCoreResult{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCore - get a core of a group
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - groupUuid: id of the group
//     - coreid: id of the core
// RETURNS:
//     - *CoreResult: the result core
//     - error: nil if ok otherwise the specific error
func GetCore(cli bce.Client, groupUuid string, coreid string) (*CoreResult, error) {
	url := PREFIX + "/" + groupUuid + "/core/" + coreid

	result := &CoreResult{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// RenewCoreAuth - renew the auth of a core
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - coreid: id of the core
// RETURNS:
//     - *CoreInfo: the result core info
//     - error: nil if ok otherwise the specific error
func RenewCoreAuth(cli bce.Client, coreid string) (*CoreInfo, error) {
	url := "/v3/core/" + coreid + "/renew-auth"

	result := &CoreInfo{}
	req := &PostHttpReq{Url: url, Result: result}
	err := Post(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetCoreStatus - get the status of a core
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - coreid: id of the core
// RETURNS:
//     - *CoreStatus: the status of the core
//     - error: nil if ok otherwise the specific error
func GetCoreStatus(cli bce.Client, coreid string) (*CoreStatus, error) {
	url := "/v3/core/" + coreid + "/online"

	result := &CoreStatus{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

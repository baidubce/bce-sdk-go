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

// group.go - the group APIs definition supported by the BIE service

// Package api defines all APIs supported by the BIE service of BCE.
package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListGroup - list all groups of the account
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *ListGroupReq: page no, page size, and name
// RETURNS:
//     - *ListGroupResult: the result group list structure
//     - error: nil if ok otherwise the specific error
func ListGroup(cli bce.Client, lgr *ListGroupReq) (*ListGroupResult, error) {
	url := PREFIX
	params := map[string]string{}
	if lgr.Name != "" {
		params["name"] = lgr.Name
	}
	if lgr.PageNo != 0 {
		params["pageNo"] = strconv.Itoa(lgr.PageNo)
	}
	if lgr.PageSize != 0 {
		params["pageSize"] = strconv.Itoa(lgr.PageSize)
	}
	result := &ListGroupResult{}
	req := &GetHttpReq{Url: url, Result: result, Params: params}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetGroup - get a group of the account
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - groupUuid: id of the group
// RETURNS:
//     - *ListGroupResult: the result group list structure
//     - error: nil if ok otherwise the specific error
func GetGroup(cli bce.Client, groupUuid string) (*Group, error) {
	url := PREFIX + "/" + groupUuid

	result := &Group{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateGroup - create a group
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cgr: parameters to create group
// RETURNS:
//     - *CreateGroupResult: the result group
//     - error: nil if ok otherwise the specific error
func CreateGroup(cli bce.Client, cgr *CreateGroupReq) (*CreateGroupResult, error) {
	url := PREFIX
	params := map[string]string{"withCore": "true"}

	result := &CreateGroupResult{}
	req := &PostHttpReq{Url: url, Result: result, Body: cgr, Params: params}
	err := Post(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// EditGroup - edit a group
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - groupId: the group to edit
//     - egr: parameters to update group
// RETURNS:
//     - *Group: the result group
//     - error: nil if ok otherwise the specific error
func EditGroup(cli bce.Client, groupId string, egr *EditGroupReq) (*Group, error) {
	url := PREFIX + "/" + groupId
	params := map[string]string{"withCore": "true"}

	result := &Group{}
	req := &PostHttpReq{Url: url, Result: result, Body: egr, Params: params}
	err := Put(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteGroup - delete a group of the account
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - groupUuid: id of the group
// RETURNS:
//     - error: nil if ok otherwise the specific error
func DeleteGroup(cli bce.Client, groupUuid string) error {
	req := &bce.BceRequest{}
	req.SetUri(PREFIX + "/" + groupUuid)
	req.SetMethod(http.DELETE)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	return nil
}

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

package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreateGroup(cli bce.Client, body *bce.Body) (*CreateGroupResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_GROUP)
	req.SetMethod(http.POST)
	req.SetBody(body)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &CreateGroupResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func GetGroup(cli bce.Client, name string) (*GetGroupResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getGroupUri(name))
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetGroupResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func UpdateGroup(cli bce.Client, name string, body *bce.Body) (*UpdateGroupResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getGroupUri(name))
	req.SetMethod(http.PUT)
	req.SetBody(body)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &UpdateGroupResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func DeleteGroup(cli bce.Client, name string) error {
	req := &bce.BceRequest{}
	req.SetUri(getGroupUri(name))
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

func ListGroup(cli bce.Client) (*ListGroupResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_GROUP)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListGroupResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func AddUserToGroup(cli bce.Client, userName string, groupName string) error {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_GROUP + "/" + groupName + URI_USER + "/" + userName)
	req.SetMethod(http.PUT)

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

func DeleteUserFromGroup(cli bce.Client, userName string, groupName string) error {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_GROUP + "/" + groupName + URI_USER + "/" + userName)
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

func ListUsersInGroup(cli bce.Client, name string) (*ListUsersInGroupResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getGroupUri(name) + URI_USER)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListUsersInGroupResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func ListGroupsForUser(cli bce.Client, name string) (*ListGroupsForUserResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(name) + URI_GROUP)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListGroupsForUserResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func getGroupUri(name string) string {
	return URI_PREFIX + URI_GROUP + "/" + name
}

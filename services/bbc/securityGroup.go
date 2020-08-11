/*
 * Copyright (C) 2020 Baidu, Inc. All Rights Reserved.
 */
package bbc

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// BindSecurityGroups - Bind Security Groups
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func BindSecurityGroups(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getBindSecurityGroupsUri())
	req.SetMethod(http.POST)
	req.SetParam("bind", "")
	req.SetBody(reqBody)

	// Send request and get response
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



// UnBindSecurityGroups - UnBind Security Groups
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - reqBody: http request body
// RETURNS:
//     - error: nil if success otherwise the specific error
func UnBindSecurityGroups(cli bce.Client, reqBody *bce.Body) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getUnBindSecurityGroupsUri())
	req.SetMethod(http.POST)
	req.SetParam("unbind", "")
	req.SetBody(reqBody)

	// Send request and get response
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

func getBindSecurityGroupsUri() string  {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + SECURITY_GROUP_URI
}

func getUnBindSecurityGroupsUri() string  {
	return URI_PREFIX_V1 + REQUEST_INSTANCE_URI + SECURITY_GROUP_URI
}
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

// instanceSnapshot.go - the instance snapshot APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateInstanceSnapshot - create an instance snapshot (consistent snapshot group)
// for the specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to create instance snapshot
// RETURNS:
//     - *CreateInstanceSnapshotResult: result of the instance snapshot newly created
//     - error: nil if success otherwise the specific error
func CreateInstanceSnapshot(cli bce.Client, args *CreateInstanceSnapshotArgs) (
	*CreateInstanceSnapshotResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getCreateInstanceSnapshotUri())
	req.SetMethod(http.POST)

	if len(args.ClientToken) != 0 {
		req.SetParam("clientToken", args.ClientToken)
	}

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateInstanceSnapshotResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// DeleteInstanceSnapshot - delete the specified instance snapshots
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to delete instance snapshots
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteInstanceSnapshot(cli bce.Client, args *DeleteInstanceSnapshotArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeleteInstanceSnapshotUri())
	req.SetMethod(http.POST)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	defer func() { resp.Body().Close() }()
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// RenameInstanceSnapshot - rename the specified instance snapshots
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to rename instance snapshots
// RETURNS:
//     - error: nil if success otherwise the specific error
func RenameInstanceSnapshot(cli bce.Client, args *RenameInstanceSnapshotArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getRenameInstanceSnapshotUri())
	req.SetMethod(http.POST)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	defer func() { resp.Body().Close() }()
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// ListInstanceSnapshot - list instance snapshots with the marker paging
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to list instance snapshots
// RETURNS:
//     - *ListInstanceSnapshotResult: result of the instance snapshot list
//     - error: nil if success otherwise the specific error
func ListInstanceSnapshot(cli bce.Client, args *ListInstanceSnapshotArgs) (
	*ListInstanceSnapshotResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getListInstanceSnapshotUri())
	req.SetMethod(http.POST)

	if args == nil {
		args = &ListInstanceSnapshotArgs{}
	}
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListInstanceSnapshotResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

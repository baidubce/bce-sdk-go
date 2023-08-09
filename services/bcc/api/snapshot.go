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

// snapshot.go - the snapshot APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"encoding/json"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateSnapshot - create a snapshot for specified volume
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to create snapshot
//
// RETURNS:
//   - *CreateSnapshotResult: result of the snapshot id newly created
//   - error: nil if success otherwise the specific error
func CreateSnapshot(cli bce.Client, args *CreateSnapshotArgs) (*CreateSnapshotResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSnapshotUri())
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

	jsonBody := &CreateSnapshotResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// ListSnapshot - list all snapshots with the specified parameters
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryArgs: arguments to list snapshots
//
// RETURNS:
//   - *ListSnapshotResult: result of the snapshot list
//   - error: nil if success otherwise the specific error
func ListSnapshot(cli bce.Client, queryArgs *ListSnapshotArgs) (*ListSnapshotResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSnapshotUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.Marker) != 0 {
			req.SetParam("marker", queryArgs.Marker)
		}
		if queryArgs.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(queryArgs.MaxKeys))
		}
	}

	if queryArgs == nil || queryArgs.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListSnapshotResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// ListSnapshotChain - list all snapshot chains with the specified parameters
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - queryArgs: arguments to list snapshot chains
//
// RETURNS:
//   - *ListSnapshotChainResult: result of the snapshot chain list
//   - error: nil if success otherwise the specific error
func ListSnapshotChain(cli bce.Client, queryArgs *ListSnapshotChainArgs) (*ListSnapshotChainResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSnapshotChainUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.OrderBy) != 0 {
			req.SetParam("orderBy", queryArgs.OrderBy)
		}
		if len(queryArgs.Order) != 0 {
			req.SetParam("order", queryArgs.Order)
		}
		if queryArgs.PageSize != 0 {
			req.SetParam("pageSize", strconv.Itoa(queryArgs.PageSize))
		}
		if queryArgs.PageNo != 0 {
			req.SetParam("pageNo", strconv.Itoa(queryArgs.PageNo))
		}
		if len(queryArgs.VolumeId) != 0 {
			req.SetParam("volumeId", queryArgs.VolumeId)
		}
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListSnapshotChainResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetSnapshotDetail - get details of the specified snapshot
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - snapshotId: id of the snapshot
//
// RETURNS:
//   - *GetSnapshotDetailResult: result of snapshot details
//   - error: nil if success otherwise the specific error
func GetSnapshotDetail(cli bce.Client, snapshotId string) (*GetSnapshotDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSnapshotUriWithId(snapshotId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetSnapshotDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// DeleteSnapshot - delete a snapshot
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - snapshotId: id of the snapshot to be deleted
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DeleteSnapshot(cli bce.Client, snapshotId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSnapshotUriWithId(snapshotId))
	req.SetMethod(http.DELETE)

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

func TagSnapshotChain(cli bce.Client, chainId string, args *TagVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getTagSnapshotChainUri(chainId))
	req.SetMethod(http.PUT)
	req.SetParam("bind", "")
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
	if resp.IsFail() {
		return resp.ServiceError()
	}

	return nil
}

func UntagSnapshotChain(cli bce.Client, chainId string, args *TagVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getUntagSnapshotChainUri(chainId))
	req.SetMethod(http.PUT)
	req.SetParam("unbind", "")
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
	if resp.IsFail() {
		return resp.ServiceError()
	}

	return nil
}

func CreateRemoteCopySnapshot(cli bce.Client, snapshotId string, args *RemoteCopySnapshotArgs) (
	*RemoteCopySnapshotResult, error) {

	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getRemoteCopySnapshotUri(snapshotId))
	req.SetMethod(http.PUT)

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

	jsonBody := &RemoteCopySnapshotResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

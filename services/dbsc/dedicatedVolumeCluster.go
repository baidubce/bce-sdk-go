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

package dbsc

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

func CreateVolumeCluster(cli bce.Client, args *CreateVolumeClusterArgs) (*CreateVolumeClusterResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getVolumeClusterUri() + REQUEST_CREATE_URI)
	req.SetMethod(http.POST)
	req.SetParam("uuidFlag", args.UuidFlag)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	resp := &bce.BceResponse{}
	err = cli.SendRequest(req, resp)
	if err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &CreateVolumeClusterResult{}
	err = resp.ParseJsonBody(jsonBody)
	if err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ListVolumeCluster(cli bce.Client, queryArgs *ListVolumeClusterArgs) (*ListVolumeClusterResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeClusterUri() + REQUEST_CREATE_URI)
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.ZoneName) != 0 {
			req.SetParam("zoneName", queryArgs.ZoneName)
		}
		if len(queryArgs.ClusterName) != 0 {
			req.SetParam("clusterName", queryArgs.ClusterName)
		}
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

	jsonBody := &ListVolumeClusterResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func GetVolumeClusterDetail(cli bce.Client, clusterId string) (*VolumeClusterDetail, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeClusterUri() + REQUEST_CREATE_URI + "/" + clusterId)
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &VolumeClusterDetail{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

func ResizeVolumeCluster(cli bce.Client, clusterId string, args *ResizeVolumeClusterArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeClusterUri() + REQUEST_CREATE_URI + "/" + clusterId)
	req.SetMethod(http.PUT)
	req.SetParam("resize", "")

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

	defer func() { _ = resp.Body().Close() }()
	return nil
}

func PurchaseReservedVolumeCluster(cli bce.Client, clusterId string, args *PurchaseReservedVolumeClusterArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeClusterUri() + REQUEST_CREATE_URI + "/" + clusterId)
	req.SetMethod(http.PUT)

	req.SetParam("purchaseReserved", "")

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

	defer func() { _ = resp.Body().Close() }()
	return nil
}

func AutoRenewVolumeCluster(cli bce.Client, args *AutoRenewVolumeClusterArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeClusterUri() + REQUEST_AUTO_RENEW_URI)
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
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { _ = resp.Body().Close() }()

	return nil
}

func CancelAutoRenewVolumeCluster(cli bce.Client, args *CancelAutoRenewVolumeClusterArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeClusterUri() + REQUEST_CANCEL_AUTO_RENEW_URI)
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
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { _ = resp.Body().Close() }()

	return nil
}

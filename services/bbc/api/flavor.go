/*
 * Copyright 2020 Baidu, Inc.
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

// flavor.go - the flavor APIs definition supported by the BBC service

// Package api defines all APIs supported by the BBC service of BCE.
package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListFlavors - list all available flavors
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - *ListFlavorsResult: the result of list all flavors
//     - error: nil if success otherwise the specific error
func ListFlavors(cli bce.Client) (*ListFlavorsResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getFlavorUri())
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListFlavorsResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetFlavorDetail - get details of the specified flavor
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - flavorId: the id of the flavor
// RETURNS:
//     - *api.GetFlavorDetailResult: the detail of the specified flavor
//     - error: nil if success otherwise the specific error
func GetFlavorDetail(cli bce.Client, flavorId string) (*GetFlavorDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getFlavorUriWithId(flavorId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetFlavorDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetFlavorRaid - get the RAID detail and disk size of the specified flavor
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - flavorId: the id of the flavor
// RETURNS:
//     - *api.GetFlavorRaidResult: the detail of the raid of the specified flavor
//     - error: nil if success otherwise the specific error
func GetFlavorRaid(cli bce.Client, flavorId string) (*GetFlavorRaidResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getFlavorRaidUriWithId(flavorId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetFlavorRaidResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

func getFlavorUri() string {
	return URI_PREFIX_V1 + REQUEST_FLAVOR_URI
}

func getFlavorUriWithId(id string) string {
	return URI_PREFIX_V1 + REQUEST_FLAVOR_URI + "/" + id
}

func getFlavorRaidUriWithId(id string) string {
	return URI_PREFIX_V1 + REQUEST_FLAVOR_RAID_URI + "/" + id
}

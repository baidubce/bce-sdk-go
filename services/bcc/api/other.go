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

// other.go - the other APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListSpec - get specification list information of the instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//
// RETURNS:
//   - *ListSpecResult: result of the specifications
//   - error: nil if success otherwise the specific error
func ListSpec(cli bce.Client) (*ListSpecResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getSpecUri())
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListSpecResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ListZone - get the available zone list in the current region
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//
// RETURNS:
//   - *ListZoneResult: result of the available zones
//   - error: nil if success otherwise the specific error
func ListZone(cli bce.Client) (*ListZoneResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getZoneUri())
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListZoneResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ListFlavorSpec - get the specified flavor list
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to list the specified flavor
//
// RETURNS:
//   - *ListFlavorSpecResult: result of the specified flavor list
//   - error: nil if success otherwise the specific error
func ListFlavorSpec(cli bce.Client, args *ListFlavorSpecArgs) (*ListFlavorSpecResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getFlavorSpecUri())
	req.SetMethod(http.GET)

	if args != nil {
		if len(args.ZoneName) != 0 {
			req.SetParam("zoneName", args.ZoneName)
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

	jsonBody := &ListFlavorSpecResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// GetPriceBySpec - get the price information of specified instance.
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - args: the arguments to get the price information of specified instance.
//
// RETURNS:
//   - *GetPriceBySpecResult: result of the specified instance's price information
//   - error: nil if success otherwise the specific error
func GetPriceBySpec(cli bce.Client, args *GetPriceBySpecArgs) (*GetPriceBySpecResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getPriceBySpecUri())
	req.SetMethod(http.POST)

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

	jsonBody := &GetPriceBySpecResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}

	return jsonBody, nil
}

// ListTypeZones - get the available zone list in the current region
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//
// RETURNS:
//   - *ListZoneResult: result of the available zones
//   - error: nil if success otherwise the specific error
func ListTypeZones(cli bce.Client, args *ListTypeZonesArgs) (*ListTypeZonesResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getInstanceTypeZoneUri())
	req.SetMethod(http.GET)
	if args != nil {
		if len(args.InstanceType) != 0 {
			req.SetParam("instanceType", args.InstanceType)
		}
		if len(args.ProductType) != 0 {
			req.SetParam("productType", args.ProductType)
		}
		if len(args.SpecId) != 0 {
			req.SetParam("specId", args.SpecId)
		}
		if len(args.Spec) != 0 {
			req.SetParam("spec", args.Spec)
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
	jsonBody := &ListTypeZonesResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// ListInstanceEni - get the eni list of the bcc instance
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - instanceId: the bcc instance id
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func ListInstanceEnis(cli bce.Client, instanceId string) (*ListInstanceEniResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getInstanceEniUri(instanceId))
	req.SetMethod(http.GET)
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListInstanceEniResult{}
	print(jsonBody)
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

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

// subnet.go - define the subnet APIs for BEC service

// Package bec defines the BEC services of BCE. The supported APIs are all defined in sub-package

package bec

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
	"strconv"
)

// CreateSubnet - create subnet with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a subnet
//
// RETURNS:
//   - *api.SubnetCommonResult: the result of create a subnet
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateSubnet(args *api.CreateSubnetRequest) (*api.SubnetCommonResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.SubnetCommonResult{}
	req := &api.PostHttpReq{Url: api.GetSubnetURI(), Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// UpdateSubnet - update subnet with the specific parameters
//
// PARAMS:
//   - args: the arguments to update a subnet
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateSubnet(subnetId string, args *api.UpdateSubnetRequest) error {
	if args == nil || subnetId == "" {
		return fmt.Errorf("please set argments")
	}

	req := &api.PostHttpReq{Url: api.GetSubnetURI() + "/" + subnetId, Body: args}
	err := api.Put(c, req)

	return err
}

// GetSubnetList - get subnet list with the specific parameters
//
// RETURNS:
//   - *api.LogicPageSubnetResult: the result of subnet list
//   - error: nil if ok otherwise the specific error
func (c *Client) GetSubnetList(args *api.ListRequest) (*api.LogicPageSubnetResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	params := make(map[string]string)
	if args.PageSize != 0 {
		params["pageSize"] = strconv.Itoa(args.PageSize)
	}
	if args.PageNo != 0 {
		params["pageNo"] = strconv.Itoa(args.PageNo)
	}
	if args.KeywordType != "" {
		params["keywordType"] = args.KeywordType
	}
	if args.Keyword != "" {
		params["keyword"] = args.Keyword
	}
	if args.VpcId != "" {
		params["vpcId"] = args.VpcId
	}

	result := &api.LogicPageSubnetResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetSubnetURI() + "/list").
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetSubnetDetail - get subnet detail with the specific parameters
//
// PARAMS:
//   - subnetId: subnet id
//
// RETURNS:
//   - *api.GetSubnetDetailResponse: the result of subnet detail
//   - error: nil if ok otherwise the specific error
func (c *Client) GetSubnetDetail(subnetId string) (*api.GetSubnetDetailResponse, error) {
	if subnetId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.GetSubnetDetailResponse{}
	req := &api.GetHttpReq{Url: api.GetSubnetURI() + "/" + subnetId, Result: result}
	err := api.Get(c, req)
	return result, err
}

// DeleteSubnet - delete subnet with the specific parameters
//
// PARAMS:
//   - subnetId: subnet id
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteSubnet(subnetId string) error {

	if subnetId == "" {
		return fmt.Errorf("please set argments")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(api.GetSubnetURI() + "/" + subnetId).
		Do()

	return err
}

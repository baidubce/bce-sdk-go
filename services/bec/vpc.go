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

// vpc.go - define the vpc APIs for BEC service

// Package bec defines the BEC services of BCE. The supported APIs are all defined in sub-package

package bec

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
	"strconv"
)

// CreateVpc - create a vpc with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a vpc
//
// RETURNS:
//   - *api.VpcCommonResult: the result of create a vpc
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateVpc(args *api.CreateVpcRequest) (*api.VpcCommonResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VpcCommonResult{}
	req := &api.PostHttpReq{Url: api.GetVpcURI(), Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// UpdateVpc - update vpc with the specific parameters
//
// PARAMS:
//   - args: the arguments to update a vpc
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateVpc(vpcId string, args *api.UpdateVpcRequest) (*api.VpcCommonResult, error) {
	if args == nil || vpcId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VpcCommonResult{}
	req := &api.PostHttpReq{Url: api.GetVpcURI() + "/" + vpcId, Body: args, Result: result}
	err := api.Put(c, req)

	return result, err
}

// GetVpcList - get vpc list with the specific parameters
// RETURNS:
//   - *api.LogicPageVpcResult: the result of vpc list
//   - error: nil if ok otherwise the specific error
func (c *Client) GetVpcList(args *api.ListRequest) (*api.LogicPageVpcResult, error) {
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
	if args.RegionId != "" {
		params["regionId"] = args.RegionId
	}

	result := &api.LogicPageVpcResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetVpcURI() + "/list").
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetVpcDetail - get vpc detail with the specific parameters
//
// PARAMS:
//   - vpcId: vpc id
//
// RETURNS:
//   - *api.GetVpcDetailResponse: the result of vpc detail
//   - error: nil if ok otherwise the specific error
func (c *Client) GetVpcDetail(vpcId string) (*api.GetVpcDetailResponse, error) {
	if vpcId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.GetVpcDetailResponse{}
	req := &api.GetHttpReq{Url: api.GetVpcURI() + "/" + vpcId, Result: result}
	err := api.Get(c, req)
	return result, err
}

// DeleteVpc - delete vpc with the specific parameters
//
// PARAMS:
//   - vpcId: vpc id
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteVpc(vpcId string) (*api.VpcCommonResult, error) {

	if vpcId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VpcCommonResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(api.GetVpcURI() + "/" + vpcId).
		WithResult(result).
		Do()

	return result, err
}

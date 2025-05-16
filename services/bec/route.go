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

// route.go - define the route APIs for BEC service

// Package bec defines the BEC services of BCE. The supported APIs are all defined in sub-package

package bec

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
	"strconv"
)

// UpdateRouteTable - update route table with the specific parameters
//
// PARAMS:
//   - tableId: route table id
//   - args: the arguments to update a route table
//
// RETURNS:
//   - *api.UpdateRouteTableResult: the result of updating route table
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateRouteTable(tableId string, args *api.UpdateRouteTableRequest) (*api.UpdateRouteTableResult, error) {
	if args == nil || tableId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	result := &api.UpdateRouteTableResult{}
	req := &api.PostHttpReq{Url: api.GetRouteURI() + "/update" + "/" + tableId, Body: args, Result: result}
	err := api.Put(c, req)

	return result, err
}

// GetRouteTableList - get route table list with the specific parameters
// RETURNS:
//   - *api.LogicPageRouteTableResult: the result of route table list
//   - error: nil if ok otherwise the specific error
func (c *Client) GetRouteTableList(args *api.ListRequest) (*api.LogicPageRouteTableResult, error) {
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
	if args.VpcId != "" {
		params["vpcId"] = args.VpcId
	}

	result := &api.LogicPageRouteTableResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetRouteURI() + "/list").
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetRouteTableDetail - get route table detail with the specific parameters
//
// PARAMS:
//   - tableId: route table id
//
// RETURNS:
//   - *api.GetRouteTableDetailResult: the result of route table detail
//   - error: nil if ok otherwise the specific error
func (c *Client) GetRouteTableDetail(tableId string) (*api.GetRouteTableDetailResult, error) {
	if tableId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.GetRouteTableDetailResult{}
	req := &api.GetHttpReq{Url: api.GetRouteURI() + "/detail" + "/" + tableId, Result: result}
	err := api.Get(c, req)
	return result, err
}

// CreateRouteRule - create route rule with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a route rule
//
// RETURNS:
//   - *api.CreateRouteRuleResult: the result of create route rule
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateRouteRule(args *api.CreateRouteRuleRequest) (*api.CreateRouteRuleResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.CreateRouteRuleResult{}
	req := &api.PostHttpReq{Url: api.GetRouteURI() + "/rule", Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// GetRouteRuleList - get route rule list with the specific parameters
// RETURNS:
//   - *LogicPageRouteRuleResult: the result of route rule list
//   - error: nil if ok otherwise the specific error
func (c *Client) GetRouteRuleList(tableId string, args *api.ListRequest) (*api.LogicPageRouteRuleResult, error) {
	if args == nil || tableId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	params := make(map[string]string)
	if args.PageSize != 0 {
		params["pageSize"] = strconv.Itoa(args.PageSize)
	}
	if args.PageNo != 0 {
		params["pageNo"] = strconv.Itoa(args.PageNo)
	}

	result := &api.LogicPageRouteRuleResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetRouteURI() + "/rule/list" + "/" + tableId).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// DeleteRouteRule - delete route rule with the specific parameters
//
// PARAMS:
//   - ruleId: rule id
//
// RETURNS:
//   - *api.VpcCommonResult: the result of deleting route rule result
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteRouteRule(ruleId string) (*api.RouteCommonResult, error) {

	if ruleId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.RouteCommonResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(api.GetRouteURI() + "/rule/delete" + "/" + ruleId).
		WithResult(result).
		Do()

	return result, err
}

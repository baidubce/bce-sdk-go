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

// client.go - define the client for BEC service

// Package bec defines the BEC services of BCE. The supported APIs are all defined in sub-package

package bec

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// CreateDeploySet - create deploy set  with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a deploy set
// RETURNS:
//     - *CreateDeploySetResponseArgs: the result of create deploy set
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateDeploySet(args *api.CreateDeploySetArgs) (*api.CreateDeploySetResponseArgs, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.CreateDeploySetResponseArgs{}
	req := &api.PostHttpReq{Url: api.GetDeploySetURI() + "/create", Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// UpdateVmInstanceDeploySet - update vm instance deploy set with the specific parameters
//
// PARAMS:
//     - args: the arguments to update a vm instance deploy set
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateVmInstanceDeploySet(args *api.UpdateVmDeploySetArgs) error {
	if args == nil {
		return fmt.Errorf("please set argments")
	}
	req := &api.PostHttpReq{Url: api.GetDeploySetURI() + "/updateRelation", Body: args}
	err := api.Post(c, req)
	return err
}

// DeleteVmInstanceFromDeploySet - remove vm instances from deploy set with the specific parameters
//
// PARAMS:
//     - args: the arguments to remove  vm instances from  deploy set
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteVmInstanceFromDeploySet(args *api.DeleteVmDeploySetArgs) error {
	if args == nil {
		return fmt.Errorf("please set argments")
	}
	req := &api.PostHttpReq{Url: api.GetDeploySetURI() + "/delRelation", Body: args}
	err := api.Post(c, req)
	return err
}

// UpdateDeploySet - update deploy set with the specific parameters
//
// PARAMS:
//     - args: the arguments to update deploy set
// RETURNS:
//     - *ListVmServiceResult: the result of get vm services
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateDeploySet(deploySetId string, args *api.CreateDeploySetArgs) error {
	if args == nil || deploySetId == "" {
		return fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	params["modifyAttribute"] = ""
	req := &api.PostHttpReq{Url: api.GetDeploySetURI() + "/" + deploySetId, Body: args, Params: params}
	err := api.Put(c, req)
	return err
}

// GetDeploySetList - get deploy set list with the specific parameters
// RETURNS:
//     - *LogicPageDeploySetResult: the result of deploy set list
//     - error: nil if ok otherwise the specific error
func (c *Client) GetDeploySetList(args *api.ListRequest) (*api.LogicPageDeploySetResult, error) {
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

	result := &api.LogicPageDeploySetResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetDeploySetURI() + "/list").
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetDeploySetDetail - get vm service detail with the specific parameters
//
// PARAMS:
//     - deploySetId: deploy set id
// RETURNS:
//     - *DeploySetDetails: the result of  deploy set detail
//     - error: nil if ok otherwise the specific error
func (c *Client) GetDeploySetDetail(deploySetId string) (*api.DeploySetDetails, error) {
	if deploySetId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.DeploySetDetails{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetDeploySetURI() + "/" + deploySetId).
		WithResult(result).
		Do()
	return result, err
}

// DeleteDeploySet - delete deploy set  with the specific parameters
//
// PARAMS:
//     - deploySetId: deploy set id
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteDeploySet(deploySetId string) error {
	if deploySetId == "" {
		return fmt.Errorf("please set argments")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(api.GetDeploySetURI() + "/" + deploySetId).
		Do()

	return err
}

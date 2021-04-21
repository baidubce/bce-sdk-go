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

// client.go - define the client for BOS service

// Package bec defines the BEC services of BCE. The supported APIs are all defined in sub-package

package bec

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// CreateVmImage - create a vm image
//
// PARAMS:
//     - args: the create vm image args
// RETURNS:
//     - *api.CreateVmImageResult: the result image
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateVmImage(args *api.CreateVmImageArgs) (*api.CreateVmImageResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.CreateVmImageResult{}
	req := &api.PostHttpReq{Url: api.GetVmImageURI(), Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// DeleteVmImage - delete a vm image
//
// PARAMS:
//     - args: the delete vm image args, spec vmId list
// RETURNS:
//     - *api.VmImageOperateResult: the result image delete
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteVmImage(args []string) (*api.VmImageOperateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmImageOperateResult{}
	params := make(map[string]string)
	params["imageIdList"] = strings.Join(args, ",")
	req := &api.PostHttpReq{Url: api.GetVmImageURI(), Result: result, Params: params}
	err := api.Delete(c, req)

	return result, err
}

// UpdateVmImage - update a vm image
//
// PARAMS:
//     - imageId: image id
//     - args: the update vm image args
// RETURNS:
//     - *api.VmImageOperateResult: the result image update
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateVmImage(imageId string, args *api.UpdateVmImageArgs) (*api.VmImageOperateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmImageOperateResult{}
	req := &api.PostHttpReq{Url: api.GetVmImageURI() + "/" + imageId, Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

// ListVmImage - image list
//
// PARAMS:
//     - args: the vm image list args
// RETURNS:
//     - *api.ListVmImageResult: the list of vm images
//     - error: nil if ok otherwise the specific error
func (c *Client) ListVmImage(args *api.ListVmImageArgs) (*api.ListVmImageResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	params := make(map[string]string)
	if args.PageNo != 0 {
		params["PageNo"] = strconv.Itoa(args.PageNo)
	}
	if args.PageSize != 0 {
		params["PageSize"] = strconv.Itoa(args.PageSize)
	}
	if args.Order != "" {
		params["Order"] = args.Order
	}
	if args.OrderBy != "" {
		params["OrderBy"] = args.OrderBy
	}
	if args.Type != "" {
		params["Type"] = args.Type
	}
	if args.KeywordType != "" {
		params["keywordType"] = args.KeywordType
	}
	if args.Keyword != "" {
		params["keyword"] = args.Keyword
	}
	if args.Status != "" {
		params["status"] = args.Status
	}
	if args.Region != "" {
		params["region"] = args.Region
	}
	if args.OsName != "" {
		params["osName"] = args.OsName
	}
	if args.ServiceId != "" {
		params["serviceId"] = args.ServiceId
	}

	result := &api.ListVmImageResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetVmImageURI()).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

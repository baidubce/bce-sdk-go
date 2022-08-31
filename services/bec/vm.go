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

// CreateVmService - create vm service with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a vm service
// RETURNS:
//     - *CreateVmServiceResult: the result of create vm service
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateVmService(args *api.CreateVmServiceArgs) (*api.CreateVmServiceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.CreateVmServiceResult{}
	req := &api.PostHttpReq{Url: api.GetVmURI(), Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// UpdateVmService - update vm service with the specific parameters
//
// PARAMS:
//     - args: the arguments to update a vm service
// RETURNS:
//     - *UpdateVmServiceResult: the result of update vm service
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateVmService(serviceId string, args *api.UpdateVmServiceArgs) (*api.UpdateVmServiceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.UpdateVmServiceResult{}
	req := &api.PostHttpReq{Url: api.GetVmURI() + "/" + serviceId, Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

// GetVmServiceList - get vm services with the specific parameters
//
// PARAMS:
//     - args: the arguments to get vm services
// RETURNS:
//     - *ListVmServiceResult: the result of get vm services
//     - error: nil if ok otherwise the specific error
func (c *Client) GetVmServiceList(args *api.ListVmServiceArgs) (*api.ListVmServiceResult, error) {
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
	if args.OrderBy != "" {
		params["orderBy"] = args.OrderBy
	}
	if args.Order != "" {
		params["order"] = args.Order
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

	result := &api.ListVmServiceResult{}
	req := &api.GetHttpReq{Url: api.GetVmURI(), Result: result, Params: params}
	err := api.Get(c, req)

	return result, err
}

// GetVmServiceDetail - get vm service detail with the specific parameters
//
// PARAMS:
//     - serviceId: vm service id
// RETURNS:
//     - *VmServiceDetailsVo: the result of vm service detail
//     - error: nil if ok otherwise the specific error
func (c *Client) GetVmServiceDetail(serviceId string) (*api.VmServiceDetailsVo, error) {
	if serviceId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmServiceDetailsVo{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetVmURI() + "/" + serviceId).
		WithResult(result).
		Do()
	return result, err
}

// GetVmServiceMetrics - get vm service metrics with the specific parameters
//
// PARAMS:
//     - serviceId: service id
//     - serviceProviderStr: service Provider
//     - offsetInSeconds:  offset Seconds
//     - metricsType: metrics Type
// RETURNS:
//     - *ServiceMetricsResult: the result of get vm service metrics
//     - error: nil if ok otherwise the specific error
func (c *Client) GetVmServiceMetrics(serviceId, serviceProviderStr string, start, end, stepInMin int, metricsType api.MetricsType) (*api.ServiceMetricsResult, error) {

	params := make(map[string]string)
	params["serviceId"] = serviceId
	if serviceProviderStr != "" {
		params["serviceProvider"] = serviceProviderStr
	}
	if metricsType != "" {
		params["metricsType"] = string(metricsType)
	}
	if start != 0 {
		params["start"] = strconv.Itoa(start)
	}
	if end != 0 {
		params["end"] = strconv.Itoa(end)
	}
	if stepInMin != 0 {
		params["stepInMin"] = strconv.Itoa(stepInMin)
	}
	result := &api.ServiceMetricsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetVmServiceMonitorURI() + "/" + serviceId).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// VmServiceAction - operate vm service with the specific parameters
//
// PARAMS:
//     - serviceId: service id
//     - action: operation action
// RETURNS:
//     - *VmServiceActionResult: the result of operate vm service
//     - error: nil if ok otherwise the specific error
func (c *Client) VmServiceAction(serviceId string, action api.VmServiceAction) (*api.VmServiceActionResult, error) {
	if serviceId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmServiceActionResult{}
	req := &api.PostHttpReq{Url: api.GetVmServiceActionURI(serviceId, string(action)), Result: result, Body: nil}
	err := api.Put(c, req)

	return result, err
}

// DeleteVmService - delete a vm service with the specific parameters
//
// PARAMS:
//     - serviceId: service id
// RETURNS:
//     - *VmServiceActionResult: the result of delete vm service
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteVmService(serviceId string) (*api.VmServiceActionResult, error) {
	if serviceId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmServiceActionResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(api.GetVmURI() + "/" + serviceId).
		WithResult(result).
		Do()

	return result, err
}

// BatchDeleteVmService - batch delete vm service with the specific parameters
//
// PARAMS:
//     - serviceIds: service id list
// RETURNS:
//     - *VmServiceBatchActionResult: the result of batch delete service id list
//     - error: nil if ok otherwise the specific error
func (c *Client) BatchDeleteVmService(serviceIds *[]string) (*api.VmServiceBatchActionResult, error) {
	if serviceIds == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmServiceBatchActionResult{}
	req := &api.PostHttpReq{Url: api.GetVmURI() + "/batch/delete", Result: result, Body: serviceIds}
	err := api.Post(c, req)

	return result, err
}

// BatchOperateVmService - batch operate vm service with the specific parameters
//
// PARAMS:
//     - args: the arguments to batch operate vm service
// RETURNS:
//     - *VmServiceBatchActionResult: the result of batch operate vm service
//     - error: nil if ok otherwise the specific error
func (c *Client) BatchOperateVmService(args *api.VmServiceBatchActionArgs) (*api.VmServiceBatchActionResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	if args.Action != "start" && args.Action != "stop" {
		return nil, fmt.Errorf("action is start|stop, please check")
	}

	result := &api.VmServiceBatchActionResult{}
	req := &api.PostHttpReq{Url: api.GetVmURI() + "/batch/operate", Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

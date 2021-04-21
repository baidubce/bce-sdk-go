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

// CreateService - create a container service with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a container service
// RETURNS:
//     - *CreateClusterResult: the result of create a container service
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateService(args *api.CreateServiceArgs) (*api.CreateServiceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.CreateServiceResult{}
	req := &api.PostHttpReq{Url: api.GetServiceURI() + "/create", Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// ListService - list container service with the specific parameters
//
// PARAMS:
//     - pageNo: page No
//     - pageSize: page Size
//     - keywordType: keyword Type
//     - keyword: keyword
//     - order: order
//     - orderBy: orderBy
//     - status: status
// RETURNS:
//     - *ListServiceResult: the result of list container service
//     - error: nil if ok otherwise the specific error
func (c *Client) ListService(pageNo, pageSize int, keywordType, keyword, order, orderBy, status string) (*api.ListServiceResult, error) {

	params := make(map[string]string)
	if pageNo != 0 {
		params["pageNo"] = strconv.Itoa(pageNo)
	}
	if pageSize != 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if keyword != "" {
		params["keyword"] = keyword
	}
	if keywordType != "" {
		params["keywordType"] = keywordType
	}
	if status != "" {
		params["status"] = status
	}
	if order != "" {
		params["order"] = order
	}
	if orderBy != "" {
		params["orderBy"] = orderBy
	}

	result := &api.ListServiceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetServiceURI()).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetService - get container service with the specific parameters
//
// PARAMS:
//     - serviceId: the service id
// RETURNS:
//     - *ServiceBriefVo: the result of get container service
//     - error: nil if ok otherwise the specific error
func (c *Client) GetService(serviceId string) (*api.ServiceDetailsVo, error) {
	result := &api.ServiceDetailsVo{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetServiceDetailURI(serviceId)).
		WithResult(result).
		Do()

	return result, err
}

// ServiceAction - operate service with the specific parameters
//
// PARAMS:
//     - serviceId: the service id
//     - action: operate action
// RETURNS:
//     - *ServiceActionResult: the result of operate service
//     - error: nil if ok otherwise the specific error
func (c *Client) ServiceAction(serviceId string, action api.ServiceAction) (*api.ServiceActionResult, error) {
	result := &api.ServiceActionResult{}
	req := &api.PostHttpReq{Url: api.GetStartServiceURI(serviceId, string(action)), Result: result, Body: nil}
	err := api.Put(c, req)

	return result, err
}

// UpdateService - update service with the specific parameters
//
// PARAMS:
//     - serviceId: the service id
//     - args: the arguments to update service
// RETURNS:
//     - *UpdateServiceResult: the result of  update service
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateService(serviceId string, args *api.UpdateServiceArgs) (*api.UpdateServiceResult, error) {
	if args == nil {
		args = &api.UpdateServiceArgs{}
	}

	result := &api.UpdateServiceResult{}
	req := &api.PostHttpReq{Url: api.GetUpdateServiceURI(serviceId), Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

// DeleteService - delete service with the specific parameters
//
// PARAMS:
//     - serviceId: the service id
// RETURNS:
//     - *ServiceActionResult: the result of delete service
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteService(serviceId string) (*api.ServiceActionResult, error) {
	result := &api.ServiceActionResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(api.GetDeleteServiceURI(serviceId)).
		WithResult(result).
		Do()

	return result, err
}

// GetServiceMetrics - get service metrics with the specific parameters
//
// PARAMS:
//     - serviceId: the service id
//     - metricsType: metrics Type
//     - serviceProviderStr: service Provider
//     - offsetInSeconds: offset Seconds
// RETURNS:
//     - *ServiceMetricsResult: the result of get service metrics
//     - error: nil if ok otherwise the specific error
func (c *Client) GetServiceMetrics(serviceId string, metricsType api.MetricsType, serviceProviderStr api.ServiceProvider, offsetInSeconds int) (*api.ServiceMetricsResult, error) {
	params := make(map[string]string)
	if serviceProviderStr != "" {
		params["serviceProvider"] = string(serviceProviderStr)
	}
	params["offsetInSeconds"] = strconv.Itoa(offsetInSeconds)

	result := &api.ServiceMetricsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetServiceMetricsURI(serviceId, string(metricsType))).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// ServiceBatchOperate - batch operate service with the specific parameters
//
// PARAMS:
//     - args: the arguments to batch operate service
// RETURNS:
//     - *ServiceBatchOperateResult: the result of batch operate service
//     - error: nil if ok otherwise the specific error
func (c *Client) ServiceBatchOperate(args *api.ServiceBatchOperateArgs) (*api.ServiceBatchOperateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	if args.Action != "start" && args.Action != "stop" {
		return nil, fmt.Errorf("action is start|stop, please check")
	}

	result := &api.ServiceBatchOperateResult{}
	req := &api.PostHttpReq{Url: api.GetBachServiceOperateURI(), Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

// ServiceBatchDelete - batch delete service with the specific parameters
//
// PARAMS:
//     - args: the arguments to batch delete service
// RETURNS:
//     - *ServiceBatchOperateResult: the result of batch delete service
//     - error: nil if ok otherwise the specific error
func (c *Client) ServiceBatchDelete(args *[]string) (*api.ServiceBatchOperateResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.ServiceBatchOperateResult{}
	req := &api.PostHttpReq{Url: api.GetBachServiceDeleteURI(), Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

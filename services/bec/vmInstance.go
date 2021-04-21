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

// GetVmInstanceList - get vm list with the specific parameters
//
// PARAMS:
//     - args: the arguments to get vm list
// RETURNS:
//     - *LogicPageVmInstanceResult: the result of get vm list
//     - error: nil if ok otherwise the specific error
func (c *Client) GetVmInstanceList(args *api.ListRequest) (*api.LogicPageVmInstanceResult, error) {
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

	result := &api.LogicPageVmInstanceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetVmInstanceURI()).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetNodeVmInstanceList - get node vm instance list with the specific parameters
//
// PARAMS:
//     - args: the arguments to get node vm instance list
//     - region: region
//     - serviceProvider: service provider
//     - city: city
// RETURNS:
//     - *GetNodeVmInstanceListResult: the result of get node vm instance list
//     - error: nil if ok otherwise the specific error
func (c *Client) GetNodeVmInstanceList(args *api.ListRequest, region, serviceProvider, city string) (*api.GetNodeVmInstanceListResult, error) {
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

	result := &api.GetNodeVmInstanceListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetVmInstanceURI() + "/regions/" + region + "/sps/" + serviceProvider + "/cities/" + city).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetVirtualMachine - get vm with the specific parameters
//
// PARAMS:
//     - vmID: vm id
// RETURNS:
//     - *VmInstanceDetailsVo: the result of get vm
//     - error: nil if ok otherwise the specific error
func (c *Client) GetVirtualMachine(vmID string) (*api.VmInstanceDetailsVo, error) {
	if vmID == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmInstanceDetailsVo{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetVmInstanceURI() + "/" + vmID).
		WithResult(result).
		Do()

	return result, err
}

// DeleteVmInstance - delete vm instance with the specific parameters
//
// PARAMS:
//     - vmID: vm id
// RETURNS:
//     - *ActionInfoVo: the result of delete vm instance
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteVmInstance(vmID string) (*api.ActionInfoVo, error) {
	if vmID == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.ActionInfoVo{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(api.GetVmInstanceURI() + "/" + vmID).
		WithResult(result).
		Do()

	return result, err
}

// UpdateVmDeployment - update vm with the specific parameters
//
// PARAMS:
//     - vmID: vm id
//     - args: the arguments to update vm
// RETURNS:
//     - *UpdateVmDeploymentResult: the result of update vm
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateVmDeployment(vmID string, args *api.UpdateVmDeploymentArgs) (*api.UpdateVmDeploymentResult, error) {
	if vmID == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.UpdateVmDeploymentResult{}
	req := &api.PostHttpReq{Url: api.GetVmInstanceURI() + "/" + vmID, Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

// ReinstallVmInstance - reinstall vm instance with the specific parameters
//
// PARAMS:
//     - vmID: vm id
//     - args: the arguments to reinstall vm instance
// RETURNS:
//     - *ReinstallVmInstanceResult: the result of reinstall vm instance
//     - error: nil if ok otherwise the specific error
func (c *Client) ReinstallVmInstance(vmID string, args *api.ReinstallVmInstanceArg) (*api.ReinstallVmInstanceResult, error) {
	if vmID == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.ReinstallVmInstanceResult{}
	req := &api.PostHttpReq{Url: api.GetVmInstanceURI() + "/" + vmID + "/system/reinstall", Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

// OperateVmDeployment - operate vm with the specific parameters
//
// PARAMS:
//     - vmID: vm id
//     - action: the arguments to operate vm
// RETURNS:
//     - *OperateVmDeploymentResult: the result of operate vm
//     - error: nil if ok otherwise the specific error
func (c *Client) OperateVmDeployment(vmID string, action api.VmInstanceBatchOperateAction) (*api.OperateVmDeploymentResult, error) {
	if vmID == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.OperateVmDeploymentResult{}
	req := &api.PostHttpReq{Url: api.GetVmInstanceURI() + "/" + vmID + "/" + string(action), Result: result, Body: nil}
	err := api.Put(c, req)

	return result, err
}

// GetVmInstanceMetrics - get vm metrics with the specific parameters
//
// PARAMS:
//     - vmId: vm id
//     - offsetInSeconds: offset Seconds
//     - serviceProvider: service provider
//     - metricsType: metrics Type
// RETURNS:
//     - *ServiceMetricsResult: the result of get vm metrics
//     - error: nil if ok otherwise the specific error
func (c *Client) GetVmInstanceMetrics(vmID string, serviceProvider api.ServiceProvider, offsetInSeconds int, metricsType api.MetricsType) (*api.ServiceMetricsResult, error) {
	if vmID == "" {
		return nil, fmt.Errorf("please set argments")
	}

	params := make(map[string]string)
	if serviceProvider != "" {
		params["serviceProvider"] = string(serviceProvider)
	}
	params["offsetInSeconds"] = strconv.Itoa(offsetInSeconds)

	result := &api.ServiceMetricsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetVmInstanceURI() + "/" + vmID + "/metrics/" + string(metricsType)).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetVmConfig - get vm config with the specific parameters
//
// PARAMS:
//     - vmID: vm id
// RETURNS:
//     - *VmConfigResult: the result of get vm config
//     - error: nil if ok otherwise the specific error
func (c *Client) GetVmConfig(vmID string) (*api.VmConfigResult, error) {
	if vmID == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmConfigResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetVmInstanceURI() + "/" + vmID + "/config").
		WithResult(result).
		Do()

	return result, err
}

// CreateVmPrivateIp - create vm private ip with the specific parameters
//
// PARAMS:
//     - vmID: vm id
//     - args: the args to create vm private ip
// RETURNS:
//     - *VmConfigResult: the result of create vm private ip
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateVmPrivateIp(vmID string, args *api.CreateVmPrivateIpForm) (*api.VmPrivateIpResult, error) {
	if vmID == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmPrivateIpResult{}
	req := &api.PostHttpReq{Url: api.GetVmInstanceURI() + "/" + vmID + "/privateIp", Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// DeleteVmPrivateIp - delete vm private ip with the specific parameters
//
// PARAMS:
//     - vmID: vm id
//     - args: the args to delete vm private ip
// RETURNS:
//     - *VmPrivateIpResult: the result of delete vm private ip
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteVmPrivateIp(vmID string, args *api.DeleteVmPrivateIpForm) (*api.VmPrivateIpResult, error) {
	if vmID == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.VmPrivateIpResult{}
	req := &api.PostHttpReq{Url: api.GetVmInstanceURI() + "/" + vmID + "/privateIp/release", Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

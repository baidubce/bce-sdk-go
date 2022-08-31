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
func (c *Client) GetServiceMetrics(serviceId string, metricsType api.MetricsType, serviceProviderStr api.ServiceProvider, start, end, stepInMin int) (*api.ServiceMetricsResult, error) {
	params := make(map[string]string)
	if serviceProviderStr != "" {
		params["serviceProvider"] = string(serviceProviderStr)
	}
	if metricsType != "" {
		params["metricsType"] = string(metricsType)
	}
	if serviceId != "" {
		params["serviceId"] = string(serviceId)
	}
	if stepInMin != 0 {
		params["stepInMin"] = strconv.Itoa(stepInMin)
	}
	if start != 0 {
		params["start"] = strconv.Itoa(start)
	}
	if end != 0 {
		params["end"] = strconv.Itoa(end)
	}

	result := &api.ServiceMetricsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetServiceMetricsURI(serviceId)).
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

// GetPodDeployment - get pod deployment with the specific parameters
//
// PARAMS:
//     - deploymentId: the deploymentId id
// RETURNS:
//     - *DeploymentResourceBriefVo: the result of get pod deployment
//     - error: nil if ok otherwise the specific error
func (c *Client) GetPodDeployment(deploymentId string) (*api.DeploymentResourceBriefVo, error) {
	result := &api.DeploymentResourceBriefVo{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetDeploymentDetailURI(deploymentId)).
		WithResult(result).
		Do()

	return result, err
}

// GetPodDeploymentMetrics - get Pod Deployment metrics with the specific parameters
//
// PARAMS:
//     - deploymentId: the pod deployment id
//     - metricsType: metrics Type
//     - serviceProviderStr: service Provider
//     - offsetInSeconds: offset Seconds
// RETURNS:
//     - *ServiceMetricsResult: the result of get Pod Deployment metrics
//     - error: nil if ok otherwise the specific error
func (c *Client) GetPodDeploymentMetrics(deploymentId string, metricsType api.MetricsType, serviceProviderStr api.ServiceProvider, start, end, stepInMin int) (*api.ServiceMetricsResult, error) {
	params := make(map[string]string)
	if serviceProviderStr != "" {
		params["serviceProvider"] = string(serviceProviderStr)
	}
	if metricsType != "" {
		params["metricsType"] = string(metricsType)
	}
	if deploymentId != "" {
		params["deploymentId"] = deploymentId
	}
	if stepInMin != 0 {
		params["stepInMin"] = strconv.Itoa(stepInMin)
	}
	if start != 0 {
		params["start"] = strconv.Itoa(start)
	}
	if end != 0 {
		params["end"] = strconv.Itoa(end)
	}

	result := &api.ServiceMetricsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetDeploymentMetricsURI(deploymentId)).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// UpdatePodDeploymentReplicas - update pod deployment replicas with the specific parameters
//
// PARAMS:
//     - deploymentId: the deploymentId id
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdatePodDeploymentReplicas(deploymentId string, args *api.UpdateDeploymentReplicasRequest) (*api.ActionInfoVo, error) {
	if args == nil || deploymentId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	res := &api.ActionInfoVo{}
	req := &api.PostHttpReq{Url: api.DEPLOYMENT_URL + "/" + deploymentId, Result: res, Body: args}
	err := api.Put(c, req)
	return res, err
}

// DeletePodDeployment - delete pod deployment with the specific parameters
//
// PARAMS:
//     - deploymentIDs: the deployment id array
// RETURNS:
//     - *ServiceActionResult: the result of delete service
//     - error: nil if ok otherwise the specific error
func (c *Client) DeletePodDeployment(args *[]string) (*api.DeleteDeploymentActionInfoVo, error) {
	result := &api.DeleteDeploymentActionInfoVo{}
	req := &api.PostHttpReq{Url: api.DEPLOYMENT_URL, Body: args, Result: result}
	err := api.Delete(c, req)
	return result, err
}

// GetPodList - list pod with the specific parameters
//
// PARAMS:
//     - pageNo: page No
//     - pageSize: page Size
//     - keyword: keyword
//     - order: order
//     - orderBy: orderBy
// RETURNS:
//     - *ListPodResult: the result of list pod
//     - error: nil if ok otherwise the specific error
func (c *Client) GetPodList(pageNo, pageSize int, keyword, order, orderBy, serviceId, deploymentId string) (*api.ListPodResult, error) {

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
	if order != "" {
		params["order"] = order
	}
	if orderBy != "" {
		params["orderBy"] = orderBy
	}
	if serviceId != "" {
		params["serviceId"] = serviceId
	}
	if deploymentId != "" {
		params["deploymentId"] = deploymentId
	}

	result := &api.ListPodResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.REQUEST_POD_URL).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetPodMetrics - get Pod metrics with the specific parameters
//
// PARAMS:
//     - deploymentId: the pod deployment id
//     - metricsType: metrics Type
//     - serviceProviderStr: service Provider
//     - offsetInSeconds: offset Seconds
// RETURNS:
//     - *ServiceMetricsResult: the result of get Pod Deployment metrics
//     - error: nil if ok otherwise the specific error
func (c *Client) GetPodMetrics(podId string, metricsType api.MetricsType, serviceProviderStr api.ServiceProvider, start, end, stepInMin int) (*api.ServiceMetricsResult, error) {
	params := make(map[string]string)
	if serviceProviderStr != "" {
		params["serviceProvider"] = string(serviceProviderStr)
	}
	if metricsType != "" {
		params["metricsType"] = string(metricsType)
	}
	if podId != "" {
		params["podId"] = podId
	}
	if stepInMin != 0 {
		params["stepInMin"] = strconv.Itoa(stepInMin)
	}
	if start != 0 {
		params["start"] = strconv.Itoa(start)
	}
	if end != 0 {
		params["end"] = strconv.Itoa(end)
	}

	result := &api.ServiceMetricsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetPodMetricsURI(podId)).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetPodDetail - get pod detail with the specific parameters
//
// PARAMS:
//     - podId: pod id
//RETURNS:
//     - *ListPodResult: the result of list pod
//     - error: nil if ok otherwise the specific error
func (c *Client) GetPodDetail(podId string) (*api.PodDetailVo, error) {

	if podId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	result := &api.PodDetailVo{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.REQUEST_POD_URL + "/" + podId).
		WithResult(result).
		Do()

	return result, err
}

// RestartPod - restart pod with the specific parameters
//
// PARAMS:
//     - podId: pod id
//RETURNS:
//     - *ListPodResult: the result of restart pod
//     - error: nil if ok otherwise the specific error
func (c *Client) RestartPod(podId string) error {

	if podId == "" {
		return fmt.Errorf("please set argments")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(api.REQUEST_POD_URL + "/" + podId + "/restart").
		Do()

	return err
}

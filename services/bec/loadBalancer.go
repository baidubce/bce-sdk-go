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

// CreateBlb - create lb
//
// PARAMS:
//     - args: the lb create args
// RETURNS:
//     - *api.CreateBlbResult: the create lb result
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateBlb(args *api.CreateBlbArgs) (*api.CreateBlbResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.CreateBlbResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI(), Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// DeleteBlb - delete lb
//
// PARAMS:
//     - blbId: lb id
// RETURNS:
//     - *api.DeleteBlbResult: delete lb result
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteBlb(blbId string) (*api.DeleteBlbResult, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.DeleteBlbResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(api.GetLoadBalancerURI() + "/" + blbId).
		WithResult(result).
		Do()

	return result, err
}

// GetBlbList - get lb list
//
// PARAMS:
//     - lbType: lb type
//     - order: list order
//     - orderBy: order by
//     - keyword: the key word
//     - keywordType: key word type
//     - status: lb status
//     - region: lb's region
//	   - pageNo: page NO
//	   - pageSize: page size
// RETURNS:
//     - *api.GetBlbListResult: the list of lb
//     - error: nil if ok otherwise the specific error
func (c *Client) GetBlbList(lbType, order, orderBy, keyword, keywordType, status, region string,
	pageNo, pageSize int) (*api.GetBlbListResult, error) {

	params := make(map[string]string)
	if order != "" {
		params["order"] = order
	}
	if orderBy != "" {
		params["orderBy"] = orderBy
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
	if region != "" {
		params["region"] = region
	}
	if pageSize != 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if pageNo != 0 {
		params["pageNo"] = strconv.Itoa(pageNo)
	}
	if lbType != "" {
		params["lbType"] = lbType
	}

	result := &api.GetBlbListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetLoadBalancerURI()).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetBlbDetail - lb detail
//
// PARAMS:
//     - blbId: lb id
// RETURNS:
//     - *api.BlbInstanceVo: lb info
//     - error: nil if ok otherwise the specific error
func (c *Client) GetBlbDetail(blbId string) (*api.BlbInstanceVo, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.BlbInstanceVo{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetLoadBalancerURI() + "/" + blbId).
		WithResult(result).
		Do()

	return result, err
}

// UpdateBlb - update lb
//
// PARAMS:
//     - blbId: lb id
//     - args: the vm image list args
// RETURNS:
//     - *api.UpdateBlbResult: update lb result
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateBlb(blbId string, args *api.UpdateBlbArgs) (*api.UpdateBlbResult, error) {
	if blbId == "" || args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.UpdateBlbResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI() + "/" + blbId, Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

// CreateBlbMonitorPort - create lb monitor port
//
// PARAMS:
//     - blbId: lb id
//     - args: create lb monitor port args
// RETURNS:
//     - *api.BlbMonitorResult: create lb monitor port result
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateBlbMonitorPort(blbId string, args *api.BlbMonitorArgs) (*api.BlbMonitorResult, error) {
	if blbId == "" || args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.BlbMonitorResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI() + "/" + blbId + "/monitor", Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// DeleteBlbMonitorPort - delete lb monitor port
//
// PARAMS:
//     - blbId: lb id
//     - args: delete lb monitor port args
// RETURNS:
//     - *api.BlbMonitorResult: delete lb monitor result
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteBlbMonitorPort(blbId string, args *[]api.Port) (*api.BlbMonitorResult, error) {
	if blbId == "" || args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.BlbMonitorResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI() + "/" + blbId + "/monitor", Body: args, Result: result}
	err := api.Delete(c, req)

	return result, err
}

// GetBlbMonitorPortList - get lb's monitor port list
//
// PARAMS:
//     - blbId: lb id
//     - pageNo: page no
//     - pageSize: page size
// RETURNS:
//     - *api.BlbMonitorListResult: the list of lb monitor ports
//     - error: nil if ok otherwise the specific error
func (c *Client) GetBlbMonitorPortList(blbId string, pageNo, pageSize int) (*api.BlbMonitorListResult, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}
	params := make(map[string]string)
	if pageSize != 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if pageNo != 0 {
		params["pageNo"] = strconv.Itoa(pageNo)
	}
	result := &api.BlbMonitorListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetLoadBalancerURI() + "/" + blbId + "/monitor").
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// UpdateBlbMonitorPort - update lb monitor port
//
// PARAMS:
//     - blbId: lb id
//     - args: monitor info args
// RETURNS:
//     - *api.BlbMonitorResult: update lb monitor result
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateBlbMonitorPort(blbId string, args *api.BlbMonitorArgs) (*api.BlbMonitorResult, error) {
	if blbId == "" || args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.BlbMonitorResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI() + "/" + blbId + "/monitor", Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

// GetBlbMonitorPortDetails - get lb monitor port detail
//
// PARAMS:
//     - blbId: lb id
//     - protocol: protocol
//     - port: port
// RETURNS:
//     - *api.BlbMonitorArgs: lb monitor info result
//     - error: nil if ok otherwise the specific error
func (c *Client) GetBlbMonitorPortDetails(blbId string, protocol api.Protocol, port int) (*api.BlbMonitorArgs, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	params := make(map[string]string)
	if protocol != "" {
		params["protocol"] = string(protocol)
	}
	if port != 0 {
		params["port"] = strconv.Itoa(port)
	}
	result := &api.BlbMonitorArgs{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetLoadBalancerURI() + "/" + blbId + "/monitor/port").
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// BatchCreateBlb - batch create lb
//
// PARAMS:
//     - args: batch create lb args
// RETURNS:
//     - *api.BatchCreateBlbResult: the result of batch create lb
//     - error: nil if ok otherwise the specific error
func (c *Client) BatchCreateBlb(args *api.BatchCreateBlbArgs) (*api.BatchCreateBlbResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.BatchCreateBlbResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI() + "/batch/create", Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// BatchDeleteBlb - batch delete lb
//
// PARAMS:
//     - blbIdList: the list of lb
// RETURNS:
//     - *api.BatchDeleteBlbResult: the result of batch delete lb
//     - error: nil if ok otherwise the specific error
func (c *Client) BatchDeleteBlb(blbIdList []string) (*api.BatchDeleteBlbResult, error) {
	if blbIdList == nil {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.BatchDeleteBlbResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI() + "/batch/delete", Result: result, Body: blbIdList}
	err := api.Post(c, req)

	return result, err
}

// BatchCreateBlbMonitor - batch create lb monitor
//
// PARAMS:
//     - blbId: lb id
//     - args: batch create lb monitor args
// RETURNS:
//     - *api.BatchCreateBlbMonitorResult: the result of batch create lb nonitor
//     - error: nil if ok otherwise the specific error
func (c *Client) BatchCreateBlbMonitor(blbId string, args *api.BatchCreateBlbMonitorArg) (*api.BatchCreateBlbMonitorResult, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.BatchCreateBlbMonitorResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerBatchURI() + "/create/" + blbId + "/monitor", Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// GetBlbBackendPodList - get lb backend list
//
// PARAMS:
//     - blbId: lb id
//     - pageNo: page NO
//     - pageSize: page size
// RETURNS:
//     - *api.GetBlbBackendPodListResult: the result of lb backend list
//     - error: nil if ok otherwise the specific error
func (c *Client) GetBlbBackendPodList(blbId string, pageNo, pageSize int) (*api.GetBlbBackendPodListResult, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	params := make(map[string]string)
	if pageSize != 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if pageNo != 0 {
		params["pageNo"] = strconv.Itoa(pageNo)
	}
	result := &api.GetBlbBackendPodListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetLoadBalancerURI() + "/" + blbId + "/binded").
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetBlbBackendBindingStsList - get lb backend's statefulset list
//
// PARAMS:
//     - blbId: lb id
//     - keyword: the key word
//     - keywordType: key word type
//	   - pageNo: page NO
//	   - pageSize: page size
// RETURNS:
//     - *api.GetBlbBackendBindingStsListResult: the list of sts result
//     - error: nil if ok otherwise the specific error
func (c *Client) GetBlbBackendBindingStsList(blbId string, pageNo, pageSize int, keywordType, keyword string) (*api.GetBlbBackendBindingStsListResult, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	params := make(map[string]string)
	if pageSize != 0 {
		params["pageSize"] = strconv.Itoa(pageSize)
	}
	if pageNo != 0 {
		params["pageNo"] = strconv.Itoa(pageNo)
	}
	if keywordType != "" {
		params["keywordType"] = keywordType
	}

	if keyword != "" {
		params["keyword"] = keyword
	}

	result := &api.GetBlbBackendBindingStsListResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetLoadBalancerURI() + "/" + blbId + "/binding").
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// GetBlbBindingPodListWithSts - image list
//
// PARAMS:
//     - blbId: lb id
//     - stsName: sts name
// RETURNS:
//     - *api.Backends: the list of backend
//     - error: nil if ok otherwise the specific error
func (c *Client) GetBlbBindingPodListWithSts(blbId, stsName string) (*[]api.Backends, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	params := make(map[string]string)
	if stsName != "" {
		params["stsName"] = stsName
	}

	result := &[]api.Backends{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(api.GetLoadBalancerURI() + "/" + blbId + "/bindingpod").
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

// CreateBlbBinding - create lb binding
//
// PARAMS:
//     - blbId: lb id
//     - args: create lb binding args
// RETURNS:
//     - *api.CreateBlbBindingResult: the result of lb binding
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateBlbBinding(blbId string, args *api.CreateBlbBindingArgs) (*api.CreateBlbBindingResult, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.CreateBlbBindingResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI() + "/" + blbId + "/binding", Result: result, Body: args}
	err := api.Post(c, req)

	return result, err
}

// DeleteBlbBindPod - delete lb bind pod
//
// PARAMS:
//     - blbId: lb id
//     - args: delete lb bind pod args
// RETURNS:
//     - *api.DeleteBlbBindPodResult: the result of delete lb pod
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteBlbBindPod(blbId string, args *api.DeleteBlbBindPodArgs) (*api.DeleteBlbBindPodResult, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.DeleteBlbBindPodResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI() + "/" + blbId + "/binded", Result: result, Body: args}
	err := api.Delete(c, req)

	return result, err
}

// UpdateBlbBindPodWeight - update bind pod weight
//
// PARAMS:
//     - blbId: lb id
//     - args: update bind pod weight args
// RETURNS:
//     - *api.UpdateBindPodWeightResult: the result of update bind pod weight
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateBlbBindPodWeight(blbId string, args *api.UpdateBindPodWeightArgs) (*api.UpdateBindPodWeightResult, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.UpdateBindPodWeightResult{}
	req := &api.PostHttpReq{Url: api.GetLoadBalancerURI() + "/" + blbId + "/binded", Result: result, Body: args}
	err := api.Put(c, req)

	return result, err
}

// GetBlbMetrics - get lb metrics
//
// PARAMS:
//     - blbId: lb id
//     - ipType: ip type
//     - port: port
//     - serviceProviderStr: service Provider
//     - offsetInSeconds:  offset Seconds
//     - metricsType: metrics Type
// RETURNS:
//     - *api.ServiceMetricsResult: the list of vm images
//     - error: nil if ok otherwise the specific error
func (c *Client) GetBlbMetrics(blbId, ipType, port, serviceProviderStr string, start, end, stepInMin int, metricsType api.MetricsType) (*api.ServiceMetricsResult, error) {
	if blbId == "" {
		return nil, fmt.Errorf("please set argments")
	}

	params := make(map[string]string)
	params["blbId"] = blbId
	if port != "" {
		params["port"] = port
	}

	if serviceProviderStr != "" {
		params["serviceProvider"] = serviceProviderStr
	}

	if ipType != "" {
		params["ipType"] = ipType
	}
	if start != 0 {
		params["start"] = strconv.Itoa(start)
	}
	if metricsType != "" {
		params["metricsType"] = string(metricsType)
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
		WithURL(api.GetLoadBalancerMonitorURI() + "/" + blbId).
		WithQueryParams(params).
		WithResult(result).
		Do()

	return result, err
}

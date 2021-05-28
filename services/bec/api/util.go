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

// util.go - define the utilities for api package of BEC service
package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"

	DEFAULT_BEC_DOMAIN = "bec." + bce.DEFAULT_REGION + bce.DEFAULT_DOMAIN

	REQUEST_SERVICE_URL = URI_PREFIX + "/service"

	REQUEST_VM_URL = URI_PREFIX + "/vm" + "/service"

	REQUEST_VM_IMAGE_URL = URI_PREFIX + "/vm" + "/image"

	REQUEST_LOADBALANCER_URL = URI_PREFIX + "/blb"

	REQUEST_VM_INSTANCE_URL = URI_PREFIX + "/vm/instance"

	REQUEST_NODE_URL = URI_PREFIX + "/node"
)

/*
var (
	MetricType = map[string]string{
		"cpu":                "cpu",
		"memory":             "memory",
		"bandwidth_receive":  "bandwidth_receive",
		"bandwidth_transmit": "bandwidth_transmit",
		"traffic_receive":    "traffic_receive",
		"traffic_transmit":   "traffic_transmit",

		"node_bw_receive":     "node_bw_receive",
		"node_bw_transmit":    "node_bw_transmit",
		"node_lb_bw_receive":  "node_lb_bw_receive",
		"node_lb_bw_transmit": "node_lb_bw_transmit",
		"request_num":         "request_num",
		"request_rate":        "request_rate",
		"request_delay":       "request_delay",
		"unknown":             "unknown",
	}
)
*/

func GetServiceURI() string {
	return REQUEST_SERVICE_URL
}

func GetServiceMetricsURI(serviceId, metricsType string) string {
	return REQUEST_SERVICE_URL + "/" + serviceId + "/metrics/" + metricsType
}

func GetServiceDetailURI(serviceId string) string {
	return REQUEST_SERVICE_URL + "/" + serviceId
}

func GetStartServiceURI(serviceId, action string) string {
	return REQUEST_SERVICE_URL + "/" + serviceId + "/" + action
}

func GetDeleteServiceURI(serviceId string) string {
	return REQUEST_SERVICE_URL + "/" + serviceId
}

func GetUpdateServiceURI(serviceId string) string {
	return REQUEST_SERVICE_URL + "/" + serviceId
}

func GetBachServiceOperateURI() string {
	return REQUEST_SERVICE_URL + "/batch/operate"
}

func GetBachServiceDeleteURI() string {
	return REQUEST_SERVICE_URL + "/batch/delete"
}

func GetVmURI() string {
	return REQUEST_VM_URL
}

func GetVmServiceActionURI(serviceId, action string) string {
	return REQUEST_VM_URL + "/" + serviceId + "/" + action
}

func GetVmImageURI() string {
	return REQUEST_VM_IMAGE_URL
}

func GetLoadBalancerURI() string {
	return REQUEST_LOADBALANCER_URL
}

func GetLoadBalancerBatchURI() string {
	return REQUEST_LOADBALANCER_URL + "/batch"
}

func GetVmServiceMetricsURI(serviceId, metricsType string) string {
	return REQUEST_VM_URL + "/" + serviceId + "/metrics/" + metricsType
}

func GetVmInstanceURI() string {
	return REQUEST_VM_INSTANCE_URL
}

func GetNodeInfoURI() string {
	return REQUEST_NODE_URL
}

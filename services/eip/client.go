/*
 * Copyright 2017 Baidu, Inc.
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

// client.go - define the client for EIP service

// Package eip defines the EIP services of BCE. The supported APIs are all defined in sub-package
package eip

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"

	DEFAULT_ENDPOINT = "eip." + bce.DEFAULT_REGION + ".baidubce.com"

	REQUEST_EIP_URL = "/eip"

	REQUEST_RECYCLE_EIP_URL = "/eip/recycle"

	REQUEST_EIP_CLUSTER_URL = "/eipcluster"

	REQUEST_EIP_TP_URL = "/eiptp"

	REQUEST_EIP_GROUP_URL = "/eipgroup"

	REQUEST_EIP_BP_URL = "/eipbp"
)

// Client of EIP service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getEipUri() string {
	return URI_PREFIX + REQUEST_EIP_URL
}

func getEipUriWithEip(eip string) string {
	return URI_PREFIX + REQUEST_EIP_URL + "/" + eip
}

func getRecycleEipUri() string {
	return URI_PREFIX + REQUEST_RECYCLE_EIP_URL
}

func getRecycleEipUriWithEip(eip string) string {
	return URI_PREFIX + REQUEST_RECYCLE_EIP_URL + "/" + eip
}

func getEipClusterUri() string {
	return URI_PREFIX + REQUEST_EIP_CLUSTER_URL
}

func getEipClusterUriWithId(clusterId string) string {
	return URI_PREFIX + REQUEST_EIP_CLUSTER_URL + "/" + clusterId
}

func getEipTpUri() string {
	return URI_PREFIX + REQUEST_EIP_TP_URL
}

func getEipTpUriWithId(id string) string {
	return URI_PREFIX + REQUEST_EIP_TP_URL + "/" + id
}

func getEipGroupUri() string {
	return URI_PREFIX + REQUEST_EIP_GROUP_URL
}

func getEipGroupUriWithId(id string) string {
	return URI_PREFIX + REQUEST_EIP_GROUP_URL + "/" + id
}

func getEipBpUrl() string {
	return URI_PREFIX + REQUEST_EIP_BP_URL
}

func getEipBpUrlWithId(id string) string {
	return URI_PREFIX + REQUEST_EIP_BP_URL + "/" + id
}

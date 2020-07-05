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

// client.go - define the client for CCE service

// Package cce defines the CCE services of BCE. The supported APIs are all defined in sub-package

package cce

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"

	DEFAULT_ENDPOINT = "cce." + bce.DEFAULT_REGION + ".baidubce.com"

	REQUEST_CLUSTER_URL = "/cluster"
	REQUEST_NODE_URL    = "/node"
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

func getClusterUri() string {
	return URI_PREFIX + REQUEST_CLUSTER_URL
}

func getClusterUriWithId(clusterUuid string) string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/" + clusterUuid
}

func getNodeUri() string {
	return URI_PREFIX + REQUEST_NODE_URL
}

func getClusterExisteNodeUri() string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/existed_node"
}

func getClusterExisteNodeListUri() string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/existed_bcc_node/list"
}

func getClusterContainerNetUri() string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/get_container_net"
}

func getClusterKubeConfigUri() string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/kubeconfig"
}

func getClusterVersionsUri() string {
	return URI_PREFIX + REQUEST_CLUSTER_URL + "/versions"
}

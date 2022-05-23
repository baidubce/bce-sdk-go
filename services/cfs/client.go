/*
 * Copyright 2022 Baidu, Inc.
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

// client.go - define the client for Application LoadBalance service

// Package cfs defines the Normal CFS services of BCE. The supported APIs are all defined in sub-package
package cfs

import "github.com/baidubce/bce-sdk-go/bce"

const (
	DEFAULT_SERVICE_DOMAIN = "cfs." + bce.DEFAULT_REGION + ".baidubce.com"
	URI_PREFIX             = bce.URI_PREFIX + "v1"
	REQUEST_CFS_URL        = "/cfs"
)

// Client of CFS service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if endPoint == "" {
		endPoint = DEFAULT_SERVICE_DOMAIN
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getCFSUri() string {
	return URI_PREFIX + REQUEST_CFS_URL
}

func getFSInstanceUri(id string) string {
	return URI_PREFIX + REQUEST_CFS_URL + "/" + id
}

func getMountTargetUri(fs_id string, mount_id string) string {
	return URI_PREFIX + REQUEST_CFS_URL + "/" + fs_id + "/" + mount_id
}

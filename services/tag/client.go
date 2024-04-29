/*
 * Copyright 2020 Baidu, Inc.
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

// client.go - define the client for tag service

package tag

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX          = bce.URI_PREFIX
	DEFAULT_ENDPOINT    = "tag.baidubce.com"
	URI_TAG_PREFIX      = "v1/tag"
	URI_RESOURCE_PREFIX = "v1/resource"
	URI_QUERY_FULL_LIST = URI_TAG_PREFIX + "/queryFullList"
	URI_TAG_RESOURCES   = URI_TAG_PREFIX + "/tagResources"
)

// Client of Group service is a kind of BceClient, so derived from BceClient
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

func queryFullListUri() string {
	return URI_PREFIX + URI_QUERY_FULL_LIST
}

func tagsResourcesUri() string {
	return URI_PREFIX + URI_TAG_RESOURCES
}

func tagBaseUri() string {
	return URI_PREFIX + URI_TAG_PREFIX
}
func resourceBaseUri() string {
	return URI_PREFIX + URI_RESOURCE_PREFIX
}

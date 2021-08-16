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

// client.go - define the client for QuotaCenter service

// Package quotacenter defines the QuotaCenter services of BCE. The supported APIs are all defined in sub-package
package quotacenter

import "github.com/baidubce/bce-sdk-go/bce"

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"

	DEFAULT_ENDPOINT = "quota-center.baidubce.com"

	BASE_QUOTA_CENTER_URL = "/quota_center"

	BASE_PRODUCT_URL = "/info/product"

	BASE_INFO_URL = "/info"

	BASE_REGION_URL = "/info/region"
)

// Client of QUOTA_CENTER service is a kind of BceClient, so derived from BceClient
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

func getQuotaCenterUri() string {
	return URI_PREFIX + BASE_QUOTA_CENTER_URL
}

func getProductUri() string {
	return getQuotaCenterUri() + BASE_PRODUCT_URL
}

func getInfoUri() string {
	return getQuotaCenterUri() + BASE_INFO_URL
}

func getRegionUri() string {
	return getQuotaCenterUri() + BASE_REGION_URL
}

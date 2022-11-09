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

// client.go - define the client for Application LBDC service

// Package lbdc defines the LBDC services of BCE. The supported APIs are all defined in sub-package
package lbdc

import "github.com/baidubce/bce-sdk-go/bce"

const (
	DEFAULT_SERVICE_DOMAIN = "blb." + bce.DEFAULT_REGION + ".baidubce.com"
	URI_PREFIX             = bce.URI_PREFIX + "v1"
	REQUEST_LBDC_URL       = "/lbdc"
)

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

func getUrlForLbdc() string {
	return URI_PREFIX + REQUEST_LBDC_URL
}

func getUrlForLbdcId(lbdcId string) string {
	return getUrlForLbdc() + "/" + lbdcId
}

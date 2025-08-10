/*
 * Copyright 2024 Baidu, Inc.
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

// Package billing: billing client for open api
package billing

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	DEFAULT_ENDPOINT              = "http://billing.baidubce.com"
	VERSION                       = "v1"
	URI_RESOURCE_MONTH_BILL       = "/bill/resource/month"
	URI_RESOURCE_CHARGE_ITEM_BILL = "/bill/resource/chargeitem"
	URI_SHARE_BILL                = "/bill/share/bill"
	URI_COST_SPLIT_BILL           = "/bill/cost/split/bill/list"
)

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
	return &Client{BceClient: client}, nil
}

func getBillingPrefix() string {
	return bce.URI_PREFIX + VERSION
}

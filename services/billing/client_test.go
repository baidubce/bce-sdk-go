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

// client_test.go - test for billing/client.go
package billing

import (
	"encoding/json"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

func TestResourceMonthBill(t *testing.T) {
	log.SetLogLevel(log.DEBUG)
	client := getClient()
	bill, err := client.ResourceMonthBill("2024-02", "", "", "postpay", "", "", 1, 10)
	if err != nil {
		log.Error(err)
	}
	log.Info(json.Marshal(bill))
}

func TestResourceChargeItemBill(t *testing.T) {
	log.SetLogLevel(log.DEBUG)
	client := getClient()
	request := ResourceChargeItemBillRequest{
		BillMonth:              "2025-06",
		QueryAccountId:         "accountId",
		PageNo:                 1,
		PageSize:               2,
		NeedSplitConfiguration: true,
	}
	bill, err := client.ResourceChargeItemBill(request)
	if err != nil {
		log.Error(err)
	}
	log.Info(json.Marshal(bill))
}

func TestShareBill(t *testing.T) {
	log.SetLogLevel(log.DEBUG)
	client := getClient()
	request := ShareBillRequest{
		Month:                  "2025-06",
		QueryAccountId:         "accountId",
		PageNo:                 1,
		PageSize:               5,
		NeedSplitConfiguration: true,
	}
	bill, err := client.ShareBill(request)
	if err != nil {
		log.Error(err)
	}
	log.Info(json.Marshal(bill))
}

func TestCostSplitBill(t *testing.T) {
	log.SetLogLevel(log.DEBUG)
	client := getClient()
	request := CostSplitBillRequest{
		Month:                  "2025-06",
		QueryAccountId:         "accountId",
		PageNo:                 1,
		PageSize:               100,
		NeedSplitConfiguration: true,
		ServiceType:            "BLB",
	}
	bill, err := client.CostSplitBill(request)
	if err != nil {
		log.Error(err)
	}
	log.Info(json.Marshal(bill))
}

func getClient() *Client {
	client, _ := NewClient("ak", "sk", "endpoint")
	return client
}

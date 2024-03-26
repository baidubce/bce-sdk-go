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

// Package billing provides the billing service.
package billing

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ResourceMonthBill - get user's resource month bill
//
// PARAMS:
//   - month: the month of bill, format: yyyy-MM
//   - beginTime: the start time of bill, format: yyyy-MM-dd
//   - endTime: the end time of bill, format: yyyy-MM-dd
//   - productType: is required. the type of bill, "postpay" or "prepay"
//   - serviceType: bill filter condition. english code, em: "BCC","CDS"...
//   - granularity: the granularity of bill, "monthly" or "daily"
//   - queryAccountId: query sub-account
//   - pageNo: the page number of result, min:1
//   - pageSize: the size of each page, max:100
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ResourceMonthBill(month string, beginTime string, endTime string, productType string, serviceType string,
	queryAccountId string, pageNo int, pageSize int) (*ResourceMonthBillResponse, error) {
	result := &ResourceMonthBillResponse{}
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 100
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBillingPrefix()+URI_RESOURCE_MONTH_BILL).
		WithQueryParamFilter("month", month).
		WithQueryParamFilter("beginTime", beginTime).
		WithQueryParamFilter("endTime", endTime).
		WithQueryParamFilter("serviceType", serviceType).
		WithQueryParamFilter("queryAccountId", queryAccountId).
		WithQueryParam("productType", productType).
		WithQueryParam("pageNo", strconv.Itoa(pageNo)).
		WithQueryParam("pageSize", strconv.Itoa(pageSize)).
		WithResult(result).
		Do()

	return result, err
}

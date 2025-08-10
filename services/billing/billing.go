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

// ResourceChargeItemBill - get user's resource charge item bill
//
// PARAMS:
//   - request: open api request body, see https://cloud.baidu.com/doc/Finance/s/qlfuqsf02
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ResourceChargeItemBill(request ResourceChargeItemBillRequest) (*ResourceChargeItemBillResponse, error) {

	if request.PageNo <= 0 {
		request.PageNo = 1
	}
	if request.PageSize <= 0 || request.PageSize > 100 {
		request.PageSize = 100
	}
	result := &ResourceChargeItemBillResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBillingPrefix() + URI_RESOURCE_CHARGE_ITEM_BILL).
		WithBody(request).
		WithResult(result).
		Do()
	return result, err
}

// ShareBill - get user's share bill
//
// PARAMS:
//   - request: open api request param, see https://cloud.baidu.com/doc/Finance/s/vmck88hhv
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ShareBill(request ShareBillRequest) (*ShareBillResponse, error) {
	result := &ShareBillResponse{}
	if request.PageNo <= 0 {
		request.PageNo = 1
	}
	if request.PageSize <= 0 || request.PageSize > 100 {
		request.PageSize = 100
	}

	httpRequest := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBillingPrefix()+URI_SHARE_BILL).
		WithQueryParam("month", request.Month).
		WithQueryParam("startDay", request.StartDay).
		WithQueryParam("endDay", request.EndDay).
		WithQueryParam("showDeductPrice", strconv.FormatBool(request.ShowDeductPrice)).
		WithQueryParam("showControversial", strconv.FormatBool(request.ShowControversial)).
		WithQueryParam("showTags", strconv.FormatBool(request.ShowTags)).
		WithQueryParam("needSplitConfiguration", strconv.FormatBool(request.NeedSplitConfiguration)).
		WithQueryParam("pageNo", strconv.Itoa(request.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(request.PageSize))

	if request.StartDay != "" {
		httpRequest = httpRequest.WithQueryParam("productType", request.ProductType)
	}
	if request.ServiceType != "" {
		httpRequest = httpRequest.WithQueryParam("serviceType", request.ServiceType)
	}
	if request.QueryAccountId != "" {
		httpRequest = httpRequest.WithQueryParam("queryAccountId", request.QueryAccountId)
	}
	if request.DisplaySystemUnit != "" {
		httpRequest = httpRequest.WithQueryParam("displaySystemUnit", request.DisplaySystemUnit)
	}
	err := httpRequest.WithResult(result).Do()

	return result, err
}

// CostSplitBill - get user's cost split bill
//
// PARAMS:
//   - request: open api request param, see https://cloud.baidu.com/doc/Finance/s/Umck9e681
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) CostSplitBill(request CostSplitBillRequest) (*ShareBillResponse, error) {
	result := &ShareBillResponse{}
	if request.PageNo <= 0 {
		request.PageNo = 1
	}
	if request.PageSize <= 0 || request.PageSize > 100 {
		request.PageSize = 100
	}

	httpRequest := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBillingPrefix()+URI_COST_SPLIT_BILL).
		WithQueryParam("month", request.Month).
		WithQueryParam("startDay", request.StartDay).
		WithQueryParam("endDay", request.EndDay).
		WithQueryParam("showTags", strconv.FormatBool(request.ShowTags)).
		WithQueryParam("needSplitConfiguration", strconv.FormatBool(request.NeedSplitConfiguration)).
		WithQueryParam("pageNo", strconv.Itoa(request.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(request.PageSize))
	if request.ServiceType != "" {
		httpRequest = httpRequest.WithQueryParam("serviceType", request.ServiceType)
	}
	if request.InstanceId != "" {
		httpRequest = httpRequest.WithQueryParam("instanceId", request.InstanceId)
	}
	if request.QueryAccountId != "" {
		httpRequest = httpRequest.WithQueryParam("queryAccountId", request.QueryAccountId)
	}
	err := httpRequest.WithResult(result).Do()
	return result, err
}

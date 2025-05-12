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

// ddos.go - the ddos APIs definition supported by the EIP service

package eip

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	VERSION_V1 = "v1"
)

/*
ListDdos listDdos

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListDdosResponse
*/
func (c *Client) ListDdos(request *ListDdosRequest) (*ListDdosResponse, error) {
	result := &ListDdosResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListDdosUri(VERSION_V1)).
		WithQueryParamFilter("ips", request.Ips).
		WithQueryParamFilter("type", request.Type).
		WithQueryParamFilter("marker", request.Marker).
		WithResult(result).
		Do()
	return result, err
}

/*
ListDdosAttackRecord listDdosAttackRecord

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ip
	@return ApiListDdosAttackRecordResponse
*/
func (c *Client) ListDdosAttackRecord(request *ListDdosAttackRecordRequest) (*ListDdosAttackRecordResponse, error) {
	if request.Ip == "" {
		return nil, fmt.Errorf("ip is required and must be specified")
	}
	result := &ListDdosAttackRecordResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListDdosAttackRecordUri(VERSION_V1, request.Ip)).
		WithQueryParamFilter("startTime", request.StartTime).
		WithQueryParamFilter("marker", request.Marker).
		WithResult(result).
		Do()
	return result, err
}

/*
ModifyDdosThreshold modifyDdosThreshold

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ip
	@return Api
*/
func (c *Client) ModifyDdosThreshold(request *ModifyDdosThresholdRequest) error {
	if request.Ip == "" {
		return fmt.Errorf("ip is required and must be specified")
	}
	if request.ThresholdType == "" {
		return fmt.Errorf("ThresholdType is required and must be specified")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getModifyDdosThresholdUri(VERSION_V1, request.Ip)).
		WithQueryParamFilter("clientToken", request.ClientToken).
		WithQueryParam("modifyThreshold", "").
		WithBody(request).
		Do()
	return err
}

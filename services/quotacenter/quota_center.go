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

// quota_center.go - the quota_center APIs definition supported by the QUOTA_CENTER service
package quotacenter

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListProducts - list quota center support products.
//
// PARAMS:
//   - args: the arguments to list products.
//
// RETURNS:
//   - *ListProductResult: the result of list products.
//   - error: nil if success otherwise the specific error
func (c *Client) ListProducts(args *ProductQueryArgs) (*ListProductResult, error) {
	if args == nil {
		args = &ProductQueryArgs{}
	}
	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}
	result := &ListProductResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getProductUri()).
		WithQueryParamFilter("productType", args.ProductType).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// ListRegions - list quota center support regions with the specific parameters.
//
// PARAMS:
//   - args: the arguments to list regions.
//
// RETURNS:
//   - *ListRegionResult: the result of regions.
//   - error: nil if success otherwise the specific error
func (c *Client) ListRegions(args *RegionQueryArgs) (*ListRegionResult, error) {
	if args == nil {
		args = &RegionQueryArgs{}
	}
	result := &ListRegionResult{}
	if len(args.ProductType) == 0 {
		return nil, errors.New("productType should not be empty")
	}
	if len(args.ServiceType) == 0 {
		return nil, errors.New("serviceType should not be empty")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRegionUri()).
		WithQueryParamFilter("productType", args.ProductType).
		WithQueryParamFilter("serviceType", args.ServiceType).
		WithQueryParamFilter("type", args.Type).
		WithResult(result).
		Do()

	return result, err
}

// QuotaCenterQuery - query from quota_center with the specific parameters
//
// PARAMS:
//   - args: the arguments to query quota_center
//
// RETURNS:
//   - *ListQuotaResult: the result of query from quota_center.
//   - error: nil if success otherwise the specific error
func (c *Client) QuotaCenterQuery(args *QuotaCenterQueryArgs) (*ListQuotaResult, error) {
	if args == nil {
		args = &QuotaCenterQueryArgs{}
	}
	if len(args.ServiceType) == 0 {
		return nil, errors.New("serviceType should not be empty")
	}
	if len(args.Region) == 0 {
		return nil, errors.New("region should not be empty")
	}
	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}
	result := &ListQuotaResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuotaCenterUri()).
		WithQueryParamFilter("type", args.Type).
		WithQueryParamFilter("serviceType", args.ServiceType).
		WithQueryParamFilter("region", args.Region).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// InfoQuery - query basic infos from quota_center with the specific parameters
//
// PARAMS:
//   - args: the arguments to query infos.
//
// RETURNS:
//   - *ListInfoResult: the result of infos from quota_center.
//   - error: nil if success otherwise the specific error
func (c *Client) InfoQuery(args *InfoQueryArgs) (*ListInfoResult, error) {
	if args == nil {
		args = &InfoQueryArgs{}
	}
	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}
	result := &ListInfoResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInfoUri()).
		WithQueryParamFilter("serviceType", args.ServiceType).
		WithQueryParamFilter("region", args.Region).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// Apply - apply quota or whitelist with the specific parameters
//
// PARAMS:
//     - args: the arguments to apply.
// RETURNS:
//     - *IdModel: the id of application.
//     - error: nil if success otherwise the specific error

func (c *Client) Apply(args *ApplicationCreateModel) (*IdModel, error) {
	if args == nil {
		args = &ApplicationCreateModel{}
	}

	result := &IdModel{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithURL(getApplyUri()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ApplicationQuery - query applications from quota_center with the specific parameters
//
// PARAMS:
//   - args: the arguments to query application.
//
// RETURNS:
//   - *ListApplicationResult: the result of applications.
//   - error: nil if success otherwise the specific error
func (c *Client) ApplicationQuery(args *ApplicationQueryArgs) (*ListApplicationResult, error) {
	if args == nil {
		args = &ApplicationQueryArgs{}
	}
	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}
	result := &ListApplicationResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getApplyUri()).
		WithQueryParamFilter("id", args.Id).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("status", args.Status).
		WithQueryParamFilter("productType", args.ProductType).
		WithQueryParamFilter("serviceType", args.ServiceType).
		WithQueryParamFilter("type", args.Type).
		WithQueryParamFilter("region", args.Region).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// ApplicationDetail - query application detail from quota_center with id.
//
// PARAMS:
//   - id: the application's id.
//
// RETURNS:
//   - *ApplicationModel: the result of application.
//   - error: nil if success otherwise the specific error
func (c *Client) ApplicationDetail(id string) (*ApplicationModel, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("please set id argment")
	}

	result := &ApplicationModel{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getApplyUriWithId(id)).
		WithResult(result).
		Do()

	return result, err
}

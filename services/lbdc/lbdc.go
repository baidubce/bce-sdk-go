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

// Package lbdc lbdc.go - the LBDC APIs definition supported by the LBDC service
package lbdc

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateLbdc - create the LBDC instance with the specific parameters
//
// PARAMS:
//   - args: the arguments to create LBDC
//
// RETURNS:
//   - *CreateLoadBalancerResult: the result of create LoadBalancer, contains new LoadBalancer's ID
//   - error: nil if success otherwise the specific error
func (c *Client) CreateLbdc(args *CreateLbdcArgs) (*CreateLbdcResult, error) {
	if args == nil {
		return nil, fmt.Errorf("the CreateLbdcArgs cannot be nil")
	}
	result := &CreateLbdcResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getUrlForLbdc()).
		WithMethod(http.POST).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// UpgradeLbdc - upgrade LBDC with the specific parameters
//
// PARAMS:
//   - args: the arguments to update LBDC
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpgradeLbdc(args *UpgradeLbdcArgs) error {
	if args == nil {
		return fmt.Errorf("the UpgradeLbdcArgs cannot be nil")
	}

	if len(args.Id) == 0 {
		return fmt.Errorf("the LbdcId cannot be empty")
	}
	return bce.NewRequestBuilder(c).
		WithURL(getUrlForLbdcId(args.Id)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		Do()
}

// RenewLbdc - renew LBDC with the specific parameters
//
// PARAMS:
//   - args: the arguments to renew LBDC
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RenewLbdc(args *RenewLbdcArgs) error {
	if args == nil {
		return fmt.Errorf("the RenewLbdcArgs cannot be nil")
	}

	if len(args.Id) == 0 {
		return fmt.Errorf("the LbdcId cannot be empty")
	}
	return bce.NewRequestBuilder(c).
		WithURL(getUrlForLbdcId(args.Id)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("purchaseReserved", "").
		Do()
}

// ListLbdc - list LBDC with the specific id and/or name
//
// PARAMS:
//   - args: the arguments to list LBDC instances
//
// RETURNS:
//   - *ListSslVpnUserResult: the result of Cluster list contains page infos
//   - error: nil if success otherwise the specific error
func (c *Client) ListLbdc(args *ListLbdcArgs) (*ListLbdcResult, error) {
	if args == nil {
		args = &ListLbdcArgs{}
	}
	result := &ListLbdcResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getUrlForLbdc()).
		WithMethod(http.GET).
		WithResult(result)
	if len(args.Id) > 0 {
		builder.WithQueryParamFilter("id", args.Id)
	}
	if len(args.Name) > 0 {
		builder.WithQueryParamFilter("name", args.Name)
	}
	err := builder.Do()
	return result, err
}

// GetLbdcDetail - get details of the specific lbdc
//
// PARAMS:
//   - lbdcId: the id of the specified lbdc
//
// RETURNS:
//   - *LbdcDetailResult: the result of the specific lbdc details
//   - error: nil if success otherwise the specific error
func (c *Client) GetLbdcDetail(lbdcId string) (*GetLbdcDetailResult, error) {
	if len(lbdcId) == 0 {
		return nil, fmt.Errorf("the LbdcId cannot be empty")
	}

	result := &GetLbdcDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getUrlForLbdcId(lbdcId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// UpdateLbdc - update lbdc with the specific parameters
//
// PARAMS:
//   - args: the arguments to update lbdc
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateLbdc(args *UpdateLbdcArgs) error {
	if args == nil {
		return fmt.Errorf("the UpdateLbdcArgs cannot be nil")
	}
	if len(args.Id) == 0 {
		return fmt.Errorf("the LbdcId cannot be empty")
	}

	if args.UpdateLbdcBody == nil {
		return fmt.Errorf("the UpdateLbdcBody cannot be nil")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getUrlForLbdcId(args.Id)).
		WithMethod(http.PUT).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args.UpdateLbdcBody).
		Do()
}

// GetBoundBlBListOfLbdc - get Bound blb list of lbdc
//
// PARAMS:
//   - lbdcId: the id of the specified lbdc
//
// RETURNS:
//   - *GetBoundBlBListOfLbdcResult: the result of the bound blb list of lbdc
//   - error: nil if success otherwise the specific error
func (c *Client) GetBoundBlBListOfLbdc(lbdcId string) (*GetBoundBlBListOfLbdcResult, error) {
	if len(lbdcId) == 0 {
		return nil, fmt.Errorf("the LbdcId cannot be empty")
	}
	result := &GetBoundBlBListOfLbdcResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getUrlForLbdcId(lbdcId) + "/blb").
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

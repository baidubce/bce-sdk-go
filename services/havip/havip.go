/*
 * Copyright  Baidu, Inc.
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

package havip

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

// CreateHaVip - create an havip with the specific parameters
//
// PARAMS:
//   - args: the arguments to create an havip
//
// RETURNS:
//   - *CreateHavipResult: the result of create havip
//   - error: nil if success otherwise the specific error
func (c *Client) CreateHaVip(args *CreateHaVipArgs) (*CreateHavipResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createHaVipArgs cannot be nil.")
	}

	result := &CreateHavipResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForHaVip()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListHaVip - list all havip with the specific parameters
//
// PARAMS:
//   - args: the arguments to list all havip
//
// RETURNS:
//   - *ListHaVipResult: the result of list all havip
//   - error: nil if success otherwise the specific error
func (c *Client) ListHaVip(args *ListHaVipArgs) (*ListHaVipResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The listHaVipArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListHaVipResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getURLForHaVip()).
		WithMethod(http.GET).
		WithQueryParamFilter("vpcId", args.VpcId).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys))

	err := builder.WithResult(result).Do()

	return result, err
}

// GetHaVipDetail - get the havip detail
//
// PARAMS:
//   - haVipId: the specific haVipId
//
// RETURNS:
//   - *HaVip: the havip
//   - error: nil if success otherwise the specific error
func (c *Client) GetHaVipDetail(haVipId string) (*HaVip, error) {
	if haVipId == "" {
		return nil, fmt.Errorf("The haVipId cannot be empty.")
	}

	result := &HaVip{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForHaVipId(haVipId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

// UpdateHaVip - update an havip
//
// PARAMS:
//   - UpdateHaVipArgs: the arguments to update an havip
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateHaVip(args *UpdateHaVipArgs) error {
	if args == nil {
		return fmt.Errorf("The updateHaVipArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForHaVipId(args.HaVipId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("modifyAttribute", "").
		Do()
}

// DeleteHaVip - delete an havip
//
// PARAMS:
//   - DeleteHaVipArgs: the arguments to delete an havip
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteHaVip(args *DeleteHaVipArgs) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForHaVipId(args.HaVipId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// HaVipAttachInstance - havip attach instance
//
// PARAMS:
//   - args: the arguments to attach instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) HaVipAttachInstance(args *HaVipInstanceArgs) error {
	if args == nil {
		return fmt.Errorf("The haVipInstanceArgs cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForHaVipId(args.HaVipId)).
		WithMethod(http.PUT).
		WithQueryParam("attach", "").
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

// HaVipDetachInstance - havip detach instance
//
// PARAMS:
//   - args: the arguments to detach instance
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) HaVipDetachInstance(args *HaVipInstanceArgs) error {
	if args == nil {
		return fmt.Errorf("The haVipInstanceArgs cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForHaVipId(args.HaVipId)).
		WithMethod(http.PUT).
		WithQueryParam("detach", "").
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

// HaVipBindPublicIp - havip bind public ip
//
// PARAMS:
//   - args: the arguments to bind public ip
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) HaVipBindPublicIp(args *HaVipBindPublicIpArgs) error {
	if args == nil {
		return fmt.Errorf("The haVipBindPublicIpArgs cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForHaVipId(args.HaVipId)).
		WithMethod(http.PUT).
		WithQueryParam("bindPublicIp", "").
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

// HaVipUnbindPublicIp - havip unbind public ip
//
// PARAMS:
//   - args: the arguments to unbind public ip
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) HaVipUnbindPublicIp(args *HaVipUnbindPublicIpArgs) error {
	if args == nil {
		return fmt.Errorf("The HaVipUnbindPublicIpArgs cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForHaVipId(args.HaVipId)).
		WithMethod(http.PUT).
		WithQueryParam("unbindPublicIp", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

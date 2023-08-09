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

package eip

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateEipGroup - create an EIP_GROUP with the specific parameters
//
// PARAMS:
//   - args: the arguments to create an eipGroup
//
// RETURNS:
//   - *CreateEipGroupResult: the result of create EIP_GROUP, contains new EIP_GROUP's id
//   - error: nil if success otherwise the specific error
func (c *Client) CreateEipGroup(args *CreateEipGroupArgs) (*CreateEipGroupResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set create eip argments")
	}

	if args.Billing == nil {
		return nil, fmt.Errorf("please set billing")
	}

	result := &CreateEipGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEipGroupUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ResizeEipGroupBandWidth - resize an EIP_GROUP with the specific parameters
//
// PARAMS:
//   - id: the eipGroup's id
//   - args: the arguments to resize an EIP_GROUP
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ResizeEipGroupBandWidth(id string, args *ResizeEipGroupArgs) error {
	if args == nil {
		return fmt.Errorf("please set resize argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipGroupUriWithId(id)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		WithBody(args).
		Do()
}

// EipGroupAddEipCount - increase EIP_GROUP's ip count with the specific parameters
//
// PARAMS:
//   - id: the eipGroup's id
//   - args: the arguments to increase EIP_GROUP's ip count
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) EipGroupAddEipCount(id string, args *GroupAddEipCountArgs) error {
	if args == nil {
		return fmt.Errorf("please set resize argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipGroupUriWithId(id)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		WithBody(args).
		Do()
}

// ReleaseEipGroupIps - release EIP_GROUP's ips with the specific parameters
//
// PARAMS:
//   - id: the eipGroup's id
//   - args: the arguments to release EIP_GROUP's ips
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ReleaseEipGroupIps(id string, args *ReleaseEipGroupIpsArgs) error {
	if args == nil {
		return fmt.Errorf("please set resize argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipGroupUriWithId(id)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		WithBody(args).
		Do()
}

// RenameEipGroup - rename EIP_GROUP's name with the specific parameters
//
// PARAMS:
//   - id: the eipGroup's id
//   - args: the arguments to rename EIP_GROUP
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RenameEipGroup(id string, args *RenameEipGroupArgs) error {
	if args == nil {
		return fmt.Errorf("please set rename argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipGroupUriWithId(id)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("update", "").
		WithBody(args).
		Do()
}

// DeleteEipGroup - delete an EIP_GROUP
//
// PARAMS:
//   - id: the specific EIP_GROUP's id
//   - clientToken: optional parameter, an Idempotent Token
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteEipGroup(id, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getEipGroupUriWithId(id)).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// ListEipGroup - list all EIP_GROUP with the specific parameters
//
// PARAMS:
//   - args: the arguments to list all eipGroup
//
// RETURNS:
//   - *ListEipGroupResult: the result of list all eipGroup
//   - error: nil if success otherwise the specific error
func (c *Client) ListEipGroup(args *ListEipGroupArgs) (*ListEipGroupResult, error) {
	if args == nil {
		args = &ListEipGroupArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListEipGroupResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipGroupUri()).
		WithQueryParamFilter("id", args.Id).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("status", args.Status).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// EipGroupDetail - get EIP_GROUP detail
//
// PARAMS:
//   - id: the eipGroup's id
//
// RETURNS:
//   - *EipGroupModel: the result of list all eip in the recycle bin
//   - error: nil if success otherwise the specific error
func (c *Client) EipGroupDetail(id string) (*EipGroupModel, error) {
	result := &EipGroupModel{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipGroupUriWithId(id)).
		WithResult(result).
		Do()

	return result, err
}

// EipGroupMoveOut - move eips out of EIP_GROUP with the specific parameters
//
// PARAMS:
//   - id: the eipGroup's id
//   - args: the arguments to move out EIP_GROUP
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) EipGroupMoveOut(id string, args *EipGroupMoveOutArgs) error {
	if args == nil {
		return fmt.Errorf("please set argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipGroupUriWithId(id)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("move_out", "").
		WithBody(args).
		Do()
}

// EipGroupMoveIn - move eips into to EIP_GROUP with the specific parameters
//
// PARAMS:
//   - id: the eipGroup's id
//   - args: the arguments to move in EIP_GROUP
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) EipGroupMoveIn(id string, args *EipGroupMoveInArgs) error {
	if args == nil {
		return fmt.Errorf("please set argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipGroupUriWithId(id)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("move_in", "").
		WithBody(args).
		Do()
}

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

// CreateEipBp - create an EIP BP with the specific parameters
//
// PARAMS:
//     - args: the arguments to create an eipbp
// RETURNS:
//     - *CreateEipBpResult: the result of create EIP BP, contains new EIP BP's id
//     - error: nil if success otherwise the specific error
func (c *Client) CreateEipBp(args *CreateEipBpArgs) (*CreateEipBpResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set create eipbp argments")
	}

	result := &CreateEipBpResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEipBpUrl()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ResizeEIpBp - resize an EIP BP with the specific parameters
//
// PARAMS:
//     - id: the id of EIP BP
//     - args: the arguments to resize an EIP BP
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ResizeEipBp(id string, args *ResizeEipBpArgs) error {
	if args == nil {
		return fmt.Errorf("please set resize eipbp argments")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipBpUrlWithId(id)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		WithBody(args).
		Do()

	return err
}

// GetEipBp - get the EIP BP detail with the id
//
// PARAMS:
//     - id: the specific eipbp id
//     - clientToken: the specific client token
// RETURNS:
//     - *EipBpDetail: the result of eipbp detail
//     - error: nil if success otherwise the specific error
func (c *Client) GetEipBp(id, clientToken string) (*EipBpDetail, error) {

	result := &EipBpDetail{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipBpUrlWithId(id)).
		WithQueryParamFilter("clientToken", clientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListEipBp - list all EIP BP with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all eipBp
// RETURNS:
//     - *EipBpListResult: the result of list all eipBp
//     - error: nil if success otherwise the specific error
func (c *Client) ListEipBp(args *ListEipBpArgs) (*ListEipBpResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set list eipbp argments")
	}

	result := &ListEipBpResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipBpUrl()).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParamFilter("id", args.Id).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("bindType", args.BindType).
		WithResult(result).
		Do()

	return result, err
}

// UpdateEipBpAutoReleaseTime - update the auto release time of an EIP BP
//
// PARAMS:
//     - id: the specific eipbp id
//     - args: the arguments to update eipbp auto release time
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateEipBpAutoReleaseTime(id string, args *UpdateEipBpAutoReleaseTimeArgs) error {
	if args == nil {
		return fmt.Errorf("please set update eipbp autoReleaseTime argments")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipBpUrlWithId(id)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("retime", "").
		WithBody(args).
		Do()

	return err
}

// UpdateEipBpName - update the Name of an EIP BP
//
// PARAMS:
//     - id: the specific eipbp id
//     - args: the arguments to update eipbp name
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateEipBpName(id string, args *UpdateEipBpNameArgs) error {
	if args == nil {
		return fmt.Errorf("please set update eipbp name argments")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipBpUrlWithId(id)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("rename", "").
		WithBody(args).
		Do()

	return err
}

// DeleteEipBp - delete an EIP BP with the specific id
//
// PARAMS:
//     - id: the specific eipbp id
//     - clientToken: the specific client token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteEipBp(id, clientToken string) error {
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getEipBpUrlWithId(id)).
		WithQueryParamFilter("clientToken", clientToken).
		Do()

	return err
}

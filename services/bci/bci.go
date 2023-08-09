/*
 * Copyright 2023 Baidu, Inc.
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

package bci

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateInstance - create a bci with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a bci
//
// RETURNS:
//   - *CreateInstanceResult: the result of create bci
//   - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(args *CreateInstanceArgs) (*CreateInstanceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The CreateInstanceArgs cannot be nil.")
	}

	result := &CreateInstanceResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForBci()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListInstance - list all bcis with the specific parameters
//
// PARAMS:
//   - args: the arguments to list all bcis
//
// RETURNS:
//   - *ListInstanceResult: the result of list all bcis
//   - error: nil if success otherwise the specific error
func (c *Client) ListInstances(args *ListInstanceArgs) (*ListInstanceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The ListInstanceArgs cannot be nil.")
	}
	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListInstanceResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForBci()).
		WithMethod(http.GET).
		WithQueryParam("keywordType", args.KeywordType).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetInstance - query bci detail with the specific parameters
//
// PARAMS:
//   - args: the arguments to list all bci
//
// RETURNS:
//   - *GetInstanceResult: the result of query bci detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetInstance(args *GetInstanceArgs) (*GetInstanceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The GetInstanceArgs cannot be nil.")
	}

	result := &GetInstanceResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForBciId(args.InstanceId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

// DeleteInstance - delete a bci
//
// PARAMS:
//   - args: the arguments to delete a bci
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(args *DeleteInstanceArgs) error {
	if args == nil {
		return fmt.Errorf("The DeleteInstanceArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForBciId(args.InstanceId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("relatedReleaseFlag", strconv.FormatBool(args.RelatedReleaseFlag)).
		Do()
}

// BatchDeleteInstance - batch delete bcis
//
// PARAMS:
//   - args: the arguments to batch delete bcis
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) BatchDeleteInstance(args *BatchDeleteInstanceArgs) error {
	if args == nil {
		return fmt.Errorf("The BatchDeleteInstanceArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForBci() + "/batchDel").
		WithMethod(http.POST).
		WithBody(args).
		Do()
}

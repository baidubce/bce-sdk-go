/*
 * Copyright 2020 Baidu, Inc.
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

// resmanager.go - the resmanager APIs definition supported by the resmanager service
package resmanager

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) BindResourceToGroup(args *BindResourceToGroupArgs) (*BindResourceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if len(args.Bindings) == 0 {
		return nil, fmt.Errorf("unset Bindings")
	}

	result := &BindResourceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAddToGroupUri()).
		WithQueryParamFilter("force", "false").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ChangeResourceGroup(args *ChangeResourceGroupArgs) (*BindResourceResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if len(args.MoveResModels) == 0 {
		return nil, fmt.Errorf("unset move res models")
	}

	result := &BindResourceResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getChangeGroupUri()).
		WithQueryParamFilter("force", "false").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) RemoveResourceFromGroup(args *BindResourceToGroupArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(args.Bindings) == 0 {
		return fmt.Errorf("unset bindings")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAddToGroupUri()).
		WithQueryParamFilter("force", "false").
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		Do()
	return err
}

func (c *Client) QueryGroupList(name string) (*GroupList, error) {
	if name == "" {
		return nil, fmt.Errorf("unset group name")
	}
	result := &GroupList{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQueryGroupUri()).
		WithQueryParamFilter("name", name).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) getResGroupBatch(args *ResGroupDetailRequest) (*ResGroupDetailResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}
	if len(args.ResourceBrief) == 0 {
		return nil, fmt.Errorf("unset ResourceBrief")
	}
	result := &ResGroupDetailResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getGroupBatchUri()).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

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

// dts.go - the dts APIs definition supported by the DTS service
package dts

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateDts - create a dtsTask with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a dtsTask
// RETURNS:
//     - *CreateDtsResult: the result of create dtsTask, contains new dtsTask's ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateDts(args *CreateDtsArgs) (*CreateDtsResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.ProductType == "" {
		return nil, fmt.Errorf("unset ProductType")
	}

	if args.Type == "" {
		return nil, fmt.Errorf("unset Type")
	}

	if args.Standard == "" {
		return nil, fmt.Errorf("unset Standard")
	}

	if args.SourceInstanceType == "" {
		return nil, fmt.Errorf("unset SourceInstanceType")
	}

	if args.TargetInstanceType == "" {
		return nil, fmt.Errorf("unset TargetInstanceType")
	}

	result := &CreateDtsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDtsUri()).
		WithQueryParamFilter("orderType","NEW").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// DeleteDts - delete a dtsTask
//
// PARAMS:
//     - taskId: the specific taskId
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteDts(taskId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getDtsUriWithTaskId(taskId)).
		Do()
}

// GetDetail - get a specific dtsTask's detail
//
// PARAMS:
//     - taskId: the specific dtsTask's ID
// RETURNS:
//     - *DtsTaskMeta: the specific dtsTask's detail
//     - error: nil if success otherwise the specific error
func (c *Client) GetDetail(taskId string) (*DtsTaskMeta, error) {
	result := &DtsTaskMeta{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getDtsUriWithTaskId(taskId)).
		WithResult(result).
		Do()

	return result, err
}

// ListDts - list all dtsTask with the specific type
//
// PARAMS:
//     - args: the arguments to list all dtsTask with the specific type
// RETURNS:
//     - *ListDtsResult: the result of list all dtsTask, contains all dtsTask' detail
//     - error: nil if success otherwise the specific error
func (c *Client) ListDts(args *ListDtsArgs) (*ListDtsResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.Type == ""{
		return nil, fmt.Errorf("unset type")
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListDtsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDtsUri()+"/list").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// PreCheck - precheck a dtsTask
//
// PARAMS:
//     - taskId: the specific dtsTask's ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) PreCheck(taskId string) (*PreCheckResult, error) {
	result:=&PreCheckResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDtsApi()+taskId+"/precheck").
		WithResult(result).
		Do()
	return result, err
}

// GetPreCheck - get a precheck result
//
// PARAMS:
//     - taskId: the specific dtsTask's ID
// RETURNS:
//     - *GetPreCheckResult: the specific dtsTask's precheck result
//     - error: nil if success otherwise the specific error
func (c *Client) GetPreCheck(taskId string) (*GetPreCheckResult, error) {
	result:=&GetPreCheckResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDtsApi()+taskId+"/get_precheck").
		WithResult(result).
		Do()

	return result, err
}

// ConfigDts - config a dtsTask with the specific parameters
//
// PARAMS:
//     - taskId: the specific dtsTask's ID
//     - args: the arguments to config a dtsTask
// RETURNS:
//     - *ConfigDtsResult: the result of config dtsTask, contains the dtsTask's ID
//     - error: nil if success otherwise the specific error
func (c *Client) ConfigDts(taskId string, args *ConfigArgs) (*ConfigDtsResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.TaskName == "" {
		return nil, fmt.Errorf("unset TaskName")
	}

	if args.DataType == nil {
		return nil, fmt.Errorf("unset DataType")
	}

	if args.SchemaMapping == nil {
		return nil, fmt.Errorf("unset SchemaMapping")
	}

	result := &ConfigDtsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDtsUriWithTaskId(taskId)+"/config").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// StartDts - start a dtsTask
//
// PARAMS:
//     - taskId: the specific dtsTask's ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) StartDts(taskId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDtsUriWithTaskId(taskId)+"/start").
		Do()
}

// PauseDts - pause a dtsTask
//
// PARAMS:
//     - taskId: the specific dtsTask's ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) PauseDts(taskId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDtsUriWithTaskId(taskId)+"/pause").
		Do()
}

// ShutdownDts - shutdown a dtsTask
//
// PARAMS:
//     - taskId: the specific dtsTask's ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ShutdownDts(taskId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getDtsUriWithTaskId(taskId)+"/shutdown").
		Do()
}
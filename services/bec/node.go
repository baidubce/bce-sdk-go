/*
 * Copyright 2021 Baidu, Inc.
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

// client.go - define the client for BEC service

// Package bec defines the BEC services of BCE. The supported APIs are all defined in sub-package

package bec

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// GetBecAvailableNodeInfoVo -  get available node
//
// PARAMS:
//     - args: the type
// RETURNS:
//     - *api.GetBecAvailableNodeInfoVoResult: get available node
//     - error: nil if ok otherwise the specific error
func (c *Client) GetBecAvailableNodeInfoVo(getType string) (*api.GetBecAvailableNodeInfoVoResult, error) {
	if getType == "" {
		return nil, fmt.Errorf("please set argments")
	}

	result := &api.GetBecAvailableNodeInfoVoResult{}
	req := &api.GetHttpReq{Url: api.GetNodeInfoURI() + "/type/" + getType, Result: result}
	err := api.Get(c, req)

	return result, err
}

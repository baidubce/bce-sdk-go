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

// cfs.go - the Normal CFS APIs definition supported by the CFS service

package cfs

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateFS - create a FS Instance
//
// PARAMS:
//   - args: parameters to create FS
//
// RETURNS:
//   - *CreateFSResult: the result of create fs, contains new FS Instance's ID
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateFS(args *CreateFSArgs) (*CreateFSResult, error) {
	if args == nil || len(args.Name) == 0 {
		return nil, fmt.Errorf("unset fs name")
	}

	if len(args.Zone) == 0 {
		return nil, fmt.Errorf("unset zone")
	}

	result := &CreateFSResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCFSUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpdateFS - update name of a FS Instance
//
// PARAMS:
//   - args: parameters to create FS
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateFS(args *UpdateFSArgs) error {
	if args == nil || len(args.FSName) == 0 {
		return fmt.Errorf("unset fs name")
	}
	if len(args.FSID) == 0 {
		return fmt.Errorf("unset fs id")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getFSInstanceUri(args.FSID)).
		WithBody(args).
		Do()

	return err
}

// DescribeFS - describe all FS Instances
//
// PARAMS:
//   - args: parameters describe all FS Instances
//
// RETURNS:
//   - *DescribeFSResult: the result FS Instances's detail
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeFS(args *DescribeFSArgs) (*DescribeFSResult, error) {
	if args == nil {
		args = &DescribeFSArgs{}
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeFSResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getCFSUri()).
		WithQueryParamFilter("fsId", args.FSID).
		WithQueryParamFilter("userId", args.UserId).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	err := request.Do()
	return result, err
}

// CreateMountTarget - create a mount target for FS Instances
//
// PARAMS:
//   - args: parameters to create mount target
//
// RETURNS:
//   - *CreateMountTargetResult: the result mount target's detail
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateMountTarget(args *CreateMountTargetArgs) (*CreateMountTargetResult, error) {
	if args == nil || len(args.FSID) == 0 {
		return nil, fmt.Errorf("unset fs id")
	}

	if len(args.SubnetId) == 0 {
		return nil, fmt.Errorf("unset subnetid")
	}

	result := &CreateMountTargetResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getFSInstanceUri(args.FSID)).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// DescribeMountTarget - describe all mount targets
//
// PARAMS:
//   - args: parameters describe all mount targets
//
// RETURNS:
//   - *DescribeMountTargetResult: the result Mount target's detail
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeMountTarget(args *DescribeMountTargetArgs) (*DescribeMountTargetResult, error) {
	if args == nil {
		args = &DescribeMountTargetArgs{}
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeMountTargetResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getFSInstanceUri(args.FSID)).
		WithQueryParamFilter("mountId", args.MountID).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	err := request.Do()
	return result, err
}

// DropMountTarget - drop a MountTarget
//
// PARAMS:
//   - args: parameters to drop  mount target
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DropMountTarget(args *DropMountTargetArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getMountTargetUri(args.FSID, args.MountId)).
		Do()
}

// DropFS - drop a fs instance
//
// PARAMS:
//   - args: parameters to drop fs
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DropFS(args *DropFSArgs) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getFSInstanceUri(args.FSID)).
		Do()
}

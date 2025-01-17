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

// client.go - define the client for aihc inference service

// Package inference defines aihc inference services of BCE. The supported APIs are all defined in sub-package

package inference

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/aihc/client"
	"github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
)

const DEFAULT_SERVICE_DOMAIN = "aihc.baidubce.com"

type Client struct {
	client.Client
}

// NewClient make the aihc inference service client with default configuration.
func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_SERVICE_DOMAIN
	}
	aihcClient, err := client.NewClient(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	newClient := Client{*aihcClient}
	return &newClient, nil
}

func (c *Client) SetBceClient(client *bce.BceClient) {
	c.DefaultClient = client
}

func (c *Client) GetBceClient() *bce.BceClient {
	return c.DefaultClient
}

// NewClientWithSTS make the aihc inference service client with STS configuration.
func NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_SERVICE_DOMAIN
	}
	aihcClient, err := client.NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint)
	if err != nil {
		return nil, err
	}
	newClient := Client{*aihcClient}
	return &newClient, nil
}

// CreateApp - create app with the specific parameters
//
// PARAMS:
//   - args: the arguments to create app
//   - region: the region of service
//   - extraInfo: the extra internal args, users can ignore it, set nil
//
// RETURNS:
//   - *api.CreateAppResult: the result of create app, contains new app id
//   - error: nil if success otherwise the specific error
func (c *Client) CreateApp(args *api.CreateAppArgs, region string, extraInfo map[string]string) (*api.CreateAppResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.CreateApp(c.DefaultClient, region, body, extraInfo)
}

// ListApp - list app with the specific parameters
//
// PARAMS:
//   - args: the arguments to list app
//   - region: the region of service
//   - extraInfo: the extra internal args, users can ignore it, set nil
//
// RETURNS:
//   - *api.ListAppResult: the result of list app, contains apps info
//   - error: nil if success otherwise the specific error
func (c *Client) ListApp(args *api.ListAppArgs, region string, extraInfo map[string]string) (*api.ListAppResult, error) {
	return api.ListApp(c.DefaultClient, region, args, extraInfo)
}

// ListAppStats - list app stats with the specific parameters
//
// PARAMS:
//   - args: the arguments to list app
//   - region: the region of service
//
// RETURNS:
//   - *api.ListAppStatsResult: the result of list app stats, contains apps status
//   - error: nil if success otherwise the specific error
func (c *Client) ListAppStats(args *api.ListAppStatsArgs, region string) (*api.ListAppStatsResult, error) {
	return api.ListAppStats(c.DefaultClient, region, args)
}

// AppDetails - the details of app with the specific parameters
//
// PARAMS:
//   - args: the arguments to app details
//   - region: the region of service
//
// RETURNS:
//   - *api.AppDetailsResult: the result of app details, contains apps status
//   - error: nil if success otherwise the specific error
func (c *Client) AppDetails(args *api.AppDetailsArgs, region string) (*api.AppDetailsResult, error) {
	return api.AppDetails(c.DefaultClient, region, args)
}

// UpdateApp - update app with the specific parameters
//
// PARAMS:
//   - args: the arguments to app details
//   - region: the region of service
//
// RETURNS:
//   - *api.UpdateAppResult: the result of update app, contains app id
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateApp(args *api.UpdateAppArgs, region string) (*api.UpdateAppResult, error) {
	jsonBytes, jsonErr := json.Marshal(args.AppConf)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.UpdateApp(c.DefaultClient, region, body, args)
}

// ScaleApp - scale app with the specific parameters
//
// PARAMS:
//   - args: the arguments to app details
//   - region: the region of service
//
// RETURNS:
//   - *api.ScaleAppResult: the result of scale app
//   - error: nil if success otherwise the specific error
func (c *Client) ScaleApp(args *api.ScaleAppArgs, region string) (*api.ScaleAppResult, error) {
	return api.ScaleApp(c.DefaultClient, region, args)
}

// PubAccess - operate public access of app with the specific parameters
//
// PARAMS:
//   - args: the arguments to operate public access
//   - region: the region of service
//
// RETURNS:
//   - *api.PubAccessResult: the result of operating public access of app
//   - error: nil if success otherwise the specific error
func (c *Client) PubAccess(args *api.PubAccessArgs, region string) (*api.PubAccessResult, error) {
	return api.PubAccess(c.DefaultClient, region, args)
}

// ListChange - list app change with the specific parameters
//
// PARAMS:
//   - args: the arguments to list app change
//   - region: the region of service
//
// RETURNS:
//   - *api.ListChangeResult: the result of list app change, contains app change info
//   - error: nil if success otherwise the specific error
func (c *Client) ListChange(args *api.ListChangeArgs, region string) (*api.ListChangeResult, error) {
	return api.ListChange(c.DefaultClient, region, args)
}

// ChangeDetail - app change detail with the specific parameters
//
// PARAMS:
//   - args: the arguments to app change details
//   - region: the region of service
//
// RETURNS:
//   - *api.ChangeDetailResult: the result of app change detail, contains app change detail info
//   - error: nil if success otherwise the specific error
func (c *Client) ChangeDetail(args *api.ChangeDetailArgs, region string) (*api.ChangeDetailResult, error) {
	return api.ChangeDetail(c.DefaultClient, region, args)
}

// DeleteApp - delete app with the specific parameters
//
// PARAMS:
//   - args: the arguments to delete app
//   - region: the region of service
//
// RETURNS:
//   - *api.DeleteAppResult: the result of delete app
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteApp(args *api.DeleteAppArgs, region string) (*api.DeleteAppResult, error) {
	return api.DeleteApp(c.DefaultClient, region, args)
}

// ListPod - list pod of app with the specific parameters
//
// PARAMS:
//   - args: the arguments to list pod of app
//   - region: the region of service
//
// RETURNS:
//   - *api.ListPodResult: the result of list pod of app, contains pod info
//   - error: nil if success otherwise the specific error
func (c *Client) ListPod(args *api.ListPodArgs, region string) (*api.ListPodResult, error) {
	return api.ListPod(c.DefaultClient, region, args)
}

// BlockPod - block pod with the specific parameters
//
// PARAMS:
//   - args: the arguments to block pod
//   - region: the region of service
//
// RETURNS:
//   - *api.BlockPodResult: the result of block pod
//   - error: nil if success otherwise the specific error
func (c *Client) BlockPod(args *api.BlockPodArgs, region string) (*api.BlockPodResult, error) {
	return api.BlockPod(c.DefaultClient, region, args)
}

// DeletePod - delete pod with the specific parameters
//
// PARAMS:
//   - args: the arguments to delete pod
//   - region: the region of service
//
// RETURNS:
//   - *api.DeletePodResult: the result of delete pod
//   - error: nil if success otherwise the specific error
func (c *Client) DeletePod(args *api.DeletePodArgs, region string) (*api.DeletePodResult, error) {
	return api.DeletePod(c.DefaultClient, region, args)
}

// ListBriefResPool - list res pool brief info with the specific parameters
//
// PARAMS:
//   - args: the arguments to list res pool brief info
//   - region: the region of service
//
// RETURNS:
//   - *api.ListBriefResPoolResult: the result of list res pool brief, contains res pool brief info
//   - error: nil if success otherwise the specific error
func (c *Client) ListBriefResPool(args *api.ListBriefResPoolArgs, region string) (*api.ListBriefResPoolResult, error) {
	return api.ListBriefResPool(c.DefaultClient, region, args)
}

// ResPoolDetail - res pool detail with the specific parameters
//
// PARAMS:
//   - args: the arguments to res pool detail
//   - region: the region of service
//
// RETURNS:
//   - *api.ResPoolDetailResult: the result of res pool detail, contains res pool detail info
//   - error: nil if success otherwise the specific error
func (c *Client) ResPoolDetail(args *api.ResPoolDetailArgs, region string) (*api.ResPoolDetailResult, error) {
	return api.ResPoolDetail(c.DefaultClient, region, args)
}

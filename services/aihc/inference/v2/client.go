/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by servicelicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// client.go - define the client for aihc inference service

// Package inference defines aihc inference services of BCE. The supported APIs are all defined in sub-package

package api

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/aihc/client"
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

// CreateService - create service with the specific parameters
//
// PARAMS:
//   - args: the arguments to create service
//   - clientToken: the client token for request, can be empty
//
// RETURNS:
//   - *CreateServiceResult: the result of create service, contains new service id
//   - error: nil if success otherwise the specific error
func (c *Client) CreateService(args *ServiceConf, clientToken string) (*CreateServiceResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return CreateService(c.DefaultClient, body, clientToken)
}

// ListService - list service with the specific parameters
//
// PARAMS:
//   - args: the arguments to list service
//
// RETURNS:
//   - *ListServiceResult: the result of list service, contains services info
//   - error: nil if success otherwise the specific error
func (c *Client) ListService(args *ListServiceArgs) (*ListServiceResult, error) {
	return ListService(c.DefaultClient, args)
}

// ListServiceStats - list service stats with the specific parameters
//
// PARAMS:
//   - args: the arguments to list service
//
// RETURNS:
//   - *ListServiceStatsResult: the result of list service stats, contains services status
//   - error: nil if success otherwise the specific error
func (c *Client) ListServiceStats(args *ListServiceStatsArgs) (*ListServiceStatsResult, error) {
	return ListServiceStats(c.DefaultClient, args)
}

// ServiceDetails - the details of service with the specific parameters
//
// PARAMS:
//   - args: the arguments to service details
//
// RETURNS:
//   - *ServiceDetailsResult: the result of service details, contains services status
//   - error: nil if success otherwise the specific error
func (c *Client) ServiceDetails(args *ServiceDetailsArgs) (*ServiceDetailsResult, error) {
	return ServiceDetails(c.DefaultClient, args)
}

// UpdateService - update service with the specific parameters
//
// PARAMS:
//   - args: the arguments to service details
//
// RETURNS:
//   - *UpdateServiceResult: the result of update service, contains service id
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateService(args *UpdateServiceArgs) (*UpdateServiceResult, error) {
	jsonBytes, jsonErr := json.Marshal(args.ServiceConf)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return UpdateService(c.DefaultClient, body, args)
}

// ScaleService - scale service with the specific parameters
//
// PARAMS:
//   - args: the arguments to service details
//
// RETURNS:
//   - *ScaleServiceResult: the result of scale service
//   - error: nil if success otherwise the specific error
func (c *Client) ScaleService(args *ScaleServiceArgs) (*ScaleServiceResult, error) {
	return ScaleService(c.DefaultClient, args)
}

// PubAccess - operate public access of service with the specific parameters
//
// PARAMS:
//   - args: the arguments to operate public access
//
// RETURNS:
//   - *PubAccessResult: the result of operating public access of service
//   - error: nil if success otherwise the specific error
func (c *Client) PubAccess(args *PubAccessArgs) (*PubAccessResult, error) {
	return PubAccess(c.DefaultClient, args)
}

// ListChange - list service change with the specific parameters
//
// PARAMS:
//   - args: the arguments to list service change
//
// RETURNS:
//   - *ListChangeResult: the result of list service change, contains service change info
//   - error: nil if success otherwise the specific error
func (c *Client) ListChange(args *ListChangeArgs) (*ListChangeResult, error) {
	return ListChange(c.DefaultClient, args)
}

// ChangeDetail - service change detail with the specific parameters
//
// PARAMS:
//   - args: the arguments to service change details
//
// RETURNS:
//   - *ChangeDetailResult: the result of service change detail, contains service change detail info
//   - error: nil if success otherwise the specific error
func (c *Client) ChangeDetail(args *ChangeDetailArgs) (*ChangeDetailResult, error) {
	return ChangeDetail(c.DefaultClient, args)
}

// DeleteService - delete service with the specific parameters
//
// PARAMS:
//   - args: the arguments to delete service
//
// RETURNS:
//   - *DeleteServiceResult: the result of delete service
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteService(args *DeleteServiceArgs) (*DeleteServiceResult, error) {
	return DeleteService(c.DefaultClient, args)
}

// ListPod - list pod of service with the specific parameters
//
// PARAMS:
//   - args: the arguments to list pod of service
//
// RETURNS:
//   - *ListPodResult: the result of list pod of service, contains pod info
//   - error: nil if success otherwise the specific error
func (c *Client) ListPod(args *ListPodArgs) (*ListPodResult, error) {
	return ListPod(c.DefaultClient, args)
}

// BlockPod - block pod with the specific parameters
//
// PARAMS:
//   - args: the arguments to block pod
//
// RETURNS:
//   - *BlockPodResult: the result of block pod
//   - error: nil if success otherwise the specific error
func (c *Client) BlockPod(args *BlockPodArgs) (*BlockPodResult, error) {
	return BlockPod(c.DefaultClient, args)
}

// DeletePod - delete pod with the specific parameters
//
// PARAMS:
//   - args: the arguments to delete pod
//
// RETURNS:
//   - *DeletePodResult: the result of delete pod
//   - error: nil if success otherwise the specific error
func (c *Client) DeletePod(args *DeletePodArgs) (*DeletePodResult, error) {
	return DeletePod(c.DefaultClient, args)
}

// ListPodGroups - list podGroups info with the specific parameters
//
// PARAMS:
//   - args: the arguments to list podGroups
//
// RETURNS:
//   - *ListPodGroupsResult: the result of list podGroups
//   - error: nil if success otherwise the specific error
func (c *Client) ListPodGroups(args *ListPodGroupsArgs) (*ListPodGroupsResult, error) {
	return ListPodGroups(c.DefaultClient, args)
}

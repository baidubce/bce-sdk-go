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

// client.go - define the client for BBC service

// Package bbc defines the BBC services of BCE. The supported APIs are all defined in sub-package
package bbc

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bbc/api"
)

const DEFAULT_SERVICE_DOMAIN = "bbc." + bce.DEFAULT_REGION + ".baidubce.com"

// Client of BBC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the BBC service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endPoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	if endPoint == "" {
		endPoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endPoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

// CreateInstance - create an instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to create instance
// RETURNS:
//     - *api.CreateInstanceResult: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(args *api.CreateInstanceArgs) (*api.CreateInstanceResult, error) {
	if len(args.AdminPass) > 0 {
		cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
		if err != nil {
			return nil, err
		}

		args.AdminPass = cryptedPass
	}

	if args.RootDiskSizeInGb <= 0 {
		args.RootDiskSizeInGb = 20
	}

	if args.PurchaseCount < 1 {
		args.PurchaseCount = 1
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return api.CreateInstance(c, args.ClientToken, body)
}

// ListInstances - list all instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all instance
// RETURNS:
//     - *api.ListInstanceResult: the result of list Instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListInstances(args *api.ListInstancesArgs) (*api.ListInstancesResult, error) {
	return api.ListInstances(c, args)
}

// GetInstanceDetail - get a specific instance detail info
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - *api.GetInstanceDetailResult: the result of get instance detail info
//     - error: nil if success otherwise the specific error
func (c *Client) GetInstanceDetail(instanceId string) (*api.InstanceModel, error) {
	return api.GetInstanceDetail(c, instanceId)
}

// StartInstance - start an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) StartInstance(instanceId string) error {
	return api.StartInstance(c, instanceId)
}

// StopInstance - stop an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - forceStop: choose to force stop an instance or not
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) StopInstance(instanceId string, forceStop bool) error {
	args := &api.StopInstanceArgs{
		ForceStop: forceStop,
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.StopInstance(c, instanceId, body)
}

// RebootInstance - restart an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - forceStop: choose to force stop an instance or not
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RebootInstance(instanceId string, forceStop bool) error {
	args := &api.StopInstanceArgs{
		ForceStop: forceStop,
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.RebootInstance(c, instanceId, body)
}

// ModifyInstanceName - modify an instance's name
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments of now instance's name
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceName(instanceId string, args *api.ModifyInstanceNameArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ModifyInstanceName(c, instanceId, body)
}

// ModifyInstanceDesc - modify an instance's description
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments of now instance's description
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceDesc(instanceId string, args *api.ModifyInstanceDescArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ModifyInstanceDesc(c, instanceId, body)
}

// RebuildInstance - rebuild an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// 	   - isPreserveData: choose to preserve data or not
//     - args: the arguments to rebuild an instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RebuildInstance(instanceId string, isPreserveData bool, args *api.RebuildInstanceArgs) error {
	cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
	if err != nil {
		return err
	}
	args.AdminPass = cryptedPass

	args.IsPreserveData = isPreserveData

	if !isPreserveData && args.SysRootSize <= 0 {
		args.SysRootSize = 20
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.RebuildInstance(c, instanceId, body)
}

// DeleteInstance - delete an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(instanceId string) error {
	return api.DeleteInstance(c, instanceId)
}

// ModifyInstancePassword - modify an instance's password
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments of now instance's password
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyInstancePassword(instanceId string, args *api.ModifyInstancePasswordArgs) error {
	cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
	if err != nil {
		return err
	}
	args.AdminPass = cryptedPass

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ModifyInstancePassword(c, instanceId, body)
}

// GetVpcSubnet - get multi instances vpc and subnet
//
// PARAMS:
//      - args: the instanceId of bbc instances
// RETURNS:
// 	   - *api.GetVpcSubnetResult: result of vpc and subnet
//     - error: nil if success otherwise the specific error
func (c *Client) GetVpcSubnet(args *api.GetVpcSubnetArgs) (*api.GetVpcSubnetResult, error) {

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return api.GetVpcSubnet(c, body)
}

// UnbindTags - unbind an instance tags
//
// PARAMS:
//     - instanceId: the id of the instance
//     - args: tags of an instance to unbind
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnbindTags(instanceId string, args *api.UnbindTagsArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return api.UnbindTags(c, instanceId, body)
}

// ListFlavors - list all available flavors
//
// RETURNS:
//     - *api.ListFlavorsResult: the result of list all flavors
//     - error: nil if success otherwise the specific error
func (c *Client) ListFlavors() (*api.ListFlavorsResult, error) {
	return api.ListFlavors(c)
}

// GetFlavorDetail - get details of the specified flavor
//
// PARAMS:
//     - flavorId: the id of the flavor
// RETURNS:
//     - *api.GetFlavorDetailResult: the detail of the specified flavor
//     - error: nil if success otherwise the specific error
func (c *Client) GetFlavorDetail(flavorId string) (*api.GetFlavorDetailResult, error) {
	return api.GetFlavorDetail(c, flavorId)
}

// GetFlavorRaid - get the RAID detail and disk size of the specified flavor
//
// PARAMS:
//     - flavorId: the id of the flavor
// RETURNS:
//     - *api.GetFlavorRaidResult: the detail of the raid of the specified flavor
//     - error: nil if success otherwise the specific error
func (c *Client) GetFlavorRaid(flavorId string) (*api.GetFlavorRaidResult, error) {
	return api.GetFlavorRaid(c, flavorId)
}

// CreateImageFromInstanceId - create image from specified instance
//
// PARAMS:
//     - args: the arguments to create image
// RETURNS:
//     - *api.CreateImageResult: the result of create Image
//     - error: nil if success otherwise the specific error
func (c *Client) CreateImageFromInstanceId(args *api.CreateImageArgs) (*api.CreateImageResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return api.CreateImageFromInstanceId(c, args.ClientToken, body)
}

//ListImage - list all images
//
// PARAMS:
//     - args: the arguments to list all images
// RETURNS:
//     - *api.ListImageResult: the result of list all images
//     - error: nil if success otherwise the specific error
func (c *Client) ListImage(args *api.ListImageArgs) (*api.ListImageResult, error) {
	return api.ListImage(c, args)
}

// GetImageDetail - get an image's detail info
//
// PARAMS:
//     - imageId: the specific image ID
// RETURNS:
//     - *api.GetImageDetailResult: the result of get image's detail
//     - error: nil if success otherwise the specific error
func (c *Client) GetImageDetail(imageId string) (*api.GetImageDetailResult, error) {
	return api.GetImageDetail(c, imageId)
}

// DeleteImage - delete an image
//
// PARAMS:
//     - imageId: the specific image ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteImage(imageId string) error {
	return api.DeleteImage(c, imageId)
}

// GetOperationLog - get operation log
//
// PARAMS:
//     - args: the arguments to get operation log
// RETURNS:
//     - *api.GetOperationLogResult: results of getting operation log
//     - error: nil if success otherwise the specific error
func (c *Client) GetOperationLog(args *api.GetOperationLogArgs) (*api.GetOperationLogResult, error) {
	return api.GetOperationLog(c, args)
}

// CreateDeploySet - create a deploy set
//
// PARAMS:
//     - args: the arguments to create a deploy set
// RETURNS:
//     - *api.CreateDeploySetResult: results of creating a deploy set
//     - error: nil if success otherwise the specific error
func (c *Client) CreateDeploySet(args *api.CreateDeploySetArgs) (*api.CreateDeploySetResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return api.CreateDeploySet(c, args.ClientToken, body)
}

// ListDeploySets - list all deploy sets
//
// RETURNS:
//     - *api.ListDeploySetsResult: the result of list all deploy sets
//     - error: nil if success otherwise the specific error
func (c *Client) ListDeploySets() (*api.ListDeploySetsResult, error) {
	return api.ListDeploySets(c)
}

// GetDeploySet - get details of the deploy set
//
// PARAMS:
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - *api.GetDeploySetResult: the detail of the deploy set
//     - error: nil if success otherwise the specific error
func (c *Client) GetDeploySet(deploySetId string) (*api.GetDeploySetResult, error) {
	return api.GetDeploySet(c, deploySetId)
}

// DeleteDeploySet - delete a deploy set
//
// PARAMS:
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteDeploySet(deploySetId string) error {
	return api.DeleteDeploySet(c, deploySetId)
}

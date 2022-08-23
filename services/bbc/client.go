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
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
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
//     - *CreateInstanceResult: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateInstance(args *CreateInstanceArgs) (*CreateInstanceResult, error) {
	if len(args.AdminPass) > 0 {
		cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
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
	return CreateInstance(c, args, body)
}

// CreateInstance - create an instance with specific parameters and support the passing in of label
//
// PARAMS:
//     - args: the arguments to create instance
// RETURNS:
//     - *CreateInstanceResult: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific erro
func (c *Client) CreateInstanceByLabel(args *CreateSpecialInstanceArgs) (*CreateInstanceResult, error) {
	if len(args.AdminPass) > 0 {
		cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
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
	return CreateInstanceByLabel(c, args, body)
}

// ListInstances - list all instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all instance
// RETURNS:
//     - *ListInstanceResult: the result of list Instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListInstances(args *ListInstancesArgs) (*ListInstancesResult, error) {
	return ListInstances(c, args)
}

// GetInstanceDetail - get a specific instance detail info
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - *GetInstanceDetailResult: the result of get instance detail info
//     - error: nil if success otherwise the specific error
func (c *Client) GetInstanceDetail(instanceId string) (*InstanceModel, error) {
	return GetInstanceDetail(c, instanceId)
}

func (c *Client) GetInstanceDetailWithDeploySet(instanceId string, isDeploySet bool) (*InstanceModel, error) {
	return GetInstanceDetailWithDeploySet(c, instanceId, isDeploySet)
}

func (c *Client) GetInstanceDetailWithDeploySetAndFailed(instanceId string,
	isDeploySet bool, containsFailed bool) (*InstanceModel, error) {
	return GetInstanceDetailWithDeploySetAndFailed(c, instanceId, isDeploySet, containsFailed)
}

// StartInstance - start an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) StartInstance(instanceId string) error {
	return StartInstance(c, instanceId)
}

// StopInstance - stop an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - forceStop: choose to force stop an instance or not
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) StopInstance(instanceId string, forceStop bool) error {
	args := &StopInstanceArgs{
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

	return StopInstance(c, instanceId, body)
}

// RebootInstance - restart an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - forceStop: choose to force stop an instance or not
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RebootInstance(instanceId string, forceStop bool) error {
	args := &StopInstanceArgs{
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

	return RebootInstance(c, instanceId, body)
}

// ModifyInstanceName - modify an instance's name
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments of now instance's name
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceName(instanceId string, args *ModifyInstanceNameArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return ModifyInstanceName(c, instanceId, body)
}

// ModifyInstanceDesc - modify an instance's description
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments of now instance's description
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceDesc(instanceId string, args *ModifyInstanceDescArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	clientToken := args.ClientToken

	return ModifyInstanceDesc(c, instanceId, clientToken, body)
}

// RebuildInstance - rebuild an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// 	   - isPreserveData: choose to preserve data or not
//     - args: the arguments to rebuild an instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RebuildInstance(instanceId string, isPreserveData bool, args *RebuildInstanceArgs) error {
	cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
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

	return RebuildInstance(c, instanceId, body)
}

// GetInstanceVNC - get an instance's VNC url
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - *api.GetInstanceVNCResult: the result of get instance's VNC url
//     - error: nil if success otherwise the specific error
func (c *Client) GetInstanceVNC(instanceId string) (*GetInstanceVNCResult, error) {
	return GetInstanceVNC(c, instanceId)
}

// RebuildBatchInstance - batch rebuild instances
//
// PARAMS:
//     - args: the arguments to batch rebuild instances
// RETURNS:
//     - *BatchRebuildResponse: the result of batch rebuild the instances
//     - error: nil if success otherwise the specific error
func (c *Client) BatchRebuildInstances(args *RebuildBatchInstanceArgs) (*BatchRebuildResponse, error) {
	encryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
	if err != nil {
		return nil, err
	}
	args.AdminPass = encryptedPass

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return BatchRebuildInstances(c, body)
}

// InstancePurchaseReserved - purchase reserve an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments to purchase reserved an instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) InstancePurchaseReserved(instanceId string, args *PurchaseReservedArgs) error {
	// this api only support Prepaid instance
	args.Billing.PaymentTiming = PaymentTimingPrePaid

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return InstancePurchaseReserved(c, instanceId, args.ClientToken, body)
}

// DeleteInstance - delete an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(instanceId string) error {
	return DeleteInstance(c, instanceId)
}

// ListInstances - list all instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all instance
// RETURNS:
//     - *ListInstanceResult: the result of list Instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListRecycledInstances(args *ListRecycledInstancesArgs) (*ListRecycledInstancesResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return ListRecycledInstances(c, body)
}

// ListInstances - list all instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all instance
// RETURNS:
//     - *ListInstanceResult: the result of list Instance
//     - error: nil if success otherwise the specific error
func (c *Client) RecoveryInstances(args *RecoveryInstancesArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return RecoveryInstances(c, body)
}

// DeleteInstance - delete an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteInstances(args *DeleteInstanceArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return DeleteInstances(c, body)
}

// ModifyInstancePassword - modify an instance's password
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments of now instance's password
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyInstancePassword(instanceId string, args *ModifyInstancePasswordArgs) error {
	cryptedPass, err := Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
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

	return ModifyInstancePassword(c, instanceId, body)
}

// InstanceChangeSubnet - change an instance's subnet
//
// PARAMS:
//     - args: the arguments to change an instance's subnet
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) InstanceChangeSubnet(args *InstanceChangeSubnetArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return InstanceChangeSubnet(c, body)
}

// InstanceChangeVpc - change an instance's vpc
//
// PARAMS:
//     - args: the arguments to change an instance's vpc
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) InstanceChangeVpc(args *InstanceChangeVpcArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return InstanceChangeVpc(c, body)
}

// GetVpcSubnet - get multi instances vpc and subnet
//
// PARAMS:
//      - args: the instanceId of bbc instances
// RETURNS:
// 	   - *GetVpcSubnetResult: result of vpc and subnet
//     - error: nil if success otherwise the specific error
func (c *Client) GetVpcSubnet(args *GetVpcSubnetArgs) (*GetVpcSubnetResult, error) {

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return GetVpcSubnet(c, body)
}

// BatchAddIP - Add ips to instance
//
// PARAMS:
//      - args: the arguments to add ips to bbc instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchAddIP(args *BatchAddIpArgs) (*BatchAddIpResponse, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return BatchAddIp(c, args, body)
}

// BatchDelIP - Delete ips of instance
//
// PARAMS:
//      - args: the arguments to add ips to bbc instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchDelIP(args *BatchDelIpArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return BatchDelIp(c, args, body)
}

// BatchAddIPCrossSubnet - Add ips to instance cross subnet
//
// PARAMS:
//      - args: the arguments to add ips to bbc instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchAddIPCrossSubnet(args *BatchAddIpCrossSubnetArgs) (*BatchAddIpResponse, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return BatchAddIpCrossSubnet(c, args, body)
}

// BindTags - bind an instance tags
//
// PARAMS:
//     - instanceId: the id of the instance
//     - args: tags of an instance to bind
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindTags(instanceId string, args *BindTagsArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return BindTags(c, instanceId, body)
}

// UnbindTags - unbind an instance tags
//
// PARAMS:
//     - instanceId: the id of the instance
//     - args: tags of an instance to unbind
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnbindTags(instanceId string, args *UnbindTagsArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return UnbindTags(c, instanceId, body)
}

// ListFlavors - list all available flavors
//
// RETURNS:
//     - *ListFlavorsResult: the result of list all flavors
//     - error: nil if success otherwise the specific error
func (c *Client) ListFlavors() (*ListFlavorsResult, error) {
	return ListFlavors(c)
}

// GetFlavorDetail - get details of the specified flavor
//
// PARAMS:
//     - flavorId: the id of the flavor
// RETURNS:
//     - *GetFlavorDetailResult: the detail of the specified flavor
//     - error: nil if success otherwise the specific error
func (c *Client) GetFlavorDetail(flavorId string) (*GetFlavorDetailResult, error) {
	return GetFlavorDetail(c, flavorId)
}

// GetFlavorRaid - get the RAID detail and disk size of the specified flavor
//
// PARAMS:
//     - flavorId: the id of the flavor
// RETURNS:
//     - *GetFlavorRaidResult: the detail of the raid of the specified flavor
//     - error: nil if success otherwise the specific error
func (c *Client) GetFlavorRaid(flavorId string) (*GetFlavorRaidResult, error) {
	return GetFlavorRaid(c, flavorId)
}

// CreateImageFromInstanceId - create image from specified instance
//
// PARAMS:
//     - args: the arguments to create image
// RETURNS:
//     - *CreateImageResult: the result of create Image
//     - error: nil if success otherwise the specific error
func (c *Client) CreateImageFromInstanceId(args *CreateImageArgs) (*CreateImageResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return CreateImageFromInstanceId(c, args.ClientToken, body)
}

//ListImage - list all images
//
// PARAMS:
//     - args: the arguments to list all images
// RETURNS:
//     - *ListImageResult: the result of list all images
//     - error: nil if success otherwise the specific error
func (c *Client) ListImage(args *ListImageArgs) (*ListImageResult, error) {
	return ListImage(c, args)
}

// GetImageDetail - get an image's detail info
//
// PARAMS:
//     - imageId: the specific image ID
// RETURNS:
//     - *GetImageDetailResult: the result of get image's detail
//     - error: nil if success otherwise the specific error
func (c *Client) GetImageDetail(imageId string) (*GetImageDetailResult, error) {
	return GetImageDetail(c, imageId)
}

// DeleteImage - delete an image
//
// PARAMS:
//     - imageId: the specific image ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteImage(imageId string) error {
	return DeleteImage(c, imageId)
}

// GetOperationLog - get operation log
//
// PARAMS:
//     - args: the arguments to get operation log
// RETURNS:
//     - *GetOperationLogResult: results of getting operation log
//     - error: nil if success otherwise the specific error
func (c *Client) GetOperationLog(args *GetOperationLogArgs) (*GetOperationLogResult, error) {
	return GetOperationLog(c, args)
}

// CreateDeploySet - create a deploy set
//
// PARAMS:
//     - args: the arguments to create a deploy set
// RETURNS:
//     - *CreateDeploySetResult: results of creating a deploy set
//     - error: nil if success otherwise the specific error
func (c *Client) CreateDeploySet(args *CreateDeploySetArgs) (*CreateDeploySetResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return CreateDeploySet(c, args.ClientToken, body)
}

// ListDeploySets - list all deploy sets
//
// RETURNS:
//     - *ListDeploySetsResult: the result of list all deploy sets
//     - error: nil if success otherwise the specific error
func (c *Client) ListDeploySets() (*ListDeploySetsResult, error) {
	return ListDeploySets(c)
}

// ListDeploySets - list all deploy sets
// PARAMS:
//     - args: the arguments to filter
// RETURNS:
//     - *ListDeploySetsResult: the result of list all deploy sets
//     - error: nil if success otherwise the specific error
func (c *Client) ListDeploySetsPage(args *ListDeploySetsArgs) (*ListDeploySetsResult, error) {
	return ListDeploySetsPage(c, args)
}

// GetDeploySet - get details of the deploy set
//
// PARAMS:
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - *GetDeploySetResult: the detail of the deploy set
//     - error: nil if success otherwise the specific error
func (c *Client) GetDeploySet(deploySetId string) (*DeploySetResult, error) {
	return GetDeploySet(c, deploySetId)
}

// DeleteDeploySet - delete a deploy set
//
// PARAMS:
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteDeploySet(deploySetId string) error {
	return DeleteDeploySet(c, deploySetId)
}

// BindSecurityGroups - Bind Security Groups
//
// PARAMS:
//     - args: the arguments of bind security groups
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindSecurityGroups(args *BindSecurityGroupsArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return BindSecurityGroups(c, body)
}

// UnBindSecurityGroups - UnBind Security Groups
//
// PARAMS:
//     - args: the arguments of bind security groups
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnBindSecurityGroups(args *UnBindSecurityGroupsArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return UnBindSecurityGroups(c, body)
}

// ListFlavorZones - get the zone list of the specified flavor which can buy
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - flavorId: the id of the flavor
// RETURNS:
//     - *ListZonesResult: the list of zone names
//     - error: nil if success otherwise the specific error
func (c *Client) ListFlavorZones(args *ListFlavorZonesArgs) (*ListZonesResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return ListFlavorZones(c, body)
}

// ListZoneFlavors - get the flavor detail of the specific zone
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - zoneName: the zone name
// RETURNS:
//     - *ListZoneResult: flavor detail of the specific zone
//     - error: nil if success otherwise the specific error
func (c *Client) ListZoneFlavors(args *ListZoneFlavorsArgs) (*ListFlavorInfosResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return ListZoneFlavors(c, body)
}

// GetCommonImage - get common flavor image list
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - flavorIds: the specific flavorIds, can be nil
// RETURNS:
//     - *GetImageDetailResult: the result of get image's detail
//     - error: nil if success otherwise the specific error
func (c *Client) GetCommonImage(args *GetFlavorImageArgs) (*GetImagesResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return GetCommonImage(c, body)
}

// GetCustomImage - get user onwer flavor image list
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - flavorIds: the specific flavorIds, can be nil
// RETURNS:
//     - *GetImageDetailResult: the result of get image's detail
//     - error: nil if success otherwise the specific error
func (c *Client) GetCustomImage(args *GetFlavorImageArgs) (*GetImagesResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return GetCustomImage(c, body)
}

// ShareImage - share an image
//
// PARAMS:
//     - imageId: the specific image ID
//     - args: the arguments to share an image
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ShareImage(imageId string, args *SharedUser) error {
	return ShareImage(c, imageId, args)
}

// UnShareImage - cancel share an image
//
// PARAMS:
//     - imageId: the specific image ID
//     - args: the arguments to cancel share an image
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnShareImage(imageId string, args *SharedUser) error {
	return UnShareImage(c, imageId, args)
}

// GetImageSharedUser - get user list use this image
//
// PARAMS:
//     - imageId: the specific image ID
// RETURNS:
//     - *api.GetImageSharedUserResult: the result of user list
//     - error: nil if success otherwise the specific error
func (c *Client) GetImageSharedUser(imageId string) (*GetImageSharedUserResult, error) {
	return GetImageSharedUser(c, imageId)
}

// RemoteCopyImage - copy a bbc image from other region
//
// PARAMS:
//     - imageId: the specific image ID
//     - args: the arguments to remote copy an image
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RemoteCopyImage(imageId string, args *RemoteCopyImageArgs) error {
	return RemoteCopyImage(c, imageId, args)
}

// RemoteCopyImageReturnImageIds - copy an image from other region
//
// PARAMS:
//     - imageId: the specific image ID
//     - args: the arguments to remote copy an image
// RETURNS:
//     - imageIds of destination region if success otherwise the specific error
func (c *Client) RemoteCopyImageReturnImageIds(imageId string, args *RemoteCopyImageArgs) (*RemoteCopyImageResult, error) {
	return RemoteCopyImageReturnImageIds(c, imageId, args)
}

// CancelRemoteCopyImage - cancel a copy image from other region operation
//
// PARAMS:
//     - imageId: the specific image ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CancelRemoteCopyImage(imageId string) error {
	return CancelRemoteCopyImage(c, imageId)
}

// GetInstanceEni - get the eni of the bbc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the bbc instance id
// RETURNS:
//     - error: nil if success otherwise the specific error

func (c *Client) GetInstanceEni(instanceId string) (*GetInstanceEniResult, error) {
	return GetInstanceEni(c, instanceId)
}

func (c *Client) GetInstanceCreateStock(args *CreateInstanceStockArgs) (*InstanceStockResult, error) {
	return GetInstanceCreateStock(c, args)
}

func (c *Client) GetSimpleFlavor(args *GetSimpleFlavorArgs) (*SimpleFlavorResult, error) {
	return GetSimpleFlavor(c, args)
}

func (c *Client) GetInstancePirce(args *InstancePirceArgs) (*InstancePirceResult, error) {
	return GetInstancePirce(c, args)
}

func (c *Client) ListRepairTasks(args *ListRepairTaskArgs) (*ListRepairTaskResult, error) {
	return ListRepairTasks(c, args)
}

func (c *Client) ListClosedRepairTasks(args *ListClosedRepairTaskArgs) (*ListClosedRepairTaskResult, error) {
	return ListClosedRepairTasks(c, args)
}

func (c *Client) GetRepairTaskDetail(taskId string) (*GetRepairTaskResult, error) {
	return GetTaskDetail(c, taskId)
}

func (c *Client) AuthorizeRepairTask(args *TaskIdArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return AuthorizeRepairTask(c, body)
}

func (c *Client) UnAuthorizeRepairTask(args *TaskIdArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return UnAuthorizeRepairTask(c, body)
}

func (c *Client) ConfirmRepairTask(args *TaskIdArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return ConfirmRepairTask(c, body)
}

func (c *Client) DisConfirmRepairTask(args *DisconfirmTaskArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return DisConfirmRepairTask(c, body)
}

func (c *Client) GetRepairTaskRecord(args *TaskIdArgs) (*GetRepairRecords, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return GetRepairTaskReocrd(c, body)
}

// ListRule - list the repair plat rules
//
// PARAMS:
//     - args: the arguments of listing the repair plat rules
// RETURNS:
//     - *ListRuleResult: results of listing the repair plat rules
//     - error: nil if success otherwise the specific error
func (c *Client) ListRule(args *ListRuleArgs) (*ListRuleResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return ListRule(c, body)
}

// GetRuleDetail - list the repair plat rules
//
// PARAMS:
//     - ruleId: the specified rule id
// RETURNS:
//     - *Rule: results of the specified repair plat rule
//     - error: nil if success otherwise the specific error
func (c *Client) GetRuleDetail(ruleId string) (*Rule, error) {
	return GetRuleDetail(c, ruleId)
}

// CreateRule - create the repair plat rule
//
// PARAMS:
//     - args: the arguments of creating the repair plat rule
// RETURNS:
//     - *CreateRuleResult: results of the id of the repair plat rule which is created
//     - error: nil if success otherwise the specific error
func (c *Client) CreateRule(args *CreateRuleArgs) (*CreateRuleResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return CreateRule(c, body)
}

// DeleteRule - delete the repair plat rule
//
// PARAMS:
//     - args: the arguments of deleting the repair plat rule
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteRule(args *DeleteRuleArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return DeleteRule(c, body)
}

// DisableRule - disable the repair plat rule
//
// PARAMS:
//     - args: the arguments of disabling the repair plat rule
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DisableRule(args *DisableRuleArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return DisableRule(c, body)
}

// EnableRule - enable the repair plat rule
//
// PARAMS:
//     - args: the arguments of enabling the repair plat rule
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) EnableRule(args *EnableRuleArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return EnableRule(c, body)
}

// BatchCreateAutoRenewRules - Batch Create AutoRenew Rules
//
// PARAMS:
//      - args: the arguments to batch create autorenew rules
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchCreateAutoRenewRules(args *BbcCreateAutoRenewArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return BatchCreateAutoRenewRules(c, body)
}

// BatchDeleteAutoRenewRules - Batch Delete AutoRenew Rules
//
// PARAMS:
//      - args: the arguments to batch delete autorenew rules
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchDeleteAutoRenewRules(args *BbcDeleteAutoRenewArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return BatchDeleteAutoRenewRules(c, body)
}

func (c *Client) DeleteInstanceIngorePayment(args *DeleteInstanceIngorePaymentArgs) (*DeleteInstanceResult, error) {
	return DeleteBbcIngorePayment(c, args)
}

// DeleteRecycledInstance - delete an recycled instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteRecycledInstance(instanceId string) error {
	return DeleteRecycledInstance(c, instanceId)
}

// ListCDSVolume - list all cds volume with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all cds
// RETURNS:
//     - *api.ListCDSVolumeResult: the result of list all CDS volume
//     - error: nil if success otherwise the specific error
func (c *Client) ListCDSVolume(queryArgs *ListCDSVolumeArgs) (*ListCDSVolumeResult, error) {
	return ListCDSVolume(c, queryArgs)
}

// GetBbcStockWithDeploySet - get the bbc's stock with deploySet
//
// RETURNS:
//     - *GetBbcStocksResult: the result of the bbc's stock
//     - error: nil if success otherwise the specific error
func (c *Client) GetBbcStockWithDeploySet(args *GetBbcStockArgs) (*GetBbcStocksResult, error) {
	return GetStockWithDeploySet(c, args)
}

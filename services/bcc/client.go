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

// client.go - define the client for BOS service

// Package bcc defines the BCC services of BCE. The supported APIs are all defined in sub-package

package bcc

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
)

const DEFAULT_SERVICE_DOMAIN = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"

// Client of BCC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the BCC service client with default configuration.
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

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.CreateInstance(c, args, body)
}

// CreateInstance - create an instance with the specific parameters and support the passing in of label
//
// PARAMS:
//     - args: the arguments to create instance
// RETURNS:
//     - *api.CreateInstanceResult: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateInstanceByLabel(args *api.CreateSpecialInstanceBySpecArgs) (*api.CreateInstanceResult, error) {
	if len(args.AdminPass) > 0 {
		cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
		if err != nil {
			return nil, err
		}

		args.AdminPass = cryptedPass
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.CreateInstanceByLabel(c, args, body)
}

// CreateInstanceBySpec - create an instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to create instance
// RETURNS:
//     - *api.CreateInstanceBySpecResult: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateInstanceBySpec(args *api.CreateInstanceBySpecArgs) (*api.CreateInstanceBySpecResult, error) {
	if len(args.AdminPass) > 0 {
		cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
		if err != nil {
			return nil, err
		}

		args.AdminPass = cryptedPass
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.CreateInstanceBySpec(c, args, body)
}

// CreateInstanceV3 - create an instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to create instance
// RETURNS:
//     - *api.CreateInstanceV3Result: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateInstanceV3(args *api.CreateInstanceV3Args) (*api.CreateInstanceV3Result, error) {
	if len(args.Password) > 0 {
		cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.Password)
		if err != nil {
			return nil, err
		}

		args.Password = cryptedPass
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.CreateInstanceV3(c, args, body)
}

// ListInstances - list all instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all instance
// RETURNS:
//     - *api.ListInstanceResult: the result of list Instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListInstances(args *api.ListInstanceArgs) (*api.ListInstanceResult, error) {
	return api.ListInstances(c, args)
}

// ListRecycleInstances - list all instance in the recycle bin with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all instance in the recycle bin
// RETURNS:
//     - *api.ListRecycleInstanceResult: the result of list Instance in the recycle bin
//     - error: nil if success otherwise the specific error
func (c *Client) ListRecycleInstances(args *api.ListRecycleInstanceArgs) (*api.ListRecycleInstanceResult, error) {
	return api.ListRecycleInstances(c, args)
}

// ListServersByMarkerV3 - list all instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all instance
// RETURNS:
//     - *api.LogicMarkerResultResponseV3: the result of list Instance
//     - error: nil if success otherwise the specific error
func (c *Client) ListServersByMarkerV3(args *api.ListServerRequestV3Args) (*api.LogicMarkerResultResponseV3, error) {
	return api.ListServersByMarkerV3(c, args)
}

// GetInstanceDetail - get a specific instance detail info
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - *api.GetInstanceDetailResult: the result of get instance detail info
//     - error: nil if success otherwise the specific error
func (c *Client) GetInstanceDetail(instanceId string) (*api.GetInstanceDetailResult, error) {
	return api.GetInstanceDetail(c, instanceId)
}

func (c *Client) GetInstanceDetailWithDeploySet(instanceId string, isDeploySet bool) (*api.GetInstanceDetailResult,
	error) {
	return api.GetInstanceDetailWithDeploySet(c, instanceId, isDeploySet)
}

func (c *Client) GetInstanceDetailWithDeploySetAndFailed(instanceId string,
	isDeploySet bool, containsFailed bool) (*api.GetInstanceDetailResult,
	error) {
	return api.GetInstanceDetailWithDeploySetAndFailed(c, instanceId, isDeploySet, containsFailed)
}

// DeleteInstance - delete a specific instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(instanceId string) error {
	return api.DeleteInstance(c, instanceId)
}

// AutoReleaseInstance - set releaseTime of a postpay instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - releaseTime: an UTC string，eg:"2021-05-01T07:58:09Z"
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) AutoReleaseInstance(instanceId string, releaseTime string) error {
	args := &api.AutoReleaseArgs{
		ReleaseTime: releaseTime,
	}
	return api.AutoReleaseInstance(c, instanceId, args)
}

// ResizeInstance - resize a specific instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments to resize a specific instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ResizeInstance(instanceId string, args *api.ResizeInstanceArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ResizeInstance(c, instanceId, args.ClientToken, body)
}

// RebuildInstance - rebuild an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments to rebuild an instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RebuildInstance(instanceId string, args *api.RebuildInstanceArgs) error {
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

	return api.RebuildInstance(c, instanceId, body)
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
//     - stopWithNoCharge: choose to stop with nocharge an instance or not
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) StopInstanceWithNoCharge(instanceId string, forceStop bool, stopWithNoCharge bool) error {
	args := &api.StopInstanceArgs{
		ForceStop:        forceStop,
		StopWithNoCharge: stopWithNoCharge,
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

func (c *Client) RecoveryInstance(args *api.RecoveryInstanceArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.RecoveryInstance(c, body)
}

// ChangeInstancePass - change an instance's password
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments to change password
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ChangeInstancePass(instanceId string, args *api.ChangeInstancePassArgs) error {
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

	return api.ChangeInstancePass(c, instanceId, body)
}

// ModifyDeletionProtection - Modify deletion protection of specified instance
//
// PARAMS:
//     - instanceId: id of the instance
//	   - args: the arguments to modify deletion protection, default 0 for deletable and 1 for deletion protection
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyDeletionProtection(instanceId string, args *api.DeletionProtectionArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return api.ModifyDeletionProtection(c, instanceId, body)
}

// ModifyInstanceAttribute - modify an instance's attribute
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments of now instance's attribute
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceAttribute(instanceId string, args *api.ModifyInstanceAttributeArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ModifyInstanceAttribute(c, instanceId, body)
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

// ModifyInstanceHostname - modify an instance's hostname
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments of now instance's hostname
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyInstanceHostname(instanceId string, args *api.ModifyInstanceHostnameArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ModifyInstanceHostname(c, instanceId, body)
}

// BindSecurityGroup - bind a security group to an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - securityGroupId: the security group ID which need to bind
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindSecurityGroup(instanceId string, securityGroupId string) error {
	args := &api.BindSecurityGroupArgs{
		SecurityGroupId: securityGroupId,
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.BindSecurityGroup(c, instanceId, body)
}

// UnBindSecurityGroup - unbind a security group ID from instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - securityGroupId: the security group ID which need to unbind
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnBindSecurityGroup(instanceId string, securityGroupId string) error {
	args := &api.BindSecurityGroupArgs{
		SecurityGroupId: securityGroupId,
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.UnBindSecurityGroup(c, instanceId, body)
}

// GetInstanceVNC - get an instance's VNC url
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - *api.GetInstanceVNCResult: the result of get instance's VNC url
//     - error: nil if success otherwise the specific error
func (c *Client) GetInstanceVNC(instanceId string) (*api.GetInstanceVNCResult, error) {
	return api.GetInstanceVNC(c, instanceId)
}

// InstancePurchaseReserved - purchase reserve an instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments to purchase reserved an instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) InstancePurchaseReserved(instanceId string, args *api.PurchaseReservedArgs) error {
	// this api only support Prepaid instance
	args.Billing.PaymentTiming = api.PaymentTimingPrePaid
	relatedRenewFlag := args.RelatedRenewFlag

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.InstancePurchaseReserved(c, instanceId, relatedRenewFlag, args.ClientToken, body)
}

// GetBidInstancePrice - get the market price of the specified bidding instance
//
// PARAMS:
//      - args: the arguments to get the bidding instance market price
// RETURNS:
//     - *GetBidInstancePriceResult: result of the market price of the specified bidding instance
//     - error: nil if success otherwise the specific error
func (c *Client) GetBidInstancePrice(args *api.GetBidInstancePriceArgs) (*api.GetBidInstancePriceResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return api.GetBidInstancePrice(c, args.ClientToken, body)
}

// ListBidFlavor - list all flavors of the bidding instance
//
// RETURNS:
//     - *ListBidFlavorResult: result of the flavor list
//     - error: nil if success otherwise the specific error
func (c *Client) ListBidFlavor() (*api.ListBidFlavorResult, error) {
	return api.ListBidFlavor(c)
}

// DeleteInstanceWithRelateResource - delete an instance and all eip/cds relate it
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments to delete instance and its relate resource
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteInstanceWithRelateResource(instanceId string, args *api.DeleteInstanceWithRelateResourceArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.DeleteInstanceWithRelatedResource(c, instanceId, body)
}

// InstanceChangeSubnet - change an instance's subnet
//
// PARAMS:
//     - args: the arguments to change an instance's subnet
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) InstanceChangeSubnet(args *api.InstanceChangeSubnetArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.InstanceChangeSubnet(c, body)
}

// InstanceChangeVpc - change an instance's vpc
//
// PARAMS:
//     - args: the arguments to change an instance's vpc
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) InstanceChangeVpc(args *api.InstanceChangeVpcArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.InstanceChangeVpc(c, body)
}

// BatchAddIP - Add ips to instance
//
// PARAMS:
//      - args: the arguments to add ips to bbc instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchAddIP(args *api.BatchAddIpArgs) (*api.BatchAddIpResponse, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.BatchAddIp(c, args, body)
}

// BatchDelIP - Delete ips of instance
//
// PARAMS:
//      - args: the arguments to add ips to bbc instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchDelIP(args *api.BatchDelIpArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return api.BatchDelIp(c, args, body)
}

//cds sdk
// CreateCDSVolume - create a CDS volume
//
// PARAMS:
//     - args: the arguments to create CDS
// RETURNS:
//     - *api.CreateCDSVolumeResult: the result of create CDS volume, contains new volume ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateCDSVolume(args *api.CreateCDSVolumeArgs) (*api.CreateCDSVolumeResult, error) {
	return api.CreateCDSVolume(c, args)
}

//cds sdk
// CreateCDSVolumeV3 - create a CDS volume
//
// PARAMS:
//     - args: the arguments to create CDS
// RETURNS:
//     - *api.CreateCDSVolumeResult: the result of create CDS volume, contains new volume ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateCDSVolumeV3(args *api.CreateCDSVolumeV3Args) (*api.CreateCDSVolumeResult, error) {
	return api.CreateCDSVolumeV3(c, args)
}

// ListCDSVolume - list all cds volume with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all cds
// RETURNS:
//     - *api.ListCDSVolumeResult: the result of list all CDS volume
//     - error: nil if success otherwise the specific error
func (c *Client) ListCDSVolume(queryArgs *api.ListCDSVolumeArgs) (*api.ListCDSVolumeResult, error) {
	return api.ListCDSVolume(c, queryArgs)
}

// ListCDSVolumeV3 - list all cds volume with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all cds
// RETURNS:
//     - *api.ListCDSVolumeResultV3: the result of list all CDS volume
//     - error: nil if success otherwise the specific error
func (c *Client) ListCDSVolumeV3(queryArgs *api.ListCDSVolumeArgs) (*api.ListCDSVolumeResultV3, error) {
	return api.ListCDSVolumeV3(c, queryArgs)
}

// GetCDSVolumeDetail - get a CDS volume's detail info
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
// RETURNS:
//     - *api.GetVolumeDetailResult: the result of get a specific CDS volume's info
//     - error: nil if success otherwise the specific error
func (c *Client) GetCDSVolumeDetail(volumeId string) (*api.GetVolumeDetailResult, error) {
	return api.GetCDSVolumeDetail(c, volumeId)
}

// GetCDSVolumeDetailV3 - get a CDS volume's detail info
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
// RETURNS:
//     - *api.GetVolumeDetailResultV3: the result of get a specific CDS volume's info
//     - error: nil if success otherwise the specific error
func (c *Client) GetCDSVolumeDetailV3(volumeId string) (*api.GetVolumeDetailResultV3, error) {
	return api.GetCDSVolumeDetailV3(c, volumeId)
}

// AttachCDSVolume - attach a CDS volume to an instance
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
//     - args: the arguments to attach a CDS volume
// RETURNS:
//     - *api.AttachVolumeResult: the result of attach a CDS volume
//     - error: nil if success otherwise the specific error
func (c *Client) AttachCDSVolume(volumeId string, args *api.AttachVolumeArgs) (*api.AttachVolumeResult, error) {
	return api.AttachCDSVolume(c, volumeId, args)
}

// DetachCDSVolume - detach a CDS volume
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
//     - args: the arguments to detach a CDS volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DetachCDSVolume(volumeId string, args *api.DetachVolumeArgs) error {
	return api.DetachCDSVolume(c, volumeId, args)
}

// DeleteCDSVolume - delete a CDS volume
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteCDSVolume(volumeId string) error {
	return api.DeleteCDSVolume(c, volumeId)
}

// DeleteCDSVolumeNew - delete a CDS volume and snapshot
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
//     - args: the arguments to delete a CDS volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteCDSVolumeNew(volumeId string, args *api.DeleteCDSVolumeArgs) error {
	return api.DeleteCDSVolumeNew(c, volumeId, args)
}

// ResizeCDSVolume - resize a CDS volume
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
//     - args: the arguments to resize CDS volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ResizeCDSVolume(volumeId string, args *api.ResizeCSDVolumeArgs) error {
	return api.ResizeCDSVolume(c, volumeId, args)
}

// RollbackCDSVolume - rollback a CDS volume
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
//     - args: the arguments to rollback a CDS volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RollbackCDSVolume(volumeId string, args *api.RollbackCSDVolumeArgs) error {
	return api.RollbackCDSVolume(c, volumeId, args)
}

// PurchaseReservedCDSVolume - purchase reserve a CDS volume
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
//     - args: the arguments to purchase reserve a CDS volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) PurchaseReservedCDSVolume(volumeId string, args *api.PurchaseReservedCSDVolumeArgs) error {
	return api.PurchaseReservedCDSVolume(c, volumeId, args)
}

// RenameCDSVolume - rename a CDS volume
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
//     - args: the arguments to rename a CDS volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RenameCDSVolume(volumeId string, args *api.RenameCSDVolumeArgs) error {
	return api.RenameCDSVolume(c, volumeId, args)
}

// ModifyCDSVolume - modify a CDS volume
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
//     - args: the arguments to modify a CDS volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyCDSVolume(volumeId string, args *api.ModifyCSDVolumeArgs) error {
	return api.ModifyCDSVolume(c, volumeId, args)
}

// ModifyChargeTypeCDSVolume - modify a CDS volume's charge type
//
// PARAMS:
//     - volumeId: the specific CDS volume ID
//     - args: the arguments to create instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyChargeTypeCDSVolume(volumeId string, args *api.ModifyChargeTypeCSDVolumeArgs) error {
	return api.ModifyChargeTypeCDSVolume(c, volumeId, args)
}

// AutoRenewCDSVolume - auto renew the specified cds volume
//
// PARAMS:
//     - args: the arguments to auto renew the cds volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) AutoRenewCDSVolume(args *api.AutoRenewCDSVolumeArgs) error {
	return api.AutoRenewCDSVolume(c, args)
}

// CancelAutoRenewCDSVolume - cancel auto renew the specified cds volume
//
// PARAMS:
//     - args: the arguments to cancel auto renew the cds volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CancelAutoRenewCDSVolume(args *api.CancelAutoRenewCDSVolumeArgs) error {
	return api.CancelAutoRenewCDSVolume(c, args)
}

//securityGroup sdk
// CreateSecurityGroup - create a security group
//
// PARAMS:
//     - args: the arguments to create security group
// RETURNS:
//     - *api.CreateSecurityGroupResult: the result of create security group
//     - error: nil if success otherwise the specific error
func (c *Client) CreateSecurityGroup(args *api.CreateSecurityGroupArgs) (*api.CreateSecurityGroupResult, error) {
	return api.CreateSecurityGroup(c, args)
}

// ListSecurityGroup - list all security group
//
// PARAMS:
//     - args: the arguments to list all security group
// RETURNS:
//     - *api.ListSecurityGroupResult: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific error
func (c *Client) ListSecurityGroup(queryArgs *api.ListSecurityGroupArgs) (*api.ListSecurityGroupResult, error) {
	return api.ListSecurityGroup(c, queryArgs)
}

// AuthorizeSecurityGroupRule - authorize a security group rule
//
// PARAMS:
//     - securityGroupId: the specific securityGroup ID
//     - args: the arguments to authorize a security group rule
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) AuthorizeSecurityGroupRule(securityGroupId string, args *api.AuthorizeSecurityGroupArgs) error {
	return api.AuthorizeSecurityGroupRule(c, securityGroupId, args)
}

// RevokeSecurityGroupRule - revoke a security group rule
//
// PARAMS:
//     - securityGroupId: the specific securityGroup ID
//     - args: the arguments to revoke security group rule
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RevokeSecurityGroupRule(securityGroupId string, args *api.RevokeSecurityGroupArgs) error {
	return api.RevokeSecurityGroupRule(c, securityGroupId, args)
}

// DeleteSecurityGroup - delete a security group
//
// PARAMS:
//     - securityGroupId: the specific securityGroup ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteSecurityGroup(securityGroupId string) error {
	return api.DeleteSecurityGroup(c, securityGroupId)
}

//image sdk
// CreateImage - create an image
//
// PARAMS:
//     - args: the arguments to create image
// RETURNS:
//     - *api.CreateImageResult: the result of create Image
//     - error: nil if success otherwise the specific error
func (c *Client) CreateImage(args *api.CreateImageArgs) (*api.CreateImageResult, error) {
	return api.CreateImage(c, args)
}

// ListImage - list all images
//
// PARAMS:
//     - args: the arguments to list all images
// RETURNS:
//     - *api.ListImageResult: the result of list all images
//     - error: nil if success otherwise the specific error
func (c *Client) ListImage(queryArgs *api.ListImageArgs) (*api.ListImageResult, error) {
	return api.ListImage(c, queryArgs)
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

// RemoteCopyImage - copy an image from other region
//
// PARAMS:
//     - imageId: the specific image ID
//     - args: the arguments to remote copy an image
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RemoteCopyImage(imageId string, args *api.RemoteCopyImageArgs) error {
	return api.RemoteCopyImage(c, imageId, args)
}

// RemoteCopyImageReturnImageIds - copy an image from other region
//
// PARAMS:
//     - imageId: the specific image ID
//     - args: the arguments to remote copy an image
// RETURNS:
//     - imageIds of destination region if success otherwise the specific error
func (c *Client) RemoteCopyImageReturnImageIds(imageId string, args *api.RemoteCopyImageArgs) (*api.RemoteCopyImageResult, error) {
	return api.RemoteCopyImageReturnImageIds(c, imageId, args)
}

// CancelRemoteCopyImage - cancel a copy image from other region operation
//
// PARAMS:
//     - imageId: the specific image ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CancelRemoteCopyImage(imageId string) error {
	return api.CancelRemoteCopyImage(c, imageId)
}

// ShareImage - share an image
//
// PARAMS:
//     - imageId: the specific image ID
//     - args: the arguments to share an image
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ShareImage(imageId string, args *api.SharedUser) error {
	return api.ShareImage(c, imageId, args)
}

// UnShareImage - cancel share an image
//
// PARAMS:
//     - imageId: the specific image ID
//     - args: the arguments to cancel share an image
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnShareImage(imageId string, args *api.SharedUser) error {
	return api.UnShareImage(c, imageId, args)
}

// GetImageSharedUser - get user list use this image
//
// PARAMS:
//     - imageId: the specific image ID
// RETURNS:
//     - *api.GetImageSharedUserResult: the result of user list
//     - error: nil if success otherwise the specific error
func (c *Client) GetImageSharedUser(imageId string) (*api.GetImageSharedUserResult, error) {
	return api.GetImageSharedUser(c, imageId)
}

// GetImageOS - get image OS
//
// PARAMS:
//     - args: the arguments to get OS info
// RETURNS:
//     - *api.GetImageOsResult: the result of get image OS info
//     - error: nil if success otherwise the specific error
func (c *Client) GetImageOS(args *api.GetImageOsArgs) (*api.GetImageOsResult, error) {
	return api.GetImageOS(c, args)
}

// CreateSnapshot - create a snapshot
//
// PARAMS:
//     - args: the arguments to create a snapshot
// RETURNS:
//     - *api.CreateSnapshotResult: the result of create snapshot
//     - error: nil if success otherwise the specific error
func (c *Client) CreateSnapshot(args *api.CreateSnapshotArgs) (*api.CreateSnapshotResult, error) {
	return api.CreateSnapshot(c, args)
}

// ListSnapshot - list all snapshots
//
// PARAMS:
//     - args: the arguments to list all snapshots
// RETURNS:
//     - *api.ListSnapshotResult: the result of list all snapshots
//     - error: nil if success otherwise the specific error
func (c *Client) ListSnapshot(args *api.ListSnapshotArgs) (*api.ListSnapshotResult, error) {
	return api.ListSnapshot(c, args)
}

// ListSnapshotChain - list all snapshot chains
//
// PARAMS:
//     - args: the arguments to list all snapshot chains
// RETURNS:
//     - *api.ListSnapshotChainResult: the result of list all snapshot chains
//     - error: nil if success otherwise the specific error
func (c *Client) ListSnapshotChain(args *api.ListSnapshotChainArgs) (*api.ListSnapshotChainResult, error) {
	return api.ListSnapshotChain(c, args)
}

// GetSnapshotDetail - get a snapshot's detail info
//
// PARAMS:
//     - snapshotId: the specific snapshot ID
// RETURNS:
//     - *api.GetSnapshotDetailResult: the result of get snapshot's detail info
//     - error: nil if success otherwise the specific error
func (c *Client) GetSnapshotDetail(snapshotId string) (*api.GetSnapshotDetailResult, error) {
	return api.GetSnapshotDetail(c, snapshotId)
}

// DeleteSnapshot - delete a snapshot
//
// PARAMS:
//     - snapshotId: the specific snapshot ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteSnapshot(snapshotId string) error {
	return api.DeleteSnapshot(c, snapshotId)
}

// CreateAutoSnapshotPolicy - create an auto snapshot policy
//
// PARAMS:
//     - args: the arguments to create an auto snapshot policy
// RETURNS:
//     - *api.CreateASPResult: the result of create an auto snapshot policy
//     - error: nil if success otherwise the specific error
func (c *Client) CreateAutoSnapshotPolicy(args *api.CreateASPArgs) (*api.CreateASPResult, error) {
	return api.CreateAutoSnapshotPolicy(c, args)
}

// AttachAutoSnapshotPolicy - attach an ASP to volumes
//
// PARAMS:
//     - aspId: the specific auto snapshot policy ID
//     - args: the arguments to attach an ASP
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) AttachAutoSnapshotPolicy(aspId string, args *api.AttachASPArgs) error {
	return api.AttachAutoSnapshotPolicy(c, aspId, args)
}

// DetachAutoSnapshotPolicy - detach an ASP
//
// PARAMS:
//     - aspId: the specific auto snapshot policy ID
//     - args: the arguments to detach an ASP
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DetachAutoSnapshotPolicy(aspId string, args *api.DetachASPArgs) error {
	return api.DetachAutoSnapshotPolicy(c, aspId, args)
}

// DeleteAutoSnapshotPolicy - delete an auto snapshot policy
//
// PARAMS:
//     - aspId: the specific auto snapshot policy ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteAutoSnapshotPolicy(aspId string) error {
	return api.DeleteAutoSnapshotPolicy(c, aspId)
}

// ListAutoSnapshotPolicy - list all auto snapshot policies
//
// PARAMS:
//     - args: the arguments to create instance
// RETURNS:
//     - *api.ListASPResult: the result of list all auto snapshot policies
//     - error: nil if success otherwise the specific error
func (c *Client) ListAutoSnapshotPolicy(args *api.ListASPArgs) (*api.ListASPResult, error) {
	return api.ListAutoSnapshotPolicy(c, args)
}

// GetAutoSnapshotPolicy - get an auto snapshot policy's meta
//
// PARAMS:
//     - aspId: the specific auto snapshot policy ID
// RETURNS:
//     - *api.GetASPDetailResult: the result of get an auto snapshot policy's meta
//     - error: nil if success otherwise the specific error
func (c *Client) GetAutoSnapshotPolicy(aspId string) (*api.GetASPDetailResult, error) {
	return api.GetAutoSnapshotPolicyDetail(c, aspId)
}

// UpdateAutoSnapshotPolicy - update an auto snapshot policy
//
// PARAMS:
//     - args: the arguments to update an auto snapshot policy
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateAutoSnapshotPolicy(args *api.UpdateASPArgs) error {
	return api.UpdateAutoSnapshotPolicy(c, args)
}

// ListSpec - list all spec
//
// RETURNS:
//     - *api.ListSpecResult: the result of all spec
//     - error: nil if success otherwise the specific error
func (c *Client) ListSpec() (*api.ListSpecResult, error) {
	return api.ListSpec(c)
}

// ListZone - list all zones
//
// RETURNS:
//     - *api.ListZoneResult: the result of list all zones
//     - error: nil if success otherwise the specific error
func (c *Client) ListZone() (*api.ListZoneResult, error) {
	return api.ListZone(c)
}

// ListFlavorSpec - get the specified flavor list
//
// PARAMS:
//	   - args: the arguments to list the specified flavor
// RETURNS:
//     - *api.ListFlavorSpecResult: result of the specified flavor list
//     - error: nil if success otherwise the specific error
func (c *Client) ListFlavorSpec(args *api.ListFlavorSpecArgs) (*api.ListFlavorSpecResult, error) {
	return api.ListFlavorSpec(c, args)
}

// GetPriceBySpec - get the price information of specified instance.
//
// PARAMS:
//	   - args: the arguments to get the price information of specified instance.
// RETURNS:
//     - *api.GetPriceBySpecResult: result of the specified instance's price information
//     - error: nil if success otherwise the specific error
func (c *Client) GetPriceBySpec(args *api.GetPriceBySpecArgs) (*api.GetPriceBySpecResult, error) {
	return api.GetPriceBySpec(c, args)
}

// CreateDeploySet - create a deploy set
//
// PARAMS:
//     - args: the arguments to create a deploy set
// RETURNS:
//     - *CreateDeploySetResult: results of creating a deploy set
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
//     - *ListDeploySetsResult: the result of list all deploy sets
//     - error: nil if success otherwise the specific error
func (c *Client) ListDeploySets() (*api.ListDeploySetsResult, error) {
	return api.ListDeploySets(c)
}

// ModifyDeploySet - modify the deploy set
//
// PARAMS:
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - *ModifyDeploySetArgs: the detail of the deploy set
//     - error: nil if success otherwise the specific error
func (c *Client) ModifyDeploySet(deploySetId string, args *api.ModifyDeploySetArgs) (error, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, nil
	}
	return api.ModifyDeploySet(c, deploySetId, args.ClientToken, body), nil
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

// GetDeploySet - get details of the deploy set
//
// PARAMS:
//     - deploySetId: the id of the deploy set
// RETURNS:
//     - *GetDeploySetResult: the detail of the deploy set
//     - error: nil if success otherwise the specific error
func (c *Client) GetDeploySet(deploySetId string) (*api.DeploySetResult, error) {
	return api.GetDeploySet(c, deploySetId)
}

// UpdateInstanceDeploySet - update deployset and instance relation
//
// PARAMS:
//     - args: the arguments to update deployset and instance relation
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateInstanceDeploySet(args *api.UpdateInstanceDeployArgs) (error, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return api.UpdateInstanceDeploy(c, args.ClientToken, body), nil
}

// DelInstanceDeploySet - delete deployset and instance relation
//
// PARAMS:
//     - args: the arguments to delete deployset and instance relation
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DelInstanceDeploySet(args *api.DelInstanceDeployArgs) (error, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	return api.DelInstanceDeploy(c, args.ClientToken, body), nil
}

// ResizeInstanceBySpec - resize a specific instance
//
// PARAMS:
//     - instanceId: the specific instance ID
//     - args: the arguments to resize a specific instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ResizeInstanceBySpec(instanceId string, args *api.ResizeInstanceArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.ResizeInstanceBySpec(c, instanceId, args.ClientToken, body)
}

// RebuildBatchInstance - batch rebuild instances
//
// PARAMS:
//     - args: the arguments to batch rebuild instances
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchRebuildInstances(args *api.RebuildBatchInstanceArgs) error {
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

	return api.BatchRebuildInstances(c, body)
}

// ChangeToPrepaid - to prepaid
//
// PARAMS:
//     - instanceId: instanceId
//     - args: the arguments to ChangeToPrepaid
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ChangeToPrepaid(instanceId string, args *api.ChangeToPrepaidRequest) (*api.ChangeToPrepaidResponse,
	error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.ChangeToPrepaid(c, instanceId, body)
}

// BindInstanceToTags - bind instance to tags
//
// PARAMS:
//     - instanceId: instanceId
//     - args: the arguments to BindInstanceToTags
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindInstanceToTags(instanceId string, args *api.BindTagsRequest) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.BindInstanceToTags(c, instanceId, body)
}

// UnBindInstanceToTags - unbind instance to tags
//
// PARAMS:
//     - instanceId: instanceId
//     - args: the arguments to unBindInstanceToTags
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnBindInstanceToTags(instanceId string, args *api.UnBindTagsRequest) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.UnBindInstanceToTags(c, instanceId, body)
}

// GetInstanceNoChargeList - get instance with nocharge list
//
// PARAMS:
//     - args: the arguments to list all nocharge instance
// RETURNS:
//     - *api.ListInstanceResult: the result of list Instance
//     - error: nil if success otherwise the specific error
func (c *Client) GetInstanceNoChargeList(args *api.ListInstanceArgs) (*api.ListInstanceResult, error) {
	return api.GetInstanceNoChargeList(c, args)
}

// CreateBidInstance - create an instance with the specific parameters
//
// PARAMS:
//     - args: the arguments to create instance
// RETURNS:
//     - *api.CreateInstanceResult: the result of create Instance, contains new Instance ID
//     - error: nil if success otherwise the specific error
func (c *Client) CreateBidInstance(args *api.CreateInstanceArgs) (*api.CreateInstanceResult, error) {
	if len(args.AdminPass) > 0 {
		cryptedPass, err := api.Aes128EncryptUseSecreteKey(c.Config.Credentials.SecretAccessKey, args.AdminPass)
		if err != nil {
			return nil, err
		}

		args.AdminPass = cryptedPass
	}

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.CreateBidInstance(c, args.ClientToken, body)
}

// CancelBidOrder - Cancel the bidding instance order.
//
// PARAMS:
//     - args: the arguments to cancel bid order
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CancelBidOrder(args *api.CancelBidOrderRequest) (*api.CreateBidInstanceResult, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.CancelBidOrder(c, args.ClientToken, body)
}

// GetAvailableDiskInfo - get available diskInfos of the specified zone
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - zoneName: the zone name eg:cn-bj-a
// RETURNS:
//     - *GetAvailableDiskInfoResult: the result of the specified zone diskInfos
//     - error: nil if success otherwise the specific error
func (c *Client) GetAvailableDiskInfo(zoneName string) (*api.GetAvailableDiskInfoResult, error) {
	return api.GetAvailableDiskInfo(c, zoneName)
}

// DeletePrepayVolume - delete the volumes for prepay
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments of method
// RETURNS:
//     - *VolumeDeleteResultResponse: the result of deleting volumes
//     - error: nil if success otherwise the specific error
func (c *Client) DeletePrepayVolume(args *api.VolumePrepayDeleteRequestArgs) (*api.VolumeDeleteResultResponse, error) {
	return api.DeletePrepayVolume(c, args)
}

// ListTypeZones - list instanceType zones
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceType: the instanceType like "N1"
// RETURNS:
//     - *api.ListTypeZonesResult: the result of list instanceType zones
//     - error: nil if success otherwise the specific error
func (c *Client) ListTypeZones(args *api.ListTypeZonesArgs) (*api.ListTypeZonesResult, error) {
	return api.ListTypeZones(c, args)
}

// ListInstanceEni - get the eni list of the bcc instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - instanceId: the bcc instance id
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ListInstanceEnis(instanceId string) (*api.ListInstanceEniResult, error) {
	return api.ListInstanceEnis(c, instanceId)
}

func (c *Client) CreateKeypair(args *api.CreateKeypairArgs) (*api.KeypairResult, error) {
	return api.CreateKeypair(c, args)
}

func (c *Client) ImportKeypair(args *api.ImportKeypairArgs) (*api.KeypairResult, error) {
	return api.ImportKeypair(c, args)
}

func (c *Client) AttachKeypair(args *api.AttackKeypairArgs) error {
	return api.AttachKeypair(c, args)
}

func (c *Client) DetachKeypair(args *api.DetachKeypairArgs) error {
	return api.DetachKeypair(c, args)
}

func (c *Client) DeleteKeypair(args *api.DeleteKeypairArgs) error {
	return api.DeleteKeypair(c, args)
}

func (c *Client) GetKeypairDetail(keypairId string) (*api.KeypairResult, error) {
	return api.GetKeypairDetail(c, keypairId)
}

func (c *Client) ListKeypairs(args *api.ListKeypairArgs) (*api.ListKeypairResult, error) {
	return api.ListKeypairs(c, args)
}

func (c *Client) RenameKeypair(args *api.RenameKeypairArgs) error {
	return api.RenameKeypair(c, args)
}

func (c *Client) UpdateKeypairDescription(args *api.KeypairUpdateDescArgs) error {
	return api.UpdateKeypairDescription(c, args)
}

// GetAllStocks - get the bcc and bbc's stock
//
// RETURNS:
//     - *GetAllStocksResult: the result of the bcc and bbc's stock
//     - error: nil if success otherwise the specific error
func (c *Client) GetAllStocks() (*api.GetAllStocksResult, error) {
	return api.GetAllStocks(c)
}

// GetStockWithDeploySet - get the bcc's stock with deploySet
//
// RETURNS:
//     - *GetStockWithDeploySetResults: the result of the bcc's stock
//     - error: nil if success otherwise the specific error
func (c *Client) GetStockWithDeploySet(args *api.GetStockWithDeploySetArgs) (*api.GetStockWithDeploySetResults, error) {
	return api.GetStockWithDeploySet(c, args)
}

// GetStockWithSpec - get the bcc's stock with spec
//
// RETURNS:
//     - *GetStockWithSpecResults: the result of the bcc's stock
//     - error: nil if success otherwise the specific error
func (c *Client) GetStockWithSpec(args *api.GetStockWithSpecArgs) (*api.GetStockWithSpecResults, error) {
	return api.GetStockWithSpec(c, args)
}

func (c *Client) GetInstanceCreateStock(args *api.CreateInstanceStockArgs) (*api.InstanceStockResult, error) {
	return api.GetInstanceCreateStock(c, args)
}

func (c *Client) GetInstanceResizeStock(args *api.ResizeInstanceStockArgs) (*api.InstanceStockResult, error) {
	return api.GetInstanceResizeStock(c, args)
}

// BatchCreateAutoRenewRules - Batch Create AutoRenew Rules
//
// PARAMS:
//      - args: the arguments to batch create autorenew rules
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchCreateAutoRenewRules(args *api.BccCreateAutoRenewArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return api.BatchCreateAutoRenewRules(c, body)
}

// BatchDeleteAutoRenewRules - Batch Delete AutoRenew Rules
//
// PARAMS:
//      - args: the arguments to batch delete autorenew rules
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchDeleteAutoRenewRules(args *api.BccDeleteAutoRenewArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return api.BatchDeleteAutoRenewRules(c, body)
}

func (c *Client) DeleteInstanceIngorePayment(args *api.DeleteInstanceIngorePaymentArgs) (*api.DeleteInstanceResult, error) {
	return api.DeleteInstanceIngorePayment(c, args)
}

// DeleteRecycledInstance - delete a recycled instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteRecycledInstance(instanceId string) error {
	return api.DeleteRecycledInstance(c, instanceId)
}

func (c *Client) ListInstanceByInstanceIds(args *api.ListInstanceByInstanceIdArgs) (*api.ListInstancesResult, error) {
	return api.ListInstanceByInstanceIds(c, args)
}

// BatchDeleteInstanceWithRelateResource - batch delete instance and all eip/cds relate it
//
// PARAMS:
//     - args: the arguments to batch delete instance and its relate resource
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchDeleteInstanceWithRelateResource(args *api.BatchDeleteInstanceWithRelateResourceArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.BatchDeleteInstanceWithRelatedResource(c, body)
}

// BatchStartInstance - batch start instance
//
// PARAMS:
//     - args: the arguments to batch start instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchStartInstance(args *api.BatchStartInstanceArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.BatchStartInstance(c, body)
}

// BatchStopInstance - batch stop instance
//
// PARAMS:
//     - args: the arguments to batch stop instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchStopInstance(args *api.BatchStopInstanceArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return api.BatchStopInstance(c, body)
}

// ListInstanceTypes - list all instance type with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all instance type
// RETURNS:
//     - *api.ListInstanceTypeResults: the result of list Instance type
//     - error: nil if success otherwise the specific error
func (c *Client) ListInstanceTypes(args *api.ListInstanceTypeArgs) (*api.ListInstanceTypeResults, error) {
	return api.ListInstanceTypes(c, args)
}

// ListIdMappings - Long and short ID conversion parameters
//
// PARAMS:
//     - args: the arguments to Long and short ID conversion
// RETURNS:
//     - *api.ListIdMappingResults: result of the Long short mapping
//     - error: nil if success otherwise the specific error
func (c *Client) ListIdMappings(args *api.ListIdMappingArgs) (*api.ListIdMappingResults, error) {
	return api.ListIdMappings(c, args)
}

// BatchResizeInstance - batch resize a specific instance
//
// PARAMS:
//     - args: the arguments to resize a specific instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchResizeInstance(args *api.BatchResizeInstanceArgs) (*api.BatchResizeInstanceResults, error) {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return nil, jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	return api.BatchResizeInstance(c, body)
}

// DeleteSecurityGroupRule - delete a security group rule
//
// PARAMS:
//	   - securityGroupRuleId: the id of the specific security group rule
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteSecurityGroupRule(args *api.DeleteSecurityGroupRuleArgs) error {
	return api.DeleteSecurityGroupRule(c, args)
}

// UpdateSecurityGroupRule - update security group rule with the specific parameters
//
// PARAMS:
//     - args: the arguments to update the specific security group rule
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateSecurityGroupRule(args *api.UpdateSecurityGroupRuleArgs) error {
	return api.UpdateSecurityGroupRule(c, args)
}

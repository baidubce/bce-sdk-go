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

	return api.CreateInstance(c, args.ClientToken, body)
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

	return api.CreateInstanceBySpec(c, args.ClientToken, body)
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

// DeleteInstance - delete a specific instance
//
// PARAMS:
//     - instanceId: the specific instance ID
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteInstance(instanceId string) error {
	return api.DeleteInstance(c, instanceId)
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

	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}

	return api.InstancePurchaseReserved(c, instanceId, args.ClientToken, body)
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

// BatchAddIP - Add ips to instance
//
// PARAMS:
//      - args: the arguments to add ips to bbc instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchAddIP(args *api.BatchAddIpArgs) error {
	jsonBytes, jsonErr := json.Marshal(args)
	if jsonErr != nil {
		return jsonErr
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	return api.BatchAddIp(c, body)
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
	return api.BatchDelIp(c, body)
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

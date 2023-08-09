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

// client.go - define the client for BIE service

// Package bie defines the bie services of BCE. The supported APIs are all defined in sub-package
package bie

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bie/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "iotedge." + bce.DEFAULT_REGION + "." + bce.DEFAULT_DOMAIN
)

// Client of BIE(iot edge) service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the BIE service client with default configuration.
// Use `cli.Config.xxx` to access the config or change it to non-default value.
func NewClient(ak, sk, endpoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
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

//////////////////////////////////////////////
// group API
//////////////////////////////////////////////

// ListGroup - list all groups
//
//   - *ListGroupReq: page no, page size, and name
//
// RETURNS:
//   - *api.ListGroupResult: the all groups
//   - error: the return error if any occurs
func (c *Client) ListGroup(lgr *api.ListGroupReq) (*api.ListGroupResult, error) {
	return api.ListGroup(c, lgr)
}

// GetGroup - get a group
//
// RETURNS:
//   - *api.Group: the group
//   - error: the return error if any occurs
func (c *Client) GetGroup(groupUuid string) (*api.Group, error) {
	return api.GetGroup(c, groupUuid)
}

// CreateGroup - create a group
//
// PARAMS:
//   - cgr: parameters to create group
//
// RETURNS:
//   - *CreateGroupResult: the result group
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateGroup(cgr *api.CreateGroupReq) (*api.CreateGroupResult, error) {
	return api.CreateGroup(c, cgr)
}

// EditGroup - edit a group
//
// PARAMS:
//   - groupId: the group to edit
//   - egr: parameters to update group
//
// RETURNS:
//   - *Group: the result group
//   - error: nil if ok otherwise the specific error
func (c *Client) EditGroup(groupid string, egr *api.EditGroupReq) (*api.Group, error) {
	return api.EditGroup(c, groupid, egr)
}

// DeleteGroup - delete a group of the account
//
// PARAMS:
//   - groupUuid: id of the group
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteGroup(groupUuid string) error {
	return api.DeleteGroup(c, groupUuid)
}

//////////////////////////////////////////////
// core API
//////////////////////////////////////////////

// ListCore - list all core of a group
//
// PARAMS:
//   - groupUuid: id of the group
//
// RETURNS:
//   - *ListCoreResult: the result core list
//   - error: nil if ok otherwise the specific error
func (c *Client) ListCore(groupUuid string) (*api.ListCoreResult, error) {
	return api.ListCore(c, groupUuid)
}

// GetCore - get a core of a group
//
// PARAMS:
//   - groupUuid: id of the group
//   - coreid: id of the core
//
// RETURNS:
//   - *CoreResult: the result core
//   - error: nil if ok otherwise the specific error
func (c *Client) GetCore(groupUuid string, coreid string) (*api.CoreResult, error) {
	return api.GetCore(c, groupUuid, coreid)
}

// RenewCoreAuth - renew the auth of a core
//
// PARAMS:
//   - coreid: id of the core
//
// RETURNS:
//   - *CoreInfo: the result core info
//   - error: nil if ok otherwise the specific error
func (c *Client) RenewCoreAuth(coreid string) (*api.CoreInfo, error) {
	return api.RenewCoreAuth(c, coreid)
}

// GetCoreStatus - get the status of a core
//
// PARAMS:
//   - coreid: id of the core
//
// RETURNS:
//   - *CoreStatus: the status of the core
//   - error: nil if ok otherwise the specific error
func (c *Client) GetCoreStatus(coreid string) (*api.CoreStatus, error) {
	return api.GetCoreStatus(c, coreid)
}

//////////////////////////////////////////////
// Config API
//////////////////////////////////////////////

// ListConfig - list all configs
//
// PARAMS:
//   - coreid: id of the core
//
// RETURNS:
//   - *ListConfigResult: the result config list
//   - error: nil if ok otherwise the specific error
func (c *Client) ListConfig(coreid string,
	lcr *api.ListConfigReq) (*api.ListConfigResult, error) {
	return api.ListConfig(c, coreid, lcr)
}

// GetConfig - get a config
//
// PARAMS:
//   - coreid: id of the core
//   - ver: version, e.g:$LATEST
//
// RETURNS:
//   - *CfgResult: the result config
//   - error: nil if ok otherwise the specific error
func (c *Client) GetConfig(coreid string, ver string) (*api.CfgResult, error) {
	return api.GetConfig(c, coreid, ver)
}

// PubConfig - pub a config
//
// PARAMS:
//   - cpr: id of the core and version
//
// RETURNS:
//   - *CfgResult: the pub result
//   - error: nil if ok otherwise the specific error
func (c *Client) PubConfig(cpr *api.CoreidVersion, cpb *api.CfgPubBody) (*api.CfgResult, error) {
	return api.PubConfig(c, cpr, cpb)
}

// DeployConfig - deploy a config
//
// PARAMS:
//   - coreid: id of the core
//   - ver: version, e.g:$LATEST
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeployConfig(coreid string, ver string) error {
	return api.DeployConfig(c, coreid, ver)
}

// DownloadConfig - download a config
//
// PARAMS:
//   - *CfgDownloadReq: id, version, with bin or not
//
// RETURNS:
//   - *CfgDownloadResult: the result
//   - error: nil if ok otherwise the specific error
func (c *Client) DownloadConfig(cdr *api.CfgDownloadReq) (*api.CfgDownloadResult, error) {
	return api.DownloadConfig(c, cdr)
}

//////////////////////////////////////////////
// Service API
//////////////////////////////////////////////

// CreateService - create a service
//
// PARAMS:
//   - *CoreidVersion: coreid, version
//   - *CreateServiceReq: request parameters
//
// RETURNS:
//   - *ServiceResult: the result
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateService(cv *api.CoreidVersion,
	sr *api.CreateServiceReq) (*api.ServiceResult, error) {
	return api.CreateService(c, cv, sr)
}

// GetService - get a service
//
// PARAMS:
//   - *IdVerName: coreid, version and name
//
// RETURNS:
//   - *ServiceResult: the result
//   - error: nil if ok otherwise the specific error
func (c *Client) GetService(ivn *api.IdVerName) (*api.ServiceResult, error) {
	return api.GetService(c, ivn)
}

// EditService - edit a service
//
// PARAMS:
//   - *IdVerName: coreid, version, and name
//
// RETURNS:
//   - *ServiceResult: the result
//   - error: nil if ok otherwise the specific error
func (c *Client) EditService(ivn *api.IdVerName,
	esr *api.EditServiceReq) (*api.ServiceResult, error) {
	return api.EditService(c, ivn, esr)
}

// DeleteService - delete a service
//
// PARAMS:
//   - *IdVerName: coreid, version, and name
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteService(ivn *api.IdVerName) error {
	return api.DeleteService(c, ivn)
}

// ReorderService - reorder service after an other service
//
// PARAMS:
//   - *IdVerName: coreid, version, and name
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) ReorderService(ivn *api.IdVerName, after string) error {
	return api.ReorderService(c, ivn, after)
}

// VolumeOp - attache of detach volumes
//
// PARAMS:
//   - *CoreidVersion: coreid, version
//   - *VolumeOpReq: request parameters
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) VolumeOp(cv *api.CoreidVersion, vor *api.VolumeOpReq) error {
	return api.VolumeOp(c, cv, vor)
}

// ////////////////////////////////////////////
// Volume API
// ////////////////////////////////////////////
// ListVolumeTpl - list all volume template
//
// PARAMS:
// RETURNS:
//   - *ListVolTemplate: the result volume template list
//   - error: nil if ok otherwise the specific error
func (c *Client) ListVolumeTpl() (*api.ListVolTemplate, error) {
	return api.ListVolumeTpl(c)
}

// CreateVolume - create a volume
//
// PARAMS:
//   - CreateVolReq: the request parameters
//
// RETURNS:
//   - *VolumeResult: the result volume
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateVolume(cvr *api.CreateVolReq) (*api.VolumeResult, error) {
	return api.CreateVolume(c, cvr)
}

// ListVolume - list a bunch of volumes
//
// PARAMS:
//   - ListVolReq: the request parameters
//
// RETURNS:
//   - *ListVolumeResult: the result list
//   - error: nil if ok otherwise the specific error
func (c *Client) ListVolume(lvr *api.ListVolumeReq) (*api.ListVolumeResult, error) {
	return api.ListVolume(c, lvr)
}

// GetVolume - get a volume
//
// PARAMS:
//   - name: the requested volume name
//
// RETURNS:
//   - *VolumeResult: the result vaolume
//   - error: nil if ok otherwise the specific error
func (c *Client) GetVolume(name string) (*api.VolumeResult, error) {
	return api.GetVolume(c, name)
}

// EditVolume - edit a volume
//
// PARAMS:
//   - name: the requested volume name
//   - *EditVolumeReq: request parameters
//
// RETURNS:
//   - *VolumeResult: the result vaolume
//   - error: nil if ok otherwise the specific error
func (c *Client) EditVolume(name string, evr *api.EditVolumeReq) (*api.VolumeResult, error) {
	return api.EditVolume(c, name, evr)
}

// DeleteVolume - delete a volume
//
// PARAMS:
//   - name: the requested volume name
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteVolume(name string) error {
	return api.DeleteVolume(c, name)
}

// ListVolumeVer - list version list of a volume
//
// PARAMS:
//   - NameVersion: the request parameters, version can be empty
//
// RETURNS:
//   - *ListVolumeVerResult: the result list
//   - error: nil if ok otherwise the specific error
func (c *Client) ListVolumeVer(nv *api.NameVersion) (*api.ListVolumeVerResult, error) {
	return api.ListVolumeVer(c, nv)
}

// PubVolumeVer - pub version list of a volume
//
// PARAMS:
//   - name: the request volume name
//
// RETURNS:
//   - *VolumeResult: the result volume
//   - error: nil if ok otherwise the specific error
func (c *Client) PubVolumeVer(name string) (*api.VolumeResult, error) {
	return api.PubVolumeVer(c, name)
}

// DownloadVolVer - get the download address of a specific version of volume
//
// PARAMS:
//   - name: the request volume name
//   - version: the request volume version
//
// RETURNS:
//   - *VolDownloadResult: the result
//   - error: nil if ok otherwise the specific error
func (c *Client) DownloadVolVer(name string,
	version string) (*api.VolDownloadResult, error) {
	return api.DownloadVolVer(c, name, version)
}

// CreateVolFile - create a volume file
//
// PARAMS:
//   - name: the request volume name
//   - *CreateVolFileReq: request parameters
//
// RETURNS:
//   - *CreateVolFileReq: the result
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateVolFile(name string,
	cvf *api.CreateVolFileReq) (*api.CreateVolFileReq, error) {
	return api.CreateVolFile(c, name, cvf)
}

// GetVolumeFile - get a volume file
//
// PARAMS:
//   - GetVolFileReq: volume name, version anf filename
//
// RETURNS:
//   - *ListVolumeVerResult: the result list
//   - error: nil if ok otherwise the specific error
func (c *Client) GetVolumeFile(cvfr *api.GetVolFileReq) (*api.CreateVolFileReq, error) {
	return api.GetVolumeFile(c, cvfr)
}

// EditVolumeFile - edit a volume file
//
// PARAMS:
//   - *Name2: the requested volume name and file name
//   - *EditVolFileReq: the content, the request body
//
// RETURNS:
//   - *CreateVolFileReq: the result
//   - error: nil if ok otherwise the specific error
func (c *Client) EditVolumeFile(names *api.Name2,
	body *api.EditVolFileReq) (*api.CreateVolFileReq, error) {
	return api.EditVolumeFile(c, names, body)
}

// DeleteVolFile - delete a volume file
//
// PARAMS:
//   - name: the requested volume name
//   - filename: the requested volume filename
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteVolFile(name string, filename string) error {
	return api.DeleteVolFile(c, name, filename)
}

// ClearVolFile - clear all volume files
//
// PARAMS:
//   - name: the requested volume name
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) ClearVolFile(name string) error {
	return api.ClearVolFile(c, name)
}

// ListVolCore - list the cores associated with the volume
//
// PARAMS:
//   - *ListVolCoreReq: the request parameters
//
// RETURNS:
//   - *ListVolumeVerResult: the result list
//   - error: nil if ok otherwise the specific error
func (c *Client) ListVolCore(lvcr *api.ListVolCoreReq) (*api.ListVolCoreResult, error) {
	return api.ListVolCore(c, lvcr)
}

// EditCoreVolVer - edit the core version
//
// PARAMS:
//   - name: the volume name
//   - *EditCoreVolVerReq: the request parameters
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) EditCoreVolVer(name string,
	ecvr *api.EditCoreVolVerReq) error {
	return api.EditCoreVolVer(c, name, ecvr)
}

// ImportCfc - import a CFC function into volume
//
// PARAMS:
//   - name: the volume name
//   - *ImportCfcReq: the request parameters
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) ImportCfc(name string,
	icr *api.ImportCfcReq) error {
	return api.ImportCfc(c, name, icr)
}

// ImportBos - import a BOS file into volume
//
// PARAMS:
//   - name: the volume name
//   - *ImportBosReq: the request parameters
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) ImportBos(name string,
	ibr *api.ImportBosReq) error {
	return api.ImportBos(c, name, ibr)
}

// ////////////////////////////////////////////
// Volume API
// ////////////////////////////////////////////
// ListImageSys - list all system docker images
//
// PARAMS:
//   - ListImageReq: list request parameters
//
// RETURNS:
//   - *ListImageResult: the result iamge list
//   - error: nil if ok otherwise the specific error
func (c *Client) ListImageSys(lir *api.ListImageReq) (*api.ListImageResult, error) {
	return api.ListImageSys(c, lir)
}

// GetImageSys - get a system docker images
//
// PARAMS:
//   - uuid: the image uuid
//
// RETURNS:
//   - *Image: the result iamge
//   - error: nil if ok otherwise the specific error
func (c *Client) GetImageSys(uuid string) (*api.Image, error) {
	return api.GetImageSys(c, uuid)
}

// ListImageUser - list all user docker images
//
// PARAMS:
//   - ListImageReq: list request parameters
//
// RETURNS:
//   - *ListImageResult: the result iamge list
//   - error: nil if ok otherwise the specific error
func (c *Client) ListImageUser(lir *api.ListImageReq) (*api.ListImageResult, error) {
	return api.ListImageUser(c, lir)
}

// GetImageUser - get a user docker image
//
// PARAMS:
//   - uuid: the image uuid
//
// RETURNS:
//   - *Image: the result iamge
//   - error: nil if ok otherwise the specific error
func (c *Client) GetImageUser(uuid string) (*api.Image, error) {
	return api.GetImageUser(c, uuid)
}

// CreateImageUser - create a user docker image
//
// PARAMS:
//   - CreateImageReq: request parameters, name, image url, description
//
// RETURNS:
//   - *Image: the result iamge
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateImageUser(cir *api.CreateImageReq) (*api.Image, error) {
	return api.CreateImageUser(c, cir)
}

// EditImageUser - edit a user docker image information
//
// PARAMS:
//   - uuid: the image uuid
//   - EditImageReq: request parameter: description
//
// RETURNS:
//   - *Image: the result iamge
//   - error: nil if ok otherwise the specific error
func (c *Client) EditImageUser(uuid string, eir *api.EditImageReq) (*api.Image, error) {
	return api.EditImageUser(c, uuid, eir)
}

// DeleteImageUser - delete a user docker image
//
// PARAMS:
//   - uuid: the image uuid
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteImageUser(uuid string) error {
	return api.DeleteImageUser(c, uuid)
}

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

// volume.go - the volume related APIs definition supported by the BIE service

// Package api defines all APIs supported by the BIE service of BCE.
package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	PREFIX_V3VOLTPL = "/v3/volumeTemplate"
	PREFIX_V3VOL    = "/v3/volume"
)

// ListVolumeTpl - list all volume template
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - *ListVolTemplate: the result volume template list
//     - error: nil if ok otherwise the specific error
func ListVolumeTpl(cli bce.Client) (*ListVolTemplate, error) {
	url := PREFIX_V3VOLTPL
	result := &ListVolTemplate{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateVolume - create a volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - CreateVolReq: the request parameters
// RETURNS:
//     - *VolumeResult: the result volume
//     - error: nil if ok otherwise the specific error
func CreateVolume(cli bce.Client, cvr *CreateVolReq) (*VolumeResult, error) {
	url := PREFIX_V3VOL
	result := &VolumeResult{}
	req := &PostHttpReq{Url: url, Result: result, Body: cvr}
	err := Post(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ListVolume - list a bunch of volumes
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - ListVolReq: the request parameters
// RETURNS:
//     - *ListVolumeResult: the result list
//     - error: nil if ok otherwise the specific error
func ListVolume(cli bce.Client, lvr *ListVolumeReq) (*ListVolumeResult, error) {
	url := PREFIX_V3VOL
	result := &ListVolumeResult{}
	params := map[string]string{
		"pageNo":   strconv.Itoa(lvr.PageNo),
		"pageSize": strconv.Itoa(lvr.PageSize)}
	if lvr.Name != "" {
		params["name"] = lvr.Name
	}
	if lvr.Tag != "" {
		params["tag"] = lvr.Tag
	}

	req := &GetHttpReq{Url: url, Result: result, Params: params}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetVolume - get a volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the requested volume name
// RETURNS:
//     - *VolumeResult: the result vaolume
//     - error: nil if ok otherwise the specific error
func GetVolume(cli bce.Client, name string) (*VolumeResult, error) {
	url := PREFIX_V3VOL + "/" + name
	result := &VolumeResult{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EditVolume - edit a volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the requested volume name
//     - *EditVolumeReq: request parameters
// RETURNS:
//     - *VolumeResult: the result vaolume
//     - error: nil if ok otherwise the specific error
func EditVolume(cli bce.Client, name string, evr *EditVolumeReq) (*VolumeResult, error) {
	url := PREFIX_V3VOL + "/" + name
	result := &VolumeResult{}
	req := &PostHttpReq{Url: url, Result: result, Body: evr}
	err := Put(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteVolume - delete a volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the requested volume name
// RETURNS:
//     - error: nil if ok otherwise the specific error
func DeleteVolume(cli bce.Client, name string) error {
	url := PREFIX_V3VOL + "/" + name
	req := &bce.BceRequest{}
	req.SetUri(url)
	req.SetMethod(http.DELETE)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	return nil
}

// ListVolumeVer - list version list of a volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - NameVersion: the request parameters, version can be empty
// RETURNS:
//     - *ListVolumeVerResult: the result list
//     - error: nil if ok otherwise the specific error
func ListVolumeVer(cli bce.Client, nv *NameVersion) (*ListVolumeVerResult, error) {
	url := PREFIX_V3VOL + "/" + nv.Name + "/version"
	params := map[string]string{}
	if nv.Version != "" {
		params["version"] = nv.Version
	}

	result := &ListVolumeVerResult{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PubVolumeVer - pub version list of a volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the request volume name
// RETURNS:
//     - *VolumeResult: the result volume
//     - error: nil if ok otherwise the specific error
func PubVolumeVer(cli bce.Client, name string) (*VolumeResult, error) {
	url := PREFIX_V3VOL + "/" + name + "/version"

	result := &VolumeResult{}
	req := &PostHttpReq{Url: url, Result: result}
	err := Post(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DownloadVolVer - get the download address of a specific version of volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the request volume name
//     - version: the request volume version
// RETURNS:
//     - *VolDownloadResult: the result
//     - error: nil if ok otherwise the specific error
func DownloadVolVer(cli bce.Client, name string,
	version string) (*VolDownloadResult, error) {
	url := PREFIX_V3VOL + "/" + name + "/version/" + version + "/zipfile"

	result := &VolDownloadResult{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateVolFile - create a volume file
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the request volume name
//     - *CreateVolFileReq: request parameters
// RETURNS:
//     - *CreateVolFileReq: the result
//     - error: nil if ok otherwise the specific error
func CreateVolFile(cli bce.Client, name string,
	cvf *CreateVolFileReq) (*CreateVolFileReq, error) {
	url := PREFIX_V3VOL + "/" + name + "/file"

	result := &CreateVolFileReq{}
	req := &PostHttpReq{Url: url, Result: result, Body: cvf}
	err := Post(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetVolumeFile - get a volume file
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - GetVolFileReq: volume name, version anf filename
// RETURNS:
//     - *ListVolumeVerResult: the result list
//     - error: nil if ok otherwise the specific error
func GetVolumeFile(cli bce.Client, cvfr *GetVolFileReq) (*CreateVolFileReq, error) {
	url := PREFIX_V3VOL + "/" + cvfr.Name + "/version/" + cvfr.Version + "/file/" + cvfr.FileName

	result := &CreateVolFileReq{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EditVolumeFile - edit a volume file
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *Name2: the requested volume name and file name
//     - *EditVolFileReq: the content, the request body
// RETURNS:
//     - *CreateVolFileReq: the result
//     - error: nil if ok otherwise the specific error
func EditVolumeFile(cli bce.Client, names *Name2,
	body *EditVolFileReq) (*CreateVolFileReq, error) {
	url := PREFIX_V3VOL + "/" + names.Name + "/file/" + names.FileName

	result := &CreateVolFileReq{}
	req := &PostHttpReq{Url: url, Result: result, Body: body}
	err := Put(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteVolFile - delete a volume file
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the requested volume name
//     - filename: the requested volume filename
// RETURNS:
//     - error: nil if ok otherwise the specific error
func DeleteVolFile(cli bce.Client, name string, filename string) error {
	url := PREFIX_V3VOL + "/" + name + "/file/" + filename
	req := &bce.BceRequest{}
	req.SetUri(url)
	req.SetMethod(http.DELETE)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	return nil
}

// ClearVolFile - clear all volume files
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the requested volume name
// RETURNS:
//     - error: nil if ok otherwise the specific error
func ClearVolFile(cli bce.Client, name string) error {
	url := PREFIX_V3VOL + "/" + name + "/file"
	req := &bce.BceRequest{}
	req.SetUri(url)
	req.SetMethod(http.DELETE)
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	return nil
}

// ListVolCore - list the cores associated with the volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *ListVolCoreReq: the request parameters
// RETURNS:
//     - *ListVolumeVerResult: the result list
//     - error: nil if ok otherwise the specific error
func ListVolCore(cli bce.Client, lvcr *ListVolCoreReq) (*ListVolCoreResult, error) {
	url := PREFIX_V3VOL + "/" + lvcr.Name + "/core"
	params := map[string]string{
		"pageNo":   strconv.Itoa(lvcr.PageNo),
		"pageSize": strconv.Itoa(lvcr.PageSize)}

	result := &ListVolCoreResult{}
	req := &GetHttpReq{Url: url, Result: result, Params: params}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EditCoreVolVer - edit core version
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the volume name
//     - *EditCoreVolVerReq: the request parameters
// RETURNS:
//     - error: nil if ok otherwise the specific error
func EditCoreVolVer(cli bce.Client, name string,
	ecvr *EditCoreVolVerReq) error {
	url := PREFIX_V3VOL + "/" + name + "/core"

	result := &ListVolCoreResult{}
	req := &PostHttpReq{Url: url, Result: result, Body: ecvr}
	err := Put(cli, req)
	if err != nil {
		return err
	}
	return nil
}

// ImportCfc - import a CFC function into volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the volume name
//     - *ImportCfcReq: the request parameters
// RETURNS:
//     - error: nil if ok otherwise the specific error
func ImportCfc(cli bce.Client, name string,
	icr *ImportCfcReq) error {
	url := PREFIX_V3VOL + "/" + name + "/cfc"
	req := &PostHttpReq{Url: url, Body: icr}
	err := Post(cli, req)
	if err != nil {
		return err
	}
	return nil
}

// ImportBos - import a BOS file into volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - name: the volume name
//     - *ImportBosReq: the request parameters
// RETURNS:
//     - error: nil if ok otherwise the specific error
func ImportBos(cli bce.Client, name string,
	ibr *ImportBosReq) error {
	url := PREFIX_V3VOL + "/" + name + "/bos"
	req := &PostHttpReq{Url: url, Body: ibr}
	err := Post(cli, req)
	if err != nil {
		return err
	}
	return nil
}

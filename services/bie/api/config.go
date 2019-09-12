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

// config.go - the config related APIs definition supported by the BIE service

// Package api defines all APIs supported by the BIE service of BCE.
package api

import (
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

const (
	PREFIX_V3CORE = "/v3/core"
)

// ListConfig - list all configs
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - coreid: id of the core
// RETURNS:
//     - *ListConfigResult: the result config list
//     - error: nil if ok otherwise the specific error
func ListConfig(cli bce.Client, coreid string, lcr *ListConfigReq) (*ListConfigResult, error) {
	params := map[string]string{}

	if lcr.Status != "" {
		params["status"] = lcr.Status
	}
	params["pageNo"] = strconv.Itoa(lcr.PageNo)
	params["pageSize"] = strconv.Itoa(lcr.PageSize)

	url := PREFIX_V3CORE + "/" + coreid + "/config"
	result := &ListConfigResult{}
	req := &GetHttpReq{Url: url, Result: result, Params: params}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetConfig - get a config
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - coreid: id of the core
//     - ver: version, e.g:$LATEST
// RETURNS:
//     - *CfgResult: the result config
//     - error: nil if ok otherwise the specific error
func GetConfig(cli bce.Client, coreid string, ver string) (*CfgResult, error) {
	url := PREFIX_V3CORE + "/" + coreid + "/config/" + ver

	result := &CfgResult{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// PubConfig - pub a config
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - cpr: id of the core and version
// RETURNS:
//     - *CfgResult: the pub result
//     - error: nil if ok otherwise the specific error
func PubConfig(cli bce.Client, cpr *CoreidVersion, cpb *CfgPubBody) (*CfgResult, error) {
	url := PREFIX_V3CORE + "/" + cpr.Coreid + "/config/" + cpr.Version + "/publish"
	result := &CfgResult{}
	req := &PostHttpReq{Url: url, Body: cpb, Result: result}
	err := Post(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeployConfig - deploy a config
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - coreid: id of the core
//     - ver: version, e.g:$LATEST
// RETURNS:
//     - error: nil if ok otherwise the specific error
func DeployConfig(cli bce.Client, coreid string, ver string) error {
	url := PREFIX_V3CORE + "/" + coreid + "/config/" + ver + "/deploy"

	req := &PostHttpReq{Url: url}
	err := Post(cli, req)
	if err != nil {
		return err
	}

	return nil
}

// DownloadConfig - download a config
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *CfgDownloadReq: id, version, with bin or not
// RETURNS:
//     - *CfgDownloadResult: the result
//     - error: nil if ok otherwise the specific error
func DownloadConfig(cli bce.Client, cdr *CfgDownloadReq) (*CfgDownloadResult, error) {
	url := PREFIX_V3CORE + "/" + cdr.Coreid + "/config/" + cdr.Version + "/download"
	params := map[string]string{"withBin": strconv.FormatBool(cdr.WithBin)}
	result := &CfgDownloadResult{}
	req := &GetHttpReq{Url: url, Params: params, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateService - create a service
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *CoreidVersion: coreid, version
//     - *CreateServiceReq: request parameters
// RETURNS:
//     - *ServiceResult: the result
//     - error: nil if ok otherwise the specific error
func CreateService(cli bce.Client, cv *CoreidVersion,
	sr *CreateServiceReq) (*ServiceResult, error) {
	url := PREFIX_V3CORE + "/" + cv.Coreid + "/config/" + cv.Version + "/service"

	result := &ServiceResult{}
	req := &PostHttpReq{Url: url, Body: sr, Result: result}
	err := Post(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// GetService - get a service
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *IdVerName: coreid, version and name
// RETURNS:
//     - *ServiceResult: the result
//     - error: nil if ok otherwise the specific error
func GetService(cli bce.Client, ivn *IdVerName) (*ServiceResult, error) {
	url := PREFIX_V3CORE + "/" + ivn.Coreid + "/config/" + ivn.Version + "/service/" + ivn.Name

	result := &ServiceResult{}
	req := &GetHttpReq{Url: url, Result: result}
	err := Get(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// EditService - edit a service
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *IdVerName: coreid, version, and name
// RETURNS:
//     - *ServiceResult: the result
//     - error: nil if ok otherwise the specific error
func EditService(cli bce.Client, ivn *IdVerName,
	esr *EditServiceReq) (*ServiceResult, error) {
	url := PREFIX_V3CORE + "/" + ivn.Coreid + "/config/" + ivn.Version + "/service/" + ivn.Name

	result := &ServiceResult{}
	req := &PostHttpReq{Url: url, Body: esr, Result: result}
	err := Put(cli, req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteService - delete a service
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *IdVerName: coreid, version, and name
// RETURNS:
//     - error: nil if ok otherwise the specific error
func DeleteService(cli bce.Client, ivn *IdVerName) error {
	url := PREFIX_V3CORE + "/" + ivn.Coreid + "/config/" + ivn.Version + "/service/" + ivn.Name

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

// ReorderService - reorder service after an other service
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *IdVerName: coreid, version, and name
// RETURNS:
//     - error: nil if ok otherwise the specific error
func ReorderService(cli bce.Client, ivn *IdVerName, after string) error {
	url := PREFIX_V3CORE + "/" + ivn.Coreid + "/config/" + ivn.Version + "/service/" + ivn.Name
	url += "/after/" + after

	req := &PostHttpReq{Url: url}
	err := Post(cli, req)
	if err != nil {
		return err
	}

	return nil
}

// VolumeOp - attache of detach volumes
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - *CoreidVersion: coreid, version
//     - *VolumeOpReq: request parameters
// RETURNS:
//     - error: nil if ok otherwise the specific error
func VolumeOp(cli bce.Client, cv *CoreidVersion, vor *VolumeOpReq) error {
	url := PREFIX_V3CORE + "/" + cv.Coreid + "/config/" + cv.Version + "/volume"

	req := &PostHttpReq{Url: url, Body: vor}
	err := Put(cli, req)
	if err != nil {
		return err
	}

	return nil
}

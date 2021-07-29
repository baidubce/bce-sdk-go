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

// cds.go - the cds APIs definition supported by the BCC service

// Package api defines all APIs supported by the BCC service of BCE.
package api

import (
	"encoding/json"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateCDSVolume - create a specified count of cds volumes
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to create cds volumes
// RETURNS:
//     - *CreateCDSVolumeResult: the result of volume ids newly created
//     - error: nil if success otherwise the specific error
func CreateCDSVolume(cli bce.Client, args *CreateCDSVolumeArgs) (*CreateCDSVolumeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUri())
	req.SetMethod(http.POST)

	if args.ClientToken != "" {
		req.SetParam("clientToken", args.ClientToken)
	}

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateCDSVolumeResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// CreateCDSVolumeV3 - create a specified count of cds volumes
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to create cds volumes
// RETURNS:
//     - *CreateCDSVolumeResult: the result of volume ids newly created
//     - error: nil if success otherwise the specific error
func CreateCDSVolumeV3(cli bce.Client, args *CreateCDSVolumeV3Args) (*CreateCDSVolumeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeV3Uri())
	req.SetMethod(http.POST)

	if args.ClientToken != "" {
		req.SetParam("clientToken", args.ClientToken)
	}

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &CreateCDSVolumeResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// ListCDSVolume - list all cds volumes with the given parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryArgs: the optional arguments to list cds volumes
// RETURNS:
//     - *ListCDSVolumeResult: the result of cds volume list
//     - error: nil if success otherwise the specific error
func ListCDSVolume(cli bce.Client, queryArgs *ListCDSVolumeArgs) (*ListCDSVolumeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.InstanceId) != 0 {
			req.SetParam("instanceId", queryArgs.InstanceId)
		}
		if len(queryArgs.ZoneName) != 0 {
			req.SetParam("zoneName", queryArgs.ZoneName)
		}
		if len(queryArgs.Marker) != 0 {
			req.SetParam("marker", queryArgs.Marker)
		}
		if queryArgs.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(queryArgs.MaxKeys))
		}
	}

	if queryArgs == nil || queryArgs.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListCDSVolumeResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// ListCDSVolumeV3 - list all cds volumes with the given parameters
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - queryArgs: the optional arguments to list cds volumes
// RETURNS:
//     - *ListCDSVolumeResultV3: the result of cds volume list
//     - error: nil if success otherwise the specific error
func ListCDSVolumeV3(cli bce.Client, queryArgs *ListCDSVolumeArgs) (*ListCDSVolumeResultV3, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeV3Uri())
	req.SetMethod(http.GET)

	if queryArgs != nil {
		if len(queryArgs.InstanceId) != 0 {
			req.SetParam("instanceId", queryArgs.InstanceId)
		}
		if len(queryArgs.ZoneName) != 0 {
			req.SetParam("zoneName", queryArgs.ZoneName)
		}
		if len(queryArgs.Marker) != 0 {
			req.SetParam("marker", queryArgs.Marker)
		}
		if queryArgs.MaxKeys != 0 {
			req.SetParam("maxKeys", strconv.Itoa(queryArgs.MaxKeys))
		}
	}

	if queryArgs == nil || queryArgs.MaxKeys == 0 {
		req.SetParam("maxKeys", "1000")
	}

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &ListCDSVolumeResultV3{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetCDSVolumeDetail - get details of the specified cds volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the cds volume
// RETURNS:
//     - *GetVolumeDetailResult: the result of the specified cds volume details
//     - error: nil if success otherwise the specific error
func GetCDSVolumeDetail(cli bce.Client, volumeId string) (*GetVolumeDetailResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetVolumeDetailResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// GetCDSVolumeDetail - get details of the specified cds volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the cds volume
// RETURNS:
//     - *GetVolumeDetailResultV3: the result of the specified cds volume details
//     - error: nil if success otherwise the specific error
func GetCDSVolumeDetailV3(cli bce.Client, volumeId string) (*GetVolumeDetailResultV3, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeV3UriWithId(volumeId))
	req.SetMethod(http.GET)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &GetVolumeDetailResultV3{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// AttachCDSVolume - attach an cds volume to a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the cds volume
//     - args: the arguments of instance id
// RETURNS:
//     - *AttachVolumeResult: the result of the attachment
//     - error: nil if success otherwise the specific error
func AttachCDSVolume(cli bce.Client, volumeId string, args *AttachVolumeArgs) (*AttachVolumeResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("attach", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &AttachVolumeResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// DetachCDSVolume - detach an cds volume for a specified instance
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the cds volume
//     - args: the arguments of instance id detached from
// RETURNS:
//     - error: nil if success otherwise the specific error
func DetachCDSVolume(cli bce.Client, volumeId string, args *DetachVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("detach", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// DeleteCDSVolume - delete a specified cds volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the cds volume to be deleted
//     - :
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteCDSVolume(cli bce.Client, volumeId string) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.DELETE)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// DeleteCDSVolumeNew - delete a specified cds volume, the difference from the above api is that \
// can control whether to delete the snapshot associated with the volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the cds volume to be deleted
//     - args: the arguments to delete cds volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteCDSVolumeNew(cli bce.Client, volumeId string, args *DeleteCDSVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.POST)

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// ResizeCDSVolume - resize a specified cds volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the cds volume to be resized
//     - args: the arguments to resize cds volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func ResizeCDSVolume(cli bce.Client, volumeId string, args *ResizeCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	if args.ClientToken != "" {
		req.SetParam("clientToken", args.ClientToken)
	}
	req.SetParam("resize", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// RollbackCDSVolume - roll back a specified cds volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the cds volume to be rolled back
//     - args: the arguments to roll back the cds volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func RollbackCDSVolume(cli bce.Client, volumeId string, args *RollbackCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("rollback", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// PurchaseReservedCDSVolume - renew a specified volume to extend expiration time.
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the volume to be renewed
//     - args: the arguments to renew cds volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func PurchaseReservedCDSVolume(cli bce.Client, volumeId string, args *PurchaseReservedCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	if args.ClientToken != "" {
		req.SetParam("clientToken", args.ClientToken)
	}
	req.SetParam("purchaseReserved", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// RenameCDSVolume - rename a specified cds volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the volume to be renamed
//     - args: the arguments to rename volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func RenameCDSVolume(cli bce.Client, volumeId string, args *RenameCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("rename", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// ModifyCDSVolume - modify attributes of the specified cds volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the volume to be modified
//     - args: arguments to modify volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyCDSVolume(cli bce.Client, volumeId string, args *ModifyCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("modify", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// ModifyChargeTypeCDSVolume - modify the volume billing method, only support Postpaid to Prepaid and Prepaid to Postpaid
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - volumeId: id of the volume to be modified
//     - args: the arguments to modify volume billing method
// RETURNS:
//     - error: nil if success otherwise the specific error
func ModifyChargeTypeCDSVolume(cli bce.Client, volumeId string, args *ModifyChargeTypeCSDVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getVolumeUriWithId(volumeId))
	req.SetMethod(http.PUT)

	req.SetParam("modifyChargeType", "")

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}

	defer func() { resp.Body().Close() }()
	return nil
}

// AutoRenewCDSVolume - auto renew the specified cds volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to auto renew the cds volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func AutoRenewCDSVolume(cli bce.Client, args *AutoRenewCDSVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getAutoRenewVolumeUri())
	req.SetMethod(http.POST)
	if args.ClientToken != "" {
		req.SetParam("clientToken", args.ClientToken)
	}

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { resp.Body().Close() }()

	return nil
}

// CancelAutoRenewCDSVolume - cancel auto renew the specified cds volume
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments to cancel auto renew the cds volume
// RETURNS:
//     - error: nil if success otherwise the specific error
func CancelAutoRenewCDSVolume(cli bce.Client, args *CancelAutoRenewCDSVolumeArgs) error {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getCancelAutoRenewVolumeUri())
	req.SetMethod(http.POST)
	if args.ClientToken != "" {
		req.SetParam("clientToken", args.ClientToken)
	}

	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { resp.Body().Close() }()

	return nil
}

// GetAvailableDiskInfo - get available diskInfos of the specified zone
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - zoneName: the zone name eg:cn-bj-a
// RETURNS:
//     - *GetAvailableDiskInfoResult: the result of the specified zone diskInfos
//     - error: nil if success otherwise the specific error
func GetAvailableDiskInfo(cli bce.Client, zoneName string) (*GetAvailableDiskInfoResult, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getAvailableDiskInfo())
	req.SetMethod(http.GET)
	req.SetParam("zoneName", zoneName)
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetAvailableDiskInfoResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

// DeletePrepayVolume - delete the volumes for prepay
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - args: the arguments of method
// RETURNS:
//     - *VolumeDeleteResultResponse: the result of deleting volumes
//     - error: nil if success otherwise the specific error
func DeletePrepayVolume(cli bce.Client, args *VolumePrepayDeleteRequestArgs) (*VolumeDeleteResultResponse, error) {
	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(getDeletePrepayVolumeUri())
	req.SetMethod(http.POST)
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)

	// Send request and get response
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	jsonBody := &VolumeDeleteResultResponse{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	return jsonBody, nil
}

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

// eip.go - the eip APIs definition supported by the EIP service
package eip

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateEip - create an EIP with the specific parameters
//
// PARAMS:
//     - args: the arguments to create an eip
// RETURNS:
//     - *CreateEipResult: the result of create EIP, contains new EIP's address
//     - error: nil if success otherwise the specific error
func (c *Client) CreateEip(args *CreateEipArgs) (*CreateEipResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set create eip argments")
	}

	if args.BandWidthInMbps <= 0 || args.BandWidthInMbps > 1000 {
		return nil, fmt.Errorf("unsupport bandwidthInMbps value: %d", args.BandWidthInMbps)
	}

	if args.Billing == nil {
		return nil, fmt.Errorf("please set billing")
	}

	result := &CreateEipResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEipUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ResizeEip - resize an EIP with the specific parameters
//
// PARAMS:
//     - eip: the specific EIP
//     - args: the arguments to resize an EIP
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) ResizeEip(eip string, args *ResizeEipArgs) error {
	if args == nil {
		return fmt.Errorf("please set resize eip argments")
	}

	if args.NewBandWidthInMbps <= 0 || args.NewBandWidthInMbps > 1000 {
		return fmt.Errorf("unsupport bandwidthInMbps value: %d", args.NewBandWidthInMbps)
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("resize", "").
		WithBody(args).
		Do()
}

// BindEip - bind an EIP to an instance with the specific parameters
//
// PARAMS:
//     - eip: the specific EIP
//     - args: the arguments to bind an EIP
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindEip(eip string, args *BindEipArgs) error {
	if args == nil {
		return fmt.Errorf("please set bind eip argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("bind", "").
		WithBody(args).
		Do()
}

// UnBindEip - unbind an EIP
//
// PARAMS:
//     - eip: the specific EIP
//     - clientToken: optional parameter, an Idempotent Token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnBindEip(eip, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", clientToken).
		WithQueryParam("unbind", "").
		Do()
}

// DeleteEip - delete an EIP
//
// PARAMS:
//     - eip: the specific EIP
//     - clientToken: optional parameter, an Idempotent Token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteEip(eip, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// ListEip - list all EIP with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all eip
// RETURNS:
//     - *ListEipResult: the result of list all eip, contains new EIP's ID
//     - error: nil if success otherwise the specific error
func (c *Client) ListEip(args *ListEipArgs) (*ListEipResult, error) {
	if args == nil {
		args = &ListEipArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListEipResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipUri()).
		WithQueryParamFilter("eip", args.Eip).
		WithQueryParamFilter("instanceType", args.InstanceType).
		WithQueryParamFilter("instanceId", args.InstanceId).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParamFilter("status", args.Status).
		WithResult(result).
		Do()

	return result, err
}

// PurchaseReservedEip - purchase reserve an eip with the specific parameters
//
// PARAMS:
//     - eip: the specific EIP
//     - args: the arguments to purchase reserve an eip
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) PurchaseReservedEip(eip string, args *PurchaseReservedEipArgs) error {
	if args == nil {
		return fmt.Errorf("please set purchase reserved eip argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// StartAutoRenew - start auto renew an eip
//
// PARAMS:
//     - eip: the specific EIP
//     - args: the arguments to start auto renew an eip
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) StartAutoRenew(eip string, args *StartAutoRenewArgs) error {
	if args == nil {
		return fmt.Errorf("please set eip auto renew argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParam("startAutoRenew", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// StopAutoRenew - stop eip auto renew
//
// PARAMS:
//     - eip: the specific EIP
//     - clientToken: optional parameter, an Idempotent Token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) StopAutoRenew(eip string, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParam("stopAutoRenew", "").
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// ListEipCluster - list all EIP Cluster with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all eip cluster
// RETURNS:
//     - *ListClusterResult: the result of list all eip cluster
//     - error: nil if success otherwise the specific error

func (c *Client) ListEipCluster(args *ListEipArgs) (*ListClusterResult, error) {
	if args == nil {
		args = &ListEipArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}
	result := &ListClusterResult{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipClusterUri()).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetEipCluster - get the eip cluster detail with the clusterId
//
// PARAMS:
//     - clusterId: the specific clusterId
// RETURNS:
//     - *ClusterDetail: the result of eip cluster detail
//     - error: nil if success otherwise the specific error

func (c *Client) GetEipCluster(clusterId string) (*ClusterDetail, error) {
	if len(clusterId) == 0 {
		return nil, fmt.Errorf("please set clusterId argment")
	}
	result := &ClusterDetail{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipClusterUriWithId(clusterId)).
		WithResult(result).
		Do()

	return result, err
}

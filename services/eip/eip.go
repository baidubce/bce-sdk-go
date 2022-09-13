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

// BatchCreateEip - create EIPs with the specific parameters
//
// PARAMS:
//     - args: the arguments to create eips
// RETURNS:
//     - *BatchCreateEipResult: the result of create EIP, contains new EIP's address
//     - error: nil if success otherwise the specific error
func (c *Client) BatchCreateEip(args *BatchCreateEipArgs) (*BatchCreateEipResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set create eip argments")
	}

	if args.Billing == nil {
		return nil, fmt.Errorf("please set billing")
	}

	result := &BatchCreateEipResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEipUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("action", "batch").
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

// OptionalDeleteEip - optionally delete an EIP
//
// PARAMS:
//     - eip: the specific EIP
//     - clientToken: optional parameter, an Idempotent Token
//     - releaseToRecycle: the parameter confirms whether to put the specific EIP in the recycle bin (true) or directly delete it (false)
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) OptionalDeleteEip(eip string, clientToken string, releaseToRecycle bool) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("releaseToRecycle", strconv.FormatBool(releaseToRecycle)).
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

// ListRecycleEip - list all EIP in the recycle bin with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all eip in the recycle bin
// RETURNS:
//     - *ListRecycleEipResult: the result of list all eip in the recycle bin
//     - error: nil if success otherwise the specific error
func (c *Client) ListRecycleEip(args *ListRecycleEipArgs) (*ListRecycleEipResult, error) {
	if args == nil {
		args = &ListRecycleEipArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &ListRecycleEipResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getRecycleEipUri()).
		WithQueryParamFilter("eip", args.Eip).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// RestoreRecycleEip - restore the specific EIP in the recycle bin
//
// PARAMS:
//     - eip: the specific EIP
//     - clientToken: optional parameter, an Idempotent Token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RestoreRecycleEip(eip string, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getRecycleEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", clientToken).
		WithQueryParam("restore", "").
		Do()
}

// DeleteRecycleEip - delete the specific EIP in the recycle bin
//
// PARAMS:
//     - eip: the specific EIP
//     - clientToken: optional parameter, an Idempotent Token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteRecycleEip(eip string, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getRecycleEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
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

// DirectEip - turn on EIP pass through with the specific parameters
//
// PARAMS:
//     - eip: the specific EIP
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DirectEip(eip, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", clientToken).
		WithQueryParam("direct", "").
		Do()
}

// UnDirectEip - turn off EIP pass through with the specific parameters
//
// PARAMS:
//     - eip: the specific EIP
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnDirectEip(eip, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getEipUriWithEip(eip)).
		WithQueryParamFilter("clientToken", clientToken).
		WithQueryParam("unDirect", "").
		Do()
}

// CreateEipTp - create an EIP TP with the specific parameters
//
// PARAMS:
//     - args: the arguments to create an eiptp
// RETURNS:
//     - *CreateEipTpResult: the created eiptp id.
//     - error: nil if success otherwise the specific error
func (c *Client) CreateEipTp(args *CreateEipTpArgs) (*CreateEipTpResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set create eip tp argments")
	}
	if len(args.Capacity) == 0 {
		return nil, fmt.Errorf("please set capacity argment")
	}
	result := &CreateEipTpResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getEipTpUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListEipTp - list all EIP TPs with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all eiptps
// RETURNS:
//     - *ListEipTpResult: the result of listing all eiptps
//     - error: nil if success otherwise the specific error
func (c *Client) ListEipTp(args *ListEipTpArgs) (*ListEipTpResult, error) {
	if args == nil {
		args = &ListEipTpArgs{}
	}
	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}
	result := &ListEipTpResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipTpUri()).
		WithQueryParamFilter("id", args.Id).
		WithQueryParamFilter("deductPolicy", args.DeductPolicy).
		WithQueryParamFilter("status", args.Status).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetEipTp - get the EIP TP detail with the id
//
// PARAMS:
//     - id: the specific eiptp id
// RETURNS:
//     - *EipTpDetail: the result of eiptp detail
//     - error: nil if success otherwise the specific error

func (c *Client) GetEipTp(id string) (*EipTpDetail, error) {
	if len(id) == 0 {
		return nil, fmt.Errorf("please set eiptp id argment")
	}
	result := &EipTpDetail{}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getEipTpUriWithId(id)).
		WithResult(result).
		Do()

	return result, err
}

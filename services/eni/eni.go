/*
 * Copyright 2021 Baidu, Inc.
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

// eni.go - the eni APIs definition supported by the eni service
package eni

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateEni - create an eni with the specific parameters
//
// PARAMS:
//     - args: the arguments to create an eni
// RETURNS:
//     - *CreateEniResult: the result of create eni
//     - error: nil if success otherwise the specific error
func (c *Client) CreateEni(args *CreateEniArgs) (*CreateEniResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createEniArgs cannot be nil.")
	}

	result := &CreateEniResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEni()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// UpdateEni - update an eni
//
// PARAMS:
//     - UpdateEniArgs: the arguments to update an eni
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateEni(args *UpdateEniArgs) error {
	if args == nil {
		return fmt.Errorf("The updateEniArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("modifyAttribute", "").
		Do()
}

// DeleteEni - delete an eni
//
// PARAMS:
//     - DeleteEniArgs: the arguments to delete an eni
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteEni(args *DeleteEniArgs) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// ListEnis - list all eni with the specific parameters
//
// PARAMS:
//     - args: the arguments to list all eni
// RETURNS:
//     - *ListEniResult: the result of list all eni
//     - error: nil if success otherwise the specific error
func (c *Client) ListEni(args *ListEniArgs) (*ListEniResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The ListEniArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListEniResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getURLForEni()).
		WithMethod(http.GET).
		WithQueryParam("vpcId", args.VpcId).
		WithQueryParamFilter("instanceId", args.InstanceId).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys))

	if len(args.PrivateIpAddress) != 0 {
		builder.WithQueryParam("privateIpAddress",
			strings.Replace(strings.Trim(fmt.Sprint(args.PrivateIpAddress), "[]"), " ", ",", -1))
	}

	err := builder.WithResult(result).Do()

	return result, err
}

// GetEniDetail - get the eni detail
//
// PARAMS:
//     - eniId: the specific eniId
// RETURNS:
//     - *Eni: the eni
//     - error: nil if success otherwise the specific error
func (c *Client) GetEniDetail(eniId string) (*Eni, error) {
	if eniId == "" {
		return nil, fmt.Errorf("The eniId cannot be empty.")
	}

	result := &Eni{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(eniId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

// AddPrivateIp - add private ip
//
// PARAMS:
//     - args: the arguments to add private ip
// RETURNS:
//     - *AddPrivateIpResult: the private ip
//     - error: nil if success otherwise the specific error
func (c *Client) AddPrivateIp(args *EniPrivateIpArgs) (*AddPrivateIpResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The EniPrivateIpArgs cannot be nil.")
	}

	result := &AddPrivateIpResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)+"/privateIp").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// BatchAddPrivateIp - batch add private ips
//
// PARAMS:
//     - args: the arguments to batch add private ips, property PrivateIpAddresses or PrivateIpAddressCount is required;
//             when PrivateIpAddressCount is set, private ips will be auto allocated,
//             and if you want assign private ips, please just set PrivateIpAddresses;
// RETURNS:
//     - *BatchAddPrivateIpResult: the private ips
//     - error: nil if success otherwise the specific error
func (c *Client) BatchAddPrivateIp(args *EniBatchPrivateIpArgs) (*BatchAddPrivateIpResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The EniBatchPrivateIpArgs cannot be nil.")
	}

	result := &BatchAddPrivateIpResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)+"/privateIp/batchAdd").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// BatchAddPrivateIpCrossSubnet - batch add private ips that support cross subnet, white list function
//
// PARAMS:
//     - args: the arguments to batch add private ips, property PrivateIps or PrivateIpAddressCount is required;
//             when PrivateIpAddressCount is set, private ips in subnet assigned by 'SubnetId' property will be auto allocated;
//             if you want assign private ips, please just set PrivateIps, and you can also assgin subnet with property 'PrivateIpArgs.SubnetId';
// RETURNS:
//     - *BatchAddPrivateIpResult: the private ips
//     - error: nil if success otherwise the specific error
func (c *Client) BatchAddPrivateIpCrossSubnet(args *EniBatchAddPrivateIpCrossSubnetArgs) (*BatchAddPrivateIpResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The EniBatchAddPrivateIpCrossSubnetArgs cannot be nil.")
	}

	result := &BatchAddPrivateIpResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)+"/privateIp/batchAddCrossSubnet").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// DeletePrivateIp - delete private ip
//
// PARAMS:
//     - args: the arguments to delete private ip
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeletePrivateIp(args *EniPrivateIpArgs) error {
	if args == nil {
		return fmt.Errorf("The EniPrivateIpArgs cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)+"/privateIp/"+args.PrivateIpAddress).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

// BatchDeletePrivateIp - batch delete private ip
//
// PARAMS:
//     - args: the arguments to batch delete private ipa
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchDeletePrivateIp(args *EniBatchPrivateIpArgs) error {
	if args == nil {
		return fmt.Errorf("The EniBatchPrivateIpArgs cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)+"/privateIp/batchDel").
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

// AttachEniInstance - eni attach instance
//
// PARAMS:
//     - args: the arguments to attach instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) AttachEniInstance(args *EniInstance) error {
	if args == nil {
		return fmt.Errorf("The EniInstance cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)).
		WithMethod(http.PUT).
		WithQueryParam("attach", "").
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

// DetachEniInstance - eni detach instance
//
// PARAMS:
//     - args: the arguments to detach instance
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DetachEniInstance(args *EniInstance) error {
	if args == nil {
		return fmt.Errorf("The EniInstance cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)).
		WithMethod(http.PUT).
		WithQueryParam("detach", "").
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

// BindEniPublicIp - eni bind public ip
//
// PARAMS:
//     - args: the arguments to bind public ip
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindEniPublicIp(args *BindEniPublicIpArgs) error {
	if args == nil {
		return fmt.Errorf("The BindEniPublicIpArgs cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)).
		WithMethod(http.PUT).
		WithQueryParam("bind", "").
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

// UnBindEniPublicIp - eni unbind public ip
//
// PARAMS:
//     - args: the arguments to bind public ip
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnBindEniPublicIp(args *UnBindEniPublicIpArgs) error {
	if args == nil {
		return fmt.Errorf("The UnBindEniPublicIpArgs cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)).
		WithMethod(http.PUT).
		WithQueryParam("unBind", "").
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

// UpdateEniSecurityGroup - update eni sg
//
// PARAMS:
//     - args: the arguments to update eni sg
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateEniSecurityGroup(args *UpdateEniSecurityGroupArgs) error {
	if args == nil {
		return fmt.Errorf("The UpdateEniSecurityGroupArgs cannot be nil.")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForEniId(args.EniId)).
		WithMethod(http.PUT).
		WithQueryParam("bindSg", "").
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()

	return err
}

func (c *Client) GetEniQuota(args *EniQuoteArgs) (*EniQuoteInfo, error) {

	result := &EniQuoteInfo{}
	request := bce.NewRequestBuilder(c).
		WithURL(getURLForEni() + "/quota").
		WithMethod(http.GET)
	if args.EniId != "" {
		request.WithQueryParam("eniId", args.EniId)
	}
	if args.InstanceId != "" {
		request.WithQueryParam("instanceId", args.InstanceId)
	}
	err := request.WithResult(result).Do()

	return result, err
}

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

// vpc.go - the vpc APIs definition supported by the VPC service

package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateVPC - create a new VPC with the specified parameters
//
// PARAMS:
//   - args: the arguments to create VPC
//
// RETURNS:
//   - *CreateVPCResult: the id of the VPC newly created
//   - error: nil if success otherwise the specific error
func (c *Client) CreateVPC(args *CreateVPCArgs) (*CreateVPCResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createVPCArgs cannot be nil.")
	}

	result := &CreateVPCResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVPC()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListVPC - list all VPCs with the specified parameters
//
// PARAMS:
//   - args: the arguments to list VPCs
//
// RETURNS:
//   - *ListVPCResult: the result of all VPCs
//   - error: nil if success otherwise the specific error
func (c *Client) ListVPC(args *ListVPCArgs) (*ListVPCResult, error) {
	if args == nil {
		args = &ListVPCArgs{}
	}
	if args.IsDefault != "" && args.IsDefault != "true" && args.IsDefault != "false" {
		return nil, fmt.Errorf("The field isDefault can only be true or false.")
	}
	if args.MaxKeys < 0 || args.MaxKeys > 1000 {
		return nil, fmt.Errorf("The field maxKeys is out of range [0, 1000]")
	}

	result := &ListVPCResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getURLForVPC()).
		WithMethod(http.GET).
		WithResult(result).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("isDefault", args.IsDefault)
	if args.MaxKeys != 0 {
		builder.WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys))
	}
	err := builder.Do()

	return result, err
}

// GetVPCDetail - get details of the specified VPC
//
// PARAMS:
//   - vpcId: the VPC id
//
// RETURNS:
//   - *GetVPCDetailResult: the details of the specified VPC
//   - error: nil if success otherwise the specific error
func (c *Client) GetVPCDetail(vpcId string) (*GetVPCDetailResult, error) {
	result := &GetVPCDetailResult{}

	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(vpcId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

// UpdateVPC - update a specified VPC
//
// PARAMS:
//   - vpcId: the id of the specified VPC
//   - updateVPCArgs: the arguments to udpate VPC
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateVPC(vpcId string, updateVPCArgs *UpdateVPCArgs) error {
	if updateVPCArgs == nil {
		return fmt.Errorf("The updateVPCArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(vpcId)).
		WithMethod(http.PUT).
		WithQueryParam("modifyAttribute", "").
		WithBody(updateVPCArgs).
		WithQueryParamFilter("clientToken", updateVPCArgs.ClientToken).
		Do()
}

// DeleteVPC - delete a specified VPC
//
// PARAMS:
//   - vpcId: the VPC id to be deleted
//   - clientToken: the idempotent token
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteVPC(vpcId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(vpcId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// GetPrivateIpAddressesInfo - get the privateIpAddressesInfo from vpc
//
// PARAMS:
//   - getVpcPrivateIpArgs: the arguments to GetPrivateIpAddressInfo
//
// RETURNS:
//   - *VpcPrivateIpAddressesResult: the privateIpAddresses info of the specified privateIps in specified vpc
//   - error: nil if success otherwise the specific error
func (c *Client) GetPrivateIpAddressesInfo(args *GetVpcPrivateIpArgs) (*VpcPrivateIpAddressesResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The GetVpcPrivateIpArgs cannot be nil.")
	}
	result := &VpcPrivateIpAddressesResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(args.VpcId)+"/privateIpAddressInfo").
		WithMethod(http.GET).WithQueryParamFilter("privateIpRange", args.PrivateIpRange)
	if len(args.PrivateIpAddresses) != 0 {
		for i := range args.PrivateIpAddresses {
			builder.WithQueryParam("privateIpAddresses", args.PrivateIpAddresses[i])
		}
	}
	err := builder.WithResult(result).Do()
	return result, err
}

// GetNetworkTopologyInfo - get the network topology info
//
// PARAMS:
//   - getNetworkTopologyArgs: the arguments to GetNetworkTopologyInfo
//
// RETURNS:
//   - *NetworkTopologyResult: the network topologies info obtained based on host ip or host id
//   - error: nil if success otherwise the specific error
func (c *Client) GetNetworkTopologyInfo(args *GetNetworkTopologyArgs) (*NetworkTopologyResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The GetNetworkTopologyArgs cannot be nil.")
	}
	result := &NetworkTopologyResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getURLForNetworkTopology()).
		WithMethod(http.GET).
		WithQueryParamFilter("hostIp", args.HostIp).
		WithQueryParamFilter("hostId", args.HostId)
	err := builder.WithResult(result).Do()
	return result, err
}

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
		WithQueryParamFilter("vpcIds", args.VpcIds).
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

// GetResourceIp - get the resource IP information from vpc
//
// PARAMS:
//   - args: the arguments to GetResourceIp
//
// RETURNS:
//   - *GetResourceIpResult: the resource IP information with pagination
//   - error: nil if success otherwise the specific error
func (c *Client) GetResourceIp(args *GetResourceIpArgs) (*GetResourceIpResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The GetResourceIpArgs cannot be nil.")
	}
	if args.VpcId == "" {
		return nil, fmt.Errorf("The vpcId field is required.")
	}

	result := &GetResourceIpResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getURLForVPC()+"/resourceIp").
		WithMethod(http.GET).
		WithQueryParam("vpcId", args.VpcId).
		WithQueryParamFilter("subnetId", args.SubnetId).
		WithQueryParamFilter("resourceType", args.ResourceType)

	if args.PageNo > 0 {
		builder.WithQueryParam("pageNo", strconv.Itoa(args.PageNo))
	}
	if args.PageSize > 0 {
		builder.WithQueryParam("pageSize", strconv.Itoa(args.PageSize))
	}

	err := builder.WithResult(result).Do()
	return result, err
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

// CreateVPCDhcp - create vpc's dhcp info with the specified parameters
//
// PARAMS:
//   - args: the arguments to create vpc's dhcp info
func (c *Client) CreateVPCDhcp(vpcId string, createVpcDhcpArgs *CreateVpcDhcpArgs) error {
	if createVpcDhcpArgs == nil {
		return fmt.Errorf("The CreateVPCDhcp cannot be nil.")
	}
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(vpcId)+"/dhcp").
		WithMethod(http.POST).
		WithBody(createVpcDhcpArgs).
		WithQueryParamFilter("clientToken", createVpcDhcpArgs.ClientToken).
		Do()
}

// UpdateVPCDhcp - update vpc's dhcp info with the specified parameters
//
//	if domainNameServers is nil, will delete vpc's dhcp.
//
// PARAMS:
//   - args: the arguments to create vpc's dhcp info
func (c *Client) UpdateVPCDhcp(vpcId string, updateVpcDhcpArgs *UpdateVpcDhcpArgs) error {
	if updateVpcDhcpArgs == nil {
		return fmt.Errorf("The UpdateVPCDhcp cannot be nil.")
	}
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(vpcId)+"/dhcp").
		WithMethod(http.PUT).
		WithBody(updateVpcDhcpArgs).
		WithQueryParamFilter("clientToken", updateVpcDhcpArgs.ClientToken).
		Do()
}

// GetVPCDhcpInfo - get the dhcp info of specified vpc
// PARAMS:
//   - args: the vpc id to get vpc's dhcp info
//
// RETURNS:
//   - *VpcDhcpInfo: the info of the VPC dhcp
//   - error: nil if success otherwise the specific error
func (c *Client) GetVPCDhcpInfo(vpcId string) (*VpcDhcpInfo, error) {
	result := &VpcDhcpInfo{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVPCId(vpcId) + "/dhcp").
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// OpenRelay open vpc relay service
// Param：
//
//	updateVpcRelayArgs
//
// return：
// Return error message; if successful, return nil.
func (c *Client) OpenRelay(updateVpcRelayArgs *UpdateVpcRelayArgs) error {
	if updateVpcRelayArgs == nil {
		return fmt.Errorf("The UpdateVPCReleyArgs cannot be nil.")
	}
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPC()+"/openRelay/"+updateVpcRelayArgs.VpcId).
		WithMethod(http.PUT).
		WithQueryParamFilter("clientToken", updateVpcRelayArgs.ClientToken).
		Do()
}

// ShutdownRelay Shutdown VPC relay service
//
// param：
// updateVpcRelayArgs:
//
//	the vpc id to shut down vpc's relay service
//
// return：
// error: Return error message; if successful, return nil.
func (c *Client) ShutdownRelay(updateVpcRelayArgs *UpdateVpcRelayArgs) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPC()+"/shutdownRelay/"+updateVpcRelayArgs.VpcId).
		WithMethod(http.PUT).
		WithQueryParamFilter("clientToken", updateVpcRelayArgs.ClientToken).
		Do()
}

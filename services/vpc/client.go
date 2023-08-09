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

// client.go - define the client for VPC service

// Package vpc defines the VPC services of BCE.
// The supported APIs are all defined in different files.
package vpc

import (
	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	URI_PREFIX = bce.URI_PREFIX + "v1"

	DEFAULT_ENDPOINT = "bcc." + bce.DEFAULT_REGION + ".baidubce.com"

	REQUEST_VPC_URL              = "/vpc"
	REQUEST_SUBNET_URL           = "/subnet"
	REQUEST_ROUTE_URL            = "/route"
	REQUEST_RULE_URL             = "/rule"
	REQUEST_ACL_URL              = "/acl"
	REQUEST_NAT_URL              = "/nat"
	REQUEST_PEERCONN_URL         = "/peerconn"
	REQUEST_NETWORK_TOPOLOGY_URL = "/topology"
)

// Client of VPC service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getURLForVPC() string {
	return URI_PREFIX + REQUEST_VPC_URL
}

func getURLForVPCId(vpcId string) string {
	return getURLForVPC() + "/" + vpcId
}

func getURLForSubnet() string {
	return URI_PREFIX + REQUEST_SUBNET_URL
}

func getURLForSubnetId(subnetId string) string {
	return getURLForSubnet() + "/" + subnetId
}

func getURLForRouteTable() string {
	return URI_PREFIX + REQUEST_ROUTE_URL
}

func getURLForRouteRule() string {
	return getURLForRouteTable() + REQUEST_RULE_URL
}

func getURLForRouteRuleId(routeRuleId string) string {
	return getURLForRouteRule() + "/" + routeRuleId
}

func getURLForAclEntry() string {
	return URI_PREFIX + REQUEST_ACL_URL
}

func getURLForAclRule() string {
	return URI_PREFIX + REQUEST_ACL_URL + REQUEST_RULE_URL
}

func getURLForAclRuleId(aclRuleId string) string {
	return URI_PREFIX + REQUEST_ACL_URL + REQUEST_RULE_URL + "/" + aclRuleId
}

func getURLForNat() string {
	return URI_PREFIX + REQUEST_NAT_URL
}

func getURLForNatId(natId string) string {
	return getURLForNat() + "/" + natId
}

func getURLForPeerConn() string {
	return URI_PREFIX + REQUEST_PEERCONN_URL
}

func getURLForPeerConnId(peerConnId string) string {
	return getURLForPeerConn() + "/" + peerConnId
}

func getURLForNetworkTopology() string {
	return getURLForVPC() + REQUEST_NETWORK_TOPOLOGY_URL
}

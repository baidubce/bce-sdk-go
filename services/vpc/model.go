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

// model.go - definitions of the request arguments and results data structure model

package vpc

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type (
	SubnetType string

	NexthopType string

	AclRuleProtocolType string

	AclRuleDirectionType string

	AclRuleActionType string

	AclRulePortType string

	NatGatewaySpecType string

	PaymentTimingType string

	PeerConnRoleType string

	NatStatusType string

	PeerConnStatusType string

	DnsStatusType string
)

const (
	SUBNET_TYPE_BCC    SubnetType = "BCC"
	SUBNET_TYPE_BCCNAT SubnetType = "BCC_NAT"
	SUBNET_TYPE_BBC    SubnetType = "BBC"

	NEXTHOP_TYPE_CUSTOM      NexthopType = "custom"
	NEXTHOP_TYPE_VPN         NexthopType = "vpn"
	NEXTHOP_TYPE_NAT         NexthopType = "nat"
	NEXTHOP_TYPE_ETGATEWAY   NexthopType = "dcGateway"
	NEXTHOP_TYPE_PEERCONN    NexthopType = "peerConn"
	NEXTHOP_TYPE_IPV6GATEWAY NexthopType = "ipv6gateway"
	NEXTHOP_TYPE_ENIC        NexthopType = "enic"
	NEXTHOP_TYPE_HAVIP       NexthopType = "havip"

	ACL_RULE_PROTOCOL_TCP  AclRuleProtocolType = "tcp"
	ACL_RULE_PROTOCOL_UDP  AclRuleProtocolType = "udp"
	ACL_RULE_PROTOCOL_ICMP AclRuleProtocolType = "icmp"

	ACL_RULE_DIRECTION_INGRESS AclRuleDirectionType = "ingress"
	ACL_RULE_DIRECTION_EGRESS  AclRuleDirectionType = "egress"

	ACL_RULE_ACTION_ALLOW AclRuleActionType = "allow"
	ACL_RULE_ACTION_DENY  AclRuleActionType = "deny"

	ACL_RULE_PORT_ALL AclRulePortType = "all"

	NAT_GATEWAY_SPEC_SMALL  NatGatewaySpecType = "small"
	NAT_GATEWAY_SPEC_MEDIUM NatGatewaySpecType = "medium"
	NAT_GATEWAY_SPEC_LARGE  NatGatewaySpecType = "large"

	PAYMENT_TIMING_PREPAID  PaymentTimingType = "Prepaid"
	PAYMENT_TIMING_POSTPAID PaymentTimingType = "Postpaid"

	PEERCONN_ROLE_INITIATOR PeerConnRoleType = "initiator"
	PEERCONN_ROLE_ACCEPTOR  PeerConnRoleType = "acceptor"

	NAT_STATUS_BUILDING     NatStatusType = "building"
	NAT_STATUS_UNCONFIGURED NatStatusType = "unconfigured"
	NAT_STATUS_CONFIGURING  NatStatusType = "configuring"
	NAT_STATUS_ACTIVE       NatStatusType = "active"
	NAT_STATUS_STOPPING     NatStatusType = "stopping"
	NAT_STATUS_DOWN         NatStatusType = "down"
	NAT_STATUS_STARTING     NatStatusType = "starting"
	NAT_STATUS_DELETING     NatStatusType = "deleting"
	NAT_STATUS_DELETED      NatStatusType = "deleted"

	PEERCONN_STATUS_CREATING       PeerConnStatusType = "creating"
	PEERCONN_STATUS_CONSULTING     PeerConnStatusType = "consulting"
	PEERCONN_STATUS_CONSULT_FAILED PeerConnStatusType = "consult_failed"
	PEERCONN_STATUS_ACTIVE         PeerConnStatusType = "active"
	PEERCONN_STATUS_DOWN           PeerConnStatusType = "down"
	PEERCONN_STATUS_STARTING       PeerConnStatusType = "starting"
	PEERCONN_STATUS_STOPPING       PeerConnStatusType = "stopping"
	PEERCONN_STATUS_DELETING       PeerConnStatusType = "deleting"
	PEERCONN_STATUS_DELETED        PeerConnStatusType = "deleted"
	PEERCONN_STATUS_EXPIRED        PeerConnStatusType = "expired"
	PEERCONN_STATUS_ERROR          PeerConnStatusType = "error"
	PEERCONN_STATUS_UPDATING       PeerConnStatusType = "updating"

	DNS_STATUS_CLOSE   DnsStatusType = "close"
	DNS_STATUS_WAIT    DnsStatusType = "wait"
	DNS_STATUS_SYNCING DnsStatusType = "syncing"
	DNS_STATUS_OPEN    DnsStatusType = "open"
	DNS_STATUS_CLOSING DnsStatusType = "closing"
)

// CreateVPCArgs defines the structure of the input parameters for the CreateVPC api
type CreateVPCArgs struct {
	ClientToken string           `json:"-"`
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Cidr        string           `json:"cidr"`
	EnableIpv6  bool             `json:"enableIpv6,omitempty"`
	Tags        []model.TagModel `json:"tags,omitempty"`
}

// CreateVPCResult defines the structure of the output parameters for the CreateVPC api
type CreateVPCResult struct {
	VPCID string `json:"vpcId"`
}

// ListVPCArgs defines the structure of the input parameters for the ListVPC api
type ListVPCArgs struct {
	Marker  string
	MaxKeys int
	VpcIds  string
	// IsDefault is a string type,
	// so we can identify if it has been setted when the value is false.
	// NOTE: it can be only true or false.
	IsDefault string
}

// ListVPCResult defines the structure of the output parameters for the ListVPC api
type ListVPCResult struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
	VPCs        []VPC  `json:"vpcs"`
}

type VPC struct {
	VPCID         string           `json:"vpcId"`
	Name          string           `json:"name"`
	Cidr          string           `json:"cidr"`
	Ipv6Cidr      string           `json:"ipv6Cidr"`
	CreatedTime   string           `json:"createdTime"`
	Description   string           `json:"description"`
	IsDefault     bool             `json:"isDefault"`
	SecondaryCidr []string         `json:"secondaryCidr"`
	Tags          []model.TagModel `json:"tags"`
}

// GetVPCDetailResult defines the structure of the output parameters for the GetVPCDetail api
type GetVPCDetailResult struct {
	VPC ShowVPCModel `json:"vpc"`
}

type ShowVPCModel struct {
	VPCId         string           `json:"vpcId"`
	Name          string           `json:"name"`
	Cidr          string           `json:"cidr"`
	Ipv6Cidr      string           `json:"ipv6Cidr"`
	CreatedTime   string           `json:"createdTime"`
	Description   string           `json:"description"`
	IsDefault     bool             `json:"isDefault"`
	Relay         bool             `json:"relay"`
	Subnets       []Subnet         `json:"subnets"`
	SecondaryCidr []string         `json:"secondaryCidr"`
	Tags          []model.TagModel `json:"tags"`
}

type Subnet struct {
	SubnetId              string           `json:"subnetId"`
	Name                  string           `json:"name"`
	ZoneName              string           `json:"zoneName"`
	Cidr                  string           `json:"cidr"`
	Ipv6Cidr              string           `json:"ipv6Cidr"`
	VPCId                 string           `json:"vpcId"`
	SubnetType            SubnetType       `json:"subnetType"`
	Description           string           `json:"description"`
	CreatedTime           string           `json:"createdTime"`
	AvailableIp           int              `json:"availableIp"`
	AvailableUnreservedIp int              `json:"availableUnreservedIp"`
	Tags                  []model.TagModel `json:"tags"`
}

// UpdateVPCArgs defines the structure of the input parameters for the UpdateVPC api
type UpdateVPCArgs struct {
	ClientToken   string   `json:"-"`
	Name          string   `json:"name"`
	Description   string   `json:"description,omitempty"`
	EnableIpv6    *bool    `json:"enableIpv6,omitempty"`
	SecondaryCidr []string `json:"secondaryCidr,omitempty"`
}

// CreateSubnetArgs defines the structure of the input parameters for the CreateSubnet api
type CreateSubnetArgs struct {
	ClientToken      string           `json:"-"`
	Name             string           `json:"name"`
	ZoneName         string           `json:"zoneName"`
	Cidr             string           `json:"cidr"`
	VpcId            string           `json:"vpcId"`
	VpcSecondaryCidr string           `json:"vpcSecondaryCidr,omitempty"`
	SubnetType       SubnetType       `json:"subnetType,omitempty"`
	Description      string           `json:"description,omitempty"`
	EnableIpv6       bool             `json:"enableIpv6,omitempty"`
	Tags             []model.TagModel `json:"tags,omitempty"`
}

// CreateSubnetResult defines the structure of the output parameters for the CreateSubnet api
type CreateSubnetResult struct {
	SubnetId string `json:"subnetId"`
}

// ListSubnetArgs defines the structure of the input parameters for the ListSubnet api
type ListSubnetArgs struct {
	Marker     string
	MaxKeys    int
	VpcId      string
	ZoneName   string
	SubnetType SubnetType
	SubnetIds  string
}

// ListSubnetResult defines the structure of the output parameters for the ListSubnet api
type ListSubnetResult struct {
	Marker      string   `json:"marker"`
	IsTruncated bool     `json:"isTruncated"`
	NextMarker  string   `json:"nextMarker"`
	MaxKeys     int      `json:"maxKeys"`
	Subnets     []Subnet `json:"subnets"`
}

// GetSubnetDetailResult defines the structure of the output parameters for the GetSubnetDetail api
type GetSubnetDetailResult struct {
	Subnet Subnet `json:"subnet"`
}

// UpdateSubnetArgs defines the structure of the input parameters for the UpdateSubnet api
type UpdateSubnetArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	EnableIpv6  *bool  `json:"enableIpv6,omitempty"`
}

// CreateIpreserveArgs - arguments for creating a reserved CIDR
type CreateIpreserveArgs struct {
	SubnetId    string `json:"subnetId"`
	IpCidr      string `json:"ipCidr"`
	IpVersion   int32  `json:"ipVersion"`
	Description string `json:"description"`
	ClientToken string `json:"-"`
}

// CreateIpreserveResult - result of creating a reserved CIDR
type CreateIpreserveResult struct {
	IpReserveId string `json:"ipReserveId"`
}

// ListIpeserveArgs - arguments for listing reserved CIDRs
type ListIpeserveArgs struct {
	SubnetId string `json:"subnetId,omitempty"`
	Marker   string `json:"marker,omitempty"`
	MaxKeys  int    `json:"maxKeys,omitempty"`
}

type IpReserve struct {
	IpReserveId string `json:"ipReserveId"`
	SubnetId    string `json:"subnetId"`
	IpCidr      string `json:"ipCidr"`
	IpVersion   string `json:"ipVersion"`
	Description string `json:"description"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

// ListIpeserveResult - result of listing reserved CIDRs
type ListIpeserveResult struct {
	IpReserves  []IpReserve `json:"ipReserves"`
	Marker      string      `json:"marker,omitempty"`
	IsTruncated bool        `json:"isTruncated"`
	NextMarker  string      `json:"nextMarker,omitempty"`
	MaxKeys     int         `json:"maxKeys,omitempty"`
}

type RouteRule struct {
	RouteRuleId        string      `json:"routeRuleId"`
	RouteTableId       string      `json:"routeTableId"`
	SourceAddress      string      `json:"sourceAddress"`
	DestinationAddress string      `json:"destinationAddress"`
	NexthopId          string      `json:"nexthopId"`
	NexthopType        NexthopType `json:"nexthopType"`
	Description        string      `json:"description"`
	PathType           string      `json:"pathType"`
}

// GetRouteTableResult defines the structure of the output parameters for the GetRouteTableDetail api
type GetRouteTableResult struct {
	RouteTableId string      `json:"routeTableId"`
	VpcId        string      `json:"vpcId"`
	RouteRules   []RouteRule `json:"routeRules"`
}

type GetRouteRuleArgs struct {
	RouteTableId string `json:"routeTableId"`
	VpcId        string `json:"vpcId"`
	Marker       string `json:"marker"`
	MaxKeys      int    `json:"maxKeys"`
}

type GetRouteRuleResult struct {
	Marker      string      `json:"marker"`
	IsTruncated bool        `json:"isTruncated"`
	NextMarker  string      `json:"nextMarker"`
	MaxKeys     int         `json:"maxKeys"`
	RouteRules  []RouteRule `json:"routeRules"`
}

type UpdateRouteRuleArgs struct {
	RouteRuleId        string `json:"routeRuleId"`
	ClientToken        string `json:"-"`
	SourceAddress      string `json:"sourceAddress"`
	DestinationAddress string `json:"destinationAddress"`
	NexthopId          string `json:"nexthopId,omitempty"`
	Description        string `json:"description,omitempty"`
}

// CreateRouteRuleArgs defines the structure of the input parameters for the CreateRouteRule api
type CreateRouteRuleArgs struct {
	ClientToken        string      `json:"-"`
	RouteTableId       string      `json:"routeTableId"`
	SourceAddress      string      `json:"sourceAddress"`
	DestinationAddress string      `json:"destinationAddress"`
	NexthopId          string      `json:"nexthopId,omitempty"`
	IpVersion          string      `json:"ipVersion,omitempty"`
	NexthopType        NexthopType `json:"nexthopType"`
	NextHopList        []NextHop   `json:"nextHopList,omitempty"`
	Description        string      `json:"description,omitempty"`
}

type NextHop struct {
	NexthopId   string      `json:"nexthopId"`
	NexthopType NexthopType `json:"nexthopType"`
	PathType    string      `json:"pathType"`
}

// CreateRouteRuleResult defines the structure of the output parameters for the CreateRouteRule api
type CreateRouteRuleResult struct {
	RouteRuleId  string   `json:"routeRuleId"`
	RouteRuleIds []string `json:"routeRuleIds,omitempty"`
}

// ListAclEntrysResult defines the structure of the output parameters for the ListAclEntrys api
type ListAclEntrysResult struct {
	VpcId     string     `json:"vpcId"`
	VpcName   string     `json:"vpcName"`
	VpcCidr   string     `json:"vpcCidr"`
	AclEntrys []AclEntry `json:"aclEntrys"`
}

type AclEntry struct {
	SubnetId   string    `json:"subnetId"`
	SubnetName string    `json:"subnetName"`
	SubnetCidr string    `json:"subnetCidr"`
	AclRules   []AclRule `json:"aclRules"`
}

type AclRule struct {
	Id                   string               `json:"id"`
	SubnetId             string               `json:"subnetId"`
	Description          string               `json:"description"`
	Protocol             AclRuleProtocolType  `json:"protocol"`
	SourceIpAddress      string               `json:"sourceIpAddress"`
	DestinationIpAddress string               `json:"destinationIpAddress"`
	SourcePort           string               `json:"sourcePort"`
	DestinationPort      string               `json:"destinationPort"`
	Position             int                  `json:"position"`
	IpVersion            int                  `json:"ipVersion"`
	Direction            AclRuleDirectionType `json:"direction"`
	Action               AclRuleActionType    `json:"action"`
}

// CreateAclRuleArgs defines the structure of the input parameters for the CreateAclRule api
type CreateAclRuleArgs struct {
	ClientToken string           `json:"-"`
	AclRules    []AclRuleRequest `json:"aclRules"`
}

type AclRuleRequest struct {
	SubnetId             string               `json:"subnetId"`
	Description          string               `json:"description,omitempty"`
	Protocol             AclRuleProtocolType  `json:"protocol"`
	SourceIpAddress      string               `json:"sourceIpAddress"`
	DestinationIpAddress string               `json:"destinationIpAddress"`
	SourcePort           string               `json:"sourcePort"`
	DestinationPort      string               `json:"destinationPort"`
	Position             int                  `json:"position"`
	Direction            AclRuleDirectionType `json:"direction"`
	Action               AclRuleActionType    `json:"action"`
}

// ListAclRulesArgs defines the structure of the input parameters for the ListAclRules api
type ListAclRulesArgs struct {
	Marker   string `json:"marker"`
	MaxKeys  int    `json:"maxKeys"`
	SubnetId string `json:"subnetId"`
}

// ListAclRulesResult defines the structure of the output parameters for the ListAclRules api
type ListAclRulesResult struct {
	Marker      string    `json:"marker"`
	IsTruncated bool      `json:"isTruncated"`
	NextMarker  string    `json:"nextMarker"`
	MaxKeys     int       `json:"maxKeys"`
	AclRules    []AclRule `json:"aclRules"`
}

// UpdateAclRuleArgs defines the structure of the input parameters for the UpdateAclRule api
type UpdateAclRuleArgs struct {
	ClientToken          string              `json:"-"`
	Description          string              `json:"description,omitempty"`
	Protocol             AclRuleProtocolType `json:"protocol,omitempty"`
	SourceIpAddress      string              `json:"sourceIpAddress,omitempty"`
	DestinationIpAddress string              `json:"destinationIpAddress,omitempty"`
	SourcePort           string              `json:"sourcePort,omitempty"`
	DestinationPort      string              `json:"destinationPort,omitempty"`
	Position             int                 `json:"position,omitempty"`
	Action               AclRuleActionType   `json:"action,omitempty"`
}

// CreateNatGatewayArgs defines the structure of the input parameters for the CreateNatGateway api
type CreateNatGatewayArgs struct {
	ClientToken     string             `json:"-"`
	Name            string             `json:"name"`
	VpcId           string             `json:"vpcId"`
	Spec            NatGatewaySpecType `json:"spec"`
	CuNum           string             `json:"cuNum,omitempty"`
	Eips            []string           `json:"eips,omitempty"`
	DnatEips        []string           `json:"dnatEips,omitempty"`
	BindEips        []string           `json:"bindEips,omitempty"`
	Billing         *Billing           `json:"billing"`
	Tags            []model.TagModel   `json:"tags,omitempty"`
	ResourceGroupId string             `json:"resourceGroupId,omitempty"`
}

type ResizeNatGatewayArgs struct {
	ClientToken string `json:"-"`
	CuNum       int    `json:"cuNum"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type Billing struct {
	PaymentTiming PaymentTimingType `json:"paymentTiming,omitempty"`
	Reservation   *Reservation      `json:"reservation,omitempty"`
}

// CreateNatGatewayResult defines the structure of the output parameters for the CreateNatGateway api
type CreateNatGatewayResult struct {
	NatId string `json:"natId"`
}

// ListNatGatewayArgs defines the structure of the input parameters for the ListNatGateway api
type ListNatGatewayArgs struct {
	VpcId   string
	NatId   string
	Name    string
	Ip      string
	Marker  string
	MaxKeys int
}

// ListNatGatewayResult defines the structure of the output parameters for the ListNatGateway api
type ListNatGatewayResult struct {
	Nats        []NAT  `json:"nats"`
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
}

// NAT is the result for getNatGatewayDetail api.
type NAT struct {
	Id            string           `json:"id"`
	Name          string           `json:"name"`
	NatType       string           `json:"natType"`
	VpcId         string           `json:"vpcId"`
	Spec          string           `json:"spec,omitempty"`
	CuNum         int              `json:"cuNum,omitempty"`
	Status        NatStatusType    `json:"status"`
	Eips          []string         `json:"eips"`
	DnatEips      []string         `json:"dnatEips"`
	BindEips      []string         `json:"bindEips"`
	PaymentTiming string           `json:"paymentTiming"`
	ExpiredTime   string           `json:"expiredTime"`
	Tags          []model.TagModel `json:"tags"`
	IpVersion     *string          `json:"ipVersion,omitempty"`
}

type SnatRule struct {
	RuleId            string   `json:"ruleId"`
	RuleName          string   `json:"ruleName"`
	PublicIpAddresses []string `json:"publicIpsAddress"`
	SourceCIDR        string   `json:"sourceCIDR"`
	Status            string   `json:"status"`
}

type SnatRuleArgs struct {
	RuleName          string   `json:"ruleName,omitempty"`
	SourceCIDR        string   `json:"sourceCIDR,omitempty"`
	PublicIpAddresses []string `json:"publicIpsAddress,omitempty"`
}

type DnatRuleArgs struct {
	RuleName         string `json:"ruleName,omitempty"`
	PublicIpAddress  string `json:"publicIpAddress,omitempty"`
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	PublicPort       string `json:"publicPort,omitempty"`
	PrivatePort      string `json:"privatePort,omitempty"`
	PublicPortRange  string `json:"publicPortRange,omitempty"`
	PrivatePortRange string `json:"privatePortRange,omitempty"`
}

type DnatRule struct {
	RuleId           string `json:"ruleId"`
	RuleName         string `json:"ruleName"`
	PublicIpAddress  string `json:"publicIpAddress"`
	PrivateIpAddress string `json:"privateIpAddress"`
	Protocol         string `json:"protocol"`
	PublicPort       int    `json:"publicPort"`
	PrivatePort      int    `json:"privatePort"`
	Status           string `json:"status"`
	PublicPortRange  string `json:"publicPortRange"`
	PrivatePortRange string `json:"privatePortRange"`
}

// UpdateNatGatewayArgs defines the structure of the input parameters for the UpdateNatGateway api
type UpdateNatGatewayArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
}

// BindEipsArgs defines the structure of the input parameters for the BindEips api
type BindEipsArgs struct {
	ClientToken string   `json:"-"`
	Eips        []string `json:"eips"`
}

// BindDnatEipsArgs defines the structure of the input parameters for the BindDnatEips api
type BindDnatEipsArgs struct {
	ClientToken string   `json:"-"`
	DnatEips    []string `json:"dnatEips"`
}

// EnhanceNatBindEipsArgs defines the structure of the input parameters for the enhanceNatBindEips api
type EnhanceNatBindEipsArgs struct {
	ClientToken string   `json:"-"`
	BindEips    []string `json:"bindEips"`
}

// UnBindEipsArgs defines the structure of the input parameters for the UnBindEips api
type UnBindEipsArgs struct {
	ClientToken string   `json:"-"`
	Eips        []string `json:"eips"`
}

// UnBindDnatEipsArgs defines the structure of the input parameters for the UnBindDnatEips api
type UnBindDnatEipsArgs struct {
	ClientToken string   `json:"-"`
	DnatEips    []string `json:"dnatEips"`
}

// EnhanceNatUnBindEipsArgs defines the structure of the input parameters for the EnhanceNatUnBindEips api
type EnhanceNatUnBindEipsArgs struct {
	ClientToken string   `json:"-"`
	BindEips    []string `json:"bindEips"`
}

// RenewNatGatewayArgs defines the structure of the input parameters for the RenewNatGateway api
type RenewNatGatewayArgs struct {
	ClientToken string   `json:"-"`
	Billing     *Billing `json:"billing"`
}

type CreateNatGatewaySnatRuleArgs struct {
	ClientToken       string   `json:"-"`
	RuleName          string   `json:"ruleName,omitempty"`
	SourceCIDR        string   `json:"sourceCIDR,omitempty"`
	PublicIpAddresses []string `json:"publicIpsAddress,omitempty"`
}

type BatchCreateNatGatewaySnatRuleArgs struct {
	ClientToken string         `json:"-"`
	NatId       string         `json:"natId"`
	SnatRules   []SnatRuleArgs `json:"snatRules"`
}

type UpdateNatGatewaySnatRuleArgs struct {
	ClientToken       string   `json:"-"`
	RuleName          string   `json:"ruleName,omitempty"`
	SourceCIDR        string   `json:"sourceCIDR,omitempty"`
	PublicIpAddresses []string `json:"publicIpsAddress,omitempty"`
}

type ListNatGatewaySnatRuleArgs struct {
	NatId   string `json:"natId"`
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListNatGatewaySnatRulesResult struct {
	Rules       []SnatRule `json:"rules"`
	Marker      string     `json:"marker"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	MaxKeys     int        `json:"maxKeys"`
}

type CreateNatGatewaySnatRuleResult struct {
	RuleId string `json:"ruleId"`
}

type BatchCreateNatGatewaySnatRuleResult struct {
	SnatRuleIds []string `json:"snatRuleIds"`
}

type CreateNatGatewayDnatRuleArgs struct {
	ClientToken      string `json:"-"`
	RuleName         string `json:"ruleName,omitempty"`
	PublicIpAddress  string `json:"publicIpAddress,omitempty"`
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	PublicPort       string `json:"publicPort,omitempty"`
	PrivatePort      string `json:"privatePort,omitempty"`
	PublicPortRange  string `json:"publicPortRange,omitempty"`
	PrivatePortRange string `json:"privatePortRange,omitempty"`
}

type BatchCreateNatGatewayDnatRuleArgs struct {
	ClientToken string         `json:"-"`
	Rules       []DnatRuleArgs `json:"rules"`
}

type UpdateNatGatewayDnatRuleArgs struct {
	ClientToken      string `json:"-"`
	RuleName         string `json:"ruleName,omitempty"`
	PublicIpAddress  string `json:"publicIpAddress,omitempty"`
	PrivateIpAddress string `json:"privateIpAddress,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	PublicPort       string `json:"publicPort,omitempty"`
	PrivatePort      string `json:"privatePort,omitempty"`
	PublicPortRange  string `json:"publicPortRange,omitempty"`
	PrivatePortRange string `json:"privatePortRange,omitempty"`
}

type ListNatGatewaDnatRuleArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListNatGatewayDnatRulesResult struct {
	Rules       []DnatRule `json:"rules"`
	Marker      string     `json:"marker"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	MaxKeys     int        `json:"maxKeys"`
}

type CreateNatGatewayDnatRuleResult struct {
	RuleId string `json:"ruleId"`
}

type BatchCreateNatGatewayDnatRuleResult struct {
	RuleIds []string `json:"ruleIds"`
}

// CreatePeerConnArgs defines the structure of the input parameters for the CreatePeerConn api
type CreatePeerConnArgs struct {
	ClientToken     string           `json:"-"`
	BandwidthInMbps int              `json:"bandwidthInMbps"`
	Description     string           `json:"description,omitempty"`
	LocalIfName     string           `json:"localIfName,omitempty"`
	LocalVpcId      string           `json:"localVpcId"`
	PeerAccountId   string           `json:"peerAccountId,omitempty"`
	PeerVpcId       string           `json:"peerVpcId"`
	PeerRegion      string           `json:"peerRegion"`
	PeerIfName      string           `json:"peerIfName,omitempty"`
	Billing         *Billing         `json:"billing"`
	Tags            []model.TagModel `json:"tags,omitempty"`
	ResourceGroupId string           `json:"resourceGroupId,omitempty"`
}

// CreatePeerConnResult defines the structure of the output parameters for the CreatePeerConn api
type CreatePeerConnResult struct {
	PeerConnId string `json:"peerConnId"`
}

// ListPeerConnsArgs defines the structure of the input parameters for the ListPeerConns api
type ListPeerConnsArgs struct {
	VpcId   string
	Marker  string
	MaxKeys int
}

// ListPeerConnsResult defines the structure of the output parameters for the ListPeerConns api
type ListPeerConnsResult struct {
	PeerConns   []PeerConn `json:"peerConns"`
	Marker      string     `json:"marker"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	MaxKeys     int        `json:"maxKeys"`
}

type PeerConn struct {
	PeerConnId      string             `json:"peerConnId"`
	Role            PeerConnRoleType   `json:"role"`
	Status          PeerConnStatusType `json:"status"`
	BandwidthInMbps int                `json:"bandwidthInMbps"`
	Description     string             `json:"description"`
	LocalIfId       string             `json:"localIfId"`
	LocalIfName     string             `json:"localIfName"`
	LocalVpcId      string             `json:"localVpcId"`
	LocalRegion     string             `json:"localRegion"`
	PeerVpcId       string             `json:"peerVpcId"`
	PeerRegion      string             `json:"peerRegion"`
	PeerAccountId   string             `json:"peerAccountId"`
	PaymentTiming   string             `json:"paymentTiming"`
	DnsStatus       DnsStatusType      `json:"dnsStatus"`
	CreatedTime     string             `json:"createdTime"`
	ExpiredTime     string             `json:"expiredTime"`
	Tags            []model.TagModel   `json:"tags"`
	DeleteProtect   bool               `json:"deleteProtect"`
}

// UpdatePeerConnArgs defines the structure of the input parameters for the UpdatePeerConn api
type UpdatePeerConnArgs struct {
	LocalIfId   string `json:"localIfId"`
	Description string `json:"description,omitempty"`
	LocalIfName string `json:"localIfName,omitempty"`
}

// ResizePeerConnArgs defines the structure of the input parameters for the ResizePeerConn api
type ResizePeerConnArgs struct {
	NewBandwidthInMbps int    `json:"newBandwidthInMbps"`
	ClientToken        string `json:"-"`
}

// RenewPeerConnArgs defines the structure of the input parameters for the RenewPeerConn api
type RenewPeerConnArgs struct {
	Billing     *Billing `json:"billing"`
	ClientToken string   `json:"-"`
}

// PeerConnSyncDNSArgs defines the structure of the input parameters for the PeerConnSyncDNS api
type PeerConnSyncDNSArgs struct {
	Role        PeerConnRoleType `json:"role"`
	ClientToken string           `json:"-"`
}

type UpdatePeerConnDeleteProtectArgs struct {
	DeleteProtect bool   `json:"deleteProtect"`
	ClientToken   string `json:"-"`
}

// GetResourceIpArgs defines the structure of the input parameters for the GetResourceIp api
type GetResourceIpArgs struct {
	VpcId        string
	SubnetId     string
	ResourceType string
	PageNo       int
	PageSize     int
}

// ResourceIp defines the structure of resource IP information
type ResourceIp struct {
	Ip           string `json:"ip"`
	ResourceType string `json:"resourceType"`
}

// GetResourceIpResult defines the structure of the output parameters for the GetResourceIp api
type GetResourceIpResult struct {
	Result     []ResourceIp `json:"result"`
	PageNo     int          `json:"pageNo"`
	PageSize   int          `json:"pageSize"`
	TotalCount int          `json:"totalCount"`
}

/*
Get VpcPrivateIpAddressedInfo args

	VpcId:the vpc you want to query ips
	PrivateIpAddresses:the privateIp list you want to query
	PrivateIpRange:the range of privateIp .ex:"192.168.0.1-192.168.0.5"
	pay attention that the size of PrivateIpAddresses and PrivateIpRange must less than 100
	if both PrivateIpAddresses and PrivateIpRange ,the PrivateIpRange will effect
*/
type GetVpcPrivateIpArgs struct {
	VpcId              string   `json:"vpcId"`
	PrivateIpAddresses []string `json:"privateIpAddresses,omitempty"`
	PrivateIpRange     string   `json:"privateIpRange,omitempty"`
}

type VpcPrivateIpAddress struct {
	PrivateIpAddress     string `json:"privateIpAddress"`
	Cidr                 string `json:"cidr"`
	PrivateIpAddressType string `json:"privateIpAddressType"`
	CreatedTime          string `json:"createdTime"`
}

// VpcPrivateIpAddressesResult defines the structure of the output parameters for the GetPrivateIpAddressInfo api
type VpcPrivateIpAddressesResult struct {
	VpcPrivateIpAddresses []VpcPrivateIpAddress `json:"vpcPrivateIpAddresses"`
}

/*
Get NetworkTopologyInfo args

	HostIp:the host ip of the network topology to be queried
	HostId:the host id of the network topology to be queried
	If both HostIp and HostId are passed in, the HostIp will effect (only need to pass in one of the two)
*/
type GetNetworkTopologyArgs struct {
	HostIp string `json:"hostIp,omitempty"`
	HostId string `json:"hostId,omitempty"`
}

type NetworkTopology struct {
	ClusterName string `json:"clusterName"`
	HostId      string `json:"hostId"`
	SwitchId    string `json:"switchId"`
	HostIp      string `json:"hostIp"`
	PodName     string `json:"podName"`
}

// NetworkTopologyResult defines the structure of the output parameters for the GetNetworkTopologyInfo api
type NetworkTopologyResult struct {
	NetworkTopologies []NetworkTopology `json:"networkTopologies"`
}

type CreateProbeArgs struct {
	ClientToken string   `json:"-"`
	Name        string   `json:"name"`
	VpcId       string   `json:"vpcId"`
	SubnetId    string   `json:"subnetId"`
	Protocol    string   `json:"protocol"`
	Frequency   int      `json:"frequency"`
	DestIp      string   `json:"destIp"`
	DestPort    string   `json:"destPort"`
	SourceIps   []string `json:"sourceIps"`
	SourceIpNum int      `json:"sourceIpNum,omitempty"`
	Description string   `json:"description,omitempty"`
	Payload     string   `json:"payload,omitempty"`
}

type CreateProbeResult struct {
	ProbeId string `json:"probeId"`
}

type UpdateProbeArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name,omitempty"`
	Frequency   int    `json:"frequency,omitempty"`
	DestIp      string `json:"destIp,omitempty"`
	DestPort    string `json:"destPort,omitempty"`
	Description string `json:"description,omitempty"`
	Payload     string `json:"payload,omitempty"`
}

type UpdateProbeResult struct {
	ProbeId string `json:"probeId"`
}

type ListProbesArgs struct {
	Marker  string
	MaxKeys int
}

type Probe struct {
	ProbeId     string   `json:"probeId"`
	Name        string   `json:"name"`
	VpcId       string   `json:"vpcId"`
	SubnetId    string   `json:"subnetId"`
	Protocol    string   `json:"protocol"`
	Frequency   int      `json:"frequency"`
	DestIp      string   `json:"destIp"`
	DestPort    string   `json:"destPort"`
	SourceIps   []string `json:"sourceIps"`
	SourceIpNum int      `json:"sourceIpNum,omitempty"`
	Description string   `json:"description,omitempty"`
	Payload     string   `json:"payload,omitempty"`
	Status      string   `json:"status"`
}

type ListProbesResult struct {
	Probes      []Probe `json:"probes"`
	Marker      string  `json:"marker"`
	IsTruncated bool    `json:"isTruncated"`
	NextMarker  string  `json:"nextMarker"`
	MaxKeys     int     `json:"maxKeys"`
}

type GetProbeDetailResult struct {
	Probe Probe `json:"probes"`
}

// CreateIPv6GatewayArgs defines the structure of the input parameters for the CreateIPv6Gateway api
type CreateIPv6GatewayArgs struct {
	ClientToken     string           `json:"-"`
	Name            string           `json:"name"`
	VpcId           string           `json:"vpcId"`
	BandwidthInMbps int              `json:"bandwidthInMbps"`
	Billing         *Billing         `json:"billing"`
	Tags            []model.TagModel `json:"tags,omitempty"`
}

// CreateIPv6GatewayResult defines the structure of the output parameters for the CreateIPv6Gateway api
type CreateIPv6GatewayResult struct {
	GatewayId string `json:"gatewayId"`
}

// ListIPv6GatewayArgs defines the structure of the input parameters for the ListIPv6Gateway api
type ListIPv6GatewayArgs struct {
	VpcId string `json:"vpcId"`
}

// DeleteIPv6GatewayArgs defines the structure of the input parameters for the DeleteIPv6Gateway api
type DeleteIPv6GatewayArgs struct {
	ClientToken string `json:"-"`
}

// ListIPv6GatewayResult defines the structure of the output parameters for the ListIPv6Gateway api
type ListIPv6GatewayResult struct {
	GatewayId       string           `json:"gatewayId"`
	Name            string           `json:"name"`
	BandwidthInMbps int              `json:"bandwidthInMbps"`
	VpcId           string           `json:"vpcId"`
	EgressOnlyRules []EgressOnlyRule `json:"egressOnlyRules"`
	RateLimitRules  []RateLimitRule  `json:"rateLimitRules"`
}

type EgressOnlyRule struct {
	EgressOnlyRuleId string `json:"egressOnlyRuleId"`
	Cidr             string `json:"cidr"`
}

type RateLimitRule struct {
	RateLimitRuleId        string `json:"rateLimitRuleId"`
	IPv6Address            string `json:"ipv6Address"`
	IngressBandwidthInMbps int    `json:"ingressBandwidthInMbps"`
	EgressBandwidthInMbps  int    `json:"egressBandwidthInMbps"`
}

// ResizeIPv6GatewayArgs defines the structure of the input parameters for the ResizeIPv6Gateway api
type ResizeIPv6GatewayArgs struct {
	ClientToken     string `json:"-"`
	BandwidthInMbps int    `json:"bandwidthInMbps"`
}

// CreateIPv6GatewayEgressOnlyRuleArgs defines the structure of the input parameters for the CreateIPv6GatewayEgressOnlyRule api
type CreateIPv6GatewayEgressOnlyRuleArgs struct {
	ClientToken string `json:"-"`
	Cidr        string `json:"cidr"`
}

// CreateIPv6GatewayEgressOnlyRuleResult defines the structure of the output parameters for the CreateIPv6GatewayEgressOnlyRule api
type CreateIPv6GatewayEgressOnlyRuleResult struct {
	EgressOnlyRuleId string `json:"egressOnlyRuleId"`
}

// ListIPv6GatewayEgressOnlyRuleArgs defines the structure of the input parameters for the ListIPv6GatewayEgressOnlyRule api
type ListIPv6GatewayEgressOnlyRuleArgs struct {
	Marker  string
	MaxKeys int
}

// DeleteIPv6GatewayEgressOnlyRuleArgs defines the structure of the input parameters for the DeleteIPv6GatewayEgressOnlyRule api
type DeleteIPv6GatewayEgressOnlyRuleArgs struct {
	ClientToken string `json:"-"`
}

// ListIPv6GatewayEgressOnlyRuleResult defines the structure of the output parameters for the ListIPv6GatewayEgressOnlyRule api
type ListIPv6GatewayEgressOnlyRuleResult struct {
	Marker          string           `json:"marker"`
	IsTruncated     bool             `json:"isTruncated"`
	NextMarker      string           `json:"nextMarker"`
	MaxKeys         int              `json:"maxKeys"`
	EgressOnlyRules []EgressOnlyRule `json:"egressOnlyRules"`
}

// CreateIPv6GatewayRateLimitRuleArgs defines the structure of the input parameters for the CreateIPv6GatewayRateLimitRule api
type CreateIPv6GatewayRateLimitRuleArgs struct {
	ClientToken            string `json:"-"`
	IPv6Address            string `json:"ipv6Address"`
	IngressBandwidthInMbps int    `json:"ingressBandwidthInMbps"`
	EgressBandwidthInMbps  int    `json:"egressBandwidthInMbps"`
}

// CreateIPv6GatewayRateLimitRuleResult defines the structure of the output parameters for the CreateIPv6GatewayRateLimitRule api
type CreateIPv6GatewayRateLimitRuleResult struct {
	RateLimitRuleId string `json:"rateLimitRuleId"`
}

// ListIPv6GatewayRateLimitRuleArgs defines the structure of the input parameters for the ListIPv6GatewayRateLimitRule api
type ListIPv6GatewayRateLimitRuleArgs struct {
	Marker  string
	MaxKeys int
}

// DeleteIPv6GatewayRateLimitRuleArgs defines the structure of the input parameters for the DeleteIPv6GatewayRateLimitRule api
type DeleteIPv6GatewayRateLimitRuleArgs struct {
	ClientToken string `json:"-"`
}

// ListIPv6GatewayRateLimitRuleResult defines the structure of the output parameters for the ListIPv6GatewayRateLimitRule api
type ListIPv6GatewayRateLimitRuleResult struct {
	Marker         string          `json:"marker"`
	IsTruncated    bool            `json:"isTruncated"`
	NextMarker     string          `json:"nextMarker"`
	MaxKeys        int             `json:"maxKeys"`
	RateLimitRules []RateLimitRule `json:"rateLimitRules"`
}

// UpdateIPv6GatewayRateLimitRuleArgs defines the structure of the input parameters for the UpdateIPv6GatewayRateLimitRule api
type UpdateIPv6GatewayRateLimitRuleArgs struct {
	ClientToken            string `json:"-"`
	IngressBandwidthInMbps int    `json:"ingressBandwidthInMbps"`
	EgressBandwidthInMbps  int    `json:"egressBandwidthInMbps"`
}

type CreateVpcDhcpArgs struct {
	ClientToken       string `json:"-"`
	DomainNameServers string `json:"domainNameServers"`
}

type UpdateVpcDhcpArgs struct {
	ClientToken       string `json:"-"`
	DomainNameServers string `json:"domainNameServers"`
}

type VpcDhcpInfo struct {
	VPCID       string      `json:"vpcId"`
	DhcpOptions DhcpOptions `json:"dhcpOptions"`
}

type DhcpOptions struct {
	DomainName        string `json:"domainName"`
	DomainNameServers string `json:"domainNameServers"`
}

type UpdateVpcRelayArgs struct {
	ClientToken string `json:"-"`
	VpcId       string `json:"vpcId"`
}

type TemplateIpAddressInfo struct {
	IpAddress   string `json:"ipAddress"`
	Description string `json:"description"`
}

type IpSet struct {
	IpSetId           string                  `json:"ipSetId"`
	Name              string                  `json:"name"`
	Description       string                  `json:"description"`
	IpVersion         string                  `json:"ipVersion"`
	IpAddressInfo     []TemplateIpAddressInfo `json:"ipAddressInfo"`
	BindedInstanceNum int                     `json:"bindedInstanceNum"`
}

type IpCollectionBindedInstance struct {
	InstanceId   string `json:"instanceId"`
	InstanceType string `json:"instanceType"`
}

type IpGroup struct {
	IpGroupId         string   `json:"ipGroupId"`
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	IpVersion         string   `json:"ipVersion"`
	IpSetIds          []string `json:"ipSetIds"`
	BindedInstanceNum int      `json:"bindedInstanceNum"`
}

type CreateIpSetArgs struct {
	ClientToken   string                  `json:"-"`
	Name          string                  `json:"name"`
	IpVersion     string                  `json:"ipVersion"`
	IpAddressInfo []TemplateIpAddressInfo `json:"ipAddressInfo"`
	Description   string                  `json:"description"`
}

type CreateIpSetResult struct {
	IpSetId string `json:"ipSetId"`
}

type AddIpAddress2IpSetArgs struct {
	ClientToken   string                  `json:"-"`
	IpAddressInfo []TemplateIpAddressInfo `json:"ipAddressInfo"`
}

type DeleteIpAddressArgs struct {
	ClientToken   string   `json:"-"`
	IpAddressInfo []string `json:"ipAddressInfo"`
}

type UpdateIpSetArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteIpSetArgs struct {
	ClientToken string `json:"-"`
}

type ListIpSetArgs struct {
	IpVersion string
	Marker    string
	MaxKeys   int
}

type ListIpSetResult struct {
	IpSets      []IpSet `json:"ipSets"`
	Marker      string  `json:"marker"`
	IsTruncated bool    `json:"isTruncated"`
	NextMarker  string  `json:"nextMarker"`
	MaxKeys     int     `json:"maxKeys"`
}

type GetIpSetDetailResult struct {
	IpSetId         string                       `json:"ipSetId"`
	Name            string                       `json:"name"`
	Description     string                       `json:"description"`
	IpVersion       string                       `json:"ipVersion"`
	IpAddressInfo   []TemplateIpAddressInfo      `json:"ipAddressInfo"`
	BindedInstances []IpCollectionBindedInstance `json:"bindedInstances"`
}

type CreateIpGroupArgs struct {
	ClientToken string   `json:"-"`
	Name        string   `json:"name"`
	IpVersion   string   `json:"ipVersion"`
	IpSetIds    []string `json:"ipSetIds"`
	Description string   `json:"description"`
}

type CreateIpGroupResult struct {
	IpGroupId string `json:"ipGroupId"`
}

type AddIpSet2IpGroupArgs struct {
	ClientToken string   `json:"-"`
	IpSetIds    []string `json:"ipSetIds"`
}

type UnbindIpSetArgs struct {
	ClientToken string   `json:"-"`
	IpSetIds    []string `json:"ipSetIds"`
}

type UpdateIpGroupArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteIpGroupArgs struct {
	ClientToken string `json:"-"`
}

type ListIpGroupArgs struct {
	IpVersion string
	Marker    string
	MaxKeys   int
}

type ListIpGroupResult struct {
	IpGroups    []IpGroup `json:"ipGroups"`
	Marker      string    `json:"marker"`
	IsTruncated bool      `json:"isTruncated"`
	NextMarker  string    `json:"nextMarker"`
	MaxKeys     int       `json:"maxKeys"`
}

type GetIpGroupDetailResult struct {
	IpGroupId       string                       `json:"ipGroupId"`
	Name            string                       `json:"name"`
	Description     string                       `json:"description"`
	IpVersion       string                       `json:"ipVersion"`
	IpSetIds        []string                     `json:"ipSetIds"`
	BindedInstances []IpCollectionBindedInstance `json:"bindedInstances"`
}

/*
 * Copyright 2020 Baidu, Inc.
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

package vpn

type (
	PaymentTimingType string
	PeerConnRoleType  string
	VpnStatusType     string
)

const (
	PAYMENT_TIMING_PREPAID  PaymentTimingType = "Prepaid"
	PAYMENT_TIMING_POSTPAID PaymentTimingType = "Postpaid"

	VPN_STATUS_BUILDING     VpnStatusType = "building"
	VPN_STATUS_UNCONFIGURED VpnStatusType = "unconfigured"
	VPN_STATUS_CONFIGURING  VpnStatusType = "configuring"
	VPN_STATUS_ACTIVE       VpnStatusType = "active"
)

// CreateVpnGatewayArgs defines the structure of the input parameters for the CreateVpnGateway api
type CreateVpnGatewayArgs struct {
	ClientToken string   `json:"-"`
	VpnName     string   `json:"vpnName"`
	VpcId       string   `json:"vpcId"`
	Description string   `json:"description,omitempty"`
	Eip         string   `json:"eip,omitempty"`
	Billing     *Billing `json:"billing"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type Billing struct {
	PaymentTiming PaymentTimingType `json:"paymentTiming,omitempty"`
	Reservation   *Reservation      `json:"reservation,omitempty"`
}

// CreateVpnGatewayResult defines the structure of the output parameters for the CreateVpnGateway api
type CreateVpnGatewayResult struct {
	VpnId string `json:"vpnId"`
}

// ListVpnGatewayArgs defines the structure of the input parameters for the ListVpnGateway api
type ListVpnGatewayArgs struct {
	VpcId   string
	Eip     string
	Marker  string
	MaxKeys int
}

// ListVpnGatewayResult defines the structure of the output parameters for the ListVpnGateway api
type ListVpnGatewayResult struct {
	Vpns        []VPN  `json:"vpns"`
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
}

// VPN is the result for getVpnDetail api.
type VPN struct {
	Status          VpnStatusType `json:"status"`
	Eip             string        `json:"eip"`
	VpnId           string        `json:"vpnId"`
	VpcId           string        `json:"vpcId"`
	Description     string        `json:"description"`
	ExpiredTime     string        `json:"expiredTime"`
	ProductType     string        `json:"paymentTiming"`
	VpnConnNum      int           `json:"vpnConnNum"`
	BandwidthInMbps int           `json:"bandwidthInMbps"`
	VpnConns        []VpnConn     `json:"vpnConns"`
	Name            string        `json:"vpnName"`
}

// UpdateVpnGatewayArgs defines the structure of the input parameters for the UpdateVpnGateway api
type UpdateVpnGatewayArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"vpnName"`
}

// BindEipArgs defines the structure of the input parameters for the BindEip api
type BindEipArgs struct {
	ClientToken string `json:"-"`
	Eip         string `json:"eip"`
}

type VpnConn struct {
	VpnId         string      `json:"vpnId"`
	VpnConnId     string      `json:"vpnConnId"`
	VpnConnName   string      `json:"vpnConnName"`
	LocalIp       string      `json:"localIp"`
	SecretKey     string      `json:"secretKey"`
	LocalSubnets  []string    `json:"localSubnets"`
	RemoteIp      string      `json:"remoteIp"`
	RemoteSubnets []string    `json:"remoteSubnets"`
	Description   string      `json:"description"`
	Status        string      `json:"status"`
	CreatedTime   string      `json:"createdTime"`
	HealthStatus  string      `json:"healthStatus"`
	IkeConfig     IkeConfig   `json:"ikeConfig"`
	IpsecConfig   IpsecConfig `json:"ipsecConfig"`
}

type IkeConfig struct {
	IkeVersion  string `json:"ikeVersion"`
	IkeMode     string `json:"ikeMode"`
	IkeEncAlg   string `json:"ikeEncAlg"`
	IkeAuthAlg  string `json:"ikeAuthAlg"`
	IkePfs      string `json:"ikePfs"`
	IkeLifeTime string `json:"ikeLifeTime"`
}
type IpsecConfig struct {
	IpsecEncAlg   string `json:"ipsecEncAlg"`
	IpsecAuthAlg  string `json:"ipsecAuthAlg"`
	IpsecPfs      string `json:"ipsecPfs"`
	IpsecLifetime string `json:"ipsecLifetime"`
}

// RenewVpnGatewayArgs defines the structure of the input parameters for the RenewVpnGateway api
type RenewVpnGatewayArgs struct {
	ClientToken string   `json:"-"`
	Billing     *Billing `json:"billing"`
}

type CreateIkeConfig struct {
	IkeVersion  string `json:"ikeVersion"`
	IkeMode     string `json:"ikeMode"`
	IkeEncAlg   string `json:"ikeEncAlg"`
	IkeAuthAlg  string `json:"ikeAuthAlg"`
	IkePfs      string `json:"ikePfs"`
	IkeLifeTime int    `json:"ikeLifeTime"`
}
type CreateIpsecConfig struct {
	IpsecEncAlg   string `json:"ipsecEncAlg"`
	IpsecAuthAlg  string `json:"ipsecAuthAlg"`
	IpsecPfs      string `json:"ipsecPfs"`
	IpsecLifetime int    `json:"ipsecLifetime"`
}

// CreateVpnConnArgs defines the structure of the input parameters for the CreateVpnGatewayConn api
type CreateVpnConnArgs struct {
	ClientToken       string             `json:"-"`
	VpnId             string             `json:"vpnId"`
	VpnConnName       string             `json:"vpnConnName"`
	LocalIp           string             `json:"localIp"`
	SecretKey         string             `json:"secretKey"`
	LocalSubnets      []string           `json:"localSubnets"`
	RemoteIp          string             `json:"remoteIp"`
	RemoteSubnets     []string           `json:"remoteSubnets"`
	Description       string             `json:"description,omitempty"`
	CreateIkeConfig   *CreateIkeConfig   `json:"ikeConfig"`
	CreateIpsecConfig *CreateIpsecConfig `json:"ipsecConfig"`
}

// CreateVpnConnResult defines the structure of the output parameters for the CreateVpnConn api
type CreateVpnConnResult struct {
	VpnConnId string `json:"vpnConnId"`
}

// UpdateVpnConnArgs defines the structure of input parameters for the UpdateVpnConn api
type UpdateVpnConnArgs struct {
	vpnConnId     string             `json:"vpnConnId"`
	updateVpnconn *CreateVpnConnArgs `json:"updateVpnconn"`
}

// ListVpnConnResult defines the structure of output parameters for the ListVpnConn api
type ListVpnConnResult struct {
	VpnConns []VpnConn `json:"vpnConns"`
}

// SslVpnServer defines the structure of the output parameters for the GetSslVpnServer api
type SslVpnServer struct {
	VpnId            string   `json:"vpnId"`
	SslVpnServerId   string   `json:"sslVpnServerId"`
	SslVpnServerName string   `json:"sslVpnServerName"`
	InterfaceType    string   `json:"interfaceType"`
	Status           string   `json:"status"`
	LocalSubnets     []string `json:"localSubnets"`
	RemoteSubnet     string   `json:"remoteSubnet"`
	ClientDns        string   `json:"clientDns"`
	MaxConnection    int      `json:"maxConnection"`
}

// CreateSslVpnServerArgs defines the structure of the input parameters for the CreateSslVpnServer api
type CreateSslVpnServerArgs struct {
	ClientToken      string   `json:"-"`
	VpnId            string   `json:"vpnId"`
	SslVpnServerName string   `json:"sslVpnServerName"`
	InterfaceType    *string  `json:"interfaceType,omitempty"`
	LocalSubnets     []string `json:"localSubnets"`
	RemoteSubnet     string   `json:"remoteSubnet"`
	ClientDns        *string  `json:"clientDns,omitempty"`
}

// UpdateSslVpnServer defines part of the structure of the input parameters for the UpdateSslVpnServer api
type UpdateSslVpnServer struct {
	SslVpnServerName string   `json:"sslVpnServerName,omitempty"`
	LocalSubnets     []string `json:"localSubnets,omitempty"`
	RemoteSubnet     string   `json:"remoteSubnet,omitempty"`
	ClientDns        *string  `json:"clientDns,omitempty"`
}

// CreateSslVpnServerResult defines the structure of the output parameters for the CreateSslVpnServer api
type CreateSslVpnServerResult struct {
	SslVpnServerId string `json:"sslVpnServerId"`
}

// UpdateSslVpnServerArgs defines the structure of input parameters for the UpdateSslVpnServer api
type UpdateSslVpnServerArgs struct {
	ClientToken        string              `json:"-"`
	VpnId              string              `json:"VpnId"`
	SslVpnServerId     string              `json:"sslVpnServerId"`
	UpdateSslVpnServer *UpdateSslVpnServer `json:"updateSslVpnServer"`
}

// ListSslVpnServerResult defines the structure of output parameters for the ListSslVpnServer api
type ListSslVpnServerResult struct {
	SslVpnServers []SslVpnServer `json:"sslVpnServers"`
}

type SslVpnUser struct {
	UserName    string  `json:"userName"`
	Password    string  `json:"password"`
	Description *string `json:"description,omitempty"`
}

type SelectSslVpnUser struct {
	UserId      string `json:"userId"`
	UserName    string `json:"userName"`
	Description string `json:"description"`
}

type UpdateSslVpnUser struct {
	Password    string  `json:"password,omitempty"`
	Description *string `json:"description,omitempty"`
}

// BatchCreateSslVpnUserArgs defines the structure of the input parameters for the CreateSslVpnUser api
type BatchCreateSslVpnUserArgs struct {
	ClientToken string       `json:"-"`
	VpnId       string       `json:"vpnId"`
	SslVpnUsers []SslVpnUser `json:"sslVpnUsers"`
}

// BatchCreateSslVpnUserResult defines the structure of the output parameters for the BatchCreateSslVpnUser api
type BatchCreateSslVpnUserResult struct {
	SslVpnUserIds []string `json:"sslVpnUserIds"`
}

// UpdateSslVpnUserArgs defines the structure of input parameters for the UpdateSslVpnUser api
type UpdateSslVpnUserArgs struct {
	ClientToken string            `json:"-"`
	VpnId       string            `json:"vpnId"`
	UserId      string            `json:"userId"`
	SslVpnUser  *UpdateSslVpnUser `json:"sslVpnUser"`
}

// ListSslVpnUserArgs defines the structure of input parameters for the ListSslVpnUser api
type ListSslVpnUserArgs struct {
	Marker   string `json:"marker"`
	MaxKeys  int    `json:"maxKeys"`
	VpnId    string `json:"vpnId"`
	UserName string `json:"username"`
}

// ListSslVpnUserResult defines the structure of output parameters for the ListSslVpnUser api
type ListSslVpnUserResult struct {
	Marker      string             `json:"marker"`
	IsTruncated bool               `json:"isTruncated"`
	NextMarker  string             `json:"nextMarker"`
	MaxKeys     int                `json:"maxKeys"`
	SslVpnUsers []SelectSslVpnUser `json:"sslVpnUsers"`
}

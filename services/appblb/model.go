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

package appblb

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type BLBStatus string

const (
	BLBStatusCreating    BLBStatus = "creating"
	BLBStatusAvailable   BLBStatus = "available"
	BLBStatusUpdating    BLBStatus = "updating"
	BLBStatusPaused      BLBStatus = "paused"
	BLBStatusUnavailable BLBStatus = "unavailable"
)

type AppRsPortModel struct {
	ListenerPort        int    `json:"listenerPort"`
	BackendPort         string `json:"backendPort"`
	PortType            string `json:"portType"`
	HealthCheckPortType string `json:"healthCheckPortType"`
	Status              string `json:"status"`
	PortId              string `json:"portId"`
	PolicyId            string `json:"policyId"`
}

type AppBackendServer struct {
	InstanceId string           `json:"instanceId,omitempty"`
	Weight     *int             `json:"weight"`
	PrivateIp  string           `json:"privateIp,omitempty"`
	PortList   []AppRsPortModel `json:"portList,omitempty"`
}

type DescribeResultMeta struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
}

type CreateAppServerGroupArgs struct {
	Name              string             `json:"name,omitempty"`
	Description       string             `json:"desc,omitempty"`
	BackendServerList []AppBackendServer `json:"backendServerList,omitempty"`
	ClientToken       string             `json:"-"`
}

type CreateAppServerGroupResult struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Status      BLBStatus `json:"status"`
}

type UpdateAppServerGroupArgs struct {
	SgId        string `json:"sgId"`
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
	ClientToken string `json:"-"`
}

type DescribeAppServerGroupArgs struct {
	Name         string
	ExactlyMatch bool
	Marker       string
	MaxKeys      int
}

type AppServerGroupPort struct {
	Id                          string    `json:"id"`
	Port                        int       `json:"port"`
	Type                        string    `json:"type"`
	Status                      BLBStatus `json:"status"`
	EnableHealthCheck           bool      `json:"enableHealthCheck"`
	HealthCheck                 string    `json:"healthCheck"`
	HealthCheckPort             int       `json:"healthCheckPort"`
	HealthCheckHost             string    `json:"healthCheckHost"`
	HealthCheckTimeoutInSecond  int       `json:"healthCheckTimeoutInSecond"`
	HealthCheckIntervalInSecond int       `json:"healthCheckIntervalInSecond"`
	HealthCheckDownRetry        int       `json:"healthCheckDownRetry"`
	HealthCheckUpRetry          int       `json:"healthCheckUpRetry"`
	HealthCheckNormalStatus     string    `json:"healthCheckNormalStatus"`
	HealthCheckUrlPath          string    `json:"healthCheckUrlPath"`
	UdpHealthCheckString        string    `json:"udpHealthCheckString"`
}

type AppServerGroup struct {
	Id          string               `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"desc"`
	Status      BLBStatus            `json:"status"`
	PortList    []AppServerGroupPort `json:"portList"`
}

type DescribeAppServerGroupResult struct {
	DescribeResultMeta
	AppServerGroupList []AppServerGroup `json:"appServerGroupList"`
}

type DeleteAppServerGroupArgs struct {
	SgId        string `json:"sgId"`
	ClientToken string `json:"-"`
}

type CreateAppServerGroupPortArgs struct {
	ClientToken                 string `json:"-"`
	SgId                        string `json:"sgId"`
	Port                        uint16 `json:"port"`
	Type                        string `json:"type"`
	EnableHealthCheck           *bool  `json:"enableHealthCheck,omitempty"`
	HealthCheck                 string `json:"healthCheck,omitempty"`
	HealthCheckHost             string `json:"healthCheckHost,omitempty"`
	HealthCheckPort             int    `json:"healthCheckPort,omitempty"`
	HealthCheckTimeoutInSecond  int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckIntervalInSecond int    `json:"healthCheckIntervalInSecond,omitempty"`
	HealthCheckDownRetry        int    `json:"healthCheckDownRetry,omitempty"`
	HealthCheckUpRetry          int    `json:"healthCheckUpRetry,omitempty"`
	HealthCheckNormalStatus     string `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckUrlPath          string `json:"healthCheckUrlPath,omitempty"`
	UdpHealthCheckString        string `json:"udpHealthCheckString,omitempty"`
}

type CreateAppServerGroupPortResult struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"desc"`
	Status      BLBStatus `json:"status"`
}

type UpdateAppServerGroupPortArgs struct {
	ClientToken                 string `json:"-"`
	SgId                        string `json:"sgId"`
	PortId                      string `json:"portId"`
	EnableHealthCheck           *bool  `json:"enableHealthCheck,omitempty"`
	HealthCheck                 string `json:"healthCheck,omitempty"`
	HealthCheckPort             int    `json:"healthCheckPort,omitempty"`
	HealthCheckHost             string `json:"healthCheckHost,omitempty"`
	HealthCheckUrlPath          string `json:"healthCheckUrlPath,omitempty"`
	HealthCheckTimeoutInSecond  int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckIntervalInSecond int    `json:"healthCheckIntervalInSecond,omitempty"`
	HealthCheckDownRetry        int    `json:"healthCheckDownRetry,omitempty"`
	HealthCheckUpRetry          int    `json:"healthCheckUpRetry,omitempty"`
	HealthCheckNormalStatus     string `json:"healthCheckNormalStatus,omitempty"`
	UdpHealthCheckString        string `json:"udpHealthCheckString,omitempty"`
}

type DeleteAppServerGroupPortArgs struct {
	SgId        string   `json:"sgId"`
	PortIdList  []string `json:"portIdList"`
	ClientToken string   `json:"-"`
}

type BlbRsWriteOpArgs struct {
	SgId              string             `json:"sgId"`
	BackendServerList []AppBackendServer `json:"backendServerList"`
	ClientToken       string             `json:"-"`
}

type CreateBlbRsArgs struct {
	BlbRsWriteOpArgs
}

type UpdateBlbRsArgs struct {
	BlbRsWriteOpArgs
}

type DescribeBlbRsArgs struct {
	Marker  string
	MaxKeys int
	SgId    string
}

type DescribeBlbRsResult struct {
	BackendServerList []AppBackendServer `json:"backendServerList"`
	DescribeResultMeta
}

type DeleteBlbRsArgs struct {
	SgId                string   `json:"sgId"`
	BackendServerIdList []string `json:"backendServerIdList"`
	ClientToken         string   `json:"-"`
}

type DescribeRsMountResult struct {
	BackendServerList []AppBackendServer `json:"backendServerList"`
}

type CreateLoadBalancerArgs struct {
	ClientToken            string           `json:"-"`
	Name                   string           `json:"name,omitempty"`
	Description            string           `json:"desc,omitempty"`
	SubnetId               string           `json:"subnetId"`
	VpcId                  string           `json:"vpcId"`
	ClusterProperty        string           `json:"clusterProperty"`
	Type                   string           `json:"type,omitempty"`
	Address                string           `json:"address,omitempty"`
	Eip                    string           `json:"eip,omitempty"`
	ResourceGroupId        string           `json:"resourceGroupId,omitempty"`
	AutoRenewLength        int              `json:"autoRenewLength,omitempty"`
	AutoRenewTimeUnit      string           `json:"autoRenewTimeUnit,omitempty"`
	PerformanceLevel       string           `json:"performanceLevel,omitempty"`
	Billing                *Billing         `json:"billing,omitempty"`
	Tags                   []model.TagModel `json:"tags,omitempty"`
	AllowDelete            *bool            `json:"allowDelete,omitempty"`
	AllocateIpv6           *bool            `json:"allocateIpv6,omitempty"`
	Layer4ClusterExclusive *bool            `json:"layer4ClusterExclusive,omitempty"`
	Layer7ClusterExclusive *bool            `json:"layer7ClusterExclusive,omitempty"`
	Layer4ClusterId        string           `json:"layer4ClusterId,omitempty"`
	Layer7ClusterId        string           `json:"layer7ClusterId,omitempty"`
}

type Billing struct {
	PaymentTiming string       `json:"paymentTiming,omitempty"`
	BillingMethod string       `json:"billingMethod,omitempty"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength,omitempty"`
	ReservationTimeUnit string `json:"reservationTimeUnit,omitempty"`
}

type CreateLoadBalanceResult struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	BlbId       string `json:"blbId"`
	Ipv6        string `json:"ipv6"`
}

type UpdateLoadBalancerArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
	AllowDelete *bool  `json:"allowDelete,omitempty"`
}

type DescribeLoadBalancersArgs struct {
	Address      string
	Name         string
	BlbId        string
	BccId        string
	Type         string
	ExactlyMatch bool
	Marker       string
	MaxKeys      int
}

type AppBLBModel struct {
	BlbId                  string           `json:"blbId"`
	Name                   string           `json:"name"`
	Description            string           `json:"desc"`
	Address                string           `json:"address"`
	Status                 BLBStatus        `json:"status"`
	VpcId                  string           `json:"vpcId"`
	SubnetId               string           `json:"subnetId"`
	PublicIp               string           `json:"publicIp"`
	Layer4ClusterId        string           `json:"layer4ClusterId"`
	Layer7ClusterId        string           `json:"layer7ClusterId"`
	Tags                   []model.TagModel `json:"tags"`
	EipRouteType           string           `json:"eipRouteType"`
	AllowDelete            bool             `json:"allowDelete"`
	Layer4ClusterExclusive bool             `json:"layer4ClusterExclusive"`
	Layer7ClusterExclusive bool             `json:"layer7ClusterExclusive"`
}

type DescribeLoadBalancersResult struct {
	BlbList []AppBLBModel `json:"blbList"`
	DescribeResultMeta
}

type ListenerModel struct {
	Port string `json:"port"`
	Type string `json:"type"`
}

type PortTypeModel struct {
	Port int    `json:"port"`
	Type string `json:"type"`
}

type DescribeLoadBalancerDetailResult struct {
	BlbId                  string           `json:"blbId"`
	Name                   string           `json:"name"`
	Status                 BLBStatus        `json:"status"`
	Description            string           `json:"desc"`
	Address                string           `json:"address"`
	PublicIp               string           `json:"publicIp"`
	Cidr                   string           `json:"cidr"`
	VpcName                string           `json:"vpcName"`
	SubnetCider            string           `json:"subnetCider"`
	SubnetName             string           `json:"subnetName"`
	CreateTime             string           `json:"createTime"`
	ReleaseTime            string           `json:"releaseTime"`
	Layer4ClusterId        string           `json:"layer4ClusterId"`
	Layer7ClusterId        string           `json:"layer7ClusterId"`
	Listener               []ListenerModel  `json:"listener"`
	Tags                   []model.TagModel `json:"tags"`
	EipRouteType           string           `json:"eipRouteType"`
	Layer4ClusterExclusive bool             `json:"layer4ClusterExclusive"`
	Layer7ClusterExclusive bool             `json:"layer7ClusterExclusive"`
	Layer4ClusterMode      string           `json:"layer4ClusterMode"`
	Layer7ClusterMode      string           `json:"layer7ClusterMode"`
	Layer4MasterAz         string           `json:"layer4MasterAz"`
	Layer7MasterAz         string           `json:"layer7MasterAz"`
	Layer4SlaveAz          string           `json:"layer4SlaveAz"`
	Layer7SlaveAz          string           `json:"layer7SlaveAz"`
	PaymentTiming          string           `json:"paymentTiming"`
	PerformanceLevel       string           `json:"performanceLevel"`
	ExpireTime             string           `json:"expireTime"`
	AllowDelete            bool             `json:"allowDelete"`
	VpcId                  string           `json:"vpcId"`
}

type CreateAppTCPListenerArgs struct {
	TcpSessionTimeout int    `json:"tcpSessionTimeout,omitempty"`
	ListenerPort      uint16 `json:"listenerPort"`
	Scheduler         string `json:"scheduler"`
	ClientToken       string `json:"-"`
}

type CreateAppUDPListenerArgs struct {
	UdpSessionTimeout int    `json:"udpSessionTimeout,omitempty"`
	ListenerPort      uint16 `json:"listenerPort"`
	Scheduler         string `json:"scheduler"`
	ClientToken       string `json:"-"`
}

type CreateAppHTTPListenerArgs struct {
	ClientToken           string `json:"-"`
	ListenerPort          uint16 `json:"listenerPort"`
	Scheduler             string `json:"scheduler"`
	KeepSession           *bool  `json:"keepSession,omitempty"`
	KeepSessionType       string `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    int    `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName string `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         *bool  `json:"xForwardedFor,omitempty"`
	XForwardedProto       *bool  `json:"xForwardedProto,omitempty"`
	ServerTimeout         int    `json:"serverTimeout,omitempty"`
	RedirectPort          uint16 `json:"redirectPort,omitempty"`
}

type CreateAppHTTPSListenerArgs struct {
	ClientToken           string                       `json:"-"`
	ListenerPort          uint16                       `json:"listenerPort"`
	Scheduler             string                       `json:"scheduler"`
	KeepSession           *bool                        `json:"keepSession,omitempty"`
	KeepSessionType       string                       `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    int                          `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName string                       `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         *bool                        `json:"xForwardedFor,omitempty"`
	XForwardedProto       *bool                        `json:"xForwardedProto,omitempty"`
	ServerTimeout         int                          `json:"serverTimeout,omitempty"`
	CertIds               []string                     `json:"certIds"`
	AdditionalCertDomains []AdditionalCertDomainsModel `json:"additionalCertDomains,omitempty"`
	EncryptionType        string                       `json:"encryptionType,omitempty"`
	EncryptionProtocols   []string                     `json:"encryptionProtocols,omitempty"`
	AppliedCiphers        string                       `json:"appliedCiphers,omitempty"`
	DualAuth              *bool                        `json:"dualAuth,omitempty"`
	ClientCertIds         []string                     `json:"clientCertIds,omitempty"`
}

type CreateAppSSLListenerArgs struct {
	ClientToken         string   `json:"-"`
	ListenerPort        uint16   `json:"listenerPort"`
	Scheduler           string   `json:"scheduler"`
	CertIds             []string `json:"certIds"`
	EncryptionType      string   `json:"encryptionType,omitempty"`
	EncryptionProtocols []string `json:"encryptionProtocols,omitempty"`
	AppliedCiphers      string   `json:"appliedCiphers,omitempty"`
	DualAuth            *bool    `json:"dualAuth,omitempty"`
	ClientCertIds       []string `json:"clientCertIds,omitempty"`
}

type UpdateAppListenerArgs struct {
	ClientToken       string `json:"-"`
	ListenerPort      uint16 `json:"-"`
	Scheduler         string `json:"scheduler,omitempty"`
	TcpSessionTimeout int    `json:"tcpSessionTimeout,omitempty"`
	UdpSessionTimeout int    `json:"udpSessionTimeout,omitempty"`
}

type UpdateAppTCPListenerArgs struct {
	UpdateAppListenerArgs
}

type UpdateAppUDPListenerArgs struct {
	UpdateAppListenerArgs
}

type UpdateAppHTTPListenerArgs struct {
	ClientToken           string `json:"-"`
	ListenerPort          uint16 `json:"-"`
	Scheduler             string `json:"scheduler"`
	KeepSession           *bool  `json:"keepSession,omitempty"`
	KeepSessionType       string `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    int    `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName string `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         *bool  `json:"xForwardedFor,omitempty"`
	XForwardedProto       *bool  `json:"xForwardedProto,omitempty"`
	ServerTimeout         int    `json:"serverTimeout,omitempty"`
	RedirectPort          uint16 `json:"redirectPort,omitempty"`
}

type UpdateAppHTTPSListenerArgs struct {
	ClientToken           string                       `json:"-"`
	ListenerPort          uint16                       `json:"listenerPort"`
	Scheduler             string                       `json:"scheduler"`
	KeepSession           *bool                        `json:"keepSession,omitempty"`
	KeepSessionType       string                       `json:"keepSessionType,omitempty"`
	KeepSessionTimeout    int                          `json:"keepSessionTimeout,omitempty"`
	KeepSessionCookieName string                       `json:"keepSessionCookieName,omitempty"`
	XForwardedFor         *bool                        `json:"xForwardedFor,omitempty"`
	XForwardedProto       *bool                        `json:"xForwardedProto,omitempty"`
	ServerTimeout         int                          `json:"serverTimeout,omitempty"`
	CertIds               []string                     `json:"certIds"`
	AdditionalCertDomains []AdditionalCertDomainsModel `json:"additionalCertDomains"`
	EncryptionType        string                       `json:"encryptionType,omitempty"`
	EncryptionProtocols   []string                     `json:"encryptionProtocols,omitempty"`
	AppliedCiphers        string                       `json:"appliedCiphers,omitempty"`
	DualAuth              *bool                        `json:"dualAuth,omitempty"`
	ClientCertIds         []string                     `json:"clientCertIds,omitempty"`
}

type UpdateAppSSLListenerArgs struct {
	ClientToken         string   `json:"-"`
	ListenerPort        uint16   `json:"-"`
	Scheduler           string   `json:"scheduler"`
	CertIds             []string `json:"certIds"`
	EncryptionType      string   `json:"encryptionType,omitempty"`
	EncryptionProtocols []string `json:"encryptionProtocols,omitempty"`
	AppliedCiphers      string   `json:"appliedCiphers,omitempty"`
	DualAuth            *bool    `json:"dualAuth,omitempty"`
	ClientCertIds       []string `json:"clientCertIds,omitempty"`
}

type AppListenerModel struct {
	Port              uint16 `json:"listenerPort"`
	Scheduler         string `json:"scheduler"`
	TcpSessionTimeout int    `json:"tcpSessionTimeout"`
	UdpSessionTimeout int    `json:"udpSessionTimeout"`
}

type AppTCPListenerModel struct {
	AppListenerModel
}

type AppUDPListenerModel struct {
	AppListenerModel
}

type AppHTTPListenerModel struct {
	ListenerPort          uint16 `json:"listenerPort"`
	Scheduler             string `json:"scheduler"`
	KeepSession           bool   `json:"keepSession"`
	KeepSessionType       string `json:"keepSessionType"`
	KeepSessionTimeout    int    `json:"keepSessionTimeout"`
	KeepSessionCookieName string `json:"keepSessionCookieName"`
	XForwardedFor         bool   `json:"xForwardedFor"`
	XForwardedProto       bool   `json:"xForwardedProto"`
	ServerTimeout         int    `json:"serverTimeout"`
	RedirectPort          int    `json:"redirectPort"`
}

type AppHTTPSListenerModel struct {
	ListenerPort          uint16                       `json:"listenerPort"`
	Scheduler             string                       `json:"scheduler"`
	KeepSession           bool                         `json:"keepSession"`
	KeepSessionType       string                       `json:"keepSessionType"`
	KeepSessionTimeout    int                          `json:"keepSessionTimeout"`
	KeepSessionCookieName string                       `json:"keepSessionCookieName"`
	XForwardedFor         bool                         `json:"xForwardedFor"`
	XForwardedProto       bool                         `json:"xForwardedProto"`
	ServerTimeout         int                          `json:"serverTimeout"`
	CertIds               []string                     `json:"certIds"`
	AdditionalCertDomains []AdditionalCertDomainsModel `json:"additionalCertDomains"`
	EncryptionType        string                       `json:"encryptionType"`
	EncryptionProtocols   []string                     `json:"encryptionProtocols"`
	AppliedCiphers        string                       `json:"appliedCiphers"`
	DualAuth              bool                         `json:"dualAuth"`
	ClientCertIds         []string                     `json:"clientCertIds"`
}

type AppSSLListenerModel struct {
	ListenerPort        uint16   `json:"listenerPort"`
	Scheduler           string   `json:"scheduler"`
	CertIds             []string `json:"certIds"`
	EncryptionType      string   `json:"encryptionType"`
	EncryptionProtocols []string `json:"encryptionProtocols"`
	AppliedCiphers      string   `json:"appliedCiphers"`
	DualAuth            bool     `json:"dualAuth"`
	ClientCertIds       []string `json:"clientCertIds"`
}

type AppAllListenerModel struct {
	ListenerPort          uint16                       `json:"listenerPort"`
	ListenerType          string                       `json:"listenerType"`
	Scheduler             string                       `json:"scheduler"`
	TcpSessionTimeout     int                          `json:"tcpSessionTimeout"`
	UdpSessionTimeout     int                          `json:"udpSessionTimeout"`
	KeepSession           bool                         `json:"keepSession"`
	KeepSessionType       string                       `json:"keepSessionType"`
	KeepSessionTimeout    int                          `json:"keepSessionTimeout"`
	KeepSessionCookieName string                       `json:"keepSessionCookieName"`
	XForwardedFor         bool                         `json:"xForwardFor"`
	XForwardedProto       bool                         `json:"xForwardedProto"`
	ServerTimeout         int                          `json:"serverTimeout"`
	RedirectPort          int                          `json:"redirectPort"`
	CertIds               []string                     `json:"certIds"`
	AdditionalCertDomains []AdditionalCertDomainsModel `json:"AdditionalCertDomains"`
	EncryptionType        string                       `json:"encryptionType"`
	EncryptionProtocols   []string                     `json:"encryptionProtocols"`
	AppliedCiphers        string                       `json:"appliedCiphers"`
	DualAuth              bool                         `json:"dualAuth"`
	ClientCertIds         []string                     `json:"clientCertIds"`
}

type DescribeAppListenerArgs struct {
	ListenerPort uint16
	Marker       string
	MaxKeys      int
}

type DescribeAppTCPListenersResult struct {
	ListenerList []AppTCPListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAppUDPListenersResult struct {
	ListenerList []AppUDPListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAppHTTPListenersResult struct {
	ListenerList []AppHTTPListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAppHTTPSListenersResult struct {
	ListenerList []AppHTTPSListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAppSSLListenersResult struct {
	ListenerList []AppSSLListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAppAllListenersResult struct {
	ListenerList []AppAllListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DeleteAppListenersArgs struct {
	ClientToken  string          `json:"-"`
	PortList     []uint16        `json:"portList"`
	PortTypeList []PortTypeModel `json:"portTypeList"`
}

type AppRule struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AppPolicy struct {
	Description      string `json:"desc"`
	AppServerGroupId string `json:"appServerGroupId,omitempty"`
	AppIpGroupId     string `json:"appIpGroupId,omitempty"`
	AppIpGroupName   string `json:"appIpGroupName,omitempty"`
	GroupType        string `json:"groupType,omitempty"`

	BackendPort uint16    `json:"backendPort,omitempty"`
	Priority    int       `json:"priority"`
	RuleList    []AppRule `json:"ruleList"`

	Id                 string `json:"id"`
	FrontendPort       int    `json:"frontendPort"`
	AppServerGroupName string `json:"appServerGroupName"`
	PortType           string `json:"portType"`
}

type CreatePolicysArgs struct {
	ClientToken  string      `json:"-"`
	ListenerPort uint16      `json:"listenerPort"`
	AppPolicyVos []AppPolicy `json:"appPolicyVos"`
	Type         string      `json:"type"`
}

type DescribePolicysArgs struct {
	Port    uint16
	Type    string
	Marker  string
	MaxKeys int
}

type DescribePolicysResult struct {
	PolicyList []AppPolicy `json:"policyList"`
	DescribeResultMeta
}

type UpdatePolicysArgs struct {
	ClientToken string            `json:"-"`
	Port        int               `json:"port"`
	Type        string            `json:"type"`
	PolicyList  []PolicyForUpdate `json:"policyList"`
}

type PolicyForUpdate struct {
	PolicyId    string `json:"policyId"`
	Priority    int    `json:"priority,omitempty"`
	Description string `json:"description,omitempty"`
}
type DeletePolicysArgs struct {
	ClientToken  string   `json:"-"`
	Port         uint16   `json:"port"`
	PolicyIdList []string `json:"policyIdList"`
	Type         string   `json:"type"`
}

type CreateAppIpGroupArgs struct {
	Name        string             `json:"name,omitempty"`
	Desc        string             `json:"desc,omitempty"`
	MemberList  []AppIpGroupMember `json:"memberList,omitempty"`
	ClientToken string             `json:"-"`
}

type AppIpGroupMember struct {
	Ip       string                      `json:"ip,omitempty"`
	Port     int                         `json:"port,omitempty"`
	Weight   int                         `json:"weight,omitempty"`
	MemberId string                      `json:"memberId,omitempty"`
	PortList []AppIpGroupMemberPortModel `json:"portList,omitempty"`
}

type AppIpGroupMemberPortModel struct {
	HealthCheckPortType string `json:"healthCheckPortType"`
	Status              string `json:"status"`
}

type CreateAppIpGroupResult struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type UpdateAppIpGroupArgs struct {
	IpGroupId   string `json:"ipGroupId"`
	Name        string `json:"name,omitempty"`
	Desc        string `json:"desc,omitempty"`
	ClientToken string `json:"-"`
}

type DescribeAppIpGroupArgs struct {
	Name         string
	ExactlyMatch bool
	Marker       string
	MaxKeys      int
}

type DescribeAppIpGroupResult struct {
	DescribeResultMeta
	AppIpGroupList []AppIpGroup `json:"appIpGroupList"`
}

type AppIpGroup struct {
	Id                string                    `json:"id"`
	Name              string                    `json:"name"`
	Desc              string                    `json:"desc"`
	BackendPolicyList []AppIpGroupBackendPolicy `json:"backendPolicyList"`
}

type AppIpGroupBackendPolicy struct {
	Id                          string `json:"id"`
	Type                        string `json:"type"`
	EnableHealthCheck           bool   `json:"enableHealthCheck"`
	HealthCheck                 string `json:"healthCheck"`
	HealthCheckPort             int    `json:"healthCheckPort"`
	HealthCheckHost             string `json:"healthCheckHost"`
	HealthCheckTimeoutInSecond  int    `json:"healthCheckTimeoutInSecond"`
	HealthCheckIntervalInSecond int    `json:"healthCheckIntervalInSecond"`
	HealthCheckDownRetry        int    `json:"healthCheckDownRetry"`
	HealthCheckUpRetry          int    `json:"healthCheckUpRetry"`
	HealthCheckNormalStatus     string `json:"healthCheckNormalStatus"`
	HealthCheckUrlPath          string `json:"healthCheckUrlPath"`
	UdpHealthCheckString        string `json:"udpHealthCheckString"`
}

type DeleteAppIpGroupArgs struct {
	IpGroupId   string `json:"ipGroupId"`
	ClientToken string `json:"-"`
}

type CreateAppIpGroupBackendPolicyArgs struct {
	ClientToken                 string `json:"-"`
	IpGroupId                   string `json:"ipGroupId"`
	Type                        string `json:"type"`
	EnableHealthCheck           *bool  `json:"enableHealthCheck,omitempty"`
	HealthCheck                 string `json:"healthCheck,omitempty"`
	HealthCheckPort             int    `json:"healthCheckPort,omitempty"`
	HealthCheckHost             string `json:"healthCheckHost,omitempty"`
	HealthCheckTimeoutInSecond  int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckIntervalInSecond int    `json:"healthCheckIntervalInSecond,omitempty"`
	HealthCheckDownRetry        int    `json:"healthCheckDownRetry,omitempty"`
	HealthCheckUpRetry          int    `json:"healthCheckUpRetry,omitempty"`
	HealthCheckNormalStatus     string `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckUrlPath          string `json:"healthCheckUrlPath,omitempty"`
	UdpHealthCheckString        string `json:"udpHealthCheckString,omitempty"`
}

type UpdateAppIpGroupBackendPolicyArgs struct {
	ClientToken                 string `json:"-"`
	IpGroupId                   string `json:"ipGroupId"`
	Id                          string `json:"id"`
	EnableHealthCheck           *bool  `json:"enableHealthCheck,omitempty"`
	HealthCheck                 string `json:"healthCheck,omitempty"`
	HealthCheckPort             int    `json:"healthCheckPort,omitempty"`
	HealthCheckHost             string `json:"healthCheckHost,omitempty"`
	HealthCheckUrlPath          string `json:"healthCheckUrlPath,omitempty"`
	HealthCheckTimeoutInSecond  int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckIntervalInSecond int    `json:"healthCheckIntervalInSecond,omitempty"`
	HealthCheckDownRetry        int    `json:"healthCheckDownRetry,omitempty"`
	HealthCheckUpRetry          int    `json:"healthCheckUpRetry,omitempty"`
	HealthCheckNormalStatus     string `json:"healthCheckNormalStatus,omitempty"`
	UdpHealthCheckString        string `json:"udpHealthCheckString,omitempty"`
}

type DeleteAppIpGroupBackendPolicyArgs struct {
	IpGroupId           string   `json:"ipGroupId"`
	BackendPolicyIdList []string `json:"backendPolicyIdList"`
	ClientToken         string   `json:"-"`
}

type AppIpGroupMemberWriteOpArgs struct {
	IpGroupId   string             `json:"ipGroupId"`
	MemberList  []AppIpGroupMember `json:"memberList"`
	ClientToken string             `json:"-"`
}

type CreateAppIpGroupMemberArgs struct {
	AppIpGroupMemberWriteOpArgs
}

type UpdateAppIpGroupMemberArgs struct {
	AppIpGroupMemberWriteOpArgs
}

type DescribeAppIpGroupMemberArgs struct {
	Marker    string
	MaxKeys   int
	IpGroupId string
}

type DescribeAppIpGroupMemberResult struct {
	MemberList []AppIpGroupMember `json:"memberList"`
	DescribeResultMeta
}

type DeleteAppIpGroupMemberArgs struct {
	IpGroupId    string   `json:"ipGroupId"`
	MemberIdList []string `json:"memberIdList"`
	ClientToken  string   `json:"-"`
}

type UpdateSecurityGroupsArgs struct {
	ClientToken      string   `json:"-"`
	SecurityGroupIds []string `json:"securityGroupIds"`
}

type UpdateEnterpriseSecurityGroupsArgs struct {
	ClientToken                string   `json:"-"`
	EnterpriseSecurityGroupIds []string `json:"enterpriseSecurityGroupIds"`
}

type DescribeSecurityGroupsResult struct {
	BlbSecurityGroups []BlbSecurityGroupModel `json:"blbSecurityGroups"`
}

type DescribeEnterpriseSecurityGroupsResult struct {
	BlbEnterpriseSecurityGroups []BlbEnterpriseSecurityGroupModel `json:"enterpriseSecurityGroups"`
}

type BlbSecurityGroupModel struct {
	SecurityGroupId    string                      `json:"securityGroupId"`
	SecurityGroupName  string                      `json:"securityGroupName"`
	SecurityGroupDesc  string                      `json:"securityGroupDesc"`
	VpcName            string                      `json:"vpcName"`
	SecurityGroupRules []BlbSecurityGroupRuleModel `json:"securityGroupRules"`
}

type BlbEnterpriseSecurityGroupModel struct {
	EnterpriseSecurityGroupId    string                                `json:"enterpriseSecurityGroupId"`
	EnterpriseSecurityGroupName  string                                `json:"enterpriseSecurityGroupName"`
	EnterpriseSecurityGroupDesc  string                                `json:"enterpriseSecurityGroupDesc"`
	EnterpriseSecurityGroupRules []BlbEnterpriseSecurityGroupRuleModel `json:"enterpriseSecurityGroupRules"`
}

type BlbSecurityGroupRuleModel struct {
	SecurityGroupRuleId string `json:"securityGroupRuleId"`
	Direction           string `json:"direction"`
	Ethertype           string `json:"ethertype,omitempty"`
	PortRange           string `json:"portRange,omitempty"`
	Protocol            string `json:"protocol,omitempty"`
	SourceGroupId       string `json:"sourceGroupId,omitempty"`
	SourceIp            string `json:"sourceIp,omitempty"`
	DestGroupId         string `json:"destGroupId,omitempty"`
	DestIp              string `json:"destIp,omitempty"`
}

type BlbEnterpriseSecurityGroupRuleModel struct {
	EnterpriseSecurityGroupRuleId string `json:"enterpriseSecurityGroupRuleId"`
	Direction                     string `json:"direction"`
	Action                        string `json:"action"`
	Priority                      int    `json:"priority"`
	Remark                        string `json:"remark"`
	Ethertype                     string `json:"ethertype,omitempty"`
	PortRange                     string `json:"portRange,omitempty"`
	Protocol                      string `json:"protocol,omitempty"`
	SourceIp                      string `json:"sourceIp,omitempty"`
	DestIp                        string `json:"destIp,omitempty"`
}

type AdditionalCertDomainsModel struct {
	CertId string `json:"certId"`
	Host   string `json:"host"`
}

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

package blb

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

type DescribeResultMeta struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
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

type CreateLoadBalancerResult struct {
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

type UpdateLoadBalancerAclArgs struct {
	ClientToken string `json:"-"`
	SupportAcl  *bool  `json:"supportAcl,omitempty"`
}

type StartLoadBalancerAutoRenewArgs struct {
	ClientToken       string `json:"-"`
	AutoRenewLength   int    `json:"autoRenewLength"`
	AutoRenewTimeUnit string `json:"autoRenewTimeUnit,omitempty"`
}

type RefundLoadBalancerArgs struct {
	ClientToken string `json:"-"`
}

type ResizeLoadBalancerArgs struct {
	ClientToken      string `json:"-"`
	PerformanceLevel string `json:"performanceLevel,omitempty"`
}

type ChangeToPostpaidArgs struct {
	ClientToken          string `json:"-"`
	BillingMethod        string `json:"billingMethod,omitempty"`
	PerformanceLevel     string `json:"performanceLevel,omitempty"`
	EffectiveImmediately *bool  `json:"effectiveImmediately,omitempty"`
}

type ChangeToPrepaidArgs struct {
	ClientToken       string `json:"-"`
	BillingMethod     string `json:"billingMethod,omitempty"`
	PerformanceLevel  string `json:"performanceLevel,omitempty"`
	ReservationLength int    `json:"reservationLength,omitempty"`
}

type OrderIdResult struct {
	OrderId string `json:"orderId"`
}

type DescribeLoadBalancersArgs struct {
	Address      string
	Name         string
	BlbId        string
	BccId        string
	ExactlyMatch bool
	Marker       string
	MaxKeys      int
	Type         string
}

type BLBModel struct {
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
	BlbList []BLBModel `json:"blbList"`
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
	Status                 BLBStatus        `json:"status"`
	Name                   string           `json:"name"`
	Description            string           `json:"desc"`
	Address                string           `json:"address"`
	PublicIp               string           `json:"publicIp"`
	Cidr                   string           `json:"cidr"`
	VpcName                string           `json:"vpcName"`
	CreateTime             string           `json:"createTime"`
	Layer4ClusterId        string           `json:"layer4ClusterId"`
	Layer7ClusterId        string           `json:"layer7ClusterId"`
	Listener               []ListenerModel  `json:"listener"`
	Tags                   []model.TagModel `json:"tags"`
	EipRouteType           string           `json:"eipRouteType"`
	Ipv6                   string           `json:"ipv6,omitempty"`
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

type CreateTCPListenerArgs struct {
	ClientToken                string `json:"-"`
	ListenerPort               uint16 `json:"listenerPort"`
	BackendPort                uint16 `json:"backendPort"`
	Scheduler                  string `json:"scheduler"`
	TcpSessionTimeout          int    `json:"tcpSessionTimeout,omitempty"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int    `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int    `json:"healthyThreshold,omitempty"`
}

type CreateUDPListenerArgs struct {
	ClientToken                string `json:"-"`
	ListenerPort               uint16 `json:"listenerPort"`
	BackendPort                uint16 `json:"backendPort"`
	Scheduler                  string `json:"scheduler"`
	HealthCheckType            string `json:"healthCheckType,omitempty"`
	HealthCheckPort            int    `json:"healthCheckPort,omitempty"`
	UdpSessionTimeout          int    `json:"udpSessionTimeout,omitempty"`
	HealthCheckString          string `json:"healthCheckString"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int    `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int    `json:"healthyThreshold,omitempty"`
}

type CreateHTTPListenerArgs struct {
	ClientToken                string `json:"-"`
	ListenerPort               uint16 `json:"listenerPort"`
	BackendPort                uint16 `json:"backendPort"`
	Scheduler                  string `json:"scheduler"`
	KeepSession                *bool  `json:"keepSession,omitempty"`
	KeepSessionType            string `json:"keepSessionType,omitempty"`
	KeepSessionDuration        int    `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      string `json:"keepSessionCookieName,omitempty"`
	XForwardedFor              *bool  `json:"xForwardFor,omitempty"`
	XForwardedProto            *bool  `json:"xForwardedProto,omitempty"`
	HealthCheckType            string `json:"healthCheckType,omitempty"`
	HealthCheckPort            uint16 `json:"healthCheckPort,omitempty"`
	HealthCheckURI             string `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int    `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int    `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    string `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckHost            string `json:"healthCheckHost,omitempty"`
	ServerTimeout              int    `json:"serverTimeout,omitempty"`
	RedirectPort               uint16 `json:"redirectPort,omitempty"`
}

type CreateHTTPSListenerArgs struct {
	ClientToken                string                       `json:"-"`
	ListenerPort               uint16                       `json:"listenerPort"`
	BackendPort                uint16                       `json:"backendPort"`
	Scheduler                  string                       `json:"scheduler"`
	CertIds                    []string                     `json:"certIds"`
	AdditionalCertDomains      []AdditionalCertDomainsModel `json:"additionalCertDomains,omitempty"`
	KeepSession                *bool                        `json:"keepSession,omitempty"`
	KeepSessionType            string                       `json:"keepSessionType,omitempty"`
	KeepSessionDuration        int                          `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      string                       `json:"keepSessionCookieName,omitempty"`
	XForwardedFor              *bool                        `json:"xForwardFor,omitempty"`
	XForwardedProto            *bool                        `json:"xForwardedProto,omitempty"`
	HealthCheckType            string                       `json:"healthCheckType,omitempty"`
	HealthCheckPort            uint16                       `json:"healthCheckPort,omitempty"`
	HealthCheckURI             string                       `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond int                          `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int                          `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int                          `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int                          `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    string                       `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckHost            string                       `json:"healthCheckHost,omitempty"`
	ServerTimeout              int                          `json:"serverTimeout,omitempty"`
	RedirectPort               uint16                       `json:"redirectPort,omitempty"`
	EncryptionType             string                       `json:"encryptionType,omitempty"`
	EncryptionProtocols        []string                     `json:"encryptionProtocols,omitempty"`
	AppliedCiphers             string                       `json:"appliedCiphers,omitempty"`
	DualAuth                   *bool                        `json:"dualAuth,omitempty"`
	ClientCertIds              []string                     `json:"clientCertIds,omitempty"`
}

type CreateSSLListenerArgs struct {
	ClientToken                string   `json:"-"`
	ListenerPort               uint16   `json:"listenerPort"`
	BackendPort                uint16   `json:"backendPort"`
	Scheduler                  string   `json:"scheduler"`
	CertIds                    []string `json:"certIds"`
	HealthCheckTimeoutInSecond int      `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int      `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int      `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int      `json:"healthyThreshold,omitempty"`
	EncryptionType             string   `json:"encryptionType,omitempty"`
	EncryptionProtocols        []string `json:"encryptionProtocols,omitempty"`
	AppliedCiphers             string   `json:"appliedCiphers,omitempty"`
	DualAuth                   *bool    `json:"dualAuth,omitempty"`
	ClientCertIds              []string `json:"clientCertIds,omitempty"`
}

type UpdateListenerArgs struct {
	ClientToken       string `json:"-"`
	ListenerPort      uint16 `json:"-"`
	Scheduler         string `json:"scheduler,omitempty"`
	TcpSessionTimeout int    `json:"tcpSessionTimeout,omitempty"`
}

type UpdateTCPListenerArgs struct {
	ClientToken                string `json:"-"`
	ListenerPort               uint16 `json:"-"`
	BackendPort                uint16 `json:"backendPort,omitempty"`
	Scheduler                  string `json:"scheduler,omitempty"`
	TcpSessionTimeout          int    `json:"tcpSessionTimeout,omitempty"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int    `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int    `json:"healthyThreshold,omitempty"`
}

type UpdateUDPListenerArgs struct {
	ClientToken                string `json:"-"`
	ListenerPort               uint16 `json:"-"`
	BackendPort                uint16 `json:"backendPort,omitempty"`
	Scheduler                  string `json:"scheduler,omitempty"`
	HealthCheckType            string `json:"healthCheckType,omitempty"`
	HealthCheckPort            int    `json:"healthCheckPort,omitempty"`
	UdpSessionTimeout          int    `json:"udpSessionTimeout,omitempty"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int    `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int    `json:"healthyThreshold,omitempty"`
	HealthCheckString          string `json:"healthCheckString,omitempty"`
}

type UpdateHTTPListenerArgs struct {
	ClientToken                string `json:"-"`
	ListenerPort               uint16 `json:"-"`
	BackendPort                uint16 `json:"backendPort,omitempty"`
	Scheduler                  string `json:"scheduler,omitempty"`
	KeepSession                *bool  `json:"keepSession,omitempty"`
	KeepSessionType            string `json:"keepSessionType,omitempty"`
	KeepSessionDuration        int    `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      string `json:"keepSessionCookieName,omitempty"`
	XForwardedFor              *bool  `json:"xForwardFor"`
	XForwardedProto            *bool  `json:"xForwardedProto"`
	HealthCheckType            string `json:"healthCheckType,omitempty"`
	HealthCheckPort            uint16 `json:"healthCheckPort,omitempty"`
	HealthCheckURI             string `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int    `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int    `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    string `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckHost            string `json:"healthCheckHost,omitempty"`
	ServerTimeout              int    `json:"serverTimeout,omitempty"`
	RedirectPort               uint16 `json:"redirectPort,omitempty"`
}

type UpdateHTTPSListenerArgs struct {
	ClientToken                string                       `json:"-"`
	ListenerPort               uint16                       `json:"listenerPort"`
	BackendPort                uint16                       `json:"backendPort,omitempty"`
	Scheduler                  string                       `json:"scheduler,omitempty"`
	KeepSession                *bool                        `json:"keepSession,omitempty"`
	KeepSessionType            string                       `json:"keepSessionType,omitempty"`
	KeepSessionDuration        int                          `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      string                       `json:"keepSessionCookieName,omitempty"`
	XForwardedFor              *bool                        `json:"xForwardFor,omitempty"`
	XForwardedProto            *bool                        `json:"xForwardedProto,omitempty"`
	HealthCheckType            string                       `json:"healthCheckType,omitempty"`
	HealthCheckPort            uint16                       `json:"healthCheckPort,omitempty"`
	HealthCheckURI             string                       `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond int                          `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int                          `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int                          `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int                          `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    string                       `json:"healthCheckNormalStatus,omitempty"`
	HealthCheckHost            string                       `json:"healthCheckHost,omitempty"`
	ServerTimeout              int                          `json:"serverTimeout,omitempty"`
	CertIds                    []string                     `json:"certIds,omitempty"`
	AdditionalCertDomains      []AdditionalCertDomainsModel `json:"additionalCertDomains"`
	EncryptionType             string                       `json:"encryptionType,omitempty"`
	EncryptionProtocols        []string                     `json:"encryptionProtocols,omitempty"`
	AppliedCiphers             string                       `json:"appliedCiphers,omitempty"`
}

type UpdateSSLListenerArgs struct {
	ClientToken                string   `json:"-"`
	ListenerPort               uint16   `json:"-"`
	BackendPort                uint16   `json:"backendPort,omitempty"`
	Scheduler                  string   `json:"scheduler,omitempty"`
	HealthCheckTimeoutInSecond int      `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int      `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int      `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int      `json:"healthyThreshold,omitempty"`
	CertIds                    []string `json:"certIds,omitempty"`
	EncryptionType             string   `json:"encryptionType,omitempty"`
	EncryptionProtocols        []string `json:"encryptionProtocols,omitempty"`
	AppliedCiphers             string   `json:"appliedCiphers,omitempty"`
	DualAuth                   *bool    `json:"dualAuth,omitempty"`
	ClientCertIds              []string `json:"clientCertIds,omitempty"`
}

type TCPListenerModel struct {
	ListenerPort               uint16 `json:"listenerPort"`
	BackendPort                uint16 `json:"backendPort"`
	Scheduler                  string `json:"scheduler"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond"`
	HealthCheckInterval        int    `json:"healthCheckInterval"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold"`
	HealthyThreshold           int    `json:"healthyThreshold"`
	GetBlbIp                   bool   `json:"getBlbIp"`
	TcpSessionTimeout          int    `json:"tcpSessionTimeout"`
}

type UDPListenerModel struct {
	ListenerPort               uint16 `json:"listenerPort"`
	BackendPort                uint16 `json:"backendPort"`
	Scheduler                  string `json:"scheduler"`
	HealthCheckType            string `json:"healthCheckType"`
	HealthCheckPort            int    `json:"healthCheckPort"`
	UdpSessionTimeout          int    `json:"udpSessionTimeout"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond"`
	HealthCheckInterval        int    `json:"healthCheckInterval"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold"`
	HealthyThreshold           int    `json:"healthyThreshold"`
	GetBlbIp                   bool   `json:"getBlbIp"`
	HealthCheckString          string `json:"healthCheckString"`
}

type HTTPListenerModel struct {
	ListenerPort               uint16 `json:"listenerPort"`
	BackendPort                uint16 `json:"backendPort"`
	Scheduler                  string `json:"scheduler"`
	KeepSession                bool   `json:"keepSession"`
	KeepSessionType            string `json:"keepSessionType"`
	KeepSessionDuration        int    `json:"keepSessionDuration"`
	KeepSessionCookieName      string `json:"keepSessionCookieName"`
	XForwardedFor              bool   `json:"xForwardFor"`
	XForwardedProto            bool   `json:"xForwardedProto"`
	HealthCheckType            string `json:"healthCheckType"`
	HealthCheckPort            uint16 `json:"healthCheckPort"`
	HealthCheckURI             string `json:"healthCheckURI"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond"`
	HealthCheckInterval        int    `json:"healthCheckInterval"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold"`
	HealthyThreshold           int    `json:"healthyThreshold"`
	GetBlbIp                   bool   `json:"getBlbIp"`
	HealthCheckNormalStatus    string `json:"healthCheckNormalStatus"`
	HealthCheckHost            string `json:"healthCheckHost"`
	ServerTimeout              int    `json:"serverTimeout"`
	RedirectPort               int    `json:"redirectPort"`
}

type HTTPSListenerModel struct {
	ListenerPort               uint16                       `json:"listenerPort"`
	BackendPort                uint16                       `json:"backendPort"`
	Scheduler                  string                       `json:"scheduler"`
	KeepSession                bool                         `json:"keepSession"`
	KeepSessionType            string                       `json:"keepSessionType"`
	KeepSessionDuration        int                          `json:"keepSessionDuration"`
	KeepSessionCookieName      string                       `json:"keepSessionCookieName"`
	XForwardedFor              bool                         `json:"xForwardFor"`
	XForwardedProto            bool                         `json:"xForwardedProto"`
	HealthCheckType            string                       `json:"healthCheckType"`
	HealthCheckPort            uint16                       `json:"healthCheckPort"`
	HealthCheckURI             string                       `json:"healthCheckURI"`
	HealthCheckTimeoutInSecond int                          `json:"healthCheckTimeoutInSecond"`
	HealthCheckInterval        int                          `json:"healthCheckInterval"`
	UnhealthyThreshold         int                          `json:"unhealthyThreshold"`
	HealthyThreshold           int                          `json:"healthyThreshold"`
	GetBlbIp                   bool                         `json:"getBlbIp"`
	HealthCheckNormalStatus    string                       `json:"healthCheckNormalStatus"`
	HealthCheckHost            string                       `json:"healthCheckHost"`
	ServerTimeout              int                          `json:"serverTimeout"`
	CertIds                    []string                     `json:"certIds"`
	AdditionalCertDomains      []AdditionalCertDomainsModel `json:"additionalCertDomains"`
	DualAuth                   bool                         `json:"dualAuth"`
	ClientCertIds              []string                     `json:"clientCertIds"`
	EncryptionType             string                       `json:"encryptionType"`
	EncryptionProtocols        []string                     `json:"encryptionProtocols"`
	AppliedCiphers             string                       `json:"appliedCiphers"`
}

type SSLListenerModel struct {
	ListenerPort               uint16   `json:"listenerPort"`
	BackendPort                uint16   `json:"backendPort"`
	Scheduler                  string   `json:"scheduler"`
	HealthCheckTimeoutInSecond int      `json:"healthCheckTimeoutInSecond"`
	HealthCheckInterval        int      `json:"healthCheckInterval"`
	UnhealthyThreshold         int      `json:"unhealthyThreshold"`
	HealthyThreshold           int      `json:"healthyThreshold"`
	GetBlbIp                   bool     `json:"getBlbIp"`
	CertIds                    []string `json:"certIds"`
	EncryptionType             string   `json:"encryptionType"`
	EncryptionProtocols        []string `json:"encryptionProtocols"`
	AppliedCiphers             string   `json:"appliedCiphers"`
	DualAuth                   bool     `json:"dualAuth"`
	ClientCertIds              []string `json:"clientCertIds"`
	ServerTimeout              int      `json:"serverTimeout"`
}

type AllListenerModel struct {
	ListenerPort               uint16                       `json:"listenerPort"`
	ListenerType               string                       `json:"listenerType"`
	BackendPort                uint16                       `json:"backendPort"`
	Scheduler                  string                       `json:"scheduler"`
	GetBlbIp                   bool                         `json:"getBlbIp"`
	TcpSessionTimeout          int                          `json:"tcpSessionTimeout"`
	UdpSessionTimeout          int                          `json:"udpSessionTimeout"`
	HealthCheckString          string                       `json:"healthCheckString"`
	KeepSession                bool                         `json:"keepSession"`
	KeepSessionType            string                       `json:"keepSessionType"`
	KeepSessionDuration        int                          `json:"keepSessionDuration"`
	KeepSessionCookieName      string                       `json:"keepSessionCookieName"`
	XForwardedFor              bool                         `json:"xForwardFor"`
	XForwardedProto            bool                         `json:"xForwardedProto"`
	HealthCheckType            string                       `json:"healthCheckType"`
	HealthCheckPort            uint16                       `json:"healthCheckPort"`
	HealthCheckURI             string                       `json:"healthCheckURI"`
	HealthCheckTimeoutInSecond int                          `json:"healthCheckTimeoutInSecond"`
	HealthCheckInterval        int                          `json:"healthCheckInterval"`
	UnhealthyThreshold         int                          `json:"unhealthyThreshold"`
	HealthyThreshold           int                          `json:"healthyThreshold"`
	HealthCheckNormalStatus    string                       `json:"healthCheckNormalStatus"`
	HealthCheckHost            string                       `json:"healthCheckHost"`
	ServerTimeout              int                          `json:"serverTimeout"`
	RedirectPort               int                          `json:"redirectPort"`
	CertIds                    []string                     `json:"certIds"`
	AdditionalCertDomains      []AdditionalCertDomainsModel `json:"additionalCertDomains"`
	DualAuth                   bool                         `json:"dualAuth"`
	ClientCertIds              []string                     `json:"clientCertIds"`
	EncryptionType             string                       `json:"encryptionType"`
	EncryptionProtocols        []string                     `json:"encryptionProtocols"`
	AppliedCiphers             string                       `json:"appliedCiphers"`
}

type DescribeListenerArgs struct {
	ListenerPort uint16
	Marker       string
	MaxKeys      int
}

type DescribeTCPListenersResult struct {
	ListenerList []TCPListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeUDPListenersResult struct {
	ListenerList []UDPListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeHTTPListenersResult struct {
	ListenerList []HTTPListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeHTTPSListenersResult struct {
	ListenerList []HTTPSListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeSSLListenersResult struct {
	ListenerList []SSLListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DescribeAllListenersResult struct {
	AllListenerList []AllListenerModel `json:"listenerList"`
	DescribeResultMeta
}

type DeleteListenersArgs struct {
	ClientToken  string          `json:"-"`
	PortList     []uint16        `json:"portList"`
	PortTypeList []PortTypeModel `json:"portTypeList"`
}

type AddBackendServersArgs struct {
	ClientToken       string               `json:"-"`
	BackendServerList []BackendServerModel `json:"backendServerList"`
}

type BackendServerModel struct {
	InstanceId string `json:"instanceId"`
	Weight     int    `json:"weight"`
	PrivateIp  string `json:"privateIp,omitempty"`
}

type BackendServerStatus struct {
	InstanceId string `json:"instanceId"`
	Weight     int    `json:"weight"`
	Status     string `json:"status"`
	PrivateIp  string `json:"privateIp"`
}

type UpdateBackendServersArgs struct {
	ClientToken       string               `json:"-"`
	BackendServerList []BackendServerModel `json:"backendServerList"`
}

type DescribeBackendServersArgs struct {
	Marker  string
	MaxKeys int
}

type DescribeBackendServersResult struct {
	BackendServerList []BackendServerModel `json:"backendServerList"`
	DescribeResultMeta
}

type DescribeHealthStatusArgs struct {
	ListenerPort uint16
	Marker       string
	MaxKeys      int
}

type DescribeHealthStatusResult struct {
	BackendServerList []BackendServerStatus `json:"backendServerList"`
	Type              string                `json:"type"`
	ListenerPort      uint16                `json:"listenerPort"`
	BackendPort       uint16                `json:"backendPort"`
	DescribeResultMeta
}

type RemoveBackendServersArgs struct {
	ClientToken       string   `json:"-"`
	BackendServerList []string `json:"backendServerList"`
}

type DescribeLbClusterDetailResult struct {
	ClusterId          string `json:"clusterId"`
	ClusterName        string `json:"clusterName"`
	ClusterType        string `json:"clusterType"`
	ClusterRegion      string `json:"clusterRegion"`
	ClusterAz          string `json:"clusterAz"`
	TotalConnectCount  uint64 `json:"totalConnectCount"`
	NewConnectCps      uint64 `json:"newConnectCps"`
	NetworkInBps       uint64 `json:"networkInBps"`
	NetworkOutBps      uint64 `json:"networkOutBps"`
	NetworkInPps       uint64 `json:"networkInPps"`
	NetworkOutPps      uint64 `json:"networkOutPps"`
	HttpsQps           uint64 `json:"httpsQps"`
	HttpQps            uint64 `json:"httpQps"`
	HttpNewConnectCps  uint64 `json:"httpNewConnectCps"`
	HttpsNewConnectCps uint64 `json:"httpsNewConnectCps"`
}

type DescribeLbClustersArgs struct {
	ClusterName  string
	ClusterId    string
	ExactlyMatch bool
	Marker       string
	MaxKeys      int
}

type DescribeLbClustersResult struct {
	ClusterList []ClusterModel `json:"clusterList"`
	DescribeResultMeta
}

type ClusterModel struct {
	ClusterId     string `json:"clusterId"`
	ClusterName   string `json:"clusterName"`
	ClusterType   string `json:"clusterType"`
	ClusterRegion string `json:"clusterRegion"`
	ClusterAz     string `json:"clusterAz"`
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

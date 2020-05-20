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
	ClientToken string           `json:"-"`
	Name        string           `json:"name,omitempty"`
	Description string           `json:"desc,omitempty"`
	SubnetId    string           `json:"subnetId"`
	VpcId       string           `json:"vpcId"`
	Tags        []model.TagModel `json:"tags,omitempty"`
}

type CreateLoadBalancerResult struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	BlbId       string `json:"blbId"`
}

type UpdateLoadBalancerArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name,omitempty"`
	Description string `json:"desc,omitempty"`
}

type DescribeLoadBalancersArgs struct {
	Address      string
	Name         string
	BlbId        string
	BccId        string
	ExactlyMatch bool
	Marker       string
	MaxKeys      int
}

type BLBModel struct {
	BlbId       string           `json:"blbId"`
	Name        string           `json:"name"`
	Description string           `json:"desc"`
	Address     string           `json:"address"`
	Status      BLBStatus        `json:"status"`
	VpcId       string           `json:"vpcId"`
	SubnetId    string           `json:"subnetId"`
	PublicIp    string           `json:"publicIp"`
	Tags        []model.TagModel `json:"tags"`
}

type DescribeLoadBalancersResult struct {
	BlbList []BLBModel `json:"blbList"`
	DescribeResultMeta
}

type ListenerModel struct {
	Port string `json:"port"`
	Type string `json:"type"`
}

type DescribeLoadBalancerDetailResult struct {
	BlbId       string           `json:"blbId"`
	Status      BLBStatus        `json:"status"`
	Name        string           `json:"name"`
	Description string           `json:"desc"`
	Address     string           `json:"address"`
	PublicIp    string           `json:"publicIp"`
	Cidr        string           `json:"cidr"`
	VpcName     string           `json:"vpcName"`
	CreateTime  string           `json:"createTime"`
	Listener    []ListenerModel  `json:"listener"`
	Tags        []model.TagModel `json:"tags"`
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
	KeepSession                bool   `json:"keepSession,omitempty"`
	KeepSessionType            string `json:"keepSessionType,omitempty"`
	KeepSessionDuration        int    `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      string `json:"keepSessionCookieName,omitempty"`
	XForwardedFor              bool   `json:"xForwardedFor,omitempty"`
	HealthCheckType            string `json:"healthCheckType,omitempty"`
	HealthCheckPort            uint16 `json:"healthCheckPort,omitempty"`
	HealthCheckURI             string `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int    `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int    `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    string `json:"healthCheckNormalStatus,omitempty"`
	ServerTimeout              int    `json:"serverTimeout,omitempty"`
	RedirectPort               uint16 `json:"redirectPort,omitempty"`
}

type CreateHTTPSListenerArgs struct {
	ClientToken                string   `json:"-"`
	ListenerPort               uint16   `json:"listenerPort"`
	BackendPort                uint16   `json:"backendPort"`
	Scheduler                  string   `json:"scheduler"`
	CertIds                    []string `json:"certIds"`
	KeepSession                bool     `json:"keepSession,omitempty"`
	KeepSessionType            string   `json:"keepSessionType,omitempty"`
	KeepSessionDuration        int      `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      string   `json:"keepSessionCookieName,omitempty"`
	XForwardedFor              bool     `json:"xForwardedFor,omitempty"`
	HealthCheckType            string   `json:"healthCheckType,omitempty"`
	HealthCheckPort            uint16   `json:"healthCheckPort,omitempty"`
	HealthCheckURI             string   `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond int      `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int      `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int      `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int      `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    string   `json:"healthCheckNormalStatus,omitempty"`
	ServerTimeout              int      `json:"serverTimeout,omitempty"`
	RedirectPort               uint16   `json:"redirectPort,omitempty"`
	Ie6Compatible              bool     `json:"ie6Compatible,omitempty"`
	EncryptionType             string   `json:"encryptionType,omitempty"`
	EncryptionProtocols        []string `json:"encryptionProtocols,omitempty"`
	DualAuth                   bool     `json:"dualAuth,omitempty"`
	ClientCertIds              []string `json:"clientCertIds,omitempty"`
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
	Ie6Compatible              bool     `json:"ie6Compatible,omitempty"`
	EncryptionType             string   `json:"encryptionType,omitempty"`
	EncryptionProtocols        []string `json:"encryptionProtocols,omitempty"`
	DualAuth                   bool     `json:"dualAuth,omitempty"`
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
	KeepSession                bool   `json:"keepSession,omitempty"`
	KeepSessionType            string `json:"keepSessionType,omitempty"`
	KeepSessionDuration        int    `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      string `json:"keepSessionCookieName,omitempty"`
	XForwardedFor              bool   `json:"xForwardedFor,omitempty"`
	HealthCheckType            string `json:"healthCheckType,omitempty"`
	HealthCheckPort            uint16 `json:"healthCheckPort,omitempty"`
	HealthCheckURI             string `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int    `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int    `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    string `json:"healthCheckNormalStatus,omitempty"`
	ServerTimeout              int    `json:"serverTimeout,omitempty"`
	RedirectPort               uint16 `json:"redirectPort,omitempty"`
}

type UpdateHTTPSListenerArgs struct {
	ClientToken                string   `json:"-"`
	ListenerPort               uint16   `json:"listenerPort"`
	BackendPort                uint16   `json:"backendPort,omitempty"`
	Scheduler                  string   `json:"scheduler,omitempty"`
	KeepSession                bool     `json:"keepSession,omitempty"`
	KeepSessionType            string   `json:"keepSessionType,omitempty"`
	KeepSessionDuration        int      `json:"keepSessionDuration,omitempty"`
	KeepSessionCookieName      string   `json:"keepSessionCookieName,omitempty"`
	XForwardedFor              bool     `json:"xForwardedFor,omitempty"`
	HealthCheckType            string   `json:"healthCheckType,omitempty"`
	HealthCheckPort            uint16   `json:"healthCheckPort,omitempty"`
	HealthCheckURI             string   `json:"healthCheckURI,omitempty"`
	HealthCheckTimeoutInSecond int      `json:"healthCheckTimeoutInSecond,omitempty"`
	HealthCheckInterval        int      `json:"healthCheckInterval,omitempty"`
	UnhealthyThreshold         int      `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold           int      `json:"healthyThreshold,omitempty"`
	HealthCheckNormalStatus    string   `json:"healthCheckNormalStatus,omitempty"`
	ServerTimeout              int      `json:"serverTimeout,omitempty"`
	CertIds                    []string `json:"certIds,omitempty"`
	Ie6Compatible              bool     `json:"ie6Compatible,omitempty"`
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
	Ie6Compatible              bool     `json:"ie6Compatible,omitempty"`
	EncryptionType             string   `json:"encryptionType,omitempty"`
	EncryptionProtocols        []string `json:"encryptionProtocols,omitempty"`
	DualAuth                   bool     `json:"dualAuth,omitempty"`
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
	XForwardedFor              bool   `json:"xForwardedFor"`
	HealthCheckType            string `json:"healthCheckType"`
	HealthCheckPort            uint16 `json:"healthCheckPort"`
	HealthCheckURI             string `json:"healthCheckURI"`
	HealthCheckTimeoutInSecond int    `json:"healthCheckTimeoutInSecond"`
	HealthCheckInterval        int    `json:"healthCheckInterval"`
	UnhealthyThreshold         int    `json:"unhealthyThreshold"`
	HealthyThreshold           int    `json:"healthyThreshold"`
	GetBlbIp                   bool   `json:"getBlbIp"`
	HealthCheckNormalStatus    string `json:"healthCheckNormalStatus"`
	ServerTimeout              int    `json:"serverTimeout"`
	RedirectPort               int    `json:"redirectPort"`
}

type HTTPSListenerModel struct {
	ListenerPort               uint16   `json:"listenerPort"`
	BackendPort                uint16   `json:"backendPort"`
	Scheduler                  string   `json:"scheduler"`
	KeepSession                bool     `json:"keepSession"`
	KeepSessionType            string   `json:"keepSessionType"`
	KeepSessionDuration        int      `json:"keepSessionDuration"`
	KeepSessionCookieName      string   `json:"keepSessionCookieName"`
	XForwardedFor              bool     `json:"xForwardedFor"`
	HealthCheckType            string   `json:"healthCheckType"`
	HealthCheckPort            uint16   `json:"healthCheckPort"`
	HealthCheckURI             string   `json:"healthCheckURI"`
	HealthCheckTimeoutInSecond int      `json:"healthCheckTimeoutInSecond"`
	HealthCheckInterval        int      `json:"healthCheckInterval"`
	UnhealthyThreshold         int      `json:"unhealthyThreshold"`
	HealthyThreshold           int      `json:"healthyThreshold"`
	GetBlbIp                   bool     `json:"getBlbIp"`
	HealthCheckNormalStatus    string   `json:"healthCheckNormalStatus"`
	ServerTimeout              int      `json:"serverTimeout"`
	CertIds                    []string `json:"certIds"`
	Ie6Compatible              bool     `json:"ie6Compatible"`
	DualAuth                   bool     `json:"dualAuth"`
	ClientCertIds              []string `json:"clientCertIds"`
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
	Ie6Compatible              bool     `json:"ie6Compatible"`
	EncryptionType             string   `json:"encryptionType"`
	EncryptionProtocols        []string `json:"encryptionProtocols"`
	DualAuth                   bool     `json:"dualAuth"`
	ClientCertIds              []string `json:"clientCertIds"`
	ServerTimeout              int      `json:"serverTimeout"`
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

type DeleteListenersArgs struct {
	ClientToken string   `json:"-"`
	PortList    []uint16 `json:"portList"`
}

type AddBackendServersArgs struct {
	ClientToken       string               `json:"-"`
	BackendServerList []BackendServerModel `json:"backendServerList"`
}

type BackendServerModel struct {
	InstanceId string `json:"instanceId"`
	Weight     int    `json:"weight"`
}

type BackendServerStatus struct {
	InstanceId string `json:"instanceId"`
	Weight     int    `json:"weight"`
	Status     string `json:"status"`
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
	ListenerPort      string                `json:"listenerPort"`
	BackendPort       string                `json:"backendPort"`
	DescribeResultMeta
}

type RemoveBackendServersArgs struct {
	ClientToken       string   `json:"-"`
	BackendServerList []string `json:"backendServerList"`
}

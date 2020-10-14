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

package scs

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type Billing struct {
	PaymentTiming string       `json:"paymentTiming"`
	Reservation   *Reservation `json:"reservation,omitempty"`
}

type Subnet struct {
	ZoneName string `json:"zoneName"`
	SubnetID string `json:"subnetId"`
}



type CreateInstanceArgs struct {
	Billing           Billing  `json:"billing"`
	PurchaseCount     int      `json:"purchaseCount"`
	InstanceName      string   `json:"instanceName"`
	NodeType          string   `json:"nodeType"`
	ShardNum          int      `json:"shardNum"`
	ProxyNum          int      `json:"proxyNum"`
	ClusterType       string   `json:"clusterType"`
	ReplicationNum    int      `json:"replicationNum"`
	Port              int      `json:"port"`
	EngineVersion     string   `json:"engineVersion"`
	VpcID             string   `json:"vpcId"`
	Subnets           []Subnet `json:"subnets,omitempty"`
	AutoRenewTimeUnit string   `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int      `json:"autoRenewTime,omitempty"`
	ClientToken       string   `json:"-"`
}

type CreateInstanceResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type InstanceModel struct {
	InstanceID         string           `json:"instanceId"`
	InstanceName       string           `json:"instanceName"`
	InstanceStatus     string           `json:"instanceStatus"`
	ClusterType        string           `json:"clusterType"`
	Engine             string           `json:"engine"`
	EngineVersion      string           `json:"engineVersion"`
	VnetIP             string           `json:"vnetIp"`
	Domain             string           `json:"domain"`
	Port               int              `json:"port"`
	InstanceCreateTime string           `json:"instanceCreateTime"`
	Capacity           int              `json:"capacity"`
	UsedCapacity       float64          `json:"usedCapacity"`
	PaymentTiming      string           `json:"paymentTiming"`
	ZoneNames          []string         `json:"zoneNames"`
	Tags               []model.TagModel `json:"tags"`
}

type ListInstancesArgs struct {
	Marker  string
	MaxKeys int
}

type ListInstancesResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Instances   []InstanceModel `json:"instances"`
}

type ResizeInstanceArgs struct {
	NodeType    string `json:"nodeType"`
	ShardNum    int    `json:"shardNum"`
	ClientToken string `json:"-"`
}

type GetInstanceDetailResult struct {
	InstanceID         string           `json:"instanceId"`
	InstanceName       string           `json:"instanceName"`
	InstanceStatus     string           `json:"instanceStatus"`
	ClusterType        string           `json:"clusterType"`
	Engine             string           `json:"engine"`
	EngineVersion      string           `json:"engineVersion"`
	VnetIP             string           `json:"vnetIp"`
	Domain             string           `json:"domain"`
	Port               int              `json:"port"`
	InstanceCreateTime string           `json:"instanceCreateTime"`
	InstanceExpireTime string           `json:"instanceExpireTime"`
	Capacity           int              `json:"capacity"`
	UsedCapacity       float64          `json:"usedCapacity"`
	PaymentTiming      string           `json:"paymentTiming"`
	VpcID              string           `json:"vpcId"`
	ZoneNames          []string         `json:"zoneNames"`
	Subnets            []Subnet         `json:"subnets"`
	AutoRenew          int              `json:"autoRenew"`
	Tags               []model.TagModel `json:"tags"`
}

type UpdateInstanceNameArgs struct {
	InstanceName string `json:"instanceName"`
	ClientToken  string `json:"-"`
}

type NodeType struct {
	InstanceFlavor          int     `json:"instanceFlavor"`
	NodeType                string  `json:"nodeType"`
	CPUNum                  int     `json:"cpuNum"`
	NetworkThroughputInGbps float64 `json:"networkThroughputInGbps"`
	PeakQPS                 int     `json:"peakQps"`
	MaxConnections          int     `json:"maxConnections"`
	AllowedNodeNumList      []int   `json:"allowedNodeNumList"`
}

type GetNodeTypeListResult struct {
	ClusterNodeTypeList []NodeType `json:"clusterNodeTypeList"`
	DefaultNodeTypeList []NodeType `json:"defaultNodeTypeList"`
	HsdbNodeTypeList    []NodeType `json:"hsdbNodeTypeList"`
}

type ListSubnetsArgs struct {
	VpcID               string         `json:"vpcId"`
	ZoneName            string         `json:"zoneName"`
}

type ListSubnetsResult struct {
	SubnetOriginals            []SubnetOriginal         `json:"subnets"`
}

type SubnetOriginal struct {
	Name     string `json:"name"`
	SubnetID string `json:"subnetId"`
	ZoneName string `json:"zoneName"`
	Cidr     string `json:"cidr"`
	VpcID    string `json:"vpcId"`
}

type UpdateInstanceDomainNameArgs struct {
	Domain              string `json:"domain"`
	ClientToken         string `json:"-"`
}

type GetZoneListResult struct {
	ZoneNames          []string         `json:"zoneNames"`
}

type FlushInstanceArgs struct {
	Password           string `json:"password"`
	ClientToken        string `json:"-"`
}

type BindingTagArgs struct {
	ChangeTags           []model.TagModel         `json:"changeTags"`
}

type GetSecurityIpResult struct {
	SecurityIps          []string         `json:"securityIps"`
}

type SecurityIpArgs struct {
	SecurityIps          []string         `json:"securityIps"`
	ClientToken          string           `json:"-"`
}

type ModifyPasswordArgs struct {
	Password             string `json:"password"`
	ClientToken          string `json:"-"`
}

type GetParametersResult struct {
	Parameters           []Parameter         `json:"parameters"`
}

type Parameter struct {
	Default              string           `json:"default"`
	ForceRestart         int              `json:"forceRestart"`
	Name                 string           `json:"name"`
	Value                string           `json:"value"`
}

type ModifyParametersArgs struct {
	Parameter            InstanceParam	  `json:"parameter"`
	ClientToken          string           `json:"-"`
}

type InstanceParam struct {
	Name                 string           `json:"name"`
	Value                string           `json:"value"`
}

type GetBackupListResult struct {
	TotalCount           string           `json:"totalCount"`
	Backups              []Backup         `json:"backups"`
}

type Backup struct {
	BackupType           string           `json:"backupType"`
	Comment              string           `json:"comment"`
	StartTime            string           `json:"startTime"`
	Records              []Record         `json:"records"`
}

type Record struct {
	BackupRecordId       string           `json:"backupRecordId"`
	BackupStatus         string           `json:"backupStatus"`
	Duration             string           `json:"duration"`
	ObjectSize           string           `json:"objectSize"`
	ShardName            string           `json:"shardName"`
	StartTime            string           `json:"startTime"`
}


type ModifyBackupPolicyArgs struct {
	BackupDays           string              `json:"backupDays"`
	BackupTime           string              `json:"backupTime"`
	ClientToken          string              `json:"clientToken"`
	ExpireDay            int                 `json:"expireDay"`
}
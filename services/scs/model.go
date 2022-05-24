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
	Billing           Billing       `json:"billing"`
	PurchaseCount     int           `json:"purchaseCount"`
	InstanceName      string        `json:"instanceName"`
	NodeType          string        `json:"nodeType"`
	ShardNum          int           `json:"shardNum"`
	ProxyNum          int           `json:"proxyNum"`
	ClusterType       string        `json:"clusterType"`
	ReplicationNum    int           `json:"replicationNum"`
	ReplicationInfo   []Replication `json:"replicationInfo"`
	Port              int           `json:"port"`
	Engine            int           `json:"engine,omitempty"`
	EngineVersion     string        `json:"engineVersion"`
	DiskFlavor        int           `json:"diskFlavor,omitempty"`
	DiskType          string        `json:"diskType,omitempty"`
	VpcID             string        `json:"vpcId"`
	Subnets           []Subnet      `json:"subnets,omitempty"`
	AutoRenewTimeUnit string        `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int           `json:"autoRenewTime,omitempty"`
	BgwGroupId        string        `json:"bgwGroupId,omitempty"`
	ClientToken       string        `json:"-"`
	ClientAuth        string        `json:"clientAuth"`
	StoreType         int           `json:"storeType"`
	EnableReadOnly    int           `json:"enableReadOnly,omitempty"`
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
	IsDefer     bool   `json:"isDefer"`
	ClientToken string `json:"-"`
	DiskFlavor  int    `json:"diskFlavor"`
	DiskType    string `json:"diskType"`
}

type ReplicationArgs struct {
	ResizeType      string        `json:"resizeType"`
	ReplicationInfo []Replication `json:"replicationInfo"`
	ClientToken     string        `json:"-"`
}

type Replication struct {
	AvailabilityZone string `json:"availabilityZone"`
	SubnetId         string `json:"subnetId"`
	IsMaster         int    `json:"isMaster"`
}
type RestartInstanceArgs struct {
	IsDefer bool `json:"isDefer"`
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
	AutoRenew          string           `json:"autoRenew"`
	Tags               []model.TagModel `json:"tags"`
	ShardNum           int              `json:"shardNum"`
	ReplicationNum     int              `json:"replicationNum"`
	NodeType           string           `json:"nodeType"`
	DiskFlavor         int              `json:"diskFlavor"`
	DiskType           string           `json:"diskType"`
	StoreType          int              `json:"storeType"`
	Eip                string           `json:"eip"`
	PublicDomain       string           `json:"publicDomain"`
	EnableReadOnly     int              `json:"enableReadOnly"`
	ReplicationInfo    []Replication    `json:"replicationInfo"`
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
	VpcID    string `json:"vpcId"`
	ZoneName string `json:"zoneName"`
}

type ListSubnetsResult struct {
	SubnetOriginals []SubnetOriginal `json:"subnets"`
}

type SubnetOriginal struct {
	Name     string `json:"name"`
	SubnetID string `json:"subnetId"`
	ZoneName string `json:"zoneName"`
	Cidr     string `json:"cidr"`
	VpcID    string `json:"vpcId"`
}

type UpdateInstanceDomainNameArgs struct {
	Domain      string `json:"domain"`
	ClientToken string `json:"-"`
}

type GetZoneListResult struct {
	Zones []ZoneNames `json:"zones"`
}

type ZoneNames struct {
	ZoneNames []string `json:"zoneNames"`
}

type FlushInstanceArgs struct {
	Password    string `json:"password"`
	ClientToken string `json:"-"`
}

type BindingTagArgs struct {
	ChangeTags []model.TagModel `json:"changeTags"`
}

type GetSecurityIpResult struct {
	SecurityIps []string `json:"securityIps"`
}

type SecurityIpArgs struct {
	SecurityIps []string `json:"securityIps"`
	ClientToken string   `json:"-"`
}

type ModifyPasswordArgs struct {
	Password    string `json:"password"`
	ClientToken string `json:"-"`
}

type GetParametersResult struct {
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Default      string `json:"default"`
	ForceRestart string `json:"forceRestart"`
	Name         string `json:"name"`
	Value        string `json:"value"`
}

type ModifyParametersArgs struct {
	Parameter   InstanceParam `json:"parameter"`
	ClientToken string        `json:"-"`
}

type InstanceParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type GetBackupListResult struct {
	TotalCount string       `json:"totalCount"`
	Backups    []BackupInfo `json:"backups"`
}

type BackupInfo struct {
	BackupType string         `json:"backupType"`
	Comment    string         `json:"comment"`
	StartTime  string         `json:"startTime"`
	Records    []BackupRecord `json:"records"`
}

type BackupRecord struct {
	BackupRecordId string `json:"backupRecordId"`
	BackupStatus   string `json:"backupStatus"`
	Duration       string `json:"duration"`
	ObjectSize     string `json:"objectSize"`
	ShardName      string `json:"shardName"`
	StartTime      string `json:"startTime"`
}

type ModifyBackupPolicyArgs struct {
	BackupDays  string `json:"backupDays"`
	BackupTime  string `json:"backupTime"`
	ClientToken string `json:"clientToken"`
	ExpireDay   int    `json:"expireDay"`
}

type ListVpcSecurityGroupsResult struct {
	Groups []SecurityGroup `json:"groups"`
}

type SecurityGroup struct {
	Name                 string `json:"name"`
	SecurityGroupID      string `json:"securityGroupId"`
	Description          string `json:"description"`
	TenantID             string `json:"tenantId"`
	AssociateNum         int    `json:"associateNum"`
	VpcID                string `json:"vpcId"`
	VpcShortID           string `json:"vpcShortId"`
	VpcName              string `json:"vpcName"`
	CreatedTime          string `json:"createdTime"`
	Version              int    `json:"version"`
	DefaultSecurityGroup int    `json:"defaultSecurityGroup"`
}

type SecurityGroupArgs struct {
	InstanceIds      []string `json:"instanceIds"`
	SecurityGroupIds []string `json:"securityGroupIds"`
}

type UnbindSecurityGroupArgs struct {
	InstanceId       string   `json:"instanceId"`
	SecurityGroupIds []string `json:"securityGroupIds"`
}

type ListSecurityGroupResult struct {
	Groups []SecurityGroupDetail `json:"groups"`
}

type SecurityGroupRule struct {
	PortRange           string `json:"portRange"`
	Protocol            string `json:"protocol"`
	RemoteGroupID       string `json:"remoteGroupId"`
	RemoteIP            string `json:"remoteIP"`
	Ethertype           string `json:"ethertype"`
	TenantID            string `json:"tenantId"`
	Name                string `json:"name"`
	ID                  string `json:"id"`
	SecurityGroupRuleID string `json:"securityGroupRuleId"`
	Direction           string `json:"direction"`
}

type SecurityGroupDetail struct {
	SecurityGroupName   string              `json:"securityGroupName"`
	SecurityGroupID     string              `json:"securityGroupId"`
	SecurityGroupRemark string              `json:"securityGroupRemark"`
	Inbound             []SecurityGroupRule `json:"inbound"`
	Outbound            []SecurityGroupRule `json:"outbound"`
	VpcName             string              `json:"vpcName"`
	VpcID               string              `json:"vpcId"`
	ProjectID           string              `json:"projectId"`
}

type RequestBuilder struct {
}

type GetPriceRequest struct {
	Engine         string `json:"engine,omitempty"`
	ShardNum       int    `json:"shardNum,omitempty"`
	Period         int    `json:"period,omitempty"`
	ChargeType     string `json:"chargeType,omitempty"`
	NodeType       string `json:"nodeType,omitempty"`
	ReplicationNum int    `json:"replicationNum,omitempty"`
	ClusterType    string `json:"clusterType,omitempty"`
}

type GetPriceResult struct {
	Price float64 `json:"price,omitempty"`
}

type Marker struct {
	Marker  string `json:"marker,omitempty"`
	MaxKeys int    `json:"maxKeys,omitempty"`
}

type ListResultWithMarker struct {
	IsTruncated bool   `json:"isTruncated"`
	Marker      string `json:"marker"`
	MaxKeys     int    `json:"maxKeys"`
	NextMarker  string `json:"nextMarker"`
}

type RecycleInstance struct {
	InstanceID         string           `json:"cacheClusterShowId"`
	InstanceName       string           `json:"instanceName"`
	InstanceStatus     string           `json:"instanceStatus"`
	IsolatedStatus     string           `json:"isolatedStatus"`
	ClusterType        string           `json:"clusterType"`
	Engine             string           `json:"engine"`
	EngineVersion      string           `json:"engineVersion"`
	VnetIP             string           `json:"vnetIp"`
	Domain             string           `json:"domain"`
	Port               string           `json:"port"`
	InstanceCreateTime string           `json:"instanceCreateTime"`
	Capacity           int              `json:"capacity"`
	UsedCapacity       float64          `json:"usedCapacity"`
	PaymentTiming      string           `json:"paymentTiming"`
	ZoneNames          []string         `json:"zoneNames"`
	Tags               []model.TagModel `json:"tags"`
}

type RecyclerInstanceList struct {
	ListResultWithMarker
	Result []RecycleInstance `json:"result"`
}

type BatchInstanceIds struct {
	InstanceIds []string `json:"cacheClusterShowIds,omitempty"`
}

type RenewInstanceArgs struct {
	Duration    int      `json:"duration,omitempty"`
	InstanceIds []string `json:"instanceIds,omitempty"`
}

type OrderIdResult struct {
	OrderId string `json:"orderId"`
}

type ListLogArgs struct {
	FileType  string `json:"fileType"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type ListLogResult struct {
	LogList []ShardLog `json:"logList"`
}
type LogItem struct {
	LogStartTime    string `json:"logStartTime"`
	LogEndTime      string `json:"logEndTime"`
	DownloadURL     string `json:"downloadUrl"`
	LogID           string `json:"logId"`
	LogSizeInBytes  int    `json:"logSizeInBytes"`
	DownloadExpires string `json:"downloadExpires"`
}
type ShardLog struct {
	ShardShowID string    `json:"shardShowId"`
	TotalNum    int       `json:"totalNum"`
	ShardID     int       `json:"shardId"`
	LogItem     []LogItem `json:"logItem"`
}

type GetLogArgs struct {
	ValidSeconds int `json:"validSeconds"`
}
type GetMaintainTimeResult struct {
	CacheClusterShowId string       `json:"cacheClusterShowId"`
	MaintainTime       MaintainTime `json:"maintainTime"`
}
type MaintainTime struct {
	StartTime string `json:"startTime"`
	Duration  int    `json:"duration"`
	Period    []int  `json:"period"`
}

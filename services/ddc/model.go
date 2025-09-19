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

package ddc

import "github.com/baidubce/bce-sdk-go/model"

type CreateInstanceArgs struct {
	ClientToken  string         `json:"-"`
	InstanceType string         `json:"instanceType"`
	Number       int            `json:"number"`
	Instance     CreateInstance `json:"instance"`
}

type CreateRdsArgs struct {
	ClientToken       string           `json:"-"`
	Billing           Billing          `json:"billing"`
	PurchaseCount     int              `json:"purchaseCount,omitempty"`
	InstanceName      string           `json:"instanceName,omitempty"`
	Engine            string           `json:"engine"`
	EngineVersion     string           `json:"engineVersion"`
	Category          string           `json:"category,omitempty"`
	CpuCount          int              `json:"cpuCount"`
	MemoryCapacity    int              `json:"memoryCapacity"`
	VolumeCapacity    int              `json:"volumeCapacity"`
	ZoneNames         []string         `json:"zoneNames,omitempty"`
	VpcId             string           `json:"vpcId,omitempty"`
	IsDirectPay       bool             `json:"isDirectPay,omitempty"`
	Subnets           []SubnetMap      `json:"subnets,omitempty"`
	Tags              []model.TagModel `json:"tags,omitempty"`
	AutoRenewTimeUnit string           `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int              `json:"autoRenewTime,omitempty"`
	DeployId          string           `json:"deployId"`
	PoolId            string           `json:"poolId"`
	SyncMode          string           `json:"syncMode"`
}

type CreateReadReplicaArgs struct {
	ClientToken         string           `json:"-"`
	Billing             Billing          `json:"billing"`
	PurchaseCount       int              `json:"purchaseCount,omitempty"`
	SourceInstanceId    string           `json:"sourceInstanceId"`
	InstanceName        string           `json:"instanceName,omitempty"`
	CpuCount            int              `json:"cpuCount"`
	MemoryCapacity      int              `json:"memoryCapacity"`
	VolumeCapacity      int              `json:"volumeCapacity"`
	ZoneNames           []string         `json:"zoneNames,omitempty"`
	VpcId               string           `json:"vpcId,omitempty"`
	IsDirectPay         bool             `json:"isDirectPay,omitempty"`
	Subnets             []SubnetMap      `json:"subnets,omitempty"`
	Tags                []model.TagModel `json:"tags,omitempty"`
	DeployId            string           `json:"deployId"`
	PoolId              string           `json:"poolId"`
	RoGroupId           string           `json:"roGroupId"`
	EnableDelayOff      bool             `json:"enableDelayOff"`
	DelayThreshold      int              `json:"delayThreshold"`
	LeastInstanceAmount int              `json:"leastInstanceAmount"`
	RoGroupWeight       int              `json:"roGroupWeight"`
}

type Instance struct {
	InstanceId         string       `json:"instanceId"`
	InstanceName       string       `json:"instanceName"`
	Engine             string       `json:"engine"`
	EngineVersion      string       `json:"engineVersion"`
	Category           string       `json:"category"`
	InstanceStatus     string       `json:"instanceStatus"`
	CpuCount           int          `json:"cpuCount"`
	MemoryCapacity     float64      `json:"allocatedMemoryInGB"`
	VolumeCapacity     int          `json:"allocatedStorageInGB"`
	NodeAmount         int          `json:"nodeAmount"`
	UsedStorage        float64      `json:"usedStorageInGB"`
	PublicAccessStatus bool         `json:"publicAccessStatus"`
	InstanceCreateTime string       `json:"instanceCreateTime"`
	InstanceExpireTime string       `json:"instanceExpireTime"`
	Endpoint           Endpoint     `json:"endpoint"`
	SyncMode           string       `json:"syncMode"`
	BackupPolicy       BackupPolicy `json:"backupPolicy"`
	Region             string       `json:"region"`
	InstanceType       string       `json:"type"`
	SourceInstanceId   string       `json:"sourceInstanceId"`
	SourceRegion       string       `json:"sourceRegion"`
	ZoneNames          []string     `json:"zoneNames"`
	VpcId              string       `json:"vpcId"`
	Subnets            []Subnet     `json:"subnets"`
	Topology           Topology     `json:"topology"`
	PaymentTiming      string       `json:"paymentTiming"`
	RoGroupList        []RoGroup    `json:"roGroupList"`
	NodeMaster         NodeInfo     `json:"nodeMaster"`
	NodeSlave          NodeInfo     `json:"nodeSlave"`
	NodeReadReplica    NodeInfo     `json:"nodeReadReplica"`
	DeployId           string       `json:"deployId"`
}

type SubnetMap struct {
	ZoneName string `json:"zoneName"`
	SubnetId string `json:"subnetId"`
}

type Billing struct {
	PaymentTiming string      `json:"paymentTiming"`
	Reservation   Reservation `json:"reservation,omitempty"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength,omitempty"`
	ReservationTimeUnit string `json:"reservationTimeUnit,omitempty"`
}

type CreateResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type InstanceModelResult struct {
	Instance InstanceModel `json:"instance"`
}

type CreateInstance struct {
	InstanceId           string           `json:"instanceId"`
	InstanceName         string           `json:"instanceName"`
	SourceInstanceId     string           `json:"sourceInstanceId"`
	Engine               string           `json:"engine"`
	EngineVersion        string           `json:"engineVersion"`
	CpuCount             int              `json:"cpuCount"`
	AllocatedMemoryInGB  int              `json:"allocatedMemoryInGB"`
	AllocatedStorageInGB int              `json:"allocatedStorageInGB"`
	AZone                string           `json:"azone"`
	VpcId                string           `json:"vpcId"`
	SubnetId             string           `json:"subnetId"`
	DiskIoType           string           `json:"diskIoType"`
	DeployId             string           `json:"deployId"`
	PoolId               string           `json:"poolId"`
	RoGroupId            string           `json:"roGroupId"`
	EnableDelayOff       bool             `json:"enableDelayOff"`
	DelayThreshold       int              `json:"delayThreshold"`
	LeastInstanceAmount  int              `json:"leastInstanceAmount"`
	RoGroupWeight        int              `json:"roGroupWeight"`
	IsDirectPay          bool             `json:"IsDirectPay"`
	Billing              Billing          `json:"billing"`
	AutoRenewTimeUnit    string           `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime        int              `json:"autoRenewTime,omitempty"`
	Category             string           `json:"category,omitempty"`
	Tags                 []model.TagModel `json:"tags,omitempty"`
	SyncMode             string           `json:"syncMode"`
}

type Pool struct {
	CPUQuotaTotal      int    `json:"cpuQuotaTotal"`
	CPUQuotaUsed       int    `json:"cpuQuotaUsed"`
	CreateTime         string `json:"createTime"`
	DeployMethod       string `json:"deployMethod"`
	DiskQuotaTotal     int    `json:"diskQuotaTotal"`
	DiskQuotaUsed      int    `json:"diskQuotaUsed"`
	Engine             string `json:"engine"`
	Hosts              []Host `json:"hosts"`
	MaxMemoryUsedRatio string `json:"maxMemoryUsedRatio"`
	MemoryQuotaTotal   int    `json:"memoryQuotaTotal"`
	MemoryQuotaUsed    int    `json:"memoryQuotaUsed"`
	PoolID             string `json:"poolId"`
	PoolName           string `json:"poolName"`
	VpcID              string `json:"vpcId"`
}

type Host struct {
	Containers       []Container `json:"containers"`
	Flavor           Flavor      `json:"flavor"`
	CPUQuotaTotal    int         `json:"cpuQuotaTotal"`
	CPUQuotaUsed     int         `json:"cpuQuotaUsed"`
	DeploymentStatus string      `json:"deploymentStatus"`
	DiskQuotaTotal   int         `json:"diskQuotaTotal"`
	DiskQuotaUsed    int         `json:"diskQuotaUsed"`
	HostID           string      `json:"hostId"`
	HostName         string      `json:"hostName"`
	ImageType        string      `json:"imageType"`
	MemoryQuotaTotal int64       `json:"memoryQuotaTotal"`
	MemoryQuotaUsed  int64       `json:"memoryQuotaUsed"`
	PnetIP           string      `json:"pnetIp"`
	Role             string      `json:"role"`
	Status           string      `json:"status"`
	SubnetID         string      `json:"subnetId"`
	VnetIP           string      `json:"vnetIp"`
	VpcID            string      `json:"vpcId"`
	Zone             string      `json:"zone"`
}

type OperateHostRequest struct {
	Action string `json:"action"`
}

type Flavor struct {
	CPUCount           int    `json:"cpuCount"`
	CPUType            string `json:"cpuType"`
	Disk               int    `json:"disk"`
	FlavorID           string `json:"flavorId"`
	MemoryCapacityInGB int    `json:"memoryCapacityInGB"`
}

type Container struct {
	ContainerID string `json:"containerId"`
	DeployID    string `json:"deployId"`
	DeployName  string `json:"deployName"`
	Engine      string `json:"engine"`
	HostID      string `json:"hostId"`
	HostName    string `json:"hostName"`
	PoolName    string `json:"poolName"`
	Role        string `json:"role"`
	Zone        string `json:"zone"`
}

type DeploySet struct {
	CreateTime          string   `json:"createTime"`
	DeployID            string   `json:"deployId"`
	DeployName          string   `json:"deployName"`
	Instances           []string `json:"instances"`
	PoolID              string   `json:"poolId"`
	Strategy            string   `json:"strategy"`
	CentralizeThreshold int      `json:"centralizeThreshold"`
}

type CreateDeployRequest struct {
	ClientToken         string `json:"-"`
	DeployName          string `json:"deployName"`
	Strategy            string `json:"strategy"`
	CentralizeThreshold int    `json:"centralizeThreshold"`
}

type UpdateDeployRequest struct {
	ClientToken         string `json:"-"`
	Strategy            string `json:"strategy"`
	CentralizeThreshold int    `json:"centralizeThreshold"`
}

type Marker struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListResultWithMarker struct {
	IsTruncated bool   `json:"isTruncated"`
	Marker      string `json:"marker"`
	MaxKeys     int    `json:"maxKeys"`
	NextMarker  string `json:"nextMarker"`
}

type ListPoolResult struct {
	ListResultWithMarker
	Result []Pool `json:"result"`
}

type ListHostResult struct {
	ListResultWithMarker
	Result []Host `json:"result"`
}

type ListDeploySetResult struct {
	ListResultWithMarker
	Result []DeploySet `json:"result"`
}

type InstanceModel struct {
	InstanceId           string       `json:"instanceId"`
	InstanceName         string       `json:"instanceName"`
	Engine               string       `json:"engine"`
	EngineVersion        string       `json:"engineVersion"`
	InstanceStatus       string       `json:"instanceStatus"`
	CpuCount             int          `json:"cpuCount"`
	AllocatedMemoryInGB  float64      `json:"allocatedMemoryInGB"`
	AllocatedStorageInGB int          `json:"allocatedStorageInGB"`
	NodeAmount           int          `json:"nodeAmount"`
	UsedStorageInGB      float64      `json:"usedStorageInGB"`
	PublicAccessStatus   bool         `json:"publicAccessStatus"`
	InstanceCreateTime   string       `json:"instanceCreateTime"`
	InstanceExpireTime   string       `json:"instanceExpireTime"`
	Endpoint             Endpoint     `json:"endpoint"`
	SyncMode             string       `json:"syncMode"`
	BackupPolicy         BackupPolicy `json:"backupPolicy"`
	Region               string       `json:"region"`
	InstanceType         string       `json:"instanceType"`
	SourceInstanceId     string       `json:"sourceInstanceId"`
	SourceRegion         string       `json:"sourceRegion"`
	ZoneNames            []string     `json:"zoneNames"`
	VpcId                string       `json:"vpcId"`
	Subnets              []Subnet     `json:"subnets"`
	NodeMaster           NodeInfo     `json:"nodeMaster"`
	NodeSlave            NodeInfo     `json:"nodeSlave"`
	NodeReadReplica      NodeInfo     `json:"nodeReadReplica"`
	DeployId             string       `json:"deployId"`
	Topology             Topology     `json:"topology"`
	DiskType             string       `json:"diskType"`
	Type                 string       `json:"type"`
	ApplicationType      string       `json:"applicationType"`
	RoGroupList          []RoGroup    `json:"roGroupList"`
	PaymentTiming        string       `json:"paymentTiming"`
}

type SubnetVo struct {
	Name     string `json:"name"`
	SubnetId string `json:"subnetId"`
	Az       string `json:"az"`
	Cidr     string `json:"cidr"`
	ShortId  string `json:"shortId"`
}

type RoGroup struct {
	RoGroupID           string    `json:"roGroupId"`
	RoGroupName         string    `json:"roGroupName"`
	VnetIP              string    `json:"vnetIp"`
	IsBalanceRoLoad     int       `json:"isBalanceRoLoad"`
	EnableDelayOff      int       `json:"enableDelayOff"`
	DelayThreshold      int       `json:"delayThreshold"`
	LeastInstanceAmount int       `json:"leastInstanceAmount"`
	ReplicaList         []Replica `json:"replicaList"`
}

type UpdateRoGroupArgs struct {
	RoGroupName         string `json:"roGroupName"`
	IsBalanceRoLoad     int    `json:"isBalanceRoLoad"`
	EnableDelayOff      int    `json:"enableDelayOff"`
	DelayThreshold      int    `json:"delayThreshold"`
	LeastInstanceAmount int    `json:"leastInstanceAmount"`
}

type UpdateRoGroupWeightArgs struct {
	IsBalanceRoLoad int             `json:"isBalanceRoLoad"`
	ReplicaList     []ReplicaWeight `json:"replicaList"`
}
type ReplicaWeight struct {
	InstanceId string `json:"instanceId"`
	Weight     int    `json:"weight"`
}

type Replica struct {
	InstanceId    string `json:"instanceId"`
	InstanceName  string `json:"instanceName"`
	Status        string `json:"status"`
	RoGroupWeight int    `json:"roGroupWeight"`
}

type NodeInfo struct {
	Id       string `json:"id"`
	Azone    string `json:"azone"`
	SubnetId string `json:"subnetId"`
	Cidr     string `json:"cidr"`
	Name     string `json:"name"`
	HostName string `json:"hostname"`
}

type Subnet struct {
	Name        string `json:"name"`
	SubnetId    string `json:"subnetId"`
	ZoneName    string `json:"zoneName"`
	Cidr        string `json:"cidr"`
	ShortId     string `json:"shortId"`
	VpcId       string `json:"vpcId"`
	VpcShortId  string `json:"vpcShortId"`
	Az          string `json:"az"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`
}

type Endpoint struct {
	Address      string `json:"address"`
	Port         int    `json:"port"`
	VnetIp       string `json:"vnetIp"`
	VnetIpBackup string `json:"vnetIpBackup"`
	InetIp       string `json:"inetIp"`
}

type BackupPolicy struct {
	BackupDays    string `json:"backupDays"`
	BackupTime    string `json:"backupTime"`
	Persistent    bool   `json:"persistent"`
	ExpireInDays  int    `json:"expireInDays"`
	FreeSpaceInGB int    `json:"freeSpaceInGb"`
}

type Topology struct {
	Rdsproxy    []string `json:"rdsproxy"`
	Master      []string `json:"master"`
	ReadReplica []string `json:"readReplica"`
}

type DeleteDdcArgs struct {
	InstanceIds []string `json:"instanceIds"`
}

type UpdateInstanceNameArgs struct {
	InstanceName string `json:"instanceName"`
}

type ListRdsResult struct {
	Marker      string     `json:"marker"`
	MaxKeys     int        `json:"maxKeys"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	Instances   []Instance `json:"result"`
}

type ListRdsArgs struct {
	Marker  string
	MaxKeys int
}

type GetBackupListResult struct {
	Snapshots     []Snapshot `json:"snapshots"`
	FreeSpaceInMB int64      `json:"freeSpaceInMB"`
	UsedSpaceInMB int64      `json:"usedSpaceInMB"`
}

type GetZoneListResult struct {
	Zones []ZoneName `json:"zones"`
}

type ZoneName struct {
	ZoneNames       []string `json:"apiZoneNames"`
	ApiZoneNames    []string `json:"zoneNames"`
	Available       bool     `json:"bool"`
	DefaultSubnetId string   `json:"defaultSubnetId"`
}

type ListSubnetsArgs struct {
	VpcId    string `json:"vpcId"`
	ZoneName string `json:"zoneName"`
}

type ListSubnetsResult struct {
	Subnets []Subnet `json:"subnets"`
}

type SecurityIpsRawResult struct {
	SecurityIps []string `json:"ip"`
}

type UpdateSecurityIpsArgs struct {
	SecurityIps []string `json:"securityIps"`
}

type ListParametersResult struct {
	Items []Parameter `json:"items"`
}

type Parameter struct {
	Name          string `json:"name"`
	DefaultValue  string `json:"defaultValue"`
	Value         string `json:"value"`
	PendingValue  string `json:"pendingValue"`
	Type          string `json:"type"`
	Dynamic       bool   `json:"dynamic"`
	Modifiable    bool   `json:"modifiable"`
	AllowedValues string `json:"allowedValues"`
	Desc          string `json:"desc"`
}

type UpdateParameterArgs struct {
	Parameters []KVParameter `json:"parameters"`
}

type KVParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Snapshot struct {
	SnapshotId          string `json:"snapshotId"`
	SnapshotSizeInBytes string `json:"snapshotSizeInBytes"`
	SnapshotType        string `json:"snapshotType"`
	SnapshotStatus      string `json:"snapshotStatus"`
	SnapshotStartTime   string `json:"snapshotStartTime"`
	SnapshotEndTime     string `json:"snapshotEndTime"`
}

type SnapshotModel struct {
	SnapshotId          string `json:"snapshotId"`
	SnapshotSizeInBytes string `json:"snapshotSizeInBytes"`
	SnapshotType        string `json:"snapshotType"`
	SnapshotStatus      string `json:"snapshotStatus"`
	SnapshotStartTime   string `json:"snapshotStartTime"`
	SnapshotEndTime     string `json:"snapshotEndTime"`
	DownloadUrl         string `json:"downloadUrl"`
	DownloadExpires     string `json:"downloadExpires"`
	Progress            string `json:"progress"`
}

type BackupDetailResult struct {
	Snapshot SnapshotModel `json:"snapshot"`
}

type Binlog struct {
	BinlogId          string `json:"binlogId"`
	BinlogSizeInBytes int64  `json:"binlogSizeInBytes"`
	BinlogStatus      string `json:"binlogStatus"`
	BinlogStartTime   string `json:"binlogStartTime"`
	BinlogEndTime     string `json:"binlogEndTime"`
}

type BinlogModel struct {
	BinlogId          string `json:"binlogId"`
	BinlogSizeInBytes int64  `json:"binlogSizeInBytes"`
	BinlogStatus      string `json:"binlogStatus"`
	BinlogStartTime   string `json:"binlogStartTime"`
	BinlogEndTime     string `json:"binlogEndTime"`
	DownloadUrl       string `json:"downloadUrl"`
	DownloadExpires   string `json:"downloadExpires"`
}

type BinlogListResult struct {
	Binlogs []Binlog `json:"binlogs"`
}

type BinlogDetailResult struct {
	Binlog BinlogModel `json:"binlog"`
}

type AuthType string

const (
	AuthType_ReadOnly  AuthType = "readOnly"
	AuthType_ReadWrite AuthType = "readWrite"
)

type AccountPrivilege struct {
	AccountName string   `json:"accountName"`
	AuthType    AuthType `json:"authType"`
}

type CreateDatabaseArgs struct {
	ClientToken      string `json:"-"`
	DbName           string `json:"dbName"`
	CharacterSetName string `json:"characterSetName"`
	Remark           string `json:"remark"`
}

type UpdateDatabaseRemarkArgs struct {
	Remark string `json:"remark"`
}

type Database struct {
	DbName            string             `json:"dbName"`
	CharacterSetName  string             `json:"characterSetName"`
	DbStatus          string             `json:"dbStatus"`
	Remark            string             `json:"remark"`
	AccountPrivileges []AccountPrivilege `json:"accountPrivileges"`
}

type DatabaseResult struct {
	Database Database `json:"database"`
}

type ListDatabaseResult struct {
	Databases []Database `json:"databases"`
}

// Account
type AccountType string

const (
	AccountType_Super  AccountType = "rdssuper"
	AccountType_Common AccountType = "common"
)

type CreateAccountArgs struct {
	ClientToken string `json:"-"`
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
	// 为了兼容 RDS 参数结构
	AccountType        AccountType         `json:"type"`
	Desc               string              `json:"remark"`
	DatabasePrivileges []DatabasePrivilege `json:"databasePrivileges,omitempty"`
}

type DatabasePrivilege struct {
	DbName   string `json:"dbName"`
	AuthType string `json:"authType"`
}

type Account struct {
	AccountName        string              `json:"accountName"`
	Desc               string              `json:"remark"`
	Status             string              `json:"accountStatus"`
	AccountType        string              `json:"accountType"`
	DatabasePrivileges []DatabasePrivilege `json:"databasePrivileges"`
}

type AccountResult struct {
	Account Account `json:"account"`
}

type ListAccountResult struct {
	Accounts []Account `json:"accounts"`
}

type UpdateAccountPasswordArgs struct {
	Password string `json:"password"`
}

type UpdateAccountDescArgs struct {
	Desc string `json:"remark"`
}

type UpdateAccountPrivilegesArgs struct {
	DatabasePrivileges []DatabasePrivilege `json:"databasePrivileges"`
}

type ListRoGroupResult struct {
	RoGroups []RoGroup `json:"roGroups"`
}

type VpcVo struct {
	VpcId         string   `json:"vpcId"`
	ShortId       string   `json:"shortId"`
	Name          string   `json:"name"`
	Cidr          string   `json:"cidr"`
	Status        int      `json:"status"`
	CreateTime    string   `json:"createTime"`
	Description   string   `json:"description"`
	DefaultVpc    bool     `json:"defaultVpc"`
	Ipv6Cidr      string   `json:"ipv6Cidr"`
	AuxiliaryCidr []string `json:"auxiliaryCidr"`
	Relay         bool     `json:"relay"`
}
type GetBackupListArgs struct {
	Marker  string
	MaxKeys int
}
type GetSecurityIpsResult struct {
	Etag        string   `json:"etag"`
	SecurityIps []string `json:"securityIps"`
}

type ResizeRdsArgs struct {
	CpuCount       int     `json:"cpuCount"`
	MemoryCapacity float64 `json:"memoryCapacity"`
	VolumeCapacity int     `json:"volumeCapacity"`
	NodeAmount     int     `json:"nodeAmount,omitempty"`
	IsDirectPay    bool    `json:"isDirectPay,omitempty"`
	IsResizeNow    bool    `json:"isResizeNow,omitempty"`
}

type RebootArgs struct {
	IsRebootNow bool `json:"isRebootNow"`
}

type SwitchArgs struct {
	IsSwitchNow bool `json:"isSwitchNow"`
}

type MaintainWindow struct {
	MaintainTime MaintainTime `json:"maintentime"`
}

type MaintainTime struct {
	Period    string `json:"period"`
	StartTime string `json:"startTime"`
	Duration  int    `json:"duration"`
}

type RecycleInstance struct {
	EngineVersion      string  `json:"engineVersion"`
	VolumeCapacity     int     `json:"volumeCapacity"`
	ApplicationType    string  `json:"applicationType"`
	InstanceName       string  `json:"instanceName"`
	PublicAccessStatus string  `json:"publicAccessStatus"`
	InstanceCreateTime string  `json:"instanceCreateTime"`
	InstanceType       string  `json:"instanceType"`
	Type               string  `json:"type"`
	InstanceStatus     string  `json:"instanceStatus"`
	MemoryCapacity     float64 `json:"memoryCapacity"`
	InstanceId         string  `json:"instanceId"`
	Engine             string  `json:"engine"`
	VpcId              string  `json:"vpcId"`
	PubliclyAccessible bool    `json:"publiclyAccessible"`
	InstanceExpireTime string  `json:"instanceExpireTime"`
	DiskType           string  `json:"diskType"`
	Region             string  `json:"region"`
	CpuCount           int     `json:"cpuCount"`
	UsedStorage        float64 `json:"usedStorage"`
}

type RecyclerInstanceList struct {
	ListResultWithMarker
	Result []RecycleInstance `json:"result"`
}

type BatchInstanceIds struct {
	InstanceIds string `json:"instanceIds"`
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

type ListLogArgs struct {
	LogType  string `json:"logType"`
	Datetime string `json:"datetime"`
}

type Log struct {
	LogStartTime   string `json:"logStartTime"`
	LogEndTime     string `json:"logEndTime"`
	LogID          string `json:"logId"`
	LogSizeInBytes int    `json:"logSizeInBytes"`
}

type LogDetail struct {
	Log
	DownloadURL     string `json:"downloadUrl"`
	DownloadExpires string `json:"downloadExpires"`
}

type GetLogArgs struct {
	ValidSeconds int `json:"downloadValidTimeInSec"`
}

type CreateTableHardLinkArgs struct {
	TableName string `json:"tableName"`
}

type ModifySyncModeArgs struct {
	SyncMode string `json:"syncMode"`
}
type Disk struct {
	CapacityRatio []string `json:"capacityRatio"`
}

type AccessDetailItem struct {
	BackupID             string `json:"backupID"`
	AccessDateTime       string `json:"accessDateTime"`
	AccessResult         string `json:"accessResult"`
	AccessSrcAddressType string `json:"accessSrcAddressType"`
	AvailableZone        string `json:"availableZone"`
	AccessSrcAddress     string `json:"accessSrcAddress"`
	AccessOperationType  string `json:"accessOperationType"`
	StorageType          string `json:"storageType"`
	StorageAddress       string `json:"storageAddress"`
	Region               string `json:"region"`
	BackupName           string `json:"backupName"`
	AccessSrcAgent       string `json:"accessSrcAgent"`
	StorageID            string `json:"storageID"`
}

type Pagination struct {
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
	Marker      string `json:"marker"`
	TotalKeys   int    `json:"totalKeys"`
}
type BackupAccessDetail struct {
	StartDateTime       string             `json:"startDateTime"`
	EndDateTime         string             `json:"endDateTime"`
	DataType            string             `json:"dataType"`
	BackupAccessDetails []AccessDetailItem `json:"backupAccessDetails"`
	Pagination          Pagination         `json:"pagination"`
}

type AccessDetailArgs struct {
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
	Marker        string `json:"marker,omitempty"`
	MaxKeys       int    `json:"maxKeys,omitempty"`
	DataType      string `json:"dataType,omitempty"`
}

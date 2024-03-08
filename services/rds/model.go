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

package rds

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type CreateRdsArgs struct {
	ClientToken          string                          `json:"-"`
	Billing              Billing                         `json:"billing"`
	PurchaseCount        int                             `json:"purchaseCount,omitempty"`
	InstanceName         string                          `json:"instanceName,omitempty"`
	Engine               string                          `json:"engine"`
	EngineVersion        string                          `json:"engineVersion"`
	Category             string                          `json:"category,omitempty"`
	CpuCount             int                             `json:"cpuCount"`
	MemoryCapacity       float64                         `json:"memoryCapacity"`
	VolumeCapacity       int                             `json:"volumeCapacity"`
	DiskIoType           string                          `json:"diskIoType"`
	ZoneNames            []string                        `json:"zoneNames,omitempty"`
	VpcId                string                          `json:"vpcId,omitempty"`
	IsDirectPay          bool                            `json:"isDirectPay,omitempty"`
	Subnets              []SubnetMap                     `json:"subnets,omitempty"`
	Tags                 []model.TagModel                `json:"tags,omitempty"`
	AutoRenewTimeUnit    string                          `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime        int                             `json:"autoRenewTime,omitempty"`
	BgwGroupId           string                          `json:"bgwGroupId,omitempty"`
	BgwGroupExclusive    bool                            `json:"bgwGroupExclusive,omitempty"`
	CharacterSetName     string                          `json:"characterSetName,omitempty"`
	LowerCaseTableNames  int                             `json:"lowerCaseTableNames,omitempty"`
	ParameterTemplateId  string                          `json:"parameterTemplateId,omitempty"`
	Ovip                 string                          `json:"ovip,omitempty"`
	EntryPort            string                          `json:"entryPort,omitempty"`
	ReplicationType      string                          `json:"replicationType,omitempty"`
	ResourceGroupId      string                          `json:"resourceGroupId,omitempty"`
	InitialDataReference *InitialData                    `json:"initialDataReference,omitempty"`
	Data                 []RecoveryToSourceInstanceModel `json:"data,omitempty"`
}

type Billing struct {
	PaymentTiming string      `json:"paymentTiming"`
	Reservation   Reservation `json:"reservation,omitempty"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength,omitempty"`
	ReservationTimeUnit string `json:"reservationTimeUnit,omitempty"`
}

type SubnetMap struct {
	ZoneName string `json:"zoneName"`
	SubnetId string `json:"subnetId"`
}

type InitialData struct {
	InstanceId    string `json:"instanceId,omitempty"`
	ReferenceType string `json:"referenceType,omitempty"`
	Datetime      string `json:"datetime,omitempty"`
	SnapshotId    string `json:"snapshotId,omitempty"`
}

type RecoveryToSourceInstanceModel struct {
	RestoreMode string  `json:"restoreMode,omitempty"`
	DbName      string  `json:"dbName,omitempty"`
	NewDbname   string  `json:"newDbname,omitempty"`
	Tables      []Table `json:"tables,omitempty"`
}

type Table struct {
	TableName    string `json:"tableName,omitempty"`
	NewTablename string `json:"newTablename,omitempty"`
}
type CreateResult struct {
	InstanceIds []string `json:"instanceIds"`
	OrderId     string   `json:"orderId"`
}

type CreateReadReplicaArgs struct {
	ClientToken      string           `json:"-"`
	Billing          Billing          `json:"billing"`
	PurchaseCount    int              `json:"purchaseCount,omitempty"`
	SourceInstanceId string           `json:"sourceInstanceId"`
	InstanceName     string           `json:"instanceName,omitempty"`
	CpuCount         int              `json:"cpuCount"`
	MemoryCapacity   float64          `json:"memoryCapacity"`
	VolumeCapacity   int              `json:"volumeCapacity"`
	ZoneNames        []string         `json:"zoneNames,omitempty"`
	VpcId            string           `json:"vpcId,omitempty"`
	IsDirectPay      bool             `json:"isDirectPay,omitempty"`
	Subnets          []SubnetMap      `json:"subnets,omitempty"`
	Tags             []model.TagModel `json:"tags,omitempty"`
	DiskIoType       string           `json:"diskIoType,omitempty"`
	Ovip             string           `json:"ovip,omitempty"`
	EntryPort        string           `json:"entryPort,omitempty"`
	ResourceGroupId  string           `json:"resourceGroupId,omitempty"`
}

type CreateRdsProxyArgs struct {
	ClientToken      string           `json:"-"`
	Billing          Billing          `json:"billing"`
	SourceInstanceId string           `json:"sourceInstanceId"`
	InstanceName     string           `json:"instanceName,omitempty"`
	NodeAmount       int              `json:"nodeAmount"`
	ZoneNames        []string         `json:"zoneNames,omitempty"`
	VpcId            string           `json:"vpcId,omitempty"`
	IsDirectPay      bool             `json:"isDirectPay,omitempty"`
	Subnets          []SubnetMap      `json:"subnets,omitempty"`
	Tags             []model.TagModel `json:"tags,omitempty"`
	Ovip             string           `json:"ovip,omitempty"`
	EntryPort        string           `json:"entryPort,omitempty"`
	ResourceGroupId  string           `json:"resourceGroupId,omitempty"`
}

type ListRdsArgs struct {
	Marker  string
	MaxKeys int
}

type Instance struct {
	InstanceId           string           `json:"instanceId"`
	InstanceName         string           `json:"instanceName"`
	Engine               string           `json:"engine"`
	EngineVersion        string           `json:"engineVersion"`
	RdsMinorVersion      string           `json:"rdsMinorVersion"`
	CharacterSetName     string           `json:"characterSetName"`
	InstanceClass        string           `json:"instanceClass"`
	AllocatedMemoryInMB  int              `json:"allocatedMemoryInMB"`
	AllocatedMemoryInGB  float64          `json:"allocatedMemoryInGB"`
	AllocatedStorageInGB int              `json:"allocatedStorageInGB"`
	Category             string           `json:"category"`
	InstanceStatus       string           `json:"instanceStatus"`
	CpuCount             int              `json:"cpuCount"`
	MemoryCapacity       float64          `json:"memoryCapacity"`
	VolumeCapacity       int              `json:"volumeCapacity"`
	TotalStorageInGB     int              `json:"totalStorageInGB"`
	NodeAmount           int              `json:"nodeAmount"`
	UsedStorage          float64          `json:"usedStorage"`
	PublicAccessStatus   string           `json:"publicAccessStatus"`
	InstanceCreateTime   string           `json:"instanceCreateTime"`
	InstanceExpireTime   string           `json:"instanceExpireTime"`
	Endpoint             Endpoint         `json:"endpoint"`
	SyncMode             string           `json:"syncMode"`
	BackupPolicy         BackupPolicy     `json:"backupPolicy"`
	Region               string           `json:"region"`
	InstanceType         string           `json:"instanceType"`
	SourceInstanceId     string           `json:"sourceInstanceId"`
	SourceRegion         string           `json:"sourceRegion"`
	ZoneNames            []string         `json:"zoneNames"`
	VpcId                string           `json:"vpcId"`
	Subnets              []Subnet         `json:"subnets"`
	Topology             Topology         `json:"topology"`
	Task                 string           `json:"task"`
	PaymentTiming        string           `json:"paymentTiming"`
	BgwGroupId           string           `json:"bgwGroupId"`
	ReadReplicaNum       int              `json:"readReplicaNum"`
	ReadReplica          []string         `json:"readReplica"`
	LockMode             string           `json:"lockMode"`
	EipStatus            string           `json:"eipStatus"`
	SuperUserFlag        string           `json:"superUserFlag"`
	ReplicationType      string           `json:"replicationType"`
	Azone                string           `json:"azone"`
	ApplicationType      string           `json:"applicationType"`
	OnlineStatus         int              `json:"onlineStatus"`
	IsSingle             bool             `json:"isSingle"`
	NodeType             string           `json:"nodeType"`
	DiskIoType           string           `json:"diskIoType"`
	GroupId              string           `json:"groupId"`
	GroupName            string           `json:"groupName"`
	DiskType             string           `json:"diskType"`
	CdsType              string           `json:"cdsType"`
	MaintainStartTime    string           `json:"maintainStartTime"`
	MaintainDuration     int              `json:"maintainDuration"`
	HaStrategy           int              `json:"haStrategy"`
	VpcName              string           `json:"vpcName"`
	Tags                 []model.TagModel `json:"tags"`
	ResourceGroupId      string           `json:"resourceGroupId"`
	ResourceGroupName    string           `json:"resourceGroupName"`
}

type ListRdsResult struct {
	Marker      string     `json:"marker"`
	MaxKeys     int        `json:"maxKeys"`
	IsTruncated bool       `json:"isTruncated"`
	NextMarker  string     `json:"nextMarker"`
	Instances   []Instance `json:"instances"`
}

type Subnet struct {
	Name     string `json:"name"`
	SubnetId string `json:"subnetId"`
	ZoneName string `json:"zoneName"`
	Cidr     string `json:"cidr"`
	VpcId    string `json:"vpcId"`
}

type Endpoint struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	VnetIp  string `json:"vnetIp"`
	InetIp  string `json:"inetIp"`
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

type ResizeRdsArgs struct {
	CpuCount            int         `json:"cpuCount"`
	MemoryCapacity      float64     `json:"memoryCapacity"`
	VolumeCapacity      int         `json:"volumeCapacity"`
	NodeAmount          int         `json:"nodeAmount,omitempty"`
	IsDirectPay         bool        `json:"isDirectPay,omitempty"`
	AllocatedMemoryInMB string      `json:"allocatedMemoryInMB,omitempty"`
	IsEnhanced          bool        `json:"isEnhanced,omitempty"`
	EffectiveTime       string      `json:"effectiveTime,omitempty"`
	MasterAzone         string      `json:"masterAzone,omitempty"`
	BackupAzone         string      `json:"backupAzone,omitempty"`
	DiskIoType          string      `json:"diskIoType,omitempty"`
	SubnetId            string      `json:"subnetId,omitempty"`
	EdgeSubnetId        string      `json:"edgeSubnetId,omitempty"`
	Subnets             []SubnetMap `json:"subnets,omitempty"`
}

type CreateAccountArgs struct {
	ClientToken        string              `json:"-"`
	AccountName        string              `json:"accountName"`
	Password           string              `json:"password"`
	AccountType        string              `json:"accountType,omitempty"`
	DatabasePrivileges []DatabasePrivilege `json:"databasePrivileges,omitempty"`
	Desc               string              `json:"desc,omitempty"`
	Type               string              `json:"type,omitempty"`
}

type DatabasePrivilege struct {
	DbName   string `json:"dbName"`
	AuthType string `json:"authType"`
}

type Account struct {
	AccountName        string              `json:"accountName"`
	Status             string              `json:"status"`
	Type               string              `json:"type"`
	AccountType        string              `json:"accountType"`
	DatabasePrivileges []DatabasePrivilege `json:"databasePrivileges"`
	Desc               string              `json:"desc"`
}

type ListAccountResult struct {
	Accounts []Account `json:"accounts"`
}

type ModifyAccountDesc struct {
	Remark string `json:"remark"`
}

type UpdateAccountPrivileges struct {
	DatabasePrivileges []DatabasePrivilege `json:"privileges"`
}

type UpdatePasswordArgs struct {
	Password string `json:"password"`
}
type UpdateInstanceNameArgs struct {
	InstanceName string `json:"instanceName"`
}

type ModifySyncModeArgs struct {
	SyncMode string `json:"syncMode"`
}

type ModifyEndpointArgs struct {
	Address string `json:"address"`
}

type ModifyPublicAccessArgs struct {
	PublicAccess bool `json:"publicAccess"`
}

type ModifyBackupPolicyArgs struct {
	BackupDays   string `json:"backupDays"`
	BackupTime   string `json:"backupTime"`
	Persistent   bool   `json:"persistent"`
	ExpireInDays int    `json:"expireInDays"`
}

type GetBackupListArgs struct {
	Marker  string
	MaxKeys int
}

type Snapshot struct {
	SnapshotId          string `json:"backupId"`
	SnapshotSizeInBytes int64  `json:"backupSize"`
	SnapshotType        string `json:"backupType"`
	SnapshotStatus      string `json:"backupStatus"`
	SnapshotStartTime   string `json:"backupStartTime"`
	SnapshotEndTime     string `json:"backupEndTime"`
	DownloadUrl         string `json:"downloadUrl"`
	DownloadExpires     string `json:"downloadExpires"`
}

type BackupDetail struct {
	BackupSize      int    `json:"backupSize"`
	BackupStatus    string `json:"backupStatus"`
	BackupId        string `json:"backupId"`
	BackupEndTime   string `json:"backupEndTime"`
	DownloadUrl     string `json:"downloadUrl"`
	BackupType      string `json:"backupType"`
	BackupStartTime string `json:"backupStartTime"`
	DownloadExpires string `json:"downloadExpires"`
}

type GetBackupListResult struct {
	Marker      string         `json:"marker"`
	MaxKeys     int            `json:"maxKeys"`
	IsTruncated bool           `json:"isTruncated"`
	NextMarker  string         `json:"nextMarker"`
	Backups     []BackupDetail `json:"backups"`
}

type GetBinlogListResult struct {
	Binlogs []Binlog `json:"binlogs"`
}

type GetBinlogInfoResult struct {
	Binlog BinlogInfo `json:"binlog"`
}

type BinlogInfo struct {
	DownloadUrl     string `json:"downloadUrl"`
	DownloadExpires string `json:"downloadExpires"`
}
type Binlog struct {
	BinlogId          string `json:"binlogId"`
	BinlogSizeInBytes int64  `json:"binlogSize"`
	BinlogStatus      string `json:"binlogStatus"`
	BinlogStartTime   string `json:"binlogStartTime"`
	BinlogEndTime     string `json:"binlogEndTime"`
}

type RecoveryByDatetimeArgs struct {
	Datetime string         `json:"dateTime"`
	Data     []RecoveryData `json:"data"`
}

type RecoveryBySnapshotArgs struct {
	SnapshotId string         `json:"snapshotId"`
	Data       []RecoveryData `json:"data"`
}
type RecoveryData struct {
	DbName      string      `json:"dbName"`
	NewDbname   string      `json:"newDbname"`
	RestoreMode string      `json:"restoreMode"`
	Tables      []TableData `json:"tables"`
}

type TableData struct {
	TableName    string `json:"tableName"`
	NewTablename string `json:"newTablename"`
}
type GetZoneListResult struct {
	Zones []ZoneName `json:"zones"`
}

type ZoneName struct {
	ZoneNames []string `json:"zoneNames"`
}

type ListSubnetsArgs struct {
	VpcId    string `json:"vpcId"`
	ZoneName string `json:"zoneName"`
}

type ListSubnetsResult struct {
	Subnets []Subnet `json:"subnets"`
}

type GetSecurityIpsResult struct {
	Etag        string   `json:"etag"`
	SecurityIps []string `json:"securityIps"`
}

type UpdateSecurityIpsArgs struct {
	SecurityIps []string `json:"securityIps"`
}

type ListParametersResult struct {
	Etag       string      `json:"etag"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name          string `json:"name"`
	DefaultValue  string `json:"defaultValue"`
	Value         string `json:"value"`
	PendingValue  string `json:"pendingValue"`
	Type          string `json:"type"`
	Dynamic       string `json:"dynamic"`
	Modifiable    string `json:"modifiable"`
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

type ParameterHistoryResult struct {
	Parameters []ParameterHistory `json:"parameters"`
}

type ParameterHistory struct {
	Name        string `json:"name"`
	BeforeValue string `json:"beforeValue"`
	AfterValue  string `json:"afterValue"`
	Status      string `json:"status"`
	UpdateTime  string `json:"updateTime"`
}
type AutoRenewArgs struct {
	InstanceIds       []string `json:"instanceIds"`
	AutoRenewTimeUnit string   `json:"autoRenewTimeUnit"`
	AutoRenewTime     int      `json:"autoRenewTime"`
}

type SlowLogDownloadTaskListResult struct {
	Slowlogs []Slowlog `json:"slowlogs"`
}

type SlowLogDownloadDetail struct {
	Slowlogs []SlowlogDetail `json:"slowlogs"`
}
type Slowlog struct {
	SlowlogId          string `json:"slowlogId"`
	SlowlogSizeInBytes int    `json:"slowlogSizeInBytes"`
	SlowlogStartTime   string `json:"slowlogStartTime"`
	SlowlogEndTime     string `json:"slowlogEndTime"`
}

type SlowlogDetail struct {
	Url             string `json:"url"`
	DownloadExpires string `json:"downloadExpires"`
}

type MaintainTimeArgs struct {
	MaintainStartTime string `json:"maintainStartTime"`
	MaintainDuration  int    `json:"maintainDuration"`
}

type DiskAutoResizeArgs struct {
	FreeSpaceThreshold int `json:"freeSpaceThreshold,omitempty"`
	DiskMaxLimit       int `json:"diskMaxLimit,omitempty"`
}

type AutoResizeConfigResult struct {
	AutoResizeDisk     int `json:"autoResizeDisk"`
	FreeSpaceThreshold int `json:"freeSpaceThreshold"`
	DiskMaxLimit       int `json:"diskMaxLimit"`
	ExtendStepPercent  int `json:"extendStepPercent"`
}

type EnableAutoExpansionResult struct {
	SupportEnableDiskAutoResize int `json:"supportEnableDiskAutoResize"`
}

type AzoneMigration struct {
	MasterAzone   string      `json:"master_azone"`
	BackupAzone   string      `json:"backup_azone"`
	ZoneNames     []string    `json:"zoneNames"`
	Subnets       []SubnetMap `json:"subnets"`
	EffectiveTime string      `json:"effectiveTime"`
}

type UpdateDatabasePortArgs struct {
	EntryPort int `json:"entryPort"`
}

type ListDatabasesResult struct {
	Databases []Database `json:"databases"`
}

type Database struct {
	DbName            string             `json:"dbName"`
	CharacterSetName  string             `json:"characterSetName"`
	Remark            string             `json:"remark"`
	DbStatus          string             `json:"dbStatus"`
	AccountPrivileges []AccountPrivilege `json:"accountPrivileges"`
}

type AccountPrivilege struct {
	AccountName string `json:"accountName"`
	AuthType    string `json:"authType"`
}

type ModifyDatabaseDesc struct {
	Remark string `json:"remark"`
}

type CreateDatabaseArgs struct {
	CharacterSetName  string             `json:"characterSetName"`
	DbName            string             `json:"dbName"`
	Remark            string             `json:"remark"`
	AccountPrivileges []AccountPrivilege `json:"accountPrivileges"`
}

type TaskListArgs struct {
	PageSize     string `json:"pageSize,omitempty"`
	PageNo       string `json:"pageNo,omitempty"`
	InstanceId   string `json:"instanceId,omitempty"`
	InstanceName string `json:"instanceName,omitempty"`
	TaskId       int    `json:"taskId,omitempty"`
	TaskType     string `json:"taskType,omitempty"`
	TaskStatus   string `json:"taskStatus,omitempty"`
	StartTime    string `json:"startTime,omitempty"`
	EndTime      string `json:"endTime,omitempty"`
}

type TaskListResult struct {
	Tasks []Task `json:"tasks"`
	Count int    `json:"count"`
}

type Task struct {
	TaskId       int            `json:"taskId"`
	TaskType     string         `json:"taskType"`
	TaskName     string         `json:"taskName"`
	InstanceId   string         `json:"instanceId"`
	InstanceName string         `json:"instanceName"`
	UserId       string         `json:"userId"`
	Region       string         `json:"region"`
	TaskStatus   string         `json:"taskStatus"`
	CreateTime   string         `json:"createTime"`
	UpdateTime   string         `json:"updateTime"`
	FinishTime   string         `json:"finishTime"`
	CancelFlag   int            `json:"cancelFlag"`
	Progress     []ProgressItem `json:"progress"`
}

type ProgressItem struct {
	Step        string `json:"step"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type RecyclerListResult struct {
	NextMarker    string     `json:"nextMarker"`
	Marker        string     `json:"marker"`
	MaxKeys       int        `json:"maxKeys"`
	IsTruncated   bool       `json:"isTruncated"`
	Instances     []Instance `json:"instances"`
	InstanceType  string     `json:"instanceType"`
	ZoneNames     []string   `json:"zoneNames"`
	PaymentTiming string     `json:"paymentTiming"`
}

type RecyclerRecoverArgs struct {
	InstanceIds []string `json:"instanceIds"`
}

type InstanceGroupArgs struct {
	Name     string `json:"name"`
	LeaderId string `json:"leaderId"`
}

type CreateInstanceGroupResult struct {
	Result int `json:"result"`
}

type ListInstanceGroupArgs struct {
	Manner   string `json:"manner"`
	Order    string `json:"order,omitempty"`
	OrderBy  string `json:"orderBy,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type InstanceGroupListResult struct {
	Result     []InstanceGroup `json:"result"`
	PageNo     int             `json:"pageNo"`
	PageSize   int             `json:"pageSize"`
	TotalCount int             `json:"totalCount"`
}

type InstanceGroup struct {
	GroupId string        `json:"groupId"`
	Name    string        `json:"name"`
	Count   int           `json:"count"`
	Leader  GroupInstance `json:"leader"`
}

type GroupInstance struct {
	InstanceIdShort string `json:"instanceIdShort"`
	Region          string `json:"region"`
	Azone           string `json:"azone"`
	Status          string `json:"status"`
	LockMode        string `json:"lockMode"`
	Name            string `json:"name"`
}

type InstanceGroupDetailResult struct {
	Group InstanceGroupDetail `json:"group"`
}

type InstanceGroupDetail struct {
	InstanceGroup
	Fllowers []GroupInstance `json:"fllowers"`
}

type CheckGtidArgs struct {
	InstanceId string `json:"instanceId"`
}

type CheckGtidResult struct {
	Result bool `json:"result"`
}

type CheckPingArgs struct {
	SourceId string `json:"sourceId"`
	TargetId string `json:"targetId"`
}

type CheckPingResult struct {
	Result bool `json:"result"`
}

type CheckDataArgs struct {
	InstanceId string `json:"instanceId"`
}

type CheckDataResult struct {
	Result bool `json:"result"`
}

type CheckVersionArgs struct {
	LeaderId   string `json:"leaderId"`
	FollowerId string `json:"followerId"`
}

type CheckVersionResult struct {
	Result bool `json:"result"`
}

type InstanceGroupNameArgs struct {
	Name string `json:"name"`
}

type InstanceGroupAddArgs struct {
	FollowerId string `json:"followerId"`
}

type InstanceGroupBatchAddArgs struct {
	FollowerIds []string `json:"followerIds"`
	Name        string   `json:"name"`
	LeaderId    string   `json:"leaderId"`
}

type ForceChangeArgs struct {
	LeaderId  string `json:"leaderId"`
	Force     int    `json:"force,omitempty"`
	MaxBehind int    `json:"maxBehind,omitempty"`
}

type ForceChangeResult struct {
	BehindMaster int `json:"behind_master"`
}

type GroupLeaderChangeArgs struct {
	LeaderId string `json:"leaderId"`
}

type MinorVersionListResult struct {
	RdsMinorVersionList []RdsMinorVersion `json:"rdsMinorVersionList"`
}

type RdsMinorVersion struct {
	DbVersion          string `json:"dbVersion"`
	MinorVersion       string `json:"minorVersion"`
	RdsMinorVersion    string `json:"rdsMinorVersion"`
	FeatureDescription string `json:"featureDescription"`
}

type UpgradeMinorVersionArgs struct {
	TargetMinorVersion string `json:"targetMinorVersion"`
	EffectiveTime      string `json:"effectiveTime"`
}

type SlowSqlFlowStatusResult struct {
	Enabled int `json:"enabled"`
}

type GetSlowSqlArgs struct {
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Sort     string `json:"sort,omitempty"`
	Schema   string `json:"schema,omitempty"`
	Digest   string `json:"digest,omitempty"`
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
}

type SlowSqlListResult struct {
	Items      []SlowSqlItem `json:"items"`
	TotalCount int           `json:"totalCount"`
}

type SlowSqlItem struct {
	AffectedRows int64   `json:"affectedRows"`
	ClientHost   string  `json:"clientHost"`
	ClientIP     string  `json:"clientIP"`
	Cluster      string  `json:"cluster"`
	ConnectionId int64   `json:"connectionId"`
	CurrentDB    string  `json:"currentDB"`
	Digest       string  `json:"digest"`
	Duration     float32 `json:"duration"`
	ExaminedRows int64   `json:"examinedRows"`
	LockTime     int64   `json:"lockTime"`
	Node         string  `json:"node"`
	NumRows      int     `json:"numRows"`
	Sql          string  `json:"sql"`
	SqlId        string  `json:"sqlId"`
	Start        string  `json:"start"`
	User         string  `json:"user"`
}

type SlowSqlExplainResult struct {
	List []SlowSqlExplainItem `json:"list"`
}

type SlowSqlExplainItem struct {
	ExplainId    int64  `json:"explainId"`
	Extra        string `json:"extra"`
	Filtered     int64  `json:"filtered"`
	Key          string `json:"key"`
	KeyLen       string `json:"keyLen"`
	Partitions   string `json:"partitions"`
	PossibleKeys string `json:"possibleKeys"`
	Ref          string `json:"ref"`
	Rows         int    `json:"rows"`
	SelectType   string `json:"selectType"`
	Table        string `json:"table"`
	Type         string `json:"type"`
}

type SlowSqlDigestResult struct {
	Items      []SlowSqlDigestItem `json:"items"`
	Summary    SlowSqlDigestItem   `json:"summary"`
	TotalCount int                 `json:"totalCount"`
}

type SlowSqlDigestItem struct {
	AvgExamRows   int64   `json:"avgExamRows"`
	AvgLockTime   int64   `json:"avgLockTime"`
	AvgNumRows    int64   `json:"avgNumRows"`
	AvgTime       float64 `json:"avgTime"`
	Digest        string  `json:"digest"`
	ExecuteTimes  int64   `json:"executeTimes"`
	MaxExamRows   int64   `json:"maxExamRows"`
	MaxLockTime   float64 `json:"maxLockTime"`
	MaxNumRows    int64   `json:"maxNumRows"`
	MaxTime       float64 `json:"maxTime"`
	NormalSql     string  `json:"normalSql"`
	Schema        string  `json:"schema"`
	TotalExamRows int64   `json:"totalExamRows"`
	TotalLockTime float64 `json:"totalLockTime"`
	TotalNumRows  int64   `json:"totalNumRows"`
	TotalTime     float64 `json:"totalTime"`
}

type GetSlowSqlDurationArgs struct {
	Schema string `json:"schema,omitempty"`
	Digest string `json:"digest,omitempty"`
	Start  string `json:"start,omitempty"`
	End    string `json:"end,omitempty"`
}

type SlowSqlDurationResult struct {
	List []SlowSqlDurationItem `json:"list"`
}

type SlowSqlDurationItem struct {
	End        int64   `json:"end"`
	Nums       int64   `json:"nums"`
	Percentage float64 `json:"percentage"`
	Start      int64   `json:"start"`
	Title      string  `json:"title"`
}

type GetSlowSqlSourceArgs struct {
	Schema string `json:"schema,omitempty"`
	Digest string `json:"digest,omitempty"`
	Start  string `json:"start,omitempty"`
	End    string `json:"end,omitempty"`
}

type SlowSqlSourceResult struct {
	List []SlowSqlSourceItem `json:"list"`
}

type SlowSqlSourceItem struct {
	Nums       int64   `json:"nums"`
	Percentage float64 `json:"percentage"`
	Host       string  `json:"host"`
	Ip         string  `json:"ip"`
}

type SlowSqlSchemaResult struct {
	List []SlowSqlSchemaItem `json:"list"`
}

type SlowSqlSchemaItem struct {
	Schema string `json:"schema"`
	Table  string `json:"table"`
}

type SlowSqlTableResult struct {
	List []SlowSqlTableItem `json:"list"`
}

type SlowSqlTableItem struct {
	Schema       string `json:"schema"`
	Table        string `json:"table"`
	Charset      string `json:"charset"`
	Collation    string `json:"collation"`
	Column       string `json:"column"`
	Comment      string `json:"comment"`
	DefaultValue string `json:"defaultValue"`
	Extra        string `json:"extra"`
	Key          string `json:"key"`
	Nullable     string `json:"nullable"`
	Position     int64  `json:"position"`
	Type         string `json:"type"`
}

type GetSlowSqlIndexArgs struct {
	SqlId  string `json:"sqlId"`
	Schema string `json:"schema"`
	Table  string `json:"table"`
	Index  string `json:"index,omitempty"`
}

type SlowSqlIndexResult struct {
	List []SlowSqlIndexItem `json:"list"`
}

type SlowSqlIndexItem struct {
	Schema      string `json:"schema"`
	Table       string `json:"table"`
	Collation   string `json:"collation"`
	Column      string `json:"column"`
	Comment     string `json:"comment"`
	Nullable    string `json:"nullable"`
	Type        string `json:"type"`
	Cardinality int64  `json:"cardinality"`
	Index       string `json:"index"`
	NonUnique   string `json:"nonUnique"`
	Sequence    string `json:"sequence"`
}

type GetSlowSqlTrendArgs struct {
	Schema   string `json:"schema,omitempty"`
	Interval string `json:"interval,omitempty"`
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
}

type SlowSqlTrendResult struct {
	Interval int                `json:"interval"`
	Items    []SlowSqlTrendItem `json:"items"`
}

type SlowSqlTrendItem struct {
	Datetime string `json:"datetime"`
	Times    int64  `json:"times"`
}

type SlowSqlAdviceResult struct {
	IndexAdvice     []SlowSqlIndexAdviceItem `json:"indexAdvice"`
	StatementAdvice []SlowSqlIndexAdviceItem `json:"statementAdvice"`
}

type SlowSqlIndexAdviceItem struct {
	Advice string `json:"advice"`
	Level  string `json:"level"`
}

type DiskInfoResult struct {
	Result DiskInfo `json:"result"`
}

type DiskInfo struct {
	Grow      float64 `json:"grow"`
	UseDay    float64 `json:"useDay"`
	DiskFree  float64 `json:"diskFree"`
	DiskUse   float64 `json:"diskUse"`
	DiskQuota float64 `json:"diskQuota"`
}

type DbListResult struct {
	Result DbInfo `json:"result"`
}

type DbInfo struct {
	DbInfoItems []DbInfoItem `json:"dbList"`
}

type DbInfoItem struct {
	Size        string `json:"size"`
	TableSchema string `json:"tableSchema"`
}

type GetTableListArgs struct {
	DbName    string `json:"dbName"`
	PageNo    int    `json:"pageNo,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
	OrderBy   string `json:"orderBy,omitempty"`
	Sort      string `json:"sort,omitempty"`
	SearchKey string `json:"searchKey,omitempty"`
}

type TableListResult struct {
	Result TableInfo `json:"result"`
}
type TableInfo struct {
	Data       []TableInfoItem `json:"data"`
	TotalCount int64           `json:"totalCount"`
}

type TableInfoItem struct {
	IndexLength    string `json:"indexLength"`
	TableSchema    string `json:"tableSchema"`
	TableName      string `json:"tableName"`
	DataLength     string `json:"dataLength"`
	Engine         string `json:"engine"`
	TableRows      string `json:"tableRows"`
	DataFree       string `json:"dataFree"`
	DataFreePer    string `json:"dataFreePer"`
	AvgRowLength   string `json:"avgRowLength"`
	TableLength    string `json:"tableLength"`
	TableLengthPer string `json:"tableLengthPer"`
}

type KillSessionTypesResult struct {
	Types Commands `json:"types"`
}
type Commands struct {
	Command []string `json:"command"`
}

type SessionSummaryResult struct {
	ActiveTotalCount   int64   `json:"activeTotalCount"`
	CpuUtilizationRate float64 `json:"cpuUtilizationRate"`
	MaxExecuteTime     float64 `json:"maxExecuteTime"`
	TotalCount         int64   `json:"totalCount"`
}

type SessionDetailArgs struct {
	Page        int    `json:"page,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	Sort        string `json:"sort,omitempty"`
	ExecuteTime int64  `json:"executeTime,omitempty"`
	Operator    string `json:"operator,omitempty"`
	IsActive    bool   `json:"isActive,omitempty"`
	SessionId   string `json:"sessionId,omitempty"`
	User        string `json:"user,omitempty"`
	Host        string `json:"host,omitempty"`
	Db          string `json:"db,omitempty"`
	Command     string `json:"command,omitempty"`
	State       string `json:"state,omitempty"`
	SqlStmt     string `json:"sqlStmt,omitempty"`
}

type SessionDetailResult struct {
	Items      []SessionDetailItem `json:"items"`
	TotalCount int64               `json:"totalCount"`
}

type SessionDetailItem struct {
	Command string `json:"command"`
	Db      string `json:"db"`
	Host    string `json:"host"`
	Id      int64  `json:"id"`
	Sqlstmt string `json:"sqlStmt"`
	State   string `json:"state"`
	Time    string `json:"time"`
	User    string `json:"user"`
}

type KillSessionAuthArgs struct {
	DbHost     string   `json:"dbHost,omitempty"`
	DbPort     int      `json:"dbPort,omitempty"`
	Items      []string `json:"items,omitempty"`
	DbUser     string   `json:"dbUser,omitempty"`
	DbPassword string   `json:"dbPassword,omitempty"`
}

type KillSessionAuthResult struct {
	Success bool `json:"success"`
}

type KillSessionHistory struct {
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
}

type KillSessionHistoryResult struct {
	Items      []KillSessionHistoryItem `json:"items"`
	TotalCount int64                    `json:"totalCount"`
}

type KillSessionHistoryItem struct {
	OperateDbUser      string `json:"operateDbUser"`
	OperateTime        string `json:"operateTime"`
	OperateUser        string `json:"operateUser"`
	SessionCommand     string `json:"sessionCommand"`
	SessionDb          string `json:"sessionDb"`
	SessionExecuteTime string `json:"sessionExecuteTime"`
	SessionHost        string `json:"sessionHost"`
	SessionId          int64  `json:"sessionId"`
	SessionSql         string `json:"sessionSql"`
	SessionState       string `json:"sessionState"`
	SessionUser        string `json:"sessionUser"`
	Status             int    `json:"status"`
	StatusDesc         string `json:"statusDesc"`
	StatusInfo         string `json:"statusInfo"`
}

type KillSessionArgs struct {
	DbHost string   `json:"dbHost,omitempty"`
	DbPort int      `json:"dbPort,omitempty"`
	Items  []string `json:"items,omitempty"`
}

type KillSessionResult struct {
	Success bool `json:"success"`
}

type SessionStatisticsResult map[string][]SessionStatisticsItem

type SessionStatisticsItem struct {
	ActiveAverageExecuteTime int64  `json:"activeAverageExecuteTime"`
	ActiveTotalCount         int64  `json:"activeTotalCount"`
	ActiveTotalExecuteTime   int64  `json:"activeTotalExecuteTime"`
	DimensionKey             string `json:"dimensionKey"`
	DimensionValue           string `json:"dimensionValue"`
	TotalCount               int64  `json:"totalCount"`
}

type ErrorLogStatusResult struct {
	Enabled int `json:"enabled"`
}

type ErrorLogResult struct {
	Success bool `json:"success"`
}

type ErrorLogListArgs struct {
	Page     int    `json:"page,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
	Label    string `json:"label,omitempty"`
}
type ErrorLogListResult struct {
	Items      []ErrorLogItem `json:"items"`
	TotalCount int64          `json:"totalCount"`
}

type ErrorLogItem struct {
	ConnectionId string `json:"connectionId"`
	ErrCode      string `json:"errCode"`
	ErrInfo      string `json:"errInfo"`
	Label        string `json:"label"`
	LogId        string `json:"logId"`
	Subsystem    string `json:"subsystem"`
	Time         string `json:"time"`
}

type SqlFilterListResult struct {
	SqlFilterList []SqlFilterItem `json:"sqlFilterList"`
}
type SqlFilterItem struct {
	Id           int64  `json:"id"`
	FilterType   string `json:"filterType"`
	FilterKey    string `json:"filterKey"`
	FilterLimit  int64  `json:"filterLimit"`
	FilterStatus string `json:"filterStatus"`
	CreateTime   string `json:"createTime"`
	UpdateTime   string `json:"updateTime"`
}

type SqlFilterArgs struct {
	FilterType  string `json:"filterType"`
	FilterKey   string `json:"filterKey"`
	FilterLimit int64  `json:"filterLimit"`
}

type StartOrStopSqlFilterArgs struct {
	Action string `json:"action"`
}

type IsAllowedResult struct {
	Allowed bool `json:"allowed"`
}

type ProcessArgs struct {
	Ids []int64 `json:"ids,omitempty"`
}

type InnodbStatusResult struct {
	Datatime string `json:"datatime"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Status   string `json:"status"`
}

type ProcessListResult struct {
	Datetime    string        `json:"datetime"`
	ProcessList []ProcessItem `json:"processList"`
}

type LockHold map[string][]int64

type LockWait struct {
	LocakId string `json:"lockId"`
	Id      int64  `json:"id"`
}
type ProcessItem struct {
	Sql      string   `json:"sql"`
	Db       string   `json:"db"`
	State    string   `json:"state"`
	Host     string   `json:"host"`
	Command  string   `json:"command"`
	User     string   `json:"user"`
	Time     int64    `json:"time"`
	ID       int64    `json:"id"`
	LockHold LockHold `json:"lockHold"`
	LockWait []LockWait
}

type TransactionListResult struct {
	Datetime      string          `json:"datetime"`
	InnodbTrxList []InnodbTrxItem `json:"innodbTrxList"`
}
type InnodbTrxItem struct {
	TrxRequestedLockId string `json:"trxRequestedLockId"`
	TrxStarted         string `json:"trxStarted"`
	TrxMysqlThreadId   int64  `json:"trxMysqlThreadId"`
	TrxRowsLocked      int64  `json:"trxRowsLocked"`
	TrxWaitStarted     string `json:"trxWaitStarted"`
	TrxState           string `json:"trxState"`
	TrxTablesInUse     int64  `json:"trxTablesInUse"`
	TrxId              string `json:"trxId"`
	TrxQuery           string `json:"trxQuery"`
	TrxTablesLocked    int    `json:"trxTablesLocked"`
}

type ConnectionListResult struct {
	Datetime    string           `json:"datetime"`
	ConnectList []ConnectionItem `json:"connectList"`
}
type ConnectionItem struct {
	LocalAddress   string `json:"localAddress"`
	Proto          string `json:"proto"`
	SendQ          int    `json:"sendQ"`
	ForeignAddress string `json:"foreignAddress"`
	RecvQ          int    `json:"recvQ"`
}

type FailInjectWhiteListResult struct {
	AppList []string `json:"appList"`
}

type FailInjectArgs struct {
	AppList []string `json:"appList"`
}

type TaskResult struct {
	TaskId int64 `json:"taskId"`
}

type OrderStatusResult struct {
	Status string `json:"status"`
}

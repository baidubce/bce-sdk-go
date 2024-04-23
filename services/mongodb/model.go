/*
 * Copyright 2024 Baidu, Inc.
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

package mongodb

import "time"

type MemberModel struct {
	MemberId    string `json:"memberId,omitempty"`    // 节点IP
	MongoStatus string `json:"mongoStatus,omitempty"` // 节点的mongo状态
	MongoType   string `json:"mongoType,omitempty"`   // 节点的mongo类型
	Readonly    string `json:"readonly,omitempty"`    // 节点是否是只读节点
	ZoneName    string `json:"zoneName,omitempty"`    // 子网所属可用区
	Ip          string `json:"ip,omitempty"`          // 节点内网IP
	MongoHost   string `json:"mongoHost,omitempty"`   // 节点所绑定域名
	Eip         string `json:"eip,omitempty"`         // 节点公网IP
}

type DbInstanceSimpleModel struct {
	DbInstanceId     string `json:"dbInstanceId,omitempty"`     // 实例短ID
	ConnectionString string `json:"connectionString,omitempty"` // 数据库链接字符串
	Port             string `json:"port,omitempty"`             // 数据库连接端口
}

type TagModel struct {
	TagKey   string `json:"tagKey,omitempty"`   // 标签键
	TagValue string `json:"tagValue,omitempty"` // 标签值
}

type NodeModel struct {
	NodeId         string        `json:"nodeId,omitempty"`         // 组件ID
	Name           string        `json:"name,omitempty"`           // 组件名称
	Status         string        `json:"status,omitempty"`         // 组件状态
	CpuCount       int           `json:"cpuCount,omitempty"`       // 组件CPU规格
	MemoryCapacity int           `json:"memoryCapacity,omitempty"` // 组件内存规格
	Storage        int           `json:"storage,omitempty"`        // 组件存储规格
	ConnectString  string        `json:"connectString,omitempty"`  // 组件链接字符串
	StorageType    string        `json:"storageType,omitempty"`    // 组件存储类型
	Members        []MemberModel `json:"members,omitempty"`        // 组件内节点信息
}

type SubnetMap struct {
	SubnetId string `json:"subnetId,omitempty"` // 子网短ID
	ZoneName string `json:"zoneName,omitempty"` // 子网所属可用区
}

type LogServiceModel struct {
	Type   string `json:"type,omitempty"`   // 日志类型
	Status string `json:"status,omitempty"` // 日志当前状态
}

type ResourceGroupModel struct {
	AccountId  string `json:"accountId,omitempty"`
	BindTime   string `json:"bindTime,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	DeleteTime string `json:"deleteTime,omitempty"`
	Extra      string `json:"extra,omitempty"`
	GroupId    string `json:"groupId,omitempty"`
	Name       string `json:"name,omitempty"`
	ParentUuid string `json:"parentUuid,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
	UserId     string `json:"userId,omitempty"`
}

type InstanceModel struct {
	DbInstanceId             string               `json:"dbInstanceId,omitempty"`
	DbInstanceName           string               `json:"dbInstanceName,omitempty"`
	DbInstanceUUID           string               `json:"dbInstanceUUID,omitempty"`
	ConnectionString         string               `json:"connectionString,omitempty"`
	Port                     string               `json:"port,omitempty"`
	StorageEngine            string               `json:"storageEngine,omitempty"`
	EngineVersion            string               `json:"engineVersion,omitempty"`
	ResourceGroups           []ResourceGroupModel `json:"resourceGroups,omitempty"`
	ResourceUuid             string               `json:"resourceUuid,omitempty"`
	DbInstanceType           string               `json:"dbInstanceType,omitempty"`
	DbInstanceStatus         string               `json:"dbInstanceStatus,omitempty"`
	DbInstanceCpuCount       int                  `json:"dbInstanceCpuCount,omitempty"`
	DbInstanceMemoryCapacity int                  `json:"dbInstanceMemoryCapacity,omitempty"`
	DbInstanceStorage        int                  `json:"dbInstanceStorage,omitempty"`
	MongosCount              int                  `json:"mongosCount,omitempty"`
	ShardCount               int                  `json:"shardCount,omitempty"`
	CreateTime               time.Time            `json:"createTime,omitempty"`
	VotingMemberNum          int                  `json:"votingMemberNum,omitempty"`
	ReadonlyNodeNum          int                  `json:"readonlyNodeNum,omitempty"`
	VpcId                    string               `json:"vpcId,omitempty"`
	PaymentTiming            string               `json:"paymentTiming,omitempty"`
	Subnets                  []SubnetMap          `json:"subnets,omitempty"`
	Tags                     []TagModel           `json:"tags,omitempty"`
}

type ListMongodbArgs struct {
	Marker         string
	MaxKeys        int
	EngineVersion  string
	StorageEngine  string
	DbInstanceType string
}

type ListMongodbResult struct {
	Marker      string          `json:"marker,omitempty"`
	MaxKeys     int             `json:"maxKeys,omitempty"`
	IsTruncated bool            `json:"isTruncated,omitempty"`
	NextMarker  string          `json:"nextMarker,omitempty"`
	DbInstances []InstanceModel `json:"dbInstances,omitempty"`
}

type InstanceDetail struct {
	DbInstanceId                   string               `json:"dbInstanceId,omitempty"`
	DbInstanceName                 string               `json:"dbInstanceName,omitempty"`
	DbInstanceUUID                 string               `json:"dbInstanceUUID,omitempty"`
	ConnectionString               string               `json:"connectionString,omitempty"`
	Port                           string               `json:"port,omitempty"`
	EngineVersion                  string               `json:"engineVersion,omitempty"`
	StorageEngine                  string               `json:"storageEngine,omitempty"`
	ResourceGroups                 []ResourceGroupModel `json:"resourceGroups,omitempty"`
	ResourceUuid                   string               `json:"resourceUuid,omitempty"`
	DbInstanceCpuCount             int                  `json:"dbInstanceCpuCount,omitempty"`
	DbInstanceMemoryCapacity       int                  `json:"dbInstanceMemoryCapacity,omitempty"`
	DbInstanceStorage              int                  `json:"dbInstanceStorage,omitempty"`
	DbInstanceStorageType          string               `json:"dbInstanceStorageType,omitempty"`
	DbInstanceType                 string               `json:"dbInstanceType,omitempty"`
	MongosCount                    int                  `json:"mongosCount,omitempty"`
	ShardCount                     int                  `json:"shardCount,omitempty"`
	MongosList                     []NodeModel          `json:"mongosList,omitempty"`
	ShardList                      []NodeModel          `json:"shardList,omitempty"`
	VotingMemberNum                int                  `json:"votingMemberNum,omitempty"`
	ReadonlyNodeNum                int                  `json:"readonlyNodeNum,omitempty"`
	DbInstanceStatus               string               `json:"dbInstanceStatus,omitempty"`
	CreateTime                     time.Time            `json:"createTime,omitempty"`
	ExpireTime                     time.Time            `json:"expireTime,omitempty"`
	VpcId                          string               `json:"vpcId,omitempty"`
	PublicConnectionString         string               `json:"publicConnectionString,omitempty"`
	PublicReadonlyConnectionString string               `json:"publicReadonlyConnectionString,omitempty"`
	ReadOnlyNodeConnectionString   string               `json:"readOnlyNodeConnectionString,omitempty"`
	PaymentTiming                  string               `json:"paymentTiming,omitempty"`
	LogServiceStatus               []LogServiceModel    `json:"logServiceStatus,omitempty"`
	Members                        []MemberModel        `json:"members,omitempty"`
	Subnets                        []SubnetMap          `json:"subnets,omitempty"`
	Tags                           []TagModel           `json:"tags,omitempty"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`   // 时长
	ReservationTimeUnit string `json:"reservationTimeUnit"` // 时间单位，Month
}

type BillingModel struct {
	PaymentTiming string      `json:"paymentTiming"` // 付费方式
	Reservation   Reservation `json:"reservation"`   // 保留信息
}

type CreateReplicaArgs struct {
	ClientToken              string       `json:"-"`
	Billing                  BillingModel `json:"billing,omitempty"`
	PurchaseCount            int          `json:"purchaseCount,omitempty"`            // 批量创建实例个数
	DbInstanceName           string       `json:"dbInstanceName,omitempty"`           // 用户自定义实例名
	StorageEngine            string       `json:"storageEngine,omitempty"`            // 存储引擎
	EngineVersion            string       `json:"engineVersion,omitempty"`            // 数据库版本
	DbInstanceType           string       `json:"dbInstanceType,omitempty"`           // 实例类型
	DbInstanceCpuCount       int          `json:"dbInstanceCpuCount,omitempty"`       // 副本集实例CPU规格
	DbInstanceMemoryCapacity int          `json:"dbInstanceMemoryCapacity,omitempty"` // 副本集实例内存规格，单位GB
	DbInstanceStorage        int          `json:"dbInstanceStorage,omitempty"`        // 副本集实例存储规格，单位GB
	AccountPassword          string       `json:"accountPassword,omitempty"`          // root账号的密码
	DbInstanceStorageType    string       `json:"dbInstanceStorageType,omitempty"`    // 实例存储类型
	VotingMemberNum          int          `json:"votingMemberNum,omitempty"`          // 副本集实例投票节点数量
	ReadonlyNodeNum          int          `json:"readonlyNodeNum,omitempty"`          // 副本集实例只读节点数量
	SrcDbInstanceId          string       `json:"srcDbInstanceId,omitempty"`          // 源实例ID
	BackupId                 string       `json:"backupId,omitempty"`                 // 备份ID
	VpcId                    string       `json:"vpcId,omitempty"`                    // VPC ID
	Subnets                  []SubnetMap  `json:"subnets,omitempty"`                  // 子网列表
	RestoreTime              string       `json:"restoreTime,omitempty"`              // 恢复时间点
	Tags                     []TagModel   `json:"tags,omitempty"`                     // 标签列表
	ResGroupId               string       `json:"resGroupId,omitempty"`               // 资源组ID
}

type CreateShardingArgs struct {
	ClientToken          string       `json:"-"`
	Billing              BillingModel `json:"billing,omitempty"`
	PurchaseCount        int          `json:"purchaseCount,omitempty"`
	DbInstanceName       string       `json:"dbInstanceName,omitempty"`
	StorageEngine        string       `json:"storageEngine,omitempty"`
	EngineVersion        string       `json:"engineVersion,omitempty"`
	DbInstanceType       string       `json:"dbInstanceType,omitempty"`
	MongosCount          int          `json:"mongosCount,omitempty"`
	MongosCpuCount       int          `json:"mongosCpuCount,omitempty"`
	MongosMemoryCapacity int          `json:"mongosMemoryCapacity,omitempty"`
	ShardCount           int          `json:"shardCount,omitempty"`
	ShardCpuCount        int          `json:"shardCpuCount,omitempty"`
	ShardMemoryCapacity  int          `json:"shardMemoryCapacity,omitempty"`
	ShardStorage         int          `json:"shardStorage,omitempty"`
	ShardStorageType     string       `json:"shardStorageType,omitempty"`
	AccountPassword      string       `json:"accountPassword,omitempty"`
	SrcDbInstanceId      string       `json:"srcDbInstanceId,omitempty"`
	BackupId             string       `json:"backupId,omitempty"`
	VpcId                string       `json:"vpcId,omitempty"`
	Subnets              []SubnetMap  `json:"subnets,omitempty"`
	RestoreTime          string       `json:"restoreTime,omitempty"`
	Tags                 []TagModel   `json:"tags,omitempty"`
	ResGroupId           string       `json:"resGroupId,omitempty"`
}

type CreateResult struct {
	DbInstanceSimpleModels []DbInstanceSimpleModel `json:"dbInstanceSimpleModels"`
}

type ShardingAddComponentArgs struct {
	ClientToken        string `json:"-"`
	DbInstanceId       string `json:"dbInstanceId,omitempty"`
	NodeCpuCount       int    `json:"nodeCpuCount,omitempty"`
	NodeMemoryCapacity int    `json:"nodeMemoryCapacity,omitempty"`
	NodeStorage        int    `json:"nodeStorage,omitempty"`
	NodeType           string `json:"nodeType,omitempty"`
	PurchaseCount      int    `json:"purchaseCount,omitempty"`
}

type ShardingAddComponentResult struct {
	NodeIds []string `json:"nodeIds"`
}

type ReplicaResizeArgs struct {
	ClientToken              string `json:"-"`
	DbInstanceCpuCount       int    `json:"dbInstanceCpuCount,omitempty"`
	DbInstanceMemoryCapacity int    `json:"dbInstanceMemoryCapacity,omitempty"`
	DbInstanceStorage        int    `json:"dbInstanceStorage,omitempty"`
}

type ShardingComponentResizeArgs struct {
	ClientToken        string `json:"-"`
	DbInstanceId       string `json:"dbInstanceId,omitempty"`
	NodeId             string `json:"nodeId,omitempty"`
	NodeCpuCount       int    `json:"nodeCpuCount,omitempty"`
	NodeMemoryCapacity int    `json:"nodeMemoryCapacity,omitempty"`
	NodeStorage        int    `json:"nodeStorage,omitempty"`
}

type RestartMongodbsArgs struct {
	DbInstanceIds []string `json:"dbInstanceIds,omitempty"`
}

type UpdateInstanceNameArgs struct {
	DbInstanceName string `json:"dbInstanceName"`
}

type UpdateComponentNameArgs struct {
	NodeName string `json:"nodeName"`
}

type MemberRoleModel struct {
	SubnetId string `json:"subnetId,omitempty"`
	Role     string `json:"role,omitempty"`
}

type MigrateAzoneArgs struct {
	Subnets []SubnetMap       `json:"subnets,omitempty"`
	Members []MemberRoleModel `json:"members,omitempty"`
}

type LogicAssignResource struct {
	ResourceId  string     `json:"resourceId,omitempty"`  // 资源短id
	ServiceType string     `json:"serviceType,omitempty"` // 资源类型
	Tags        []TagModel `json:"tags,omitempty"`        // 标签
}

type UpdateTagArgs struct {
	DbInstanceId string     `json:"dbInstanceId,omitempty"`
	Tags         []TagModel `json:"tags,omitempty"` // 标签
}

type AssignTagArgs struct {
	Resources []LogicAssignResource `json:"resources,omitempty"` // 资源列表
}

type BackupModel struct {
	BackupId          string    `json:"backupId,omitempty"`          // 备份ID
	BackupSize        string    `json:"backupSize,omitempty"`        // 备份大小。单位Byte
	BackupMethod      string    `json:"backupMethod,omitempty"`      // 备份方式，取值参考
	BackupMode        string    `json:"backupMode,omitempty"`        // 备份模式，取值参考
	BackupType        string    `json:"backupType,omitempty"`        // 备份类型，取值参考
	BackupStatus      string    `json:"backupStatus,omitempty"`      // 备份状态，取值参考
	BackupStartTime   time.Time `json:"backupStartTime,omitempty"`   // 备份开始时间
	BackupEndTime     time.Time `json:"backupEndTime,omitempty"`     // 备份结束时间
	BackupDescription string    `json:"backupDescription,omitempty"` // 备份详情
}

type CreateBackupArgs struct {
	BackupMethod      string `json:"backupMethod,omitempty"`
	BackupDescription string `json:"backupDescription,omitempty"`
}

type CreateBackupResult struct {
	BackupId string `json:"backupId,omitempty"`
}

type ListBackupArgs struct {
	Marker  string `json:"marker,omitempty"`
	MaxKeys int    `json:"maxKeys,omitempty"`
}

type ListBackupResult struct {
	Marker      string        `json:"marker,omitempty"`
	MaxKeys     int           `json:"maxKeys,omitempty"`
	IsTruncated bool          `json:"isTruncated,omitempty"`
	NextMarker  string        `json:"nextMarker,omitempty"`
	Backups     []BackupModel `json:"backups,omitempty"`
}

type BackupDetail struct {
	BackupId          string    `json:"backupId,omitempty"`
	BackupSize        string    `json:"backupSize,omitempty"`
	BackupMethod      string    `json:"backupMethod,omitempty"`
	BackupMode        string    `json:"backupMode,omitempty"`
	BackupType        string    `json:"backupType,omitempty"`
	BackupStatus      string    `json:"backupStatus,omitempty"`
	BackupStartTime   time.Time `json:"backupStartTime,omitempty"`
	BackupEndTime     time.Time `json:"backupEndTime,omitempty"`
	DownloadURL       string    `json:"downloadUrl,omitempty"`
	DownloadExpires   string    `json:"downloadExpires,omitempty"`
	BackupDescription string    `json:"backupDescription,omitempty"`
}

type ModifyBackupDescriptionArgs struct {
	BackupDescription string `json:"backupDescription,omitempty"`
}

type BackupPolicy struct {
	AutoBackupEnable          string `json:"autoBackupEnable,omitempty"`
	PreferredBackupPeriod     string `json:"preferredBackupPeriod,omitempty"`
	PreferredBackupTime       string `json:"preferredBackupTime,omitempty"`
	BackupRetentionPeriod     int    `json:"backupRetentionPeriod,omitempty"`
	EnableIncrementBackup     int    `json:"enableIncrementBackup,omitempty"`
	BackupMethod              string `json:"backupMethod,omitempty"`
	IncrBackupRetentionPeriod int    `json:"incrBackupRetentionPeriod,omitempty"`
}

type SecurityIpModel struct {
	SecurityIps []string `json:"securityIps,omitempty"`
}

type StartLoggingArgs struct {
	Type string `json:"type,omitempty"`
}

type ListLogFilesArgs struct {
	MemberId  string `json:"memberId,omitempty"`
	Type      string `json:"type,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
}

type ListLogFilesResult struct {
	Logs []LogFile `json:"logs,omitempty"`
}

type LogFile struct {
	Name               string `json:"name,omitempty"`
	StartTime          string `json:"startTime,omitempty"`
	EndTime            string `json:"endTime,omitempty"`
	Size               int    `json:"size,omitempty"`
	DownloadUrl        string `json:"downloadUrl,omitempty"`
	DownloadExpires    int    `json:"downloadExpires,omitempty"`
	DownloadExpireTime string `json:"downloadExpireTime,omitempty"`
}

type UpdatePasswordArgs struct {
	AccountPassword string `json:"accountPassword"`
}

type ReplicaAddReadonlyNodesArgs struct {
	ClientToken     string    `json:"-"`
	DbInstanceId    string    `json:"dbInstanceId,omitempty"`
	ReadonlyNodeNum int       `json:"readonlyNodeNum,omitempty"`
	Subnet          SubnetMap `json:"subnet,omitempty"`
}

type ReplicaAddReadonlyNodesResult struct {
	OrderId           string   `json:"orderId,omitempty"`
	ReadonlyMemberIds []string `json:"readonlyMemberIds,omitempty"`
}

type ReadonlyNodesList struct {
	CompId  string   `json:"compId,omitempty"`
	NodeIds []string `json:"nodeIds,omitempty"`
}

type GetReadonlyNodesResult struct {
	ReadOnlyList []ReadonlyNodesList `json:"readOnlyList,omitempty"`
}

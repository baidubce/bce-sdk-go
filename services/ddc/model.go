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

type CreateInstanceArgs struct {
	ClientToken  string   `json:"-"`
	InstanceType string   `json:"instanceType"`
	Number       int      `json:"number"`
	Instance     Instance `json:"instance"`
}

type CreateResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type InstanceModelResult struct {
	Instance InstanceModel `json:"instance"`
}

type Instance struct {
	InstanceName         string `json:"instanceName"`
	SourceInstanceId     string `json:"sourceInstanceId"`
	Engine               string `json:"engine"`
	EngineVersion        string `json:"engineVersion"`
	CpuCount             int    `json:"cpuCount"`
	AllocatedMemoryInGB  int    `json:"allocatedMemoryInGB"`
	AllocatedStorageInGB int    `json:"allocatedStorageInGB"`
	AZone                string `json:"azone"`
	VpcId                string `json:"vpcId"`
	SubnetId             string `json:"subnetId"`
	DiskIoType           string `json:"diskIoType"`
	DeployId             string `json:"deployId"`
	PoolId               string `json:"poolId"`
	RoGroupId            string `json:"roGroupId"`
	EnableDelayOff       string `json:"enableDelayOff"`
	DelayThreshold       int    `json:"delayThreshold"`
	LeastInstanceAmount  int    `json:"leastInstanceAmount"`
	RoGroupWeight        int    `json:"roGroupWeight"`
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
	CreateTime string   `json:"createTime"`
	DeployID   string   `json:"deployId"`
	DeployName string   `json:"deployName"`
	Instances  []string `json:"instances"`
	PoolID     string   `json:"poolId"`
	Strategy   string   `json:"strategy"`
}

type CreateDeployRequest struct {
	ClientToken string `json:"-"`
	DeployName  string `json:"deployName"`
	Strategy    string `json:"strategy"`
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
	PublicAccessStatus   string       `json:"publicAccessStatus"`
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
	Subnets              []SubnetVo   `json:"subnets"`
	NodeMaster           NodeInfo     `json:"nodeMaster"`
	NodeSlave            NodeInfo     `json:"nodeSlave"`
	NodeReadReplica      NodeInfo     `json:"nodeReadReplica"`
	DeployId             string       `json:"deployId"`
	Topology             Topology     `json:"topology"`
	DiskType             string       `json:"diskType"`
	Type                 string       `json:"type"`
	ApplicationType      string       `json:"applicationType"`
	RoGroupList          []RoGroup    `json:"roGroupList"`
}

type SubnetVo struct {
	Name     string `json:"name"`
	SubnetId string `json:"subnetId"`
	Az       string `json:"az"`
	Cidr     string `json:"cidr"`
	ShortId  string `json:"shortId"`
}

type RoGroup struct {
	RoGroupId   string    `json:"roGroupId"`
	VnetIp      string    `json:"vnetIp"`
	ReplicaList []Replica `json:"replicaList"`
}

type Replica struct {
	InstanceId    string `json:"instanceId"`
	Status        string `json:"status"`
	RoGroupWeight int    `json:"roGroupWeight"`
}

type NodeInfo struct {
	Id       string `json:"id"`
	Azone    string `json:"azone"`
	SubnetId string `json:"subnetId"`
	Cidr     string `json:"cidr"`
	Name     string `json:"name"`
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

type DeleteDdcArgs struct {
	InstanceIds []string `json:"instanceIds"`
}

type UpdateInstanceNameArgs struct {
	InstanceName string `json:"instanceName"`
}

type ListDdcResult struct {
	Marker      string          `json:"marker"`
	MaxKeys     int             `json:"maxKeys"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	Result      []InstanceModel `json:"result"`
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

type GetSecurityIpsResult struct {
	SecurityIps []string `json:"ip"`
}

type UpdateSecurityIpsArgs struct {
	InstanceId  string   `json:"instanceId"`
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
	ClientToken        string              `json:"-"`
	AccountName        string              `json:"accountName"`
	Password           string              `json:"password"`
	Type               AccountType         `json:"type"`
	Remark             string              `json:"remark"`
	DatabasePrivileges []DatabasePrivilege `json:"databasePrivileges,omitempty"`
}

type DatabasePrivilege struct {
	DbName   string   `json:"dbName"`
	AuthType AuthType `json:"authType"`
}

type Account struct {
	AccountName        string              `json:"accountName"`
	Remark             string              `json:"remark"`
	AccountStatus      string              `json:"accountStatus"`
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

type UpdateAccountRemarkArgs struct {
	Remark string `json:"remark"`
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

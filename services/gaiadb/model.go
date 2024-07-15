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

package gaiadb

type AvailableSubnet struct {
	VpcId    string `json:"vpcId"`
	SubnetId string `json:"subnetId"`
}

type CreateClusterArgs struct {
	ClientToken       string        `json:"-"`
	ProductType       string        `json:"productType"`
	Duration          string        `json:"duration,omitempty"`
	AutoRenewTimeUnit string        `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int           `json:"autoRenewTime,omitempty"`
	Number            int           `json:"number"`
	InstanceParam     InstanceParam `json:"instanceParam"`
}
type InstanceParam struct {
	ReleaseVersion       string `json:"releaseVersion"`
	AllocatedCpuInCore   int    `json:"allocatedCpuInCore"`
	AllocatedMemoryInMB  int    `json:"allocatedMemoryInMB"`
	AllocatedStorageInGB int    `json:"allocatedStorageInGB"`
	InstanceAmount       int    `json:"instanceAmount"`
	ProxyAmount          int    `json:"proxyAmount"`
	VpcId                string `json:"vpcId"`
	SubnetId             string `json:"subnetId"`
	SrcClusterId         string `json:"srcClusterId,omitempty"`
	SnapshotId           string `json:"snapshotId,omitempty"`
	Pit                  string `json:"pit,omitempty"`
	LowerCaseTableNames  string `json:"lowercaseTableNames,omitempty"`
	ComputeTplId         string `json:"computeTplId,omitempty"`
}
type CreateResult struct {
	OrderId    string   `json:"orderId"`
	ClusterIds []string `json:"clusterIds"`
}

type ClusterName struct {
	ClusterName string `json:"clusterName"`
}

type ResizeClusterArgs struct {
	ResizeType          string     `json:"resizeType"`
	SlaveId             string     `json:"slaveId,omitempty"`
	AllocatedCpuInCore  int        `json:"allocatedCpuInCore,omitempty"`
	AllocatedMemoryInMB int        `json:"allocatedMemoryInMB,omitempty"`
	ProxyAmount         int        `json:"proxyAmount,omitempty"`
	InterfaceId         int        `json:"interfaceId,omitempty"`
	Interfacee          Interfacee `json:"interfacee"`
}
type Interfacee struct {
	AddressName     string   `json:"addressName,omitempty"`
	InstanceBinding []string `json:"instanceBinding,omitempty"`
	ProxyAmount     int      `json:"proxyAmount,omitempty"`
	ReadWriteMode   string   `json:"readWriteMode,omitempty"`
	MasterReadable  int      `json:"masterReadable,omitempty"`
}

type OrderId struct {
	OrderId string `json:"orderId"`
}

type Marker struct {
	Marker  string
	MaxKeys int
}

type ClusterListResult struct {
	Marker      string    `json:"marker"`
	MaxKeys     int       `json:"maxKeys"`
	NextMarker  string    `json:"nextMarker"`
	Istruncated bool      `json:"istruncated"`
	Clusters    []Cluster `json:"clusters"`
}

type Cluster struct {
	ClusterId            string   `json:"clusterId"`
	ClusterName          string   `json:"clusterName"`
	Endpoint             Endpoint `json:"endpoint"`
	AllocatedCpuInCore   int      `json:"allocatedCpuInCore"`
	AllocatedMemoryInMB  int      `json:"allocatedMemoryInMB"`
	InstanceStatus       string   `json:"instanceStatus"`
	PubliclyAccessible   bool     `json:"publiclyAccessible"`
	InstanceCreateTime   string   `json:"instanceCreateTime"`
	InstanceExpireTime   string   `json:"instanceExpireTime"`
	InstanceAmount       int      `json:"instanceAmount"`
	ProxyAmount          int      `json:"proxyAmount"`
	ProductType          string   `json:"productType"`
	MultiActiveGroupId   string   `json:"multiActiveGroupId"`
	MultiActiveGroupRole string   `json:"multiActiveGroupRole"`
	MultiActiveGroupName string   `json:"multiActiveGroupName"`
}

type Endpoint struct {
	Port    int    `json:"port"`
	Address string `json:"address"`
	VnetIp  string `json:"vnetIp,omitempty"`
	InetIp  string `json:"inetIp,omitempty"`
}

type ClusterDetailResult struct {
	ClusterId            string        `json:"clusterId"`
	ClusterName          string        `json:"clusterName"`
	Endpoint             Endpoint      `json:"endpoint"`
	AllocatedCpuInCore   int           `json:"allocatedCpuInCore"`
	AllocatedMemoryInMB  int           `json:"allocatedMemoryInMB"`
	InstanceStatus       string        `json:"instanceStatus"`
	PubliclyAccessible   bool          `json:"publiclyAccessible"`
	InstanceCreateTime   string        `json:"instanceCreateTime"`
	InstanceExpireTime   string        `json:"instanceExpireTime"`
	InstanceAmount       int           `json:"instanceAmount"`
	ProductType          string        `json:"productType"`
	MultiActiveGroupId   string        `json:"multiActiveGroupId"`
	MultiActiveGroupRole string        `json:"multiActiveGroupRole"`
	MultiActiveGroupName string        `json:"multiActiveGroupName"`
	ComputeList          []ComputeNode `json:"computeList"`
}
type ComputeNode struct {
	Role             string `json:"role"`
	InstanceId       string `json:"instanceId"`
	InstanceShortId  string `json:"instanceShortId"`
	InstanceUniqueId string `json:"instanceUniqueId"`
	Status           string `json:"status"`
}

type ClusterCapacityResult struct {
	TotalFree     int `json:"totalFree"`
	TotalCapacity int `json:"totalCapacity"`
}

type QueryPriceArgs struct {
	Number        int          `json:"number"`
	InstanceParam InstanceInfo `json:"instanceParam"`
	ProductType   string       `json:"productType"`
	Duration      int          `json:"duration"`
}

type InstanceInfo struct {
	ReleaseVersion       string `json:"releaseVersion"`
	AllocatedStorageInGB int    `json:"allocatedStorageInGB"`
	AllocatedMemoryInMB  int    `json:"allocatedMemoryInMB"`
	AllocatedCpuInCore   int    `json:"allocatedCpuInCore"`
	InstanceAmount       int    `json:"instanceAmount"`
	ProxyAmount          int    `json:"proxyAmount"`
}

type PriceResult struct {
	Price        float32 `json:"price"`
	CatalogPrice float32 `json:"catalogPrice,omitempty"`
}

type QueryResizePriceArgs struct {
	ClusterId           string     `json:"clusterId"`
	ResizeType          string     `json:"resizeType"`
	SlaveId             string     `json:"slaveId"`
	AllocatedCpuInCore  int        `json:"allocatedCpuInCore"`
	AllocatedMemoryInMB int        `json:"allocatedMemoryInMB"`
	ProxyAmount         int        `json:"proxyAmount"`
	InterfaceId         string     `json:"interfaceId"`
	Interfacee          Interfacee `json:"interfacee"`
}

type RebootInstanceArgs struct {
	ExecuteAction string `json:"executeAction"`
}

type BindTagsArgs struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	ResourceId string `json:"resourceId"`
	Tags       []Tag  `json:"tags"`
}

type Tag struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

type InterfaceListResult struct {
	Interfaces []Interface `json:"interfaces"`
}
type Interface struct {
	AppId               string   `json:"appId"`
	InterfaceId         string   `json:"interfaceId"`
	Status              string   `json:"status"`
	CreateTime          string   `json:"createTime"`
	InstanceBinding     []string `json:"instanceBinding"`
	NewInstanceAutoJoin int      `json:"newInstanceAutoJoin"`
	ReadWriteMode       string   `json:"readWriteMode"`
	InterfaceType       string   `json:"interfaceType"`
	MasterReadable      int      `json:"masterReadable"`
	Access              Access   `json:"access"`
	Flavor              Flavor   `json:"flavor"`
	Proxies             []Proxy  `json:"proxies"`
}
type Proxy struct {
	Id      int    `json:"id"`
	ProxyId string `json:"proxyId"`
}
type Flavor struct {
	RamInMB   int    `json:"ramInMB"`
	DiskInGB  string `json:"diskInGB"`
	CpuInCore int    `json:"cpuInCore"`
}
type Access struct {
	Name        string `json:"name"`
	DnsName     string `json:"dnsName"`
	AddressName string `json:"addressName"`
	Eip         string `json:"eip"`
}

type UpdateDnsNameArgs struct {
	InterfaceId string `json:"instanceId"`
	DnsName     string `json:"dnsName"`
}

type UpdateInterfaceArgs struct {
	InterfaceId string        `json:"instanceId"`
	Interface   InterfaceInfo `json:"interface"`
}
type InterfaceInfo struct {
	MasterReadable  int      `json:"masterReadable"`
	AddressName     string   `json:"addressName"`
	InstanceBinding []string `json:"instanceBinding"`
}

type NewInstanceAutoJoinArgs struct {
	AutoJoinRequestItems []AutoJoinRequestItem `json:"autoJoinRequestItems"`
}
type AutoJoinRequestItem struct {
	NewInstanceAutoJoin string `json:"newInstanceAutoJoin"`
	InterfaceId         string `json:"interfaceId"`
}
type WhiteList struct {
	AuthIps []string `json:"authIps"`
	Etag    string   `json:"etag"`
}

type ClusterSwitchArgs struct {
	ExecuteAction       string `json:"executeAction"`
	SecondaryInstanceId string `json:"secondaryInstanceId"`
}

type ClusterSwitchResult struct {
	Switch TaskId `json:"switch"`
}
type TaskId struct {
	TaskId int `json:"taskId"`
}

type CreateAccountArgs struct {
	AccountName string `json:"accountName"`
	Password    string `json:"password"`
	AccountType string `json:"accountType"`
	Remark      string `json:"remark"`
}

type AccountDetail struct {
	Account AccountInfo `json:"account"`
}
type AccountInfo struct {
	AccountName        string              `json:"accountName"`
	Remark             string              `json:"remark"`
	AccountStatus      string              `json:"accountStatus"`
	AccountType        string              `json:"accountType"`
	DatabasePrivileges []DatabasePrivilege `json:"databasePrivileges"`
	Etag               string              `json:"etag"`
}
type DatabasePrivilege struct {
	DbName     string   `json:"dbName"`
	AuthType   string   `json:"authType"`
	Privileges []string `json:"privileges"`
}

type AccountList struct {
	Accounts []AccountInfo `json:"accounts"`
}

type RemarkArgs struct {
	Remark string `json:"remark"`
	Etag   string `json:"etag"`
}

type AuthIpArgs struct {
	Action string `json:"action"`
	Value  AuthIp `json:"value"`
	Etag   string `json:"etag"`
}
type AuthIp struct {
	Authip  []string `json:"authip"`
	Authbns []string `json:"authbns"`
}

type PrivilegesArgs struct {
	DatabasePrivileges []DatabasePrivilege `json:"databasePrivileges"`
	Etag               string              `json:"etag"`
}
type PasswordArgs struct {
	Password string `json:"password"`
	Etag     string `json:"etag"`
}

type CreateDatabaseArgs struct {
	DbName           string `json:"dbName"`
	CharacterSetName string `json:"characterSetName"`
	Remark           string `json:"remark"`
}

type DatabaseList struct {
	Databases []DatabaseInfo `json:"databases"`
}
type DatabaseInfo struct {
	DbName            string             `json:"dbName"`
	Remark            string             `json:"remark"`
	DbStatus          string             `json:"dbStatus"`
	CharacterSetName  string             `json:"characterSetName"`
	AccountPrivileges []AccountPrivilege `json:"accountPrivileges"`
}
type AccountPrivilege struct {
	AccountName string `json:"accountName"`
	AuthType    string `json:"authType"`
}

type SnapshotList struct {
	Snapshots []Snapshot `json:"snapshots"`
}
type Snapshot struct {
	SnapshotId          string `json:"snapshotId"`
	AppId               string `json:"appId"`
	SnapshotSizeInBytes int    `json:"snapshotSizeInBytes"`
	SnapshotType        string `json:"snapshotType"`
	SnapshotStatus      string `json:"snapshotStatus"`
	SnapshotStartTime   string `json:"snapshotStartTime"`
	SnapshotEndTime     string `json:"snapshotEndTime"`
	SnapshotDataTime    string `json:"snapshotDataTime"`
}

type UpdateSnapshotPolicyArgs struct {
	DataBackupWeekDay         []string                   `json:"dataBackupWeekDay"`
	DataBackupRetainStrategys []DataBackupRetainStrategy `json:"dataBackupRetainStrategys"`
	DataBackupTime            string                     `json:"dataBackupTime"`
}
type DataBackupRetainStrategy struct {
	StartSeconds int    `json:"startSeconds"`
	RetainCount  string `json:"retainCount"`
	Precision    int    `json:"precision"`
	EndSeconds   int    `json:"endSeconds"`
}
type SnapshotPolicy struct {
	Policys []SnapshotPolicyPolicy `json:"policys"`
}
type SnapshotPolicyPolicy struct {
	PolicyID                  string                     `json:"policyID"`
	PolicyName                string                     `json:"policyName"`
	DataBackupWeekDay         []string                   `json:"dataBackupWeekDay"`
	DataBackupTime            string                     `json:"dataBackupTime"`
	DataBackupRetainStrategys []DataBackupRetainStrategy `json:"dataBackupRetainStrategys"`
	LogBackupRetainDays       int                        `json:"logBackupRetainDays"`
	SpeedLimitBytesPerSec     int                        `json:"speedLimitBytesPerSec"`
	EncryptStrategy           EncryptStrategy            `json:"encryptStrategy"`
}
type EncryptStrategy struct {
	EncryptEnable            bool   `json:"encryptEnable"`
	KeyManagementType        string `json:"keyManagementType"`
	KeyManagementServiceName string `json:"keyManagementServiceName"`
	SecretKeyID              string `json:"secretKeyID"`
}
type CreateMultiactiveGroupArgs struct {
	LeaderClusterId      string `json:"leaderClusterId"`
	MultiActiveGroupName string `json:"multiActiveGroupName"`
}
type CreateMultiactiveGroupResult struct {
	MultiActiveGroupId string `json:"multiActiveGroupId"`
}

type RenameMultiactiveGroupArgs struct {
	MultiActiveGroupName string `json:"multiActiveGroupName"`
}

type MultiactiveGroupListResult struct {
	MultiActiveGroupInfoList []MultiactiveGroupInfo `json:"multiActiveGroupInfoList"`
}
type MultiactiveGroupInfo struct {
	MultiActiveGroupId   string            `json:"multiActiveGroupId"`
	MultiActiveGroupName string            `json:"multiActiveGroupName"`
	Status               string            `json:"status"`
	CreateTime           string            `json:"createTime"`
	LeaderCluster        LeaderCluster     `json:"leaderCluster"`
	FollowerClusterList  []FollowerCluster `json:"followerClusterList"`
}
type LeaderCluster struct {
	ClusterId    string `json:"clusterId"`
	ClusterName  string `json:"clusterName"`
	Region       string `json:"region"`
	Azone        string `json:"azone"`
	Status       string `json:"status"`
	ScaleOutFlag bool   `json:"scaleOutFlag"`
}
type FollowerCluster struct {
	ClusterId   string `json:"clusterId"`
	ClusterName string `json:"clusterName"`
	Region      string `json:"region"`
	Azone       string `json:"azone"`
	Status      string `json:"status"`
}

type MultiactiveGroupDetailResult struct {
	MultiActiveGroupInfo MultiactiveGroupInfo `json:"multiActiveGroupInfo"`
}

type GetSyncStatusResult struct {
	DelayTime int    `json:"delayTime"`
	Status    string `json:"status"`
}

type ExchangeArgs struct {
	ExecuteAction      string `json:"executeAction"`
	NewLeaderClusterId string `json:"newLeaderClusterId"`
}

type GetParamsListResult struct {
	Params []Param `json:"params"`
}
type Param struct {
	Status        string `json:"status"`
	UpdateTime    string `json:"updateTime"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	CaseSensitive bool   `json:"caseSensitive"`
	Modifiable    bool   `json:"modifiable"`
	CreateTime    string `json:"createTime"`
	Attention     string `json:"attention"`
	Precision     int    `json:"precision"`
	Value         string `json:"value"`
	AllowedValues string `json:"allowedValues"`
	BestValue     string `json:"bestValue"`
	Id            int64  `json:"id"`
	Description   string `json:"description"`
	AllowEmpty    bool   `json:"allowEmpty"`
}

type GetParamsHistoryResult struct {
	Records []Record `json:"records"`
}
type Record struct {
	Status      string `json:"status"`
	FinishTime  string `json:"finishTime"`
	Name        string `json:"name"`
	StartTime   string `json:"startTime"`
	AfterValue  string `json:"afterValue"`
	BeforeValue string `json:"beforeValue"`
	TaskId      int64  `json:"taskId"`
	Id          int64  `json:"id"`
}

type UpdateParamsArgs struct {
	Params ParamsItem `json:"params"`
	Timing string     `json:"timing"`
}

type ParamsItem map[string]interface{}

type ListParamTempArgs struct {
	Detail   int    `json:"detail,omitempty"`
	Type     string `json:"type,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}
type ListParamTempResult struct {
	ParamTemplates []ParamTemplate `json:"paramTemplates"`
}
type ParamTemplate struct {
	Uuid        string `json:"uuid"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Version     string `json:"version"`
	Description string `json:"description"`
	IsReboot    bool   `json:"isReboot"`
	ParamAmount int    `json:"paramAmount"`
	Status      string `json:"status"`
	IsSystem    bool   `json:"isSystem"`
	UserId      string `json:"userId,omitempty"`
	UserName    string `json:"userName,omitempty"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type ParamTempArgs struct {
	Type        string `json:"type"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Source      string `json:"source"`
}

type GetTemplateApplyRecordsResult struct {
	Records []ApplyRecord `json:"records"`
}
type ApplyRecord struct {
	Status        string `json:"status"`
	FinishTime    string `json:"finishTime"`
	ClusterId     string `json:"clusterId"`
	ClusterName   string `json:"clusterName"`
	ClusterStatus string `json:"clusterStatus"`
	ClusterRegion string `json:"clusterRegion"`
	StartTime     string `json:"startTime"`
	ErrMsg        string `json:"errMsg"`
}
type Params struct {
	Params []string `json:"params"`
}

type UpdateParamTplArgs struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ModifyParamsArgs struct {
	Params ParamsItem `json:"params"`
}

type CreateParamTemplateArgs struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Version     string `json:"version"`
	Description string `json:"description"`
}
type ParamTemplateDetail struct {
	ParamTemplate ParamTemplate `json:"paramTemplate"`
}
type ParamTemplateHistory struct {
	Records []ParamTemplateRecord `json:"records"`
}
type ParamTemplateRecord struct {
	Action          string `json:"action"`
	Name            string `json:"name"`
	ParamTemplateId string `json:"paramTemplateId"`
	RequestId       string `json:"requestId"`
	Id              int64  `json:"id"`
	ErrMsg          string `json:"errMsg"`
	AfterValue      string `json:"afterValue"`
	StartTime       string `json:"startTime"`
	FinishTime      string `json:"finishTime"`
}

type ApplyParamTemplateArgs struct {
	Timing   string     `json:"timing"`
	Clusters ParamsItem `json:"clusters"`
}

type UpdateMaintenTimeArgs struct {
	Period    string `json:"period"`
	StartTime string `json:"startTime"`
	Duration  int    `json:"duration"`
}

type MaintenTimeDetail struct {
	MaintenTime MaintenTime `json:"maintenTime"`
}
type MaintenTime struct {
	ClusterId string `json:"clusterId"`
	Period    string `json:"period"`
	StartTime string `json:"startTime"`
	Duration  int    `json:"duration"`
}

type GetSlowSqlArgs struct {
	Page     string `json:"page,omitempty"`
	PageSize string `json:"pageSize,omitempty"`
	Offset   string `json:"offset,omitempty"`
	Sort     string `json:"sort,omitempty"`
	NodeId   string `json:"nodeId,omitempty"`
	Engine   string `json:"engine,omitempty"`
	Schema   string `json:"schema,omitempty"`
	Digest   string `json:"digest,omitempty"`
	Start    string `json:"start,omitempty"`
	End      string `json:"end,omitempty"`
}

type SlowSqlDetailDetail struct {
	Items      []SlowSqlDetailItem `json:"items"`
	TotalCount int64               `json:"totalCount"`
}
type SlowSqlDetailItem struct {
	AffectedRows int64  `json:"affectedRows"`
	App          string `json:"app"`
	ClientHost   string `json:"clientHost"`
	ClientIP     string `json:"clientIP"`
	Cluster      string `json:"cluster"`
	ConnectionId int    `json:"connectionId"`
	CurrentDB    string `json:"currentDB"`
	Digest       string `json:"digest"`
	Duration     string `json:"duration"`
	ExaminedRows int64  `json:"examinedRows"`
	LockTime     string `json:"lockTime"`
	Node         string `json:"node"`
	NumRows      int64  `json:"numRows"`
	SqlId        string `json:"sqlId"`
	Start        string `json:"start"`
	User         string `json:"user"`
}

type SlowSqlAdviceDetail struct {
	IndexAdvice     []SlowSqlAdvice `json:"indexAdvice"`
	StatementAdvice []SlowSqlAdvice `json:"statementAdvice"`
}
type SlowSqlAdvice struct {
	Advice string `json:"advice"`
	Title  string `json:"title"`
}

type GetBinlogArgs struct {
	AppID         string `json:"appID"`
	LogBackupType string `json:"logBackupType"`
}
type BinlogDetail struct {
	Name             string   `json:"name"`
	LogStartDateTime string   `json:"logStartDateTime"`
	LogEndDateTime   string   `json:"logEndDateTime"`
	LogSizeBytes     int64    `json:"logSizeBytes"`
	OuterLinks       []string `json:"outerLinks"`
}
type BinlogList struct {
	LogBackups []BinlogItem `json:"logBackups"`
	Pagination Pagination   `json:"pagination"`
}
type BinlogItem struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	LogStartDateTime string `json:"logStartDateTime"`
	LogEndDateTime   string `json:"logEndDateTime"`
	LogSizeBytes     int64  `json:"logSizeBytes"`
}
type Pagination struct {
	TotalKeys int64 `json:"totalKeys"`
}

type GetBinlogListArgs struct {
	AppID         string `json:"appID"`
	LogBackupType string `json:"logBackupType"`
	PageNo        int    `json:"pageNo"`
	PageSize      int    `json:"pageSize"`
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
}

type TaskListArgs struct {
	Region    string `json:"region,omitempty"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime,omitempty"`
}
type TaskList struct {
	Task []TaskItem `json:"task"`
}
type TaskItem struct {
	TaskId           int64    `json:"taskId"`
	Region           string   `json:"region"`
	StartTime        string   `json:"startTime"`
	EndTime          string   `json:"endTime"`
	TaskName         string   `json:"taskName"`
	ClusterId        string   `json:"clusterId"`
	ClusterName      string   `json:"clusterName"`
	TaskStatus       string   `json:"taskStatus"`
	SupportedOperate []string `json:"supportedOperate"`
}

type ClusterList struct {
	Clusters []ClusterItem `json:"clusters"`
}
type ClusterItem struct {
	ClusterId string    `json:"clusterId"`
	Name      string    `json:"name"`
	LbInfos   []LbInfos `json:"lbInfos"`
}
type LbInfos struct {
	LbId string `json:"lbId"`
	LbIp string `json:"lbIp"`
	Eip  string `json:"eip"`
}
type OrderInfo struct {
	Status string `json:"status"`
}

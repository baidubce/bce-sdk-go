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
	Billing           Billing          `json:"billing"`
	PurchaseCount     int              `json:"purchaseCount"`
	InstanceName      string           `json:"instanceName"`
	NodeType          string           `json:"nodeType"`
	ShardNum          int              `json:"shardNum"`
	ProxyNum          int              `json:"proxyNum"`
	ClusterType       string           `json:"clusterType"`
	ReplicationNum    int              `json:"replicationNum"`
	ReplicationInfo   []Replication    `json:"replicationInfo"`
	Port              int              `json:"port"`
	Engine            int              `json:"engine,omitempty"`
	EngineVersion     string           `json:"engineVersion"`
	DiskFlavor        int              `json:"diskFlavor,omitempty"`
	DiskType          string           `json:"diskType,omitempty"`
	VpcID             string           `json:"vpcId"`
	Subnets           []Subnet         `json:"subnets,omitempty"`
	AutoRenewTimeUnit string           `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int              `json:"autoRenewTime,omitempty"`
	BgwGroupId        string           `json:"bgwGroupId,omitempty"`
	ClientToken       string           `json:"-"`
	ClientAuth        string           `json:"clientAuth"`
	ResourceGroupId   string           `json:"resourceGroupId"`
	StoreType         int              `json:"storeType"`
	EnableReadOnly    int              `json:"enableReadOnly,omitempty"`
	Tags              []model.TagModel `json:"tags"`
}

type CreateInstanceResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type InstanceModel struct {
	InstanceID         string           `json:"instanceId"`
	InstanceName       string           `json:"instanceName"`
	InstanceStatus     string           `json:"instanceStatus"`
	InstanceExpireTime string           `json:"instanceExpireTime"`
	ShardNum           int              `json:"shardNum"`
	ReplicationNum     int              `json:"replicationNum"`
	ClusterType        string           `json:"clusterType"`
	Engine             string           `json:"engine"`
	EngineVersion      string           `json:"engineVersion"`
	VnetIP             string           `json:"vnetIp"`
	Domain             string           `json:"domain"`
	Port               int              `json:"port"`
	InstanceCreateTime string           `json:"instanceCreateTime"`
	Capacity           float64          `json:"capacity"`
	UsedCapacity       float64          `json:"usedCapacity"`
	PaymentTiming      string           `json:"paymentTiming"`
	ZoneNames          []string         `json:"zoneNames"`
	Tags               []model.TagModel `json:"tags"`
	ResourceGroupId    string           `json:"resourceGroupId"`
	ResourceGroupName  string           `json:"resourceGroupName"`
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
	Capacity           float64          `json:"capacity"`
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
	BnsGroup           string           `json:"bnsGroup"`
	ResourceGroupId    string           `json:"resourceGroupId"`
	ResourceGroupName  string           `json:"resourceGroupName"`
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
	Capacity           float64          `json:"capacity"`
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

type CreatePriceArgs struct {
	Engine         int    `json:"engine,omitempty"`
	ClusterType    string `json:"clusterType,omitempty"`
	NodeType       string `json:"nodeType,omitempty"`
	ShardNum       int    `json:"shardNum,omitempty"`
	ReplicationNum int    `json:"replicationNum,omitempty"`
	InstanceNum    int    `json:"instanceNum,omitempty"`
	DiskType       int    `json:"diskType,omitempty"`
	DiskFlavor     int    `json:"diskFlavor,omitempty"`
	ChargeType     string `json:"chargeType,omitempty"`
	Period         int    `json:"period,omitempty"`
}
type CreatePriceResult struct {
	Price float64 `json:"price,omitempty"`
}

type ResizePriceArgs struct {
	NodeType       string `json:"nodeType"`
	ShardNum       int    `json:"shardNum,omitempty"`
	ReplicationNum int    `json:"replicationNum,omitempty"`
	DiskFlavor     int    `json:"diskFlavor,omitempty"`
	ChargeType     string `json:"chargeType,omitempty"`
	Period         int    `json:"period,omitempty"`
	ChangeType     string `json:"changeType,omitempty"`
}
type ResizePriceResult struct {
	Price float64 `json:"price,omitempty"`
}

type SetAsSlaveArgs struct {
	MasterDomain string `json:"masterDomain"`
	MasterPort   int    `json:"masterPort"`
}

type RenameDomainArgs struct {
	Domain      string `json:"domain"`
	ClientToken string `json:"clientToken,omitempty"`
}

type SwapDomainArgs struct {
	SourceInstanceId string `json:"sourceInstanceId"`
	TargetInstanceId string `json:"targetInstanceId"`
	ClientToken      string `json:"clientToken,omitempty"`
}

type GetBackupDetailResult struct {
	Url           string `json:"url"`
	UrlExpiration string `json:"urlExpiration"`
}

type GroupPreCheckArgs struct {
	Leader    GroupLeader     `json:"leader"`
	Followers []GroupFollower `json:"followers"`
}
type GroupLeader struct {
	LeaderRegion string `json:"leaderRegion"`
	LeaderId     string `json:"leaderId"`
}
type GroupFollower struct {
	FollowerId     string `json:"followerId"`
	FollowerRegion string `json:"followerRegion"`
}
type GroupPreCheckResult struct {
	LeaderResult      GroupLeaderResult       `json:"leaderResult"`
	FollowerResult    []GroupFollowerResult   `json:"followerResult"`
	ConnectionResults []GroupConnectionResult `json:"connectionResults"`
}
type GroupLeaderResult struct {
	Version         bool `json:"version"`
	ClusterStatus   bool `json:"clusterStatus"`
	ReplicationNum  bool `json:"replicationNum"`
	Flavor          bool `json:"flavor"`
	Joined          bool `json:"joined"`
	NoPasswd        bool `json:"noPasswd"`
	NoSecurityGroup bool `json:"noSecurityGroup"`
	IsHitX1         bool `json:"isHitX1"`
}
type GroupFollowerResult struct {
	FollowerId      string `json:"followerId"`
	NoData          bool   `json:"noData"`
	Version         bool   `json:"version"`
	EngineVersion   bool   `json:"engineVersion"`
	ClusterStatus   bool   `json:"clusterStatus"`
	ShardNum        bool   `json:"shardNum"`
	ReplicationNum  bool   `json:"replicationNum"`
	Flavor          bool   `json:"flavor"`
	Joined          bool   `json:"joined"`
	NoPasswd        bool   `json:"noPasswd"`
	NoSecurityGroup bool   `json:"noSecurityGroup"`
	IsHitX1         bool   `json:"isHitX1"`
}
type GroupConnectionResult struct {
	SourceId    string `json:"sourceId"`
	SourceRole  string `json:"sourceRole"`
	TargetId    string `json:"targetId"`
	TargetRole  string `json:"targetRole"`
	Connectable bool   `json:"connectable"`
}

type CreateGroupArgs struct {
	Leader CreateGroupLeader `json:"leader"`
}
type CreateGroupLeader struct {
	GroupName    string `json:"groupName"`
	LeaderRegion string `json:"leaderRegion"`
	LeaderId     string `json:"leaderId"`
}
type CreateGroupResult struct {
	GroupId string `json:"groupId"`
}
type GroupListResult struct {
	TotalCount int           `json:"totalCount"`
	PageNo     int           `json:"pageNo"`
	PageSize   int           `json:"pageSize"`
	Result     []GroupResult `json:"result"`
}
type GroupResult struct {
	GroupId         string `json:"groupId"`
	GroupName       string `json:"groupName"`
	GroupStatus     string `json:"groupStatus"`
	ClusterNum      int    `json:"clusterNum"`
	GroupCreateTime string `json:"groupCreateTime"`
	ForbidWrite     int    `json:"forbidWrite"`
	GroupType       string `json:"groupType"`
	LeaderName      string `json:"leaderName"`
	LeaderShowId    string `json:"leaderShowId"`
	LeaderRegion    string `json:"leaderRegion"`
}
type GetGroupListArgs struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

type GroupDetailResult struct {
	GroupId         string              `json:"groupId"`
	GroupName       string              `json:"groupName"`
	GroupStatus     string              `json:"groupStatus"`
	ClusterNum      int                 `json:"clusterNum"`
	GroupCreateTime string              `json:"groupCreateTime"`
	ForbidWrite     int                 `json:"forbidWrite"`
	GroupType       string              `json:"groupType"`
	Leader          GroupLeaderInfo     `json:"leader"`
	Followers       []GroupFollowerInfo `json:"followers"`
}
type GroupLeaderInfo struct {
	ClusterName       string  `json:"clusterName"`
	ClusterShowId     string  `json:"clusterShowId"`
	Region            string  `json:"region"`
	Status            string  `json:"status"`
	TotalCapacityInGB float64 `json:"totalCapacityInGB"`
	UsedCapacityInGB  int     `json:"usedCapacityInGB"`
	ShardNum          int     `json:"shardNum"`
	Flavor            int     `json:"flavor"`
	QpsWrite          int64   `json:"qpsWrite"`
	QpsRead           int64   `json:"qpsRead"`
	StableReadable    bool    `json:"stableReadable"`
	ForbidWrite       int     `json:"forbidWrite"`
	AvailabilityZone  string  `json:"availabilityZone"`
	ExpiredTime       string  `json:"expiredTime"`
}
type GroupFollowerInfo struct {
	ClusterName       string  `json:"clusterName"`
	ClusterShowId     string  `json:"clusterShowId"`
	Region            string  `json:"region"`
	Status            string  `json:"status"`
	TotalCapacityInGB float64 `json:"totalCapacityInGB"`
	UsedCapacityInGB  int     `json:"usedCapacityInGB"`
	ShardNum          int     `json:"shardNum"`
	Flavor            int     `json:"flavor"`
	QpsWrite          int64   `json:"qpsWrite"`
	QpsRead           int64   `json:"qpsRead"`
	StableReadable    bool    `json:"stableReadable"`
	ForbidWrite       int     `json:"forbidWrite"`
	AvailabilityZone  string  `json:"availabilityZone"`
	ExpiredTime       string  `json:"expiredTime"`
}

type FollowerInfo struct {
	FollowerId     string `json:"followerId"`
	FollowerRegion string `json:"followerRegion"`
	SyncMaster     string `json:"syncMaster"`
}

type GroupNameArgs struct {
	GroupName string `json:"groupName"`
}

type ForbidWriteArgs struct {
	ForbidWriteFlag bool `json:"forbidWriteFlag"`
}

type GroupSetQpsArgs struct {
	ClusterShowId string `json:"clusterShowId"`
	QpsWrite      int    `json:"qpsWrite"`
	QpsRead       int    `json:"qpsRead"`
}

type GroupSyncStatusResult struct {
	Followers []FollowerSyncInfo `json:"followers"`
}

type FollowerSyncInfo struct {
	ClusterShowId string `json:"clusterShowId"`
	SyncStatus    string `json:"syncStatus"`
	MaxOffset     int    `json:"maxOffset"`
	Lag           int    `json:"lag"`
}

type GroupWhiteList struct {
	WhiteLists []string `json:"whiteLists"`
}

type StaleReadableArgs struct {
	FollowerId    string `json:"followerId"`
	StaleReadable bool   `json:"staleReadable"`
}

type CreateTemplateArgs struct {
	EngineVersion string          `json:"engineVersion"`
	TemplateType  int             `json:"templateType"`
	ClusterType   string          `json:"clusterType"`
	Engine        string          `json:"engine"`
	Name          string          `json:"name"`
	Comment       string          `json:"comment"`
	Parameters    []ParameterItem `json:"parameters"`
}
type ParameterItem struct {
	ConfName   string `json:"confName"`
	ConfModule int    `json:"confModule"`
	ConfValue  string `json:"confValue"`
	ConfType   int    `json:"confType"`
}
type CreateParamsTemplateResult struct {
	TemplateId     int    `json:"templateId"`
	TemplateShowId string `json:"templateShowId"`
}

type ParamsTemplateListResult struct {
	Marker      string       `json:"marker"`
	MaxKeys     int          `json:"maxKeys"`
	NextMarker  string       `json:"nextMarker"`
	IsTruncated bool         `json:"isTruncated"`
	Result      []ResultItem `json:"result"`
}
type ResultItem struct {
	EngineVersion  string      `json:"engineVersion"`
	TemplateType   int         `json:"templateType"`
	ClusterType    string      `json:"clusterType"`
	NeedReboot     int         `json:"needReboot"`
	TemplateShowId string      `json:"templateShowId"`
	UpdateTime     string      `json:"updateTime"`
	TemplateId     int         `json:"templateId"`
	ParameterNum   int         `json:"parameterNum"`
	TemplateName   string      `json:"templateName"`
	Engine         string      `json:"engine"`
	CreateTime     string      `json:"createTime"`
	Comment        string      `json:"comment"`
	Parameters     []ParamItem `json:"parameters"`
}

type ParamItem struct {
	ConfName         string `json:"confName"`
	ConfModule       int    `json:"confModule"`
	ConfCacheVersion int    `json:"confCacheVersion"`
	ConfValue        string `json:"confValue"`
	NeedReboot       int    `json:"needReboot"`
	ConfRedisVersion string `json:"confRedisVersion"`
	ConfDefault      string `json:"confDefault"`
	ConfType         int    `json:"confType"`
	ConfRange        string `json:"confRange"`
	ConfDesc         string `json:"confDesc"`
	ConfUserVisible  int    `json:"confUserVisible"`
}

type RenameTemplateArgs struct {
	Name string `json:"name"`
}

type ApplyTemplateArgs struct {
	RebootType             int                  `json:"rebootType"`
	Extra                  string               `json:"extra"`
	CacheClusterShowIdItem []CacheClusterShowId `json:"cacheClusterShowId"`
	Parameters             []ParameterItem      `json:"parameters"`
}
type CacheClusterShowId struct {
	CacheClusterShowId string `json:"cacheClusterShowId"`
	Region             string `json:"region"`
}

type AddParamsArgs struct {
	Parameters []ParameterItem `json:"parameters"`
}

type ModifyParamsArgs struct {
	Parameters []ParameterItem `json:"parameters"`
}

type DeleteParamsArgs struct {
	Parameters []string `json:"parameters"`
}

type GetSystemTemplateArgs struct {
	Engine        string `json:"engine"`
	EngineVersion string `json:"engineVersion"`
	ClusterType   string `json:"clusterType"`
}

type SystemTemplateResult struct {
	Success bool             `json:"success"`
	Result  []SystemTemplate `json:"result"`
}

type SystemTemplate struct {
	ConfName         string `json:"confName"`
	ConfDefault      string `json:"confDefault"`
	ConfValue        string `json:"confValue"`
	ConfType         int    `json:"confType"`
	ConfRange        string `json:"confRange"`
	ConfModule       int    `json:"confModule"`
	ConfDesc         string `json:"confDesc"`
	NeedReboot       int    `json:"needReboot"`
	ConfRedisVersion string `json:"confRedisVersion"`
	ConfCacheVersion int    `json:"confCacheVersion"`
}

type GetApplyRecordsResult struct {
	Marker      string        `json:"marker"`
	MaxKeys     int           `json:"maxKeys"`
	NextMarker  string        `json:"nextMarker"`
	IsTruncated bool          `json:"isTruncated"`
	Result      []ApplyRecord `json:"result"`
}

type ApplyRecord struct {
	CacheClusterShowId string `json:"cacheClusterShowId"`
	CacheClusterName   string `json:"cacheClusterName"`
	AvailabilityZone   string `json:"availabilityZone"`
	Version            int    `json:"version"`
	Status             string `json:"status"`
	Engine             string `json:"engine"`
	EngineVersion      string `json:"engineVersion"`
	ClusterType        string `json:"clusterType"`
	CreateTime         string `json:"createTime"`
	ApplyTime          string `json:"applyTime"`
}

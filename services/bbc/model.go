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

package bbc

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type PaymentTimingType string

const (
	PaymentTimingPrePaid  PaymentTimingType = "Prepaid"
	PaymentTimingPostPaid PaymentTimingType = "Postpaid"
)

type InstanceStatus string

const (
	InstanceStatusRunning         InstanceStatus = "Running"
	InstanceStatusStarting        InstanceStatus = "Starting"
	InstanceStatusStopping        InstanceStatus = "Stopping"
	InstanceStatusStopped         InstanceStatus = "Stopped"
	InstanceStatusDeleted         InstanceStatus = "Deleted"
	InstanceStatusExpired         InstanceStatus = "Expired"
	InstanceStatusError           InstanceStatus = "Error"
	InstanceStatusImageProcessing InstanceStatus = "ImageProcessing"
)

type ImageType string

const (
	ImageTypeIntegration ImageType = "Integration"
	ImageTypeSystem      ImageType = "System"
	ImageTypeCustom      ImageType = "Custom"
)

type ImageStatus string

const (
	ImageStatusCreating     ImageStatus = "Creating"
	ImageStatusCreateFailed ImageStatus = "CreateFailed"
	ImageStatusAvailable    ImageStatus = "Available"
	ImageStatusNotAvailable ImageStatus = "NotAvailable"
	ImageStatusError        ImageStatus = "Error"
)

type CreateInstanceArgs struct {
	FlavorId          string           `json:"flavorId"`
	ImageId           string           `json:"imageId"`
	RaidId            string           `json:"raidId"`
	RootDiskSizeInGb  int              `json:"rootDiskSizeInGb"`
	PurchaseCount     int              `json:"purchaseCount"`
	ZoneName          string           `json:"zoneName"`
	SubnetId          string           `json:"subnetId"`
	AutoRenewTimeUnit string           `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int              `json:"autoRenewTime,omitempty"`
	Billing           Billing          `json:"billing"`
	Name              string           `json:"name,omitempty"`
	AdminPass         string           `json:"adminPass,omitempty"`
	DeploySetId       string           `json:"deploySetId,omitempty"`
	ClientToken       string           `json:"-"`
	SecurityGroupId   string           `json:"securityGroupId,omitempty"`
	Tags              []model.TagModel `json:"tags,omitempty"`
	InternalIps       []string         `json:"internalIps,omitempty"`
	RequestToken      string           `json:"requestToken"`
	EnableNuma        bool             `json:"enableNuma"`
}

type Billing struct {
	PaymentTiming PaymentTimingType `json:"paymentTiming,omitempty"`
	Reservation   Reservation       `json:"reservation,omitempty"`
}

type Reservation struct {
	Length   int    `json:"reservationLength"`
	TimeUnit string `json:"reservationTimeUnit"`
}

type CreateInstanceResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type ListInstancesArgs struct {
	Marker     string
	MaxKeys    int
	InternalIp string
	VpcId      string `json:"vpcId"`
}

type ListInstancesResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Instances   []InstanceModel `json:"instances"`
}

type InstanceModel struct {
	Id                    string           `json:"id"`
	Name                  string           `json:"name"`
	Uuid                  string           `json:"uuid"`
	Desc                  string           `json:"desc"`
	Status                InstanceStatus   `json:"status"`
	PaymentTiming         string           `json:"paymentTiming"`
	CreateTime            string           `json:"createTime"`
	ExpireTime            string           `json:"expireTime"`
	PublicIp              string           `json:"publicIp"`
	InternalIp            string           `json:"internalIp"`
	ImageId               string           `json:"imageId"`
	FlavorId              string           `json:"flavorId"`
	Zone                  string           `json:"zone"`
	Region                string           `json:"region"`
	NetworkCapacityInMbps int              `json:"networkCapacityInMbps"`
	Tags                  []model.TagModel `json:"tags"`
	SwitchId              string           `json:"switchId"`
	HostId                string           `json:"hostId"`
	RackId                string           `json:"rackId"`
}

type StopInstanceArgs struct {
	ForceStop bool `json:"forceStop,omitempty"`
}

type ModifyInstanceNameArgs struct {
	Name string `json:"name"`
}

type ModifyInstanceDescArgs struct {
	Description string `json:"desc"`
}

type RebuildInstanceArgs struct {
	ImageId        string `json:"imageId"`
	AdminPass      string `json:"adminPass"`
	IsPreserveData bool   `json:"isPreserveData"`
	RaidId         string `json:"raidId,omitempty"`
	SysRootSize    int    `json:"sysRootSize,omitempty"`
}

type GetVpcSubnetArgs struct {
	BbcIds []string `json:"bbcIds"`
}

type GetVpcSubnetResult struct {
	NetworkInfo []BbcNetworkModel `json:"networkInfo"`
}

type BbcNetworkModel struct {
	BbcId  string      `json:"bbcId"`
	Vpc    VpcModel    `json:"vpc"`
	Subnet SubnetModel `json:"subnet"`
}

type VpcModel struct {
	VpcId       string `json:"vpcId"`
	Cidr        string `json:"cidr"`
	Name        string `json:"name"`
	IsDefault   bool   `json:"isDefault"`
	Description string `json:"description"`
}

type SubnetModel struct {
	VpcId      string `json:"vpcId"`
	Name       string `json:"name"`
	SubnetType string `json:"subnetType"`
	SubnetId   string `json:"subnetId"`
	Cidr       string `json:"cidr"`
	ZoneName   string `json:"zoneName"`
}

type ModifyInstancePasswordArgs struct {
	AdminPass string `json:"adminPass"`
}

type BatchAddIpArgs struct {
	InstanceId                     string   `json:"instanceId"`
	PrivateIps                     []string `json:"privateIps"`
	SecondaryPrivateIpAddressCount int      `json:"secondaryPrivateIpAddressCount"`
}

type BatchAddIpResponse struct {
	PrivateIps []string `json:"privateIps"`
}

type BatchDelIpArgs struct {
	InstanceId string   `json:"instanceId"`
	PrivateIps []string `json:"privateIps"`
}

type BindTagsArgs struct {
	ChangeTags []model.TagModel `json:"changeTags"`
}

type UnbindTagsArgs struct {
	ChangeTags []model.TagModel `json:"changeTags"`
}

type ListImageArgs struct {
	Marker    string
	MaxKeys   int
	ImageType string
}

type ListImageResult struct {
	Marker      string       `json:"marker"`
	IsTruncated bool         `json:"isTruncated"`
	NextMarker  string       `json:"nextMarker"`
	MaxKeys     int          `json:"maxKeys"`
	Images      []ImageModel `json:"images"`
}

type ImageModel struct {
	OsVersion      string      `json:"osVersion"`
	OsArch         string      `json:"osArch"`
	Status         ImageStatus `json:"status"`
	Desc           string      `json:"desc"`
	Id             string      `json:"id"`
	Name           string      `json:"name"`
	OsName         string      `json:"osName"`
	OsBuild        string      `json:"osBuild"`
	CreateTime     string      `json:"createTime"`
	Type           ImageType   `json:"type"`
	OsType         string      `json:"osType"`
	SpecialVersion string      `json:"specialVersion"`
}

type FlavorImageModel struct {
	FlavorId string       `json:"flavorId"`
	Images   []ImageModel `json:"images"`
}

type GetImageDetailResult struct {
	Result *ImageModel `json:"image"`
}

type GetImagesResult struct {
	Result []FlavorImageModel `json:"result"`
}

type ListFlavorsResult struct {
	Flavors []FlavorModel `json:"flavors"`
}

type FlavorModel struct {
	FlavorId           string `json:"flavorID"`
	CpuCount           int    `json:"cpuCount"`
	CpuType            string `json:"cpuType"`
	MemoryCapacityInGB int    `json:"memoryCapacityInGb"`
	Disk               string `json:"disk"`
	NetworkCard        string `json:"networkCard"`
	Others             string `json:"others"`
}

type GetFlavorDetailResult struct {
	FlavorModel
}

type GetFlavorRaidResult struct {
	FlavorId string      `json:"flavorId"`
	Raids    []RaidModel `json:"raids"`
}

type RaidModel struct {
	RaidId       string  `json:"raidId"`
	Raid         string  `json:"raid"`
	SysSwapSize  int     `json:"sysSwapSize"`
	SysRootSize  int     `json:"sysRootSize"`
	SysHomeSize  int     `json:"sysHomeSize"`
	SysDiskSize  int     `json:"sysDiskSize"`
	DataDiskSize float64 `json:"dataDiskSize"`
}

type CreateImageArgs struct {
	ImageName   string `json:"imageName"`
	InstanceId  string `json:"instanceId"`
	ClientToken string `json:"-"`
}

type CreateImageResult struct {
	ImageId string `json:"imageId"`
}

type GetOperationLogArgs struct {
	Marker    string
	MaxKeys   int
	StartTime string
	EndTime   string
}

type GetOperationLogResult struct {
	Marker        string              `json:"marker"`
	IsTruncated   bool                `json:"isTruncated"`
	NextMarker    string              `json:"nextMarker"`
	MaxKeys       int                 `json:"maxKeys"`
	OperationLogs []OperationLogModel `json:"operationLogs"`
}

type OperationLogModel struct {
	OperationStatus bool   `json:"operationStatus"`
	OperationTime   string `json:"operationTime"`
	OperationDesc   string `json:"operationDesc"`
	OperationIp     string `json:"operationIp"`
}

type CreateDeploySetArgs struct {
	Strategy    string `json:"strategy"`
	Concurrency int    `json:"concurrency"`
	Name        string `json:"name,omitempty"`
	Desc        string `json:"desc,omitempty"`
	ClientToken string `json:"-"`
}

type GetFlavorImageArgs struct {
	FlavorIds   []string `json:"flavorIds"`
	ClientToken string   `json:"-"`
}

type CreateDeploySetResult struct {
	DeploySetId string `json:"deploySetId"`
}

type ListDeploySetsArgs struct {
	Marker   string `json:"marker"`
	MaxKeys  int    `json:"MaxKeys"`
	Strategy string `json:"strategy"`
}

type ListDeploySetsResult struct {
	Marker        string           `json:"marker"`
	IsTruncated   bool             `json:"isTruncated"`
	NextMarker    string           `json:"nextMarker"`
	MaxKeys       int              `json:"maxKeys"`
	DeploySetList []DeploySetModel `json:"deploySetList"`
}

type AzIntstanceStatis struct {
	ZoneName string `json:"zoneName"`
	Count    int    `json:"instanceCount"`
	Total    int    `json:"instanceTotal"`
}

type DeploySetModel struct {
	Strategy              string              `json:"strategy"`
	AzIntstanceStatisList []AzIntstanceStatis `json:"azIntstanceStatisList"`
	Name                  string              `json:"name"`
	Desc                  string              `json:"desc"`
	DeploySetId           string              `json:"deploysetId"`
	Concurrency           int                 `json:"concurrency"`
}

type GetDeploySetResult struct {
	DeploySetModel
}

type BindSecurityGroupsArgs struct {
	InstanceIds      []string `json:"instanceIds"`
	SecurityGroupIds []string `json:"securityGroups"`
}

type UnBindSecurityGroupsArgs struct {
	InstanceId      string `json:"instanceId"`
	SecurityGroupId string `json:"securityGroupId"`
}
type ListZonesResult struct {
	ZoneNames []string `json:"zoneNames"`
}

type DiskInfo struct {
	Raid           string  `json:"raid"`
	Description    string  `json:"description"`
	DataDiskName   string  `json:"dataDiskName"`
	RaidDisplay    string  `json:"raidDisplay"`
	SysAndHomeSize float64 `json:"sysAndHomeSize"`
	DataDiskSize   float64 `json:"dataDiskSize"`
	RaidId         string  `json:"raidId"`
}

type BbcFlavorInfo struct {
	Count       int                 `json:"count"`
	SataInfo    string              `json:"sataInfo"`
	Cpu         int                 `json:"cpu"`
	CpuGhz      string              `json:"cpuGhz"`
	Memory      int                 `json:"memory"`
	StorageType string              `json:"type"`
	FlavorId    string              `json:"id"`
	DiskInfos   map[string]DiskInfo `json:"diskInfos"`
}

type ListFlavorInfosResult struct {
	BbcFlavorInfoList []BbcFlavorInfo `json:"bbcFlavorInfoList"`
}

type ListFlavorZonesArgs struct {
	FlavorId    string            `json:"flavorId"`
	ProductType PaymentTimingType `json:"productType"`
}

type ListZoneFlavorsArgs struct {
	ZoneName    string            `json:"zoneName"`
	ProductType PaymentTimingType `json:"productType"`
}

type PrivateIP struct {
	PublicIpAddress  string `json:"publicIpAddress"`
	Primary          bool   `json:"primary"`
	PrivateIpAddress string `json:"privateIpAddress"`
	Ipv6Address      string `json:"ipv6Address"`
}

type GetInstanceEniResult struct {
	Id           string      `json:"eniId"`
	Name         string      `json:"name"`
	ZoneName     string      `json:"zoneName"`
	Description  string      `json:"description"`
	InstanceId   string      `json:"instanceId"`
	MacAddress   string      `json:"macAddress"`
	VpcId        string      `json:"vpcId"`
	SubnetId     string      `json:"subnetId"`
	Status       string      `json:"status"`
	PrivateIpSet []PrivateIP `json:"privateIpSet"`
}

type CreateInstanceStockArgs struct {
	FlaovrId string `json:"flavorId"`
	ZoneName string `json:"zoneName,omitempty"`
}

type InstanceStockResult struct {
	FlaovrId string `json:"flavorId"`
	Count    int    `json:"Count"`
}

type GetSimpleFlavorArgs struct {
	InstanceIds []string `json:"instanceIds"`
}

type SimpleFlavorResult struct {
	SimpleFlavorModel []SimpleFlavorModel `json:"flavorInfo"`
}

type SimpleFlavorModel struct {
	GpuCard         string `json:"gpuCard"`
	DiskDescription string `json:"diskDescription"`
	InstanceId      string `json:"instanceId"`
	MemDescription  string `json:"memDescription"`
	NicDescription  string `json:"nicDescription"`
	RamType         string `json:"ramType"`
	RamRate         string `json:"ramRate"`
	CpuDescription  string `json:"cpuDescription"`
	RaidDescription string `json:"raidDescription"`
}

type InstancePirceArgs struct {
	FlaovrId      string  `json:"flavorId"`
	PurchaseCount int     `json:"purchaseCount"`
	Billing       Billing `json:"billing"`
}

type InstancePirceResult struct {
	Pirce string `json:"price"`
}

type ListRepairTaskArgs struct {
	Marker     string `json:"marker"`
	MaxKeys    int    `json:"MaxKeys"`
	ErrResult  string `json:"errResult"`
	InstanceId string `json:"instanceId"`
}

type RepairTask struct {
	TaskId     string `json:"taskId"`
	InstanceId string `json:"instanceId"`
	ErrResult  string `json:"errResult"`
	Status     string `json:"status"`
}

type ListRepairTaskResult struct {
	Marker      string       `json:"marker"`
	IsTruncated bool         `json:"isTruncated"`
	NextMarker  string       `json:"nextMarker"`
	MaxKeys     int          `json:"maxKeys"`
	RepairTasks []RepairTask `json:"RepairTask"`
}

type ListClosedRepairTaskArgs struct {
	Marker     string `json:"marker"`
	MaxKeys    int    `json:"MaxKeys"`
	ErrResult  string `json:"errResult"`
	InstanceId string `json:"instanceId"`
	TaskId     string `json:"taskId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ClosedRepairTask struct {
	TaskId     string `json:"taskId"`
	InstanceId string `json:"instanceId"`
	ErrResult  string `json:"errResult"`
	CreateTime string `json:"createTime"`
	EndTime    string `json:"endTime"`
}

type ListClosedRepairTaskResult struct {
	Marker      string             `json:"marker"`
	IsTruncated bool               `json:"isTruncated"`
	NextMarker  string             `json:"nextMarker"`
	MaxKeys     int                `json:"maxKeys"`
	RepairTasks []ClosedRepairTask `json:"RepairTask"`
}

type GetRepairTaskResult struct {
	TaskId       string `json:"taskId"`
	InstanceId   string `json:"instanceId"`
	InstanceName string `json:"instanceName"`
	ErrResult    string `json:"errResult"`
	Status       string `json:"status"`
	ServerStatus string `json:"serverStatus"`
	Region       string `json:"region"`
	InternalIp   string `json:"internalIp"`
	FloatingIp   string `json:"floatingIp"`
}

type TaskIdArgs struct {
	TaskId string `json:"taskId"`
}

type DisconfirmTaskArgs struct {
	TaskId       string `json:"taskId"`
	NewErrResult string `json:"newErrResult"`
}

type RepairRecord struct {
	Name        string `json:"name"`
	Operator    string `json:"operator"`
	OperateTime string `json:"operateTime"`
}

type GetRepairRecords struct {
	RepairRecords []RepairRecord `json:"RepairRecord"`
}

type ListRuleArgs struct {
	Marker   string `json:"marker"`
	MaxKeys  int    `json:"maxKeys"`
	RuleName string `json:"ruleName"`
	RuleId   string `json:"ruleId"`
}

type ListRuleResult struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
	RuleList    []Rule `json:"RuleList"`
}

type Rule struct {
	RuleId           string           `json:"ruleId"`
	RuleName         string           `json:"ruleName"`
	TagCount         int              `json:"tagCount"`
	AssociateBbcNum  int              `json:"associateBbcNum"`
	ErrorBbcNum      int              `json:"errorBbcNum"`
	ErrResult        string           `json:"errResult"`
	Limit            int              `json:"limit"`
	Status           string           `json:"status"`
	AssociateBbcList []string         `json:"associateBbcList"`
	Tags             []model.TagModel `json:"tags"`
}

type CreateRuleArgs struct {
	RuleName string `json:"ruleName"`
	Limit    int    `json:"limit"`
	Enabled  int    `json:"enabled"`
	TagStr   string `json:"tagStr"`
	Extra    string `json:"extra"`
}

type CreateRuleResult struct {
	RuleId string `json:"ruleId"`
}

type DeleteRuleArgs struct {
	RuleId string `json:"ruleId"`
}

type DisableRuleArgs struct {
	RuleId string `json:"ruleId"`
}

type EnableRuleArgs struct {
	RuleId string `json:"ruleId"`
}

type DeploySetResult struct {
	Strategy     string                    `json:"strategy"`
	Name         string                    `json:"name"`
	Desc         string                    `json:"desc"`
	DeploySetId  string                    `json:"deploySetId"`
	InstanceList []AzIntstanceStatisDetail `json:"azIntstanceStatisList"`
	Concurrency  int                       `json:"concurrency"`
}

type AzIntstanceStatisDetail struct {
	ZoneName    string   `json:"zoneName"`
	Count       int      `json:"instanceCount"`
	Total       int      `json:"instanceTotal"`
	InstanceIds []string `json:"instanceIds"`
}

type BbcCreateAutoRenewArgs struct {
	InstanceId    string `json:"instanceId"`
	RenewTimeUnit string `json:"renewTimeUnit"`
	RenewTime     int    `json:"renewTime"`
}

type BbcDeleteAutoRenewArgs struct {
	InstanceId string `json:"instanceId"`
}

type DeleteInstanceIngorePaymentArgs struct {
	InstanceId            string `json:"instanceId"`
	RelatedReleaseFlag    bool   `json:"relatedReleaseFlag"`
	DeleteCdsSnapshotFlag bool   `json:"deleteCdsSnapshotFlag"`
	DeleteRelatedEnisFlag bool   `json:"deleteRelatedEnisFlag"`
}

type DeleteInstanceModel struct {
	InstanceId string `json:"instanceId"`
	Eip        string `json:"eip"`
}

type DeleteInstanceResult struct {
	SuccessResources *DeleteInstanceModel `json:"successResources"`
	FailResources    *DeleteInstanceModel `json:"failResources"`
}

type SharedUser struct {
	AccountId string `json:"accountId,omitempty"`
	Account   string `json:"account,omitempty"`
}

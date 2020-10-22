/*
 * Copyright 2017 Baidu, Inc.
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

package api

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type InstanceStatus string

const (
	InstanceStatusRunning            InstanceStatus = "Running"
	InstanceStatusStarting           InstanceStatus = "Starting"
	InstanceStatusStopping           InstanceStatus = "Stopping"
	InstanceStatusStopped            InstanceStatus = "Stopped"
	InstanceStatusDeleted            InstanceStatus = "Deleted"
	InstanceStatusScaling            InstanceStatus = "Scaling"
	InstanceStatusExpired            InstanceStatus = "Expired"
	InstanceStatusError              InstanceStatus = "Error"
	InstanceStatusSnapshotProcessing InstanceStatus = "SnapshotProcessing"
	InstanceStatusImageProcessing    InstanceStatus = "ImageProcessing"
)

type InstanceType string

const (
	InstanceTypeN1 InstanceType = "N1"
	InstanceTypeN2 InstanceType = "N2"
	InstanceTypeN3 InstanceType = "N3"
	InstanceTypeC1 InstanceType = "C1"
	InstanceTypeC2 InstanceType = "C2"
	InstanceTypeS1 InstanceType = "S1"
	InstanceTypeG1 InstanceType = "G1"
	InstanceTypeF1 InstanceType = "F1"

	// InstanceTypeN4 网络增强型 BCC 实例: 通用网络增强型g3ne、计算网络增强型c3ne、内存网络增强型m3ne
	InstanceTypeN4 InstanceType = "N4"
	// InstanceTypeN5 普通型Ⅳ BCC实例: 通用型g4、密集计算型ic4、计算型c4、内存型m4
	InstanceTypeN5 InstanceType = "N5"
)

type StorageType string

const (
	StorageTypeStd1          StorageType = "std1"
	StorageTypeHP1           StorageType = "hp1"
	StorageTypeCloudHP1      StorageType = "cloud_hp1"
	StorageTypeLocal         StorageType = "local"
	StorageTypeSATA          StorageType = "sata"
	StorageTypeSSD           StorageType = "ssd"
	StorageTypeHDDThroughput StorageType = "HDD_Throughput"
	StorageTypeHdd           StorageType = "hdd"
)

type PaymentTimingType string

const (
	PaymentTimingPrePaid  PaymentTimingType = "Prepaid"
	PaymentTimingPostPaid PaymentTimingType = "Postpaid"
	PaymentTimingBidding  PaymentTimingType = "bidding"
)

// Instance define instance model
type InstanceModel struct {
	InstanceId            string           `json:"id"`
	SerialNumber          string           `json:"serialNumber"`
	InstanceName          string           `json:"name"`
	InstanceType          InstanceType     `json:"instanceType"`
	Description           string           `json:"desc"`
	Status                InstanceStatus   `json:"status"`
	PaymentTiming         string           `json:"paymentTiming"`
	CreationTime          string           `json:"createTime"`
	ExpireTime            string           `json:"expireTime"`
	PublicIP              string           `json:"publicIp"`
	InternalIP            string           `json:"internalIp"`
	CpuCount              int              `json:"cpuCount"`
	GpuCard               string           `json:"gpuCard"`
	FpgaCard              string           `json:"fpgaCard"`
	CardCount             string           `json:"cardCount"`
	MemoryCapacityInGB    int              `json:"memoryCapacityInGB"`
	LocalDiskSizeInGB     int              `json:"localDiskSizeInGB"`
	ImageId               string           `json:"imageId"`
	NetworkCapacityInMbps int              `json:"networkCapacityInMbps"`
	PlacementPolicy       string           `json:"placementPolicy"`
	ZoneName              string           `json:"zoneName"`
	SubnetId              string           `json:"subnetId"`
	VpcId                 string           `json:"vpcId"`
	AutoRenew             bool             `json:"autoRenew"`
	KeypairId             string           `json:"keypairId"`
	KeypairName           string           `json:"keypairName"`
	DedicatedHostId       string           `json:"dedicatedHostId"`
	Tags                  []model.TagModel `json:"tags"`
	Ipv6                  string           `json:"ipv6"`
	SwitchId              string           `json:"switchId"`
	HostId                string           `json:"hostId"`
	RackId                string           `json:"rackId"`
	NicInfo               NicInfo          `json:"nicInfo"`
}

type NicInfo struct {
	Status         string    `json:"status"`
	MacAddress     string    `json:"macAddress"`
	DeviceId       string    `json:"deviceId"`
	VpcId          string    `json:"vpcId"`
	EniId          string    `json:"eniId"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	CreatedTime    string    `json:"createdTime"`
	SubnetType     string    `json:"subnetType"`
	SubnetId       string    `json:"subnetId"`
	EniNum         int       `json:"eniNum"`
	Az             string    `json:"az"`
	EniUuid        string    `json:"eniUuid"`
	Description    string    `json:"description"`
	Ips            []IpModel `json:"ips"`
	SecurityGroups []string  `json:"securityGroups"`
}

type IpModel struct {
	Eip             string `json:"eip"`
	EipStatus       string `json:"eipStatus"`
	EipSize         string `json:"eipSize"`
	EipId           string `json:"eipId"`
	Primary         string `json:"primary"`
	PrivateIp       string `json:"privateIp"`
	EipAllocationId string `json:"eipAllocationId"`
	EipType         string `json:"eipType"`
	EipGroupId      string `json:"eipGroupId"`
}

type Reservation struct {
	ReservationLength   int    `json:"reservationLength"`
	ReservationTimeUnit string `json:"reservationTimeUnit"`
}

type Billing struct {
	PaymentTiming PaymentTimingType `json:"paymentTiming,omitempty"`
	Reservation   *Reservation      `json:"reservation,omitempty"`
}

type EphemeralDisk struct {
	StorageType  StorageType `json:"storageType"`
	SizeInGB     int         `json:"sizeInGB"`
	FreeSizeInGB int         `json:"freeSizeInGB"`
}

type CreateCdsModel struct {
	CdsSizeInGB int         `json:"cdsSizeInGB"`
	StorageType StorageType `json:"storageType"`
	SnapShotId  string      `json:"snapshotId,omitempty"`
}

type DiskInfo struct {
	StorageType StorageType `json:"storageType"`
	MinDiskSize int         `json:"minDiskSize"`
	MaxDiskSize int         `json:"maxDiskSize"`
}

type DiskZoneResource struct {
	ZoneName  string     `json:"zoneName"`
	DiskInfos []DiskInfo `json:"diskInfos"`
}

type CreateInstanceArgs struct {
	ImageId               string           `json:"imageId"`
	Billing               Billing          `json:"billing"`
	InstanceType          InstanceType     `json:"instanceType,omitempty"`
	CpuCount              int              `json:"cpuCount"`
	MemoryCapacityInGB    int              `json:"memoryCapacityInGB"`
	RootDiskSizeInGb      int              `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType   StorageType      `json:"rootDiskStorageType,omitempty"`
	LocalDiskSizeInGB     int              `json:"localDiskSizeInGB,omitempty"`
	EphemeralDisks        []EphemeralDisk  `json:"ephemeralDisks,omitempty"`
	CreateCdsList         []CreateCdsModel `json:"createCdsList,omitempty"`
	NetWorkCapacityInMbps int              `json:"networkCapacityInMbps,omitempty"`
	DedicateHostId        string           `json:"dedicatedHostId,omitempty"`
	PurchaseCount         int              `json:"purchaseCount,omitempty"`
	Name                  string           `json:"name,omitempty"`
	AdminPass             string           `json:"adminPass,omitempty"`
	ZoneName              string           `json:"zoneName,omitempty"`
	SubnetId              string           `json:"subnetId,omitempty"`
	SecurityGroupId       string           `json:"securityGroupId,omitempty"`
	GpuCard               string           `json:"gpuCard,omitempty"`
	FpgaCard              string           `json:"fpgaCard,omitempty"`
	CardCount             string           `json:"cardCount,omitempty"`
	AutoRenewTimeUnit     string           `json:"autoRenewTimeUnit"`
	AutoRenewTime         int              `json:"autoRenewTime"`
	CdsAutoRenew          bool             `json:"cdsAutoRenew"`
	RelationTag           bool             `json:"relationTag,omitempty"`
	Tags                  []model.TagModel `json:"tags,omitempty"`
	DeployId              string           `json:"deployId,omitempty"`
	BidModel              string           `json:"bidModel,omitempty"`
	BidPrice              string           `json:"bidPrice,omitempty"`
	KeypairId             string           `json:"keypairId,omitempty"`
	AspId                 string           `json:"aspId,omitempty"`
	InternetChargeType    string           `json:"internetChargeType,omitempty"`
	InternalIps           []string         `json:"internalIps,omitempty"`
	ClientToken           string           `json:"-"`
	RequestToken          string           `json:"requestToken"`
}

type CreateInstanceStockArgs struct {
	EphemeralDisks     []EphemeralDisk `json:"ephemeralDisks,omitempty"`
	ZoneName           string          `json:"zoneName,omitempty"`
	CardCount          string          `json:"cardCount"`
	InstanceType       InstanceType    `json:"instanceType"`
	CpuCount           int             `json:"cpuCount"`
	MemoryCapacityInGB int             `json:"memoryCapacityInGB"`
	GpuCard            string          `json:"gpuCard"`
}

type ResizeInstanceStockArgs struct {
	EphemeralDisks     []EphemeralDisk `json:"ephemeralDisks,omitempty"`
	CpuCount           int             `json:"cpuCount"`
	MemoryCapacityInGB int             `json:"memoryCapacityInGB"`
	InstanceId         string          `json:"instanceId"`
}

type InstanceStockResult struct {
	FlaovrId string `json:"flavorId"`
	Count    int    `json:"Count"`
}

type GetBidInstancePriceArgs struct {
	InstanceType          InstanceType     `json:"instanceType"`
	CpuCount              int              `json:"cpuCount"`
	MemoryCapacityInGB    int              `json:"memoryCapacityInGB"`
	RootDiskSizeInGb      int              `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType   StorageType      `json:"rootDiskStorageType,omitempty"`
	CreateCdsList         []CreateCdsModel `json:"createCdsList,omitempty"`
	PurchaseCount         int              `json:"purchaseCount,omitempty"`
	Name                  string           `json:"name,omitempty"`
	AdminPass             string           `json:"adminPass,omitempty"`
	KeypairId             string           `json:"keypairId,omitempty"`
	AspId                 string           `json:"aspId,omitempty"`
	ImageId               string           `json:"imageId,omitempty"`
	BidModel              string           `json:"bidModel,omitempty"`
	BidPrice              string           `json:"bidPrice,omitempty"`
	NetWorkCapacityInMbps int              `json:"networkCapacityInMbps,omitempty"`
	RelationTag           bool             `json:"relationTag,omitempty"`
	Tags                  []model.TagModel `json:"tags,omitempty"`
	SecurityGroupId       string           `json:"securityGroupId,omitempty"`
	SubnetId              string           `json:"subnetId,omitempty"`
	ZoneName              string           `json:"zoneName,omitempty"`
	InternetChargeType    string           `json:"internetChargeType,omitempty"`
	ClientToken           string           `json:"-"`
}

type CreateInstanceResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type CreateInstanceBySpecArgs struct {
	ImageId               string           `json:"imageId"`
	Spec                  string           `json:"spec"`
	RootDiskSizeInGb      int              `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType   StorageType      `json:"rootDiskStorageType,omitempty"`
	EphemeralDisks        []EphemeralDisk  `json:"ephemeralDisks,omitempty"`
	CreateCdsList         []CreateCdsModel `json:"createCdsList,omitempty"`
	NetWorkCapacityInMbps int              `json:"networkCapacityInMbps,omitempty"`
	InternetChargeType    string           `json:"internetChargeType,omitempty"`
	PurchaseCount         int              `json:"purchaseCount,omitempty"`
	Name                  string           `json:"name,omitempty"`
	AdminPass             string           `json:"adminPass,omitempty"`
	Billing               Billing          `json:"billing"`
	ZoneName              string           `json:"zoneName,omitempty"`
	SubnetId              string           `json:"subnetId,omitempty"`
	SecurityGroupId       string           `json:"securityGroupId,omitempty"`
	RelationTag           bool             `json:"relationTag,omitempty"`
	Tags                  []model.TagModel `json:"tags,omitempty"`
	KeypairId             string           `json:"keypairId"`
	AutoRenewTimeUnit     string           `json:"autoRenewTimeUnit"`
	AutoRenewTime         int              `json:"autoRenewTime"`
	CdsAutoRenew          bool             `json:"cdsAutoRenew"`
	AspId                 string           `json:"aspId"`
	InternalIps           []string         `json:"internalIps,omitempty"`
	DeployId              string           `json:"deployId,omitempty"`
	ClientToken           string           `json:"-"`
	RequestToken          string           `json:"requestToken"`
}

type CreateInstanceBySpecResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type ListInstanceArgs struct {
	Marker          string
	MaxKeys         int
	InternalIp      string
	DedicatedHostId string
	ZoneName        string
	KeypairId       string
}

type ListInstanceResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Instances   []InstanceModel `json:"instances"`
}

type GetInstanceDetailResult struct {
	Instance InstanceModel `json:"instance"`
}

type ResizeInstanceArgs struct {
	CpuCount           int             `json:"cpuCount"`
	MemoryCapacityInGB int             `json:"memoryCapacityInGB"`
	EphemeralDisks     []EphemeralDisk `json:"ephemeralDisks,omitempty"`
	Spec               string          `json:"spec"`
	ClientToken        string          `json:"-"`
}

type RebuildInstanceArgs struct {
	ImageId   string `json:"imageId"`
	AdminPass string `json:"adminPass"`
	KeypairId string `json:"keypairId"`
}

type StopInstanceArgs struct {
	ForceStop        bool `json:"forceStop"`
	StopWithNoCharge bool `json:"stopWithNoCharge"`
}

type ChangeInstancePassArgs struct {
	AdminPass string `json:"adminPass"`
}

type ModifyInstanceAttributeArgs struct {
	Name string `json:"name"`
}

type ModifyInstanceDescArgs struct {
	Description string `json:"desc"`
}

type BindSecurityGroupArgs struct {
	SecurityGroupId string `json:"securityGroupId"`
}

type GetInstanceVNCResult struct {
	VNCUrl string `json:"vncUrl"`
}

type GetBidInstancePriceResult struct {
	Money    string `json:"money"`
	Count    string `json:"count"`
	PerMoney string `json:"perMoney"`
}

type ListBidFlavorResult struct {
	ZoneResources []ZoneResource `json:"zoneResources"`
}

type ZoneResource struct {
	ZoneName     string        `json:"zoneName"`
	BccResources []BccResource `json:"bccResources"`
}

type BccResource struct {
	InstanceType InstanceType `json:"instanceType"`
	Flavors      []Flavor     `json:"flavors"`
}

type Flavor struct {
	SpecId             string `json:"specId"`
	CpuCount           int    `json:"cpuCount"`
	MemoryCapacityInGB int    `json:"memoryCapacityInGB"`
	ProductType        string `json:"productType"`
	Spec               string `json:"spec"`
}

type PurchaseReservedArgs struct {
	RelatedRenewFlag string  `json:"relatedRenewFlag"`
	Billing          Billing `json:"billing"`
	ClientToken      string  `json:"-"`
}

const (
	RelatedRenewFlagCDS       string = "CDS"
	RelatedRenewFlagEIP       string = "EIP"
	RelatedRenewFlagMKT       string = "MKT"
	RelatedRenewFlagCDSEIP    string = "CDS_EIP"
	RelatedRenewFlagCDSMKT    string = "CDS_MKT"
	RelatedRenewFlagEIPMKT    string = "EIP_MKT"
	RelatedRenewFlagCDSEIPMKT string = "CDS_EIP_MKT"
)

type DeleteInstanceWithRelateResourceArgs struct {
	RelatedReleaseFlag    bool `json:"relatedReleaseFlag"`
	DeleteCdsSnapshotFlag bool `json:"deleteCdsSnapshotFlag"`
}

type InstanceChangeSubnetArgs struct {
	InstanceId string `json:"instanceId"`
	SubnetId   string `json:"subnetId"`
	Reboot     bool   `json:"reboot"`
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

type VolumeStatus string

const (
	VolumeStatusAVAILABLE          VolumeStatus = "Available"
	VolumeStatusINUSE              VolumeStatus = "InUse"
	VolumeStatusSNAPSHOTPROCESSING VolumeStatus = "SnapshotProcessing"
	VolumeStatusRECHARGING         VolumeStatus = "Recharging"
	VolumeStatusDETACHING          VolumeStatus = "Detaching"
	VolumeStatusDELETING           VolumeStatus = "Deleting"
	VolumeStatusEXPIRED            VolumeStatus = "Expired"
	VolumeStatusNOTAVAILABLE       VolumeStatus = "NotAvailable"
	VolumeStatusDELETED            VolumeStatus = "Deleted"
	VolumeStatusSCALING            VolumeStatus = "Scaling"
	VolumeStatusIMAGEPROCESSING    VolumeStatus = "ImageProcessing"
	VolumeStatusCREATING           VolumeStatus = "Creating"
	VolumeStatusATTACHING          VolumeStatus = "Attaching"
	VolumeStatusERROR              VolumeStatus = "Error"
)

type VolumeType string

const (
	VolumeTypeSYSTEM    VolumeType = "System"
	VolumeTypeEPHEMERAL VolumeType = "Ephemeral"
	VolumeTypeCDS       VolumeType = "Cds"
)

type RenameCSDVolumeArgs struct {
	Name string `json:"name"`
}

type ModifyCSDVolumeArgs struct {
	CdsName string `json:"cdsName,omitempty"`
	Desc    string `json:"desc,omitempty"`
}

type DetachVolumeArgs struct {
	InstanceId string `json:"instanceId"`
}

type PurchaseReservedCSDVolumeArgs struct {
	Billing     *Billing `json:"billing"`
	ClientToken string   `json:"-"`
}

type DeleteCDSVolumeArgs struct {
	ManualSnapshot string `json:"manualSnapshot,omitempty"`
	AutoSnapshot   string `json:"autoSnapshot,omitempty"`
}

type ModifyChargeTypeCSDVolumeArgs struct {
	Billing *Billing `json:"billing"`
}

type ListCDSVolumeResult struct {
	Marker      string        `json:"marker"`
	IsTruncated bool          `json:"isTruncated"`
	NextMarker  string        `json:"nextMarker"`
	MaxKeys     int           `json:"maxKeys"`
	Volumes     []VolumeModel `json:"volumes"`
}

type VolumeModel struct {
	Type               VolumeType               `json:"type"`
	StorageType        StorageType              `json:"storageType"`
	Id                 string                   `json:"id"`
	Name               string                   `json:"name"`
	DiskSizeInGB       int                      `json:"diskSizeInGB"`
	PaymentTiming      string                   `json:"paymentTiming"`
	ExpireTime         string                   `json:"expireTime"`
	Status             VolumeStatus             `json:"status"`
	Desc               string                   `json:"desc"`
	Attachments        []VolumeAttachmentModel  `json:"attachments"`
	ZoneName           string                   `json:"zoneName"`
	AutoSnapshotPolicy *AutoSnapshotPolicyModel `json:"autoSnapshotPolicy"`
	CreateTime         string                   `json:"createTime"`
	IsSystemVolume     bool                     `json:"isSystemVolume"`
	RegionId           string                   `json:"regionId"`
	SourceSnapshotId   string                   `json:"sourceSnapshotId"`
	SnapshotNum        string                   `json:"snapshotNum"`
	Tags               []model.TagModel         `json:"tags"`
	Encrypted          bool                     `json:"encrypted"`
}

type VolumeAttachmentModel struct {
	VolumeId   string `json:"volumeId"`
	InstanceId string `json:"instanceId"`
	Device     string `json:"device"`
	Serial     string `json:"serial"`
}

type AttachVolumeResult struct {
	VolumeAttachment *VolumeAttachmentModel `json:"volumeAttachment"`
}

type CreateCDSVolumeArgs struct {
	Name          string      `json:"name,omitempty"`
	Description   string      `json:"description,omitempty"`
	SnapshotId    string      `json:"snapshotId,omitempty"`
	ZoneName      string      `json:"zoneName,omitempty"`
	PurchaseCount int         `json:"purchaseCount,omitempty"`
	CdsSizeInGB   int         `json:"cdsSizeInGB,omitempty"`
	StorageType   StorageType `json:"storageType,omitempty"`
	Billing       *Billing    `json:"billing"`
	EncryptKey    string      `json:"encryptKey"`
	ClientToken   string      `json:"-"`
}

type CreateCDSVolumeResult struct {
	VolumeIds []string `json:"volumeIds"`
}

type GetVolumeDetailResult struct {
	Volume *VolumeModel `json:"volume"`
}

type GetAvailableDiskInfoResult struct {
	CdsUsedCapacityGB  string             `json:"cdsUsedCapacityGB"`
	CdsCreated         string             `json:"cdsCreated"`
	CdsTotalCapacityGB string             `json:"cdsTotalCapacityGB"`
	CdsTotal           string             `json:"cdsTotal"`
	CdsRatio           string             `json:"cdsRatio"`
	DiskZoneResources  []DiskZoneResource `json:"diskZoneResources"`
}

type AttachVolumeArgs struct {
	InstanceId string `json:"instanceId"`
}

type ResizeCSDVolumeArgs struct {
	NewCdsSizeInGB int         `json:"newCdsSizeInGB"`
	NewVolumeType  StorageType `json:"newVolumeType"`
	ClientToken    string      `json:"-"`
}

type RollbackCSDVolumeArgs struct {
	SnapshotId string `json:"snapshotId"`
}

type ListCDSVolumeArgs struct {
	MaxKeys    int
	InstanceId string
	ZoneName   string
	Marker     string
}

type AutoRenewCDSVolumeArgs struct {
	VolumeId      string `json:"volumeId"`
	RenewTimeUnit string `json:"renewTimeUnit"`
	RenewTime     int    `json:"renewTime"`
	ClientToken   string `json:"-"`
}

type CancelAutoRenewCDSVolumeArgs struct {
	VolumeId    string `json:"volumeId"`
	ClientToken string `json:"-"`
}

type AutoSnapshotPolicyModel struct {
	CreatedTime     string `json:"createdTime"`
	Id              string `json:"id"`
	Status          string `json:"status"`
	RetentionDays   int    `json:"retentionDays"`
	UpdatedTime     string `json:"updatedTime"`
	DeletedTime     string `json:"deletedTime"`
	LastExecuteTime string `json:"lastExecuteTime"`
	VolumeCount     int    `json:"volumeCount"`
	Name            string `json:"name"`
	TimePoints      []int  `json:"timePoints"`
	RepeatWeekdays  []int  `json:"repeatWeekdays"`
}

type SecurityGroupRuleModel struct {
	SourceIp        string `json:"sourceIp,omitempty"`
	DestIp          string `json:"destIp,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	SourceGroupId   string `json:"sourceGroupId,omitempty"`
	Ethertype       string `json:"ethertype,omitempty"`
	PortRange       string `json:"portRange,omitempty"`
	DestGroupId     string `json:"destGroupId,omitempty"`
	SecurityGroupId string `json:"securityGroupId,omitempty"`
	Remark          string `json:"remark,omitempty"`
	Direction       string `json:"direction"`
}

type SecurityGroupModel struct {
	Id    string                   `json:"id"`
	Name  string                   `json:"name"`
	Desc  string                   `json:"desc"`
	VpcId string                   `json:"vpcId"`
	Rules []SecurityGroupRuleModel `json:"rules"`
	Tags  []model.TagModel         `json:"tags"`
}

type CreateSecurityGroupArgs struct {
	ClientToken string                   `json:"-"`
	Name        string                   `json:"name"`
	Desc        string                   `json:"desc,omitempty"`
	VpcId       string                   `json:"vpcId,omitempty"`
	Rules       []SecurityGroupRuleModel `json:"rules"`
	Tags        []model.TagModel         `json:"tags,omitempty"`
}

type ListSecurityGroupArgs struct {
	Marker     string
	MaxKeys    int
	InstanceId string
	VpcId      string
}

type CreateSecurityGroupResult struct {
	SecurityGroupId string `json:"securityGroupId"`
}

type ListSecurityGroupResult struct {
	Marker         string               `json:"marker"`
	IsTruncated    bool                 `json:"isTruncated"`
	NextMarker     string               `json:"nextMarker"`
	MaxKeys        int                  `json:"maxKeys"`
	SecurityGroups []SecurityGroupModel `json:"securityGroups"`
}

type AuthorizeSecurityGroupArgs struct {
	ClientToken string                  `json:"-"`
	Rule        *SecurityGroupRuleModel `json:"rule"`
}

type RevokeSecurityGroupArgs struct {
	Rule *SecurityGroupRuleModel `json:"rule"`
}

type ImageType string

const (
	ImageTypeIntegration ImageType = "Integration"
	ImageTypeSystem      ImageType = "System"
	ImageTypeCustom      ImageType = "Custom"

	// ImageTypeAll 所有镜像类型
	ImageTypeAll ImageType = "All"
	// ImageTypeSharing 共享镜像
	ImageTypeSharing ImageType = "Sharing"
	// ImageTypeGPUSystem gpu公有
	ImageTypeGPUSystem ImageType = "GpuBccSystem"
	// ImageTypeGPUCustom gpu 自定义
	ImageTypeGPUCustom ImageType = "GpuBccCustom"
	// ImageTypeBBCSystem BBC 公有
	ImageTypeBBCSystem ImageType = "BbcSystem"
	// ImageTypeBBCCustom BBC 自定义
	ImageTypeBBCCustom ImageType = "BbcCustom"
)

type ImageStatus string

const (
	ImageStatusCreating     ImageStatus = "Creating"
	ImageStatusCreateFailed ImageStatus = "CreateFailed"
	ImageStatusAvailable    ImageStatus = "Available"
	ImageStatusNotAvailable ImageStatus = "NotAvailable"
	ImageStatusError        ImageStatus = "Error"
)

type SharedUser struct {
	AccountId string `json:"accountId,omitempty"`
	Account   string `json:"account,omitempty"`
}

type GetImageSharedUserResult struct {
	Users []SharedUser `json:"users"`
}

type GetImageOsResult struct {
	OsInfo []OsModel `json:"osInfo"`
}

type CreateImageResult struct {
	ImageId string `json:"imageId"`
}

type ListImageResult struct {
	Marker      string       `json:"marker"`
	IsTruncated bool         `json:"isTruncated"`
	NextMarker  string       `json:"nextMarker"`
	MaxKeys     int          `json:"maxKeys"`
	Images      []ImageModel `json:"images"`
}

type ImageModel struct {
	OsVersion      string          `json:"osVersion"`
	OsArch         string          `json:"osArch"`
	Status         ImageStatus     `json:"status"`
	Desc           string          `json:"desc"`
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	OsName         string          `json:"osName"`
	OsBuild        string          `json:"osBuild"`
	CreateTime     string          `json:"createTime"`
	Type           ImageType       `json:"type"`
	OsType         string          `json:"osType"`
	SpecialVersion string          `json:"specialVersion"`
	Package        bool            `json:"package"`
	Snapshots      []SnapshotModel `json:"snapshots"`
}

type GetImageDetailResult struct {
	Image *ImageModel `json:"image"`
}

type RemoteCopyImageArgs struct {
	Name       string   `json:"name,omitempty"`
	DestRegion []string `json:"destRegion"`
}

type CreateImageArgs struct {
	InstanceId  string `json:"instanceId,omitempty"`
	SnapshotId  string `json:"snapshotId,omitempty"`
	ImageName   string `json:"imageName"`
	IsRelateCds bool   `json:"relateCds"`
	ClientToken string `json:"-"`
}

type ListImageArgs struct {
	Marker    string
	MaxKeys   int
	ImageType string
}

type OsModel struct {
	OsVersion      string `json:"osVersion"`
	OsType         string `json:"osType"`
	InstanceId     string `json:"instanceId"`
	OsArch         string `json:"osArch"`
	OsName         string `json:"osName"`
	OsLang         string `json:"osLang"`
	SpecialVersion string `json:"specialVersion"`
}

type GetImageOsArgs struct {
	InstanceIds []string `json:"instanceIds"`
}

type CreateSnapshotArgs struct {
	ClientToken  string `json:"-"`
	VolumeId     string `json:"volumeId"`
	SnapshotName string `json:"snapshotName"`
	Description  string `json:"desc,omitempty"`
}

type CreateSnapshotResult struct {
	SnapshotId string `json:"snapshotId"`
}

type ListSnapshotArgs struct {
	Marker   string
	MaxKeys  int
	VolumeId string
}

type ListSnapshotChainArgs struct {
	OrderBy  string `json:"orderBy,omitempty"`
	Order    string `json:"order,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	VolumeId string `json:"volumeId,omitempty"`
}

type SnapshotStatus string

const (
	SnapshotStatusCreating      SnapshotStatus = "Creating"
	SnapshotStatusCreatedFailed SnapshotStatus = "CreatedFailed"
	SnapshotStatusAvailable     SnapshotStatus = "Available"
	SnapshotStatusNotAvailable  SnapshotStatus = "NotAvailable"
)

type SnapshotModel struct {
	Id           string         `json:"id"`
	Name         string         `json:"name"`
	SizeInGB     int            `json:"sizeInGB"`
	CreateTime   string         `json:"createTime"`
	Status       SnapshotStatus `json:"status"`
	CreateMethod string         `json:"createMethod"`
	VolumeId     string         `json:"volumeId"`
	Description  string         `json:"desc"`
	ExpireTime   string         `json:"expireTime"`
	Package      bool           `json:"package"`
	TemplateId   string         `json:"templateId"`
	InsnapId     string         `json:"insnapId"`
	Encrypted    bool           `json:"encrypted"`
}

type ListSnapshotResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Snapshots   []SnapshotModel `json:"snapshots"`
}

type ListSnapshotChainResult struct {
	OrderBy     string           `json:"orderBy"`
	TotalCount  int              `json:"totalCount"`
	PageSize    int              `json:"pageSize"`
	PageNo      int              `json:"pageNo"`
	IsTruncated bool             `json:"isTruncated"`
	Snapchains  []SnapchainModel `json:"snapchains"`
}

type SnapchainModel struct {
	Status          string `json:"status"`
	ChainSize       string `json:"chainSize"`
	ChainId         string `json:"chainId"`
	InstanceId      string `json:"instanceId"`
	UserId          string `json:"userId"`
	VolumeId        string `json:"volumeId"`
	VolumeSize      int    `json:"volumeSize"`
	ManualSnapCount int    `json:"manualSnapCount"`
	AutoSnapCount   int    `json:"autoSnapCount"`
	CreateTime      string `json:"createTime"`
}

type GetSnapshotDetailResult struct {
	Snapshot SnapshotModel `json:"snapshot"`
}

type CreateASPArgs struct {
	ClientToken    string   `json:"-"`
	Name           string   `json:"name"`
	TimePoints     []string `json:"timePoints"`
	RepeatWeekdays []string `json:"repeatWeekdays"`
	RetentionDays  string   `json:"retentionDays"`
}

type CreateASPResult struct {
	AspId string `json:"aspId"`
}

type AttachASPArgs struct {
	VolumeIds []string `json:"volumeIds"`
}

type DetachASPArgs struct {
	VolumeIds []string `json:"volumeIds"`
}

type ListASPArgs struct {
	Marker     string
	MaxKeys    int
	AspName    string
	VolumeName string
}

type ListASPResult struct {
	Marker              string                    `json:"marker"`
	IsTruncated         bool                      `json:"isTruncated"`
	NextMarker          string                    `json:"nextMarker"`
	MaxKeys             int                       `json:"maxKeys"`
	AutoSnapshotPolicys []AutoSnapshotPolicyModel `json:"autoSnapshotPolicys"`
}

type GetASPDetailResult struct {
	AutoSnapshotPolicy AutoSnapshotPolicyModel `json:"autoSnapshotPolicy"`
}

type UpdateASPArgs struct {
	Name           string   `json:"name"`
	TimePoints     []string `json:"timePoints"`
	RepeatWeekdays []string `json:"repeatWeekdays"`
	RetentionDays  string   `json:"retentionDays"`
	AspId          string   `json:"aspId"`
}

type InstanceTypeModel struct {
	Type              string `json:"type"`
	Name              string `json:"name"`
	CpuCount          int    `json:"cpuCount"`
	MemorySizeInGB    int    `json:"memorySizeInGB"`
	LocalDiskSizeInGB int    `json:"localDiskSizeInGB"`
}

type ListSpecResult struct {
	InstanceTypes []InstanceTypeModel `json:"instanceTypes"`
}

type ZoneModel struct {
	ZoneName string `json:"zoneName"`
}

type ListZoneResult struct {
	Zones []ZoneModel `json:"zones"`
}

type ListTypeZonesResult struct {
	ZoneNames []string `json:"zoneNames"`
}

type CreateDeploySetArgs struct {
	Strategy    string `json:"strategy"`
	Name        string `json:"name,omitempty"`
	Desc        string `json:"desc,omitempty"`
	Concurrency int    `json:"concurrency,omitempty"`
	ClientToken string `json:"-"`
}

type ModifyDeploySetArgs struct {
	Name        string `json:"name,omitempty"`
	Desc        string `json:"desc,omitempty"`
	ClientToken string `json:"-"`
}

type CreateDeploySetResp struct {
	DeploySetIds []string `json:"deploySetIds"`
}

type CreateDeploySetResult struct {
	DeploySetId string `json:"deploySetIds"`
}

type ListDeploySetsResult struct {
	DeploySetList []DeploySetModel `json:"deploySets"`
}

type DeploySetModel struct {
	Strategy     string              `json:"strategy"`
	InstanceList []AzIntstanceStatis `json:"azIntstanceStatisList"`
	Name         string              `json:"name"`
	Desc         string              `json:"desc"`
	DeploySetId  string              `json:"deploysetId"`
	Concurrency  int                 `json:"concurrency"`
}

type DeploySetResult struct {
	Strategy     string                    `json:"strategy"`
	Name         string                    `json:"name"`
	Desc         string                    `json:"desc"`
	DeploySetId  string                    `json:"shortId"`
	Concurrency  int                       `json:"concurrency"`
	InstanceList []AzIntstanceStatisDetail `json:"azIntstanceStatisList"`
}

type AzIntstanceStatisDetail struct {
	ZoneName    string   `json:"zoneName"`
	Count       int      `json:"instanceCount"`
	Total       int      `json:"instanceTotal"`
	InstanceIds []string `json:"instanceIds"`
}

type AzIntstanceStatis struct {
	ZoneName string `json:"zoneName"`
	Count    int    `json:"instanceCount"`
	Total    int    `json:"instanceTotal"`
}

type GetDeploySetResult struct {
	DeploySetModel
}

type RebuildBatchInstanceArgs struct {
	ImageId     string   `json:"imageId"`
	AdminPass   string   `json:"adminPass"`
	KeypairId   string   `json:"keypairId"`
	InstanceIds []string `json:"instanceIds"`
}

type ChangeToPrepaidRequest struct {
	Duration    int  `json:"duration"`
	RelationCds bool `json:"relationCds"`
}

type ChangeToPrepaidResponse struct {
	OrderId string `json:"orderId"`
}

type BindTagsRequest struct {
	ChangeTags []model.TagModel `json:"changeTags"`
}

type UnBindTagsRequest struct {
	ChangeTags []model.TagModel `json:"changeTags"`
}

type CancelBidOrderRequest struct {
	OrderId     string `json:"orderId"`
	ClientToken string `json:"-"`
}

type CreateBidInstanceResult struct {
	OrderId string `json:"orderId"`
}

type ListFlavorSpecArgs struct {
	ZoneName string `json:"zoneName,omitempty"`
}

type ListFlavorSpecResult struct {
	ZoneResources []ZoneResourceDetailSpec `json:"zoneResources"`
}

type ZoneResourceDetailSpec struct {
	ZoneName     string       `json:"zoneName"`
	BccResources BccResources `json:"bccResources"`
}

type BccResources struct {
	FlavorGroups []FlavorGroup `json:"flavorGroups"`
}

type FlavorGroup struct {
	GroupId string      `json:"groupId"`
	Flavors []BccFlavor `json:"flavors"`
}

type BccFlavor struct {
	CpuCount           int    `json:"cpuCount"`
	MemoryCapacityInGB int    `json:"memoryCapacityInGB"`
	EphemeralDiskInGb  int    `json:"ephemeralDiskInGb"`
	EphemeralDiskCount int    `json:"ephemeralDiskCount"`
	EphemeralDiskType  string `json:"ephemeralDiskType"`
	GpuCardType        string `json:"gpuCardType"`
	GpuCardCount       int    `json:"gpuCardCount"`
	FpgaCardType       string `json:"fpgaCardType"`
	FpgaCardCount      int    `json:"fpgaCardCount"`
	ProductType        string `json:"productType"`
	Spec               string `json:"spec"`
	SpecId             string `json:"specId"`
	CpuModel           string `json:"cpuModel"`
	CpuGHz             string `json:"cpuGHz"`
	NetworkBandwidth   string `json:"networkBandwidth"`
	NetworkPackage     string `json:"networkPackage"`
}

type GetPriceBySpecArgs struct {
	SpecId         string `json:"specId"`
	Spec           string `json:"spec"`
	PaymentTiming  string `json:"paymentTiming"`
	ZoneName       string `json:"zoneName"`
	PurchaseCount  int    `json:"purchaseCount,omitempty"`
	PurchaseLength int    `json:"purchaseLength"`
}

type GetPriceBySpecResult struct {
	Price []SpecIdPrices `json:"price"`
}

type SpecIdPrices struct {
	SpecId     string       `json:"specId"`
	SpecPrices []SpecPrices `json:"specPrices"`
}

type SpecPrices struct {
	Spec      string `json:"spec"`
	Status    string `json:"status"`
	SpecPrice string `json:"specPrice"`
}

type PrivateIP struct {
	PublicIpAddress  string `json:"publicIpAddress"`
	Primary          bool   `json:"primary"`
	PrivateIpAddress string `json:"privateIpAddress"`
	Ipv6Address      string `json:"ipv6Address"`
}
type Eni struct {
	EniId        string      `json:"eniId"`
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
type ListInstanceEniResult struct {
	EniList []Eni `json:"enis"`
}

type CreateKeypairArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ImportKeypairArgs struct {
	ClientToken string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PublicKey   string `json:"publicKey"`
}

type KeypairModel struct {
	KeypairId     string `json:"keypairId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	PublicKey     string `json:"publicKey"`
	RegionId      string `json:"regionId"`
	FingerPrint   string `json:"fingerPrint"`
	PrivateKey    string `json:"privateKey"`
	InstanceCount int    `json:"instanceCount"`
	CreatedTime   string `json:"createdTime"`
}

type KeypairResult struct {
	Keypair KeypairModel `json:"keypair"`
}

type AttackKeypairArgs struct {
	KeypairId   string   `json:"keypairId"`
	InstanceIds []string `json:"instanceIds"`
}

type DetachKeypairArgs struct {
	KeypairId   string   `json:"keypairId"`
	InstanceIds []string `json:"instanceIds"`
}

type DeleteKeypairArgs struct {
	KeypairId string `json:"keypairId"`
}

type ListKeypairArgs struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
}

type ListKeypairResult struct {
	Marker      string         `json:"marker"`
	IsTruncated bool           `json:"isTruncated"`
	NextMarker  string         `json:"nextMarker"`
	MaxKeys     int            `json:"maxKeys"`
	Keypairs    []KeypairModel `json:"keypairs"`
}

type RenameKeypairArgs struct {
	Name      string `json:"name"`
	KeypairId string `json:"keypairId"`
}

type KeypairUpdateDescArgs struct {
	Description string `json:"description"`
	KeypairId   string `json:"keypairId"`
}

type ListTypeZonesArgs struct {
	InstanceType string `json:"instanceType"`
	ProductType  string `json:"productType"`
	Spec         string `json:"spec"`
	SpecId       string `json:"specId"`
}

type BccCreateAutoRenewArgs struct {
	InstanceId    string `json:"instanceId"`
	RenewTimeUnit string `json:"renewTimeUnit"`
	RenewTime     int    `json:"renewTime"`
}

type BccDeleteAutoRenewArgs struct {
	InstanceId    string `json:"instanceId"`
}

type DeleteInstanceIngorePaymentArgs struct {
	InstanceId            string `json:"instanceId"`
	RelatedReleaseFlag    bool   `json:"relatedReleaseFlag"`
	DeleteCdsSnapshotFlag bool   `json:"deleteCdsSnapshotFlag"`
	DeleteRelatedEnisFlag bool   `json:"deleteRelatedEnisFlag"`
}

type DeleteInstanceModel struct {
	InstanceId  string   `json:"instanceId"`
	Eip         string   `json:"eip"`
	InsnapIds   []string `json:"insnapIds"`
	SnapshotIds []string `json:"snapshotIds"`
	VolumeIds   []string `json:"volumeIds"`
}

type DeleteInstanceResult struct {
	SuccessResources *DeleteInstanceModel `json:"successResources"`
	FailResources *DeleteInstanceModel `json:"failResources"`
}

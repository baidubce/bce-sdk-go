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
	InstanceStatusRunning             InstanceStatus = "Running"
	InstanceStatusStarting            InstanceStatus = "Starting"
	InstanceStatusStopping            InstanceStatus = "Stopping"
	InstanceStatusStopped             InstanceStatus = "Stopped"
	InstanceStatusDeleted             InstanceStatus = "Deleted"
	InstanceStatusScaling             InstanceStatus = "Scaling"
	InstanceStatusExpired             InstanceStatus = "Expired"
	InstanceStatusError               InstanceStatus = "Error"
	InstanceStatusSnapshotProcessing  InstanceStatus = "SnapshotProcessing"
	InstanceStatusImageProcessing     InstanceStatus = "ImageProcessing"
	InstanceStatusChangeVpcProcessing InstanceStatus = "ChangeVpc"
	InstanceStatusRecycled            InstanceStatus = "Recycled"
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
	StorageTypeLocalSSD      StorageType = "local-ssd"
	StorageTypeLocalHDD      StorageType = "local-hdd"
	StorageTypeLocalNVME     StorageType = "local-nvme"
	StorageTypeEnhancedPl1   StorageType = "enhanced_ssd_pl1"
	StorageTypeEnhancedPl2   StorageType = "enhanced_ssd_pl2"
)

type StorageTypeV3 string

const (
	StorageTypeV3CloudSATA          StorageTypeV3 = "Cloud_Sata"
	StorageTypeV3CloudHDDGeneral    StorageTypeV3 = "Cloud_HDD_General"
	StorageTypeV3CloudHDDThroughput StorageTypeV3 = "Cloud_HDD_Throughput"
	StorageTypeV3CloudPremium       StorageTypeV3 = "Cloud_Premium"
	StorageTypeV3CloudSSDGeneral    StorageTypeV3 = "Cloud_SSD_General"
	StorageTypeV3CloudSSDEnhanced   StorageTypeV3 = "Cloud_SSD_Enhanced"
	StorageTypeV3LocalHDD           StorageTypeV3 = "Local_HDD"
	StorageTypeV3LocalSSD           StorageTypeV3 = "Local_SSD"
	StorageTypeV3LocalNVME          StorageTypeV3 = "Local_NVME"
	StorageTypeV3LocalPVHDD         StorageTypeV3 = "Local_PV_HDD"
	StorageTypeV3LocalPVSSD         StorageTypeV3 = "Local_PV_SSD"
	StorageTypeV3LocalPVNVME        StorageTypeV3 = "Local_PV_NVME"
	StorageTypeV3EnhancedPl2        StorageTypeV3 = "enhanced_ssd_pl2"
)

type PaymentTimingType string

const (
	PaymentTimingPrePaid  PaymentTimingType = "Prepaid"
	PaymentTimingPostPaid PaymentTimingType = "Postpaid"

	// v3
	PaymentTimingSpotPaid PaymentTimingType = "Spotpaid"
	PaymentTimingBidding  PaymentTimingType = "bidding"
)

// Instance define instance model
type InstanceModel struct {
	InstanceId            string                 `json:"id"`
	SerialNumber          string                 `json:"serialNumber"`
	InstanceName          string                 `json:"name"`
	Hostname              string                 `json:"hostname"`
	InstanceType          InstanceType           `json:"instanceType"`
	Spec                  string                 `json:"spec"`
	Description           string                 `json:"desc"`
	Status                InstanceStatus         `json:"status"`
	PaymentTiming         string                 `json:"paymentTiming"`
	CreationTime          string                 `json:"createTime"`
	ExpireTime            string                 `json:"expireTime"`
	ReleaseTime           string                 `json:"releaseTime"`
	PublicIP              string                 `json:"publicIp"`
	InternalIP            string                 `json:"internalIp"`
	CpuCount              int                    `json:"cpuCount"`
	IsomerismCard         string                 `json:"isomerismCard"`
	NpuVideoMemory        string                 `json:"npuVideoMemory"`
	GpuCard               string                 `json:"gpuCard"`
	FpgaCard              string                 `json:"fpgaCard"`
	CardCount             string                 `json:"cardCount"`
	MemoryCapacityInGB    int                    `json:"memoryCapacityInGB"`
	LocalDiskSizeInGB     int                    `json:"localDiskSizeInGB"`
	ImageId               string                 `json:"imageId"`
	NetworkCapacityInMbps int                    `json:"networkCapacityInMbps"`
	PlacementPolicy       string                 `json:"placementPolicy"`
	ZoneName              string                 `json:"zoneName"`
	SubnetId              string                 `json:"subnetId"`
	VpcId                 string                 `json:"vpcId"`
	AutoRenew             bool                   `json:"autoRenew"`
	KeypairId             string                 `json:"keypairId"`
	KeypairName           string                 `json:"keypairName"`
	DedicatedHostId       string                 `json:"dedicatedHostId"`
	Tags                  []model.TagModel       `json:"tags"`
	Ipv6                  string                 `json:"ipv6"`
	EniQuota              int                    `json:"eniQuota"`
	EriQuota              int                    `json:"eriQuota"`
	RdmaTypeApi           string                 `json:"rdmaTypeApi"`
	SwitchId              string                 `json:"switchId"`
	HostId                string                 `json:"hostId"`
	DeploysetId           string                 `json:"deploysetId"`
	RackId                string                 `json:"rackId"`
	NicInfo               NicInfo                `json:"nicInfo"`
	EniNum                string                 `json:"eniNum"`
	DeploySetList         []DeploySetSimpleModel `json:"deploysetList"`
	DeletionProtection    int                    `json:"deletionProtection"`
	NetEthQueueCount      string                 `json:"netEthQueueCount"`
	Volumes               []VolumeModel          `json:"volumes"`
}

type DeploySetSimpleModel struct {
	Strategy    string `json:"strategy"`
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	DeploySetId string `json:"deploysetId"`
	Concurrency int    `json:"concurrency"`
}

type GetAllStocksResult struct {
	BccStocks []BccStock `json:"bccStocks"`
	BbcStocks []BbcStock `json:"bbcStocks"`
}

type BccStock struct {
	Spec              string `json:"spec"`
	SpecId            string `json:"specId"`
	InventoryQuantity int    `json:"inventoryQuantity"`
	RootOnLocal       bool   `json:"rootOnLocal"`
	UpdatedTime       string `json:"updatedTime"`
	CollectionTime    string `json:"collectionTime"`
	ZoneName          string `json:"logicalZone"`
}

type BbcStock struct {
	FlavorId          string `json:"flavorId"`
	InventoryQuantity int    `json:"inventoryQuantity"`
	UpdatedTime       string `json:"updatedTime"`
	CollectionTime    string `json:"collectionTime"`
	ZoneName          string `json:"logicalZone"`
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

type EphemeralDiskV3 struct {
	StorageType  StorageTypeV3 `json:"storageType"`
	SizeInGB     int           `json:"sizeInGB"`
	FreeSizeInGB int           `json:"freeSizeInGB"`
}

type CreateCdsModel struct {
	CdsSizeInGB int         `json:"cdsSizeInGB"`
	StorageType StorageType `json:"storageType"`
	SnapShotId  string      `json:"snapshotId,omitempty"`
}

type CreateCdsModelV3 struct {
	CdsSizeInGB int           `json:"cdsSizeInGB"`
	StorageType StorageTypeV3 `json:"storageType"`
	SnapShotId  string        `json:"snapshotId,omitempty"`
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
	ImageId                    string            `json:"imageId"`
	Billing                    Billing           `json:"billing"`
	InstanceType               InstanceType      `json:"instanceType,omitempty"`
	CpuCount                   int               `json:"cpuCount"`
	MemoryCapacityInGB         int               `json:"memoryCapacityInGB"`
	RootDiskSizeInGb           int               `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType        StorageType       `json:"rootDiskStorageType,omitempty"`
	LocalDiskSizeInGB          int               `json:"localDiskSizeInGB,omitempty"`
	EphemeralDisks             []EphemeralDisk   `json:"ephemeralDisks,omitempty"`
	CreateCdsList              []CreateCdsModel  `json:"createCdsList,omitempty"`
	NetWorkCapacityInMbps      int               `json:"networkCapacityInMbps,omitempty"`
	EipName                    string            `json:"eipName,omitempty"`
	DedicateHostId             string            `json:"dedicatedHostId,omitempty"`
	PurchaseCount              int               `json:"purchaseCount,omitempty"`
	Name                       string            `json:"name,omitempty"`
	Hostname                   string            `json:"hostname,omitempty"`
	IsOpenHostnameDomain       bool              `json:"isOpenHostnameDomain,omitempty"`
	AutoSeqSuffix              bool              `json:"autoSeqSuffix,omitempty"`
	AdminPass                  string            `json:"adminPass,omitempty"`
	ZoneName                   string            `json:"zoneName,omitempty"`
	SubnetId                   string            `json:"subnetId,omitempty"`
	SecurityGroupId            string            `json:"securityGroupId,omitempty"`
	EnterpriseSecurityGroupId  string            `json:"enterpriseSecurityGroupId,omitempty"`
	SecurityGroupIds           []string          `json:"securityGroupIds,omitempty"`
	EnterpriseSecurityGroupIds []string          `json:"enterpriseSecurityGroupIds,omitempty"`
	GpuCard                    string            `json:"gpuCard,omitempty"`
	FpgaCard                   string            `json:"fpgaCard,omitempty"`
	KunlunCard                 string            `json:"kunlunCard,omitempty"`
	IsomerismCard              string            `json:"isomerismCard,omitempty"`
	CardCount                  string            `json:"cardCount,omitempty"`
	AutoRenewTimeUnit          string            `json:"autoRenewTimeUnit"`
	AutoRenewTime              int               `json:"autoRenewTime"`
	CdsAutoRenew               bool              `json:"cdsAutoRenew"`
	RelationTag                bool              `json:"relationTag,omitempty"`
	IsOpenIpv6                 bool              `json:"isOpenIpv6,omitempty"`
	Tags                       []model.TagModel  `json:"tags,omitempty"`
	DeployId                   string            `json:"deployId,omitempty"`
	BidModel                   string            `json:"bidModel,omitempty"`
	BidPrice                   string            `json:"bidPrice,omitempty"`
	KeypairId                  string            `json:"keypairId,omitempty"`
	AspId                      string            `json:"aspId,omitempty"`
	InternetChargeType         string            `json:"internetChargeType,omitempty"`
	UserData                   string            `json:"userData,omitempty"`
	InternalIps                []string          `json:"internalIps,omitempty"`
	ClientToken                string            `json:"-"`
	RequestToken               string            `json:"requestToken"`
	DeployIdList               []string          `json:"deployIdList"`
	DetetionProtection         int               `json:"deletionProtection"`
	FileSystems                []FileSystemModel `json:"fileSystems,omitempty"`
	IsOpenHostEye              bool              `json:"isOpenHostEye,omitempty"`
	ResGroupId                 string            `json:"resGroupId,omitempty"`
}

type CreateInstanceArgsV2 struct {
	ImageId                    string            `json:"imageId"`
	Billing                    Billing           `json:"billing"`
	InstanceType               InstanceType      `json:"instanceType,omitempty"`
	CpuCount                   int               `json:"cpuCount"`
	MemoryCapacityInGB         int               `json:"memoryCapacityInGB"`
	RootDiskSizeInGb           int               `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType        StorageType       `json:"rootDiskStorageType,omitempty"`
	LocalDiskSizeInGB          int               `json:"localDiskSizeInGB,omitempty"`
	EphemeralDisks             []EphemeralDisk   `json:"ephemeralDisks,omitempty"`
	CreateCdsList              []CreateCdsModel  `json:"createCdsList,omitempty"`
	NetWorkCapacityInMbps      int               `json:"networkCapacityInMbps,omitempty"`
	EipName                    string            `json:"eipName,omitempty"`
	DedicateHostId             string            `json:"dedicatedHostId,omitempty"`
	PurchaseCount              int               `json:"purchaseCount,omitempty"`
	Name                       string            `json:"name,omitempty"`
	Hostname                   string            `json:"hostname,omitempty"`
	IsOpenHostnameDomain       *bool             `json:"isOpenHostnameDomain"`
	AutoSeqSuffix              *bool             `json:"autoSeqSuffix"`
	AdminPass                  string            `json:"adminPass,omitempty"`
	ZoneName                   string            `json:"zoneName,omitempty"`
	SubnetId                   string            `json:"subnetId,omitempty"`
	SecurityGroupId            string            `json:"securityGroupId,omitempty"`
	EnterpriseSecurityGroupId  string            `json:"enterpriseSecurityGroupId,omitempty"`
	SecurityGroupIds           []string          `json:"securityGroupIds,omitempty"`
	EnterpriseSecurityGroupIds []string          `json:"enterpriseSecurityGroupIds,omitempty"`
	GpuCard                    string            `json:"gpuCard,omitempty"`
	FpgaCard                   string            `json:"fpgaCard,omitempty"`
	KunlunCard                 string            `json:"kunlunCard,omitempty"`
	IsomerismCard              string            `json:"isomerismCard,omitempty"`
	CardCount                  string            `json:"cardCount,omitempty"`
	AutoRenewTimeUnit          string            `json:"autoRenewTimeUnit"`
	AutoRenewTime              int               `json:"autoRenewTime"`
	CdsAutoRenew               *bool             `json:"cdsAutoRenew"`
	RelationTag                *bool             `json:"relationTag"`
	IsOpenIpv6                 *bool             `json:"isOpenIpv6"`
	Tags                       []model.TagModel  `json:"tags,omitempty"`
	DeployId                   string            `json:"deployId,omitempty"`
	BidModel                   string            `json:"bidModel,omitempty"`
	BidPrice                   string            `json:"bidPrice,omitempty"`
	KeypairId                  string            `json:"keypairId,omitempty"`
	AspId                      string            `json:"aspId,omitempty"`
	InternetChargeType         string            `json:"internetChargeType,omitempty"`
	UserData                   string            `json:"userData,omitempty"`
	InternalIps                []string          `json:"internalIps,omitempty"`
	ClientToken                string            `json:"-"`
	RequestToken               string            `json:"requestToken"`
	DeployIdList               []string          `json:"deployIdList"`
	DetetionProtection         int               `json:"deletionProtection"`
	FileSystems                []FileSystemModel `json:"fileSystems,omitempty"`
	IsOpenHostEye              *bool             `json:"isOpenHostEye"`
	ResGroupId                 string            `json:"resGroupId,omitempty"`
}

type FileSystemModel struct {
	FsID     string `json:"fsId"`
	MountAds string `json:"mountAds"`
	Path     string `json:"path"`
	Protocol string `json:"protocol"`
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

type GetStockWithDeploySetArgs struct {
	Spec         string   `json:"spec"`
	DeploySetIds []string `json:"deploySetIds"`
}

type GetStockWithDeploySetResults struct {
	BccStocks []BccStock `json:"bccStocks"`
}

type GetStockWithSpecArgs struct {
	Spec         string   `json:"spec"`
	DeploySetIds []string `json:"deploySetIds"`
}

type GetStockWithSpecResults struct {
	BccStocks []BccStock `json:"bccStocks"`
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
	WarningList []string `json:"warningList"`
}

type CreateInstanceBySpecArgs struct {
	ImageId                    string           `json:"imageId"`
	Spec                       string           `json:"spec"`
	RootDiskSizeInGb           int              `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType        StorageType      `json:"rootDiskStorageType,omitempty"`
	EphemeralDisks             []EphemeralDisk  `json:"ephemeralDisks,omitempty"`
	CreateCdsList              []CreateCdsModel `json:"createCdsList,omitempty"`
	NetWorkCapacityInMbps      int              `json:"networkCapacityInMbps,omitempty"`
	EipName                    string           `json:"eipName,omitempty"`
	InternetChargeType         string           `json:"internetChargeType,omitempty"`
	PurchaseCount              int              `json:"purchaseCount,omitempty"`
	PurchaseMinCount           int              `json:"purchaseMinCount,omitempty"`
	Name                       string           `json:"name,omitempty"`
	Hostname                   string           `json:"hostname,omitempty"`
	IsOpenHostnameDomain       bool             `json:"isOpenHostnameDomain,omitempty"`
	AutoSeqSuffix              bool             `json:"autoSeqSuffix,omitempty"`
	AdminPass                  string           `json:"adminPass,omitempty"`
	Billing                    Billing          `json:"billing"`
	ZoneName                   string           `json:"zoneName,omitempty"`
	SubnetId                   string           `json:"subnetId,omitempty"`
	SecurityGroupId            string           `json:"securityGroupId,omitempty"`
	EnterpriseSecurityGroupId  string           `json:"enterpriseSecurityGroupId,omitempty"`
	SecurityGroupIds           []string         `json:"securityGroupIds,omitempty"`
	EnterpriseSecurityGroupIds []string         `json:"enterpriseSecurityGroupIds,omitempty"`
	RelationTag                bool             `json:"relationTag,omitempty"`
	Tags                       []model.TagModel `json:"tags,omitempty"`
	KeypairId                  string           `json:"keypairId"`
	AutoRenewTimeUnit          string           `json:"autoRenewTimeUnit"`
	AutoRenewTime              int              `json:"autoRenewTime"`
	RaidId                     string           `json:"raidId,omitempty"`
	EnableNuma                 bool             `json:"enableNuma,omitempty"`
	DataPartitionType          string           `json:"dataPartitionType,omitempty"`
	RootPartitionType          string           `json:"rootPartitionType,omitempty"`
	CdsAutoRenew               bool             `json:"cdsAutoRenew"`
	AspId                      string           `json:"aspId"`
	InternalIps                []string         `json:"internalIps,omitempty"`
	DeployId                   string           `json:"deployId,omitempty"`
	UserData                   string           `json:"userData,omitempty"`
	ClientToken                string           `json:"-"`
	RequestToken               string           `json:"requestToken"`
	DeployIdList               []string         `json:"deployIdList"`
	DetetionProtection         int              `json:"deletionProtection"`
	IsOpenIpv6                 bool             `json:"isOpenIpv6,omitempty"`
	SpecId                     string           `json:"specId,omitempty"`
	IsOpenHostEye              bool             `json:"isOpenHostEye,omitempty"`
	BidModel                   string           `json:"bidModel,omitempty"`
	BidPrice                   string           `json:"bidPrice,omitempty"`
	ResGroupId                 string           `json:"resGroupId,omitempty"`
}

type CreateInstanceBySpecArgsV2 struct {
	ImageId                    string           `json:"imageId"`
	Spec                       string           `json:"spec"`
	RootDiskSizeInGb           int              `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType        StorageType      `json:"rootDiskStorageType,omitempty"`
	EphemeralDisks             []EphemeralDisk  `json:"ephemeralDisks,omitempty"`
	CreateCdsList              []CreateCdsModel `json:"createCdsList,omitempty"`
	NetWorkCapacityInMbps      int              `json:"networkCapacityInMbps,omitempty"`
	EipName                    string           `json:"eipName,omitempty"`
	InternetChargeType         string           `json:"internetChargeType,omitempty"`
	PurchaseCount              int              `json:"purchaseCount,omitempty"`
	PurchaseMinCount           int              `json:"purchaseMinCount,omitempty"`
	Name                       string           `json:"name,omitempty"`
	Hostname                   string           `json:"hostname,omitempty"`
	IsOpenHostnameDomain       *bool            `json:"isOpenHostnameDomain"`
	AutoSeqSuffix              *bool            `json:"autoSeqSuffix"`
	AdminPass                  string           `json:"adminPass,omitempty"`
	Billing                    Billing          `json:"billing"`
	ZoneName                   string           `json:"zoneName,omitempty"`
	SubnetId                   string           `json:"subnetId,omitempty"`
	SecurityGroupId            string           `json:"securityGroupId,omitempty"`
	EnterpriseSecurityGroupId  string           `json:"enterpriseSecurityGroupId,omitempty"`
	SecurityGroupIds           []string         `json:"securityGroupIds,omitempty"`
	EnterpriseSecurityGroupIds []string         `json:"enterpriseSecurityGroupIds,omitempty"`
	RelationTag                *bool            `json:"relationTag"`
	Tags                       []model.TagModel `json:"tags,omitempty"`
	KeypairId                  string           `json:"keypairId"`
	AutoRenewTimeUnit          string           `json:"autoRenewTimeUnit"`
	AutoRenewTime              int              `json:"autoRenewTime"`
	RaidId                     string           `json:"raidId,omitempty"`
	EnableNuma                 *bool            `json:"enableNuma"`
	DataPartitionType          string           `json:"dataPartitionType,omitempty"`
	RootPartitionType          string           `json:"rootPartitionType,omitempty"`
	CdsAutoRenew               *bool            `json:"cdsAutoRenew"`
	AspId                      string           `json:"aspId"`
	InternalIps                []string         `json:"internalIps,omitempty"`
	DeployId                   string           `json:"deployId,omitempty"`
	UserData                   string           `json:"userData,omitempty"`
	ClientToken                string           `json:"-"`
	RequestToken               string           `json:"requestToken"`
	DeployIdList               []string         `json:"deployIdList"`
	DetetionProtection         int              `json:"deletionProtection"`
	IsOpenIpv6                 *bool            `json:"isOpenIpv6"`
	SpecId                     string           `json:"specId,omitempty"`
	IsOpenHostEye              *bool            `json:"isOpenHostEye"`
	BidModel                   string           `json:"bidModel,omitempty"`
	BidPrice                   string           `json:"bidPrice,omitempty"`
	ResGroupId                 string           `json:"resGroupId,omitempty"`
	EnableHt                   *bool            `json:"enableHt"`
}

const (
	LabelOperatorEqual    LabelOperator = "equal"
	LabelOperatorNotEqual LabelOperator = "not_equal"
	LabelOperatorExist    LabelOperator = "exist"
	LabelOperatorNotExist LabelOperator = "not_exist"
)

type LabelOperator string

type LabelConstraint struct {
	Key      string        `json:"labelKey,omitempty"`
	Value    string        `json:"labelValue,omitempty"`
	Operator LabelOperator `json:"operatorName,omitempty"`
}

// --- 创建 BCC 的新接口的参数和返回值

type CreateSpecialInstanceBySpecArgs struct {
	ImageId                    string           `json:"imageId"`
	Spec                       string           `json:"spec"`
	RootDiskSizeInGb           int              `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType        StorageType      `json:"rootDiskStorageType,omitempty"`
	EphemeralDisks             []EphemeralDisk  `json:"ephemeralDisks,omitempty"`
	CreateCdsList              []CreateCdsModel `json:"createCdsList,omitempty"`
	NetWorkCapacityInMbps      int              `json:"networkCapacityInMbps,omitempty"`
	InternetChargeType         string           `json:"internetChargeType,omitempty"`
	PurchaseCount              int              `json:"purchaseCount,omitempty"`
	Name                       string           `json:"name,omitempty"`
	Hostname                   string           `json:"hostname,omitempty"`
	IsOpenHostnameDomain       bool             `json:"isOpenHostnameDomain,omitempty"`
	UserData                   string           `json:"userData,omitempty"`
	AutoSeqSuffix              bool             `json:"autoSeqSuffix,omitempty"`
	AdminPass                  string           `json:"adminPass,omitempty"`
	Billing                    Billing          `json:"billing"`
	ZoneName                   string           `json:"zoneName,omitempty"`
	SubnetId                   string           `json:"subnetId,omitempty"`
	SecurityGroupId            string           `json:"securityGroupId,omitempty"`
	EnterpriseSecurityGroupId  string           `json:"enterpriseSecurityGroupId,omitempty"`
	SecurityGroupIds           []string         `json:"securityGroupIds,omitempty"`
	EnterpriseSecurityGroupIds []string         `json:"enterpriseSecurityGroupIds,omitempty"`
	RelationTag                bool             `json:"relationTag,omitempty"`
	Tags                       []model.TagModel `json:"tags,omitempty"`
	KeypairId                  string           `json:"keypairId"`
	AutoRenewTimeUnit          string           `json:"autoRenewTimeUnit"`
	AutoRenewTime              int              `json:"autoRenewTime"`
	CdsAutoRenew               bool             `json:"cdsAutoRenew"`
	AspId                      string           `json:"aspId"`
	InternalIps                []string         `json:"internalIps,omitempty"`
	DeployId                   string           `json:"deployId,omitempty"`
	ClientToken                string           `json:"-"`
	RequestToken               string           `json:"requestToken"`
	DeployIdList               []string         `json:"deployIdList"`
	DetetionProtection         int              `json:"deletionProtection"`

	// CreateInstanceBySpecArgs 的基础上增加的参数
	LabelConstraints []LabelConstraint `json:"labelConstraints,omitempty"`

	ResGroupId string `json:"resGroupId,omitempty"`
}

type CreateSpecialInstanceBySpecResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type CreateInstanceV3Args struct {
	InstanceSpec          string                `json:"instanceSpec,omitempty"`
	SystemVolume          SystemVolume          `json:"systemVolume,omitempty"`
	DataVolumes           []DataVolume          `json:"dataVolumes,omitempty"`
	PurchaseCount         int                   `json:"purchaseCount,omitempty"`
	InstanceName          string                `json:"instanceName,omitempty"`
	HostName              string                `json:"hostName,omitempty"`
	AutoSeqSuffix         bool                  `json:"autoSeqSuffix,omitempty"`
	HostNameDomain        bool                  `json:"hostNameDomain,omitempty"`
	Password              string                `json:"password,omitempty"`
	Billing               Billing               `json:"billing"`
	ZoneName              string                `json:"zoneName,omitempty"`
	SubnetId              string                `json:"subnetId,omitempty"`
	SecurityGroupIds      []string              `json:"securityGroupIds,omitempty"`
	AssociatedResourceTag bool                  `json:"associatedResourceTag,omitempty"`
	Tags                  []model.TagModel      `json:"tags,omitempty"`
	KeypairId             string                `json:"keypairId,omitempty"`
	AutoRenewTime         int                   `json:"autoRenewTime,omitempty"`
	CdsAutoRenew          bool                  `json:"cdsAutoRenew,omitempty"`
	AutoSnapshotPolicyId  string                `json:"autoSnapshotPolicyId,omitempty"`
	PrivateIpAddresses    []string              `json:"privateIpAddresses,omitempty"`
	DeploymentSetId       string                `json:"deploymentSetId,omitempty"`
	DeployIdList          []string              `json:"deployIdList"`
	ImageId               string                `json:"imageId,omitempty"`
	UserData              string                `json:"userData,omitempty"`
	InstanceMarketOptions InstanceMarketOptions `json:"instanceMarketOptions,omitempty"`
	Ipv6                  bool                  `json:"ipv6,omitempty"`
	DedicatedHostId       string                `json:"dedicatedHostId,omitempty"`
	InternetAccessible    InternetAccessible    `json:"internetAccessible,omitempty"`
	ClientToken           string                `json:"-"`
	RequestToken          string                `json:"requestToken"`
	ResGroupId            string                `json:"resGroupId,omitempty"`
}

type CreateInstanceV3Result struct {
	InstanceIds []string `json:"instanceIds"`
}

type SystemVolume struct {
	StorageType StorageTypeV3 `json:"storageType,omitempty"`
	VolumeSize  int           `json:"volumeSize,omitempty"`
}

type DataVolume struct {
	StorageType StorageTypeV3 `json:"storageType,omitempty"`
	VolumeSize  int           `json:"volumeSize,omitempty"`
	SnapshotId  string        `json:"snapshotId,omitempty"`
	EncryptKey  string        `json:"encryptKey,omitempty"`
}

type InstanceMarketOptions struct {
	SpotOption string `json:"spotOption,omitempty"`
	SpotPrice  string `json:"spotPrice,omitempty"`
}

type InternetAccessible struct {
	InternetMaxBandwidthOut int                `json:"internetMaxBandwidthOut,omitempty"`
	InternetChargeType      InternetChargeType `json:"internetChargeType,omitempty"`
}

type InternetChargeType string

const (
	BandwidthPrepaid        InternetChargeType = "BANDWIDTH_PREPAID"
	TrafficPostpaidByHour   InternetChargeType = "TRAFFIC_POSTPAID_BY_HOUR"
	BandwidthPostpaidByHour InternetChargeType = "BANDWIDTH_POSTPAID_BY_HOUR"
)

type CreateInstanceBySpecResult struct {
	InstanceIds []string `json:"instanceIds"`
	OrderId     string   `json:"orderId,omitempty"`
}

type ListInstanceArgs struct {
	Marker           string
	MaxKeys          int
	InternalIp       string
	DedicatedHostId  string
	ZoneName         string
	KeypairId        string
	AutoRenew        bool
	InstanceIds      string
	InstanceNames    string
	CdsIds           string
	DeploySetIds     string
	SecurityGroupIds string
	PaymentTiming    string
	Status           string
	Tags             string
	VpcId            string
	PrivateIps       string
}

type ListInstanceResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Instances   []InstanceModel `json:"instances"`
}

type ListRecycleInstanceArgs struct {
	Marker        string `json:"marker,omitempty"`
	MaxKeys       int    `json:"maxKeys,omitempty"`
	InstanceId    string `json:"instanceId,omitempty"`
	Name          string `json:"name,omitempty"`
	PaymentTiming string `json:"paymentTiming,omitempty"`
	RecycleBegin  string `json:"recycleBegin,omitempty"`
	RecycleEnd    string `json:"recycleEnd,omitempty"`
}

type ListServerRequestV3Args struct {
	Marker            string         `json:"marker,omitempty"`
	MaxKeys           int            `json:"maxKeys,omitempty"`
	InstanceId        string         `json:"instanceId,omitempty"`
	InstanceName      string         `json:"instanceName,omitempty"`
	PrivateIpAddress  string         `json:"privateIpAddress,omitempty"`
	PublicIpAddress   string         `json:"publicIpAddress,omitempty"`
	VpcName           string         `json:"vpcName,omitempty"`
	SubnetName        string         `json:"subnetName,omitempty"`
	SubnetId          string         `json:"subnetId,omitempty"`
	DedicatedHostId   string         `json:"dedicatedHostId,omitempty"`
	ZoneName          string         `json:"zoneName,omitempty"`
	AutoRenew         bool           `json:"autoRenew,omitempty"`
	KeypairId         string         `json:"keypairId,omitempty"`
	KeypairName       string         `json:"keypairName,omitempty"`
	DeploymentSetId   string         `json:"deploymentSetId,omitempty"`
	DeploymentSetName string         `json:"deploymentSetName,omitempty"`
	ResGroupId        string         `json:"resGroupId,omitempty"`
	Tag               model.TagModel `json:"tag,omitempty"`
}

type LogicMarkerResultResponseV3 struct {
	Marker      string            `json:"marker"`
	IsTruncated bool              `json:"isTruncated"`
	NextMarker  string            `json:"nextMarker"`
	MaxKeys     int               `json:"maxKeys"`
	Instances   []InstanceModelV3 `json:"instances"`
}

type ListRecycleInstanceResult struct {
	Marker      string                 `json:"marker"`
	IsTruncated bool                   `json:"isTruncated"`
	NextMarker  string                 `json:"nextMarker"`
	MaxKeys     int                    `json:"maxKeys"`
	Instances   []RecycleInstanceModel `json:"instances"`
}

type InstanceModelV3 struct {
	InstanceId            string             `json:"instanceId"`
	InstanceName          string             `json:"instanceName"`
	HostId                string             `json:"hostId"`
	HostName              string             `json:"hostName"`
	InstanceSpec          string             `json:"instanceSpec"`
	Status                InstanceStatus     `json:"status"`
	Description           string             `json:"description"`
	PaymentTiming         string             `json:"paymentTiming"`
	CreateTime            string             `json:"createTime"`
	ExpireTime            string             `json:"expireTime"`
	ReleaseTime           string             `json:"releaseTime"`
	PrivateIpAddress      string             `json:"privateIpAddress"`
	PublicIpAddress       string             `json:"publicIpAddress"`
	Cpu                   int                `json:"cpu"`
	Memory                int                `json:"memory"`
	GpuCard               string             `json:"gpuCard"`
	FpgaCard              string             `json:"fpgaCard"`
	CardCount             int                `json:"cardCount"`
	DataVolumes           []DataVolumeV3     `json:"dataVolumes"`
	ImageId               string             `json:"imageId"`
	NetworkCapacityInMbps InternetAccessible `json:"networkCapacityInMbps"`
	ZoneName              string             `json:"zoneName"`
	SubnetId              string             `json:"subnetId"`
	VpcId                 string             `json:"vpcId"`
	AutoRenew             bool               `json:"autoRenew"`
	KeypairId             string             `json:"keypairId"`
	KeypairName           string             `json:"keypairName"`
	HypervisorDedicatedId string             `json:"hypervisorDedicatedId"`
	Ipv6                  string             `json:"ipv6"`
	Tags                  []model.TagModel   `json:"tags"`
	DeployId              []string           `json:"deployId"`
	SerialNumber          string             `json:"serialNumber"`
	SwitchId              string             `json:"switchId"`
	RackId                string             `json:"rackId"`
	NicInfo               NicInfoV3          `json:"nicInfo"`
	OsName                string             `json:"osName"`
	OsType                string             `json:"osType"`
}

type NicInfoV3 struct {
	MacAddress     string      `json:"macAddress"`
	EniId          string      `json:"eniId"`
	Type           string      `json:"type"`
	Ips            []IpModelV3 `json:"ips"`
	SecurityGroups []string    `json:"securityGroups"`
}

type IpModelV3 struct {
	Primary   bool   `json:"primary"`
	PrivateIp string `json:"privateIp"`
}

type DataVolumeV3 struct {
	VolumeId       string `json:"volumeId"`
	VolumeType     string `json:"volumeType"`
	VolumeSizeInGb int    `json:"volumeSizeInGb"`
	StorageType    string `json:"storageType"`
	SnapshotId     string `json:"snapshotId"`
	EncryptKey     string `json:"encryptKey"`
}

type RecycleInstanceModel struct {
	InstanceId    string                         `json:"id"`
	SerialNumber  string                         `json:"serialNumber"`
	InstanceName  string                         `json:"name"`
	RecycleTime   string                         `json:"recycleTime"`
	DeleteTime    string                         `json:"deleteTime"`
	PaymentTiming string                         `json:"paymentTiming"`
	ServiceName   string                         `json:"serviceName"`
	ServiceType   string                         `json:"serviceType"`
	ConfigItems   []string                       `json:"configItems"`
	ConfigItem    RecycleInstanceModelConfigItem `json:"configItem"`
}

type RecycleInstanceModelConfigItem struct {
	Cpu      int    `json:"cpu"`
	Memory   int    `json:"memory"`
	Type     string `json:"type"`
	SpecId   string `json:"specId"`
	Spec     string `json:"spec"`
	ZoneName string `json:"zoneName"`
}

type ModifyInstanceHostnameArgs struct {
	Hostname             string `json:"hostname"`
	IsOpenHostnameDomain bool   `json:"isOpenHostnameDomain"`
	Reboot               bool   `json:"reboot"`
}

type GetInstanceDetailResult struct {
	Instance InstanceModel `json:"instance"`
}

type AutoReleaseArgs struct {
	ReleaseTime string `json:"releaseTime"`
}

type ResizeInstanceArgs struct {
	CpuCount           int             `json:"cpuCount"`
	GpuCardCount       int             `json:"gpuCardCount"`
	MemoryCapacityInGB int             `json:"memoryCapacityInGB"`
	EphemeralDisks     []EphemeralDisk `json:"ephemeralDisks,omitempty"`
	Spec               string          `json:"spec"`
	LiveResize         bool            `json:"liveResize"`
	ClientToken        string          `json:"-"`
}

type RebuildInstanceArgs struct {
	ImageId           string `json:"imageId"`
	AdminPass         string `json:"adminPass"`
	KeypairId         string `json:"keypairId"`
	IsOpenHostEye     bool   `json:"isOpenHostEye"`
	IsPreserveData    bool   `json:"isPreserveData"`
	RaidId            string `json:"raidId,omitempty"`
	SysRootSize       int    `json:"sysRootSize,omitempty"`
	RootPartitionType string `json:"rootPartitionType,omitempty"`
	DataPartitionType string `json:"dataPartitionType,omitempty"`
}

type RebuildInstanceArgsV2 struct {
	ImageId           string `json:"imageId"`
	AdminPass         string `json:"adminPass"`
	KeypairId         string `json:"keypairId"`
	IsOpenHostEye     *bool  `json:"isOpenHostEye"`
	IsPreserveData    *bool  `json:"isPreserveData"`
	RaidId            string `json:"raidId,omitempty"`
	SysRootSize       int    `json:"sysRootSize,omitempty"`
	RootPartitionType string `json:"rootPartitionType,omitempty"`
	DataPartitionType string `json:"dataPartitionType,omitempty"`
}

type StopInstanceArgs struct {
	ForceStop        bool `json:"forceStop"`
	StopWithNoCharge bool `json:"stopWithNoCharge"`
}

type ChangeInstancePassArgs struct {
	AdminPass string `json:"adminPass"`
}

type ModifyInstanceAttributeArgs struct {
	Name             string `json:"name"`
	NetEthQueueCount string `json:"netEthQueueCount"`
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

type InstancePurchaseReservedResult struct {
	OrderId string `json:"orderId"`
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
	BccRecycleFlag        bool `json:"bccRecycleFlag"`
	DeleteRelatedEnisFlag bool `json:"deleteRelatedEnisFlag"`
	DeleteImmediate       bool `json:"deleteImmediate"`
}

type DeletePrepaidInstanceWithRelateResourceArgs struct {
	InstanceId            string `json:"instanceId"`
	RelatedReleaseFlag    bool   `json:"relatedReleaseFlag"`
	DeleteCdsSnapshotFlag bool   `json:"deleteCdsSnapshotFlag"`
	DeleteRelatedEnisFlag bool   `json:"deleteRelatedEnisFlag"`
}

type ReleasePrepaidInstanceResponse struct {
	SuccessResources   InstanceDeleteResultModel `json:"successResources"`
	FailResources      InstanceDeleteResultModel `json:"failResources"`
	InstanceRefundFlag bool                      `json:"instanceRefundFlag"`
}

type InstanceDeleteResultModel struct {
	InstanceId  string   `json:"instanceId"`
	Eip         string   `json:"eip"`
	InsnapIds   []string `json:"insnapIds"`
	SnapshotIds []string `json:"snapshotIds"`
	VolumeIds   []string `json:"volumeIds"`
}

type InstanceChangeVpcArgs struct {
	InstanceId                 string   `json:"instanceId"`
	SubnetId                   string   `json:"subnetId"`
	InternalIp                 string   `json:"internalIp"`
	Reboot                     bool     `json:"reboot"`
	SecurityGroupIds           []string `json:"securityGroupIds"`
	EnterpriseSecurityGroupIds []string `json:"enterpriseSecurityGroupIds"`
}

type InstanceChangeSubnetArgs struct {
	InstanceId string `json:"instanceId"`
	SubnetId   string `json:"subnetId"`
	InternalIp string `json:"internalIp"`
	Reboot     bool   `json:"reboot"`
}

type BatchAddIpArgs struct {
	InstanceId                     string   `json:"instanceId"`
	PrivateIps                     []string `json:"privateIps"`
	SecondaryPrivateIpAddressCount int      `json:"secondaryPrivateIpAddressCount"`
	AllocateMultiIpv6Addr          bool     `json:"allocateMultiIpv6Addr"`
	ClientToken                    string   `json:"-"`
}

type BatchAddIpResponse struct {
	PrivateIps []string `json:"privateIps"`
}

type BatchDelIpArgs struct {
	InstanceId  string   `json:"instanceId"`
	PrivateIps  []string `json:"privateIps"`
	ClientToken string   `json:"-"`
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

type VolumeStatusV3 string

const (
	VolumeStatusV3AVAILABLE          VolumeStatusV3 = "Available"
	VolumeStatusV3INUSE              VolumeStatusV3 = "InUse"
	VolumeStatusV3SNAPSHOTPROCESSING VolumeStatusV3 = "SnapshotProcessing"
	VolumeStatusV3RECHARGING         VolumeStatusV3 = "Recharging"
	VolumeStatusV3DETACHING          VolumeStatusV3 = "Detaching"
	VolumeStatusV3DELETING           VolumeStatusV3 = "Deleting"
	VolumeStatusV3EXPIRED            VolumeStatusV3 = "Expired"
	VolumeStatusV3NOTAVAILABLE       VolumeStatusV3 = "NotAvailable"
	VolumeStatusV3DELETED            VolumeStatusV3 = "Deleted"
	VolumeStatusV3SCALING            VolumeStatusV3 = "Scaling"
	VolumeStatusV3IMAGEPROCESSING    VolumeStatusV3 = "ImageProcessing"
	VolumeStatusV3CREATING           VolumeStatusV3 = "Creating"
	VolumeStatusV3ATTACHING          VolumeStatusV3 = "Attaching"
	VolumeStatusV3ERROR              VolumeStatusV3 = "Error"
	VolumeStatusV3Recycled           VolumeStatusV3 = "Recycled"
)

type VolumeType string

const (
	VolumeTypeSYSTEM    VolumeType = "System"
	VolumeTypeEPHEMERAL VolumeType = "Ephemeral"
	VolumeTypeCDS       VolumeType = "Cds"
)

type VolumeTypeV3 string

const (
	VolumeTypeV3SYSTEM VolumeTypeV3 = "SYSTEM"
	VolumeTypeV3DATA   VolumeTypeV3 = "DATA"
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
	Recycle        string `json:"recycle,omitempty"`
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

type ListCDSVolumeResultV3 struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Volumes     []VolumeModelV3 `json:"volumes"`
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
	EbcDiskSize        int                      `json:"ebcDiskSize"`
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
	EnableAutoRenew    bool                     `json:"enableAutoRenew"`
	AutoRenewTime      int                      `json:"autoRenewTime"`
	Encrypted          bool                     `json:"encrypted"`
	ClusterId          string                   `json:"clusterId"`
	RoleName           string                   `json:"roleName"`
	CreatedFrom        string                   `json:"createdFrom"`
	ReleaseTime        string                   `json:"releaseTime"`
	VolumeId           string                   `json:"volumeId"`
}

type VolumeModelV3 struct {
	Id                   string                   `json:"volumeId"`
	Name                 string                   `json:"volumeName"`
	VolumeSize           int                      `json:"volumeSizeInGB"`
	VolumeStatus         VolumeStatusV3           `json:"volumeStatus"`
	VolumeType           VolumeTypeV3             `json:"volumeType"`
	StorageType          StorageTypeV3            `json:"storageType"`
	CreateTime           string                   `json:"createTime"`
	ExpireTime           string                   `json:"expireTime"`
	Desc                 string                   `json:"description"`
	PaymentTiming        string                   `json:"paymentTiming"`
	EnableAutoRenew      bool                     `json:"enableAutoRenew"`
	AutoRenewTime        int                      `json:"autoRenewTime"`
	ZoneName             string                   `json:"zoneName"`
	SourceSnapshotId     string                   `json:"sourceSnapshotId"`
	Region               string                   `json:"region"`
	SnapshotCount        int                      `json:"snapshotCount"`
	AutoSnapshotPolicyId string                   `json:"autoSnapshotPolicyId"`
	Encrypted            bool                     `json:"encrypted"`
	Tags                 []model.TagModel         `json:"tags"`
	Attachments          []VolumeAttachmentsModel `json:"volumeAttachments"`
}

type VolumeAttachmentsModel struct {
	InstanceId string `json:"instanceId"`
	Device     string `json:"device"`
	AttachTime string `json:"attachTime"`
}

type VolumeAttachmentModel struct {
	VolumeId   string `json:"volumeId"`
	InstanceId string `json:"instanceId"`
	Device     string `json:"device"`
	Serial     string `json:"serial"`
}

type AttachVolumeResult struct {
	VolumeAttachment *VolumeAttachmentModel `json:"volumeAttachment"`
	WarningList      []string               `json:"warningList"`
}

type CreateCDSVolumeArgs struct {
	Name               string               `json:"name,omitempty"`
	Description        string               `json:"description,omitempty"`
	SnapshotId         string               `json:"snapshotId,omitempty"`
	ZoneName           string               `json:"zoneName,omitempty"`
	PurchaseCount      int                  `json:"purchaseCount,omitempty"`
	CdsSizeInGB        int                  `json:"cdsSizeInGB,omitempty"`
	StorageType        StorageType          `json:"storageType,omitempty"`
	Billing            *Billing             `json:"billing"`
	EncryptKey         string               `json:"encryptKey"`
	RenewTimeUnit      string               `json:"renewTimeUnit"`
	RenewTime          int                  `json:"renewTime"`
	InstanceId         string               `json:"instanceId"`
	ClusterId          string               `json:"clusterId"`
	Tags               []model.TagModel     `json:"tags"`
	AutoSnapshotPolicy []AutoSnapshotPolicy `json:"autoSnapshotPolicy"`
	ClientToken        string               `json:"-"`
	ChargeType         string               `json:"chargeType"`
}

type AutoSnapshotPolicy struct {
	AutoSnapshotPolicyId string `json:"autoSnapshotPolicyId"`
}

type CreateCDSVolumeV3Args struct {
	VolumeName           string        `json:"volumeName,omitempty"`
	Description          string        `json:"description,omitempty"`
	SnapshotId           string        `json:"snapshotId,omitempty"`
	ZoneName             string        `json:"zoneName,omitempty"`
	PurchaseCount        int           `json:"purchaseCount,omitempty"`
	VolumeSize           int           `json:"volumeSizeInGB,omitempty"`
	StorageType          StorageTypeV3 `json:"storageType,omitempty"`
	Billing              *Billing      `json:"billing"`
	EncryptKey           string        `json:"encryptKey"`
	AutoSnapshotPolicyId string        `json:"autoSnapshotPolicyId"`
	InstanceId           string        `json:"instanceId"`
	RenewTimeUnit        string        `json:"renewTimeUnit"`
	RenewTime            int           `json:"renewTime"`
	ClientToken          string        `json:"-"`
	ChargeType           string        `json:"chargeType"`
}

type CreateCDSVolumeResult struct {
	VolumeIds   []string `json:"volumeIds"`
	WarningList []string `json:"warningList"`
}

type GetVolumeDetailResult struct {
	Volume *VolumeModel `json:"volume"`
}

type GetVolumeDetailResultV3 struct {
	Volume *VolumeModelV3 `json:"volume"`
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
	NewCdsSizeInGB int         `json:"newCdsSizeInGB,omitempty"`
	NewVolumeType  StorageType `json:"newVolumeType,omitempty"`
	ClientToken    string      `json:"-"`
}

type ResizeCDSVolumeResult struct {
	WarningList []string `json:"warningList"`
}

type RollbackCSDVolumeArgs struct {
	SnapshotId string `json:"snapshotId"`
}

type ListCDSVolumeArgs struct {
	MaxKeys      int    `json:"maxKeys"`
	InstanceId   string `json:"instanceId"`
	ZoneName     string `json:"zoneName"`
	Marker       string `json:"marker"`
	ClusterId    string `json:"clusterId"`
	ChargeFilter string `json:"chargeFilter"`
	UsageFilter  string `json:"usageFilter"`
	Name         string `json:"name"`
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
	UcAccount string `json:"ucAccount,omitempty"`
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
	OsLang         string          `json:"osLang"`
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
	Encrypted      bool            `json:"encrypted"`
	MinDiskGb      int             `json:"minDiskGb"`
	DiskSize       int             `json:"diskSize"`
	Snapshots      []SnapshotModel `json:"snapshots"`
}

type GetImageDetailResult struct {
	Image *ImageModel `json:"image"`
}

type RemoteCopyImageArgs struct {
	Name       string   `json:"name,omitempty"`
	DestRegion []string `json:"destRegion"`
}

type RemoteCopyImageResult struct {
	RemoteCopyImages []RemoteCopyImageModel `json:"result"`
}

type RemoteCopyImageModel struct {
	Region  string `json:"region"`
	ImageId string `json:"imageId"`
	ErrMsg  string `json:"errMsg"`
	Code    string `json:"code"`
}

type CreateImageArgs struct {
	InstanceId  string `json:"instanceId,omitempty"`
	SnapshotId  string `json:"snapshotId,omitempty"`
	ImageName   string `json:"imageName"`
	IsRelateCds bool   `json:"relateCds"`
	EncryptKey  string `json:"encryptKey"`
	ClientToken string `json:"-"`
}

type ListImageArgs struct {
	Marker    string
	MaxKeys   int
	ImageType string
	ImageName string
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
	ClientToken  string           `json:"-"`
	VolumeId     string           `json:"volumeId"`
	SnapshotName string           `json:"snapshotName"`
	Description  string           `json:"desc,omitempty"`
	Tags         []model.TagModel `json:"tags"`
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
	VolumeId    string           `json:"volumeId"`
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
	InstanceCount string              `json:"instanceCount"`
	Strategy      string              `json:"strategy"`
	InstanceList  []AzIntstanceStatis `json:"azIntstanceStatisList"`
	Name          string              `json:"name"`
	Desc          string              `json:"desc"`
	DeploySetId   string              `json:"deploysetId"`
	Concurrency   int                 `json:"concurrency"`
}

type DeploySetResult struct {
	Strategy     string                    `json:"strategy"`
	Name         string                    `json:"name"`
	Desc         string                    `json:"desc"`
	DeploySetId  string                    `json:"shortId"`
	Concurrency  int                       `json:"concurrency"`
	InstanceList []AzIntstanceStatisDetail `json:"azIntstanceStatisList"`
}

type UpdateInstanceDeployArgs struct {
	ClientToken  string   `json:"-"`
	InstanceId   string   `json:"instanceId,omitempty"`
	DeploySetIds []string `json:"deploysetIdList,omitempty"`
}

type DelInstanceDeployArgs struct {
	ClientToken string   `json:"-"`
	InstanceIds []string `json:"instanceIdList,omitempty"`
	DeploySetId string   `json:"deployId,omitempty"`
}

type AzIntstanceStatisDetail struct {
	ZoneName       string   `json:"zoneName"`
	Count          int      `json:"instanceCount"`
	BccCount       int      `json:"bccInstanceCnt"`
	BbcCount       int      `json:"bbcInstanceCnt"`
	Total          int      `json:"instanceTotal"`
	InstanceIds    []string `json:"instanceIds"`
	BccInstanceIds []string `json:"bccInstanceIds"`
	BbcInstanceIds []string `json:"bbcInstanceIds"`
}

type AzIntstanceStatis struct {
	ZoneName string `json:"zoneName"`
	Count    int    `json:"instanceCount"`
	BbcCount int    `json:"bbcInstanceCnt"`
	BccCount int    `json:"bccInstanceCnt"`
	Total    int    `json:"instanceTotal"`
}

type GetDeploySetResult struct {
	DeploySetModel
}

type RebuildBatchInstanceArgs struct {
	ImageId           string   `json:"imageId"`
	AdminPass         string   `json:"adminPass"`
	KeypairId         string   `json:"keypairId"`
	InstanceIds       []string `json:"instanceIds"`
	IsOpenHostEye     bool     `json:"isOpenHostEye"`
	IsPreserveData    bool     `json:"isPreserveData"`
	RaidId            string   `json:"raidId,omitempty"`
	SysRootSize       int      `json:"sysRootSize,omitempty"`
	RootPartitionType string   `json:"rootPartitionType,omitempty"`
	DataPartitionType string   `json:"dataPartitionType,omitempty"`
}

type RebuildBatchInstanceArgsV2 struct {
	ImageId           string   `json:"imageId"`
	AdminPass         string   `json:"adminPass"`
	KeypairId         string   `json:"keypairId"`
	InstanceIds       []string `json:"instanceIds"`
	IsOpenHostEye     *bool    `json:"isOpenHostEye"`
	IsPreserveData    *bool    `json:"isPreserveData"`
	RaidId            string   `json:"raidId,omitempty"`
	SysRootSize       int      `json:"sysRootSize,omitempty"`
	RootPartitionType string   `json:"rootPartitionType,omitempty"`
	DataPartitionType string   `json:"dataPartitionType,omitempty"`
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
	EbcResources EbcResources `json:"ebcResources"`
}

type BccResources struct {
	FlavorGroups []FlavorGroup `json:"flavorGroups"`
}

type FlavorGroup struct {
	GroupId string      `json:"groupId"`
	Flavors []BccFlavor `json:"flavors"`
}

type BccFlavor struct {
	CpuCount            int    `json:"cpuCount"`
	MemoryCapacityInGB  int    `json:"memoryCapacityInGB"`
	EphemeralDiskInGb   int    `json:"ephemeralDiskInGb"`
	EphemeralDiskCount  int    `json:"ephemeralDiskCount"`
	EphemeralDiskType   string `json:"ephemeralDiskType"`
	GpuCardType         string `json:"gpuCardType"`
	GpuCardCount        int    `json:"gpuCardCount"`
	FpgaCardType        string `json:"fpgaCardType"`
	FpgaCardCount       int    `json:"fpgaCardCount"`
	ProductType         string `json:"productType"`
	Spec                string `json:"spec"`
	SpecId              string `json:"specId"`
	CpuModel            string `json:"cpuModel"`
	CpuGHz              string `json:"cpuGHz"`
	NetworkBandwidth    string `json:"networkBandwidth"`
	NetworkPackage      string `json:"networkPackage"`
	NetEthQueueCount    string `json:"netEthQueueCount"`
	NetEthMaxQueueCount string `json:"netEthMaxQueueCount"`
}

type EbcResources struct {
	FlavorGroups []EbcFlavorGroup `json:"flavorGroups"`
}

type EbcFlavorGroup struct {
	GroupId string      `json:"groupId"`
	Flavors []EbcFlavor `json:"flavors"`
}

type EbcFlavor struct {
	CpuCount           int    `json:"cpuCount"`
	MemoryCapacityInGB int    `json:"memoryCapacityInGB"`
	EphemeralDiskInGb  int    `json:"ephemeralDiskInGb"`
	EphemeralDiskCount string `json:"ephemeralDiskCount"`
	EphemeralDiskType  string `json:"ephemeralDiskType"`
	GpuCardType        string `json:"gpuCardType"`
	GpuCardCount       string `json:"gpuCardCount"`
	FpgaCardType       string `json:"fpgaCardType"`
	FpgaCardCount      string `json:"fpgaCardCount"`
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
	Name    string `json:"name,omitempty"`
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
	RenewCds      bool   `json:"renewCds"`
	RenewEip      bool   `json:"renewEip"`
}

type BccDeleteAutoRenewArgs struct {
	InstanceId string `json:"instanceId"`
	RenewCds   bool   `json:"renewCds"`
	RenewEip   bool   `json:"renewEip"`
}

type DeleteInstanceIngorePaymentArgs struct {
	InstanceId            string `json:"instanceId"`
	RelatedReleaseFlag    bool   `json:"relatedReleaseFlag"`
	DeleteCdsSnapshotFlag bool   `json:"deleteCdsSnapshotFlag"`
	DeleteRelatedEnisFlag bool   `json:"deleteRelatedEnisFlag"`
	DeleteImmediate       bool   `json:"deleteImmediate"`
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
	FailResources    *DeleteInstanceModel `json:"failResources"`
}

type RecoveryInstanceArgs struct {
	InstanceIds []RecoveryInstanceModel `json:"instanceIds"`
}

type RecoveryInstanceModel struct {
	InstanceId string `json:"instanceId"`
}

type ListInstanceByInstanceIdArgs struct {
	Marker      string
	MaxKeys     int
	InstanceIds []string `json:"instanceIds"`
}

type ListInstancesResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Instances   []InstanceModel `json:"instances"`
}

type VolumePrepayDeleteRequestArgs struct {
	VolumeId           string `json:"volumeId"`
	RelatedReleaseFlag bool   `json:"relatedReleaseFlag"`
	DeleteImmediate    bool   `json:"deleteImmediate"`
}

type VolumeDeleteResultResponse struct {
	SuccessResources VolumeDeleteResultModel `json:"successResources"`
	FailResources    VolumeDeleteResultModel `json:"failResources"`
}

type VolumeDeleteResultModel struct {
	VolumeIds []string `json:"volumeIds"`
}

type DeletionProtectionArgs struct {
	DeletionProtection int `json:"deletionProtection"`
}

type BatchDeleteInstanceWithRelateResourceArgs struct {
	RelatedReleaseFlag    bool     `json:"relatedReleaseFlag"`
	DeleteCdsSnapshotFlag bool     `json:"deleteCdsSnapshotFlag"`
	BccRecycleFlag        bool     `json:"bccRecycleFlag"`
	DeleteRelatedEnisFlag bool     `json:"deleteRelatedEnisFlag"`
	InstanceIds           []string `json:"instanceIds"`
}

type BatchStartInstanceArgs struct {
	InstanceIds []string `json:"instanceIds"`
}

type BatchStopInstanceArgs struct {
	ForceStop        bool     `json:"forceStop"`
	StopWithNoCharge bool     `json:"stopWithNoCharge"`
	InstanceIds      []string `json:"instanceIds"`
}

type ListInstanceTypeArgs struct {
	ZoneName string `json:"zoneName"`
}

type ListInstanceTypeResults struct {
	ZoneInstanceTypes []ZoneInstanceTypes `json:"zoneInstanceTypes"`
}

type ZoneInstanceTypes struct {
	ZoneName      string   `json:"zoneName"`
	InstanceTypes []string `json:"instanceTypes"`
}

type ListIdMappingArgs struct {
	IdType      string   `json:"idType"`
	ObjectType  string   `json:"objectType"`
	InstanceIds []string `json:"instanceIds"`
}

type ListIdMappingResults struct {
	IdMapping []IdMapping `json:"mappings"`
}

type IdMapping struct {
	Uuid string `json:"uuid"`
	Id   string `json:"id"`
}

type BatchResizeInstanceArgs struct {
	Spec           string   `json:"spec"`
	InstanceIdList []string `json:"instanceIdList"`
}

type BatchResizeInstanceResults struct {
	OrderUuidResults []string `json:"orderUuidResults"`
}

// UpdateSecurityGroupRuleArgs defines the structure of input parameters for the UpdateSecurityGroupRule api
type UpdateSecurityGroupRuleArgs struct {
	SgVersion           int64   `json:"sgVersion,omitempty"`
	SecurityGroupRuleId string  `json:"securityGroupRuleId"`
	Remark              *string `json:"remark,omitempty"`
	PortRange           *string `json:"portRange,omitempty"`
	SourceIp            *string `json:"sourceIp,omitempty"`
	SourceGroupId       *string `json:"sourceGroupId,omitempty"`
	DestIp              *string `json:"destIp,omitempty"`
	DestGroupId         *string `json:"destGroupId,omitempty"`
	Protocol            *string `json:"protocol,omitempty"`
}

// DeleteSecurityGroupRuleArgs defines the structure of input parameters for the DeleteSecurityGroupRule api
type DeleteSecurityGroupRuleArgs struct {
	SgVersion           int64  `json:"sgVersion,omitempty"`
	SecurityGroupRuleId string `json:"securityGroupRuleId"`
}

type GetInstanceDeleteProgressArgs struct {
	InstanceIds []string `json:"instanceIds"`
}

type Tag struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

type TagVolumeArgs struct {
	ChangeTags  []Tag `json:"changeTags"`
	RelationTag bool  `json:"relationTag"`
}

type Volume struct {
	VolumeId string `json:"volumeId"`
	SizeInGB int    `json:"sizeInGb"`
}

type ListAvailableResizeSpecsArgs struct {
	Spec           string   `json:"spec"`
	SpecId         string   `json:"specId"`
	Zone           string   `json:"zone"`
	InstanceIdList []string `json:"instanceIdList"`
}

type ListAvailableResizeSpecResults struct {
	SpecList []string `json:"specList"`
}

type BatchChangeInstanceToPrepayArgs struct {
	Config []PrepayConfig `json:"config"`
}

type PrepayConfig struct {
	InstanceId  string   `json:"instanceId"`
	Duration    int      `json:"duration"`
	RelationCds bool     `json:"relationCds"`
	CdsList     []string `json:"cdsList"`
	AutoPay     bool     `json:"autoPay"`
}

type BatchChangeInstanceToPostpayArgs struct {
	Config []PostpayConfig `json:"config"`
}

type PostpayConfig struct {
	InstanceId    string   `json:"instanceId"`
	RelationCds   bool     `json:"relationCds"`
	CdsList       []string `json:"cdsList"`
	EffectiveType string   `json:"effectiveType"`
}

type BatchChangeInstanceBillingMethodResult struct {
	OrderId string `json:"orderId"`
}

type Role struct {
	RoleName string `json:"roleName"`
}

type ListInstanceRolesResult struct {
	Roles []Role `json:"roles"`
}

type BindInstanceRoleResult struct {
	FailInstances            []FailInstances            `json:"failInstances"`
	InstanceRoleAssociations []InstanceRoleAssociations `json:"instanceRoleAssociations"`
}

type FailInstances struct {
	InstanceId  string `json:"instanceId"`
	FailMessage string `json:"failMessage"`
}

type InstanceRoleAssociations struct {
	InstanceID string `json:"instanceId"`
}

type BindInstanceRoleArgs struct {
	RoleName  string      `json:"roleName"`
	Instances []Instances `json:"instances"`
}

type Instances struct {
	InstanceId string `json:"instanceId"`
}

type UnBindInstanceRoleArgs struct {
	RoleName  string      `json:"roleName"`
	Instances []Instances `json:"instances"`
}

type UnBindInstanceRoleResult struct {
	FailInstances            []FailInstances            `json:"failInstances"`
	InstanceRoleAssociations []InstanceRoleAssociations `json:"instanceRoleAssociations"`
}

type DeleteIpv6Args struct {
	InstanceId string `json:"instanceId"`
	Reboot     bool   `json:"reboot"`
}

type AddIpv6Args struct {
	InstanceId  string `json:"instanceId"`
	Ipv6Address string `json:"ipv6Address"`
	Reboot      bool   `json:"reboot"`
}

type AddIpv6Result struct {
	Ipv6Address string `json:"ipv6Address"`
}

type RemoteCopySnapshotArgs struct {
	ClientToken     string           `json:"-"`
	DestRegionInfos []DestRegionInfo `json:"destRegionInfos"`
}

type DestRegionInfo struct {
	Name       string `json:"name"`
	DestRegion string `json:"destRegion"`
}

type RemoteCopySnapshotResult struct {
	Result []RemoteCopySnapshotResultItem `json:"result"`
}

type RemoteCopySnapshotResultItem struct {
	Region     string `json:"region"`
	SnapshotID string `json:"snapshotId"`
}

type ImportCustomImageArgs struct {
	OsName    string `json:"osName"`
	OsArch    string `json:"osArch"`
	OsType    string `json:"osType"`
	OsVersion string `json:"osVersion"`
	Name      string `json:"name"`
	BosURL    string `json:"bosUrl"`
}

type ImportCustomImageResult struct {
	Id string `json:"id"`
}

type GetAvailableImagesBySpecArg struct {
	Marker  string `json:"marker"`
	MaxKeys int    `json:"maxKeys"`
	Spec    string `json:"spec"`
	OsName  string `json:"osName"`
}

type GetAvailableImagesBySpecResult struct {
	IsTruncated bool     `json:"isTruncated"`
	Marker      string   `json:"marker"`
	MaxKeys     int      `json:"maxKeys"`
	NextMarker  string   `json:"nextMarker"`
	Images      ImageArg `json:"images"`
}

type BatchRefundResourceArg struct {
	InstanceIds           []string `json:"instanceIds"`
	RelatedReleaseFlag    bool     `json:"relatedReleaseFlag"`
	DeleteCdsSnapshotFlag bool     `json:"deleteCdsSnapshotFlag"`
	DeleteRelatedEnisFlag bool     `json:"deleteRelatedEnisFlag"`
}

type BatchRefundResourceResult struct {
	FailedInstanceIds []string `json:"failedInstanceIds"`
}

type ImageArg []struct {
	ImageID      string `json:"imageId"`
	ImageName    string `json:"imageName"`
	OsType       string `json:"osType"`
	OsVersion    string `json:"osVersion"`
	OsArch       string `json:"osArch"`
	OsName       string `json:"osName"`
	OsLang       string `json:"osLang"`
	MinSizeInGiB int    `json:"minSizeInGiB"`
}

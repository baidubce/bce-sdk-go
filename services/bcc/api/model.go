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
)

type StorageType string

const (
	StorageTypeStd1     StorageType = "std1"
	StorageTypeHP1      StorageType = "hp1"
	StorageTypeCloudHP1 StorageType = "cloud_hp1"
	StorageTypeLocal    StorageType = "local"
	StorageTypeSATA     StorageType = "sata"
	StorageTypeSSD      StorageType = "ssd"
)

type PaymentTimingType string

const (
	PaymentTimingPrePaid  PaymentTimingType = "Prepaid"
	PaymentTimingPostPaid PaymentTimingType = "Postpaid"
)

// Instance define instance model
type InstanceModel struct {
	InstanceId            string           `json:"id"`
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
	ClientToken           string           `json:"-"`
}

type CreateInstanceResult struct {
	InstanceIds []string `json:"instanceIds"`
}

type ListInstanceArgs struct {
	Marker          string
	MaxKeys         int
	InternalIp      string
	DedicatedHostId string
	ZoneName        string
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
	ClientToken        string          `json:"-"`
}

type RebuildInstanceArgs struct {
	ImageId   string `json:"imageId"`
	AdminPass string `json:"adminPass"`
}

type StopInstanceArgs struct {
	ForceStop bool `json:"forceStop"`
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

type PurchaseReservedArgs struct {
	Billing     Billing `json:"billing"`
	ClientToken string  `json:"-"`
}

type DeleteInstanceWithRelateResourceArgs struct {
	RelatedReleaseFlag    bool `json:"relatedReleaseFlag"`
	DeleteCdsSnapshotFlag bool `json:"deleteCdsSnapshotFlag"`
}

type InstanceChangeSubnetArgs struct {
	InstanceId string `json:"instanceId"`
	SubnetId   string `json:"subnetId"`
	Reboot     bool   `json:"reboot"`
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
	ClientToken   string      `json:"-"`
}

type CreateCDSVolumeResult struct {
	VolumeIds []string `json:"volumeIds"`
}

type GetVolumeDetailResult struct {
	Volume *VolumeModel `json:"volume"`
}

type AttachVolumeArgs struct {
	InstanceId string `json:"instanceId"`
}

type ResizeCSDVolumeArgs struct {
	NewCdsSizeInGB int    `json:"newCdsSizeInGB"`
	ClientToken    string `json:"-"`
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
	ClientToken string `json:"-"`
}

type ListImageArgs struct {
	Marker    string
	MaxKeys   int
	ImageType string
}

type OsModel struct {
	OsVersion  string `json:"osVersion"`
	OsType     string `json:"osType"`
	InstanceId string `json:"instanceId"`
	OsArch     string `json:"osArch"`
	OsName     string `json:"osName"`
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
}

type ListSnapshotResult struct {
	Marker      string          `json:"marker"`
	IsTruncated bool            `json:"isTruncated"`
	NextMarker  string          `json:"nextMarker"`
	MaxKeys     int             `json:"maxKeys"`
	Snapshots   []SnapshotModel `json:"snapshots"`
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

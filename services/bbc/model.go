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

import "github.com/baidubce/bce-sdk-go/model"

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
	FlavorId         string   `json:"flavorId"`
	ImageId          string   `json:"imageId"`
	RaidId           string   `json:"raidId"`
	RootDiskSizeInGb int      `json:"rootDiskSizeInGb"`
	PurchaseCount    int      `json:"purchaseCount"`
	ZoneName         string   `json:"zoneName"`
	SubnetId         string   `json:"subnetId"`
	Billing          Billing  `json:"billing"`
	Name             string   `json:"name,omitempty"`
	AdminPass        string   `json:"adminPass,omitempty"`
	DeploySetId      string   `json:"deploySetId,omitempty"`
	ClientToken      string   `json:"-"`
	SecurityGroupId  string   `json:"securityGroupId,omitempty"`
	InternalIps      []string `json:"internalIps,omitempty"`
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
	IsPreserveData bool   `json:"isPreserveData,omitempty"`
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
	InstanceId string   `json:"instanceId"`
	PrivateIps []string `json:"privateIps"`
}

type BatchDelIpArgs struct {
	InstanceId string   `json:"instanceId"`
	PrivateIps []string `json:"privateIps"`
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

type GetImageDetailResult struct {
	Image *ImageModel `json:"image"`
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
	RaidId       string `json:"raidId"`
	Raid         string `json:"raid"`
	SysSwapSize  int    `json:"sysSwapSize"`
	SysRootSize  int    `json:"sysRootSize"`
	SysHomeSize  int    `json:"sysHomeSize"`
	SysDiskSize  int    `json:"sysDiskSize"`
	DataDiskSize int    `json:"dataDiskSize"`
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

type CreateDeploySetResult struct {
	DeploySetId string `json:"deploySetId"`
}

type ListDeploySetsResult struct {
	DeploySetList []DeploySetModel `json:"deploySetList"`
}

type DeploySetModel struct {
	Strategy     string   `json:"strategy"`
	InstanceList []string `json:"instanceList"`
	Concurrency  int      `json:"concurrency"`
	DeploySetId  string   `json:"deploySetId"`
}

type GetDeploySetResult struct {
	DeploySetModel
}

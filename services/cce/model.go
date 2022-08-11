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

package cce

import (
	"time"
)

type ProductType string

const (
	ProductTypePostpay ProductType = "postpay"
	ProductTypePrepay  ProductType = "prepay"
)

type ClusterStatus string

const (
	ClusterStatusRunning             ClusterStatus = "RUNNING"
	ClusterStatusCreating            ClusterStatus = "CREATING"
	ClusterStatusCreateFailed        ClusterStatus = "CREATE_FAILED"
	ClusterStatusDeleting            ClusterStatus = "DELETING"
	ClusterStatusDeletingFailed      ClusterStatus = "DELETE_FAILED"
	ClusterStatusMasterUpgrading     ClusterStatus = "MASTER_UPGRADING"
	ClusterStatusMasterUpgradeFailed ClusterStatus = "MASTER_UPGRADE_FAILED"
	ClusterStatusError               ClusterStatus = "ERROR"
	ClusterStatusDeleted             ClusterStatus = "DELETED"
)

type SimpleNode struct {
	InstanceShortId string        `json:"instanceShortId"`
	InstanceUuid    string        `json:"instanceUuid"`
	InstanceName    string        `json:"instanceName"`
	ClusterUuid     string        `json:"clusterUuid"`
	Status          ClusterStatus `json:"status"`
}

type Cluster struct {
	ClusterUuid       string        `json:"clusterUuid"`
	ClusterName       string        `json:"clusterName"`
	SlaveVmCount      int           `json:"slaveVmCount"`
	MasterVmCount     int           `json:"masterVmCount"`
	ContainerNet      string        `json:"containerNet"`
	Status            ClusterStatus `json:"status"`
	Region            string        `json:"region"`
	CreateTime        time.Time     `json:"createTime"`
	DeleteTime        time.Time     `json:"deleteTime"`
	AllInstanceNormal bool          `json:"allInstanceNormal"`
	InstanceList      []SimpleNode  `json:"instanceList"`
	DccUuid           string        `json:"dccUuid"`
	HasPrepay         bool          `json:"hasPrepay"`
	VpcId             string        `json:"vpcId"`
	InstanceMode      string        `json:"instanceMode"`
	MasterExposed     bool          `json:"masterExposed"`
}

type ImageType string

const (
	ImageTypeCommon    ImageType = "common"
	ImageTypeCustom    ImageType = "custom"
	ImageTypeGpu       ImageType = "gpuBccImage"
	ImageTypeGpuCustom ImageType = "gpuBccCustom"
	ImageTypeSharing   ImageType = "sharing"
)

type ServiceType string

const (
	ServiceTypeBCC ServiceType = "BCC"
	ServiceTypeCDS ServiceType = "CDS"
	ServiceTypeEIP ServiceType = "EIP"
)

type InstanceType string

const (
	InstanceTypeG1     InstanceType = "0"
	InstanceTypeDCC    InstanceType = "1"
	InstanceTypeBCC    InstanceType = "2"
	InstanceTypeC1     InstanceType = "4"
	InstanceTypeG2     InstanceType = "7"
	InstanceTypeGPU    InstanceType = "9"
	InstanceTypeG3     InstanceType = "10"
	InstanceTypeC2     InstanceType = "11"
	InstanceTypeG4     InstanceType = "13"
	InstanceTypeVGPU   InstanceType = "15"
	InstanceTypeKunlun InstanceType = "25"
)

type BccConfig struct {
	Name                string       `json:"name,omitempty"`
	KeypairId           string       `json:"keypairId,omitempty"`
	ProductType         ProductType  `json:"productType"`
	LogicalZone         string       `json:"logicalZone,omitempty"`
	InstanceType        InstanceType `json:"instanceType"`
	GpuCard             string       `json:"gpuCard,omitempty"`
	GpuCount            int          `json:"gpuCount"`
	Cpu                 int          `json:"cpu"`
	Memory              int          `json:"memory"`
	ImageType           ImageType    `json:"imageType"`
	SubnetUuid          string       `json:"subnetUuid"`
	SecurityGroupId     string       `json:"securityGroupId"`
	AdminPass           string       `json:"adminPass,omitempty"`
	PurchaseLength      int          `json:"purchaseLength,omitempty"`
	PurchaseNum         int          `json:"purchaseNum"`
	RootDiskSizeInGb    int          `json:"rootDiskSizeInGb,omitempty"`
	RootDiskStorageType VolumeType   `json:"rootDiskStorageType,omitempty"`
	AutoRenewTimeUnit   string       `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime       int          `json:"autoRenewTime,omitempty"`
	AutoRenew           bool         `json:"autoRenew,omitempty"`
	ImageId             string       `json:"imageId"`
	ServiceType         ServiceType  `json:"serviceType"`
}

type VolumeType string

const (
	VolumeTypeSata       VolumeType = "sata"
	VolumeTypeSsd        VolumeType = "ssd"
	VolumeTypePremiumSsd VolumeType = "premium_ssd"
)

type DiskSizeConfig struct {
	Size       string     `json:"size"`
	VolumeType VolumeType `json:"volumeType"`
	SnapshotId string     `json:"snapshotId"`
}

type CdsConfig struct {
	ProductType       ProductType      `json:"productType"`
	LogicalZone       string           `json:"logicalZone"`
	PurchaseNum       int              `json:"purchaseNum"`
	PurchaseLength    int              `json:"purchaseLength,omitempty"`
	AutoRenewTimeUnit string           `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime     int              `json:"autoRenewTime,omitempty"`
	CdsDiskSize       []DiskSizeConfig `json:"cdsDiskSize"`
	ServiceType       ServiceType      `json:"serviceType"`
}

type EipType string

const (
	EipTypeBandwidth EipType = "bandwidth"
	EipTypeNetraffic EipType = "netraffic"
)

type EipConfig struct {
	ProductType       ProductType `json:"productType"`
	BandwidthInMbps   int         `json:"bandwidthInMbps"`
	SubProductType    EipType     `json:"subProductType"`
	PurchaseNum       int         `json:"purchaseNum"`
	PurchaseLength    int         `json:"purchaseLength,omitempty"`
	AutoRenewTime     int         `json:"autoRenewTime,omitempty"`
	AutoRenewTimeUnit string      `json:"autoRenewTimeUnit,omitempty"`
	Name              string      `json:"name,omitempty"`
	ServiceType       ServiceType `json:"serviceType"`
}

type Item struct {
	Config interface{} `json:"config"`
}

type BaseCreateOrderRequestVo struct {
	Items []Item `json:"items"`
}

type CdsPreMountInfo struct {
	MountPath string           `json:"mountPath"`
	CdsConfig []DiskSizeConfig `json:"cdsConfig"`
}

type CniMode string

const (
	CniModeKubenet CniMode = "kubenet"
	CniModeCni     CniMode = "cni"
)

type CniType string

const (
	CniTypeEmpty                 CniType = ""
	CniTypeRouteVeth             CniType = "VPC_ROUTE_VETH"
	CniTypeRouteIpvlan           CniType = "VPC_ROUTE_IPVLAN"
	CniTypeRouteAutoDetect       CniType = "VPC_ROUTE_AUTODETECT"
	CniTypeSecondaryIpVeth       CniType = "VPC_SECONDARY_IP_VETH"
	CniTypeSecondaryIpIpvlan     CniType = "VPC_SECONDARY_IP_IPVLAN"
	CniTypeSecondaryIpAutoDetect CniType = "VPC_SECONDARY_IP_AUTODETECT"
)

type DNSMode string

const (
	DNSModeKubeDNS DNSMode = "kubeDNS"
	DNSModeCoreDNS DNSMode = "coreDNS"
)

type KubeProxyMode string

const (
	KubeProxyModeIptables KubeProxyMode = "iptables"
	KubeProxyModeIpvs     KubeProxyMode = "ipvs"
)

type AdvancedOptions struct {
	KubeProxyMode         KubeProxyMode `json:"kubeProxyMode,omitempty"`
	SecureContainerEnable bool          `json:"secureContainerEnable,omitempty"`
	SetOSSecurity         bool          `json:"setOSSecurity,omitempty"`
	CniMode               CniMode       `json:"cniMode,omitempty"`
	CniType               CniType       `json:"cniType,omitempty"`
	DnsMode               DNSMode       `json:"dnsMode,omitempty"`
	MaxPodNum             int           `json:"maxPodNum,omitempty"`
}

type NodeInfo struct {
	InstanceId string `json:"instanceId"`
}

type Node struct {
	InstanceShortId string    `json:"instanceShortId"`
	InstanceUuid    string    `json:"instanceUuid"`
	InstanceName    string    `json:"instanceName"`
	ClusterUuid     string    `json:"clusterUuid"`
	AvailableZone   string    `json:"availableZone"`
	VpcId           string    `json:"vpcId"`
	VpcCidr         string    `json:"vpcCidr"`
	SubnetId        string    `json:"subnetId"`
	SubnetType      string    `json:"subnetType"`
	Eip             string    `json:"eip"`
	EipBandwidth    int       `json:"eipBandwidth"`
	Cpu             int       `json:"cpu"`
	Memory          int       `json:"memory"`
	DiskSize        int       `json:"diskSize"`
	SysDisk         int       `json:"sysDisk"`
	InstanceType    string    `json:"instanceType"`
	Blb             string    `json:"blb"`
	FloatingIp      string    `json:"floatingIp"`
	FixIp           string    `json:"fixIp"`
	CreateTime      time.Time `json:"createTime"`
	DeleteTime      time.Time `json:"deleteTime"`
	Status          string    `json:"status"`
	ExpireTime      time.Time `json:"expireTime"`
	PaymentMethod   string    `json:"paymentMethod"`
	RuntimeVersion  string    `json:"runtimeVersion"`
}

type CceNodeInfo struct {
	InstanceId string `json:"instanceId"`
	AdminPass  string `json:"adminPass,omitempty"`
}

type ServerForDisplay struct {
	InstanceId string `json:"instanceId"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Payment    string `json:"payment"`
	InternalIp string `json:"internalIp"`
}

type DeployMode string

const (
	DeployModeBcc DeployMode = "BCC"
	DeployModeDcc DeployMode = "DCC"
)

type CreateClusterArgs struct {
	ClusterName        string                    `json:"clusterName"`
	Version            string                    `json:"version"`
	MainAvailableZone  string                    `json:"mainAvailableZone"`
	ContainerNet       string                    `json:"containerNet"`
	AdvancedOptions    *AdvancedOptions          `json:"advancedOptions,omitempty"`
	CdsPreMountInfo    *CdsPreMountInfo          `json:"cdsPreMountInfo,omitempty"`
	Comment            string                    `json:"comment,omitempty"`
	DeployMode         DeployMode                `json:"deployMode"`
	DccUuid            string                    `json:"dccUuid,omitempty"`
	MasterExposed      bool                      `json:"masterExposed,omitempty"`
	OrderContent       *BaseCreateOrderRequestVo `json:"orderContent"`
	MasterOrderContent *BaseCreateOrderRequestVo `json:"masterOrderContent,omitempty"`
}

type CreateClusterResult struct {
	ClusterUuid string   `json:"clusterUuid"`
	OrderId     []string `json:"orderId"`
}

type ListClusterArgs struct {
	Status  ClusterStatus
	Marker  string
	MaxKeys int
}

type ListClusterResult struct {
	Marker      string    `json:"marker"`
	IsTruncated bool      `json:"isTruncated"`
	NextMarker  string    `json:"nextMarker"`
	MaxKeys     int       `json:"maxKeys"`
	Clusters    []Cluster `json:"clusters"`
}

type GetClusterResult struct {
	ClusterUuid           string            `json:"clusterUuid"`
	ClusterName           string            `json:"clusterName"`
	Version               string            `json:"version"`
	Region                string            `json:"region"`
	SlaveVmCount          int               `json:"slaveVmCount"`
	MasterVmCount         int               `json:"masterVmCount"`
	VpcId                 string            `json:"vpcId"`
	VpcUuid               string            `json:"vpcUuid"`
	VpcCidr               string            `json:"vpcCidr"`
	ZoneSubnetMap         map[string]string `json:"zoneSubnetMap"`
	ContainerNet          string            `json:"containerNet"`
	AdvancedOptions       *AdvancedOptions  `json:"advancedOptions"`
	Status                ClusterStatus     `json:"status"`
	CreateStartTime       time.Time         `json:"createStartTime"`
	DeleteTime            time.Time         `json:"deleteTime"`
	Comment               string            `json:"comment"`
	InstanceMode          string            `json:"instanceMode"`
	HasPrepay             bool              `json:"hasPrepay"`
	VpcName               string            `json:"vpcName"`
	SecureContainerEnable bool              `json:"secureContainerEnable"`
	MasterZoneSubnetMap   map[string]string `json:"masterZoneSubnetMap"`
	MasterExposed         bool              `json:"masterExposed"`
}

type DeleteClusterArgs struct {
	ClusterUuid  string
	DeleteEipCds bool
	DeleteSnap   bool
}

type ScalingUpArgs struct {
	ClusterUuid     string                    `json:"clusterUuid"`
	DccUuid         string                    `json:"dccUuid,omitempty"`
	CdsPreMountInfo *CdsPreMountInfo          `json:"cdsPreMountInfo,omitempty"`
	OrderContent    *BaseCreateOrderRequestVo `json:"orderContent"`
}

type ScalingUpResult struct {
	ClusterUuid string   `json:"clusterUuid"`
	OrderId     []string `json:"orderId"`
}

type ScalingDownArgs struct {
	ClusterUuid  string     `json:"clusterUuid"`
	DeleteEipCds bool       `json:"-"`
	DeleteSnap   bool       `json:"-"`
	NodeInfo     []NodeInfo `json:"nodeInfo"`
}

type ListNodeArgs struct {
	ClusterUuid string
	Marker      string
	MaxKeys     int
}

type ListNodeResult struct {
	Marker      string `json:"marker"`
	IsTruncated bool   `json:"isTruncated"`
	NextMarker  string `json:"nextMarker"`
	MaxKeys     int    `json:"maxKeys"`
	Nodes       []Node `json:"nodes"`
}

type ShiftInstanceType string

const (
	ShiftInstanceTypeBcc ShiftInstanceType = "BCC"
	ShiftInstanceTypeBBC ShiftInstanceType = "BBC"
)

type ShiftInNodeArgs struct {
	ClusterUuid  string            `json:"clusterUuid"`
	NeedRebuild  bool              `json:"needRebuild"`
	ImageId      string            `json:"imageId,omitempty"`
	AdminPass    string            `json:"adminPass"`
	InstanceType ShiftInstanceType `json:"instanceType"`
	NodeInfoList []CceNodeInfo     `json:"nodeInfoList"`
}

type ShiftOutNodeArgs struct {
	ClusterUuid  string        `json:"clusterUuid"`
	NodeInfoList []CceNodeInfo `json:"nodeInfoList"`
}

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

type KeywordType string

const (
	KeywordTypeName       KeywordType = "name"
	KeywordTypeInstanceId KeywordType = "instanceId"
)

type ListExistedNodeArgs struct {
	ClusterUuid  string            `json:"clusterUuid"`
	VpcId        string            `json:"vpcId,omitempty"`
	VpcCidr      string            `json:"vpcCidr,omitempty"`
	InstanceType ShiftInstanceType `json:"instanceType,omitempty"`
	BBCFlavorId  string            `json:"bbcFlavorId,omitempty"`
	KeywordType  KeywordType       `json:"keywordType,omitempty"`
	Keyword      string            `json:"keyword,omitempty"`
	OrderBy      string            `json:"orderBy,omitempty"`
	Order        Order             `json:"order,omitempty"`
	PageNo       int               `json:"pageNo,omitempty"`
	PageSize     int               `json:"pageSize,omitempty"`
}

type ListExistedNodeResult struct {
	ClusterUuid string             `json:"clusterUuid"`
	OrderBy     string             `json:"orderBy"`
	Order       string             `json:"order"`
	PageNo      int                `json:"pageNo"`
	PageSize    int                `json:"pageSize"`
	TotalCount  int                `json:"totalCount"`
	NodeList    []ServerForDisplay `json:"nodeList"`
}

type GetContainerNetArgs struct {
	VpcShortId string `json:"vpcShortId"`
	VpcCidr    string `json:"vpcCidr"`
	Size       int    `json:"size"`
}

type GetContainerNetResult struct {
	ContainerNet string `json:"containerNet"`
	Capacity     int    `json:"capacity"`
}

type KubeConfigType string

const (
	KubeConfigTypeDefault  KubeConfigType = "default"
	KubeConfigTypeInternal KubeConfigType = "internal"
	KubeConfigTypeIntraVpc KubeConfigType = "intraVpc"
)

type GetKubeConfigArgs struct {
	ClusterUuid string
	Type        KubeConfigType
}

type GetKubeConfigResult struct {
	Data string `json:"data"`
}

type ListVersionsResult struct {
	Data []string `json:"data"`
}

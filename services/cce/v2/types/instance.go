// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package types

import (
	bccapi "github.com/baidubce/bce-sdk-go/services/bcc/api"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 已有节点需要用户提供：ClusterRole 、短ID，密码，镜像ID,镜像类型, docker storage(可选); BBC要额外加preservedData、raidId、sysRootSize
type InstanceSpec struct {

	// 用于 CCE 唯一标识 Instance
	CCEInstanceID string `json:"cceInstanceID,omitempty"`
	InstanceName  string `json:"instanceName"`

	RuntimeType    RuntimeType `json:"runtimeType,omitempty"`
	RuntimeVersion string      `json:"runtimeVersion,omitempty"`

	ClusterID   string      `json:"clusterID,omitempty"`
	ClusterRole ClusterRole `json:"clusterRole,omitempty"`

	InstanceGroupID   string `json:"instanceGroupID,omitempty"`
	InstanceGroupName string `json:"instanceGroupName,omitempty"`

	// 初始化 DelProvider 使用
	MasterType MasterType `json:"masterType,omitempty"`

	// 是否为已有实例
	Existed       bool          `json:"existed,omitempty"`
	ExistedOption ExistedOption `json:"existedOption,omitempty"`

	// BCC, BBC, 裸金属
	MachineType MachineType `json:"machineType,omitempty"`
	// 机器规格: 普通一, 普通二 ...
	InstanceType bccapi.InstanceType `json:"instanceType"`
	// BBC 选项
	BBCOption *BBCOption `json:"bbcOption,omitempty"`

	// VPC 相关配置
	VPCConfig VPCConfig `json:"vpcConfig,omitempty"`

	// 集群规格相关配置
	InstanceResource InstanceResource `json:"instanceResource,omitempty"`

	// 优先使用 ImageID, 如果用户传入 InstanceOS 信息, 由 service 计算 ImageID
	ImageID    string     `json:"imageID,omitempty"`
	InstanceOS InstanceOS `json:"instanceOS,omitempty"`

	// EIP
	NeedEIP   bool `json:"needEIP,omitempty"`
	EIPOption *EIPOption `json:"eipOption,omitempty"`

	// AdminPassword
	AdminPassword string `json:"adminPassword,omitempty"`
	SSHKeyID      string `json:"sshKeyID,omitempty"`

	// Charging Type, 通常只支持后付费
	InstanceChargingType      bccapi.PaymentTimingType `json:"instanceChargingType,omitempty"` // 后付费或预付费
	InstancePreChargingOption InstancePreChargingOption `json:"instancePreChargingOption,omitempty"`

	// 删除节点选项
	DeleteOption *DeleteOption `json:"deleteOption,omitempty"`

	DeployCustomConfig DeployCustomConfig `json:"deployCustomConfig,omitempty"` // 部署相关高级配置

	Tags TagList `json:"tags,omitempty"`

	Labels InstanceLabels `json:"labels,omitempty"`
	Taints InstanceTaints `json:"taints,omitempty"`

	CCEInstancePriority int `json:"cceInstancePriority,omitempty"`
}

// VPCConfig 定义 Instance VPC
type VPCConfig struct {
	VPCID           string `json:"vpcID,omitempty"`
	VPCSubnetID     string `json:"vpcSubnetID,omitempty"`
	SecurityGroupID string `json:"securityGroupID,omitempty"`

	VPCSubnetType     vpc.SubnetType `json:"vpcSubnetType,omitempty"`
	VPCSubnetCIDR     string         `json:"vpcSubnetCIDR,omitempty"`
	VPCSubnetCIDRIPv6 string         `json:"vpcSubnetCIDRIPv6,omitempty"`

	AvailableZone AvailableZone `json:"availableZone,omitempty"`
}

// InstanceResource 定义 Instance CPU/MEM/Disk 配置
type InstanceResource struct {
	CPU int `json:"cpu,omitempty"` // unit: Core
	MEM int `json:"mem,omitempty"` // unit: GB

	NodeCPUQuota int `json:"nodeCPUQuota,omitempty"` // unit: Core
	NodeMEMQuota int `json:"nodeMEMQuota,omitempty"` // unit: GB

	// RootDisk
	RootDiskType bccapi.StorageType `json:"rootDiskType,omitempty"`
	RootDiskSize int             `json:"rootDiskSize,omitempty"` // unit: GB

	// GPU 机器必须指定, 其他机器不用
	LocalDiskSize int `json:"localDiskSize,omitempty"` // unit: GB

	// CDS 列表, 默认第一块盘作为 docker 和 kubelet 数据盘
	CDSList CDSConfigList `json:"cdsList,omitempty"`

	// Only necessary when InstanceType = GPU
	GPUType  GPUType `json:"gpuType,omitempty"`
	GPUCount int         `json:"gpuCount,omitempty"`
}

// EIPOption 定义 Instance EIP 相关配置
type EIPOption struct {
	EIPName         string            `json:"eipName,omitempty"`
	EIPChargingType BillingMethod `json:"eipChargeType,omitempty"`
	EIPBandwidth    int               `json:"eipBandwidth,omitempty"`
}

// InstancePreChargingOption 定义付费相关配置
type InstancePreChargingOption struct {
	PurchaseTime      int    `json:"purchaseTime,omitempty"`            //  预付费才生效：单位月，12 = 12 月
	AutoRenew         bool   `json:"autoRenew,omitempty"`                  // 是否自动续费
	AutoRenewTimeUnit string `json:"autoRenewTimeUnit,omitempty"` // 续费单位：月
	AutoRenewTime     int    `json:"autoRenewTime,omitempty"`         // 12 = 12 个月
}

// DeleteOption 删除节点选项
type DeleteOption struct {
	MoveOut           bool `json:"moveOut,omitempty"`
	DeleteResource    bool `json:"deleteResource,omitempty"`
	DeleteCDSSnapshot bool `json:"deleteCDSSnapshot,omitempty"`
}

// BBCOption BBC 相关配置
type BBCOption struct {
	// 是否保留数据
	ReserveData bool `json:"reserveData,omitempty"`
	// 磁盘阵列类型 ID
	RaidID string `json:"raidID,omitempty"`
	// 系统盘分配大小
	SysDiskSize int `json:"sysDiskSize,omitempty"`
}

// DeployCustomConfig - 部署自定义配置
type DeployCustomConfig struct {
	// Docker相关配置
	DockerConfig DockerConfig `json:"dockerConfig,omitempty"`

	// kubelet数据目录
	KubeletRootDir string `json:"kubeletRootDir,omitempty"`
	// 是否开启资源预留
	EnableResourceReserved bool `json:"EnableResourceReserved,omitempty"`
	// 资源预留配额,
	// key:value: cpu: 100m, memory: 1000Mi
	KubeReserved map[string]string `json:"kubeReserved,omitempty"`

	// 是否封锁节点
	EnableCordon bool `json:"enableCordon,omitempty"`

	// 部署前执行脚本, 前端 base64编码后传参
	PreUserScript string `json:"preUserScript,omitempty"`
	// 部署后执行脚本, 前端 base64编码后传参
	PostUserScript string `json:"postUserScript,omitempty"`
}


// DockerConfig docker相关配置
type DockerConfig struct {
	DockerDataRoot     string   `json:"dockerDataRoot,omitempty"`     // 自定义 docker 数据目录
	RegistryMirrors    []string `json:"registryMirrors,omitempty"`    // 自定义 RegistryMirrors
	InsecureRegistries []string `json:"insecureRegistries,omitempty"` // 自定义 InsecureRegistries
	DockerLogMaxSize   string   `json:"dockerLogMaxSize,omitempty"`   // docker日志大小，default: 20m
	DockerLogMaxFile   string   `json:"dockerLogMaxFile,omitempty"`   // docker日志保留数，default: 10
	BIP                string   `json:"dockerBIP,omitempty"`          // docker0网桥网段， default: 169.254.30.1/28
}


// ExistedOption 已有实例相关配置
type ExistedOption struct {
	ExistedInstanceID string `json:"existedInstanceID,omitempty"`

	// nil 为默认: 重装系统
	Rebuild *bool `json:"rebuild,omitempty"`
}

// MachineType 机器类型: BCC, BBC
type MachineType string

const (
	// MachineTypeBCC 机器类型 BCC
	MachineTypeBCC MachineType = "BCC"

	// MachineTypeBBC 机器类型 BBC
	MachineTypeBBC MachineType = "BBC"

	// MachineTypeMetal 机器类型 裸金属
	MachineTypeMetal MachineType = "Metal"
)

// CDSConfig clone from BCC
type CDSConfig struct {
	Path        string          `json:"diskPath,omitempty"`
	StorageType bccapi.StorageType `json:"storageType,omitempty"`
	CDSSize     int             `json:"cdsSize,omitempty"`
	SnapshotID  string          `json:"snapshotID,omitempty"`
}

// MountConfig - 磁盘挂载信息
type MountConfig struct {
	Path        string          `json:"diskPath,omitempty"` // "/data"
	CDSID       string          `json:"cdsID,omitempty"`
	Device      string          `json:"device,omitempty"` // "/dev/vdb"
	CDSSize     int             `json:"cdsSize,omitempty"`
	StorageType bccapi.StorageType `json:"storageType,omitempty"`
}

// ClusterRole master & slave
type ClusterRole string

const (
	// ClusterRoleMaster K8S master
	ClusterRoleMaster ClusterRole = "master"

	// ClusterRoleNode K8S node
	ClusterRoleNode ClusterRole = "node"
)

// InstanceOS defines the OS of BCC
type InstanceOS struct {
	ImageType  bccapi.ImageType `json:"imageType,omitempty"` // 镜像类型
	ImageName  string             `json:"imageName,omitempty"` // 镜像名字: ubuntu-14.04.1-server-amd64-201506171832
	OSType     OSType    `json:"osType,omitempty"`       // e.g. linux
	OSName     OSName    `json:"osName,omitempty"`       // e.g. Ubuntu
	OSVersion  string             `json:"osVersion,omitempty"` // e.g. 14.04.1 LTS
	OSArch     string             `json:"osArch,omitempty"`       // e.g. x86_64 (64bit)
	OSBuild    string             `json:"osBuild,omitempty"`     // e.g. 2015061700
}

// InstancePhase CCE InstancePhase
type InstancePhase string

const (
	// InstancePhasePending 创建节点时默认状态
	InstancePhasePending InstancePhase = "pending"

	// InstancePhaseProvisioning IaaS 相关资源正在创建中
	InstancePhaseProvisioning InstancePhase = "provisioning"

	// InstancePhaseProvisioned IaaS 相关资源已经 Ready
	InstancePhaseProvisioned InstancePhase = "provisioned"

	// InstancePhaseRunning 节点运行正常
	InstancePhaseRunning InstancePhase = "running"

	// InstancePhaseCreateFailed 节点异常
	InstancePhaseCreateFailed InstancePhase = "create_failed"

	// InstancePhaseDeleting 节点正在删除
	InstancePhaseDeleting InstancePhase = "deleting"

	// InstancePhaseDeleted 节点删除完成
	InstancePhaseDeleted InstancePhase = "deleted"

	// InstancePhaseDeleteFailed 节点删除失败
	InstancePhaseDeleteFailed InstancePhase = "delete_failed"
)

type CDSConfigList []CDSConfig

type TagList []Tag

type InstanceLabels map[string]string

type InstanceTaints []Taint
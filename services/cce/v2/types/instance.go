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

// InstanceSpec 已有节点需要用户提供：ClusterRole 、短ID，密码，镜像ID,镜像类型, docker storage(可选); BBC要额外加preservedData、raidId、sysRootSize
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

	// 是否为竞价实例
	Bid       bool      `json:"bid,omitempty"`
	BidOption BidOption `json:"bidOption,omitempty"`

	// VPC 相关配置
	VPCConfig VPCConfig `json:"vpcConfig,omitempty"`

	HPASOption *HPASOption `json:"hpasOption,omitempty"`

	// 集群规格相关配置
	InstanceResource InstanceResource `json:"instanceResource,omitempty"`

	// 优先使用 ImageID, 如果用户传入 InstanceOS 信息, 由 service 计算 ImageID
	ImageID    string     `json:"imageID,omitempty"`
	InstanceOS InstanceOS `json:"instanceOS,omitempty"`

	// 实例自定义数据, 支持安装驱动
	UserData string `json:"userData,omitempty"`

	// EIP
	NeedEIP   bool       `json:"needEIP,omitempty"`
	EIPOption *EIPOption `json:"eipOption,omitempty"`

	// AdminPassword
	AdminPassword string `json:"adminPassword,omitempty"`
	SSHKeyID      string `json:"sshKeyID,omitempty"`

	// Charging Type, 通常只支持后付费
	InstanceChargingType      bccapi.PaymentTimingType  `json:"instanceChargingType,omitempty"` // 后付费或预付费
	InstancePreChargingOption InstancePreChargingOption `json:"instancePreChargingOption,omitempty"`

	// 创建虚机的时候，是否需要绑定虚机的tag到虚机的附加的资源上
	RelationTag bool `json:"relationTag,omitempty"`

	// 删除节点选项
	DeleteOption *DeleteOption `json:"deleteOption,omitempty"`

	DeployCustomConfig DeployCustomConfig `json:"deployCustomConfig,omitempty"` // 部署相关高级配置

	Tags TagList `json:"tags,omitempty"`

	Labels      InstanceLabels      `json:"labels,omitempty"`
	Taints      InstanceTaints      `json:"taints,omitempty"`
	Annotations InstanceAnnotations `json:"annotations,omitempty"`

	CCEInstancePriority int `json:"cceInstancePriority,omitempty"`

	AutoSnapshotID string `json:"autoSnapshotID,omitempty"` // 自动快照策略   ID

	IsOpenHostnameDomain bool `json:"isOpenHostnameDomain,omitempty"`

	ResourceGroupID string `json:"resourceGroupID,omitempty"`
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

	SecurityGroup SecurityGroup `json:"securityGroup,omitempty"`

	SecurityGroupType string `json:"securityGroupType"`

	SecurityGroups []SecurityGroupV2 `json:"securityGroups"`
}

// HPASOption 定义 HPAS 配置
type HPASOption struct {
	AppType             string `json:"appType,omitempty"`
	AppPerformanceLevel string `json:"appPerformanceLevel,omitempty"`
}

// SecurityGroup 定义 Instance 安全组配置
type SecurityGroup struct {
	// 是否附加 CCE 必须安全组
	EnableCCERequiredSecurityGroup bool `json:"enableCCERequiredSecurityGroup"`
	// 是否附加 CCE 可选安全组
	EnableCCEOptionalSecurityGroup bool `json:"enableCCEOptionalSecurityGroup"`
	// 用户自定义安全组 ID 列表
	CustomSecurityGroupIDs []string `json:"customSecurityGroups,omitempty"`
}

type SecurityGroupV2 struct {
	Name string            `json:"name"`
	Type SecurityGroupType `json:"type"`
	ID   string            `json:"id"`
}

type SecurityGroupType string

const (
	// 普通安全组
	SecurityGroupTypeNormal SecurityGroupType = "normal"
	// 企业安全组
	SecurityGroupTypeEnterprise         SecurityGroupType = "enterprise"
	SecurityGroupTypeNormalIDPrefix     string            = "g-"
	SecurityGroupTypeEnterpriseIDPrefix string            = "esg-"
)

// InstanceResource 定义 Instance CPU/MEM/Disk 配置
type InstanceResource struct {
	MachineSpec string `json:"machineSpec,omitempty"` // 机器规格，例：bcc.g5.c2m8

	CPU int `json:"cpu,omitempty"` // unit: Core
	MEM int `json:"mem,omitempty"` // unit: GB

	NodeCPUQuota int `json:"nodeCPUQuota,omitempty"` // unit: Core
	NodeMEMQuota int `json:"nodeMEMQuota,omitempty"` // unit: GB

	// RootDisk
	RootDiskType bccapi.StorageType `json:"rootDiskType,omitempty"`
	RootDiskSize int                `json:"rootDiskSize,omitempty"` // unit: GB

	EphemeralDiskList []EphemeralDisk `json:"ephemeralDiskList,omitempty"`

	// GPU 机器必须指定, 其他机器不用
	LocalDiskSize int `json:"localDiskSize,omitempty"` // unit: GB

	// CDS 列表, 默认第一块盘作为 docker 和 kubelet 数据盘
	CDSList CDSConfigList `json:"cdsList,omitempty"`

	// Only necessary when InstanceType = GPU
	GPUType  GPUType `json:"gpuType,omitempty"`
	GPUCount int     `json:"gpuCount,omitempty"`
}

type EphemeralDisk struct {
	StorageType StorageType `json:"storageType,omitempty"`
	SizeInGB    int         `json:"sizeInGB,omitempty"`
}

// StorageType 存储类型
type StorageType string

const (
	// StorageTypeSTD1 上一代云磁盘
	StorageTypeSTD1 StorageType = "sata"

	// StorageTypeHP1 高性能型
	StorageTypeHP1 StorageType = "ssd"

	// StorageTypeCloudHP1 SSD 型
	StorageTypeCloudHP1 StorageType = "premium_ssd"

	StorageTypeNNME StorageType = "nvme"

	// TODO: 下面的值待确定，目前 CCE 不支持

	// StorageTypeHDD 普通型
	StorageTypeHDD StorageType = "hdd"

	// StorageTypeLocal 本地盘
	StorageTypeLocal StorageType = "local"

	// StorageTypeDCCSATA Sata 盘, 创建 DCC 实例专用
	StorageTypeDCCSATA StorageType = "SATA"

	// StorageTypeDCCSSD SSD 盘, 创建 DCC 实例专用
	StorageTypeDCCSSD StorageType = "SSD"

	// StorageTypeEnhancedSSD 增强型SSD
	StorageTypeEnhancedSSD = "enhanced_ssd_pl1"
)

// EIPOption 定义 Instance EIP 相关配置
type EIPOption struct {
	EIPName         string        `json:"eipName,omitempty"`
	EIPChargingType BillingMethod `json:"eipChargeType,omitempty"`
	EIPPurchaseType PurchaseType  `json:"eipPurchaseType,omitempty" gorm:"column:eip_purchase_type"`
	EIPBandwidth    int           `json:"eipBandwidth,omitempty"`
}

// InstancePreChargingOption 定义付费相关配置
type InstancePreChargingOption struct {
	PurchaseTime      int    `json:"purchaseTime,omitempty"` //  预付费才生效：单位月，12 = 12 月
	PurchaseTimeUnit  string `json:"purchaseTimeUnit,omitempty"`
	AutoRenew         bool   `json:"autoRenew,omitempty"`         // 是否自动续费
	AutoRenewTimeUnit string `json:"autoRenewTimeUnit,omitempty"` // 续费单位：月
	AutoRenewTime     int    `json:"autoRenewTime,omitempty"`     // 12 = 12 个月
}

// DeleteOption 删除节点选项
type DeleteOption struct {
	MoveOut           bool `json:"moveOut,omitempty"`
	DeleteResource    bool `json:"deleteResource,omitempty"`
	DeleteCDSSnapshot bool `json:"deleteCDSSnapshot,omitempty"`
	DrainNode         bool `json:"drainNode,omitempty"`
}

// BBCOption BBC 相关配置
type BBCOption struct {
	Flavor   string `json:"flavor,omitempty"`
	DiskInfo string `json:"diskInfo,omitempty"`
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
	// containerd相关配置
	ContainerdConfig ContainerdConfig `json:"containerdConfig,omitempty"`

	// kubelet数据目录
	KubeletRootDir string `json:"kubeletRootDir,omitempty"`
	// 是否开启资源预留
	EnableResourceReserved bool `json:"EnableResourceReserved,omitempty"`
	// k8s进程资源预留配额
	// key:value: cpu: 50m, memory: 100Mi
	KubeReserved map[string]string `json:"kubeReserved,omitempty"`
	// 系统进程资源预留配额
	// key:value: cpu: 50m, memory: 100Mi
	SystemReserved map[string]string `json:"systemReserved,omitempty"`

	// RegistryPullQPS, default: 5
	RegistryPullQPS int `json:"registryPullQPS,omitempty"`
	// RegistryBurst, default: 10
	RegistryBurst int `json:"registryBurst,omitempty"`
	// PodPidsLimit, default: -1
	PodPidsLimit int `json:"podPidsLimit,omitempty"`

	EventRecordQPS *int32 `json:"eventRecordQPS,omitempty"`
	EventBurst     *int32 `json:"eventBurst,omitempty"`
	KubeAPIQPS     *int32 `json:"kubeAPIQPS,omitempty"`   // 自定义 KubeAPIQPS
	KubeAPIBurst   *int32 `json:"kubeAPIBurst,omitempty"` // 自定义 KubeAPIBurst
	MaxPods        *int32 `json:"maxPods,omitempty"`      // 自定义 MaxPods

	// https://kubernetes.io/zh/docs/tasks/administer-cluster/topology-manager/
	CPUManagerPolicy      K8SCPUManagerPolicy      `json:"cpuManagerPolicy,omitempty"`
	TopologyManagerScope  K8STopologyManagerScope  `json:"topologyManagerScope,omitempty"`
	TopologyManagerPolicy K8STopologyManagerPolicy `json:"topologyManagerPolicy,omitempty"`
	CPUCFSQuota           *bool                    `json:"cpuCFSQuota,omitempty"`

	// 是否封锁节点
	EnableCordon bool `json:"enableCordon,omitempty"`

	// 部署前执行脚本, 前端 base64编码后传参
	PreUserScript string `json:"preUserScript,omitempty"`
	// 部署后执行脚本, 前端 base64编码后传参
	PostUserScript string `json:"postUserScript,omitempty"`

	// KubeletBindAddressType, kubelet bind address
	KubeletBindAddressType KubeletBindAddressType `json:"kubeletBindAddressType,omitempty"`

	// PostUserScriptFailedAutoCordon 部署后执行脚本失败自动封锁节点
	PostUserScriptFailedAutoCordon bool `json:"postUserScriptFailedAutoCordon,omitempty"`
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

// ContainerdConfig containerd相关配置
type ContainerdConfig struct {
	DataRoot           string   `json:"dataRoot,omitempty"`           // 自定义 containerd 数据目录
	RegistryMirrors    []string `json:"registryMirrors,omitempty"`    // 自定义 RegistryMirrors
	InsecureRegistries []string `json:"insecureRegistries,omitempty"` // 自定义 InsecureRegistries
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

	// MachineTypeEBC 机器类型 EBC
	MachineTypeEBC MachineType = "EBC"

	// MachineTypeHPAS 机器类型 HPAS
	MachineTypeHPAS MachineType = "HPAS"
)

// CDSConfig clone from BCC
type CDSConfig struct {
	Path        string             `json:"diskPath,omitempty"`
	StorageType bccapi.StorageType `json:"storageType,omitempty"`
	CDSSize     int                `json:"cdsSize,omitempty"`
	SnapshotID  string             `json:"snapshotID,omitempty"`
	DataDevice  string             `json:"dataDevice,omitempty"`
	NeedFormat  bool               `json:"needFormat,omitempty"`
}

// MountConfig - 磁盘挂载信息
type MountConfig struct {
	Path        string             `json:"diskPath,omitempty"` // "/data"
	CDSID       string             `json:"cdsID,omitempty"`
	Device      string             `json:"device,omitempty"` // "/dev/vdb"
	CDSSize     int                `json:"cdsSize,omitempty"`
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

// K8SCPUManagerPolicy - K8S CPUManagerPolicy
type K8SCPUManagerPolicy string

const (
	// K8SCPUManagerPolicyNone - none
	K8SCPUManagerPolicyNone K8SCPUManagerPolicy = "none"

	// K8SCPUManagerPolicyStatic - static
	K8SCPUManagerPolicyStatic K8SCPUManagerPolicy = "static"
)

// K8STopologyManagerScope - K8S topologyManagerScope
type K8STopologyManagerScope string

const (
	// K8STopologyManagerScopePod - pod
	K8STopologyManagerScopePod K8STopologyManagerScope = "pod"

	// K8STopologyManagerScopeContainer - container
	K8STopologyManagerScopeContainer K8STopologyManagerScope = "container"
)

// K8STopologyManagerPolicy - K8S topologyManagerPolicy
type K8STopologyManagerPolicy string

const (
	// K8STopologyManagerPolicyNone - none
	K8STopologyManagerPolicyNone K8STopologyManagerPolicy = "none"

	// K8STopologyManagerPolicyBestEffort - best-effort
	K8STopologyManagerPolicyBestEffort K8STopologyManagerPolicy = "best-effort"

	// K8STopologyManagerPolicyRestricted - restricted
	K8STopologyManagerPolicyRestricted K8STopologyManagerPolicy = "restricted"

	// K8STopologyManagerPolicySingleNumaNode - single-numa-node
	K8STopologyManagerPolicySingleNumaNode K8STopologyManagerPolicy = "single-numa-node"
)

type PurchaseType string

const (
	// EIPPurchaseTypeBGP 标准型BGP
	EIPPurchaseTypeBGP PurchaseType = "BGP"
	// EIPPurchaseTypeBGP_S 增强型BGP
	EIPPurchaseTypeBGP_S PurchaseType = "BGP_S"
	// EIPPurchaseTypeChinaTelcom  电信单线
	EIPPurchaseTypeChinaTelcom PurchaseType = "ChinaTelcom"
	// EIPPurchaseTypeChinaUnicom  联通单线
	EIPPurchaseTypeChinaUnicom PurchaseType = "ChinaUnicom"
	// EIPPurchaseTypeChinaMobile  移动单线
	EIPPurchaseTypeChinaMobile PurchaseType = "ChinaMobile"
)

// InstanceOS defines the OS of BCC
type InstanceOS struct {
	ImageType bccapi.ImageType `json:"imageType,omitempty"` // 镜像类型
	ImageName string           `json:"imageName,omitempty"` // 镜像名字: ubuntu-14.04.1-server-amd64-201506171832
	OSType    OSType           `json:"osType,omitempty"`    // e.g. linux
	OSName    OSName           `json:"osName,omitempty"`    // e.g. Ubuntu
	OSVersion string           `json:"osVersion,omitempty"` // e.g. 14.04.1 LTS
	OSArch    string           `json:"osArch,omitempty"`    // e.g. x86_64 (64bit)
	OSBuild   string           `json:"osBuild,omitempty"`   // e.g. 2015061700
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

// KubeletBindAddressType - kubelet bind address 类型
type KubeletBindAddressType string

const (
	// KubeletBindAddressTypeAll - 0.0.0.0
	KubeletBindAddressTypeAll KubeletBindAddressType = "all"

	// KubeletBindAddressTypeLocal - 127.0.0.1
	KubeletBindAddressTypeLocal KubeletBindAddressType = "local"

	// KubeletBindAddressTypeHostIP - 主网卡 IP
	KubeletBindAddressTypeHostIP KubeletBindAddressType = "hostip"
)

type BidOption struct {
	// BidMode 竞价模式
	BidMode BidMode `json:"bidMode,omitempty"`

	// BidPrice 用户的出价, 仅在 BidMode=BidModeCustomPrice 模式下生效
	BidPrice string `json:"bidPrice,omitempty"`

	// BidTime 竞价超时时间, 单位: minute, 超时会取消该竞价实例订单
	BidTimeout int `json:"bidTimeout,omitempty"`

	// BidReleaseEIP 竞价实例被动释放时, 是否联动释放实例 EIP
	BidReleaseEIP bool `json:"bidReleaseEIP,omitempty"`

	// BidReleaseEIP 竞价实例被动释放时, 是否联动释放实例 CDS
	BidReleaseCDS bool `json:"bidReleaseCDS,omitempty"`
}

type CDSConfigList []CDSConfig

type TagList []Tag

type InstanceLabels map[string]string

type InstanceTaints []Taint

type InstanceAnnotations map[string]string

type BBCFlavorID string

type BidMode string

const (
	// BidModeMarketPrice 跟随市场价出价
	BidModeMarketPrice BidMode = "MARKET_PRICE_BID"

	// BidModeCustomPrice 用户自定义出价
	BidModeCustomPrice BidMode = "CUSTOM_BID"
)

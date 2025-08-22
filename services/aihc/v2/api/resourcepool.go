package api

import "time"

type ResourcePoolPhase string

const (
	// 资源池创建中
	ResourcePoolPhaseCreating ResourcePoolPhase = "creating"

	// 资源池运行正常
	ResourcePoolPhaseRunning ResourcePoolPhase = "running"

	// 资源池正在删除
	ResourcePoolPhaseDeleting ResourcePoolPhase = "deleting"

	// 资源池伸缩中
	ResourcePoolPhaseScaling ResourcePoolPhase = "scaling"

	// 资源池创建失败
	ResourcePoolPhaseCreatedFailed ResourcePoolPhase = "created_failed"

	// 资源池删除失败
	ResourcePoolPhaseDeletedFailed ResourcePoolPhase = "deleted_failed"

	// 资源池已删除
	ResourcePoolPhaseDeleted ResourcePoolPhase = "deleted"

	// 资源池未知状态
	ResourcePoolPhaseUnknown ResourcePoolPhase = "unknown"
)

// RuntimeType defines the runtime on each node
type RuntimeType string

const (
	// RuntimeTypeDocker = docker
	RuntimeTypeDocker RuntimeType = "docker"

	// RuntimeTypeContainerd = containerd
	RuntimeTypeContainerd RuntimeType = "containerd"
)

type OrderType string

const (
	OrderTypeAsc  = "ASC"
	OrderTypeDesc = "DESC"
)

type ResourcePoolAction string

const (
	ResourcePoolActionDescribeResourcePools        ResourcePoolAction = "DescribeResourcePools"
	ResourcePoolActionDescribeResourcePool         ResourcePoolAction = "DescribeResourcePool"
	ResourcePoolActionDescribeResourcePoolOverview ResourcePoolAction = "DescribeResourcePoolOverview"
)

type ResourcePoolType string

const (
	ResourcePoolTypeCommon      ResourcePoolType = "common"      // 通用资源池
	ResourcePoolTypeDedicatedV2 ResourcePoolType = "dedicatedV2" // 托管资源池v2
	ResourcePoolTypeBHCMP       ResourcePoolType = "bhcmp"       // bhcmp资源池
)

type ResourcePoolKeywordType string

const (
	ResourcePoolKeywordTypeResourcePoolID   ResourcePoolKeywordType = "resourcePoolId"
	ResourcePoolKeywordTypeResourcePoolName ResourcePoolKeywordType = "resourcePoolName"
)

type ResorucePoolOrderBy string

const (
	ResorucePoolOrderByResourcePoolID   ResorucePoolOrderBy = "resourcePoolId"
	ResorucePoolOrderByResourcePoolName ResorucePoolOrderBy = "resourcePoolName"
	ResorucePoolOrderByCreatedAt        ResorucePoolOrderBy = "CreatedAt"
)

type DescribeResourcePoolsRequest struct {
	ResourcePoolType ResourcePoolType        `json:"resourcePoolType"`
	KeywordType      ResourcePoolKeywordType `json:"keywordType"`
	Keyword          string                  `json:"keyword"`
	OrderBy          ResorucePoolOrderBy     `json:"orderBy"`
	Order            OrderType               `json:"order"`
	PageNumber       int                     `json:"pageNumber"`
	PageSize         int                     `json:"pageSize"`
}

type DescribeResourcePoolsResponse struct {
	KeywordType   string          `json:"keywordType,omitempty"`
	Keyword       string          `json:"keyword,omitempty"`
	OrderBy       string          `json:"orderBy,omitempty"`
	Order         string          `json:"order,omitempty"`
	PageSize      int             `json:"pageSize,omitempty"`
	PageNumber    int             `json:"pageNumber,omitempty"`
	TotalCount    int             `json:"totalCount"`
	ResourcePools []*ResourcePool `json:"resourcePools"`
}

type DescribeResourcePoolResponse struct {
	ResourcePool
}

type DescribeResourcePoolOverviewResponse struct {
	ResourcePoolsStatisticWithType []*ResourcePoolsStatisticWithType `json:"resourcePoolsStatisticWithType"`
}

type ResourcePool struct {
	AccountID           string               `json:"-"`
	UserID              string               `json:"-"`
	CreateBy            string               `json:"createdBy,omitempty"`
	ResourcePoolID      string               `json:"resourcePoolId,omitempty"`
	Type                string               `json:"type,omitempty"`
	Name                string               `json:"name,omitempty"`
	Region              string               `json:"region"`                // 地域
	Description         string               `json:"description,omitempty"` // 描述
	Configuration       *Configuration       `json:"configuration,omitempty"`
	CreatedAt           time.Time            `json:"createdAt"`           // 创建时间，readonly
	UpdatedAt           time.Time            `json:"updatedAt,omitempty"` // 更新时间，readonly
	Phase               ResourcePoolPhase    `json:"phase,omitempty"`
	K8SVersion          string               `json:"k8sVersion,omitempty"`
	RuntimeType         RuntimeType          `json:"runtimeType,omitempty"`
	RuntimeVersion      string               `json:"runtimeVersion,omitempty"`
	AssociatedResources []*AssociateResource `json:"associatedResources,omitempty"`
	Network             *Network             `json:"network,omitempty"`
	PluginNames         []string             `json:"pluginNames,omitempty"`
	BindingStorages     []*ProviderInfo      `json:"bindingStorages,omitempty"`
	BindingMonitor      []*ProviderInfo      `json:"bindingMonitor,omitempty"`
	NodeNum             int                  `json:"nodeNum,omitempty"`
	UpgradeWorkflowID   string               `json:"upgradeWorkflowID,omitempty"`
}

type Configuration struct {
	ExposedPublic                   *bool `json:"exposedPublic,omitempty"`
	ForbidDelete                    *bool `json:"forbidDelete,omitempty"`
	DeschedulerEnabled              *bool `json:"deschedulerEnabled,omitempty"`
	UnifiedSchedulerEnabled         *bool `json:"unifiedSchedulerEnabled,omitempty"`
	DatasetPermissionEnabled        *bool `json:"datasetPermissionEnabled,omitempty"`
	VolumePermissionEnabled         *bool `json:"volumePermissionEnabled,omitempty"`
	ImageNoAuthPullEnabled          *bool `json:"imageNoAuthPullEnabled,omitempty"`
	PublicNetInferenceServiceEnable *bool `json:"publicNetInferenceServiceEnable,omitempty"`
	//AutoCreateMonitorInstance *bool `json:"autoCreateMonitorInstance,omitempty"`
}

type AssociateResource struct {
	Provider string `json:"provider"`
	ID       string `json:"id"`
	UUID     string `json:"uuid,omitempty"`
	Region   string `json:"region,omitempty"`
	Zone     string `json:"zone,omitempty"`
}

// Network 集群网络信息
type Network struct {
	Mode string `json:"mode"` // support: vpc-route, vpc-eni

	// master
	Master *NetworkInfo `json:"master"`
	// nodes
	Nodes *NetworkInfo `json:"nodes"`

	// pods
	Pods *NetworkInfo `json:"pods,omitempty"`

	// clusterIPCIDR
	ClusterIPCIDR string `json:"clusterIPCidr,omitempty"`

	// loadBalanceService
	LoadBalanceService *NetworkInfo `json:"loadBalanceService,omitempty"`

	// maxPodsPerNode
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty"`
}

type NetworkInfo struct {
	Region string `json:"region"`

	// vpcID
	VPCID   string `json:"vpcId"`
	VPCUUID string `json:"vpcUuid,omitempty"`
	VPCCIDR string `json:"vpcCidr,omitempty"`

	// subnetIDs
	SubnetIDs []string `json:"subnetIds,omitempty"`
	// subnets
	Subnets []*Subnet `json:"subnets,omitempty"`
	// SubnetCidr
	SubnetCidr string `json:"subnetCidr,omitempty"`
	// securityGroups
	SecurityGroups []*SecurityGroup `json:"securityGroups,omitempty"`
}

type Subnet struct {
	SubnetID   string `json:"subnetId"`
	SubnetCIDR string `json:"subnetCIDR"`
	Zone       string `json:"zone"`
}

type SecurityGroup struct {
	Type string `json:"type"` // support: common, enterprise
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProviderInfo struct {
	Provider string            `json:"provider"`
	Type     string            `json:"type,omitempty"`
	ID       string            `json:"id"`
	Options  map[string]string `json:"options,omitempty"`
	Region   string            `json:"region,omitempty"`
	Zone     string            `json:"zone,omitempty"`
}

type ResourcePoolsStatisticWithType struct {
	Type                  string                 `json:"type,omitempty"`
	NodesStatistic        *NodesStatistic        `json:"nodesStatistic,omitempty"`
	AcceleratorsStatistic *AcceleratorsStatistic `json:"acceleratorsStatistic,omitempty"`
	ResourcesStatistic    []*ResourceStatistic   `json:"resourcesStatistic,omitempty"`
	// ResourcePoolStatistics 资源池统计信息列表
	// 每个资源池统计信息存储在缓存中，列表中包含多个资源池统计信息，需要按照请求的租户信息进行汇聚
	ResourcePoolStatistics []*ResourcePoolStatistic `json:"resourcePoolStatistics,omitempty"`
}

// NodesStatistic 节点统计信息
type NodesStatistic struct {
	// TotalCount 节点总数
	TotalCount int64 `json:"totalCount"`
	// 扩展其他节点相关统计，比如分节点状态的统计等
}

// AcceleratorsStatistic 加速器统计信息
type AcceleratorsStatistic struct {
	ResourceStatistic `json:",inline"`
	// 扩展其他芯片相关统计，比如按照卡芯片统计
	AcceleratorStatistics []*AcceleratorStatistic `json:"acceleratorStatistics,omitempty"`
}

// ResourceStatistic 资源统计信息，原生资源+k8s扩展资源
type ResourceStatistic struct {
	// ResourceType 资源类型
	ResourceType string `json:"resourceType,omitempty"`
	// Capacity 容量总数
	Capacity float64 `json:"capacity"`
	// Allocated 已分配总数
	Allocated float64 `json:"allocated"`
	// Idle 空闲总数
	Idle float64 `json:"idle"`

	// AllocationRate 分配率
	AllocationRate float64 `json:"allocationRate"`
	// UtilizationRate 使用率
	UtilizationRate float64 `json:"utilizationRate"`
}

// AcceleratorStatistic 加速器统计信息
type AcceleratorStatistic struct {
	// AcceleratorType 加速器类型
	AcceleratorType   string `json:"acceleratorType,omitempty"`
	ResourceStatistic `json:",inline"`
}

// ResourcePoolStatistic 资源池统计信息
type ResourcePoolStatistic struct {
	// ResourcePoolID 资源池ID
	ResourcePoolID string `json:"resourcePoolId,omitempty"`
	// Type 类型
	Type string `json:"type,omitempty"`
	// Name 名称
	Name string `json:"name,omitempty"`
	// NodesStatistic 节点统计信息
	NodesStatistic *NodesStatistic `json:"nodesStatistic,omitempty"`

	// ResourcesStatistic 原生 or k8s扩展资源统计信息
	ResourcesStatistic []*ResourceStatistic `json:"resourcesStatistic,omitempty"`

	//AcceleratorsStatistic 加速器统计信息
	AcceleratorsStatistic *AcceleratorsStatistic `json:"acceleratorsStatistic,omitempty"`

	NodeResourceStatistic map[string]*NodeResource `json:"nodeResourceStatistic,omitempty"`
}

type NodeResource struct {
	IsAcclerator         bool                `json:"isAcclerator,omitempty"`
	AccleratorType       string              `json:"accleratorType,omitempty"`
	AccleratorDescriptor string              `json:"accleratorDescriptor,omitempty"`
	Capacity             *NodeResourceAmount `json:"capacity,omitempty"`
	Allocatable          *NodeResourceAmount `json:"allocatable,omitempty"`
	Allocated            *NodeResourceAmount `json:"allocated,omitempty"`
	TaskPodList          []*TaskPod          `json:"taskPodList,omitempty"`
	ResourceConfig       string              `json:"resourceConfig,omitempty"`
}

type NodeResourceAmount struct {
	MilliCPUcores       int64              `json:"milliCPUcores,omitempty"`
	MemoryBytes         int64              `json:"memoryBytes,omitempty"`
	GPUNum              int64              `json:"gpuNum,omitempty"`
	GPUMemoryGi         int64              `json:"gpuMemoryGi,omitempty"`
	AcceleratorCardList []*AcceleratorCard `json:"acceleratorCardList,omitempty"`
}

type AcceleratorCard struct {
	AcceleratorCount string `json:"acceleratorCount"`
	AcceleratorType  string `json:"acceleratorType"`
}

type TaskPod struct {
	PodName   string `json:"podName"`
	Namespace string `json:"namespace"`
}

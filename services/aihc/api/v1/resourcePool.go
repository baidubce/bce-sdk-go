package v1

import "time"

// 是否干预，枚举：未干预，封锁，污点
type Intervention string

const (
	Cordoned     Intervention = "cordoned"
	Tainted      Intervention = "tainted"
	Unintervened Intervention = "unintervened"
)

type ListResourcePoolRequest struct {
	OrderBy     string `json:"orderBy"`
	Order       string `json:"order"`
	Keyword     string `json:"keyword"`
	KeywordType string `json:"keywordType"`
	PageNo      int    `json:"pageNo"`
	PageSize    int    `json:"pageSize"`
}

type ListResourcePoolNodeRequest struct {
	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
}

type ResourcePool struct {
	Metadata Metadata           `json:"metadata"` // 资源池元信息
	Spec     ResourcePoolSpec   `json:"spec"`     // 资源池属性
	Status   ResourcePoolStatus `json:"status"`   // 资源池状态
}

// 资源池元信息
type Metadata struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`      // 资源池名称
	Id        string    `json:"id"`        // 资源池 ID，CCE Cluster ID
	UpdatedAt time.Time `json:"updatedAt"` // 资源池修改时间
}

// 资源池属性
type ResourcePoolSpec struct {
	K8SVersion         string   `json:"k8sVersion"`         // k8s版本
	AssociatedCpromIds []string `json:"associatedCpromIds"` // 关联 Cprom Id
	AssociatedPfsID    string   `json:"associatedPfsId"`    // 关联 PFS ID
	CreatedBy          string   `json:"createdBy"`          // 创建人
	Description        string   `json:"description"`        // 描述
	ForbidDelete       bool     `json:"forbidDelete"`       // 资源池保护，设为 true 禁止删除
	Region             string   `json:"region"`             // 地域
}

// 资源池状态
type ResourcePoolStatus struct {
	GPUDetails []*GPUDetail `json:"gpuDetail"` // GPU/NPU 资源详情
	GPUCount   *GPUCount    `json:"gpuCount"`  // 资源池计算卡状态 ，统计 GPU，昆仑 or 晟腾卡情况
	NodeCount  *NodeCount   `json:"nodeCount"` // 资源池节点状态
	Phase      string       `json:"phase"`     // 资源池状态，pending，running，deleting等
}

// 资源池计算卡状态 ，统计 GPU，昆仑 or 晟腾卡情况
type GPUCount struct {
	Total int64 `json:"total"` // 计算卡总数
	Used  int64 `json:"used"`  // 已分配
}

type GPUDetail struct {
	GPUDescriptor string `json:"gpuDescriptor"` // GPU/NPU 资源描述符
	Total         int64  `json:"total"`         // 计算卡总数
	Used          int64  `json:"used"`          // 已分配
}

// 资源池节点状态
type NodeCount struct {
	Total int64 `json:"total"` // 节点总数
	Used  int64 `json:"used"`  // 已分满节点
}

type GetResourcePoolResponse struct {
	RequestId string                 `json:"requestId"`
	Result    *GetResourcePoolResult `json:"result"`
}

type GetResourcePoolResult struct {
	ResourcePool *ResourcePool `json:"resourcePool"`
}

type ListNodeByResourcePoolResponse struct {
	RequestId string    `json:"requestId"`
	Result    *NodePage `json:"result"`
}

type NodePage struct {
	Nodes    []*Node `json:"nodes"`
	Total    int     `json:"total"`
	OrderBy  string  `json:"orderBy"`
	Order    string  `json:"order"`
	PageNo   int     `json:"pageNo"`
	PageSize int     `json:"pageSize"`
}

// 节点模型
type Node struct {
	ChargingType    string       `json:"chargingType"` // 节点付费方式
	GPUAllocated    int64        `json:"gpuAllocated"` // GPU/NPU 分配卡数
	GPUTotal        int64        `json:"gpuTotal"`     // GPU/NPU 总卡数
	GPUType         string       `json:"gpuType"`      // GPU/NPU 类型
	InstanceID      string       `json:"instanceId"`   // BCC 实例短 ID
	InstanceName    string       `json:"instanceName"` // BCC 实例名称
	Intervention    Intervention `json:"intervention"` // 是否干预，枚举：未干预，封锁，污点
	NodeName        string       `json:"nodeName"`     // CCE 节点名称
	StatusPhase     string       `json:"statusPhase"`  // 节点状态
	Zone            string       `json:"zone"`         // 可用区
	Region          string       `json:"region"`       // 地域
	CreatedAt       time.Time    `json:"createdAt"`
	GPUDescriptor   string       `json:"gpuDescriptor"` // GPU/NPU 资源描述符
	CPUTotal        int64        `json:"cpuTotal"`
	CPUAllocated    int64        `json:"cpuAllocated"`
	MemoryTotal     int64        `json:"memoryTotal"`
	MemoryAllocated int64        `json:"memoryAllocated"`
}

type ListResourcePoolResponse struct {
	RequestId string            `json:"requestId"`
	Result    *ResourcePoolPage `json:"result"`
}

type ResourcePoolPage struct {
	ResourcePools []*ResourcePool `json:"resourcePools"`
	Total         int             `json:"total"`
	OrderBy       string          `json:"orderBy"`
	Order         string          `json:"order"`
	PageNo        int             `json:"pageNo"`
	PageSize      int             `json:"pageSize"`
}

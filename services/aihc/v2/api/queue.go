package api

import (
	"time"
)

type ResourceQueueAction string

const (
	ResourceQueueActionDescribeResourceQueues ResourceQueueAction = "DescribeQueues"
	ResourceQueueActionDescribeResourceQueue  ResourceQueueAction = "DescribeQueue"
)

type QueueType string
type QueueingStrategy string

const (
	QueueTypeRegular  QueueType = "Regular"
	QueueTypeElastic  QueueType = "Elastic"
	QueueTypePhysical QueueType = "Physical"
)

type ResourceList map[string]string

type DescribeResourceQueuesRequest struct {
	OrderBy        string `json:"orderBy,omitempty"` //排序字段不生效
	Order          string `json:"order,omitempty"`
	Keyword        string `json:"keyword,omitempty"`
	KeywordType    string `json:"keywordType,omitempty"`
	PageNumber     int    `json:"pageNumber,omitempty"` //分页字段自运维队列不生效，以树状形式返回
	PageSize       int    `json:"pageSize,omitempty"`
	ResourcePoolID string `json:"resourcePoolId"`
}

type DescribeResourceQueuesResponse struct {
	PageSize    int64        `json:"pageSize,omitempty"`
	PageNumber  int64        `json:"pageNumber,omitempty"`
	Keyword     string       `json:"keyword,omitempty"`
	KeywordType string       `json:"keywordType,omitempty"`
	TotalCount  int64        `json:"totalCount"`
	Queues      []*QueueSpec `json:"queues"`
}

type DescribeResourceQueueResponse struct {
	QueueSpec
}

type QueueSpec struct {
	// 基础属性
	QueueID        string    `json:"queueId"`
	QueueName      string    `json:"queueName"`
	Description    string    `json:"description,omitempty"` // 描述
	QueueType      QueueType `json:"queueType"`             // 队列类型，enum: elastic，physical
	ResourcePoolID string    `json:"resourcePoolId"`        // 来源资源池Id列表
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`

	// 层级关系属性
	ParentQueue string       `json:"parentQueue"`        // 父队列，默认为 root 队列
	Children    []*QueueSpec `json:"children,omitempty"` // 子队列列表

	// 调度策略属性
	Opened           bool             `json:"opened"`                // 队列是否开启
	Reclaimable      bool             `json:"reclaimable"`           // 是否可回收
	Preemptable      bool             `json:"preemptable,omitempty"` // 是否开启队列内优先级抢占
	DisableOversell  bool             `json:"disableOversell"`
	QueueingStrategy QueueingStrategy `json:"queueingStrategy,omitempty"` // 设置队列调度策略
	EnableVGPU       bool             `json:"enableVGPU,omitempty"`

	// 队列配额属性
	Capability *ResourceAmount `json:"capability,omitempty"` // 队列上线配额
	Deserved   *ResourceAmount `json:"deserved,omitempty"`   // 队列申请配额
	Guarantee  *ResourceAmount `json:"guarantee,omitempty"`  // 队列预留配额
	Allocated  *ResourceAmount `json:"allocated,omitempty"`  // 队列分配情况

	// 队列统计
	RunningJobs             int                `json:"runningJobs,omitempty"`
	InqueueJobs             int                `json:"inqueueJobs,omitempty"`
	PendingJobs             int                `json:"pendingJobs,omitempty"`
	MaxDeservedForSubqueue  *ResourceAmount    `json:"maxDeservedForSubqueue,omitempty"`
	MaxGuaranteeForSubqueue *ResourceAmount    `json:"maxGuaranteeForSubqueue,omitempty"`
	BindingNodes            []*BindingNodeInfo `json:"bindingNodes,omitempty"`
}

type ResourceAmount struct {
	MilliCPUcores       string             `json:"milliCPUcores,omitempty"`
	MemoryGi            string             `json:"memoryGi,omitempty"`
	AcceleratorCardList []*AcceleratorCard `json:"acceleratorCardList,omitempty"`
}

type SpecInfo struct {
	MachineSpec  string   `json:"machineSpec"`
	K8sNodeNames []string `json:"k8sNodeNames"`
	Count        int      `json:"count"`
}

type BindingNodeInfo struct {
	MachineSpec     string   `json:"machineSpec"`
	NodeNameList    []string `json:"nodeNameList"`
	Count           int      `json:"count"`
	AcceleratorType string   `json:"acceleratorType"`
}

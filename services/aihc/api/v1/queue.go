package v1

import (
	"time"
)

type ListQueueRequest struct {
	OrderBy     string `json:"orderBy"`
	Order       string `json:"order"`
	Keyword     string `json:"keyword"`
	KeywordType string `json:"keywordType"`
	PageNo      int    `json:"pageNo"`
	PageSize    int    `json:"pageSize"`
}

type GetQueuesResponse struct {
	Result    *GetQueuesResult `json:"result"`
	RequestId string           `json:"requestID"`
}

type GetQueuesResult struct {
	Queue *Queue `json:"queue"`
}

type ListQueuesResponse struct {
	Result    *QueuePage `json:"result"`
	RequestId string     `json:"requestID"`
}

type QueuePage struct {
	Total    int      `json:"total"`
	Queues   []*Queue `json:"queues"`
	OrderBy  string   `json:"orderBy"`
	Order    string   `json:"order"`
	PageNo   int      `json:"pageNo"`
	PageSize int      `json:"pageSize"`
}

type Queue struct {
	Children         []*Queue               `json:"children,omitempty"`         // 子队列列表
	CreatedTime      time.Time              `json:"createdTime"`                // 创建时间
	Description      string                 `json:"description,omitempty"`      // 描述
	Name             string                 `json:"name"`                       // 名称
	NodeList         []string               `json:"nodeList,omitempty"`         // 物理队列绑定节点列表，物理队列指定节点范围时，创建时为必须字段
	Users            []string               `json:"users,omitempty"`            // 队列成员-用户
	Groups           []string               `json:"groups,omitempty"`           // 队列成员-用户组
	ManagerUsers     []string               `json:"managerUsers,omitempty"`     // 队列成员-用户
	ManagerGroups    []string               `json:"managerGroups,omitempty"`    // 队列成员-用户组
	ParentQueue      string                 `json:"parentQueue"`                // 父队列，默认为 root 队列
	QueueType        QueueType              `json:"queueType"`                  // 队列类型，enum: regular， elastic，physical
	State            string                 `json:"state"`                      // 状态，Open or Close
	Reclaimable      *bool                  `json:"reclaimable"`                // 是否可回收
	Preemptable      *bool                  `json:"preemptable,omitempty"`      // 是否开启队列内优先级抢占
	QueueingStrategy string                 `json:"queueingStrategy,omitempty"` // 设置队列调度策略
	Capability       ResourceList           `json:"capability,omitempty"`       // 队列上线配额
	Deserved         ResourceList           `json:"deserved,omitempty"`         // 队列申请配额
	Guarantee        ResourceList           `json:"guarantee,omitempty"`        // 队列预留配额
	Allocated        ResourceList           `json:"allocated,omitempty"`        // 队列分配情况
	DisableOversell  bool                   `json:"disableOversell"`
	Running          int32                  `json:"running"`
	Inqueue          int32                  `json:"inqueue"`
	Pending          int32                  `json:"pending"`
	Nodes            map[string][]*SpecInfo `json:"nodes,omitempty"` // 队列节点列表
	Opened           *bool                  `json:"opened"`
}

type SpecInfo struct {
	MachineSpec  string   `json:"machineSpec"`
	K8sNodeNames []string `json:"k8sNodeNames"`
	Count        int      `json:"count"`
}

type QueueType string

const (
	QueueTypeRegular  QueueType = "Regular"
	QueueTypeElastic  QueueType = "Elastic"
	QueueTypePhysical QueueType = "Physical"
)

type ResourceList map[string]string

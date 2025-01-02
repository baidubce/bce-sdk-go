/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

type ConfigType string

const (
	ConfigTypeSystemPresetConfig ConfigType = "SystemPresetConfig"
	ConfigTypeCustomConfig       ConfigType = "CustomConfig"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// NodeRemedierSpec defines the desired state of NodeRemedier
type NodeRemedierSpec struct {
	// 是否启用
	// +kubebuilder:validation:Optional
	Paused bool `json:"paused,omitempty"`
	// 集群名称
	ClusterName string `json:"clusterName"`
	// 集群节点
	// +kubebuilder:validation:Optional
	NodeSelector LabelSelector `json:"nodeSelector,omitempty"`
	// 不健康 condition
	// +kubebuilder:validation:Optional
	UnhealthyConditions []UnhealthyCondition `json:"unhealthyConditions,omitempty"`
	// 机器维修设置
	NodeRemediation *NodeRemediation `json:"nodeRemediation,omitempty"`

	// The number of remedyTasks to retain.
	// This is a pointer to distinguish between explicit zero and not specified.
	// Defaults to 10.
	// // +kubebuilder:validation:Optional
	RemedyTaskLimit *int32 `json:"remedyTaskLimit,omitempty"`
}

// NodeRemedierStatus defines the observed state of NodeRemedier
type NodeRemedierStatus struct {
	UpdateTime          *Time `json:"updateTime,omitempty"`
	ExpectedNodes       int32 `json:"expectedNodes,omitempty"`
	CurrentHealthy      int32 `json:"currentHealthy,omitempty"`
	RemediationsAllowed int32 `json:"remediationsAllowed,omitempty"`
	// +kubebuilder:validation:Optional
	UnhealthyTarget []UnhealthyTarget `json:"unhealthyTarget,omitempty"`
	// +kubebuilder:validation:Optional
	ProcessTarget []UnhealthyTarget `json:"processTarget,omitempty"`
}

// NodeRemedier is the Schema for the noderemediers API
type NodeRemedier struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeRemedierSpec   `json:"spec,omitempty"`
	Status NodeRemedierStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NodeRemedierList contains a list of NodeRemedier
type NodeRemedierList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`
	Items    []NodeRemedier `json:"items,omitempty"`
}

type NodeRemediation struct {
	MachineRemediationTemplate *ObjectReference `json:"machineRemediationTemplate,omitempty"`

	NodeRemedierPolicy *NodeRemedierPolicy `json:"nodeRemedierPolicy,omitempty"`
}

// Node 故障状态
type UnhealthyCondition struct {
	// 故障维修优先级, 默认为 0, 0 ～ 1000 为软件故障，1000 ～ 2000 为 硬件故障；2000 以上为未定义
	// +kubebuilder:validation:Optional
	Priority int64 `json:"priority,omitempty"`
	// 配置模版类型 自定义 or 预制
	// +kubebuilder:validation:Optional
	ConfigType ConfigType `json:"configType,omitempty"`
	// 开启检测
	// +kubebuilder:validation:Optional
	EnableCheck bool `json:"enableCheck,omitempty"`
	// 开启维修
	// +kubebuilder:validation:Optional
	EnableRemediation bool `json:"enableRemediation,omitempty"`
	// node 状态类型
	// +kubebuilder:validation:Optional
	Type NodeConditionType `json:"type,omitempty"`
	// node 状态
	// +kubebuilder:validation:Optional
	Status ConditionStatus `json:"status,omitempty"`
	// 故障持续时间
	// +kubebuilder:validation:Optional
	Timeout string `json:"timeout,omitempty"`
	// 机器维修模板
	// +kubebuilder:validation:Optional
	MachineRemediationTemplate *ObjectReference `json:"machineRemediationTemplate,omitempty"`
}

// nodeSelector 级别，不同CR的node之间做聚合。例如，一条cr选定十台机器，已经开始维修一台node3。另一条cr选定十台机器，也包括node3，node2也触发故障，此时已经计算维修了一台机器。
type NodeRemedierPolicy struct {
	// 每天最大维修数
	// +kubebuilder:validation:Optional
	MaxNodesPerDay *IntOrString `json:"maxNodesPerDay,omitempty"`
	// 每小时最大维修数
	// +kubebuilder:validation:Optional
	MaxNodesPerHour *IntOrString `json:"maxNodesPerHour,omitempty"`
	// 不健康比例，node达到不健康比例后暂停维修

	// 当前维修最大并发
	// +kubebuilder:validation:Optional
	MaxProcessingNodes *IntOrString `json:"maxProcessingNodes,omitempty"`
	// 同一个node维修最小间隔
	// +kubebuilder:validation:Optional
	MinIntervalPerNode string `json:"minIntervalPerNode,omitempty"`

	// 不同node维修最小间隔
	//MinIntervalEveryNode string `json:"minIntervalEveryNode,omitempty"`

	// 不健康比例，node达到不健康比例后暂停维修
	// +kubebuilder:validation:Optional
	MaxUnhealthy *IntOrString `json:"maxUnhealthy,omitempty"`
	// [3-5] 只有故障在 3-5 台可以修
	// +kubebuilder:validation:Optional
	UnhealthyRange *string `json:"unhealthyRange,omitempty"`
}

type UnhealthyTarget struct {
	// +kubebuilder:validation:Optional
	Name string `json:"name,omitempty"`
	// +kubebuilder:validation:Optional
	UnhealthyConditions []NodeCondition `json:"unhealthyConditions,omitempty"`
}

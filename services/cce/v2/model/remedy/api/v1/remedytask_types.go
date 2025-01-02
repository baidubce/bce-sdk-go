package v1

import (
	"time"
)

// RemedyTask is the Schema for the RemedyTask API
type RemedyTask struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	Spec   RemedyTaskSpec   `json:"spec,omitempty"`
	Status RemedyTaskStatus `json:"status,omitempty"`
}

// RemedyTaskSpec defines the desired state of RemedyTask
type RemedyTaskSpec struct {
	NodeProvider string `json:"nodeProvider,omitempty"`
	// 节点名称
	NodeName string `json:"nodeName"`
	// 集群名称
	ClusterName string `json:"clusterName"`
	// 是否暂停维修
	Paused bool `json:"paused"`
	// Task 处理的 ConditionType 列表
	ConditionsTypes []string `json:"conditionsTypes,omitempty"`
	// 维修步骤
	ReconcileSteps []StepName `json:"reconcileSteps,omitempty"`
	// 封锁设置
	MachineCordonConfig MachineCordonConfig `json:"machineCordonConfig,omitempty"`
	// 排水配置
	MachineDrainConfig DrainConfig `json:"machineDrainConfig,omitempty"`
	// 执行步骤间的最小事件间隔
	MinIntervalByStep time.Duration `json:"minIntervalByStep,omitempty"`
	// 维修密码
	MachineRepairConfig MachineRepairConfig `json:"machineRepairConfig,omitempty"`
	// 机器 重启设置
	MachineRebootConfig MachineRebootConfig `json:"machineRebootConfig,omitempty"`
	// 机器 添加配置
	MachineAddConfig MachineAddConfig `json:"machineAddConfig,omitempty"`
	// 机器 下线设置
	MachineDeleteConfig MachineDeleteConfig `json:"machineDeleteConfig,omitempty"`
	// 机器 重启设置
	MachineReInstallConfig MachineReInstallConfig `json:"machineInstallConfig,omitempty"`
	// 机器 扩容设置
	MachineScaleUpConfig MachineScaleUpConfig `json:"machineScaleUpConfig,omitempty"`
	// 机器 检测故障恢复配置
	MachineDetectRecoveryConfig MachineDetectRecoveryConfig `json:"machineDetectRecoveryConfig,omitempty"`
	// 维修步骤的通知配置
	NotificationConfig *NotificationConfig `json:"notificationConfig,omitempty"`
}

// RemedyTaskStatus defines the observed state of RemedyTask
type RemedyTaskStatus struct {
	ReconcileSteps       map[StepName]Step `json:"reconcileSteps,omitempty"`
	ReconcileDeleteSteps map[StepName]Step `json:"reconcileDeleteSteps,omitempty"`
	InstancePhase        InstancePhase     `json:"instancePhase,omitempty"`
	LastFinished         *Time             `json:"lastFinished,omitempty"`
	// 最近发起维修时间
	LastStarted   *Time  `json:"lastStarted,omitempty"`
	Message       string `json:"message,omitempty"`
	NodeIsHealthy bool   `json:"nodeIsHealthy,omitempty"`
}

// RemedyTaskList contains a list of RemedyTasks
type RemedyTaskList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`
	Items    []RemedyTask `json:"items"`
}

type RemedyTasksByCreateTimestamp []RemedyTask

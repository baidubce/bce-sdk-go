/*
Copyright 2022.

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

import (
	"time"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type RepairMachineConfig struct {
	BaseStepConfig        `json:",inline"`
	AutomaticAuth         bool   `json:"automaticAuth,omitempty"`
	RequestAuthWebhookURL string `json:"requestAuthWebhookURL,omitempty"`
}

// MachineRebootConfig define the machine reboot config
type MachineRebootConfig struct {
	BaseStepConfig `json:",inline"`
	// 重启设置，最多重启次数
	MaxRebootTimes int64 `json:"maxRestartTimes"`
}

type MachineRepairConfig struct {
	BaseStepConfig `json:",inline"`
	AuthConfig     MachineRepairAuthConfig `json:"authConfig,omitempty"`
}

type MachineRepairAuthConfig struct {
	Automatic             bool          `json:"automatic,omitempty"`
	RetryInterval         time.Duration `json:"retryInterval,omitempty"`
	RequestAuthWebhookURL string        `json:"requestAuthWebhookURL,omitempty"`
}

type SSHKey struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

type MachineAddConfig struct {
	BaseStepConfig `json:",inline"`
}

// MachineDeleteConfig define the machine offline config
type MachineDeleteConfig struct {
	BaseStepConfig       `json:",inline"`
	KeepInstanceCount    bool `json:"keepInstanceCount,omitempty"`
	KeepPostPaidInstance bool `json:"KeepPostPaidInstance,omitempty"`
}

// MachineReInstallConfig define the machine reinstall config
type MachineReInstallConfig struct {
	BaseStepConfig `json:",inline"`
	// 机器重装模板
	MachineTemplate string `json:"machineTemplate,omitempty"`
}

// MachineScaleUpConfig define machine increase config
type MachineScaleUpConfig struct {
	BaseStepConfig `json:",inline"`
	// 机器扩容模板
	NodeTemplate string `json:"nodeTemplate,omitempty"`
}

type MachineCordonConfig struct {
	BaseStepConfig `json:",inline"`
}

type MachineDetectRecoveryConfig struct {
	BaseStepConfig `json:",inline"`
}

// RemedyMachineSpec defines the desired state of RemedyMachine
type RemedyMachineSpec struct {
	NodeProvider string `json:"nodeProvider,omitempty"`
	// 节点名称
	NodeName string `json:"nodeName"`
	// 集群名称
	ClusterName string `json:"clusterName"`
	// 是否暂停维修
	Paused bool `json:"paused"`
	// 维修步骤
	ReconcileSteps []StepName `json:"reconcileSteps,omitempty"`
	// 排水配置
	DrainConfig DrainConfig `json:"drainConfig,omitempty"`
	// 暂时未使用
	MinIntervalByStep int64 `json:"minIntervalByStep,omitempty"`
	// 维修密码
	MachineRepairConfig MachineRepairConfig `json:"machineRepairConfig,omitempty"`
	// 机器 重启设置
	MachineRebootConfig MachineRebootConfig `json:"machineRebootConfig,omitempty"`
	// 机器 下线设置
	MachineDeleteConfig MachineDeleteConfig `json:"machineDeleteConfig,omitempty"`
	// 机器 重启设置
	MachineReInstallConfig MachineReInstallConfig `json:"machineInstallConfig,omitempty"`
	// 机器 扩容设置
	MachineScaleUpConfig MachineScaleUpConfig `json:"machineScaleUpConfig,omitempty"`
}

type NotificationConfig struct {
	WebhookURL string `json:"webhookURL,omitempty"`
}

type DrainConfig struct {
	BaseStepConfig `json:",inline"`
	// 是否跳过排水
	SkipDrain bool `json:"skipDrain,omitempty"`
	// 是否驱逐 Daemonset 创建的 pod
	EvictDaemonsetPods bool `json:"evictDaemonsetPods,omitempty"`
	// 是否驱逐 StatefulSet 创建的 pod
	EvictStatefulsetPods bool `json:"evictStatefulsetPods,omitempty"`
	// 排水超时时间
	TimeOut int `json:"timeout,omitempty"`
	// 是否强制删除
	Force bool `json:"force,omitempty"`
	// 优雅删除时间
	GracePeriodSeconds int `json:"gracePeriodSeconds,omitempty"`

	SkipWaitForDeleteTimeoutSeconds int `json:"skipWaitForDeleteTimeoutSeconds,omitempty"`
	// 是否删除使用empty dir的pod
	DeleteEmptyDirData bool `json:"deleteEmptyDirData,omitempty"`
	// 是否删除使用host path的pod
	DeleteHostPathData bool `json:"deleteHostPathData,omitempty"`
}

// RemedyMachineStatus defines the observed state of RemedyMachine
type RemedyMachineStatus struct {
	ReconcileSteps       map[StepName]Step `json:"reconcileSteps,omitempty"`
	ReconcileDeleteSteps map[StepName]Step `json:"reconcileDeleteSteps,omitempty"`
	//corev1.NodeStatus    `json:"nodeStatus,omitempty"`
	InstancePhase InstancePhase `json:"instancePhase,omitempty"`
	LastFinished  *Time         `json:"lastFinished,omitempty"`
	// 最近发起维修时间
	LastStarted *Time `json:"lastStarted,omitempty"`

	Message string `json:"message,omitempty" protobuf:"bytes,3,opt,name=message"`
	Healthy bool   `json:"healthy,omitempty"`
}

// RemedyMachine is the Schema for the remedyMachines API
type RemedyMachine struct {
	TypeMeta   `json:",inline"`
	ObjectMeta `json:"metadata,omitempty"`

	Spec   RemedyMachineSpec   `json:"spec,omitempty"`
	Status RemedyMachineStatus `json:"status,omitempty"`
}

// RemedyMachineList contains a list of RemedyMachines
type RemedyMachineList struct {
	TypeMeta `json:",inline"`
	ListMeta `json:"metadata,omitempty"`
	Items    []RemedyMachine `json:"items"`
}

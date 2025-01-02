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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RemediationTemplateSpec defines the desired state of MachineRemediationTemplate
type RemediationTemplateSpec struct {
	// 节点provider
	NodeProvider string `json:"nodeProvider"`
	// 是否暂停维修
	// +kubebuilder:validation:Optional
	Paused bool `json:"paused,omitempty"`
	// 步骤间最小间隔，每个步骤维修后，需等待大于NPD上报周期的时间后，才能判断是否真正修复
	// MinIntervalByStep int64 `json:"minIntervalByStep,omitempty"`
	// 封锁设置
	// +kubebuilder:validation:Optional
	MachineCordonConfig MachineCordonConfig `json:"machineCordonConfig,omitempty"`
	// 排水设置
	// +kubebuilder:validation:Optional
	DrainConfig DrainConfig `json:"drainConfig,omitempty"`
	// 维修设置
	// +kubebuilder:validation:Optional
	MachineRepairConfig MachineRepairConfig `json:"machineRepairConfig,omitempty"`
	// 机器 重启设置
	// +kubebuilder:validation:Optional
	MachineRebootConfig MachineRebootConfig `json:"machineRebootConfig,omitempty"`
	// 机器 重启设置
	// +kubebuilder:validation:Optional
	MachineAddConfig MachineAddConfig `json:"machineAddConfig,omitempty"`
	// 机器 下线设置
	// +kubebuilder:validation:Optional
	MachineDeleteConfig MachineDeleteConfig `json:"machineDeleteConfig,omitempty"`
	// 机器 重装设置
	// +kubebuilder:validation:Optional
	MachineReInstallConfig MachineReInstallConfig `json:"machineReInstallConfig,omitempty"`
	// 机器 扩容设置
	// +kubebuilder:validation:Optional
	MachineScaleUpConfig MachineScaleUpConfig `json:"machineScaleUpConfig,omitempty"`
	// 机器 检测故障恢复配置
	// +kubebuilder:validation:Optional
	MachineDetectRecoveryConfig MachineDetectRecoveryConfig `json:"machineDetectRecoveryConfig,omitempty"`
	// 维修步骤的通知配置
	// +kubebuilder:validation:Optional
	NotificationConfig NotificationConfig `json:"notificationConfig,omitempty"`
	// 维修步骤
	// +kubebuilder:validation:Optional
	ReconcileSteps []StepName `json:"reconcileSteps,omitempty"`
	// 自定义故障热维修
	// +kubebuilder:validation:Optional
	//ConditionReconcileSteps map[corev1.NodeConditionType]ConditionReconcileStepsConfig `json:"conditionReconcileSteps,omitempty"`
}

type HTTPHeader struct {
	// The header field name.
	// This will be canonicalized upon output, so case-variant names will be understood as the same header.
	Name string `json:"name" protobuf:"bytes,1,opt,name=name"`
	// The header field value
	Value string `json:"value" protobuf:"bytes,2,opt,name=value"`
}

type BaseStepConfig struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:false
	EnableNotification bool `json:"enableNotification,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default:false
	LifeCycle *StepLifecycle `json:"lifeCycle,omitempty"`
}

type StepLifecycle struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:false
	PreExec *StepLifecycleHandler `json:"preExec,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default:false
	PostExec *StepLifecycleHandler `json:"postExec,omitempty"`
}

type LifecycleHandlerType string

const (
	LifecycleHandlerTypePreExec  LifecycleHandlerType = "preExec"
	LifecycleHandlerTypePostExec LifecycleHandlerType = "postExec"
)

// StepLifecycleHandler defines a specific action that should be taken in a lifecycle
// hook. One and only one of the fields.
type StepLifecycleHandler struct {
	// HTTPPost specifies the http request to perform.
	// +kubebuilder:validation:Optional
	HTTPPost *HTTPostAction `json:"httpPost,omitempty"`
}

type IntOrString struct {
	Type   Type   `protobuf:"varint,1,opt,name=type,casttype=Type"`
	IntVal int32  `protobuf:"varint,2,opt,name=intVal"`
	StrVal string `protobuf:"bytes,3,opt,name=strVal"`
}

// Type represents the stored type of IntOrString.
type Type int64

const (
	Int    Type = iota // The IntOrString holds an int.
	String             // The IntOrString holds a string.
)

// HTTPostAction describes an action based on HTTP Get requests.
type HTTPostAction struct {
	// Path to access on the HTTP server.
	// +kubebuilder:validation:Optional
	Path string `json:"path,omitempty"`
	// Name or number of the port to access on the container.
	// Number must be in the range 1 to 65535.
	// Name must be an IANA_SVC_NAME.
	// +kubebuilder:validation:Optional
	Port IntOrString `json:"port,omitempty"`
	// Host name to connect to, defaults to the pod IP. You probably want to set
	// "Host" in httpHeaders instead.
	// +kubebuilder:validation:Optional
	Host string `json:"host,omitempty"`
	// Scheme to use for connecting to the host.
	// Defaults to HTTP.
	// +kubebuilder:validation:Optional
	Scheme string `json:"scheme,omitempty"`
	// Custom headers to set in the request. HTTP allows repeated headers.
	// +kubebuilder:validation:Optional
	HTTPHeaders []HTTPHeader `json:"httpHeaders,omitempty"`
}

type HTTPPostActionPayload struct {
	// ClusterID of k8s cluster
	// +kubebuilder:validation:Optional
	ClusterID string `json:"clusterID,omitempty"`

	// NodeName of k8s node
	// +kubebuilder:validation:Optional
	NodeName string `json:"nodeName,omitempty"`

	// InstanceID of BCC/EBC instance from IaaS
	// +kubebuilder:validation:Optional
	InstanceID string `json:"instanceID,omitempty"`

	// +kubebuilder:validation:Optional
	RemedyTaskName string `json:"remedyTaskName,omitempty"`

	// +kubebuilder:validation:Optional
	RemedyTaskNamespace string `json:"remedyTaskNamespace,omitempty"`

	// StepName of remedy action
	// +kubebuilder:validation:Optional
	StepName string `json:"stepName,omitempty"`

	// HandlerType of executor
	// +kubebuilder:validation:Optional
	HandlerType LifecycleHandlerType `json:"handlerType,omitempty"`
}

type HTTPPostActionResponse struct {
	// +kubebuilder:validation:Optional
	Result bool `json:"result,omitempty"`
	// +kubebuilder:validation:Optional
	Reason string `json:"reason,omitempty"`
}

type ConditionReconcileStepsConfig struct {
	// +kubebuilder:validation:Optional
	Steps []StepName `json:"steps,omitempty"`
	// +kubebuilder:validation:Optional
	MachineRepairContents []string `json:"machineRepairContents,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:shortName=mrtpl

// MachineRemediationTemplate is the Schema for the machineremediationtemplates API
type MachineRemediationTemplate struct {
	//metav1.TypeMeta   `json:",inline"`
	//metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RemediationTemplateSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// MachineRemediationTemplateList contains a list of MachineRemediationTemplate
type MachineRemediationTemplateList struct {
	//metav1.TypeMeta `json:",inline"`
	//metav1.ListMeta `json:"metadata,omitempty"`
	Items []MachineRemediationTemplate `json:"items,omitempty"`
}

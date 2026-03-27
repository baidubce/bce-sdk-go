package types

import (
	"encoding/json"
	"time"
)

type WorkflowType string

const (
	WorkflowTypeUpgradeNodes         WorkflowType = "UpgradeNodes"
	WorkflowTypeUpgradeNodesPreCheck WorkflowType = "UpgradeNodesPreCheck"
	WorkflowTypeUpgradeKubeletConfig WorkflowType = "UpgradeKubeletConfig"
)

type WorkflowPhase string

const (
	WorkflowPhasePending   WorkflowPhase = "Pending"
	WorkflowPhaseUpgrading WorkflowPhase = "Upgrading"
	WorkflowPhaseSucceeded WorkflowPhase = "Succeeded"
	WorkflowPhaseFailed    WorkflowPhase = "Failed"
	WorkflowPhaseVerifying WorkflowPhase = "verifying"
	WorkflowPhaseWarning   WorkflowPhase = "Warning"
	WorkflowPhaseSkipped   WorkflowPhase = "Skipped"
	WorkflowPhaseConfirmed WorkflowPhase = "Confirmed"
	WorkflowPhasePaused    WorkflowPhase = "Paused"
	WorkflowPhaseDeleting  WorkflowPhase = "Deleting"
	WorkflowPhaseDeleted   WorkflowPhase = "Deleted"
	WorkflowPhaseUnknown   WorkflowPhase = "Unknown"
)

type WorkflowTaskType string

const (
	WorkflowTaskTypePause WorkflowTaskType = "Pause"
)

type TaskGroupName string

const (
	TaskGroupNamePreCheck  TaskGroupName = "PreCheck"
	TaskGroupNameBackup    TaskGroupName = "Backup"
	TaskGroupNameOperate   TaskGroupName = "Operate"
	TaskGroupNamePostCheck TaskGroupName = "PostCheck"
)

type PausePolicy string

const (
	NotPause   PausePolicy = "NotPause"
	FirstBatch PausePolicy = "FirstBatch"
	EveryBatch PausePolicy = "EveryBatch"
)

type K8SMemoryManagerPolicy string

const (
	MemoryManagerPolicyNone   K8SMemoryManagerPolicy = "None"
	MemoryManagerPolicyStatic K8SMemoryManagerPolicy = "Static"
)

type MemoryReservation struct {
	NumaNode *int32 `json:"NumaNode,omitempty"`
	Limits   string `json:"Limits,omitempty"`
}

type GPUVersion struct {
	CUDA   string `json:"cuda,omitempty"`
	Driver string `json:"driver,omitempty"`
	CuDNN  string `json:"cuDNN,omitempty"`
}

type ComponentName string

const (
	ComponentKubelet                ComponentName = "kubelet"
	ComponentContainerRuntime       ComponentName = "containerRuntime"
	ComponentNvidiaContainerToolkit ComponentName = "nvidiaContainerToolkit"
	ComponentXPUContainerToolkit    ComponentName = "xpuContainerToolkit"
	ComponentGPUDriver              ComponentName = "gpuDriver"
)

type Component struct {
	Name              ComponentName `json:"name,omitempty"`
	CurrentVersion    string        `json:"currentVersion,omitempty"`
	TargetVersion     string        `json:"targetVersion,omitempty"`
	GPUCurrentVersion *GPUVersion   `json:"gpuCurrentVersion,omitempty"`
	GPUTargetVersion  *GPUVersion   `json:"gpuTargetVersion,omitempty"`
}

type UpgradeNodesWorkflowConfig struct {
	CCEInstanceIDList      []string        `json:"cceInstanceIDList,omitempty"`
	InstanceGroupID        string          `json:"instanceGroupID,omitempty"`
	NodeUpgradeBatchSize   int             `json:"nodeUpgradeBatchSize,omitempty"`
	DrainNodeBeforeUpgrade *bool           `json:"drainNodeBeforeUpgrade,omitempty"`
	PausePolicy            *PausePolicy    `json:"pausePolicy,omitempty"`
	BatchIntervalMinutes   *int            `json:"batchIntervalMinutes,omitempty"`
	Components             []Component     `json:"components,omitempty"`
	InstanceIdToNeedDrain  map[string]bool `json:"instanceIdToNeedDrain,omitempty"`
	IsPreCheck             bool            `json:"isPreCheck,omitempty"`
}

func (c *UpgradeNodesWorkflowConfig) SetInstanceNeedDrain(instanceID string) {
	if c == nil {
		return
	}
	if c.InstanceIdToNeedDrain == nil {
		c.InstanceIdToNeedDrain = make(map[string]bool)
	}
	c.InstanceIdToNeedDrain[instanceID] = true
}

func (c *UpgradeNodesWorkflowConfig) IsInstanceDrain(instanceID string) bool {
	if c == nil || len(c.InstanceIdToNeedDrain) == 0 {
		return false
	}
	return c.InstanceIdToNeedDrain[instanceID]
}

type UpgradeKubeletConfig struct {
	DeployCustomConfig   KubeletDeployCustomConfig `json:"deployCustomConfig,omitempty"`
	InstanceGroupID      string                    `json:"instanceGroupID,omitempty"`
	NodeUpgradeBatchSize int                       `json:"nodeUpgradeBatchSize,omitempty"`
}

type KubeletDeployCustomConfig struct {
	KubeletRootDir              string                   `json:"kubeletRootDir,omitempty"`
	KubeReserved                map[string]string        `json:"kubeReserved,omitempty"`
	SystemReserved              map[string]string        `json:"systemReserved,omitempty"`
	RegistryPullQPS             *int32                   `json:"registryPullQPS,omitempty"`
	RegistryBurst               *int32                   `json:"registryBurst,omitempty"`
	PodPidsLimit                *int64                   `json:"podPidsLimit,omitempty"`
	EventRecordQPS              *int32                   `json:"eventRecordQPS,omitempty"`
	EventBurst                  *int32                   `json:"eventBurst,omitempty"`
	KubeAPIQPS                  *int32                   `json:"kubeAPIQPS,omitempty"`
	KubeAPIBurst                *int32                   `json:"kubeAPIBurst,omitempty"`
	MaxPods                     *int32                   `json:"maxPods,omitempty"`
	ResolvConf                  *string                  `json:"resolvConf,omitempty"`
	AllowedUnsafeSysctls        string                   `json:"allowedUnsafeSysctls,omitempty"`
	SerializeImagePulls         bool                     `json:"serializeImagePulls,omitempty"`
	EvictionHard                map[string]string        `json:"evictionHard,omitempty"`
	EvictionSoft                map[string]string        `json:"evictionSoft,omitempty"`
	EvictionSoftGracePeriod     map[string]string        `json:"evictionSoftGracePeriod,omitempty"`
	ContainerLogMaxFiles        *int32                   `json:"containerLogMaxFiles,omitempty"`
	ContainerLogMaxSize         string                   `json:"containerLogMaxSize,omitempty"`
	FeatureGates                map[string]bool          `json:"featureGates,omitempty"`
	ReadOnlyPort                *int32                   `json:"readOnlyPort,omitempty"`
	CpuCFSQuotaPeriod           *int32                   `json:"cpuCFSQuotaPeriod,omitempty"`
	MemoryManagerPolicy         K8SMemoryManagerPolicy   `json:"memoryManagerPolicy,omitempty"`
	ImageGCHighThresholdPercent *int32                   `json:"imageGCHighThresholdPercent,omitempty"`
	ImageGCLowThresholdPercent  *int32                   `json:"imageGCLowThresholdPercent,omitempty"`
	ReservedMemory              []MemoryReservation      `json:"reservedMemory,omitempty"`
	Runtimerequesttimeout       *int32                   `json:"runtimerequesttimeout,omitempty"`
	CPUManagerPolicy            K8SCPUManagerPolicy      `json:"cpuManagerPolicy,omitempty"`
	TopologyManagerScope        K8STopologyManagerScope  `json:"topologyManagerScope,omitempty"`
	TopologyManagerPolicy       K8STopologyManagerPolicy `json:"topologyManagerPolicy,omitempty"`
	CPUCFSQuota                 *bool                    `json:"cpuCFSQuota,omitempty"`
}

type WorkflowConfig struct {
	UpgradeNodesWorkflowConfig *UpgradeNodesWorkflowConfig `json:"upgradeNodesWorkflowConfig,omitempty"`
	UpgradeKubeletConfig       *UpgradeKubeletConfig       `json:"upgradeKubeletConfig,omitempty"`
}

type WatchDogConfig struct {
	UnhealthyPodsPercent int `json:"unhealthyPodsPercent,omitempty"`
}

type WatchDogStatus struct {
	Healthy           bool       `json:"healthy,omitempty"`
	PausedReason      string     `json:"pausedReason,omitempty"`
	TotalPodCount     int        `json:"totalPodCount,omitempty"`
	UnhealthyPodCount int        `json:"unhealthyPodCount,omitempty"`
	UpdateTime        *time.Time `json:"updateTime,omitempty"`
}

type WorkflowTask struct {
	TaskName          string           `json:"taskName,omitempty"`
	WorkflowTaskType  WorkflowTaskType `json:"workflowTaskType,omitempty"`
	WorkflowTaskPhase WorkflowPhase    `json:"workflowTaskPhase,omitempty"`
	ErrorMessage      string           `json:"errorMessage,omitempty"`
	TaskConfig        json.RawMessage  `json:"taskConfig,omitempty"`
	TaskExecuteResult json.RawMessage  `json:"taskExecuteResult,omitempty"`
	StartTime         *time.Time       `json:"startTime,omitempty"`
	FinishedTime      *time.Time       `json:"finishedTime,omitempty"`
}

type TaskGroup struct {
	TaskGroupName  TaskGroupName   `json:"taskGroupName,omitempty"`
	TaskList       []*WorkflowTask `json:"taskList,omitempty"`
	TaskGroupPhase WorkflowPhase   `json:"taskGroupPhase,omitempty"`
	Pause          *bool           `json:"pause,omitempty"`
}

type WorkflowSpec struct {
	Handler        string         `json:"handler,omitempty"`
	WorkflowID     string         `json:"workflowID,omitempty"`
	ClusterID      string         `json:"clusterID,omitempty"`
	AccountID      string         `json:"accountID,omitempty"`
	UserID         string         `json:"userID,omitempty"`
	Paused         bool           `json:"paused,omitempty"`
	WorkflowType   WorkflowType   `json:"workflowType,omitempty"`
	WorkflowConfig WorkflowConfig `json:"config,omitempty"`
	WatchDogConfig WatchDogConfig `json:"watchDogConfig,omitempty"`
}

type WorkflowStatus struct {
	StartTime         *time.Time     `json:"startTime,omitempty"`
	FinishedTime      *time.Time     `json:"finishedTime,omitempty"`
	WorkflowPhase     WorkflowPhase  `json:"phase,omitempty"`
	TotalTaskCount    int            `json:"totalTaskCount"`
	FinishedTaskCount int            `json:"finishedTaskCount"`
	TaskGroupList     []*TaskGroup   `json:"taskGroupList,omitempty"`
	RetryCount        int            `json:"retryCount,omitempty"`
	ErrorMessage      string         `json:"errorMessage,omitempty"`
	WatchDogStatus    WatchDogStatus `json:"watchDogStatus,omitempty"`
}

type UpgradeComponents struct {
	Kubelet                UpgradeVersionList `json:"kubelet"`
	ContainerRuntime       UpgradeVersionList `json:"containerRuntime"`
	NvidiaContainerToolkit UpgradeVersionList `json:"nvidiaContainerToolkit"`
	XPUContainerToolkit    UpgradeVersionList `json:"xpuContainerToolkit"`
	GPUDriver              UpgradeVersionList `json:"gpuDriver"`
}

type UpgradeVersionList struct {
	CurrentVersion    string             `json:"currentVersion"`
	ComponentVersions []ComponentVersion `json:"componentVersions"`
	GPUCurrentVersion *GPUVersion        `json:"gpuCurrentVersion"`
}

type ComponentVersion struct {
	TargetVersion    string      `json:"targetVersion"`
	NeedDrainNode    bool        `json:"needDrainNode"`
	GPUTargetVersion *GPUVersion `json:"gpuTargetVersion,omitempty"`
}

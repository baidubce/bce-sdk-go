package v1

// Step 定义 Reconcile/ReconcileDelete 过程中耗时和状态
/**
1. 发起步骤 StartTime，Ready为false
2. 步骤成功
*/
type Step struct {
	Finish       bool          `json:"finish"`                 // 整个 step 完成，包含成功或者失败
	Successful   bool          `json:"successful"`             // 步骤是否成功，比如重启，失败或者成功
	StepStatus   InstancePhase `json:"stepStatus"`             // 当前步骤状态
	StartTime    *Time         `json:"startTime,omitempty"`    // 第一次开始时间
	FinishedTime *Time         `json:"finishedTime,omitempty"` // 最后一次成功时间
	CostSeconds  int           `json:"costSeconds,omitempty"`  // 花费时间
	RetryCount   int           `json:"retryCount,omitempty"`   // 重试次数
	ErrMsg       string        `json:"errMsg,omitempty"`       // 失败信息
	NeedCheck    bool          `json:"needCheck,omitempty"`    // 已经发起请求，需要checking执行状态
	CheckingTime *Time         `json:"checkingTime,omitempty"` // 最近一次检查时间

	// task 授权结果
	AuthorizedStatus *TaskAuthorizedStatus `json:"authorizedStatus,omitempty"`
	// 用户确认机器恢复
	UserConfirmedStatus *UserConfirmedStatus `json:"userConfirmedStatus,omitempty"`

	// 添加节点状态
	AddNodeStatus *AddNodeStatus `json:"addNodeStatus,omitempty"`
	// 删除节点状态
	DeleteNodeStatus *DeleteNodeStatus `json:"deleteNodeStatus,omitempty"`
}

type TaskAuthorizedStatus struct {
	Authorized      bool  `json:"authorized"`
	LastRequestTime *Time `json:"lastRequestTime,omitempty"`
}

type UserConfirmedStatus struct {
	Confirmed bool `json:"confirmed"`
}

type AddNodeStatus struct {
	AddInstanceID     string `json:"addInstanceID,omitempty"`
	AddInstanceStatus string `json:"addInstanceStatus,omitempty"`
	ErrMsg            string `json:"errMsg,omitempty"`
}

type DeleteNodeStatus struct {
	DeleteInstanceID     string `json:"deleteInstanceID,omitempty"`
	DeleteInstanceStatus string `json:"deleteInstanceStatus,omitempty"`
	ErrMsg               string `json:"errMsg,omitempty"`
}

// StepName 定义操作步骤名字
type StepName string

// 节点维修
const (
	RemedyMachineStepCordonNode              StepName = "CordonNode"
	RemedyMachineStepCordon                  StepName = "Cordon"
	RemedyMachineStepUnCordonNode            StepName = "UnCordonNode"
	RemedyMachineStepUnCordon                StepName = "UnCordon"
	RemedyMachineStepDrainNode               StepName = "DrainNode"
	RemedyMachineStepDrain                   StepName = "Drain"
	RemedyMachineStepRebootNode              StepName = "RebootNode"
	RemedyMachineStepReboot                  StepName = "Reboot"
	RemedyMachineStepDelete                  StepName = "Delete"
	RemedyMachineStepDeleteNode              StepName = "DeleteNode"
	RemedyMachineStepReinstallNode           StepName = "ReInstallNode"
	RemedyMachineStepReinstall               StepName = "ReInstall"
	RemedyMachineStepRestartKubelet          StepName = "RestartKubelet"
	RemedyMachineStepRestartContainerRuntime StepName = "RestartContainerRuntime"
	RemedyMachineStepRepairNode              StepName = "RepairNode"
	RemedyMachineStepDetectRecovery          StepName = "DetectRecovery"
	RemedyMachineStepAddNode                 StepName = "AddNode"
)

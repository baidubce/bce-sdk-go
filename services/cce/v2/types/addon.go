package types

// AddOnInfo 表示集群中某个组件的完整信息 分为元信息和安装实例信息(如果有)
// 如果组件允许多实例部署,所以已安装实例信息同时会放在 MultiInstances 字段中
type AddOnInfo struct {
	// 组件的元信息
	Meta Meta `json:"meta,omitempty"`

	// 组件的安装实例信息 某些组件可以多实例在集群中安装 如 Ingress NGINX Controller
	Instance *AddOnInstance `json:"instance,omitempty"`

	// 组件的安装实例信息 某些组件可以多实例在集群中安装 如 Ingress NGINX Controller
	// 多实例安装的组件会把实例信息都展示在这里
	// 如果组件仅有一个部署实例 这个字段不会返回
	MultiInstances []*AddOnInstance `json:"multiInstances,omitempty"`
}

// Meta 表示集群中某个组件的元信息 包括组件名称、类型、最新版本、简介与详细介绍、默认参数
type Meta struct {
	// 组件名称
	Name string `json:"name,omitempty"`

	// 组件所述的类型: AI/网络 等
	Type AddOnType `json:"type,omitempty"`

	//是否是托管组件
	Managed bool `json:"managed,omitempty"`

	//是否是系统组件
	Required bool `json:"required,omitempty"`

	// 组件的最新版本
	LatestVersion string `json:"latestVersion,omitempty"`

	LatestImageTag string `json:"latestImageTag,omitempty"`

	// 组件简要介绍 一句话介绍
	ShortIntroduction string `json:"shortIntroduction,omitempty"`

	// 组件的详细介绍
	DetailIntroduction string `json:"detailIntroduction,omitempty"`

	// 默认参数 用户可在此基础上编辑
	DefaultParams string `json:"defaultParams,omitempty"`

	// 集群是否满足组件安装条件
	InstallInfo InstallInfo `json:"installInfo,omitempty"`
}

// AddOnInstance 表示集群中某个组件的实例部署信息 包括实例名称、安装版本、部署参数、当前状态以及卸载、更新参数、升级方面的信息
type AddOnInstance struct {
	// 组件实例的名称
	AddOnInstanceName string `json:"name,omitempty"`

	// 当前组件实例的版本
	InstalledVersion string `json:"installedVersion,omitempty"`

	// 当前组件实例的参数
	Params string `json:"params,omitempty"`

	// Status 组件当前的状态
	Status AddonInstanceStatus `json:"status,omitempty"`

	// UninstallInfo 该组件实例卸载方面的信息(能否卸载,原因)
	UninstallInfo UninstallInfo `json:"uninstallInfo,omitempty"`

	// UpgradeInfo 该组件实例升级方面的信息(能否升级,原因)
	UpgradeInfo UpgradeInfo `json:"upgradeInfo,omitempty"`

	// UpdateInfo 该组件实例更新方面的信息(能否升级,原因)
	UpdateInfo UpdateInfo `json:"updateInfo,omitempty"`

	//RollbackInfo 该组件实例是否支持回滚 默认不支持
	RollbackInfo RollbackInfo `json:"rollback,omitempty"`

	// 组件实例的部署yaml
	Manifest string `json:"manifest,omitempty"`
}

type AddonInstanceStatus struct {
	Phase AddOnInstancePhase `json:"phase,omitempty"`

	// 如果组件实例有异常 这里尽可能返回详细的信息
	Code    string `json:"code,omitempty"`
	TraceID string `json:"traceID,omitempty"`
	Message string `json:"message,omitempty"`
}

type InstallInfo struct {
	AllowInstall bool   `json:"allowInstall"`      // 是否允许安装
	Message      string `json:"message,omitempty"` // 原因解释 或是其他想返回给用户的信息
}

type UninstallInfo struct {
	AllowUninstall bool   `json:"allowUninstall"`    // 是否允许卸载
	Message        string `json:"message,omitempty"` // 原因解释 或是其他想返回给用户的信息
}

type UpgradeInfo struct {
	AllowUpgrade bool   `json:"allowUpgrade"`      // 是否允许升级
	NextVersion  string `json:"nextVersion"`       // 若允许升级 升级到哪个版本
	Message      string `json:"message,omitempty"` // 原因解释 或是其他想返回给用户的信息
}

type UpdateInfo struct {
	AllowUpdate bool   `json:"allowUpdate"`       // 是否允许更新参数
	Message     string `json:"message,omitempty"` // 原因解释 或是其他想返回给用户的信息
}

type RollbackInfo struct {
	AllowRollback bool   `json:"allowRollback"`     // 是否允许回滚
	Message       string `json:"message,omitempty"` // 原因解释 或是其他想返回给用户的信息
}

type AddOnType string

const (
	TypeCloudNativeAI AddOnType = "CloudNativeAI"

	TypeNetworking AddOnType = "Networking"

	TypeHybridSchedule AddOnType = "HybridSchedule"

	TypeImage AddOnType = "Image"

	TypeStorage AddOnType = "Storage"

	TypeObservability AddOnType = "Observability"

	TypeOthers = "Others"
)

type AddOnInstancePhase string

const (
	AddOnInstancePhaseRunning      AddOnInstancePhase = "Running"
	AddOnInstancePhaseAbnormal     AddOnInstancePhase = "Abnormal"
	AddOnInstancePhaseInstalling   AddOnInstancePhase = "Installing"
	AddOnInstancePhaseUninstalling AddOnInstancePhase = "Uninstalling"
	AddOnInstancePhaseUpgrading    AddOnInstancePhase = "Upgrading"
	AddOnInstancePhaseFailed       AddOnInstancePhase = "Failed"
	AddOnInstancePhaseDeleting     AddOnInstancePhase = "Deleting"
)

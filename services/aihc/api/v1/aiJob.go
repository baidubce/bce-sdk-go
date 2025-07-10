package v1

type OpenAPIGetJobResponse struct {
	Result    OpenAPIGetJobDetailResponseResult `json:"result"`
	RequestId string                            `json:"requestId"`
}

type OpenAPIGetJobDetailResponseResult struct {
	OpenAPIGetJobResponseResult
	PodList *OpenAPIPodList `json:"podList"`
}

type OpenAPIGetJobResponseResult struct {
	JobID                       string              `json:"jobId"`
	Name                        string              `json:"name"`
	ResourcePoolID              string              `json:"resourcePoolId"`
	Command                     string              `json:"command"`
	CreatedAt                   string              `json:"createdAt"`
	FinishedAt                  string              `json:"finishedAt"`
	RunningAt                   string              `json:"runningAt"`
	ScheduledAt                 string              `json:"scheduledAt"`
	Datasources                 []OpenAPIDatasource `json:"datasources"`
	EnableFaultTolerance        bool                `json:"enableFaultTolerance"`
	CustomFaultTolerancePattern []string            `json:"customFaultTolerancePattern"`
	Labels                      []OpenAPILabel      `json:"labels"`
	Priority                    string              `json:"priority"`
	Queue                       string              `json:"queue"`
	Status                      string              `json:"status"`
	Image                       string              `json:"image"`
	Resources                   []OpenAPIResource   `json:"resources"`
	EnableRDMA                  bool                `json:"enableRDMA"`
	HostNetwork                 bool                `json:"hostNetwork"`
	Privileged                  *bool               `json:"privileged,omitempty"`
	Replicas                    int32               `json:"replicas"`
	Envs                        []OpenAPIEnv        `json:"envs"`
	JobFramework                string              `json:"jobFramework"`
	QueueingSequence            *int                `json:"queueingSequence,omitempty"`
	EnableBccl                  bool                `json:"enableBccl"`
	EnableBcclStatus            string              `json:"enableBcclStatus"`
	EnableBcclErrorReason       string              `json:"enableBcclErrorReason"`
	K8sUID                      string              `json:"k8sUID"`
	K8sNamespace                string              `json:"k8sNamespace"`
	FaultToleranceArgs          string              `json:"faultToleranceArgs,omitempty"`
	LogCollectionFilePatterns   []string            `json:"logCollectionFilePatterns,omitempty"`
}

type OpenAPIEnv struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type OpenAPILabel struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type OpenAPIResource struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
}

type OpenAPIPodList struct {
	ListMeta OpenAPIPodListMeta `json:"listMeta"`
	Pods     []OpenAPIPod       `json:"pods"`
}

type OpenAPIPodListMeta struct {
	TotalItems int `json:"totalItems"`
}

type OpenAPIPod struct {
	PodIP        string            `json:"PodIP"`
	NodeName     string            `json:"nodeName"`
	ObjectMeta   OpenAPIObjectMeta `json:"objectMeta"`
	PodStatus    OpenAPIPodStatus  `json:"podStatus"`
	ReplicaType  string            `json:"replicaType"`
	RestartCount int32             `json:"restartCount"`
	Envs         []OpenAPIEnv      `json:"envs"`
	FinishedAt   string            `json:"finishedAt"`
	Reason       string            `json:"reason"`
}

type OpenAPIDatasource struct {
	// Type 数据源类型 dataset/pfsl1/pfsl2/emptydir/hostpath/cfs
	Type string `json:"type"`
	// SourcePath为源路径，存储类型为pfs时，代表用户传入的pfs路径；存储类型为hostPath时，代表节点上的本地路径
	SourcePath string `json:"sourcePath,omitempty"`

	// MountPath代表挂载到容器内的路径
	MountPath string `json:"mountPath"`

	// 当数据源类型为pfsl1或pfsl2时，name为对应的pfsId
	Name    string                 `json:"name"`
	Options AIJobDatasourceOptions `json:"options"`
}

type AIJobDatasourceOptions struct {
	// SizeLimit emptydir 限制
	SizeLimit int `json:"sizeLimit"`

	// Medium emptydir 介质
	Medium StorageMedium `json:"medium"`

	HostPath     string       `json:"hostPath,omitempty"`
	HostPathType HostPathType `json:"hostPathType,omitempty"`

	// ReadOnly 是否以只读挂载，默认 false 非只读
	ReadOnly bool `json:"readOnly"`

	// PFS的连接地址，仅限L1
	PfsL1ClusterIP string `yaml:"pfsL1ClusterIp" json:"pfsL1ClusterIp,omitempty"`
	// PFS连接端口，仅限L1
	PfsL1ClusterPort string `yaml:"pfsL1ClusterPort" json:"pfsL1ClusterPort,omitempty"`
	// PV的根目录，仅限L1
	PfsL1ParentDir string `yaml:"pfsL1ParentDir" json:"pfsL1ParentDir,omitempty"`
	// PV的子目录，与PFSParentDir合并为PV完整路径，仅L1使用，
	PFSPath string `yaml:"pfsPath" json:"pfsPath,omitempty"`
	// Medium emptydir 介质
	PfsL2MountTargetID []string `json:"pfsL2MountTargetId,omitempty"`
	PfsL2HostMountPath string   `json:"pfsL2HostMountPath,omitempty"`
	//cfs 挂载点信息
	CfsInstanceID string `json:"cfsInstanceId"`
	CfsMountPoint string `json:"cfsMountPoint"`
}

type OpenAPIDatasourceOptions struct {
	Medium    StorageMedium `json:"medium"`
	SizeLimit int           `json:"sizeLimit"`
	ReadOnly  bool          `json:"readOnly"`
}

type OpenAPIObjectMeta struct {
	Annotations       map[string]string `json:"annotations"`
	CreationTimestamp string            `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	OwnerReferences   []OwnerReference  `json:"ownerReferences"`
}

type OpenAPIPodStatus struct {
	PodPhase PodPhase `json:"podPhase"`
	Status   string   `json:"status"`
}

// delete Job
type OpenAPIJobDeleteResponse struct {
	RequestID string                         `json:"requestId"`
	Result    OpenAPIJobDeleteResponseResult `json:"result"`
}

type OpenAPIJobDeleteResponseResult struct {
	JobID   string `json:"jobId"`
	JobName string `json:"jobName"`
}

// create Job
type OpenAPIJobCreateRequest struct {
	Name           string              `json:"name"`
	Queue          string              `json:"queue"`
	JobFramework   string              `json:"jobFramework"`
	JobSpec        OpenAPIAIJobSpec    `json:"jobSpec"`
	FaultTolerance bool                `json:"faultTolerance"`
	Labels         []OpenAPILabel      `json:"labels"`
	Priority       string              `json:"priority"`
	Datasources    []OpenAPIDatasource `json:"datasources"`
	// 两者选其一
	FaultToleranceConfig *OpenAPIJobFaultToleranceConfig `json:"faultToleranceConfig,omitempty"`
	FaultToleranceArgs   *string                         `json:"faultToleranceArgs,omitempty"`
	CodeSource           CodeSourceV3                    `json:"codeSource"`
	// 创建任务告警规则对应的cprom 实例id

	AlertConfig *AlertRuleReq `json:"alertConfig"` // 告警详情信息列表

	EnableBccl bool `json:"enableBccl"` // 是否开启 bccl 注入
	// 容器日志采集文件路径
	LogCollectionFilePatterns []string `json:"logCollectionFilePatterns"`
}

type CodeSourceV3 struct {
	FilePath          string `json:"filePath"`
	MountPath         string `json:"mountPath"`
	ID                string `json:"id"`
	BosObjectName     string `json:"bosObjectName"`
	BosTemporaryToken string `json:"bosTemporaryToken"`
	BosEndpoint       string `json:"bosEndpoint"`
	BosBucket         string `json:"bosBucket"`
}

type OpenAPIAIJobSpec struct {
	Command     string             `json:"command"`
	Image       string             `json:"image"`
	ImageConfig OpenAPIImageConfig `json:"imageConfig"`
	Replicas    int32              `json:"replicas"`
	Resources   []OpenAPIResource  `json:"resources"`
	Envs        []OpenAPIEnv       `json:"envs"`
	EnableRDMA  bool               `json:"enableRDMA"`
	HostNetwork *bool              `json:"hostNetwork,omitempty"`
	Privileged  *bool              `json:"privileged,omitempty"`
}

type OpenAPIImageConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type OpenAPIJobFaultToleranceConfig struct {
	EnabledHangDetection        bool     `json:"enabledHangDetection"`
	HangDetectionTimeoutMinutes int      `json:"hangDetectionTimeoutMinutes"`
	FaultToleranceLimit         int      `json:"faultToleranceLimit"`
	CustomFaultTolerancePattern []string `json:"customFaultTolerancePattern"`
}

type AlertRuleReq struct {
	AlertIDs        []string `json:"alertIds"`
	CpromInstanceID string   `json:"instanceId"`
	AlertName       string   `json:"alertName"   `
	AlertItems      []string `json:"alertItems"        `
	For             string   `json:"for"          `
	Description     string   `json:"description"  `
	NotifyRuleID    string   `json:"notifyRuleId" `
	Severity        string   `json:"severity,omitempty"`
}

type OpenAPIJobCreateResponse struct {
	RequestID string                         `json:"requestId"`
	Result    OpenAPIJobCreateResponseResult `json:"result"`
}

type OpenAPIJobCreateResponseResult struct {
	JobID   string `json:"jobId"`
	JobName string `json:"jobName"`
}

//list Job

type OpenAPIJobListRequest struct {
	ResourcePoolID string `json:"resourcePoolId"`
	Queue          string `json:"queue"`
	OrderBy        string `json:"orderBy"`
	Order          string `json:"order"`
	PageNo         int    `json:"pageNo"`
	PageSize       int    `json:"pageSize"`
}

type OpenAPIJobListResponse struct {
	RequestId string                       `json:"requestId"`
	Result    OpenAPIJobListResponseResult `json:"result"`
}

type OpenAPIJobListResponseResult struct {
	Total int                           `json:"total"`
	Jobs  []OpenAPIGetJobResponseResult `json:"jobs"`

	OrderBy  string `json:"orderBy"`
	Order    string `json:"order"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
}

// update Job
type OpenAPIJobUpdateRequest struct {
	Priority string `json:"priority"`
}

type OpenAPIJobUpdateResponse struct {
	RequestID string                         `json:"requestId"`
	Result    OpenAPIJobUpdateResponseResult `json:"result"`
}

type OpenAPIJobUpdateResponseResult struct {
	JobID string `json:"jobId"`
}

// stop Job
type OpenAPIJobStopResponse struct {
	RequestID string                       `json:"requestId"`
	Result    OpenAPIJobStopResponseResult `json:"result"`
}

type OpenAPIJobStopResponseResult struct {
	JobID string `json:"jobId"`
}

// get job nodes lists
type JobNodesListResponse struct {
	RequestId string             `json:"requestId"`
	Result    JobNodesListResult `json:"result"`
}
type JobNodesListResult struct {
	JobID     string   `json:"jobId"`
	NodeNames []string `json:"nodeName"`
}

//get job events

type GetJobEventsRequest struct {
	Namespace      string `json:"nameSpace"`
	JobFramework   string `json:"jobFramework"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	JobID          string `json:"jobId"`
	ResourcePoolID string `json:"resourcePoolId"`
}

type Event struct {
	Reason         string `json:"reason"`
	Message        string `json:"message"`
	FirstTimestamp string `json:"firstTimestamp"`
	LastTimestamp  string `json:"lastTimestamp"`
	Count          int32  `json:"count"`
	Type           string `json:"type"`
}

type GetJobEventsResponse struct {
	Events []Event `json:"events"`
	Total  int     `json:"total"`
}

// get pod events
type GetPodEventsRequest struct {
	JobID          string `json:"jobId"`
	Namespace      string `json:"nameSpace"`
	JobFramework   string `json:"jobFramework"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	ResourcePoolID string `json:"resourcePoolId"`
	PodName        string `json:"podName"`
}

type GetPodEventsResponse struct {
	Events []Event `json:"events"`
	Total  int     `json:"total"`
}

// get job logs
type GetPodLogsRequest struct {
	JobID          string `json:"jobId"`
	ResourcePoolID string `json:"resourcePoolId"`
	PodName        string `json:"podName"`
	Namespace      string `json:"namespace"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	MaxLines       string `json:"maxLines"`
	Container      string `json:"container"`
	Chunk          string `json:"chunck"`
	Marker         string `json:"marker"`
	FilePath       string `json:"filePath"`
	LogSource      string `json:"logSource"`
}

type GetPodLogResponse struct {
	RequestID string       `json:"requestId"`
	Result    PodLogResult `json:"result"`
}

type PodLogResult struct {
	JobID      string   `json:"jobId"`
	PodName    string   `json:"podName"`
	Logs       []string `json:"logs"`
	NextMarker string   `json:"nextMarker"`
}

// get task metrics
type GetTaskMetricsRequest struct {
	StartTime      string `json:"startTime"`
	ResourcePoolID string `json:"resourcePoolId"`
	EndTime        string `json:"endTime"`
	TimeStep       string `json:"timeStep"`
	MetricType     string `json:"metricType"`
	JobID          string `json:"jobId"`
	Namespace      string `json:"namespace"`
	RateInterval   string `json:"rateInterval"`
}

type GetTaskMetricsResponse struct {
	RequestID string                `json:"requestId"`
	Result    TaskPromMetricsResult `json:"result"`
}

type TaskPromMetricsResult struct {
	JobID      string          `json:"jobId"`
	PodMetrics []PodPromMetric `json:"podMetrics"`
}

type PodPromMetric struct {
	PodName string   `json:"podName"`
	Metrics []Metric `json:"metrics"`
}

type Metric struct {
	Time  interface{} `json:"time"`
	Value interface{} `json:"value"`
}

// get webshell url
type GetWebShellURLRequest struct {
	JobID                  string `json:"jobId"`
	ResourcePoolID         string `json:"resourcePoolId"`
	PodName                string `json:"podName"`
	Namespace              string `json:"namespace"`
	PingTimeoutSecond      string `json:"pingTimeoutSecond"`
	HandshakeTimeoutSecond string `json:"handshakeTimeoutSecond"`
	QueueID                string `json:"queueId"`
}

// Result 包含WebTerminalUrl字段
type SSHResult struct {
	WebTerminalURL string `json:"WebTerminalUrl"`
}

// Response 用于接收JSON响应
type GetWebShellURLResponse struct {
	RequestID string    `json:"requestId"`
	Result    SSHResult `json:"result"`
}

// create alert policy
type CreateNotifyRuleReq struct {
	NotifyRuleName    string          `json:"notifyRuleName"`
	StartTime         string          `json:"startTime"`
	EndTime           string          `json:"endTime"`
	ReceiverType      string          `json:"receiverType"`
	Channel           []string        `json:"channel"`
	Users             []User          `json:"users"`
	UserGroups        []UserGroup     `json:"userGroups"`
	Enable            bool            `json:"enable"`
	EnableCallback    bool            `json:"enableCallback"`
	WebhookConfigList []WebhookConfig `json:"webhookConfigList"`
}

type WebhookConfig struct {
	WebhookType string        `json:"webhookType"`
	WebhookList []WebhookItem `json:"webhookList"`
}

type WebhookItem struct {
	HookName   string      `json:"hookName"`
	HookMethod string      `json:"hookMethod"`
	HookURL    string      `json:"hookUrl"`
	Headers    interface{} `json:"headers"`
	Params     interface{} `json:"params"`
}

type UserGroup struct {
	GroupID     string `json:"groupId"`     // 用户组ID
	GroupName   string `json:"groupName"`   // 用户组名称
	Description string `json:"description"` // 用户组描述
}

type User struct {
	UserID      string `json:"userId"`      // 用户ID
	UserName    string `json:"userName"`    // 用户名称
	UserType    string `json:"userType"`    // 用户类型
	PhoneNumber string `json:"phoneNumber"` // 用户手机号码
	Email       string `json:"email"`       // 用户邮箱
}

type CreateNotifyRuleResp struct {
	NotifyRuleID string `json:"notifyRuleId"` // 通知策略列表
}

// file upload
type FileUploadRequest struct {
	FilePaths      []string `json:"filePath"`
	ResourcePoolID string   `json:"resourcePoolId"`
}

// FileUploadResponse 用于接收JSON响应
type FileUploaderResponse struct {
	RequestId string       `json:"requestId"`
	Result    FileUploader `json:"result"`
}

type FileUploader struct {
	FilePath string `json:"filePath"`
	Token    string `json:"token"`
	FileID   string `json:"fileID"`
	AK       string `json:"ak"`
	SK       string `json:"sk"`
	Bucket   string `json:"bucket"`
	Endpoint string `json:"endpoint"`
}

// k8s
type HostPathType string

const (
	// For backwards compatible, leave it empty if unset
	HostPathUnset HostPathType = ""
	// If nothing exists at the given path, an empty directory will be created there
	// as needed with file mode 0755, having the same group and ownership with Kubelet.
	HostPathDirectoryOrCreate HostPathType = "DirectoryOrCreate"
	// A directory must exist at the given path
	HostPathDirectory HostPathType = "Directory"
	// If nothing exists at the given path, an empty file will be created there
	// as needed with file mode 0644, having the same group and ownership with Kubelet.
	HostPathFileOrCreate HostPathType = "FileOrCreate"
	// A file must exist at the given path
	HostPathFile HostPathType = "File"
	// A UNIX socket must exist at the given path
	HostPathSocket HostPathType = "Socket"
	// A character device must exist at the given path
	HostPathCharDev HostPathType = "CharDevice"
	// A block device must exist at the given path
	HostPathBlockDev HostPathType = "BlockDevice"
)

// StorageMedium defines ways that storage can be allocated to a volume.
type StorageMedium string

const (
	StorageMediumDefault         StorageMedium = ""           // use whatever the default is for the node, assume anything we don't explicitly handle is this
	StorageMediumMemory          StorageMedium = "Memory"     // use memory (e.g. tmpfs on linux)
	StorageMediumHugePages       StorageMedium = "HugePages"  // use hugepages
	StorageMediumHugePagesPrefix StorageMedium = "HugePages-" // prefix for full medium notation HugePages-<size>
)

type OwnerReference struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion" protobuf:"bytes,5,opt,name=apiVersion"`
	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind" protobuf:"bytes,1,opt,name=kind"`
	// Name of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#names
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`
	// UID of the referent.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names#uids
	UID string `json:"uid" protobuf:"bytes,4,opt,name=uid,casttype=k8s.io/apimachinery/pkg/types.UID"`
	// If true, this reference points to the managing controller.
	// +optional
	Controller *bool `json:"controller,omitempty" protobuf:"varint,6,opt,name=controller"`
	// If true, AND if the owner has the "foregroundDeletion" finalizer, then
	// the owner cannot be deleted from the key-value store until this
	// reference is removed.
	// See https://kubernetes.io/docs/concepts/architecture/garbage-collection/#foreground-deletion
	// for how the garbage collector interacts with this field and enforces the foreground deletion.
	// Defaults to false.
	// To set this field, a user needs "delete" permission of the owner,
	// otherwise 422 (Unprocessable Entity) will be returned.
	// +optional
	BlockOwnerDeletion *bool `json:"blockOwnerDeletion,omitempty" protobuf:"varint,7,opt,name=blockOwnerDeletion"`
}

type PodPhase string

// These are the valid statuses of pods.
const (
	// PodPending means the pod has been accepted by the system, but one or more of the containers
	// has not been started. This includes time before being bound to a node, as well as time spent
	// pulling images onto the host.
	PodPending PodPhase = "Pending"
	// PodRunning means the pod has been bound to a node and all of the containers have been started.
	// At least one container is still running or is in the process of being restarted.
	PodRunning PodPhase = "Running"
	// PodSucceeded means that all containers in the pod have voluntarily terminated
	// with a container exit code of 0, and the system is not going to restart any of these containers.
	PodSucceeded PodPhase = "Succeeded"
	// PodFailed means that all containers in the pod have terminated, and at least one container has
	// terminated in a failure (exited with a non-zero exit code or was stopped by the system).
	PodFailed PodPhase = "Failed"
	// PodUnknown means that for some reason the state of the pod could not be obtained, typically due
	// to an error in communicating with the host of the pod.
	// Deprecated: It isn't being set since 2015 (74da3b14b0c0f658b3bb8d2def5094686d0e9095)
	PodUnknown PodPhase = "Unknown"
)

type GetAIJobOptions struct {
	JobID          string
	ResourcePoolID string
	QueueID        string
}

type DeleteAIJobOptions struct {
	JobID          string
	ResourcePoolID string
	QueueID        string
}

type CreateAIJobOptions struct {
	ResourcePoolID string
}

type UpdateAIJobOptions struct {
	JobID          string
	ResourcePoolID string
	QueueID        string
}

type StopAIJobOptions struct {
	JobID          string
	ResourcePoolID string
	QueueID        string
}

type GetJobNodesListOptions struct {
	JobID          string
	ResourcePoolID string
	Namespace      string
}

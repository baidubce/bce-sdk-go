package api

type BaseRes struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type ServiceConf struct {
	Name            string           `json:"name"`
	AcceleratorType string           `json:"acceleratorType"`
	InstanceCount   int32            `json:"instanceCount"`
	WorkloadType    string           `json:"workloadType"` // "" | fed
	ResourcePool    ResourcePoolConf `json:"resourcePool"`
	Storage         StorageConf      `json:"storage"`
	Containers      []ContainerConf  `json:"containers"`
	Access          AccessConf       `json:"access"`
	Log             LogConf          `json:"log"`
	Deploy          DeployConf       `json:"deploy"`
	Misc            Misc             `json:"misc"`
	Hpa             HPAConf          `json:"hpa,omitempty"`
}

type HPAConf struct {
	TimeBased   *TimeBasedScaleConf   `json:"timeBased"`
	MetricBased *MetricBasedScaleConf `json:"metricBased"`
}
type TimeBasedScaleConf struct {
	Enable bool                 `json:"enable"`
	Items  []TimeBasedScaleInfo `json:"items"`
}

type TimeBasedScaleInfo struct {
	DesiredReplicas int32  `json:"desiredReplicas"`
	Type            string `json:"type"`
	Hour            int32  `json:"hour"`
	Minute          int32  `json:"minute"`
	Second          int32  `json:"second"`
	DayOfWeek       int32  `json:"dayOfWeek"`
	DayOfMonth      int32  `json:"dayOfMonth"`
	Cron            string `json:"cron"`
}

type MetricBasedScaleConf struct {
	Enable           bool            `json:"enable"`
	MinInstances     int32           `json:"minInstances"`
	MaxInstances     int32           `json:"maxInstances"`
	ScaleIndicators  ScaleIndicators `json:"scaleIndicators"`
	ScaleOutDuration int32           `json:"scaleOutDuration"`
	ScaleInDuration  int32           `json:"scaleInDuration"`
}

type ScaleIndicators struct {
	DefaultIndicators map[string]int32   `json:"defaultIndicators"`
	CustomIndicators  map[string]float32 `json:"customIndicators"`
}

type CreateServiceResult struct {
	ServiceId string `json:"serviceId"`
	BaseRes
}

type ListServiceArgs struct {
	PageNumber int32  `json:"pageNumber"`
	PageSize   int32  `json:"pageSize"`
	OrderBy    string `json:"orderBy"`
	Order      string `json:"order"`
	Source     string `json:"source"`
}

type ListServiceResult struct {
	Services   []ServiceBriefInfo `json:"services"`
	TotalCount int32              `json:"totalCount"`
	PageNumber int32              `json:"pageNumber"`
	PageSize   int32              `json:"pageSize"`
	OrderBy    string             `json:"orderBy"`
	Order      string             `json:"order"`
	BaseRes
}

type ListServiceStatsArgs struct {
	ServiceId string `json:"serviceId"`
}

type ListServiceStatsResult struct {
	Service map[string]ServiceBriefStat `json:"service"`
	BaseRes
}

type ServiceDetailsArgs struct {
	ServiceId string `json:"serviceId"`
}

type ServiceDetailsResult struct {
	Service   ServiceConf   `json:"service"`
	Status    ServiceStatus `json:"status"`
	Instances []InsInfo     `json:"instances"`
	Creator   string        `json:"creator"`
	CreatedAt uint32        `json:"createdAt"`
	UpdatedAt uint32        `json:"updatedAt"`
	BaseRes
}

type UpdateServiceArgs struct {
	ServiceId   string      `json:"serviceId"`
	ServiceConf ServiceConf `json:"serviceConf"`
	Description string      `json:"description"`
}

type UpdateServiceResult struct {
	ServiecId string `json:"serviceId"`
	BaseRes
}

type ScaleServiceArgs struct {
	ServiceId     string `json:"serviceId"`
	InstanceCount int32  `json:"instanceCount"`
}

type ScaleServiceResult struct {
	BaseRes
}

type PubAccessArgs struct {
	ServiceId    string `json:"serviceId"`
	PublicAccess bool   `json:"publicAccess"`
	Eip          string `json:"eip"`
}

type PubAccessResult struct {
	BaseRes
}

type EIPInfo struct {
	Name            string `json:"name"`
	Eip             string `json:"eip"`
	EipId           string `json:"eipId"`
	Status          string `json:"status"`
	EipInstanceType string `json:"eipInstanceType"`
	InstanceType    string `json:"instanceType"`
	InstanceId      string `json:"instanceId"`
	ClusterId       string `json:"clusterId"`
	BandwidthInMbps int32  `json:"bandwidthInMbps"`
	PaymentTiming   string `json:"paymentTiming"`
	BillingMethod   string `json:"billingMethod"`
	CreatedAt       string `json:"createdAt"`
	ExpiredAt       string `json:"expiredAt"`
}
type ListChangeArgs struct {
	ServiceId  string `json:"serviceId"`
	ChangeType int32  `json:"changeType"`
	PageNumber int32  `json:"pageNumber"`
	PageSize   int32  `json:"pageSize"`
	OrderBy    string `json:"orderBy"`
	Order      string `json:"order"`
}

type ListChangeResult struct {
	TotalCount int32                 `json:"totalCount"`
	ChangeLogs []ServiceChangeRecord `json:"changeLogs"`
	PageNumber int32                 `json:"pageNumber"`
	PageSize   int32                 `json:"pageSize"`
	OrderBy    string                `json:"orderBy"`
	Order      string                `json:"order"`
	BaseRes
}

type ChangeDetailArgs struct {
	ChangeId string `json:"changeId"`
}

type ChangeDetailResult struct {
	ChangeId    string `json:"changeId"`
	Prev        string `json:"prev"`
	ChangeType  int32  `json:"changeType"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	CreatedAt   uint32 `json:"createdAt"`
	BaseRes
}

type DeleteServiceArgs struct {
	ServiceId string `json:"serviceId"`
}

type DeleteServiceResult struct {
	BaseRes
}

type ListPodArgs struct {
	ServiceId string `json:"serviceId"`
}

type ListPodResult struct {
	Pods []InsInfo `json:"pods"`
	BaseRes
}

type BlockPodArgs struct {
	ServiceId  string `json:"serviceId"`
	InstanceId string `json:"instanceId"`
	Block      bool   `json:"block"`
}

type BlockPodResult struct {
	BaseRes
}

type DeletePodArgs struct {
	ServiceId  string `json:"serviceId"`
	InstanceId string `json:"instanceId"`
}

type DeletePodResult struct {
	BaseRes
}

type ListPodGroupsArgs struct {
	ServiceId string `json:"serviceId"`
}

type ListPodGroupsResult struct {
	ResourcePoolId   string           `json:"resourcePoolId"`
	ResourcePoolName string           `json:"resourcePoolName"`
	QueueName        string           `json:"queueName"`
	ServicePodGroups []InsGroupStatus `json:"servicePodGroups"`
	BaseRes
}

type InsGroupStatus struct {
	Id            string    `json:"id"`
	Masters       []InsInfo `json:"masters"`
	Workers       []InsInfo `json:"workers"`
	Status        string    `json:"status"`
	Blocked       bool      `json:"blocked"`
	CreatedAt     uint32    `json:"createdAt"`
	AvailablePods uint32    `json:"availablePods"`
	TotalPods     uint32    `json:"totalPods"`
}

type InsInfo struct {
	InstanceId string          `json:"instanceId"`
	Containers []ContainerInfo `json:"containers"`
	Status     InsStatus       `json:"status"`
	Hazard     bool            `json:"hazard,omitempty"`
}

type ContainerInfo struct {
	ContainerId string          `json:"containerId"`
	Container   ContainerConf   `json:"container"`
	Status      ContainerStatus `json:"status"`
}

type ContainerStatus struct {
	ContainerStatus string `json:"containerStatus"`
	CreatedAt       uint32 `json:"createdAt"`
	Reason          string `json:"reason"`
}

type InsStatus struct {
	Blocked             bool   `json:"blocked"`
	Status              string `json:"status"`
	CreatedAt           uint32 `json:"createdAt"`
	AvailableContainers int32  `json:"availableContainers"`
	TotalContainers     int32  `json:"totalContainers"`
	PodIP               string `json:"podIP"`
	NodeIP              string `json:"nodeIP"`
	Reason              string `json:"reason"` // pod 级别的异常信息
}

type ChangeDetailData struct {
	ChangeId   string `json:"changeId"`
	Prev       string `json:"prev"`
	ChangeType int32  `json:"changeType"`
	ShortDesc  string `json:"shortDesc"`
	Creator    string `json:"creator"`
	CreatedAt  uint32 `json:"createdAt"`
}

type ServiceChangeRecord struct {
	ChangeId    string `json:"changeId"`
	Prev        string `json:"prev"`
	ChangeType  int32  `json:"changeType"`
	Description string `json:"description"`
	Creator     string `json:"creator"`
	CreatedAt   uint32 `json:"createdAt"`
}

type ServiceStatus struct {
	AccessIPs   AccessIPConf     `json:"accessIPs"`
	AccessPorts []AccessPortConf `json:"accessPorts"`
	BlbShortId  string           `json:"blbShortId"`
	BriefStat   ServiceBriefStat `json:"briefStat"`
}

type AccessIPConf struct {
	Internal string `json:"internal"`
	External string `json:"external"`
}

type AccessPortConf struct {
	Name          string `json:"name"`
	ContainerPort int32  `json:"containerPort"`
	ServicePort   int32  `json:"servicePort"`
}

type ServiceBriefStat struct {
	Status       int32  `json:"status"`
	AvailableIns int32  `json:"availableIns"`
	TotalIns     int32  `json:"totalIns"`
	Hazard       bool   `json:"hazard,omitempty"`
	Reason       string `json:"reason"`
}

type ServiceBriefInfo struct {
	Id               string       `json:"id"`
	Name             string       `json:"name"`
	ResourcePoolId   string       `json:"resourcePoolId"`
	ResourcePoolName string       `json:"resourcePoolName"`
	QueueName        string       `json:"queueName"`
	Region           string       `json:"region"`
	PublicAccess     bool         `json:"publicAccess"`
	NetType          string       `json:"networkType"`
	Creator          string       `json:"creator"`
	CreatedAt        uint32       `json:"createdAt"`
	UpdatedAt        uint32       `json:"updatedAt"`
	ResourcePoolType string       `json:"resourcePoolType"`
	Hpa              HPAConf      `json:"hpa,omitempty"`
	ResourceSpec     ResourceSpec `json:"resourceSpec"`
	WorkloadType     string       `json:"workloadType"`
}
type ResourceSpec struct {
	Cpus             int32  `json:"cpus"`
	Memory           int32  `json:"memory"`
	AcceleratorCount int32  `json:"acceleratorCount"`
	AcceleratorType  string `json:"acceleratorType"`
}

type ContainerConf struct {
	Name             string            `json:"name"`
	Cpus             int32             `json:"cpus"`
	Memory           int32             `json:"memory"`
	AcceleratorCount int32             `json:"acceleratorCount"`
	Command          []string          `json:"command"`
	RunArgs          []string          `json:"runArgs"`
	Ports            []PortConf        `json:"ports"`
	Envs             map[string]string `json:"envs"`
	Image            ImageConf         `json:"image"`
	VolumeMounts     []VolumnMountConf `json:"volumeMounts"`
	StartupsProbe    *ProbeConf        `json:"startupsProbe"`
	ReadinessProbe   *ProbeConf        `json:"readinessProbe"`
	LivenessProbe    *ProbeConf        `json:"livenessProbe"`

	PostStart       *LifecycleHandlerConf `json:"postStart,omitempty"`
	PreStop         *LifecycleHandlerConf `json:"preStop,omitempty"`
	IsInitContainer bool                  `json:"isInitContainer,omitempty"`
}

type LifecycleHandlerConf struct {
	Exec      *ExecAction      `protobuf:"bytes,1,opt,name=exec,proto3,oneof" json:"exec"`
	HttpGet   *HTTPGetAction   `protobuf:"bytes,2,opt,name=httpGet,proto3,oneof" json:"httpGet"`
	TcpSocket *TCPSocketAction `protobuf:"bytes,3,opt,name=tcpSocket,proto3,oneof" json:"tcpSocket"`
	SleepSec  int32            `protobuf:"varint,4,opt,name=sleepSec,proto3,oneof" json:"sleepSec"`
}

type PortConf struct {
	Protocol string `json:"protocol,omitempty"`
	Name     string `json:"name"`
	Port     int32  `json:"port"`
	Prefix   string `json:"prefix,omitempty"`
	Rewrite  string `json:"rewrite,omitempty"`
}

type ImageConf struct {
	ImageType int32  `json:"imageType"`
	ImageUrl  string `json:"imageUrl"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type VolumnMountConf struct {
	VolumnName string `json:"volumnName"`
	MountPath  string `json:"mountPath"`
	ReadOnly   bool   `json:"readOnly"`
}

type ProbeConf struct {
	InitialDelaySeconds int32            `json:"initialDelaySeconds"`
	TimeoutSeconds      int32            `json:"timeoutSeconds"`
	PeriodSeconds       int32            `json:"periodSeconds"`
	SuccessThreshold    int32            `json:"successThreshold"`
	FailureThreshold    int32            `json:"failureThreshold"`
	Handler             ProbeHandlerConf `json:"handler"`
}

type ProbeHandlerConf struct {
	Exec            *ExecAction      `json:"exec"`
	HttpGet         *HTTPGetAction   `json:"httpGet"`
	TcpSocketAction *TCPSocketAction `json:"tcpSocketAction"`
}

type ExecAction struct {
	Command []string `json:"command"`
}

type HTTPGetAction struct {
	Path string `json:"path"`
	Port int32  `json:"port"`
}

type TCPSocketAction struct {
	Port int32 `json:"port"`
}

type ResourcePoolConf struct {
	ResourcePoolId   string `json:"resourcePoolId"`
	ResourcePoolName string `json:"resourcePoolName"`
	QueueName        string `json:"queueName"`
	ResourcePoolType string `json:"resourcePoolType"`
}

type StorageConf struct {
	ShmSize int32        `json:"shmSize"`
	Volumns []VolumnConf `json:"volumns"`
}

type VolumnConf struct {
	VolumeType string          `json:"volumeType"`
	VolumnName string          `json:"volumnName"`
	Dynamic    bool            `json:"dynamic,omitempty"` // 动态挂载类型，动态挂载类型的卷，不创建 pv，而是通过 storageClass 动态创建 pv
	Pfs        *PFSConfig      `json:"pfs"`
	Hostpath   *HostPathConfig `json:"hostpath"`
	Cfs        *CFSConfig      `json:"cfs,omitempty"`
	Dataset    *DataSetConfig  `json:"dataset,omitempty"`
	Cds        *CDSConfig      `json:"cds,omitempty"`
	Bos        *BOSConfig      `json:"bos,omitempty"`
}

type BOSConfig struct {
	Secret     SecretRef `json:"secret,omitempty"`
	SourcePath string    `json:"sourcePath,omitempty"`
}

type SecretRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type CDSConfig struct {
	InstanceId  string `json:"instanceId"`
	StorageType string `json:"storageType"`
	Capacity    int32  `json:"capacity"` // 动态挂载时必填
	SourcePath  string `json:"sourcePath"`
}

type CFSConfig struct {
	InstanceId string `json:"instanceId,omitempty"`
	SourcePath string `json:"sourcePath,omitempty"`
	MountPoint string `json:"mountPoint,omitempty"`
}

type DataSetConfig struct {
	DatasetId   string `json:"datasetId,omitempty"`
	VersionId   string `json:"versionId,omitempty"`
	Source      string `json:"source,omitempty"` // 固定私有类型
	StorageType string `json:"storageType,omitempty"`

	// DefaultMountPath string     `json:"DefaultMountPath"`
	Pfs *PFSConfig `json:"pfs,omitempty"`
}

type PFSConfig struct {
	InstanceId    string   `json:"instanceId"`
	InstanceType  string   `json:"instanceType"`
	HostMountPath string   `json:"hostMountPath"`
	MountTargetId []string `json:"mountTargetId"`
	ClusterIP     string   `json:"clusterIP"`
	ClientID      string   `json:"clientID"`
	ClusterPort   string   `json:"clusterPort"`
	SourcePath    string   `json:"sourcePath"`
}

type HostPathConfig struct {
	SourcePath string `json:"sourcePath"`
}

type AccessConf struct {
	PublicAccess       bool              `json:"publicAccess"`
	Eip                string            `json:"eip"`
	NetworkType        string            `json:"networkType"`
	AiGateway          AiGatewayConf     `json:"aiGateway"`
	ServiceLabels      map[string]string `json:"serviceLabels"`
	ServiceAnnotations map[string]string `json:"serviceAnnotations"`
}

type AiGatewayConf struct {
	EnableAuth bool `json:"enableAuth"`
}

type LogConf struct {
	Persistent bool `json:"persistent"`
}

type DeployConf struct {
	CanaryStrategy *CanaryStrategyConf `json:"canaryStrategy"`
	Schedule       ScheduleConf        `json:"schedule"`
}

type ScheduleConf struct {
	Priority string `json:"priority"`
}

type CanaryStrategyConf struct {
	MaxSurge       int32 `json:"maxSurge"`
	MaxUnavailable int32 `json:"maxUnavailable"`
}

type Misc struct {
	PodLabels      map[string]string `json:"podLabels"`
	PodAnnotations map[string]string `json:"podAnnotations,omitempty"`
	GracePeriodSec int32             `json:"gracePeriodSec"`
	FedPodsPerIns  int32             `json:"fedPodsPerIns"`
	EnableRDMA     bool              `json:"enableRDMA"`
}

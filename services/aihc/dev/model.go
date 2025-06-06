package dev

type CreateDevInstanceResult struct {
	DevInstanceId string `json:"devInstanceId"`
	RequestId     string `json:"requestId"`
}

type ListDevInstanceArgs struct {
	QueryKey       string `json:"queryKey"`
	QueryVal       string `json:"queryVal"`
	ResourcePoolId string `json:"resourcePoolID"`
	QueueName      string `json:"queueName"`
	OnlyMyDevs     bool   `json:"onlyMyDevs"`
	PageNumber     int    `json:"pageNumber"`
	PageSize       int    `json:"pageSize"`
}

type ListDevInstanceResult struct {
	DevInstances []*DevInstanceBrief `json:"devInstances"`
	TotalCount   int32               `json:"totalCount"`
	RequestId    string              `json:"requestId"`
}

type QueryDevInstanceDetailArgs struct {
	DevInstanceId string `json:"devInstanceId"`
}

type QueryDevInstanceDetailResult struct {
	DevInstance DevInstanceDetail `json:"devInstance"`
	RequestId   string            `json:"requestId"`
}

type StartDevInstanceArgs struct {
	DevInstanceId string `json:"devInstanceId"`
}

type StartDevInstanceResult struct {
	DevInstanceId string `json:"devInstanceId"`
	RequestId     string `json:"requestId"`
}

type StopDevInstanceArgs struct {
	DevInstanceId string `json:"devInstanceId"`
}

type StopDevInstanceResult struct {
	DevInstanceId string `json:"devInstanceId"`
	RequestId     string `json:"requestId"`
}

type DeleteDevInstanceArgs struct {
	DevInstanceId string `json:"devInstanceId"`
}

type DeleteDevInstanceResult struct {
	DevInstanceId string `json:"devInstanceId"`
	RequestId     string `json:"requestId"`
}

type TimedStopDevInstanceResult struct {
	DevInstanceId string `json:"devInstanceId"`
	RequestId     string `json:"requestId"`
}

type ListDevInstanceEventArgs struct {
	DevInstanceId string `json:"devInstanceId"`
	PageNumber    int    `json:"pageNumber"`
	PageSize      int    `json:"pageSize"`
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	EventType     string `json:"eventType"`
	Message       string `json:"message"`
}

type ListDevInstanceEventResult struct {
	Events     []*Event `json:"events"`
	TotalCount int32    `json:"totalCount"`
	RequestId  string   `json:"requestId"`
}

type CreateDevInstanceImagePackJobResult struct {
	ImagePackJobId string `json:"imagePackJobId"`
	DevInstanceId  string `json:"devInstanceId"`
	RequestId      string `json:"requestId"`
}

type DevInstanceImagePackJobDetailArgs struct {
	ImagePackJobId string `json:"imagePackJobId"`
	DevInstanceId  string `json:"devInstanceId"`
}

type DevInstanceImagePackJobDetailResult struct {
	ImagePackJobDetail ImagePackJobDetail `json:"devInstanceImagePackJob"`

	RequestId string `json:"requestId"`
}

type ImagePackJobDetail struct {
	Id            string `json:"id"`
	DevInstanceId string `json:"devInstanceId"`
	ImageName     string `json:"imageName"`
	ImageTag      string `json:"imageTag"`
	Registry      string `json:"registry"`
	Namespace     string `json:"namespace"`
	Status        int32  `json:"status"`
	CreatedAt     int32  `json:"createdAt"`
	UpdatedAt     int32  `json:"updatedAt"`
}

type CreateDevInstanceArgs struct {
	Name         string           `json:"name" validate:"required"`
	Id           string           `json:"id"`
	Conf         *DevInstanceConf `json:"conf" validate:"required"`
	VisibleScope *VisibleScope    `json:"visibleScope"`
	Notify       *Notify          `json:"notify"`
	IsPublicMgmt bool             `json:"isPublicMgmt"`
	Creator      string           `json:"creator" validate:"required"`
	CreatorID    string           `json:"creatorId" validate:"required"`
}

type DevInstanceConf struct {
	ResourcePool *ResourcePool `json:"resourcePool" validate:"required"`
	Resources    *Resources    `json:"resources" validate:"required"`
	Image        *Image        `json:"image" validate:"required"`
	Access       *Access       `json:"access"`
	ScheduleConf *ScheduleConf `json:"scheduleConf"`
	VolumnConfs  []*VolumnConf `json:"volumnConfs" validate:"required"`
	UpdatedAt    int32         `json:"updatedAt"`
}

type VolumnConf struct {
	VolumnType string   `json:"volumnType" validate:"required"`
	PFS        *PFS     `json:"pfs"`
	Dataset    *Dataset `json:"dataset"`
	CFS        *CFS     `json:"cfs"`
	CDS        *CDS     `json:"cds"`
	BOS        *BOS     `json:"bos"`
	MountPath  string   `json:"mountPath" validate:"required"`
	ReadOnly   bool     `json:"readOnly"`
}

type PFS struct {
	InstanceID string `json:"instanceId" validate:"required"`
	SourcePath string `json:"sourcePath" validate:"required"`
}

type DatasetPFS struct {
	ClientID      string   `json:"clientID"`
	InstanceType  string   `json:"instanceType"`
	Region        string   `json:"region"`
	InstanceID    string   `json:"instanceId"`
	ClusterIP     string   `json:"clusterIP"`
	MountTargetID []string `json:"mountTargetId"`
	HostMountPath string   `json:"hostMountPath"`
	SrcPath       string   `json:"srcPath"`
}

type Dataset struct {
	DatasetID        string      `json:"datasetId"`
	VersionID        string      `json:"versionId"`
	StorageType      string      `json:"storageType"`
	PFS              *DatasetPFS `json:"pfs"`
	DefaultMountPath string      `json:"defaultMountPath"`
}

type CFS struct {
	InstanceID string `json:"instanceId" validate:"required"`
	SourcePath string `json:"sourcePath" validate:"required"`
	MountPoint string `json:"mountPoint" validate:"required"`
}

type CDS struct {
	Capacity int32 `json:"capacity" validate:"required"`
}

type BOS struct {
	SourcePath string `json:"sourcePath" validate:"required"`
}

type VisibleScope struct {
	Type int32 `json:"type"`
}

type Notify struct {
	NotifyRuleID string `json:"notifyRuleId" validate:"required"`
	IsOpen       bool   `json:"isOpen"`
}

type ResourcePool struct {
	ResourcePoolType string `json:"resourcePoolType"`
	ResourcePoolID   string `json:"resourcePoolId"`
	ResourcePoolName string `json:"resourcePoolName"`
	QueueName        string `json:"queueName" validate:"required"`
}

type Resources struct {
	AcceleratorType  string `json:"acceleratorType"`
	AcceleratorCount int32  `json:"acceleratorCount"`
	CPUs             int32  `json:"cpus" validate:"required"`
	Memory           int32  `json:"memory" validate:"required"`
	ShmSize          int32  `json:"shmSize"`
	// EnableVGPU       bool   `json:"enableVGPU"`
}

type Image struct {
	ImageType int32  `json:"imageType"`
	ImageURL  string `json:"imageUrl" validate:"required"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type Access struct {
	BLBID     string     `json:"blbId" validate:"required"`
	SSHEnable bool       `json:"sshEnable"`
	SSHRSAKey string     `json:"sshRSAPubKey"`
	PortInfo  []PortInfo `json:"portInfo" validate:"required"`
}

type PortInfo struct {
	Name       string `json:"name" validate:"required"`
	Port       int32  `json:"port" validate:"required"`
	AccessPort int32  `json:"accessPort" validate:"required"`
}

type ScheduleConf struct {
	CPUNodeAffinity bool   `json:"cpuNodeAffinity"`
	Priority        string `json:"priority"`
}

type TimedStopDevInstanceArgs struct {
	DevInstanceId string `json:"devInstanceId" validate:"required"`
	DelaySec      int32  `json:"delaySec"`
	Enable        bool   `json:"enable"`
}

type CreateDevInstanceImagePackJobArgs struct {
	ImageName     string `json:"imageName" validate:"required"`
	ImageTag      string `json:"imageTag" validate:"required"`
	Registry      string `json:"registry" validate:"required"`
	Namespace     string `json:"namespace" validate:"required"`
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required"`
	DevInstanceID string `json:"devInstanceId" validate:"required"`
}

type Event struct {
	Reason         string `json:"reason"`         // 原因
	Message        string `json:"message"`        // 详细信息
	FirstTimestamp int32  `json:"firstTimestamp"` // 首次出现时间(建议使用time.Time类型)
	LastTimestamp  int32  `json:"lastTimestamp"`  // 最后出现时间(建议使用time.Time类型)
	Count          int32  `json:"count"`          // 出现次数
	Type           string `json:"type"`           // 事件类型
}

type DevInstanceDetail struct {
	Name string `json:"name"`
	ID   string `json:"id"`

	Conf                 *DevInstanceConf            `json:"conf"`
	ServiceInstanceInfo  *ServiceInstanceInfo        `json:"serviceInstanceInfo"`
	LoginInfo            *LoginInfo                  `json:"loginInfo"`
	Notify               *Notify                     `json:"notify"`
	VisibleScope         *VisibleScope               `json:"visibleScope"`
	TimedStopDevInstance *TimedStopDevInstanceDetail `json:"timedStopDevInstance"`
	IsPublicMgmt         bool                        `json:"isPublicMgmt"`

	AccountId    string `json:"accountId"`
	Creator      string `json:"creator"`
	CreatorID    string `json:"creatorId"`
	Region       string `json:"region"`
	Version      string `json:"version"`
	Status       int32  `json:"status"`
	StatusReason string `json:"statusReason"`
	CreatedAt    int32  `json:"createdAt"`
	UpdatedAt    int32  `json:"updatedAt"`
}

type ServiceInstanceInfo struct {
	NodeIP            string `json:"nodeIP"`
	PodName           string `json:"podName"`
	ServiceInstanceId string `json:"serviceInstanceId"`
	InternalIP        string `json:"internalIP"`
	PublicIP          string `json:"publicIP"`
}

type LoginInfo struct {
	Jupyter Url `json:"jupyter"`
	Vscode  Url `json:"vscode"`
}

type Url struct {
	URL string `json:"url"`
}

type TimedStopDevInstanceDetail struct {
	StartTime int32 `json:"startTime"`
	DelaySec  int32 `json:"delaySec"`
}

type DevInstanceBrief struct {
	Name             string     `json:"name"`
	ID               string     `json:"id"`
	ResourcePoolId   string     `json:"resourcePoolId"`
	ResourcePoolName string     `json:"resourcePoolName"`
	QueueName        string     `json:"queueName"`
	ImageURL         string     `json:"imageUrl"`
	Resources        *Resources `json:"resources"`
	AccountId        string     `json:"accountId"`
	Creator          string     `json:"creator"`
	CreatorID        string     `json:"creatorId"`
	Region           string     `json:"region"`
	Version          string     `json:"version"`
	Status           int32      `json:"status"`
	StatusReason     string     `json:"statusReason"`
	CreatedAt        int32      `json:"createdAt"`
	UpdatedAt        int32      `json:"updatedAt"`
}

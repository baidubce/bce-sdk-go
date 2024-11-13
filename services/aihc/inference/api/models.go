/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// models.go - definitions of the request arguments and results data structure model

package api

type BaseRes struct {
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestId string `json:"requestId,omitempty"`
}

type CreateAppArgs struct {
	AppName    string           `json:"appName,omitempty"`
	ChipType   string           `json:"chipType,omitempty"`
	InsCount   int32            `json:"insCount,omitempty"`
	ResPool    *ResPoolConf     `json:"resPool,omitempty"`
	Storage    *StorageConf     `json:"storage,omitempty"`
	Containers []*ContainerConf `json:"containers,omitempty"`
	Access     *AccessConf      `json:"access,omitempty"`
	Log        *LogConf         `json:"log,omitempty"`
	Deploy     *DeployConf      `json:"deploy,omitempty"`
	Misc       *Misc            `json:"misc,omitempty"`
}

type CreateAppResult struct {
	Data CreateAppData `json:"data,omitempty"`
	BaseRes
}

type CreateAppData struct {
	AppId string `json:"appId,omitempty"`
}

type ListAppArgs struct {
	PageNo   int32  `json:"pageNo,omitempty"`
	PageSize int32  `json:"pageSize,omitempty"`
	OrderBy  string `json:"orderBy,omitempty"`
	Order    string `json:"order,omitempty"`
	Keyword  string `json:"keyword,omitempty"`
}

type ListAppResult struct {
	Data ListAppData `json:"data,omitempty"`
	BaseRes
}

type ListAppData struct {
	AppList []*AppBriefInfo `json:"list,omitempty"`
	Count   int32           `json:"count,omitempty"`
}

type ListAppStatsArgs struct {
	AppIds []string `json:"appIds,omitempty"`
}

type ListAppStatsResult struct {
	Data ListAppStatsData `json:"data,omitempty"`
	BaseRes
}

type ListAppStatsData struct {
	Apps map[string]*AppBriefStat `json:"apps,omitempty"`
}

type AppDetailsArgs struct {
	AppId string `json:"appId,omitempty"`
}

type AppDetailsResult struct {
	Data AppDetailsData `json:"data,omitempty"`
	BaseRes
}

type AppDetailsData struct {
	App     *CreateAppArgs `json:"app,omitempty"`
	Status  *AppStatus     `json:"status,omitempty"`
	Creator string         `json:"creator,omitempty"`
	Ctime   uint32         `json:"ctime,omitempty"`
	Mtime   uint32         `json:"mtime,omitempty"`
}

type UpdateAppArgs struct {
	AppId     string         `json:"appId,omitempty"`
	AppConf   *CreateAppArgs `json:"appConf,omitempty"`
	ShortDesc string         `json:"shortDesc,omitempty"`
}

type UpdateAppResult struct {
	Data UpdateAppData `json:"data,omitempty"`
	BaseRes
}

type UpdateAppData struct {
	AppId string `json:"appId,omitempty"`
}

type ScaleAppArgs struct {
	AppId    string `json:"appId,omitempty"`
	InsCount int32  `json:"insCount,omitempty"`
}

type ScaleAppResult struct {
	BaseRes
}

type PubAccessArgs struct {
	AppId        string `json:"appId,omitempty"`
	PublicAccess bool   `json:"publicAccess,omitempty"`
	Eip          string `json:"eip,omitempty"`
}

type PubAccessResult struct {
	BaseRes
}

type ListChangeArgs struct {
	AppId      string `json:"appId,omitempty"`
	ChangeType int32  `json:"changeType,omitempty"`
	PageNo     int32  `json:"pageNo,omitempty"`
	PageSize   int32  `json:"pageSize,omitempty"`
	OrderBy    string `json:"orderBy,omitempty"`
	Order      string `json:"order,omitempty"`
}

type ListChangeResult struct {
	Data ListChangeData `json:"data,omitempty"`
	BaseRes
}

type ListChangeData struct {
	Count int32              `json:"count,omitempty"`
	List  []*AppChangeRecord `json:"list,omitempty"`
}

type ChangeDetailArgs struct {
	ChangeId string `json:"changeId,omitempty"`
}

type ChangeDetailResult struct {
	Data ChangeDetailData `json:"data,omitempty"`
	BaseRes
}

type DeleteAppArgs struct {
	AppId string `json:"appId,omitempty"`
}

type DeleteAppResult struct {
	BaseRes
}

type ListPodArgs struct {
	AppId string `json:"appId,omitempty"`
}

type ListPodResult struct {
	Data ListPodData `json:"data,omitempty"`
	BaseRes
}

type ListPodData struct {
	List []*InsInfo `json:"list,omitempty"`
}

type BlockPodArgs struct {
	AppId string `json:"appId,omitempty"`
	InsID string `json:"insID,omitempty"`
	Block bool   `json:"block,omitempty"`
}

type BlockPodResult struct {
	BaseRes
}

type DeletePodArgs struct {
	AppId string `json:"appId,omitempty"`
	InsID string `json:"insID,omitempty"`
}

type DeletePodResult struct {
	BaseRes
}

type ListBriefResPoolArgs struct {
	PageNo   int32 `json:"pageNo,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
}

type ListBriefResPoolResult struct {
	Data ListBriefResPoolData `json:"data,omitempty"`
	BaseRes
}

type ResPoolDetailArgs struct {
	ResPoolId string `json:"resPoolId,omitempty"`
}

type ResPoolDetailResult struct {
	Data ResPoolDetailData `json:"data,omitempty"`
	BaseRes
}

type ResPoolDetailData struct {
	ResPool  *ResPoolInfo    `json:"resPool,omitempty"`
	ResQueue []*ResQueueInfo `json:"resQueue,omitempty"`
}

type ResPoolInfo struct {
	Meta   *ResPoolMeta   `json:"meta,omitempty"`
	Spec   *ResPoolSpec   `json:"spec,omitempty"`
	Status *ResPoolStatus `json:"status,omitempty"`
}

type ResQueueInfo struct {
	Name            string            `json:"name,omitempty"`
	NameSpace       string            `json:"nameSpace,omitempty"`
	ParentQueue     string            `json:"parentQueue,omitempty"`
	QueueType       string            `json:"queueType,omitempty"`
	State           string            `json:"state,omitempty"`
	Reclaimable     bool              `json:"reclaimable,omitempty"`
	Preemptable     bool              `json:"preemptable,omitempty"`
	Capability      map[string]string `json:"capability,omitempty"`
	Allocated       map[string]string `json:"allocated,omitempty"`
	DisableOversell bool              `json:"disableOversell,omitempty"`
	CreatedTime     string            `json:"createdTime,omitempty"`
}

type ResPoolMeta struct {
	ResPoolId   string `json:"resPoolId,omitempty"`
	ResPoolName string `json:"resPoolName,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type ResPoolSpec struct {
	K8SVersion         string   `json:"k8sVersion,omitempty"`
	AssociatedPfsID    string   `json:"associatedPfsID,omitempty"`
	Region             string   `json:"region,omitempty"`
	AssociatedCpromIDs []string `json:"associatedCpromIDs,omitempty"`
	CreatedBy          string   `json:"createdBy,omitempty"`
	Description        string   `json:"description,omitempty"`
	ForbidDelete       bool     `json:"forbidDelete,omitempty"`
}

type ResPoolStatus struct {
	GPUCount  *ResPoolCountInfo `json:"GPUCount,omitempty"`
	NodeCount *ResPoolCountInfo `json:"NodeCount,omitempty"`
	Phase     string            `json:"phase,omitempty"`
}

type ResPoolCountInfo struct {
	Total int32 `json:"total,omitempty"`
	Used  int32 `json:"used,omitempty"`
}

type ListBriefResPoolData struct {
	ResPools []*ResPoolBriefInfo `json:"resPools,omitempty"`
}

type ResPoolBriefInfo struct {
	ResPoolId       string `json:"resPoolId,omitempty"`
	ResPoolName     string `json:"resPoolName,omitempty"`
	AssociatedPfsID string `json:"associatedPfsID,omitempty"`
	ClusterType     string `json:"clusterType,omitempty"`
	Description     string `json:"description,omitempty"`
	ForbidDelete    bool   `json:"forbidDelete,omitempty"`
	K8SVersion      string `json:"k8sVersion,omitempty"`
	CreatedAt       string `json:"createdAt,omitempty"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
	Phase           string `json:"phase,omitempty"`
}

type InsInfo struct {
	InsID      string           `json:"insID,omitempty"`
	Containers []*ContainerInfo `json:"containers,omitempty"`
	Status     *InsStauts       `json:"status,omitempty"`
}

type ContainerInfo struct {
	ContainerId string           `json:"containerId,omitempty"`
	Container   *ContainerConf   `json:"container,omitempty"`
	Status      *ContainerStatus `json:"status,omitempty"`
}

type ContainerStatus struct {
	ContainerStatus string `json:"containerStatus,omitempty"`
	Ctime           uint32 `json:"ctime,omitempty"`
	Reason          string `json:"reason,omitempty"`
}

type InsStauts struct {
	Blocked             bool   `json:"blocked,omitempty"`
	InsStatus           string `json:"insStatus,omitempty"`
	Ctime               uint32 `json:"ctime,omitempty"`
	AvailableContainers int32  `json:"availableContainers,omitempty"`
	TotalContainers     int32  `json:"totalContainers,omitempty"`
	PodIP               string `json:"podIP,omitempty"`
	NodeIP              string `json:"nodeIP,omitempty"`
}

type ChangeDetailData struct {
	ChangeId   string `json:"changeId,omitempty"`
	Prev       string `json:"prev,omitempty"`
	ChangeType int32  `json:"changeType,omitempty"`
	ShortDesc  string `json:"shortDesc,omitempty"`
	Creator    string `json:"creator,omitempty"`
	Ctime      uint32 `json:"ctime,omitempty"`
}

type AppChangeRecord struct {
	ChangeId   string `json:"changeId,omitempty"`
	Prev       string `json:"prev,omitempty"`
	ChangeType int32  `json:"changeType,omitempty"`
	ShortDesc  string `json:"shortDesc,omitempty"`
	Creator    string `json:"creator,omitempty"`
	Ctime      uint32 `json:"ctime,omitempty"`
}

type AppStatus struct {
	AccessIPs   *AccessIPConf     `json:"accessIPs,omitempty"`
	AccessPorts []*AccessPortConf `json:"accessPorts,omitempty"`
	BlbShortId  string            `json:"blbShortId,omitempty"`
}

type AccessIPConf struct {
	Internal string `json:"internal,omitempty"`
	External string `json:"external,omitempty"`
}

type AccessPortConf struct {
	Name          string `json:"name,omitempty"`
	ContainerPort int32  `json:"containerPort,omitempty"`
	ServicePort   int32  `json:"servicePort,omitempty"`
}

type AppBriefStat struct {
	Status       int32 `json:"status,omitempty"`
	AvailableIns int32 `json:"availableIns,omitempty"`
	TotalIns     int32 `json:"totalIns,omitempty"`
}

type AppBriefInfo struct {
	AppId        string `json:"appId,omitempty"`
	AppName      string `json:"appName,omitempty"`
	ResPoolId    string `json:"resPoolId,omitempty"`
	ResPoolName  string `json:"resPoolName,omitempty"`
	ResQueue     string `json:"resQueue,omitempty"`
	Region       string `json:"region,omitempty"`
	PublicAccess bool   `json:"publicAccess,omitempty"`
	Creator      string `json:"creator,omitempty"`
	Ctime        uint32 `json:"ctime,omitempty"`
	Mtime        uint32 `json:"mtime,omitempty"`
}

type ContainerConf struct {
	Name           string             `json:"name,omitempty"`
	Cpus           int32              `json:"cpus,omitempty"`
	Memory         int32              `json:"memory,omitempty"`
	Cards          int32              `json:"cards,omitempty"`
	RunCmd         []string           `json:"runCmd,omitempty"`
	RunArgs        []string           `json:"runArgs,omitempty"`
	Ports          []*PortConf        `json:"ports,omitempty"`
	Env            map[string]string  `json:"env,omitempty"`
	Image          *ImageConf         `json:"image,omitempty"`
	VolumeMounts   []*VolumnMountConf `json:"volumeMounts,omitempty"`
	StartupsProbe  *ProbeConf         `json:"startupsProbe,omitempty"`
	ReadinessProbe *ProbeConf         `json:"readinessProbe,omitempty"`
	LivenessProbe  *ProbeConf         `json:"livenessProbe,omitempty"`
}

type PortConf struct {
	Port int32 `json:"port,omitempty"`
}

type ImageConf struct {
	ImageType     int32  `json:"imageType,omitempty"`
	ImageAddr     string `json:"imageAddr,omitempty"`
	ImagePullUser string `json:"imagePullUser,omitempty"`
	ImagePullPass string `json:"imagePullPass,omitempty"`
}

type VolumnMountConf struct {
	VolumnName string `json:"volumnName,omitempty"`
	DstPath    string `json:"dstPath,omitempty"`
	ReadOnly   bool   `json:"readOnly,omitempty"`
}

type ProbeConf struct {
	InitialDelaySeconds int32             `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      int32             `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       int32             `json:"periodSeconds,omitempty"`
	SuccessThreshold    int32             `json:"successThreshold,omitempty"`
	FailureThreshold    int32             `json:"failureThreshold,omitempty"`
	Handler             *ProbeHandlerConf `json:"handler,omitempty"`
}

type ProbeHandlerConf struct {
	Exec            *ExecAction      `json:"exec,omitempty"`
	HttpGet         *HTTPGetAction   `json:"httpGet,omitempty"`
	TcpSocketAction *TCPSocketAction `json:"tcpSocketAction,omitempty"`
}

type ExecAction struct {
	Command []string `json:"command,omitempty"`
}

type HTTPGetAction struct {
	Path string `json:"path,omitempty"`
	Port int32  `json:"port,omitempty"`
}

type TCPSocketAction struct {
	Port int32 `json:"port,omitempty"`
}

type ResPoolConf struct {
	ResPoolId string `json:"resPoolId,omitempty"`
	Queue     string `json:"queue,omitempty"`
}

type StorageConf struct {
	ShmSize int32         `json:"shmSize,omitempty"`
	Volumns []*VolumnConf `json:"volumns,omitempty"`
}

type VolumnConf struct {
	VolumeType string          `json:"volumeType,omitempty"`
	VolumnName string          `json:"volumnName,omitempty"`
	Pfs        *PFSConfig      `json:"pfs,omitempty"`
	Hostpath   *HostPathConfig `json:"hostpath,omitempty"`
}

type PFSConfig struct {
	SrcPath string `json:"srcPath,omitempty"`
}

type HostPathConfig struct {
	SrcPath string `json:"srcPath,omitempty"`
}

type AccessConf struct {
	PublicAccess bool   `json:"publicAccess,omitempty"`
	Eip          string `json:"eip,omitempty"`
}

type LogConf struct {
	Persistent bool `json:"persistent,omitempty"`
}

type DeployConf struct {
	CanaryStrategy *CanaryStrategyConf `protobuf:"bytes,1,opt,name=canaryStrategy,proto3" json:"canaryStrategy,omitempty"`
}

type CanaryStrategyConf struct {
	MaxSurge       int32 `json:"maxSurge,omitempty"`
	MaxUnavailable int32 `json:"maxUnavailable,omitempty"`
}

type Misc struct {
	PodLabels      map[string]string `json:"podLabels,omitempty"`
	PodAnnotations map[string]string `json:"podAnnotations,omitempty"`
}

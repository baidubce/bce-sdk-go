/*
 * Copyright 2021 Baidu, Inc.
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

// model.go - definitions of the request arguments and results data structure model

package api

type V1VolumeMount struct {
	Name             string `json:"name,omitempty"`
	MountPath        string `json:"mountPath,omitempty"`
	ReadOnly         bool   `json:"readOnly,omitempty"`
	MountPropagation string `json:"mountPropagation,omitempty"`
	SubPath          string `json:"subPath,omitempty"`
}

type V1ContainerPort struct {
	Protocol      string `json:"protocol,omitempty"`
	ContainerPort int    `json:"containerPort,omitempty"`
	HostIP        string `json:"hostIP,omitempty"`
	HostPort      int    `json:"hostPort,omitempty"`
	Name          string `json:"name,omitempty"`
}

type V1ConfigMapKeySelector struct {
	Key      string `json:"key,omitempty"`
	Name     string `json:"name,omitempty"`
	Optional bool   `json:"optional,omitempty"`
}

type V1ObjectFieldSelector struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	FieldPath  string `json:"fieldPath,omitempty"`
}

type V1ResourceFieldSelector struct {
	ContainerName string `json:"containerName,omitempty"`
	Divisor       string `json:"divisor,omitempty"`
	Resource      string `json:"resource,omitempty"`
}

type V1SecretKeySelector struct {
	Key      string `json:"key,omitempty"`
	Name     string `json:"name,omitempty"`
	Optional string `json:"optional,omitempty"`
}

type V1EnvVarSource struct {
	ConfigMapKeyRef  *V1ConfigMapKeySelector  `json:"configMapKeyRef,omitempty"`
	FieldRef         *V1ObjectFieldSelector   `json:"fieldRef,omitempty"`
	ResourceFieldRef *V1ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
	SecretKeyRef     *V1SecretKeySelector     `json:"secretKeyRef,omitempty"`
}

type V1EnvVar struct {
	Name      string          `json:"name,omitempty"`
	Value     string          `json:"value,omitempty"`
	ValueFrom *V1EnvVarSource `json:"valueFrom,omitempty"`
}

type ImageRegistrySecret struct {
	Name string `json:"name,omitempty"`
}

type EmptyDir struct {
	Name string `json:"name,omitempty"`
}

type ConfigFile EmptyDir

type Secret EmptyDir

type VolumeClaimTemplates struct {
	Name        string `json:"name,omitempty"`
	StorageSize int    `json:"storageSize,omitempty"`
	DiskType    string `json:"diskType,omitempty"`
}

type Volume struct {
	EmptyDir             *[]EmptyDir             `json:"emptyDir,omitempty"`
	ConfigMap            *[]ConfigFile           `json:"configMap,omitempty"`
	Secret               *[]Secret               `json:"secret,omitempty"`
	VolumeClaimTemplates *[]VolumeClaimTemplates `json:"volumeClaimTemplates,omitempty"`
}

type Tag struct {
	TagKey   string `json:"tagKey,omitempty"`
	TagValue string `json:"tagValue,omitempty"`
}

type Region string

const (
	RegionCentralChina Region = "CENTRAL_CHINA"
	RegionEastChina    Region = "EAST_CHINA"
	RegionNorthChina   Region = "NORTH_CHINA"
	RegionSouthChina   Region = "SOUTH_CHINA"
	RegionNorthEast    Region = "NORTH_EAST"
	RegionNorthWest    Region = "NORTH_WEST"
	RegionSouthWest    Region = "SOUTH_WEST"
)

type ServiceProvider string

const (
	ServiceChinaMobile  ServiceProvider = "CHINA_MOBILE"
	ServiceChinaUnicom  ServiceProvider = "CHINA_UNICOM"
	ServiceChinaTelecom ServiceProvider = "CHINA_TELECOM"
	ServiceTripleLine   ServiceProvider = "TRIPLE_LINE"
)

type DeploymentInstance struct {
	Region          Region          `json:"region,omitempty"`
	ServiceProvider ServiceProvider `json:"serviceProvider,omitempty"`
	Replicas        int             `json:"replicas,omitempty"`
	City            string          `json:"city,omitempty"`
}

type ResourceBriefVo struct {
	ServiceId            string                `json:"serviceId"`
	ServiceName          string                `json:"serviceName"`
	ResourceId           string                `json:"resourceId"`
	ResourceName         string                `json:"resourceName"`
	Labels               map[string]string     `json:"labels"`
	TotalCpu             int                   `json:"totalCpu"`
	TotalMem             int                   `json:"totalMem"`
	TotalGpu             int                   `json:"totalGpu"`
	TotalPods            int                   `json:"totalPods"`
	RunningPods          int                   `json:"runningPods"`
	TotalDeploy          int                   `json:"totalDeploy"`
	IngressBandwidth     string                `json:"ingressBandwidth"`
	DeployInstance       DeploymentInstance    `json:"deployInstance"`
	ImageList            []string              `json:"imageList"`
	Containers           []ContainerDetails    `json:"containers"`
	ImageRegistrySecrets []ImageRegistrySecret `json:"imageRegistrySecrets"`
}

type ServiceDetailsVo struct {
	ServiceId            string                `json:"serviceId"`
	ServiceName          string                `json:"serviceName"`
	Status               string                `json:"status"`
	TotalCpu             int                   `json:"totalCpu"`
	TotalMem             int                   `json:"totalMem"`
	TotalGpu             int                   `json:"totalGpu"`
	TotalDisk            int                   `json:"totalDisk"`
	TotalPods            int                   `json:"totalPods"`
	RunningPods          int                   `json:"runningPods"`
	RegionSize           int                   `json:"regionSize"`
	TagsMap              []Tag                 `json:"tagsMap"`
	DeployInstances      []DeploymentInstance  `json:"deployInstances"`
	ResourceBriefVos     []ResourceBriefVo     `json:"resourceBriefVos"`
	ImageRegistrySecrets []ImageRegistrySecret `json:"imageRegistrySecrets"`
	LogCollectDetail     LogCollectDetail      `json:"logCollectDetail"`
	CreateTime           string                `json:"createTime"`
	LastUpdateTime       string                `json:"lastUpdateTime"`
}

type ServiceBriefVo struct {
	ServiceId       string               `json:"serviceId"`
	ServiceName     string               `json:"serviceName"`
	Level           string               `json:"level"`
	Status          string               `json:"status"`
	TotalCpu        int                  `json:"totalCpu"`
	TotalMen        int                  `json:"totalMem"`
	TotalGpu        int                  `json:"totalGpu"`
	TotalDisk       int                  `json:"totalDisk"`
	Regions         int                  `json:"regions"`
	TotalPods       int                  `json:"totalPods"`
	RunningPods     int                  `json:"runningPods"`
	TagsMap         []Tag                `json:"tagMap"`
	DeployInstances []DeploymentInstance `json:"deployInstances"`
	CreateTime      string               `json:"createTime"`
	LastUpdateTime  string               `json:"lastUpdateTime"`
}

type ContainerDetails struct {
	Name         string             `json:"name,omitempty"`
	ImageVersion string             `json:"imageVersion,omitempty"`
	ImageAddress string             `json:"imageAddress,omitempty"`
	Memory       int                `json:"memory,omitempty"`
	Cpu          int                `json:"cpu,omitempty"`
	Gpu          int                `json:"gpu,omitempty"`
	WorkingDir   string             `json:"workingDir,omitempty"`
	Commands     []string           `json:"commands,omitempty"`
	Args         []string           `json:"args,omitempty"`
	VolumeMounts *[]V1VolumeMount   `json:"volumeMounts,omitempty"`
	Ports        *[]V1ContainerPort `json:"ports,omitempty"`
	Envs         *[]V1EnvVar        `json:"envs,omitempty"`
}

type LogCollectDetail struct {
	ServiceId      string `json:"serviceId,omitempty"`
	LogCollect     bool   `json:"logCollect,omitempty"`
	LogPath        string `json:"logPath,omitempty"`
	JsonAnalysis   bool   `json:"jsonAnalysis,omitempty"`
	PushLog        bool   `json:"pushLog,omitempty"`
	Standard       bool   `json:"standard,omitempty"`
	Custom         bool   `json:"custom,omitempty"`
	LogOutputType  string `json:"logOutputType,omitempty"`
	EsIP           string `json:"esIP,omitempty"`
	EsPort         int    `json:"esPort,omitempty"`
	EsIndex        string `json:"esIndex,omitempty"`
	Encrypted      bool   `json:"encrypted,omitempty"`
	EsUserName     string `json:"esUserName,omitempty"`
	EsUserPassword string `json:"esUserPassword,omitempty"`
}

type CreateServiceArgs struct {
	ServiceName          string                 `json:"serviceName,omitempty"`
	PaymentMethod        string                 `json:"paymentMethod,omitempty"`
	ContainerGroupName   string                 `json:"containerGroupName,omitempty"`
	Containers           *[]ContainerDetails    `json:"containers,omitempty"`
	ImageRegistrySecrets *[]ImageRegistrySecret `json:"imageRegistrySecrets,omitempty"`
	Volumes              *Volume                `json:"volumes,omitempty"`
	NeedPublicIp         bool                   `json:"needPublicIp,omitempty"`
	Bandwidth            int                    `json:"bandwidth,omitempty"`
	Tags                 *[]Tag                 `json:"tags,omitempty"`
	DeployInstances      *[]DeploymentInstance  `json:"deployInstances,omitempty"`
	LogCollectDetail     *LogCollectDetail      `json:"logCollectDetail,omitempty"`
}

type CreateServiceResult struct {
	Details ServiceBriefVo `json:"details"`
	Result  bool           `json:"result"`
	Action  string         `json:"action"`
}

type OrderModel struct {
	OrderBy string `json:"orderBy"`
	Order   string `json:"order"`
}

type ListServiceResult struct {
	Result     []ServiceBriefVo `json:"result"`
	OrderBy    string           `json:"orderBy"`
	Order      string           `json:"order"`
	PageNo     int              `json:"pageNo"`
	PageSize   int              `json:"pageSize"`
	TotalCount int              `json:"totalCount"`
}

type MetricsType string

const (
	MetricsTypeCpu               MetricsType = "CPU"
	MetricsTypeMemory            MetricsType = "MEMORY"
	MetricsTypeBandwidthReceive  MetricsType = "BANDWIDTH_RECEIVE"
	MetricsTypeBandwidthTransmit MetricsType = "BANDWIDTH_TRANSMIT"
	MetricsTypeTrafficReceive    MetricsType = "TRAFFIC_RECEIVE"
	MetricsTypeTrafficTransmit   MetricsType = "TRAFFIC_TRANSMIT"

	MetricsTypeNodeBwReceive    MetricsType = "NODE_BW_RECEIVE"
	MetricsTypeNodeBwTransmit   MetricsType = "NODE_BW_TRANSMIT"
	MetricsTypeNodeLbBwReceive  MetricsType = "NODE_LB_BW_RECEIVE"
	MetricsTypeNodeLbBwTransmit MetricsType = "NODE_LB_BW_TRANSMIT"

	MetricsTypeRequestNum   MetricsType = "REQUEST_NUMBER"
	MetricsTypeRequestRate  MetricsType = "REQUEST_RATE"
	MetricsTypeRequestDelay MetricsType = "REQUEST_DELAY"

	MetricsTypeUnknown MetricsType = "UNKNOWN"
)

type Metric struct {
	TimeInSecond int     `json:"timeInSecond"`
	Value        float64 `json:"value"`
}

type ServiceMetricsResult struct {
	Metrics    []Metric `json:"metrics"`
	MaxValue   float64  `json:"maxValue"`
	AvgValue   float64  `json:"avgValue"`
	TotalValue float64  `json:"totalValue"`
}

type GetServiceArgs struct {
	ServiceId string
}

type ServiceAction string

const (
	ServiceActionStart ServiceAction = "start"
	ServiceActionStop  ServiceAction = "stop"
)

type ServiceActionResult struct {
	Result bool              `json:"result"`
	Action string            `json:"action"`
	Detail map[string]string `json:"detail"`
}

type UpdateServiceType string

const (
	UpdateServiceTypeName         UpdateServiceType = "NAME"
	UpdateServiceTypeReplicas     UpdateServiceType = "REPLICAS"
	UpdateServiceTypeNameResource UpdateServiceType = "RESOURCE"
)

type UpdateServiceArgs struct {
	Type                 UpdateServiceType      `json:"type,omitempty"`
	DeployInstances      *[]DeploymentInstance  `json:"deployInstances,omitempty"`
	ServiceName          string                 `json:"serviceName,omitempty"`
	Containers           *[]ContainerDetails    `json:"containers,omitempty"`
	ImageRegistrySecrets *[]ImageRegistrySecret `json:"imageRegistrySecrets,omitempty"`
	Bandwidth            float32                `json:"bandwidth,omitempty"`
}

type UpdateServiceResult struct {
	Result  bool             `json:"result"`
	Action  string           `json:"action"`
	Details []ServiceBriefVo `json:"details"`
}

type ServiceBatchOperateArgs struct {
	IdList []string `json:"idList,omitempty"`
	Action string   `json:"action,omitempty"`
}

type OperationVo struct {
	ResourceId string `json:"resourceId"`
	Success    bool   `json:"success"`
	Error      string `json:"error"`
}

type ServiceBatchOperateResult struct {
	Result  bool          `json:"result"`
	Action  string        `json:"action"`
	Details []OperationVo `json:"details"`
}

type ListDeploymentArgs struct {
	DeploymentID string `json:"deploymentID"`
}

type Networks struct {
	NetType string `json:"netType,omitempty"`
	NetName string `json:"netName,omitempty"`
}
type NetworkConfig struct {
	NodeType     string      `json:"nodeType,omitempty"` //NoneType
	NetworksList *[]Networks `json:"networksList,omitempty"`
}

type GpuRequest struct {
	Type string `json:"type,omitempty"`
	Num  int    `json:"num,omitempty"`
}

type DiskType string

const (
	DiskTypeNVME             DiskType = "NVME"
	DiskTypeSATA             DiskType = "SATA"
	DiskTypeCDSHDD           DiskType = "CDS_HDD"
	DiskTypeCDSSSD           DiskType = "CDS_SSD"
	DiskTypeHDDPASSTHROUGH4T DiskType = "HDD_PASSTHROUGH_4T"
	DiskTypeSSDPASSTHROUGH4T DiskType = "SSD_PASSTHROUGH_4T"
)

type VolumeConfig struct {
	Name            string   `json:"name,omitempty"`
	VolumeType      DiskType `json:"volumeType,omitempty"`
	SizeInGB        int      `json:"sizeInGB,omitempty"`
	PvcName         string   `json:"pvcName,omitempty"`
	PassthroughCode string   `json:"passthroughCode,omitempty"`
}

type SystemVolumeConfig struct {
	VolumeType DiskType `json:"volumeType,omitempty"`
	SizeInGB   int      `json:"sizeInGB,omitempty"`
	Name       string   `json:"name,omitempty"`
	PvcName    string   `json:"pvcName,omitempty"`
}

type DnsConfig struct {
	DnsType    string `json:"dnsType,omitempty"`
	DnsAddress string `json:"dnsAddress,omitempty"`
}

type KeyConfig struct {
	Type             string   `json:"type,omitempty"`
	AdminPass        string   `json:"adminPass,omitempty"`
	BccKeyPairIdList []string `json:"bccKeyPairIdList,omitempty"`
}

type CreateVmServiceArgs struct {
	ServiceName       string                `json:"serviceName,omitempty"`
	VmName            string                `json:"vmName,omitempty"`
	NeedIpv6PublicIp  bool                  `json:"needIpv6PublicIp,omitempty"`
	DisableIntranet   bool                  `json:"disableIntranet,omitempty"`
	DisableCloudInit  bool                  `json:"disableCloudInit,omitempty"`
	PaymentMethod     string                `json:"paymentMethod,omitempty"`
	NeedPublicIp      bool                  `json:"needPublicIp,omitempty"`
	Bandwidth         int                   `json:"bandwidth,omitempty"`
	DeployInstances   *[]DeploymentInstance `json:"deployInstances,omitempty"`
	Cpu               int                   `json:"cpu,omitempty"`
	Memory            int                   `json:"memory,omitempty"`
	Gpu               *GpuRequest           `json:"gpu,omitempty"`
	ImageId           string                `json:"imageId,omitempty"`
	ImageType         ImageType             `json:"imageType,omitempty"`
	DataVolumeList    *[]VolumeConfig       `json:"dataVolumeList,omitempty"`
	SystemVolume      *SystemVolumeConfig   `json:"systemVolume,omitempty"`
	NetworkConfigList *[]NetworkConfig      `json:"networkConfigList,omitempty"`
	AdminPass         string                `json:"adminPass,omitempty"`
	DnsConfig         *DnsConfig            `json:"dnsConfig,omitempty"`
	KeyConfig         *KeyConfig            `json:"keyConfig,omitempty"`
}

type ImageDetail struct {
	Id      string `json:"id"`
	ImageId string `json:"imageId"`
	Name    string `json:"name"`

	NameFri    string `json:"nameFri"`
	ImageType  string `json:"imageType"`
	SnapshotId string `json:"snapshotId"`
	Cpu        int    `json:"cpu"`
	Memory     int    `json:"memory"`
	OsType     string `json:"osType"`
	OsVersion  string `json:"osVersion"`
	OsName     string `json:"osName"`
	OsBuild    string `json:"osBuild"`
	OsLang     string `json:"osLang"`
	DiskSize   int    `json:"diskSize"`

	CreateTime          string `json:"createTime"`
	Status              string `json:"status"`
	MinMem              int    `json:"minMem"`
	MinCpu              int    `json:"minCpu"`
	MinDiskGb           int    `json:"minDiskGb"`
	Desc                string `json:"desc"`
	OsArch              string `json:"osArch"`
	EphemeralSize       int    `json:"ephemeralSize"`
	ImageDescription    string `json:"imageDescription"`
	ShareToUserNumLimit int    `json:"shareToUserNumLimit"`
	SharedToUserNum     int    `json:"sharedToUserNum"`
	FpgaType            string `json:"fpgaType"`
}
type VmServiceBriefVo struct {
	ServiceId        string               `json:"serviceId"`
	ServiceName      string               `json:"serviceName"`
	Status           string               `json:"status"`
	TotalCpu         int                  `json:"totalCpu"`
	TotalMem         int                  `json:"totalMem"`
	TotalDisk        int                  `json:"totalDisk"`
	TotalRootDisk    int                  `json:"totalRootDisk"`
	Regions          int                  `json:"regions"`
	DeployInstances  []DeploymentInstance `json:"deployInstances"`
	TotalInstances   int                  `json:"totalInstances"`
	RunningInstances int                  `json:"runningInstances"`
	OsImage          ImageDetail          `json:"osImage"`
	CreateTime       string               `json:"createTime"`
	TotalGpu         int                  `json:"totalGpu"`
}

type CreateVmServiceResult struct {
	Details VmServiceBriefVo `json:"details"`
	Result  bool             `json:"result"`
	Action  string           `json:"action"`
}

type DeleteVmServiceArgs struct {
	ServiceId string `json:"serviceId"`
}

type DeleteVmServiceResult struct {
	Details map[string]string `json:"details"`
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
}

type UpdateVmType string

const (
	UpdateVmTypeServiceName UpdateVmType = "serviceName"
	UpdateVmTypeVmName      UpdateVmType = "vmName"
	UpdateVmPassWord        UpdateVmType = "password"
	UpdateVmReplicas        UpdateVmType = "replicas"
	UpdateVmResource        UpdateVmType = "resource"
)

type UpdateBecVmForm struct {
	Type           UpdateVmType        `json:"type,omitempty"`
	Cpu            int                 `json:"cpu,omitempty"`
	Memory         int                 `json:"memory,omitempty"`
	NeedRestart    bool                `json:"needRestart,omitempty"`
	AdminPass      string              `json:"adminPass,omitempty"`
	ImageId        string              `json:"imageId,omitempty"`
	Bandwidth      int                 `json:"bandwidth,omitempty"`
	ImageType      ImageType           `json:"imageType,omitempty"`
	VmName         string              `json:"vmName,omitempty"`
	DataVolumeList *[]VolumeConfig     `json:"dataVolumeList,omitempty"`
	SystemVolume   *SystemVolumeConfig `json:"systemVolume,omitempty"`
}

type UpdateVmServiceArgs struct {
	UpdateBecVmForm
	ServiceName     string                `json:"serviceName,omitempty"`
	DeployInstances *[]DeploymentInstance `json:"deployInstances,omitempty"`
}

type UpdateVmServiceResult struct {
	Details VmServiceBriefVo `json:"details"`
	Result  bool             `json:"result"`
	Action  string           `json:"action"`
}

type ListVmServiceArgs struct {
	KeywordType string `json:"keywordType,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	PageNo      int    `json:"pageNo,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	Order       string `json:"order,omitempty"`
	OrderBy     string `json:"orderBy,omitempty"`
	Status      string `json:"status,omitempty"`
	Region      string `json:"region,omitempty"`
	OsName      string `json:"osName,omitempty"`
	ServiceId   string `json:"serviceId,omitempty"`
}

type ListVmServiceResult struct {
	Orders     []OrderModel       `json:"orders"`
	OrderBy    string             `json:"orderBy"`
	Order      string             `json:"order"`
	PageNo     int                `json:"pageNo"`
	PageSize   int                `json:"pageSize"`
	TotalCount int                `json:"totalCount"`
	Result     []VmServiceBriefVo `json:"result"`
}

type GetVmServiceDetailArgs struct {
	ServiceId string `json:"serviceId"`
}

type VmServiceDetailsVo struct {
	Bandwidth        string         `json:"bandwidth"`
	TotalBandwidth   string         `json:"totalBandwidth"`
	DataVolumeList   []VolumeConfig `json:"dataVolumeList"`
	SystemVolumeList []VolumeConfig `json:"systemVolumeList"`
}

type VmServiceAction string

const (
	VmServiceActionStart VmServiceAction = "start"
	VmServiceActionStop  VmServiceAction = "stop"
)

type VmServiceActionResult struct {
	Details map[string]string `json:"details"`
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
}

type VmServiceBatchActionResult struct {
	Result  bool          `json:"result"`
	Action  string        `json:"action"`
	Details []OperationVo `json:"details"`
}

type VmServiceBatchAction string

const (
	VmServiceBatchStart VmServiceBatchAction = "start"
	VmServiceBatchStop  VmServiceBatchAction = "stop"
)

type VmServiceBatchActionArgs struct {
	IdList []string             `json:"idList,omitempty"`
	Action VmServiceBatchAction `json:"action,omitempty"`
}

type CreateVmImageArgs struct {
	VmId string `json:"vmId,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreateVmImageResult struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

type VmImageOperateResult struct {
	Success bool `json:"success"`
	Result  bool `json:"result"`
}

type UpdateVmImageArgs struct {
	Name string `json:"name,omitempty"`
}

type ListVmImageArgs struct {
	KeywordType string `json:"keywordType"`
	Keyword     string `json:"keyword"`
	PageNo      int    `json:"pageNo,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	Order       string `json:"order,omitempty"`
	OrderBy     string `json:"orderBy,omitempty"`
	Status      string `json:"status"`
	Region      string `json:"region"`
	OsName      string `json:"osName"`
	ServiceId   string `json:"serviceId"`
	Type        string `json:"type,omitempty"`
}

type VmImageVo struct {
	ImageId    string `json:"imageId"`
	Status     string `json:"status"`
	BccImageId string `json:"bccImageId"`
	Name       string `json:"name"`
	AccountId  string `json:"accountId"`
	ImageType  string `json:"imageType"`
	SystemDisk int    `json:"systemDisk"`
	OsType     string `json:"osType"`
	OsVersion  string `json:"osVersion"`
	OsName     string `json:"osName"`
	OsBuild    string `json:"osBuild"`
	OsLang     string `json:"osLang"`
	OsArch     string `json:"osArch"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type ListVmImageResult struct {
	Orders     []OrderModel `json:"orders"`
	OrderBy    string       `json:"orderBy"`
	Order      string       `json:"order"`
	PageNo     int          `json:"pageNo"`
	PageSize   int          `json:"pageSize"`
	TotalCount int          `json:"totalCount"`
	Result     []VmImageVo  `json:"result"`
}

type CreateBlbArgs struct {
	LbType               string          `json:"lbType,omitempty"`
	PaymentMethod        string          `json:"paymentMethod,omitempty"`
	RegionSelection      string          `json:"regionSelection,omitempty"`
	Region               Region          `json:"region,omitempty"`
	City                 string          `json:"city,omitempty"`
	ServiceProvider      ServiceProvider `json:"serviceProvider,omitempty"`
	BlbName              string          `json:"blbName,omitempty"`
	NeedPublicIp         bool            `json:"needPublicIp,omitempty"`
	BandwidthInMbpsLimit int             `json:"bandwidthInMbpsLimit,omitempty"`
	Tags                 *[]Tag          `json:"tags,omitempty"`
	Listeners            *[]Listeners    `json:"listeners,omitempty"`
}

type Protocol string

const (
	ProtocolTcp   Protocol = "TCP"
	ProtocolUdp   Protocol = "UDP"
	ProtocolHttp  Protocol = "HTTP"
	ProtocolHttps Protocol = "HTTPS"
	ProtocolSsl   Protocol = "SSL"
)

type Listeners struct {
	Protocol         Protocol `json:"protocol,omitempty"`
	Port             int      `json:"port,omitempty"`
	BackendPort      int      `json:"backendPort,omitempty"`
	KeepaliveTimeout int      `json:"keepaliveTimeout,omitempty"`
	Scheduler        LbMode   `json:"scheduler,omitempty"`
	EnableCipTTM     bool     `json:"enableCipTTM,omitempty"`
	EnableVipTTM     bool     `json:"enableVipTTM,omitempty"`

	// health check config
	HealthCheckInterval  int    `json:"healthCheckInterval,omitempty"`
	HealthCheckRetry     int    `json:"healthCheckRetry,omitempty"`
	HealthCheckTimeout   int    `json:"healthCheckTimeout,omitempty"`
	UdpHealthCheckString string `json:"udpHealthCheckString,omitempty"`
	HealthCheckType      string `json:"healthCheckType,omitempty"`
}

type BlbInstanceVo struct {
	BlbId                string          `json:"blbId"`
	BlbName              string          `json:"blbName"`
	Status               string          `json:"status"`
	LbType               string          `json:"lbType"`
	Region               Region          `json:"region"`
	ServiceProvider      ServiceProvider `json:"serviceProvider"`
	City                 string          `json:"city"`
	PublicIp             string          `json:"publicIp"`
	CmPublicIP           string          `json:"cmPublicIP"`
	CtPublicIP           string          `json:"ctPublicIP"`
	UnPublicIP           string          `json:"unPublicIP"`
	InternalIp           string          `json:"internalIp"`
	Ports                []Listeners     `json:"ports"`
	PodCount             int             `json:"podCount"`
	BandwidthInMbpsLimit int             `json:"bandwidthInMbpsLimit"`
	CreateTime           string          `json:"createTime"`
}

type CreateBlbResult struct {
	Result  bool          `json:"result"`
	Action  string        `json:"action"`
	Details BlbInstanceVo `json:"details"`
}

type DeleteBlbResult struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details map[string]string `json:"details"`
}

type GetBlbListResult struct {
	Orders     []OrderModel    `json:"orders"`
	OrderBy    string          `json:"orderBy"`
	PageNo     int             `json:"pageNo"`
	PageSize   int             `json:"pageSize"`
	TotalCount int             `json:"totalCount"`
	Result     []BlbInstanceVo `json:"result"`
}

type UpdateBlbArgs struct {
	BlbName              string `json:"blbName,omitempty"`
	BandwidthInMbpsLimit int    `json:"bandwidthInMbpsLimit,omitempty"`
	Type                 string `json:"type,omitempty"`
}

type UpdateBlbResult struct {
	Result  bool          `json:"result"`
	Action  string        `json:"action"`
	Details BlbInstanceVo `json:"details"`
}

type LbMode string

const (
	LbModeWrr     LbMode = "wrr"
	LbModeMinConn LbMode = "minconn"
	LbModeSrch    LbMode = "srch"
)

type BlbMonitorArgs struct {
	FrontendPort     *Port        `json:"frontendPort,omitempty"`
	BackendPort      int          `json:"backendPort,omitempty"`
	LbMode           LbMode       `json:"lbMode,omitempty"`
	KeepaliveTimeout int          `json:"keepaliveTimeout,omitempty"`
	HealthCheck      *HealthCheck `json:"healthCheck,omitempty"`
	EnableCipTTM     bool         `json:"enableCipTTM,omitempty"`
	EnableVipTTM     bool         `json:"enableVipTTM,omitempty"`
}

type BlbMonitorResult struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details map[string]string `json:"details"`
}

type BlbMonitorListResult struct {
	Orders     []OrderModel `json:"orders"`
	OrderBy    string       `json:"orderBy"`
	Order      string       `json:"order"`
	PageNo     int          `json:"pageNo"`
	PageSize   int          `json:"pageSize"`
	TotalCount int          `json:"totalCount"`
	Result     []Listeners  `json:"result"`
}

type Port struct {
	Protocol Protocol `json:"protocol,omitempty"`
	Port     int      `json:"port,omitempty"`
}

type Stats struct {
	Health   bool     `json:"health"`
	Port     int      `json:"port"`
	Protocol Protocol `json:"protocol"`
}

type BlbBackendPodBriefVo struct {
	PodName     string  `json:"podName"`
	PodStatus   string  `json:"podStatus"`
	PodIp       string  `json:"podIp"`
	BackendPort []Stats `json:"backendPort"`
	Weight      int     `json:"weight"`
}

type GetBlbBackendPodListResult struct {
	Orders     []OrderModel           `json:"orders"`
	OrderBy    string                 `json:"orderBy"`
	Order      string                 `json:"order"`
	PageNo     int                    `json:"pageNo"`
	PageSize   int                    `json:"pageSize"`
	TotalCount int                    `json:"totalCount"`
	Result     []BlbBackendPodBriefVo `json:"result"`
}

type BatchCreateBlbArgs struct {
	LbType               string                `json:"lbType,omitempty"`
	PaymentMethod        string                `json:"paymentMethod,omitempty"`
	RegionSelection      string                `json:"regionSelection,omitempty"`
	DeployInstances      *[]DeploymentInstance `json:"deployInstances,omitempty"`
	BlbName              string                `json:"blbName,omitempty"`
	NeedPublicIp         bool                  `json:"needPublicIp,omitempty"`
	BandwidthInMbpsLimit int                   `json:"bandwidthInMbpsLimit,omitempty"`
	Tags                 *[]Tag                `json:"tags,omitempty"`
	Listeners            *[]Listeners          `json:"listeners,omitempty"`
}

type BatchCreateBlbResult struct {
	Result  bool            `json:"result"`
	Action  string          `json:"action"`
	Details []BlbInstanceVo `json:"details"`
}

type BatchDeleteBlbResult struct {
	Result  bool          `json:"result"`
	Action  string        `json:"action"`
	Details []OperationVo `json:"details"`
}

type PortGroup struct {
	Port        int `json:"port,omitempty"`
	BackendPort int `json:"backendPort,omitempty"`
}
type HealthCheck struct {
	TimeoutInSeconds   int    `json:"timeoutInSeconds,omitempty"`
	IntervalInSeconds  int    `json:"intervalInSeconds,omitempty"`
	UnhealthyThreshold int    `json:"unhealthyThreshold,omitempty"`
	HealthyThreshold   int    `json:"healthyThreshold,omitempty"`
	HealthCheckString  string `json:"healthCheckString,omitempty"`
	HealthCheckType    string `json:"healthCheckType,omitempty"`
}
type BatchCreateBlbMonitorArg struct {
	Protocol         Protocol     `json:"protocol,omitempty"`
	PortGroups       *[]PortGroup `json:"portGroups,omitempty"`
	LbMode           LbMode       `json:"lbMode,omitempty"`
	KeepaliveTimeout int          `json:"keepaliveTimeout,omitempty"`
	HealthCheck      *HealthCheck `json:"healthCheck,omitempty"`
	EnableCipTTM     bool         `json:"enableCipTTM,omitempty"`
	EnableVipTTM     bool         `json:"enableVipTTM,omitempty"`
}

type BatchCreateBlbMonitorResult struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details map[string]string `json:"details"`
}

type ListRequest struct {
	KeywordType string `json:"keywordType"`
	Keyword     string `json:"keyword"`
	PageNo      int    `json:"pageNo,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	Order       string `json:"order,omitempty"`
	OrderBy     string `json:"orderBy,omitempty"`
	Status      string `json:"status"`
	Region      string `json:"region"`
	OsName      string `json:"osName"`
	ServiceId   string `json:"serviceId"`
}

type VmInstanceDetailsVo struct {
	VmInstanceBriefVo
	RootDiskSize   int                `json:"rootDiskSize"`
	DataStorage    int                `json:"dataStorage"`
	DataVolumeList []VolumeConfig     `json:"dataVolumeList"`
	SystemVolume   SystemVolumeConfig `json:"systemVolume"`
	RackId         string             `json:"rackId,omitempty"`
	HostId         string             `json:"hostId,omitempty"`
	SwitchId       string             `json:"switchId"`
	PrivateIps     []string           `json:"privateIps"`
}

type LogicPageVmInstanceResult struct {
	Orders     []OrderModel          `json:"orders"`
	OrderBy    string                `json:"orderBy"`
	Order      string                `json:"order"`
	PageNo     int                   `json:"pageNo"`
	PageSize   int                   `json:"pageSize"`
	TotalCount int                   `json:"totalCount"`
	Result     []VmInstanceDetailsVo `json:"result"`
}

type GetNodeVmInstanceListResult struct {
	Result  []VmInstanceBriefVo `json:"result"`
	Success bool                `json:"success"`
}

type ActionInfoVo struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details map[string]string `json:"details"`
}

type ImageType string

const (
	ImageTypeBcc ImageType = "bcc"
	ImageTypeBec ImageType = "bec"
)

type UpdateVmDeploymentArgs struct {
	Type           string              `json:"type,omitempty"`
	Cpu            int                 `json:"cpu,omitempty"`
	Memory         int                 `json:"memory,omitempty"`
	NeedRestart    bool                `json:"needRestart,omitempty"`
	AdminPass      string              `json:"adminPass,omitempty"`
	ImageId        string              `json:"imageId,omitempty"`
	Bandwidth      int                 `json:"bandwidth,omitempty"`
	ImageType      ImageType           `json:"imageType,omitempty"`
	VmName         string              `json:"vmName,omitempty"`
	DataVolumeList *[]VolumeConfig     `json:"dataVolumeList,omitempty"`
	SystemVolume   *SystemVolumeConfig `json:"systemVolume,omitempty"`
	KeyConfig      *KeyConfig          `json:"keyConfig,omitempty"`
}

type IpInfo struct {
	ServiceProvider ServiceProvider `json:"serviceProvider"`
	Ip              string          `json:"ip"`
	Ipv6            string          `json:"ipv6"`
}

type IpPackageVo struct {
	PublicIp         string          `json:"publicIp"`
	Ipv6PublicIp     string          `json:"ipv6PublicIp"`
	InternalIp       string          `json:"internalIp"`
	MultiplePublicIp []IpInfo        `json:"multiplePublicIp"`
	ServiceProvider  ServiceProvider `json:"serviceProvider"`
}

type VmInstanceBriefVo struct {
	IpPackageVo
	VmId             string      `json:"vmId"`
	Uuid             string      `json:"uuid"`
	VmName           string      `json:"vmName"`
	Status           string      `json:"status"`
	Cpu              int         `json:"cpu"`
	Mem              int         `json:"mem"`
	Gpu              int         `json:"gpu"`
	Region           Region      `json:"region"`
	City             string      `json:"city"`
	NeedPublicIp     bool        `json:"needPublicIp"`
	NeedIpv6PublicIp bool        `json:"needIpv6PublicIp"`
	Bandwidth        string      `json:"bandwidth"`
	OsImage          ImageDetail `json:"osImage"`
	ServiceId        string      `json:"serviceId"`
	CreateTime       string      `json:"createTime"`
}

type UpdateVmDeploymentResult struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details VmInstanceBriefVo `json:"details"`
}

type ReinstallVmInstanceArg struct {
	AdminPass     string     `json:"adminPass,omitempty"`
	ImageId       string     `json:"imageId,omitempty"`
	ImageType     ImageType  `json:"imageType,omitempty"`
	ResetDataDisk bool       `json:"resetDataDisk,omitempty"`
	KeyConfig     *KeyConfig `json:"keyConfig,omitempty"`
}

type ReinstallVmInstanceResult struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details VmInstanceBriefVo `json:"details"`
}

type VmInstanceBatchOperateAction string

const (
	VmInstanceBatchOperateStart   VmInstanceBatchOperateAction = "start"
	VmInstanceBatchOperateStop    VmInstanceBatchOperateAction = "stop"
	VmInstanceBatchOperateRestart VmInstanceBatchOperateAction = "restart"
)

type OperateVmDeploymentResult struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details map[string]string `json:"details"`
}

type VmConfigResult struct {
	Cpu             int                `json:"cpu"`
	Mem             int                `json:"mem"`
	Region          Region             `json:"region"`
	ServiceProvider ServiceProvider    `json:"serviceProvider"`
	City            string             `json:"city"`
	NeedPublicIp    bool               `json:"needPublicIp"`
	Bandwidth       string             `json:"bandwidth"`
	OsImage         ImageDetail        `json:"osImage"`
	DataVolumeList  []VolumeConfig     `json:"dataVolumeList"`
	SystemVolume    SystemVolumeConfig `json:"systemVolume"`
}

type Backends struct {
	Name   string `json:"name,omitempty"`
	Ip     string `json:"ip,omitempty"`
	Weight int    `json:"weight,omitempty"`
}

type LbDeployPo struct {
	ServiceName string     `json:"serviceName"`
	Backends    []Backends `json:"backends"`
}

type GetBlbBackendBindingStsListResult struct {
	Orders     []OrderModel `json:"orders"`
	OrderBy    string       `json:"orderBy"`
	Order      string       `json:"order"`
	PageNo     int          `json:"pageNo"`
	PageSize   int          `json:"pageSize"`
	TotalCount int          `json:"totalCount"`
	Result     []LbDeployPo `json:"result"`
}

type BlbBindingForm struct {
	DeploymentId  string      `json:"deploymentId,omitempty"`
	DefaultWeight int         `json:"defaultWeight,omitempty"`
	PodWeight     *[]Backends `json:"podWeight,omitempty"`
}
type CreateBlbBindingArgs struct {
	BindingForms *[]BlbBindingForm `json:"bindingForms,omitempty"`
}

type DeleteBlbBindPodArgs struct {
	PodWeightList *[]Backends `json:"podWeightList,omitempty"`
	DeploymentIds []string    `json:"deploymentIds,omitempty"`
}

type DeleteBlbBindPodResult struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details map[string]string `json:"details"`
}

type CreateBlbBindingResult struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details map[string]string `json:"details"`
}
type UpdateBindPodWeightArgs struct {
	PodWeightList *[]Backends `json:"podWeightList,omitempty"`
	DeploymentIds []string    `json:"deploymentIds,omitempty"`
}

type UpdateBindPodWeightResult struct {
	Result  bool              `json:"result"`
	Action  string            `json:"action"`
	Details map[string]string `json:"details"`
}

type CreateVmPrivateIpForm struct {
	SecondaryPrivateIpAddressCount int      `json:"secondaryPrivateIpAddressCount,omitempty"`
	PrivateIps                     []string `json:"privateIps,omitempty"`
}

type DeleteVmPrivateIpForm struct {
	PrivateIps []string `json:"privateIps,omitempty"`
}

type IpamResultVo struct {
	Success bool     `json:"success"`
	ErrCode string   `json:"errCode"`
	ErrMsg  string   `json:"errMsg"`
	Ips     []string `json:"ips"`
	ErrIPs  []string `json:"errIPs"`
}

type VmPrivateIpResult struct {
	Result  IpamResultVo `json:"result"`
	Success bool         `json:"success"`
}

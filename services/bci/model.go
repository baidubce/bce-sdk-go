/*
 * Copyright 2023 Baidu, Inc.
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

package bci

type CreateInstanceArgs struct {
	ClientToken                   string                    `json:"-"`
	Name                          string                    `json:"name"`
	ZoneName                      string                    `json:"zoneName,omitempty"`
	SecurityGroupIds              []string                  `json:"securityGroupIds,omitempty"`
	SubnetIds                     []string                  `json:"subnetIds,omitempty"`
	RestartPolicy                 string                    `json:"restartPolicy,omitempty"`
	EipIp                         string                    `json:"eipIp,omitempty"`
	AutoCreateEip                 bool                      `json:"autoCreateEip,omitempty"`
	EipName                       string                    `json:"eipName,omitempty"`
	EipRouteType                  string                    `json:"eipRouteType,omitempty"`
	EipBandwidthInMbps            int                       `json:"eipBandwidthInMbps,omitempty"`
	EipBillingMethod              string                    `json:"eipBillingMethod,omitempty"`
	GPUType                       string                    `json:"gpuType,omitempty"`
	TerminationGracePeriodSeconds int64                     `json:"terminationGracePeriodSeconds,omitempty"`
	HostName                      string                    `json:"hostName,omitempty"`
	Tags                          []Tag                     `json:"tags,omitempty"`
	ImageRegistryCredentials      []ImageRegistryCredential `json:"imageRegistryCredentials,omitempty"`
	Containers                    []Container               `json:"containers,omitempty"`
	InitContainers                []Container               `json:"initContainers,omitempty"`
	Volume                        *Volume                   `json:"volume,omitempty"`
}

type CreateInstanceResult struct {
	InstanceId string `json:"instanceId,omitempty"`
}

type ListInstanceArgs struct {
	KeywordType string `json:"keywordType,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	Marker      string `json:"marker,omitempty"`
	MaxKeys     int    `json:"maxKeys,omitempty"`
}

type ListInstanceResult struct {
	Marker      string          `json:"marker,omitempty"`
	IsTruncated bool            `json:"isTruncated,omitempty"`
	NextMarker  string          `json:"nextMarker,omitempty"`
	MaxKeys     int             `json:"maxKeys,omitempty"`
	Result      []InstanceModel `json:"result,omitempty"`
}

type GetInstanceArgs struct {
	InstanceId string `json:"instanceId"`
}

type GetInstanceResult struct {
	Instance *InstanceDetailModel `json:"instance,omitempty"`
}

type DeleteInstanceArgs struct {
	InstanceId         string `json:"instanceId"`
	RelatedReleaseFlag bool   `json:"relatedReleaseFlag,omitempty"`
}

type BatchDeleteInstanceArgs struct {
	InstanceIds        []string `json:"instanceIds,omitempty"`
	RelatedReleaseFlag bool     `json:"relatedReleaseFlag,omitempty"`
}

type InstanceDetailModel struct {
	InstanceModel
	Volume         *Volume                `json:"volume,omitempty"`
	Containers     []ContainerDetailModel `json:"containers,omitempty"`
	InitContainers []ContainerDetailModel `json:"initContainers,omitempty"`
	SecurityGroups []SecurityGroupModel   `json:"securityGroups,omitempty"`
	Vpc            *VpcModel              `json:"vpc,omitempty"`
	Subnet         *SubnetModel           `json:"subnet,omitempty"`
}

type ContainerDetailModel struct {
	Name            string           `json:"name,omitempty"`
	Image           string           `json:"image,omitempty"`
	CPU             float64          `json:"cpu,omitempty"`
	GPU             float64          `json:"gpu,omitempty"`
	Memory          float64          `json:"memory,omitempty"`
	WorkingDir      string           `json:"workingDir"`
	ImagePullPolicy string           `json:"imagePullPolicy,omitempty"`
	Commands        []string         `json:"commands,omitempty"`
	Args            []string         `json:"args,omitempty"`
	Ports           []Port           `json:"ports,omitempty"`
	VolumeMounts    []VolumeMount    `json:"volumeMounts,omitempty"`
	Envs            []Environment    `json:"envs,omitempty"`
	CreateTime      string           `json:"createTime,omitempty"`
	UpdateTime      string           `json:"updateTime,omitempty"`
	DeleteTime      string           `json:"deleteTime,omitempty"`
	PreviousState   *ContainerStatus `json:"previousState,omitempty"`
	CurrentState    *ContainerStatus `json:"currentState,omitempty"`
	Ready           bool             `json:"ready,omitempty"`
	RestartCount    int              `json:"restartCount,omitempty"`
}

type ContainerStatus struct {
	State        string `json:"state,omitempty"`
	Reason       string `json:"reason,omitempty"`
	Message      string `json:"message,omitempty"`
	StartTime    string `json:"startTime,omitempty"`
	FinishTime   string `json:"finishTime,omitempty"`
	DetailStatus string `json:"detailStatus,omitempty"`
	ExitCode     int    `json:"exitCode,omitempty"`
}

type SecurityGroupModel struct {
	SecurityGroupId string `json:"securityGroupId,omitempty"`
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	VpcId           string `json:"vpcId,omitempty"`
}

type VpcModel struct {
	VpcId       string `json:"vpcId,omitempty"`
	Name        string `json:"name,omitempty"`
	Cidr        string `json:"cidr,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault   bool   `json:"isDefault,omitempty"`
}

type SubnetModel struct {
	SubnetId    string `json:"subnetId,omitempty"`
	Name        string `json:"name,omitempty"`
	Cidr        string `json:"cidr,omitempty"`
	VpcId       string `json:"vpcId,omitempty"`
	SubnetType  string `json:"subnetType,omitempty"`
	Description string `json:"description,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
}

type InstanceModel struct {
	InstanceId      string `json:"instanceId,omitempty"`
	InstanceName    string `json:"instanceName,omitempty"`
	Status          string `json:"status,omitempty"`
	ZoneName        string `json:"zoneName,omitempty"`
	CPUType         string `json:"cpuType,omitempty"`
	GPUType         string `json:"gpuType,omitempty"`
	BandwidthInMbps int    `json:"bandwidthInMbps,omitempty"`
	InternalIp      string `json:"internalIp,omitempty"`
	PublicIp        string `json:"publicIp,omitempty"`
	CreateTime      string `json:"createTime,omitempty"`
	UpdateTime      string `json:"updateTime,omitempty"`
	DeleteTime      string `json:"deleteTime,omitempty"`
	RestartPolicy   string `json:"restartPolicy,omitempty"`
	Tags            []Tag  `json:"tags,omitempty"`
}

type Tag struct {
	TagKey   string `json:"tagKey,omitempty"`
	TagValue string `json:"tagValue,omitempty"`
}

type ImageRegistryCredential struct {
	Server   string `json:"server,omitempty"`
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

type Container struct {
	Name            string                    `json:"name"`
	Image           string                    `json:"image"`
	Memory          float64                   `json:"memory,omitempty"`
	CPU             float64                   `json:"cpu,omitempty"`
	GPU             float64                   `json:"gpu,omitempty"`
	WorkingDir      string                    `json:"workingDir"`
	ImagePullPolicy string                    `json:"imagePullPolicy,omitempty"`
	Commands        []string                  `json:"commands,omitempty"`
	Args            []string                  `json:"args,omitempty"`
	VolumeMounts    []VolumeMount             `json:"volumeMounts,omitempty"`
	Ports           []Port                    `json:"ports,omitempty"`
	EnvironmentVars []Environment             `json:"environmentVars,omitempty"`
	LivenessProbe   *Probe                    `json:"livenessProbe,omitempty"`
	ReadinessProbe  *Probe                    `json:"readinessProbe,omitempty"`
	StartupProbe    *Probe                    `json:"startupProbe,omitempty"`
	Stdin           bool                      `json:"stdin,omitempty"`
	StdinOnce       bool                      `json:"stdinOnce,omitempty"`
	Tty             bool                      `json:"tty,omitempty"`
	SecurityContext *ContainerSecurityContext `json:"securityContext,omitempty"`
}

type VolumeMount struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	MountPath string `json:"mountPath,omitempty"`
	ReadOnly  bool   `json:"readOnly,omitiempty"`
}

type Port struct {
	Name     string `json:"name,omitempty"`
	Port     int    `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}

type Environment struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Probe struct {
	InitialDelaySeconds           int              `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds                int              `json:"timeoutSeconds,omitempty"`
	PeriodSeconds                 int              `json:"periodSeconds,omitempty"`
	SuccessThreshold              int              `json:"successThreshold,omitempty"`
	FailureThreshold              int              `json:"failureThreshold,omitempty"`
	TerminationGracePeriodSeconds int64            `json:"terminationGracePeriodSeconds,omitempty"`
	HTTPGet                       *HTTPGetAction   `json:"httpGet,omitempty"`
	TCPSocket                     *TCPSocketAction `json:"tcpSocket,omitempty"`
	Exec                          *ExecAction      `json:"exec,omitempty"`
	GRPC                          *GRPCAction      `json:"grpc,omitempty"`
}

type HTTPGetAction struct {
	Path        string      `json:"path,omitempty"`
	Port        int         `json:"port,omitempty"`
	Scheme      string      `json:"scheme,omitempty"`
	Host        string      `json:"host,omitempty"`
	HTTPHeaders *HTTPHeader `json:"httpHeaders,omitempty"`
}

type HTTPHeader struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type TCPSocketAction struct {
	Port int    `json:"port,omitempty"`
	Host string `json:"host,omitempty"`
}

type ExecAction struct {
	Command []string `json:"command,omitempty"`
}

type GRPCAction struct {
	Port    int    `json:"port,omitempty"`
	Service string `josn:"service,omitempty"`
}

type ContainerSecurityContext struct {
	Capabilities           *Capabilities `json:"capabilities,omitempty"`
	RunAsUser              int64         `json:"runAsUser,omitempty"`
	RunAsGroup             int64         `json:"runAsGroup,omitempty"`
	RunAsNonRoot           bool          `json:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem bool          `json:"readOnlyRootFilesystem,omitempty"`
}

type Capabilities struct {
	Add  []string `json:"add,omitempty"`
	Drop []string `json:"drop,omitempty"`
}

type Volume struct {
	Nfs        []NfsVolume        `json:"nfs,omitempty"`
	EmptyDir   []EmptyDirVolume   `json:"emptyDir,omitempty"`
	ConfigFile []ConfigFileVolume `json:"configFile,omitempty"`
}

type NfsVolume struct {
	Name     string `json:"name,omitempty"`
	Server   string `json:"server,omitempty"`
	Path     string `json:"path,omitempty"`
	ReadOnly bool   `json:"readOnly,omitempty"`
}

type EmptyDirVolume struct {
	Name      string  `json:"name,omitempty"`
	Medium    string  `json:"medium,omitempty"`
	SizeLimit float64 `json:"sizeLimit,omitempty"`
}

type ConfigFileVolume struct {
	Name        string             `json:"name,omitempty"`
	DefaultMode int                `json:"defaultMode,omitempty"`
	ConfigFiles []ConfigFileDetail `json:"configFiles,omitempty"`
}

type ConfigFileDetail struct {
	Path string `json:"path,omitempty"`
	File string `json:"file,omitempty"`
}

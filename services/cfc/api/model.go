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

package api

import (
	"time"
)

type InvocationType string
type LogType string
type SourceType string
type TriggerType string

const (
	InvocationTypeEvent           InvocationType = "Event"
	InvocationTypeRequestResponse InvocationType = "RequestResponse"
	InvocationTypeDryRun          InvocationType = "DryRun"
	// 同步调用函数最大超时时间
	DefaultMaxFunctionTimeout = 300

	LogTypeTail LogType = "Tail"
	LogTypeNone LogType = "None"

	SourceTypeDuerOS  SourceType = "dueros"
	SourceTypeDuEdge  SourceType = "duedge"
	SourceTypeHTTP    SourceType = "cfc-http-trigger/v1/CFCAPI"
	SourceTypeCrontab SourceType = "cfc-crontab-trigger/v1/"
	SourceTypeCDN     SourceType = "cdn"

	TriggerTypeHTTP             TriggerType = "cfc-http-trigger"
	TriggerTypeGeneric          TriggerType = "generic"
	TypeEventSourceDatahubTopic             = "datahub_topic"
	TypeEventSourceBms                      = "bms"

	StartingPositionTriHorizon  = "TRIM_HORIZON"
	StartingPositionLatest      = "LATEST"
	StartingPositionAtTimeStamp = "AT_TIMESTAMP"

	DatahubTopicStartPointLatest = int64(-1)
	DatahubTopicStartPointOldest = int64(-2)
)

type Function struct {
	Uid                string             `json:"Uid"`
	Description        string             `json:"Description"`
	FunctionBrn        string             `json:"FunctionBrn"`
	Region             string             `json:"Region"`
	Timeout            int                `json:"Timeout"`
	VersionDesc        string             `json:"VersionDesc"`
	UpdatedAt          time.Time          `json:"UpdatedAt"`
	LastModified       time.Time          `json:"LastModified"`
	BlueprintTag       string             `json:"BlueprintTag"`
	CodeSha256         string             `json:"CodeSha256"`
	CodeSize           int32              `json:"CodeSize"`
	FunctionArn        string             `json:"FunctionArn"`
	FunctionName       string             `json:"FunctionName"`
	ServiceName        string             `json:"ServiceName"`
	Handler            string             `json:"Handler"`
	Version            string             `json:"Version"`
	Runtime            string             `json:"Runtime"`
	MemorySize         int                `json:"MemorySize"`
	Environment        *Environment       `json:"Environment"`
	CommitID           string             `json:"CommitID"`
	CodeID             string             `json:"CodeID"`
	Role               string             `json:"Role"`
	VpcConfig          *VpcConfig         `json:"VpcConfig"`
	LogType            string             `json:"LogType"`
	LogBosDir          string             `json:"LogBosDir"`
	BlsLogSet          string             `json:"BlsLogSet"`
	SourceTag          string             `json:"SourceTag"`
	DeadLetterTopic    string             `json:"DeadLetterTopic"`
	LayerList          []*LayerSample     `json:"Layers"`
	PodConcurrentQuota int                `json:"PodConcurrentQuota"`
	AsyncInvokeConfig  *AsyncInvokeConfig `json:"AsyncInvokeConfig"`
	CFSConfig          *CFSConfig         `json:"CFSConfig"`
}

type LayerSample struct {
	Brn         string                     `json:"Brn,omitempty"`
	CodeSize    int64                      `json:"CodeSize,omitempty"`
	Description string                     `json:"Description,omitempty"`
	Version     int64                      `json:"Version,omitempty"`
	LayerName   string                     `json:"LayerName,omitempty"`
	Content     *LayerVersionContentOutput `json:"Content,omitempty"`
}

type LayerVersionContentOutput struct {
	CodeSha256 string `json:"CodeSha256,omitempty"`
	CodeSize   int64  `json:"CodeSize,omitempty"`
	Location   string `json:"Location,omitempty"`
}

type AsyncInvokeConfig struct {
	MaxRetryIntervalInSeconds *int64             `json:"MaxRetryIntervalInSeconds"` // 消息最大保留时间
	MaxRetryAttempts          *int               `json:"MaxRetryAttempts"`          // 最大失败重试次数
	OnSuccess                 *DestinationConfig `json:"OnSuccess"`                 // 异步调用成功触发目标服务
	OnFailure                 *DestinationConfig `json:"OnFailure"`                 // 异步调用失败触发目标服务
}

type DestinationConfig struct {
	Type        string `json:"Type"`        // 触发目标服务类型，kafka of cfc
	Destination string `json:"Destination"` // 目标服务，topic or 函数brn
}
type CFSConfig struct {
	FsName     *string `json:"FsName"`     // 文件系统名称
	FsId       *string `json:"FsId"`       // 文件系统id
	SubnetID   *string `json:"SubnetID"`   // 子网ID
	Domain     *string `json:"Domain"`     // 挂载域名
	RemotePath *string `json:"RemotePath"` // CFS侧挂载路径
	LocalPath  *string `json:"LocalPath"`  // 本地目标路径
	Ovip       *string `json:"Ovip"`       // domain对应的外部虚拟ip
	VpcId      *string `json:"VpcId"`      // VPC Id
}
type Destination struct {
	Destination string `json:"destination,omitempty"`
}

// functionInfo
type FunctionInfo struct {
	Code          *CodeStorage `json:"Code"`
	Publish       bool         `json:"Publish"`
	Configuration *Function    `json:"Configuration"`
}

type Alias struct {
	AliasBrn        string    `json:"AliasBrn"`
	AliasArn        string    `json:"AliasArn"`
	FunctionName    string    `json:"FunctionName"`
	FunctionVersion string    `json:"FunctionVersion"`
	Name            string    `json:"Name"`
	Description     string    `json:"Description"`
	Uid             string    `json:"Uid"`
	UpdatedAt       time.Time `json:"UpdatedAt"`
	CreatedAt       time.Time `json:"CreatedAt"`
}

type RelationInfo struct {
	RelationId string      `json:"RelationId"`
	Sid        string      `json:"Sid"`
	Source     SourceType  `json:"Source"`
	Target     string      `json:"Target"`
	Data       interface{} `json:"Data"`
}

type CodeStorage struct {
	Location       string `json:"Location"`
	RepositoryType string `json:"RepositoryType"`
}

type Environment struct {
	Variables map[string]string
}

type CodeFile struct {
	Publish   bool
	DryRun    bool
	ZipFile   []byte
	BosBucket string
	BosObject string
}

type VpcConfig struct {
	VpcId            string
	SubnetIds        []string
	SecurityGroupIds []string
}

type InvocationsArgs struct {
	FunctionName   string
	InvocationType InvocationType
	LogType        LogType
	Qualifier      string
	Payload        interface{}
	RequestId      string
}

type InvocationsResult struct {
	Payload       string
	FunctionError string
	LogResult     string
}

type GetFunctionArgs struct {
	FunctionName string
	Qualifier    string
}

type GetFunctionResult struct {
	Code          CodeStorage
	Configuration Function
}

type ListFunctionsArgs struct {
	FunctionVersion string
	Marker          int
	MaxItems        int
}

type ListFunctionsResult struct {
	Functions  []*Function
	NextMarker string
}

type CreateFunctionArgs struct {
	Code               *CodeFile
	Publish            bool
	FunctionName       string
	Handler            string
	Runtime            string
	MemorySize         int
	Timeout            int
	Description        string
	Environment        *Environment
	VpcConfig          *VpcConfig
	LogType            string
	LogBosDir          string
	BlsLogSet          string
	ServiceName        string
	Region             string
	Version            string
	VersionDesc        string
	DeadLetterTopic    string
	LayerList          []*LayerSample `json:"Layers,omitempty"`
	PodConcurrentQuota int
	AsyncInvokeConfig  *AsyncInvokeConfig
	CFSConfig          *CFSConfig
	SourceTag          string
}

type CreateFunctionByBlueprintArgs struct {
	BlueprintID  string
	ServiceName  string
	FunctionName string
	Environment  *Environment
}

type CreateFunctionResult Function

type DeleteFunctionArgs struct {
	FunctionName string
	Qualifier    string
}

type GetFunctionConfigurationArgs struct {
	FunctionName string
	Qualifier    string
}

type GetFunctionConfigurationResult Function

type UpdateFunctionConfigurationArgs struct {
	FunctionName string
	Timeout      int `json:"Timeout,omitempty"`
	MemorySize   int `json:"MemorySize,omitempty"`
	Description  string
	Handler      string
	Runtime      string
	Environment  *Environment
	VpcConfig    *VpcConfig
	LogType      string
	LogBosDir    string
}

type UpdateFunctionConfigurationResult Function

type UpdateFunctionCodeArgs struct {
	FunctionName string
	ZipFile      []byte
	Publish      bool
	DryRun       bool
	BosBucket    string
	BosObject    string
}

type UpdateFunctionCodeResult Function

type ReservedConcurrentExecutionsArgs struct {
	FunctionName                 string
	ReservedConcurrentExecutions int
}

type DeleteReservedConcurrentExecutionsArgs struct {
	FunctionName string
}

type ListVersionsByFunctionArgs struct {
	FunctionName string
	Marker       int
	MaxItems     int
}
type ListVersionsByFunctionResult struct {
	Versions   []*Function
	NextMarker string
}

type PublishVersionArgs struct {
	FunctionName string
	Description  string
	CodeSha256   string
}

type PublishVersionResult Function

type ListAliasesArgs struct {
	FunctionName    string
	FunctionVersion string
	Marker          int
	MaxItems        int
}

type ListAliasesResult struct {
	Aliases    []*Alias
	NextMarker string
}

type GetAliasArgs struct {
	FunctionName string
	AliasName    string
}

type GetAliasResult Alias

type CreateAliasArgs struct {
	FunctionName    string
	FunctionVersion string
	Name            string
	Description     string
}

type CreateAliasResult Alias

type UpdateAliasArgs struct {
	FunctionName    string
	AliasName       string
	FunctionVersion string
	Description     string
}
type UpdateAliasResult Alias

type DeleteAliasArgs struct {
	FunctionName string
	AliasName    string
}

type ListTriggersArgs struct {
	FunctionBrn string
	ScopeType   string
}

type ListTriggersResult struct {
	Relation []*RelationInfo
}

type BosEventType string

const (
	BosEventTypePutObject               BosEventType = "PutObject"
	BosEventTypePostObject              BosEventType = "PostObject"
	BosEventTypeAppendObject            BosEventType = "AppendObject"
	BosEventTypeCopyObject              BosEventType = "CopyObject"
	BosEventTypeCompleteMultipartObject BosEventType = "CompleteMultipartUpload"
)

type BosTriggerData struct {
	Resource  string
	Status    string
	EventType []BosEventType
	Name      string
}

type HttpTriggerData struct {
	ResourcePath string
	Method       string
	AuthType     string
}

type CDNEventType string

const (
	CDNEventTypeCachedObjectsBlocked   CDNEventType = "CachedObjectsBlocked"
	CDNEventTypeCachedObjectsPushed    CDNEventType = "CachedObjectsPushed"
	CDNEventTypeCachedObjectsRefreshed CDNEventType = "CachedObjectsRefreshed"
	CDNEventTypeCdnDomainCreated       CDNEventType = "CdnDomainCreated"
	CDNEventTypeCdnDomainDeleted       CDNEventType = "CdnDomainDeleted"
	CDNEventTypeLogFileCreated         CDNEventType = "LogFileCreated"
	CDNEventTypeCdnDomainStarted       CDNEventType = "CdnDomainStarted"
	CDNEventTypeCdnDomainStopped       CDNEventType = "CdnDomainStopped"
)

type CDNTriggerData struct {
	EventType CDNEventType
	Domains   []string
	Remark    string
	Status    string
}

type CrontabTriggerData struct {
	Brn                string
	Enabled            string
	Input              string
	Name               string
	ScheduleExpression string
}

type CFCEdgeTriggerData struct {
	Domain    string
	EventType string
	Path      string
}

type CreateTriggerArgs struct {
	Target string
	Source SourceType
	Data   interface{}
}
type CreateTriggerResult struct {
	Relation *RelationInfo
}

type UpdateTriggerArgs struct {
	RelationId string
	Target     string
	Source     SourceType
	Data       interface{}
}
type UpdateTriggerResult struct {
	Relation *RelationInfo
}

type DeleteTriggerArgs struct {
	RelationId string
	Target     string
	Source     SourceType
}

type CreateEventSourceArgs FuncEventSource

type CreateEventSourceResult FuncEventSource

type DeleteEventSourceArgs struct {
	UUID string
}
type UpdateEventSourceArgs struct {
	UUID            string
	FuncEventSource FuncEventSource
}

type UpdateEventSourceResult FuncEventSource

type GetEventSourceArgs struct {
	UUID string
}

type GetEventSourceResult FuncEventSource

type ListEventSourceArgs struct {
	FunctionName string
	Marker       int
	MaxItems     int
}

type ListEventSourceResult struct {
	EventSourceMappings []FuncEventSource
}

// EventSource
type FuncEventSource struct {
	Uuid                      string     `json:"UUID"`
	BatchSize                 int        //  一次最多消费多少条消息
	Enabled                   *bool      `json:"Enabled,omitempty"` //是否开启消息触发器
	FunctionBrn               string     //  绑定的function brn
	EventSourceBrn            string     //  百度消息触发器bms kafka的topic名；Datahub触发器的配置唯一标识符，无需用户传入，服务端自动生成
	FunctionArn               string     //  兼容aws,与FunctionBrn相同
	EventSourceArn            string     //  兼容aws,与EventSourceBrn相同
	Type                      string     `json:"Type,omitempty"`                      // 类型 bms/datahub_topic
	FunctionName              string     `json:"FunctionName,omitempty"`              // 函数brn或者函数名
	StartingPosition          string     `json:"StartingPosition,omitempty"`          // 百度消息触发器bms kalfka topic 起始位置
	StartingPositionTimestamp *time.Time `json:"StartingPositionTimestamp,omitempty"` // 百度消息触发器bms kalfka topic 起始时间
	StateTransitionReason     string     // 状态变更原因
	DatahubConfig                        // Datahub触发器相关配置

	State                string    `json:"State"`                  // 消息触发器状态，开启或关闭，与aws兼容
	LastProcessingResult string    `json:"LastProcessingResult"`   // 最新一次触发器的执行结果
	LastModified         time.Time `json:"LastModified,omitempty"` // 上次修改时间
}

type DatahubConfig struct {
	MetaHostEndpoint string `json:"MetaHostEndpoint,omitempty"` // MetaHost endpoint
	MetaHostPort     int    `json:"MetaHostPort,omitempty"`     // MetaHost port
	ClusterName      string `json:"ClusterName,omitempty"`      // 集群名
	PipeName         string `json:"PipeName,omitempty"`         // pipe名
	PipeletNum       uint32 `json:"PipeletNum,omitempty"`       // 订阅PipiletNum
	StartPoint       int64  `json:"StartPoint,omitempty"`       // 起始订阅点  正常情况下id为正整数, 2个特殊的点 -1: 表示pipelet内的最新一条消息；-2: 表示pipelet内最旧的一条消息
	AclName          string `json:"ACLName,omitempty"`          // ACL name
	AclPassword      string `json:"ACLPassword,omitempty"`      // ACL passwd
}

type StartExecutionArgs struct {
	FlowName      string `json:"flowName,omitempty"`      // flowName
	ExecutionName string `json:"executionName,omitempty"` // executionName
	Input         string `json:"input,omitempty"`         // input
}

type Execution struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	FlowName       string `json:"flowName"`
	FlowDefinition string `json:"flowDefinition"`
	Input          string `json:"input"`
	Output         string `json:"output"`
	StartedTime    int64  `json:"startedTime"`
	StoppedTime    int64  `json:"stoppedTime"`
}

type ListExecutionsResult struct {
	Total      int         `json:"total"`
	Executions []Execution `json:"executions"`
}

type GetWorkflowResult struct {
	Code          CodeStorage
	Configuration Function
}

type GetExecutionHistoryResult struct {
	Events []*ExecutionEvent `json:"events"`
}

type ExecutionEvent struct {
	CostTime    int64  `json:"costTime"`
	EventDetail string `json:"eventDetail"`
	EventId     string `json:"eventId"`
	ExecutionId string `json:"executionId"`
	StateName   string `json:"stateName"`
	Time        int64  `json:"time"`
	Type        string `json:"type"`
}

type Flow struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Definition  string `json:"definition"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ListFlowResult struct {
	Flows []*struct {
		Name             string `json:"name"`
		Type             string `json:"type"`
		LastModifiedTime string `json:"lastModifiedTime"`
		Description      string `json:"description"`
		Definition       string `json:"definition"`
		CreatedTime      string `json:"createdTime"`
	} `json:"flows"`
}

type CreateUpdateFlowArgs struct {
	Name        string `json:"name"`
	Type        string `json:"type"` // 当前只支持 "FDL"
	Definition  string `json:"definition"`
	Description string `json:"description"`
}

type GetExecutionHistoryArgs struct {
	FlowName      string `json:"name"`
	ExecutionName string `json:"type"`
	Limit         int    `json:"limit"`
}

type PublishLayerVersionInput struct {
	CompatibleRuntimes []string                  `json:"CompatibleRuntimes,omitempty"`
	Content            *LayerVersionContentInput `json:"Content" valid:"required"`
	Description        string                    `json:"Description,omitempty" valid:"optional,runelength(0|256)"`
	LayerName          string                    `json:"LayerName" valid:"required,matches(^[a-zA-Z0-9-_]+$),runelength(0|140)"`
	LicenseInfo        string                    `json:"LicenseInfo,omitempty" valid:"optional,runelength(0|512)"`
	SourceTag          string                    `json:"SourceTag,omitempty" valid:"optional,runelength(0|128)"`
	Version            int64                     `json:"Version,omitempty" valid:"optional"`
}

type LayerVersionContentInput struct {
	BosBucket string `json:"BosBucket,omitempty"`
	BosObject string `json:"BosObject,omitempty"`
	ZipFile   string `json:"ZipFile,omitempty"`
}

type GetLayerVersionOutput struct {
	CompatibleRuntimes []string                   `json:"CompatibleRuntimes,omitempty"`
	Content            *LayerVersionContentOutput `json:"Content,omitempty"`
	CreatedDate        string                     `json:"CreatedDate,omitempty"`
	Description        string                     `json:"Description,omitempty"`
	LayerBrn           string                     `json:"LayerBrn,omitempty"`
	LayerVersionBrn    string                     `json:"LayerVersionBrn,omitempty"`
	LicenseInfo        string                     `json:"LicenseInfo,omitempty"`
	Version            int64                      `json:"Version,omitempty"`
}

type PublishLayerVersionOutput struct {
	CompatibleRuntimes []string                   `json:"CompatibleRuntimes,omitempty"`
	Content            *LayerVersionContentOutput `json:"Content,omitempty"`
	CreatedDate        string                     `json:"CreatedDate,omitempty"`
	Description        string                     `json:"Description,omitempty"`
	LayerBrn           string                     `json:"LayerBrn,omitempty"`
	LayerVersionBrn    string                     `json:"LayerVersionBrn,omitempty"`
	LicenseInfo        string                     `json:"LicenseInfo,omitempty"`
	Version            int64                      `json:"Version,omitempty"`
}

type GetLayerVersionArgs struct {
	find          string
	Brn           string
	LayerName     string
	VersionNumber string
}

type ListLayerVersionsInput struct {
	CompatibleRuntime string `json:"CompatibleRuntime,omitempty"`
	LayerName         string `json:"LayerName" valid:"required"`
	*ListCondition
}

type ListLayerVersionsOutput struct {
	LayerVersions []*LayerVersionsListItem
	NextMarker    string
	Total         int64
	PageNo        int64 `json:"pageNo"`
	PageSize      int64 `json:"pageSize"`
}
type ListCondition struct {
	PageNo   int64 `json:"PageNo,omitempty"`
	PageSize int64 `json:"PageSize,omitempty"`
	Marker   int64 `json:"Marker,omitempty"`
	MaxItems int64 `json:"MaxItems,omitempty"`
}

type ListLayerInput struct {
	CompatibleRuntime string `json:"CompatibleRuntime,omitempty"`
	*ListCondition
}

type ListLayersOutput struct {
	Layers     []*LayersListItem `json:"Layers,omitempty"`
	NextMarker string            `json:"NextMarker,omitempty"`
	Total      int64             `json:"Total,omitempty"`
	PageNo     int64             `json:"PageNo"`
	PageSize   int64             `json:"PageSize"`
}

type LayersListItem struct {
	LatestMatchingVersion *LayerVersionsListItem `json:"LatestMatchingVersion,omitempty"`
	LayerBrn              string                 `json:"LayerBrn,omitempty"`
	LayerName             string                 `json:"LayerName,omitempty"`
}

type LayerVersionsListItem struct {
	CompatibleRuntimes []string `json:"CompatibleRuntimes,omitempty"`
	CreatedDate        string   `json:"CreatedDate,omitempty"`
	Description        string   `json:"Description,omitempty"`
	LayerVersionBrn    string   `json:"LayerVersionBrn,omitempty"`
	LicenseInfo        string   `json:"LicenseInfo,omitempty"`
	Version            int64    `json:"Version,omitempty"`
}

type DeleteLayerVersionArgs struct {
	LayerName     string
	VersionNumber string
}

type DeleteLayerArgs struct {
	LayerName string
}

type Service struct {
	Uid           string                 `json:"Uid,omitempty"`
	ServiceName   string                 `json:"ServiceName" valid:"optional,matches(^[a-zA-Z0-9-_]+$),runelength(1|50)"`
	ServiceDesc   *string                `json:"ServiceDesc,omitempty"`
	ServiceConf   string                 `json:"ServiceConf,omitempty"`
	ServiceConfig map[string]interface{} `json:"ServiceConfig,omitempty"`
	Region        string                 `json:"Region,omitempty"`
	Status        int                    `json:"Status,omitempty"`
	UpdatedAt     time.Time              `json:"UpdatedAt,omitempty"`
	CreatedAt     time.Time              `json:"CreatedAt,omitempty"`
}

// ServiceWithFun represents a service with function count information
type ServiceWithFun struct {
	Service
	FuncCount int `json:"FuncCount,omitempty"`
}

type GetServiceArgs struct {
	ServiceName string `json:"ServiceName"`
}

type GetServiceResult Service

type CreateServiceArgs struct {
	ServiceName string  `json:"ServiceName"`
	ServiceDesc *string `json:"ServiceDesc,omitempty"`
	ServiceConf string  `json:"ServiceConf,omitempty"`
	Region      string  `json:"Region,omitempty"`
}

type CreateServiceResult Service

type UpdateServiceArgs struct {
	ServiceName string  `json:"ServiceName"`
	ServiceDesc *string `json:"ServiceDesc,omitempty"`
	ServiceConf string  `json:"ServiceConf,omitempty"`
	Region      string  `json:"Region,omitempty"`
}

type UpdateServiceResult Service

type DeleteServiceArgs struct {
	ServiceName string `json:"ServiceName"`
}

type DeleteServiceResult Service

type ListServicesResult struct {
	Services []*Service `json:"Services"`
}

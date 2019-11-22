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

	LogTypeTail LogType = "Tail"
	LogTypeNone LogType = "None"

	SourceTypeDuerOS  SourceType = "dueros"
	SourceTypeDuEdge  SourceType = "duedge"
	SourceTypeHTTP    SourceType = "cfc-http-trigger/v1/CFCAPI"
	SourceTypeCrontab SourceType = "cfc-crontab-trigger/v1/"
	SourceTypeCDN     SourceType = "cdn"

	TriggerTypeHTTP    TriggerType = "cfc-http-trigger"
	TriggerTypeGeneric TriggerType = "generic"
)

type Function struct {
	Uid          string       `json:"Uid"`
	Description  string       `json:"Description"`
	FunctionBrn  string       `json:"FunctionBrn"`
	Region       string       `json:"Region"`
	Timeout      int          `json:"Timeout"`
	VersionDesc  string       `json:"VersionDesc"`
	UpdatedAt    time.Time    `json:"UpdatedAt"`
	LastModified time.Time    `json:"LastModified"`
	CodeSha256   string       `json:"CodeSha256"`
	CodeSize     int32        `json:"CodeSize"`
	FunctionArn  string       `json:"FunctionArn"`
	FunctionName string       `json:"FunctionName"`
	Handler      string       `json:"Handler"`
	Version      string       `json:"Version"`
	Runtime      string       `json:"Runtime"`
	MemorySize   int          `json:"MemorySize"`
	Environment  *Environment `json:"Environment"`
	CommitID     string       `json:"CommitID"`
	CodeID       string       `json:"CodeID"`
	Role         string       `json:"Role"`
	VpcConfig    *VpcConfig   `json:"VpcConfig"`
	LogType      string       `json:"LogType"`
	LogBosDir    string       `json:"LogBosDir"`
	SourceTag    string       `json:"SourceTag"`
}

//functionInfo
type FunctionInfo struct {
	Code          *CodeStorage `json:"Code"`
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
	Code         *CodeFile
	FunctionName string
	Handler      string
	Runtime      string
	MemorySize   int
	Timeout      int
	Description  string
	Environment  *Environment
	VpcConfig    *VpcConfig
	LogType      string
	LogBosDir    string
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
	BosEventTypeCompleteMultipartObject BosEventType = "CompleteMultipartObject"
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

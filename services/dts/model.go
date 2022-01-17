/*
 * Copyright 2020 Baidu, Inc.
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

package dts

type CreateDtsArgs struct {
	ClientToken        string `json:"-"`
	ProductType        string `json:"productType"`
	Type               string `json:"type"`
	Standard           string `json:"standard"`
	SourceInstanceType string `json:"sourceInstanceType"`
	TargetInstanceType string `json:"targetInstanceType"`
	CrossRegionTag     int    `json:"crossRegionTag"`
	DirectionType      string `json:"directionType"`
}

type CreateDtsResult struct {
	DtsTasks []DtsId `json:"dtsTasks"`
}

type DtsId struct {
	DtsId string `json:"dtsId"`
}

type ConfigDtsResult struct {
	DtsId string `json:"dtsId"`
}

type DtsTaskMeta struct {
	DtsId               string       `json:"dtsId"`
	TaskName            string       `json:"taskName"`
	Status              string       `json:"status"`
	DataType            []string     `json:"dataType"`
	Region              string       `json:"region"`
	CreateTime          string       `json:"createTime"`
	SrcConnection       Connection   `json:"srcConnection"`
	DstConnection       Connection   `json:"dstConnection"`
	SchemaMapping       []Schema     `json:"schemaMapping,omitempty"`
	RunningTime         int          `json:"runningTime"`
	SubStatus           []SubStatus  `json:"subStatus,omitempty"`
	DynamicInfo         DynamicInfo  `json:"dynamicInfo,omitempty"`
	Errmsg              string       `json:"errmsg,omitempty"`
	SdkRealtimeProgress string       `json:"sdkRealtimeProgress,omitempty"`
	Granularity         string       `json:"granularity,omitempty"`
	SubDataScope        SubDataScope `json:"subDataScope,omitempty"`
	PayInfo             PayInfo      `json:"payInfo,omitempty"`
	LockStatus          string       `json:"lockStatus,omitempty"`
	DtsIdPos            string       `json:"dtsIdPos,omitempty"`
	DtsIdNeg            string       `json:"dtsIdNeg,omitempty"`
	DtsTaskPos          *DtsTaskMeta `json:"dtsTaskPos"`
	DtsTaskNeg          *DtsTaskMeta `json:"dtsTaskNeg"`
}

type Connection struct {
	Region          string `json:"region"`
	DbType          string `json:"dbType"`
	DbUser          string `json:"dbUser"`
	DbPass          string `json:"dbPass"`
	DbPort          int    `json:"dbPort"`
	DbHost          string `json:"dbHost"`
	InstanceId      string `json:"instanceId"`
	DbServer        string `json:"dbServer,omitempty"`
	InstanceType    string `json:"instanceType"`
	InstanceShortId string `json:"instanceShortId,omitempty"`
	FieldWhitelist  string `json:"field_whitelist,omitempty"`
	FieldBlacklist  string `json:"field_blacklist,omitempty"`
	StartTime       string `json:"startTime,omitempty"`
	EndTime         string `json:"endTime,omitempty"`
	SqlType         string `json:"sqlType,omitempty"`
}

type Schema struct {
	Type  string `json:"type"`
	Src   string `json:"src"`
	Dst   string `json:"dst"`
	Where string `json:"where"`
}

type SubStatus struct {
	S string `json:"s"`
	B string `json:"b"`
	I string `json:"i"`
}

type DynamicInfo struct {
	Schema    []SchemaInfo `json:"schema"`
	Base      []SchemaInfo `json:"base"`
	Increment Increment    `json:"increment"`
}

type Increment struct {
	Delay      int64  `json:"delay"`
	Position   string `json:"position"`
	SyncStatus string `json:"syncStatus"`
}

type SchemaInfo struct {
	Current          string `json:"current"`
	Count            string `json:"count"`
	Speed            string `json:"speed"`
	ExpectFinishTime string `json:"expectFinishTime"`
}

type SubDataScope struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type PayInfo struct {
	ProductType        string `json:"productType"`
	SourceInstanceType string `json:"sourceInstanceType"`
	TargetInstanceType string `json:"targetInstanceType"`
	CrossRegionTag     int    `json:"crossRegionTag"`
	CreateTime         int    `json:"createTime"`
	Standard           string `json:"standard"`
	EndTime            string `json:"endTime"`
}

type ListDtsArgs struct {
	Type        string `json:"type"`
	Status      string `json:"status,omitempty"`
	Marker      string `json:"marker,omitempty"`
	MaxKeys     int    `json:"maxKeys,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	KeywordType string `json:"keywordType,omitempty"`
}

type ListDtsWithPageArgs struct {
	Types    []string     `json:"types"`
	Filters  []ListFilter `json:"filters"`
	Order    string       `json:"order"`
	OrderBy  string       `json:"orderBy"`
	PageNo   int          `json:"pageNo"`
	PageSize int          `json:"pageSize"`
}

type ListFilter struct {
	KeywordType string `json:"keywordType"`
	Keyword     string `json:"keyword"`
}

type ListDtsResult struct {
	Marker      string        `json:"marker"`
	MaxKeys     int           `json:"maxKeys"`
	IsTruncated bool          `json:"isTruncated"`
	NextMarker  string        `json:"nextMarker"`
	Task        []DtsTaskMeta `json:"task"`
}

type ListDtsWithPageResult struct {
	OrderBy    string        `json:"orderBy"`
	Order      string        `json:"order"`
	PageNo     int           `json:"pageNo"`
	PageSize   int           `json:"pageSize"`
	TotalCount int           `json:"totalCount"`
	Result     []DtsTaskMeta `json:"result"`
}

type CheckResult struct {
	Name         string `json:"name"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	Subscription string `json:"subscription"`
}

type GetPreCheckResult struct {
	Success bool          `json:"success"`
	Result  []CheckResult `json:"result"`
}

type ConfigArgs struct {
	Type          string       `json:"type,omitempty"`
	DtsId         string       `json:"dtsId,omitempty"`
	TaskName      string       `json:"taskName"`
	DataType      []string     `json:"dataType"`
	SrcConnection Connection   `json:"srcConnection"`
	DstConnection Connection   `json:"dstConnection"`
	SchemaMapping []Schema     `json:"schemaMapping"`
	Granularity   string       `json:"granularity,omitempty"`
	ProductType   string       `json:"productType,omitempty"`
	QueueType     string       `json:"queueType,omitempty"`
	InitPosition  InitPosition `json:"initPosition,omitempty"`
	NetType       string       `json:"netType,omitempty"`
	Admin         string       `json:"admin,omitempty"`
}

type InitPosition struct {
	Type     string `json:"type"`
	Position string `json:"position"`
}

type PreCheckResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type SkipPreCheckResponse struct {
	Success bool   `json:"success"`
	Result  string `json:"result"`
}

type GetSchemaArgs struct {
	Connection Connection `json:"connection"`
}

type GetSchemaResponse struct {
	Success bool            `json:"success"`
	Result  GetSchemaResult `json:"result"`
}

type GetSchemaResult struct {
	SchemaAll map[string]ObjectInDb `json:"schemaAll"`
}

type ObjectInDb struct {
	Tables     []string `json:"tables"`
	Views      []string `json:"views"`
	Procedures []string `json:"procedures"`
	Functions  []string `json:"functions"`
}

type UpdateTaskNameArgs struct {
	TaskName string `json:"taskName"`
}

type ResizeTaskStandardArgs struct {
	ClientToken string `json:"-"`
	Standard    string `json:"standard"`
}

type ResizeTaskStandardResponse struct {
	OrderId string `json:"orderId"`
}

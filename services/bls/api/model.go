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

import (
	"github.com/baidubce/bce-sdk-go/model"
)

type DateTime string

type Project struct {
	CreatedTime DateTime `json:"createdTime"`
	UpdatedTime DateTime `json:"updatedTime"`
	UUID        string   `json:"uuid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Top         bool     `json:"top"`
}

type DescribeProjectResult struct {
	Project *Project `json:"project"`
}

type DescribeProjectResponse struct {
	Code    string                `json:"code"`
	Message string                `json:"message"`
	Result  DescribeProjectResult `json:"result"`
}

type ListProjectResult struct {
	Order          string    `json:"order"`
	OrderBy        string    `json:"orderBy"`
	PageNo         int       `json:"pageNo"`
	PageSize       int       `json:"pageSize"`
	TotalCount     int       `json:"totalCount"`
	DefaultProject Project   `json:"default"`
	Projects       []Project `json:"projects"`
}

type ListProjectResponse struct {
	Code    string             `json:"code"`
	Message string             `json:"message"`
	Result  *ListProjectResult `json:"result"`
}

type LogRecord struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Sequence  int    `json:"sequence"`
}

type LogStream struct {
	CreationDateTime DateTime `json:"creationDateTime"`
	LogStreamName    string   `json:"logStreamName"`
}

type LogStore struct {
	CreationDateTime DateTime         `json:"creationDateTime"`
	LastModifiedTime DateTime         `json:"lastModifiedTime"`
	Project          string           `json:"project"`
	LogStoreName     string           `json:"logStoreName"`
	Retention        int              `json:"retention"`
	Tags             []model.TagModel `json:"tags,omitempty"`
}

type LogShipper struct {
	Status         string             `json:"status"`
	LogShipperName string             `json:"logShipperName"`
	Project        string             `json:"project"`
	LogStoreName   string             `json:"logStoreName"`
	StartTime      string             `json:"startTime"`
	DestType       string             `json:"destType"`
	DestConfig     *ShipperDestConfig `json:"destConfig"`
}

type QueryConditions struct {
	NamePattern string `json:"namePattern"`
	Order       string `json:"order"`
	OrderBy     string `json:"orderBy"`
	PageNo      int    `json:"pageNo"`
	PageSize    int    `json:"pageSize"`
}

type PushLogRecordBody struct {
	LogStreamName string      `json:"logStreamName"`
	Type          string      `json:"type"`
	LogRecords    []LogRecord `json:"logRecords"`
}

type QueryLogRecordArgs struct {
	LogStreamName    string   `json:"logStreamName"`
	Query            string   `json:"query"`
	StartDateTime    DateTime `json:"startDatetime"`
	EndDateTime      DateTime `json:"endDateTime"`
	Limit            int      `json:"limit"`
	Marker           string   `json:"marker"`
	Sort             string   `json:"sort"`
	SamplePercentage float64  `json:"samplePercentage"`
	SampleSeed       int      `json:"sampleSeed"`
}

type PullLogRecordArgs struct {
	LogStreamName string   `json:"logStreamName"`
	StartDateTime DateTime `json:"startDatetime"`
	EndDateTime   DateTime `json:"endDateTime"`
	Limit         int      `json:"limit"`
	Marker        string   `json:"marker"`
}

type PullLogRecordResult struct {
	Result      []LogRecord `json:"result"`
	IsTruncated bool        `json:"isTruncated"`
	Marker      string      `json:"marker"`
	NextMarker  string      `json:"nextMarker"`
}

type Histogram struct {
	Interval      int      `json:"interval"`
	StartDateTime DateTime `json:"startDatetime"`
	EndDateTime   DateTime `json:"endDateTime"`
	Counts        []int    `json:"counts"`
}

type Statistics struct {
	ExecutionTimeInMs int        `json:"executionTimeInMs"`
	ScanCount         int        `json:"scanCount"`
	Histogram         *Histogram `json:"histogram"`
}

type DataSetScanInfo struct {
	IsTruncated     bool        `json:"isTruncated"`
	TruncatedReason string      `json:"truncatedReason"`
	Statistics      *Statistics `json:"statistics"`
}

type ResultSet struct {
	QueryType       string              `json:"queryType"`
	Columns         []string            `json:"columns"`
	ColumnTypes     []string            `json:"columnTypes"`
	Rows            [][]interface{}     `json:"rows"`
	Tags            []map[string]string `json:"tags,omitempty"`
	IsTruncated     bool                `json:"isTruncated"`
	TruncatedReason string              `json:"truncatedReason"`
}

type QueryLogResult struct {
	ResultSet       *ResultSet       `json:"resultSet"`
	DataSetScanInfo *DataSetScanInfo `json:"datasetScanInfo"`
	NextMarker      string           `json:"nextMarker"`
}

type ListLogStreamResult struct {
	Order      string      `json:"order"`
	OrderBy    string      `json:"orderBy"`
	PageNumber int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	Result     []LogStream `json:"result"`
}

type ListLogStoreResult struct {
	Order      string     `json:"order"`
	OrderBy    string     `json:"orderBy"`
	PageNo     int        `json:"pageNo"`
	PageSize   int        `json:"pageSize"`
	TotalCount int        `json:"totalCount"`
	Result     []LogStore `json:"result"`
}

type BatchLogStoreResult struct {
	Result  []LogStore `json:"result"`
	Code    string     `json:"code"`
	Success bool       `json:"success"`
}

type FastQuery struct {
	CreationDateTime DateTime `json:"creationDateTime"`
	LastModifiedTime DateTime `json:"lastModifiedTime"`
	FastQueryName    string   `json:"fastQueryName"`
	Description      string   `json:"description"`
	Query            string   `json:"query"`
	Project          string   `json:"project"`
	LogStoreName     string   `json:"logStoreName"`
	LogStreamName    string   `json:"logStreamName"`
	StartDateTime    DateTime `json:"startDateTime"`
	EndDateTime      DateTime `json:"endDateTime"`
}

type CreateFastQueryBody struct {
	FastQueryName string `json:"fastQueryName"`
	Query         string `json:"query"`
	Description   string `json:"description"`
	LogStoreName  string `json:"logStoreName"`
	LogStreamName string `json:"logStreamName"`
}

type UpdateFastQueryBody struct {
	Query         string `json:"query"`
	Description   string `json:"description"`
	LogStoreName  string `json:"logStoreName"`
	LogStreamName string `json:"logStreamName"`
}

type ListFastQueryResult struct {
	Order      string      `json:"order"`
	OrderBy    string      `json:"orderBy"`
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	Result     []FastQuery `json:"result"`
}

type LogField struct {
	Type           string              `json:"type"`
	CaseSensitive  bool                `json:"caseSensitive"`
	Separators     string              `json:"separators"`
	IncludeChinese bool                `json:"includeChinese"`
	Fields         map[string]LogField `json:"fields,omitempty"`
	DynamicMapping bool                `json:"dynamicMapping,omitempty"`
}

type IndexFields struct {
	FullText       bool                `json:"fulltext"`
	CaseSensitive  bool                `json:"caseSensitive"`
	Separators     string              `json:"separators"`
	IncludeChinese bool                `json:"includeChinese"`
	Fields         map[string]LogField `json:"fields"`
}

type CreateLogShipperBody struct {
	LogShipperName string             `json:"logShipperName"`
	LogStoreName   string             `json:"logStoreName"`
	StartTime      string             `json:"startTime"`
	DestType       string             `json:"destType"`
	DestConfig     *ShipperDestConfig `json:"destConfig"`
}

type CreateLogShipperResponse struct {
	LogShipperID string `json:"logShipperID"`
}

type ShipperDestConfig struct {
	BOSPath                  string `json:"BOSPath"`
	PartitionFormatTS        string `json:"partitionFormatTS"`
	PartitionFormatLogStream bool   `json:"partitionFormatLogStream"`
	MaxObjectSize            int64  `json:"maxObjectSize"`
	CompressType             string `json:"compressType"`
	DeliverInterval          int64  `json:"deliverInterval"`
	StorageFormat            string `json:"storageFormat"`
	ShipperType              string `json:"shipperType"`
	CsvHeadline              bool   `json:"csvHeadline"`
	CsvDelimiter             string `json:"csvDelimiter"`
	CsvQuote                 string `json:"csvQuote"`
	NullIdentifier           string `json:"nullIdentifier"`
	SelectedColumnName       string `json:"selectedColumnName"`
	SelectedColumnType       string `json:"selectedColumnType"`
}

type ListShipperRecordCondition struct {
	SinceHours int `json:"sinceHours"`
	PageNo     int `json:"pageNo"`
	PageSize   int `json:"pageSize"`
}

type ListShipperRecordResult struct {
	TotalCount int                `json:"totalCount"`
	Result     []LogShipperRecord `json:"result"`
}

type LogShipperRecord struct {
	StartTime     string `json:"startTime"`
	EndTime       string `json:"endTime"`
	FinishedCount int    `json:"finishedCount"`
}

type ListShipperResult struct {
	TotalCount int              `json:"totalCount"`
	Result     []ShipperSummary `json:"result"`
}

type ShipperSummary struct {
	LogShipperID   string `json:"logShipperID"`
	LogShipperName string `json:"logShipperName"`
	Project        string `json:"project"`
	LogStoreName   string `json:"logStoreName"`
	DestType       string `json:"destType"`
	Status         string `json:"status"`
	ErrMessage     string `json:"errMessage"`
	CreateDateTime string `json:"createDateTime"`
}

type ListLogShipperCondition struct {
	LogShipperID   string `json:"logShipperID"`
	LogShipperName string `json:"logShipperName"`
	Project        string `json:"project"`
	LogStoreName   string `json:"logStoreName"`
	DestType       string `json:"destType"`
	Status         string `json:"status"`
	Order          string `json:"order"`
	OrderBy        string `json:"orderBy"`
	PageNo         int    `json:"pageNo"`
	PageSize       int    `json:"pageSize"`
}

type UpdateLogShipperBody struct {
	LogShipperName string             `json:"logShipperName"`
	DestConfig     *ShipperDestConfig `json:"destConfig"`
}

type BulkDeleteShipperCondition struct {
	LogShipperIDs []string `json:"logShipperIDs"`
}

type SetSingleShipperStatusCondition struct {
	DesiredStatus string `json:"desiredStatus"`
}

type BulkSetShipperStatusCondition struct {
	LogShipperIDs []string `json:"logShipperIDs"`
	DesiredStatus string   `json:"desiredStatus"`
}

type DownloadTask struct {
	UUID           string `json:"uuid"`
	Name           string `json:"name"`
	Project        string `json:"project"`
	LogStoreName   string `json:"logStoreName"`
	LogStreamName  string `json:"logStreamName"`
	Query          string `json:"query"`
	QueryStartTime string `json:"queryStartTime"`
	QueryEndTime   string `json:"queryEndTime"`
	Format         string `json:"format"`
	Limit          int64  `json:"limit"`
	OrderBy        string `json:"orderBy"`
	Order          string `json:"order"`
	State          string `json:"state"`
	FailedCode     string `json:"failedCode"`
	FailedMessage  string `json:"failedMessage"`
	Retry          int    `json:"retry"`
	WrittenRows    int64  `json:"writtenRows"`
	FileDir        string `json:"fileDir"`
	FileName       string `json:"fileName"`
	ExecStartTime  string `json:"execStartTime"`
	ExecEndTime    string `json:"execEndTime"`
	CreatedTime    string `json:"createdTime"`
	UpdatedTime    string `json:"updatedTime"`
}

type CreateDownloadTaskResult struct {
	UUID string `json:"uuid"`
}

type CreateDownloadResponse struct {
	Code    string                   `json:"code"`
	Message string                   `json:"message"`
	Result  CreateDownloadTaskResult `json:"result"`
}

type DescribeDownloadTaskResult struct {
	Task *DownloadTask `json:"task"`
}

type DescribeDownloadTaskResponse struct {
	Code    string                     `json:"code"`
	Message string                     `json:"message"`
	Result  DescribeDownloadTaskResult `json:"result"`
}

type GetDownloadTaskLinkResult struct {
	FileDir  string `json:"fileDir"`
	FileName string `json:"fileName"`
	Link     string `json:"link"`
}

type GetDownloadTaskLinkResponse struct {
	Code    string                     `json:"code"`
	Message string                     `json:"message"`
	Result  *GetDownloadTaskLinkResult `json:"result"`
}

type ListDownloadTaskResult struct {
	Order    string         `json:"order"`
	OrderBy  string         `json:"orderBy"`
	PageNo   int            `json:"pageNo"`
	PageSize int            `json:"pageSize"`
	Total    int            `json:"total"`
	Tasks    []DownloadTask `json:"tasks"`
}

type ListDownloadTaskResponse struct {
	Code    string                  `json:"code"`
	Message string                  `json:"message"`
	Result  *ListDownloadTaskResult `json:"result"`
}

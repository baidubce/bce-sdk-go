package bls

import (
	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/bls/api"
)

const (
	DefaultProject = "default"
)

type CreateProjectRequest struct {
	// 日志组名称，选填，默认为default
	Name string `json:"name"`
	// 日志组描述，选填
	Description string `json:"description"`
}

type UpdateProjectRequest struct {
	// 日志组UUID 必填
	UUID string `json:"uuid"`
	// 日志组描述，选填
	Description string `json:"description"`
	// 日志组是否置顶， 选填
	Top bool `json:"top"`
}

type DescribeProjectRequest struct {
	// 日志组UUID 必填
	UUID string `json:"uuid"`
}

type DeleteProjectRequest struct {
	// 日志组UUID 必填
	UUID string `json:"uuid"`
}

type ListProjectRequest struct {
	// 日志组名称关键字查询，模糊匹配，选填
	Name string `json:"name"`
	// 日志组描述关键字查询，模糊匹配，选填
	Description string `json:"description"`
	// 排序方式 选填 支持desc和asc，默认desc
	Order string `json:"order"`
	// 排序字段 选填 支持 createdAt, updatedAt, name  默认按照createdAt创建时间排序
	OrderBy string `json:"orderBy"`
	// 第几页 选填， 默认为1
	PageNo int `json:"pageNo"`
	// 每页大小 选填， 默认为10
	PageSize int `json:"pageSize"`
}

type CreateLogStoreRequest struct {
	// 日志组名称，选填，默认为default
	Project string `json:"project"`
	// 日志集名称，必填
	LogStoreName string `json:"logStoreName"`
	// 日志集存储周期，必填
	Retention int `json:"retention"`
	// 日志集标签，选填，默认没有标签
	Tags []model.TagModel `json:"tags"`
}

type UpdateLogStoreRequest struct {
	// 日志组名称，选填，默认为default
	Project string `json:"project"`
	// 日志集名称，必填
	LogStoreName string `json:"logStoreName"`
	// 日志集存储周期，必填
	Retention int `json:"retention"`
}

type DescribeLogStoreRequest struct {
	// 日志组名称，选填，默认为default
	Project string `json:"project"`
	// 日志集名称，必填
	LogStoreName string `json:"logStoreName"`
}

type DeleteLogStoreRequest struct {
	// 日志组名称，选填，默认为default
	Project string `json:"project"`
	// 日志集名称，必填
	LogStoreName string `json:"logStoreName"`
}

type ListLogStoreRequest struct {
	// 日志组名称，选填，默认全部日志组中搜索日志集
	Project string `json:"project"`
	// 日志集名称关键字查询，模糊匹配，选填
	NamePattern string `json:"namePattern"`
	// 排序方式 选填 支持desc和asc，默认desc
	Order string `json:"order"`
	// 排序字段 选填 支持 creationDateTime, lastModifiedTime, name  默认按照creationDateTime创建时间排序
	OrderBy string `json:"orderBy"`
	// 第几页 选填， 默认为1
	PageNo int `json:"pageNo"`
	// 每页大小 选填， 默认为10
	PageSize int `json:"pageSize"`
}

type ListLogStreamRequest struct {
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
	// 日志流名称关键字查询，模糊匹配，选填
	NamePattern string `json:"namePattern"`
	// 排序方式 选填 支持desc和asc，默认desc
	Order string `json:"order"`
	// 排序字段 选填 取值为：logStreamName  默认按照logStreamName名称排序
	OrderBy string `json:"orderBy"`
	// 第几页 选填， 默认为1
	PageNo int `json:"pageNo"`
	// 每页大小 选填， 默认为10
	PageSize int `json:"pageSize"`
}

type PushLogRecordRequest struct {
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
	// 日志流名称 选填，默认在全部日志流中查询
	LogStreamName string `json:"logStreamName"`
	// 日志类型，选填，取值为TEXT, JSON  默认为TEXT格式
	LogType string `json:"logType"`
	// 日志数据，必填
	LogRecords []api.LogRecord `json:"logRecords"`
}

type PullLogRecordRequest struct {
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
	// 日志流名称 选填，默认在全部日志流中查询
	LogStreamName string `json:"logStreamName"`
	// 日志的开始时间， 必填，UTC时间，格式ISO8601，例如：2020-01-10T13:23:34Z
	StartDateTime string `json:"startDateTime"`
	// 日志的结束时间，必填，UTC时间，格式ISO8601，例如：2020-01-10T14:23:34Z
	EndDateTime string `json:"endDateTime"`
	// 返回的数据条数，选填，默认为100
	Limit int `json:"limit"`
	// 指定查看的开始位置标记，选填
	Marker string `json:"marker"`
}

type QueryLogRecordRequest struct {
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
	// 日志流名称 选填，默认在全部日志流中查询
	LogStreamName string `json:"logStreamName"`
	// 查询语句，必填
	Query string `json:"query"`
	// 日志的开始时间， 必填，UTC时间，格式ISO8601，例如：2020-01-10T13:23:34Z
	StartDateTime string `json:"startDateTime"`
	// 日志的结束时间，必填，UTC时间，格式ISO8601，例如：2020-01-10T14:23:34Z
	EndDateTime string `json:"endDateTime"`
	// 返回的数据条数，选填，默认为100
	Limit int `json:"limit"`
}

type CreateFastQueryRequest struct {
	// 快速查询名称，必填
	FastQueryName string `json:"fastQueryName"`
	// 查询语句，必填
	Query string `json:"query"`
	// 快速查询描述 选填
	Description string `json:"description"`
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
	// 日志流名称 选填，默认在全部日志流中查询
	LogStreamName string `json:"logStreamName"`
	// 日志的开始时间， 必填，UTC时间，格式ISO8601，例如：2020-01-10T13:23:34Z
	StartDateTime string `json:"startDateTime"`
	// 日志的结束时间，必填，UTC时间，格式ISO8601，例如：2020-01-10T14:23:34Z
	EndDateTime string `json:"endDateTime"`
}

type UpdateFastQueryRequest struct {
	// 快速查询名称，必填
	FastQueryName string `json:"fastQueryName"`
	// 查询语句，必填
	Query string `json:"query"`
	// 快速查询描述 选填
	Description string `json:"description"`
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
	// 日志流名称 选填，默认在全部日志流中查询
	LogStreamName string `json:"logStreamName"`
	// 日志的开始时间， 必填，UTC时间，格式ISO8601，例如：2020-01-10T13:23:34Z
	StartDateTime string `json:"startDateTime"`
	// 日志的结束时间，必填，UTC时间，格式ISO8601，例如：2020-01-10T14:23:34Z
	EndDateTime string `json:"endDateTime"`
}

type DescribeFastQueryRequest struct {
	// 快速查询名称，必填
	FastQueryName string `json:"fastQueryName"`
}

type DeleteFastQueryRequest struct {
	// 快速查询名称，必填
	FastQueryName string `json:"fastQueryName"`
}

type ListFastQueryRequest struct {
	// 日志组名称关键字查询，精确匹配，选填，默认全部
	Project string `json:"project"`
	// 日志集名称关键字查询，精确匹配，选填，默认全部
	LogStoreName string `json:"logStoreName"`
	// 快速查询名称关键字查询，模糊匹配，选填
	NamePattern string `json:"namePattern"`
	// 排序方式 选填 支持desc和asc，默认desc
	Order string `json:"order"`
	// 排序字段 选填 取值为：creationDateTime,lastModifiedTime,name  默认按照creationDateTime创建时间排序
	OrderBy string `json:"orderBy"`
	// 第几页 选填， 默认为1
	PageNo int `json:"pageNo"`
	// 每页大小 选填， 默认为10
	PageSize int `json:"pageSize"`
}

type CreateIndexRequest struct {
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
	// 是否开启全文索引，默认false，不开启全文索引
	Fulltext bool `json:"fulltext"`
	// 字段索引信息
	Fields map[string]api.LogField `json:"fields"`
}

type UpdateIndexRequest struct {
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
	// 是否开启全文索引，默认false，不开启全文索引
	Fulltext bool `json:"fulltext"`
	// 字段索引信息
	Fields map[string]api.LogField `json:"fields"`
}

type DeleteIndexRequest struct {
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
}

type DescribeIndexRequest struct {
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
}

type CreateLogShipperRequest struct {
	// 投递任务名称，必填
	LogShipperName string `json:"logShipperName"`
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称 必填
	LogStoreName string `json:"logStoreName"`
	// 投递的开始时间，选填 UTC时间，格式ISO8601，例如：2020-01-10T14:23:34Z 默认为当前时间
	StartTime string `json:"startTime"`
	// 投递类型 选填 取值为：BOS 目前只支持BOS， 默认为BOS
	DestType string `json:"destType"`
	// 投递目的端配置， 必填
	DestConfig *api.ShipperDestConfig `json:"destConfig"`
}

type UpdateLogShipperRequest struct {
	// 投递任务ID 必填
	LogShipperID string `json:"logShipperID"`
	// 投递任务名称，必填
	LogShipperName string `json:"logShipperName"`
	// 投递目的端配置， 必填
	DestConfig *api.ShipperDestConfig `json:"destConfig"`
}

type GetLogShipperRequest struct {
	// 投递任务ID 必填
	LogShipperID string `json:"logShipperID"`
}

type DeleteLogShipperRequest struct {
	// 投递任务ID 必填
	LogShipperID string `json:"logShipperID"`
}

type BulkDeleteLogShipperRequest struct {
	// 投递任务ID集合 必填
	LogShipperIDs []string `json:"logShipperIDs"`
}

type UpdateLogShipperStatusRequest struct {
	// 投递任务ID 必填
	LogShipperID string `json:"logShipperID"`
	// 投递任务状态 必填 取值: Running（启动），Paused（暂停）
	DesiredStatus string `json:"status"`
}

type BulkUpdateLogShipperStatusRequest struct {
	// 投递任务ID集合 必填
	LogShipperIDs []string `json:"logShipperIDs"`
	// 投递任务状态 必填 取值: Running（启动），Paused（暂停）
	DesiredStatus string `json:"desiredStatus"`
}

type ListLogShipperRequest struct {
	// 投递任务ID 选填
	LogShipperID string `json:"logShipperID"`
	// 投递任务名称关键字查询， 选填 模糊匹配
	LogShipperName string `json:"logShipperName"`
	// 日志组名称，选填，默认全部
	Project string `json:"project"`
	// 日志集名称名称关键字查询 选填  模糊匹配
	LogStoreName string `json:"logStoreName"`
	// 投递任务类型查询 选填 精确匹配 取值BOS 目前只支持BOS类型 默认BOS类型
	DestType string `json:"destType"`
	// 投递任务状态查询 选填 精确匹配 取值：运行中（Running）、异常（Abnormal）、已暂停（Paused）默认全部状态
	Status string `json:"status"`
	// 排序方式 选填 支持desc和asc，默认desc
	Order string `json:"order"`
	// 排序字段 选填 取值为：creationDateTime  默认按照creationDateTime创建时间排序
	OrderBy string `json:"orderBy"`
	// 第几页 必填
	PageNo int `json:"pageNo"`
	// 每页大小 必填
	PageSize int `json:"pageSize"`
}

type ListShipperRecordRequest struct {
	// 投递任务ID 必填
	LogShipperID string `json:"logShipperID"`
	// 投递记录的开始时间，查询范围为：(now-SinceHours, now) 选填，默认为1
	SinceHours int `json:"sinceHours"`
	// 第几页 选填， 默认为1
	PageNo int `json:"pageNo"`
	// 每页大小 选填， 默认为10
	PageSize int `json:"pageSize"`
}

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
	// 返回的数据条数，选填，默认为100，最大支持1000
	Limit int `json:"limit"`
	// 置顶查询的游标，默认为空，从头开始查询；如果不为，将从游标位置开始查询
	Marker string `json:"marker"`
	// 排序字段，默认desc，按照时间倒序排序；asc，按照时间顺序排序
	Sort string `json:"sort"`
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
	// 全文索引是否开启大小写敏感，默认false，不开启大小写敏感
	CaseSensitive bool `json:"caseSensitive"`
	// 全文分词符，将字段内容按照分词符拆分成若干个分词用于检索
	Separators string `json:"separators"`
	// 是否包含中文，默认为false，不包含中文
	IncludeChinese bool `json:"includeChinese"`
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
	// 全文索引是否开启大小写敏感，默认false，不开启大小写敏感
	CaseSensitive bool `json:"caseSensitive"`
	// 全文分词符，将字段内容按照分词符拆分成若干个分词用于检索
	Separators string `json:"separators"`
	// 是否包含中文，默认为false，不包含中文
	IncludeChinese bool `json:"includeChinese"`
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

type CreateDownloadTaskRequest struct {
	// 下载任务名称，选填
	Name string `json:"name"`
	// 日志组名称，选填，默认default
	Project string `json:"project"`
	// 日志集名称，必填
	LogStoreName string `json:"logStoreName"`
	// 日志流名称 选填，默认在全部日志流中下载数据
	LogStreamName string `json:"logStreamName"`
	// 查询语句，选填，默认match *
	Query string `json:"query"`
	// 日志的开始时间， 必填，UTC时间，格式ISO8601，例如：2020-01-10T13:23:34Z
	QueryStartTime string `json:"queryStartTime"`
	// 日志的结束时间，必填，UTC时间，格式ISO8601，例如：2020-01-10T14:23:34Z
	QueryEndTime string `json:"queryEndTime"`
	// 下载文件的格式，选填， 默认json，支持 json,csv
	Format string `json:"format"`
	// 下载日志的行数，选填，默认1000000，最大1000000
	Limit int64 `json:"limit"`
	// 排序方式 选填 支持desc和asc，默认desc，按照时间倒序排序
	Order string `json:"order"`
	// 下载文件的bos目录，选填，默认放到bls资源账号的下载目录
	// 如果不为空，表示放到用户自己的bos目录 需要确保bos的bucket存在，目录可以不存在，会自动创建
	FileDir string `json:"fileDir"`
}

type DescribeDownloadRequest struct {
	// 下载任务的UUID 必填
	UUID string `json:"uuid"`
}

type GetDownloadTaskLinkRequest struct {
	// 下载任务的UUID 必填
	UUID string `json:"uuid"`
}

type DeleteDownloadTaskRequest struct {
	// 下载任务的UUID 必填
	UUID string `json:"uuid"`
}

type ListDownloadTaskRequest struct {
	// 日志组名称关键字查询，精确匹配，选填
	Project string `json:"project"`
	// 日志集关键字查询，模糊匹配，选填
	LogStoreName string `json:"logStoreName"`
	// 排序方式 选填 支持desc和asc，默认desc
	Order string `json:"order"`
	// 排序字段 选填 支持 createdTime, updatedTime, name  默认按照createdTime创建时间排序
	OrderBy string `json:"orderBy"`
	// 第几页 选填， 默认为1
	PageNo int `json:"pageNo"`
	// 每页大小 选填， 默认为10
	PageSize int `json:"pageSize"`
}

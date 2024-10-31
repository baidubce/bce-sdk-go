package model

type GetMetricDataRequest struct {
	UserId         string            `json:"userId,omitempty"`
	Scope          string            `json:"scope,omitempty"`
	MetricName     string            `json:"metricName,omitempty"`
	Dimensions     map[string]string `json:"dimensions,omitempty"`
	Statistics     []string          `json:"statistics,omitempty"`
	StartTime      string            `json:"startTime,omitempty"`
	EndTime        string            `json:"endTime,omitempty"`
	PeriodInSecond int               `json:"periodInSecond,omitempty"`
}

type GetMetricDataResponse struct {
	RequestId  string        `json:"requestId,omitempty"`
	Code       string        `json:"code,omitempty"`
	Message    string        `json:"message,omitempty"`
	DataPoints []*DataPoints `json:"dataPoints,omitempty"`
}

type DataPoints struct {
	Average     *float64 `json:"average,omitempty"`
	Sum         *float64 `json:"sum,omitempty"`
	Minimum     *float64 `json:"minimum,omitempty"`
	Maximum     *float64 `json:"maximum,omitempty"`
	SampleCount *int64   `json:"sampleCount,omitempty"`
	Value       *float64 `json:"value,omitempty"`
	Timestamp   string   `json:"timestamp,omitempty"`
}

type BatchGetMetricDataRequest struct {
	UserId         string            `json:"userId,omitempty"`
	Scope          string            `json:"scope,omitempty"`
	MetricNames    []string          `json:"metricNames,omitempty"`
	Dimensions     map[string]string `json:"dimensions,omitempty"`
	Statistics     []string          `json:"statistics,omitempty"`
	StartTime      string            `json:"startTime,omitempty"`
	EndTime        string            `json:"endTime,omitempty"`
	PeriodInSecond int               `json:"periodInSecond,omitempty"`
}

type BatchGetMetricDataResponse struct {
	RequestId   string                       `json:"requestId,omitempty"`
	Code        string                       `json:"code,omitempty"`
	Message     string                       `json:"message,omitempty"`
	SuccessList []*SuccessBatchGetMetricData `json:"successList,omitempty"`
	ErrorList   []*ErrorBatchGetMetricData   `json:"errorList,omitempty"`
}

type SuccessBatchGetMetricData struct {
	MetricName string        `json:"metricName,omitempty"`
	Dimensions []*Dimension  `json:"dimensions,omitempty"`
	DataPoints []*DataPoints `json:"dataPoints,omitempty"`
}

type ErrorBatchGetMetricData struct {
	MetricName string       `json:"metricName,omitempty"`
	Dimensions []*Dimension `json:"dimensions,omitempty"`
	Message    string       `json:"message,omitempty"`
}

type Dimension struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Namespace struct {
	Name           string `json:"name,omitempty"`
	NamespaceAlias string `json:"namespaceAlias,omitempty"`
	UserId         string `json:"userId,omitempty"`
	Comment        string `json:"comment,omitempty"`
}

type CustomBatchNames struct {
	UserId string   `json:"userId,omitempty"`
	Names  []string `json:"names,omitempty"`
}

type ListNamespacesRequest struct {
	UserId   string `json:"userId,omitempty"`
	Name     string `json:"name,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type ListNamespacesResponse struct {
	PageNo     int                           `json:"pageNo,omitempty"`
	PageSize   int                           `json:"pageSize,omitempty"`
	TotalCount int                           `json:"totalCount,omitempty"`
	Result     []NamespaceWithMetricAndEvent `json:"result,omitempty"`
}

type NamespaceWithMetricAndEvent struct {
	Name           string              `json:"name,omitempty"`
	NamespaceAlias string              `json:"namespaceAlias,omitempty"`
	UserId         string              `json:"userId,omitempty"`
	Comment        string              `json:"comment,omitempty"`
	Metrics        []NamespaceItemView `json:"metrics,omitempty"`
	EventConfigs   []NamespaceItemView `json:"eventConfigs,omitempty"`
}

type NamespaceItemView struct {
	Name  string `json:"name,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type NamespaceMetric struct {
	Id          int64                      `json:"id,omitempty"`
	UserId      string                     `json:"userId,omitempty"`
	Namespace   string                     `json:"namespace,omitempty"`
	MetricName  string                     `json:"metricName,omitempty"`
	MetricAlias string                     `json:"metricAlias,omitempty"`
	Unit        string                     `json:"unit,omitempty"`
	Cycle       int                        `json:"cycle,omitempty"`
	Dimensions  []NamespaceMetricDimension `json:"dimensions"`
}

type NamespaceMetricDimension struct {
	Order int    `json:"order,omitempty"`
	Name  string `json:"name,omitempty"`
	Alias string `json:"alias,omitempty"`
}

type CustomBatchIds struct {
	UserId    string  `json:"userId,omitempty"`
	Namespace string  `json:"namespace,omitempty"`
	Ids       []int64 `json:"ids,omitempty"`
}

type ListNamespaceMetricsRequest struct {
	UserId      string `json:"userId,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
	MetricName  string `json:"metricName,omitempty"`
	MetricAlias string `json:"metricAlias,omitempty"`
	PageNo      int    `json:"pageNo,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

type ListNamespaceMetricsResponse struct {
	PageNo     int               `json:"pageNo,omitempty"`
	PageSize   int               `json:"pageSize,omitempty"`
	TotalCount int               `json:"totalCount,omitempty"`
	Result     []NamespaceMetric `json:"result,omitempty"`
}

type NamespaceEvent struct {
	UserId         string `json:"userId,omitempty"`
	Namespace      string `json:"namespace,omitempty"`
	EventName      string `json:"eventName,omitempty"`
	EventNameAlias string `json:"eventNameAlias,omitempty"`
	EventLevel     string `json:"eventLevel,omitempty"`
	Comment        string `json:"comment,omitempty"`
}

type CustomBatchEventNames struct {
	UserId    string   `json:"userId,omitempty"`
	Namespace string   `json:"namespace,omitempty"`
	Names     []string `json:"names,omitempty"`
}

type ListNamespaceEventsRequest struct {
	UserId     string `json:"userId,omitempty"`
	Namespace  string `json:"namespace,omitempty"`
	Name       string `json:"name,omitempty"`
	EventLevel string `json:"eventLevel,omitempty"`
	PageNo     int    `json:"pageNo,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
}

type ListNamespaceEventsResponse struct {
	PageNo     int              `json:"pageNo,omitempty"`
	PageSize   int              `json:"pageSize,omitempty"`
	TotalCount int              `json:"totalCount,omitempty"`
	Result     []NamespaceEvent `json:"result,omitempty"`
}

type EventDataRequest struct {
	PageNo       int    `json:"pageNo,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
	StartTime    string `json:"startTime,omitempty"`
	EndTime      string `json:"endTime,omitempty"`
	AccountID    string `json:"accountId,omitempty"`
	Ascending    bool   `json:"ascending,omitempty"`
	Scope        string `json:"scope,omitempty"`
	Region       string `json:"region,omitempty"`
	EventLevel   string `json:"eventLevel,omitempty"`
	EventName    string `json:"eventName,omitempty"`
	EventAlias   string `json:"eventAlias,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	ResourceID   string `json:"resourceId,omitempty"`
	EventID      string `json:"eventId,omitempty"`
}

type CloudEventResponse struct {
	PageNumber    int              `json:"pageNumber,omitempty"`
	PageSize      int              `json:"pageSize,omitempty"`
	PageElements  int              `json:"pageElements,omitempty"`
	Last          bool             `json:"last,omitempty"`
	First         bool             `json:"first,omitempty"`
	TotalPages    int              `json:"totalPages,omitempty"`
	TotalElements int              `json:"totalElements,omitempty"`
	Content       []CloudEventData `json:"content,omitempty"`
}

type CloudEventData struct {
	AccountID    string `json:"accountId,omitempty"`
	ServiceName  string `json:"serviceName,omitempty"`
	Region       string `json:"region,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	ResourceID   string `json:"resourceId,omitempty"`
	EventID      string `json:"eventId,omitempty"`
	EventType    string `json:"eventType,omitempty"`
	EventLevel   string `json:"eventLevel,omitempty"`
	EventAlias   string `json:"eventAlias,omitempty"`
	Timestamp    string `json:"timestamp,omitempty"`
	Content      string `json:"content,omitempty"`
}

type CloudEvent struct {
	EventID      string `json:"eventId,omitempty"`
	EventName    string `json:"eventName,omitempty"`
	EventAlias   string `json:"eventAlias,omitempty"`
	EventLevel   string `json:"eventLevel,omitempty"`
	EventTime    string `json:"eventTime,omitempty"`
	EventContent string `json:"eventContent,omitempty"`
	EventSource  string `json:"eventSource,omitempty"`
	EventStatus  string `json:"eventStatus,omitempty"`
	EventDetail  string `json:"eventDetail,omitempty"`
}

type PlatformEventResponse struct {
	Content       []PlatformEventData `json:"content,omitempty"`
	PageNumber    int                 `json:"pageNumber,omitempty"`
	PageSize      int                 `json:"pageSize,omitempty"`
	PageElements  int                 `json:"pageElements,omitempty"`
	Last          bool                `json:"last,omitempty"`
	First         bool                `json:"first,omitempty"`
	TotalPages    int                 `json:"totalPages,omitempty"`
	TotalElements int64               `json:"totalElements,omitempty"`
}

type PlatformEventData struct {
	EventID      string `json:"eventId,omitempty"`
	EventName    string `json:"eventName,omitempty"`
	EventAlias   string `json:"eventAlias,omitempty"`
	EventLevel   string `json:"eventLevel,omitempty"`
	EventTime    string `json:"eventTime,omitempty"`
	EventContent string `json:"eventContent,omitempty"`
	EventSource  string `json:"eventSource,omitempty"`
	EventStatus  string `json:"eventStatus,omitempty"`
	EventDetail  string `json:"eventDetail,omitempty"`
}

type EventPolicy struct {
	AccountID       string              `json:"accountId,omitempty"`
	ServiceName     string              `json:"serviceName,omitempty"`
	Name            string              `json:"name,omitempty"`
	BlockStatus     string              `json:"blockStatus,omitempty"`
	EventFilter     EventFilter         `json:"eventFilter,omitempty"`
	Resource        EventResourceFilter `json:"resource,omitempty"`
	IncidentActions []string            `json:"incidentActions,omitempty"`
}

type EventFilter struct {
	EventLevel      string   `json:"eventLevel,omitempty"`
	EventTypeList   []string `json:"eventTypeList,omitempty"`
	EventAliasNames []string `json:"eventAliasNames,omitempty"`
}

type EventResourceFilter struct {
	Region            string          `json:"region,omitempty"`
	Type              string          `json:"type,omitempty"`
	MonitorObjectType string          `json:"monitorObjectType,omitempty"`
	Resources         []EventResource `json:"resources,omitempty"`
}

type EventResource struct {
	Identifiers []Dimension `json:"identifiers,omitempty"`
}

type MergedGroup struct {
	ID             string            `json:"id,omitempty"`
	UserId         string            `json:"userId,omitempty"`
	Region         string            `json:"region,omitempty"`
	ServiceName    string            `json:"serviceName,omitempty"`
	TypeName       string            `json:"typeName,omitempty"`
	Name           string            `json:"name,omitempty"`
	ResourceIDList []MonitorResource `json:"resourceIdList,omitempty"`
}

type MonitorResource struct {
	UserId        string      `json:"userId,omitempty"`
	Region        string      `json:"region,omitempty"`
	ServiceName   string      `json:"serviceName,omitempty"`
	TypeName      string      `json:"typeName,omitempty"`
	ResourceID    string      `json:"resourceId,omitempty"`
	ErrUpdateTime string      `json:"errUpdateTime,omitempty"`
	Identifiers   []Dimension `json:"identifiers,omitempty"`
	Properties    []Dimension `json:"properties,omitempty"`
	Tags          []Dimension `json:"tags,omitempty"`
}

type InstanceGroup struct {
	ID               int64  `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	ServiceName      string `json:"serviceName,omitempty"`
	TypeName         string `json:"typeName,omitempty"`
	Region           string `json:"region,omitempty"`
	UserID           string `json:"userId,omitempty"`
	UUID             string `json:"uuid,omitempty"`
	Count            int    `json:"count,omitempty"`
	ServiceNameAlias string `json:"serviceNameAlias,omitempty"`
	TypeNameAlias    string `json:"typeNameAlias,omitempty"`
	RegionAlias      string `json:"regionAlias,omitempty"`
	TagKey           string `json:"tagKey,omitempty"`
}

type InstanceGroupBase struct {
	ID     string `json:"id,omitempty"`
	UserID string `json:"userId,omitempty"`
}

type InstanceGroupQuery struct {
	UserID      string `json:"userId,omitempty"`
	Name        string `json:"name,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
	Region      string `json:"region,omitempty"`
	TypeName    string `json:"typeName,omitempty"`
	PageNo      int    `json:"pageNo,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

type InstanceGroupListResponse struct {
	OrderBy    string          `json:"orderBy,omitempty"`
	Order      string          `json:"order,omitempty"`
	PageNo     int             `json:"pageNo,omitempty"`
	PageSize   int             `json:"pageSize,omitempty"`
	TotalCount int             `json:"totalCount,omitempty"`
	Result     []InstanceGroup `json:"result,omitempty"`
}

type IGInstanceQuery struct {
	UserID      string `json:"userId,omitempty"`
	ID          string `json:"id,omitempty"`
	UUID        string `json:"uuid,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
	TypeName    string `json:"typeName,omitempty"`
	Region      string `json:"region,omitempty"`
	ViewType    string `json:"viewType,omitempty"`
	PageNo      int    `json:"pageNo,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
	KeywordType string `json:"keywordType,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
}

type IGInstanceListResponse struct {
	OrderBy    string             `json:"orderBy,omitempty"`
	Order      string             `json:"order,omitempty"`
	PageNo     int                `json:"pageNo,omitempty"`
	PageSize   int                `json:"pageSize,omitempty"`
	TotalCount int                `json:"totalCount,omitempty"`
	Result     [][]IGInstanceItem `json:"result,omitempty"`
}

type IGInstanceItem struct {
	ItemName        string `json:"itemName,omitempty"`
	ItemAlias       string `json:"itemAlias,omitempty"`
	ItemValue       string `json:"itemValue,omitempty"`
	ItemIdentitable bool   `json:"itemIdentitable,omitempty"`
	ItemDimension   string `json:"itemDimension,omitempty"`
}

type MultiDimensionalLatestMetricsRequest struct {
	UserID       string      `json:"userId,omitempty"`
	Region       string      `json:"region,omitempty"`
	Scope        string      `json:"scope,omitempty"`
	ResourceType string      `json:"resourceType,omitempty"`
	Dimensions   []Dimension `json:"dimensions,omitempty"`
	MetricNames  []string    `json:"metricNames,omitempty"`
	Timestamp    string      `json:"timestamp,omitempty"`
	Statistics   []string    `json:"statistics,omitempty"`
	Cycle        int         `json:"cycle,omitempty"`
}

type MetricsByPartialDimensionsRequest struct {
	UserID       string      `json:"userId,omitempty"`
	Scope        string      `json:"scope,omitempty"`
	StartTime    string      `json:"startTime,omitempty"`
	EndTime      string      `json:"endTime,omitempty"`
	Statistics   []string    `json:"statistics,omitempty"`
	Cycle        int         `json:"cycle,omitempty"`
	Dimensions   []Dimension `json:"dimensions,omitempty"`
	ResourceType string      `json:"resourceType,omitempty"`
	MetricName   string      `json:"metricName,omitempty"`
	Region       string      `json:"region,omitempty"`
	PageNo       int         `json:"pageNo,omitempty"`
	PageSize     int         `json:"pageSize,omitempty"`
}

type DataPoint struct {
	Average     float64 `json:"average,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Minimum     float64 `json:"minimum,omitempty"`
	Maximum     float64 `json:"maximum,omitempty"`
	SampleCount int     `json:"sampleCount,omitempty"`
	Value       string  `json:"value,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
}

type MultiDimensionalMetric struct {
	Region     string       `json:"region,omitempty"`
	Scope      string       `json:"scope,omitempty"`
	UserID     string       `json:"userId,omitempty"`
	ResourceID string       `json:"resourceId,omitempty"`
	MetricName string       `json:"metricName,omitempty"`
	Dimensions []Dimension  `json:"dimensions,omitempty"`
	DataPoints []*DataPoint `json:"dataPoints,omitempty"`
}

type MultiDimensionalMetricsResponse struct {
	RequestID string                    `json:"requestId,omitempty"`
	Code      string                    `json:"code,omitempty"`
	Message   interface{}               `json:"message,omitempty"`
	Metrics   []*MultiDimensionalMetric `json:"metrics,omitempty"`
}

type MetricsByPartialDimensionsPageResponse struct {
	RequestID string `json:"requestId,omitempty"`
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
	Result    struct {
		OrderBy    string `json:"orderBy,omitempty"`
		Order      string `json:"order,omitempty"`
		PageNo     int    `json:"pageNo,omitempty"`
		PageSize   int    `json:"pageSize,omitempty"`
		TotalCount int    `json:"totalCount,omitempty"`
		Result     []struct {
			Region     string      `json:"region,omitempty"`
			Scope      string      `json:"scope,omitempty"`
			UserID     string      `json:"userId,omitempty"`
			ResourceID string      `json:"resourceId,omitempty"`
			MetricName string      `json:"metricName,omitempty"`
			Dimensions []Dimension `json:"dimensions,omitempty"`
			DataPoints []struct {
				Average     *float64 `json:"average,omitempty"`
				Sum         *float64 `json:"sum,omitempty"`
				Minimum     *float64 `json:"minimum,omitempty"`
				Maximum     *float64 `json:"maximum,omitempty"`
				SampleCount int      `json:"sampleCount,omitempty"`
				Value       string   `json:"value,omitempty"`
				Timestamp   string   `json:"timestamp,omitempty"`
			} `json:"dataPoints,omitempty"`
		} `json:"result,omitempty"`
	} `json:"result,omitempty"`
}

type TsdbMetricAllDataQueryRequest struct {
	UserID      string        `json:"userId,omitempty"`
	Region      string        `json:"region,omitempty"`
	Scope       string        `json:"scope,omitempty"`
	Type        string        `json:"type,omitempty"`
	Dimensions  [][]Dimension `json:"dimensions,omitempty"`
	MetricNames []string      `json:"metricNames,omitempty"`
	Statistics  []string      `json:"statistics,omitempty"`
	Cycle       int           `json:"cycle,omitempty"`
	StartTime   string        `json:"startTime,omitempty"`
	EndTime     string        `json:"endTime,omitempty"`
}

type TsdbDimensionTopQuery struct {
	UserID     string            `json:"userId,omitempty"`
	Region     string            `json:"region,omitempty"`
	Scope      string            `json:"scope,omitempty"`
	Dimensions map[string]string `json:"dimensions,omitempty"`
	MetricName string            `json:"metricName,omitempty"`
	Statistics string            `json:"statistics,omitempty"`
	StartTime  string            `json:"startTime,omitempty"`
	EndTime    string            `json:"endTime,omitempty"`
	Order      string            `json:"order,omitempty"`
	TopNum     int               `json:"topNum,omitempty"`
	Labels     []string          `json:"labels,omitempty"`
	Cycle      int               `json:"cycle,omitempty"`
}

type TsdbDimensionTopResult struct {
	RequestId string    `json:"requestId,omitempty"`
	TopDatas  []TopData `json:"topDatas,omitempty"`
}

type TopData struct {
	Order      int         `json:"order,omitempty"`
	Dimensions []Dimension `json:"dimensions,omitempty"`
}

type TsdbQueryMetaData struct {
	RequestId   string        `json:"requestId,omitempty"`
	UserId      string        `json:"userId,omitempty"`
	ServiceName string        `json:"serviceName,omitempty"`
	MetricName  string        `json:"metricName,omitempty"`
	Statistics  []string      `json:"statistics,omitempty"`
	ResourceID  string        `json:"resourceId,omitempty"`
	DataPoints  []*DataPoints `json:"dataPoints,omitempty"`
}

type AlarmListQuery struct {
	UserID         string              `json:"userId,omitempty"`
	States         []string            `json:"states,omitempty"`
	AlarmType      string              `json:"alarmType,omitempty"`
	StartTime      int64               `json:"startTime,omitempty"`
	EndTime        int64               `json:"endTime,omitempty"`
	Sort           string              `json:"sort,omitempty"`
	Ascending      bool                `json:"ascending,omitempty"`
	PageNo         int                 `json:"pageNo"`
	PageSize       int                 `json:"pageSize"`
	Scope          string              `json:"scope,omitempty"`
	Region         string              `json:"region,omitempty"`
	ResourceType   string              `json:"resourceType,omitempty"`
	Level          string              `json:"level,omitempty"`
	AlarmAliasName string              `json:"alarmAliasName,omitempty"`
	Resource       map[string]string   `json:"resource,omitempty"`
	Resources      []map[string]string `json:"resources,omitempty"`
}

type AlarmListResponse struct {
	Success bool   `json:"success,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Result  struct {
		Alarms     []Alarm `json:"alarms,omitempty"`
		PageNo     int     `json:"pageNo,omitempty"`
		PageSize   int     `json:"pageSize,omitempty"`
		TotalCount int     `json:"totalCount,omitempty"`
	} `json:"result,omitempty"`
}

type AlarmDetailQuery struct {
	UserID  string `json:"userId,omitempty"`
	AlarmID string `json:"alarmId,omitempty"`
}

type AlarmDetailResponse struct {
	Success bool   `json:"success,omitempty"`
	Msg     string `json:"msg,omitempty"`
	Result  Alarm  `json:"result,omitempty"`
}

type AlertMetric struct {
	Metric struct {
		Name        string            `json:"name,omitempty"`
		Value       float64           `json:"value,omitempty"`
		Dimensions  map[string]string `json:"dimensions,omitempty"`
		AliasName   string            `json:"aliasName,omitempty"`
		AliasNameEn string            `json:"aliasNameEn,omitempty"`
		Unit        string            `json:"unit,omitempty"`
	} `json:"metric,omitempty"`
	Rule struct {
		Seq             int32               `json:"seq"`
		MetricName      string              `json:"metricName"`
		MetricDimension map[string][]string `json:"metricDimension,omitempty"`
		Operator        string              `json:"operator"`
		Threshold       float64             `json:"threshold"`
		Statistics      string              `json:"statistics"`
		Window          int32               `json:"window"`
	} `json:"rule,omitempty"`
}

type Alarm struct {
	ID          string `json:"id,omitempty"`
	UserID      string `json:"userId"`
	SeriesID    string `json:"seriesId"`
	State       string `json:"state"`
	InitState   string `json:"initState"`
	CloseReason string `json:"closeReason,omitempty"`
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime,omitempty"`
	AlarmType   string `json:"alarmType,omitempty"`
	Resource    struct {
		Scope            string            `json:"scope,omitempty"`
		ResourceType     string            `json:"resourceType,omitempty"`
		Region           string            `json:"region,omitempty"`
		Identifiers      map[string]string `json:"identifiers,omitempty"`
		Properties       map[string]string `json:"properties,omitempty"`
		MetricDimensions map[string]string `json:"metricDimensions,omitempty"`
	} `json:"resource,omitempty"`
	Policy struct {
		ID          int64             `json:"id,omitempty"`
		Name        string            `json:"name,omitempty"`
		IndexedName string            `json:"indexedName,omitempty"`
		AliasName   string            `json:"aliasName,omitempty"`
		UpdateTime  int64             `json:"updateTime,omitempty"`
		Content     string            `json:"content,omitempty"`
		ContentEn   string            `json:"contentEn,omitempty"`
		Level       string            `json:"level,omitempty"`
		Extra       map[string]string `json:"extra,omitempty"`
	} `json:"policy,omitempty"`
	Actions []struct {
		Name          string    `json:"name,omitempty"`
		Type          string    `json:"type,omitempty"`
		ExecutedTime  int64     `json:"executedTime,omitempty"`
		Alias         string    `json:"alias"`         // 通知模板名称，也用标识是否 hover 展示
		Notifications *[]string `json:"notifications"` // 通知方式,
		CallBacks     *[]string `json:"callBacks"`     // 回调
		Members       *[]string `json:"members"`       // 用户/用户组}
	} `json:"actions,omitempty"`
	AlertMetrics []AlertMetric `json:"alertMetrics,omitempty"`
}

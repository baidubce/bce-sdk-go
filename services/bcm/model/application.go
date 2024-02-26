package model

// Application information management

type ApplicationInfoRequest struct {
	Alias       string `json:"alias,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	UserID      string `json:"userId,omitempty"`
	Description string `json:"description,omitempty"`
}

type ApplicationInfoResponse struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Alias       string `json:"alias,omitempty"`
	Type        string `json:"type,omitempty"`
	UserID      string `json:"userId,omitempty"`
	Description string `json:"description,omitempty"`
}

type ApplicationInfoListResponse struct {
	Content          []*ApplicationInfoResponse `json:"content,omitempty"`
	Pageable         *Pageable                  `json:"pageable,omitempty"`
	Last             bool                       `json:"last,omitempty"`
	TotalElements    int                        `json:"totalElements,omitempty"`
	TotalPages       int                        `json:"totalPages,omitempty"`
	First            bool                       `json:"first,omitempty"`
	Sort             *Sort                      `json:"sort,omitempty"`
	Number           int                        `json:"number,omitempty"`
	NumberOfElements int                        `json:"numberOfElements,omitempty"`
	Size             int                        `json:"size,omitempty"`
}

type Sort struct {
	Unsorted bool `json:"unsorted,omitempty"`
	Sorted   bool `json:"sorted,omitempty"`
}

type Pageable struct {
	Sort       *Sort `json:"sort,omitempty"`
	PageSize   int   `json:"pageSize,omitempty"`
	PageNumber int   `json:"pageNumber,omitempty"`
	Offset     int   `json:"offset,omitempty"`
	Unpaged    bool  `json:"unpaged,omitempty"`
	Paged      bool  `json:"paged,omitempty"`
}

type ApplicationInfoUpdateRequest struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Alias       string `json:"alias,omitempty"`
	Type        string `json:"type,omitempty"`
	UserID      string `json:"userId,omitempty"`
	Description string `json:"description,omitempty"`
}

type ApplicationInfoDeleteRequest struct {
	Name string `json:"name,omitempty"`
}

// Application instance information management

type ApplicationInstanceListRequest struct {
	PageSize    int    `json:"pageSize,omitempty"`
	PageNo      int    `json:"pageNo,omitempty"`
	SearchName  string `json:"searchName,omitempty"`
	SearchValue string `json:"searchValue,omitempty"`
	Region      string `json:"region,omitempty"`
	AppName     string `json:"appName,omitempty"`
}

type ApplicationInstanceInfo struct {
	FloatingIP   string `json:"floatingIp,omitempty"`
	HasBinded    bool   `json:"hasBinded,omitempty"`
	ID           string `json:"id,omitempty"`
	InstanceID   string `json:"instanceId,omitempty"`
	InstanceUUID string `json:"instanceUuid,omitempty"`
	InternalIP   string `json:"internalIp,omitempty"`
	Name         string `json:"name,omitempty"`
	PublicIP     string `json:"publicIp,omitempty"`
}

type ApplicationInstanceListResponse struct {
	Content       []*ApplicationInstanceInfo `json:"content,omitempty"`
	Fields        []interface{}              `json:"fields,omitempty"`
	First         bool                       `json:"first,omitempty"`
	Last          bool                       `json:"last,omitempty"`
	OrderBy       []interface{}              `json:"orderBy,omitempty"`
	PageElements  int                        `json:"pageElements,omitempty"`
	PageNumber    int                        `json:"pageNumber,omitempty"`
	PageSize      int                        `json:"pageSize,omitempty"`
	Query         interface{}                `json:"query,omitempty"`
	TotalElements int                        `json:"totalElements,omitempty"`
	TotalPages    int                        `json:"totalPages,omitempty"`
}

type ApplicationInstanceCreateRequest struct {
	AppName  string              `json:"appName,omitempty"`
	UserID   string              `json:"userId,omitempty"`
	HostList []*HostInstanceInfo `json:"hostList,omitempty"`
}

type HostInstanceInfo struct {
	InstanceID string `json:"instanceId,omitempty"`
	Region     string `json:"region,omitempty"`
}

type ApplicationInstanceCreatedListRequest struct {
	UserID  string `json:"userId,omitempty"`
	AppName string `json:"appName,omitempty"`
	Region  string `json:"region,omitempty"`
}

type ApplicationInstanceCreatedInfo struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	InstanceOffSet int    `json:"instanceOffSet,omitempty"`
	InstanceID     string `json:"instanceId,omitempty"`
	Region         string `json:"region,omitempty"`
	UserID         string `json:"userId,omitempty"`
	AppName        string `json:"appName,omitempty"`
	HostName       string `json:"hostName,omitempty"`
	IP             string `json:"ip,omitempty"`
}

type ApplicationInstanceCreatedListResponse struct {
	Content       []*ApplicationInstanceCreatedInfo `json:"content,omitempty"`
	Query         interface{}                       `json:"query,omitempty"`
	Fields        []interface{}                     `json:"fields,omitempty"`
	OrderBy       []interface{}                     `json:"orderBy,omitempty"`
	PageNumber    int                               `json:"pageNumber,omitempty"`
	PageSize      int                               `json:"pageSize,omitempty"`
	PageElements  int                               `json:"pageElements,omitempty"`
	Last          bool                              `json:"last,omitempty"`
	First         bool                              `json:"first,omitempty"`
	TotalPages    int                               `json:"totalPages,omitempty"`
	TotalElements int                               `json:"totalElements,omitempty"`
}

type ApplicationInstanceDeleteRequest struct {
	ID      string `json:"id,omitempty"`
	AppName string `json:"appName,omitempty"`
}

// Application monitoring task management

type ApplicationMonitorTaskInfoLogRequest struct {
	AppName       string              `json:"appName,omitempty"`
	AliasName     string              `json:"aliasName,omitempty"`
	Type          int                 `json:"type,omitempty"`
	Description   string              `json:"description,omitempty"`
	Target        string              `json:"target,omitempty"`
	Cycle         int                 `json:"cycle,omitempty"`
	Rate          int                 `json:"rate,omitempty"`
	ExtractResult []*LogExtractResult `json:"extractResult,omitempty"`
	LogExample    string              `json:"logExample,omitempty"`
	MatchRule     string              `json:"matchRule,omitempty"`
	UserID        string              `json:"userId,omitempty"`
	Metrics       []*Metric           `json:"metrics,omitempty"`
}

type ApplicationMonitorTaskInfoRequest struct {
	AppName     string `json:"appName,omitempty"`
	AliasName   string `json:"aliasName,omitempty"`
	Type        int    `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Target      string `json:"target,omitempty"`
	Cycle       int    `json:"cycle,omitempty"`
}

type ApplicationMonitorTaskResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	AliasName      string `json:"aliasName"`
	AppName        string `json:"appName"`
	UserID         string `json:"userId"`
	Cycle          int    `json:"cycle"`
	Target         string `json:"target"`
	Type           int    `json:"type"`
	Description    string `json:"description"`
	HasAlarmConfig bool   `json:"hasAlarmConfig"`
}

type LogExtractResult struct {
	ID                int    `json:"id,omitempty"`
	TaskID            int    `json:"taskId,omitempty"`
	ExtractFieldName  string `json:"extractFieldName,omitempty"`
	ExtractFieldValue string `json:"extractFieldValue,omitempty"`
	MetricEnable      int    `json:"metricEnable,omitempty"`
	DimensionMapTable string `json:"dimensionMapTable,omitempty"`
}

type Metric struct {
	ID               int       `json:"id,omitempty"`
	TaskID           int       `json:"taskId,omitempty"`
	MetricName       string    `json:"metricName"`
	SaveInstanceData int       `json:"saveInstanceData"`
	ValueFieldType   int       `json:"valueFieldType"`
	AggrTags         []*AggTag `json:"aggrTags"`
	MetricAlias      string    `json:"metricAlias"`
	MetricUnit       string    `json:"metricUnit"`
	ValueFieldName   string    `json:"valueFieldName"`
}

type AggTag struct {
	Range string `json:"range"`
	Tags  string `json:"tags"`
}
type ApplicationMonitorTaskInfoListResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	AliasName      string `json:"aliasName"`
	AppName        string `json:"appName"`
	UserID         string `json:"userId"`
	Cycle          int    `json:"cycle"`
	Target         string `json:"target"`
	Type           int    `json:"type"`
	Description    string `json:"description"`
	HasAlarmConfig bool   `json:"hasAlarmConfig"`
}

type ApplicationMonitorTaskInfoNormalResponse struct {
	AppName     string `json:"appName"`
	Type        int    `json:"type"`
	AliasName   string `json:"aliasName"`
	Cycle       int    `json:"cycle"`
	Target      string `json:"target"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type ApplicationMonitorTaskInfoLogResponse struct {
	ID             int                 `json:"id,omitempty"`
	Name           string              `json:"name,omitempty"`
	AliasName      string              `json:"aliasName,omitempty"`
	AppName        string              `json:"appName,omitempty"`
	UserID         string              `json:"userId,omitempty"`
	Cycle          int                 `json:"cycle,omitempty"`
	Target         string              `json:"target,omitempty"`
	Type           int                 `json:"type"`
	Description    string              `json:"description,omitempty"`
	HasAlarmConfig bool                `json:"hasAlarmConfig"`
	LogExample     string              `json:"logExample,omitempty"`
	MatchRule      string              `json:"matchRule,omitempty"`
	Rate           int                 `json:"rate,omitempty"`
	ExtractResult  []*LogExtractResult `json:"extractResult,omitempty"`
	Metrics        []*Metric           `json:"metrics,omitempty"`
}

type ApplicationMonitorTaskDetailRequest struct {
	UserID   string `json:"userId,omitempty"`
	AppName  string `json:"appName,omitempty"`
	TaskName string `json:"taskName,omitempty"`
}

type ApplicationMonitorTaskListRequest struct {
	UserID  string `json:"userId,omitempty"`
	AppName string `json:"appName,omitempty"`
	Type    string `json:"type,omitempty"`
}

type ApplicationMonitorTaskInfoUpdateRequest struct {
	ID            string              `json:"id,omitempty"`
	Name          string              `json:"name,omitempty"`
	AppName       string              `json:"appName,omitempty"`
	AliasName     string              `json:"aliasName,omitempty"`
	Type          int                 `json:"type,omitempty"`
	Description   string              `json:"description,omitempty"`
	Target        string              `json:"target,omitempty"`
	Cycle         int                 `json:"cycle,omitempty"`
	Rate          int                 `json:"rate,omitempty"`
	ExtractResult []*LogExtractResult `json:"extractResult,omitempty"`
	LogExample    string              `json:"logExample,omitempty"`
	MatchRule     string              `json:"matchRule,omitempty"`
	UserID        string              `json:"userId,omitempty"`
	Metrics       []*Metric           `json:"metrics,omitempty"`
}

type ApplicationMonitorTaskDeleteRequest struct {
	UserID  string `json:"userId,omitempty"`
	Name    string `json:"name,omitempty"`
	AppName string `json:"appName,omitempty"`
}

// Dimension mapping table management

type ApplicationDimensionTableInfoRequest struct {
	UserID         string `json:"userId,omitempty"`
	AppName        string `json:"appName,omitempty"`
	TableName      string `json:"tableName,omitempty"`
	MapContentJSON string `json:"mapContentJson,omitempty"`
}

type ApplicationDimensionTableInfoResponse struct {
	UserID         string `json:"userId,omitempty"`
	AppName        string `json:"appName,omitempty"`
	MapContentJSON string `json:"mapContentJson,omitempty"`
	TableName      string `json:"tableName,omitempty"`
}

type ApplicationDimensionTableListRequest struct {
	UserID     string `json:"userId,omitempty"`
	AppName    string `json:"appName,omitempty"`
	SearchName string `json:"searchName,omitempty"`
}

type ApplicationDimensionTableDeleteRequest struct {
	UserID    string `json:"userId,omitempty"`
	AppName   string `json:"appName,omitempty"`
	TableName string `json:"tableName,omitempty"`
}

package model

type Level string

const (
	LevelNotice   Level = "NOTICE"
	LevelWarning  Level = "WARNING"
	LevelCritical Level = "CRITICAL"
	LevelMajor    Level = "MAJOR"
	LevelCustom   Level = "CUSTOM"
)

type CreateHTTPTask struct {
	UserID        string `json:"userId,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	TaskName      string `json:"taskName,omitempty"`
	Address       string `json:"address,omitempty"`
	Method        string `json:"method,omitempty"`
	PostContent   string `json:"postContent,omitempty"`
	Cycle         int    `json:"cycle,omitempty"`
	Idc           string `json:"idc,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	AdvanceConfig string `json:"advanceConfig,omitempty"`
	IPType        string `json:"ipType,omitempty"`

	Cookies       string `json:"cookies,omitempty"`
	Host          string `json:"host,omitempty"`
	UserAgent     string `json:"userAgent,omitempty"`
	ResponseCode  string `json:"responseCode,omitempty"`
	ResponseCheck string `json:"responseCheck,omitempty"`
	UserName      string `json:"userName,omitempty"`
	Password      string `json:"password,omitempty"`
}

type CreateHTTPTaskResponse struct {
	UserID        string `json:"userId,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	TaskName      string `json:"taskName,omitempty"`
	Address       string `json:"address,omitempty"`
	Method        string `json:"method,omitempty"`
	PostContent   string `json:"postContent,omitempty"`
	Cycle         int    `json:"cycle,omitempty"`
	Idc           string `json:"idc,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	AdvanceConfig bool   `json:"advanceConfig,omitempty"`
	IPType        string `json:"ipType,omitempty"`

	Cookies       string `json:"cookies,omitempty"`
	Host          string `json:"host,omitempty"`
	UserAgent     string `json:"userAgent,omitempty"`
	ResponseCode  string `json:"responseCode,omitempty"`
	ResponseCheck string `json:"responseCheck,omitempty"`
	UserName      string `json:"userName,omitempty"`
	Password      string `json:"password,omitempty"`
}

type CreateHTTPSTask struct {
	UserID        string `json:"userId,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	TaskName      string `json:"taskName,omitempty"`
	Address       string `json:"address,omitempty"`
	Method        string `json:"method,omitempty"`
	PostContent   string `json:"postContent,omitempty"`
	Cycle         int    `json:"cycle,omitempty"`
	Idc           string `json:"idc,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	AdvanceConfig string `json:"advanceConfig,omitempty"`
	IPType        string `json:"ipType,omitempty"`

	Cookies       string `json:"cookies,omitempty"`
	Host          string `json:"host,omitempty"`
	UserAgent     string `json:"userAgent,omitempty"`
	ResponseCode  string `json:"responseCode,omitempty"`
	ResponseCheck string `json:"responseCheck,omitempty"`
	UserName      string `json:"userName,omitempty"`
	Password      string `json:"password,omitempty"`
}

type CreateHTTPSTaskResponse struct {
	UserID        string `json:"userId,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	TaskName      string `json:"taskName,omitempty"`
	Address       string `json:"address,omitempty"`
	Method        string `json:"method,omitempty"`
	PostContent   string `json:"postContent,omitempty"`
	Cycle         int    `json:"cycle,omitempty"`
	Idc           string `json:"idc,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	AdvanceConfig bool   `json:"advanceConfig,omitempty"`
	IPType        string `json:"ipType,omitempty"`

	Cookies       string `json:"cookies,omitempty"`
	Host          string `json:"host,omitempty"`
	UserAgent     string `json:"userAgent,omitempty"`
	ResponseCode  string `json:"responseCode,omitempty"`
	ResponseCheck string `json:"responseCheck,omitempty"`
	UserName      string `json:"userName,omitempty"`
	Password      string `json:"password,omitempty"`
}

type CreatePingTask struct {
	UserID         string `json:"userId,omitempty"`
	TaskID         string `json:"taskId,omitempty"`
	TaskName       string `json:"taskName,omitempty"`
	Address        string `json:"address,omitempty"`
	PacketCount    int    `json:"packetCount,omitempty"`
	PacketLossRate int    `json:"packetLossRate,omitempty"`
	Cycle          int    `json:"cycle,omitempty"`
	Idc            string `json:"idc,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	IPType         string `json:"ipType,omitempty"`
}

type CreateTCPTask struct {
	UserID        string `json:"userId,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	TaskName      string `json:"taskName,omitempty"`
	Address       string `json:"address,omitempty"`
	Cycle         int    `json:"cycle,omitempty"`
	Idc           string `json:"idc,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	Port          int    `json:"port,omitempty"`
	AdvanceConfig bool   `json:"advanceConfig,omitempty"`
	IPType        string `json:"ipType,omitempty"`

	InputType      int    `json:"inputType,omitempty"`
	OutputType     int    `json:"outputType,omitempty"`
	Input          string `json:"input,omitempty"`
	ExpectedOutput string `json:"expectedOutput,omitempty"`
}

type CreateUDPTask struct {
	UserID   string `json:"userId,omitempty"`
	TaskID   string `json:"taskId,omitempty"`
	TaskName string `json:"taskName,omitempty"`
	Address  string `json:"address,omitempty"`
	Cycle    int    `json:"cycle,omitempty"`
	Idc      string `json:"idc,omitempty"`
	Timeout  int    `json:"timeout,omitempty"`
	Port     int    `json:"port,omitempty"`
	IPType   string `json:"ipType,omitempty"`

	InputType      int    `json:"inputType,omitempty"`
	OutputType     int    `json:"outputType,omitempty"`
	Input          string `json:"input,omitempty"`
	ExpectedOutput string `json:"expectedOutput,omitempty"`
}

type CreateFtpTask struct {
	UserID         string `json:"userId,omitempty"`
	TaskID         string `json:"taskId,omitempty"`
	TaskName       string `json:"taskName,omitempty"`
	Address        string `json:"address,omitempty"`
	Cycle          int    `json:"cycle,omitempty"`
	Idc            string `json:"idc,omitempty"`
	Timeout        int    `json:"timeout,omitempty"`
	IPType         string `json:"ipType,omitempty"`
	Port           int    `json:"port,omitempty"`
	AnonymousLogin bool   `json:"anonymousLogin,omitempty"`
	UserName       string `json:"userName,omitempty"`
	Password       string `json:"password,omitempty"`
}

type CreateDNSTask struct {
	UserID      string `json:"userId,omitempty"`
	TaskID      string `json:"taskId,omitempty"`
	TaskName    string `json:"taskName,omitempty"`
	Address     string `json:"address,omitempty"`
	Cycle       int    `json:"cycle,omitempty"`
	Idc         string `json:"idc,omitempty"`
	Timeout     int    `json:"timeout,omitempty"`
	IPType      string `json:"ipType,omitempty"`
	Server      string `json:"server,omitempty"`
	ResolveType string `json:"resolveType,omitempty"`
	KidnapWhite string `json:"kidnapWhite,omitempty"`
}

type GetTaskDetailRequest struct {
	UserID string `json:"userId,omitempty"`
	TaskID string `json:"taskId,omitempty"`
	Isp    string `json:"isp,omitempty"`
}

type CreateTaskResponse struct {
	TaskID string `json:"taskId,omitempty"`
	JobID  string `json:"jobId,omitempty"`
}

type GetTaskListRequest struct {
	UserID   string `json:"userId,omitempty"`
	Query    string `json:"query,omitempty"`
	Type     string `json:"type,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

type GetTaskList struct {
	TaskID   string `json:"taskId,omitempty"`
	TaskName string `json:"taskName,omitempty"`
	Cycle    int    `json:"cycle,omitempty"`
	Type     string `json:"type,omitempty"`
	Address  string `json:"address,omitempty"`
	Status   string `json:"status,omitempty"`
}

type GetTaskListResponse struct {
	PageSize      int           `json:"pageSize,omitempty"`
	PageNumber    int           `json:"pageNumber,omitempty"`
	TotalElements int           `json:"totalElements,omitempty"`
	Content       []GetTaskList `json:"content,omitempty"`
}

type DeleteTaskResponse struct {
	RequestID string `json:"requestId,omitempty"`
	Code      string `json:"code,omitempty"`
	Message   string `json:"message,omitempty"`
}

type GetTaskDetailResponse struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
	Method  string `json:"method,omitempty"`
	SiteID  string `json:"siteId,omitempty"`
}

type CreateSiteAlarmConfigRequest struct {
	UserID              string          `json:"userId,omitempty"`
	TaskID              string          `json:"taskId,omitempty"`
	Comment             string          `json:"comment,omitempty"`
	AlarmName           string          `json:"alarmName,omitempty"`
	AliasName           string          `json:"aliasName,omitempty"`
	Namespace           string          `json:"namespace,omitempty"`
	Level               Level           `json:"level,omitempty"`
	ActionEnabled       string          `json:"actionEnabled,omitempty"`
	ResumeActions       []string        `json:"resumeActions,omitempty"`
	InsufficientActions []string        `json:"insufficientActions,omitempty"`
	IncidentAction      []string        `json:"incidentAction,omitempty"`
	InsufficientCycle   int             `json:"insufficientCycle,omitempty"`
	Rules               []SiteAlarmRule `json:"rules,omitempty"`
	Region              string          `json:"region,omitempty"`
	CallbackURL         string          `json:"callbackUrl,omitempty"`
	CallbackToken       string          `json:"callbackToken,omitempty"`
	Method              string          `json:"method,omitempty"`
	SiteMonitor         string          `json:"siteMonitor,omitempty"`
	Tag                 string          `json:"tag,omitempty"`
	Cycle               int             `json:"cycle,omitempty"`
}

type CreateSiteAlarmConfigResponse struct {
	UserID              string          `json:"userId,omitempty"`
	TaskID              string          `json:"taskId,omitempty"`
	Comment             string          `json:"comment,omitempty"`
	AlarmName           string          `json:"alarmName,omitempty"`
	AliasName           string          `json:"aliasName,omitempty"`
	Namespace           string          `json:"namespace,omitempty"`
	Level               Level           `json:"level,omitempty"`
	ActionEnabled       bool            `json:"actionEnabled,omitempty"`
	ResumeActions       []string        `json:"resumeActions,omitempty"`
	InsufficientActions []string        `json:"insufficientActions,omitempty"`
	IncidentAction      []string        `json:"incidentAction,omitempty"`
	InsufficientCycle   int             `json:"insufficientCycle,omitempty"`
	Rules               []SiteAlarmRule `json:"rules,omitempty"`
	Region              string          `json:"region,omitempty"`
	CallbackURL         string          `json:"callbackUrl,omitempty"`
	CallbackToken       string          `json:"callbackToken,omitempty"`
	Method              string          `json:"method,omitempty"`
	SiteMonitor         string          `json:"siteMonitor,omitempty"`
	Tag                 string          `json:"tag,omitempty"`
	Cycle               int             `json:"cycle,omitempty"`
}

type SiteAlarmRule struct {
	Metric             string   `json:"metric,omitempty"`
	MetricAlias        string   `json:"metricAlias,omitempty"`
	Statistics         string   `json:"statistics,omitempty"`
	Threshold          string   `json:"threshold,omitempty"`
	ComparisonOperator string   `json:"comparisonOperator,omitempty"`
	Cycle              int      `json:"cycle,omitempty"`
	Count              int      `json:"count,omitempty"`
	Function           string   `json:"function,omitempty"`
	ActOnIdcs          []string `json:"actOnIdcs,omitempty"`
	ActOnIsps          []string `json:"actOnIsps,omitempty"`
	VersionSite        string   `json:"versionSite,omitempty"`
}

type DeleteSiteAlarmConfigRequest struct {
	UserID     string   `json:"userId,omitempty"`
	AlarmNames []string `json:"alarmNames,omitempty"`
}

type GetSiteAlarmConfigRequest struct {
	UserID    string `json:"userId,omitempty"`
	AlarmName string `json:"alarmName,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

type GetSiteAlarmConfigListRequest struct {
	UserID        string `json:"userId,omitempty"`
	TaskID        string `json:"taskId,omitempty"`
	AliasName     string `json:"aliasName,omitempty"`
	ActionEnabled bool   `json:"actionEnabled,omitempty"`
	PageNo        int    `json:"pageNo,omitempty"`
	PageSize      int    `json:"pageSize,omitempty"`
}

type GetSiteAlarmConfigListResponse struct {
	OrderBy    string                          `json:"orderBy,omitempty"`
	Order      string                          `json:"order,omitempty"`
	PageNo     int                             `json:"pageNo,omitempty"`
	PageSize   int                             `json:"pageSize,omitempty"`
	TotalCount int                             `json:"totalCount,omitempty"`
	Result     []CreateSiteAlarmConfigResponse `json:"result,omitempty"`
}

type GetSiteMetricDataRequest struct {
	UserID     string   `json:"userId,omitempty"`
	MetricName string   `json:"metricName,omitempty"`
	Statistics []string `json:"statistics,omitempty"`
	StartTime  string   `json:"startTime,omitempty"`
	EndTime    string   `json:"endTime,omitempty"`
	Cycle      int      `json:"cycle,omitempty"`
	TaskID     string   `json:"taskId,omitempty"`
	Dimensions string   `json:"dimensions,omitempty"`
}

type GetSiteMetricDataResponse struct {
	Namespace  string       `json:"namespace,omitempty"`
	Dimensions []Dimension  `json:"dimensions,omitempty"`
	DataPoints []DataPoints `json:"dataPoints,omitempty"`
}

type GetSiteViewResponse struct {
	ID           string  `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Availability string  `json:"availability,omitempty"`
	ResponseTime float64 `json:"responseTime,omitempty"`
}

type GetSiteAgentListRequest struct {
	UserID string `json:"userId,omitempty"`
}

type GetSiteAgentListResponse struct {
	AgentID   string `json:"agentId,omitempty"`
	AgentName string `json:"agentName,omitempty"`
}

type GetSiteAgentsResponse struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type GetSiteAgentByTaskIDResponse struct {
	Idcs []GetSiteAgentsResponse `json:"idcs,omitempty"`
	Isps []GetSiteAgentsResponse `json:"isps,omitempty"`
}

package model

type MonitorObjectType string

const (
	MonitorObjectApp      MonitorObjectType = "APP"
	MonitorObjectService  MonitorObjectType = "SERVICE"
	MonitorObjectInstance MonitorObjectType = "INSTANCE"
	MonitorObjectCustom   MonitorObjectType = "CUSTOM"
	MonitorObjectDomain   MonitorObjectType = "DOMAIN"
)

type SrcType string

const (
	SrcTypeProc SrcType = "PROC"
	SrcTypeLog  SrcType = "LOG"
	SrcTypePort SrcType = "PORT"
	SrcTypeSCR  SrcType = "SCR"
)

type MonitorObject struct {
	Id                int64                    `json:"id,omitempty"`
	MonitorObjectView []MonitorObjectViewModel `json:"monitorObjectView,omitempty"`
	MonitorObjectType string                   `json:"monitorObjectType,omitempty"`
	TypeName          string                   `json:"typeName,omitempty"`
}

type MonitorObjectViewModel struct {
	MonitorObjectName     string `json:"monitorObjectName,omitempty"`
	MonitorObjectNameView string `json:"monitorObjectNameView"`
	MetricDimensionView   string `json:"metricDimensionView"`
}

type AppMonitorAlarmRule struct {
	Metric             string      `json:"metric,omitempty"`
	MetricAlias        string      `json:"metricAlias,omitempty"`
	Cycle              int         `json:"cycle,omitempty"`
	Statistics         string      `json:"statistics,omitempty"`
	Threshold          float64     `json:"threshold,omitempty"`
	ComparisonOperator string      `json:"comparisonOperator,omitempty"`
	Count              int         `json:"count,omitempty"`
	Function           string      `json:"function,omitempty"`
	MetricDimensions   []Dimension `json:"metricDimensions,omitempty"`
	Sequence           int         `json:"sequence,omitempty"`
	FormulaV2Alias     string      `json:"formulaV2Alias,omitempty"`
	MetricTags         string      `json:"metricTags,omitempty"`
}

type AppMonitorAlarmConfig struct {
	AlarmDescription    string                  `json:"alarmDescription"`
	AlarmName           string                  `json:"alarmName,omitempty"`
	UserId              string                  `json:"userId,omitempty"`
	AppName             string                  `json:"appName,omitempty"`
	MonitorObjectType   MonitorObjectType       `json:"monitorObjectType,omitempty"`
	MonitorObject       MonitorObject           `json:"monitorObject,omitempty"`
	SrcName             string                  `json:"srcName,omitempty"`
	SrcType             SrcType                 `json:"srcType,omitempty"`
	Type                AlarmType               `json:"type,omitempty"`
	Level               AlarmLevel              `json:"level,omitempty"`
	ActionEnabled       bool                    `json:"actionEnabled,omitempty"`
	PolicyEnabled       bool                    `json:"policyEnabled,omitempty"`
	Rules               [][]AppMonitorAlarmRule `json:"rules,omitempty"`
	IncidentActions     []string                `json:"incidentActions,omitempty"`
	ResumeAction        []string                `json:"resumeAction,omitempty"`
	InsufficientActions []string                `json:"insufficientActions,omitempty"`
	InsufficientCycle   int                     `json:"insufficientCycle,omitempty"`
	RepeatAlarmCycle    int                     `json:"repeatAlarmCycle,omitempty"`
	MaxRepeatCount      int                     `json:"maxRepeatCount,omitempty"`
}

type GetAppMonitorAlarmMetricsRequest struct {
	UserId     string `json:"userId,omitempty"`
	AppName    string `json:"appName,omitempty"`
	TaskName   string `json:"taskName,omitempty"`
	SearchName string `json:"searchName,omitempty"`
}

type DeleteAppMonitorAlarmConfigRequest struct {
	UserId    string `json:"userId,omitempty"`
	AppName   string `json:"appName,omitempty"`
	AlarmName string `json:"alarmName,omitempty"`
}

type ListAppMonitorAlarmConfigsRequest struct {
	UserId        string  `json:"userId,omitempty"`
	AppName       string  `json:"appName,omitempty"`
	AlarmName     string  `json:"alarmName,omitempty"`
	ActionEnabled *bool   `json:"actionEnabled,omitempty"`
	SrcType       SrcType `json:"srcType,omitempty"`
	TaskName      string  `json:"taskName,omitempty"`
	PageNo        int     `json:"pageNo,omitempty"`
	PageSize      int     `json:"pageSize,omitempty"`
}

type ListAppMonitorAlarmConfigsResponse struct {
	TotalCount int                     `json:"totalCount,omitempty"`
	OrderBy    string                  `json:"orderBy,omitempty"`
	Order      string                  `json:"order,omitempty"`
	PageNo     int                     `json:"pageNo,omitempty"`
	PageSize   int                     `json:"pageSize,omitempty"`
	Result     []AppMonitorAlarmConfig `json:"result,omitempty"`
}

type GetAppMonitorAlarmConfigDetailRequest struct {
	UserId    string `json:"userId,omitempty"`
	AlarmName string `json:"alarmName,omitempty"`
	AppName   string `json:"appName,omitempty"`
}

type LogExtractRequest struct {
	UserId      string `json:"userId,omitempty"`
	ExtractRule string `json:"extractRule,omitempty"`
	LogExample  string `json:"logExample,omitempty"`
}

package model

type AlarmConfig struct {
	AliasName           string              `json:"aliasName,omitempty"`
	AlarmName           string              `json:"alarmName,omitempty"`
	Level               AlarmLevel          `json:"level,omitempty"`
	MonitorObject       *AlarmMonitorObject `json:"monitorObject,omitempty"`
	IncidentActions     []string            `json:"alarmActions,omitempty"`
	ResumeAction        []string            `json:"okActions,omitempty"`
	InsufficientActions []string            `json:"insufficientActions,omitempty"`
	Region              string              `json:"region,omitempty"`
	Scope               string              `json:"scope,omitempty"`
	UserId              string              `json:"userId,omitempty"`
	Rules               [][]*AlarmRule      `json:"rules,omitempty"`
	SrcName             string              `json:"srcName,omitempty"`
	SrcType             string              `json:"srcType,omitempty"`
	AlarmType           string              `json:"type,omitempty"`
	AlarmDescription    string              `json:"alarmDescription"`
	EventTypeList       []string            `json:"eventTypeList,omitempty"`
	InsufficientCycle   int                 `json:"insufficientCycle,omitempty"`
	RepeatAlarmCycle    int                 `json:"repeatAlarmCycle,omitempty"`
	MaxRepeatCount      int                 `json:"maxRepeatCount,omitempty"`
	CallbackURL         string              `json:"callbackUrl,omitempty"`
	CallbackToken       string              `json:"callbackToken,omitempty"`
}

type AlarmConfigV2 struct {
	UserId                        string           `json:"userId,omitempty"`
	AlarmName                     string           `json:"alarmName,omitempty"`
	AliasName                     string           `json:"aliasName,omitempty"`
	Region                        string           `json:"region,omitempty"`
	Scope                         string           `json:"scope,omitempty"`
	InsufficientDataPendingPeriod int              `json:"insufficientDataPendingPeriod,omitempty"`
	AlarmRepeatInterval           int              `json:"alarmRepeatInterval,omitempty"`
	AlarmRepeatCount              int              `json:"alarmRepeatCount,omitempty"`
	ResourceType                  string           `json:"resourceType,omitempty"`
	AlarmLevel                    AlarmLevel       `json:"alarmLevel,omitempty"`
	TargetType                    TargetType       `json:"targetType,omitempty"`
	TargetInstanceGroups          []string         `json:"targetInstanceGroups,omitempty"`
	TargetInstances               []*AlarmInstance `json:"targetInstances,omitempty"`
	TargetInstanceTags            []*KV            `json:"targetInstanceTags,omitempty"`
	Policies                      []*AlarmPolicy   `json:"policies,omitempty"`
	Actions                       []*AlarmAction   `json:"actions,omitempty"`
}

type AlarmPolicy struct {
	AlarmPendingPeriodCount int            `json:"alarmPendingPeriodCount,omitempty"`
	Rules                   []*AlarmRuleV2 `json:"rules,omitempty"`
}

type AlarmAction struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

type AlarmInstance struct {
	Region           string `json:"region,omitempty"`
	Identifiers      []*KV  `json:"identifiers,omitempty"`
	MetricDimensions []*KV  `json:"metricDimensions,omitempty"`
}

type AlarmRule struct {
	Index                 int          `json:"index,omitempty"`
	Metric                string       `json:"metric,omitempty"`
	PeriodInSecond        int          `json:"periodInSecond,omitempty"`
	Statistics            string       `json:"statistics,omitempty"`
	Threshold             string       `json:"threshold,omitempty"`
	ComparisonOperator    string       `json:"comparisonOperator,omitempty"`
	EvaluationPeriodCount int          `json:"evaluationPeriodCount,omitempty"`
	MetricDimensions      []*Dimension `json:"metricDimensions,omitempty"`
}

type AlarmRuleV2 struct {
	MetricName       string  `json:"metricName,omitempty"`
	MetricDimensions []*KV   `json:"metricDimensions,omitempty"`
	Operator         string  `json:"operator,omitempty"`
	Statistics       string  `json:"statistics,omitempty"`
	Threshold        float64 `json:"threshold,omitempty"`
	Window           int     `json:"window,omitempty"`
}

type AlarmMetric struct {
	Alias            string     `json:"alias,omitempty"`
	Name             string     `json:"name,omitempty"`
	UnitCategory     string     `json:"unitCategory,omitempty"`
	UnitName         string     `json:"unitName,omitempty"`
	Cycle            int        `json:"cycle,omitempty"`
	Scope            string     `json:"scope,omitempty"`
	TypeName         string     `json:"typeName,omitempty"`
	MetricDimensions [][]string `json:"metricDimensions,omitempty"`
}

type AlarmMonitorObject struct {
	MonitorType MonitorObjectType `json:"type,omitempty"`
	Names       []string          `json:"names,omitempty"`
	Resources   []*PolicyResource `json:"resources,omitempty"`
	TypeName    string            `json:"typeName,omitempty"`
}

type PolicyResource struct {
	Identifiers      []*Dimension `json:"identifiers,omitempty"`
	MetricDimensions []*Dimension `json:"metricDimensions,omitempty"`
}

type CommonAlarmConfigRequest struct {
	UserId    string `json:"userId,omitempty"`
	Scope     string `json:"scope,omitempty"`
	AlarmName string `json:"alarmName,omitempty"`
}

type ListSingleInstanceAlarmConfigsRequest struct {
	UserId          string `json:"userId,omitempty"`
	Scope           string `json:"scope,omitempty"`
	Region          string `json:"region,omitempty"`
	Dimensions      string `json:"dimensions,omitempty"`
	Order           string `json:"order,omitempty"`
	PageNo          int    `json:"pageNo,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	ActionEnabled   *bool  `json:"actionEnabled,omitempty"`
	AlarmNamePrefix string `json:"alarmNamePrefix,omitempty"`
}

type ListSingleInstanceAlarmConfigsResponse struct {
	TotalCount int            `json:"totalCount,omitempty"`
	OrderBy    string         `json:"orderBy,omitempty"`
	Order      string         `json:"order,omitempty"`
	PageNo     int            `json:"pageNo,omitempty"`
	PageSize   int            `json:"pageSize,omitempty"`
	Result     []*AlarmConfig `json:"result,omitempty"`
}

type ListAlarmMetricsRequest struct {
	UserId     string `json:"userId,omitempty"`
	Scope      string `json:"scope,omitempty"`
	Region     string `json:"region,omitempty"`
	Dimensions string `json:"dimensions,omitempty"`
	Type       string `json:"type,omitempty"`
	Locale     string `json:"locale,omitempty"`
}

type CreateAlarmPolicyV2Response struct {
	Success bool           `json:"success,omitempty"`
	Msg     string         `json:"msg,omitempty"`
	Result  *AlarmConfigV2 `json:"result,omitempty"`
}

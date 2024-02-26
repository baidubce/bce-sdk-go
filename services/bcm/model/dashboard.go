package model

type DashboardBaseResponse struct {
	Code      int                    `json:"code,omitempty"`
	Data      map[string]interface{} `json:"data,omitempty"`
	Message   string                 `json:"message,omitempty"`
	TraceInfo string                 `json:"traceInfo,omitempty"`
	Success   bool                   `json:"success,omitempty"`
}

type DashboardRequest struct {
	DashboardName string `json:"dashboardName,omitempty"`
	UserID        string `json:"userId,omitempty"`
	Title         string `json:"title,omitempty"`
	Configure     string `json:"configure,omitempty"`
	WidgetName    string `json:"widgetName,omitempty"`
	Type          string `json:"type,omitempty"`
}

type DashboardResponse struct {
	Code      int    `json:"code,omitempty"`
	Data      string `json:"data,omitempty"`
	Message   string `json:"message,omitempty"`
	TraceInfo string `json:"traceInfo,omitempty"`
	Success   bool   `json:"success,omitempty"`
}

type DashboardWidgetRequest struct {
	DashboardName string                   `json:"dashboardName,omitempty"`
	UserID        string                   `json:"userId,omitempty"`
	Title         string                   `json:"title,omitempty"`
	Configure     DashboardWidgetConfigure `json:"configure,omitempty"`
	WidgetName    string                   `json:"widgetName,omitempty"`
	Type          string                   `json:"type,omitempty"`
}

type DashboardWidgetConfigure struct {
	Data        []Data    `json:"data,omitempty"`
	Style       Style     `json:"style,omitempty"`
	Title       string    `json:"title,omitempty"`
	TimeRange   TimeRange `json:"timeRange,omitempty"`
	Time        string    `json:"time,omitempty"`
	MonitorType string    `json:"monitorType,omitempty"`
}

type DashboardMonitorObject struct {
	InstanceName string `json:"instanceName,omitempty"`
	Id           string `json:"id,omitempty"`
}

type Data struct {
	Metric        []DashboardMetric        `json:"metric,omitempty"`
	MonitorObject []DashboardMonitorObject `json:"monitorObject,omitempty"`
	Scope         string                   `json:"scope,omitempty"`
	SubService    string                   `json:"subService,omitempty"`
	Region        string                   `json:"region,omitempty"`
	ScopeValue    ScopeValue               `json:"scopeValue,omitempty"`
	ResourceType  string                   `json:"resourceType,omitempty"`
	MonitorType   string                   `json:"monitorType,omitempty"`
	Namespace     []DashboardNamespace     `json:"namespace,omitempty"`
	Product       string                   `json:"product,omitempty"`
}

type DashboardMetric struct {
	Name             string             `json:"name,omitempty"`
	Unit             string             `json:"unit,omitempty"`
	Alias            string             `json:"alias,omitempty"`
	Contrast         []string           `json:"contrast,omitempty"`
	TimeContrast     []string           `json:"timeContrast,omitempty"`
	Statistics       string             `json:"statistics,omitempty"`
	Dimensions       []string           `json:"dimensions,omitempty"`
	MetricDimensions []MetricDimensions `json:"metricDimensions,omitempty"`
	Cycle            int                `json:"cycle,omitempty"`
	DisplayName      string             `json:"displayName,omitempty"`
}

type MetricDimensions struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"values,omitempty"`
}

type ScopeValue struct {
	Name        string `json:"name,omitempty"`
	Value       string `json:"value,omitempty"`
	HasChildren bool   `json:"hasChildren,omitempty"`
}

type DashboardNamespace struct {
	NamespaceType string       `json:"namespaceType,omitempty"`
	Transfer      string       `json:"transfer"`
	Filter        string       `json:"filter,omitempty"`
	Name          string       `json:"name,omitempty"`
	InstanceName  string       `json:"instanceName,omitempty"`
	Region        string       `json:"region,omitempty"`
	BcmService    string       `json:"bcmService,omitempty"`
	SubService    []SubService `json:"subService,omitempty"`
}

type TimeRange struct {
	TimeType string `json:"timeType,omitempty"`
	Unit     string `json:"unit,omitempty"`
	Number   int    `json:"number,omitempty"`
	Relative string `json:"relative,omitempty"`
}

type Style struct {
	DisplayType   string `json:"displayType,omitempty"`
	NullPointMode string `json:"nullPointMode,omitempty"`
	Threshold     int    `json:"threshold,omitempty"`
	Decimals      int    `json:"decimals,omitempty"`
	IsEdit        bool   `json:"isEdit,omitempty"`
	Unit          string `json:"unit,omitempty"`
}

type SubService struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type DashboardDataRequest struct {
	Data []Data `json:"data,omitempty"`
	Time string `json:"time,omitempty"`
}

type DashboardTrendData struct {
	Data                     [][]float64 `json:"data,omitempty"`
	Denominator              int         `json:"denominator,omitempty"`
	Dimensions               string      `json:"dimensions,omitempty"`
	Legend                   interface{} `json:"legend,omitempty"`
	Metric                   string      `json:"metric,omitempty"`
	MetricType               string      `json:"metricType,omitempty"`
	MetricUnit               string      `json:"metricUnit,omitempty"`
	MetricUnitTransformation string      `json:"metricUnitTransformation,omitempty"`
	Name                     string      `json:"name,omitempty"`
	Namespace                string      `json:"namespace,omitempty"`
	Numerator                int         `json:"numerator,omitempty"`
	Product                  interface{} `json:"product,omitempty"`
	Scope                    string      `json:"scope,omitempty"`
	Statistics               interface{} `json:"statistics,omitempty"`
	Time                     interface{} `json:"time,omitempty"`
	TransPolicy              string      `json:"transPolicy,omitempty"`
	HostName                 interface{} `json:"hostName,omitempty"`
	InternalIP               interface{} `json:"internalIp,omitempty"`
}

type DashboardTrendResponse struct {
	Code      int                  `json:"code,omitempty"`
	Data      []DashboardTrendData `json:"data,omitempty"`
	Message   string               `json:"message,omitempty"`
	TraceInfo string               `json:"traceInfo,omitempty"`
	Success   bool                 `json:"success,omitempty"`
}

type DashboardReportData struct {
	Alias    string             `json:"alias,omitempty"`
	Children []string           `json:"children,omitempty"`
	Metrics  map[string]float64 `json:"metrics,omitempty"`
	Name     string             `json:"name,omitempty"`
	Value    string             `json:"value,omitempty"`
}

type DashboardReportDataResponse struct {
	Code      int                   `json:"code,omitempty"`
	Data      []DashboardReportData `json:"data,omitempty"`
	Message   string                `json:"message,omitempty"`
	TraceInfo interface{}           `json:"traceInfo,omitempty"`
	Success   bool                  `json:"success,omitempty"`
}

type DashboardBillboardResponse struct {
	Code      int                      `json:"code,omitempty"`
	Data      []DashboardBillboardData `json:"data,omitempty"`
	Message   string                   `json:"message,omitempty"`
	TraceInfo string                   `json:"traceInfo,omitempty"`
	Success   bool                     `json:"success,omitempty"`
}

type DashboardBillboardData struct {
	Data            [][]float64 `json:"data,omitempty"`
	Decimals        interface{} `json:"decimals,omitempty"`
	DisplayName     string      `json:"displayName,omitempty"`
	InstanceName    string      `json:"instanceName,omitempty"`
	MetricDimension string      `json:"metricDimension,omitempty"`
	Name            string      `json:"name,omitempty"`
	Unit            string      `json:"unit,omitempty"`
}

type DashboardTrendSeniorResponse struct {
	Code      int                        `json:"code,omitempty"`
	Data      []DashboardSeniorTrendData `json:"data,omitempty"`
	Message   string                     `json:"message,omitempty"`
	TraceInfo string                     `json:"traceInfo,omitempty"`
	Success   bool                       `json:"success,omitempty"`
}

type DashboardSeniorTrendData struct {
	Items   []Item  `json:"items,omitempty"`
	Job     Job     `json:"job,omitempty"`
	Numeric Numeric `json:"numeric,omitempty"`
}

type Item struct {
	StatisticsValue Numeric `json:"statisticsValue,omitempty"`
	Timestamp       int     `json:"timestamp,omitempty"`
}

type TagForTsdb struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Job struct {
	Alias          string       `json:"alias,omitempty"`
	BcmSource      bool         `json:"bcmSource,omitempty"`
	Contrast       string       `json:"contrast,omitempty"`
	Decimals       interface{}  `json:"decimals,omitempty"`
	DisplayName    string       `json:"displayName,omitempty"`
	EndTime        int          `json:"endTime,omitempty"`
	Flatten        bool         `json:"flatten,omitempty"`
	InstanceID     interface{}  `json:"instanceId,omitempty"`
	InstanceName   string       `json:"instanceName,omitempty"`
	IntranetIP     interface{}  `json:"intranetIp,omitempty"`
	Items          []Item       `json:"items,omitempty"`
	MetricName     string       `json:"metricName,omitempty"`
	Namespace      string       `json:"namespace,omitempty"`
	Offset         int          `json:"offset,omitempty"`
	OriginalPeriod int          `json:"originalPeriod,omitempty"`
	Period         int          `json:"period,omitempty"`
	Product        string       `json:"product,omitempty"`
	StartTime      int          `json:"startTime,omitempty"`
	Statistics     string       `json:"statistics,omitempty"`
	Tags           string       `json:"tags,omitempty"`
	TagsForTsdb    []TagForTsdb `json:"tagsForTsdb,omitempty"`
	Unit           string       `json:"unit,omitempty"`
}

type Numeric struct {
	Avg float64 `json:"avg,omitempty"`
	Cnt float64 `json:"cnt,omitempty"`
	Max float64 `json:"max,omitempty"`
	Min float64 `json:"min,omitempty"`
	Sum float64 `json:"sum,omitempty"`
}

type DashboardDimension struct {
	Namespace string   `json:"namespace,omitempty"`
	Metrics   []string `json:"metrics,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

type DimensionFilter struct {
	Name   string   `json:"name,omitempty"`
	Values []string `json:"value,omitempty"`
}

type DashboardDimensionsRequest struct {
	UserID     string `json:"userId,omitempty"`
	Service    string `json:"service,omitempty"`
	Region     string `json:"region,omitempty"`
	ResourceID string `json:"showId,omitempty"`
	Dimensions string `json:"dimensions,omitempty"`
	MetricName string `json:"metricName,omitempty"`
}

type DimensionValue struct {
	Name      string `json:"name,omitempty"`
	Comment   string `json:"comment,omitempty"`
	Available bool   `json:"available,omitempty"`
}

type DashboardDimensionsData struct {
	Name            string           `json:"name,omitempty"`
	Product         string           `json:"product,omitempty"`
	Comment         string           `json:"comment,omitempty"`
	DimensionValues []DimensionValue `json:"dimensionValues,omitempty"`
}

type DashboardDimensionsResponse struct {
	Code      int                       `json:"code,omitempty"`
	Data      []DashboardDimensionsData `json:"data,omitempty"`
	Message   string                    `json:"message,omitempty"`
	TraceInfo string                    `json:"traceInfo,omitempty"`
	Success   bool                      `json:"success,omitempty"`
}

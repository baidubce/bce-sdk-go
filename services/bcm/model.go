package bcm

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
	Average     float64 `json:"average,omitempty"`
	Sum         float64 `json:"sum,omitempty"`
	Minimum     float64 `json:"minimum,omitempty"`
	Maximum     float64 `json:"maximum,omitempty"`
	SampleCount int64   `json:"sampleCount,omitempty"`
	Timestamp   string  `json:"timestamp,omitempty"`
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

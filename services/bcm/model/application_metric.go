package model

type AppAggrTag struct {
	Range string `json:"range,omitempty"`
	Tags  string `json:"tags,omitempty"`
}

type AppMetric struct {
	Id             int64        `json:"id,omitempty"`
	TaskId         int64        `json:"taskId,omitempty"`
	MetricName     string       `json:"metricName,omitempty"`
	MetricAlias    string       `json:"metricAlias,omitempty"`
	MetricUnit     string       `json:"metricUnit,omitempty"`
	ValueFieldType int          `json:"valueFieldType,omitempty"`
	ValueFieldName string       `json:"valueFieldName,omitempty"`
	ValueMatchRule string       `json:"valueMatchRule,omitempty"`
	AggrTags       []AppAggrTag `json:"aggrTags,omitempty"`

	SaveInstanceData int `json:"saveInstanceData,omitempty"`
}

type GetMetricMetaForApplicationRequest struct {
	UserId        string   `json:"userId,omitempty"`
	AppName       string   `json:"appName,omitempty"`
	TaskName      string   `json:"taskName,omitempty"`
	MetricName    string   `json:"metricName,omitempty"`
	Instances     []string `json:"instances,omitempty"`
	DimensionKeys []string `json:"dimensionKeys,omitempty"`
}

type GetMetricDataForApplicationRequest struct {
	UserId     string              `json:"userId,omitempty"`
	AppName    string              `json:"appName,omitempty"`
	TaskName   string              `json:"taskName,omitempty"`
	MetricName string              `json:"metricName,omitempty"`
	Instances  []string            `json:"instances,omitempty"`
	StartTime  string              `json:"startTime,omitempty"`
	EndTime    string              `json:"endTime,omitempty"`
	Cycle      int                 `json:"cycle,omitempty"`
	Statistics []string            `json:"statistics,omitempty"`
	Dimensions map[string][]string `json:"dimensions,omitempty"`
	AggrData   bool                `json:"aggrData,omitempty"`
}

type MetricDataDimension struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type GetMetricDataForApplicationResult struct {
	Namespace  string                `json:"namespace,omitempty"`
	Dimensions []MetricDataDimension `json:"dimensions,omitempty"`
	DataPoints []DataPoints          `json:"dataPoints"`
}

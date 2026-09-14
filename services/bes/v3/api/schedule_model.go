package api

// ScheduleItem describes a single scheduled task as returned by the list API.
type ScheduleItem struct {
	ScheduleId       string                 `json:"scheduleId,omitempty"`
	ScheduleName     string                 `json:"scheduleName,omitempty"`
	ScheduleType     string                 `json:"scheduleType,omitempty"`
	CronExpr         string                 `json:"cronExpr,omitempty"`
	Enabled          bool                   `json:"enabled,omitempty"`
	CreateTime       string                 `json:"createTime,omitempty"`
	LastTriggerTime  string                 `json:"lastTriggerTime,omitempty"`
	LastStatus       string                 `json:"lastStatus,omitempty"`
	LastErrorMessage string                 `json:"lastErrorMessage,omitempty"`
	ExecutedCount    int                    `json:"executedCount,omitempty"`
	FailedCount      int                    `json:"failedCount,omitempty"`
	Params           map[string]interface{} `json:"params,omitempty"`
}

// ListSchedulesRequest contains the path and query parameters for
// GET /v3/clusters/{clusterId}/schedules.
type ListSchedulesRequest struct {
	Request
	ClusterId string `json:"-"`
	PageNo    int    `json:"-"`
	PageSize  int    `json:"-"`
}

// ListSchedulesResponse contains the response body for
// GET /v3/clusters/{clusterId}/schedules.
type ListSchedulesResponse struct {
	Schedules  []ScheduleItem `json:"schedules,omitempty"`
	PageNo     int            `json:"pageNo,omitempty"`
	PageSize   int            `json:"pageSize,omitempty"`
	TotalCount int            `json:"totalCount,omitempty"`
}

// CreateScheduleRequest contains the request body for POST /v3/clusters/{clusterId}/schedules.
type CreateScheduleRequest struct {
	Request
	ClusterId    string                 `json:"-"`
	ScheduleName string                 `json:"scheduleName"`
	ScheduleType string                 `json:"scheduleType"`
	CronExpr     string                 `json:"cronExpr"`
	Params       map[string]interface{} `json:"params"`
}

// CreateScheduleResponse contains the response body for
// POST /v3/clusters/{clusterId}/schedules.
type CreateScheduleResponse struct {
	Success bool `json:"success,omitempty"`
}

// UpdateScheduleRequest contains the request body for
// PUT /v3/clusters/{clusterId}/schedules/{scheduleId}.
type UpdateScheduleRequest struct {
	Request
	ClusterId  string                 `json:"-"`
	ScheduleId string                 `json:"-"`
	CronExpr   string                 `json:"cronExpr"`
	Params     map[string]interface{} `json:"params"`
}

// UpdateScheduleResponse contains the response body for
// PUT /v3/clusters/{clusterId}/schedules/{scheduleId}.
type UpdateScheduleResponse struct {
	Success bool `json:"success,omitempty"`
}

// DeleteScheduleRequest contains the path parameters for
// DELETE /v3/clusters/{clusterId}/schedules/{scheduleId}.
type DeleteScheduleRequest struct {
	Request
	ClusterId  string `json:"-"`
	ScheduleId string `json:"-"`
}

// DeleteScheduleResponse contains the response body for
// DELETE /v3/clusters/{clusterId}/schedules/{scheduleId}.
type DeleteScheduleResponse struct {
	Success bool `json:"success,omitempty"`
}

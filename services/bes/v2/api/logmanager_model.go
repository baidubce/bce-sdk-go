package api

// UpdateLogSettingsRequest is the request of toggling a cluster's slow-log/audit-log switches.
type UpdateLogSettingsRequest struct {
	ClusterId     string `json:"clusterId"`
	IndexSlowLog  bool   `json:"indexSlowLog"`
	SearchSlowLog bool   `json:"searchSlowLog"`
	AuditLog      bool   `json:"auditLog"`
}

// UpdateLogSettingsResponse is the response of toggling a cluster's slow-log/audit-log switches.
type UpdateLogSettingsResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// SearchLogRequest is the request of searching a cluster's logs.
type SearchLogRequest struct {
	ClusterId string `json:"clusterId"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
	Keyword   string `json:"keyword,omitempty"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Type      string `json:"type"`
}

// SearchLogResponse is the response of searching a cluster's logs.
type SearchLogResponse struct {
	Success bool     `json:"success"`
	Status  int      `json:"status"`
	Page    *LogPage `json:"page"`
}

// LogPage is the paged log entry list carried by SearchLogResponse.
type LogPage struct {
	PageNo     int        `json:"pageNo"`
	PageSize   int        `json:"pageSize"`
	TotalCount int        `json:"totalCount"`
	OrderBy    string     `json:"orderBy"`
	Order      string     `json:"order"`
	Result     []LogEntry `json:"result"`
}

// LogEntry is a single log record inside LogPage.
type LogEntry struct {
	Level   string `json:"level"`
	Host    string `json:"host"`
	Time    string `json:"time"`
	Content string `json:"content"`
}

// CreateLogExportTaskRequest is the request of creating a log export task.
type CreateLogExportTaskRequest struct {
	ClusterId string   `json:"clusterId,omitempty"`
	StartTime string   `json:"startTime,omitempty"`
	EndTime   string   `json:"endTime,omitempty"`
	Types     []string `json:"types,omitempty"`
	BosBucket string   `json:"bosBucket,omitempty"`
	BosPath   string   `json:"bosPath,omitempty"`
}

// CreateLogExportTaskResponse is the response of creating a log export task.
type CreateLogExportTaskResponse struct {
	Success bool                  `json:"success"`
	Status  int                   `json:"status"`
	Result  *LogExportTasksResult `json:"result"`
}

// LogExportTasksResult is the result envelope carrying the created log export tasks.
type LogExportTasksResult struct {
	Tasks []LogExportTask `json:"tasks"`
}

// LogExportTask describes a single log export task created for one log type.
type LogExportTask struct {
	LogType string `json:"logType"`
	TaskId  string `json:"taskId"`
}

// GetLogExportRecordRequest is the request of getting a log export task's record.
type GetLogExportRecordRequest struct {
	ClusterId string `json:"clusterId,omitempty"`
	TaskId    string `json:"taskId,omitempty"`
}

// GetLogExportRecordResponse is the response of getting a log export task's record.
type GetLogExportRecordResponse struct {
	Success bool                   `json:"success"`
	Status  int                    `json:"status"`
	Result  *LogExportRecordResult `json:"result"`
}

// LogExportRecordResult is the result envelope carrying a log export task's progress.
type LogExportRecordResult struct {
	Progress int                  `json:"progress"`
	State    string               `json:"state"`
	Config   *LogExportTaskConfig `json:"config,omitempty"`
}

// LogExportTaskConfig describes the configuration of a log export task.
type LogExportTaskConfig struct {
	LogType   string `json:"logType"`
	BosBucket string `json:"bosBucket"`
	BosPath   string `json:"bosPath"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	Region    string `json:"region"`
}

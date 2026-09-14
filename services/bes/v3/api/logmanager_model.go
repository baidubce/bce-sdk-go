package api

// GetLogSwitchRequest contains the path variable for GET /v3/clusters/{clusterId}/logs/switch.
type GetLogSwitchRequest struct {
	Request
	ClusterId string `json:"-"`
}

// UpdateLogSwitchRequest contains the path variable and body for PUT /v3/clusters/{clusterId}/logs/switch.
type UpdateLogSwitchRequest struct {
	Request
	ClusterId string          `json:"-"`
	Enabled   *bool           `json:"enabled,omitempty"`
	Types     map[string]bool `json:"types,omitempty"`
}

// LogSwitchResponse contains the log collection switch state of a cluster.
type LogSwitchResponse struct {
	Enabled bool            `json:"enabled,omitempty"`
	Types   map[string]bool `json:"types,omitempty"`
}

// GetLogUsageRequest contains the path variable for GET /v3/clusters/{clusterId}/logs/usage.
type GetLogUsageRequest struct {
	Request
	ClusterId string `json:"-"`
}

// LogTypeUsage contains usage details for one log type.
type LogTypeUsage struct {
	SizeBytes int64 `json:"sizeBytes,omitempty"`
	Shards    int   `json:"shards,omitempty"`
	DocCount  int64 `json:"docCount,omitempty"`
}

// LogUsageResponse contains aggregated log index usage for a cluster.
type LogUsageResponse struct {
	TotalSizeBytes int64                   `json:"totalSizeBytes,omitempty"`
	TotalShards    int                     `json:"totalShards,omitempty"`
	Indices        map[string]LogTypeUsage `json:"indices,omitempty"`
}

// GetLogCollectorStatusRequest contains the path variable for GET /v3/clusters/{clusterId}/logs/collector/status.
type GetLogCollectorStatusRequest struct {
	Request
	ClusterId string `json:"-"`
}

// LogCollectorStatusResponse contains the log collector deployment and operation state.
type LogCollectorStatusResponse struct {
	Deployed        bool   `json:"deployed,omitempty"`
	OperationStatus string `json:"operationStatus,omitempty"`
	OperationType   string `json:"operationType,omitempty"`
	OperationId     string `json:"operationId,omitempty"`
}

// SearchLogsRequest contains the path variable and body for POST /v3/clusters/{clusterId}/logs/search.
type SearchLogsRequest struct {
	Request
	ClusterId string `json:"-"`
	LogType   string `json:"logType"`
	Keyword   string `json:"keyword,omitempty"`
	Level     string `json:"level,omitempty"`
	NodeName  string `json:"nodeName,omitempty"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Page      int    `json:"page,omitempty"`
	PageSize  int    `json:"pageSize,omitempty"`
}

// SearchLogsResponse contains a page of raw log records.
type SearchLogsResponse struct {
	Total    int64                    `json:"total,omitempty"`
	Page     int                      `json:"page,omitempty"`
	PageSize int                      `json:"pageSize,omitempty"`
	Logs     []map[string]interface{} `json:"logs,omitempty"`
}

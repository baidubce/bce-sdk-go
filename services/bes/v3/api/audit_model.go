package api

// GetAuditStatusRequest contains the path variable for GET /v3/clusters/{clusterId}/audit.
type GetAuditStatusRequest struct {
	Request
	ClusterId string `json:"-"`
}

// AuditStatusResponse contains the response body for GET /v3/clusters/{clusterId}/audit.
type AuditStatusResponse struct {
	Enabled bool `json:"enabled,omitempty"`
}

// EnableAuditRequest contains the request body for PUT /v3/clusters/{clusterId}/audit.
type EnableAuditRequest struct {
	Request
	ClusterId       string `json:"-"`
	TargetClusterId string `json:"targetClusterId,omitempty"`
}

// EnableAuditResponse contains the response body for PUT /v3/clusters/{clusterId}/audit.
type EnableAuditResponse struct {
	Enabled         bool   `json:"enabled,omitempty"`
	TargetClusterId string `json:"targetClusterId,omitempty"`
}

// ListAuditEventTypesRequest contains the path variable for GET /v3/clusters/{clusterId}/audit/event-types.
type ListAuditEventTypesRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ListAuditEventTypesResponse contains the response body for GET /v3/clusters/{clusterId}/audit/event-types.
type ListAuditEventTypesResponse struct {
	EventTypes []string `json:"eventTypes,omitempty"`
}

// SearchAuditEventsRequest contains the path variable and query parameters for
// GET /v3/clusters/{clusterId}/audit/events.
type SearchAuditEventsRequest struct {
	Request
	ClusterId string `json:"-"`
	StartTime string `json:"-"`
	EndTime   string `json:"-"`
	Category  string `json:"-"`
	IndexName string `json:"-"`
	PageSize  int    `json:"-"`
	Cursor    string `json:"-"`
}

// SearchAuditEventsResponse contains the response body for GET /v3/clusters/{clusterId}/audit/events.
type SearchAuditEventsResponse struct {
	Logs   []AuditEvent `json:"logs,omitempty"`
	Cursor string       `json:"cursor,omitempty"`
}

// GetBctAuthSwitchResponse contains the response body for GET /v3/clusters/bct/auth/switch.
type GetBctAuthSwitchResponse struct {
	Enabled bool `json:"enabled,omitempty"`
}

// UpdateBctAuthSwitchRequest contains the request body for PUT /v3/clusters/bct/auth/switch.
type UpdateBctAuthSwitchRequest struct {
	Request
	Enabled *bool `json:"enabled"`
}

// BctAuthSwitchResponse contains the response body for PUT /v3/clusters/bct/auth/switch.
type BctAuthSwitchResponse struct {
	Enabled bool `json:"enabled,omitempty"`
}

// SearchBctEventsRequest contains the request body for POST /v3/bct/events/query.
type SearchBctEventsRequest struct {
	Request
	StartTime  string           `json:"startTime"`
	EndTime    string           `json:"endTime"`
	PageSize   int              `json:"pageSize,omitempty"`
	DomainId   string           `json:"domainId,omitempty"`
	NextMarker string           `json:"nextMarker,omitempty"`
	Filters    []BctQueryFilter `json:"filters,omitempty"`
}

// SearchBctEventsResponse contains the response body for POST /v3/bct/events/query.
type SearchBctEventsResponse struct {
	PageSize   int        `json:"pageSize,omitempty"`
	Data       []BctEvent `json:"data,omitempty"`
	NextMarker string     `json:"nextMarker,omitempty"`
	Truncated  bool       `json:"truncated,omitempty"`
}

// BctQueryFilter defines a query filter for searching BCT events.
type BctQueryFilter struct {
	Field string `json:"field,omitempty"`
	Value string `json:"value,omitempty"`
}

// BctEvent describes a BCT (cloud audit) event.
type BctEvent struct {
	EventType               string           `json:"eventType,omitempty"`
	EventSource             string           `json:"eventSource,omitempty"`
	EventName               string           `json:"eventName,omitempty"`
	EventTimeInMilliseconds int64            `json:"eventTimeInMilliseconds,omitempty"`
	EventTime               string           `json:"eventTime,omitempty"`
	UserIpAddress           string           `json:"userIpAddress,omitempty"`
	UserAgent               string           `json:"userAgent,omitempty"`
	RegionId                string           `json:"regionId,omitempty"`
	RequestId               string           `json:"requestId,omitempty"`
	OrderId                 string           `json:"orderId,omitempty"`
	ApiVersion              string           `json:"apiVersion,omitempty"`
	Description             string           `json:"description,omitempty"`
	ErrorCode               string           `json:"errorCode,omitempty"`
	ErrorMessage            string           `json:"errorMessage,omitempty"`
	Success                 bool             `json:"success,omitempty"`
	UserIdentity            *BctUserIdentity `json:"userIdentity,omitempty"`
}

// BctUserIdentity describes the user identity of a BCT event.
type BctUserIdentity struct {
	IamDomainId     string `json:"iamDomainId,omitempty"`
	IamUserId       string `json:"iamUserId,omitempty"`
	LoginUserId     string `json:"loginUserId,omitempty"`
	UserDisplayName string `json:"userDisplayName,omitempty"`
	Accesskey       string `json:"accesskey,omitempty"`
}

// AuditEvent describes a data-plane audit event.
type AuditEvent struct {
	Category  string `json:"category,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Content   string `json:"content,omitempty"`
}

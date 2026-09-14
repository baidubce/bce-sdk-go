package api

// WeeklyRiskItem describes a risk item found within the last 7 days.
type WeeklyRiskItem struct {
	Id            string `json:"id,omitempty"`
	HighRiskCount int    `json:"highRiskCount,omitempty"`
	LowRiskCount  int    `json:"lowRiskCount,omitempty"`
}

// ListWeeklyOverviewRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/diagnosis/overview/weekly.
type ListWeeklyOverviewRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ListWeeklyOverviewResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/overview/weekly.
type ListWeeklyOverviewResponse struct {
	Count     int              `json:"count,omitempty"`
	RiskItems []WeeklyRiskItem `json:"riskItems,omitempty"`
}

// GetManualDiagnosisCountRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/diagnosis/manual/count.
type GetManualDiagnosisCountRequest struct {
	Request
	ClusterId string `json:"-"`
}

// GetManualDiagnosisCountResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/manual/count.
type GetManualDiagnosisCountResponse struct {
	Count int `json:"count,omitempty"`
}

// GetManualDiagnosisConfigRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/diagnosis/manual.
type GetManualDiagnosisConfigRequest struct {
	Request
	ClusterId string `json:"-"`
}

// GetManualDiagnosisConfigResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/manual.
type GetManualDiagnosisConfigResponse struct {
	Indices string   `json:"indices,omitempty"`
	Items   []string `json:"items,omitempty"`
}

// UpdateManualDiagnosisConfigRequest contains the request body for
// PUT /v3/clusters/{clusterId}/diagnosis/manual.
type UpdateManualDiagnosisConfigRequest struct {
	Request
	ClusterId string   `json:"-"`
	Indices   string   `json:"indices,omitempty"`
	Items     []string `json:"items"`
}

// UpdateManualDiagnosisConfigResponse contains the response body for
// PUT /v3/clusters/{clusterId}/diagnosis/manual.
type UpdateManualDiagnosisConfigResponse struct {
	Success bool `json:"success,omitempty"`
}

// CreateManualDiagnosisRequest contains the path parameters for
// POST /v3/clusters/{clusterId}/diagnosis/manual.
type CreateManualDiagnosisRequest struct {
	Request
	ClusterId string `json:"-"`
}

// CreateManualDiagnosisResponse contains the response body for
// POST /v3/clusters/{clusterId}/diagnosis/manual.
type CreateManualDiagnosisResponse struct {
	Success bool `json:"success,omitempty"`
}

// DiagnosisReportListItem describes a single diagnosis report as returned by the list API.
type DiagnosisReportListItem struct {
	ReportId      string `json:"reportId,omitempty"`
	Type          string `json:"type,omitempty"`
	ExecTime      string `json:"execTime,omitempty"`
	LowRiskCount  int    `json:"lowRiskCount,omitempty"`
	HighRiskCount int    `json:"highRiskCount,omitempty"`
	ErrorCount    int    `json:"errorCount,omitempty"`
}

// ListDiagnosisReportsRequest contains the path and query parameters for
// GET /v3/clusters/{clusterId}/diagnosis/reports.
type ListDiagnosisReportsRequest struct {
	Request
	ClusterId string `json:"-"`
	PageNo    int    `json:"-"`
	PageSize  int    `json:"-"`
}

// ListDiagnosisReportsResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/reports.
type ListDiagnosisReportsResponse struct {
	PageNo     int                       `json:"pageNo,omitempty"`
	PageSize   int                       `json:"pageSize,omitempty"`
	TotalCount int                       `json:"totalCount,omitempty"`
	Reports    []DiagnosisReportListItem `json:"reports,omitempty"`
}

// DiagnosisReportResultItem describes a single diagnosis conclusion within a report group.
type DiagnosisReportResultItem struct {
	Id               string `json:"id,omitempty"`
	Type             string `json:"type,omitempty"`
	Conclusion       string `json:"conclusion,omitempty"`
	HasLowRiskAdvice bool   `json:"hasLowRiskAdvice,omitempty"`
	NotExistIndex    bool   `json:"notExistIndex,omitempty"`
}

// DiagnosisReportGroupItem groups diagnosis results by risk grade.
type DiagnosisReportGroupItem struct {
	Grade     string                      `json:"grade,omitempty"`
	ItemCount int                         `json:"itemCount,omitempty"`
	Results   []DiagnosisReportResultItem `json:"results,omitempty"`
}

// GetDiagnosisReportRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/diagnosis/reports/{reportId}.
type GetDiagnosisReportRequest struct {
	Request
	ClusterId string `json:"-"`
	ReportId  string `json:"-"`
}

// GetDiagnosisReportResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/reports/{reportId}.
type GetDiagnosisReportResponse struct {
	State      string                     `json:"state,omitempty"`
	Type       string                     `json:"type,omitempty"`
	Indices    string                     `json:"indices,omitempty"`
	TotalCount int                        `json:"totalCount,omitempty"`
	ExecTime   string                     `json:"execTime,omitempty"`
	Groups     []DiagnosisReportGroupItem `json:"groups,omitempty"`
}

// GetDiagnosisBusyStatusRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/diagnosis/busy.
type GetDiagnosisBusyStatusRequest struct {
	Request
	ClusterId string `json:"-"`
}

// GetDiagnosisBusyStatusResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/busy.
type GetDiagnosisBusyStatusResponse struct {
	Busy bool `json:"busy,omitempty"`
}

// GetDiagnosisAuthorizationStatusRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/diagnosis/authorize.
type GetDiagnosisAuthorizationStatusRequest struct {
	Request
	ClusterId string `json:"-"`
}

// GetDiagnosisAuthorizationStatusResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/authorize.
type GetDiagnosisAuthorizationStatusResponse struct {
	Enabled bool `json:"enabled,omitempty"`
}

// AuthorizeDiagnosisRequest contains the path parameters for
// POST /v3/clusters/{clusterId}/diagnosis/authorize.
type AuthorizeDiagnosisRequest struct {
	Request
	ClusterId string `json:"-"`
}

// AuthorizeDiagnosisResponse contains the response body for
// POST /v3/clusters/{clusterId}/diagnosis/authorize.
type AuthorizeDiagnosisResponse struct {
	Authorized bool `json:"authorized,omitempty"`
}

// DiagnosisItem describes a single diagnosis item supported by the service.
type DiagnosisItem struct {
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

// ListDiagnosisItemsRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/diagnosis/items.
type ListDiagnosisItemsRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ListDiagnosisItemsResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/items.
type ListDiagnosisItemsResponse struct {
	Items []DiagnosisItem `json:"items,omitempty"`
}

// GetAutoDiagnosisStatusRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/diagnosis/auto.
type GetAutoDiagnosisStatusRequest struct {
	Request
	ClusterId string `json:"-"`
}

// GetAutoDiagnosisStatusResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/auto.
type GetAutoDiagnosisStatusResponse struct {
	Enabled bool `json:"enabled,omitempty"`
}

// UpdateAutoDiagnosisRequest contains the request body for
// PUT /v3/clusters/{clusterId}/diagnosis/auto.
type UpdateAutoDiagnosisRequest struct {
	Request
	ClusterId string `json:"-"`
	Enabled   *bool  `json:"enabled"`
}

// UpdateAutoDiagnosisResponse contains the response body for
// PUT /v3/clusters/{clusterId}/diagnosis/auto.
type UpdateAutoDiagnosisResponse struct {
	Enabled bool `json:"enabled,omitempty"`
}

// GetLatestOverviewRequest contains the path parameters for
// GET /v3/clusters/{clusterId}/diagnosis/overview/latest.
type GetLatestOverviewRequest struct {
	Request
	ClusterId string `json:"-"`
}

// GetLatestOverviewResponse contains the response body for
// GET /v3/clusters/{clusterId}/diagnosis/overview/latest.
type GetLatestOverviewResponse struct {
	ReportId      string `json:"reportId,omitempty"`
	Status        string `json:"status,omitempty"`
	ExecTime      string `json:"execTime,omitempty"`
	CountSafe     int    `json:"countSafe,omitempty"`
	CountLowRisk  int    `json:"countLowRisk,omitempty"`
	CountHighRisk int    `json:"countHighRisk,omitempty"`
	CountError    int    `json:"countError,omitempty"`
}

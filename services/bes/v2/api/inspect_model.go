package api

// AuthorizeInspectRequest is the request of authorizing intelligent inspection for a cluster.
type AuthorizeInspectRequest struct {
	ClusterId string `json:"clusterId"`
}

// AuthorizeInspectResponse is the response of authorizing intelligent inspection for a cluster.
type AuthorizeInspectResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// SwitchAutoInspectRequest is the request of turning auto inspection on or off.
type SwitchAutoInspectRequest struct {
	ClusterId string `json:"clusterId"`
	SwitchOn  bool   `json:"switchOn"`
}

// InspectStringSuccessCommonResponse is shared by inspect APIs.
type InspectStringSuccessCommonResponse struct {
	Success SuccessFlag `json:"success"`
	Status  int         `json:"status"`
	Result  interface{} `json:"result"`
}

// CheckAutoInspectRequest is the request of checking whether auto inspection is enabled.
type CheckAutoInspectRequest struct {
	ClusterId string `json:"clusterId"`
}

// CheckAutoInspectResponse is the response of checking whether auto inspection is enabled.
type CheckAutoInspectResponse struct {
	Success SuccessFlag        `json:"success"`
	Status  int                `json:"status"`
	Result  *AutoInspectStatus `json:"result"`
}

// AutoInspectStatus describes whether auto inspection is currently enabled.
type AutoInspectStatus struct {
	SwitchOn bool `json:"switchOn"`
}

// CreateManualInspectTaskRequest is the request of submitting a manual inspection task.
type CreateManualInspectTaskRequest struct {
	ClusterId string `json:"clusterId"`
}

// CheckInspectBusyRequest is the request of checking whether a new inspection task can be submitted.
type CheckInspectBusyRequest struct {
	ClusterId string `json:"clusterId"`
}

// CheckInspectBusyResponse is the response of checking whether a new inspection task can be submitted.
type CheckInspectBusyResponse struct {
	Success SuccessFlag        `json:"success"`
	Status  int                `json:"status"`
	Result  *InspectBusyStatus `json:"result"`
}

// InspectBusyStatus describes whether the cluster is currently busy running an inspection task.
type InspectBusyStatus struct {
	Busy bool `json:"busy"`
}

// GetManualInspectCountRequest is the request of getting today's completed manual inspection count.
type GetManualInspectCountRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetManualInspectCountResponse is the response of getting today's completed manual inspection count.
type GetManualInspectCountResponse struct {
	Success SuccessFlag         `json:"success"`
	Status  int                 `json:"status"`
	Result  *ManualInspectCount `json:"result"`
}

// ManualInspectCount carries today's completed manual inspection count.
type ManualInspectCount struct {
	Count int `json:"count"`
}

// GetManualInspectConfigRequest is the request of viewing a cluster's manual inspection configuration.
type GetManualInspectConfigRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetManualInspectConfigResponse is the response of viewing a cluster's manual inspection configuration.
type GetManualInspectConfigResponse struct {
	Success SuccessFlag          `json:"success"`
	Status  int                  `json:"status"`
	Result  *ManualInspectConfig `json:"result"`
}

// ManualInspectConfig describes a cluster's manual inspection configuration.
type ManualInspectConfig struct {
	Indices string   `json:"indices"`
	Items   []string `json:"items"`
}

// UpdateManualInspectConfigRequest is the request of updating a cluster's manual inspection configuration.
type UpdateManualInspectConfigRequest struct {
	ClusterId string   `json:"clusterId"`
	Indices   string   `json:"indices"`
	Items     []string `json:"items,omitempty"`
}

// ListInspectItemsResponse is the response of listing all selectable inspection items.
type ListInspectItemsResponse struct {
	Success SuccessFlag      `json:"success"`
	Status  int              `json:"status"`
	Result  *InspectItemList `json:"result"`
}

// InspectItemList carries the list of all selectable inspection items.
type InspectItemList struct {
	Items []InspectItem `json:"items"`
}

// InspectItem describes a single selectable inspection item.
type InspectItem struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

// GetInspectTaskRequest is the request of getting a single inspection task's status and result.
type GetInspectTaskRequest struct {
	ClusterId string `json:"clusterId"`
	TaskId    string `json:"taskId"`
}

// GetInspectTaskResponse is the response of getting a single inspection task's status and result.
type GetInspectTaskResponse struct {
	Success SuccessFlag        `json:"success"`
	Status  int                `json:"status"`
	Result  *InspectTaskDetail `json:"result"`
}

// InspectTaskDetail describes a single inspection task's execution status and result.
type InspectTaskDetail struct {
	State      string               `json:"state"`
	Type       string               `json:"type"`
	Indices    string               `json:"indices"`
	TotalCount int                  `json:"totalCount"`
	ExecTime   string               `json:"execTime"`
	Groups     []InspectResultGroup `json:"groups"`
}

// InspectResultGroup groups inspection results by risk grade.
type InspectResultGroup struct {
	Grade     string              `json:"grade"`
	ItemCount int                 `json:"itemCount"`
	Results   []InspectItemResult `json:"results"`
}

// InspectItemResult describes a single inspection item's result within a risk group.
type InspectItemResult struct {
	Id               string `json:"id"`
	Type             string `json:"type"`
	Conclusion       string `json:"conclusion"`
	HasLowRiskAdvice bool   `json:"hasLowRiskAdvice"`
	NotExistIndex    bool   `json:"notExistIndex"`
}

// ListInspectTasksRequest is the request of listing completed inspection tasks in the last 7 days.
type ListInspectTasksRequest struct {
	ClusterId string `json:"clusterId"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
}

// ListInspectTasksResponse is the response of listing completed inspection tasks in the last 7 days.
type ListInspectTasksResponse struct {
	Success SuccessFlag      `json:"success"`
	Status  int              `json:"status"`
	Page    *InspectTaskPage `json:"page"`
}

// InspectTaskPage carries a page of inspection tasks.
type InspectTaskPage struct {
	OrderBy    string            `json:"orderBy"`
	Order      string            `json:"order"`
	PageNo     int               `json:"pageNo"`
	PageSize   int               `json:"pageSize"`
	TotalCount int               `json:"totalCount"`
	Result     []InspectTaskItem `json:"result"`
}

// InspectTaskItem describes a single inspection task in the list.
type InspectTaskItem struct {
	TaskId        string `json:"taskId"`
	Type          string `json:"type"`
	ExecTime      string `json:"execTime"`
	LowRiskCount  int    `json:"lowRiskCount"`
	HighRiskCount int    `json:"highRiskCount"`
	ErrorCount    int    `json:"errorCount"`
}

// GetLatestInspectOverviewRequest is the request of getting the latest inspection overview.
type GetLatestInspectOverviewRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetLatestInspectOverviewResponse is the response of getting the latest inspection overview.
type GetLatestInspectOverviewResponse struct {
	Success SuccessFlag            `json:"success"`
	Status  int                    `json:"status"`
	Result  *LatestInspectOverview `json:"result"`
}

// LatestInspectOverview describes the latest inspection task's overview.
type LatestInspectOverview struct {
	TaskId        string `json:"taskId"`
	TaskStatus    string `json:"taskStatus"`
	CountSafe     int    `json:"countSafe"`
	CountLowRisk  int    `json:"countLowRisk"`
	CountHighRisk int    `json:"countHighRisk"`
	CountError    int    `json:"countError"`
	ExecTime      string `json:"execTime"`
}

// GetWeeklyInspectOverviewRequest is the request of getting the last 7 days' inspection overview.
type GetWeeklyInspectOverviewRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetWeeklyInspectOverviewResponse is the response of getting the last 7 days' inspection overview.
type GetWeeklyInspectOverviewResponse struct {
	Success SuccessFlag            `json:"success"`
	Status  int                    `json:"status"`
	Result  *WeeklyInspectOverview `json:"result"`
}

// WeeklyInspectOverview describes the last 7 days' inspection overview.
type WeeklyInspectOverview struct {
	Count     int               `json:"count"`
	RiskItems []InspectRiskItem `json:"riskItems"`
}

// InspectRiskItem describes a single inspection item's accumulated risk counts over 7 days.
type InspectRiskItem struct {
	Id            string `json:"id"`
	HighRiskCount int    `json:"highRiskCount"`
	LowRiskCount  int    `json:"lowRiskCount"`
}

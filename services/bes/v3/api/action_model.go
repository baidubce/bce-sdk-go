package api

// ListActionsRequest contains filters for querying BES task actions.
type ListActionsRequest struct {
	Request
	PageNo      int    `json:"-"`
	PageSize    int    `json:"-"`
	Name        string `json:"-"`
	Status      string `json:"-"`
	OrderBy     string `json:"-"`
	Order       string `json:"-"`
	ClusterName string `json:"-"`
	ClusterId   string `json:"-"`
}

// ListActionsResponse contains paged BES task actions.
type ListActionsResponse struct {
	Result     []ActionItem `json:"result,omitempty"`
	PageNo     int          `json:"pageNo,omitempty"`
	PageSize   int          `json:"pageSize,omitempty"`
	TotalCount int          `json:"totalCount,omitempty"`
}

// ActionItem describes a BES task action.
type ActionItem struct {
	ClusterId   string `json:"clusterId,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	Name        string `json:"name,omitempty"`
	NameCN      string `json:"nameCN,omitempty"`
	ActionId    string `json:"actionId,omitempty"`
	Status      string `json:"status,omitempty"`
	CreateTime  string `json:"createTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	RunningTime string `json:"runningTime,omitempty"`
}

// ListActionTypesRequest contains request metadata for querying action types.
type ListActionTypesRequest struct {
	Request
}

// ListActionTypesResponse contains BES task action types.
type ListActionTypesResponse struct {
	ActionTypes []ActionTypeItem `json:"actionTypes,omitempty"`
}

// ActionTypeItem describes a BES task action type.
type ActionTypeItem struct {
	Name   string `json:"name,omitempty"`
	NameCN string `json:"nameCN,omitempty"`
}

// ListOperationsRequest identifies a BES task action whose operations are queried.
type ListOperationsRequest struct {
	Request
	ClusterId string `json:"-"`
	ActionId  string `json:"-"`
}

// ListOperationsResponse contains operations under a BES task action.
type ListOperationsResponse struct {
	ClusterId  string          `json:"clusterId,omitempty"`
	ActionId   string          `json:"actionId,omitempty"`
	Name       string          `json:"name,omitempty"`
	NameCN     string          `json:"nameCN,omitempty"`
	Status     string          `json:"status,omitempty"`
	Operations []OperationItem `json:"operations,omitempty"`
}

// OperationItem describes one operation in a BES task action.
type OperationItem struct {
	ActionId    string `json:"actionId,omitempty"`
	OperationId string `json:"operationId,omitempty"`
	Type        string `json:"type,omitempty"`
	TypeCN      string `json:"typeCN,omitempty"`
	State       string `json:"state,omitempty"`
	Process     int    `json:"process,omitempty"`
	StartTime   string `json:"startTime,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	RunningTime string `json:"runningTime,omitempty"`
}

// GetOperationRequest identifies one operation in a BES task action.
type GetOperationRequest struct {
	Request
	ClusterId   string `json:"-"`
	ActionId    string `json:"-"`
	OperationId string `json:"-"`
}

// GetOperationResponse contains details for one BES task operation.
type GetOperationResponse struct {
	ActionId      string      `json:"actionId,omitempty"`
	OperationId   string      `json:"operationId,omitempty"`
	ClusterId     string      `json:"clusterId,omitempty"`
	Process       int         `json:"process,omitempty"`
	Groups        []GroupItem `json:"groups,omitempty"`
	SourceContext string      `json:"sourceContext,omitempty"`
	TargetContext string      `json:"targetContext,omitempty"`
}

// GroupItem describes a group in a BES task operation.
type GroupItem struct {
	GroupName   string `json:"groupName,omitempty"`
	GroupNameCN string `json:"groupNameCN,omitempty"`
	State       string `json:"state,omitempty"`
	Analysis    *bool  `json:"analysis,omitempty"`
	Diagnosis   string `json:"diagnosis,omitempty"`
	DiagnosisCN string `json:"diagnosisCN,omitempty"`
}

// ListOperationAnalysisDetailsRequest contains filters for operation analysis details.
type ListOperationAnalysisDetailsRequest struct {
	Request
	ClusterId   string `json:"-"`
	ActionId    string `json:"-"`
	OperationId string `json:"-"`
	GroupNames  string `json:"-"`
	PageNo      int    `json:"-"`
	PageSize    int    `json:"-"`
}

// ListOperationAnalysisDetailsResponse contains paged operation analysis details.
type ListOperationAnalysisDetailsResponse struct {
	Result     []KindItem `json:"result,omitempty"`
	PageNo     int        `json:"pageNo,omitempty"`
	PageSize   int        `json:"pageSize,omitempty"`
	TotalCount int        `json:"totalCount,omitempty"`
}

// KindItem describes one operation analysis detail item.
type KindItem struct {
	Name   string `json:"name,omitempty"`
	NameCN string `json:"nameCN,omitempty"`
	State  string `json:"state,omitempty"`
}

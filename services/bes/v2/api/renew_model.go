package api

import "encoding/json"

// CreateAutoRenewRuleRequest is the request of creating an auto-renew rule for one or more clusters.
type CreateAutoRenewRuleRequest struct {
	ClusterIds    []string `json:"clusterIds"`
	RenewTimeUnit string   `json:"renewTimeUnit"`
	RenewTime     int      `json:"renewTime"`
	ServiceType   string   `json:"serviceType"`
}

// OrderIdResult carries the order id of a successfully submitted order.
// Shared by CreateAutoRenewRuleResponse and RenewClusterResponse.
type OrderIdResult struct {
	OrderId string `json:"orderId"`
}

// UnmarshalJSON tolerates the server answering result as an empty string instead of an object.
func (r *OrderIdResult) UnmarshalJSON(data []byte) error {
	if isEmptyJSONResult(data) {
		return nil
	}
	type plain OrderIdResult
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*r = OrderIdResult(decoded)
	return nil
}

// CreateAutoRenewRuleResponse is the response of creating an auto-renew rule.
type CreateAutoRenewRuleResponse struct {
	Success bool           `json:"success"`
	Status  int            `json:"status"`
	Result  *OrderIdResult `json:"result"`
}

// GetAutoRenewRuleDetailRequest is the request of getting a cluster's auto-renew rule detail.
type GetAutoRenewRuleDetailRequest struct {
	ClusterId string `json:"clusterId"`
}

// GetAutoRenewRuleDetailResponse is the response of getting a cluster's auto-renew rule detail.
type GetAutoRenewRuleDetailResponse struct {
	Success bool                 `json:"success"`
	Status  int                  `json:"status"`
	Result  *AutoRenewRuleDetail `json:"result"`
}

// AutoRenewRuleDetail describes a cluster's auto-renew rule.
type AutoRenewRuleDetail struct {
	ClusterId     string  `json:"clusterId"`
	RenewTimeUnit string  `json:"renewTimeUnit"`
	RenewTime     float64 `json:"renewTime"`
	ExpireTime    float64 `json:"expireTime"`
	NextRenewTime float64 `json:"nextRenewTime"`
}

// UnmarshalJSON tolerates the server answering result as an empty string instead of an object,
// which happens when the cluster has no auto-renew rule at all.
func (d *AutoRenewRuleDetail) UnmarshalJSON(data []byte) error {
	if isEmptyJSONResult(data) {
		return nil
	}
	type plain AutoRenewRuleDetail
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*d = AutoRenewRuleDetail(decoded)
	return nil
}

// ListAutoRenewRulesRequest is the request of listing auto-renew rules.
type ListAutoRenewRulesRequest struct {
	ClusterId   string `json:"clusterId,omitempty"`
	ServiceType string `json:"serviceType"`
}

// ListAutoRenewRulesResponse is the response of listing auto-renew rules.
// The result is directly a list, not wrapped in an object.
type ListAutoRenewRulesResponse struct {
	Success bool            `json:"success"`
	Status  int             `json:"status"`
	Result  []AutoRenewRule `json:"result"`
}

// AutoRenewRule describes a single auto-renew rule.
type AutoRenewRule struct {
	Uuid          string `json:"uuid"`
	UserId        string `json:"userId"`
	ClusterId     string `json:"clusterId"`
	Region        string `json:"region"`
	RenewTimeUnit string `json:"renewTimeUnit"`
	RenewTime     int    `json:"renewTime"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
}

// UpdateAutoRenewRuleRequest is the request of updating a cluster's auto-renew rule.
type UpdateAutoRenewRuleRequest struct {
	ClusterId     string `json:"clusterId"`
	RenewTimeUnit string `json:"renewTimeUnit"`
	RenewTime     int    `json:"renewTime"`
}

// RenewStringResultResponse is shared by renew APIs whose response is
// {success: Boolean, status: Integer, result: String}.
type RenewStringResultResponse struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Result  string `json:"result"`
}

// DeleteAutoRenewRuleRequest is the request of deleting a cluster's auto-renew rule.
type DeleteAutoRenewRuleRequest struct {
	ClusterId string `json:"clusterId,omitempty"`
}

// DeleteAutoRenewRuleResponse is the response of deleting a cluster's auto-renew rule.
type DeleteAutoRenewRuleResponse struct {
	Success bool              `json:"success"`
	Status  int               `json:"status"`
	Result  *ExpireTimeResult `json:"result"`
}

// ExpireTimeResult carries the expire time left after an auto-renew rule is deleted.
type ExpireTimeResult struct {
	ExpireTime float64 `json:"expireTime"`
}

// UnmarshalJSON tolerates the server answering result as an empty string instead of an object.
func (r *ExpireTimeResult) UnmarshalJSON(data []byte) error {
	if isEmptyJSONResult(data) {
		return nil
	}
	type plain ExpireTimeResult
	var decoded plain
	if err := json.Unmarshal(data, &decoded); err != nil {
		return err
	}
	*r = ExpireTimeResult(decoded)
	return nil
}

// RenewClusterRequest is the request of renewing a cluster.
type RenewClusterRequest struct {
	ClusterId string `json:"clusterId"`
	Time      int    `json:"time"`
}

// RenewClusterResponse is the response of renewing a cluster.
type RenewClusterResponse struct {
	Success bool           `json:"success"`
	Status  int            `json:"status"`
	Result  *OrderIdResult `json:"result"`
}

// ListRenewalsRequest is the request of listing clusters approaching renewal.
type ListRenewalsRequest struct {
	Order            string `json:"order"`
	OrderBy          string `json:"orderBy"`
	PageNo           int    `json:"pageNo"`
	PageSize         int    `json:"pageSize"`
	DaysToExpiration int    `json:"daysToExpiration"`
}

// ListRenewalsResponse is the response of listing clusters approaching renewal.
// The response has a top-level "page" field instead of "result", and no "status" field.
type ListRenewalsResponse struct {
	Success bool         `json:"success"`
	Page    *RenewalPage `json:"page"`
}

// RenewalPage carries a page of clusters approaching renewal.
type RenewalPage struct {
	OrderBy    string           `json:"orderBy"`
	Order      string           `json:"order"`
	PageNo     int              `json:"pageNo"`
	PageSize   int              `json:"pageSize"`
	TotalCount int              `json:"totalCount"`
	Result     []RenewalCluster `json:"result"`
}

// RenewalCluster describes a single cluster approaching renewal.
type RenewalCluster struct {
	ClusterId     string `json:"clusterId"`
	ClusterName   string `json:"clusterName"`
	Region        string `json:"region"`
	ExpiredTime   string `json:"expiredTime"`
	ClusterStatus string `json:"clusterStatus"`
}

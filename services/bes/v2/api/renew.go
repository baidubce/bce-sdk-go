package api

import (
	"errors"
	nethttp "net/http"

	bcehttp "github.com/baidubce/bce-sdk-go/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// CreateAutoRenewRule creates an auto-renew rule for one or more clusters.
func CreateAutoRenewRule(cli bce.Client, region string, request *CreateAutoRenewRuleRequest) (*CreateAutoRenewRuleResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create auto renew rule request should not be nil")
	}
	if len(request.ClusterIds) == 0 {
		return nil, errors.New("create auto renew rule request clusterIds should not be empty")
	}
	if request.RenewTimeUnit == "" {
		return nil, errors.New("create auto renew rule request renewTimeUnit should not be empty")
	}
	if request.RenewTime == 0 {
		return nil, errors.New("create auto renew rule request renewTime should not be empty")
	}
	if request.ServiceType == "" {
		return nil, errors.New("create auto renew rule request serviceType should not be empty")
	}

	uri := URI_AUTO_RENEW_RULE + "/create"

	result := &CreateAutoRenewRuleResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAutoRenewRuleDetail gets a cluster's auto-renew rule detail.
func GetAutoRenewRuleDetail(cli bce.Client, region string, request *GetAutoRenewRuleDetailRequest) (*GetAutoRenewRuleDetailResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get auto renew rule detail request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get auto renew rule detail request clusterId should not be empty")
	}

	uri := URI_AUTO_RENEW_RULE + "/detail"

	result := &GetAutoRenewRuleDetailResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListAutoRenewRules lists auto-renew rules.
func ListAutoRenewRules(cli bce.Client, region string, request *ListAutoRenewRulesRequest) (*ListAutoRenewRulesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list auto renew rules request should not be nil")
	}
	if request.ServiceType == "" {
		return nil, errors.New("list auto renew rules request serviceType should not be empty")
	}

	uri := URI_AUTO_RENEW_RULE + "/list"

	result := &ListAutoRenewRulesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateAutoRenewRule updates a cluster's auto-renew rule.
func UpdateAutoRenewRule(cli bce.Client, region string, request *UpdateAutoRenewRuleRequest) (*RenewStringResultResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update auto renew rule request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update auto renew rule request clusterId should not be empty")
	}
	if request.RenewTimeUnit == "" {
		return nil, errors.New("update auto renew rule request renewTimeUnit should not be empty")
	}
	if request.RenewTime == 0 {
		return nil, errors.New("update auto renew rule request renewTime should not be empty")
	}

	uri := URI_AUTO_RENEW_RULE + "/update"

	result := &RenewStringResultResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteAutoRenewRule deletes a cluster's auto-renew rule.
func DeleteAutoRenewRule(cli bce.Client, region string, request *DeleteAutoRenewRuleRequest) (*DeleteAutoRenewRuleResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete auto renew rule request should not be nil")
	}

	uri := URI_AUTO_RENEW_RULE + "/delete"

	result := &DeleteAutoRenewRuleResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RenewCluster renews a cluster for a given number of months.
func RenewCluster(cli bce.Client, region string, request *RenewClusterRequest) (*RenewClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("renew cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("renew cluster request clusterId should not be empty")
	}
	if request.Time == 0 {
		return nil, errors.New("renew cluster request time should not be empty")
	}

	result := &RenewClusterResponse{}
	builder := bce.NewRequestBuilder(cli).
		WithMethod(nethttp.MethodPost).
		WithURL(URI_RENEW).
		WithQueryParam("orderType", "RENEW").
		WithHeader(bcehttp.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithHeader(HEADER_REGION, region).
		WithBody(request).
		WithResult(result)
	if err := builder.Do(); err != nil {
		return nil, err
	}
	return result, nil
}

// ListRenewals lists clusters approaching renewal.
func ListRenewals(cli bce.Client, region string, request *ListRenewalsRequest) (*ListRenewalsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list renewals request should not be nil")
	}
	if request.Order == "" {
		return nil, errors.New("list renewals request order should not be empty")
	}
	if request.OrderBy == "" {
		return nil, errors.New("list renewals request orderBy should not be empty")
	}
	if request.PageNo == 0 {
		return nil, errors.New("list renewals request pageNo should not be empty")
	}
	if request.PageSize == 0 {
		return nil, errors.New("list renewals request pageSize should not be empty")
	}

	uri := URI_RENEW + "/list"

	result := &ListRenewalsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

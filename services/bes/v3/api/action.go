package api

import (
	"errors"
	"net/url"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	URI_ACTIONS          = URI_CLUSTERS + "/actions"
	URI_ACTION_TYPES     = URI_ACTIONS + "/action-types"
	URI_OPERATIONS       = "actions"
	URI_OPERATION_GROUPS = "groups"
)

func actionURI(parts ...string) string {
	if len(parts) == 0 {
		return URI_ACTIONS
	}
	uri := URI_CLUSTERS
	for _, part := range parts {
		uri += "/" + url.PathEscape(part)
	}
	return uri
}

// ListActions lists BES task actions.
func ListActions(cli bce.Client, region string, request *ListActionsRequest) (*ListActionsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list actions request should not be nil")
	}
	region = setRegion(region, request.Region)

	query := map[string]string{}
	if request.PageNo > 0 {
		query["pageNo"] = intString(request.PageNo)
	}
	if request.PageSize > 0 {
		query["pageSize"] = intString(request.PageSize)
	}
	if request.Name != "" {
		query["name"] = request.Name
	}
	if request.Status != "" {
		query["status"] = request.Status
	}
	if request.OrderBy != "" {
		query["orderBy"] = request.OrderBy
	}
	if request.Order != "" {
		query["order"] = request.Order
	}
	if request.ClusterName != "" {
		query["clusterName"] = request.ClusterName
	}
	if request.ClusterId != "" {
		query["clusterId"] = request.ClusterId
	}

	result := &ListActionsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, URI_ACTIONS, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListActionTypes lists the visible task action types in BES OpenAPI.
func ListActionTypes(cli bce.Client, region string, request *ListActionTypesRequest) (*ListActionTypesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list action types request should not be nil")
	}
	region = setRegion(region, request.Region)

	result := &ListActionTypesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, URI_ACTION_TYPES, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListOperations lists operations under a BES task action.
func ListOperations(cli bce.Client, region string, request *ListOperationsRequest) (*ListOperationsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list operations request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list operations request clusterId should not be empty")
	}
	if request.ActionId == "" {
		return nil, errors.New("list operations request actionId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := actionURI(request.ClusterId, "actions", request.ActionId, "operations")
	result := &ListOperationsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetOperation gets the detail of one BES task operation.
func GetOperation(cli bce.Client, region string, request *GetOperationRequest) (*GetOperationResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get operation request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get operation request clusterId should not be empty")
	}
	if request.ActionId == "" {
		return nil, errors.New("get operation request actionId should not be empty")
	}
	if request.OperationId == "" {
		return nil, errors.New("get operation request operationId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := actionURI(request.ClusterId, "actions", request.ActionId, "operations", request.OperationId)
	result := &GetOperationResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListOperationAnalysisDetails lists analysis details for a task operation group.
func ListOperationAnalysisDetails(cli bce.Client, region string, request *ListOperationAnalysisDetailsRequest) (*ListOperationAnalysisDetailsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list operation analysis details request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list operation analysis details request clusterId should not be empty")
	}
	if request.ActionId == "" {
		return nil, errors.New("list operation analysis details request actionId should not be empty")
	}
	if request.OperationId == "" {
		return nil, errors.New("list operation analysis details request operationId should not be empty")
	}
	if request.GroupNames == "" {
		return nil, errors.New("list operation analysis details request groupNames should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := actionURI(request.ClusterId, "actions", request.ActionId, "operations", request.OperationId, "groups", request.GroupNames)
	query := map[string]string{}
	if request.PageNo > 0 {
		query["pageNo"] = intString(request.PageNo)
	}
	if request.PageSize > 0 {
		query["pageSize"] = intString(request.PageSize)
	}

	result := &ListOperationAnalysisDetailsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

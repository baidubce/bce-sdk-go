package api

import (
	"errors"
	"strconv"

	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// CreateIndex creates an index in a BES cluster.
func CreateIndex(cli bce.Client, region string, request *CreateIndexRequest) (*CreateIndexResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create index request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("create index request clusterId should not be empty")
	}
	if request.IndexName == "" {
		return nil, errors.New("create index request indexName should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", request.IndexName)

	result := &CreateIndexResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListIndices lists indices in a BES cluster.
func ListIndices(cli bce.Client, region string, request *ListIndicesRequest) (*ListIndicesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list indices request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list indices request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices")
	query := map[string]string{}
	if request.IndexName != "" {
		query["indexName"] = request.IndexName
	}
	if request.Health != "" {
		query["health"] = request.Health
	}
	if request.IncludeSystem != nil {
		query["includeSystem"] = strconv.FormatBool(*request.IncludeSystem)
	}
	if request.PageNo > 0 {
		query["pageNo"] = intString(request.PageNo)
	}
	if request.PageSize > 0 {
		query["pageSize"] = intString(request.PageSize)
	}
	if request.OrderBy != "" {
		query["orderBy"] = request.OrderBy
	}
	if request.Order != "" {
		query["order"] = request.Order
	}

	result := &ListIndicesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetIndex gets the detail of an index in a BES cluster.
func GetIndex(cli bce.Client, region string, request *GetIndexRequest) (*GetIndexResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get index request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get index request clusterId should not be empty")
	}
	if request.IndexName == "" {
		return nil, errors.New("get index request indexName should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", request.IndexName)

	result := &GetIndexResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExistIndex checks whether an index exists in a BES cluster.
func ExistIndex(cli bce.Client, region string, request *ExistIndexRequest) (*ExistResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("exist index request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("exist index request clusterId should not be empty")
	}
	if request.IndexName == "" {
		return nil, errors.New("exist index request indexName should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", request.IndexName, "_exist")

	result := &ExistResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetIndexStats gets the stats of an index in a BES cluster.
func GetIndexStats(cli bce.Client, region string, request *GetIndexStatsRequest) (*IndexStatsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get index stats request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get index stats request clusterId should not be empty")
	}
	if request.IndexName == "" {
		return nil, errors.New("get index stats request indexName should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", request.IndexName, "stats")

	result := &IndexStatsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListIndexFieldTypes lists the supported index field types for a BES cluster.
func ListIndexFieldTypes(cli bce.Client, region string, request *ListIndexFieldTypesRequest) (*ListIndexFieldTypesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list index field types request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list index field types request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "index-field-types")

	result := &ListIndexFieldTypesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// OpenIndices opens the given indices in a BES cluster.
func OpenIndices(cli bce.Client, region string, request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("open indices request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("open indices request clusterId should not be empty")
	}
	if len(request.IndexNames) == 0 {
		return nil, errors.New("open indices request indexNames should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", "_open")

	result := &OperationSuccessResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CloseIndices closes the given indices in a BES cluster.
func CloseIndices(cli bce.Client, region string, request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("close indices request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("close indices request clusterId should not be empty")
	}
	if len(request.IndexNames) == 0 {
		return nil, errors.New("close indices request indexNames should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", "_close")

	result := &OperationSuccessResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteIndices deletes the given indices in a BES cluster.
func DeleteIndices(cli bce.Client, region string, request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete indices request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete indices request clusterId should not be empty")
	}
	if len(request.IndexNames) == 0 {
		return nil, errors.New("delete indices request indexNames should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", "_delete")

	result := &OperationSuccessResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RefreshIndices refreshes the given indices in a BES cluster.
func RefreshIndices(cli bce.Client, region string, request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("refresh indices request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("refresh indices request clusterId should not be empty")
	}
	if len(request.IndexNames) == 0 {
		return nil, errors.New("refresh indices request indexNames should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", "_refresh")

	result := &OperationSuccessResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// FlushIndices flushes the given indices in a BES cluster.
func FlushIndices(cli bce.Client, region string, request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("flush indices request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("flush indices request clusterId should not be empty")
	}
	if len(request.IndexNames) == 0 {
		return nil, errors.New("flush indices request indexNames should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", "_flush")

	result := &OperationSuccessResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ForceMergeIndices force-merges the given indices in a BES cluster.
func ForceMergeIndices(cli bce.Client, region string, request *ForceMergeIndicesRequest) (*OperationSuccessResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("force merge indices request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("force merge indices request clusterId should not be empty")
	}
	if len(request.IndexNames) == 0 {
		return nil, errors.New("force merge indices request indexNames should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", "_forcemerge")

	result := &OperationSuccessResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ClearIndicesCache clears the cache of the given indices in a BES cluster.
func ClearIndicesCache(cli bce.Client, region string, request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("clear indices cache request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("clear indices cache request clusterId should not be empty")
	}
	if len(request.IndexNames) == 0 {
		return nil, errors.New("clear indices cache request indexNames should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", "_clear_cache")

	result := &OperationSuccessResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateIndexSettings updates the settings of an index in a BES cluster.
func UpdateIndexSettings(cli bce.Client, region string, request *UpdateIndexSettingsRequest) (*IndexNameResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update index settings request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update index settings request clusterId should not be empty")
	}
	if request.IndexName == "" {
		return nil, errors.New("update index settings request indexName should not be empty")
	}
	if len(request.Settings) == 0 {
		return nil, errors.New("update index settings request settings should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", request.IndexName, "settings")

	result := &IndexNameResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateIndexMappings updates the mappings of an index in a BES cluster.
func UpdateIndexMappings(cli bce.Client, region string, request *UpdateIndexMappingsRequest) (*IndexNameResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update index mappings request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update index mappings request clusterId should not be empty")
	}
	if request.IndexName == "" {
		return nil, errors.New("update index mappings request indexName should not be empty")
	}
	if len(request.Mappings) == 0 {
		return nil, errors.New("update index mappings request mappings should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", request.IndexName, "mappings")

	result := &IndexNameResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateIndexAliases updates the aliases of an index in a BES cluster.
func UpdateIndexAliases(cli bce.Client, region string, request *UpdateIndexAliasesRequest) (*IndexNameResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update index aliases request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update index aliases request clusterId should not be empty")
	}
	if request.IndexName == "" {
		return nil, errors.New("update index aliases request indexName should not be empty")
	}
	if len(request.Actions) == 0 {
		return nil, errors.New("update index aliases request actions should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "indices", request.IndexName, "aliases")

	result := &IndexNameResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListIndexTemplates lists index templates in a BES cluster.
func ListIndexTemplates(cli bce.Client, region string, request *ListIndexTemplatesRequest) (*ListIndexTemplatesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list index templates request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list index templates request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "index-templates")

	result := &ListIndexTemplatesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetIndexTemplate gets the detail of an index template in a BES cluster.
func GetIndexTemplate(cli bce.Client, region string, request *GetIndexTemplateRequest) (*GetIndexTemplateResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get index template request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get index template request clusterId should not be empty")
	}
	if request.TemplateName == "" {
		return nil, errors.New("get index template request templateName should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "index-templates", request.TemplateName)

	result := &GetIndexTemplateResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExistIndexTemplate checks whether an index template exists in a BES cluster.
func ExistIndexTemplate(cli bce.Client, region string, request *ExistIndexTemplateRequest) (*ExistResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("exist index template request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("exist index template request clusterId should not be empty")
	}
	if request.TemplateName == "" {
		return nil, errors.New("exist index template request templateName should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "index-templates", request.TemplateName, "_exist")

	result := &ExistResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIndexTemplate creates an index template in a BES cluster.
func CreateIndexTemplate(cli bce.Client, region string, request *CreateIndexTemplateRequest) (*TemplateNameResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create index template request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("create index template request clusterId should not be empty")
	}
	if len(request.IndexPatterns) == 0 {
		return nil, errors.New("create index template request indexPatterns should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "index-templates")

	result := &TemplateNameResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateIndexTemplate updates an index template in a BES cluster.
func UpdateIndexTemplate(cli bce.Client, region string, request *UpdateIndexTemplateRequest) (*TemplateNameResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update index template request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update index template request clusterId should not be empty")
	}
	if request.TemplateName == "" {
		return nil, errors.New("update index template request templateName should not be empty")
	}
	if len(request.IndexPatterns) == 0 {
		return nil, errors.New("update index template request indexPatterns should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "index-templates", request.TemplateName)

	result := &TemplateNameResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteIndexTemplate deletes an index template in a BES cluster.
func DeleteIndexTemplate(cli bce.Client, region string, request *DeleteIndexTemplateRequest) (*TemplateNameResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete index template request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete index template request clusterId should not be empty")
	}
	if request.TemplateName == "" {
		return nil, errors.New("delete index template request templateName should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "index-templates", request.TemplateName)
	query := map[string]string{}
	if request.IsLegacy != "" {
		query["isLegacy"] = request.IsLegacy
	}

	result := &TemplateNameResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodDelete, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListIsmPolicies lists ISM policies in a BES cluster.
func ListIsmPolicies(cli bce.Client, region string, request *ListIsmPoliciesRequest) (*ListIsmPoliciesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list ism policies request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list ism policies request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "ism-policies")

	result := &ListIsmPoliciesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetIsmPolicy gets the detail of an ISM policy in a BES cluster.
func GetIsmPolicy(cli bce.Client, region string, request *GetIsmPolicyRequest) (*GetIsmPolicyResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get ism policy request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get ism policy request clusterId should not be empty")
	}
	if request.PolicyId == "" {
		return nil, errors.New("get ism policy request policyId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "ism-policies", request.PolicyId)

	result := &GetIsmPolicyResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ExistIsmPolicy checks whether an ISM policy exists in a BES cluster.
func ExistIsmPolicy(cli bce.Client, region string, request *ExistIsmPolicyRequest) (*ExistResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("exist ism policy request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("exist ism policy request clusterId should not be empty")
	}
	if request.PolicyId == "" {
		return nil, errors.New("exist ism policy request policyId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "ism-policies", request.PolicyId, "_exist")

	result := &ExistResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateIsmPolicy creates an ISM policy in a BES cluster.
func CreateIsmPolicy(cli bce.Client, region string, request *IsmPolicyRequest) (*PolicyIDResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create ism policy request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("create ism policy request clusterId should not be empty")
	}
	if request.PolicyId == "" {
		return nil, errors.New("create ism policy request policyId should not be empty")
	}
	if request.Description == "" {
		return nil, errors.New("create ism policy request description should not be empty")
	}
	if len(request.IsmTemplate) == 0 {
		return nil, errors.New("create ism policy request ismTemplate should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "ism-policies", request.PolicyId)

	result := &PolicyIDResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateIsmPolicy updates an ISM policy in a BES cluster.
func UpdateIsmPolicy(cli bce.Client, region string, request *IsmPolicyRequest) (*PolicyIDResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("update ism policy request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("update ism policy request clusterId should not be empty")
	}
	if request.PolicyId == "" {
		return nil, errors.New("update ism policy request policyId should not be empty")
	}
	if request.Description == "" {
		return nil, errors.New("update ism policy request description should not be empty")
	}
	if len(request.IsmTemplate) == 0 {
		return nil, errors.New("update ism policy request ismTemplate should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "ism-policies", request.PolicyId)

	result := &PolicyIDResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPut, uri, nil, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteIsmPolicy deletes an ISM policy in a BES cluster.
func DeleteIsmPolicy(cli bce.Client, region string, request *DeleteIsmPolicyRequest) (*PolicyIDResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete ism policy request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete ism policy request clusterId should not be empty")
	}
	if request.PolicyId == "" {
		return nil, errors.New("delete ism policy request policyId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "index-management", "ism-policies", request.PolicyId)

	result := &PolicyIDResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodDelete, uri, nil, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListMonitorIndices lists monitored indices in a BES cluster.
func ListMonitorIndices(cli bce.Client, region string, request *ListMonitorIndicesRequest) (*ListMonitorIndicesResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list monitor indices request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("list monitor indices request clusterId should not be empty")
	}
	region = setRegion(region, request.Region)

	uri := clusterURI(request.ClusterId, "monitor", "indices")
	query := map[string]string{}
	if request.IndexName != "" {
		query["indexName"] = request.IndexName
	}
	if request.IncludeSystem != nil {
		query["includeSystem"] = strconv.FormatBool(*request.IncludeSystem)
	}
	if request.PageNo > 0 {
		query["pageNo"] = intString(request.PageNo)
	}
	if request.PageSize > 0 {
		query["pageSize"] = intString(request.PageSize)
	}

	result := &ListMonitorIndicesResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodGet, uri, query, nil, result); err != nil {
		return nil, err
	}
	return result, nil
}

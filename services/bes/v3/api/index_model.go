package api

// CreateIndexRequest contains the request body for POST /v3/clusters/{clusterId}/index-management/indices/{indexName}.
type CreateIndexRequest struct {
	Request
	ClusterId string                 `json:"-"`
	IndexName string                 `json:"-"`
	Settings  map[string]interface{} `json:"settings,omitempty"`
	Mappings  map[string]interface{} `json:"mappings,omitempty"`
	Aliases   map[string]interface{} `json:"aliases,omitempty"`
}

// CreateIndexResponse contains the response body for POST /v3/clusters/{clusterId}/index-management/indices/{indexName}.
type CreateIndexResponse struct {
	IndexName string `json:"indexName,omitempty"`
}

// ListIndexItem describes a single index as returned by the list/detail index APIs.
type ListIndexItem struct {
	IndexName     string `json:"indexName,omitempty"`
	Health        string `json:"health,omitempty"`
	Status        string `json:"status,omitempty"`
	PrimaryShards string `json:"primaryShards,omitempty"`
	Replicas      string `json:"replicas,omitempty"`
	DocumentCount string `json:"documentCount,omitempty"`
	StorageSize   string `json:"storageSize,omitempty"`
	CreateTime    string `json:"createTime,omitempty"`
	Category      string `json:"category,omitempty"`
}

// ListIndicesRequest contains the optional query filters for
// GET /v3/clusters/{clusterId}/index-management/indices.
type ListIndicesRequest struct {
	Request
	ClusterId     string `json:"-"`
	IndexName     string `json:"-"`
	Health        string `json:"-"`
	IncludeSystem *bool  `json:"-"`
	PageNo        int    `json:"-"`
	PageSize      int    `json:"-"`
	OrderBy       string `json:"-"`
	Order         string `json:"-"`
}

// ListIndicesResponse contains the response body for
// GET /v3/clusters/{clusterId}/index-management/indices.
type ListIndicesResponse struct {
	Indices    []ListIndexItem `json:"indices,omitempty"`
	PageNo     int             `json:"pageNo,omitempty"`
	PageSize   int             `json:"pageSize,omitempty"`
	TotalCount int             `json:"totalCount,omitempty"`
}

// GetIndexRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/indices/{indexName}.
type GetIndexRequest struct {
	Request
	ClusterId string `json:"-"`
	IndexName string `json:"-"`
}

// GetIndexResponse contains the response body for
// GET /v3/clusters/{clusterId}/index-management/indices/{indexName}.
type GetIndexResponse struct {
	Index    ListIndexItem `json:"index,omitempty"`
	Mappings string        `json:"mappings,omitempty"`
	Settings string        `json:"settings,omitempty"`
	Aliases  string        `json:"aliases,omitempty"`
}

// ExistIndexRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/indices/{indexName}/_exist.
type ExistIndexRequest struct {
	Request
	ClusterId string `json:"-"`
	IndexName string `json:"-"`
}

// ExistResponse is shared by the index/index-template/ism-policy existence check APIs.
type ExistResponse struct {
	Exist bool `json:"exist,omitempty"`
}

// GetIndexStatsRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/indices/{indexName}/stats.
type GetIndexStatsRequest struct {
	Request
	ClusterId string `json:"-"`
	IndexName string `json:"-"`
}

// IndexStatsResponse contains the response body for
// GET /v3/clusters/{clusterId}/index-management/indices/{indexName}/stats.
type IndexStatsResponse struct {
	Stats map[string]interface{} `json:"stats,omitempty"`
}

// FieldTypeItemVO describes a supported index field type.
type FieldTypeItemVO struct {
	TypeName       string `json:"typeName,omitempty"`
	IsVector       string `json:"isVector,omitempty"`
	IsAnalyzable   string `json:"isAnalyzable,omitempty"`
	IsAggregatable string `json:"isAggregatable,omitempty"`
}

// ListIndexFieldTypesRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/index-field-types.
type ListIndexFieldTypesRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ListIndexFieldTypesResponse contains the response body for
// GET /v3/clusters/{clusterId}/index-management/index-field-types.
type ListIndexFieldTypesResponse struct {
	FieldTypes []FieldTypeItemVO `json:"fieldTypes,omitempty"`
}

// IndexNamesRequest is shared by the batch index operation APIs whose request body
// is only an index name list: open, close, batch delete, refresh, flush, clear cache.
type IndexNamesRequest struct {
	Request
	ClusterId  string   `json:"-"`
	IndexNames []string `json:"indexNames"`
}

// OperationSuccessResponse is shared by the batch index operation APIs whose response
// body is only a success flag: open, close, batch delete, refresh, flush, force merge,
// clear cache.
type OperationSuccessResponse struct {
	Success bool `json:"success,omitempty"`
}

// ForceMergeIndicesRequest contains the request body for
// POST /v3/clusters/{clusterId}/index-management/indices/_forcemerge.
type ForceMergeIndicesRequest struct {
	Request
	ClusterId          string   `json:"-"`
	IndexNames         []string `json:"indexNames"`
	MaxNumSegments     *int     `json:"maxNumSegments,omitempty"`
	OnlyExpungeDeletes *bool    `json:"onlyExpungeDeletes,omitempty"`
}

// IndexNameResponse is shared by the index mutation APIs whose response body is only
// the index name: create index, update settings, update mappings, update aliases.
type IndexNameResponse struct {
	IndexName string `json:"indexName,omitempty"`
}

// UpdateIndexSettingsRequest contains the request body for
// PUT /v3/clusters/{clusterId}/index-management/indices/{indexName}/settings.
type UpdateIndexSettingsRequest struct {
	Request
	ClusterId string                 `json:"-"`
	IndexName string                 `json:"-"`
	Settings  map[string]interface{} `json:"settings"`
}

// UpdateIndexMappingsRequest contains the request body for
// PUT /v3/clusters/{clusterId}/index-management/indices/{indexName}/mappings.
type UpdateIndexMappingsRequest struct {
	Request
	ClusterId string                 `json:"-"`
	IndexName string                 `json:"-"`
	Type      string                 `json:"type,omitempty"`
	Mappings  map[string]interface{} `json:"mappings"`
}

// AliasAction describes a single alias mutation operation.
type AliasAction struct {
	Action        string                 `json:"action,omitempty"`
	Alias         string                 `json:"alias,omitempty"`
	IsWriteIndex  *bool                  `json:"isWriteIndex,omitempty"`
	IndexRouting  string                 `json:"indexRouting,omitempty"`
	SearchRouting string                 `json:"searchRouting,omitempty"`
	Filter        map[string]interface{} `json:"filter,omitempty"`
}

// UpdateIndexAliasesRequest contains the request body for
// PUT /v3/clusters/{clusterId}/index-management/indices/{indexName}/aliases.
type UpdateIndexAliasesRequest struct {
	Request
	ClusterId string        `json:"-"`
	IndexName string        `json:"-"`
	Actions   []AliasAction `json:"actions"`
}

// IndexTemplateItem describes a single index template as returned by the template list API.
type IndexTemplateItem struct {
	TemplateName  string `json:"templateName,omitempty"`
	IndexPatterns string `json:"indexPatterns,omitempty"`
	Priority      string `json:"priority,omitempty"`
	Order         string `json:"order,omitempty"`
}

// ListIndexTemplatesRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/index-templates.
type ListIndexTemplatesRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ListIndexTemplatesResponse contains the response body for
// GET /v3/clusters/{clusterId}/index-management/index-templates.
type ListIndexTemplatesResponse struct {
	Templates       []IndexTemplateItem `json:"templates,omitempty"`
	LegacyTemplates []IndexTemplateItem `json:"legacyTemplates,omitempty"`
}

// GetIndexTemplateRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/index-templates/{templateName}.
type GetIndexTemplateRequest struct {
	Request
	ClusterId    string `json:"-"`
	TemplateName string `json:"-"`
}

// GetIndexTemplateResponse contains the response body for
// GET /v3/clusters/{clusterId}/index-management/index-templates/{templateName}.
type GetIndexTemplateResponse struct {
	TemplateName  string                 `json:"templateName,omitempty"`
	IsLegacy      string                 `json:"isLegacy,omitempty"`
	Type          string                 `json:"type,omitempty"`
	IndexPatterns string                 `json:"indexPatterns,omitempty"`
	Priority      string                 `json:"priority,omitempty"`
	Order         string                 `json:"order,omitempty"`
	Template      map[string]interface{} `json:"template,omitempty"`
}

// ExistIndexTemplateRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/index-templates/{templateName}/_exist.
type ExistIndexTemplateRequest struct {
	Request
	ClusterId    string `json:"-"`
	TemplateName string `json:"-"`
}

// TemplateNameResponse is shared by the index template mutation APIs whose response
// body is only the template name: create, update, delete index template.
type TemplateNameResponse struct {
	TemplateName string `json:"templateName,omitempty"`
}

// CreateIndexTemplateRequest contains the request body for
// POST /v3/clusters/{clusterId}/index-management/index-templates.
type CreateIndexTemplateRequest struct {
	Request
	ClusterId     string                 `json:"-"`
	TemplateName  string                 `json:"templateName,omitempty"`
	IndexPatterns []string               `json:"indexPatterns"`
	Priority      string                 `json:"priority,omitempty"`
	Template      map[string]interface{} `json:"template,omitempty"`
}

// UpdateIndexTemplateRequest contains the request body for
// PUT /v3/clusters/{clusterId}/index-management/index-templates/{templateName}.
type UpdateIndexTemplateRequest struct {
	Request
	ClusterId     string                 `json:"-"`
	TemplateName  string                 `json:"-"`
	IndexPatterns []string               `json:"indexPatterns"`
	Priority      string                 `json:"priority,omitempty"`
	Template      map[string]interface{} `json:"template,omitempty"`
}

// DeleteIndexTemplateRequest contains the path variables and optional query filter for
// DELETE /v3/clusters/{clusterId}/index-management/index-templates/{templateName}.
type DeleteIndexTemplateRequest struct {
	Request
	ClusterId    string `json:"-"`
	TemplateName string `json:"-"`
	IsLegacy     string `json:"-"`
}

// IsmTemplate describes the index pattern matching scope of an ISM policy.
type IsmTemplate struct {
	IndexPatterns []string `json:"indexPatterns,omitempty"`
	Priority      string   `json:"priority,omitempty"`
}

// IsmRollover describes the rollover stage configuration of an ISM policy.
type IsmRollover struct {
	Enabled     *bool  `json:"enabled,omitempty"`
	MinIndexAge string `json:"minIndexAge,omitempty"`
	MinDocCount string `json:"minDocCount,omitempty"`
	MinSize     string `json:"minSize,omitempty"`
}

// IsmWarm describes the warm stage configuration of an ISM policy.
type IsmWarm struct {
	Enabled     *bool  `json:"enabled,omitempty"`
	MinIndexAge string `json:"minIndexAge,omitempty"`
}

// IsmCold describes the cold stage configuration of an ISM policy.
type IsmCold struct {
	Enabled     *bool  `json:"enabled,omitempty"`
	MinIndexAge string `json:"minIndexAge,omitempty"`
}

// IsmForceMerge describes the force-merge stage configuration of an ISM policy.
type IsmForceMerge struct {
	Enabled        *bool  `json:"enabled,omitempty"`
	MinIndexAge    string `json:"minIndexAge,omitempty"`
	MaxNumSegments *int   `json:"maxNumSegments,omitempty"`
}

// IsmDelete describes the delete stage configuration of an ISM policy.
type IsmDelete struct {
	Enabled     *bool  `json:"enabled,omitempty"`
	MinIndexAge string `json:"minIndexAge,omitempty"`
}

// IsmPolicyItem describes a single ISM policy as returned by the policy list API.
type IsmPolicyItem struct {
	PolicyId        string `json:"policyId,omitempty"`
	Description     string `json:"description,omitempty"`
	LastUpdatedTime string `json:"lastUpdatedTime,omitempty"`
}

// ListIsmPoliciesRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/ism-policies.
type ListIsmPoliciesRequest struct {
	Request
	ClusterId string `json:"-"`
}

// ListIsmPoliciesResponse contains the response body for
// GET /v3/clusters/{clusterId}/index-management/ism-policies.
type ListIsmPoliciesResponse struct {
	Policies []IsmPolicyItem `json:"policies,omitempty"`
}

// GetIsmPolicyRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/ism-policies/{policyId}.
type GetIsmPolicyRequest struct {
	Request
	ClusterId string `json:"-"`
	PolicyId  string `json:"-"`
}

// GetIsmPolicyResponse contains the response body for
// GET /v3/clusters/{clusterId}/index-management/ism-policies/{policyId}.
type GetIsmPolicyResponse struct {
	PolicyId        string        `json:"policyId,omitempty"`
	RawContent      string        `json:"rawContent,omitempty"`
	Description     string        `json:"description,omitempty"`
	IsmTemplate     []IsmTemplate `json:"ismTemplate,omitempty"`
	LastUpdatedTime string        `json:"lastUpdatedTime,omitempty"`
	Rollover        IsmRollover   `json:"rollover,omitempty"`
	Cold            IsmCold       `json:"cold,omitempty"`
	ForceMerge      IsmForceMerge `json:"forceMerge,omitempty"`
	Delete          IsmDelete     `json:"delete,omitempty"`
}

// ExistIsmPolicyRequest contains the path variables for
// GET /v3/clusters/{clusterId}/index-management/ism-policies/{policyId}/_exist.
type ExistIsmPolicyRequest struct {
	Request
	ClusterId string `json:"-"`
	PolicyId  string `json:"-"`
}

// PolicyIDResponse is shared by the ISM policy mutation APIs whose response body is
// only the policy id: create, update, delete ISM policy.
type PolicyIDResponse struct {
	PolicyId string `json:"policyId,omitempty"`
}

// IsmPolicyRequest is shared by the create and update ISM policy APIs, whose request
// body fields are identical.
type IsmPolicyRequest struct {
	Request
	ClusterId   string         `json:"-"`
	PolicyId    string         `json:"-"`
	Description string         `json:"description"`
	IsmTemplate []IsmTemplate  `json:"ismTemplate"`
	Rollover    *IsmRollover   `json:"rollover,omitempty"`
	Warm        *IsmWarm       `json:"warm,omitempty"`
	Cold        *IsmCold       `json:"cold,omitempty"`
	ForceMerge  *IsmForceMerge `json:"forceMerge,omitempty"`
	Delete      *IsmDelete     `json:"delete,omitempty"`
}

// DeleteIsmPolicyRequest contains the path variables for
// DELETE /v3/clusters/{clusterId}/index-management/ism-policies/{policyId}.
type DeleteIsmPolicyRequest struct {
	Request
	ClusterId string `json:"-"`
	PolicyId  string `json:"-"`
}

// IndexOptionItem describes a single monitored index as returned by the monitor index list API.
type IndexOptionItem struct {
	IndexName string `json:"indexName,omitempty"`
	Category  string `json:"category,omitempty"`
}

// ListMonitorIndicesRequest contains the path variables and optional query filters for
// GET /v3/clusters/{clusterId}/monitor/indices.
type ListMonitorIndicesRequest struct {
	Request
	ClusterId     string `json:"-"`
	IndexName     string `json:"-"`
	IncludeSystem *bool  `json:"-"`
	PageNo        int    `json:"-"`
	PageSize      int    `json:"-"`
}

// ListMonitorIndicesResponse contains the response body for
// GET /v3/clusters/{clusterId}/monitor/indices.
type ListMonitorIndicesResponse struct {
	Indices    []IndexOptionItem `json:"indices,omitempty"`
	PageNo     int               `json:"pageNo,omitempty"`
	PageSize   int               `json:"pageSize,omitempty"`
	TotalCount int               `json:"totalCount,omitempty"`
}

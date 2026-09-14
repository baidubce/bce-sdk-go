package v3

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bes/v3/api"
)

const DEFAULT_ENDPOINT = "bes." + bce.DEFAULT_REGION + ".baidubce.com"

// Client is the BES v3 service client.
type Client struct {
	*bce.BceClient
	region string
}

// NewClient creates a BES v3 service client with AK/SK credentials.
func NewClient(ak, sk, endPoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	return newClientWithCredentials(credentials, endPoint)
}

func NewClientWithSTS(ak, sk, token, endPoint string) (*Client, error) {
	credentials, err := auth.NewSessionBceCredentials(ak, sk, token)
	if err != nil {
		return nil, err
	}
	return newClientWithCredentials(credentials, endPoint)
}

func newClientWithCredentials(credentials *auth.BceCredentials, endPoint string) (*Client, error) {
	region := inferRegionFromEndpoint(endPoint)
	if region == "" {
		region = bce.DEFAULT_REGION
	}
	if endPoint == "" {
		endPoint = fmt.Sprintf("bes.%s.baidubce.com", region)
	}

	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS,
	}

	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endPoint,
		Region:                    region,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
	}

	return &Client{
		BceClient: bce.NewBceClient(defaultConf, &auth.BceV1Signer{}),
		region:    region,
	}, nil
}

/** ========================================= Cluster API ================================================================== */

// CreateCluster creates a BES cluster.
func (c *Client) CreateCluster(request *CreateClusterRequest) (*CreateClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateCluster(c.BceClient, c.region, request)
}

// ListClusters lists BES clusters.
func (c *Client) ListClusters(request *ListClustersRequest) (*ListClustersResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListClusters(c.BceClient, c.region, request)
}

// GetCluster gets BES cluster detail.
func (c *Client) GetCluster(request *GetClusterRequest) (*GetClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetCluster(c.BceClient, c.region, request)
}

// ListClusterNodes lists nodes for a BES cluster.
func (c *Client) ListClusterNodes(request *ListClusterNodesRequest) (*ListClusterNodesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListClusterNodes(c.BceClient, c.region, request)
}

// QueryAvailableSpecs queries available cluster specs.
func (c *Client) QueryAvailableSpecs(request *QueryAvailableSpecsRequest) (*QueryAvailableSpecsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.QueryAvailableSpecs(c.BceClient, c.region, request)
}

// QueryAvailableKernels queries available cluster kernels.
func (c *Client) QueryAvailableKernels(request *QueryAvailableKernelsRequest) (*QueryAvailableKernelsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.QueryAvailableKernels(c.BceClient, c.region, request)
}

// DownloadClusterCert downloads the cluster certificate zip stream. The returned response's
// Body must be closed by the caller.
func (c *Client) DownloadClusterCert(request *DownloadClusterCertRequest) (*DownloadClusterCertResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DownloadClusterCert(c.BceClient, c.region, request)
}

// UpdateClusterName updates a cluster name.
func (c *Client) UpdateClusterName(request *UpdateClusterNameRequest) (*UpdateClusterNameResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateClusterName(c.BceClient, c.region, request)
}

// UpdateClusterMaintenance updates a cluster maintenance window.
func (c *Client) UpdateClusterMaintenance(request *UpdateClusterMaintenanceRequest) (*UpdateClusterMaintenanceResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateClusterMaintenance(c.BceClient, c.region, request)
}

// UpdateClusterDeletionProtection updates deletion protection.
func (c *Client) UpdateClusterDeletionProtection(request *UpdateClusterDeletionProtectionRequest) (*UpdateClusterDeletionProtectionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateClusterDeletionProtection(c.BceClient, c.region, request)
}

// RecoverCluster recovers a deleted cluster.
func (c *Client) RecoverCluster(request *RecoverClusterRequest) (*RecoverClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.RecoverCluster(c.BceClient, c.region, request)
}

// DeleteCluster deletes a BES cluster.
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (*DeleteClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteCluster(c.BceClient, c.region, request)
}

// ResizeCluster performs a partial resize of a BES cluster.
func (c *Client) ResizeCluster(request *ResizeClusterRequest) (*ResizeClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ResizeCluster(c.BceClient, c.region, request)
}

// ListSuggestedMigrationNodes lists suggested nodes for data migration.
func (c *Client) ListSuggestedMigrationNodes(request *ListSuggestedMigrationNodesRequest) (*ListMigrationNodesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListSuggestedMigrationNodes(c.BceClient, c.region, request)
}

// ListMigratableNodes lists nodes that are eligible for data migration.
func (c *Client) ListMigratableNodes(request *ListMigratableNodesRequest) (*ListMigrationNodesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListMigratableNodes(c.BceClient, c.region, request)
}

// QueryAvailableUpgrades queries the versions and kernels a cluster can upgrade to.
func (c *Client) QueryAvailableUpgrades(request *QueryAvailableUpgradesRequest) (*QueryAvailableUpgradesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.QueryAvailableUpgrades(c.BceClient, c.region, request)
}

// UpdateClusterPublicAccess switches the cluster public access.
func (c *Client) UpdateClusterPublicAccess(request *UpdateClusterPublicAccessRequest) (*ClusterActionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateClusterPublicAccess(c.BceClient, c.region, request)
}

// UpdateClusterCerebro enables Cerebro for a cluster.
func (c *Client) UpdateClusterCerebro(request *UpdateClusterCerebroRequest) (*UpdateClusterCerebroResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateClusterCerebro(c.BceClient, c.region, request)
}

// StartCluster resumes a suspended BES cluster.
func (c *Client) StartCluster(request *StartClusterRequest) (*ClusterActionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.StartCluster(c.BceClient, c.region, request)
}

// StartClusterNodes starts the given nodes of a cluster.
func (c *Client) StartClusterNodes(request *ClusterNodesRequest) (*ClusterActionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.StartClusterNodes(c.BceClient, c.region, request)
}

// MigrateClusterNodeData migrates data off the given nodes.
func (c *Client) MigrateClusterNodeData(request *MigrateClusterNodeDataRequest) (*ClusterActionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.MigrateClusterNodeData(c.BceClient, c.region, request)
}

// UpdateClusterAccessWhitelist sets the access whitelist for a cluster.
func (c *Client) UpdateClusterAccessWhitelist(request *UpdateClusterAccessWhitelistRequest) (*UpdateClusterAccessWhitelistResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateClusterAccessWhitelist(c.BceClient, c.region, request)
}

// UpgradeCluster checks or executes an engine/kernel upgrade for a cluster.
func (c *Client) UpgradeCluster(request *UpgradeClusterRequest) (*ClusterActionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpgradeCluster(c.BceClient, c.region, request)
}

// StopCluster suspends a running BES cluster.
func (c *Client) StopCluster(request *StopClusterRequest) (*ClusterActionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.StopCluster(c.BceClient, c.region, request)
}

// StopClusterNodes stops the given nodes of a cluster.
func (c *Client) StopClusterNodes(request *ClusterNodesRequest) (*ClusterActionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.StopClusterNodes(c.BceClient, c.region, request)
}

// UpdateClusterProtocol switches HTTPS on or off for a cluster.
func (c *Client) UpdateClusterProtocol(request *UpdateClusterProtocolRequest) (*ClusterActionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateClusterProtocol(c.BceClient, c.region, request)
}

// RestartCluster restarts a cluster or a subset of its nodes.
func (c *Client) RestartCluster(request *RestartClusterRequest) (*ClusterActionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.RestartCluster(c.BceClient, c.region, request)
}

/** ========================================= Index API ==================================================================== */

// CreateIndex creates an index in a BES cluster.
func (c *Client) CreateIndex(request *CreateIndexRequest) (*CreateIndexResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateIndex(c.BceClient, c.region, request)
}

// ListIndices lists indices in a BES cluster.
func (c *Client) ListIndices(request *ListIndicesRequest) (*ListIndicesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListIndices(c.BceClient, c.region, request)
}

// GetIndex gets the detail of an index in a BES cluster.
func (c *Client) GetIndex(request *GetIndexRequest) (*GetIndexResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetIndex(c.BceClient, c.region, request)
}

// ExistIndex checks whether an index exists in a BES cluster.
func (c *Client) ExistIndex(request *ExistIndexRequest) (*ExistResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ExistIndex(c.BceClient, c.region, request)
}

// GetIndexStats gets the stats of an index in a BES cluster.
func (c *Client) GetIndexStats(request *GetIndexStatsRequest) (*IndexStatsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetIndexStats(c.BceClient, c.region, request)
}

// ListIndexFieldTypes lists the supported index field types for a BES cluster.
func (c *Client) ListIndexFieldTypes(request *ListIndexFieldTypesRequest) (*ListIndexFieldTypesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListIndexFieldTypes(c.BceClient, c.region, request)
}

// OpenIndices opens the given indices in a BES cluster.
func (c *Client) OpenIndices(request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.OpenIndices(c.BceClient, c.region, request)
}

// CloseIndices closes the given indices in a BES cluster.
func (c *Client) CloseIndices(request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CloseIndices(c.BceClient, c.region, request)
}

// DeleteIndices deletes the given indices in a BES cluster.
func (c *Client) DeleteIndices(request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteIndices(c.BceClient, c.region, request)
}

// RefreshIndices refreshes the given indices in a BES cluster.
func (c *Client) RefreshIndices(request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.RefreshIndices(c.BceClient, c.region, request)
}

// FlushIndices flushes the given indices in a BES cluster.
func (c *Client) FlushIndices(request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.FlushIndices(c.BceClient, c.region, request)
}

// ForceMergeIndices force-merges the given indices in a BES cluster.
func (c *Client) ForceMergeIndices(request *ForceMergeIndicesRequest) (*OperationSuccessResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ForceMergeIndices(c.BceClient, c.region, request)
}

// ClearIndicesCache clears the cache of the given indices in a BES cluster.
func (c *Client) ClearIndicesCache(request *IndexNamesRequest) (*OperationSuccessResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ClearIndicesCache(c.BceClient, c.region, request)
}

// UpdateIndexSettings updates the settings of an index in a BES cluster.
func (c *Client) UpdateIndexSettings(request *UpdateIndexSettingsRequest) (*IndexNameResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateIndexSettings(c.BceClient, c.region, request)
}

// UpdateIndexMappings updates the mappings of an index in a BES cluster.
func (c *Client) UpdateIndexMappings(request *UpdateIndexMappingsRequest) (*IndexNameResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateIndexMappings(c.BceClient, c.region, request)
}

// UpdateIndexAliases updates the aliases of an index in a BES cluster.
func (c *Client) UpdateIndexAliases(request *UpdateIndexAliasesRequest) (*IndexNameResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateIndexAliases(c.BceClient, c.region, request)
}

// ListIndexTemplates lists index templates in a BES cluster.
func (c *Client) ListIndexTemplates(request *ListIndexTemplatesRequest) (*ListIndexTemplatesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListIndexTemplates(c.BceClient, c.region, request)
}

// GetIndexTemplate gets the detail of an index template in a BES cluster.
func (c *Client) GetIndexTemplate(request *GetIndexTemplateRequest) (*GetIndexTemplateResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetIndexTemplate(c.BceClient, c.region, request)
}

// ExistIndexTemplate checks whether an index template exists in a BES cluster.
func (c *Client) ExistIndexTemplate(request *ExistIndexTemplateRequest) (*ExistResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ExistIndexTemplate(c.BceClient, c.region, request)
}

// CreateIndexTemplate creates an index template in a BES cluster.
func (c *Client) CreateIndexTemplate(request *CreateIndexTemplateRequest) (*TemplateNameResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateIndexTemplate(c.BceClient, c.region, request)
}

// UpdateIndexTemplate updates an index template in a BES cluster.
func (c *Client) UpdateIndexTemplate(request *UpdateIndexTemplateRequest) (*TemplateNameResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateIndexTemplate(c.BceClient, c.region, request)
}

// DeleteIndexTemplate deletes an index template in a BES cluster.
func (c *Client) DeleteIndexTemplate(request *DeleteIndexTemplateRequest) (*TemplateNameResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteIndexTemplate(c.BceClient, c.region, request)
}

// ListIsmPolicies lists ISM policies in a BES cluster.
func (c *Client) ListIsmPolicies(request *ListIsmPoliciesRequest) (*ListIsmPoliciesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListIsmPolicies(c.BceClient, c.region, request)
}

// GetIsmPolicy gets the detail of an ISM policy in a BES cluster.
func (c *Client) GetIsmPolicy(request *GetIsmPolicyRequest) (*GetIsmPolicyResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetIsmPolicy(c.BceClient, c.region, request)
}

// ExistIsmPolicy checks whether an ISM policy exists in a BES cluster.
func (c *Client) ExistIsmPolicy(request *ExistIsmPolicyRequest) (*ExistResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ExistIsmPolicy(c.BceClient, c.region, request)
}

// CreateIsmPolicy creates an ISM policy in a BES cluster.
func (c *Client) CreateIsmPolicy(request *IsmPolicyRequest) (*PolicyIDResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateIsmPolicy(c.BceClient, c.region, request)
}

// UpdateIsmPolicy updates an ISM policy in a BES cluster.
func (c *Client) UpdateIsmPolicy(request *IsmPolicyRequest) (*PolicyIDResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateIsmPolicy(c.BceClient, c.region, request)
}

// DeleteIsmPolicy deletes an ISM policy in a BES cluster.
func (c *Client) DeleteIsmPolicy(request *DeleteIsmPolicyRequest) (*PolicyIDResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteIsmPolicy(c.BceClient, c.region, request)
}

// ListMonitorIndices lists monitored indices in a BES cluster.
func (c *Client) ListMonitorIndices(request *ListMonitorIndicesRequest) (*ListMonitorIndicesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListMonitorIndices(c.BceClient, c.region, request)
}

/** ========================================= Billing API ================================================================== */

// ConvertToPrepay converts a postpaid cluster to prepaid billing.
func (c *Client) ConvertToPrepay(request *ConvertToPrepayRequest) (*OrderIdResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ConvertToPrepay(c.BceClient, c.region, request)
}

// ConvertToPostpay converts a prepaid cluster to postpaid billing.
func (c *Client) ConvertToPostpay(request *ConvertToPostpayRequest) (*OrderIdResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ConvertToPostpay(c.BceClient, c.region, request)
}

// CancelConvertToPostpay cancels a pending convert-to-postpay order.
func (c *Client) CancelConvertToPostpay(request *CancelConvertToPostpayRequest) (*OrderIdResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CancelConvertToPostpay(c.BceClient, c.region, request)
}

// RenewCluster creates a renewal order for a prepaid cluster.
func (c *Client) RenewCluster(request *RenewClusterRequest) (*OrderIdResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.RenewCluster(c.BceClient, c.region, request)
}

// QueryConfigPrice queries the price of a hypothetical cluster configuration before purchase.
func (c *Client) QueryConfigPrice(request *QueryConfigPriceRequest) (*QueryConfigPriceResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.QueryConfigPrice(c.BceClient, c.region, request)
}

// QueryClusterMarginPrice queries the price delta for resizing an existing cluster.
func (c *Client) QueryClusterMarginPrice(request *QueryClusterMarginPriceRequest) (*QueryClusterMarginPriceResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.QueryClusterMarginPrice(c.BceClient, c.region, request)
}

/** ========================================= Action API =================================================================== */

// ListActions lists BES task actions.
func (c *Client) ListActions(request *ListActionsRequest) (*ListActionsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListActions(c.BceClient, c.region, request)
}

// ListActionTypes lists BES task action types.
func (c *Client) ListActionTypes(request *ListActionTypesRequest) (*ListActionTypesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListActionTypes(c.BceClient, c.region, request)
}

// ListOperations lists operations for a BES task action.
func (c *Client) ListOperations(request *ListOperationsRequest) (*ListOperationsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListOperations(c.BceClient, c.region, request)
}

// GetOperation gets the detail of one BES task operation.
func (c *Client) GetOperation(request *GetOperationRequest) (*GetOperationResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetOperation(c.BceClient, c.region, request)
}

// ListOperationAnalysisDetails lists analysis details for a task operation group.
func (c *Client) ListOperationAnalysisDetails(request *ListOperationAnalysisDetailsRequest) (*ListOperationAnalysisDetailsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListOperationAnalysisDetails(c.BceClient, c.region, request)
}

/** ========================================= Audit API ==================================================================== */

// GetAuditStatus queries the data-plane audit status of a cluster.
func (c *Client) GetAuditStatus(request *GetAuditStatusRequest) (*AuditStatusResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetAuditStatus(c.BceClient, c.region, request)
}

// EnableAudit enables data-plane audit for a cluster.
func (c *Client) EnableAudit(request *EnableAuditRequest) (*EnableAuditResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.EnableAudit(c.BceClient, c.region, request)
}

// ListAuditEventTypes lists the searchable audit event types of a cluster.
func (c *Client) ListAuditEventTypes(request *ListAuditEventTypesRequest) (*ListAuditEventTypesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListAuditEventTypes(c.BceClient, c.region, request)
}

// SearchAuditEvents searches data-plane audit events of a cluster.
func (c *Client) SearchAuditEvents(request *SearchAuditEventsRequest) (*SearchAuditEventsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.SearchAuditEvents(c.BceClient, c.region, request)
}

// GetBctAuthSwitch queries the cloud audit (BCT) authorization switch status.
func (c *Client) GetBctAuthSwitch(request *Request) (*GetBctAuthSwitchResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetBctAuthSwitch(c.BceClient, c.region, request)
}

// UpdateBctAuthSwitch updates the cloud audit (BCT) authorization switch status.
func (c *Client) UpdateBctAuthSwitch(request *UpdateBctAuthSwitchRequest) (*BctAuthSwitchResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateBctAuthSwitch(c.BceClient, c.region, request)
}

// SearchBctEvents searches cloud audit (BCT) control-plane events.
func (c *Client) SearchBctEvents(request *SearchBctEventsRequest) (*SearchBctEventsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.SearchBctEvents(c.BceClient, c.region, request)
}

/** ========================================= Log API ======================================================================= */

// GetLogSwitch queries the log collection switch state of a cluster.
func (c *Client) GetLogSwitch(request *GetLogSwitchRequest) (*LogSwitchResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetLogSwitch(c.BceClient, c.region, request)
}

// UpdateLogSwitch updates the log collection switch state of a cluster.
func (c *Client) UpdateLogSwitch(request *UpdateLogSwitchRequest) (*LogSwitchResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateLogSwitch(c.BceClient, c.region, request)
}

// GetLogUsage queries aggregated log index usage for a cluster.
func (c *Client) GetLogUsage(request *GetLogUsageRequest) (*LogUsageResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetLogUsage(c.BceClient, c.region, request)
}

// GetLogCollectorStatus queries the log collector deployment and operation state.
func (c *Client) GetLogCollectorStatus(request *GetLogCollectorStatusRequest) (*LogCollectorStatusResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetLogCollectorStatus(c.BceClient, c.region, request)
}

// SearchLogs searches raw log records of a cluster.
func (c *Client) SearchLogs(request *SearchLogsRequest) (*SearchLogsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.SearchLogs(c.BceClient, c.region, request)
}

/** ========================================= Plugin API =================================================================== */

// ListSystemPlugins lists the system plugins available to a BES cluster.
func (c *Client) ListSystemPlugins(request *ListSystemPluginsRequest) (*ListSystemPluginsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListSystemPlugins(c.BceClient, c.region, request)
}

// UpdateSystemPlugin installs or uninstalls system plugins on a BES cluster.
func (c *Client) UpdateSystemPlugin(request *UpdateSystemPluginRequest) (*UpdateSystemPluginResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateSystemPlugin(c.BceClient, c.region, request)
}

// ListCustomPlugins lists the custom plugins installed on a BES cluster.
func (c *Client) ListCustomPlugins(request *ListCustomPluginsRequest) (*ListCustomPluginsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListCustomPlugins(c.BceClient, c.region, request)
}

// UpdateCustomPlugin installs or uninstalls custom plugins on a BES cluster.
func (c *Client) UpdateCustomPlugin(request *UpdateCustomPluginRequest) (*UpdateCustomPluginResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateCustomPlugin(c.BceClient, c.region, request)
}

// DeleteCustomPluginVersion deletes a specific version of a custom plugin from a BES cluster.
func (c *Client) DeleteCustomPluginVersion(request *DeleteCustomPluginVersionRequest) (*DeleteCustomPluginVersionResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteCustomPluginVersion(c.BceClient, c.region, request)
}

// UploadPluginFile uploads a plugin package or other plugin-management file to a BES cluster.
func (c *Client) UploadPluginFile(request *UploadPluginFileRequest) (*UploadPluginFileResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UploadPluginFile(c.BceClient, c.region, request)
}

/** ========================================= Diagnosis API ================================================================ */

// ListWeeklyOverview lists the diagnosis risk items found within the last 7 days.
func (c *Client) ListWeeklyOverview(request *ListWeeklyOverviewRequest) (*ListWeeklyOverviewResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListWeeklyOverview(c.BceClient, c.region, request)
}

// GetManualDiagnosisCount gets the number of manual diagnoses that can be or have been created.
func (c *Client) GetManualDiagnosisCount(request *GetManualDiagnosisCountRequest) (*GetManualDiagnosisCountResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetManualDiagnosisCount(c.BceClient, c.region, request)
}

// GetManualDiagnosisConfig gets the manual diagnosis configuration of a cluster.
func (c *Client) GetManualDiagnosisConfig(request *GetManualDiagnosisConfigRequest) (*GetManualDiagnosisConfigResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetManualDiagnosisConfig(c.BceClient, c.region, request)
}

// UpdateManualDiagnosisConfig updates the manual diagnosis configuration of a cluster.
func (c *Client) UpdateManualDiagnosisConfig(request *UpdateManualDiagnosisConfigRequest) (*UpdateManualDiagnosisConfigResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateManualDiagnosisConfig(c.BceClient, c.region, request)
}

// CreateManualDiagnosis creates a manual diagnosis task on a cluster.
func (c *Client) CreateManualDiagnosis(request *CreateManualDiagnosisRequest) (*CreateManualDiagnosisResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateManualDiagnosis(c.BceClient, c.region, request)
}

// ListDiagnosisReports lists the diagnosis reports of a cluster.
func (c *Client) ListDiagnosisReports(request *ListDiagnosisReportsRequest) (*ListDiagnosisReportsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListDiagnosisReports(c.BceClient, c.region, request)
}

// GetDiagnosisReport gets the detail of a diagnosis report.
func (c *Client) GetDiagnosisReport(request *GetDiagnosisReportRequest) (*GetDiagnosisReportResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetDiagnosisReport(c.BceClient, c.region, request)
}

// GetDiagnosisBusyStatus gets whether the diagnosis service is busy for a cluster.
func (c *Client) GetDiagnosisBusyStatus(request *GetDiagnosisBusyStatusRequest) (*GetDiagnosisBusyStatusResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetDiagnosisBusyStatus(c.BceClient, c.region, request)
}

// GetDiagnosisAuthorizationStatus gets whether diagnosis is authorized for a cluster.
func (c *Client) GetDiagnosisAuthorizationStatus(request *GetDiagnosisAuthorizationStatusRequest) (*GetDiagnosisAuthorizationStatusResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetDiagnosisAuthorizationStatus(c.BceClient, c.region, request)
}

// AuthorizeDiagnosis authorizes diagnosis for a cluster.
func (c *Client) AuthorizeDiagnosis(request *AuthorizeDiagnosisRequest) (*AuthorizeDiagnosisResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.AuthorizeDiagnosis(c.BceClient, c.region, request)
}

// ListDiagnosisItems lists the diagnosis items supported by the service.
func (c *Client) ListDiagnosisItems(request *ListDiagnosisItemsRequest) (*ListDiagnosisItemsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListDiagnosisItems(c.BceClient, c.region, request)
}

// GetAutoDiagnosisStatus gets the auto diagnosis switch status of a cluster.
func (c *Client) GetAutoDiagnosisStatus(request *GetAutoDiagnosisStatusRequest) (*GetAutoDiagnosisStatusResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetAutoDiagnosisStatus(c.BceClient, c.region, request)
}

// UpdateAutoDiagnosis updates the auto diagnosis switch of a cluster.
func (c *Client) UpdateAutoDiagnosis(request *UpdateAutoDiagnosisRequest) (*UpdateAutoDiagnosisResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateAutoDiagnosis(c.BceClient, c.region, request)
}

// GetLatestOverview gets the latest diagnosis overview of a cluster.
func (c *Client) GetLatestOverview(request *GetLatestOverviewRequest) (*GetLatestOverviewResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetLatestOverview(c.BceClient, c.region, request)
}

/** ========================================= Snapshot API ================================================================= */

// ListIndicesByPattern lists snapshottable indices matching a pattern.
func (c *Client) ListIndicesByPattern(request *ListIndicesByPatternRequest) (*ListIndicesByPatternResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListIndicesByPattern(c.BceClient, c.region, request)
}

// ListRestores lists restore tasks of a snapshot.
func (c *Client) ListRestores(request *ListRestoresRequest) (*ListRestoresResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListRestores(c.BceClient, c.region, request)
}

// ListRestoreClusters lists clusters eligible as restore targets for a snapshot.
func (c *Client) ListRestoreClusters(request *ListRestoreClustersRequest) (*ListRestoreClustersResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListRestoreClusters(c.BceClient, c.region, request)
}

// GetSnapshotRule queries the snapshot rule detail of a snapshot.
func (c *Client) GetSnapshotRule(request *GetSnapshotRuleRequest) (*GetSnapshotRuleResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetSnapshotRule(c.BceClient, c.region, request)
}

// ListSnapshots lists snapshots of a cluster.
func (c *Client) ListSnapshots(request *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListSnapshots(c.BceClient, c.region, request)
}

// GetSnapshotConfig queries a snapshot config's detail.
func (c *Client) GetSnapshotConfig(request *GetSnapshotConfigRequest) (*GetSnapshotConfigResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetSnapshotConfig(c.BceClient, c.region, request)
}

// GetSnapshot queries a snapshot's detail.
func (c *Client) GetSnapshot(request *GetSnapshotRequest) (*GetSnapshotResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetSnapshot(c.BceClient, c.region, request)
}

// ListAutoSnapshotConfigs lists auto snapshot configs of a cluster.
func (c *Client) ListAutoSnapshotConfigs(request *ListAutoSnapshotConfigsRequest) (*ListAutoSnapshotConfigsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListAutoSnapshotConfigs(c.BceClient, c.region, request)
}

// CreateRestore creates a restore task from a snapshot.
func (c *Client) CreateRestore(request *CreateRestoreRequest) (*CreateRestoreResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateRestore(c.BceClient, c.region, request)
}

// CreateManualSnapshot creates a manual snapshot.
func (c *Client) CreateManualSnapshot(request *CreateManualSnapshotRequest) (*CreateManualSnapshotResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateManualSnapshot(c.BceClient, c.region, request)
}

// CreateAutoSnapshotConfig creates an auto snapshot config.
func (c *Client) CreateAutoSnapshotConfig(request *CreateAutoSnapshotConfigRequest) (*CreateAutoSnapshotConfigResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateAutoSnapshotConfig(c.BceClient, c.region, request)
}

// UpdateAutoSnapshotConfig updates an auto snapshot config.
func (c *Client) UpdateAutoSnapshotConfig(request *UpdateAutoSnapshotConfigRequest) (*UpdateAutoSnapshotConfigResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateAutoSnapshotConfig(c.BceClient, c.region, request)
}

// DeleteSnapshot deletes a snapshot.
func (c *Client) DeleteSnapshot(request *DeleteSnapshotRequest) (*DeleteSnapshotResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteSnapshot(c.BceClient, c.region, request)
}

// DeleteAutoSnapshotConfig deletes an auto snapshot config.
func (c *Client) DeleteAutoSnapshotConfig(request *DeleteAutoSnapshotConfigRequest) (*DeleteAutoSnapshotConfigResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteAutoSnapshotConfig(c.BceClient, c.region, request)
}

// BatchEnableAutoSnapshotConfigs batch enables or disables auto snapshot configs of a cluster.
func (c *Client) BatchEnableAutoSnapshotConfigs(request *BatchEnableAutoSnapshotConfigsRequest) (*BatchEnableAutoSnapshotConfigsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.BatchEnableAutoSnapshotConfigs(c.BceClient, c.region, request)
}

/** ========================================= User API ===================================================================== */

// ResetAdminPassword resets the admin (superuser) password of a cluster.
func (c *Client) ResetAdminPassword(request *ResetAdminPasswordRequest) (*ResetAdminPasswordResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ResetAdminPassword(c.BceClient, c.region, request)
}

/** ========================================= Schedule API ================================================================= */

// ListSchedules lists the scheduled tasks of a cluster.
func (c *Client) ListSchedules(request *ListSchedulesRequest) (*ListSchedulesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListSchedules(c.BceClient, c.region, request)
}

// CreateSchedule creates a scheduled task on a cluster.
func (c *Client) CreateSchedule(request *CreateScheduleRequest) (*CreateScheduleResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateSchedule(c.BceClient, c.region, request)
}

// UpdateSchedule updates a scheduled task on a cluster.
func (c *Client) UpdateSchedule(request *UpdateScheduleRequest) (*UpdateScheduleResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateSchedule(c.BceClient, c.region, request)
}

// DeleteSchedule deletes a scheduled task from a cluster.
func (c *Client) DeleteSchedule(request *DeleteScheduleRequest) (*DeleteScheduleResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteSchedule(c.BceClient, c.region, request)
}

/** ========================================= General Function ================================================================= */

func inferRegionFromEndpoint(endpoint string) string {
	if endpoint == "" {
		return ""
	}

	host := endpoint
	if parsed, err := url.Parse(endpoint); err == nil && parsed.Host != "" {
		host = parsed.Host
	}
	if idx := strings.Index(host, ":"); idx != -1 {
		host = host[:idx]
	}

	parts := strings.Split(host, ".")
	if len(parts) >= 4 && parts[0] == "bes" && parts[len(parts)-2] == "baidubce" {
		return parts[1]
	}
	return ""
}

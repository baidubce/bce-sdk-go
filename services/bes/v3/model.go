package v3

import "github.com/baidubce/bce-sdk-go/services/bes/v3/api"

type Request = api.Request

type NodeSpec = api.NodeSpec
type Tag = api.Tag

type ClusterActionResponse = api.ClusterActionResponse
type ClusterBilling = api.ClusterBilling
type ClusterDetail = api.ClusterDetail
type ClusterDetailNodeSpec = api.ClusterDetailNodeSpec
type ClusterHealth = api.ClusterHealth
type ClusterNode = api.ClusterNode
type ClusterSummary = api.ClusterSummary
type DiskSpecInfo = api.DiskSpecInfo
type KernelUpgrade = api.KernelUpgrade
type KernelVersion = api.KernelVersion
type NodeAvailableSpec = api.NodeAvailableSpec
type NodeListNodeSpec = api.NodeListNodeSpec
type QueryNodeSpec = api.QueryNodeSpec
type RunningTime = api.RunningTime

type CreateClusterRequest = api.CreateClusterRequest
type CreateClusterResponse = api.CreateClusterResponse

type ListClustersRequest = api.ListClustersRequest
type ListClustersResponse = api.ListClustersResponse

type GetClusterResponse = api.GetClusterResponse
type GetClusterRequest = api.GetClusterRequest

type ListClusterNodesRequest = api.ListClusterNodesRequest
type ListClusterNodesResponse = api.ListClusterNodesResponse

type QueryAvailableSpecsRequest = api.QueryAvailableSpecsRequest
type QueryAvailableSpecsResponse = api.QueryAvailableSpecsResponse

type QueryAvailableKernelsRequest = api.QueryAvailableKernelsRequest
type QueryAvailableKernelsResponse = api.QueryAvailableKernelsResponse

type DownloadClusterCertRequest = api.DownloadClusterCertRequest
type DownloadClusterCertResponse = api.DownloadClusterCertResponse

type UpdateClusterNameRequest = api.UpdateClusterNameRequest
type UpdateClusterNameResponse = api.UpdateClusterNameResponse

type UpdateClusterMaintenanceRequest = api.UpdateClusterMaintenanceRequest
type UpdateClusterMaintenanceResponse = api.UpdateClusterMaintenanceResponse

type UpdateClusterDeletionProtectionRequest = api.UpdateClusterDeletionProtectionRequest
type UpdateClusterDeletionProtectionResponse = api.UpdateClusterDeletionProtectionResponse

type RecoverClusterRequest = api.RecoverClusterRequest
type RecoverClusterResponse = api.RecoverClusterResponse

type DeleteClusterRequest = api.DeleteClusterRequest
type DeleteClusterResponse = api.DeleteClusterResponse

type ResizeClusterRequest = api.ResizeClusterRequest
type ResizeClusterResponse = api.ResizeClusterResponse

type MigrationNode = api.MigrationNode
type ListSuggestedMigrationNodesRequest = api.ListSuggestedMigrationNodesRequest
type ListMigratableNodesRequest = api.ListMigratableNodesRequest
type ListMigrationNodesResponse = api.ListMigrationNodesResponse

type VersionUpgrade = api.VersionUpgrade
type QueryAvailableUpgradesRequest = api.QueryAvailableUpgradesRequest
type QueryAvailableUpgradesResponse = api.QueryAvailableUpgradesResponse

type UpdateClusterPublicAccessRequest = api.UpdateClusterPublicAccessRequest

type UpdateClusterCerebroRequest = api.UpdateClusterCerebroRequest
type UpdateClusterCerebroResponse = api.UpdateClusterCerebroResponse

type StartClusterRequest = api.StartClusterRequest

type ClusterNodesRequest = api.ClusterNodesRequest

type MigrateClusterNodeDataRequest = api.MigrateClusterNodeDataRequest

type UpdateClusterAccessWhitelistRequest = api.UpdateClusterAccessWhitelistRequest
type UpdateClusterAccessWhitelistResponse = api.UpdateClusterAccessWhitelistResponse

type UpgradeClusterRequest = api.UpgradeClusterRequest

type UpdateClusterProtocolRequest = api.UpdateClusterProtocolRequest

type StopClusterRequest = api.StopClusterRequest

type RestartClusterRequest = api.RestartClusterRequest

type ListIndexItem = api.ListIndexItem
type FieldTypeItemVO = api.FieldTypeItemVO
type AliasAction = api.AliasAction

type ExistResponse = api.ExistResponse
type OperationSuccessResponse = api.OperationSuccessResponse
type IndexNameResponse = api.IndexNameResponse
type IndexNamesRequest = api.IndexNamesRequest

type CreateIndexRequest = api.CreateIndexRequest
type CreateIndexResponse = api.CreateIndexResponse

type ListIndicesRequest = api.ListIndicesRequest
type ListIndicesResponse = api.ListIndicesResponse

type GetIndexRequest = api.GetIndexRequest
type GetIndexResponse = api.GetIndexResponse

type ExistIndexRequest = api.ExistIndexRequest

type GetIndexStatsRequest = api.GetIndexStatsRequest
type IndexStatsResponse = api.IndexStatsResponse

type ListIndexFieldTypesRequest = api.ListIndexFieldTypesRequest
type ListIndexFieldTypesResponse = api.ListIndexFieldTypesResponse

type ForceMergeIndicesRequest = api.ForceMergeIndicesRequest

type UpdateIndexSettingsRequest = api.UpdateIndexSettingsRequest

type UpdateIndexMappingsRequest = api.UpdateIndexMappingsRequest

type UpdateIndexAliasesRequest = api.UpdateIndexAliasesRequest

type IndexTemplateItem = api.IndexTemplateItem
type ListIndexTemplatesRequest = api.ListIndexTemplatesRequest
type ListIndexTemplatesResponse = api.ListIndexTemplatesResponse

type GetIndexTemplateRequest = api.GetIndexTemplateRequest
type GetIndexTemplateResponse = api.GetIndexTemplateResponse

type ExistIndexTemplateRequest = api.ExistIndexTemplateRequest

type TemplateNameResponse = api.TemplateNameResponse

type CreateIndexTemplateRequest = api.CreateIndexTemplateRequest

type UpdateIndexTemplateRequest = api.UpdateIndexTemplateRequest

type DeleteIndexTemplateRequest = api.DeleteIndexTemplateRequest

type IsmTemplate = api.IsmTemplate
type IsmRollover = api.IsmRollover
type IsmWarm = api.IsmWarm
type IsmCold = api.IsmCold
type IsmForceMerge = api.IsmForceMerge
type IsmDelete = api.IsmDelete
type IsmPolicyItem = api.IsmPolicyItem

type ListIsmPoliciesRequest = api.ListIsmPoliciesRequest
type ListIsmPoliciesResponse = api.ListIsmPoliciesResponse

type GetIsmPolicyRequest = api.GetIsmPolicyRequest
type GetIsmPolicyResponse = api.GetIsmPolicyResponse

type ExistIsmPolicyRequest = api.ExistIsmPolicyRequest

type PolicyIDResponse = api.PolicyIDResponse

type IsmPolicyRequest = api.IsmPolicyRequest

type DeleteIsmPolicyRequest = api.DeleteIsmPolicyRequest

type IndexOptionItem = api.IndexOptionItem
type ListMonitorIndicesRequest = api.ListMonitorIndicesRequest
type ListMonitorIndicesResponse = api.ListMonitorIndicesResponse

type OrderIdResponse = api.OrderIdResponse

type ConvertToPrepayRequest = api.ConvertToPrepayRequest

type ConvertToPostpayRequest = api.ConvertToPostpayRequest

type CancelConvertToPostpayRequest = api.CancelConvertToPostpayRequest

type RenewClusterRequest = api.RenewClusterRequest

type BillingConfig = api.BillingConfig
type ComponentPrice = api.ComponentPrice
type QueryNodeTypePrice = api.QueryNodeTypePrice

type QueryConfigPriceRequest = api.QueryConfigPriceRequest
type QueryConfigPriceResponse = api.QueryConfigPriceResponse

type QueryClusterMarginPriceRequest = api.QueryClusterMarginPriceRequest
type QueryClusterMarginPriceResponse = api.QueryClusterMarginPriceResponse

type ListActionsRequest = api.ListActionsRequest
type ListActionsResponse = api.ListActionsResponse

type ActionItem = api.ActionItem

type ListActionTypesRequest = api.ListActionTypesRequest
type ListActionTypesResponse = api.ListActionTypesResponse

type ActionTypeItem = api.ActionTypeItem

type ListOperationsRequest = api.ListOperationsRequest
type ListOperationsResponse = api.ListOperationsResponse

type OperationItem = api.OperationItem

type GetOperationRequest = api.GetOperationRequest
type GetOperationResponse = api.GetOperationResponse

type GroupItem = api.GroupItem

type ListOperationAnalysisDetailsRequest = api.ListOperationAnalysisDetailsRequest
type ListOperationAnalysisDetailsResponse = api.ListOperationAnalysisDetailsResponse

type KindItem = api.KindItem

type GetAuditStatusRequest = api.GetAuditStatusRequest
type AuditStatusResponse = api.AuditStatusResponse

type EnableAuditRequest = api.EnableAuditRequest
type EnableAuditResponse = api.EnableAuditResponse

type ListAuditEventTypesRequest = api.ListAuditEventTypesRequest
type ListAuditEventTypesResponse = api.ListAuditEventTypesResponse

type SearchAuditEventsRequest = api.SearchAuditEventsRequest
type SearchAuditEventsResponse = api.SearchAuditEventsResponse

type GetBctAuthSwitchResponse = api.GetBctAuthSwitchResponse

type UpdateBctAuthSwitchRequest = api.UpdateBctAuthSwitchRequest
type BctAuthSwitchResponse = api.BctAuthSwitchResponse

type SearchBctEventsRequest = api.SearchBctEventsRequest
type SearchBctEventsResponse = api.SearchBctEventsResponse

type BctQueryFilter = api.BctQueryFilter
type BctEvent = api.BctEvent
type BctUserIdentity = api.BctUserIdentity
type AuditEvent = api.AuditEvent

type GetLogSwitchRequest = api.GetLogSwitchRequest
type UpdateLogSwitchRequest = api.UpdateLogSwitchRequest
type LogSwitchResponse = api.LogSwitchResponse

type GetLogUsageRequest = api.GetLogUsageRequest
type LogTypeUsage = api.LogTypeUsage
type LogUsageResponse = api.LogUsageResponse

type GetLogCollectorStatusRequest = api.GetLogCollectorStatusRequest
type LogCollectorStatusResponse = api.LogCollectorStatusResponse

type SearchLogsRequest = api.SearchLogsRequest
type SearchLogsResponse = api.SearchLogsResponse

type IndexItem = api.IndexItem
type RestoreIndexInfo = api.RestoreIndexInfo
type RestoreDetailItem = api.RestoreDetailItem
type RestoreItem = api.RestoreItem
type RestoreClusterItem = api.RestoreClusterItem
type SnapshotRuleIndexItem = api.SnapshotRuleIndexItem
type SnapshotDetailItem = api.SnapshotDetailItem
type SnapshotItem = api.SnapshotItem
type AutoSnapshotConfigItem = api.AutoSnapshotConfigItem

type ListIndicesByPatternRequest = api.ListIndicesByPatternRequest
type ListIndicesByPatternResponse = api.ListIndicesByPatternResponse

type ListRestoresRequest = api.ListRestoresRequest
type ListRestoresResponse = api.ListRestoresResponse

type ListRestoreClustersRequest = api.ListRestoreClustersRequest
type ListRestoreClustersResponse = api.ListRestoreClustersResponse

type GetSnapshotRuleRequest = api.GetSnapshotRuleRequest
type GetSnapshotRuleResponse = api.GetSnapshotRuleResponse

type ListSnapshotsRequest = api.ListSnapshotsRequest
type ListSnapshotsResponse = api.ListSnapshotsResponse

type GetSnapshotConfigRequest = api.GetSnapshotConfigRequest
type GetSnapshotConfigResponse = api.GetSnapshotConfigResponse

type GetSnapshotRequest = api.GetSnapshotRequest
type GetSnapshotResponse = api.GetSnapshotResponse

type ListAutoSnapshotConfigsRequest = api.ListAutoSnapshotConfigsRequest
type ListAutoSnapshotConfigsResponse = api.ListAutoSnapshotConfigsResponse

type CreateRestoreRequest = api.CreateRestoreRequest
type CreateRestoreResponse = api.CreateRestoreResponse

type CreateManualSnapshotRequest = api.CreateManualSnapshotRequest
type CreateManualSnapshotResponse = api.CreateManualSnapshotResponse

type CreateAutoSnapshotConfigRequest = api.CreateAutoSnapshotConfigRequest
type CreateAutoSnapshotConfigResponse = api.CreateAutoSnapshotConfigResponse

type UpdateAutoSnapshotConfigRequest = api.UpdateAutoSnapshotConfigRequest
type UpdateAutoSnapshotConfigResponse = api.UpdateAutoSnapshotConfigResponse

type DeleteSnapshotRequest = api.DeleteSnapshotRequest
type DeleteSnapshotResponse = api.DeleteSnapshotResponse

type DeleteAutoSnapshotConfigRequest = api.DeleteAutoSnapshotConfigRequest
type DeleteAutoSnapshotConfigResponse = api.DeleteAutoSnapshotConfigResponse

type BatchEnableAutoSnapshotConfigsRequest = api.BatchEnableAutoSnapshotConfigsRequest
type BatchEnableAutoSnapshotConfigsResponse = api.BatchEnableAutoSnapshotConfigsResponse

type PluginItem = api.PluginItem

type SystemPluginVersionInfo = api.SystemPluginVersionInfo
type SystemPluginInfo = api.SystemPluginInfo

type ListSystemPluginsRequest = api.ListSystemPluginsRequest
type ListSystemPluginsResponse = api.ListSystemPluginsResponse

type UpdateSystemPluginRequest = api.UpdateSystemPluginRequest
type UpdateSystemPluginResponse = api.UpdateSystemPluginResponse

type CustomPluginVersionInfo = api.CustomPluginVersionInfo
type CustomPluginInfo = api.CustomPluginInfo

type ListCustomPluginsRequest = api.ListCustomPluginsRequest
type ListCustomPluginsResponse = api.ListCustomPluginsResponse

type UpdateCustomPluginRequest = api.UpdateCustomPluginRequest
type UpdateCustomPluginResponse = api.UpdateCustomPluginResponse

type DeleteCustomPluginVersionRequest = api.DeleteCustomPluginVersionRequest
type DeleteCustomPluginVersionResponse = api.DeleteCustomPluginVersionResponse

type UploadPluginFileRequest = api.UploadPluginFileRequest
type UploadPluginFileResponse = api.UploadPluginFileResponse

type WeeklyRiskItem = api.WeeklyRiskItem

type ListWeeklyOverviewRequest = api.ListWeeklyOverviewRequest
type ListWeeklyOverviewResponse = api.ListWeeklyOverviewResponse

type GetManualDiagnosisCountRequest = api.GetManualDiagnosisCountRequest
type GetManualDiagnosisCountResponse = api.GetManualDiagnosisCountResponse

type GetManualDiagnosisConfigRequest = api.GetManualDiagnosisConfigRequest
type GetManualDiagnosisConfigResponse = api.GetManualDiagnosisConfigResponse

type UpdateManualDiagnosisConfigRequest = api.UpdateManualDiagnosisConfigRequest
type UpdateManualDiagnosisConfigResponse = api.UpdateManualDiagnosisConfigResponse

type CreateManualDiagnosisRequest = api.CreateManualDiagnosisRequest
type CreateManualDiagnosisResponse = api.CreateManualDiagnosisResponse

type DiagnosisReportListItem = api.DiagnosisReportListItem

type ListDiagnosisReportsRequest = api.ListDiagnosisReportsRequest
type ListDiagnosisReportsResponse = api.ListDiagnosisReportsResponse

type DiagnosisReportResultItem = api.DiagnosisReportResultItem
type DiagnosisReportGroupItem = api.DiagnosisReportGroupItem

type GetDiagnosisReportRequest = api.GetDiagnosisReportRequest
type GetDiagnosisReportResponse = api.GetDiagnosisReportResponse

type GetDiagnosisBusyStatusRequest = api.GetDiagnosisBusyStatusRequest
type GetDiagnosisBusyStatusResponse = api.GetDiagnosisBusyStatusResponse

type GetDiagnosisAuthorizationStatusRequest = api.GetDiagnosisAuthorizationStatusRequest
type GetDiagnosisAuthorizationStatusResponse = api.GetDiagnosisAuthorizationStatusResponse

type AuthorizeDiagnosisRequest = api.AuthorizeDiagnosisRequest
type AuthorizeDiagnosisResponse = api.AuthorizeDiagnosisResponse

type DiagnosisItem = api.DiagnosisItem

type ListDiagnosisItemsRequest = api.ListDiagnosisItemsRequest
type ListDiagnosisItemsResponse = api.ListDiagnosisItemsResponse

type GetAutoDiagnosisStatusRequest = api.GetAutoDiagnosisStatusRequest
type GetAutoDiagnosisStatusResponse = api.GetAutoDiagnosisStatusResponse

type UpdateAutoDiagnosisRequest = api.UpdateAutoDiagnosisRequest
type UpdateAutoDiagnosisResponse = api.UpdateAutoDiagnosisResponse

type GetLatestOverviewRequest = api.GetLatestOverviewRequest
type GetLatestOverviewResponse = api.GetLatestOverviewResponse

type ResetAdminPasswordRequest = api.ResetAdminPasswordRequest
type ResetAdminPasswordResponse = api.ResetAdminPasswordResponse

type ScheduleItem = api.ScheduleItem

type ListSchedulesRequest = api.ListSchedulesRequest
type ListSchedulesResponse = api.ListSchedulesResponse

type CreateScheduleRequest = api.CreateScheduleRequest
type CreateScheduleResponse = api.CreateScheduleResponse

type UpdateScheduleRequest = api.UpdateScheduleRequest
type UpdateScheduleResponse = api.UpdateScheduleResponse

type DeleteScheduleRequest = api.DeleteScheduleRequest
type DeleteScheduleResponse = api.DeleteScheduleResponse

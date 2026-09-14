package v2

import "github.com/baidubce/bce-sdk-go/services/bes/v2/api"

type ModuleInfo = api.ModuleInfo
type DiskSlotInfo = api.DiskSlotInfo
type AutoRenewInfo = api.AutoRenewInfo
type Billing = api.Billing
type Tag = api.Tag
type OrderResult = api.OrderResult

type ClusterListItem = api.ClusterListItem
type ClusterHealthSummary = api.ClusterHealthSummary
type BillingSummary = api.BillingSummary
type ResGroup = api.ResGroup
type ClusterSubnetSummary = api.ClusterSubnetSummary
type NetworkSummary = api.NetworkSummary
type NetworkInfo = api.NetworkInfo
type ClusterPage = api.ClusterPage

type ClusterDetail = api.ClusterDetail
type ClusterModuleDetail = api.ClusterModuleDetail
type ClusterInstance = api.ClusterInstance
type ClusterHealthDetail = api.ClusterHealthDetail
type NodeDiskMaxSizeInGB = api.NodeDiskMaxSizeInGB
type ClusterLogSetting = api.ClusterLogSetting

type ResizeModuleInfo = api.ResizeModuleInfo
type ResizeDiskSlotInfo = api.ResizeDiskSlotInfo
type AddModuleInfo = api.AddModuleInfo
type AddModuleDiskSlotInfo = api.AddModuleDiskSlotInfo

type ClusterTasksResult = api.ClusterTasksResult
type AppTask = api.AppTask
type SubTask = api.SubTask
type TaskCheckItem = api.TaskCheckItem

type DataSizeTendencyResult = api.DataSizeTendencyResult
type DataSizeItem = api.DataSizeItem
type DataSizeIndex = api.DataSizeIndex

type CouponsResult = api.CouponsResult
type Coupon = api.Coupon

type AssessResult = api.AssessResult
type AssessZone = api.AssessZone
type AssessPackageStatus = api.AssessPackageStatus
type AssessOtherConfig = api.AssessOtherConfig

type CreateClusterRequest = api.CreateClusterRequest
type CreateClusterResponse = api.CreateClusterResponse

type ListClustersRequest = api.ListClustersRequest
type ListClustersResponse = api.ListClustersResponse

type GetClusterDetailRequest = api.GetClusterDetailRequest
type GetClusterDetailResponse = api.GetClusterDetailResponse

type DeleteClusterRequest = api.DeleteClusterRequest
type DeleteClusterResponse = api.DeleteClusterResponse

type StartClusterRequest = api.StartClusterRequest
type StartClusterResponse = api.StartClusterResponse

type StopClusterRequest = api.StopClusterRequest
type StopClusterResponse = api.StopClusterResponse

type RestartClusterRequest = api.RestartClusterRequest
type RestartClusterResponse = api.RestartClusterResponse

type ResizeClusterRequest = api.ResizeClusterRequest
type ResizeClusterResponse = api.ResizeClusterResponse

type AddClusterModuleRequest = api.AddClusterModuleRequest
type AddClusterModuleResponse = api.AddClusterModuleResponse

type ResetClusterPasswordRequest = api.ResetClusterPasswordRequest
type ResetClusterPasswordResponse = api.ResetClusterPasswordResponse

type ToggleClusterHTTPSRequest = api.ToggleClusterHTTPSRequest
type ToggleClusterHTTPSResponse = api.ToggleClusterHTTPSResponse

type BindClusterEIPRequest = api.BindClusterEIPRequest
type BindClusterEIPResponse = api.BindClusterEIPResponse

type UnbindClusterEIPRequest = api.UnbindClusterEIPRequest
type UnbindClusterEIPResponse = api.UnbindClusterEIPResponse

type ToggleClusterMonitorRequest = api.ToggleClusterMonitorRequest
type ToggleClusterMonitorResponse = api.ToggleClusterMonitorResponse

type GetClusterTasksRequest = api.GetClusterTasksRequest
type GetClusterTasksResponse = api.GetClusterTasksResponse

type GetClusterDataSizeTendencyRequest = api.GetClusterDataSizeTendencyRequest
type GetClusterDataSizeTendencyResponse = api.GetClusterDataSizeTendencyResponse

type ListAvailableCouponsResponse = api.ListAvailableCouponsResponse

type AssessClusterSourceRequest = api.AssessClusterSourceRequest
type AssessClusterSourceResponse = api.AssessClusterSourceResponse

type TagListResponse = api.TagListResponse

type UpdateClusterTagsRequest = api.UpdateClusterTagsRequest
type UpdateClusterTagsResponse = api.UpdateClusterTagsResponse

type BatchInsertTagsRequest = api.BatchInsertTagsRequest
type BatchInsertTagsResponse = api.BatchInsertTagsResponse

type ClusterConfigInfo = api.ClusterConfigInfo

type GetClusterConfigRequest = api.GetClusterConfigRequest
type GetClusterConfigResponse = api.GetClusterConfigResponse

type UpdateClusterConfigRequest = api.UpdateClusterConfigRequest
type UpdateClusterConfigResponse = api.UpdateClusterConfigResponse

type SynonymDict = api.SynonymDict

type ListSynonymDictsRequest = api.ListSynonymDictsRequest
type ListSynonymDictsResponse = api.ListSynonymDictsResponse

type DeleteSynonymDictRequest = api.DeleteSynonymDictRequest
type DeleteSynonymDictResponse = api.DeleteSynonymDictResponse

type UploadSynonymDictRequest = api.UploadSynonymDictRequest
type UploadSynonymDictResponse = api.UploadSynonymDictResponse

type InstanceOperationRequest = api.InstanceOperationRequest
type InstanceOperationResponse = api.InstanceOperationResponse

type BatchInstanceOperationRequest = api.BatchInstanceOperationRequest
type BatchInstanceOperationResponse = api.BatchInstanceOperationResponse

type DeleteInstancesRequest = api.DeleteInstancesRequest
type DeleteInstancesResponse = api.DeleteInstancesResponse

type ListScaleInInstancesRequest = api.ListScaleInInstancesRequest
type ListScaleInInstancesResponse = api.ListScaleInInstancesResponse
type ScaleInInstancesResult = api.ScaleInInstancesResult
type ScaleInInstanceInfo = api.ScaleInInstanceInfo
type ScaleInInstanceItem = api.ScaleInInstanceItem

type ConfirmDataMigrationRequest = api.ConfirmDataMigrationRequest
type ConfirmDataMigrationResponse = api.ConfirmDataMigrationResponse

type RollbackDataMigrationRequest = api.RollbackDataMigrationRequest
type RollbackDataMigrationResponse = api.RollbackDataMigrationResponse

type ListDataMigrationInstancesRequest = api.ListDataMigrationInstancesRequest
type ListDataMigrationInstancesResponse = api.ListDataMigrationInstancesResponse
type MigrationInstancesResult = api.MigrationInstancesResult
type MigrationInstanceInfo = api.MigrationInstanceInfo
type MigrationInstanceItem = api.MigrationInstanceItem

type SuggestDataMigrationInstancesRequest = api.SuggestDataMigrationInstancesRequest
type SuggestDataMigrationInstancesResponse = api.SuggestDataMigrationInstancesResponse
type MigrationSuggestResult = api.MigrationSuggestResult
type MigrationSuggestInfo = api.MigrationSuggestInfo
type MigrationSuggestItem = api.MigrationSuggestItem

type UpdateLogSettingsRequest = api.UpdateLogSettingsRequest
type UpdateLogSettingsResponse = api.UpdateLogSettingsResponse

type SearchLogRequest = api.SearchLogRequest
type SearchLogResponse = api.SearchLogResponse
type LogPage = api.LogPage
type LogEntry = api.LogEntry

type CreateLogExportTaskRequest = api.CreateLogExportTaskRequest
type CreateLogExportTaskResponse = api.CreateLogExportTaskResponse
type LogExportTasksResult = api.LogExportTasksResult
type LogExportTask = api.LogExportTask

type GetLogExportRecordRequest = api.GetLogExportRecordRequest
type GetLogExportRecordResponse = api.GetLogExportRecordResponse
type LogExportRecordResult = api.LogExportRecordResult
type LogExportTaskConfig = api.LogExportTaskConfig

type DefaultPluginRequest = api.DefaultPluginRequest
type DefaultPluginResponse = api.DefaultPluginResponse

type CustomPluginRequest = api.CustomPluginRequest
type CustomPluginResponse = api.CustomPluginResponse

type DeleteCustomPluginResponse = api.DeleteCustomPluginResponse

type UploadCustomPluginRequest = api.UploadCustomPluginRequest
type UploadCustomPluginResponse = api.UploadCustomPluginResponse

type GetPluginInfoRequest = api.GetPluginInfoRequest
type GetPluginInfoResponse = api.GetPluginInfoResponse
type PluginInfoResult = api.PluginInfoResult
type PluginItem = api.PluginItem
type DefaultPluginItem = api.DefaultPluginItem
type PluginOperation = api.PluginOperation

type GetNLPDictRequest = api.GetNLPDictRequest
type GetNLPDictResponse = api.GetNLPDictResponse
type NLPDictResult = api.NLPDictResult

type UpdateNLPDictRequest = api.UpdateNLPDictRequest
type UpdateNLPDictResponse = api.UpdateNLPDictResponse
type NLPDictUpdateResult = api.NLPDictUpdateResult

type AuthorizeInspectRequest = api.AuthorizeInspectRequest
type AuthorizeInspectResponse = api.AuthorizeInspectResponse

type SuccessFlag = api.SuccessFlag

const (
	SuccessFlagTrue  = api.SuccessFlagTrue
	SuccessFlagFalse = api.SuccessFlagFalse
)

type SwitchAutoInspectRequest = api.SwitchAutoInspectRequest
type InspectStringSuccessCommonResponse = api.InspectStringSuccessCommonResponse

type CheckAutoInspectRequest = api.CheckAutoInspectRequest
type CheckAutoInspectResponse = api.CheckAutoInspectResponse
type AutoInspectStatus = api.AutoInspectStatus

type CreateManualInspectTaskRequest = api.CreateManualInspectTaskRequest

type CheckInspectBusyRequest = api.CheckInspectBusyRequest
type CheckInspectBusyResponse = api.CheckInspectBusyResponse
type InspectBusyStatus = api.InspectBusyStatus

type GetManualInspectCountRequest = api.GetManualInspectCountRequest
type GetManualInspectCountResponse = api.GetManualInspectCountResponse
type ManualInspectCount = api.ManualInspectCount

type GetManualInspectConfigRequest = api.GetManualInspectConfigRequest
type GetManualInspectConfigResponse = api.GetManualInspectConfigResponse
type ManualInspectConfig = api.ManualInspectConfig

type UpdateManualInspectConfigRequest = api.UpdateManualInspectConfigRequest

type ListInspectItemsResponse = api.ListInspectItemsResponse
type InspectItemList = api.InspectItemList
type InspectItem = api.InspectItem

type GetInspectTaskRequest = api.GetInspectTaskRequest
type GetInspectTaskResponse = api.GetInspectTaskResponse
type InspectTaskDetail = api.InspectTaskDetail
type InspectResultGroup = api.InspectResultGroup
type InspectItemResult = api.InspectItemResult

type ListInspectTasksRequest = api.ListInspectTasksRequest
type ListInspectTasksResponse = api.ListInspectTasksResponse
type InspectTaskPage = api.InspectTaskPage
type InspectTaskItem = api.InspectTaskItem

type GetLatestInspectOverviewRequest = api.GetLatestInspectOverviewRequest
type GetLatestInspectOverviewResponse = api.GetLatestInspectOverviewResponse
type LatestInspectOverview = api.LatestInspectOverview

type GetWeeklyInspectOverviewRequest = api.GetWeeklyInspectOverviewRequest
type GetWeeklyInspectOverviewResponse = api.GetWeeklyInspectOverviewResponse
type WeeklyInspectOverview = api.WeeklyInspectOverview
type InspectRiskItem = api.InspectRiskItem

type CreateScheduleRequest = api.CreateScheduleRequest
type ScheduleStringResultResponse = api.ScheduleStringResultResponse

type UpdateScheduleRequest = api.UpdateScheduleRequest

type ListSchedulesRequest = api.ListSchedulesRequest
type ListSchedulesResponse = api.ListSchedulesResponse
type ScheduleListResult = api.ScheduleListResult
type ScheduleItem = api.ScheduleItem

type DeleteScheduleRequest = api.DeleteScheduleRequest

type CreateAutoRenewRuleRequest = api.CreateAutoRenewRuleRequest
type CreateAutoRenewRuleResponse = api.CreateAutoRenewRuleResponse
type OrderIdResult = api.OrderIdResult

type GetAutoRenewRuleDetailRequest = api.GetAutoRenewRuleDetailRequest
type GetAutoRenewRuleDetailResponse = api.GetAutoRenewRuleDetailResponse
type AutoRenewRuleDetail = api.AutoRenewRuleDetail

type ListAutoRenewRulesRequest = api.ListAutoRenewRulesRequest
type ListAutoRenewRulesResponse = api.ListAutoRenewRulesResponse
type AutoRenewRule = api.AutoRenewRule

type UpdateAutoRenewRuleRequest = api.UpdateAutoRenewRuleRequest
type RenewStringResultResponse = api.RenewStringResultResponse

type DeleteAutoRenewRuleRequest = api.DeleteAutoRenewRuleRequest
type DeleteAutoRenewRuleResponse = api.DeleteAutoRenewRuleResponse
type ExpireTimeResult = api.ExpireTimeResult

type RenewClusterRequest = api.RenewClusterRequest
type RenewClusterResponse = api.RenewClusterResponse

type ListRenewalsRequest = api.ListRenewalsRequest
type ListRenewalsResponse = api.ListRenewalsResponse
type RenewalPage = api.RenewalPage
type RenewalCluster = api.RenewalCluster

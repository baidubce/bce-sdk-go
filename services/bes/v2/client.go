package v2

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bes/v2/api"
)

const DEFAULT_ENDPOINT = "bes." + bce.DEFAULT_REGION + ".baidubce.com"

// Client is the BES v2 service client.
type Client struct {
	*bce.BceClient
	region string
}

// NewClient creates a BES v2 service client with AK/SK credentials.
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

// GetClusterDetail gets the detail of a BES cluster.
func (c *Client) GetClusterDetail(request *GetClusterDetailRequest) (*GetClusterDetailResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetClusterDetail(c.BceClient, c.region, request)
}

// DeleteCluster deletes a BES cluster.
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (*DeleteClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteCluster(c.BceClient, c.region, request)
}

// StartCluster starts a BES cluster.
func (c *Client) StartCluster(request *StartClusterRequest) (*StartClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.StartCluster(c.BceClient, c.region, request)
}

// StopCluster stops a BES cluster.
func (c *Client) StopCluster(request *StopClusterRequest) (*StopClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.StopCluster(c.BceClient, c.region, request)
}

// RestartCluster restarts a BES cluster.
func (c *Client) RestartCluster(request *RestartClusterRequest) (*RestartClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.RestartCluster(c.BceClient, c.region, request)
}

// ResizeCluster resizes a BES cluster.
func (c *Client) ResizeCluster(request *ResizeClusterRequest) (*ResizeClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ResizeCluster(c.BceClient, c.region, request)
}

// AddClusterModule adds a new node type to a BES cluster.
func (c *Client) AddClusterModule(request *AddClusterModuleRequest) (*AddClusterModuleResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.AddClusterModule(c.BceClient, c.region, request)
}

// ResetClusterPassword resets the admin password of a BES cluster.
func (c *Client) ResetClusterPassword(request *ResetClusterPasswordRequest) (*ResetClusterPasswordResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ResetClusterPassword(c.BceClient, c.region, request)
}

// ToggleClusterHTTPS enables or disables HTTPS for a BES cluster.
func (c *Client) ToggleClusterHTTPS(request *ToggleClusterHTTPSRequest) (*ToggleClusterHTTPSResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ToggleClusterHTTPS(c.BceClient, c.region, request)
}

// BindClusterEIP binds an EIP to a cluster module.
func (c *Client) BindClusterEIP(request *BindClusterEIPRequest) (*BindClusterEIPResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.BindClusterEIP(c.BceClient, c.region, request)
}

// UnbindClusterEIP unbinds an EIP from a cluster module.
func (c *Client) UnbindClusterEIP(request *UnbindClusterEIPRequest) (*UnbindClusterEIPResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UnbindClusterEIP(c.BceClient, c.region, request)
}

// ToggleClusterMonitor enables or disables Grafana monitor for a BES cluster.
func (c *Client) ToggleClusterMonitor(request *ToggleClusterMonitorRequest) (*ToggleClusterMonitorResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ToggleClusterMonitor(c.BceClient, c.region, request)
}

// GetClusterTasks gets the operation history of a BES cluster.
func (c *Client) GetClusterTasks(request *GetClusterTasksRequest) (*GetClusterTasksResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetClusterTasks(c.BceClient, c.region, request)
}

// GetClusterDataSizeTendency gets the data size tendency of a BES cluster.
func (c *Client) GetClusterDataSizeTendency(request *GetClusterDataSizeTendencyRequest) (*GetClusterDataSizeTendencyResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetClusterDataSizeTendency(c.BceClient, c.region, request)
}

// ListAvailableCoupons lists the user's available coupons.
func (c *Client) ListAvailableCoupons() (*ListAvailableCouponsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListAvailableCoupons(c.BceClient, c.region)
}

// AssessClusterSource runs an intelligent capacity assessment for a cluster plan.
func (c *Client) AssessClusterSource(request *AssessClusterSourceRequest) (*AssessClusterSourceResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.AssessClusterSource(c.BceClient, c.region, request)
}

/** ========================================= Tag API ======================================================================= */

// ListTags lists all tags used across the user's clusters.
func (c *Client) ListTags() (*TagListResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListTags(c.BceClient, c.region)
}

// UpdateClusterTags updates the tags of a single cluster.
func (c *Client) UpdateClusterTags(request *UpdateClusterTagsRequest) (*UpdateClusterTagsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateClusterTags(c.BceClient, c.region, request)
}

// BatchInsertTags inserts the same set of tags into multiple clusters at once.
func (c *Client) BatchInsertTags(request *BatchInsertTagsRequest) (*BatchInsertTagsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.BatchInsertTags(c.BceClient, c.region, request)
}

/** ========================================= Config API ==================================================================== */

// GetClusterConfig views a cluster's extra ES/Kibana configuration.
func (c *Client) GetClusterConfig(request *GetClusterConfigRequest) (*GetClusterConfigResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetClusterConfig(c.BceClient, c.region, request)
}

// UpdateClusterConfig performs a full update of a cluster's extra ES/Kibana configuration.
func (c *Client) UpdateClusterConfig(request *UpdateClusterConfigRequest) (*UpdateClusterConfigResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateClusterConfig(c.BceClient, c.region, request)
}

// ListSynonymDicts lists a cluster's uploaded synonym dict files.
func (c *Client) ListSynonymDicts(request *ListSynonymDictsRequest) (*ListSynonymDictsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListSynonymDicts(c.BceClient, c.region, request)
}

// DeleteSynonymDict deletes a cluster's synonym dict file.
func (c *Client) DeleteSynonymDict(request *DeleteSynonymDictRequest) (*DeleteSynonymDictResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteSynonymDict(c.BceClient, c.region, request)
}

// UploadSynonymDict uploads a synonym dict file to a cluster.
func (c *Client) UploadSynonymDict(request *UploadSynonymDictRequest) (*UploadSynonymDictResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UploadSynonymDict(c.BceClient, c.region, request)
}

/** ========================================= Log API ======================================================================= */

// UpdateLogSettings toggles a cluster's slow-log/audit-log switches.
func (c *Client) UpdateLogSettings(request *UpdateLogSettingsRequest) (*UpdateLogSettingsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateLogSettings(c.BceClient, c.region, request)
}

// SearchLog searches a cluster's logs by time range and type.
func (c *Client) SearchLog(request *SearchLogRequest) (*SearchLogResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.SearchLog(c.BceClient, c.region, request)
}

// CreateLogExportTask creates a log export task for one or more log types.
func (c *Client) CreateLogExportTask(request *CreateLogExportTaskRequest) (*CreateLogExportTaskResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateLogExportTask(c.BceClient, c.region, request)
}

// GetLogExportRecord gets the progress and configuration of a log export task.
func (c *Client) GetLogExportRecord(request *GetLogExportRecordRequest) (*GetLogExportRecordResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetLogExportRecord(c.BceClient, c.region, request)
}

/** ========================================= Instance API ================================================================== */

// StartInstance starts a single cluster instance.
func (c *Client) StartInstance(request *InstanceOperationRequest) (*InstanceOperationResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.StartInstance(c.BceClient, c.region, request)
}

// StopInstance stops a single cluster instance.
func (c *Client) StopInstance(request *InstanceOperationRequest) (*InstanceOperationResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.StopInstance(c.BceClient, c.region, request)
}

// BatchStartInstances starts multiple cluster instances at once.
func (c *Client) BatchStartInstances(request *BatchInstanceOperationRequest) (*BatchInstanceOperationResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.BatchStartInstances(c.BceClient, c.region, request)
}

// BatchStopInstances stops multiple cluster instances at once.
func (c *Client) BatchStopInstances(request *BatchInstanceOperationRequest) (*BatchInstanceOperationResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.BatchStopInstances(c.BceClient, c.region, request)
}

// DeleteInstances deletes (scales in) cluster instances.
func (c *Client) DeleteInstances(request *DeleteInstancesRequest) (*DeleteInstancesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteInstances(c.BceClient, c.region, request)
}

// ListScaleInInstances lists scale-in candidate instances for a cluster module.
func (c *Client) ListScaleInInstances(request *ListScaleInInstancesRequest) (*ListScaleInInstancesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListScaleInInstances(c.BceClient, c.region, request)
}

// ConfirmDataMigration confirms a pending data migration for the given instances.
func (c *Client) ConfirmDataMigration(request *ConfirmDataMigrationRequest) (*ConfirmDataMigrationResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ConfirmDataMigration(c.BceClient, c.region, request)
}

// RollbackDataMigration rolls back a previously confirmed data migration.
func (c *Client) RollbackDataMigration(request *RollbackDataMigrationRequest) (*RollbackDataMigrationResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.RollbackDataMigration(c.BceClient, c.region, request)
}

// ListDataMigrationInstances lists data migration candidate instances for a cluster module.
func (c *Client) ListDataMigrationInstances(request *ListDataMigrationInstancesRequest) (*ListDataMigrationInstancesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListDataMigrationInstances(c.BceClient, c.region, request)
}

// SuggestDataMigrationInstances gets the system-suggested instances for a data migration plan.
func (c *Client) SuggestDataMigrationInstances(request *SuggestDataMigrationInstancesRequest) (*SuggestDataMigrationInstancesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.SuggestDataMigrationInstances(c.BceClient, c.region, request)
}

/** ========================================= Plugin API ==================================================================== */

// InstallDefaultPlugin installs a system default plugin on a cluster module.
func (c *Client) InstallDefaultPlugin(request *DefaultPluginRequest) (*DefaultPluginResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.InstallDefaultPlugin(c.BceClient, c.region, request)
}

// UninstallDefaultPlugin uninstalls a system default plugin from a cluster module.
func (c *Client) UninstallDefaultPlugin(request *DefaultPluginRequest) (*DefaultPluginResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UninstallDefaultPlugin(c.BceClient, c.region, request)
}

// InstallCustomPlugin installs a previously uploaded custom plugin on a cluster module.
func (c *Client) InstallCustomPlugin(request *CustomPluginRequest) (*CustomPluginResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.InstallCustomPlugin(c.BceClient, c.region, request)
}

// UninstallCustomPlugin uninstalls a custom plugin from a cluster module.
func (c *Client) UninstallCustomPlugin(request *CustomPluginRequest) (*CustomPluginResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UninstallCustomPlugin(c.BceClient, c.region, request)
}

// DeleteCustomPlugin deletes an uploaded custom plugin package.
func (c *Client) DeleteCustomPlugin(request *CustomPluginRequest) (*DeleteCustomPluginResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteCustomPlugin(c.BceClient, c.region, request)
}

// UploadCustomPlugin uploads a custom plugin package (zip) to a cluster module.
func (c *Client) UploadCustomPlugin(request *UploadCustomPluginRequest) (*UploadCustomPluginResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UploadCustomPlugin(c.BceClient, c.region, request)
}

// GetPluginInfo gets a cluster's default and custom plugin lists.
func (c *Client) GetPluginInfo(request *GetPluginInfoRequest) (*GetPluginInfoResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetPluginInfo(c.BceClient, c.region, request)
}

// GetNLPDict views a cluster's NLP dict configuration.
func (c *Client) GetNLPDict(request *GetNLPDictRequest) (*GetNLPDictResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetNLPDict(c.BceClient, c.region, request)
}

// UpdateNLPDict uploads an NLP dict file to a cluster.
func (c *Client) UpdateNLPDict(request *UpdateNLPDictRequest) (*UpdateNLPDictResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateNLPDict(c.BceClient, c.region, request)
}

/** ========================================= Inspect API ==================================================================== */

// AuthorizeInspect authorizes intelligent inspection for a cluster.
func (c *Client) AuthorizeInspect(request *AuthorizeInspectRequest) (*AuthorizeInspectResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.AuthorizeInspect(c.BceClient, c.region, request)
}

// SwitchAutoInspect turns auto inspection on or off for a cluster.
func (c *Client) SwitchAutoInspect(request *SwitchAutoInspectRequest) (*InspectStringSuccessCommonResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.SwitchAutoInspect(c.BceClient, c.region, request)
}

// CheckAutoInspect checks whether auto inspection is enabled for a cluster.
func (c *Client) CheckAutoInspect(request *CheckAutoInspectRequest) (*CheckAutoInspectResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CheckAutoInspect(c.BceClient, c.region, request)
}

// CreateManualInspectTask submits a manual inspection task for a cluster.
func (c *Client) CreateManualInspectTask(request *CreateManualInspectTaskRequest) (*InspectStringSuccessCommonResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateManualInspectTask(c.BceClient, c.region, request)
}

// CheckInspectBusy checks whether a new inspection task can be submitted for a cluster.
func (c *Client) CheckInspectBusy(request *CheckInspectBusyRequest) (*CheckInspectBusyResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CheckInspectBusy(c.BceClient, c.region, request)
}

// GetManualInspectCount gets today's completed manual inspection count for a cluster.
func (c *Client) GetManualInspectCount(request *GetManualInspectCountRequest) (*GetManualInspectCountResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetManualInspectCount(c.BceClient, c.region, request)
}

// GetManualInspectConfig views a cluster's manual inspection configuration.
func (c *Client) GetManualInspectConfig(request *GetManualInspectConfigRequest) (*GetManualInspectConfigResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetManualInspectConfig(c.BceClient, c.region, request)
}

// UpdateManualInspectConfig updates a cluster's manual inspection configuration.
func (c *Client) UpdateManualInspectConfig(request *UpdateManualInspectConfigRequest) (*InspectStringSuccessCommonResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateManualInspectConfig(c.BceClient, c.region, request)
}

// ListInspectItems lists all selectable inspection items.
func (c *Client) ListInspectItems() (*ListInspectItemsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListInspectItems(c.BceClient, c.region)
}

// GetInspectTask gets a single inspection task's execution status and result.
func (c *Client) GetInspectTask(request *GetInspectTaskRequest) (*GetInspectTaskResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetInspectTask(c.BceClient, c.region, request)
}

// ListInspectTasks lists completed inspection tasks in the last 7 days.
func (c *Client) ListInspectTasks(request *ListInspectTasksRequest) (*ListInspectTasksResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListInspectTasks(c.BceClient, c.region, request)
}

// GetLatestInspectOverview gets the latest inspection overview for a cluster.
func (c *Client) GetLatestInspectOverview(request *GetLatestInspectOverviewRequest) (*GetLatestInspectOverviewResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetLatestInspectOverview(c.BceClient, c.region, request)
}

// GetWeeklyInspectOverview gets the last 7 days' inspection overview for a cluster.
func (c *Client) GetWeeklyInspectOverview(request *GetWeeklyInspectOverviewRequest) (*GetWeeklyInspectOverviewResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetWeeklyInspectOverview(c.BceClient, c.region, request)
}

/** ========================================= Renew API ===================================================================== */

// CreateAutoRenewRule creates an auto-renew rule for one or more clusters.
func (c *Client) CreateAutoRenewRule(request *CreateAutoRenewRuleRequest) (*CreateAutoRenewRuleResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateAutoRenewRule(c.BceClient, c.region, request)
}

// GetAutoRenewRuleDetail gets a cluster's auto-renew rule detail.
func (c *Client) GetAutoRenewRuleDetail(request *GetAutoRenewRuleDetailRequest) (*GetAutoRenewRuleDetailResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.GetAutoRenewRuleDetail(c.BceClient, c.region, request)
}

// ListAutoRenewRules lists auto-renew rules.
func (c *Client) ListAutoRenewRules(request *ListAutoRenewRulesRequest) (*ListAutoRenewRulesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListAutoRenewRules(c.BceClient, c.region, request)
}

// UpdateAutoRenewRule updates a cluster's auto-renew rule.
func (c *Client) UpdateAutoRenewRule(request *UpdateAutoRenewRuleRequest) (*RenewStringResultResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateAutoRenewRule(c.BceClient, c.region, request)
}

// DeleteAutoRenewRule deletes a cluster's auto-renew rule.
func (c *Client) DeleteAutoRenewRule(request *DeleteAutoRenewRuleRequest) (*DeleteAutoRenewRuleResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteAutoRenewRule(c.BceClient, c.region, request)
}

// RenewCluster renews a cluster for a given number of months.
func (c *Client) RenewCluster(request *RenewClusterRequest) (*RenewClusterResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.RenewCluster(c.BceClient, c.region, request)
}

// ListRenewals lists clusters approaching renewal.
func (c *Client) ListRenewals(request *ListRenewalsRequest) (*ListRenewalsResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListRenewals(c.BceClient, c.region, request)
}

/** ========================================= Schedule API ================================================================== */

// CreateSchedule creates a scheduled task on a cluster.
func (c *Client) CreateSchedule(request *CreateScheduleRequest) (*ScheduleStringResultResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.CreateSchedule(c.BceClient, c.region, request)
}

// UpdateSchedule updates a scheduled task on a cluster.
func (c *Client) UpdateSchedule(request *UpdateScheduleRequest) (*ScheduleStringResultResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.UpdateSchedule(c.BceClient, c.region, request)
}

// ListSchedules lists a cluster's scheduled tasks.
func (c *Client) ListSchedules(request *ListSchedulesRequest) (*ListSchedulesResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.ListSchedules(c.BceClient, c.region, request)
}

// DeleteSchedule deletes a scheduled task from a cluster.
func (c *Client) DeleteSchedule(request *DeleteScheduleRequest) (*ScheduleStringResultResponse, error) {
	if c == nil {
		return nil, api.ErrNilClient
	}
	return api.DeleteSchedule(c.BceClient, c.region, request)
}

/** ========================================= General Functions ============================================================= */

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

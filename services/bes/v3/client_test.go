package v3

import (
	"encoding/json"
	"errors"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bes/v3/api"
)

const (
	testAK = "ak-test-0000000000000000"
	testSK = "sk-test-0000000000000000"
)

func newTestClient(t *testing.T, handler nethttp.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)
	client, err := NewClient(testAK, testSK, server.URL)
	if err != nil {
		server.Close()
		t.Fatalf("new client failed: %v", err)
	}
	return client, server
}

func writeJSONResponse(t *testing.T, w nethttp.ResponseWriter, value interface{}) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	if err := json.NewEncoder(w).Encode(value); err != nil {
		t.Fatalf("write response failed: %v", err)
	}
}

// clientMethodCalls invokes every *Client method exactly once. It is used twice: against a real
// client to cover the `return api.Xxx(...)` delegation line of each method, and against a nil
// *Client to cover each method's `if c == nil` branch. Requests are intentionally empty — the
// delegation line executes regardless of whether the api layer then rejects the request, and the
// api layer's own validation is covered by the tests in services/bes/v3/api.
func clientMethodCalls(c *Client) []func() error {
	return []func() error{
		func() error { _, err := c.CreateCluster(&CreateClusterRequest{}); return err },
		func() error { _, err := c.ListClusters(&ListClustersRequest{}); return err },
		func() error { _, err := c.GetCluster(&GetClusterRequest{}); return err },
		func() error { _, err := c.ListClusterNodes(&ListClusterNodesRequest{}); return err },
		func() error { _, err := c.QueryAvailableSpecs(&QueryAvailableSpecsRequest{}); return err },
		func() error { _, err := c.QueryAvailableKernels(&QueryAvailableKernelsRequest{}); return err },
		func() error { _, err := c.DownloadClusterCert(&DownloadClusterCertRequest{}); return err },
		func() error { _, err := c.UpdateClusterName(&UpdateClusterNameRequest{}); return err },
		func() error { _, err := c.UpdateClusterMaintenance(&UpdateClusterMaintenanceRequest{}); return err },
		func() error {
			_, err := c.UpdateClusterDeletionProtection(&UpdateClusterDeletionProtectionRequest{})
			return err
		},
		func() error { _, err := c.RecoverCluster(&RecoverClusterRequest{}); return err },
		func() error { _, err := c.DeleteCluster(&DeleteClusterRequest{}); return err },
		func() error { _, err := c.ResizeCluster(&ResizeClusterRequest{}); return err },
		func() error {
			_, err := c.ListSuggestedMigrationNodes(&ListSuggestedMigrationNodesRequest{})
			return err
		},
		func() error { _, err := c.ListMigratableNodes(&ListMigratableNodesRequest{}); return err },
		func() error { _, err := c.QueryAvailableUpgrades(&QueryAvailableUpgradesRequest{}); return err },
		func() error { _, err := c.UpdateClusterPublicAccess(&UpdateClusterPublicAccessRequest{}); return err },
		func() error { _, err := c.UpdateClusterCerebro(&UpdateClusterCerebroRequest{}); return err },
		func() error { _, err := c.StartCluster(&StartClusterRequest{}); return err },
		func() error { _, err := c.StartClusterNodes(&ClusterNodesRequest{}); return err },
		func() error { _, err := c.MigrateClusterNodeData(&MigrateClusterNodeDataRequest{}); return err },
		func() error {
			_, err := c.UpdateClusterAccessWhitelist(&UpdateClusterAccessWhitelistRequest{})
			return err
		},
		func() error { _, err := c.UpgradeCluster(&UpgradeClusterRequest{}); return err },
		func() error { _, err := c.StopCluster(&StopClusterRequest{}); return err },
		func() error { _, err := c.StopClusterNodes(&ClusterNodesRequest{}); return err },
		func() error { _, err := c.UpdateClusterProtocol(&UpdateClusterProtocolRequest{}); return err },
		func() error { _, err := c.RestartCluster(&RestartClusterRequest{}); return err },
		func() error { _, err := c.CreateIndex(&CreateIndexRequest{}); return err },
		func() error { _, err := c.ListIndices(&ListIndicesRequest{}); return err },
		func() error { _, err := c.GetIndex(&GetIndexRequest{}); return err },
		func() error { _, err := c.ExistIndex(&ExistIndexRequest{}); return err },
		func() error { _, err := c.GetIndexStats(&GetIndexStatsRequest{}); return err },
		func() error { _, err := c.ListIndexFieldTypes(&ListIndexFieldTypesRequest{}); return err },
		func() error { _, err := c.OpenIndices(&IndexNamesRequest{}); return err },
		func() error { _, err := c.CloseIndices(&IndexNamesRequest{}); return err },
		func() error { _, err := c.DeleteIndices(&IndexNamesRequest{}); return err },
		func() error { _, err := c.RefreshIndices(&IndexNamesRequest{}); return err },
		func() error { _, err := c.FlushIndices(&IndexNamesRequest{}); return err },
		func() error { _, err := c.ForceMergeIndices(&ForceMergeIndicesRequest{}); return err },
		func() error { _, err := c.ClearIndicesCache(&IndexNamesRequest{}); return err },
		func() error { _, err := c.UpdateIndexSettings(&UpdateIndexSettingsRequest{}); return err },
		func() error { _, err := c.UpdateIndexMappings(&UpdateIndexMappingsRequest{}); return err },
		func() error { _, err := c.UpdateIndexAliases(&UpdateIndexAliasesRequest{}); return err },
		func() error { _, err := c.ListIndexTemplates(&ListIndexTemplatesRequest{}); return err },
		func() error { _, err := c.GetIndexTemplate(&GetIndexTemplateRequest{}); return err },
		func() error { _, err := c.ExistIndexTemplate(&ExistIndexTemplateRequest{}); return err },
		func() error { _, err := c.CreateIndexTemplate(&CreateIndexTemplateRequest{}); return err },
		func() error { _, err := c.UpdateIndexTemplate(&UpdateIndexTemplateRequest{}); return err },
		func() error { _, err := c.DeleteIndexTemplate(&DeleteIndexTemplateRequest{}); return err },
		func() error { _, err := c.ListIsmPolicies(&ListIsmPoliciesRequest{}); return err },
		func() error { _, err := c.GetIsmPolicy(&GetIsmPolicyRequest{}); return err },
		func() error { _, err := c.ExistIsmPolicy(&ExistIsmPolicyRequest{}); return err },
		func() error { _, err := c.CreateIsmPolicy(&IsmPolicyRequest{}); return err },
		func() error { _, err := c.UpdateIsmPolicy(&IsmPolicyRequest{}); return err },
		func() error { _, err := c.DeleteIsmPolicy(&DeleteIsmPolicyRequest{}); return err },
		func() error { _, err := c.ListMonitorIndices(&ListMonitorIndicesRequest{}); return err },
		func() error { _, err := c.ConvertToPrepay(&ConvertToPrepayRequest{}); return err },
		func() error { _, err := c.ConvertToPostpay(&ConvertToPostpayRequest{}); return err },
		func() error { _, err := c.CancelConvertToPostpay(&CancelConvertToPostpayRequest{}); return err },
		func() error { _, err := c.RenewCluster(&RenewClusterRequest{}); return err },
		func() error { _, err := c.QueryConfigPrice(&QueryConfigPriceRequest{}); return err },
		func() error { _, err := c.QueryClusterMarginPrice(&QueryClusterMarginPriceRequest{}); return err },
		func() error { _, err := c.ListActions(&ListActionsRequest{}); return err },
		func() error { _, err := c.ListActionTypes(&ListActionTypesRequest{}); return err },
		func() error { _, err := c.ListOperations(&ListOperationsRequest{}); return err },
		func() error { _, err := c.GetOperation(&GetOperationRequest{}); return err },
		func() error {
			_, err := c.ListOperationAnalysisDetails(&ListOperationAnalysisDetailsRequest{})
			return err
		},
		func() error { _, err := c.GetAuditStatus(&GetAuditStatusRequest{}); return err },
		func() error { _, err := c.EnableAudit(&EnableAuditRequest{}); return err },
		func() error { _, err := c.ListAuditEventTypes(&ListAuditEventTypesRequest{}); return err },
		func() error { _, err := c.SearchAuditEvents(&SearchAuditEventsRequest{}); return err },
		func() error { _, err := c.GetBctAuthSwitch(&Request{}); return err },
		func() error { _, err := c.UpdateBctAuthSwitch(&UpdateBctAuthSwitchRequest{}); return err },
		func() error { _, err := c.SearchBctEvents(&SearchBctEventsRequest{}); return err },
		func() error { _, err := c.GetLogSwitch(&GetLogSwitchRequest{}); return err },
		func() error { _, err := c.UpdateLogSwitch(&UpdateLogSwitchRequest{}); return err },
		func() error { _, err := c.GetLogUsage(&GetLogUsageRequest{}); return err },
		func() error { _, err := c.GetLogCollectorStatus(&GetLogCollectorStatusRequest{}); return err },
		func() error { _, err := c.SearchLogs(&SearchLogsRequest{}); return err },
		func() error { _, err := c.ListSystemPlugins(&ListSystemPluginsRequest{}); return err },
		func() error { _, err := c.UpdateSystemPlugin(&UpdateSystemPluginRequest{}); return err },
		func() error { _, err := c.ListCustomPlugins(&ListCustomPluginsRequest{}); return err },
		func() error { _, err := c.UpdateCustomPlugin(&UpdateCustomPluginRequest{}); return err },
		func() error {
			_, err := c.DeleteCustomPluginVersion(&DeleteCustomPluginVersionRequest{})
			return err
		},
		func() error { _, err := c.UploadPluginFile(&UploadPluginFileRequest{}); return err },
		func() error { _, err := c.ListWeeklyOverview(&ListWeeklyOverviewRequest{}); return err },
		func() error { _, err := c.GetManualDiagnosisCount(&GetManualDiagnosisCountRequest{}); return err },
		func() error { _, err := c.GetManualDiagnosisConfig(&GetManualDiagnosisConfigRequest{}); return err },
		func() error {
			_, err := c.UpdateManualDiagnosisConfig(&UpdateManualDiagnosisConfigRequest{})
			return err
		},
		func() error { _, err := c.CreateManualDiagnosis(&CreateManualDiagnosisRequest{}); return err },
		func() error { _, err := c.ListDiagnosisReports(&ListDiagnosisReportsRequest{}); return err },
		func() error { _, err := c.GetDiagnosisReport(&GetDiagnosisReportRequest{}); return err },
		func() error { _, err := c.GetDiagnosisBusyStatus(&GetDiagnosisBusyStatusRequest{}); return err },
		func() error {
			_, err := c.GetDiagnosisAuthorizationStatus(&GetDiagnosisAuthorizationStatusRequest{})
			return err
		},
		func() error { _, err := c.AuthorizeDiagnosis(&AuthorizeDiagnosisRequest{}); return err },
		func() error { _, err := c.ListDiagnosisItems(&ListDiagnosisItemsRequest{}); return err },
		func() error { _, err := c.GetAutoDiagnosisStatus(&GetAutoDiagnosisStatusRequest{}); return err },
		func() error { _, err := c.UpdateAutoDiagnosis(&UpdateAutoDiagnosisRequest{}); return err },
		func() error { _, err := c.GetLatestOverview(&GetLatestOverviewRequest{}); return err },
		func() error { _, err := c.ListIndicesByPattern(&ListIndicesByPatternRequest{}); return err },
		func() error { _, err := c.ListRestores(&ListRestoresRequest{}); return err },
		func() error { _, err := c.ListRestoreClusters(&ListRestoreClustersRequest{}); return err },
		func() error { _, err := c.GetSnapshotRule(&GetSnapshotRuleRequest{}); return err },
		func() error { _, err := c.ListSnapshots(&ListSnapshotsRequest{}); return err },
		func() error { _, err := c.GetSnapshotConfig(&GetSnapshotConfigRequest{}); return err },
		func() error { _, err := c.GetSnapshot(&GetSnapshotRequest{}); return err },
		func() error { _, err := c.ListAutoSnapshotConfigs(&ListAutoSnapshotConfigsRequest{}); return err },
		func() error { _, err := c.CreateRestore(&CreateRestoreRequest{}); return err },
		func() error { _, err := c.CreateManualSnapshot(&CreateManualSnapshotRequest{}); return err },
		func() error { _, err := c.CreateAutoSnapshotConfig(&CreateAutoSnapshotConfigRequest{}); return err },
		func() error { _, err := c.UpdateAutoSnapshotConfig(&UpdateAutoSnapshotConfigRequest{}); return err },
		func() error { _, err := c.DeleteSnapshot(&DeleteSnapshotRequest{}); return err },
		func() error { _, err := c.DeleteAutoSnapshotConfig(&DeleteAutoSnapshotConfigRequest{}); return err },
		func() error {
			_, err := c.BatchEnableAutoSnapshotConfigs(&BatchEnableAutoSnapshotConfigsRequest{})
			return err
		},
		func() error { _, err := c.ResetAdminPassword(&ResetAdminPasswordRequest{}); return err },
		func() error { _, err := c.ListSchedules(&ListSchedulesRequest{}); return err },
		func() error { _, err := c.CreateSchedule(&CreateScheduleRequest{}); return err },
		func() error { _, err := c.UpdateSchedule(&UpdateScheduleRequest{}); return err },
		func() error { _, err := c.DeleteSchedule(&DeleteScheduleRequest{}); return err },
	}
}

// TestClientDelegatesToAPI covers the `return api.Xxx(...)` line of every Client method. The stub
// server answers everything with an empty JSON object; whether the api layer then accepts the
// empty request is irrelevant here, the delegation line runs either way.
func TestClientDelegatesToAPI(t *testing.T) {
	client, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		writeJSONResponse(t, w, map[string]interface{}{})
	}))
	defer server.Close()

	calls := clientMethodCalls(client)
	if want := 119; len(calls) != want {
		t.Fatalf("clientMethodCalls covers %d methods, want %d — add the newly declared ones", len(calls), want)
	}
	for _, call := range calls {
		_ = call()
	}
}

// TestNilClientReturnsError covers the `if c == nil` branch of every Client method.
func TestNilClientReturnsError(t *testing.T) {
	var client *Client
	for i, call := range clientMethodCalls(client) {
		if err := call(); !errors.Is(err, api.ErrNilClient) {
			t.Fatalf("call %d: got %v, want ErrNilClient", i, err)
		}
	}
}

func TestNewClientConstructors(t *testing.T) {
	if _, err := NewClient("", testSK, ""); err == nil {
		t.Fatal("NewClient expected error for empty access key")
	}
	if _, err := NewClientWithSTS(testAK, testSK, "", ""); err == nil {
		t.Fatal("NewClientWithSTS expected error for empty session token")
	}

	// Empty endpoint falls back to the default endpoint and the default region.
	client, err := NewClient(testAK, testSK, "")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	if client.Config.Endpoint != DEFAULT_ENDPOINT {
		t.Fatalf("Endpoint = %q, want %q", client.Config.Endpoint, DEFAULT_ENDPOINT)
	}
	if client.region != "bj" {
		t.Fatalf("region = %q, want bj", client.region)
	}
	if client.Config.Credentials.SessionToken != "" {
		t.Fatalf("SessionToken = %q, want empty", client.Config.Credentials.SessionToken)
	}

	stsClient, err := NewClientWithSTS(testAK, testSK, "session-token", "https://bes.sh.baidubce.com")
	if err != nil {
		t.Fatalf("NewClientWithSTS returned error: %v", err)
	}
	if stsClient.Config.Credentials.SessionToken != "session-token" {
		t.Fatalf("SessionToken = %q, want session-token", stsClient.Config.Credentials.SessionToken)
	}
	if stsClient.region != "sh" {
		t.Fatalf("region = %q, want sh", stsClient.region)
	}
}

func TestInferRegionFromEndpoint(t *testing.T) {
	cases := map[string]string{
		"":                            "",
		"https://bes.bj.baidubce.com": "bj",
		"bes.sh.baidubce.com":         "sh",
		"bes.gz.baidubce.com:443":     "gz",
		"http://127.0.0.1:8080":       "",
		"https://example.com":         "",
		"bes.baidubce.com":            "",
	}
	for endpoint, want := range cases {
		if got := inferRegionFromEndpoint(endpoint); got != want {
			t.Fatalf("inferRegionFromEndpoint(%q) = %q, want %q", endpoint, got, want)
		}
	}
}

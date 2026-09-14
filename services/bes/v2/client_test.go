package v2

import (
	"encoding/json"
	"errors"
	nethttp "net/http"
	"net/http/httptest"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bes/v2/api"
)

const (
	testAK = "ak-test-0000000000000000"
	testSK = "sk-test-0000000000000000"
)

func newTestClient(t *testing.T, handler nethttp.HandlerFunc) (*Client, *httptest.Server) {
	t.Helper()
	server := httptest.NewServer(handler)
	client, err := NewClient("ak", "sk", server.URL)
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
// api layer's own validation is covered by the tests in services/bes/v2/api.
func clientMethodCalls(c *Client) []func() error {
	return []func() error{
		func() error { _, err := c.CreateCluster(&CreateClusterRequest{}); return err },
		func() error { _, err := c.ListClusters(&ListClustersRequest{}); return err },
		func() error { _, err := c.GetClusterDetail(&GetClusterDetailRequest{}); return err },
		func() error { _, err := c.DeleteCluster(&DeleteClusterRequest{}); return err },
		func() error { _, err := c.StartCluster(&StartClusterRequest{}); return err },
		func() error { _, err := c.StopCluster(&StopClusterRequest{}); return err },
		func() error { _, err := c.RestartCluster(&RestartClusterRequest{}); return err },
		func() error { _, err := c.ResizeCluster(&ResizeClusterRequest{}); return err },
		func() error { _, err := c.AddClusterModule(&AddClusterModuleRequest{}); return err },
		func() error { _, err := c.ResetClusterPassword(&ResetClusterPasswordRequest{}); return err },
		func() error { _, err := c.ToggleClusterHTTPS(&ToggleClusterHTTPSRequest{}); return err },
		func() error { _, err := c.BindClusterEIP(&BindClusterEIPRequest{}); return err },
		func() error { _, err := c.UnbindClusterEIP(&UnbindClusterEIPRequest{}); return err },
		func() error { _, err := c.ToggleClusterMonitor(&ToggleClusterMonitorRequest{}); return err },
		func() error { _, err := c.GetClusterTasks(&GetClusterTasksRequest{}); return err },
		func() error {
			_, err := c.GetClusterDataSizeTendency(&GetClusterDataSizeTendencyRequest{})
			return err
		},
		func() error { _, err := c.ListAvailableCoupons(); return err },
		func() error { _, err := c.AssessClusterSource(&AssessClusterSourceRequest{}); return err },
		func() error { _, err := c.ListTags(); return err },
		func() error { _, err := c.UpdateClusterTags(&UpdateClusterTagsRequest{}); return err },
		func() error { _, err := c.BatchInsertTags(&BatchInsertTagsRequest{}); return err },
		func() error { _, err := c.GetClusterConfig(&GetClusterConfigRequest{}); return err },
		func() error { _, err := c.UpdateClusterConfig(&UpdateClusterConfigRequest{}); return err },
		func() error { _, err := c.ListSynonymDicts(&ListSynonymDictsRequest{}); return err },
		func() error { _, err := c.DeleteSynonymDict(&DeleteSynonymDictRequest{}); return err },
		func() error { _, err := c.UploadSynonymDict(&UploadSynonymDictRequest{}); return err },
		func() error { _, err := c.UpdateLogSettings(&UpdateLogSettingsRequest{}); return err },
		func() error { _, err := c.SearchLog(&SearchLogRequest{}); return err },
		func() error { _, err := c.CreateLogExportTask(&CreateLogExportTaskRequest{}); return err },
		func() error { _, err := c.GetLogExportRecord(&GetLogExportRecordRequest{}); return err },
		func() error { _, err := c.StartInstance(&InstanceOperationRequest{}); return err },
		func() error { _, err := c.StopInstance(&InstanceOperationRequest{}); return err },
		func() error { _, err := c.BatchStartInstances(&BatchInstanceOperationRequest{}); return err },
		func() error { _, err := c.BatchStopInstances(&BatchInstanceOperationRequest{}); return err },
		func() error { _, err := c.DeleteInstances(&DeleteInstancesRequest{}); return err },
		func() error { _, err := c.ListScaleInInstances(&ListScaleInInstancesRequest{}); return err },
		func() error { _, err := c.ConfirmDataMigration(&ConfirmDataMigrationRequest{}); return err },
		func() error { _, err := c.RollbackDataMigration(&RollbackDataMigrationRequest{}); return err },
		func() error {
			_, err := c.ListDataMigrationInstances(&ListDataMigrationInstancesRequest{})
			return err
		},
		func() error {
			_, err := c.SuggestDataMigrationInstances(&SuggestDataMigrationInstancesRequest{})
			return err
		},
		func() error { _, err := c.InstallDefaultPlugin(&DefaultPluginRequest{}); return err },
		func() error { _, err := c.UninstallDefaultPlugin(&DefaultPluginRequest{}); return err },
		func() error { _, err := c.InstallCustomPlugin(&CustomPluginRequest{}); return err },
		func() error { _, err := c.UninstallCustomPlugin(&CustomPluginRequest{}); return err },
		func() error { _, err := c.DeleteCustomPlugin(&CustomPluginRequest{}); return err },
		func() error { _, err := c.UploadCustomPlugin(&UploadCustomPluginRequest{}); return err },
		func() error { _, err := c.GetPluginInfo(&GetPluginInfoRequest{}); return err },
		func() error { _, err := c.GetNLPDict(&GetNLPDictRequest{}); return err },
		func() error { _, err := c.UpdateNLPDict(&UpdateNLPDictRequest{}); return err },
		func() error { _, err := c.AuthorizeInspect(&AuthorizeInspectRequest{}); return err },
		func() error { _, err := c.SwitchAutoInspect(&SwitchAutoInspectRequest{}); return err },
		func() error { _, err := c.CheckAutoInspect(&CheckAutoInspectRequest{}); return err },
		func() error {
			_, err := c.CreateManualInspectTask(&CreateManualInspectTaskRequest{})
			return err
		},
		func() error { _, err := c.CheckInspectBusy(&CheckInspectBusyRequest{}); return err },
		func() error { _, err := c.GetManualInspectCount(&GetManualInspectCountRequest{}); return err },
		func() error {
			_, err := c.GetManualInspectConfig(&GetManualInspectConfigRequest{})
			return err
		},
		func() error {
			_, err := c.UpdateManualInspectConfig(&UpdateManualInspectConfigRequest{})
			return err
		},
		func() error { _, err := c.ListInspectItems(); return err },
		func() error { _, err := c.GetInspectTask(&GetInspectTaskRequest{}); return err },
		func() error { _, err := c.ListInspectTasks(&ListInspectTasksRequest{}); return err },
		func() error {
			_, err := c.GetLatestInspectOverview(&GetLatestInspectOverviewRequest{})
			return err
		},
		func() error {
			_, err := c.GetWeeklyInspectOverview(&GetWeeklyInspectOverviewRequest{})
			return err
		},
		func() error { _, err := c.CreateAutoRenewRule(&CreateAutoRenewRuleRequest{}); return err },
		func() error { _, err := c.GetAutoRenewRuleDetail(&GetAutoRenewRuleDetailRequest{}); return err },
		func() error { _, err := c.ListAutoRenewRules(&ListAutoRenewRulesRequest{}); return err },
		func() error { _, err := c.UpdateAutoRenewRule(&UpdateAutoRenewRuleRequest{}); return err },
		func() error { _, err := c.DeleteAutoRenewRule(&DeleteAutoRenewRuleRequest{}); return err },
		func() error { _, err := c.RenewCluster(&RenewClusterRequest{}); return err },
		func() error { _, err := c.ListRenewals(&ListRenewalsRequest{}); return err },
		func() error { _, err := c.CreateSchedule(&CreateScheduleRequest{}); return err },
		func() error { _, err := c.UpdateSchedule(&UpdateScheduleRequest{}); return err },
		func() error { _, err := c.ListSchedules(&ListSchedulesRequest{}); return err },
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
	if want := 73; len(calls) != want {
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

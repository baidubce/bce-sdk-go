package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	inspectTestClusterId = "218185657699405824"
	inspectTestTaskId    = "10"
)

// inspectWriteBadRequest writes a 400 service error. A 4xx status is used on purpose: the test
// client mirrors the production DEFAULT_RETRY_POLICY, so a 5xx would trigger backoff retries.
func inspectWriteBadRequest(t *testing.T, w nethttp.ResponseWriter) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(nethttp.StatusBadRequest)
	if _, err := w.Write([]byte(`{"code":"BadRequest","message":"invalid inspect request"}`)); err != nil {
		t.Fatalf("write error response failed: %v", err)
	}
}

// inspectUnreachableHandler fails the test when a request reaches the server, used by cases that
// must fail during local validation.
func inspectUnreachableHandler(t *testing.T) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Errorf("unexpected request reached the server: %s", r.URL.Path)
		inspectWriteBadRequest(t, w)
	}
}

// assertInspectRequest checks the parts of the request every inspect API builds the same way.
func assertInspectRequest(t *testing.T, r *nethttp.Request, path string) {
	t.Helper()
	if r.URL.Path != path {
		t.Fatalf("unexpected path: %s", r.URL.Path)
	}
	if r.Method != nethttp.MethodPost {
		t.Fatalf("unexpected method: %s", r.Method)
	}
	if got := r.Header.Get(HEADER_REGION); got != testRegion {
		t.Fatalf("unexpected region header: %s", got)
	}
}

// inspectFullUpdateConfigRequest / inspectFullGetTaskRequest / inspectFullListTasksRequest return
// fully populated requests, so the required-field cases below only have to clear one field each.
func inspectFullUpdateConfigRequest() *UpdateManualInspectConfigRequest {
	return &UpdateManualInspectConfigRequest{
		ClusterId: inspectTestClusterId,
		Indices:   "*",
		Items:     []string{"ClusterHealth", "ClusterPayload"},
	}
}

func inspectFullGetTaskRequest() *GetInspectTaskRequest {
	return &GetInspectTaskRequest{ClusterId: inspectTestClusterId, TaskId: inspectTestTaskId}
}

func inspectFullListTasksRequest() *ListInspectTasksRequest {
	return &ListInspectTasksRequest{ClusterId: inspectTestClusterId, PageNo: 1, PageSize: 10}
}

// TestAuthorizeInspect is the only inspect API whose response success is a plain JSON bool instead
// of a SuccessFlag, so the bool is asserted directly here.
func TestAuthorizeInspect(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/authorize")
		if got := readJSONBody(t, r)["clusterId"]; got != inspectTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{},
		})
	}))
	defer server.Close()

	result, err := AuthorizeInspect(cli, testRegion, &AuthorizeInspectRequest{
		ClusterId: inspectTestClusterId,
	})
	if err != nil {
		t.Fatalf("authorize inspect failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result == nil {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestSwitchAutoInspect(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/switch_auto")
		body := readJSONBody(t, r)
		if body["clusterId"] != inspectTestClusterId || body["switchOn"] != true {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result":  map[string]interface{}{},
		})
	}))
	defer server.Close()

	result, err := SwitchAutoInspect(cli, testRegion, &SwitchAutoInspectRequest{
		ClusterId: inspectTestClusterId,
		SwitchOn:  true,
	})
	if err != nil {
		t.Fatalf("switch auto inspect failed: %v", err)
	}
	if !result.Success.Bool() || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestCheckAutoInspect(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/check_auto")
		if got := readJSONBody(t, r)["clusterId"]; got != inspectTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result":  map[string]interface{}{"switchOn": true},
		})
	}))
	defer server.Close()

	result, err := CheckAutoInspect(cli, testRegion, &CheckAutoInspectRequest{
		ClusterId: inspectTestClusterId,
	})
	if err != nil {
		t.Fatalf("check auto inspect failed: %v", err)
	}
	if result.Result == nil || !result.Result.SwitchOn {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if result.Success != SuccessFlagTrue || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestCreateManualInspectTask(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/create_manual")
		if got := readJSONBody(t, r)["clusterId"]; got != inspectTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result":  map[string]interface{}{},
		})
	}))
	defer server.Close()

	result, err := CreateManualInspectTask(cli, testRegion, &CreateManualInspectTaskRequest{
		ClusterId: inspectTestClusterId,
	})
	if err != nil {
		t.Fatalf("create manual inspect task failed: %v", err)
	}
	if result.Success != SuccessFlagTrue || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestCheckInspectBusy(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/check_busy")
		if got := readJSONBody(t, r)["clusterId"]; got != inspectTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result":  map[string]interface{}{"busy": true},
		})
	}))
	defer server.Close()

	result, err := CheckInspectBusy(cli, testRegion, &CheckInspectBusyRequest{
		ClusterId: inspectTestClusterId,
	})
	if err != nil {
		t.Fatalf("check inspect busy failed: %v", err)
	}
	if result.Result == nil || !result.Result.Busy {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestGetManualInspectCount(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/manual_count")
		if got := readJSONBody(t, r)["clusterId"]; got != inspectTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result":  map[string]interface{}{"count": 3},
		})
	}))
	defer server.Close()

	result, err := GetManualInspectCount(cli, testRegion, &GetManualInspectCountRequest{
		ClusterId: inspectTestClusterId,
	})
	if err != nil {
		t.Fatalf("get manual inspect count failed: %v", err)
	}
	if result.Result == nil || result.Result.Count != 3 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestGetManualInspectConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/get_manual_conf")
		if got := readJSONBody(t, r)["clusterId"]; got != inspectTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result": map[string]interface{}{
				"indices": "*",
				"items":   []string{"ClusterHealth", "ClusterPayload"},
			},
		})
	}))
	defer server.Close()

	result, err := GetManualInspectConfig(cli, testRegion, &GetManualInspectConfigRequest{
		ClusterId: inspectTestClusterId,
	})
	if err != nil {
		t.Fatalf("get manual inspect config failed: %v", err)
	}
	if result.Result == nil || result.Result.Indices != "*" || len(result.Result.Items) != 2 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if result.Result.Items[0] != "ClusterHealth" {
		t.Fatalf("unexpected items: %+v", result.Result.Items)
	}
}

// TestUpdateManualInspectConfig also fills the optional Items list, so its omitempty body field is
// exercised instead of being dropped.
func TestUpdateManualInspectConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/update_manual_conf")
		body := readJSONBody(t, r)
		if body["clusterId"] != inspectTestClusterId || body["indices"] != "*" {
			t.Fatalf("unexpected body: %+v", body)
		}
		items, ok := body["items"].([]interface{})
		if !ok || len(items) != 2 || items[0] != "ClusterHealth" || items[1] != "ClusterPayload" {
			t.Fatalf("unexpected items: %v", body["items"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result":  map[string]interface{}{},
		})
	}))
	defer server.Close()

	result, err := UpdateManualInspectConfig(cli, testRegion, inspectFullUpdateConfigRequest())
	if err != nil {
		t.Fatalf("update manual inspect config failed: %v", err)
	}
	if result.Success != SuccessFlagTrue || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
}

// TestListInspectItems is the only inspect API that takes no request: it always posts an empty JSON
// object, which is asserted here.
func TestListInspectItems(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/list_items")
		if body := readJSONBody(t, r); len(body) != 0 {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result": map[string]interface{}{
				"items": []map[string]interface{}{
					{"id": "ClusterHealth", "type": "cluster"},
					{"id": "NodeShardNum", "type": "node"},
				},
			},
		})
	}))
	defer server.Close()

	result, err := ListInspectItems(cli, testRegion)
	if err != nil {
		t.Fatalf("list inspect items failed: %v", err)
	}
	if result.Result == nil || len(result.Result.Items) != 2 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if result.Result.Items[1].Id != "NodeShardNum" || result.Result.Items[1].Type != "node" {
		t.Fatalf("unexpected items: %+v", result.Result.Items)
	}
}

func TestGetInspectTask(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/get_task")
		body := readJSONBody(t, r)
		if body["clusterId"] != inspectTestClusterId || body["taskId"] != inspectTestTaskId {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result": map[string]interface{}{
				"state":      "doing",
				"type":       "manual",
				"indices":    "index1",
				"totalCount": 14,
				"execTime":   "2024-07-20 18:31:42",
				"groups": []map[string]interface{}{
					{
						"grade":     "lowRisk",
						"itemCount": 1,
						"results": []map[string]interface{}{
							{
								"id":               "ClusterHealth",
								"type":             "cluster",
								"conclusion":       "low risk",
								"hasLowRiskAdvice": true,
								"notExistIndex":    false,
							},
						},
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := GetInspectTask(cli, testRegion, inspectFullGetTaskRequest())
	if err != nil {
		t.Fatalf("get inspect task failed: %v", err)
	}
	detail := result.Result
	if detail == nil || detail.State != "doing" || detail.Type != "manual" {
		t.Fatalf("unexpected result: %+v", detail)
	}
	if detail.Indices != "index1" || detail.TotalCount != 14 || detail.ExecTime != "2024-07-20 18:31:42" {
		t.Fatalf("unexpected detail: %+v", detail)
	}
	if len(detail.Groups) != 1 || detail.Groups[0].Grade != "lowRisk" || detail.Groups[0].ItemCount != 1 {
		t.Fatalf("unexpected groups: %+v", detail.Groups)
	}
	if len(detail.Groups[0].Results) != 1 {
		t.Fatalf("unexpected group results: %+v", detail.Groups[0].Results)
	}
	item := detail.Groups[0].Results[0]
	if item.Id != "ClusterHealth" || item.Type != "cluster" || item.Conclusion != "low risk" {
		t.Fatalf("unexpected item result: %+v", item)
	}
	if !item.HasLowRiskAdvice || item.NotExistIndex {
		t.Fatalf("unexpected item flags: %+v", item)
	}
}

// TestListInspectTasks asserts the page-shaped response: this API answers with a top-level "page"
// field instead of "result".
func TestListInspectTasks(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/list_tasks")
		body := readJSONBody(t, r)
		if body["clusterId"] != inspectTestClusterId {
			t.Fatalf("unexpected clusterId: %v", body["clusterId"])
		}
		if body["pageNo"] != float64(1) || body["pageSize"] != float64(10) {
			t.Fatalf("unexpected page fields: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"page": map[string]interface{}{
				"orderBy":    "execTime",
				"order":      "desc",
				"pageNo":     1,
				"pageSize":   10,
				"totalCount": 20,
				"result": []map[string]interface{}{
					{
						"taskId":        "11",
						"type":          "auto",
						"execTime":      "2024-07-20 18:31:42",
						"lowRiskCount":  1,
						"highRiskCount": 2,
						"errorCount":    3,
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := ListInspectTasks(cli, testRegion, inspectFullListTasksRequest())
	if err != nil {
		t.Fatalf("list inspect tasks failed: %v", err)
	}
	if result.Page == nil || result.Page.TotalCount != 20 || len(result.Page.Result) != 1 {
		t.Fatalf("unexpected page: %+v", result.Page)
	}
	if result.Page.OrderBy != "execTime" || result.Page.Order != "desc" {
		t.Fatalf("unexpected page order: %+v", result.Page)
	}
	if result.Page.PageNo != 1 || result.Page.PageSize != 10 {
		t.Fatalf("unexpected page numbers: %+v", result.Page)
	}
	task := result.Page.Result[0]
	if task.TaskId != "11" || task.Type != "auto" || task.ExecTime != "2024-07-20 18:31:42" {
		t.Fatalf("unexpected task: %+v", task)
	}
	if task.LowRiskCount != 1 || task.HighRiskCount != 2 || task.ErrorCount != 3 {
		t.Fatalf("unexpected task counts: %+v", task)
	}
}

func TestGetLatestInspectOverview(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/overview/latest")
		if got := readJSONBody(t, r)["clusterId"]; got != inspectTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result": map[string]interface{}{
				"taskId":        "11",
				"taskStatus":    "success",
				"countSafe":     1,
				"countLowRisk":  2,
				"countHighRisk": 3,
				"countError":    4,
				"execTime":      "2024-07-20 18:31:42",
			},
		})
	}))
	defer server.Close()

	result, err := GetLatestInspectOverview(cli, testRegion, &GetLatestInspectOverviewRequest{
		ClusterId: inspectTestClusterId,
	})
	if err != nil {
		t.Fatalf("get latest inspect overview failed: %v", err)
	}
	overview := result.Result
	if overview == nil || overview.TaskId != "11" || overview.TaskStatus != "success" {
		t.Fatalf("unexpected result: %+v", overview)
	}
	if overview.CountSafe != 1 || overview.CountLowRisk != 2 {
		t.Fatalf("unexpected counts: %+v", overview)
	}
	if overview.CountHighRisk != 3 || overview.CountError != 4 {
		t.Fatalf("unexpected risk counts: %+v", overview)
	}
	if overview.ExecTime != "2024-07-20 18:31:42" {
		t.Fatalf("unexpected execTime: %s", overview.ExecTime)
	}
}

func TestGetWeeklyInspectOverview(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertInspectRequest(t, r, "/api/bes/cluster/inspect/overview/weekly")
		if got := readJSONBody(t, r)["clusterId"]; got != inspectTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": "true",
			"status":  200,
			"result": map[string]interface{}{
				"count": 1,
				"riskItems": []map[string]interface{}{
					{"id": "ClusterHealth", "highRiskCount": 1, "lowRiskCount": 3},
				},
			},
		})
	}))
	defer server.Close()

	result, err := GetWeeklyInspectOverview(cli, testRegion, &GetWeeklyInspectOverviewRequest{
		ClusterId: inspectTestClusterId,
	})
	if err != nil {
		t.Fatalf("get weekly inspect overview failed: %v", err)
	}
	if result.Result == nil || result.Result.Count != 1 || len(result.Result.RiskItems) != 1 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	risk := result.Result.RiskItems[0]
	if risk.Id != "ClusterHealth" || risk.HighRiskCount != 1 || risk.LowRiskCount != 3 {
		t.Fatalf("unexpected risk item: %+v", risk)
	}
}

func TestInspectAPIRejectsNilClient(t *testing.T) {
	cases := []struct {
		name string
		call func() error
	}{
		{"AuthorizeInspect", func() error {
			_, err := AuthorizeInspect(nil, testRegion, &AuthorizeInspectRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"SwitchAutoInspect", func() error {
			_, err := SwitchAutoInspect(nil, testRegion, &SwitchAutoInspectRequest{
				ClusterId: inspectTestClusterId,
				SwitchOn:  true,
			})
			return err
		}},
		{"CheckAutoInspect", func() error {
			_, err := CheckAutoInspect(nil, testRegion, &CheckAutoInspectRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"CreateManualInspectTask", func() error {
			_, err := CreateManualInspectTask(nil, testRegion, &CreateManualInspectTaskRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"CheckInspectBusy", func() error {
			_, err := CheckInspectBusy(nil, testRegion, &CheckInspectBusyRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"GetManualInspectCount", func() error {
			_, err := GetManualInspectCount(nil, testRegion, &GetManualInspectCountRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"GetManualInspectConfig", func() error {
			_, err := GetManualInspectConfig(nil, testRegion, &GetManualInspectConfigRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"UpdateManualInspectConfig", func() error {
			_, err := UpdateManualInspectConfig(nil, testRegion, inspectFullUpdateConfigRequest())
			return err
		}},
		{"ListInspectItems", func() error {
			_, err := ListInspectItems(nil, testRegion)
			return err
		}},
		{"GetInspectTask", func() error {
			_, err := GetInspectTask(nil, testRegion, inspectFullGetTaskRequest())
			return err
		}},
		{"ListInspectTasks", func() error {
			_, err := ListInspectTasks(nil, testRegion, inspectFullListTasksRequest())
			return err
		}},
		{"GetLatestInspectOverview", func() error {
			_, err := GetLatestInspectOverview(nil, testRegion, &GetLatestInspectOverviewRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"GetWeeklyInspectOverview", func() error {
			_, err := GetWeeklyInspectOverview(nil, testRegion, &GetWeeklyInspectOverviewRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(); !errors.Is(err, ErrNilClient) {
				t.Fatalf("expected ErrNilClient, got %v", err)
			}
		})
	}
}

// TestInspectAPIRejectsNilRequest covers the nil-request guard of every inspect API that takes a
// request. ListInspectItems is absent on purpose: it takes no request parameter at all.
func TestInspectAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, inspectUnreachableHandler(t))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"AuthorizeInspect", func(c bce.Client) error {
			_, err := AuthorizeInspect(c, testRegion, nil)
			return err
		}},
		{"SwitchAutoInspect", func(c bce.Client) error {
			_, err := SwitchAutoInspect(c, testRegion, nil)
			return err
		}},
		{"CheckAutoInspect", func(c bce.Client) error {
			_, err := CheckAutoInspect(c, testRegion, nil)
			return err
		}},
		{"CreateManualInspectTask", func(c bce.Client) error {
			_, err := CreateManualInspectTask(c, testRegion, nil)
			return err
		}},
		{"CheckInspectBusy", func(c bce.Client) error {
			_, err := CheckInspectBusy(c, testRegion, nil)
			return err
		}},
		{"GetManualInspectCount", func(c bce.Client) error {
			_, err := GetManualInspectCount(c, testRegion, nil)
			return err
		}},
		{"GetManualInspectConfig", func(c bce.Client) error {
			_, err := GetManualInspectConfig(c, testRegion, nil)
			return err
		}},
		{"UpdateManualInspectConfig", func(c bce.Client) error {
			_, err := UpdateManualInspectConfig(c, testRegion, nil)
			return err
		}},
		{"GetInspectTask", func(c bce.Client) error {
			_, err := GetInspectTask(c, testRegion, nil)
			return err
		}},
		{"ListInspectTasks", func(c bce.Client) error {
			_, err := ListInspectTasks(c, testRegion, nil)
			return err
		}},
		{"GetLatestInspectOverview", func(c bce.Client) error {
			_, err := GetLatestInspectOverview(c, testRegion, nil)
			return err
		}},
		{"GetWeeklyInspectOverview", func(c bce.Client) error {
			_, err := GetWeeklyInspectOverview(c, testRegion, nil)
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(cli); err == nil {
				t.Fatal("expected error for nil request")
			}
		})
	}
}

// TestInspectAPIRejectsMissingRequiredFields covers all 16 required-field validation branches in
// inspect.go. ListInspectItems is absent on purpose: it takes no request and therefore validates
// nothing beyond the nil client.
func TestInspectAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, inspectUnreachableHandler(t))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"AuthorizeInspect empty clusterId", func(c bce.Client) error {
			_, err := AuthorizeInspect(c, testRegion, &AuthorizeInspectRequest{})
			return err
		}},
		{"SwitchAutoInspect empty clusterId", func(c bce.Client) error {
			_, err := SwitchAutoInspect(c, testRegion, &SwitchAutoInspectRequest{SwitchOn: true})
			return err
		}},
		{"CheckAutoInspect empty clusterId", func(c bce.Client) error {
			_, err := CheckAutoInspect(c, testRegion, &CheckAutoInspectRequest{})
			return err
		}},
		{"CreateManualInspectTask empty clusterId", func(c bce.Client) error {
			_, err := CreateManualInspectTask(c, testRegion, &CreateManualInspectTaskRequest{})
			return err
		}},
		{"CheckInspectBusy empty clusterId", func(c bce.Client) error {
			_, err := CheckInspectBusy(c, testRegion, &CheckInspectBusyRequest{})
			return err
		}},
		{"GetManualInspectCount empty clusterId", func(c bce.Client) error {
			_, err := GetManualInspectCount(c, testRegion, &GetManualInspectCountRequest{})
			return err
		}},
		{"GetManualInspectConfig empty clusterId", func(c bce.Client) error {
			_, err := GetManualInspectConfig(c, testRegion, &GetManualInspectConfigRequest{})
			return err
		}},
		{"UpdateManualInspectConfig empty clusterId", func(c bce.Client) error {
			request := inspectFullUpdateConfigRequest()
			request.ClusterId = ""
			_, err := UpdateManualInspectConfig(c, testRegion, request)
			return err
		}},
		{"UpdateManualInspectConfig empty indices", func(c bce.Client) error {
			request := inspectFullUpdateConfigRequest()
			request.Indices = ""
			_, err := UpdateManualInspectConfig(c, testRegion, request)
			return err
		}},
		{"GetInspectTask empty clusterId", func(c bce.Client) error {
			request := inspectFullGetTaskRequest()
			request.ClusterId = ""
			_, err := GetInspectTask(c, testRegion, request)
			return err
		}},
		{"GetInspectTask empty taskId", func(c bce.Client) error {
			request := inspectFullGetTaskRequest()
			request.TaskId = ""
			_, err := GetInspectTask(c, testRegion, request)
			return err
		}},
		{"ListInspectTasks empty clusterId", func(c bce.Client) error {
			request := inspectFullListTasksRequest()
			request.ClusterId = ""
			_, err := ListInspectTasks(c, testRegion, request)
			return err
		}},
		{"ListInspectTasks zero pageNo", func(c bce.Client) error {
			request := inspectFullListTasksRequest()
			request.PageNo = 0
			_, err := ListInspectTasks(c, testRegion, request)
			return err
		}},
		{"ListInspectTasks zero pageSize", func(c bce.Client) error {
			request := inspectFullListTasksRequest()
			request.PageSize = 0
			_, err := ListInspectTasks(c, testRegion, request)
			return err
		}},
		{"GetLatestInspectOverview empty clusterId", func(c bce.Client) error {
			_, err := GetLatestInspectOverview(c, testRegion, &GetLatestInspectOverviewRequest{})
			return err
		}},
		{"GetWeeklyInspectOverview empty clusterId", func(c bce.Client) error {
			_, err := GetWeeklyInspectOverview(c, testRegion, &GetWeeklyInspectOverviewRequest{})
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(cli); err == nil {
				t.Fatal("expected a required-field validation error")
			}
		})
	}
}

func TestInspectAPIPropagatesServerError(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		inspectWriteBadRequest(t, w)
	}))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"AuthorizeInspect", func(c bce.Client) error {
			_, err := AuthorizeInspect(c, testRegion, &AuthorizeInspectRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"SwitchAutoInspect", func(c bce.Client) error {
			_, err := SwitchAutoInspect(c, testRegion, &SwitchAutoInspectRequest{
				ClusterId: inspectTestClusterId,
				SwitchOn:  true,
			})
			return err
		}},
		{"CheckAutoInspect", func(c bce.Client) error {
			_, err := CheckAutoInspect(c, testRegion, &CheckAutoInspectRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"CreateManualInspectTask", func(c bce.Client) error {
			_, err := CreateManualInspectTask(c, testRegion, &CreateManualInspectTaskRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"CheckInspectBusy", func(c bce.Client) error {
			_, err := CheckInspectBusy(c, testRegion, &CheckInspectBusyRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"GetManualInspectCount", func(c bce.Client) error {
			_, err := GetManualInspectCount(c, testRegion, &GetManualInspectCountRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"GetManualInspectConfig", func(c bce.Client) error {
			_, err := GetManualInspectConfig(c, testRegion, &GetManualInspectConfigRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"UpdateManualInspectConfig", func(c bce.Client) error {
			_, err := UpdateManualInspectConfig(c, testRegion, inspectFullUpdateConfigRequest())
			return err
		}},
		{"ListInspectItems", func(c bce.Client) error {
			_, err := ListInspectItems(c, testRegion)
			return err
		}},
		{"GetInspectTask", func(c bce.Client) error {
			_, err := GetInspectTask(c, testRegion, inspectFullGetTaskRequest())
			return err
		}},
		{"ListInspectTasks", func(c bce.Client) error {
			_, err := ListInspectTasks(c, testRegion, inspectFullListTasksRequest())
			return err
		}},
		{"GetLatestInspectOverview", func(c bce.Client) error {
			_, err := GetLatestInspectOverview(c, testRegion, &GetLatestInspectOverviewRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
		{"GetWeeklyInspectOverview", func(c bce.Client) error {
			_, err := GetWeeklyInspectOverview(c, testRegion, &GetWeeklyInspectOverviewRequest{
				ClusterId: inspectTestClusterId,
			})
			return err
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(cli); err == nil {
				t.Fatal("expected the server error to be propagated")
			}
		})
	}
}

// TestInspectSuccessFlagAcceptsBoolAndString is migrated from the v2 client tests. The
// intelligent-inspection docs declare success as a String for 12 of the 13 interfaces, but the
// server answers an unquoted JSON boolean on all of them. SuccessFlag (util.go) has to accept
// either shape, otherwise every successful inspection call fails with
// "cannot unmarshal bool into Go struct field ...success of type string".
func TestInspectSuccessFlagAcceptsBoolAndString(t *testing.T) {
	cases := []struct {
		name     string
		success  interface{}
		expected SuccessFlag
		wantBool bool
	}{
		{"JSONBoolTrue", true, SuccessFlagTrue, true},
		{"JSONBoolFalse", false, SuccessFlagFalse, false},
		{"QuotedTrue", "true", SuccessFlagTrue, true},
		{"QuotedFalse", "false", SuccessFlagFalse, false},
		{"QuotedOther", "unknown", SuccessFlag("unknown"), false},
		{"Null", nil, SuccessFlag(""), false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
				defer r.Body.Close()
				writeJSONResponse(t, w, map[string]interface{}{
					"success": c.success,
					"status":  200,
					"result": map[string]interface{}{
						"items": []map[string]interface{}{{"id": "ClusterHealth", "type": "cluster"}},
					},
				})
			}))
			defer server.Close()

			result, err := ListInspectItems(cli, testRegion)
			if err != nil {
				t.Fatalf("list inspect items failed: %v", err)
			}
			if result.Success != c.expected {
				t.Fatalf("expected success %q, got %q", c.expected, result.Success)
			}
			if result.Success.Bool() != c.wantBool {
				t.Fatalf("expected Bool() %v, got %v", c.wantBool, result.Success.Bool())
			}
			if result.Result == nil || len(result.Result.Items) != 1 {
				t.Fatalf("unexpected result: %+v", result.Result)
			}
		})
	}
}

// TestInspectSuccessFlagRejectsUnsupportedShape covers the error return of
// SuccessFlag.UnmarshalJSON: a success value that is neither a boolean nor a string cannot be
// mapped to the flag, so the whole response must fail to parse.
func TestInspectSuccessFlagRejectsUnsupportedShape(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		writeJSONResponse(t, w, map[string]interface{}{"success": 42, "status": 200})
	}))
	defer server.Close()

	if _, err := ListInspectItems(cli, testRegion); err == nil {
		t.Fatal("expected a decoding error for a numeric success flag")
	}
}

// TestInspectSuccessFlagUnmarshalJSONEmptyPayload exercises the empty-input branch of
// SuccessFlag.UnmarshalJSON directly: encoding/json never hands an empty slice to a custom
// unmarshaller, so the branch is only reachable through a direct call.
func TestInspectSuccessFlagUnmarshalJSONEmptyPayload(t *testing.T) {
	for _, payload := range []string{``, `   `, `null`} {
		flag := SuccessFlagTrue
		if err := flag.UnmarshalJSON([]byte(payload)); err != nil {
			t.Fatalf("%q: unmarshal failed: %v", payload, err)
		}
		if flag != SuccessFlag("") || flag.Bool() {
			t.Fatalf("%q: unexpected flag: %q", payload, flag)
		}
	}
}

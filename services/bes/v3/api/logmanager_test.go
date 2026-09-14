package api

import (
	"errors"
	nethttp "net/http"
	"testing"
)

// assertLogManagerRequest checks the method, path and region header of a log manager request.
func assertLogManagerRequest(t *testing.T, r *nethttp.Request, method, path, region string) {
	t.Helper()
	if r.Method != method {
		t.Fatalf("unexpected method: got %s want %s", r.Method, method)
	}
	if r.URL.Path != path {
		t.Fatalf("unexpected path: got %s want %s", r.URL.Path, path)
	}
	if got := r.Header.Get("X-Region"); got != region {
		t.Fatalf("unexpected X-Region: got %s want %s", got, region)
	}
}

func TestGetLogSwitch(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogManagerRequest(t, r, nethttp.MethodGet, "/v3/clusters/search-xxxxxxxx/logs/switch", "sh")
		writeJSONResponse(t, w, map[string]interface{}{
			"enabled": true,
			"types":   map[string]bool{"main": true, "audit": false},
		})
	}))
	defer server.Close()

	result, err := GetLogSwitch(cli, testRegion, &GetLogSwitchRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("get log switch failed: %v", err)
	}
	if !result.Enabled || !result.Types["main"] || result.Types["audit"] {
		t.Fatalf("unexpected log switch result: %+v", result)
	}
}

func TestUpdateLogSwitch(t *testing.T) {
	enabled := false
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogManagerRequest(t, r, nethttp.MethodPut, "/v3/clusters/search-xxxxxxxx/logs/switch", testRegion)
		body := readJSONBody(t, r)
		if got, ok := body["enabled"].(bool); !ok || got {
			t.Fatalf("expected enabled=false in request body: %+v", body)
		}
		if body["types"].(map[string]interface{})["audit"] != true {
			t.Fatalf("expected audit switch in request body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"enabled": false,
			"types":   map[string]bool{"audit": true},
		})
	}))
	defer server.Close()

	result, err := UpdateLogSwitch(cli, testRegion, &UpdateLogSwitchRequest{
		ClusterId: "search-xxxxxxxx",
		Enabled:   &enabled,
		Types:     map[string]bool{"audit": true},
	})
	if err != nil {
		t.Fatalf("update log switch failed: %v", err)
	}
	if result.Enabled || !result.Types["audit"] {
		t.Fatalf("unexpected log switch result: %+v", result)
	}
}

func TestGetLogUsage(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogManagerRequest(t, r, nethttp.MethodGet, "/v3/clusters/search-xxxxxxxx/logs/usage", testRegion)
		writeJSONResponse(t, w, map[string]interface{}{
			"totalSizeBytes": 32052578,
			"totalShards":    28,
			"indices": map[string]interface{}{
				"main": map[string]interface{}{"sizeBytes": 14006681, "shards": 12, "docCount": 38244},
			},
		})
	}))
	defer server.Close()

	result, err := GetLogUsage(cli, testRegion, &GetLogUsageRequest{ClusterId: "search-xxxxxxxx"})
	if err != nil {
		t.Fatalf("get log usage failed: %v", err)
	}
	mainUsage := result.Indices["main"]
	if result.TotalSizeBytes != 32052578 || result.TotalShards != 28 || mainUsage.SizeBytes != 14006681 ||
		mainUsage.Shards != 12 || mainUsage.DocCount != 38244 {
		t.Fatalf("unexpected log usage result: %+v", result)
	}
}

func TestGetLogCollectorStatus(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogManagerRequest(t, r, nethttp.MethodGet, "/v3/clusters/search-xxxxxxxx/logs/collector/status", testRegion)
		writeJSONResponse(t, w, map[string]interface{}{
			"deployed":        true,
			"operationStatus": "RUNNING",
			"operationType":   "DEPLOY_LOG_COLLECTOR",
			"operationId":     "operation-1",
		})
	}))
	defer server.Close()

	result, err := GetLogCollectorStatus(cli, testRegion, &GetLogCollectorStatusRequest{ClusterId: "search-xxxxxxxx"})
	if err != nil {
		t.Fatalf("get log collector status failed: %v", err)
	}
	if !result.Deployed || result.OperationStatus != "RUNNING" ||
		result.OperationType != "DEPLOY_LOG_COLLECTOR" || result.OperationId != "operation-1" {
		t.Fatalf("unexpected log collector status result: %+v", result)
	}
}

func TestSearchLogs(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogManagerRequest(t, r, nethttp.MethodPost, "/v3/clusters/search-xxxxxxxx/logs/search", testRegion)
		body := readJSONBody(t, r)
		if body["logType"] != "main" || body["keyword"] != "error" || body["level"] != "ERROR" ||
			body["nodeName"] != "node-1" || body["startTime"] == nil || body["endTime"] == nil ||
			body["page"] != float64(1) || body["pageSize"] != float64(10) {
			t.Fatalf("unexpected search request body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"total":    1,
			"page":     1,
			"pageSize": 10,
			"logs": []map[string]interface{}{{
				"timestamp": "2026-07-24T02:02:18.006Z",
				"message":   "error",
				"node.id":   "node-1",
			}},
		})
	}))
	defer server.Close()

	result, err := SearchLogs(cli, testRegion, &SearchLogsRequest{
		ClusterId: "search-xxxxxxxx",
		LogType:   "main",
		Keyword:   "error",
		Level:     "ERROR",
		NodeName:  "node-1",
		StartTime: "2026-07-24T01:02:21.046541Z",
		EndTime:   "2026-07-24T02:02:21.046541Z",
		Page:      1,
		PageSize:  10,
	})
	if err != nil {
		t.Fatalf("search logs failed: %v", err)
	}
	if result.Total != 1 || result.Page != 1 || result.PageSize != 10 ||
		result.Logs[0]["message"] != "error" || result.Logs[0]["node.id"] != "node-1" {
		t.Fatalf("unexpected search logs result: %+v", result)
	}
}

func TestLogManagerAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"GetLogSwitch":    func() error { _, err := GetLogSwitch(nil, testRegion, &GetLogSwitchRequest{}); return err },
		"UpdateLogSwitch": func() error { _, err := UpdateLogSwitch(nil, testRegion, &UpdateLogSwitchRequest{}); return err },
		"GetLogUsage":     func() error { _, err := GetLogUsage(nil, testRegion, &GetLogUsageRequest{}); return err },
		"GetLogCollectorStatus": func() error {
			_, err := GetLogCollectorStatus(nil, testRegion, &GetLogCollectorStatusRequest{})
			return err
		},
		"SearchLogs": func() error { _, err := SearchLogs(nil, testRegion, &SearchLogsRequest{}); return err },
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

func TestLogManagerAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"GetLogSwitch":          func() error { _, err := GetLogSwitch(cli, testRegion, nil); return err },
		"UpdateLogSwitch":       func() error { _, err := UpdateLogSwitch(cli, testRegion, nil); return err },
		"GetLogUsage":           func() error { _, err := GetLogUsage(cli, testRegion, nil); return err },
		"GetLogCollectorStatus": func() error { _, err := GetLogCollectorStatus(cli, testRegion, nil); return err },
		"SearchLogs":            func() error { _, err := SearchLogs(cli, testRegion, nil); return err },
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

func TestLogManagerAPIRejectsMissingClusterId(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"GetLogSwitch":    func() error { _, err := GetLogSwitch(cli, testRegion, &GetLogSwitchRequest{}); return err },
		"UpdateLogSwitch": func() error { _, err := UpdateLogSwitch(cli, testRegion, &UpdateLogSwitchRequest{}); return err },
		"GetLogUsage":     func() error { _, err := GetLogUsage(cli, testRegion, &GetLogUsageRequest{}); return err },
		"GetLogCollectorStatus": func() error {
			_, err := GetLogCollectorStatus(cli, testRegion, &GetLogCollectorStatusRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestSearchLogsRejectsInvalidRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := SearchLogsRequest{
		ClusterId: "search-xxxxxxxx",
		LogType:   "main",
		StartTime: "2026-07-24T01:02:21Z",
		EndTime:   "2026-07-24T02:02:21Z",
		Page:      1,
		PageSize:  10,
	}
	cases := map[string]func(*SearchLogsRequest){
		"clusterId":        func(r *SearchLogsRequest) { r.ClusterId = "" },
		"logType":          func(r *SearchLogsRequest) { r.LogType = "" },
		"startTime":        func(r *SearchLogsRequest) { r.StartTime = "" },
		"endTime":          func(r *SearchLogsRequest) { r.EndTime = "" },
		"pageNegative":     func(r *SearchLogsRequest) { r.Page = -1 },
		"pageTooLarge":     func(r *SearchLogsRequest) { r.Page = 51 },
		"pageSizeNegative": func(r *SearchLogsRequest) { r.PageSize = -1 },
		"pageSizeTooLarge": func(r *SearchLogsRequest) { r.PageSize = 201 },
	}
	for name, mutate := range cases {
		request := full
		mutate(&request)
		if _, err := SearchLogs(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestLogManagerAPIPropagatesServerError uses 400 rather than 500 on purpose: the helper mirrors
// the production retry policy, which retries 5xx with backoff and would make this test take seconds.
func TestLogManagerAPIPropagatesServerError(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		w.WriteHeader(nethttp.StatusBadRequest)
		writeJSONResponse(t, w, map[string]interface{}{
			"code":      "InvalidParameter",
			"message":   "invalid parameter",
			"requestId": "req-xxxx",
		})
	}))
	defer server.Close()

	cases := map[string]func() error{
		"GetLogSwitch": func() error {
			_, err := GetLogSwitch(cli, testRegion, &GetLogSwitchRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"UpdateLogSwitch": func() error {
			_, err := UpdateLogSwitch(cli, testRegion, &UpdateLogSwitchRequest{
				ClusterId: "search-xxxxxxxx",
				Enabled:   boolPtr(true),
			})
			return err
		},
		"GetLogUsage": func() error {
			_, err := GetLogUsage(cli, testRegion, &GetLogUsageRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"GetLogCollectorStatus": func() error {
			_, err := GetLogCollectorStatus(cli, testRegion, &GetLogCollectorStatusRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"SearchLogs": func() error {
			_, err := SearchLogs(cli, testRegion, &SearchLogsRequest{
				ClusterId: "search-xxxxxxxx",
				LogType:   "main",
				StartTime: "2026-07-24T01:02:21Z",
				EndTime:   "2026-07-24T02:02:21Z",
			})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for 400 response", name)
		}
	}
}

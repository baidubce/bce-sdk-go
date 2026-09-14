package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

// assertLogAPIRequest checks the method, path and region header of a log manager request.
// It is a log-manager-local copy of the assertion helper used by the v2 client tests.
func assertLogAPIRequest(t *testing.T, r *nethttp.Request, method, path string) {
	t.Helper()
	if r.Method != method {
		t.Errorf("unexpected method: got %s, want %s", r.Method, method)
	}
	if r.URL.Path != path {
		t.Errorf("unexpected path: got %s, want %s", r.URL.Path, path)
	}
	if got := r.Header.Get(HEADER_REGION); got != testRegion {
		t.Errorf("unexpected region header: got %s, want %s", got, testRegion)
	}
}

// logUnusedClient returns a client whose server fails the test when it is called. It is used by
// the validation tests to prove no HTTP request is issued when a request is rejected locally.
func logUnusedClient(t *testing.T) (bce.Client, func()) {
	t.Helper()
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Errorf("unexpected request to %s", r.URL.Path)
	}))
	return cli, server.Close
}

// logFailingClient returns a client whose server always answers with an HTTP 400 service error.
// 400 is used on purpose: newTestClient mirrors bce.DEFAULT_RETRY_POLICY, so a 5xx status would
// trigger the retry back-off and slow the package tests down by an order of magnitude.
func logFailingClient(t *testing.T) (bce.Client, func()) {
	t.Helper()
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(nethttp.StatusBadRequest)
		_, _ = w.Write([]byte(`{"code":"BadRequest","message":"invalid log request","requestId":"req-log-1"}`))
	}))
	return cli, server.Close
}

func TestUpdateLogSettings(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogAPIRequest(t, r, nethttp.MethodPost, "/api/bes/cluster/es_log_view_new/settings")
		body := readJSONBody(t, r)
		if body["clusterId"] != "811453865437302781" || body["auditLog"] != true ||
			body["indexSlowLog"] != false || body["searchSlowLog"] != false {
			t.Errorf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := UpdateLogSettings(cli, testRegion, &UpdateLogSettingsRequest{
		ClusterId:     "811453865437302781",
		IndexSlowLog:  false,
		SearchSlowLog: false,
		AuditLog:      true,
	})
	if err != nil {
		t.Fatalf("update log settings failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestSearchLog(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogAPIRequest(t, r, nethttp.MethodPost, "/api/bes/cluster/es_log_view_new/search")
		body := readJSONBody(t, r)
		// keyword is the only optional body field; the happy path fills it so the omitempty
		// branch of the marshalled payload is exercised too.
		if body["clusterId"] != "811453865437302784" || body["keyword"] != "timeout" ||
			body["startTime"] != "2023-11-02 14:24:21" || body["endTime"] != "2023-11-09 14:24:21" ||
			body["type"] != "mainLog" {
			t.Errorf("unexpected body: %+v", body)
		}
		if body["pageNo"] != float64(1) || body["pageSize"] != float64(10) {
			t.Errorf("unexpected paging in body: %+v", body)
		}
		// 实际响应是 success/status/page 三层信封，日志在 page.result 数组里
		// （文档的响应表只列了内层四个字段，见 SearchLogResponse 注释）。
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"page": map[string]interface{}{
				"pageNo":     1,
				"pageSize":   10,
				"totalCount": 2595,
				"orderBy":    "",
				"order":      "",
				"result": []map[string]interface{}{
					{
						"level":   "WARN",
						"host":    "0.0.0.0",
						"time":    "2023-11-09T17:11:43,013",
						"content": "xxxx",
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := SearchLog(cli, testRegion, &SearchLogRequest{
		ClusterId: "811453865437302784",
		PageNo:    1,
		PageSize:  10,
		Keyword:   "timeout",
		StartTime: "2023-11-02 14:24:21",
		EndTime:   "2023-11-09 14:24:21",
		Type:      "mainLog",
	})
	if err != nil {
		t.Fatalf("search log failed: %v", err)
	}
	if !result.Success || result.Page == nil || result.Page.TotalCount != 2595 {
		t.Fatalf("unexpected result: %+v", result)
	}
	if len(result.Page.Result) != 1 || result.Page.Result[0].Level != "WARN" ||
		result.Page.Result[0].Content != "xxxx" {
		t.Fatalf("unexpected log entries: %+v", result.Page.Result)
	}
}

func TestCreateLogExportTask(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogAPIRequest(t, r, nethttp.MethodPost, "/api/bes/cluster/es_log_view_new/export")
		body := readJSONBody(t, r)
		if body["clusterId"] != "811453865437302784" || body["startTime"] != "2023-11-02 14:24:21" ||
			body["endTime"] != "2023-11-03 14:24:21" || body["bosBucket"] != "my-bucket" ||
			body["bosPath"] != "my-path" {
			t.Errorf("unexpected body: %+v", body)
		}
		types, ok := body["types"].([]interface{})
		if !ok || len(types) != 1 || types[0] != "mainLog" {
			t.Errorf("unexpected types in body: %+v", body["types"])
		}
		// tasks 在 result 信封里面，不是平铺在顶层（见 CreateLogExportTaskResponse 注释）。
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"tasks": []map[string]interface{}{{"logType": "mainLog", "taskId": "12345"}},
			},
		})
	}))
	defer server.Close()

	result, err := CreateLogExportTask(cli, testRegion, &CreateLogExportTaskRequest{
		ClusterId: "811453865437302784",
		Types:     []string{"mainLog"},
		StartTime: "2023-11-02 14:24:21",
		EndTime:   "2023-11-03 14:24:21",
		BosBucket: "my-bucket",
		BosPath:   "my-path",
	})
	if err != nil {
		t.Fatalf("create log export task failed: %v", err)
	}
	if result.Result == nil || len(result.Result.Tasks) != 1 || result.Result.Tasks[0].TaskId != "12345" ||
		result.Result.Tasks[0].LogType != "mainLog" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

// CreateLogExportTask treats the request as optional: a nil request is replaced by an empty
// struct, so the call must succeed and send an empty JSON object instead of returning an error.
func TestCreateLogExportTaskAcceptsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogAPIRequest(t, r, nethttp.MethodPost, "/api/bes/cluster/es_log_view_new/export")
		if body := readJSONBody(t, r); len(body) != 0 {
			t.Errorf("unexpected body for nil request: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200})
	}))
	defer server.Close()

	result, err := CreateLogExportTask(cli, testRegion, nil)
	if err != nil {
		t.Fatalf("create log export task with nil request failed: %v", err)
	}
	if !result.Success || result.Result != nil {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetLogExportRecord(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogAPIRequest(t, r, nethttp.MethodPost, "/api/bes/cluster/es_log_view_new/export_record")
		body := readJSONBody(t, r)
		if body["clusterId"] != "811453865437302784" || body["taskId"] != "123456" {
			t.Errorf("unexpected body: %+v", body)
		}
		// progress/state/config 同样在 result 信封里面；
		// config.startTime/endTime 是毫秒时间戳数字，不是文档写的 String。
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"progress": 85,
				"state":    "进行中",
				"config": map[string]interface{}{
					"logType":   "gcLog",
					"bosBucket": "mybucket",
					"bosPath":   "bes_log_repository",
					"startTime": int64(1787709708000),
					"endTime":   int64(1787713308000),
					"region":    "bj",
				},
			},
		})
	}))
	defer server.Close()

	result, err := GetLogExportRecord(cli, testRegion, &GetLogExportRecordRequest{
		ClusterId: "811453865437302784",
		TaskId:    "123456",
	})
	if err != nil {
		t.Fatalf("get log export record failed: %v", err)
	}
	if result.Result == nil || result.Result.Progress != 85 || result.Result.State != "进行中" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if result.Result.Config == nil || result.Result.Config.LogType != "gcLog" ||
		result.Result.Config.StartTime != 1787709708000 || result.Result.Config.EndTime != 1787713308000 ||
		result.Result.Config.BosBucket != "mybucket" || result.Result.Config.BosPath != "bes_log_repository" ||
		result.Result.Config.Region != "bj" {
		t.Fatalf("unexpected config: %+v", result.Result.Config)
	}
}

// GetLogExportRecord also treats the request as optional and substitutes an empty struct.
func TestGetLogExportRecordAcceptsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertLogAPIRequest(t, r, nethttp.MethodPost, "/api/bes/cluster/es_log_view_new/export_record")
		if body := readJSONBody(t, r); len(body) != 0 {
			t.Errorf("unexpected body for nil request: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200})
	}))
	defer server.Close()

	result, err := GetLogExportRecord(cli, testRegion, nil)
	if err != nil {
		t.Fatalf("get log export record with nil request failed: %v", err)
	}
	if !result.Success || result.Result != nil {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestLogManagerAPIRejectsNilClient(t *testing.T) {
	tests := []struct {
		name string
		call func() error
	}{
		{"UpdateLogSettings", func() error {
			_, err := UpdateLogSettings(nil, testRegion, &UpdateLogSettingsRequest{ClusterId: "c"})
			return err
		}},
		{"SearchLog", func() error {
			_, err := SearchLog(nil, testRegion, &SearchLogRequest{ClusterId: "c"})
			return err
		}},
		{"CreateLogExportTask", func() error {
			_, err := CreateLogExportTask(nil, testRegion, &CreateLogExportTaskRequest{})
			return err
		}},
		{"GetLogExportRecord", func() error {
			_, err := GetLogExportRecord(nil, testRegion, &GetLogExportRecordRequest{})
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); !errors.Is(err, ErrNilClient) {
				t.Fatalf("expected ErrNilClient, got %v", err)
			}
		})
	}
}

// Only UpdateLogSettings and SearchLog require a request. CreateLogExportTask and
// GetLogExportRecord treat it as optional and are covered by the *AcceptsNilRequest tests above.
func TestLogManagerAPIRejectsNilRequest(t *testing.T) {
	cli, closeServer := logUnusedClient(t)
	defer closeServer()

	tests := []struct {
		name string
		call func() error
		want string
	}{
		{"UpdateLogSettings", func() error {
			_, err := UpdateLogSettings(cli, testRegion, nil)
			return err
		}, "update log settings request should not be nil"},
		{"SearchLog", func() error {
			_, err := SearchLog(cli, testRegion, nil)
			return err
		}, "search log request should not be nil"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.call()
			if err == nil || err.Error() != tt.want {
				t.Fatalf("expected error %q, got %v", tt.want, err)
			}
		})
	}
}

func TestLogManagerAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, closeServer := logUnusedClient(t)
	defer closeServer()

	// fullSearchRequest returns a request with every required field set; each case blanks exactly
	// one field so the matching guard in SearchLog is the one that fires.
	fullSearchRequest := func() *SearchLogRequest {
		return &SearchLogRequest{
			ClusterId: "811453865437302784",
			PageNo:    1,
			PageSize:  10,
			Keyword:   "timeout",
			StartTime: "2023-11-02 14:24:21",
			EndTime:   "2023-11-09 14:24:21",
			Type:      "mainLog",
		}
	}
	searchCase := func(name string, mutate func(*SearchLogRequest), want string) struct {
		name string
		call func() error
		want string
	} {
		return struct {
			name string
			call func() error
			want string
		}{name, func() error {
			request := fullSearchRequest()
			mutate(request)
			_, err := SearchLog(cli, testRegion, request)
			return err
		}, want}
	}

	tests := []struct {
		name string
		call func() error
		want string
	}{
		{"UpdateLogSettings missing clusterId", func() error {
			_, err := UpdateLogSettings(cli, testRegion, &UpdateLogSettingsRequest{AuditLog: true})
			return err
		}, "update log settings request clusterId should not be empty"},
		searchCase("SearchLog missing clusterId", func(r *SearchLogRequest) { r.ClusterId = "" },
			"search log request clusterId should not be empty"),
		searchCase("SearchLog missing pageNo", func(r *SearchLogRequest) { r.PageNo = 0 },
			"search log request pageNo should not be empty"),
		searchCase("SearchLog missing pageSize", func(r *SearchLogRequest) { r.PageSize = 0 },
			"search log request pageSize should not be empty"),
		searchCase("SearchLog missing startTime", func(r *SearchLogRequest) { r.StartTime = "" },
			"search log request startTime should not be empty"),
		searchCase("SearchLog missing endTime", func(r *SearchLogRequest) { r.EndTime = "" },
			"search log request endTime should not be empty"),
		searchCase("SearchLog missing type", func(r *SearchLogRequest) { r.Type = "" },
			"search log request type should not be empty"),
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.call()
			if err == nil || err.Error() != tt.want {
				t.Fatalf("expected error %q, got %v", tt.want, err)
			}
		})
	}
}

func TestLogManagerAPIPropagatesServerError(t *testing.T) {
	cli, closeServer := logFailingClient(t)
	defer closeServer()

	tests := []struct {
		name string
		call func() error
	}{
		{"UpdateLogSettings", func() error {
			_, err := UpdateLogSettings(cli, testRegion, &UpdateLogSettingsRequest{ClusterId: "c"})
			return err
		}},
		{"SearchLog", func() error {
			_, err := SearchLog(cli, testRegion, &SearchLogRequest{
				ClusterId: "c",
				PageNo:    1,
				PageSize:  10,
				StartTime: "2023-11-02 14:24:21",
				EndTime:   "2023-11-09 14:24:21",
				Type:      "mainLog",
			})
			return err
		}},
		{"CreateLogExportTask", func() error {
			_, err := CreateLogExportTask(cli, testRegion, &CreateLogExportTaskRequest{ClusterId: "c"})
			return err
		}},
		{"GetLogExportRecord", func() error {
			_, err := GetLogExportRecord(cli, testRegion, &GetLogExportRecordRequest{ClusterId: "c"})
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.call(); err == nil {
				t.Fatal("expected server error, got nil")
			}
		})
	}
}

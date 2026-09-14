package api

import (
	"errors"
	nethttp "net/http"
	"testing"
)

func TestListSchedules(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/schedules" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		if got := r.URL.Query().Get("pageNo"); got != "1" {
			t.Fatalf("unexpected pageNo: %s", got)
		}
		if got := r.URL.Query().Get("pageSize"); got != "20" {
			t.Fatalf("unexpected pageSize: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"schedules": []map[string]interface{}{
				{
					"scheduleId":    "9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e",
					"scheduleName":  "my-delete-task",
					"scheduleType":  "DELETE",
					"cronExpr":      "0 30 1 * * ?",
					"enabled":       true,
					"createTime":    "2026-07-28 02:00:00",
					"executedCount": 12,
					"failedCount":   0,
					"params":        map[string]interface{}{"indexPattern": "logs-*"},
				},
			},
			"pageNo":     1,
			"pageSize":   20,
			"totalCount": 1,
		})
	}))
	defer server.Close()

	result, err := ListSchedules(cli, testRegion, &ListSchedulesRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
		PageNo:    1,
		PageSize:  20,
	})
	if err != nil {
		t.Fatalf("list schedules failed: %v", err)
	}
	if result.TotalCount != 1 || len(result.Schedules) != 1 || result.Schedules[0].ScheduleId != "9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e" {
		t.Fatalf("unexpected result: %+v", result)
	}
	if result.Schedules[0].ScheduleType != SCHEDULE_TYPE_DELETE {
		t.Fatalf("unexpected schedule type: %+v", result.Schedules[0])
	}
}

// TestListSchedulesWithoutPagination covers the request path where the optional pageNo and
// pageSize parameters are omitted from the query string entirely.
func TestListSchedulesWithoutPagination(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if _, ok := r.URL.Query()["pageNo"]; ok {
			t.Fatalf("unexpected pageNo in query: %s", r.URL.RawQuery)
		}
		if _, ok := r.URL.Query()["pageSize"]; ok {
			t.Fatalf("unexpected pageSize in query: %s", r.URL.RawQuery)
		}
		if got := r.Header.Get("X-Region"); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"totalCount": 0})
	}))
	defer server.Close()

	result, err := ListSchedules(cli, testRegion, &ListSchedulesRequest{ClusterId: "search-xxxxxxxx"})
	if err != nil {
		t.Fatalf("list schedules failed: %v", err)
	}
	if result.TotalCount != 0 || len(result.Schedules) != 0 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestCreateSchedule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/schedules" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if body["scheduleName"] != "my-delete-task" || body["scheduleType"] != "DELETE" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["cronExpr"] != "0 30 1 * * ?" {
			t.Fatalf("unexpected body: %+v", body)
		}
		params, ok := body["params"].(map[string]interface{})
		if !ok || params["indexPattern"] != "logs-*" {
			t.Fatalf("unexpected params: %+v", body["params"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := CreateSchedule(cli, testRegion, &CreateScheduleRequest{
		Request:      Request{Region: "sh"},
		ClusterId:    "search-xxxxxxxx",
		ScheduleName: "my-delete-task",
		ScheduleType: SCHEDULE_TYPE_DELETE,
		CronExpr:     "0 30 1 * * ?",
		Params:       map[string]interface{}{"indexPattern": "logs-*"},
	})
	if err != nil {
		t.Fatalf("create schedule failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestUpdateSchedule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/schedules/9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if body["cronExpr"] != "0 30 1 * * ?" {
			t.Fatalf("unexpected body: %+v", body)
		}
		params, ok := body["params"].(map[string]interface{})
		if !ok || params["indexPattern"] != "logs-*" {
			t.Fatalf("unexpected params: %+v", body["params"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := UpdateSchedule(cli, testRegion, &UpdateScheduleRequest{
		Request:    Request{Region: "sh"},
		ClusterId:  "search-xxxxxxxx",
		ScheduleId: "9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e",
		CronExpr:   "0 30 1 * * ?",
		Params:     map[string]interface{}{"indexPattern": "logs-*"},
	})
	if err != nil {
		t.Fatalf("update schedule failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestDeleteSchedule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/schedules/9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := DeleteSchedule(cli, testRegion, &DeleteScheduleRequest{
		Request:    Request{Region: "sh"},
		ClusterId:  "search-xxxxxxxx",
		ScheduleId: "9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e",
	})
	if err != nil {
		t.Fatalf("delete schedule failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestScheduleAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"ListSchedules":  func() error { _, err := ListSchedules(nil, testRegion, &ListSchedulesRequest{}); return err },
		"CreateSchedule": func() error { _, err := CreateSchedule(nil, testRegion, &CreateScheduleRequest{}); return err },
		"UpdateSchedule": func() error { _, err := UpdateSchedule(nil, testRegion, &UpdateScheduleRequest{}); return err },
		"DeleteSchedule": func() error { _, err := DeleteSchedule(nil, testRegion, &DeleteScheduleRequest{}); return err },
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

func TestScheduleAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ListSchedules":  func() error { _, err := ListSchedules(cli, testRegion, nil); return err },
		"CreateSchedule": func() error { _, err := CreateSchedule(cli, testRegion, nil); return err },
		"UpdateSchedule": func() error { _, err := UpdateSchedule(cli, testRegion, nil); return err },
		"DeleteSchedule": func() error { _, err := DeleteSchedule(cli, testRegion, nil); return err },
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

func TestListSchedulesRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	if _, err := ListSchedules(cli, testRegion, &ListSchedulesRequest{}); err == nil {
		t.Fatal("clusterId: expected validation error")
	}
}

func TestCreateScheduleRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := CreateScheduleRequest{
		ClusterId:    "search-xxxxxxxx",
		ScheduleName: "my-delete-task",
		ScheduleType: SCHEDULE_TYPE_DELETE,
		CronExpr:     "0 30 1 * * ?",
		Params:       map[string]interface{}{"indexPattern": "logs-*"},
	}
	cases := map[string]func(*CreateScheduleRequest){
		"clusterId":    func(r *CreateScheduleRequest) { r.ClusterId = "" },
		"scheduleName": func(r *CreateScheduleRequest) { r.ScheduleName = "" },
		"scheduleType": func(r *CreateScheduleRequest) { r.ScheduleType = "" },
		"cronExpr":     func(r *CreateScheduleRequest) { r.CronExpr = "" },
		"params":       func(r *CreateScheduleRequest) { r.Params = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := CreateSchedule(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateScheduleRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := UpdateScheduleRequest{
		ClusterId:  "search-xxxxxxxx",
		ScheduleId: "9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e",
		CronExpr:   "0 30 1 * * ?",
		Params:     map[string]interface{}{"indexPattern": "logs-*"},
	}
	cases := map[string]func(*UpdateScheduleRequest){
		"clusterId":  func(r *UpdateScheduleRequest) { r.ClusterId = "" },
		"scheduleId": func(r *UpdateScheduleRequest) { r.ScheduleId = "" },
		"cronExpr":   func(r *UpdateScheduleRequest) { r.CronExpr = "" },
		"params":     func(r *UpdateScheduleRequest) { r.Params = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateSchedule(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestDeleteScheduleRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := DeleteScheduleRequest{
		ClusterId:  "search-xxxxxxxx",
		ScheduleId: "9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e",
	}
	cases := map[string]func(*DeleteScheduleRequest){
		"clusterId":  func(r *DeleteScheduleRequest) { r.ClusterId = "" },
		"scheduleId": func(r *DeleteScheduleRequest) { r.ScheduleId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := DeleteSchedule(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestScheduleAPIPropagatesServerError uses 400 rather than 500 on purpose: the helper mirrors the
// production retry policy, which retries 5xx with backoff and would make this test take seconds.
func TestScheduleAPIPropagatesServerError(t *testing.T) {
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
		"ListSchedules": func() error {
			_, err := ListSchedules(cli, testRegion, &ListSchedulesRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"CreateSchedule": func() error {
			_, err := CreateSchedule(cli, testRegion, &CreateScheduleRequest{
				ClusterId:    "search-xxxxxxxx",
				ScheduleName: "my-delete-task",
				ScheduleType: SCHEDULE_TYPE_DELETE,
				CronExpr:     "0 30 1 * * ?",
				Params:       map[string]interface{}{"indexPattern": "logs-*"},
			})
			return err
		},
		"UpdateSchedule": func() error {
			_, err := UpdateSchedule(cli, testRegion, &UpdateScheduleRequest{
				ClusterId:  "search-xxxxxxxx",
				ScheduleId: "9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e",
				CronExpr:   "0 30 1 * * ?",
				Params:     map[string]interface{}{"indexPattern": "logs-*"},
			})
			return err
		},
		"DeleteSchedule": func() error {
			_, err := DeleteSchedule(cli, testRegion, &DeleteScheduleRequest{
				ClusterId:  "search-xxxxxxxx",
				ScheduleId: "9e1f6b2a-1c34-4e7a-9d3a-2f6a7b0c1d2e",
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

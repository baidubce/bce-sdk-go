package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

// scheduleFullCreateRequest returns a create request with every field populated.
func scheduleFullCreateRequest() CreateScheduleRequest {
	return CreateScheduleRequest{
		ClusterId:    "665018667943202816",
		Schedule:     "0 1 0 * * ?",
		ScheduleName: "CreateIndexApi1",
		TaskType:     "CREATE_INDEX",
		Task:         map[string]interface{}{"indexPatterns": []interface{}{"apiindex*"}},
	}
}

// scheduleFullUpdateRequest returns an update request with every field populated.
func scheduleFullUpdateRequest() UpdateScheduleRequest {
	return UpdateScheduleRequest{
		ClusterId:    "665018667943202816",
		Schedule:     "0 1 0 * * ?",
		ScheduleName: "CreateIndexApi1",
		TaskType:     "CREATE_INDEX",
		Task:         map[string]interface{}{"indexPatterns": []interface{}{"apiindex*"}},
	}
}

// TestCreateSchedule is migrated from v2/client_test.go TestClientCreateSchedule and additionally
// asserts the region header and every body field.
func TestCreateSchedule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/bes/cluster/schedule/create" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get(HEADER_REGION); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if body["clusterId"] != "665018667943202816" || body["schedule"] != "0 1 0 * * ?" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["scheduleName"] != "CreateIndexApi1" || body["taskType"] != "CREATE_INDEX" {
			t.Fatalf("unexpected body: %+v", body)
		}
		task, ok := body["task"].(map[string]interface{})
		if !ok {
			t.Fatalf("unexpected task: %+v", body["task"])
		}
		patterns, ok := task["indexPatterns"].([]interface{})
		if !ok || len(patterns) != 1 || patterns[0] != "apiindex*" {
			t.Fatalf("unexpected indexPatterns: %+v", task["indexPatterns"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  "result",
		})
	}))
	defer server.Close()

	request := scheduleFullCreateRequest()
	result, err := CreateSchedule(cli, testRegion, &request)
	if err != nil {
		t.Fatalf("create schedule failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestUpdateSchedule is migrated from v2/client_test.go TestClientUpdateSchedule.
func TestUpdateSchedule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/bes/cluster/schedule/update" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get(HEADER_REGION); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if body["clusterId"] != "665018667943202816" || body["schedule"] != "0 1 0 * * ?" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["scheduleName"] != "CreateIndexApi1" || body["taskType"] != "CREATE_INDEX" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if _, ok := body["task"].(map[string]interface{}); !ok {
			t.Fatalf("unexpected task: %+v", body["task"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  "result",
		})
	}))
	defer server.Close()

	request := scheduleFullUpdateRequest()
	result, err := UpdateSchedule(cli, testRegion, &request)
	if err != nil {
		t.Fatalf("update schedule failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestListSchedules is migrated from v2/client_test.go TestClientListSchedules.
func TestListSchedules(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/bes/cluster/schedule/list" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get(HEADER_REGION); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}
		if body := readJSONBody(t, r); body["clusterId"] != "665018667943202816" {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"schedules": []map[string]interface{}{
					{
						"clusterId":         "665018667943202816",
						"schedule":          "0 4 0 * * ?",
						"scheduleName":      "IndexColdApi100",
						"taskType":          "MIGRATE_COLD",
						"task":              map[string]interface{}{"minIndexAge": "48h"},
						"taskRunTimes":      0,
						"taskFailTimes":     0,
						"lastExecuteMillis": 0,
						"modifiedMillis":    1658912883143,
						"scheduleId":        "schedule-id-1",
						"modifiedVersion":   1,
						"status":            "RUNNING",
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := ListSchedules(cli, testRegion, &ListSchedulesRequest{ClusterId: "665018667943202816"})
	if err != nil {
		t.Fatalf("list schedules failed: %v", err)
	}
	if result.Result == nil || len(result.Result.Schedules) != 1 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	item := result.Result.Schedules[0]
	if item.TaskType != "MIGRATE_COLD" || item.ScheduleName != "IndexColdApi100" {
		t.Fatalf("unexpected schedule item: %+v", item)
	}
	if item.ModifiedMillis != 1658912883143 || item.ModifiedVersion != 1 || item.Status != "RUNNING" {
		t.Fatalf("unexpected schedule item: %+v", item)
	}
}

// TestDeleteSchedule is migrated from v2/client_test.go TestClientDeleteSchedule.
func TestDeleteSchedule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/bes/cluster/schedule/delete" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get(HEADER_REGION); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if body["clusterId"] != "665018667943202816" || body["scheduleName"] != "BackupApi1" {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  "result",
		})
	}))
	defer server.Close()

	result, err := DeleteSchedule(cli, testRegion, &DeleteScheduleRequest{
		ClusterId:    "665018667943202816",
		ScheduleName: "BackupApi1",
	})
	if err != nil {
		t.Fatalf("delete schedule failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestListSchedulesWithEmptyStringResult covers ScheduleListResult.UnmarshalJSON tolerating the
// server answering result as an empty string instead of an object.
func TestListSchedulesWithEmptyStringResult(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  "",
		})
	}))
	defer server.Close()

	result, err := ListSchedules(cli, testRegion, &ListSchedulesRequest{ClusterId: "665018667943202816"})
	if err != nil {
		t.Fatalf("list schedules failed: %v", err)
	}
	if result.Result == nil || len(result.Result.Schedules) != 0 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

// TestListSchedulesWithNullResult pins the asymmetry between an empty-string result and a null
// result: encoding/json never calls the custom UnmarshalJSON for null, so the pointer stays nil and
// callers must keep their nil checks. Migrated from the NullResultStaysNil subtest of
// TestEmptyStringResultIsTolerated in services/bes/v2/client_test.go, which asserted this api-layer
// contract from the v2 package.
func TestListSchedulesWithNullResult(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  nil,
		})
	}))
	defer server.Close()

	result, err := ListSchedules(cli, testRegion, &ListSchedulesRequest{ClusterId: "665018667943202816"})
	if err != nil {
		t.Fatalf("list schedules failed: %v", err)
	}
	if result.Result != nil {
		t.Fatalf("expected a nil result, got %+v", result.Result)
	}
}

// TestScheduleListResultUnmarshalJSON exercises the three paths of the custom unmarshaller,
// including the malformed-payload path that only surfaces as a decoding error.
func TestScheduleListResultUnmarshalJSON(t *testing.T) {
	emptyCases := map[string]string{
		"empty":        ``,
		"null":         `null`,
		"emptyString":  `""`,
		"paddedString": `  ""  `,
	}
	for name, data := range emptyCases {
		var result ScheduleListResult
		if err := result.UnmarshalJSON([]byte(data)); err != nil {
			t.Fatalf("%s: unmarshal failed: %v", name, err)
		}
		if result.Schedules != nil {
			t.Fatalf("%s: unexpected schedules: %+v", name, result.Schedules)
		}
	}

	var decoded ScheduleListResult
	if err := decoded.UnmarshalJSON([]byte(`{"schedules":[{"scheduleName":"a"}]}`)); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if len(decoded.Schedules) != 1 || decoded.Schedules[0].ScheduleName != "a" {
		t.Fatalf("unexpected schedules: %+v", decoded.Schedules)
	}

	var broken ScheduleListResult
	if err := broken.UnmarshalJSON([]byte(`"not-empty"`)); err == nil {
		t.Fatal("expected error for non-object result")
	}
}

func TestScheduleAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"CreateSchedule": func() error { _, err := CreateSchedule(nil, testRegion, &CreateScheduleRequest{}); return err },
		"UpdateSchedule": func() error { _, err := UpdateSchedule(nil, testRegion, &UpdateScheduleRequest{}); return err },
		"ListSchedules":  func() error { _, err := ListSchedules(nil, testRegion, &ListSchedulesRequest{}); return err },
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
		t.Error("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"CreateSchedule": func() error { _, err := CreateSchedule(cli, testRegion, nil); return err },
		"UpdateSchedule": func() error { _, err := UpdateSchedule(cli, testRegion, nil); return err },
		"ListSchedules":  func() error { _, err := ListSchedules(cli, testRegion, nil); return err },
		"DeleteSchedule": func() error { _, err := DeleteSchedule(cli, testRegion, nil); return err },
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

// TestScheduleAPIRejectsMissingRequiredFields fills every struct field and then clears one field
// per case, so each required-field guard is reached exactly once.
func TestScheduleAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Error("request should not be sent")
	}))
	defer server.Close()

	t.Run("CreateSchedule", func(t *testing.T) {
		cases := map[string]func(*CreateScheduleRequest){
			"clusterId":    func(r *CreateScheduleRequest) { r.ClusterId = "" },
			"schedule":     func(r *CreateScheduleRequest) { r.Schedule = "" },
			"scheduleName": func(r *CreateScheduleRequest) { r.ScheduleName = "" },
			"taskType":     func(r *CreateScheduleRequest) { r.TaskType = "" },
			"task":         func(r *CreateScheduleRequest) { r.Task = nil },
			"emptyTask":    func(r *CreateScheduleRequest) { r.Task = map[string]interface{}{} },
		}
		for name, clear := range cases {
			request := scheduleFullCreateRequest()
			clear(&request)
			if _, err := CreateSchedule(cli, testRegion, &request); err == nil {
				t.Fatalf("%s: expected validation error", name)
			}
		}
	})

	t.Run("UpdateSchedule", func(t *testing.T) {
		cases := map[string]func(*UpdateScheduleRequest){
			"clusterId":    func(r *UpdateScheduleRequest) { r.ClusterId = "" },
			"schedule":     func(r *UpdateScheduleRequest) { r.Schedule = "" },
			"scheduleName": func(r *UpdateScheduleRequest) { r.ScheduleName = "" },
			"taskType":     func(r *UpdateScheduleRequest) { r.TaskType = "" },
			"task":         func(r *UpdateScheduleRequest) { r.Task = nil },
		}
		for name, clear := range cases {
			request := scheduleFullUpdateRequest()
			clear(&request)
			if _, err := UpdateSchedule(cli, testRegion, &request); err == nil {
				t.Fatalf("%s: expected validation error", name)
			}
		}
	})

	t.Run("ListSchedules", func(t *testing.T) {
		if _, err := ListSchedules(cli, testRegion, &ListSchedulesRequest{}); err == nil {
			t.Fatal("clusterId: expected validation error")
		}
	})

	t.Run("DeleteSchedule", func(t *testing.T) {
		cases := map[string]func(*DeleteScheduleRequest){
			"clusterId":    func(r *DeleteScheduleRequest) { r.ClusterId = "" },
			"scheduleName": func(r *DeleteScheduleRequest) { r.ScheduleName = "" },
		}
		for name, clear := range cases {
			request := DeleteScheduleRequest{ClusterId: "665018667943202816", ScheduleName: "BackupApi1"}
			clear(&request)
			if _, err := DeleteSchedule(cli, testRegion, &request); err == nil {
				t.Fatalf("%s: expected validation error", name)
			}
		}
	})
}

// TestScheduleCheckScheduleTaskRequestAcceptsCompleteRequest covers the private validator's
// success path directly, so the helper is exercised independently of the transport.
func TestScheduleCheckScheduleTaskRequestAcceptsCompleteRequest(t *testing.T) {
	err := checkScheduleTaskRequest("cluster", "0 1 0 * * ?", "name", "CREATE_INDEX",
		map[string]interface{}{"indexPatterns": []interface{}{"a*"}}, "create schedule")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// TestScheduleAPIPropagatesServerError uses 400 rather than 500 on purpose: newTestClient mirrors
// the production retry policy, which retries 5xx with backoff and would make this test take seconds.
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

	cases := map[string]func(bce.Client) error{
		"CreateSchedule": func(cli bce.Client) error {
			request := scheduleFullCreateRequest()
			_, err := CreateSchedule(cli, testRegion, &request)
			return err
		},
		"UpdateSchedule": func(cli bce.Client) error {
			request := scheduleFullUpdateRequest()
			_, err := UpdateSchedule(cli, testRegion, &request)
			return err
		},
		"ListSchedules": func(cli bce.Client) error {
			_, err := ListSchedules(cli, testRegion, &ListSchedulesRequest{ClusterId: "665018667943202816"})
			return err
		},
		"DeleteSchedule": func(cli bce.Client) error {
			_, err := DeleteSchedule(cli, testRegion, &DeleteScheduleRequest{
				ClusterId:    "665018667943202816",
				ScheduleName: "BackupApi1",
			})
			return err
		},
	}
	for name, call := range cases {
		if err := call(cli); err == nil {
			t.Fatalf("%s: expected error for 400 response", name)
		}
	}
}

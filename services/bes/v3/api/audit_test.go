package api

import (
	"errors"
	nethttp "net/http"
	"testing"
)

func TestGetAuditStatus(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/audit" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"enabled": true,
		})
	}))
	defer server.Close()

	result, err := GetAuditStatus(cli, testRegion, &GetAuditStatusRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("get audit status failed: %v", err)
	}
	if !result.Enabled {
		t.Fatalf("unexpected enabled: %v", result.Enabled)
	}
}

func TestEnableAudit(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/audit" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		body := readJSONBody(t, r)
		if got := body["targetClusterId"]; got != "search-yyyyyyyy" {
			t.Fatalf("unexpected targetClusterId: %v", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"enabled":         true,
			"targetClusterId": "search-yyyyyyyy",
		})
	}))
	defer server.Close()

	result, err := EnableAudit(cli, testRegion, &EnableAuditRequest{
		ClusterId:       "search-xxxxxxxx",
		TargetClusterId: "search-yyyyyyyy",
	})
	if err != nil {
		t.Fatalf("enable audit failed: %v", err)
	}
	if !result.Enabled || result.TargetClusterId != "search-yyyyyyyy" {
		t.Fatalf("unexpected enable audit result: %+v", result)
	}
}

func TestListAuditEventTypes(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/audit/event-types" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"eventTypes": []string{"FAILED_LOGIN", "SSL_EXCEPTION"},
		})
	}))
	defer server.Close()

	result, err := ListAuditEventTypes(cli, testRegion, &ListAuditEventTypesRequest{ClusterId: "search-xxxxxxxx"})
	if err != nil {
		t.Fatalf("list audit event types failed: %v", err)
	}
	if len(result.EventTypes) != 2 {
		t.Fatalf("unexpected event types: %v", result.EventTypes)
	}
}

// TestSearchAuditEvents populates every optional query field so each conditional query-parameter
// branch in SearchAuditEvents is exercised.
func TestSearchAuditEvents(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/audit/events" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		query := r.URL.Query()
		expected := map[string]string{
			"startTime": "2026-07-17T02:35:31.234891Z",
			"endTime":   "2026-07-18T02:35:31.234891Z",
			"category":  "SSL_EXCEPTION",
			"indexName": "audit-index",
			"pageSize":  "20",
			"cursor":    "WzE3ODQ2MzMzMDkzNDks",
		}
		for key, want := range expected {
			if got := query.Get(key); got != want {
				t.Fatalf("unexpected %s: got %s want %s", key, got, want)
			}
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"logs": []map[string]interface{}{{
				"category":  "SSL_EXCEPTION",
				"timestamp": "2026-07-21T11:28:29.349Z",
				"content":   "{}",
			}},
			"cursor": "WzE3ODQ2MzMzMDkzNDksIlZoZ2xpWjhCWS1SRUVwQWpzR2dvIl0",
		})
	}))
	defer server.Close()

	result, err := SearchAuditEvents(cli, testRegion, &SearchAuditEventsRequest{
		ClusterId: "search-xxxxxxxx",
		StartTime: "2026-07-17T02:35:31.234891Z",
		EndTime:   "2026-07-18T02:35:31.234891Z",
		Category:  "SSL_EXCEPTION",
		IndexName: "audit-index",
		PageSize:  20,
		Cursor:    "WzE3ODQ2MzMzMDkzNDks",
	})
	if err != nil {
		t.Fatalf("search audit events failed: %v", err)
	}
	if len(result.Logs) != 1 || result.Logs[0].Category != "SSL_EXCEPTION" || result.Cursor == "" {
		t.Fatalf("unexpected search audit events result: %+v", result)
	}
}

func TestSearchAuditEventsWithoutOptionalFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.RawQuery != "" {
			t.Fatalf("unexpected query: %s", r.URL.RawQuery)
		}

		writeJSONResponse(t, w, map[string]interface{}{})
	}))
	defer server.Close()

	result, err := SearchAuditEvents(cli, testRegion, &SearchAuditEventsRequest{ClusterId: "search-xxxxxxxx"})
	if err != nil {
		t.Fatalf("search audit events failed: %v", err)
	}
	if len(result.Logs) != 0 || result.Cursor != "" {
		t.Fatalf("unexpected search audit events result: %+v", result)
	}
}

func TestGetBctAuthSwitch(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/bct/auth/switch" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "gz" {
			t.Fatalf("unexpected region: %s", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"enabled": false,
		})
	}))
	defer server.Close()

	result, err := GetBctAuthSwitch(cli, testRegion, &Request{Region: "gz"})
	if err != nil {
		t.Fatalf("get bct auth switch failed: %v", err)
	}
	if result.Enabled {
		t.Fatalf("unexpected enabled: %v", result.Enabled)
	}
}

// TestGetBctAuthSwitchWithNilRequest covers the optional-request branch: GetBctAuthSwitch accepts a
// nil request and falls back to the region argument.
func TestGetBctAuthSwitchWithNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/bct/auth/switch" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"enabled": true,
		})
	}))
	defer server.Close()

	result, err := GetBctAuthSwitch(cli, testRegion, nil)
	if err != nil {
		t.Fatalf("get bct auth switch failed: %v", err)
	}
	if !result.Enabled {
		t.Fatalf("unexpected enabled: %v", result.Enabled)
	}
}

func TestUpdateBctAuthSwitch(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/bct/auth/switch" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		body := readJSONBody(t, r)
		if got := body["enabled"]; got != true {
			t.Fatalf("unexpected enabled: %v", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"enabled": true,
		})
	}))
	defer server.Close()

	result, err := UpdateBctAuthSwitch(cli, testRegion, &UpdateBctAuthSwitchRequest{Enabled: boolPtr(true)})
	if err != nil {
		t.Fatalf("update bct auth switch failed: %v", err)
	}
	if !result.Enabled {
		t.Fatalf("unexpected enabled: %v", result.Enabled)
	}
}

func TestSearchBctEvents(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/bct/events/query" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}

		body := readJSONBody(t, r)
		if got := body["startTime"]; got != "2026-07-14T01:51:33.355981Z" {
			t.Fatalf("unexpected startTime: %v", got)
		}
		if got := body["endTime"]; got != "2026-07-24T01:51:33.355981Z" {
			t.Fatalf("unexpected endTime: %v", got)
		}
		if got := body["domainId"]; got != "domain-xxxx" {
			t.Fatalf("unexpected domainId: %v", got)
		}
		if got := body["nextMarker"]; got != "CNiCgtH2Mw" {
			t.Fatalf("unexpected nextMarker: %v", got)
		}
		filters, ok := body["filters"].([]interface{})
		if !ok || len(filters) != 1 {
			t.Fatalf("unexpected filters: %v", body["filters"])
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"pageSize": 20,
			"data": []map[string]interface{}{{
				"eventType":   "console",
				"eventSource": "bes",
				"eventName":   "createCluster",
				"success":     true,
				"userIdentity": map[string]interface{}{
					"iamDomainId": "domain-xxxx",
					"iamUserId":   "user-xxxx",
				},
			}},
			"nextMarker": "CNiCgtH2MxCdgM6khNH2MyAB",
			"truncated":  false,
		})
	}))
	defer server.Close()

	result, err := SearchBctEvents(cli, testRegion, &SearchBctEventsRequest{
		StartTime:  "2026-07-14T01:51:33.355981Z",
		EndTime:    "2026-07-24T01:51:33.355981Z",
		PageSize:   20,
		DomainId:   "domain-xxxx",
		NextMarker: "CNiCgtH2Mw",
		Filters:    []BctQueryFilter{{Field: "eventSource", Value: "bes"}},
	})
	if err != nil {
		t.Fatalf("search bct events failed: %v", err)
	}
	if len(result.Data) != 1 || result.Data[0].EventName != "createCluster" {
		t.Fatalf("unexpected search bct events result: %+v", result)
	}
	if result.Data[0].UserIdentity == nil || result.Data[0].UserIdentity.IamUserId != "user-xxxx" {
		t.Fatalf("unexpected user identity: %+v", result.Data[0].UserIdentity)
	}
	if result.NextMarker == "" {
		t.Fatalf("expected nextMarker to be set")
	}
}

func TestAuditAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"GetAuditStatus": func() error { _, err := GetAuditStatus(nil, testRegion, &GetAuditStatusRequest{}); return err },
		"EnableAudit":    func() error { _, err := EnableAudit(nil, testRegion, &EnableAuditRequest{}); return err },
		"ListAuditEventTypes": func() error {
			_, err := ListAuditEventTypes(nil, testRegion, &ListAuditEventTypesRequest{})
			return err
		},
		"SearchAuditEvents": func() error {
			_, err := SearchAuditEvents(nil, testRegion, &SearchAuditEventsRequest{})
			return err
		},
		"GetBctAuthSwitch": func() error { _, err := GetBctAuthSwitch(nil, testRegion, &Request{}); return err },
		"UpdateBctAuthSwitch": func() error {
			_, err := UpdateBctAuthSwitch(nil, testRegion, &UpdateBctAuthSwitchRequest{})
			return err
		},
		"SearchBctEvents": func() error {
			_, err := SearchBctEvents(nil, testRegion, &SearchBctEventsRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

// TestAuditAPIRejectsNilRequest excludes GetBctAuthSwitch on purpose: its request is optional.
func TestAuditAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"GetAuditStatus":      func() error { _, err := GetAuditStatus(cli, testRegion, nil); return err },
		"EnableAudit":         func() error { _, err := EnableAudit(cli, testRegion, nil); return err },
		"ListAuditEventTypes": func() error { _, err := ListAuditEventTypes(cli, testRegion, nil); return err },
		"SearchAuditEvents":   func() error { _, err := SearchAuditEvents(cli, testRegion, nil); return err },
		"UpdateBctAuthSwitch": func() error { _, err := UpdateBctAuthSwitch(cli, testRegion, nil); return err },
		"SearchBctEvents":     func() error { _, err := SearchBctEvents(cli, testRegion, nil); return err },
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

func TestAuditAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"GetAuditStatus/clusterId": func() error {
			_, err := GetAuditStatus(cli, testRegion, &GetAuditStatusRequest{})
			return err
		},
		"EnableAudit/clusterId": func() error {
			_, err := EnableAudit(cli, testRegion, &EnableAuditRequest{TargetClusterId: "search-yyyyyyyy"})
			return err
		},
		"ListAuditEventTypes/clusterId": func() error {
			_, err := ListAuditEventTypes(cli, testRegion, &ListAuditEventTypesRequest{})
			return err
		},
		"SearchAuditEvents/clusterId": func() error {
			_, err := SearchAuditEvents(cli, testRegion, &SearchAuditEventsRequest{StartTime: "2026-07-17T02:35:31Z"})
			return err
		},
		"UpdateBctAuthSwitch/enabled": func() error {
			_, err := UpdateBctAuthSwitch(cli, testRegion, &UpdateBctAuthSwitchRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestSearchBctEventsRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := SearchBctEventsRequest{
		StartTime: "2026-07-14T01:51:33.355981Z",
		EndTime:   "2026-07-24T01:51:33.355981Z",
	}
	cases := map[string]func(*SearchBctEventsRequest){
		"startTime": func(r *SearchBctEventsRequest) { r.StartTime = "" },
		"endTime":   func(r *SearchBctEventsRequest) { r.EndTime = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := SearchBctEvents(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestAuditAPIPropagatesServerError uses 400 rather than 500 on purpose: the helper mirrors the
// production retry policy, which retries 5xx with backoff and would make this test take seconds.
func TestAuditAPIPropagatesServerError(t *testing.T) {
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
		"GetAuditStatus": func() error {
			_, err := GetAuditStatus(cli, testRegion, &GetAuditStatusRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"EnableAudit": func() error {
			_, err := EnableAudit(cli, testRegion, &EnableAuditRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"ListAuditEventTypes": func() error {
			_, err := ListAuditEventTypes(cli, testRegion, &ListAuditEventTypesRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"SearchAuditEvents": func() error {
			_, err := SearchAuditEvents(cli, testRegion, &SearchAuditEventsRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"GetBctAuthSwitch": func() error {
			_, err := GetBctAuthSwitch(cli, testRegion, &Request{})
			return err
		},
		"UpdateBctAuthSwitch": func() error {
			_, err := UpdateBctAuthSwitch(cli, testRegion, &UpdateBctAuthSwitchRequest{Enabled: boolPtr(false)})
			return err
		},
		"SearchBctEvents": func() error {
			_, err := SearchBctEvents(cli, testRegion, &SearchBctEventsRequest{
				StartTime: "2026-07-14T01:51:33.355981Z",
				EndTime:   "2026-07-24T01:51:33.355981Z",
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

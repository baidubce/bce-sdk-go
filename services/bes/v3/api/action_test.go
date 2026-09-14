package api

import (
	"errors"
	nethttp "net/http"
	"testing"
)

func TestListActions(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/actions" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		if got := r.URL.Query().Get("pageNo"); got != "2" {
			t.Fatalf("unexpected pageNo: %s", got)
		}
		if got := r.URL.Query().Get("pageSize"); got != "10" {
			t.Fatalf("unexpected pageSize: %s", got)
		}
		if got := r.URL.Query().Get("name"); got != "DEPLOY_CLUSTER" {
			t.Fatalf("unexpected name: %s", got)
		}
		if got := r.URL.Query().Get("status"); got != "RUNNING" {
			t.Fatalf("unexpected status: %s", got)
		}
		if got := r.URL.Query().Get("orderBy"); got != "createTime" {
			t.Fatalf("unexpected orderBy: %s", got)
		}
		if got := r.URL.Query().Get("order"); got != "desc" {
			t.Fatalf("unexpected order: %s", got)
		}
		if got := r.URL.Query().Get("clusterName"); got != "test-cluster" {
			t.Fatalf("unexpected clusterName: %s", got)
		}
		if got := r.URL.Query().Get("clusterId"); got != "search-xxxx" {
			t.Fatalf("unexpected clusterId: %s", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"result": []map[string]interface{}{{
				"clusterId":   "search-xxxx",
				"clusterName": "test-cluster",
				"name":        "DEPLOY_CLUSTER",
				"nameCN":      "创建集群",
				"actionId":    "action-xxxx",
				"status":      "RUNNING",
				"createTime":  "2026-07-09T03:31:29Z",
				"endTime":     "2026-07-09T03:41:29Z",
				"runningTime": "10m0s",
			}},
			"pageNo":     2,
			"pageSize":   10,
			"totalCount": 1,
		})
	}))
	defer server.Close()

	result, err := ListActions(cli, testRegion, &ListActionsRequest{
		Request:     Request{Region: "sh"},
		PageNo:      2,
		PageSize:    10,
		Name:        "DEPLOY_CLUSTER",
		Status:      "RUNNING",
		OrderBy:     "createTime",
		Order:       "desc",
		ClusterName: "test-cluster",
		ClusterId:   "search-xxxx",
	})
	if err != nil {
		t.Fatalf("list actions failed: %v", err)
	}
	if len(result.Result) != 1 || result.Result[0].ActionId != "action-xxxx" || result.TotalCount != 1 {
		t.Fatalf("unexpected list actions result: %+v", result)
	}
}

func TestListActionTypes(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/actions/action-types" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"actionTypes": []map[string]interface{}{{
				"name":   "DEPLOY_CLUSTER",
				"nameCN": "创建集群",
			}},
		})
	}))
	defer server.Close()

	result, err := ListActionTypes(cli, testRegion, &ListActionTypesRequest{})
	if err != nil {
		t.Fatalf("list action types failed: %v", err)
	}
	if len(result.ActionTypes) != 1 || result.ActionTypes[0].Name != "DEPLOY_CLUSTER" {
		t.Fatalf("unexpected list action types result: %+v", result)
	}
}

func TestListOperations(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/actions/action-xxxx/operations" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-xxxx",
			"name":      "DEPLOY_CLUSTER",
			"nameCN":    "创建集群",
			"status":    "RUNNING",
			"operations": []map[string]interface{}{{
				"actionId":    "action-xxxx",
				"operationId": "operation-xxxx",
				"type":        "CREATE",
				"typeCN":      "创建",
				"state":       "RUNNING",
				"process":     80,
				"startTime":   "2026-07-09T03:31:29Z",
				"endTime":     "2026-07-09T03:41:29Z",
				"runningTime": "10m0s",
			}},
		})
	}))
	defer server.Close()

	result, err := ListOperations(cli, testRegion, &ListOperationsRequest{
		ClusterId: "search-xxxx",
		ActionId:  "action-xxxx",
	})
	if err != nil {
		t.Fatalf("list operations failed: %v", err)
	}
	if len(result.Operations) != 1 || result.Operations[0].OperationId != "operation-xxxx" {
		t.Fatalf("unexpected list operations result: %+v", result)
	}
}

func TestGetOperation(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/actions/action-xxxx/operations/operation-xxxx" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"actionId":      "action-xxxx",
			"operationId":   "operation-xxxx",
			"clusterId":     "search-xxxx",
			"process":       80,
			"sourceContext": "source",
			"targetContext": "target",
			"groups": []map[string]interface{}{{
				"groupName":   "CHECK_DISK",
				"groupNameCN": "磁盘检查",
				"state":       "RUNNING",
				"analysis":    true,
				"diagnosis":   "ok",
				"diagnosisCN": "正常",
			}},
		})
	}))
	defer server.Close()

	result, err := GetOperation(cli, testRegion, &GetOperationRequest{
		ClusterId:   "search-xxxx",
		ActionId:    "action-xxxx",
		OperationId: "operation-xxxx",
	})
	if err != nil {
		t.Fatalf("get operation failed: %v", err)
	}
	if len(result.Groups) != 1 || result.Groups[0].GroupName != "CHECK_DISK" || result.Groups[0].Analysis == nil || !*result.Groups[0].Analysis {
		t.Fatalf("unexpected get operation result: %+v", result)
	}
}

func TestListOperationAnalysisDetails(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/actions/action-xxxx/operations/operation-xxxx/groups/CHECK_DISK" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("pageNo"); got != "1" {
			t.Fatalf("unexpected pageNo: %s", got)
		}
		if got := r.URL.Query().Get("pageSize"); got != "20" {
			t.Fatalf("unexpected pageSize: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"result": []map[string]interface{}{{
				"name":   "DISK_FULL",
				"nameCN": "磁盘已满",
				"state":  "RUNNING",
			}},
			"pageNo":     1,
			"pageSize":   20,
			"totalCount": 1,
		})
	}))
	defer server.Close()

	result, err := ListOperationAnalysisDetails(cli, testRegion, &ListOperationAnalysisDetailsRequest{
		ClusterId:   "search-xxxx",
		ActionId:    "action-xxxx",
		OperationId: "operation-xxxx",
		GroupNames:  "CHECK_DISK",
		PageNo:      1,
		PageSize:    20,
	})
	if err != nil {
		t.Fatalf("list operation analysis details failed: %v", err)
	}
	if len(result.Result) != 1 || result.Result[0].Name != "DISK_FULL" || result.TotalCount != 1 {
		t.Fatalf("unexpected list operation analysis details result: %+v", result)
	}
}

func TestActionAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"ListActions":     func() error { _, err := ListActions(nil, testRegion, &ListActionsRequest{}); return err },
		"ListActionTypes": func() error { _, err := ListActionTypes(nil, testRegion, &ListActionTypesRequest{}); return err },
		"ListOperations":  func() error { _, err := ListOperations(nil, testRegion, &ListOperationsRequest{}); return err },
		"GetOperation":    func() error { _, err := GetOperation(nil, testRegion, &GetOperationRequest{}); return err },
		"ListOperationAnalysisDetails": func() error {
			_, err := ListOperationAnalysisDetails(nil, testRegion, &ListOperationAnalysisDetailsRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

func TestActionAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ListActions":     func() error { _, err := ListActions(cli, testRegion, nil); return err },
		"ListActionTypes": func() error { _, err := ListActionTypes(cli, testRegion, nil); return err },
		"ListOperations":  func() error { _, err := ListOperations(cli, testRegion, nil); return err },
		"GetOperation":    func() error { _, err := GetOperation(cli, testRegion, nil); return err },
		"ListOperationAnalysisDetails": func() error {
			_, err := ListOperationAnalysisDetails(cli, testRegion, nil)
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

func TestActionAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ListOperations/clusterId": func() error {
			_, err := ListOperations(cli, testRegion, &ListOperationsRequest{ActionId: "action-xxxx"})
			return err
		},
		"ListOperations/actionId": func() error {
			_, err := ListOperations(cli, testRegion, &ListOperationsRequest{ClusterId: "search-xxxx"})
			return err
		},
		"GetOperation/clusterId": func() error {
			_, err := GetOperation(cli, testRegion, &GetOperationRequest{ActionId: "action-xxxx", OperationId: "operation-xxxx"})
			return err
		},
		"GetOperation/actionId": func() error {
			_, err := GetOperation(cli, testRegion, &GetOperationRequest{ClusterId: "search-xxxx", OperationId: "operation-xxxx"})
			return err
		},
		"GetOperation/operationId": func() error {
			_, err := GetOperation(cli, testRegion, &GetOperationRequest{ClusterId: "search-xxxx", ActionId: "action-xxxx"})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestListOperationAnalysisDetailsRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := ListOperationAnalysisDetailsRequest{
		ClusterId:   "search-xxxx",
		ActionId:    "action-xxxx",
		OperationId: "operation-xxxx",
		GroupNames:  "CHECK_DISK",
	}
	cases := map[string]func(*ListOperationAnalysisDetailsRequest){
		"clusterId":   func(r *ListOperationAnalysisDetailsRequest) { r.ClusterId = "" },
		"actionId":    func(r *ListOperationAnalysisDetailsRequest) { r.ActionId = "" },
		"operationId": func(r *ListOperationAnalysisDetailsRequest) { r.OperationId = "" },
		"groupNames":  func(r *ListOperationAnalysisDetailsRequest) { r.GroupNames = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := ListOperationAnalysisDetails(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestActionAPIPropagatesServerError uses 400 rather than 500 on purpose: the helper mirrors the
// production retry policy, which retries 5xx with backoff and would make this test take seconds.
func TestActionAPIPropagatesServerError(t *testing.T) {
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
		"ListActions":     func() error { _, err := ListActions(cli, testRegion, &ListActionsRequest{}); return err },
		"ListActionTypes": func() error { _, err := ListActionTypes(cli, testRegion, &ListActionTypesRequest{}); return err },
		"ListOperations": func() error {
			_, err := ListOperations(cli, testRegion, &ListOperationsRequest{ClusterId: "search-xxxx", ActionId: "action-xxxx"})
			return err
		},
		"GetOperation": func() error {
			_, err := GetOperation(cli, testRegion, &GetOperationRequest{
				ClusterId: "search-xxxx", ActionId: "action-xxxx", OperationId: "operation-xxxx",
			})
			return err
		},
		"ListOperationAnalysisDetails": func() error {
			_, err := ListOperationAnalysisDetails(cli, testRegion, &ListOperationAnalysisDetailsRequest{
				ClusterId: "search-xxxx", ActionId: "action-xxxx", OperationId: "operation-xxxx", GroupNames: "CHECK_DISK",
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

func TestActionURIWithoutParts(t *testing.T) {
	if got := actionURI(); got != URI_ACTIONS {
		t.Fatalf("actionURI() = %q, want %q", got, URI_ACTIONS)
	}
}

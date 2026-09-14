package api

import (
	"errors"
	nethttp "net/http"
	"testing"
)

func TestListWeeklyOverview(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/overview/weekly" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"count": 7,
			"riskItems": []map[string]interface{}{
				{"id": "ClusterPublicAccess", "highRiskCount": 7, "lowRiskCount": 0},
			},
		})
	}))
	defer server.Close()

	result, err := ListWeeklyOverview(cli, testRegion, &ListWeeklyOverviewRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("list weekly overview failed: %v", err)
	}
	if result.Count != 7 || len(result.RiskItems) != 1 || result.RiskItems[0].Id != "ClusterPublicAccess" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetManualDiagnosisCount(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/manual/count" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"count": 3})
	}))
	defer server.Close()

	result, err := GetManualDiagnosisCount(cli, testRegion, &GetManualDiagnosisCountRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("get manual diagnosis count failed: %v", err)
	}
	if result.Count != 3 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetManualDiagnosisConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/manual" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"indices": "*",
			"items":   []string{"ClusterHealth"},
		})
	}))
	defer server.Close()

	result, err := GetManualDiagnosisConfig(cli, testRegion, &GetManualDiagnosisConfigRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("get manual diagnosis config failed: %v", err)
	}
	if result.Indices != "*" || len(result.Items) != 1 || result.Items[0] != "ClusterHealth" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestUpdateManualDiagnosisConfig populates the optional indices field so the request body carries
// every field the API accepts.
func TestUpdateManualDiagnosisConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/manual" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if body["indices"] != "*" {
			t.Fatalf("unexpected body: %+v", body)
		}
		items, ok := body["items"].([]interface{})
		if !ok || len(items) != 1 || items[0] != "ClusterHealth" {
			t.Fatalf("unexpected items: %+v", body["items"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := UpdateManualDiagnosisConfig(cli, testRegion, &UpdateManualDiagnosisConfigRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
		Indices:   "*",
		Items:     []string{"ClusterHealth"},
	})
	if err != nil {
		t.Fatalf("update manual diagnosis config failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestCreateManualDiagnosis(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/manual" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := CreateManualDiagnosis(cli, testRegion, &CreateManualDiagnosisRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("create manual diagnosis failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestListDiagnosisReports populates both optional pagination fields so each conditional
// query-parameter branch in ListDiagnosisReports is exercised.
func TestListDiagnosisReports(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/reports" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		if got := r.URL.Query().Get("pageNo"); got != "1" {
			t.Fatalf("unexpected pageNo: %s", got)
		}
		if got := r.URL.Query().Get("pageSize"); got != "10" {
			t.Fatalf("unexpected pageSize: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"pageNo":     1,
			"pageSize":   10,
			"totalCount": 1,
			"reports": []map[string]interface{}{
				{"reportId": "report-xxxxxxxx", "type": "auto", "execTime": "2026-07-28 02:00:00", "lowRiskCount": 1, "highRiskCount": 1, "errorCount": 0},
			},
		})
	}))
	defer server.Close()

	result, err := ListDiagnosisReports(cli, testRegion, &ListDiagnosisReportsRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
		PageNo:    1,
		PageSize:  10,
	})
	if err != nil {
		t.Fatalf("list diagnosis reports failed: %v", err)
	}
	if result.TotalCount != 1 || len(result.Reports) != 1 || result.Reports[0].ReportId != "report-xxxxxxxx" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestListDiagnosisReportsWithoutPagination covers the request path where the optional pageNo and
// pageSize parameters are omitted from the query string entirely.
func TestListDiagnosisReportsWithoutPagination(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.RawQuery != "" {
			t.Fatalf("unexpected query: %s", r.URL.RawQuery)
		}
		if got := r.Header.Get("X-Region"); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"totalCount": 0})
	}))
	defer server.Close()

	result, err := ListDiagnosisReports(cli, testRegion, &ListDiagnosisReportsRequest{ClusterId: "search-xxxxxxxx"})
	if err != nil {
		t.Fatalf("list diagnosis reports failed: %v", err)
	}
	if result.TotalCount != 0 || len(result.Reports) != 0 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetDiagnosisReport(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/reports/report-xxxxxxxx" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"state":      "success",
			"type":       "auto",
			"indices":    "*",
			"totalCount": 1,
			"execTime":   "2026-07-28 02:00:00",
			"groups": []map[string]interface{}{
				{
					"grade":     "HIGH_RISK",
					"itemCount": 1,
					"results": []map[string]interface{}{
						{"id": "ClusterPublicAccess", "type": "cluster", "conclusion": "危险", "hasLowRiskAdvice": false, "notExistIndex": false},
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := GetDiagnosisReport(cli, testRegion, &GetDiagnosisReportRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
		ReportId:  "report-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("get diagnosis report failed: %v", err)
	}
	if result.State != "success" || len(result.Groups) != 1 || result.Groups[0].Grade != DIAGNOSIS_GRADE_HIGH_RISK {
		t.Fatalf("unexpected result: %+v", result)
	}
	if len(result.Groups[0].Results) != 1 || result.Groups[0].Results[0].Id != "ClusterPublicAccess" {
		t.Fatalf("unexpected results: %+v", result.Groups[0].Results)
	}
}

func TestGetDiagnosisBusyStatus(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/busy" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"busy": true})
	}))
	defer server.Close()

	result, err := GetDiagnosisBusyStatus(cli, testRegion, &GetDiagnosisBusyStatusRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("get diagnosis busy status failed: %v", err)
	}
	if !result.Busy {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetDiagnosisAuthorizationStatus(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/authorize" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		// GET 返回的是 enabled，不是 authorized（见 GetDiagnosisAuthorizationStatusResponse 注释）。
		writeJSONResponse(t, w, map[string]interface{}{"enabled": true})
	}))
	defer server.Close()

	result, err := GetDiagnosisAuthorizationStatus(cli, testRegion, &GetDiagnosisAuthorizationStatusRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("get diagnosis authorization status failed: %v", err)
	}
	if !result.Enabled {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestAuthorizeDiagnosis(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/authorize" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"authorized": true})
	}))
	defer server.Close()

	result, err := AuthorizeDiagnosis(cli, testRegion, &AuthorizeDiagnosisRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("authorize diagnosis failed: %v", err)
	}
	if !result.Authorized {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestListDiagnosisItems(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/items" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"items": []map[string]interface{}{
				{"id": "ClusterHealth", "type": "cluster"},
			},
		})
	}))
	defer server.Close()

	result, err := ListDiagnosisItems(cli, testRegion, &ListDiagnosisItemsRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("list diagnosis items failed: %v", err)
	}
	if len(result.Items) != 1 || result.Items[0].Id != "ClusterHealth" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetAutoDiagnosisStatus(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/auto" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"enabled": true})
	}))
	defer server.Close()

	result, err := GetAutoDiagnosisStatus(cli, testRegion, &GetAutoDiagnosisStatusRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("get auto diagnosis status failed: %v", err)
	}
	if !result.Enabled {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestUpdateAutoDiagnosis(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/auto" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if body["enabled"] != true {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"enabled": true})
	}))
	defer server.Close()

	result, err := UpdateAutoDiagnosis(cli, testRegion, &UpdateAutoDiagnosisRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
		Enabled:   boolPtr(true),
	})
	if err != nil {
		t.Fatalf("update auto diagnosis failed: %v", err)
	}
	if !result.Enabled {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetLatestOverview(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/diagnosis/overview/latest" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"reportId":      "report-xxxxxxxx",
			"status":        "success",
			"execTime":      "2026-07-28 02:00:00",
			"countSafe":     16,
			"countLowRisk":  1,
			"countHighRisk": 1,
			"countError":    0,
		})
	}))
	defer server.Close()

	result, err := GetLatestOverview(cli, testRegion, &GetLatestOverviewRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("get latest overview failed: %v", err)
	}
	if result.ReportId != "report-xxxxxxxx" || result.CountSafe != 16 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestDiagnosisAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"ListWeeklyOverview": func() error {
			_, err := ListWeeklyOverview(nil, testRegion, &ListWeeklyOverviewRequest{})
			return err
		},
		"GetManualDiagnosisCount": func() error {
			_, err := GetManualDiagnosisCount(nil, testRegion, &GetManualDiagnosisCountRequest{})
			return err
		},
		"GetManualDiagnosisConfig": func() error {
			_, err := GetManualDiagnosisConfig(nil, testRegion, &GetManualDiagnosisConfigRequest{})
			return err
		},
		"UpdateManualDiagnosisConfig": func() error {
			_, err := UpdateManualDiagnosisConfig(nil, testRegion, &UpdateManualDiagnosisConfigRequest{})
			return err
		},
		"CreateManualDiagnosis": func() error {
			_, err := CreateManualDiagnosis(nil, testRegion, &CreateManualDiagnosisRequest{})
			return err
		},
		"ListDiagnosisReports": func() error {
			_, err := ListDiagnosisReports(nil, testRegion, &ListDiagnosisReportsRequest{})
			return err
		},
		"GetDiagnosisReport": func() error {
			_, err := GetDiagnosisReport(nil, testRegion, &GetDiagnosisReportRequest{})
			return err
		},
		"GetDiagnosisBusyStatus": func() error {
			_, err := GetDiagnosisBusyStatus(nil, testRegion, &GetDiagnosisBusyStatusRequest{})
			return err
		},
		"GetDiagnosisAuthorizationStatus": func() error {
			_, err := GetDiagnosisAuthorizationStatus(nil, testRegion, &GetDiagnosisAuthorizationStatusRequest{})
			return err
		},
		"AuthorizeDiagnosis": func() error {
			_, err := AuthorizeDiagnosis(nil, testRegion, &AuthorizeDiagnosisRequest{})
			return err
		},
		"ListDiagnosisItems": func() error {
			_, err := ListDiagnosisItems(nil, testRegion, &ListDiagnosisItemsRequest{})
			return err
		},
		"GetAutoDiagnosisStatus": func() error {
			_, err := GetAutoDiagnosisStatus(nil, testRegion, &GetAutoDiagnosisStatusRequest{})
			return err
		},
		"UpdateAutoDiagnosis": func() error {
			_, err := UpdateAutoDiagnosis(nil, testRegion, &UpdateAutoDiagnosisRequest{})
			return err
		},
		"GetLatestOverview": func() error {
			_, err := GetLatestOverview(nil, testRegion, &GetLatestOverviewRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

// TestDiagnosisAPIRejectsNilRequest covers every diagnosis API: all of them treat the request as
// mandatory, so none of them has an optional-request success path.
func TestDiagnosisAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ListWeeklyOverview":       func() error { _, err := ListWeeklyOverview(cli, testRegion, nil); return err },
		"GetManualDiagnosisCount":  func() error { _, err := GetManualDiagnosisCount(cli, testRegion, nil); return err },
		"GetManualDiagnosisConfig": func() error { _, err := GetManualDiagnosisConfig(cli, testRegion, nil); return err },
		"UpdateManualDiagnosisConfig": func() error {
			_, err := UpdateManualDiagnosisConfig(cli, testRegion, nil)
			return err
		},
		"CreateManualDiagnosis":  func() error { _, err := CreateManualDiagnosis(cli, testRegion, nil); return err },
		"ListDiagnosisReports":   func() error { _, err := ListDiagnosisReports(cli, testRegion, nil); return err },
		"GetDiagnosisReport":     func() error { _, err := GetDiagnosisReport(cli, testRegion, nil); return err },
		"GetDiagnosisBusyStatus": func() error { _, err := GetDiagnosisBusyStatus(cli, testRegion, nil); return err },
		"GetDiagnosisAuthorizationStatus": func() error {
			_, err := GetDiagnosisAuthorizationStatus(cli, testRegion, nil)
			return err
		},
		"AuthorizeDiagnosis":     func() error { _, err := AuthorizeDiagnosis(cli, testRegion, nil); return err },
		"ListDiagnosisItems":     func() error { _, err := ListDiagnosisItems(cli, testRegion, nil); return err },
		"GetAutoDiagnosisStatus": func() error { _, err := GetAutoDiagnosisStatus(cli, testRegion, nil); return err },
		"UpdateAutoDiagnosis":    func() error { _, err := UpdateAutoDiagnosis(cli, testRegion, nil); return err },
		"GetLatestOverview":      func() error { _, err := GetLatestOverview(cli, testRegion, nil); return err },
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

// TestDiagnosisAPIRejectsMissingClusterId covers the clusterId check of every diagnosis API whose
// only required field is clusterId. The three APIs with a second required field get their own
// full-struct/clear-one-field tables below.
func TestDiagnosisAPIRejectsMissingClusterId(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ListWeeklyOverview": func() error {
			_, err := ListWeeklyOverview(cli, testRegion, &ListWeeklyOverviewRequest{})
			return err
		},
		"GetManualDiagnosisCount": func() error {
			_, err := GetManualDiagnosisCount(cli, testRegion, &GetManualDiagnosisCountRequest{})
			return err
		},
		"GetManualDiagnosisConfig": func() error {
			_, err := GetManualDiagnosisConfig(cli, testRegion, &GetManualDiagnosisConfigRequest{})
			return err
		},
		"CreateManualDiagnosis": func() error {
			_, err := CreateManualDiagnosis(cli, testRegion, &CreateManualDiagnosisRequest{})
			return err
		},
		"ListDiagnosisReports": func() error {
			_, err := ListDiagnosisReports(cli, testRegion, &ListDiagnosisReportsRequest{PageNo: 1, PageSize: 10})
			return err
		},
		"GetDiagnosisBusyStatus": func() error {
			_, err := GetDiagnosisBusyStatus(cli, testRegion, &GetDiagnosisBusyStatusRequest{})
			return err
		},
		"GetDiagnosisAuthorizationStatus": func() error {
			_, err := GetDiagnosisAuthorizationStatus(cli, testRegion, &GetDiagnosisAuthorizationStatusRequest{})
			return err
		},
		"AuthorizeDiagnosis": func() error {
			_, err := AuthorizeDiagnosis(cli, testRegion, &AuthorizeDiagnosisRequest{})
			return err
		},
		"ListDiagnosisItems": func() error {
			_, err := ListDiagnosisItems(cli, testRegion, &ListDiagnosisItemsRequest{})
			return err
		},
		"GetAutoDiagnosisStatus": func() error {
			_, err := GetAutoDiagnosisStatus(cli, testRegion, &GetAutoDiagnosisStatusRequest{})
			return err
		},
		"GetLatestOverview": func() error {
			_, err := GetLatestOverview(cli, testRegion, &GetLatestOverviewRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateManualDiagnosisConfigRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := UpdateManualDiagnosisConfigRequest{
		ClusterId: "search-xxxxxxxx",
		Indices:   "*",
		Items:     []string{"ClusterHealth"},
	}
	cases := map[string]func(*UpdateManualDiagnosisConfigRequest){
		"clusterId": func(r *UpdateManualDiagnosisConfigRequest) { r.ClusterId = "" },
		"items":     func(r *UpdateManualDiagnosisConfigRequest) { r.Items = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateManualDiagnosisConfig(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestGetDiagnosisReportRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := GetDiagnosisReportRequest{
		ClusterId: "search-xxxxxxxx",
		ReportId:  "report-xxxxxxxx",
	}
	cases := map[string]func(*GetDiagnosisReportRequest){
		"clusterId": func(r *GetDiagnosisReportRequest) { r.ClusterId = "" },
		"reportId":  func(r *GetDiagnosisReportRequest) { r.ReportId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := GetDiagnosisReport(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateAutoDiagnosisRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := UpdateAutoDiagnosisRequest{
		ClusterId: "search-xxxxxxxx",
		Enabled:   boolPtr(true),
	}
	cases := map[string]func(*UpdateAutoDiagnosisRequest){
		"clusterId": func(r *UpdateAutoDiagnosisRequest) { r.ClusterId = "" },
		"enabled":   func(r *UpdateAutoDiagnosisRequest) { r.Enabled = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateAutoDiagnosis(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestDiagnosisAPIPropagatesServerError uses 400 rather than 500 on purpose: the helper mirrors the
// production retry policy, which retries 5xx with backoff and would make this test take seconds.
func TestDiagnosisAPIPropagatesServerError(t *testing.T) {
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
		"ListWeeklyOverview": func() error {
			_, err := ListWeeklyOverview(cli, testRegion, &ListWeeklyOverviewRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"GetManualDiagnosisCount": func() error {
			_, err := GetManualDiagnosisCount(cli, testRegion, &GetManualDiagnosisCountRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"GetManualDiagnosisConfig": func() error {
			_, err := GetManualDiagnosisConfig(cli, testRegion, &GetManualDiagnosisConfigRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"UpdateManualDiagnosisConfig": func() error {
			_, err := UpdateManualDiagnosisConfig(cli, testRegion, &UpdateManualDiagnosisConfigRequest{
				ClusterId: "search-xxxxxxxx",
				Indices:   "*",
				Items:     []string{"ClusterHealth"},
			})
			return err
		},
		"CreateManualDiagnosis": func() error {
			_, err := CreateManualDiagnosis(cli, testRegion, &CreateManualDiagnosisRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"ListDiagnosisReports": func() error {
			_, err := ListDiagnosisReports(cli, testRegion, &ListDiagnosisReportsRequest{
				ClusterId: "search-xxxxxxxx",
				PageNo:    1,
				PageSize:  10,
			})
			return err
		},
		"GetDiagnosisReport": func() error {
			_, err := GetDiagnosisReport(cli, testRegion, &GetDiagnosisReportRequest{
				ClusterId: "search-xxxxxxxx",
				ReportId:  "report-xxxxxxxx",
			})
			return err
		},
		"GetDiagnosisBusyStatus": func() error {
			_, err := GetDiagnosisBusyStatus(cli, testRegion, &GetDiagnosisBusyStatusRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"GetDiagnosisAuthorizationStatus": func() error {
			_, err := GetDiagnosisAuthorizationStatus(cli, testRegion, &GetDiagnosisAuthorizationStatusRequest{
				ClusterId: "search-xxxxxxxx",
			})
			return err
		},
		"AuthorizeDiagnosis": func() error {
			_, err := AuthorizeDiagnosis(cli, testRegion, &AuthorizeDiagnosisRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"ListDiagnosisItems": func() error {
			_, err := ListDiagnosisItems(cli, testRegion, &ListDiagnosisItemsRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"GetAutoDiagnosisStatus": func() error {
			_, err := GetAutoDiagnosisStatus(cli, testRegion, &GetAutoDiagnosisStatusRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"UpdateAutoDiagnosis": func() error {
			_, err := UpdateAutoDiagnosis(cli, testRegion, &UpdateAutoDiagnosisRequest{
				ClusterId: "search-xxxxxxxx",
				Enabled:   boolPtr(true),
			})
			return err
		},
		"GetLatestOverview": func() error {
			_, err := GetLatestOverview(cli, testRegion, &GetLatestOverviewRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for 400 response", name)
		}
	}
}

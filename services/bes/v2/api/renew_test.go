package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	renewTestClusterId = "443755843180171264"
	renewTestOrderId   = "186f2566f2884534a65ed8cdc2499d82"
)

// renewWriteBadRequest writes a 400 service error. A 4xx status is used on purpose: the test
// client mirrors the production DEFAULT_RETRY_POLICY, so a 5xx would trigger backoff retries.
func renewWriteBadRequest(t *testing.T, w nethttp.ResponseWriter) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(nethttp.StatusBadRequest)
	if _, err := w.Write([]byte(`{"code":"BadRequest","message":"invalid renew request"}`)); err != nil {
		t.Fatalf("write error response failed: %v", err)
	}
}

// renewUnreachableHandler fails the test when a request reaches the server, used by cases that
// must fail during local validation.
func renewUnreachableHandler(t *testing.T) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Errorf("unexpected request reached the server: %s", r.URL.Path)
		renewWriteBadRequest(t, w)
	}
}

// assertRenewRequest checks the parts of the request every renew API builds the same way.
func assertRenewRequest(t *testing.T, r *nethttp.Request, path string) {
	t.Helper()
	if r.URL.Path != path {
		t.Fatalf("unexpected path: %s", r.URL.Path)
	}
	if r.Method != nethttp.MethodPost {
		t.Fatalf("unexpected method: %s", r.Method)
	}
	if got := r.Header.Get("x-Region"); got != testRegion {
		t.Fatalf("unexpected region header: %s", got)
	}
}

// renewFullCreateRequest / renewFullUpdateRequest / renewFullRenewClusterRequest /
// renewFullListRenewalsRequest return fully populated requests, so the required-field cases
// below only have to clear one field each.
func renewFullCreateRequest() *CreateAutoRenewRuleRequest {
	return &CreateAutoRenewRuleRequest{
		ClusterIds:    []string{renewTestClusterId},
		RenewTimeUnit: "month",
		RenewTime:     1,
		ServiceType:   "BES",
	}
}

func renewFullUpdateRequest() *UpdateAutoRenewRuleRequest {
	return &UpdateAutoRenewRuleRequest{
		ClusterId:     renewTestClusterId,
		RenewTimeUnit: "month",
		RenewTime:     1,
	}
}

func renewFullRenewClusterRequest() *RenewClusterRequest {
	return &RenewClusterRequest{ClusterId: renewTestClusterId, Time: 1}
}

func renewFullListRenewalsRequest() *ListRenewalsRequest {
	return &ListRenewalsRequest{
		Order:            "desc",
		OrderBy:          "expireTime",
		PageNo:           1,
		PageSize:         15,
		DaysToExpiration: 60,
	}
}

func TestCreateAutoRenewRule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertRenewRequest(t, r, "/api/bes/cluster/auto_renew_rule/create")
		body := readJSONBody(t, r)
		clusterIds, ok := body["clusterIds"].([]interface{})
		if !ok || len(clusterIds) != 1 || clusterIds[0] != renewTestClusterId {
			t.Fatalf("unexpected clusterIds: %v", body["clusterIds"])
		}
		if body["renewTimeUnit"] != "month" || body["renewTime"] != float64(1) {
			t.Fatalf("unexpected renew fields: %+v", body)
		}
		if body["serviceType"] != "BES" {
			t.Fatalf("unexpected serviceType: %v", body["serviceType"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"orderId": renewTestOrderId},
		})
	}))
	defer server.Close()

	result, err := CreateAutoRenewRule(cli, testRegion, renewFullCreateRequest())
	if err != nil {
		t.Fatalf("create auto renew rule failed: %v", err)
	}
	if result.Result == nil || result.Result.OrderId != renewTestOrderId {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestGetAutoRenewRuleDetail(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertRenewRequest(t, r, "/api/bes/cluster/auto_renew_rule/detail")
		if got := readJSONBody(t, r)["clusterId"]; got != renewTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"clusterId":     renewTestClusterId,
				"renewTimeUnit": "month",
				"renewTime":     1,
				"expireTime":    1658216735000,
				"nextRenewTime": 1658216735000,
			},
		})
	}))
	defer server.Close()

	result, err := GetAutoRenewRuleDetail(cli, testRegion, &GetAutoRenewRuleDetailRequest{
		ClusterId: renewTestClusterId,
	})
	if err != nil {
		t.Fatalf("get auto renew rule detail failed: %v", err)
	}
	if result.Result == nil || result.Result.ClusterId != renewTestClusterId {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if result.Result.RenewTimeUnit != "month" || result.Result.RenewTime != 1 {
		t.Fatalf("unexpected renew fields: %+v", result.Result)
	}
	if result.Result.ExpireTime != 1658216735000 || result.Result.NextRenewTime != 1658216735000 {
		t.Fatalf("unexpected time fields: %+v", result.Result)
	}
}

// TestListAutoRenewRules also fills the optional ClusterId filter, so its omitempty body field
// is exercised instead of being dropped.
func TestListAutoRenewRules(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertRenewRequest(t, r, "/api/bes/cluster/auto_renew_rule/list")
		body := readJSONBody(t, r)
		if body["serviceType"] != "BES" || body["clusterId"] != renewTestClusterId {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": []map[string]interface{}{
				{
					"uuid":          "rule-f38743879d554275a71d2d737362ab40",
					"userId":        "9485f297f8a64d8da79dde31791c08a6",
					"clusterId":     renewTestClusterId,
					"region":        "bd",
					"renewTimeUnit": "month",
					"renewTime":     1,
					"createTime":    "2021-10-27T06:51:18Z",
					"updateTime":    "1971-01-01T00:00:01Z",
				},
			},
		})
	}))
	defer server.Close()

	result, err := ListAutoRenewRules(cli, testRegion, &ListAutoRenewRulesRequest{
		ServiceType: "BES",
		ClusterId:   renewTestClusterId,
	})
	if err != nil {
		t.Fatalf("list auto renew rules failed: %v", err)
	}
	if len(result.Result) != 1 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	rule := result.Result[0]
	if rule.Uuid != "rule-f38743879d554275a71d2d737362ab40" || rule.Region != "bd" {
		t.Fatalf("unexpected rule: %+v", rule)
	}
	if rule.RenewTimeUnit != "month" || rule.RenewTime != 1 {
		t.Fatalf("unexpected renew fields: %+v", rule)
	}
}

func TestUpdateAutoRenewRule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertRenewRequest(t, r, "/api/bes/cluster/auto_renew_rule/update")
		body := readJSONBody(t, r)
		if body["clusterId"] != renewTestClusterId || body["renewTimeUnit"] != "month" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["renewTime"] != float64(1) {
			t.Fatalf("unexpected renewTime: %v", body["renewTime"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  "result",
		})
	}))
	defer server.Close()

	result, err := UpdateAutoRenewRule(cli, testRegion, renewFullUpdateRequest())
	if err != nil {
		t.Fatalf("update auto renew rule failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestDeleteAutoRenewRule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertRenewRequest(t, r, "/api/bes/cluster/auto_renew_rule/delete")
		if got := readJSONBody(t, r)["clusterId"]; got != renewTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"expireTime": 1658216735000},
		})
	}))
	defer server.Close()

	result, err := DeleteAutoRenewRule(cli, testRegion, &DeleteAutoRenewRuleRequest{
		ClusterId: renewTestClusterId,
	})
	if err != nil {
		t.Fatalf("delete auto renew rule failed: %v", err)
	}
	if result.Result == nil || result.Result.ExpireTime != 1658216735000 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

// TestRenewCluster is the only renew API that builds its request inline instead of going through
// createJSONRequest, so the orderType query parameter and the content type are asserted here.
func TestRenewCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertRenewRequest(t, r, "/api/bes/cluster/renew")
		if got := r.URL.Query().Get("orderType"); got != "RENEW" {
			t.Fatalf("unexpected orderType: %s (raw query %s)", got, r.URL.RawQuery)
		}
		if got := r.Header.Get("Content-Type"); got != bce.DEFAULT_CONTENT_TYPE {
			t.Fatalf("unexpected content type: %s", got)
		}
		body := readJSONBody(t, r)
		if body["clusterId"] != renewTestClusterId || body["time"] != float64(1) {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"orderId": renewTestOrderId},
		})
	}))
	defer server.Close()

	result, err := RenewCluster(cli, testRegion, renewFullRenewClusterRequest())
	if err != nil {
		t.Fatalf("renew cluster failed: %v", err)
	}
	if result.Result == nil || result.Result.OrderId != renewTestOrderId {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
}

// TestListRenewals asserts the page-shaped response: this API answers with a top-level "page"
// field and no "status", unlike every other renew API.
func TestListRenewals(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		assertRenewRequest(t, r, "/api/bes/cluster/renew/list")
		body := readJSONBody(t, r)
		if body["order"] != "desc" || body["orderBy"] != "expireTime" {
			t.Fatalf("unexpected order fields: %+v", body)
		}
		if body["pageNo"] != float64(1) || body["pageSize"] != float64(15) {
			t.Fatalf("unexpected page fields: %+v", body)
		}
		if body["daysToExpiration"] != float64(60) {
			t.Fatalf("unexpected daysToExpiration: %v", body["daysToExpiration"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"page": map[string]interface{}{
				"orderBy":    "expireTime",
				"order":      "desc",
				"pageNo":     1,
				"pageSize":   15,
				"totalCount": 1,
				"result": []map[string]interface{}{
					{
						"clusterName":   "postpaynew2",
						"clusterId":     renewTestClusterId,
						"region":        "bd",
						"expiredTime":   "2021-11-27T03:03:54Z",
						"clusterStatus": "RUNNING",
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := ListRenewals(cli, testRegion, renewFullListRenewalsRequest())
	if err != nil {
		t.Fatalf("list renewals failed: %v", err)
	}
	if !result.Success || result.Page == nil {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Page.TotalCount != 1 || result.Page.PageNo != 1 || result.Page.PageSize != 15 {
		t.Fatalf("unexpected page: %+v", result.Page)
	}
	if len(result.Page.Result) != 1 || result.Page.Result[0].ClusterStatus != "RUNNING" {
		t.Fatalf("unexpected page result: %+v", result.Page.Result)
	}
}

func TestRenewAPIRejectsNilClient(t *testing.T) {
	cases := []struct {
		name string
		call func() error
	}{
		{"CreateAutoRenewRule", func() error {
			_, err := CreateAutoRenewRule(nil, testRegion, renewFullCreateRequest())
			return err
		}},
		{"GetAutoRenewRuleDetail", func() error {
			_, err := GetAutoRenewRuleDetail(nil, testRegion, &GetAutoRenewRuleDetailRequest{
				ClusterId: renewTestClusterId,
			})
			return err
		}},
		{"ListAutoRenewRules", func() error {
			_, err := ListAutoRenewRules(nil, testRegion, &ListAutoRenewRulesRequest{ServiceType: "BES"})
			return err
		}},
		{"UpdateAutoRenewRule", func() error {
			_, err := UpdateAutoRenewRule(nil, testRegion, renewFullUpdateRequest())
			return err
		}},
		{"DeleteAutoRenewRule", func() error {
			_, err := DeleteAutoRenewRule(nil, testRegion, &DeleteAutoRenewRuleRequest{
				ClusterId: renewTestClusterId,
			})
			return err
		}},
		{"RenewCluster", func() error {
			_, err := RenewCluster(nil, testRegion, renewFullRenewClusterRequest())
			return err
		}},
		{"ListRenewals", func() error {
			_, err := ListRenewals(nil, testRegion, renewFullListRenewalsRequest())
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

func TestRenewAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, renewUnreachableHandler(t))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"CreateAutoRenewRule", func(c bce.Client) error {
			_, err := CreateAutoRenewRule(c, testRegion, nil)
			return err
		}},
		{"GetAutoRenewRuleDetail", func(c bce.Client) error {
			_, err := GetAutoRenewRuleDetail(c, testRegion, nil)
			return err
		}},
		{"ListAutoRenewRules", func(c bce.Client) error {
			_, err := ListAutoRenewRules(c, testRegion, nil)
			return err
		}},
		{"UpdateAutoRenewRule", func(c bce.Client) error {
			_, err := UpdateAutoRenewRule(c, testRegion, nil)
			return err
		}},
		{"DeleteAutoRenewRule", func(c bce.Client) error {
			_, err := DeleteAutoRenewRule(c, testRegion, nil)
			return err
		}},
		{"RenewCluster", func(c bce.Client) error {
			_, err := RenewCluster(c, testRegion, nil)
			return err
		}},
		{"ListRenewals", func(c bce.Client) error {
			_, err := ListRenewals(c, testRegion, nil)
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

// TestRenewAPIRejectsMissingRequiredFields covers every local validation branch in renew.go.
// DeleteAutoRenewRule is absent on purpose: its only field is optional, so it has no
// required-field branch beyond the nil-request guard.
func TestRenewAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, renewUnreachableHandler(t))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"CreateAutoRenewRule empty clusterIds", func(c bce.Client) error {
			request := renewFullCreateRequest()
			request.ClusterIds = nil
			_, err := CreateAutoRenewRule(c, testRegion, request)
			return err
		}},
		{"CreateAutoRenewRule empty renewTimeUnit", func(c bce.Client) error {
			request := renewFullCreateRequest()
			request.RenewTimeUnit = ""
			_, err := CreateAutoRenewRule(c, testRegion, request)
			return err
		}},
		{"CreateAutoRenewRule zero renewTime", func(c bce.Client) error {
			request := renewFullCreateRequest()
			request.RenewTime = 0
			_, err := CreateAutoRenewRule(c, testRegion, request)
			return err
		}},
		{"CreateAutoRenewRule empty serviceType", func(c bce.Client) error {
			request := renewFullCreateRequest()
			request.ServiceType = ""
			_, err := CreateAutoRenewRule(c, testRegion, request)
			return err
		}},
		{"GetAutoRenewRuleDetail empty clusterId", func(c bce.Client) error {
			_, err := GetAutoRenewRuleDetail(c, testRegion, &GetAutoRenewRuleDetailRequest{})
			return err
		}},
		{"ListAutoRenewRules empty serviceType", func(c bce.Client) error {
			_, err := ListAutoRenewRules(c, testRegion, &ListAutoRenewRulesRequest{
				ClusterId: renewTestClusterId,
			})
			return err
		}},
		{"UpdateAutoRenewRule empty clusterId", func(c bce.Client) error {
			request := renewFullUpdateRequest()
			request.ClusterId = ""
			_, err := UpdateAutoRenewRule(c, testRegion, request)
			return err
		}},
		{"UpdateAutoRenewRule empty renewTimeUnit", func(c bce.Client) error {
			request := renewFullUpdateRequest()
			request.RenewTimeUnit = ""
			_, err := UpdateAutoRenewRule(c, testRegion, request)
			return err
		}},
		{"UpdateAutoRenewRule zero renewTime", func(c bce.Client) error {
			request := renewFullUpdateRequest()
			request.RenewTime = 0
			_, err := UpdateAutoRenewRule(c, testRegion, request)
			return err
		}},
		{"RenewCluster empty clusterId", func(c bce.Client) error {
			request := renewFullRenewClusterRequest()
			request.ClusterId = ""
			_, err := RenewCluster(c, testRegion, request)
			return err
		}},
		{"RenewCluster zero time", func(c bce.Client) error {
			request := renewFullRenewClusterRequest()
			request.Time = 0
			_, err := RenewCluster(c, testRegion, request)
			return err
		}},
		{"ListRenewals empty order", func(c bce.Client) error {
			request := renewFullListRenewalsRequest()
			request.Order = ""
			_, err := ListRenewals(c, testRegion, request)
			return err
		}},
		{"ListRenewals empty orderBy", func(c bce.Client) error {
			request := renewFullListRenewalsRequest()
			request.OrderBy = ""
			_, err := ListRenewals(c, testRegion, request)
			return err
		}},
		{"ListRenewals zero pageNo", func(c bce.Client) error {
			request := renewFullListRenewalsRequest()
			request.PageNo = 0
			_, err := ListRenewals(c, testRegion, request)
			return err
		}},
		{"ListRenewals zero pageSize", func(c bce.Client) error {
			request := renewFullListRenewalsRequest()
			request.PageSize = 0
			_, err := ListRenewals(c, testRegion, request)
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

func TestRenewAPIPropagatesServerError(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		renewWriteBadRequest(t, w)
	}))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"CreateAutoRenewRule", func(c bce.Client) error {
			_, err := CreateAutoRenewRule(c, testRegion, renewFullCreateRequest())
			return err
		}},
		{"GetAutoRenewRuleDetail", func(c bce.Client) error {
			_, err := GetAutoRenewRuleDetail(c, testRegion, &GetAutoRenewRuleDetailRequest{
				ClusterId: renewTestClusterId,
			})
			return err
		}},
		{"ListAutoRenewRules", func(c bce.Client) error {
			_, err := ListAutoRenewRules(c, testRegion, &ListAutoRenewRulesRequest{ServiceType: "BES"})
			return err
		}},
		{"UpdateAutoRenewRule", func(c bce.Client) error {
			_, err := UpdateAutoRenewRule(c, testRegion, renewFullUpdateRequest())
			return err
		}},
		{"DeleteAutoRenewRule", func(c bce.Client) error {
			_, err := DeleteAutoRenewRule(c, testRegion, &DeleteAutoRenewRuleRequest{
				ClusterId: renewTestClusterId,
			})
			return err
		}},
		{"RenewCluster", func(c bce.Client) error {
			_, err := RenewCluster(c, testRegion, renewFullRenewClusterRequest())
			return err
		}},
		{"ListRenewals", func(c bce.Client) error {
			_, err := ListRenewals(c, testRegion, renewFullListRenewalsRequest())
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

// renewEmptyResultPayloads are the payloads isEmptyJSONResult treats as "no result": the BES v2
// server has been observed to answer result as an empty string where an object is documented.
var renewEmptyResultPayloads = map[string]string{
	"empty":        ``,
	"null":         `null`,
	"emptyString":  `""`,
	"paddedString": `  ""  `,
}

// TestOrderIdResultUnmarshalJSON exercises the three paths of the unmarshaller in renew_model.go:
// tolerated empty result, successful decode, and malformed payload.
func TestOrderIdResultUnmarshalJSON(t *testing.T) {
	for name, payload := range renewEmptyResultPayloads {
		var result OrderIdResult
		if err := result.UnmarshalJSON([]byte(payload)); err != nil {
			t.Fatalf("%s: unmarshal failed: %v", name, err)
		}
		if result.OrderId != "" {
			t.Fatalf("%s: unexpected orderId: %s", name, result.OrderId)
		}
	}

	var decoded OrderIdResult
	if err := decoded.UnmarshalJSON([]byte(`{"orderId":"` + renewTestOrderId + `"}`)); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.OrderId != renewTestOrderId {
		t.Fatalf("unexpected orderId: %s", decoded.OrderId)
	}

	var malformed OrderIdResult
	if err := malformed.UnmarshalJSON([]byte(`{"orderId":42}`)); err == nil {
		t.Fatal("expected a decoding error for a non-string orderId")
	}
}

// TestAutoRenewRuleDetailUnmarshalJSON covers the same three paths. The empty result is the real
// server answer when the cluster has no auto-renew rule at all.
func TestAutoRenewRuleDetailUnmarshalJSON(t *testing.T) {
	for name, payload := range renewEmptyResultPayloads {
		var detail AutoRenewRuleDetail
		if err := detail.UnmarshalJSON([]byte(payload)); err != nil {
			t.Fatalf("%s: unmarshal failed: %v", name, err)
		}
		if detail.ClusterId != "" || detail.RenewTime != 0 || detail.ExpireTime != 0 {
			t.Fatalf("%s: unexpected detail: %+v", name, detail)
		}
	}

	var decoded AutoRenewRuleDetail
	payload := `{"clusterId":"` + renewTestClusterId + `","renewTimeUnit":"month","renewTime":1,` +
		`"expireTime":1658216735000,"nextRenewTime":1658216735000}`
	if err := decoded.UnmarshalJSON([]byte(payload)); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.ClusterId != renewTestClusterId || decoded.RenewTimeUnit != "month" {
		t.Fatalf("unexpected detail: %+v", decoded)
	}
	if decoded.RenewTime != 1 || decoded.ExpireTime != 1658216735000 {
		t.Fatalf("unexpected detail numbers: %+v", decoded)
	}

	var malformed AutoRenewRuleDetail
	if err := malformed.UnmarshalJSON([]byte(`{"renewTime":"month"}`)); err == nil {
		t.Fatal("expected a decoding error for a non-numeric renewTime")
	}
}

// TestExpireTimeResultUnmarshalJSON covers the same three paths.
func TestExpireTimeResultUnmarshalJSON(t *testing.T) {
	for name, payload := range renewEmptyResultPayloads {
		var result ExpireTimeResult
		if err := result.UnmarshalJSON([]byte(payload)); err != nil {
			t.Fatalf("%s: unmarshal failed: %v", name, err)
		}
		if result.ExpireTime != 0 {
			t.Fatalf("%s: unexpected expireTime: %v", name, result.ExpireTime)
		}
	}

	var decoded ExpireTimeResult
	if err := decoded.UnmarshalJSON([]byte(`{"expireTime":1658216735000}`)); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if decoded.ExpireTime != 1658216735000 {
		t.Fatalf("unexpected expireTime: %v", decoded.ExpireTime)
	}

	var malformed ExpireTimeResult
	if err := malformed.UnmarshalJSON([]byte(`{"expireTime":"soon"}`)); err == nil {
		t.Fatal("expected a decoding error for a non-numeric expireTime")
	}
}

// TestGetAutoRenewRuleDetailWithEmptyStringResult is the end-to-end counterpart of the
// unmarshaller tests: a successful response whose result is an empty string must decode to a zero
// value rather than failing the whole call.
func TestGetAutoRenewRuleDetailWithEmptyStringResult(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": ""})
	}))
	defer server.Close()

	result, err := GetAutoRenewRuleDetail(cli, testRegion, &GetAutoRenewRuleDetailRequest{
		ClusterId: renewTestClusterId,
	})
	if err != nil {
		t.Fatalf("get auto renew rule detail failed: %v", err)
	}
	if result.Result == nil || result.Result.ClusterId != "" || result.Result.ExpireTime != 0 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

// TestRenewClusterRejectsMalformedResult pins the other half of the empty-string tolerance: the
// guard in isEmptyJSONResult only covers the empty string, it must not swallow a genuine decode
// failure. Migrated from the MalformedResultStillFails subtest of TestEmptyStringResultIsTolerated
// in services/bes/v2/client_test.go, which asserted an api-layer contract from the v2 package.
func TestRenewClusterRejectsMalformedResult(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"orderId": 12345},
		})
	}))
	defer server.Close()

	if _, err := RenewCluster(cli, testRegion, renewFullRenewClusterRequest()); err == nil {
		t.Fatal("expected an error for an orderId of the wrong type")
	}
}

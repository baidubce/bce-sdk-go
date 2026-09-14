package api

import (
	"errors"
	"io"
	nethttp "net/http"
	"strings"
	"testing"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	bcehttp "github.com/baidubce/bce-sdk-go/http"
)

// clusterBatchCountPtr returns a pointer to a restart batch count. RestartClusterRequest.BatchCount
// is a *float64, so the shared intPtr helper does not fit.
func clusterBatchCountPtr(v float64) *float64 { return &v }

// newClusterRejectingClient returns a client backed by a server that fails the test if it is ever
// reached. Every validation case below must be rejected before any HTTP call, so this keeps each of
// them to a single line of setup.
func newClusterRejectingClient(t *testing.T) bce.Client {
	t.Helper()
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	t.Cleanup(server.Close)
	return cli
}

// newClusterOfflineClient builds a client with the given credentials pointed at a closed port, with
// retries disabled. It covers the failures that happen before or instead of a successful HTTP
// exchange: missing credentials, an unusable secret key, and a transport-level send failure.
func newClusterOfflineClient(t *testing.T, credentials *auth.BceCredentials) bce.Client {
	t.Helper()
	return bce.NewBceClient(&bce.BceClientConfiguration{
		Endpoint:    "http://127.0.0.1:1",
		Region:      testRegion,
		UserAgent:   bce.DEFAULT_USER_AGENT,
		Credentials: credentials,
		SignOption: &auth.SignOptions{
			HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
			ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS,
		},
		Retry:                     bce.NewNoRetryPolicy(),
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
	}, &auth.BceV1Signer{})
}

// TestClusterHelperSetRegion covers both branches of setRegion, which every cluster function uses to
// let a per-request region override the client region.
func TestClusterHelperSetRegion(t *testing.T) {
	if got := setRegion("bj", "sh"); got != "sh" {
		t.Fatalf("request region should win, got %s", got)
	}
	if got := setRegion("bj", ""); got != "bj" {
		t.Fatalf("client region should be the fallback, got %s", got)
	}
}

// TestClusterHelperIntString covers the non-positive branch of intString, which callers never reach
// because they all guard the call with a `> 0` check.
func TestClusterHelperIntString(t *testing.T) {
	if got := intString(7); got != "7" {
		t.Fatalf("unexpected value: %s", got)
	}
	for _, v := range []int{0, -1} {
		if got := intString(v); got != "" {
			t.Fatalf("intString(%d) should be empty, got %s", v, got)
		}
	}
}

// TestClusterHelperClusterURI covers the zero-parts branch of clusterURI, plus path escaping of the
// segments, which no production caller exercises today.
func TestClusterHelperClusterURI(t *testing.T) {
	if got := clusterURI(); got != URI_CLUSTERS {
		t.Fatalf("unexpected collection uri: %s", got)
	}
	if got := clusterURI("search-xxxx"); got != "/v3/clusters/search-xxxx" {
		t.Fatalf("unexpected uri: %s", got)
	}
	if got := clusterURI("search-xxxx", "nodes", "start"); got != "/v3/clusters/search-xxxx/nodes/start" {
		t.Fatalf("unexpected uri: %s", got)
	}
	if got := clusterURI("a b/c"); got != "/v3/clusters/a%20b%2Fc" {
		t.Fatalf("path segment should be escaped, got %s", got)
	}
}

// TestClusterHelperCredentialsOf covers both branches of credentialsOf: a configured client, and a
// client whose configuration carries no credentials.
func TestClusterHelperCredentialsOf(t *testing.T) {
	credentials, err := credentialsOf(newClusterRejectingClient(t))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if credentials.AccessKeyId != testAK || credentials.SecretAccessKey != testSK {
		t.Fatalf("unexpected credentials: %+v", credentials)
	}
	if _, err := credentialsOf(newClusterOfflineClient(t, nil)); err == nil {
		t.Fatal("expected error when credentials are not configured")
	}
}

// TestCreateCluster populates every optional body field as well, and asserts that the admin password
// is encrypted with the client secret key, that the matching access key is sent in the header, and
// that the caller's own request struct is left untouched.
func TestCreateCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Content-Type"); got != "application/json;charset=utf-8" {
			t.Fatalf("unexpected content-type: %s", got)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		if got := r.Header.Get("X-Bce-Accesskey"); got != testAK {
			t.Fatalf("unexpected X-Bce-Accesskey: %s", got)
		}

		body := readJSONBody(t, r)
		for key, want := range map[string]interface{}{
			"payment":            "Postpaid",
			"name":               "bes-test",
			"engine":             "OPENSEARCH",
			"engineType":         "COMMUNITY",
			"version":            "2.11.0",
			"subnetId":           "sbn-xxxx",
			"vpcId":              "vpc-xxxx",
			"enableHttps":        true,
			"deletionProtection": true,
			"enableDeploySet":    true,
		} {
			if got := body[key]; got != want {
				t.Fatalf("unexpected %s: got %v want %v", key, got, want)
			}
		}
		if got := body["logicalZones"]; got == nil {
			t.Fatalf("logicalZones should not be nil")
		}
		if paymentInfo, ok := body["paymentInfo"].(map[string]interface{}); !ok || paymentInfo["timeUnit"] != "month" {
			t.Fatalf("unexpected paymentInfo: %v", body["paymentInfo"])
		}
		if couponIds, ok := body["couponIds"].([]interface{}); !ok || len(couponIds) != 1 {
			t.Fatalf("unexpected couponIds: %v", body["couponIds"])
		}
		if tags, ok := body["tags"].([]interface{}); !ok || len(tags) != 1 {
			t.Fatalf("unexpected tags: %v", body["tags"])
		}
		if got, _ := body["adminPassword"].(string); got == "" || got == testPassword {
			t.Fatalf("adminPassword should be encrypted, not sent as plaintext: %v", got)
		} else if decryptAES128WithFirst16Char(t, got, testSK) != testPassword {
			t.Fatalf("adminPassword did not decrypt to the original plaintext: %v", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"orderId":   "order-xxxx",
			"actionId":  "action-xxxx",
		})
	}))
	defer server.Close()

	request := &CreateClusterRequest{
		Request:            Request{Region: "sh"},
		Payment:            PAYMENT_POSTPAID,
		PaymentInfo:        map[string]interface{}{"timeUnit": "month", "timeLength": 1},
		CouponIds:          []string{"coupon-xxxx"},
		Name:               "bes-test",
		Engine:             ENGINE_OPENSEARCH,
		EngineType:         ENGINE_TYPE_COMMUNITY,
		Version:            "2.11.0",
		AdminPassword:      testPassword,
		EnableHttps:        boolPtr(true),
		DeletionProtection: boolPtr(true),
		EnableDeploySet:    boolPtr(true),
		LogicalZones:       []string{"cn-bj-a"},
		VpcId:              "vpc-xxxx",
		SubnetId:           "sbn-xxxx",
		NodeSpecs: []NodeSpec{{
			Type:      NODE_SPEC_TYPE_DATA,
			NodeType:  "opensearch.c2m8",
			NodeCount: intPtr(3),
			DiskType:  "enhanced_ssd_pl1",
			DiskSize:  intPtr(50),
			DiskCount: intPtr(1),
		}},
		Tags: []Tag{{TagKey: "env", TagValue: "test"}},
	}

	result, err := CreateCluster(cli, testRegion, request)
	if err != nil {
		t.Fatalf("create cluster failed: %v", err)
	}
	if result.ClusterId != "search-xxxx" || result.OrderId != "order-xxxx" || result.ActionId != "action-xxxx" {
		t.Fatalf("unexpected result: %+v", result)
	}
	if request.AdminPassword != testPassword {
		t.Fatalf("caller request should keep the plaintext password, got %s", request.AdminPassword)
	}
}

// TestListClusters populates every optional filter so each conditional query-parameter branch in
// ListClusters is exercised.
func TestListClusters(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		query := r.URL.Query()
		expected := map[string]string{
			"pageNo":        "2",
			"pageSize":      "20",
			"order":         "desc",
			"orderBy":       "createTime",
			"payment":       PAYMENT_POSTPAID,
			"engine":        ENGINE_OPENSEARCH,
			"status":        "ACTIVE",
			"clusterHealth": "GREEN",
			"logicalZone":   "cn-bj-a",
			"clusterName":   "test-cluster",
			"clusterId":     "search-xxxx",
			"tagKey":        "env",
			"tagValue":      "test",
		}
		for key, want := range expected {
			if got := query.Get(key); got != want {
				t.Fatalf("unexpected %s: got %s want %s", key, got, want)
			}
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"clusters": []map[string]interface{}{{
				"clusterId": "search-xxxx",
				"name":      "test-cluster-1",
				"status":    "ACTIVE",
				"region":    "bj",
			}},
			"pageNo":     2,
			"pageSize":   20,
			"totalCount": 1,
		})
	}))
	defer server.Close()

	result, err := ListClusters(cli, testRegion, &ListClustersRequest{
		Request:       Request{Region: "sh"},
		PageNo:        2,
		PageSize:      20,
		Order:         "desc",
		OrderBy:       "createTime",
		Payment:       PAYMENT_POSTPAID,
		Engine:        ENGINE_OPENSEARCH,
		Status:        "ACTIVE",
		ClusterHealth: "GREEN",
		LogicalZone:   "cn-bj-a",
		ClusterName:   "test-cluster",
		ClusterId:     "search-xxxx",
		TagKey:        "env",
		TagValue:      "test",
	})
	if err != nil {
		t.Fatalf("list clusters failed: %v", err)
	}
	if result.PageNo != 2 || result.PageSize != 20 || result.TotalCount != 1 {
		t.Fatalf("unexpected paging result: %+v", result)
	}
	if len(result.Clusters) != 1 || result.Clusters[0].ClusterId != "search-xxxx" {
		t.Fatalf("unexpected cluster list: %+v", result)
	}
}

// TestListClustersWithoutOptionalFields covers the request path where every optional filter is
// omitted from the query string entirely.
func TestListClustersWithoutOptionalFields(t *testing.T) {
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

	result, err := ListClusters(cli, testRegion, &ListClustersRequest{})
	if err != nil {
		t.Fatalf("list clusters failed: %v", err)
	}
	if result.TotalCount != 0 || len(result.Clusters) != 0 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId":     "search-xxxx",
			"name":          "test-cluster-1",
			"engine":        "OPENSEARCH",
			"engineType":    "COMMUNITY",
			"version":       "2.19.4",
			"kernelVersion": "1.0.3",
			"status":        "ACTIVE",
		})
	}))
	defer server.Close()

	result, err := GetCluster(cli, testRegion, &GetClusterRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("get cluster failed: %v", err)
	}
	if result.ClusterId != "search-xxxx" || result.Name != "test-cluster-1" {
		t.Fatalf("unexpected detail result: %+v", result)
	}
}

func TestListClusterNodes(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/nodes" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"nodeSpecs": []map[string]interface{}{{
				"type":      "data",
				"nodeType":  "opensearch.c2m8",
				"cpu":       2,
				"memory":    8,
				"nodeCount": 3,
			}},
			"nodes": []map[string]interface{}{{
				"nodeId":   "i-FAn0oxI5",
				"nodeName": "search-xxxx-data-cn-bj-a-VmL8yp1q-3",
				"type":     "data",
				"status":   "ALIVE",
			}},
		})
	}))
	defer server.Close()

	result, err := ListClusterNodes(cli, testRegion, &ListClusterNodesRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("list cluster nodes failed: %v", err)
	}
	if len(result.NodeSpecs) != 1 || len(result.Nodes) != 1 {
		t.Fatalf("unexpected node list result: %+v", result)
	}
}

func TestQueryAvailableSpecs(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/specs" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		for key, want := range map[string]interface{}{
			"engine":     "OPENSEARCH",
			"engineType": "COMMUNITY",
			"version":    "2.19.4",
		} {
			if got := body[key]; got != want {
				t.Fatalf("unexpected %s: got %v want %v", key, got, want)
			}
		}
		if features, ok := body["features"].([]interface{}); !ok || len(features) != 2 {
			t.Fatalf("unexpected features: %v", body["features"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"nodeAvailableSpecs": []map[string]interface{}{{
				"type":        "data",
				"nodeMinSize": 2,
				"nodeMaxSize": 200,
				"diskMinSize": map[string]int{"enhanced_ssd_pl1": 50},
			}},
			"nodeSpecs": []map[string]interface{}{{"nodeType": "opensearch.c2m8", "cpu": 2, "memory": 8}},
			"diskSpecs": []map[string]interface{}{{"diskType": "enhanced_ssd_pl1"}},
		})
	}))
	defer server.Close()

	result, err := QueryAvailableSpecs(cli, testRegion, &QueryAvailableSpecsRequest{
		Request:      Request{Region: "sh"},
		Engine:       ENGINE_OPENSEARCH,
		EngineType:   ENGINE_TYPE_COMMUNITY,
		Version:      "2.19.4",
		Features:     []string{"a", "b"},
		LogicalZones: []string{"cn-bj-a"},
	})
	if err != nil {
		t.Fatalf("query available specs failed: %v", err)
	}
	if len(result.NodeAvailableSpecs) != 1 || len(result.NodeSpecs) != 1 || len(result.DiskSpecs) != 1 {
		t.Fatalf("unexpected spec result: %+v", result)
	}
}

func TestQueryAvailableKernels(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/kernels" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"versions": []map[string]interface{}{{
				"engine":     "OPENSEARCH",
				"engineType": "COMMUNITY",
				"version":    "1.0.3",
				"status":     "RECOMMENDED",
				"features":   []string{"a"},
				"scenes":     []string{"1"},
			}},
		})
	}))
	defer server.Close()

	result, err := QueryAvailableKernels(cli, testRegion, &QueryAvailableKernelsRequest{
		Request: Request{Region: "sh"},
	})
	if err != nil {
		t.Fatalf("query available kernels failed: %v", err)
	}
	if len(result.Versions) != 1 || result.Versions[0].Version != "1.0.3" {
		t.Fatalf("unexpected kernels result: %+v", result)
	}
}

// TestClusterAPIAcceptsNilRequest covers the two cluster functions that treat the request as
// optional: they only read a region override from it, so a nil request must succeed with the client
// region rather than return a validation error.
func TestClusterAPIAcceptsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if got := r.Header.Get("X-Region"); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}
		switch r.URL.Path {
		case "/v3/clusters/specs":
			writeJSONResponse(t, w, map[string]interface{}{
				"nodeSpecs": []map[string]interface{}{{"nodeType": "opensearch.c2m8"}},
			})
		case "/v3/clusters/kernels":
			writeJSONResponse(t, w, map[string]interface{}{
				"versions": []map[string]interface{}{{"version": "1.0.3"}},
			})
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	specs, err := QueryAvailableSpecs(cli, testRegion, nil)
	if err != nil {
		t.Fatalf("query available specs with nil request failed: %v", err)
	}
	if len(specs.NodeSpecs) != 1 {
		t.Fatalf("unexpected spec result: %+v", specs)
	}

	kernels, err := QueryAvailableKernels(cli, testRegion, nil)
	if err != nil {
		t.Fatalf("query available kernels with nil request failed: %v", err)
	}
	if len(kernels.Versions) != 1 {
		t.Fatalf("unexpected kernels result: %+v", kernels)
	}
}

// TestDownloadClusterCert asserts the zip stream and every response header the API surfaces. The
// returned Body must be closed by the caller.
func TestDownloadClusterCert(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/certs" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("Accept"); got != CONTENT_TYPE_ZIP {
			t.Fatalf("unexpected accept header: %s", got)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", `attachment; filename="bes-search-xxxx-https-certs-v0001.zip"`)
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, private")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Write([]byte("fake zip bytes"))
	}))
	defer server.Close()

	result, err := DownloadClusterCert(cli, testRegion, &DownloadClusterCertRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxx",
	})
	if err != nil {
		t.Fatalf("download cluster cert failed: %v", err)
	}
	defer result.Body.Close()

	content, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatalf("read cert body failed: %v", err)
	}
	if string(content) != "fake zip bytes" {
		t.Fatalf("unexpected cert content: %s", content)
	}
	if result.ContentType != "application/zip" {
		t.Fatalf("unexpected content type: %s", result.ContentType)
	}
	if result.ContentDisposition != `attachment; filename="bes-search-xxxx-https-certs-v0001.zip"` {
		t.Fatalf("unexpected content disposition: %s", result.ContentDisposition)
	}
	if result.CacheControl != "no-store, no-cache, must-revalidate, private" {
		t.Fatalf("unexpected cache control: %s", result.CacheControl)
	}
	if result.Pragma != "no-cache" || result.Expires != "0" {
		t.Fatalf("unexpected pragma/expires: %s / %s", result.Pragma, result.Expires)
	}
	if result.XContentTypeOptions != "nosniff" {
		t.Fatalf("unexpected x-content-type-options: %s", result.XContentTypeOptions)
	}
	if result.ContentLength != int64(len("fake zip bytes")) {
		t.Fatalf("unexpected content length: %d", result.ContentLength)
	}
}

// TestDownloadClusterCertPropagatesErrors covers the two failure paths of the streaming download,
// which does not go through the shared JSON request helper: a transport failure from SendRequest and
// a service error carried by a failed response. 400 is used rather than 500 because the test client
// mirrors the production retry policy, which retries 5xx with backoff.
func TestDownloadClusterCertPropagatesErrors(t *testing.T) {
	credentials, err := auth.NewBceCredentials(testAK, testSK)
	if err != nil {
		t.Fatalf("new credentials failed: %v", err)
	}
	offline := newClusterOfflineClient(t, credentials)
	if _, err := DownloadClusterCert(offline, testRegion, &DownloadClusterCertRequest{
		ClusterId: "search-xxxx",
	}); err == nil {
		t.Fatal("expected transport error when the endpoint is unreachable")
	}

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

	if _, err := DownloadClusterCert(cli, testRegion, &DownloadClusterCertRequest{
		ClusterId: "search-xxxx",
	}); err == nil {
		t.Fatal("expected error for 400 response")
	}
}

// clusterFailingResponseClient reports a failed response without returning an error from SendRequest.
// *bce.BceClient turns a failing status into an error itself, so a stub client is the only way to reach
// the resp.IsFail() guard that DownloadClusterCert keeps for other bce.Client implementations. The
// embedded nil interface supplies the rest of bce.Client; none of those methods is reached here.
type clusterFailingResponseClient struct {
	bce.Client
}

func (clusterFailingResponseClient) SendRequest(req *bce.BceRequest, resp *bce.BceResponse) error {
	httpResponse := &bcehttp.Response{}
	httpResponse.SetHttpResponse(&nethttp.Response{
		StatusCode: nethttp.StatusForbidden,
		Status:     "403 Forbidden",
		Header:     nethttp.Header{},
		Body:       io.NopCloser(strings.NewReader(`{"code":"AccessDenied","message":"denied"}`)),
	})
	resp.SetHttpResponse(httpResponse)
	resp.ParseResponse()
	return nil
}

// TestDownloadClusterCertRejectsFailedResponse covers the resp.IsFail() branch of the streaming
// download, which is unreachable through *bce.BceClient.
func TestDownloadClusterCertRejectsFailedResponse(t *testing.T) {
	_, err := DownloadClusterCert(clusterFailingResponseClient{}, testRegion, &DownloadClusterCertRequest{
		ClusterId: "search-xxxx",
	})
	if err == nil {
		t.Fatal("expected the service error of the failed response")
	}
}

func TestUpdateClusterName(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/name" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if got := body["newName"]; got != "bes-new-name" {
			t.Fatalf("unexpected newName: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"name":      "bes-new-name",
		})
	}))
	defer server.Close()

	result, err := UpdateClusterName(cli, testRegion, &UpdateClusterNameRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxx",
		NewName:   "bes-new-name",
	})
	if err != nil {
		t.Fatalf("update cluster name failed: %v", err)
	}
	if result.ClusterId != "search-xxxx" || result.Name != "bes-new-name" {
		t.Fatalf("unexpected updated name result: %+v", result)
	}
}

// TestUpdateClusterMaintenance populates every optional body field of the maintenance window.
func TestUpdateClusterMaintenance(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/maintenance-duration" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		for key, want := range map[string]interface{}{
			"maintenanceTimeZone":  "Asia/Shanghai",
			"maintenanceStartTime": "02:00",
			"maintenanceEndTime":   "04:00",
		} {
			if got := body[key]; got != want {
				t.Fatalf("unexpected %s: got %v want %v", key, got, want)
			}
		}
		if periods, ok := body["maintenancePeriods"].([]interface{}); !ok || len(periods) != 2 {
			t.Fatalf("unexpected maintenancePeriods: %v", body["maintenancePeriods"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := UpdateClusterMaintenance(cli, testRegion, &UpdateClusterMaintenanceRequest{
		ClusterId:            "search-xxxx",
		MaintenanceTimeZone:  "Asia/Shanghai",
		MaintenancePeriods:   []string{"MONDAY", "FRIDAY"},
		MaintenanceStartTime: "02:00",
		MaintenanceEndTime:   "04:00",
	})
	if err != nil {
		t.Fatalf("update cluster maintenance failed: %v", err)
	}
	if result.Success == nil || !*result.Success {
		t.Fatalf("unexpected maintenance update result: %+v", result)
	}
}

func TestUpdateClusterDeletionProtection(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/deletion-protection" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if got := body["deletionProtection"]; got != true {
			t.Fatalf("unexpected deletionProtection: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId":          "search-xxxx",
			"deletionProtection": true,
		})
	}))
	defer server.Close()

	result, err := UpdateClusterDeletionProtection(cli, testRegion, &UpdateClusterDeletionProtectionRequest{
		ClusterId:          "search-xxxx",
		DeletionProtection: boolPtr(true),
	})
	if err != nil {
		t.Fatalf("update cluster deletion protection failed: %v", err)
	}
	if result.DeletionProtection == nil || !*result.DeletionProtection {
		t.Fatalf("unexpected deletion protection result: %+v", result)
	}
}

func TestRecoverCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/recovery" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-xxxx",
		})
	}))
	defer server.Close()

	result, err := RecoverCluster(cli, testRegion, &RecoverClusterRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("recover cluster failed: %v", err)
	}
	if result.ClusterId != "search-xxxx" || result.ActionId != "action-xxxx" {
		t.Fatalf("unexpected recover result: %+v", result)
	}
}

func TestDeleteCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("recycle"); got != "true" {
			t.Fatalf("unexpected recycle: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"clusterId": "search-xxxx"})
	}))
	defer server.Close()

	result, err := DeleteCluster(cli, testRegion, &DeleteClusterRequest{
		ClusterId: "search-xxxx",
		Recycle:   boolPtr(true),
	})
	if err != nil {
		t.Fatalf("delete cluster failed: %v", err)
	}
	if result.ClusterId != "search-xxxx" {
		t.Fatalf("unexpected delete result: %+v", result)
	}
}

// TestDeleteClusterWithoutOptionalFields covers the path where the optional recycle flag is left
// unset and no query parameter is sent at all. A false flag still has to be sent, so the two cases
// are asserted together.
func TestDeleteClusterWithoutOptionalFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		writeJSONResponse(t, w, map[string]interface{}{"clusterId": r.URL.RawQuery})
	}))
	defer server.Close()

	result, err := DeleteCluster(cli, testRegion, &DeleteClusterRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("delete cluster failed: %v", err)
	}
	if result.ClusterId != "" {
		t.Fatalf("unexpected query for an unset recycle flag: %s", result.ClusterId)
	}

	result, err = DeleteCluster(cli, testRegion, &DeleteClusterRequest{
		ClusterId: "search-xxxx",
		Recycle:   boolPtr(false),
	})
	if err != nil {
		t.Fatalf("delete cluster failed: %v", err)
	}
	if result.ClusterId != "recycle=false" {
		t.Fatalf("unexpected query for recycle=false: %s", result.ClusterId)
	}
}

// TestResizeCluster populates every optional body field of a scale-out resize.
func TestResizeCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		for key, want := range map[string]interface{}{
			"mode":            "DIRECT",
			"vpcId":           "vpc-xxxx",
			"subnetId":        "sbn-xxxx",
			"enableDeploySet": true,
		} {
			if got := body[key]; got != want {
				t.Fatalf("unexpected %s: got %v want %v", key, got, want)
			}
		}
		if zones, ok := body["logicalZones"].([]interface{}); !ok || len(zones) != 1 {
			t.Fatalf("unexpected logicalZones: %v", body["logicalZones"])
		}
		if specs, ok := body["nodeSpecs"].([]interface{}); !ok || len(specs) != 1 {
			t.Fatalf("unexpected nodeSpecs: %v", body["nodeSpecs"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"orderId":   "order-resize-example",
		})
	}))
	defer server.Close()

	result, err := ResizeCluster(cli, testRegion, &ResizeClusterRequest{
		ClusterId:       "search-xxxx",
		Mode:            RESIZE_MODE_DIRECT,
		LogicalZones:    []string{"cn-bj-a"},
		VpcId:           "vpc-xxxx",
		SubnetId:        "sbn-xxxx",
		EnableDeploySet: boolPtr(true),
		NodeSpecs: []NodeSpec{{
			Type:      NODE_SPEC_TYPE_COORDINATING,
			NodeType:  "opensearch.c2m8",
			NodeCount: intPtr(2),
		}},
	})
	if err != nil {
		t.Fatalf("resize cluster failed: %v", err)
	}
	if result.ClusterId != "search-xxxx" || result.OrderId != "order-resize-example" {
		t.Fatalf("unexpected resize result: %+v", result)
	}
}

// TestResizeClusterDeletingNodeTypes covers the other side of the mutually exclusive resize inputs:
// deleteNodeTypes set with no nodeSpecs.
func TestResizeClusterDeletingNodeTypes(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		body := readJSONBody(t, r)
		if body["nodeSpecs"] != nil {
			t.Fatalf("nodeSpecs should be omitted: %v", body["nodeSpecs"])
		}
		types, ok := body["deleteNodeTypes"].([]interface{})
		if !ok || len(types) != 1 || types[0] != "coordinating" {
			t.Fatalf("unexpected deleteNodeTypes: %v", body["deleteNodeTypes"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"clusterId": "search-xxxx"})
	}))
	defer server.Close()

	result, err := ResizeCluster(cli, testRegion, &ResizeClusterRequest{
		ClusterId:       "search-xxxx",
		Mode:            RESIZE_MODE_DIRECT,
		DeleteNodeTypes: []string{NODE_SPEC_TYPE_COORDINATING},
	})
	if err != nil {
		t.Fatalf("resize cluster failed: %v", err)
	}
	if result.ClusterId != "search-xxxx" {
		t.Fatalf("unexpected resize result: %+v", result)
	}
}

func TestListSuggestedMigrationNodes(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/migrations/suggested-nodes" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("type"); got != "data" {
			t.Fatalf("unexpected type: %s", got)
		}
		if got := r.URL.Query().Get("migrateCount"); got != "2" {
			t.Fatalf("unexpected migrateCount: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"nodes": []map[string]interface{}{{
				"nodeId":      "i-xxxx1",
				"nodeName":    "node-xxxx1",
				"status":      "ALIVE",
				"logicalZone": "cn-bj-a",
			}},
		})
	}))
	defer server.Close()

	result, err := ListSuggestedMigrationNodes(cli, testRegion, &ListSuggestedMigrationNodesRequest{
		ClusterId:    "search-xxxx",
		Type:         NODE_SPEC_TYPE_DATA,
		MigrateCount: 2,
	})
	if err != nil {
		t.Fatalf("list suggested migration nodes failed: %v", err)
	}
	if len(result.Nodes) != 1 || result.Nodes[0].NodeId != "i-xxxx1" {
		t.Fatalf("unexpected suggested migration nodes result: %+v", result)
	}
}

func TestListMigratableNodes(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/migrations/nodes" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("type"); got != "data" {
			t.Fatalf("unexpected type: %s", got)
		}
		if got := r.URL.Query().Get("migrateCount"); got != "" {
			t.Fatalf("migrateCount should not be present, got: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"nodes": []map[string]interface{}{{"nodeId": "i-xxxx1", "status": "ALIVE"}},
		})
	}))
	defer server.Close()

	result, err := ListMigratableNodes(cli, testRegion, &ListMigratableNodesRequest{
		ClusterId: "search-xxxx",
		Type:      NODE_SPEC_TYPE_DATA,
	})
	if err != nil {
		t.Fatalf("list migratable nodes failed: %v", err)
	}
	if len(result.Nodes) != 1 || result.Nodes[0].NodeId != "i-xxxx1" {
		t.Fatalf("unexpected migratable nodes result: %+v", result)
	}
}

func TestQueryAvailableUpgrades(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/upgrades/versions" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId":       "search-xxxx",
			"currentVersion":  "2.19.4",
			"currentKernel":   "1.0.3",
			"versionUpgrades": []map[string]interface{}{{"version": "2.19.5", "upgradeModes": []string{"ROLLING"}}},
			"kernelUpgrades":  []map[string]interface{}{{"kernel": "1.0.4", "upgradeModes": []string{"ROLLING"}}},
		})
	}))
	defer server.Close()

	result, err := QueryAvailableUpgrades(cli, testRegion, &QueryAvailableUpgradesRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("query available upgrades failed: %v", err)
	}
	if len(result.VersionUpgrades) != 1 || len(result.KernelUpgrades) != 1 {
		t.Fatalf("unexpected upgrades result: %+v", result)
	}
}

// TestUpdateClusterPublicAccess populates every optional body field of the public access switch.
func TestUpdateClusterPublicAccess(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/public-access" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		for key, want := range map[string]interface{}{
			"enabled":    true,
			"accessType": "OPENSEARCH",
			"publicIp":   "10.0.0.1",
		} {
			if got := body[key]; got != want {
				t.Fatalf("unexpected %s: got %v want %v", key, got, want)
			}
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-public-access-example",
		})
	}))
	defer server.Close()

	result, err := UpdateClusterPublicAccess(cli, testRegion, &UpdateClusterPublicAccessRequest{
		ClusterId:  "search-xxxx",
		Enabled:    boolPtr(true),
		PublicIp:   "10.0.0.1",
		AccessType: ACCESS_TYPE_OPENSEARCH,
	})
	if err != nil {
		t.Fatalf("update cluster public access failed: %v", err)
	}
	if result.ActionId != "action-public-access-example" {
		t.Fatalf("unexpected public access result: %+v", result)
	}
}

func TestUpdateClusterCerebro(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/cerebro" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if got := body["enabled"]; got != true {
			t.Fatalf("unexpected enabled: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"enabled":   true,
		})
	}))
	defer server.Close()

	result, err := UpdateClusterCerebro(cli, testRegion, &UpdateClusterCerebroRequest{
		ClusterId: "search-xxxx",
		Enabled:   boolPtr(true),
	})
	if err != nil {
		t.Fatalf("update cluster cerebro failed: %v", err)
	}
	if result.ClusterId != "search-xxxx" || result.Enabled == nil || !*result.Enabled {
		t.Fatalf("unexpected cerebro result: %+v", result)
	}
}

func TestStartCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/start" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-start-cluster-example",
		})
	}))
	defer server.Close()

	result, err := StartCluster(cli, testRegion, &StartClusterRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("start cluster failed: %v", err)
	}
	if result.ActionId != "action-start-cluster-example" {
		t.Fatalf("unexpected start cluster result: %+v", result)
	}
}

func TestStartClusterNodes(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/nodes/start" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if nodes, ok := body["nodes"].([]interface{}); !ok || len(nodes) != 2 {
			t.Fatalf("unexpected nodes: %v", body["nodes"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-start-nodes-example",
		})
	}))
	defer server.Close()

	result, err := StartClusterNodes(cli, testRegion, &ClusterNodesRequest{
		ClusterId: "search-xxxx",
		Nodes:     []string{"i-xxxx1", "i-xxxx2"},
	})
	if err != nil {
		t.Fatalf("start cluster nodes failed: %v", err)
	}
	if result.ActionId != "action-start-nodes-example" {
		t.Fatalf("unexpected start nodes result: %+v", result)
	}
}

func TestMigrateClusterNodeData(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/migrations" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if got := body["type"]; got != "data" {
			t.Fatalf("unexpected type: %v", got)
		}
		if nodes, ok := body["nodes"].([]interface{}); !ok || len(nodes) != 1 {
			t.Fatalf("unexpected nodes: %v", body["nodes"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-migrate-data-example",
		})
	}))
	defer server.Close()

	result, err := MigrateClusterNodeData(cli, testRegion, &MigrateClusterNodeDataRequest{
		ClusterId: "search-xxxx",
		Nodes:     []string{"i-xxxx1"},
		Type:      NODE_SPEC_TYPE_DATA,
	})
	if err != nil {
		t.Fatalf("migrate cluster node data failed: %v", err)
	}
	if result.ActionId != "action-migrate-data-example" {
		t.Fatalf("unexpected migrate data result: %+v", result)
	}
}

func TestUpdateClusterAccessWhitelist(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/access-white-ips" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		for key, want := range map[string]interface{}{
			"accessType":  "OPENSEARCH",
			"networkType": "PRIVATE",
		} {
			if got := body[key]; got != want {
				t.Fatalf("unexpected %s: got %v want %v", key, got, want)
			}
		}
		if ips, ok := body["ipWhitelist"].([]interface{}); !ok || len(ips) != 1 {
			t.Fatalf("unexpected ipWhitelist: %v", body["ipWhitelist"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId":                "search-xxxx",
			"privateIpWhitelist":       []string{"192.168.0.1"},
			"visualPrivateIpWhitelist": []string{"192.168.0.2"},
		})
	}))
	defer server.Close()

	result, err := UpdateClusterAccessWhitelist(cli, testRegion, &UpdateClusterAccessWhitelistRequest{
		ClusterId:   "search-xxxx",
		AccessType:  ACCESS_TYPE_OPENSEARCH,
		NetworkType: NETWORK_TYPE_PRIVATE,
		IpWhitelist: []string{"192.168.0.1"},
	})
	if err != nil {
		t.Fatalf("update cluster access whitelist failed: %v", err)
	}
	if len(result.PrivateIpWhitelist) != 1 || len(result.VisualPrivateIpWhitelist) != 1 {
		t.Fatalf("unexpected whitelist result: %+v", result)
	}
}

// TestUpgradeCluster populates every optional body field of an upgrade request.
func TestUpgradeCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/upgrades" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		for key, want := range map[string]interface{}{
			"upgradeType":          "KERNEL",
			"operation":            "CHECK",
			"targetVersion":        "2.19.5",
			"targetKernel":         "1.0.4",
			"upgradeMode":          "ROLLING",
			"skipHealthCheck":      true,
			"skipDeprecationCheck": true,
			"upgradeVisual":        true,
			"enableDeploySet":      true,
		} {
			if got := body[key]; got != want {
				t.Fatalf("unexpected %s: got %v want %v", key, got, want)
			}
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-upgrade-cluster-example",
		})
	}))
	defer server.Close()

	result, err := UpgradeCluster(cli, testRegion, &UpgradeClusterRequest{
		ClusterId:            "search-xxxx",
		UpgradeType:          UPGRADE_TYPE_KERNEL,
		TargetVersion:        "2.19.5",
		TargetKernel:         "1.0.4",
		Operation:            UPGRADE_OPERATION_CHECK,
		UpgradeMode:          RESTART_MODE_ROLLING,
		SkipHealthCheck:      boolPtr(true),
		SkipDeprecationCheck: boolPtr(true),
		UpgradeVisual:        boolPtr(true),
		EnableDeploySet:      boolPtr(true),
	})
	if err != nil {
		t.Fatalf("upgrade cluster failed: %v", err)
	}
	if result.ActionId != "action-upgrade-cluster-example" {
		t.Fatalf("unexpected upgrade cluster result: %+v", result)
	}
}

func TestStopCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/stop" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-stop-cluster-example",
		})
	}))
	defer server.Close()

	result, err := StopCluster(cli, testRegion, &StopClusterRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("stop cluster failed: %v", err)
	}
	if result.ActionId != "action-stop-cluster-example" {
		t.Fatalf("unexpected stop cluster result: %+v", result)
	}
}

func TestStopClusterNodes(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/nodes/stop" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if nodes, ok := body["nodes"].([]interface{}); !ok || len(nodes) != 2 {
			t.Fatalf("unexpected nodes: %v", body["nodes"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-stop-nodes-example",
		})
	}))
	defer server.Close()

	result, err := StopClusterNodes(cli, testRegion, &ClusterNodesRequest{
		ClusterId: "search-xxxx",
		Nodes:     []string{"i-xxxx1", "i-xxxx2"},
	})
	if err != nil {
		t.Fatalf("stop cluster nodes failed: %v", err)
	}
	if result.ActionId != "action-stop-nodes-example" {
		t.Fatalf("unexpected stop nodes result: %+v", result)
	}
}

func TestUpdateClusterProtocol(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/protocol" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if got := body["enableHttps"]; got != true {
			t.Fatalf("unexpected enableHttps: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-change-protocol-example",
		})
	}))
	defer server.Close()

	result, err := UpdateClusterProtocol(cli, testRegion, &UpdateClusterProtocolRequest{
		ClusterId:   "search-xxxx",
		EnableHttps: boolPtr(true),
	})
	if err != nil {
		t.Fatalf("update cluster protocol failed: %v", err)
	}
	if result.ActionId != "action-change-protocol-example" {
		t.Fatalf("unexpected protocol result: %+v", result)
	}
}

// TestRestartCluster restarts the whole cluster in rolling mode, the combination where batchCount and
// batchUnit must both stay unset.
func TestRestartCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/restart" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		for key, want := range map[string]interface{}{
			"restartType": "CLUSTER",
			"mode":        "ROLLING",
		} {
			if got := body[key]; got != want {
				t.Fatalf("unexpected %s: got %v want %v", key, got, want)
			}
		}
		if body["nodes"] != nil || body["batchUnit"] != nil || body["batchCount"] != nil {
			t.Fatalf("unexpected optional fields: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-restart-cluster-example",
		})
	}))
	defer server.Close()

	result, err := RestartCluster(cli, testRegion, &RestartClusterRequest{
		ClusterId:   "search-xxxx",
		RestartType: RESTART_TYPE_CLUSTER,
		Mode:        RESTART_MODE_ROLLING,
	})
	if err != nil {
		t.Fatalf("restart cluster failed: %v", err)
	}
	if result.ActionId != "action-restart-cluster-example" {
		t.Fatalf("unexpected restart cluster result: %+v", result)
	}
}

// TestRestartClusterNodesInForceMode covers the opposite validation path: a node-scoped restart in
// force mode, which requires nodes plus both batch fields, with a percent batch inside (0, 100].
func TestRestartClusterNodesInForceMode(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/restart" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		for key, want := range map[string]interface{}{
			"restartType": "NODE",
			"mode":        "FORCE",
			"batchUnit":   "PERCENT",
			"batchCount":  float64(50),
		} {
			if got := body[key]; got != want {
				t.Fatalf("unexpected %s: got %v want %v", key, got, want)
			}
		}
		if nodes, ok := body["nodes"].([]interface{}); !ok || len(nodes) != 1 {
			t.Fatalf("unexpected nodes: %v", body["nodes"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId": "search-xxxx",
			"actionId":  "action-restart-nodes-example",
		})
	}))
	defer server.Close()

	result, err := RestartCluster(cli, testRegion, &RestartClusterRequest{
		ClusterId:   "search-xxxx",
		RestartType: RESTART_TYPE_NODE,
		Mode:        RESTART_MODE_FORCE,
		Nodes:       []string{"i-xxxx1"},
		BatchUnit:   BATCH_UNIT_PERCENT,
		BatchCount:  clusterBatchCountPtr(50),
	})
	if err != nil {
		t.Fatalf("restart cluster nodes failed: %v", err)
	}
	if result.ActionId != "action-restart-nodes-example" {
		t.Fatalf("unexpected restart nodes result: %+v", result)
	}
}

// TestClusterAPIRejectsNilClient covers the ErrNilClient guard of all 27 cluster functions.
func TestClusterAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"CreateCluster": func() error { _, err := CreateCluster(nil, testRegion, &CreateClusterRequest{}); return err },
		"ListClusters":  func() error { _, err := ListClusters(nil, testRegion, &ListClustersRequest{}); return err },
		"GetCluster":    func() error { _, err := GetCluster(nil, testRegion, &GetClusterRequest{}); return err },
		"ListClusterNodes": func() error {
			_, err := ListClusterNodes(nil, testRegion, &ListClusterNodesRequest{})
			return err
		},
		"QueryAvailableSpecs": func() error {
			_, err := QueryAvailableSpecs(nil, testRegion, &QueryAvailableSpecsRequest{})
			return err
		},
		"QueryAvailableKernels": func() error {
			_, err := QueryAvailableKernels(nil, testRegion, &QueryAvailableKernelsRequest{})
			return err
		},
		"DownloadClusterCert": func() error {
			_, err := DownloadClusterCert(nil, testRegion, &DownloadClusterCertRequest{})
			return err
		},
		"UpdateClusterName": func() error {
			_, err := UpdateClusterName(nil, testRegion, &UpdateClusterNameRequest{})
			return err
		},
		"UpdateClusterMaintenance": func() error {
			_, err := UpdateClusterMaintenance(nil, testRegion, &UpdateClusterMaintenanceRequest{})
			return err
		},
		"UpdateClusterDeletionProtection": func() error {
			_, err := UpdateClusterDeletionProtection(nil, testRegion, &UpdateClusterDeletionProtectionRequest{})
			return err
		},
		"RecoverCluster": func() error { _, err := RecoverCluster(nil, testRegion, &RecoverClusterRequest{}); return err },
		"DeleteCluster":  func() error { _, err := DeleteCluster(nil, testRegion, &DeleteClusterRequest{}); return err },
		"ResizeCluster":  func() error { _, err := ResizeCluster(nil, testRegion, &ResizeClusterRequest{}); return err },
		"ListSuggestedMigrationNodes": func() error {
			_, err := ListSuggestedMigrationNodes(nil, testRegion, &ListSuggestedMigrationNodesRequest{})
			return err
		},
		"ListMigratableNodes": func() error {
			_, err := ListMigratableNodes(nil, testRegion, &ListMigratableNodesRequest{})
			return err
		},
		"QueryAvailableUpgrades": func() error {
			_, err := QueryAvailableUpgrades(nil, testRegion, &QueryAvailableUpgradesRequest{})
			return err
		},
		"UpdateClusterPublicAccess": func() error {
			_, err := UpdateClusterPublicAccess(nil, testRegion, &UpdateClusterPublicAccessRequest{})
			return err
		},
		"UpdateClusterCerebro": func() error {
			_, err := UpdateClusterCerebro(nil, testRegion, &UpdateClusterCerebroRequest{})
			return err
		},
		"StartCluster": func() error { _, err := StartCluster(nil, testRegion, &StartClusterRequest{}); return err },
		"StartClusterNodes": func() error {
			_, err := StartClusterNodes(nil, testRegion, &ClusterNodesRequest{})
			return err
		},
		"MigrateClusterNodeData": func() error {
			_, err := MigrateClusterNodeData(nil, testRegion, &MigrateClusterNodeDataRequest{})
			return err
		},
		"UpdateClusterAccessWhitelist": func() error {
			_, err := UpdateClusterAccessWhitelist(nil, testRegion, &UpdateClusterAccessWhitelistRequest{})
			return err
		},
		"UpgradeCluster": func() error { _, err := UpgradeCluster(nil, testRegion, &UpgradeClusterRequest{}); return err },
		"StopCluster":    func() error { _, err := StopCluster(nil, testRegion, &StopClusterRequest{}); return err },
		"StopClusterNodes": func() error {
			_, err := StopClusterNodes(nil, testRegion, &ClusterNodesRequest{})
			return err
		},
		"UpdateClusterProtocol": func() error {
			_, err := UpdateClusterProtocol(nil, testRegion, &UpdateClusterProtocolRequest{})
			return err
		},
		"RestartCluster": func() error { _, err := RestartCluster(nil, testRegion, &RestartClusterRequest{}); return err },
	}
	if len(cases) != 27 {
		t.Fatalf("expected all 27 cluster functions to be covered, got %d", len(cases))
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

// TestClusterAPIRejectsNilRequest covers the 24 cluster functions that require a request. The other
// three are deliberately absent: QueryAvailableSpecs and QueryAvailableKernels accept nil (see
// TestClusterAPIAcceptsNilRequest), and ListClusters dereferences the request without a nil guard, so
// calling it with nil panics instead of returning an error.
func TestClusterAPIRejectsNilRequest(t *testing.T) {
	cli := newClusterRejectingClient(t)

	cases := map[string]func() error{
		"CreateCluster":    func() error { _, err := CreateCluster(cli, testRegion, nil); return err },
		"GetCluster":       func() error { _, err := GetCluster(cli, testRegion, nil); return err },
		"ListClusterNodes": func() error { _, err := ListClusterNodes(cli, testRegion, nil); return err },
		"DownloadClusterCert": func() error {
			_, err := DownloadClusterCert(cli, testRegion, nil)
			return err
		},
		"UpdateClusterName": func() error { _, err := UpdateClusterName(cli, testRegion, nil); return err },
		"UpdateClusterMaintenance": func() error {
			_, err := UpdateClusterMaintenance(cli, testRegion, nil)
			return err
		},
		"UpdateClusterDeletionProtection": func() error {
			_, err := UpdateClusterDeletionProtection(cli, testRegion, nil)
			return err
		},
		"RecoverCluster": func() error { _, err := RecoverCluster(cli, testRegion, nil); return err },
		"DeleteCluster":  func() error { _, err := DeleteCluster(cli, testRegion, nil); return err },
		"ResizeCluster":  func() error { _, err := ResizeCluster(cli, testRegion, nil); return err },
		"ListSuggestedMigrationNodes": func() error {
			_, err := ListSuggestedMigrationNodes(cli, testRegion, nil)
			return err
		},
		"ListMigratableNodes": func() error { _, err := ListMigratableNodes(cli, testRegion, nil); return err },
		"QueryAvailableUpgrades": func() error {
			_, err := QueryAvailableUpgrades(cli, testRegion, nil)
			return err
		},
		"UpdateClusterPublicAccess": func() error {
			_, err := UpdateClusterPublicAccess(cli, testRegion, nil)
			return err
		},
		"UpdateClusterCerebro": func() error { _, err := UpdateClusterCerebro(cli, testRegion, nil); return err },
		"StartCluster":         func() error { _, err := StartCluster(cli, testRegion, nil); return err },
		"StartClusterNodes":    func() error { _, err := StartClusterNodes(cli, testRegion, nil); return err },
		"MigrateClusterNodeData": func() error {
			_, err := MigrateClusterNodeData(cli, testRegion, nil)
			return err
		},
		"UpdateClusterAccessWhitelist": func() error {
			_, err := UpdateClusterAccessWhitelist(cli, testRegion, nil)
			return err
		},
		"UpgradeCluster":        func() error { _, err := UpgradeCluster(cli, testRegion, nil); return err },
		"StopCluster":           func() error { _, err := StopCluster(cli, testRegion, nil); return err },
		"StopClusterNodes":      func() error { _, err := StopClusterNodes(cli, testRegion, nil); return err },
		"UpdateClusterProtocol": func() error { _, err := UpdateClusterProtocol(cli, testRegion, nil); return err },
		"RestartCluster":        func() error { _, err := RestartCluster(cli, testRegion, nil); return err },
	}
	if len(cases) != 24 {
		t.Fatalf("expected 24 nil-request guards, got %d", len(cases))
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

// TestClusterAPIRejectsMissingClusterId covers the eight cluster functions whose only required field
// is clusterId.
func TestClusterAPIRejectsMissingClusterId(t *testing.T) {
	cli := newClusterRejectingClient(t)

	cases := map[string]func() error{
		"GetCluster":       func() error { _, err := GetCluster(cli, testRegion, &GetClusterRequest{}); return err },
		"ListClusterNodes": func() error { _, err := ListClusterNodes(cli, testRegion, &ListClusterNodesRequest{}); return err },
		"DownloadClusterCert": func() error {
			_, err := DownloadClusterCert(cli, testRegion, &DownloadClusterCertRequest{})
			return err
		},
		"RecoverCluster": func() error { _, err := RecoverCluster(cli, testRegion, &RecoverClusterRequest{}); return err },
		"DeleteCluster": func() error {
			_, err := DeleteCluster(cli, testRegion, &DeleteClusterRequest{Recycle: boolPtr(true)})
			return err
		},
		"QueryAvailableUpgrades": func() error {
			_, err := QueryAvailableUpgrades(cli, testRegion, &QueryAvailableUpgradesRequest{})
			return err
		},
		"StartCluster": func() error { _, err := StartCluster(cli, testRegion, &StartClusterRequest{}); return err },
		"StopCluster":  func() error { _, err := StopCluster(cli, testRegion, &StopClusterRequest{}); return err },
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestCreateClusterRejectsMissingRequiredFields covers all 11 required-field checks of CreateCluster,
// including the non-string ones (a nil enableHttps flag and the two empty slices).
func TestCreateClusterRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := CreateClusterRequest{
		Payment:       PAYMENT_POSTPAID,
		Name:          "bes-test",
		Engine:        ENGINE_OPENSEARCH,
		EngineType:    ENGINE_TYPE_COMMUNITY,
		Version:       "2.11.0",
		AdminPassword: testPassword,
		EnableHttps:   boolPtr(true),
		LogicalZones:  []string{"cn-bj-a"},
		VpcId:         "vpc-xxxx",
		SubnetId:      "sbn-xxxx",
		NodeSpecs:     []NodeSpec{{Type: NODE_SPEC_TYPE_DATA, NodeCount: intPtr(2)}},
	}
	cases := map[string]func(*CreateClusterRequest){
		"payment":       func(r *CreateClusterRequest) { r.Payment = "" },
		"name":          func(r *CreateClusterRequest) { r.Name = "" },
		"engine":        func(r *CreateClusterRequest) { r.Engine = "" },
		"engineType":    func(r *CreateClusterRequest) { r.EngineType = "" },
		"version":       func(r *CreateClusterRequest) { r.Version = "" },
		"adminPassword": func(r *CreateClusterRequest) { r.AdminPassword = "" },
		"enableHttps":   func(r *CreateClusterRequest) { r.EnableHttps = nil },
		"logicalZones":  func(r *CreateClusterRequest) { r.LogicalZones = nil },
		"vpcId":         func(r *CreateClusterRequest) { r.VpcId = "" },
		"subnetId":      func(r *CreateClusterRequest) { r.SubnetId = "" },
		"nodeSpecs":     func(r *CreateClusterRequest) { r.NodeSpecs = nil },
	}
	if len(cases) != 11 {
		t.Fatalf("expected 11 required-field checks, got %d", len(cases))
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := CreateCluster(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestCreateClusterRejectsUnusableCredentials covers the two credential error paths that precede the
// HTTP call: no credentials configured on the client, and a secret key too short to key the AES-128
// encryption of the admin password.
func TestCreateClusterRejectsUnusableCredentials(t *testing.T) {
	shortKeyCredentials, err := auth.NewBceCredentials(testAK, "short")
	if err != nil {
		t.Fatalf("new credentials failed: %v", err)
	}

	cases := map[string]bce.Client{
		"missingCredentials": newClusterOfflineClient(t, nil),
		"shortSecretKey":     newClusterOfflineClient(t, shortKeyCredentials),
	}
	for name, cli := range cases {
		_, err := CreateCluster(cli, testRegion, &CreateClusterRequest{
			Payment:       PAYMENT_POSTPAID,
			Name:          "bes-test",
			Engine:        ENGINE_OPENSEARCH,
			EngineType:    ENGINE_TYPE_COMMUNITY,
			Version:       "2.11.0",
			AdminPassword: testPassword,
			EnableHttps:   boolPtr(true),
			LogicalZones:  []string{"cn-bj-a"},
			VpcId:         "vpc-xxxx",
			SubnetId:      "sbn-xxxx",
			NodeSpecs:     []NodeSpec{{Type: NODE_SPEC_TYPE_DATA, NodeCount: intPtr(2)}},
		})
		if err == nil {
			t.Fatalf("%s: expected error", name)
		}
	}
}

func TestUpdateClusterNameRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := UpdateClusterNameRequest{ClusterId: "search-xxxx", NewName: "bes-new-name"}
	cases := map[string]func(*UpdateClusterNameRequest){
		"clusterId": func(r *UpdateClusterNameRequest) { r.ClusterId = "" },
		"newName":   func(r *UpdateClusterNameRequest) { r.NewName = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateClusterName(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateClusterMaintenanceRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := UpdateClusterMaintenanceRequest{ClusterId: "search-xxxx", MaintenanceTimeZone: "Asia/Shanghai"}
	cases := map[string]func(*UpdateClusterMaintenanceRequest){
		"clusterId":           func(r *UpdateClusterMaintenanceRequest) { r.ClusterId = "" },
		"maintenanceTimeZone": func(r *UpdateClusterMaintenanceRequest) { r.MaintenanceTimeZone = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateClusterMaintenance(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateClusterDeletionProtectionRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := UpdateClusterDeletionProtectionRequest{ClusterId: "search-xxxx", DeletionProtection: boolPtr(true)}
	cases := map[string]func(*UpdateClusterDeletionProtectionRequest){
		"clusterId":          func(r *UpdateClusterDeletionProtectionRequest) { r.ClusterId = "" },
		"deletionProtection": func(r *UpdateClusterDeletionProtectionRequest) { r.DeletionProtection = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateClusterDeletionProtection(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestResizeClusterRejectsInvalidRequests covers the two required fields plus the mutual exclusion of
// nodeSpecs and deleteNodeTypes.
func TestResizeClusterRejectsInvalidRequests(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := ResizeClusterRequest{
		ClusterId: "search-xxxx",
		Mode:      RESIZE_MODE_DIRECT,
		NodeSpecs: []NodeSpec{{Type: NODE_SPEC_TYPE_COORDINATING}},
	}
	cases := map[string]func(*ResizeClusterRequest){
		"clusterId": func(r *ResizeClusterRequest) { r.ClusterId = "" },
		"mode":      func(r *ResizeClusterRequest) { r.Mode = "" },
		"nodeSpecsAndDeleteNodeTypesTogether": func(r *ResizeClusterRequest) {
			r.DeleteNodeTypes = []string{NODE_SPEC_TYPE_COORDINATING}
		},
	}
	for name, mutate := range cases {
		request := full
		mutate(&request)
		if _, err := ResizeCluster(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestListSuggestedMigrationNodesRejectsMissingRequiredFields covers the two string fields plus the
// non-positive migrateCount check, which rejects both zero and a negative count.
func TestListSuggestedMigrationNodesRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := ListSuggestedMigrationNodesRequest{
		ClusterId:    "search-xxxx",
		Type:         NODE_SPEC_TYPE_DATA,
		MigrateCount: 2,
	}
	cases := map[string]func(*ListSuggestedMigrationNodesRequest){
		"clusterId":            func(r *ListSuggestedMigrationNodesRequest) { r.ClusterId = "" },
		"type":                 func(r *ListSuggestedMigrationNodesRequest) { r.Type = "" },
		"zeroMigrateCount":     func(r *ListSuggestedMigrationNodesRequest) { r.MigrateCount = 0 },
		"negativeMigrateCount": func(r *ListSuggestedMigrationNodesRequest) { r.MigrateCount = -1 },
	}
	for name, mutate := range cases {
		request := full
		mutate(&request)
		if _, err := ListSuggestedMigrationNodes(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestListMigratableNodesRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := ListMigratableNodesRequest{ClusterId: "search-xxxx", Type: NODE_SPEC_TYPE_DATA}
	cases := map[string]func(*ListMigratableNodesRequest){
		"clusterId": func(r *ListMigratableNodesRequest) { r.ClusterId = "" },
		"type":      func(r *ListMigratableNodesRequest) { r.Type = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := ListMigratableNodes(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateClusterPublicAccessRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := UpdateClusterPublicAccessRequest{ClusterId: "search-xxxx", Enabled: boolPtr(true)}
	cases := map[string]func(*UpdateClusterPublicAccessRequest){
		"clusterId": func(r *UpdateClusterPublicAccessRequest) { r.ClusterId = "" },
		"enabled":   func(r *UpdateClusterPublicAccessRequest) { r.Enabled = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateClusterPublicAccess(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestUpdateClusterCerebroRejectsInvalidRequests covers the missing clusterId as well as both halves of
// the enable-only rule: the API cannot be used to disable Cerebro, so a nil and a false flag are both
// rejected.
func TestUpdateClusterCerebroRejectsInvalidRequests(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := UpdateClusterCerebroRequest{ClusterId: "search-xxxx", Enabled: boolPtr(true)}
	cases := map[string]func(*UpdateClusterCerebroRequest){
		"clusterId":      func(r *UpdateClusterCerebroRequest) { r.ClusterId = "" },
		"nilEnabled":     func(r *UpdateClusterCerebroRequest) { r.Enabled = nil },
		"disableCerebro": func(r *UpdateClusterCerebroRequest) { r.Enabled = boolPtr(false) },
	}
	for name, mutate := range cases {
		request := full
		mutate(&request)
		if _, err := UpdateClusterCerebro(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestClusterNodeActionsRejectMissingRequiredFields covers the shared ClusterNodesRequest validation of
// StartClusterNodes and StopClusterNodes, which each reject an empty clusterId and an empty node list.
func TestClusterNodeActionsRejectMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := ClusterNodesRequest{ClusterId: "search-xxxx", Nodes: []string{"i-xxxx1"}}
	mutations := map[string]func(*ClusterNodesRequest){
		"clusterId": func(r *ClusterNodesRequest) { r.ClusterId = "" },
		"nodes":     func(r *ClusterNodesRequest) { r.Nodes = nil },
	}
	actions := map[string]func(*ClusterNodesRequest) error{
		"StartClusterNodes": func(r *ClusterNodesRequest) error { _, err := StartClusterNodes(cli, testRegion, r); return err },
		"StopClusterNodes":  func(r *ClusterNodesRequest) error { _, err := StopClusterNodes(cli, testRegion, r); return err },
	}
	for action, call := range actions {
		for name, clear := range mutations {
			request := full
			clear(&request)
			if err := call(&request); err == nil {
				t.Fatalf("%s/%s: expected validation error", action, name)
			}
		}
	}
}

func TestMigrateClusterNodeDataRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := MigrateClusterNodeDataRequest{
		ClusterId: "search-xxxx",
		Nodes:     []string{"i-xxxx1"},
		Type:      NODE_SPEC_TYPE_DATA,
	}
	cases := map[string]func(*MigrateClusterNodeDataRequest){
		"clusterId": func(r *MigrateClusterNodeDataRequest) { r.ClusterId = "" },
		"nodes":     func(r *MigrateClusterNodeDataRequest) { r.Nodes = nil },
		"type":      func(r *MigrateClusterNodeDataRequest) { r.Type = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := MigrateClusterNodeData(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateClusterAccessWhitelistRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := UpdateClusterAccessWhitelistRequest{
		ClusterId:   "search-xxxx",
		AccessType:  ACCESS_TYPE_OPENSEARCH,
		NetworkType: NETWORK_TYPE_PRIVATE,
		IpWhitelist: []string{"192.168.0.1"},
	}
	cases := map[string]func(*UpdateClusterAccessWhitelistRequest){
		"clusterId":   func(r *UpdateClusterAccessWhitelistRequest) { r.ClusterId = "" },
		"accessType":  func(r *UpdateClusterAccessWhitelistRequest) { r.AccessType = "" },
		"networkType": func(r *UpdateClusterAccessWhitelistRequest) { r.NetworkType = "" },
		"ipWhitelist": func(r *UpdateClusterAccessWhitelistRequest) { r.IpWhitelist = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateClusterAccessWhitelist(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpgradeClusterRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := UpgradeClusterRequest{
		ClusterId:   "search-xxxx",
		UpgradeType: UPGRADE_TYPE_KERNEL,
		Operation:   UPGRADE_OPERATION_CHECK,
	}
	cases := map[string]func(*UpgradeClusterRequest){
		"clusterId":   func(r *UpgradeClusterRequest) { r.ClusterId = "" },
		"upgradeType": func(r *UpgradeClusterRequest) { r.UpgradeType = "" },
		"operation":   func(r *UpgradeClusterRequest) { r.Operation = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpgradeCluster(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateClusterProtocolRejectsMissingRequiredFields(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := UpdateClusterProtocolRequest{ClusterId: "search-xxxx", EnableHttps: boolPtr(true)}
	cases := map[string]func(*UpdateClusterProtocolRequest){
		"clusterId":   func(r *UpdateClusterProtocolRequest) { r.ClusterId = "" },
		"enableHttps": func(r *UpdateClusterProtocolRequest) { r.EnableHttps = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateClusterProtocol(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestRestartClusterRejectsInvalidRequests covers every validation branch of RestartCluster: the three
// required fields, the restartType/nodes agreement, the batch fields each mode allows or demands, and
// the percent range check.
func TestRestartClusterRejectsInvalidRequests(t *testing.T) {
	cli := newClusterRejectingClient(t)

	full := RestartClusterRequest{
		ClusterId:   "search-xxxx",
		RestartType: RESTART_TYPE_CLUSTER,
		Mode:        RESTART_MODE_ROLLING,
	}
	cases := map[string]func(*RestartClusterRequest){
		"clusterId":   func(r *RestartClusterRequest) { r.ClusterId = "" },
		"restartType": func(r *RestartClusterRequest) { r.RestartType = "" },
		"mode":        func(r *RestartClusterRequest) { r.Mode = "" },
		"nodesSetForClusterRestart": func(r *RestartClusterRequest) {
			r.Nodes = []string{"i-xxxx1"}
		},
		"nodesMissingForNodeRestart": func(r *RestartClusterRequest) {
			r.RestartType = RESTART_TYPE_NODE
		},
		"batchCountSetForRollingMode": func(r *RestartClusterRequest) {
			r.BatchCount = clusterBatchCountPtr(2)
		},
		"batchUnitSetForBlueGreenMode": func(r *RestartClusterRequest) {
			r.Mode = RESTART_MODE_BLUE_GREEN
			r.BatchUnit = BATCH_UNIT_COUNT
		},
		"batchUnitMissingForForceMode": func(r *RestartClusterRequest) {
			r.Mode = RESTART_MODE_FORCE
		},
		"batchCountMissingForForceMode": func(r *RestartClusterRequest) {
			r.Mode = RESTART_MODE_FORCE
			r.BatchUnit = BATCH_UNIT_PERCENT
		},
		"percentBatchCountTooSmall": func(r *RestartClusterRequest) {
			r.Mode = RESTART_MODE_FORCE
			r.BatchUnit = BATCH_UNIT_PERCENT
			r.BatchCount = clusterBatchCountPtr(0)
		},
		"percentBatchCountTooLarge": func(r *RestartClusterRequest) {
			r.Mode = RESTART_MODE_FORCE
			r.BatchUnit = BATCH_UNIT_PERCENT
			r.BatchCount = clusterBatchCountPtr(101)
		},
	}
	for name, mutate := range cases {
		request := full
		mutate(&request)
		if _, err := RestartCluster(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestClusterAPIPropagatesServerError covers the error-propagation branch of the 26 cluster functions
// that go through the shared JSON request helper; DownloadClusterCert streams its response and is
// covered by TestDownloadClusterCertPropagatesErrors. 400 is used rather than 500 on purpose: the test
// client mirrors the production retry policy, which retries 5xx with backoff.
func TestClusterAPIPropagatesServerError(t *testing.T) {
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
		"CreateCluster": func() error {
			_, err := CreateCluster(cli, testRegion, &CreateClusterRequest{
				Payment:       PAYMENT_POSTPAID,
				Name:          "bes-test",
				Engine:        ENGINE_OPENSEARCH,
				EngineType:    ENGINE_TYPE_COMMUNITY,
				Version:       "2.11.0",
				AdminPassword: testPassword,
				EnableHttps:   boolPtr(true),
				LogicalZones:  []string{"cn-bj-a"},
				VpcId:         "vpc-xxxx",
				SubnetId:      "sbn-xxxx",
				NodeSpecs:     []NodeSpec{{Type: NODE_SPEC_TYPE_DATA, NodeCount: intPtr(2)}},
			})
			return err
		},
		"ListClusters": func() error { _, err := ListClusters(cli, testRegion, &ListClustersRequest{}); return err },
		"GetCluster": func() error {
			_, err := GetCluster(cli, testRegion, &GetClusterRequest{ClusterId: "search-xxxx"})
			return err
		},
		"ListClusterNodes": func() error {
			_, err := ListClusterNodes(cli, testRegion, &ListClusterNodesRequest{ClusterId: "search-xxxx"})
			return err
		},
		"QueryAvailableSpecs": func() error {
			_, err := QueryAvailableSpecs(cli, testRegion, &QueryAvailableSpecsRequest{
				Engine:       ENGINE_OPENSEARCH,
				LogicalZones: []string{"cn-bj-a"},
			})
			return err
		},
		"QueryAvailableKernels": func() error {
			_, err := QueryAvailableKernels(cli, testRegion, &QueryAvailableKernelsRequest{})
			return err
		},
		"UpdateClusterName": func() error {
			_, err := UpdateClusterName(cli, testRegion, &UpdateClusterNameRequest{
				ClusterId: "search-xxxx",
				NewName:   "bes-new-name",
			})
			return err
		},
		"UpdateClusterMaintenance": func() error {
			_, err := UpdateClusterMaintenance(cli, testRegion, &UpdateClusterMaintenanceRequest{
				ClusterId:           "search-xxxx",
				MaintenanceTimeZone: "Asia/Shanghai",
			})
			return err
		},
		"UpdateClusterDeletionProtection": func() error {
			_, err := UpdateClusterDeletionProtection(cli, testRegion, &UpdateClusterDeletionProtectionRequest{
				ClusterId:          "search-xxxx",
				DeletionProtection: boolPtr(true),
			})
			return err
		},
		"RecoverCluster": func() error {
			_, err := RecoverCluster(cli, testRegion, &RecoverClusterRequest{ClusterId: "search-xxxx"})
			return err
		},
		"DeleteCluster": func() error {
			_, err := DeleteCluster(cli, testRegion, &DeleteClusterRequest{ClusterId: "search-xxxx"})
			return err
		},
		"ResizeCluster": func() error {
			_, err := ResizeCluster(cli, testRegion, &ResizeClusterRequest{
				ClusterId: "search-xxxx",
				Mode:      RESIZE_MODE_DIRECT,
			})
			return err
		},
		"ListSuggestedMigrationNodes": func() error {
			_, err := ListSuggestedMigrationNodes(cli, testRegion, &ListSuggestedMigrationNodesRequest{
				ClusterId:    "search-xxxx",
				Type:         NODE_SPEC_TYPE_DATA,
				MigrateCount: 2,
			})
			return err
		},
		"ListMigratableNodes": func() error {
			_, err := ListMigratableNodes(cli, testRegion, &ListMigratableNodesRequest{
				ClusterId: "search-xxxx",
				Type:      NODE_SPEC_TYPE_DATA,
			})
			return err
		},
		"QueryAvailableUpgrades": func() error {
			_, err := QueryAvailableUpgrades(cli, testRegion, &QueryAvailableUpgradesRequest{
				ClusterId: "search-xxxx",
			})
			return err
		},
		"UpdateClusterPublicAccess": func() error {
			_, err := UpdateClusterPublicAccess(cli, testRegion, &UpdateClusterPublicAccessRequest{
				ClusterId: "search-xxxx",
				Enabled:   boolPtr(true),
			})
			return err
		},
		"UpdateClusterCerebro": func() error {
			_, err := UpdateClusterCerebro(cli, testRegion, &UpdateClusterCerebroRequest{
				ClusterId: "search-xxxx",
				Enabled:   boolPtr(true),
			})
			return err
		},
		"StartCluster": func() error {
			_, err := StartCluster(cli, testRegion, &StartClusterRequest{ClusterId: "search-xxxx"})
			return err
		},
		"StartClusterNodes": func() error {
			_, err := StartClusterNodes(cli, testRegion, &ClusterNodesRequest{
				ClusterId: "search-xxxx",
				Nodes:     []string{"i-xxxx1"},
			})
			return err
		},
		"MigrateClusterNodeData": func() error {
			_, err := MigrateClusterNodeData(cli, testRegion, &MigrateClusterNodeDataRequest{
				ClusterId: "search-xxxx",
				Nodes:     []string{"i-xxxx1"},
				Type:      NODE_SPEC_TYPE_DATA,
			})
			return err
		},
		"UpdateClusterAccessWhitelist": func() error {
			_, err := UpdateClusterAccessWhitelist(cli, testRegion, &UpdateClusterAccessWhitelistRequest{
				ClusterId:   "search-xxxx",
				AccessType:  ACCESS_TYPE_OPENSEARCH,
				NetworkType: NETWORK_TYPE_PRIVATE,
				IpWhitelist: []string{"192.168.0.1"},
			})
			return err
		},
		"UpgradeCluster": func() error {
			_, err := UpgradeCluster(cli, testRegion, &UpgradeClusterRequest{
				ClusterId:   "search-xxxx",
				UpgradeType: UPGRADE_TYPE_KERNEL,
				Operation:   UPGRADE_OPERATION_CHECK,
			})
			return err
		},
		"StopCluster": func() error {
			_, err := StopCluster(cli, testRegion, &StopClusterRequest{ClusterId: "search-xxxx"})
			return err
		},
		"StopClusterNodes": func() error {
			_, err := StopClusterNodes(cli, testRegion, &ClusterNodesRequest{
				ClusterId: "search-xxxx",
				Nodes:     []string{"i-xxxx1"},
			})
			return err
		},
		"UpdateClusterProtocol": func() error {
			_, err := UpdateClusterProtocol(cli, testRegion, &UpdateClusterProtocolRequest{
				ClusterId:   "search-xxxx",
				EnableHttps: boolPtr(true),
			})
			return err
		},
		"RestartCluster": func() error {
			_, err := RestartCluster(cli, testRegion, &RestartClusterRequest{
				ClusterId:   "search-xxxx",
				RestartType: RESTART_TYPE_CLUSTER,
				Mode:        RESTART_MODE_ROLLING,
			})
			return err
		},
	}
	if len(cases) != 26 {
		t.Fatalf("expected 26 error-propagation cases, got %d", len(cases))
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for 400 response", name)
		}
	}
}

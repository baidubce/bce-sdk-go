package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

// newIndexRejectingClient returns a client backed by a server that fails the test if it is ever
// reached. Every index validation case below has to be rejected before any HTTP call is made, so
// this keeps each of them to a single line of setup.
func newIndexRejectingClient(t *testing.T) bce.Client {
	t.Helper()
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	t.Cleanup(server.Close)
	return cli
}

// TestCreateIndex populates every optional body field so each of them is serialised at least once.
func TestCreateIndex(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/openapi_test_index" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if _, ok := body["settings"].(map[string]interface{}); !ok {
			t.Fatalf("expected settings in body: %+v", body)
		}
		if _, ok := body["mappings"].(map[string]interface{}); !ok {
			t.Fatalf("expected mappings in body: %+v", body)
		}
		if _, ok := body["aliases"].(map[string]interface{}); !ok {
			t.Fatalf("expected aliases in body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"indexName": "openapi_test_index"})
	}))
	defer server.Close()

	result, err := CreateIndex(cli, testRegion, &CreateIndexRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
		Settings: map[string]interface{}{
			"index": map[string]interface{}{"number_of_shards": "1", "number_of_replicas": "0"},
		},
		Mappings: map[string]interface{}{
			"properties": map[string]interface{}{"title": map[string]interface{}{"type": "text"}},
		},
		Aliases: map[string]interface{}{"openapi_test_alias": map[string]interface{}{}},
	})
	if err != nil {
		t.Fatalf("create index failed: %v", err)
	}
	if result.IndexName != "openapi_test_index" {
		t.Fatalf("unexpected index name: %s", result.IndexName)
	}
}

// TestListIndices populates every optional query field so each conditional query-parameter branch
// in ListIndices is exercised.
func TestListIndices(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		query := r.URL.Query()
		expected := map[string]string{
			"indexName":     "top_queries-*",
			"health":        "green",
			"includeSystem": "true",
			"pageNo":        "1",
			"pageSize":      "10",
			"orderBy":       "indexName",
			"order":         "desc",
		}
		for key, want := range expected {
			if got := query.Get(key); got != want {
				t.Fatalf("unexpected %s: got %s want %s", key, got, want)
			}
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"indices": []map[string]interface{}{{
				"indexName":     "top_queries-2026.07.09-44484",
				"health":        "green",
				"status":        "open",
				"primaryShards": "1",
				"replicas":      "2",
				"documentCount": "11",
				"storageSize":   "291.2kb",
				"createTime":    "1783567502389",
				"category":      "default",
			}},
			"pageNo":     1,
			"pageSize":   10,
			"totalCount": 1,
		})
	}))
	defer server.Close()

	result, err := ListIndices(cli, testRegion, &ListIndicesRequest{
		ClusterId:     "search-xxxx",
		IndexName:     "top_queries-*",
		Health:        "green",
		IncludeSystem: boolPtr(true),
		PageNo:        1,
		PageSize:      10,
		OrderBy:       "indexName",
		Order:         "desc",
	})
	if err != nil {
		t.Fatalf("list indices failed: %v", err)
	}
	if len(result.Indices) != 1 || result.Indices[0].IndexName != "top_queries-2026.07.09-44484" {
		t.Fatalf("unexpected indices: %+v", result.Indices)
	}
	if result.TotalCount != 1 {
		t.Fatalf("unexpected total count: %d", result.TotalCount)
	}
}

// TestListIndicesWithoutOptionalFields covers the request path where every optional query
// parameter is omitted from the query string entirely.
func TestListIndicesWithoutOptionalFields(t *testing.T) {
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

	result, err := ListIndices(cli, testRegion, &ListIndicesRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("list indices failed: %v", err)
	}
	if result.TotalCount != 0 || len(result.Indices) != 0 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetIndex(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/openapi_test_index" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"index": map[string]interface{}{
				"indexName":     "openapi_test_index",
				"health":        "green",
				"status":        "open",
				"primaryShards": "1",
				"replicas":      "0",
				"documentCount": "0",
				"storageSize":   "208b",
				"createTime":    "1783567858578",
				"category":      "default",
			},
			"mappings": `{"properties":{"title":{"type":"text"}}}`,
			"settings": `{"index.number_of_shards":"1"}`,
			"aliases":  "{}",
		})
	}))
	defer server.Close()

	result, err := GetIndex(cli, testRegion, &GetIndexRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
	})
	if err != nil {
		t.Fatalf("get index failed: %v", err)
	}
	if result.Index.IndexName != "openapi_test_index" {
		t.Fatalf("unexpected index name: %s", result.Index.IndexName)
	}
	if result.Mappings != `{"properties":{"title":{"type":"text"}}}` {
		t.Fatalf("unexpected mappings: %s", result.Mappings)
	}
	if result.Settings != `{"index.number_of_shards":"1"}` {
		t.Fatalf("unexpected settings: %s", result.Settings)
	}
}

func TestExistIndex(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/openapi_test_index/_exist" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"exist": true})
	}))
	defer server.Close()

	result, err := ExistIndex(cli, testRegion, &ExistIndexRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
	})
	if err != nil {
		t.Fatalf("exist index failed: %v", err)
	}
	if !result.Exist {
		t.Fatalf("expected exist to be true")
	}
}

func TestGetIndexStats(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/openapi_test_index/stats" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"stats": map[string]interface{}{"_shards": map[string]interface{}{"total": 1}},
		})
	}))
	defer server.Close()

	result, err := GetIndexStats(cli, testRegion, &GetIndexStatsRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
	})
	if err != nil {
		t.Fatalf("get index stats failed: %v", err)
	}
	if result.Stats == nil {
		t.Fatalf("expected stats to be present")
	}
}

func TestListIndexFieldTypes(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/index-field-types" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"fieldTypes": []map[string]interface{}{{
				"typeName":       "text",
				"isVector":       "false",
				"isAnalyzable":   "true",
				"isAggregatable": "false",
			}},
		})
	}))
	defer server.Close()

	result, err := ListIndexFieldTypes(cli, testRegion, &ListIndexFieldTypesRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("list index field types failed: %v", err)
	}
	if len(result.FieldTypes) != 1 || result.FieldTypes[0].TypeName != "text" {
		t.Fatalf("unexpected field types: %+v", result.FieldTypes)
	}
}

func TestOpenIndices(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/_open" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if names, _ := body["indexNames"].([]interface{}); len(names) != 2 {
			t.Fatalf("unexpected indexNames: %v", body["indexNames"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := OpenIndices(cli, testRegion, &IndexNamesRequest{
		ClusterId:  "search-xxxx",
		IndexNames: []string{"index-a", "index-b"},
	})
	if err != nil {
		t.Fatalf("open indices failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("expected success to be true")
	}
}

func TestCloseIndices(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/_close" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := CloseIndices(cli, testRegion, &IndexNamesRequest{
		ClusterId:  "search-xxxx",
		IndexNames: []string{"index-a"},
	})
	if err != nil {
		t.Fatalf("close indices failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("expected success to be true")
	}
}

func TestDeleteIndices(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/_delete" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := DeleteIndices(cli, testRegion, &IndexNamesRequest{
		ClusterId:  "search-xxxx",
		IndexNames: []string{"index-a"},
	})
	if err != nil {
		t.Fatalf("delete indices failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("expected success to be true")
	}
}

func TestRefreshIndices(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/_refresh" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := RefreshIndices(cli, testRegion, &IndexNamesRequest{
		ClusterId:  "search-xxxx",
		IndexNames: []string{"index-a"},
	})
	if err != nil {
		t.Fatalf("refresh indices failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("expected success to be true")
	}
}

func TestFlushIndices(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/_flush" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := FlushIndices(cli, testRegion, &IndexNamesRequest{
		ClusterId:  "search-xxxx",
		IndexNames: []string{"index-a"},
	})
	if err != nil {
		t.Fatalf("flush indices failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("expected success to be true")
	}
}

// TestForceMergeIndices populates both optional body fields so each of them is serialised.
func TestForceMergeIndices(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/_forcemerge" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if got := body["maxNumSegments"]; got != float64(1) {
			t.Fatalf("unexpected maxNumSegments: %v", got)
		}
		if got := body["onlyExpungeDeletes"]; got != true {
			t.Fatalf("unexpected onlyExpungeDeletes: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := ForceMergeIndices(cli, testRegion, &ForceMergeIndicesRequest{
		ClusterId:          "search-xxxx",
		IndexNames:         []string{"index-a", "index-b"},
		MaxNumSegments:     intPtr(1),
		OnlyExpungeDeletes: boolPtr(true),
	})
	if err != nil {
		t.Fatalf("force merge indices failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("expected success to be true")
	}
}

func TestClearIndicesCache(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/_clear_cache" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true})
	}))
	defer server.Close()

	result, err := ClearIndicesCache(cli, testRegion, &IndexNamesRequest{
		ClusterId:  "search-xxxx",
		IndexNames: []string{"index-a"},
	})
	if err != nil {
		t.Fatalf("clear indices cache failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("expected success to be true")
	}
}

func TestUpdateIndexSettings(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/openapi_test_index/settings" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"indexName": "openapi_test_index"})
	}))
	defer server.Close()

	result, err := UpdateIndexSettings(cli, testRegion, &UpdateIndexSettingsRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
		Settings:  map[string]interface{}{"index": map[string]interface{}{"refresh_interval": "1s"}},
	})
	if err != nil {
		t.Fatalf("update index settings failed: %v", err)
	}
	if result.IndexName != "openapi_test_index" {
		t.Fatalf("unexpected index name: %s", result.IndexName)
	}
}

// TestUpdateIndexMappings sets the optional type field so it is serialised in the body.
func TestUpdateIndexMappings(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/openapi_test_index/mappings" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if got := body["type"]; got != "_doc" {
			t.Fatalf("unexpected type: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"indexName": "openapi_test_index"})
	}))
	defer server.Close()

	result, err := UpdateIndexMappings(cli, testRegion, &UpdateIndexMappingsRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
		Type:      "_doc",
		Mappings: map[string]interface{}{
			"properties": map[string]interface{}{"summary": map[string]interface{}{"type": "text"}},
		},
	})
	if err != nil {
		t.Fatalf("update index mappings failed: %v", err)
	}
	if result.IndexName != "openapi_test_index" {
		t.Fatalf("unexpected index name: %s", result.IndexName)
	}
}

// TestUpdateIndexAliases populates every optional field of the alias action so all of them are
// serialised in the request body.
func TestUpdateIndexAliases(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/indices/openapi_test_index/aliases" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		actions, _ := body["actions"].([]interface{})
		if len(actions) != 1 {
			t.Fatalf("unexpected actions: %v", body["actions"])
		}
		action, _ := actions[0].(map[string]interface{})
		if action["action"] != "add" || action["alias"] != "openapi_test_alias" {
			t.Fatalf("unexpected action: %+v", action)
		}
		if action["isWriteIndex"] != true || action["indexRouting"] != "r1" {
			t.Fatalf("unexpected action: %+v", action)
		}
		writeJSONResponse(t, w, map[string]interface{}{"indexName": "openapi_test_index"})
	}))
	defer server.Close()

	result, err := UpdateIndexAliases(cli, testRegion, &UpdateIndexAliasesRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
		Actions: []AliasAction{{
			Action:        "add",
			Alias:         "openapi_test_alias",
			IsWriteIndex:  boolPtr(true),
			IndexRouting:  "r1",
			SearchRouting: "r2",
			Filter:        map[string]interface{}{"term": map[string]interface{}{"category": "default"}},
		}},
	})
	if err != nil {
		t.Fatalf("update index aliases failed: %v", err)
	}
	if result.IndexName != "openapi_test_index" {
		t.Fatalf("unexpected index name: %s", result.IndexName)
	}
}

func TestListIndexTemplates(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/index-templates" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"templates": []map[string]interface{}{{
				"templateName":  "tenant_template",
				"indexPatterns": `[".kibana_-*_*"]`,
				"priority":      "2147483647",
			}},
			"legacyTemplates": []map[string]interface{}{},
		})
	}))
	defer server.Close()

	result, err := ListIndexTemplates(cli, testRegion, &ListIndexTemplatesRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("list index templates failed: %v", err)
	}
	if len(result.Templates) != 1 || result.Templates[0].TemplateName != "tenant_template" {
		t.Fatalf("unexpected templates: %+v", result.Templates)
	}
}

func TestGetIndexTemplate(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/index-templates/openapi_test_template" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"templateName":  "openapi_test_template",
			"isLegacy":      "false",
			"type":          "default",
			"indexPatterns": `["openapi_test_*"]`,
			"priority":      "1",
			"template": map[string]interface{}{
				"settings": map[string]interface{}{"index": map[string]interface{}{"number_of_shards": "1"}},
			},
		})
	}))
	defer server.Close()

	result, err := GetIndexTemplate(cli, testRegion, &GetIndexTemplateRequest{
		ClusterId:    "search-xxxx",
		TemplateName: "openapi_test_template",
	})
	if err != nil {
		t.Fatalf("get index template failed: %v", err)
	}
	if result.TemplateName != "openapi_test_template" {
		t.Fatalf("unexpected template name: %s", result.TemplateName)
	}
	if result.IndexPatterns != `["openapi_test_*"]` {
		t.Fatalf("unexpected indexPatterns: %s", result.IndexPatterns)
	}
}

func TestExistIndexTemplate(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/index-templates/openapi_test_template/_exist" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"exist": true})
	}))
	defer server.Close()

	result, err := ExistIndexTemplate(cli, testRegion, &ExistIndexTemplateRequest{
		ClusterId:    "search-xxxx",
		TemplateName: "openapi_test_template",
	})
	if err != nil {
		t.Fatalf("exist index template failed: %v", err)
	}
	if !result.Exist {
		t.Fatalf("expected exist to be true")
	}
}

// TestCreateIndexTemplate populates every optional body field so each of them is serialised.
func TestCreateIndexTemplate(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/index-templates" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if got := body["templateName"]; got != "openapi_test_template" {
			t.Fatalf("unexpected templateName: %v", got)
		}
		if got := body["priority"]; got != "1" {
			t.Fatalf("unexpected priority: %v", got)
		}
		if _, ok := body["template"].(map[string]interface{}); !ok {
			t.Fatalf("expected template in body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"templateName": "openapi_test_template"})
	}))
	defer server.Close()

	result, err := CreateIndexTemplate(cli, testRegion, &CreateIndexTemplateRequest{
		ClusterId:     "search-xxxx",
		TemplateName:  "openapi_test_template",
		IndexPatterns: []string{"openapi_test_*"},
		Priority:      "1",
		Template: map[string]interface{}{
			"settings": map[string]interface{}{"index": map[string]interface{}{"number_of_shards": "1"}},
		},
	})
	if err != nil {
		t.Fatalf("create index template failed: %v", err)
	}
	if result.TemplateName != "openapi_test_template" {
		t.Fatalf("unexpected template name: %s", result.TemplateName)
	}
}

// TestUpdateIndexTemplate populates every optional body field; templateName is a path variable for
// the update API and must not appear in the body.
func TestUpdateIndexTemplate(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/index-templates/openapi_test_template" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if _, ok := body["templateName"]; ok {
			t.Fatalf("templateName should not be sent in update request body: %+v", body)
		}
		if got := body["priority"]; got != "1" {
			t.Fatalf("unexpected priority: %v", got)
		}
		if _, ok := body["template"].(map[string]interface{}); !ok {
			t.Fatalf("expected template in body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"templateName": "openapi_test_template"})
	}))
	defer server.Close()

	result, err := UpdateIndexTemplate(cli, testRegion, &UpdateIndexTemplateRequest{
		ClusterId:     "search-xxxx",
		TemplateName:  "openapi_test_template",
		IndexPatterns: []string{"openapi_test_*"},
		Priority:      "1",
		Template: map[string]interface{}{
			"settings": map[string]interface{}{"index": map[string]interface{}{"number_of_shards": "1"}},
		},
	})
	if err != nil {
		t.Fatalf("update index template failed: %v", err)
	}
	if result.TemplateName != "openapi_test_template" {
		t.Fatalf("unexpected template name: %s", result.TemplateName)
	}
}

// TestDeleteIndexTemplate sets the optional isLegacy query filter so its branch is exercised.
func TestDeleteIndexTemplate(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/index-templates/openapi_test_template" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("isLegacy"); got != "true" {
			t.Fatalf("unexpected isLegacy: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"templateName": "openapi_test_template"})
	}))
	defer server.Close()

	result, err := DeleteIndexTemplate(cli, testRegion, &DeleteIndexTemplateRequest{
		ClusterId:    "search-xxxx",
		TemplateName: "openapi_test_template",
		IsLegacy:     "true",
	})
	if err != nil {
		t.Fatalf("delete index template failed: %v", err)
	}
	if result.TemplateName != "openapi_test_template" {
		t.Fatalf("unexpected template name: %s", result.TemplateName)
	}
}

// TestDeleteIndexTemplateWithoutOptionalFields covers the request path where the optional isLegacy
// query parameter is omitted from the query string entirely.
func TestDeleteIndexTemplateWithoutOptionalFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.RawQuery != "" {
			t.Fatalf("unexpected query: %s", r.URL.RawQuery)
		}
		writeJSONResponse(t, w, map[string]interface{}{"templateName": "openapi_test_template"})
	}))
	defer server.Close()

	result, err := DeleteIndexTemplate(cli, testRegion, &DeleteIndexTemplateRequest{
		ClusterId:    "search-xxxx",
		TemplateName: "openapi_test_template",
	})
	if err != nil {
		t.Fatalf("delete index template failed: %v", err)
	}
	if result.TemplateName != "openapi_test_template" {
		t.Fatalf("unexpected template name: %s", result.TemplateName)
	}
}

func TestListIsmPolicies(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/ism-policies" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"policies": []map[string]interface{}{{
				"policyId":        "openapi_test_policy",
				"description":     "OpenAPI smoke policy",
				"lastUpdatedTime": "1784812307373",
			}},
		})
	}))
	defer server.Close()

	result, err := ListIsmPolicies(cli, testRegion, &ListIsmPoliciesRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("list ism policies failed: %v", err)
	}
	if len(result.Policies) != 1 || result.Policies[0].PolicyId != "openapi_test_policy" {
		t.Fatalf("unexpected policies: %+v", result.Policies)
	}
}

func TestGetIsmPolicy(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/ism-policies/openapi_test_policy" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"policyId":    "openapi_test_policy",
			"rawContent":  `{"policy_id":"openapi_test_policy"}`,
			"description": "OpenAPI smoke policy",
			"ismTemplate": []map[string]interface{}{
				{"indexPatterns": []string{"openapi_test_*"}, "priority": "1"},
			},
			"lastUpdatedTime": "1783567864780",
			"delete":          map[string]interface{}{"enabled": true, "minIndexAge": "30d"},
		})
	}))
	defer server.Close()

	result, err := GetIsmPolicy(cli, testRegion, &GetIsmPolicyRequest{
		ClusterId: "search-xxxx",
		PolicyId:  "openapi_test_policy",
	})
	if err != nil {
		t.Fatalf("get ism policy failed: %v", err)
	}
	if result.PolicyId != "openapi_test_policy" {
		t.Fatalf("unexpected policy id: %s", result.PolicyId)
	}
	if result.Delete.MinIndexAge != "30d" {
		t.Fatalf("unexpected delete.minIndexAge: %s", result.Delete.MinIndexAge)
	}
}

func TestExistIsmPolicy(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/ism-policies/openapi_test_policy/_exist" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"exist": true})
	}))
	defer server.Close()

	result, err := ExistIsmPolicy(cli, testRegion, &ExistIsmPolicyRequest{
		ClusterId: "search-xxxx",
		PolicyId:  "openapi_test_policy",
	})
	if err != nil {
		t.Fatalf("exist ism policy failed: %v", err)
	}
	if !result.Exist {
		t.Fatalf("expected exist to be true")
	}
}

// newFullIsmPolicyRequest returns an ISM policy request with every required and optional field set,
// so the create/update happy paths serialise all of the optional stage configurations.
func newFullIsmPolicyRequest() *IsmPolicyRequest {
	return &IsmPolicyRequest{
		ClusterId:   "search-xxxx",
		PolicyId:    "openapi_test_policy",
		Description: "OpenAPI smoke policy",
		IsmTemplate: []IsmTemplate{{IndexPatterns: []string{"openapi_test_*"}, Priority: "1"}},
		Rollover:    &IsmRollover{Enabled: boolPtr(true), MinIndexAge: "7d", MinDocCount: "1000", MinSize: "5gb"},
		Warm:        &IsmWarm{Enabled: boolPtr(true), MinIndexAge: "10d"},
		Cold:        &IsmCold{Enabled: boolPtr(true), MinIndexAge: "20d"},
		ForceMerge:  &IsmForceMerge{Enabled: boolPtr(true), MinIndexAge: "25d", MaxNumSegments: intPtr(1)},
		Delete:      &IsmDelete{Enabled: boolPtr(true), MinIndexAge: "30d"},
	}
}

func TestCreateIsmPolicy(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/ism-policies/openapi_test_policy" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if got := body["description"]; got != "OpenAPI smoke policy" {
			t.Fatalf("unexpected description: %v", got)
		}
		for _, stage := range []string{"rollover", "warm", "cold", "forceMerge", "delete"} {
			if _, ok := body[stage].(map[string]interface{}); !ok {
				t.Fatalf("expected %s in body: %+v", stage, body)
			}
		}
		writeJSONResponse(t, w, map[string]interface{}{"policyId": "openapi_test_policy"})
	}))
	defer server.Close()

	result, err := CreateIsmPolicy(cli, testRegion, newFullIsmPolicyRequest())
	if err != nil {
		t.Fatalf("create ism policy failed: %v", err)
	}
	if result.PolicyId != "openapi_test_policy" {
		t.Fatalf("unexpected policy id: %s", result.PolicyId)
	}
}

func TestUpdateIsmPolicy(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/ism-policies/openapi_test_policy" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if templates, _ := body["ismTemplate"].([]interface{}); len(templates) != 1 {
			t.Fatalf("unexpected ismTemplate: %v", body["ismTemplate"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"policyId": "openapi_test_policy"})
	}))
	defer server.Close()

	result, err := UpdateIsmPolicy(cli, testRegion, newFullIsmPolicyRequest())
	if err != nil {
		t.Fatalf("update ism policy failed: %v", err)
	}
	if result.PolicyId != "openapi_test_policy" {
		t.Fatalf("unexpected policy id: %s", result.PolicyId)
	}
}

func TestDeleteIsmPolicy(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxx/index-management/ism-policies/openapi_test_policy" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"policyId": "openapi_test_policy"})
	}))
	defer server.Close()

	result, err := DeleteIsmPolicy(cli, testRegion, &DeleteIsmPolicyRequest{
		ClusterId: "search-xxxx",
		PolicyId:  "openapi_test_policy",
	})
	if err != nil {
		t.Fatalf("delete ism policy failed: %v", err)
	}
	if result.PolicyId != "openapi_test_policy" {
		t.Fatalf("unexpected policy id: %s", result.PolicyId)
	}
}

// TestListMonitorIndices populates every optional query field so each conditional query-parameter
// branch in ListMonitorIndices is exercised.
func TestListMonitorIndices(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.Path != "/v3/clusters/search-xxxx/monitor/indices" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		query := r.URL.Query()
		expected := map[string]string{
			"indexName":     "openapi_test_*",
			"includeSystem": "false",
			"pageNo":        "1",
			"pageSize":      "100",
		}
		for key, want := range expected {
			if got := query.Get(key); got != want {
				t.Fatalf("unexpected %s: got %s want %s", key, got, want)
			}
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"indices": []map[string]interface{}{
				{"indexName": "openapi_test_index", "category": "default"},
				{"indexName": "top_queries-2026.07.09-44484", "category": "default"},
			},
			"pageNo":     1,
			"pageSize":   100,
			"totalCount": 2,
		})
	}))
	defer server.Close()

	result, err := ListMonitorIndices(cli, testRegion, &ListMonitorIndicesRequest{
		ClusterId:     "search-xxxx",
		IndexName:     "openapi_test_*",
		IncludeSystem: boolPtr(false),
		PageNo:        1,
		PageSize:      100,
	})
	if err != nil {
		t.Fatalf("list monitor indices failed: %v", err)
	}
	if len(result.Indices) != 2 || result.TotalCount != 2 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestListMonitorIndicesWithoutOptionalFields covers the request path where every optional query
// parameter is omitted from the query string entirely.
func TestListMonitorIndicesWithoutOptionalFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.URL.RawQuery != "" {
			t.Fatalf("unexpected query: %s", r.URL.RawQuery)
		}
		writeJSONResponse(t, w, map[string]interface{}{"totalCount": 0})
	}))
	defer server.Close()

	result, err := ListMonitorIndices(cli, testRegion, &ListMonitorIndicesRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("list monitor indices failed: %v", err)
	}
	if result.TotalCount != 0 || len(result.Indices) != 0 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestIndexAPIRejectsNilClient covers all 29 index functions: every one of them rejects a nil
// client with ErrNilClient before touching the request.
func TestIndexAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"CreateIndex": func() error { _, err := CreateIndex(nil, testRegion, &CreateIndexRequest{}); return err },
		"ListIndices": func() error { _, err := ListIndices(nil, testRegion, &ListIndicesRequest{}); return err },
		"GetIndex":    func() error { _, err := GetIndex(nil, testRegion, &GetIndexRequest{}); return err },
		"ExistIndex":  func() error { _, err := ExistIndex(nil, testRegion, &ExistIndexRequest{}); return err },
		"GetIndexStats": func() error {
			_, err := GetIndexStats(nil, testRegion, &GetIndexStatsRequest{})
			return err
		},
		"ListIndexFieldTypes": func() error {
			_, err := ListIndexFieldTypes(nil, testRegion, &ListIndexFieldTypesRequest{})
			return err
		},
		"OpenIndices":    func() error { _, err := OpenIndices(nil, testRegion, &IndexNamesRequest{}); return err },
		"CloseIndices":   func() error { _, err := CloseIndices(nil, testRegion, &IndexNamesRequest{}); return err },
		"DeleteIndices":  func() error { _, err := DeleteIndices(nil, testRegion, &IndexNamesRequest{}); return err },
		"RefreshIndices": func() error { _, err := RefreshIndices(nil, testRegion, &IndexNamesRequest{}); return err },
		"FlushIndices":   func() error { _, err := FlushIndices(nil, testRegion, &IndexNamesRequest{}); return err },
		"ForceMergeIndices": func() error {
			_, err := ForceMergeIndices(nil, testRegion, &ForceMergeIndicesRequest{})
			return err
		},
		"ClearIndicesCache": func() error {
			_, err := ClearIndicesCache(nil, testRegion, &IndexNamesRequest{})
			return err
		},
		"UpdateIndexSettings": func() error {
			_, err := UpdateIndexSettings(nil, testRegion, &UpdateIndexSettingsRequest{})
			return err
		},
		"UpdateIndexMappings": func() error {
			_, err := UpdateIndexMappings(nil, testRegion, &UpdateIndexMappingsRequest{})
			return err
		},
		"UpdateIndexAliases": func() error {
			_, err := UpdateIndexAliases(nil, testRegion, &UpdateIndexAliasesRequest{})
			return err
		},
		"ListIndexTemplates": func() error {
			_, err := ListIndexTemplates(nil, testRegion, &ListIndexTemplatesRequest{})
			return err
		},
		"GetIndexTemplate": func() error {
			_, err := GetIndexTemplate(nil, testRegion, &GetIndexTemplateRequest{})
			return err
		},
		"ExistIndexTemplate": func() error {
			_, err := ExistIndexTemplate(nil, testRegion, &ExistIndexTemplateRequest{})
			return err
		},
		"CreateIndexTemplate": func() error {
			_, err := CreateIndexTemplate(nil, testRegion, &CreateIndexTemplateRequest{})
			return err
		},
		"UpdateIndexTemplate": func() error {
			_, err := UpdateIndexTemplate(nil, testRegion, &UpdateIndexTemplateRequest{})
			return err
		},
		"DeleteIndexTemplate": func() error {
			_, err := DeleteIndexTemplate(nil, testRegion, &DeleteIndexTemplateRequest{})
			return err
		},
		"ListIsmPolicies": func() error {
			_, err := ListIsmPolicies(nil, testRegion, &ListIsmPoliciesRequest{})
			return err
		},
		"GetIsmPolicy":   func() error { _, err := GetIsmPolicy(nil, testRegion, &GetIsmPolicyRequest{}); return err },
		"ExistIsmPolicy": func() error { _, err := ExistIsmPolicy(nil, testRegion, &ExistIsmPolicyRequest{}); return err },
		"CreateIsmPolicy": func() error {
			_, err := CreateIsmPolicy(nil, testRegion, &IsmPolicyRequest{})
			return err
		},
		"UpdateIsmPolicy": func() error {
			_, err := UpdateIsmPolicy(nil, testRegion, &IsmPolicyRequest{})
			return err
		},
		"DeleteIsmPolicy": func() error {
			_, err := DeleteIsmPolicy(nil, testRegion, &DeleteIsmPolicyRequest{})
			return err
		},
		"ListMonitorIndices": func() error {
			_, err := ListMonitorIndices(nil, testRegion, &ListMonitorIndicesRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

// TestIndexAPIRejectsNilRequest covers all 29 index functions: none of them treats the request as
// optional, so every one must reject nil before any HTTP call is made.
func TestIndexAPIRejectsNilRequest(t *testing.T) {
	cli := newIndexRejectingClient(t)

	cases := map[string]func() error{
		"CreateIndex":         func() error { _, err := CreateIndex(cli, testRegion, nil); return err },
		"ListIndices":         func() error { _, err := ListIndices(cli, testRegion, nil); return err },
		"GetIndex":            func() error { _, err := GetIndex(cli, testRegion, nil); return err },
		"ExistIndex":          func() error { _, err := ExistIndex(cli, testRegion, nil); return err },
		"GetIndexStats":       func() error { _, err := GetIndexStats(cli, testRegion, nil); return err },
		"ListIndexFieldTypes": func() error { _, err := ListIndexFieldTypes(cli, testRegion, nil); return err },
		"OpenIndices":         func() error { _, err := OpenIndices(cli, testRegion, nil); return err },
		"CloseIndices":        func() error { _, err := CloseIndices(cli, testRegion, nil); return err },
		"DeleteIndices":       func() error { _, err := DeleteIndices(cli, testRegion, nil); return err },
		"RefreshIndices":      func() error { _, err := RefreshIndices(cli, testRegion, nil); return err },
		"FlushIndices":        func() error { _, err := FlushIndices(cli, testRegion, nil); return err },
		"ForceMergeIndices":   func() error { _, err := ForceMergeIndices(cli, testRegion, nil); return err },
		"ClearIndicesCache":   func() error { _, err := ClearIndicesCache(cli, testRegion, nil); return err },
		"UpdateIndexSettings": func() error { _, err := UpdateIndexSettings(cli, testRegion, nil); return err },
		"UpdateIndexMappings": func() error { _, err := UpdateIndexMappings(cli, testRegion, nil); return err },
		"UpdateIndexAliases":  func() error { _, err := UpdateIndexAliases(cli, testRegion, nil); return err },
		"ListIndexTemplates":  func() error { _, err := ListIndexTemplates(cli, testRegion, nil); return err },
		"GetIndexTemplate":    func() error { _, err := GetIndexTemplate(cli, testRegion, nil); return err },
		"ExistIndexTemplate":  func() error { _, err := ExistIndexTemplate(cli, testRegion, nil); return err },
		"CreateIndexTemplate": func() error { _, err := CreateIndexTemplate(cli, testRegion, nil); return err },
		"UpdateIndexTemplate": func() error { _, err := UpdateIndexTemplate(cli, testRegion, nil); return err },
		"DeleteIndexTemplate": func() error { _, err := DeleteIndexTemplate(cli, testRegion, nil); return err },
		"ListIsmPolicies":     func() error { _, err := ListIsmPolicies(cli, testRegion, nil); return err },
		"GetIsmPolicy":        func() error { _, err := GetIsmPolicy(cli, testRegion, nil); return err },
		"ExistIsmPolicy":      func() error { _, err := ExistIsmPolicy(cli, testRegion, nil); return err },
		"CreateIsmPolicy":     func() error { _, err := CreateIsmPolicy(cli, testRegion, nil); return err },
		"UpdateIsmPolicy":     func() error { _, err := UpdateIsmPolicy(cli, testRegion, nil); return err },
		"DeleteIsmPolicy":     func() error { _, err := DeleteIsmPolicy(cli, testRegion, nil); return err },
		"ListMonitorIndices":  func() error { _, err := ListMonitorIndices(cli, testRegion, nil); return err },
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

// TestIndexAPIRejectsMissingClusterId covers the five index functions whose only required field is
// clusterId. Optional filters are populated to prove the clusterId check runs first.
func TestIndexAPIRejectsMissingClusterId(t *testing.T) {
	cli := newIndexRejectingClient(t)

	cases := map[string]func() error{
		"ListIndices": func() error {
			_, err := ListIndices(cli, testRegion, &ListIndicesRequest{PageNo: 1, PageSize: 10})
			return err
		},
		"ListIndexFieldTypes": func() error {
			_, err := ListIndexFieldTypes(cli, testRegion, &ListIndexFieldTypesRequest{})
			return err
		},
		"ListIndexTemplates": func() error {
			_, err := ListIndexTemplates(cli, testRegion, &ListIndexTemplatesRequest{})
			return err
		},
		"ListIsmPolicies": func() error {
			_, err := ListIsmPolicies(cli, testRegion, &ListIsmPoliciesRequest{})
			return err
		},
		"ListMonitorIndices": func() error {
			_, err := ListMonitorIndices(cli, testRegion, &ListMonitorIndicesRequest{PageNo: 1, PageSize: 100})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestCreateIndexRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := CreateIndexRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
		Settings:  map[string]interface{}{"index": map[string]interface{}{"number_of_shards": "1"}},
	}
	cases := map[string]func(*CreateIndexRequest){
		"clusterId": func(r *CreateIndexRequest) { r.ClusterId = "" },
		"indexName": func(r *CreateIndexRequest) { r.IndexName = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := CreateIndex(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestGetIndexRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := GetIndexRequest{ClusterId: "search-xxxx", IndexName: "openapi_test_index"}
	cases := map[string]func(*GetIndexRequest){
		"clusterId": func(r *GetIndexRequest) { r.ClusterId = "" },
		"indexName": func(r *GetIndexRequest) { r.IndexName = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := GetIndex(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestExistIndexRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := ExistIndexRequest{ClusterId: "search-xxxx", IndexName: "openapi_test_index"}
	cases := map[string]func(*ExistIndexRequest){
		"clusterId": func(r *ExistIndexRequest) { r.ClusterId = "" },
		"indexName": func(r *ExistIndexRequest) { r.IndexName = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := ExistIndex(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestGetIndexStatsRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := GetIndexStatsRequest{ClusterId: "search-xxxx", IndexName: "openapi_test_index"}
	cases := map[string]func(*GetIndexStatsRequest){
		"clusterId": func(r *GetIndexStatsRequest) { r.ClusterId = "" },
		"indexName": func(r *GetIndexStatsRequest) { r.IndexName = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := GetIndexStats(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestIndexNamesAPIRejectsMissingRequiredFields covers the six batch index operations that share
// IndexNamesRequest: clusterId and a non-empty indexNames list are required by all of them.
func TestIndexNamesAPIRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	operations := map[string]func(*IndexNamesRequest) error{
		"OpenIndices":       func(r *IndexNamesRequest) error { _, err := OpenIndices(cli, testRegion, r); return err },
		"CloseIndices":      func(r *IndexNamesRequest) error { _, err := CloseIndices(cli, testRegion, r); return err },
		"DeleteIndices":     func(r *IndexNamesRequest) error { _, err := DeleteIndices(cli, testRegion, r); return err },
		"RefreshIndices":    func(r *IndexNamesRequest) error { _, err := RefreshIndices(cli, testRegion, r); return err },
		"FlushIndices":      func(r *IndexNamesRequest) error { _, err := FlushIndices(cli, testRegion, r); return err },
		"ClearIndicesCache": func(r *IndexNamesRequest) error { _, err := ClearIndicesCache(cli, testRegion, r); return err },
	}
	full := IndexNamesRequest{ClusterId: "search-xxxx", IndexNames: []string{"index-a"}}
	cases := map[string]func(*IndexNamesRequest){
		"clusterId":  func(r *IndexNamesRequest) { r.ClusterId = "" },
		"indexNames": func(r *IndexNamesRequest) { r.IndexNames = nil },
	}
	for operation, call := range operations {
		for name, clear := range cases {
			request := full
			clear(&request)
			if err := call(&request); err == nil {
				t.Fatalf("%s/%s: expected validation error", operation, name)
			}
		}
	}
}

func TestForceMergeIndicesRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := ForceMergeIndicesRequest{
		ClusterId:      "search-xxxx",
		IndexNames:     []string{"index-a"},
		MaxNumSegments: intPtr(1),
	}
	cases := map[string]func(*ForceMergeIndicesRequest){
		"clusterId":  func(r *ForceMergeIndicesRequest) { r.ClusterId = "" },
		"indexNames": func(r *ForceMergeIndicesRequest) { r.IndexNames = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := ForceMergeIndices(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateIndexSettingsRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := UpdateIndexSettingsRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
		Settings:  map[string]interface{}{"index": map[string]interface{}{"refresh_interval": "1s"}},
	}
	cases := map[string]func(*UpdateIndexSettingsRequest){
		"clusterId": func(r *UpdateIndexSettingsRequest) { r.ClusterId = "" },
		"indexName": func(r *UpdateIndexSettingsRequest) { r.IndexName = "" },
		"settings":  func(r *UpdateIndexSettingsRequest) { r.Settings = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateIndexSettings(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateIndexMappingsRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := UpdateIndexMappingsRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
		Type:      "_doc",
		Mappings:  map[string]interface{}{"properties": map[string]interface{}{}},
	}
	cases := map[string]func(*UpdateIndexMappingsRequest){
		"clusterId": func(r *UpdateIndexMappingsRequest) { r.ClusterId = "" },
		"indexName": func(r *UpdateIndexMappingsRequest) { r.IndexName = "" },
		"mappings":  func(r *UpdateIndexMappingsRequest) { r.Mappings = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateIndexMappings(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateIndexAliasesRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := UpdateIndexAliasesRequest{
		ClusterId: "search-xxxx",
		IndexName: "openapi_test_index",
		Actions:   []AliasAction{{Action: "add", Alias: "openapi_test_alias"}},
	}
	cases := map[string]func(*UpdateIndexAliasesRequest){
		"clusterId": func(r *UpdateIndexAliasesRequest) { r.ClusterId = "" },
		"indexName": func(r *UpdateIndexAliasesRequest) { r.IndexName = "" },
		"actions":   func(r *UpdateIndexAliasesRequest) { r.Actions = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateIndexAliases(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestGetIndexTemplateRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := GetIndexTemplateRequest{ClusterId: "search-xxxx", TemplateName: "openapi_test_template"}
	cases := map[string]func(*GetIndexTemplateRequest){
		"clusterId":    func(r *GetIndexTemplateRequest) { r.ClusterId = "" },
		"templateName": func(r *GetIndexTemplateRequest) { r.TemplateName = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := GetIndexTemplate(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestExistIndexTemplateRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := ExistIndexTemplateRequest{ClusterId: "search-xxxx", TemplateName: "openapi_test_template"}
	cases := map[string]func(*ExistIndexTemplateRequest){
		"clusterId":    func(r *ExistIndexTemplateRequest) { r.ClusterId = "" },
		"templateName": func(r *ExistIndexTemplateRequest) { r.TemplateName = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := ExistIndexTemplate(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestCreateIndexTemplateRejectsMissingRequiredFields also proves templateName is optional for the
// create API: only clusterId and indexPatterns are validated.
func TestCreateIndexTemplateRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := CreateIndexTemplateRequest{
		ClusterId:     "search-xxxx",
		TemplateName:  "openapi_test_template",
		IndexPatterns: []string{"openapi_test_*"},
	}
	cases := map[string]func(*CreateIndexTemplateRequest){
		"clusterId":     func(r *CreateIndexTemplateRequest) { r.ClusterId = "" },
		"indexPatterns": func(r *CreateIndexTemplateRequest) { r.IndexPatterns = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := CreateIndexTemplate(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateIndexTemplateRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := UpdateIndexTemplateRequest{
		ClusterId:     "search-xxxx",
		TemplateName:  "openapi_test_template",
		IndexPatterns: []string{"openapi_test_*"},
	}
	cases := map[string]func(*UpdateIndexTemplateRequest){
		"clusterId":     func(r *UpdateIndexTemplateRequest) { r.ClusterId = "" },
		"templateName":  func(r *UpdateIndexTemplateRequest) { r.TemplateName = "" },
		"indexPatterns": func(r *UpdateIndexTemplateRequest) { r.IndexPatterns = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateIndexTemplate(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestDeleteIndexTemplateRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := DeleteIndexTemplateRequest{
		ClusterId:    "search-xxxx",
		TemplateName: "openapi_test_template",
		IsLegacy:     "true",
	}
	cases := map[string]func(*DeleteIndexTemplateRequest){
		"clusterId":    func(r *DeleteIndexTemplateRequest) { r.ClusterId = "" },
		"templateName": func(r *DeleteIndexTemplateRequest) { r.TemplateName = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := DeleteIndexTemplate(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestGetIsmPolicyRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := GetIsmPolicyRequest{ClusterId: "search-xxxx", PolicyId: "openapi_test_policy"}
	cases := map[string]func(*GetIsmPolicyRequest){
		"clusterId": func(r *GetIsmPolicyRequest) { r.ClusterId = "" },
		"policyId":  func(r *GetIsmPolicyRequest) { r.PolicyId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := GetIsmPolicy(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestExistIsmPolicyRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := ExistIsmPolicyRequest{ClusterId: "search-xxxx", PolicyId: "openapi_test_policy"}
	cases := map[string]func(*ExistIsmPolicyRequest){
		"clusterId": func(r *ExistIsmPolicyRequest) { r.ClusterId = "" },
		"policyId":  func(r *ExistIsmPolicyRequest) { r.PolicyId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := ExistIsmPolicy(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestIsmPolicyAPIRejectsMissingRequiredFields covers create and update ISM policy, which share
// IsmPolicyRequest and validate the same four fields.
func TestIsmPolicyAPIRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	operations := map[string]func(*IsmPolicyRequest) error{
		"CreateIsmPolicy": func(r *IsmPolicyRequest) error { _, err := CreateIsmPolicy(cli, testRegion, r); return err },
		"UpdateIsmPolicy": func(r *IsmPolicyRequest) error { _, err := UpdateIsmPolicy(cli, testRegion, r); return err },
	}
	cases := map[string]func(*IsmPolicyRequest){
		"clusterId":   func(r *IsmPolicyRequest) { r.ClusterId = "" },
		"policyId":    func(r *IsmPolicyRequest) { r.PolicyId = "" },
		"description": func(r *IsmPolicyRequest) { r.Description = "" },
		"ismTemplate": func(r *IsmPolicyRequest) { r.IsmTemplate = nil },
	}
	for operation, call := range operations {
		for name, clear := range cases {
			request := newFullIsmPolicyRequest()
			clear(request)
			if err := call(request); err == nil {
				t.Fatalf("%s/%s: expected validation error", operation, name)
			}
		}
	}
}

func TestDeleteIsmPolicyRejectsMissingRequiredFields(t *testing.T) {
	cli := newIndexRejectingClient(t)

	full := DeleteIsmPolicyRequest{ClusterId: "search-xxxx", PolicyId: "openapi_test_policy"}
	cases := map[string]func(*DeleteIsmPolicyRequest){
		"clusterId": func(r *DeleteIsmPolicyRequest) { r.ClusterId = "" },
		"policyId":  func(r *DeleteIsmPolicyRequest) { r.PolicyId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := DeleteIsmPolicy(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestIndexAPIPropagatesServerError uses 400 rather than 500 on purpose: the helper mirrors the
// production retry policy, which retries 5xx with backoff and would make this test take seconds.
func TestIndexAPIPropagatesServerError(t *testing.T) {
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

	indexNames := &IndexNamesRequest{ClusterId: "search-xxxx", IndexNames: []string{"index-a"}}
	cases := map[string]func() error{
		"CreateIndex": func() error {
			_, err := CreateIndex(cli, testRegion, &CreateIndexRequest{
				ClusterId: "search-xxxx",
				IndexName: "openapi_test_index",
			})
			return err
		},
		"ListIndices": func() error {
			_, err := ListIndices(cli, testRegion, &ListIndicesRequest{ClusterId: "search-xxxx"})
			return err
		},
		"GetIndex": func() error {
			_, err := GetIndex(cli, testRegion, &GetIndexRequest{
				ClusterId: "search-xxxx",
				IndexName: "openapi_test_index",
			})
			return err
		},
		"ExistIndex": func() error {
			_, err := ExistIndex(cli, testRegion, &ExistIndexRequest{
				ClusterId: "search-xxxx",
				IndexName: "openapi_test_index",
			})
			return err
		},
		"GetIndexStats": func() error {
			_, err := GetIndexStats(cli, testRegion, &GetIndexStatsRequest{
				ClusterId: "search-xxxx",
				IndexName: "openapi_test_index",
			})
			return err
		},
		"ListIndexFieldTypes": func() error {
			_, err := ListIndexFieldTypes(cli, testRegion, &ListIndexFieldTypesRequest{ClusterId: "search-xxxx"})
			return err
		},
		"OpenIndices":    func() error { _, err := OpenIndices(cli, testRegion, indexNames); return err },
		"CloseIndices":   func() error { _, err := CloseIndices(cli, testRegion, indexNames); return err },
		"DeleteIndices":  func() error { _, err := DeleteIndices(cli, testRegion, indexNames); return err },
		"RefreshIndices": func() error { _, err := RefreshIndices(cli, testRegion, indexNames); return err },
		"FlushIndices":   func() error { _, err := FlushIndices(cli, testRegion, indexNames); return err },
		"ClearIndicesCache": func() error {
			_, err := ClearIndicesCache(cli, testRegion, indexNames)
			return err
		},
		"ForceMergeIndices": func() error {
			_, err := ForceMergeIndices(cli, testRegion, &ForceMergeIndicesRequest{
				ClusterId:  "search-xxxx",
				IndexNames: []string{"index-a"},
			})
			return err
		},
		"UpdateIndexSettings": func() error {
			_, err := UpdateIndexSettings(cli, testRegion, &UpdateIndexSettingsRequest{
				ClusterId: "search-xxxx",
				IndexName: "openapi_test_index",
				Settings:  map[string]interface{}{"index": map[string]interface{}{"refresh_interval": "1s"}},
			})
			return err
		},
		"UpdateIndexMappings": func() error {
			_, err := UpdateIndexMappings(cli, testRegion, &UpdateIndexMappingsRequest{
				ClusterId: "search-xxxx",
				IndexName: "openapi_test_index",
				Mappings:  map[string]interface{}{"properties": map[string]interface{}{}},
			})
			return err
		},
		"UpdateIndexAliases": func() error {
			_, err := UpdateIndexAliases(cli, testRegion, &UpdateIndexAliasesRequest{
				ClusterId: "search-xxxx",
				IndexName: "openapi_test_index",
				Actions:   []AliasAction{{Action: "add", Alias: "openapi_test_alias"}},
			})
			return err
		},
		"ListIndexTemplates": func() error {
			_, err := ListIndexTemplates(cli, testRegion, &ListIndexTemplatesRequest{ClusterId: "search-xxxx"})
			return err
		},
		"GetIndexTemplate": func() error {
			_, err := GetIndexTemplate(cli, testRegion, &GetIndexTemplateRequest{
				ClusterId:    "search-xxxx",
				TemplateName: "openapi_test_template",
			})
			return err
		},
		"ExistIndexTemplate": func() error {
			_, err := ExistIndexTemplate(cli, testRegion, &ExistIndexTemplateRequest{
				ClusterId:    "search-xxxx",
				TemplateName: "openapi_test_template",
			})
			return err
		},
		"CreateIndexTemplate": func() error {
			_, err := CreateIndexTemplate(cli, testRegion, &CreateIndexTemplateRequest{
				ClusterId:     "search-xxxx",
				TemplateName:  "openapi_test_template",
				IndexPatterns: []string{"openapi_test_*"},
			})
			return err
		},
		"UpdateIndexTemplate": func() error {
			_, err := UpdateIndexTemplate(cli, testRegion, &UpdateIndexTemplateRequest{
				ClusterId:     "search-xxxx",
				TemplateName:  "openapi_test_template",
				IndexPatterns: []string{"openapi_test_*"},
			})
			return err
		},
		"DeleteIndexTemplate": func() error {
			_, err := DeleteIndexTemplate(cli, testRegion, &DeleteIndexTemplateRequest{
				ClusterId:    "search-xxxx",
				TemplateName: "openapi_test_template",
			})
			return err
		},
		"ListIsmPolicies": func() error {
			_, err := ListIsmPolicies(cli, testRegion, &ListIsmPoliciesRequest{ClusterId: "search-xxxx"})
			return err
		},
		"GetIsmPolicy": func() error {
			_, err := GetIsmPolicy(cli, testRegion, &GetIsmPolicyRequest{
				ClusterId: "search-xxxx",
				PolicyId:  "openapi_test_policy",
			})
			return err
		},
		"ExistIsmPolicy": func() error {
			_, err := ExistIsmPolicy(cli, testRegion, &ExistIsmPolicyRequest{
				ClusterId: "search-xxxx",
				PolicyId:  "openapi_test_policy",
			})
			return err
		},
		"CreateIsmPolicy": func() error {
			_, err := CreateIsmPolicy(cli, testRegion, newFullIsmPolicyRequest())
			return err
		},
		"UpdateIsmPolicy": func() error {
			_, err := UpdateIsmPolicy(cli, testRegion, newFullIsmPolicyRequest())
			return err
		},
		"DeleteIsmPolicy": func() error {
			_, err := DeleteIsmPolicy(cli, testRegion, &DeleteIsmPolicyRequest{
				ClusterId: "search-xxxx",
				PolicyId:  "openapi_test_policy",
			})
			return err
		},
		"ListMonitorIndices": func() error {
			_, err := ListMonitorIndices(cli, testRegion, &ListMonitorIndicesRequest{ClusterId: "search-xxxx"})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for 400 response", name)
		}
	}
}

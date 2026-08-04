// Copyright 2026 Baidu, Inc. All rights reserved.
// Based on Hindsight Go Client - MIT License

package dumemory_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/dumemory/api"
)

func TestEntityScopeTags(t *testing.T) {
	scope := dumemory.EntityScope{UserID: "u1", AgentID: "a1", AppID: "app1", RunID: "r1"}
	tags, err := scope.Tags()
	if err != nil {
		t.Fatalf("Tags: %v", err)
	}
	want := []string{"user_id:u1", "agent_id:a1", "app_id:app1", "run_id:r1"}
	if len(tags) != len(want) {
		t.Fatalf("tags length = %d, want %d: %#v", len(tags), len(want), tags)
	}
	for i := range want {
		if tags[i] != want[i] {
			t.Fatalf("tags[%d] = %q, want %q", i, tags[i], want[i])
		}
	}
}

func TestEntityScopeRequiresAtLeastOneID(t *testing.T) {
	_, err := (dumemory.EntityScope{}).Tags()
	if !errors.Is(err, dumemory.ErrMissingEntityScope) {
		t.Fatalf("Tags error = %v, want ErrMissingEntityScope", err)
	}
}

func TestScopedAPIsApplyTagsAndDefaultTagsMatch(t *testing.T) {
	ctx := context.Background()
	scope := dumemory.EntityScope{UserID: "u1", AppID: "app1"}
	wantTags := []string{"topic:coffee", "user_id:u1", "app_id:app1"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/default/banks/bank/memories":
			if r.Method != http.MethodPost {
				t.Fatalf("retain method = %s", r.Method)
			}
			var body map[string]interface{}
			decodeJSONBody(t, r, &body)
			items := body["items"].([]interface{})
			item := items[0].(map[string]interface{})
			assertJSONTags(t, item["tags"], wantTags)
			assertJSONTags(t, body["document_tags"], []string{"doc:batch", "user_id:u1", "app_id:app1"})
			_, _ = w.Write([]byte(`{"success":true,"bank_id":"bank","items_count":1,"async":false}`))
		case "/v1/default/banks/bank/memories/recall":
			if r.Method != http.MethodPost {
				t.Fatalf("recall method = %s", r.Method)
			}
			var body map[string]interface{}
			decodeJSONBody(t, r, &body)
			assertJSONTags(t, body["tags"], wantTags)
			if body["tags_match"] != "all_strict" {
				t.Fatalf("recall tags_match = %#v, want all_strict", body["tags_match"])
			}
			_, _ = w.Write([]byte(`{"results":[]}`))
		case "/v1/default/banks/bank/reflect":
			if r.Method != http.MethodPost {
				t.Fatalf("reflect method = %s", r.Method)
			}
			var body map[string]interface{}
			decodeJSONBody(t, r, &body)
			assertJSONTags(t, body["tags"], wantTags)
			if body["tags_match"] != "all_strict" {
				t.Fatalf("reflect tags_match = %#v, want all_strict", body["tags_match"])
			}
			_, _ = w.Write([]byte(`{"text":"ok"}`))
		case "/v1/default/banks/bank/memories/list":
			if r.Method != http.MethodGet {
				t.Fatalf("memories list method = %s", r.Method)
			}
			query := r.URL.Query()
			if query.Get("tags_match") != "all_strict" {
				t.Fatalf("memories list tags_match = %q, want all_strict", query.Get("tags_match"))
			}
			assertQueryTags(t, r, wantTags)
			if query.Get("state") != "active" || query.Get("document_id") != "doc-1" || query.Get("entity_id") != "entity-1" {
				t.Fatalf("memory filters = %v", query)
			}
			_, _ = w.Write([]byte(`{"items":[],"total":0,"limit":10,"offset":0}`))
		case "/v1/default/banks/bank/documents":
			if r.Method != http.MethodGet {
				t.Fatalf("documents method = %s", r.Method)
			}
			assertQueryTags(t, r, wantTags)
			if got := r.URL.Query().Get("tags_match"); got != "all_strict" {
				t.Fatalf("documents tags_match = %q, want all_strict", got)
			}
			_, _ = w.Write([]byte(`{"items":[],"total":0,"limit":10,"offset":0}`))
		case "/v1/default/banks/bank/documents/doc1":
			if r.Method != http.MethodPatch {
				t.Fatalf("update document method = %s", r.Method)
			}
			var body map[string]interface{}
			decodeJSONBody(t, r, &body)
			assertJSONTags(t, body["tags"], wantTags)
			_, _ = w.Write([]byte(`{"success":true}`))
		case "/v1/default/banks/bank/directives":
			switch r.Method {
			case http.MethodGet:
				assertQueryTags(t, r, wantTags)
				if got := r.URL.Query().Get("tags_match"); got != "all_strict" {
					t.Fatalf("directives tags_match = %q, want all_strict", got)
				}
				_, _ = w.Write([]byte(`{"items":[]}`))
			case http.MethodPost:
				var body map[string]interface{}
				decodeJSONBody(t, r, &body)
				assertJSONTags(t, body["tags"], wantTags)
				_, _ = w.Write([]byte(`{"id":"dir1","bank_id":"bank","name":"n","content":"c"}`))
			default:
				t.Fatalf("directives method = %s", r.Method)
			}
		case "/v1/default/banks/bank/directives/dir1":
			if r.Method != http.MethodPatch {
				t.Fatalf("update directive method = %s", r.Method)
			}
			var body map[string]interface{}
			decodeJSONBody(t, r, &body)
			assertJSONTags(t, body["tags"], wantTags)
			_, _ = w.Write([]byte(`{"id":"dir1","bank_id":"bank","name":"n","content":"c"}`))
		case "/v1/default/banks/bank/mental-models":
			switch r.Method {
			case http.MethodGet:
				assertQueryTags(t, r, wantTags)
				if got := r.URL.Query().Get("tags_match"); got != "all_strict" {
					t.Fatalf("mental models tags_match = %q, want all_strict", got)
				}
				_, _ = w.Write([]byte(`{"items":[]}`))
			case http.MethodPost:
				var body map[string]interface{}
				decodeJSONBody(t, r, &body)
				assertJSONTags(t, body["tags"], wantTags)
				_, _ = w.Write([]byte(`{"mental_model_id":"mm1","operation_id":"op1"}`))
			default:
				t.Fatalf("mental models method = %s", r.Method)
			}
		case "/v1/default/banks/bank/mental-models/mm1":
			if r.Method != http.MethodPatch {
				t.Fatalf("update mental model method = %s", r.Method)
			}
			var body map[string]interface{}
			decodeJSONBody(t, r, &body)
			assertJSONTags(t, body["tags"], wantTags)
			_, _ = w.Write([]byte(`{"id":"mm1","bank_id":"bank","name":"m"}`))
		case "/v1/default/banks/bank/tags":
			if r.Method != http.MethodGet {
				t.Fatalf("tags method = %s", r.Method)
			}
			if got := r.URL.Query().Get("q"); got != "user_id:u1*" {
				t.Fatalf("tags q = %q, want user_id:u1*", got)
			}
			if got := r.URL.Query().Get("source"); got != "memories" {
				t.Fatalf("tags source = %q, want memories", got)
			}
			_, _ = w.Write([]byte(`{"items":[{"tag":"user_id:u1","count":1}],"total":1,"limit":20,"offset":0}`))
		case "/v1/default/banks/bank/memories/m1":
			if r.Method != http.MethodGet {
				t.Fatalf("get memory method = %s", r.Method)
			}
			_, _ = w.Write([]byte(`{"id":"m1","tags":["user_id:u1"]}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.String())
		}
	}))
	defer server.Close()

	client := dumemory.New(server.URL, "test-token")

	item := *dumemory.NewMemoryItem("hello")
	item.Tags = []string{"topic:coffee"}
	retainReq := *dumemory.NewRetainRequest([]dumemory.MemoryItem{item})
	retainReq.DocumentTags = []string{"doc:batch"}
	if _, err := client.RetainWithScope(ctx, "bank", scope, retainReq); err != nil {
		t.Fatalf("RetainWithScope: %v", err)
	}

	recallReq := *dumemory.NewRecallRequest("coffee")
	recallReq.Tags = []string{"topic:coffee"}
	if _, err := client.RecallWithScope(ctx, "bank", scope, recallReq); err != nil {
		t.Fatalf("RecallWithScope: %v", err)
	}

	reflectReq := *dumemory.NewReflectRequest("coffee")
	reflectReq.Tags = []string{"topic:coffee"}
	if _, err := client.ReflectWithScope(ctx, "bank", scope, reflectReq); err != nil {
		t.Fatalf("ReflectWithScope: %v", err)
	}

	if _, err := client.ListMemoriesWithScope(ctx, "bank", scope, dumemory.ListMemoriesOptions{
		Q:          "coffee",
		State:      "active",
		DocumentID: "doc-1",
		EntityID:   "entity-1",
		Tags:       []string{"topic:coffee"},
		Limit:      10,
	}); err != nil {
		t.Fatalf("ListMemoriesWithScope: %v", err)
	}

	if _, err := client.ListDocumentsWithScope(ctx, "bank", scope, dumemory.ListDocumentsOptions{Tags: []string{"topic:coffee"}, Limit: 10}); err != nil {
		t.Fatalf("ListDocumentsWithScope: %v", err)
	}

	updateDocReq := *dumemory.NewUpdateDocumentRequest()
	updateDocReq.Tags = []string{"topic:coffee"}
	if _, err := client.UpdateDocumentTagsWithScope(ctx, "bank", "doc1", scope, updateDocReq); err != nil {
		t.Fatalf("UpdateDocumentTagsWithScope: %v", err)
	}

	if _, err := client.ListDirectivesWithScope(ctx, "bank", scope, dumemory.ListDirectivesOptions{Tags: []string{"topic:coffee"}}); err != nil {
		t.Fatalf("ListDirectivesWithScope: %v", err)
	}

	createDirectiveReq := *dumemory.NewCreateDirectiveRequest("n", "c")
	createDirectiveReq.Tags = []string{"topic:coffee"}
	if _, err := client.CreateDirectiveWithScope(ctx, "bank", scope, createDirectiveReq); err != nil {
		t.Fatalf("CreateDirectiveWithScope: %v", err)
	}

	updateDirectiveReq := *dumemory.NewUpdateDirectiveRequest()
	updateDirectiveReq.Tags = []string{"topic:coffee"}
	if _, err := client.UpdateDirectiveWithScope(ctx, "bank", "dir1", scope, updateDirectiveReq); err != nil {
		t.Fatalf("UpdateDirectiveWithScope: %v", err)
	}

	if _, err := client.ListMentalModelsWithScope(ctx, "bank", scope, dumemory.ListMentalModelsOptions{Tags: []string{"topic:coffee"}}); err != nil {
		t.Fatalf("ListMentalModelsWithScope: %v", err)
	}

	createMentalModelReq := *dumemory.NewCreateMentalModelRequest("m", "coffee")
	createMentalModelReq.Tags = []string{"topic:coffee"}
	if _, err := client.CreateMentalModelWithScope(ctx, "bank", scope, createMentalModelReq); err != nil {
		t.Fatalf("CreateMentalModelWithScope: %v", err)
	}

	updateMentalModelReq := *dumemory.NewUpdateMentalModelRequest()
	updateMentalModelReq.Tags = []string{"topic:coffee"}
	if _, err := client.UpdateMentalModelWithScope(ctx, "bank", "mm1", scope, updateMentalModelReq); err != nil {
		t.Fatalf("UpdateMentalModelWithScope: %v", err)
	}

	if _, err := client.ListTagsWithScope(ctx, "bank", scope, dumemory.ListTagsOptions{Source: "memories", Limit: 20}); err != nil {
		t.Fatalf("ListTagsWithScope: %v", err)
	}

	if _, err := client.GetMemoryWithScope(ctx, "bank", "m1", scope); err != nil {
		t.Fatalf("GetMemoryWithScope: %v", err)
	}
}

func TestScopedAPIsPreserveExplicitTagsMatch(t *testing.T) {
	ctx := context.Background()
	scope := dumemory.EntityScope{UserID: "u1"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path != "/v1/default/banks/bank/documents" {
			t.Fatalf("unexpected path: %s", r.URL.String())
		}
		if got := r.URL.Query().Get("tags_match"); got != "exact" {
			t.Fatalf("tags_match = %q, want exact", got)
		}
		_, _ = w.Write([]byte(`{"items":[],"total":0,"limit":0,"offset":0}`))
	}))
	defer server.Close()

	client := dumemory.New(server.URL, "test-token")
	_, err := client.ListDocumentsWithScope(ctx, "bank", scope, dumemory.ListDocumentsOptions{TagsMatch: "exact"})
	if err != nil {
		t.Fatalf("ListDocumentsWithScope: %v", err)
	}
}

func decodeJSONBody(t *testing.T, r *http.Request, out interface{}) {
	t.Helper()
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(out); err != nil {
		t.Fatalf("decode request body: %v", err)
	}
}

func assertQueryTags(t *testing.T, r *http.Request, want []string) {
	t.Helper()
	got := r.URL.Query()["tags"]
	assertStringSet(t, got, want)
}

func assertJSONTags(t *testing.T, value interface{}, want []string) {
	t.Helper()
	items, ok := value.([]interface{})
	if !ok {
		t.Fatalf("tags = %#v, want JSON array", value)
	}
	got := make([]string, 0, len(items))
	for _, item := range items {
		got = append(got, item.(string))
	}
	assertStringSet(t, got, want)
}

func assertStringSet(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("tags = %#v, want %#v", got, want)
	}
	seen := make(map[string]int, len(got))
	for _, tag := range got {
		seen[tag]++
	}
	for _, tag := range want {
		if seen[tag] != 1 {
			t.Fatalf("tags = %#v, want exactly one %q", got, tag)
		}
	}
}

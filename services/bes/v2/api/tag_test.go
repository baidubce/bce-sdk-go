package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

// tagErrorServerClient builds a client whose server always answers with HTTP 400 and a BCE
// service error body. 400 is used on purpose: newTestClient mirrors the production
// DEFAULT_RETRY_POLICY, so a 5xx status would trigger backoff retries and slow the package down.
func tagErrorServerClient(t *testing.T) bce.Client {
	t.Helper()
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(nethttp.StatusBadRequest)
		if _, err := w.Write([]byte(`{"code":"BadRequest","message":"tag request rejected","requestId":"tag-req-id"}`)); err != nil {
			t.Errorf("write error response failed: %v", err)
		}
	}))
	t.Cleanup(server.Close)
	return cli
}

func TestListTags(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if r.Method != nethttp.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/bes/cluster/tagList" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get(HEADER_REGION); got != testRegion {
			t.Errorf("unexpected region header: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  []map[string]interface{}{{"tagKey": "key", "tagValue": "value"}},
		})
	}))
	defer server.Close()

	result, err := ListTags(cli, testRegion)
	if err != nil {
		t.Fatalf("list tags failed: %v", err)
	}
	if len(result.Result) != 1 || result.Result[0].TagKey != "key" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestUpdateClusterTags(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if r.Method != nethttp.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/bes/cluster/updateTags" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get(HEADER_REGION); got != testRegion {
			t.Errorf("unexpected region header: %s", got)
		}
		body := readJSONBody(t, r)
		if got := body["clusterId"]; got != "1111111" {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		tags, ok := body["tags"].([]interface{})
		if !ok || len(tags) != 1 {
			t.Fatalf("unexpected tags: %v", body["tags"])
		}
		if tag, ok := tags[0].(map[string]interface{}); !ok || tag["tagKey"] != "key" || tag["tagValue"] != "value" {
			t.Fatalf("unexpected tag entry: %v", tags[0])
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": map[string]interface{}{}})
	}))
	defer server.Close()

	result, err := UpdateClusterTags(cli, testRegion, &UpdateClusterTagsRequest{
		ClusterId: "1111111",
		Tags:      []Tag{{TagKey: "key", TagValue: "value"}},
	})
	if err != nil {
		t.Fatalf("update cluster tags failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected success: %v", result.Success)
	}
}

// TestUpdateClusterTagsAcceptsEmptyTagList pins the required-field rule: Tags is rejected only
// when it is nil, so an explicitly empty (tag-clearing) list must still reach the server.
func TestUpdateClusterTagsAcceptsEmptyTagList(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		body := readJSONBody(t, r)
		tags, ok := body["tags"].([]interface{})
		if !ok || len(tags) != 0 {
			t.Fatalf("unexpected tags: %v", body["tags"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": nil})
	}))
	defer server.Close()

	result, err := UpdateClusterTags(cli, testRegion, &UpdateClusterTagsRequest{
		ClusterId: "1111111",
		Tags:      []Tag{},
	})
	if err != nil {
		t.Fatalf("update cluster tags failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected success: %v", result.Success)
	}
}

func TestBatchInsertTags(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if r.Method != nethttp.MethodPost {
			t.Errorf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/api/bes/cluster/batchInsertTags" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get(HEADER_REGION); got != testRegion {
			t.Errorf("unexpected region header: %s", got)
		}
		body := readJSONBody(t, r)
		idList, ok := body["clusterIdList"].([]interface{})
		if !ok || len(idList) != 2 || idList[0] != "111" || idList[1] != "222" {
			t.Fatalf("unexpected clusterIdList: %v", body["clusterIdList"])
		}
		insertTags, ok := body["insertTags"].([]interface{})
		if !ok || len(insertTags) != 1 {
			t.Fatalf("unexpected insertTags: %v", body["insertTags"])
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": map[string]interface{}{}})
	}))
	defer server.Close()

	result, err := BatchInsertTags(cli, testRegion, &BatchInsertTagsRequest{
		ClusterIdList: []string{"111", "222"},
		InsertTags:    []Tag{{TagKey: "key", TagValue: "value"}},
	})
	if err != nil {
		t.Fatalf("batch insert tags failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected success: %v", result.Success)
	}
}

func TestTagAPIRejectsNilClient(t *testing.T) {
	cases := []struct {
		name string
		call func() error
	}{
		{"ListTags", func() error {
			_, err := ListTags(nil, testRegion)
			return err
		}},
		{"UpdateClusterTags", func() error {
			_, err := UpdateClusterTags(nil, testRegion, &UpdateClusterTagsRequest{})
			return err
		}},
		{"BatchInsertTags", func() error {
			_, err := BatchInsertTags(nil, testRegion, &BatchInsertTagsRequest{})
			return err
		}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := c.call(); !errors.Is(err, ErrNilClient) {
				t.Fatalf("expected ErrNilClient, got %v", err)
			}
		})
	}
}

func TestTagAPIRejectsNilRequest(t *testing.T) {
	cli := tagErrorServerClient(t)

	cases := []struct {
		name string
		call func() error
	}{
		{"UpdateClusterTags", func() error {
			_, err := UpdateClusterTags(cli, testRegion, nil)
			return err
		}},
		{"BatchInsertTags", func() error {
			_, err := BatchInsertTags(cli, testRegion, nil)
			return err
		}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.call()
			if err == nil {
				t.Fatal("expected error for nil request")
			}
			if errors.Is(err, ErrNilClient) {
				t.Fatalf("nil request must not report a nil client: %v", err)
			}
		})
	}
}

func TestTagAPIRejectsMissingRequiredFields(t *testing.T) {
	cli := tagErrorServerClient(t)

	cases := []struct {
		name string
		call func() error
	}{
		{"UpdateClusterTags/clusterId", func() error {
			_, err := UpdateClusterTags(cli, testRegion, &UpdateClusterTagsRequest{
				ClusterId: "",
				Tags:      []Tag{{TagKey: "key", TagValue: "value"}},
			})
			return err
		}},
		{"UpdateClusterTags/tags", func() error {
			_, err := UpdateClusterTags(cli, testRegion, &UpdateClusterTagsRequest{
				ClusterId: "1111111",
				Tags:      nil,
			})
			return err
		}},
		{"BatchInsertTags/clusterIdList", func() error {
			_, err := BatchInsertTags(cli, testRegion, &BatchInsertTagsRequest{
				ClusterIdList: nil,
				InsertTags:    []Tag{{TagKey: "key", TagValue: "value"}},
			})
			return err
		}},
		{"BatchInsertTags/insertTags", func() error {
			_, err := BatchInsertTags(cli, testRegion, &BatchInsertTagsRequest{
				ClusterIdList: []string{"111"},
				InsertTags:    nil,
			})
			return err
		}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := c.call(); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func TestTagAPIPropagatesServerError(t *testing.T) {
	cli := tagErrorServerClient(t)

	cases := []struct {
		name string
		call func() error
	}{
		{"ListTags", func() error {
			_, err := ListTags(cli, testRegion)
			return err
		}},
		{"UpdateClusterTags", func() error {
			_, err := UpdateClusterTags(cli, testRegion, &UpdateClusterTagsRequest{
				ClusterId: "1111111",
				Tags:      []Tag{{TagKey: "key", TagValue: "value"}},
			})
			return err
		}},
		{"BatchInsertTags", func() error {
			_, err := BatchInsertTags(cli, testRegion, &BatchInsertTagsRequest{
				ClusterIdList: []string{"111"},
				InsertTags:    []Tag{{TagKey: "key", TagValue: "value"}},
			})
			return err
		}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if err := c.call(); err == nil {
				t.Fatal("expected server error to be propagated")
			}
		})
	}
}

package api

import (
	"errors"
	nethttp "net/http"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

func TestListIndicesByPattern(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/indices/_list_by_pattern" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if body["pattern"] != "openapi-test-*" {
			t.Fatalf("unexpected pattern: %v", body["pattern"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"indices": []map[string]interface{}{{"index": "openapi-test-1", "indexSize": "9.9kb", "indexPriSize": "3.4kb"}},
		})
	}))
	defer server.Close()

	result, err := ListIndicesByPattern(cli, testRegion, &ListIndicesByPatternRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
		Pattern:   "openapi-test-*",
	})
	if err != nil {
		t.Fatalf("list indices by pattern failed: %v", err)
	}
	if len(result.Indices) != 1 || result.Indices[0].Index != "openapi-test-1" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestListRestores populates every optional query field so each conditional query-parameter
// branch in ListRestores is exercised.
func TestListRestores(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/snapshots/27/restores" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		query := r.URL.Query()
		expected := map[string]string{
			"pageNo":        "1",
			"pageSize":      "10",
			"order":         "desc",
			"orderBy":       "startTimestamp",
			"restoreStatus": "IN_PROGRESS",
		}
		for key, want := range expected {
			if got := query.Get(key); got != want {
				t.Fatalf("unexpected %s: got %s want %s", key, got, want)
			}
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"restores": []map[string]interface{}{{
				"historyId":       "27",
				"restoreId":       "4",
				"destClusterId":   "search-xxxx",
				"destClusterName": "testOpenToPost1",
				"restoreStatus":   "IN_PROGRESS",
			}},
			"pageNo":     1,
			"pageSize":   10,
			"totalCount": 1,
			"orderBy":    "startTimestamp",
			"order":      "desc",
		})
	}))
	defer server.Close()

	result, err := ListRestores(cli, testRegion, &ListRestoresRequest{
		ClusterId:     "search-xxxxxxxx",
		SnapshotId:    "27",
		PageNo:        1,
		PageSize:      10,
		Order:         "desc",
		OrderBy:       "startTimestamp",
		RestoreStatus: "IN_PROGRESS",
	})
	if err != nil {
		t.Fatalf("list restores failed: %v", err)
	}
	if result.TotalCount != 1 || result.Restores[0].RestoreId != "4" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestListRestoresWithoutOptionalFields covers the request path where every optional query
// parameter is omitted from the query string entirely.
func TestListRestoresWithoutOptionalFields(t *testing.T) {
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

	result, err := ListRestores(cli, testRegion, &ListRestoresRequest{
		ClusterId:  "search-xxxxxxxx",
		SnapshotId: "27",
	})
	if err != nil {
		t.Fatalf("list restores failed: %v", err)
	}
	if result.TotalCount != 0 || len(result.Restores) != 0 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestListRestoreClusters(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/snapshots/27/restore-clusters" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("includeSelf"); got != "true" {
			t.Fatalf("unexpected includeSelf: %s", r.URL.RawQuery)
		}
		if got := r.URL.Query().Get("keyword"); got != "testOpen" {
			t.Fatalf("unexpected keyword: %s", r.URL.RawQuery)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusters": []map[string]interface{}{{"clusterId": "search-xxxx", "name": "testOpenToPost1", "available": true}},
		})
	}))
	defer server.Close()

	result, err := ListRestoreClusters(cli, testRegion, &ListRestoreClustersRequest{
		ClusterId:   "search-xxxxxxxx",
		SnapshotId:  "27",
		IncludeSelf: boolPtr(true),
		Keyword:     "testOpen",
	})
	if err != nil {
		t.Fatalf("list restore clusters failed: %v", err)
	}
	if len(result.Clusters) != 1 || !result.Clusters[0].Available {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestGetSnapshotRule populates every optional query field so each conditional query-parameter
// branch in GetSnapshotRule is exercised.
func TestGetSnapshotRule(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/snapshots/26/rule" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		query := r.URL.Query()
		expected := map[string]string{
			"pageNo":              "1",
			"pageSize":            "10",
			"orderBy":             "indexName",
			"order":               "asc",
			"indexName":           "openapi-test-2",
			"restoreIndexPattern": "openapi-test-*",
		}
		for key, want := range expected {
			if got := query.Get(key); got != want {
				t.Fatalf("unexpected %s: got %s want %s", key, got, want)
			}
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"snapshotId":     "26",
			"configId":       "29",
			"snapshotName":   "openapi_test_manual_snapshot",
			"snapshotStatus": "SUCCESS",
			"totalCount":     6,
			"indices":        []map[string]interface{}{{"indexName": "openapi-test-2", "totalSize": "3.41KB"}},
		})
	}))
	defer server.Close()

	result, err := GetSnapshotRule(cli, testRegion, &GetSnapshotRuleRequest{
		ClusterId:           "search-xxxxxxxx",
		SnapshotId:          "26",
		PageNo:              1,
		PageSize:            10,
		OrderBy:             "indexName",
		Order:               "asc",
		IndexName:           "openapi-test-2",
		RestoreIndexPattern: "openapi-test-*",
	})
	if err != nil {
		t.Fatalf("get snapshot rule failed: %v", err)
	}
	if result.ConfigId != "29" || result.TotalCount != 6 || len(result.Indices) != 1 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestListSnapshots populates every optional query field so each conditional query-parameter
// branch in ListSnapshots is exercised.
func TestListSnapshots(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/snapshots" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		query := r.URL.Query()
		expected := map[string]string{
			"pageNo":         "1",
			"pageSize":       "10",
			"order":          "desc",
			"orderBy":        "startTimestamp",
			"snapshotName":   "openapi_test_manual_snapshot",
			"type":           "MANUAL",
			"startTimestamp": "1784633309349",
			"endTimestamp":   "1784719709349",
			"snapshotStatus": "SUCCESS",
			"restoreStatus":  "IN_PROGRESS",
		}
		for key, want := range expected {
			if got := query.Get(key); got != want {
				t.Fatalf("unexpected %s: got %s want %s", key, got, want)
			}
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"snapshots": []map[string]interface{}{{
				"historyId":      "27",
				"configId":       "31",
				"type":           "MANUAL",
				"snapshotStatus": "SUCCESS",
				"totalSize":      "2.1MB",
			}},
			"pageNo":     1,
			"pageSize":   10,
			"totalCount": 8,
		})
	}))
	defer server.Close()

	result, err := ListSnapshots(cli, testRegion, &ListSnapshotsRequest{
		ClusterId:      "search-xxxxxxxx",
		PageNo:         1,
		PageSize:       10,
		Order:          "desc",
		OrderBy:        "startTimestamp",
		SnapshotName:   "openapi_test_manual_snapshot",
		Type:           "MANUAL",
		StartTimestamp: 1784633309349,
		EndTimestamp:   1784719709349,
		SnapshotStatus: "SUCCESS",
		RestoreStatus:  "IN_PROGRESS",
	})
	if err != nil {
		t.Fatalf("list snapshots failed: %v", err)
	}
	if result.TotalCount != 8 || result.Snapshots[0].HistoryId != "27" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetSnapshotConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/snapshot-configs/30" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"configId":     "30",
			"type":         "AUTO",
			"name":         "openapi_test_auto_snapshot6_updated",
			"snapshotType": "SNAPSHOT_BOS",
			"period":       "0 2 * * *",
			"enabled":      true,
			"allIndex":     true,
		})
	}))
	defer server.Close()

	result, err := GetSnapshotConfig(cli, testRegion, &GetSnapshotConfigRequest{
		ClusterId: "search-xxxxxxxx",
		ConfigId:  "30",
	})
	if err != nil {
		t.Fatalf("get snapshot config failed: %v", err)
	}
	if result.ConfigId != "30" || !result.Enabled || !result.AllIndex {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestGetSnapshot(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/snapshots/27" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"historyId":      "27",
			"configId":       "31",
			"type":           "MANUAL",
			"snapshotName":   "openapi_test_manual_snapshot",
			"snapshotStatus": "SUCCESS",
			"totalSize":      "2.1MB",
		})
	}))
	defer server.Close()

	result, err := GetSnapshot(cli, testRegion, &GetSnapshotRequest{
		ClusterId:  "search-xxxxxxxx",
		SnapshotId: "27",
	})
	if err != nil {
		t.Fatalf("get snapshot failed: %v", err)
	}
	if result.HistoryId != "27" || result.SnapshotStatus != "SUCCESS" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestListAutoSnapshotConfigs(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/auto-snapshot-configs" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"configList": []map[string]interface{}{{
				"configId":     "28",
				"type":         "AUTO",
				"name":         "openapi_test_auto_snapshot5_updated",
				"snapshotType": "SNAPSHOT_BOS",
				"period":       "0 2 * * *",
				"enabled":      true,
				"allIndex":     true,
			}},
		})
	}))
	defer server.Close()

	result, err := ListAutoSnapshotConfigs(cli, testRegion, &ListAutoSnapshotConfigsRequest{ClusterId: "search-xxxxxxxx"})
	if err != nil {
		t.Fatalf("list auto snapshot configs failed: %v", err)
	}
	if len(result.ConfigList) != 1 || result.ConfigList[0].ConfigId != "28" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestCreateRestore(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/snapshots/27/restores" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if body["destClusterId"] != "search-xxxx" || body["restoreIndexPattern"] != "openapi-test-*" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["renamePattern"] != "(.+)" || body["renameReplacement"] != "restored-$1" {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"restoreId": 4})
	}))
	defer server.Close()

	result, err := CreateRestore(cli, testRegion, &CreateRestoreRequest{
		ClusterId:           "search-xxxxxxxx",
		SnapshotId:          "27",
		DestClusterId:       "search-xxxx",
		AllIndex:            boolPtr(false),
		RestoreIndices:      "openapi-test-1",
		RestoreIndexPattern: "openapi-test-*",
		RenamePattern:       "(.+)",
		RenameReplacement:   "restored-$1",
	})
	if err != nil {
		t.Fatalf("create restore failed: %v", err)
	}
	if result.RestoreId != 4 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestCreateManualSnapshot(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/manual-snapshot-configs" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if body["name"] != "openapi_test_manual_snapshot" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["snapshotType"] != "SNAPSHOT_BOS" || body["expireAfter"] != "7d" {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"configId": 31})
	}))
	defer server.Close()

	result, err := CreateManualSnapshot(cli, testRegion, &CreateManualSnapshotRequest{
		ClusterId:            "search-xxxxxxxx",
		Name:                 "openapi_test_manual_snapshot",
		SnapshotType:         "SNAPSHOT_BOS",
		ExpireAfter:          "7d",
		Remark:               "created by unit test",
		AllIndex:             boolPtr(true),
		SnapshotIndices:      "openapi-test-1",
		SnapshotIndexPattern: "openapi-test-*",
	})
	if err != nil {
		t.Fatalf("create manual snapshot failed: %v", err)
	}
	if result.ConfigId != 31 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestCreateAutoSnapshotConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/auto-snapshot-configs" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if body["name"] != "openapi_test_auto_snapshot7" || body["period"] != "0 2 * * *" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["enabled"] != true || body["allIndex"] != true {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"configId": 32})
	}))
	defer server.Close()

	result, err := CreateAutoSnapshotConfig(cli, testRegion, &CreateAutoSnapshotConfigRequest{
		ClusterId:            "search-xxxxxxxx",
		Name:                 "openapi_test_auto_snapshot7",
		SnapshotType:         "SNAPSHOT_BOS",
		Period:               "0 2 * * *",
		ExpireAfter:          "7d",
		Remark:               "created by unit test",
		AllIndex:             boolPtr(true),
		SnapshotIndices:      "openapi-test-1",
		SnapshotIndexPattern: "openapi-test-*",
		Enabled:              boolPtr(true),
	})
	if err != nil {
		t.Fatalf("create auto snapshot config failed: %v", err)
	}
	if result.ConfigId != 32 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestUpdateAutoSnapshotConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPut {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/auto-snapshot-configs/28" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		if body["name"] != "openapi_test_auto_snapshot5_updated" {
			t.Fatalf("unexpected body: %+v", body)
		}
		if body["period"] != "0 2 * * *" || body["enabled"] != true {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{"configId": 1})
	}))
	defer server.Close()

	result, err := UpdateAutoSnapshotConfig(cli, testRegion, &UpdateAutoSnapshotConfigRequest{
		ClusterId:            "search-xxxxxxxx",
		ConfigId:             "28",
		Name:                 "openapi_test_auto_snapshot5_updated",
		SnapshotType:         "SNAPSHOT_BOS",
		Period:               "0 2 * * *",
		ExpireAfter:          "7d",
		Remark:               "updated by unit test",
		AllIndex:             boolPtr(true),
		SnapshotIndices:      "openapi-test-1",
		SnapshotIndexPattern: "openapi-test-*",
		Enabled:              boolPtr(true),
	})
	if err != nil {
		t.Fatalf("update auto snapshot config failed: %v", err)
	}
	if result.ConfigId != 1 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestDeleteSnapshot(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/snapshots/27" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"snapshotId": 27})
	}))
	defer server.Close()

	result, err := DeleteSnapshot(cli, testRegion, &DeleteSnapshotRequest{
		ClusterId:  "search-xxxxxxxx",
		SnapshotId: "27",
	})
	if err != nil {
		t.Fatalf("delete snapshot failed: %v", err)
	}
	if result.SnapshotId != 27 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestDeleteAutoSnapshotConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/auto-snapshot-configs/1" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		writeJSONResponse(t, w, map[string]interface{}{"configId": 1})
	}))
	defer server.Close()

	result, err := DeleteAutoSnapshotConfig(cli, testRegion, &DeleteAutoSnapshotConfigRequest{
		ClusterId: "search-xxxxxxxx",
		ConfigId:  "1",
	})
	if err != nil {
		t.Fatalf("delete auto snapshot config failed: %v", err)
	}
	if result.ConfigId != 1 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

// TestBatchEnableAutoSnapshotConfigs is written from scratch: BatchEnableAutoSnapshotConfigs
// PATCHes the auto snapshot config collection URI with the enabled flag and the target config ids
// in the JSON body.
func TestBatchEnableAutoSnapshotConfigs(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPatch {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/backup/auto-snapshot-configs" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		if body["enabled"] != false {
			t.Fatalf("unexpected enabled: %v", body["enabled"])
		}
		configIdList, ok := body["configIdList"].([]interface{})
		if !ok || len(configIdList) != 2 {
			t.Fatalf("unexpected configIdList: %v", body["configIdList"])
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"enabled":      false,
			"configIdList": []int64{28, 30},
		})
	}))
	defer server.Close()

	result, err := BatchEnableAutoSnapshotConfigs(cli, testRegion, &BatchEnableAutoSnapshotConfigsRequest{
		Request:      Request{Region: "sh"},
		ClusterId:    "search-xxxxxxxx",
		Enabled:      boolPtr(false),
		ConfigIdList: []int64{28, 30},
	})
	if err != nil {
		t.Fatalf("batch enable auto snapshot configs failed: %v", err)
	}
	if result.Enabled || len(result.ConfigIdList) != 2 || result.ConfigIdList[0] != 28 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestSnapshotAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"ListIndicesByPattern": func() error {
			_, err := ListIndicesByPattern(nil, testRegion, &ListIndicesByPatternRequest{})
			return err
		},
		"ListRestores": func() error { _, err := ListRestores(nil, testRegion, &ListRestoresRequest{}); return err },
		"ListRestoreClusters": func() error {
			_, err := ListRestoreClusters(nil, testRegion, &ListRestoreClustersRequest{})
			return err
		},
		"GetSnapshotRule": func() error { _, err := GetSnapshotRule(nil, testRegion, &GetSnapshotRuleRequest{}); return err },
		"ListSnapshots":   func() error { _, err := ListSnapshots(nil, testRegion, &ListSnapshotsRequest{}); return err },
		"GetSnapshotConfig": func() error {
			_, err := GetSnapshotConfig(nil, testRegion, &GetSnapshotConfigRequest{})
			return err
		},
		"GetSnapshot": func() error { _, err := GetSnapshot(nil, testRegion, &GetSnapshotRequest{}); return err },
		"ListAutoSnapshotConfigs": func() error {
			_, err := ListAutoSnapshotConfigs(nil, testRegion, &ListAutoSnapshotConfigsRequest{})
			return err
		},
		"CreateRestore": func() error { _, err := CreateRestore(nil, testRegion, &CreateRestoreRequest{}); return err },
		"CreateManualSnapshot": func() error {
			_, err := CreateManualSnapshot(nil, testRegion, &CreateManualSnapshotRequest{})
			return err
		},
		"CreateAutoSnapshotConfig": func() error {
			_, err := CreateAutoSnapshotConfig(nil, testRegion, &CreateAutoSnapshotConfigRequest{})
			return err
		},
		"UpdateAutoSnapshotConfig": func() error {
			_, err := UpdateAutoSnapshotConfig(nil, testRegion, &UpdateAutoSnapshotConfigRequest{})
			return err
		},
		"DeleteSnapshot": func() error { _, err := DeleteSnapshot(nil, testRegion, &DeleteSnapshotRequest{}); return err },
		"DeleteAutoSnapshotConfig": func() error {
			_, err := DeleteAutoSnapshotConfig(nil, testRegion, &DeleteAutoSnapshotConfigRequest{})
			return err
		},
		"BatchEnableAutoSnapshotConfigs": func() error {
			_, err := BatchEnableAutoSnapshotConfigs(nil, testRegion, &BatchEnableAutoSnapshotConfigsRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

// TestSnapshotAPIRejectsNilRequest covers all 15 snapshot functions: none of them treats the
// request as optional, so every one must reject nil before any HTTP call is made.
func TestSnapshotAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ListIndicesByPattern": func() error { _, err := ListIndicesByPattern(cli, testRegion, nil); return err },
		"ListRestores":         func() error { _, err := ListRestores(cli, testRegion, nil); return err },
		"ListRestoreClusters":  func() error { _, err := ListRestoreClusters(cli, testRegion, nil); return err },
		"GetSnapshotRule":      func() error { _, err := GetSnapshotRule(cli, testRegion, nil); return err },
		"ListSnapshots":        func() error { _, err := ListSnapshots(cli, testRegion, nil); return err },
		"GetSnapshotConfig":    func() error { _, err := GetSnapshotConfig(cli, testRegion, nil); return err },
		"GetSnapshot":          func() error { _, err := GetSnapshot(cli, testRegion, nil); return err },
		"ListAutoSnapshotConfigs": func() error {
			_, err := ListAutoSnapshotConfigs(cli, testRegion, nil)
			return err
		},
		"CreateRestore":        func() error { _, err := CreateRestore(cli, testRegion, nil); return err },
		"CreateManualSnapshot": func() error { _, err := CreateManualSnapshot(cli, testRegion, nil); return err },
		"CreateAutoSnapshotConfig": func() error {
			_, err := CreateAutoSnapshotConfig(cli, testRegion, nil)
			return err
		},
		"UpdateAutoSnapshotConfig": func() error {
			_, err := UpdateAutoSnapshotConfig(cli, testRegion, nil)
			return err
		},
		"DeleteSnapshot": func() error { _, err := DeleteSnapshot(cli, testRegion, nil); return err },
		"DeleteAutoSnapshotConfig": func() error {
			_, err := DeleteAutoSnapshotConfig(cli, testRegion, nil)
			return err
		},
		"BatchEnableAutoSnapshotConfigs": func() error {
			_, err := BatchEnableAutoSnapshotConfigs(cli, testRegion, nil)
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

// newSnapshotRejectingClient returns a client backed by a server that fails the test if it is ever
// reached. The many snapshot validation cases below all have to be rejected before any HTTP call,
// so this keeps each of them to a single line of setup.
func newSnapshotRejectingClient(t *testing.T) bce.Client {
	t.Helper()
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	t.Cleanup(server.Close)
	return cli
}

// TestSnapshotAPIRejectsMissingClusterId covers the three snapshot functions whose only required
// field is clusterId.
func TestSnapshotAPIRejectsMissingClusterId(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	cases := map[string]func() error{
		"ListIndicesByPattern": func() error {
			_, err := ListIndicesByPattern(cli, testRegion, &ListIndicesByPatternRequest{Pattern: "openapi-test-*"})
			return err
		},
		"ListSnapshots": func() error {
			_, err := ListSnapshots(cli, testRegion, &ListSnapshotsRequest{PageNo: 1, PageSize: 10})
			return err
		},
		"ListAutoSnapshotConfigs": func() error {
			_, err := ListAutoSnapshotConfigs(cli, testRegion, &ListAutoSnapshotConfigsRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestListRestoresRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := ListRestoresRequest{ClusterId: "search-xxxxxxxx", SnapshotId: "27"}
	cases := map[string]func(*ListRestoresRequest){
		"clusterId":  func(r *ListRestoresRequest) { r.ClusterId = "" },
		"snapshotId": func(r *ListRestoresRequest) { r.SnapshotId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := ListRestores(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestListRestoreClustersRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := ListRestoreClustersRequest{ClusterId: "search-xxxxxxxx", SnapshotId: "27"}
	cases := map[string]func(*ListRestoreClustersRequest){
		"clusterId":  func(r *ListRestoreClustersRequest) { r.ClusterId = "" },
		"snapshotId": func(r *ListRestoreClustersRequest) { r.SnapshotId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := ListRestoreClusters(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestGetSnapshotRuleRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := GetSnapshotRuleRequest{ClusterId: "search-xxxxxxxx", SnapshotId: "26"}
	cases := map[string]func(*GetSnapshotRuleRequest){
		"clusterId":  func(r *GetSnapshotRuleRequest) { r.ClusterId = "" },
		"snapshotId": func(r *GetSnapshotRuleRequest) { r.SnapshotId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := GetSnapshotRule(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestGetSnapshotConfigRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := GetSnapshotConfigRequest{ClusterId: "search-xxxxxxxx", ConfigId: "30"}
	cases := map[string]func(*GetSnapshotConfigRequest){
		"clusterId": func(r *GetSnapshotConfigRequest) { r.ClusterId = "" },
		"configId":  func(r *GetSnapshotConfigRequest) { r.ConfigId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := GetSnapshotConfig(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestGetSnapshotRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := GetSnapshotRequest{ClusterId: "search-xxxxxxxx", SnapshotId: "27"}
	cases := map[string]func(*GetSnapshotRequest){
		"clusterId":  func(r *GetSnapshotRequest) { r.ClusterId = "" },
		"snapshotId": func(r *GetSnapshotRequest) { r.SnapshotId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := GetSnapshot(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestCreateRestoreRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := CreateRestoreRequest{
		ClusterId:     "search-xxxxxxxx",
		SnapshotId:    "27",
		DestClusterId: "search-xxxx",
		AllIndex:      boolPtr(true),
	}
	cases := map[string]func(*CreateRestoreRequest){
		"clusterId":     func(r *CreateRestoreRequest) { r.ClusterId = "" },
		"snapshotId":    func(r *CreateRestoreRequest) { r.SnapshotId = "" },
		"destClusterId": func(r *CreateRestoreRequest) { r.DestClusterId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := CreateRestore(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestCreateManualSnapshotRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := CreateManualSnapshotRequest{
		ClusterId:   "search-xxxxxxxx",
		Name:        "openapi_test_manual_snapshot",
		ExpireAfter: "7d",
		AllIndex:    boolPtr(true),
	}
	cases := map[string]func(*CreateManualSnapshotRequest){
		"clusterId": func(r *CreateManualSnapshotRequest) { r.ClusterId = "" },
		"name":      func(r *CreateManualSnapshotRequest) { r.Name = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := CreateManualSnapshot(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestCreateAutoSnapshotConfigRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := CreateAutoSnapshotConfigRequest{
		ClusterId: "search-xxxxxxxx",
		Name:      "openapi_test_auto_snapshot7",
		Period:    "0 2 * * *",
		Enabled:   boolPtr(true),
	}
	cases := map[string]func(*CreateAutoSnapshotConfigRequest){
		"clusterId": func(r *CreateAutoSnapshotConfigRequest) { r.ClusterId = "" },
		"name":      func(r *CreateAutoSnapshotConfigRequest) { r.Name = "" },
		"period":    func(r *CreateAutoSnapshotConfigRequest) { r.Period = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := CreateAutoSnapshotConfig(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUpdateAutoSnapshotConfigRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := UpdateAutoSnapshotConfigRequest{
		ClusterId: "search-xxxxxxxx",
		ConfigId:  "28",
		Name:      "openapi_test_auto_snapshot5_updated",
		Period:    "0 2 * * *",
		Enabled:   boolPtr(true),
	}
	cases := map[string]func(*UpdateAutoSnapshotConfigRequest){
		"clusterId": func(r *UpdateAutoSnapshotConfigRequest) { r.ClusterId = "" },
		"configId":  func(r *UpdateAutoSnapshotConfigRequest) { r.ConfigId = "" },
		"name":      func(r *UpdateAutoSnapshotConfigRequest) { r.Name = "" },
		"period":    func(r *UpdateAutoSnapshotConfigRequest) { r.Period = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UpdateAutoSnapshotConfig(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestDeleteSnapshotRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := DeleteSnapshotRequest{ClusterId: "search-xxxxxxxx", SnapshotId: "27"}
	cases := map[string]func(*DeleteSnapshotRequest){
		"clusterId":  func(r *DeleteSnapshotRequest) { r.ClusterId = "" },
		"snapshotId": func(r *DeleteSnapshotRequest) { r.SnapshotId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := DeleteSnapshot(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestDeleteAutoSnapshotConfigRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := DeleteAutoSnapshotConfigRequest{ClusterId: "search-xxxxxxxx", ConfigId: "1"}
	cases := map[string]func(*DeleteAutoSnapshotConfigRequest){
		"clusterId": func(r *DeleteAutoSnapshotConfigRequest) { r.ClusterId = "" },
		"configId":  func(r *DeleteAutoSnapshotConfigRequest) { r.ConfigId = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := DeleteAutoSnapshotConfig(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestBatchEnableAutoSnapshotConfigsRejectsMissingRequiredFields(t *testing.T) {
	cli := newSnapshotRejectingClient(t)

	full := BatchEnableAutoSnapshotConfigsRequest{
		ClusterId:    "search-xxxxxxxx",
		Enabled:      boolPtr(true),
		ConfigIdList: []int64{28, 30},
	}
	cases := map[string]func(*BatchEnableAutoSnapshotConfigsRequest){
		"clusterId":    func(r *BatchEnableAutoSnapshotConfigsRequest) { r.ClusterId = "" },
		"enabled":      func(r *BatchEnableAutoSnapshotConfigsRequest) { r.Enabled = nil },
		"configIdList": func(r *BatchEnableAutoSnapshotConfigsRequest) { r.ConfigIdList = nil },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := BatchEnableAutoSnapshotConfigs(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestSnapshotAPIPropagatesServerError uses 400 rather than 500 on purpose: the helper mirrors the
// production retry policy, which retries 5xx with backoff and would make this test take seconds.
func TestSnapshotAPIPropagatesServerError(t *testing.T) {
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
		"ListIndicesByPattern": func() error {
			_, err := ListIndicesByPattern(cli, testRegion, &ListIndicesByPatternRequest{
				ClusterId: "search-xxxxxxxx",
				Pattern:   "openapi-test-*",
			})
			return err
		},
		"ListRestores": func() error {
			_, err := ListRestores(cli, testRegion, &ListRestoresRequest{
				ClusterId:  "search-xxxxxxxx",
				SnapshotId: "27",
			})
			return err
		},
		"ListRestoreClusters": func() error {
			_, err := ListRestoreClusters(cli, testRegion, &ListRestoreClustersRequest{
				ClusterId:  "search-xxxxxxxx",
				SnapshotId: "27",
			})
			return err
		},
		"GetSnapshotRule": func() error {
			_, err := GetSnapshotRule(cli, testRegion, &GetSnapshotRuleRequest{
				ClusterId:  "search-xxxxxxxx",
				SnapshotId: "26",
			})
			return err
		},
		"ListSnapshots": func() error {
			_, err := ListSnapshots(cli, testRegion, &ListSnapshotsRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"GetSnapshotConfig": func() error {
			_, err := GetSnapshotConfig(cli, testRegion, &GetSnapshotConfigRequest{
				ClusterId: "search-xxxxxxxx",
				ConfigId:  "30",
			})
			return err
		},
		"GetSnapshot": func() error {
			_, err := GetSnapshot(cli, testRegion, &GetSnapshotRequest{
				ClusterId:  "search-xxxxxxxx",
				SnapshotId: "27",
			})
			return err
		},
		"ListAutoSnapshotConfigs": func() error {
			_, err := ListAutoSnapshotConfigs(cli, testRegion, &ListAutoSnapshotConfigsRequest{
				ClusterId: "search-xxxxxxxx",
			})
			return err
		},
		"CreateRestore": func() error {
			_, err := CreateRestore(cli, testRegion, &CreateRestoreRequest{
				ClusterId:     "search-xxxxxxxx",
				SnapshotId:    "27",
				DestClusterId: "search-xxxx",
			})
			return err
		},
		"CreateManualSnapshot": func() error {
			_, err := CreateManualSnapshot(cli, testRegion, &CreateManualSnapshotRequest{
				ClusterId: "search-xxxxxxxx",
				Name:      "openapi_test_manual_snapshot",
			})
			return err
		},
		"CreateAutoSnapshotConfig": func() error {
			_, err := CreateAutoSnapshotConfig(cli, testRegion, &CreateAutoSnapshotConfigRequest{
				ClusterId: "search-xxxxxxxx",
				Name:      "openapi_test_auto_snapshot7",
				Period:    "0 2 * * *",
			})
			return err
		},
		"UpdateAutoSnapshotConfig": func() error {
			_, err := UpdateAutoSnapshotConfig(cli, testRegion, &UpdateAutoSnapshotConfigRequest{
				ClusterId: "search-xxxxxxxx",
				ConfigId:  "28",
				Name:      "openapi_test_auto_snapshot5_updated",
				Period:    "0 2 * * *",
			})
			return err
		},
		"DeleteSnapshot": func() error {
			_, err := DeleteSnapshot(cli, testRegion, &DeleteSnapshotRequest{
				ClusterId:  "search-xxxxxxxx",
				SnapshotId: "27",
			})
			return err
		},
		"DeleteAutoSnapshotConfig": func() error {
			_, err := DeleteAutoSnapshotConfig(cli, testRegion, &DeleteAutoSnapshotConfigRequest{
				ClusterId: "search-xxxxxxxx",
				ConfigId:  "1",
			})
			return err
		},
		"BatchEnableAutoSnapshotConfigs": func() error {
			_, err := BatchEnableAutoSnapshotConfigs(cli, testRegion, &BatchEnableAutoSnapshotConfigsRequest{
				ClusterId:    "search-xxxxxxxx",
				Enabled:      boolPtr(true),
				ConfigIdList: []int64{28},
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

package api

import (
	"errors"
	nethttp "net/http"
	"path/filepath"
	"testing"
)

func TestListSystemPlugins(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/system-plugins" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"devPlugins": []map[string]interface{}{
				{
					"pluginName":  "baidu-metrics",
					"engine":      "OPENSEARCH",
					"description": "百度 OpenSearch 指标插件",
					"order":       10,
					"versionInfos": []map[string]interface{}{
						{"version": "2.19.4", "status": "PreInstalled"},
					},
				},
			},
			"prePlugins": []map[string]interface{}{
				{
					"pluginName": "opensearch-alerting",
					"engine":     "OPENSEARCH",
					"order":      30,
					"versionInfos": []map[string]interface{}{
						{"version": "2.19.4.0", "status": "PreInstalled"},
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := ListSystemPlugins(cli, testRegion, &ListSystemPluginsRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
	})
	if err != nil {
		t.Fatalf("list system plugins failed: %v", err)
	}
	if len(result.DevPlugins) != 1 || result.DevPlugins[0].PluginName != "baidu-metrics" {
		t.Fatalf("unexpected dev plugins: %+v", result.DevPlugins)
	}
	if len(result.PrePlugins) != 1 || result.PrePlugins[0].PluginName != "opensearch-alerting" {
		t.Fatalf("unexpected pre plugins: %+v", result.PrePlugins)
	}
}

func TestUpdateSystemPlugin(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/system-plugins" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		body := readJSONBody(t, r)
		install, ok := body["install"].([]interface{})
		if !ok || len(install) != 1 {
			t.Fatalf("unexpected body: %+v", body)
		}
		uninstall, ok := body["uninstall"].([]interface{})
		if !ok || len(uninstall) != 1 {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId":   "search-xxxxxxxx",
			"clusterName": "test-cluster",
			"actionId":    "action-xxxxxxxx",
		})
	}))
	defer server.Close()

	result, err := UpdateSystemPlugin(cli, testRegion, &UpdateSystemPluginRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
		Install:   []PluginItem{{Name: "analysis-ik", Version: "1.0.0"}},
		Uninstall: []PluginItem{{Name: "analysis-pinyin", Version: "2.0.0"}},
	})
	if err != nil {
		t.Fatalf("update system plugin failed: %v", err)
	}
	if result.ActionId != "action-xxxxxxxx" || result.ClusterName != "test-cluster" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestListCustomPlugins(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/custom-plugins" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"customPlugins": []map[string]interface{}{
				{
					"pluginName": "custom-analyzer",
					"versionInfos": []map[string]interface{}{
						{"version": "1.0.0", "description": "自定义分析器插件", "status": "Idle", "uploadTime": 1783699200000},
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := ListCustomPlugins(cli, testRegion, &ListCustomPluginsRequest{ClusterId: "search-xxxxxxxx"})
	if err != nil {
		t.Fatalf("list custom plugins failed: %v", err)
	}
	if len(result.CustomPlugins) != 1 || result.CustomPlugins[0].PluginName != "custom-analyzer" {
		t.Fatalf("unexpected result: %+v", result.CustomPlugins)
	}
	if len(result.CustomPlugins[0].VersionInfos) != 1 || result.CustomPlugins[0].VersionInfos[0].UploadTime != 1783699200000 {
		t.Fatalf("unexpected version infos: %+v", result.CustomPlugins[0].VersionInfos)
	}
}

func TestUpdateCustomPlugin(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/custom-plugins" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		body := readJSONBody(t, r)
		install, ok := body["install"].([]interface{})
		if !ok || len(install) != 1 {
			t.Fatalf("unexpected body: %+v", body)
		}
		uninstall, ok := body["uninstall"].([]interface{})
		if !ok || len(uninstall) != 1 {
			t.Fatalf("unexpected body: %+v", body)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"clusterId":   "search-xxxxxxxx",
			"clusterName": "test-cluster",
			"actionId":    "action-xxxxxxxx",
		})
	}))
	defer server.Close()

	result, err := UpdateCustomPlugin(cli, testRegion, &UpdateCustomPluginRequest{
		ClusterId: "search-xxxxxxxx",
		Install:   []PluginItem{{Name: "custom-analyzer", Version: "1.0.0"}},
		Uninstall: []PluginItem{{Name: "custom-tokenizer", Version: "2.0.0"}},
	})
	if err != nil {
		t.Fatalf("update custom plugin failed: %v", err)
	}
	if result.ActionId != "action-xxxxxxxx" || result.ClusterId != "search-xxxxxxxx" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestDeleteCustomPluginVersion(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodDelete {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/custom-plugins/custom-analyzer/versions/1.0.0" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "gz" {
			t.Fatalf("unexpected region: %s", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"pluginName": "custom-analyzer",
			"version":    "1.0.0",
		})
	}))
	defer server.Close()

	result, err := DeleteCustomPluginVersion(cli, testRegion, &DeleteCustomPluginVersionRequest{
		Request:       Request{Region: "gz"},
		ClusterId:     "search-xxxxxxxx",
		PluginName:    "custom-analyzer",
		PluginVersion: "1.0.0",
	})
	if err != nil {
		t.Fatalf("delete custom plugin version failed: %v", err)
	}
	if result.PluginName != "custom-analyzer" || result.Version != "1.0.0" {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestUploadPluginFile(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/search-xxxxxxxx/files" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}
		fields, fileContent, fileName := readMultipartBody(t, r, "file")
		if fields["type"] != "PluginZip" || fields["subtype"] != "CustomPlugin" {
			t.Fatalf("unexpected fields: %+v", fields)
		}
		if fileName != "custom-analyzer.zip" || string(fileContent) != "zip-bytes" {
			t.Fatalf("unexpected file: name=%s content=%s", fileName, fileContent)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"url":  "bos://bucket/custom-analyzer.zip",
			"size": 102400,
		})
	}))
	defer server.Close()

	path := writeTempUploadFile(t, "custom-analyzer.zip", []byte("zip-bytes"))
	result, err := UploadPluginFile(cli, testRegion, &UploadPluginFileRequest{
		Request:   Request{Region: "sh"},
		ClusterId: "search-xxxxxxxx",
		FilePath:  path,
		Type:      "PluginZip",
		Subtype:   "CustomPlugin",
	})
	if err != nil {
		t.Fatalf("upload plugin file failed: %v", err)
	}
	if result.Url != "bos://bucket/custom-analyzer.zip" || result.Size != 102400 {
		t.Fatalf("unexpected result: %+v", result)
	}
}

func TestPluginAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"ListSystemPlugins": func() error {
			_, err := ListSystemPlugins(nil, testRegion, &ListSystemPluginsRequest{})
			return err
		},
		"UpdateSystemPlugin": func() error {
			_, err := UpdateSystemPlugin(nil, testRegion, &UpdateSystemPluginRequest{})
			return err
		},
		"ListCustomPlugins": func() error {
			_, err := ListCustomPlugins(nil, testRegion, &ListCustomPluginsRequest{})
			return err
		},
		"UpdateCustomPlugin": func() error {
			_, err := UpdateCustomPlugin(nil, testRegion, &UpdateCustomPluginRequest{})
			return err
		},
		"DeleteCustomPluginVersion": func() error {
			_, err := DeleteCustomPluginVersion(nil, testRegion, &DeleteCustomPluginVersionRequest{})
			return err
		},
		"UploadPluginFile": func() error {
			_, err := UploadPluginFile(nil, testRegion, &UploadPluginFileRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

func TestPluginAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ListSystemPlugins":  func() error { _, err := ListSystemPlugins(cli, testRegion, nil); return err },
		"UpdateSystemPlugin": func() error { _, err := UpdateSystemPlugin(cli, testRegion, nil); return err },
		"ListCustomPlugins":  func() error { _, err := ListCustomPlugins(cli, testRegion, nil); return err },
		"UpdateCustomPlugin": func() error { _, err := UpdateCustomPlugin(cli, testRegion, nil); return err },
		"DeleteCustomPluginVersion": func() error {
			_, err := DeleteCustomPluginVersion(cli, testRegion, nil)
			return err
		},
		"UploadPluginFile": func() error { _, err := UploadPluginFile(cli, testRegion, nil); return err },
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

func TestPluginAPIRejectsMissingClusterId(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ListSystemPlugins": func() error {
			_, err := ListSystemPlugins(cli, testRegion, &ListSystemPluginsRequest{})
			return err
		},
		"UpdateSystemPlugin": func() error {
			_, err := UpdateSystemPlugin(cli, testRegion, &UpdateSystemPluginRequest{
				Install: []PluginItem{{Name: "analysis-ik", Version: "1.0.0"}},
			})
			return err
		},
		"ListCustomPlugins": func() error {
			_, err := ListCustomPlugins(cli, testRegion, &ListCustomPluginsRequest{})
			return err
		},
		"UpdateCustomPlugin": func() error {
			_, err := UpdateCustomPlugin(cli, testRegion, &UpdateCustomPluginRequest{
				Install: []PluginItem{{Name: "custom-analyzer", Version: "1.0.0"}},
			})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestDeleteCustomPluginVersionRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := DeleteCustomPluginVersionRequest{
		ClusterId:     "search-xxxxxxxx",
		PluginName:    "custom-analyzer",
		PluginVersion: "1.0.0",
	}
	cases := map[string]func(*DeleteCustomPluginVersionRequest){
		"clusterId":     func(r *DeleteCustomPluginVersionRequest) { r.ClusterId = "" },
		"pluginName":    func(r *DeleteCustomPluginVersionRequest) { r.PluginName = "" },
		"pluginVersion": func(r *DeleteCustomPluginVersionRequest) { r.PluginVersion = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := DeleteCustomPluginVersion(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestUploadPluginFileRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	full := UploadPluginFileRequest{
		ClusterId: "search-xxxxxxxx",
		FilePath:  writeTempUploadFile(t, "custom-analyzer.zip", []byte("zip-bytes")),
		Type:      "PluginZip",
		Subtype:   "CustomPlugin",
	}
	cases := map[string]func(*UploadPluginFileRequest){
		"clusterId": func(r *UploadPluginFileRequest) { r.ClusterId = "" },
		"filePath":  func(r *UploadPluginFileRequest) { r.FilePath = "" },
	}
	for name, clear := range cases {
		request := full
		clear(&request)
		if _, err := UploadPluginFile(cli, testRegion, &request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestUploadPluginFileRejectsUnopenableFile covers the openUploadFile error path: the file must
// exist and must be a regular file, otherwise no request is sent at all.
func TestUploadPluginFileRejectsUnopenableFile(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	dir := t.TempDir()
	cases := map[string]string{
		"nonexistent": filepath.Join(dir, "missing-plugin.zip"),
		"directory":   dir,
	}
	for name, path := range cases {
		if _, err := UploadPluginFile(cli, testRegion, &UploadPluginFileRequest{
			ClusterId: "search-xxxxxxxx",
			FilePath:  path,
			Type:      "PluginZip",
		}); err == nil {
			t.Fatalf("%s: expected error for file path %q", name, path)
		}
	}
}

// TestPluginAPIPropagatesServerError uses 400 rather than 500 on purpose: the helper mirrors the
// production retry policy, which retries 5xx with backoff and would make this test take seconds.
func TestPluginAPIPropagatesServerError(t *testing.T) {
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

	uploadPath := writeTempUploadFile(t, "custom-analyzer.zip", []byte("zip-bytes"))
	cases := map[string]func() error{
		"ListSystemPlugins": func() error {
			_, err := ListSystemPlugins(cli, testRegion, &ListSystemPluginsRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"UpdateSystemPlugin": func() error {
			_, err := UpdateSystemPlugin(cli, testRegion, &UpdateSystemPluginRequest{
				ClusterId: "search-xxxxxxxx",
				Install:   []PluginItem{{Name: "analysis-ik", Version: "1.0.0"}},
			})
			return err
		},
		"ListCustomPlugins": func() error {
			_, err := ListCustomPlugins(cli, testRegion, &ListCustomPluginsRequest{ClusterId: "search-xxxxxxxx"})
			return err
		},
		"UpdateCustomPlugin": func() error {
			_, err := UpdateCustomPlugin(cli, testRegion, &UpdateCustomPluginRequest{
				ClusterId: "search-xxxxxxxx",
				Install:   []PluginItem{{Name: "custom-analyzer", Version: "1.0.0"}},
			})
			return err
		},
		"DeleteCustomPluginVersion": func() error {
			_, err := DeleteCustomPluginVersion(cli, testRegion, &DeleteCustomPluginVersionRequest{
				ClusterId: "search-xxxxxxxx", PluginName: "custom-analyzer", PluginVersion: "1.0.0",
			})
			return err
		},
		"UploadPluginFile": func() error {
			_, err := UploadPluginFile(cli, testRegion, &UploadPluginFileRequest{
				ClusterId: "search-xxxxxxxx", FilePath: uploadPath, Type: "PluginZip",
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

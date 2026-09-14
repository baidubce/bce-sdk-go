package api

import (
	"errors"
	nethttp "net/http"
	"path/filepath"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

const (
	pluginTestClusterId         = "644734225693675520"
	pluginTestModuleType        = "es_node"
	pluginTestDefaultPluginName = "baidu-rate-limiting"
	pluginTestCustomPluginName  = "xxx.zip"
	pluginTestSeparator         = "tab"
)

// pluginWriteBadRequest writes a 400 service error. A 4xx status is used on purpose: the test
// client mirrors the production DEFAULT_RETRY_POLICY, so a 5xx would trigger backoff retries.
func pluginWriteBadRequest(t *testing.T, w nethttp.ResponseWriter) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(nethttp.StatusBadRequest)
	if _, err := w.Write([]byte(`{"code":"BadRequest","message":"invalid plugin request"}`)); err != nil {
		t.Fatalf("write error response failed: %v", err)
	}
}

// pluginUnreachableHandler fails the test when a request reaches the server, used by cases that
// must fail during local validation or while opening the upload file.
func pluginUnreachableHandler(t *testing.T) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Errorf("unexpected request reached the server: %s", r.URL.Path)
		pluginWriteBadRequest(t, w)
	}
}

// pluginAssertRequest checks the parts of the request every plugin API builds the same way.
func pluginAssertRequest(t *testing.T, r *nethttp.Request, path string) {
	t.Helper()
	if r.URL.Path != path {
		t.Fatalf("unexpected path: %s", r.URL.Path)
	}
	if r.Method != nethttp.MethodPost {
		t.Fatalf("unexpected method: %s", r.Method)
	}
	if got := r.Header.Get(HEADER_REGION); got != testRegion {
		t.Fatalf("unexpected region header: %s", got)
	}
}

// pluginFullDefaultRequest / pluginFullCustomRequest / pluginFullUploadRequest /
// pluginFullUpdateNLPDictRequest return fully populated requests, so the required-field cases
// below only have to clear one field each.
func pluginFullDefaultRequest() *DefaultPluginRequest {
	return &DefaultPluginRequest{
		ClusterId:  pluginTestClusterId,
		PluginName: pluginTestDefaultPluginName,
		ModuleType: pluginTestModuleType,
	}
}

func pluginFullCustomRequest() *CustomPluginRequest {
	return &CustomPluginRequest{
		ClusterId:  pluginTestClusterId,
		PluginName: pluginTestCustomPluginName,
		ModuleType: pluginTestModuleType,
	}
}

func pluginFullUploadRequest(filePath string) *UploadCustomPluginRequest {
	return &UploadCustomPluginRequest{
		ClusterId:  pluginTestClusterId,
		ModuleType: pluginTestModuleType,
		FilePath:   filePath,
	}
}

func pluginFullUpdateNLPDictRequest(filePath string) *UpdateNLPDictRequest {
	return &UpdateNLPDictRequest{
		ClusterId: pluginTestClusterId,
		Separator: pluginTestSeparator,
		FilePath:  filePath,
	}
}

// pluginAssertDefaultBody asserts the JSON body shared by the two default plugin APIs.
func pluginAssertDefaultBody(t *testing.T, r *nethttp.Request) {
	t.Helper()
	body := readJSONBody(t, r)
	if body["clusterId"] != pluginTestClusterId || body["pluginName"] != pluginTestDefaultPluginName {
		t.Fatalf("unexpected body: %+v", body)
	}
	if body["moduleType"] != pluginTestModuleType {
		t.Fatalf("unexpected moduleType: %v", body["moduleType"])
	}
}

// pluginAssertCustomBody asserts the JSON body shared by the custom plugin
// install/uninstall/delete APIs.
func pluginAssertCustomBody(t *testing.T, r *nethttp.Request) {
	t.Helper()
	body := readJSONBody(t, r)
	if body["clusterId"] != pluginTestClusterId || body["pluginName"] != pluginTestCustomPluginName {
		t.Fatalf("unexpected body: %+v", body)
	}
	if body["moduleType"] != pluginTestModuleType {
		t.Fatalf("unexpected moduleType: %v", body["moduleType"])
	}
}

func TestInstallDefaultPlugin(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		pluginAssertRequest(t, r, "/api/bes/cluster/default_plugin/install")
		pluginAssertDefaultBody(t, r)
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := InstallDefaultPlugin(cli, testRegion, pluginFullDefaultRequest())
	if err != nil {
		t.Fatalf("install default plugin failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestUninstallDefaultPlugin(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		pluginAssertRequest(t, r, "/api/bes/cluster/default_plugin/uninstall")
		pluginAssertDefaultBody(t, r)
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := UninstallDefaultPlugin(cli, testRegion, pluginFullDefaultRequest())
	if err != nil {
		t.Fatalf("uninstall default plugin failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

// TestInstallCustomPlugin also asserts the JSON content type, since the custom plugin
// install/uninstall/delete APIs go through createJSONRequest while the upload API does not.
func TestInstallCustomPlugin(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		pluginAssertRequest(t, r, "/api/bes/cluster/plugin/install")
		if got := r.Header.Get("Content-Type"); got != bce.DEFAULT_CONTENT_TYPE {
			t.Fatalf("unexpected content type: %s", got)
		}
		pluginAssertCustomBody(t, r)
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"taskId": "task-1"},
		})
	}))
	defer server.Close()

	result, err := InstallCustomPlugin(cli, testRegion, pluginFullCustomRequest())
	if err != nil {
		t.Fatalf("install custom plugin failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	payload, ok := result.Result.(map[string]interface{})
	if !ok || payload["taskId"] != "task-1" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestUninstallCustomPlugin(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		pluginAssertRequest(t, r, "/api/bes/cluster/plugin/uninstall")
		pluginAssertCustomBody(t, r)
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": map[string]interface{}{}})
	}))
	defer server.Close()

	result, err := UninstallCustomPlugin(cli, testRegion, pluginFullCustomRequest())
	if err != nil {
		t.Fatalf("uninstall custom plugin failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result == nil {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestDeleteCustomPlugin(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		pluginAssertRequest(t, r, "/api/bes/cluster/plugin/delete")
		pluginAssertCustomBody(t, r)
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
	}))
	defer server.Close()

	result, err := DeleteCustomPlugin(cli, testRegion, pluginFullCustomRequest())
	if err != nil {
		t.Fatalf("delete custom plugin failed: %v", err)
	}
	if !result.Success || result.Status != 200 || result.Result != "result" {
		t.Fatalf("unexpected response: %+v", result)
	}
}

// TestUploadCustomPlugin asserts the multipart body: the clusterId and moduleType travel in the
// path instead of form fields, so the body must carry the file part only. The file names cover the
// three guessContentType branches (known extension, unknown extension, no extension at all).
func TestUploadCustomPlugin(t *testing.T) {
	cases := []struct {
		name     string
		fileName string
	}{
		{"known extension", "plugin.zip"},
		{"unknown extension", "plugin.besplugin"},
		{"no extension", "plugin"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
				defer r.Body.Close()
				pluginAssertRequest(t, r,
					"/api/bes/cluster/plugin/upload/"+pluginTestClusterId+"/"+pluginTestModuleType)
				fields, fileContent, fileName := readMultipartBody(t, r, "file")
				if len(fields) != 0 {
					t.Fatalf("unexpected extra form fields: %+v", fields)
				}
				if fileName != tc.fileName || string(fileContent) != "zip-bytes" {
					t.Fatalf("unexpected file: name=%s content=%s", fileName, fileContent)
				}
				writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": "result"})
			}))
			defer server.Close()

			path := writeTempUploadFile(t, tc.fileName, []byte("zip-bytes"))
			result, err := UploadCustomPlugin(cli, testRegion, pluginFullUploadRequest(path))
			if err != nil {
				t.Fatalf("upload custom plugin failed: %v", err)
			}
			if !result.Success || result.Status != 200 || result.Result != "result" {
				t.Fatalf("unexpected response: %+v", result)
			}
		})
	}
}

func TestGetPluginInfo(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		pluginAssertRequest(t, r, "/api/bes/cluster/plugin/info")
		if got := readJSONBody(t, r)["clusterId"]; got != pluginTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"plugins": []map[string]interface{}{
					{
						"moduleType":   pluginTestModuleType,
						"pluginDesc":   "Baidu NLP中文分词插件",
						"pluginName":   pluginTestCustomPluginName,
						"pluginStatus": "UNINSTALLED",
					},
				},
				"defaultPlugins": []map[string]interface{}{
					{
						"moduleType":   pluginTestModuleType,
						"pluginDesc":   "custom",
						"pluginName":   pluginTestDefaultPluginName,
						"pluginStatus": "INSTALLED",
						"pluginOperation": []map[string]interface{}{
							{"enable": true, "operation": "uninstall", "text": "卸载"},
							{"enable": false, "operation": "set-dictonary", "text": "配置词典"},
						},
					},
				},
			},
		})
	}))
	defer server.Close()

	result, err := GetPluginInfo(cli, testRegion, &GetPluginInfoRequest{ClusterId: pluginTestClusterId})
	if err != nil {
		t.Fatalf("get plugin info failed: %v", err)
	}
	if result.Result == nil || len(result.Result.Plugins) != 1 || len(result.Result.DefaultPlugins) != 1 {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	plugin := result.Result.Plugins[0]
	if plugin.ModuleType != pluginTestModuleType || plugin.PluginName != pluginTestCustomPluginName ||
		plugin.PluginStatus != "UNINSTALLED" || plugin.PluginDesc == "" {
		t.Fatalf("unexpected plugin: %+v", plugin)
	}
	operations := result.Result.DefaultPlugins[0].PluginOperation
	if len(operations) != 2 || operations[0].Operation != "uninstall" || !operations[0].Enable {
		t.Fatalf("unexpected plugin operation: %+v", operations)
	}
	if operations[1].Operation != "set-dictonary" || operations[1].Enable || operations[1].Text == "" {
		t.Fatalf("unexpected plugin operation: %+v", operations)
	}
}

func TestGetNLPDict(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		pluginAssertRequest(t, r, "/api/bes/cluster/nlp_dict/display")
		if got := readJSONBody(t, r)["clusterId"]; got != pluginTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"operationType": "upload_dict",
				"separator":     pluginTestSeparator,
				"fileName":      "test.txt",
				"dictContent":   "中华\t民国\n",
			},
		})
	}))
	defer server.Close()

	result, err := GetNLPDict(cli, testRegion, &GetNLPDictRequest{ClusterId: pluginTestClusterId})
	if err != nil {
		t.Fatalf("get nlp dict failed: %v", err)
	}
	if result.Result == nil || result.Result.OperationType != "upload_dict" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if result.Result.Separator != pluginTestSeparator || result.Result.FileName != "test.txt" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if result.Result.DictContent != "中华\t民国\n" {
		t.Fatalf("unexpected dict content: %q", result.Result.DictContent)
	}
}

// TestUpdateNLPDict is the only plugin API whose multipart body carries plain form fields next to
// the file part, so both the fields and the file are asserted here.
func TestUpdateNLPDict(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		pluginAssertRequest(t, r, "/api/bes/cluster/v2/nlp_dict/update")
		fields, fileContent, fileName := readMultipartBody(t, r, "file")
		if len(fields) != 2 || fields["clusterId"] != pluginTestClusterId {
			t.Fatalf("unexpected form fields: %+v", fields)
		}
		if fields["separator"] != pluginTestSeparator {
			t.Fatalf("unexpected separator field: %+v", fields)
		}
		if fileName != "dict.txt" || string(fileContent) != "中华\t民国\n" {
			t.Fatalf("unexpected file: name=%s content=%s", fileName, fileContent)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  map[string]interface{}{"code": "Success", "message": ""},
		})
	}))
	defer server.Close()

	path := writeTempUploadFile(t, "dict.txt", []byte("中华\t民国\n"))
	result, err := UpdateNLPDict(cli, testRegion, pluginFullUpdateNLPDictRequest(path))
	if err != nil {
		t.Fatalf("update nlp dict failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
	if result.Result == nil || result.Result.Code != "Success" || result.Result.Message != "" {
		t.Fatalf("unexpected result payload: %+v", result.Result)
	}
}

// TestPluginAPIRejectsUnopenableFile covers the openUploadFile error propagation in the two
// multipart plugin APIs. chmod-based unreadable files are avoided on purpose: a root test runner
// can still read them, which makes such a case flaky.
func TestPluginAPIRejectsUnopenableFile(t *testing.T) {
	paths := []struct {
		name     string
		filePath string
	}{
		{name: "missing file", filePath: filepath.Join(t.TempDir(), "does-not-exist.zip")},
		{name: "directory", filePath: t.TempDir()},
	}
	apis := []struct {
		name string
		call func(bce.Client, string) error
	}{
		{"UploadCustomPlugin", func(c bce.Client, filePath string) error {
			_, err := UploadCustomPlugin(c, testRegion, pluginFullUploadRequest(filePath))
			return err
		}},
		{"UpdateNLPDict", func(c bce.Client, filePath string) error {
			_, err := UpdateNLPDict(c, testRegion, pluginFullUpdateNLPDictRequest(filePath))
			return err
		}},
	}
	for _, api := range apis {
		for _, tc := range paths {
			t.Run(api.name+" "+tc.name, func(t *testing.T) {
				cli, server := newTestClient(t, pluginUnreachableHandler(t))
				defer server.Close()

				if err := api.call(cli, tc.filePath); err == nil {
					t.Fatalf("expected error for %s", tc.filePath)
				}
			})
		}
	}
}

func TestPluginAPIRejectsNilClient(t *testing.T) {
	uploadPath := writeTempUploadFile(t, "plugin.zip", []byte("zip-bytes"))
	dictPath := writeTempUploadFile(t, "dict.txt", []byte("中华\t民国\n"))

	cases := []struct {
		name string
		call func() error
	}{
		{"InstallDefaultPlugin", func() error {
			_, err := InstallDefaultPlugin(nil, testRegion, pluginFullDefaultRequest())
			return err
		}},
		{"UninstallDefaultPlugin", func() error {
			_, err := UninstallDefaultPlugin(nil, testRegion, pluginFullDefaultRequest())
			return err
		}},
		{"InstallCustomPlugin", func() error {
			_, err := InstallCustomPlugin(nil, testRegion, pluginFullCustomRequest())
			return err
		}},
		{"UninstallCustomPlugin", func() error {
			_, err := UninstallCustomPlugin(nil, testRegion, pluginFullCustomRequest())
			return err
		}},
		{"DeleteCustomPlugin", func() error {
			_, err := DeleteCustomPlugin(nil, testRegion, pluginFullCustomRequest())
			return err
		}},
		{"UploadCustomPlugin", func() error {
			_, err := UploadCustomPlugin(nil, testRegion, pluginFullUploadRequest(uploadPath))
			return err
		}},
		{"GetPluginInfo", func() error {
			_, err := GetPluginInfo(nil, testRegion, &GetPluginInfoRequest{ClusterId: pluginTestClusterId})
			return err
		}},
		{"GetNLPDict", func() error {
			_, err := GetNLPDict(nil, testRegion, &GetNLPDictRequest{ClusterId: pluginTestClusterId})
			return err
		}},
		{"UpdateNLPDict", func() error {
			_, err := UpdateNLPDict(nil, testRegion, pluginFullUpdateNLPDictRequest(dictPath))
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

func TestPluginAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, pluginUnreachableHandler(t))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"InstallDefaultPlugin", func(c bce.Client) error {
			_, err := InstallDefaultPlugin(c, testRegion, nil)
			return err
		}},
		{"UninstallDefaultPlugin", func(c bce.Client) error {
			_, err := UninstallDefaultPlugin(c, testRegion, nil)
			return err
		}},
		{"InstallCustomPlugin", func(c bce.Client) error {
			_, err := InstallCustomPlugin(c, testRegion, nil)
			return err
		}},
		{"UninstallCustomPlugin", func(c bce.Client) error {
			_, err := UninstallCustomPlugin(c, testRegion, nil)
			return err
		}},
		{"DeleteCustomPlugin", func(c bce.Client) error {
			_, err := DeleteCustomPlugin(c, testRegion, nil)
			return err
		}},
		{"UploadCustomPlugin", func(c bce.Client) error {
			_, err := UploadCustomPlugin(c, testRegion, nil)
			return err
		}},
		{"GetPluginInfo", func(c bce.Client) error {
			_, err := GetPluginInfo(c, testRegion, nil)
			return err
		}},
		{"GetNLPDict", func(c bce.Client) error {
			_, err := GetNLPDict(c, testRegion, nil)
			return err
		}},
		{"UpdateNLPDict", func(c bce.Client) error {
			_, err := UpdateNLPDict(c, testRegion, nil)
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

// TestPluginAPIRejectsMissingRequiredFields covers the 14 required-field branches of plugin.go:
// three in checkDefaultPluginRequest, three in checkCustomPluginRequest, and eight inline in
// UploadCustomPlugin (3), GetPluginInfo (1), GetNLPDict (1) and UpdateNLPDict (3). The two shared
// validators are entered through one API each here to prove the error is propagated end to end;
// TestPluginCheckDefaultPluginRequest and TestPluginCheckCustomPluginRequest then cover them
// directly for both action strings, instead of duplicating five more end-to-end cases.
func TestPluginAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, pluginUnreachableHandler(t))
	defer server.Close()
	uploadPath := writeTempUploadFile(t, "plugin.zip", []byte("zip-bytes"))
	dictPath := writeTempUploadFile(t, "dict.txt", []byte("中华\t民国\n"))

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"InstallDefaultPlugin empty clusterId", func(c bce.Client) error {
			request := pluginFullDefaultRequest()
			request.ClusterId = ""
			_, err := InstallDefaultPlugin(c, testRegion, request)
			return err
		}},
		{"InstallDefaultPlugin empty pluginName", func(c bce.Client) error {
			request := pluginFullDefaultRequest()
			request.PluginName = ""
			_, err := InstallDefaultPlugin(c, testRegion, request)
			return err
		}},
		{"InstallDefaultPlugin empty moduleType", func(c bce.Client) error {
			request := pluginFullDefaultRequest()
			request.ModuleType = ""
			_, err := InstallDefaultPlugin(c, testRegion, request)
			return err
		}},
		{"InstallCustomPlugin empty clusterId", func(c bce.Client) error {
			request := pluginFullCustomRequest()
			request.ClusterId = ""
			_, err := InstallCustomPlugin(c, testRegion, request)
			return err
		}},
		{"InstallCustomPlugin empty pluginName", func(c bce.Client) error {
			request := pluginFullCustomRequest()
			request.PluginName = ""
			_, err := InstallCustomPlugin(c, testRegion, request)
			return err
		}},
		{"InstallCustomPlugin empty moduleType", func(c bce.Client) error {
			request := pluginFullCustomRequest()
			request.ModuleType = ""
			_, err := InstallCustomPlugin(c, testRegion, request)
			return err
		}},
		{"UploadCustomPlugin empty clusterId", func(c bce.Client) error {
			request := pluginFullUploadRequest(uploadPath)
			request.ClusterId = ""
			_, err := UploadCustomPlugin(c, testRegion, request)
			return err
		}},
		{"UploadCustomPlugin empty moduleType", func(c bce.Client) error {
			request := pluginFullUploadRequest(uploadPath)
			request.ModuleType = ""
			_, err := UploadCustomPlugin(c, testRegion, request)
			return err
		}},
		{"UploadCustomPlugin empty filePath", func(c bce.Client) error {
			request := pluginFullUploadRequest(uploadPath)
			request.FilePath = ""
			_, err := UploadCustomPlugin(c, testRegion, request)
			return err
		}},
		{"GetPluginInfo empty clusterId", func(c bce.Client) error {
			_, err := GetPluginInfo(c, testRegion, &GetPluginInfoRequest{})
			return err
		}},
		{"GetNLPDict empty clusterId", func(c bce.Client) error {
			_, err := GetNLPDict(c, testRegion, &GetNLPDictRequest{})
			return err
		}},
		{"UpdateNLPDict empty clusterId", func(c bce.Client) error {
			request := pluginFullUpdateNLPDictRequest(dictPath)
			request.ClusterId = ""
			_, err := UpdateNLPDict(c, testRegion, request)
			return err
		}},
		{"UpdateNLPDict empty separator", func(c bce.Client) error {
			request := pluginFullUpdateNLPDictRequest(dictPath)
			request.Separator = ""
			_, err := UpdateNLPDict(c, testRegion, request)
			return err
		}},
		{"UpdateNLPDict empty filePath", func(c bce.Client) error {
			request := pluginFullUpdateNLPDictRequest(dictPath)
			request.FilePath = ""
			_, err := UpdateNLPDict(c, testRegion, request)
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

func TestPluginAPIPropagatesServerError(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		pluginWriteBadRequest(t, w)
	}))
	defer server.Close()
	uploadPath := writeTempUploadFile(t, "plugin.zip", []byte("zip-bytes"))
	dictPath := writeTempUploadFile(t, "dict.txt", []byte("中华\t民国\n"))

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"InstallDefaultPlugin", func(c bce.Client) error {
			_, err := InstallDefaultPlugin(c, testRegion, pluginFullDefaultRequest())
			return err
		}},
		{"UninstallDefaultPlugin", func(c bce.Client) error {
			_, err := UninstallDefaultPlugin(c, testRegion, pluginFullDefaultRequest())
			return err
		}},
		{"InstallCustomPlugin", func(c bce.Client) error {
			_, err := InstallCustomPlugin(c, testRegion, pluginFullCustomRequest())
			return err
		}},
		{"UninstallCustomPlugin", func(c bce.Client) error {
			_, err := UninstallCustomPlugin(c, testRegion, pluginFullCustomRequest())
			return err
		}},
		{"DeleteCustomPlugin", func(c bce.Client) error {
			_, err := DeleteCustomPlugin(c, testRegion, pluginFullCustomRequest())
			return err
		}},
		{"UploadCustomPlugin", func(c bce.Client) error {
			_, err := UploadCustomPlugin(c, testRegion, pluginFullUploadRequest(uploadPath))
			return err
		}},
		{"GetPluginInfo", func(c bce.Client) error {
			_, err := GetPluginInfo(c, testRegion, &GetPluginInfoRequest{ClusterId: pluginTestClusterId})
			return err
		}},
		{"GetNLPDict", func(c bce.Client) error {
			_, err := GetNLPDict(c, testRegion, &GetNLPDictRequest{ClusterId: pluginTestClusterId})
			return err
		}},
		{"UpdateNLPDict", func(c bce.Client) error {
			_, err := UpdateNLPDict(c, testRegion, pluginFullUpdateNLPDictRequest(dictPath))
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

// TestPluginCheckDefaultPluginRequest covers checkDefaultPluginRequest directly. The validator is
// shared by InstallDefaultPlugin and UninstallDefaultPlugin, so both action strings are exercised
// here rather than through duplicated end-to-end cases; the message prefix is asserted because it
// is the only thing that tells the two callers apart.
func TestPluginCheckDefaultPluginRequest(t *testing.T) {
	for _, action := range []string{"install default plugin", "uninstall default plugin"} {
		cases := []struct {
			name    string
			request *DefaultPluginRequest
			want    string
		}{
			{"nil request", nil, action + " request should not be nil"},
			{"empty clusterId", &DefaultPluginRequest{
				PluginName: pluginTestDefaultPluginName,
				ModuleType: pluginTestModuleType,
			}, action + " request clusterId should not be empty"},
			{"empty pluginName", &DefaultPluginRequest{
				ClusterId:  pluginTestClusterId,
				ModuleType: pluginTestModuleType,
			}, action + " request pluginName should not be empty"},
			{"empty moduleType", &DefaultPluginRequest{
				ClusterId:  pluginTestClusterId,
				PluginName: pluginTestDefaultPluginName,
			}, action + " request moduleType should not be empty"},
		}
		for _, tc := range cases {
			t.Run(action+" "+tc.name, func(t *testing.T) {
				err := checkDefaultPluginRequest(tc.request, action)
				if err == nil {
					t.Fatal("expected a validation error")
				}
				if err.Error() != tc.want {
					t.Fatalf("unexpected message: %q, want %q", err.Error(), tc.want)
				}
			})
		}
	}

	if err := checkDefaultPluginRequest(pluginFullDefaultRequest(), "install default plugin"); err != nil {
		t.Fatalf("unexpected error for a complete request: %v", err)
	}
}

// TestPluginCheckCustomPluginRequest covers checkCustomPluginRequest directly. This validator is
// shared by three APIs (install, uninstall and delete custom plugin), so all three action strings
// are exercised here.
func TestPluginCheckCustomPluginRequest(t *testing.T) {
	actions := []string{"install custom plugin", "uninstall custom plugin", "delete custom plugin"}
	for _, action := range actions {
		cases := []struct {
			name    string
			request *CustomPluginRequest
			want    string
		}{
			{"nil request", nil, action + " request should not be nil"},
			{"empty clusterId", &CustomPluginRequest{
				PluginName: pluginTestCustomPluginName,
				ModuleType: pluginTestModuleType,
			}, action + " request clusterId should not be empty"},
			{"empty pluginName", &CustomPluginRequest{
				ClusterId:  pluginTestClusterId,
				ModuleType: pluginTestModuleType,
			}, action + " request pluginName should not be empty"},
			{"empty moduleType", &CustomPluginRequest{
				ClusterId:  pluginTestClusterId,
				PluginName: pluginTestCustomPluginName,
			}, action + " request moduleType should not be empty"},
		}
		for _, tc := range cases {
			t.Run(action+" "+tc.name, func(t *testing.T) {
				err := checkCustomPluginRequest(tc.request, action)
				if err == nil {
					t.Fatal("expected a validation error")
				}
				if err.Error() != tc.want {
					t.Fatalf("unexpected message: %q, want %q", err.Error(), tc.want)
				}
			})
		}
	}

	if err := checkCustomPluginRequest(pluginFullCustomRequest(), "install custom plugin"); err != nil {
		t.Fatalf("unexpected error for a complete request: %v", err)
	}
}

// TestPluginBuildMultipartBodyRejectsNilFile covers the only buildMultipartBody guard that no
// plugin API can reach: openUploadFile always hands it a non-nil file, so the branch is exercised
// directly instead of adding a production seam.
func TestPluginBuildMultipartBodyRejectsNilFile(t *testing.T) {
	body, contentType, err := buildMultipartBody(nil, "file", nil, "plugin.zip", 0)
	if err == nil {
		t.Fatal("expected an error for a nil file")
	}
	if body != nil || contentType != "" {
		t.Fatalf("unexpected body/contentType on error: %v %q", body, contentType)
	}
}

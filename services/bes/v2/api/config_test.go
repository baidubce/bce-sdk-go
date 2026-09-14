package api

import (
	"errors"
	nethttp "net/http"
	"path/filepath"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

const configTestClusterId = "644734225693675520"

// configWriteBadRequest writes a 400 service error. A 4xx status is used on purpose: the test
// client mirrors the production DEFAULT_RETRY_POLICY, so a 5xx would trigger backoff retries.
func configWriteBadRequest(t *testing.T, w nethttp.ResponseWriter) {
	t.Helper()
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(nethttp.StatusBadRequest)
	if _, err := w.Write([]byte(`{"code":"BadRequest","message":"invalid config request"}`)); err != nil {
		t.Fatalf("write error response failed: %v", err)
	}
}

// configUnreachableHandler fails the test when a request reaches the server, used by cases that
// must fail during local validation.
func configUnreachableHandler(t *testing.T) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Errorf("unexpected request reached the server: %s", r.URL.Path)
		configWriteBadRequest(t, w)
	}
}

func TestGetClusterConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if r.URL.Path != "/api/bes/cluster/config/info" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("x-Region"); got != testRegion {
			t.Fatalf("unexpected region header: %s", got)
		}
		if got := readJSONBody(t, r)["clusterId"]; got != configTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result": map[string]interface{}{
				"esNodeExtraConfig": "cluster.publish.timeout: 180s",
				"kibanaExtraConfig": "ops.interval: 100",
			},
		})
	}))
	defer server.Close()

	result, err := GetClusterConfig(cli, testRegion, &GetClusterConfigRequest{ClusterId: configTestClusterId})
	if err != nil {
		t.Fatalf("get cluster config failed: %v", err)
	}
	if result.Result == nil || result.Result.EsNodeExtraConfig != "cluster.publish.timeout: 180s" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
	if result.Result.KibanaExtraConfig != "ops.interval: 100" {
		t.Fatalf("unexpected kibana config: %s", result.Result.KibanaExtraConfig)
	}
}

func TestUpdateClusterConfig(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if r.URL.Path != "/api/bes/cluster/config/update" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("x-Region"); got != testRegion {
			t.Fatalf("unexpected region header: %s", got)
		}
		body := readJSONBody(t, r)
		esConfigs, ok := body["esConfigsMap"].(map[string]interface{})
		if !ok {
			t.Fatalf("unexpected esConfigsMap: %v", body["esConfigsMap"])
		}
		if esConfigs["cluster.publish.timeout"] != "180s" {
			t.Fatalf("unexpected esConfigsMap content: %+v", esConfigs)
		}
		kibanaConfigs, ok := body["kibanaConfigsMap"].(map[string]interface{})
		if !ok {
			t.Fatalf("unexpected kibanaConfigsMap: %v", body["kibanaConfigsMap"])
		}
		if kibanaConfigs["ops.interval"] != float64(100) {
			t.Fatalf("unexpected kibanaConfigsMap content: %+v", kibanaConfigs)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": map[string]interface{}{}})
	}))
	defer server.Close()

	result, err := UpdateClusterConfig(cli, testRegion, &UpdateClusterConfigRequest{
		ClusterId:        configTestClusterId,
		EsConfigsMap:     map[string]interface{}{"cluster.publish.timeout": "180s"},
		KibanaConfigsMap: map[string]interface{}{"ops.interval": 100},
	})
	if err != nil {
		t.Fatalf("update cluster config failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestListSynonymDicts(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if r.URL.Path != "/api/bes/cluster/synonym_dict/list" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("x-Region"); got != testRegion {
			t.Fatalf("unexpected region header: %s", got)
		}
		if got := readJSONBody(t, r)["clusterId"]; got != configTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{
			"success": true,
			"status":  200,
			"result":  []map[string]interface{}{{"synonym_dict_name": "xxx.txt"}},
		})
	}))
	defer server.Close()

	result, err := ListSynonymDicts(cli, testRegion, &ListSynonymDictsRequest{ClusterId: configTestClusterId})
	if err != nil {
		t.Fatalf("list synonym dicts failed: %v", err)
	}
	if len(result.Result) != 1 || result.Result[0].SynonymDictName != "xxx.txt" {
		t.Fatalf("unexpected result: %+v", result.Result)
	}
}

func TestDeleteSynonymDict(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if r.URL.Path != "/api/bes/cluster/synonym_dict/delete" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("x-Region"); got != testRegion {
			t.Fatalf("unexpected region header: %s", got)
		}
		body := readJSONBody(t, r)
		if got := body["clusterId"]; got != configTestClusterId {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		if got := body["synonymDictName"]; got != "xxx.txt" {
			t.Fatalf("unexpected synonymDictName: %v", got)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": map[string]interface{}{}})
	}))
	defer server.Close()

	result, err := DeleteSynonymDict(cli, testRegion, &DeleteSynonymDictRequest{
		ClusterId:       configTestClusterId,
		SynonymDictName: "xxx.txt",
	})
	if err != nil {
		t.Fatalf("delete synonym dict failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
}

func TestUploadSynonymDict(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		if r.URL.Path != "/api/bes/cluster/"+configTestClusterId+"/synonym_dict/upload" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("x-Region"); got != testRegion {
			t.Fatalf("unexpected region header: %s", got)
		}
		fields, fileContent, fileName := readMultipartBody(t, r, "file")
		if len(fields) != 0 {
			t.Fatalf("unexpected extra form fields: %+v", fields)
		}
		if fileName != "syn.txt" || string(fileContent) != "hello=hi\n" {
			t.Fatalf("unexpected file: name=%s content=%s", fileName, fileContent)
		}
		writeJSONResponse(t, w, map[string]interface{}{"success": true, "status": 200, "result": map[string]interface{}{}})
	}))
	defer server.Close()

	path := writeTempUploadFile(t, "syn.txt", []byte("hello=hi\n"))
	result, err := UploadSynonymDict(cli, testRegion, &UploadSynonymDictRequest{
		ClusterId: configTestClusterId,
		FilePath:  path,
	})
	if err != nil {
		t.Fatalf("upload synonym dict failed: %v", err)
	}
	if !result.Success || result.Status != 200 {
		t.Fatalf("unexpected response: %+v", result)
	}
}

// TestUploadSynonymDictRejectsUnopenableFile covers the openUploadFile error propagation in
// UploadSynonymDict. chmod-based unreadable files are avoided on purpose: a root test runner can
// still read them, which makes such a case flaky.
func TestUploadSynonymDictRejectsUnopenableFile(t *testing.T) {
	cases := []struct {
		name     string
		filePath string
	}{
		{name: "missing file", filePath: filepath.Join(t.TempDir(), "does-not-exist.txt")},
		{name: "directory", filePath: t.TempDir()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cli, server := newTestClient(t, configUnreachableHandler(t))
			defer server.Close()

			if _, err := UploadSynonymDict(cli, testRegion, &UploadSynonymDictRequest{
				ClusterId: configTestClusterId,
				FilePath:  tc.filePath,
			}); err == nil {
				t.Fatalf("expected error for %s", tc.filePath)
			}
		})
	}
}

func TestConfigAPIRejectsNilClient(t *testing.T) {
	cases := []struct {
		name string
		call func() error
	}{
		{"GetClusterConfig", func() error {
			_, err := GetClusterConfig(nil, testRegion, &GetClusterConfigRequest{ClusterId: configTestClusterId})
			return err
		}},
		{"UpdateClusterConfig", func() error {
			_, err := UpdateClusterConfig(nil, testRegion, &UpdateClusterConfigRequest{
				ClusterId:        configTestClusterId,
				EsConfigsMap:     map[string]interface{}{},
				KibanaConfigsMap: map[string]interface{}{},
			})
			return err
		}},
		{"ListSynonymDicts", func() error {
			_, err := ListSynonymDicts(nil, testRegion, &ListSynonymDictsRequest{ClusterId: configTestClusterId})
			return err
		}},
		{"DeleteSynonymDict", func() error {
			_, err := DeleteSynonymDict(nil, testRegion, &DeleteSynonymDictRequest{
				ClusterId:       configTestClusterId,
				SynonymDictName: "xxx.txt",
			})
			return err
		}},
		{"UploadSynonymDict", func() error {
			_, err := UploadSynonymDict(nil, testRegion, &UploadSynonymDictRequest{
				ClusterId: configTestClusterId,
				FilePath:  "syn.txt",
			})
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

func TestConfigAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, configUnreachableHandler(t))
	defer server.Close()

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"GetClusterConfig", func(c bce.Client) error {
			_, err := GetClusterConfig(c, testRegion, nil)
			return err
		}},
		{"UpdateClusterConfig", func(c bce.Client) error {
			_, err := UpdateClusterConfig(c, testRegion, nil)
			return err
		}},
		{"ListSynonymDicts", func(c bce.Client) error {
			_, err := ListSynonymDicts(c, testRegion, nil)
			return err
		}},
		{"DeleteSynonymDict", func(c bce.Client) error {
			_, err := DeleteSynonymDict(c, testRegion, nil)
			return err
		}},
		{"UploadSynonymDict", func(c bce.Client) error {
			_, err := UploadSynonymDict(c, testRegion, nil)
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

// configFullUpdateRequest / configFullDeleteRequest / configFullUploadRequest return fully
// populated requests, so the required-field cases below only have to clear one field each.
func configFullUpdateRequest() *UpdateClusterConfigRequest {
	return &UpdateClusterConfigRequest{
		ClusterId:        configTestClusterId,
		EsConfigsMap:     map[string]interface{}{"cluster.publish.timeout": "180s"},
		KibanaConfigsMap: map[string]interface{}{"ops.interval": 100},
	}
}

func configFullDeleteRequest() *DeleteSynonymDictRequest {
	return &DeleteSynonymDictRequest{ClusterId: configTestClusterId, SynonymDictName: "xxx.txt"}
}

func configFullUploadRequest(filePath string) *UploadSynonymDictRequest {
	return &UploadSynonymDictRequest{ClusterId: configTestClusterId, FilePath: filePath}
}

func TestConfigAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, configUnreachableHandler(t))
	defer server.Close()

	uploadPath := writeTempUploadFile(t, "syn.txt", []byte("hello=hi\n"))

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"GetClusterConfig empty clusterId", func(c bce.Client) error {
			_, err := GetClusterConfig(c, testRegion, &GetClusterConfigRequest{})
			return err
		}},
		{"UpdateClusterConfig empty clusterId", func(c bce.Client) error {
			request := configFullUpdateRequest()
			request.ClusterId = ""
			_, err := UpdateClusterConfig(c, testRegion, request)
			return err
		}},
		{"UpdateClusterConfig nil esConfigsMap", func(c bce.Client) error {
			request := configFullUpdateRequest()
			request.EsConfigsMap = nil
			_, err := UpdateClusterConfig(c, testRegion, request)
			return err
		}},
		{"UpdateClusterConfig nil kibanaConfigsMap", func(c bce.Client) error {
			request := configFullUpdateRequest()
			request.KibanaConfigsMap = nil
			_, err := UpdateClusterConfig(c, testRegion, request)
			return err
		}},
		{"ListSynonymDicts empty clusterId", func(c bce.Client) error {
			_, err := ListSynonymDicts(c, testRegion, &ListSynonymDictsRequest{})
			return err
		}},
		{"DeleteSynonymDict empty clusterId", func(c bce.Client) error {
			request := configFullDeleteRequest()
			request.ClusterId = ""
			_, err := DeleteSynonymDict(c, testRegion, request)
			return err
		}},
		{"DeleteSynonymDict empty synonymDictName", func(c bce.Client) error {
			request := configFullDeleteRequest()
			request.SynonymDictName = ""
			_, err := DeleteSynonymDict(c, testRegion, request)
			return err
		}},
		{"UploadSynonymDict empty clusterId", func(c bce.Client) error {
			request := configFullUploadRequest(uploadPath)
			request.ClusterId = ""
			_, err := UploadSynonymDict(c, testRegion, request)
			return err
		}},
		{"UploadSynonymDict empty filePath", func(c bce.Client) error {
			request := configFullUploadRequest(uploadPath)
			request.FilePath = ""
			_, err := UploadSynonymDict(c, testRegion, request)
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

func TestConfigAPIPropagatesServerError(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		configWriteBadRequest(t, w)
	}))
	defer server.Close()

	uploadPath := writeTempUploadFile(t, "syn.txt", []byte("hello=hi\n"))

	cases := []struct {
		name string
		call func(bce.Client) error
	}{
		{"GetClusterConfig", func(c bce.Client) error {
			_, err := GetClusterConfig(c, testRegion, &GetClusterConfigRequest{ClusterId: configTestClusterId})
			return err
		}},
		{"UpdateClusterConfig", func(c bce.Client) error {
			_, err := UpdateClusterConfig(c, testRegion, configFullUpdateRequest())
			return err
		}},
		{"ListSynonymDicts", func(c bce.Client) error {
			_, err := ListSynonymDicts(c, testRegion, &ListSynonymDictsRequest{ClusterId: configTestClusterId})
			return err
		}},
		{"DeleteSynonymDict", func(c bce.Client) error {
			_, err := DeleteSynonymDict(c, testRegion, configFullDeleteRequest())
			return err
		}},
		{"UploadSynonymDict", func(c bce.Client) error {
			_, err := UploadSynonymDict(c, testRegion, configFullUploadRequest(uploadPath))
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

package eval

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
)

// capturedRequest 假服务端收到的请求
type capturedRequest struct {
	method  string
	path    string
	query   map[string]string
	version string
	body    map[string]interface{}
}

// testServers 本包测试起的假服务端，TestMain 结束时统一关闭（流水线工具链是 Go 1.13，不能用 1.14 起才有的测试清理钩子）
var testServers []*httptest.Server

func TestMain(m *testing.M) {
	code := m.Run()
	for _, s := range testServers {
		s.Close()
	}
	os.Exit(code)
}

// newTestClient 起一个假服务端：记录请求并返回 status/body，client 不签名、不重试
func newTestClient(t *testing.T, status int, respBody string) (*Client, *capturedRequest) {
	t.Helper()
	got := &capturedRequest{query: map[string]string{}}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.method = r.Method
		got.path = r.URL.Path
		for k := range r.URL.Query() {
			got.query[k] = r.URL.Query().Get(k)
		}
		got.version = r.Header.Get("version")
		if b, _ := ioutil.ReadAll(r.Body); len(b) > 0 {
			if err := json.Unmarshal(b, &got.body); err != nil {
				t.Errorf("request body is not json: %s", string(b))
			}
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, respBody)
	}))
	testServers = append(testServers, server)

	c, err := NewClient("ak", "sk", server.URL)
	if err != nil {
		t.Fatalf("NewClient failed: %v", err)
	}
	c.DefaultClient.Config.Credentials = nil
	c.DefaultClient.Config.Retry = bce.NewNoRetryPolicy()
	return c, got
}

// expectRequest 校验方法、路径、version 头和 query；query 里没列出的键必须不存在
func expectRequest(t *testing.T, got *capturedRequest, method string, query map[string]string) {
	t.Helper()
	if got.method != method {
		t.Errorf("method = %s, want %s", got.method, method)
	}
	if got.path != DefaultUrl {
		t.Errorf("path = %s, want %s", got.path, DefaultUrl)
	}
	if got.version != DefaultVersion {
		t.Errorf("version header = %q, want %q", got.version, DefaultVersion)
	}
	if len(got.query) != len(query) {
		t.Errorf("query = %v, want %v", got.query, query)
	}
	for k, want := range query {
		if got.query[k] != want {
			t.Errorf("query %s = %q, want %q", k, got.query[k], want)
		}
	}
}

func TestCreateEmbodiedEval(t *testing.T) {
	c, got := newTestClient(t, 200, `{"requestId":"r-1","evalId":"em-0123456789ab"}`)
	scope := int32(0)
	res, err := c.CreateEmbodiedEval(&CreateEmbodiedEvalArgs{
		ClientToken:    "token-1",
		Name:           "eval-1",
		ResourcePoolId: "aihc-serverless",
		QueueId:        "q-1",
		VisibleScope:   &scope,
		ModelConfig: &ModelConfig{
			ModelType: "customImage", Replicas: 1, MaxBatchSize: 1,
			Resources: []Resource{{Name: "cpu", Quantity: 8}},
			CustomImage: &CustomImage{Image: "ccr-registry.baidubce.com/ns/vla:v1", Command: "serve",
				ImageConfig: &ImageConfig{Username: "u", Password: "p"}},
		},
		BenchmarkConfigs: []BenchmarkConfig{{
			Name: "libero", SubBench: "libero_spatial", ParallelShards: 1, EpisodesPerTask: 10,
			Params:        map[string]interface{}{"seed": 7},
			ModelOverride: &ModelOverride{ChunkSize: 1},
		}},
	})
	if err != nil {
		t.Fatalf("CreateEmbodiedEval failed: %v", err)
	}
	if res.RequestId != "r-1" || res.EvalId != "em-0123456789ab" {
		t.Errorf("unexpected result: %+v", res)
	}
	expectRequest(t, got, http.MethodPost, map[string]string{"action": ApiCreateEmbodiedEval, "clientToken": "token-1"})
	if _, ok := got.body["clientToken"]; ok {
		t.Errorf("clientToken must only be sent in query, body=%v", got.body)
	}
	if got.body["name"] != "eval-1" || got.body["queueId"] != "q-1" {
		t.Errorf("unexpected body: %v", got.body)
	}
	// 显式传 0 必须下发，否则服务端按默认 1 处理
	if v, ok := got.body["visibleScope"]; !ok || v != float64(0) {
		t.Errorf("visibleScope = %v, want 0", v)
	}
	bench := got.body["benchmarkConfigs"].([]interface{})[0].(map[string]interface{})
	if params, ok := bench["params"].(map[string]interface{}); !ok || params["seed"] != float64(7) {
		t.Errorf("params should be sent as object, got %v", bench["params"])
	}
}

func TestCreateEmbodiedEvalWithoutClientToken(t *testing.T) {
	saved := newClientToken
	defer func() { newClientToken = saved }()
	newClientToken = func() string { return "auto-token" }

	c, got := newTestClient(t, 200, `{"requestId":"r-1","evalId":"em-1"}`)
	args := &CreateEmbodiedEvalArgs{Name: "eval-1"}
	if _, err := c.CreateEmbodiedEval(args); err != nil {
		t.Fatalf("CreateEmbodiedEval failed: %v", err)
	}
	// 不传时自动生成，SDK 重试复用同一个请求，不会重复建任务
	expectRequest(t, got, http.MethodPost, map[string]string{"action": ApiCreateEmbodiedEval, "clientToken": "auto-token"})
	if args.ClientToken != "" {
		t.Errorf("generated clientToken must not be written back to args, got %q", args.ClientToken)
	}
	if _, ok := got.body["visibleScope"]; ok {
		t.Errorf("visibleScope not set should be omitted, body=%v", got.body)
	}
}

func TestDescribeEmbodiedEvals(t *testing.T) {
	c, got := newTestClient(t, 200,
		`{"requestId":"r-1","totalCount":1,"evals":[{"evalId":"em-1","status":"Running","benchmarks":[{"name":"libero","subBench":"libero_10"}]}]}`)
	res, err := c.DescribeEmbodiedEvals(&DescribeEmbodiedEvalsArgs{
		PageNumber: 2, PageSize: 20, KeywordType: "name", Keyword: "eval",
		Status: "Running,Failed", ModelName: "openvla_7b", Priority: "high", OrderBy: "createdAt", Order: "DESC",
	})
	if err != nil {
		t.Fatalf("DescribeEmbodiedEvals failed: %v", err)
	}
	expectRequest(t, got, http.MethodGet, map[string]string{
		"action": ApiDescribeEmbodiedEvals, "pageNumber": "2", "pageSize": "20", "keywordType": "name", "keyword": "eval",
		"status": "Running,Failed", "modelName": "openvla_7b", "priority": "high", "orderBy": "createdAt", "order": "DESC",
	})
	if res.TotalCount != 1 || len(res.Evals) != 1 || res.Evals[0].Benchmarks[0].SubBench != "libero_10" {
		t.Errorf("unexpected result: %+v", res)
	}
}

func TestDescribeEmbodiedEvalsDefaults(t *testing.T) {
	c, got := newTestClient(t, 200, `{"requestId":"r-1","totalCount":0,"evals":[]}`)
	if _, err := c.DescribeEmbodiedEvals(nil); err != nil {
		t.Fatalf("DescribeEmbodiedEvals failed: %v", err)
	}
	// 不传的参数不下发，由服务端取默认值
	expectRequest(t, got, http.MethodGet, map[string]string{"action": ApiDescribeEmbodiedEvals})
}

func TestDescribeEmbodiedEval(t *testing.T) {
	c, got := newTestClient(t, 200, `{"requestId":"r-1","evalId":"em-1","status":"Succeeded",
		"modelConfig":{"modelType":"customImage","customImage":{"image":"img","imageConfig":{"username":"u","password":"******"}}},
		"benchmarkConfigs":[{"name":"libero","params":{"seed":7},"modelOverride":{"chunkSize":1,"extraArgs":{"k":"v"}}}]}`)
	res, err := c.DescribeEmbodiedEval("em-1")
	if err != nil {
		t.Fatalf("DescribeEmbodiedEval failed: %v", err)
	}
	expectRequest(t, got, http.MethodGet, map[string]string{"action": ApiDescribeEmbodiedEval, "evalId": "em-1"})
	if res.EvalId != "em-1" || res.ModelConfig.CustomImage.ImageConfig.Username != "u" {
		t.Errorf("unexpected result: %+v", res)
	}
	if res.BenchmarkConfigs[0].ModelOverride.ExtraArgs.(map[string]interface{})["k"] != "v" {
		t.Errorf("extraArgs not parsed: %+v", res.BenchmarkConfigs[0].ModelOverride)
	}
}

func TestStopAndDeleteEmbodiedEval(t *testing.T) {
	c, got := newTestClient(t, 200, `{"requestId":"r-1","evalId":"em-1"}`)
	if _, err := c.StopEmbodiedEval(&StopEmbodiedEvalArgs{EvalId: "em-1", ClientToken: "t-1"}); err != nil {
		t.Fatalf("StopEmbodiedEval failed: %v", err)
	}
	expectRequest(t, got, http.MethodPost, map[string]string{"action": ApiStopEmbodiedEval, "evalId": "em-1", "clientToken": "t-1"})
	if got.body != nil {
		t.Errorf("stop should not send body, got %v", got.body)
	}

	c, got = newTestClient(t, 200, `{"requestId":"r-2","evalId":"em-1"}`)
	res, err := c.DeleteEmbodiedEval(&DeleteEmbodiedEvalArgs{EvalId: "em-1"})
	if err != nil {
		t.Fatalf("DeleteEmbodiedEval failed: %v", err)
	}
	expectRequest(t, got, http.MethodPost, map[string]string{"action": ApiDeleteEmbodiedEval, "evalId": "em-1"})
	if res.RequestId != "r-2" || res.EvalId != "em-1" {
		t.Errorf("unexpected result: %+v", res)
	}
}

func TestDescribeEmbodiedEvalResultAndVideo(t *testing.T) {
	c, got := newTestClient(t, 200, `{"requestId":"r-1","evalId":"em-1","result":null}`)
	res, err := c.DescribeEmbodiedEvalResult("em-1")
	if err != nil {
		t.Fatalf("DescribeEmbodiedEvalResult failed: %v", err)
	}
	expectRequest(t, got, http.MethodGet, map[string]string{"action": ApiDescribeEmbodiedEvalResult, "evalId": "em-1"})
	if res.Result != nil {
		t.Errorf("result not ready should be nil, got %v", res.Result)
	}

	c, got = newTestClient(t, 200,
		`{"requestId":"r-1","videoId":"v-1","streamUrl":"https://bos/x.mp4","expiresIn":900,"mimeType":"video/mp4","sizeBytes":1024}`)
	video, err := c.DescribeEmbodiedEvalVideo(&DescribeEmbodiedEvalVideoArgs{EvalId: "em-1", VideoId: "v-1"})
	if err != nil {
		t.Fatalf("DescribeEmbodiedEvalVideo failed: %v", err)
	}
	expectRequest(t, got, http.MethodGet, map[string]string{"action": ApiDescribeEmbodiedEvalVideo, "evalId": "em-1", "videoId": "v-1"})
	if video.ExpiresIn != 900 || video.SizeBytes != 1024 || video.StreamUrl == "" {
		t.Errorf("unexpected result: %+v", video)
	}
}

func TestDescribeEmbodiedEvalMetadata(t *testing.T) {
	c, got := newTestClient(t, 200,
		`{"requestId":"r-1","models":[{"modelName":"openvla_7b","benchmarks":[{"name":"libero","subBenches":["libero_spatial"]}]}]}`)
	models, err := c.DescribeEmbodiedEvalModels(&DescribeEmbodiedEvalModelsArgs{Source: "custom"})
	if err != nil {
		t.Fatalf("DescribeEmbodiedEvalModels failed: %v", err)
	}
	expectRequest(t, got, http.MethodGet, map[string]string{"action": ApiDescribeEmbodiedEvalModels, "source": "custom"})
	if models.Models[0].Benchmarks[0].SubBenches[0] != "libero_spatial" {
		t.Errorf("unexpected result: %+v", models)
	}

	c, got = newTestClient(t, 200, `{"requestId":"r-1","modelName":"openvla_7b",
		"minHardware":{"acceleratorTypes":["baidu.com/a800_80g_cgpu"],"acceleratorCount":1,"cpu":8,"memory":64},
		"defaultModelConfig":{"modelType":"public","replicas":1},"chunkSizeLimit":8}`)
	model, err := c.DescribeEmbodiedEvalModel(&DescribeEmbodiedEvalModelArgs{ModelName: "openvla_7b"})
	if err != nil {
		t.Fatalf("DescribeEmbodiedEvalModel failed: %v", err)
	}
	expectRequest(t, got, http.MethodGet, map[string]string{"action": ApiDescribeEmbodiedEvalModel, "modelName": "openvla_7b"})
	if model.MinHardware.Memory != 64 || model.DefaultModelConfig.ModelType != "public" || model.ChunkSizeLimit != 8 {
		t.Errorf("unexpected result: %+v", model)
	}

	c, got = newTestClient(t, 200, `{"requestId":"r-1","benchmarks":[{"modelName":"openvla_7b","name":"libero",
		"subBench":"libero_spatial","defaultBenchmarkConfig":{"name":"libero","parallelShards":1},"maxEpisodesPerTask":50}]}`)
	benches, err := c.DescribeEmbodiedEvalBenchmarks(&DescribeEmbodiedEvalBenchmarksArgs{
		ModelName: "openvla_7b", BenchmarkName: "libero", SubBench: "libero_spatial", Source: "public",
	})
	if err != nil {
		t.Fatalf("DescribeEmbodiedEvalBenchmarks failed: %v", err)
	}
	expectRequest(t, got, http.MethodGet, map[string]string{
		"action": ApiDescribeEmbodiedEvalBenchmarks, "modelName": "openvla_7b", "benchmarkName": "libero",
		"subBench": "libero_spatial", "source": "public",
	})
	if benches.Benchmarks[0].MaxEpisodesPerTask != 50 || benches.Benchmarks[0].DefaultBenchmarkConfig.ParallelShards != 1 {
		t.Errorf("unexpected result: %+v", benches)
	}
}

// 错误体是 {requestId, code, message}：保留 HTTP 状态码、code、message 和 requestId
func TestEmbodiedEvalServiceError(t *testing.T) {
	c, _ := newTestClient(t, 403,
		`{"requestId":"r-err","code":"IdempotentParameterMismatch","message":"clientToken token-1 was used with different name"}`)
	_, err := c.CreateEmbodiedEval(&CreateEmbodiedEvalArgs{ClientToken: "token-1", Name: "eval-2"})
	serviceErr, ok := err.(*bce.BceServiceError)
	if !ok {
		t.Fatalf("err = %T %v, want *bce.BceServiceError", err, err)
	}
	if serviceErr.StatusCode != 403 || serviceErr.Code != "IdempotentParameterMismatch" || serviceErr.RequestId != "r-err" ||
		!strings.Contains(serviceErr.Message, "different name") {
		t.Errorf("unexpected service error: %+v", serviceErr)
	}
}

func TestEmbodiedEvalRequiredArgs(t *testing.T) {
	c, got := newTestClient(t, 200, `{}`)
	cases := map[string]func() error{
		"create nil": func() error { _, err := c.CreateEmbodiedEval(nil); return err },
		"describe":   func() error { _, err := c.DescribeEmbodiedEval(""); return err },
		"stop nil":   func() error { _, err := c.StopEmbodiedEval(nil); return err },
		"stop":       func() error { _, err := c.StopEmbodiedEval(&StopEmbodiedEvalArgs{}); return err },
		"delete nil": func() error { _, err := c.DeleteEmbodiedEval(nil); return err },
		"delete":     func() error { _, err := c.DeleteEmbodiedEval(&DeleteEmbodiedEvalArgs{}); return err },
		"result":     func() error { _, err := c.DescribeEmbodiedEvalResult(""); return err },
		"video nil":  func() error { _, err := c.DescribeEmbodiedEvalVideo(nil); return err },
		"video evalId": func() error {
			_, err := c.DescribeEmbodiedEvalVideo(&DescribeEmbodiedEvalVideoArgs{VideoId: "v"})
			return err
		},
		"video videoId": func() error {
			_, err := c.DescribeEmbodiedEvalVideo(&DescribeEmbodiedEvalVideoArgs{EvalId: "e"})
			return err
		},
		"model nil":      func() error { _, err := c.DescribeEmbodiedEvalModel(nil); return err },
		"model":          func() error { _, err := c.DescribeEmbodiedEvalModel(&DescribeEmbodiedEvalModelArgs{}); return err },
		"benchmarks nil": func() error { _, err := c.DescribeEmbodiedEvalBenchmarks(nil); return err },
		"benchmarks model": func() error {
			_, err := c.DescribeEmbodiedEvalBenchmarks(&DescribeEmbodiedEvalBenchmarksArgs{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Errorf("%s: want error for missing required arg", name)
		}
	}
	// 参数不全时不能发出请求
	if got.method != "" {
		t.Errorf("request should not be sent, got %s", got.method)
	}
}

func TestImageConfigStringMasksPassword(t *testing.T) {
	ic := ImageConfig{Username: "u", Password: "secret-pass"}
	// 覆盖值和指针两种 %v / %+v 格式化，日志里最常见的就是这两种写法
	for _, s := range []string{fmt.Sprintf("%v", ic), fmt.Sprintf("%+v", ic), fmt.Sprintf("%+v", &ic)} {
		if strings.Contains(s, "secret-pass") {
			t.Errorf("String() leaks password: %s", s)
		}
	}
	// JSON 序列化仍要带上密码，创建请求依赖它
	b, _ := json.Marshal(ic)
	if !strings.Contains(string(b), "secret-pass") {
		t.Errorf("json should keep password: %s", string(b))
	}
}

func TestNewClient(t *testing.T) {
	if _, err := NewClient("ak", "sk", ""); err != nil {
		t.Errorf("empty endpoint should use default, got %v", err)
	}
	c, err := NewClientWithSTS("ak", "sk", "token", "aihc.bj.baidubce.com")
	if err != nil {
		t.Fatalf("NewClientWithSTS failed: %v", err)
	}
	bceClient := c.GetBceClient()
	c.SetBceClient(bceClient)
	if c.GetBceClient() != bceClient {
		t.Errorf("SetBceClient/GetBceClient mismatch")
	}
}

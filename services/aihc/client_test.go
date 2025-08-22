package aihc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	v1 "github.com/baidubce/bce-sdk-go/services/aihc/api/v1"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	AIHC_CLIENT Interface

	RESOURCE_POOL_ID string
	QUEUE_NAME       string
	QUEUE_ID         string
	AIJobID          string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	AIHC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}, testName string) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		equal = true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		errorMsg := fmt.Sprintf("%s:%d: mismatch, expect %v but got %v", file, line, expected, actual)
		alert(errorMsg)
		return false
	}
	return true
}

func TestListResourcePool(t *testing.T) {
	params := &v1.ListResourcePoolRequest{
		PageNo:   1,
		PageSize: 1,
	}
	testName := "TestListResPool"
	result, err := AIHC_CLIENT.ListResourcePool(params)
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}

	size := params.PageSize
	if result.Result.Total < size {
		size = result.Result.Total
	}

	if len(result.Result.ResourcePools) != size {
		t.Fatalf("the number of resourcePools in response is unexpected")
	}

	RESOURCE_POOL_ID = result.Result.ResourcePools[0].Metadata.Id

	params.PageNo = 2
	result, err = AIHC_CLIENT.ListResourcePool(params)
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}

	if len(result.Result.ResourcePools) != size {
		t.Fatalf("the number of resourcePools in response is unexpected")
	}
	if result.Result.ResourcePools[0].Metadata.Id == RESOURCE_POOL_ID {
		t.Fatalf("the pageNo is uneffective")
	}

	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("ListResPool success:\n%s", string(resultJSON))
}

func TestGetResourcePool(t *testing.T) {
	testName := "TestGetResPool"
	result, err := AIHC_CLIENT.GetResourcePool(RESOURCE_POOL_ID)
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}

	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("GetResPool success:\n%s", string(resultJSON))
}

func TestListNodeByResourcePoolID(t *testing.T) {
	testName := "TestListNodeByResourcePoolID"

	params := &v1.ListResourcePoolNodeRequest{
		PageNo:   1,
		PageSize: 3,
	}

	result, err := AIHC_CLIENT.ListNodeByResourcePoolID(RESOURCE_POOL_ID, params)
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}

	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("ListNodeByResourcePoolID success:\n%s", string(resultJSON))
}

func TestListQueue(t *testing.T) {
	testName := "TestListQueue"

	params := &v1.ListQueueRequest{
		PageNo:   1,
		PageSize: 3,
	}

	result, err := AIHC_CLIENT.ListQueue(RESOURCE_POOL_ID, params)
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}

	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("ListQueue success:\n%s", string(resultJSON))
}

func TestGetQueue(t *testing.T) {
	testName := "TestGetQueue"
	result, err := AIHC_CLIENT.GetQueue(RESOURCE_POOL_ID, QUEUE_NAME)
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}

	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("GetQueue success:\n%s", string(resultJSON))
}

func TestListJobs(t *testing.T) {
	testName := "TestListJobs"
	req := &v1.OpenAPIJobListRequest{
		ResourcePoolID: RESOURCE_POOL_ID,
		PageNo:         1,
		PageSize:       3,
	}
	result, err := AIHC_CLIENT.ListJobs(req)
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("ListJobs success:\n%s", string(resultJSON))
}

func TestGetAIJob(t *testing.T) {
	testName := "TestGetAIJob"
	result, err := AIHC_CLIENT.GetJob(&v1.GetAIJobOptions{
		JobID:          AIJobID,
		ResourcePoolID: RESOURCE_POOL_ID,
		QueueID:        QUEUE_ID,
	})
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		t.Errorf("GetAIJob failed: %v", err)
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		t.Errorf("GetAIJob failed1: %v", err)
		return
	}
	t.Logf("GetAIJob success:\n%s", string(resultJSON))
}

func TestDeleteJob(t *testing.T) {
	testName := "TestDeleteJob"
	result, err := AIHC_CLIENT.DeleteJob(&v1.DeleteAIJobOptions{
		JobID:          AIJobID,
		ResourcePoolID: RESOURCE_POOL_ID,
		QueueID:        QUEUE_ID,
	})
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("DeleteAIJob success:\n%s", string(resultJSON))
}

func TestCreateJob(t *testing.T) {
	// 测试用例结构体：包含测试名称、输入参数、预期结果检查函数
	type testCase struct {
		name        string
		jobReq      *v1.OpenAPIJobCreateRequest
		opts        *v1.CreateAIJobOptions
		checkResult func(t *testing.T, result *v1.OpenAPIJobCreateResponse, err error)
	}

	// 基础请求参数（可复用）
	baseJobReq := &v1.OpenAPIJobCreateRequest{
		Name:         "test",
		Queue:        "default",
		JobFramework: "",
		JobSpec: v1.OpenAPIAIJobSpec{
			Command:  "sleep 2000",
			Image:    "",
			Replicas: 1,
		},
		FaultTolerance:       false,
		Labels:               []v1.OpenAPILabel{{Key: "key", Value: "value"}},
		Priority:             "",
		Datasources:          nil,
		FaultToleranceConfig: &v1.OpenAPIJobFaultToleranceConfig{},
		AlertConfig:          nil,
		EnableBccl:           false,
	}

	baseOpts := &v1.CreateAIJobOptions{
		ResourcePoolID: "",
		Username:       "",
	}

	// 测试用例集合（覆盖正常场景和异常场景）
	testCases := []testCase{
		{
			name:   "normal_case",
			jobReq: baseJobReq,
			opts:   baseOpts,
			checkResult: func(t *testing.T, result *v1.OpenAPIJobCreateResponse, err error) {
				t.Helper() // 标记为辅助函数，错误定位更准确

				// 检查无错误
				if err != nil {
					t.Fatalf("预期无错误，实际错误: %v", err)
				}

				// 检查结果非空且包含必要字段
				if result == nil {
					t.Fatal("返回结果不应为nil")
				}
				if result.Result.JobID == "" {
					t.Error("返回结果缺少JobID")
				}

				// 输出格式化结果（便于调试）
				resultJSON, err := json.MarshalIndent(result, "", "    ")
				if err != nil {
					t.Errorf("序列化结果失败: %v", err)
					return
				}
				t.Logf("创建任务成功，详情:\n%s", resultJSON)
			},
		},
		{
			name:   "invalid_resource_pool",
			jobReq: baseJobReq,
			opts: &v1.CreateAIJobOptions{
				ResourcePoolID: "invalid-pool", // 无效资源池ID
				Username:       "",
			},
			checkResult: func(t *testing.T, result *v1.OpenAPIJobCreateResponse, err error) {
				t.Helper()

				// 预期应返回错误
				if err == nil {
					t.Fatal("预期有错误，实际无错误")
				}
				t.Logf("捕获预期错误: %v", err) // 记录预期错误，便于确认错误类型是否正确
			},
		},
		{
			name: "empty_job_name",
			jobReq: func() *v1.OpenAPIJobCreateRequest {
				req := *baseJobReq
				req.Name = "" // 空任务名称（非法参数）
				return &req
			}(),
			opts: baseOpts,
			checkResult: func(t *testing.T, result *v1.OpenAPIJobCreateResponse, err error) {
				t.Helper()
				if err == nil {
					t.Fatal("任务名称为空时应返回错误")
				}
			},
		},
	}

	// 执行所有测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 为每个子测试创建独立的客户端实例（避免测试间状态污染）
			client := AIHC_CLIENT // 若有必要可在此处初始化新实例

			// 调用被测试方法
			result, err := client.CreateJob(tc.jobReq, tc.opts)

			// 执行结果检查
			tc.checkResult(t, result, err)
		})
	}
}

func TestUpadateJob(t *testing.T) {
	testName := "TestUpdateJob"
	result, err := AIHC_CLIENT.UpdateJob(&v1.OpenAPIJobUpdateRequest{
		Priority: "",
	}, &v1.UpdateAIJobOptions{
		JobID:          AIJobID,
		ResourcePoolID: RESOURCE_POOL_ID,
		QueueID:        QUEUE_ID,
	})
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("UpdateAIJob success:\n%s", string(resultJSON))
}

func TestStopJob(t *testing.T) {
	testName := "TestStopJob"
	result, err := AIHC_CLIENT.StopJob(&v1.StopAIJobOptions{
		JobID:          AIJobID,
		ResourcePoolID: RESOURCE_POOL_ID,
		QueueID:        QUEUE_ID,
	})
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("StopAIJob success:\n%s", string(resultJSON))
}

func TestGetTaskEvents(t *testing.T) {
	testName := "TestCreateTaskEvents"
	result, err := AIHC_CLIENT.GetTaskEvent(&v1.GetJobEventsRequest{
		Namespace:      "",
		JobFramework:   "",
		StartTime:      "",
		EndTime:        "",
		JobID:          "",
		ResourcePoolID: "",
	})
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("GetTaskEvent success:\n%s", string(resultJSON))
}

func TestGetPodEvents(t *testing.T) {

	testName := "TestGetPodEvents"
	result, err := AIHC_CLIENT.GetPodEvents(&v1.GetPodEventsRequest{
		JobID:          "",
		Namespace:      "",
		JobFramework:   "",
		StartTime:      "",
		EndTime:        "",
		ResourcePoolID: "",
		PodName:        "",
	})
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("GetPodEvent success:\n%s", string(resultJSON))
}

func TestGetPodLogs(t *testing.T) {
	testName := "TestGetPodLogs"
	result, err := AIHC_CLIENT.GetPodLogs(&v1.GetPodLogsRequest{
		JobID:          "",
		ResourcePoolID: "",
		PodName:        "",
		Namespace:      "",
		StartTime:      "",
		EndTime:        "",
		MaxLines:       "",
		Container:      "",
		Chunk:          "",
	})
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("GetPodLogs success:\n%s", string(resultJSON))
}

func TestGetJobNodesList(t *testing.T) {
	testName := "TestGetJobNodesList"
	result, err := AIHC_CLIENT.GetJobNodesList(&v1.GetJobNodesListOptions{
		JobID:          AIJobID,
		ResourcePoolID: RESOURCE_POOL_ID,
		Namespace:      "",
	})
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("GetJobNodesList success:\n%s", string(resultJSON))
}

func TestGetTaskMetrics(t *testing.T) {
	testName := "TestGetTaskMetrics"
	result, err := AIHC_CLIENT.GetTaskMetrics(&v1.GetTaskMetricsRequest{
		StartTime:      "",
		ResourcePoolID: "",
		EndTime:        "",
		TimeStep:       "",
		MetricType:     "",
		JobID:          "",
		Namespace:      "",
		RateInterval:   "",
	})
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	if result == nil {
		t.Error("Result should not be nil")
		return
	}
	// 使用 t.Logf 输出详细信息
	resultJSON, err := json.MarshalIndent(result, "", "    ")
	if !ExpectEqual(t.Errorf, err, nil, testName) {
		return
	}
	t.Logf("GetTaskMetrics success:\n%s", string(resultJSON))
}

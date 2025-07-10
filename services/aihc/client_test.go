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
	testName := "TestCreateJob"
	result, err := AIHC_CLIENT.CreateJob(&v1.OpenAPIJobCreateRequest{
		Name:                 "",
		Queue:                "",
		JobFramework:         "",
		JobSpec:              v1.OpenAPIAIJobSpec{},
		FaultTolerance:       false,
		Labels:               nil,
		Priority:             "",
		Datasources:          nil,
		FaultToleranceConfig: &v1.OpenAPIJobFaultToleranceConfig{},
		AlertConfig:          nil,
		EnableBccl:           false,
	}, &v1.CreateAIJobOptions{ResourcePoolID: RESOURCE_POOL_ID})
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
	t.Logf("CreateAIJob success:\n%s", string(resultJSON))
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

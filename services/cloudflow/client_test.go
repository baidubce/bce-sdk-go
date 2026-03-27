package cloudflow

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/cloudflow/api"
	"github.com/baidubce/bce-sdk-go/util"
)

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
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
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

var (
	bceServiceError404 = bce.NewBceServiceError("NOTFound", "404 NOT Found", "", http.StatusNotFound)
	bceServiceError500 = bce.NewBceServiceError("InternalError", "Internal Server Error", "", http.StatusInternalServerError)
)

// NewMockCloudFlowClient creates a mock cloudflow client for testing
func NewMockCloudFlowClient(ak, sk, endpoint, respBody string, options ...util.MockRoundTripperOption) (*Client, error) {
	// First create the client
	client, err := NewClient(ak, sk, endpoint)
	if err != nil {
		return nil, err
	}

	// Setup default options if not provided
	if len(options) == 0 {
		options = []util.MockRoundTripperOption{
			util.SetStatusCode(http.StatusOK),
			util.SetStatusMsg(http.StatusText(http.StatusOK)),
		}
	}
	if len(respBody) > 0 {
		options = append(options, util.SetRespBody(respBody))
	}

	// Create mock HTTP client and inject it
	mockHttpClient := util.NewMockHTTPClient(options...)
	if mockHttpClient == nil {
		return nil, fmt.Errorf("util.NewMockHTTPClient fail")
	}
	client.Config.Retry = bce.NewNoRetryPolicy()
	client.HTTPClient = mockHttpClient
	return client, nil
}

// newPostMigrationArgs creates a new PostMigrationArgs for testing
// Note: AK/SK must be short (< 117 bytes) due to RSA 1024-bit key encryption limit
// This function creates a new instance each time because PostMigration modifies the args in place
func newPostMigrationArgs() *api.PostMigrationArgs {
	return &api.PostMigrationArgs{
		PostMigrationArgsCommon: api.PostMigrationArgsCommon{
			Name:     "test-migration",
			Strategy: api.MIGFRATION_STRATEGY_FORCE_OVERWRITE,
			DestinationConfig: api.MigrationDestinationConfig{
				MigrationConfigCommon: api.MigrationConfigCommon{
					Provider: "BOS",
					Endpoint: "bj.bcebos.com",
					Bucket:   "dest-bucket",
					Ak:       "dak",
					Sk:       "dsk",
				},
				StorageClass: api.STORAGE_CLASS_STANDARD,
			},
			MigrationType: api.MigrationType{
				Type: api.MIGRATION_TYPE_STOCK,
			},
		},
		SourceConfig: api.MigrationPrefixSourceConfig{
			MigrationConfigCommon: api.MigrationConfigCommon{
				Provider: "AWS",
				Endpoint: "s3.amazonaws.com",
				Bucket:   "source-bucket",
				Ak:       "sak",
				Sk:       "ssk",
			},
			Prefixes: []string{"prefix1/", "prefix2/"},
		},
	}
}

// newPostMigrationFromListArgs creates a new PostMigrationFromListArgs for testing
func newPostMigrationFromListArgs() *api.PostMigrationFromListArgs {
	return &api.PostMigrationFromListArgs{
		PostMigrationArgsCommon: api.PostMigrationArgsCommon{
			Name:     "test-migration-from-list",
			Strategy: api.MIGFRATION_STRATEGY_KEEP_DESTINATION,
			DestinationConfig: api.MigrationDestinationConfig{
				MigrationConfigCommon: api.MigrationConfigCommon{
					Provider: "BOS",
					Endpoint: "bj.bcebos.com",
					Bucket:   "dest-bucket",
					Ak:       "dak",
					Sk:       "dsk",
				},
			},
		},
		SourceConfig: api.MigrationListSourceConfig{
			MigrationConfigCommon: api.MigrationConfigCommon{
				Provider: "AWS",
				Endpoint: "s3.amazonaws.com",
				Bucket:   "source-bucket",
				Ak:       "sak",
				Sk:       "ssk",
			},
			ListFileURL: []string{"https://example.com/list.txt"},
		},
	}
}

func TestNewClient(t *testing.T) {
	// case1: NewClient(ak, sk, endpoint)
	ak, sk, endpoint := "test-ak", "test-sk", "test-endpoint"
	_, err := NewClient(ak, sk, endpoint)
	ExpectEqual(t.Errorf, nil, err)

	// case2: NewClientWithConfig
	config := &ClientConfiguration{
		Ak:       ak,
		Sk:       sk,
		Endpoint: endpoint,
	}
	_, err = NewClientWithConfig(config)
	ExpectEqual(t.Errorf, nil, err)

	// case3: empty ak/sk for public-read-write
	config2 := &ClientConfiguration{
		Endpoint: endpoint,
	}
	_, err = NewClientWithConfig(config2)
	ExpectEqual(t.Errorf, nil, err)

	// case4: empty endpoint uses default
	config3 := NewClientConfig(ak, sk, "")
	client, err := NewClientWithConfig(config3)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, DEFAULT_SERVICE_DOMAIN, client.Config.Endpoint)

	// case5: With methods chain
	config4 := NewClientConfig(ak, sk, endpoint).
		WithAk("new-ak").
		WithSk("new-sk").
		WithEndpoint("new-endpoint").
		WithRedirectDisabled(true).
		WithDisableKeepAlives(true).
		WithNoVerifySSL(true).
		WithDialTimeout(100).
		WithKeepAlive(100).
		WithReadTimeout(100).
		WithWriteTimeout(100).
		WithTLSHandshakeTimeout(100).
		WithIdleConnectionTimeout(100).
		WithResponseHeaderTimeout(100).
		WithHttpClientTimeout(100).
		WithRetryPolicy(&bce.BackOffRetryPolicy{})
	ExpectEqual(t.Errorf, "new-ak", config4.Ak)
	ExpectEqual(t.Errorf, "new-sk", config4.Sk)
	ExpectEqual(t.Errorf, "new-endpoint", config4.Endpoint)
	ExpectEqual(t.Errorf, true, config4.RedirectDisabled)
	ExpectEqual(t.Errorf, true, config4.DisableKeepAlives)
	ExpectEqual(t.Errorf, true, config4.NoVerifySSL)
	_, err = NewClientWithConfig(config4)
	ExpectEqual(t.Errorf, nil, err)
}

func TestPostMigration(t *testing.T) {
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id",
		"result": {
			"taskID": ["task-id-001", "task-id-002"]
		}
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockCloudFlowClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	res, err := client.PostMigration(newPostMigrationArgs())
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)
	ExpectEqual(t.Errorf, 2, len(res.Result.TaskId))
	ExpectEqual(t.Errorf, "task-id-001", res.Result.TaskId[0])

	// case2: args is nil
	res, err = client.PostMigration(nil)
	ExpectEqual(t.Errorf, bce.NewBceClientError("PostMigrationArgs is nil"), err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: HTTP 404 error
	client404, err := NewMockCloudFlowClient(ak, sk, endpoint, "", util.RoundTripperOpts404...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client404.PostMigration(newPostMigrationArgs())
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case4: network error
	netError := fmt.Errorf("net error")
	options := []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(netError),
	}
	clientNetErr, err := NewMockCloudFlowClient(ak, sk, endpoint, "", options...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientNetErr.PostMigration(newPostMigrationArgs())
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)

	// case5: response with error code
	errorRespBody := `{
		"success": false,
		"code": "InvalidParameter",
		"message": "invalid parameter",
		"requestId": "test-request-id"
	}`
	clientError, err := NewMockCloudFlowClient(ak, sk, endpoint, errorRespBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientError.PostMigration(newPostMigrationArgs())
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "InvalidParameter", serviceErr.Code)
}

func TestPostMigrationFromList(t *testing.T) {
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id",
		"result": {
			"taskID": ["task-id-003"]
		}
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockCloudFlowClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	res, err := client.PostMigrationFromList(newPostMigrationFromListArgs())
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)
	ExpectEqual(t.Errorf, 1, len(res.Result.TaskId))

	// case2: args is nil
	res, err = client.PostMigrationFromList(nil)
	ExpectEqual(t.Errorf, bce.NewBceClientError("PostMigrationFromListArgs is nil"), err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: HTTP 404 error
	client404, err := NewMockCloudFlowClient(ak, sk, endpoint, "", util.RoundTripperOpts404...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client404.PostMigrationFromList(newPostMigrationFromListArgs())
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestGetMigration(t *testing.T) {
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id",
		"result": [
			{
				"taskID": "task-id-001",
				"name": "test-migration",
				"runningStatus": "MIGRATING",
				"createTime": 1679000000,
				"totalCount": 1000,
				"finishedCount": 500,
				"failedCount": 10
			}
		]
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockCloudFlowClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	res, err := client.GetMigration("task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)
	ExpectEqual(t.Errorf, 1, len(res.TaskInfos))
	ExpectEqual(t.Errorf, "task-id-001", res.TaskInfos[0].TaskId)
	ExpectEqual(t.Errorf, "MIGRATING", res.TaskInfos[0].RunningStatus)
	ExpectEqual(t.Errorf, int64(1000), res.TaskInfos[0].TotalCount)
	ExpectEqual(t.Errorf, int64(500), res.TaskInfos[0].FinishedCount)

	// case2: HTTP 404 error
	client404, err := NewMockCloudFlowClient(ak, sk, endpoint, "", util.RoundTripperOpts404...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client404.GetMigration("task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: network error
	netError := fmt.Errorf("net error")
	options := []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(netError),
	}
	clientNetErr, err := NewMockCloudFlowClient(ak, sk, endpoint, "", options...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientNetErr.GetMigration("task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)

	// case4: response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskNotFound",
		"message": "task not found",
		"requestId": "test-request-id"
	}`
	clientError, err := NewMockCloudFlowClient(ak, sk, endpoint, errorRespBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientError.GetMigration("non-exist-task")
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "TaskNotFound", serviceErr.Code)
}

func TestListMigration(t *testing.T) {
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id",
		"result": [
			{
				"taskID": "task-id-001",
				"name": "migration-1",
				"runningStatus": "FINISHED"
			},
			{
				"taskID": "task-id-002",
				"name": "migration-2",
				"runningStatus": "MIGRATING"
			}
		]
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockCloudFlowClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	res, err := client.ListMigration()
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)
	ExpectEqual(t.Errorf, 2, len(res.TaskInfos))
	ExpectEqual(t.Errorf, "task-id-001", res.TaskInfos[0].TaskId)
	ExpectEqual(t.Errorf, "FINISHED", res.TaskInfos[0].RunningStatus)
	ExpectEqual(t.Errorf, "task-id-002", res.TaskInfos[1].TaskId)
	ExpectEqual(t.Errorf, "MIGRATING", res.TaskInfos[1].RunningStatus)

	// case2: HTTP 404 error
	client404, err := NewMockCloudFlowClient(ak, sk, endpoint, "", util.RoundTripperOpts404...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client404.ListMigration()
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: network error
	netError := fmt.Errorf("net error")
	options := []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(netError),
	}
	clientNetErr, err := NewMockCloudFlowClient(ak, sk, endpoint, "", options...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientNetErr.ListMigration()
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestGetMigrationResult(t *testing.T) {
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id",
		"result": {
			"failObjectListURLs": [
				"https://example.com/fail-list-1.txt",
				"https://example.com/fail-list-2.txt"
			]
		}
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockCloudFlowClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	res, err := client.GetMigrationResult("task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)
	ExpectEqual(t.Errorf, 2, len(res.Result.FailObjectListurl))
	ExpectEqual(t.Errorf, "https://example.com/fail-list-1.txt", res.Result.FailObjectListurl[0])

	// case2: HTTP 404 error
	client404, err := NewMockCloudFlowClient(ak, sk, endpoint, "", util.RoundTripperOpts404...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client404.GetMigrationResult("task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskNotFound",
		"message": "task not found",
		"requestId": "test-request-id"
	}`
	clientError, err := NewMockCloudFlowClient(ak, sk, endpoint, errorRespBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientError.GetMigrationResult("non-exist-task")
	ExpectEqual(t.Errorf, true, err != nil)
}

func TestPauseMigration(t *testing.T) {
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id"
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockCloudFlowClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	res, err := client.PauseMigration("task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)

	// case2: HTTP 404 error
	client404, err := NewMockCloudFlowClient(ak, sk, endpoint, "", util.RoundTripperOpts404...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client404.PauseMigration("task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskNotRunning",
		"message": "task is not running",
		"requestId": "test-request-id"
	}`
	clientError, err := NewMockCloudFlowClient(ak, sk, endpoint, errorRespBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientError.PauseMigration("task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "TaskNotRunning", serviceErr.Code)
}

func TestResumeMigration(t *testing.T) {
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id"
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockCloudFlowClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	res, err := client.ResumeMigration("task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)

	// case2: HTTP 404 error
	client404, err := NewMockCloudFlowClient(ak, sk, endpoint, "", util.RoundTripperOpts404...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client404.ResumeMigration("task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskNotPaused",
		"message": "task is not paused",
		"requestId": "test-request-id"
	}`
	clientError, err := NewMockCloudFlowClient(ak, sk, endpoint, errorRespBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientError.ResumeMigration("task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
}

func TestRetryMigration(t *testing.T) {
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id"
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockCloudFlowClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	res, err := client.RetryMigration("task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)

	// case2: HTTP 404 error
	client404, err := NewMockCloudFlowClient(ak, sk, endpoint, "", util.RoundTripperOpts404...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client404.RetryMigration("task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: network error
	netError := fmt.Errorf("net error")
	options := []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(netError),
	}
	clientNetErr, err := NewMockCloudFlowClient(ak, sk, endpoint, "", options...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientNetErr.RetryMigration("task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestDeleteMigration(t *testing.T) {
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id"
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockCloudFlowClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	res, err := client.DeleteMigration("task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)

	// case2: HTTP 404 error
	client404, err := NewMockCloudFlowClient(ak, sk, endpoint, "", util.RoundTripperOpts404...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client404.DeleteMigration("task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskRunning",
		"message": "cannot delete running task",
		"requestId": "test-request-id"
	}`
	clientError, err := NewMockCloudFlowClient(ak, sk, endpoint, errorRespBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientError.DeleteMigration("task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "TaskRunning", serviceErr.Code)

	// case4: network error
	netError := fmt.Errorf("net error")
	options := []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(netError),
	}
	clientNetErr, err := NewMockCloudFlowClient(ak, sk, endpoint, "", options...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = clientNetErr.DeleteMigration("task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)
}

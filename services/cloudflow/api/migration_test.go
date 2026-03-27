package api

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
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

// NewMockBceClient creates a mock BceClient for API testing
func NewMockBceClient() (*bce.BceClient, error) {
	credentials, err := auth.NewBceCredentials("ak", "sk")
	if err != nil {
		return nil, err
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS,
	}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:    "192.168.1.1:8080",
		Region:      bce.DEFAULT_REGION,
		UserAgent:   bce.DEFAULT_USER_AGENT,
		Credentials: credentials,
		SignOption:  defaultSignOptions,
		Retry:       bce.NewNoRetryPolicy(),
	}
	v1Signer := &auth.BceV1Signer{}
	client := bce.NewBceClient(defaultConf, v1Signer)
	return client, nil
}

// AttachMockHttpClient attaches a mock HTTP client to the BceClient
func AttachMockHttpClient(client *bce.BceClient, respBody string, options ...util.MockRoundTripperOption) {
	if len(options) == 0 {
		options = []util.MockRoundTripperOption{
			util.SetStatusCode(http.StatusOK),
			util.SetStatusMsg(http.StatusText(http.StatusOK)),
		}
	}
	if len(respBody) > 0 {
		options = append(options, util.SetRespBody(respBody))
	}
	mockHttpClient := util.NewMockHTTPClient(options...)
	client.HTTPClient = mockHttpClient
}

// newPostMigrationArgs creates a new PostMigrationArgs for testing
func newPostMigrationArgs() *PostMigrationArgs {
	return &PostMigrationArgs{
		PostMigrationArgsCommon: PostMigrationArgsCommon{
			Name:     "test-migration",
			Strategy: MIGFRATION_STRATEGY_FORCE_OVERWRITE,
			DestinationConfig: MigrationDestinationConfig{
				MigrationConfigCommon: MigrationConfigCommon{
					Provider: "BOS",
					Endpoint: "bj.bcebos.com",
					Bucket:   "dest-bucket",
					Ak:       "dak",
					Sk:       "dsk",
				},
				StorageClass: STORAGE_CLASS_STANDARD,
			},
			MigrationType: MigrationType{
				Type: MIGRATION_TYPE_STOCK,
			},
		},
		SourceConfig: MigrationPrefixSourceConfig{
			MigrationConfigCommon: MigrationConfigCommon{
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
func newPostMigrationFromListArgs() *PostMigrationFromListArgs {
	return &PostMigrationFromListArgs{
		PostMigrationArgsCommon: PostMigrationArgsCommon{
			Name:     "test-migration-from-list",
			Strategy: MIGFRATION_STRATEGY_KEEP_DESTINATION,
			DestinationConfig: MigrationDestinationConfig{
				MigrationConfigCommon: MigrationConfigCommon{
					Provider: "BOS",
					Endpoint: "bj.bcebos.com",
					Bucket:   "dest-bucket",
					Ak:       "dak",
					Sk:       "dsk",
				},
			},
		},
		SourceConfig: MigrationListSourceConfig{
			MigrationConfigCommon: MigrationConfigCommon{
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

func TestPostMigrationAPI(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// case1: args is nil
	res, err := PostMigration(client, nil)
	ExpectEqual(t.Errorf, bce.NewBceClientError("PostMigrationArgs is nil"), err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case2: success
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id",
		"result": {
			"taskID": ["task-id-001"]
		}
	}`
	AttachMockHttpClient(client, respBody)
	res, err = PostMigration(client, newPostMigrationArgs())
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)
	ExpectEqual(t.Errorf, 1, len(res.Result.TaskId))

	// case3: HTTP 404 error
	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)
	res, err = PostMigration(client, newPostMigrationArgs())
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case4: network error
	netError := fmt.Errorf("net error")
	AttachMockHttpClient(client, "", util.SetHTTPClientDoError(netError))
	res, err = PostMigration(client, newPostMigrationArgs())
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)

	// case5: response with error code
	errorRespBody := `{
		"success": false,
		"code": "InvalidParameter",
		"message": "invalid parameter",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, errorRespBody)
	res, err = PostMigration(client, newPostMigrationArgs())
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "InvalidParameter", serviceErr.Code)
}

func TestPostMigrationFromListAPI(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// case1: args is nil
	res, err := PostMigrationFromList(client, nil)
	ExpectEqual(t.Errorf, bce.NewBceClientError("PostMigrationFromListArgs is nil"), err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case2: success
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id",
		"result": {
			"taskID": ["task-id-001"]
		}
	}`
	AttachMockHttpClient(client, respBody)
	res, err = PostMigrationFromList(client, newPostMigrationFromListArgs())
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)

	// case3: HTTP 404 error
	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)
	res, err = PostMigrationFromList(client, newPostMigrationFromListArgs())
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestGetMigrationAPI(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
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
				"totalCount": 1000,
				"finishedCount": 500
			}
		]
	}`
	AttachMockHttpClient(client, respBody)
	res, err := GetMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)
	ExpectEqual(t.Errorf, 1, len(res.TaskInfos))
	ExpectEqual(t.Errorf, "task-id-001", res.TaskInfos[0].TaskId)
	ExpectEqual(t.Errorf, "MIGRATING", res.TaskInfos[0].RunningStatus)

	// case2: HTTP 404 error
	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)
	res, err = GetMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskNotFound",
		"message": "task not found",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, errorRespBody)
	res, err = GetMigration(client, "non-exist-task")
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "TaskNotFound", serviceErr.Code)
}

func TestListMigrationAPI(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
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
	AttachMockHttpClient(client, respBody)
	res, err := ListMigration(client)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)
	ExpectEqual(t.Errorf, 2, len(res.TaskInfos))
	ExpectEqual(t.Errorf, "task-id-001", res.TaskInfos[0].TaskId)
	ExpectEqual(t.Errorf, "task-id-002", res.TaskInfos[1].TaskId)

	// case2: HTTP 404 error
	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)
	res, err = ListMigration(client)
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestGetMigrationResultAPI(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id",
		"result": {
			"failObjectListURLs": [
				"https://example.com/fail-list-1.txt"
			]
		}
	}`
	AttachMockHttpClient(client, respBody)
	res, err := GetMigrationResult(client, "task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)
	ExpectEqual(t.Errorf, 1, len(res.Result.FailObjectListurl))

	// case2: HTTP 404 error
	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)
	res, err = GetMigrationResult(client, "task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestPauseMigrationAPI(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, respBody)
	res, err := PauseMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)

	// case2: HTTP 404 error
	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)
	res, err = PauseMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskNotRunning",
		"message": "task is not running",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, errorRespBody)
	res, err = PauseMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "TaskNotRunning", serviceErr.Code)
}

func TestResumeMigrationAPI(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, respBody)
	res, err := ResumeMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)

	// case2: HTTP 404 error
	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)
	res, err = ResumeMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestRetryMigrationAPI(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, respBody)
	res, err := RetryMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)

	// case2: HTTP 404 error
	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)
	res, err = RetryMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestDeleteMigrationAPI(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// case1: success
	respBody := `{
		"success": true,
		"code": "",
		"message": "",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, respBody)
	res, err := DeleteMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Success)

	// case2: HTTP 404 error
	AttachMockHttpClient(client, "", util.RoundTripperOpts404...)
	res, err = DeleteMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, bceServiceError404, err)
	ExpectEqual(t.Errorf, true, res == nil)

	// case3: response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskRunning",
		"message": "cannot delete running task",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, errorRespBody)
	res, err = DeleteMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "TaskRunning", serviceErr.Code)
}

func TestRsaEncrypt(t *testing.T) {
	// case1: normal short string
	data := "short-string"
	encrypted, err := RsaEncrypt([]byte(data))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(encrypted) > 0)

	// case2: empty string
	encrypted, err = RsaEncrypt([]byte(""))
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(encrypted) > 0)
}

func TestEncryptAndBase64Encode(t *testing.T) {
	// case1: normal short string
	str := "test-ak"
	err := EncryptAndBase64Encode(&str)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(str) > len("test-ak"))

	// case2: string too long for RSA
	longStr := "this-is-a-very-long-string-that-exceeds-the-maximum-length-for-rsa-1024-bit-key-encryption-which-is-117-bytes-and-should-fail"
	err = EncryptAndBase64Encode(&longStr)
	ExpectEqual(t.Errorf, true, err != nil)
}

func TestMarkAuthentication(t *testing.T) {
	// case1: success
	config := &MigrationConfigCommon{
		Provider: "BOS",
		Endpoint: "bj.bcebos.com",
		Bucket:   "bucket",
		Ak:       "ak",
		Sk:       "sk",
	}
	err := MarkAuthentication(config)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, config.Ak != "ak")
	ExpectEqual(t.Errorf, true, config.Sk != "sk")
}

func TestSetUriAndEndpoint(t *testing.T) {
	// case1: http endpoint
	req := &bce.BceRequest{}
	setUriAndEndpoint(req, "http://192.168.1.1:8080")
	ExpectEqual(t.Errorf, "192.168.1.1:8080", req.Host())
	ExpectEqual(t.Errorf, "/v1/", req.Uri())

	// case2: https endpoint
	req2 := &bce.BceRequest{}
	setUriAndEndpoint(req2, "https://bj.bcebos.com")
	ExpectEqual(t.Errorf, "bj.bcebos.com", req2.Host())
	ExpectEqual(t.Errorf, "/v1/", req2.Uri())

	// case3: endpoint without protocol
	req3 := &bce.BceRequest{}
	setUriAndEndpoint(req3, "192.168.1.1")
	ExpectEqual(t.Errorf, "192.168.1.1", req3.Host())
}

func TestListMigrationWithErrorCode(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// response with error code
	errorRespBody := `{
		"success": false,
		"code": "InternalError",
		"message": "internal error",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, errorRespBody)
	res, err := ListMigration(client)
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "InternalError", serviceErr.Code)
	ExpectEqual(t.Errorf, true, res != nil)
}

func TestGetMigrationResultWithErrorCode(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskNotFound",
		"message": "task not found",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, errorRespBody)
	res, err := GetMigrationResult(client, "non-exist-task")
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "TaskNotFound", serviceErr.Code)
	ExpectEqual(t.Errorf, true, res != nil)
}

func TestResumeMigrationWithErrorCode(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskNotPaused",
		"message": "task is not paused",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, errorRespBody)
	res, err := ResumeMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "TaskNotPaused", serviceErr.Code)
	ExpectEqual(t.Errorf, true, res != nil)
}

func TestRetryMigrationWithErrorCode(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// response with error code
	errorRespBody := `{
		"success": false,
		"code": "TaskNotFinished",
		"message": "task is not finished",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, errorRespBody)
	res, err := RetryMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "TaskNotFinished", serviceErr.Code)
	ExpectEqual(t.Errorf, true, res != nil)
}

func TestPostMigrationFromListWithErrorCode(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// response with error code
	errorRespBody := `{
		"success": false,
		"code": "InvalidParameter",
		"message": "invalid list file url",
		"requestId": "test-request-id"
	}`
	AttachMockHttpClient(client, errorRespBody)
	res, err := PostMigrationFromList(client, newPostMigrationFromListArgs())
	ExpectEqual(t.Errorf, true, err != nil)
	serviceErr, ok := err.(*bce.BceServiceError)
	ExpectEqual(t.Errorf, true, ok)
	ExpectEqual(t.Errorf, "InvalidParameter", serviceErr.Code)
	ExpectEqual(t.Errorf, true, res != nil)
}

func TestMarkAuthenticationWithLongSk(t *testing.T) {
	// This test verifies MarkAuthentication fails when SK is too long for RSA encryption
	// RSA 1024-bit key can only encrypt max 117 bytes
	config := &MigrationConfigCommon{
		Provider: "BOS",
		Endpoint: "bj.bcebos.com",
		Bucket:   "bucket",
		Ak:       "short-ak",
		Sk:       "this-is-a-very-long-sk-that-exceeds-the-maximum-length-for-rsa-1024-bit-key-encryption-which-is-117-bytes-and-should-fail",
	}
	err := MarkAuthentication(config)
	ExpectEqual(t.Errorf, true, err != nil)
}

func TestPostMigrationWithInvalidJson(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// response with invalid JSON
	invalidJsonResp := `{invalid json}`
	AttachMockHttpClient(client, invalidJsonResp)
	res, err := PostMigration(client, newPostMigrationArgs())
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestGetMigrationWithInvalidJson(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// response with invalid JSON
	invalidJsonResp := `{invalid json}`
	AttachMockHttpClient(client, invalidJsonResp)
	res, err := GetMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestListMigrationWithInvalidJson(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// response with invalid JSON
	invalidJsonResp := `{invalid json}`
	AttachMockHttpClient(client, invalidJsonResp)
	res, err := ListMigration(client)
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)
}

func TestPauseMigrationWithInvalidJson(t *testing.T) {
	client, err := NewMockBceClient()
	ExpectEqual(t.Errorf, nil, err)

	// response with invalid JSON
	invalidJsonResp := `{invalid json}`
	AttachMockHttpClient(client, invalidJsonResp)
	res, err := PauseMigration(client, "task-id-001")
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, true, res == nil)
}

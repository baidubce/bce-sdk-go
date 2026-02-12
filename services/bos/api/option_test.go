// nolint
package api

import (
	"fmt"
	"hash/crc64"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	my_http "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

func TestAddCrc64Check(t *testing.T) {
	// new bce client
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	config := &bce.BceClientConfiguration{
		Endpoint:                  "endpoint",
		Region:                    "bce.DEFAULT_REGION",
		UserAgent:                 "bce.DEFAULT_USER_AGENT",
		Credentials:               &auth.BceCredentials{AccessKeyId: "ak", SecretAccessKey: "sk"},
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
	}
	client, err := bce.NewBceClientWithExclusiveHTTPClient(config, &auth.BceV1Signer{})
	ExpectEqual(t, nil, err)

	// prepare request / response
	req1 := &BosRequest{}
	resp1 := &BosResponse{}
	host := "hostname:8080"
	uri := "bucket/object"
	req1.SetMethod(http.MethodGet)
	req1.SetUri(uri)
	req1.SetHost(host)
	req1.SetProtocol("http")
	bodyStr1 := "this is a test-string"
	// calc crc64
	crc64Sum := crc64.New(crc64.MakeTable(crc64.ECMA))
	n, _ := crc64Sum.Write([]byte(bodyStr1))
	ExpectEqual(t, len(bodyStr1), n)
	crc64Str := strconv.FormatUint(crc64Sum.Sum64(), 10)

	// mock http client, crc64 match
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str,
		}),
	}
	mockHttpClient := util.NewMockHTTPClient(options...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient))
	// mock http client, crc64 mismatch
	crc64ErrorStr := "xxxxxxxxxxxxx"
	options12 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64ErrorStr,
		}),
	}
	mockHttpClient12 := util.NewMockHTTPClient(options12...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient12))

	// case1: body v2, AddCrc64Check, crc64 match
	body1, err := bce.NewBodyFromStringV2(bodyStr1, false)
	ExpectEqual(t, nil, err)
	req1.SetBody(body1)
	AddCrc64Check(req1, resp1)
	for _, tracker := range req1.Tracker {
		tracker(req1)
	}
	client.HTTPClient = mockHttpClient
	err = client.SendRequest(&req1.BceRequest, &resp1.BceResponse)
	ExpectEqual(t, nil, err)
	for _, handler := range resp1.Handler {
		err := handler(resp1)
		ExpectEqual(t, nil, err)
	}

	//case1.2: body v2, AddCrc64Check, crc64 mismatch
	body12, err := bce.NewBodyFromStringV2(bodyStr1, false)
	ExpectEqual(t, nil, err)
	req1.SetBody(body12)
	// reset request / response
	req1.Tracker = []RequestTracker{}
	resp1.Handler = []ResponseHandler{}
	AddCrc64Check(req1, resp1)
	for _, tracker := range req1.Tracker {
		tracker(req1)
	}
	client.HTTPClient = mockHttpClient12
	err = client.SendRequest(&req1.BceRequest, &resp1.BceResponse)
	ExpectEqual(t, nil, err)
	for _, handler := range resp1.Handler {
		err := handler(resp1)
		ExpectEqual(t, fmt.Errorf("crc64 is not consistent, client: %s, server: %s", crc64Str, crc64ErrorStr), err)
	}

	//case2.1: body v2, AddWriter/Crc64Handler, match
	client.HTTPClient = mockHttpClient
	body2, err := bce.NewBodyFromStringV2(bodyStr1, false)
	ExpectEqual(t, nil, err)
	req1.SetBody(body2)
	crc64Hasher2 := crc64.New(crc64.MakeTable(crc64.ECMA))
	tracker2 := AddWriter(crc64Hasher2)
	handler2 := Crc64Handler(crc64Hasher2)
	// reset request / response
	req1.Tracker = []RequestTracker{tracker2}
	resp1.Handler = []ResponseHandler{handler2}
	for _, tracker := range req1.Tracker {
		tracker(req1)
	}
	err = client.SendRequest(&req1.BceRequest, &resp1.BceResponse)
	ExpectEqual(t, nil, err)
	for _, handler := range resp1.Handler {
		err := handler(resp1)
		ExpectEqual(t, nil, err)
	}

	//case2.2: body v2, AddWriter/Crc64Handler, mismatch
	client.HTTPClient = mockHttpClient12
	body22, err := bce.NewBodyFromStringV2(bodyStr1, false)
	ExpectEqual(t, nil, err)
	req1.SetBody(body22)
	crc64Hasher22 := crc64.New(crc64.MakeTable(crc64.ECMA))
	tracker22 := AddWriter(crc64Hasher22)
	handler22 := Crc64Handler(crc64Hasher22)
	// reset request / response
	req1.Tracker = []RequestTracker{tracker22}
	resp1.Handler = []ResponseHandler{handler22}
	for _, tracker := range req1.Tracker {
		tracker(req1)
	}
	err = client.SendRequest(&req1.BceRequest, &resp1.BceResponse)
	ExpectEqual(t, nil, err)
	for _, handler := range resp1.Handler {
		err := handler(resp1)
		ExpectEqual(t, fmt.Errorf("crc64 is not consistent, client: %s, server: %s", crc64Str, crc64ErrorStr), err)
	}

	// case3.1: body v1, AddCrc64Check, crc64 match
	body31, err := bce.NewBodyFromString(bodyStr1)
	ExpectEqual(t, nil, err)
	req1.SetBody(body31)
	// reset request / response
	req1.Tracker = []RequestTracker{}
	resp1.Handler = []ResponseHandler{}
	AddCrc64Check(req1, resp1)
	for _, tracker := range req1.Tracker {
		tracker(req1)
	}
	client.HTTPClient = mockHttpClient
	err = client.SendRequest(&req1.BceRequest, &resp1.BceResponse)
	ExpectEqual(t, nil, err)
	for _, handler := range resp1.Handler {
		err := handler(resp1)
		ExpectEqual(t, nil, err)
	}

	//case3.2: body v1, AddCrc64Check, crc64 mismatch
	body32, err := bce.NewBodyFromString(bodyStr1)
	ExpectEqual(t, nil, err)
	req1.SetBody(body32)
	// reset request / response
	req1.Tracker = []RequestTracker{}
	resp1.Handler = []ResponseHandler{}
	AddCrc64Check(req1, resp1)
	for _, tracker := range req1.Tracker {
		tracker(req1)
	}
	client.HTTPClient = mockHttpClient12
	err = client.SendRequest(&req1.BceRequest, &resp1.BceResponse)
	ExpectEqual(t, nil, err)
	for _, handler := range resp1.Handler {
		err := handler(resp1)
		ExpectEqual(t, fmt.Errorf("crc64 is not consistent, client: %s, server: %s", crc64Str, crc64ErrorStr), err)
	}
}

func TestHandleBosClientOption(t *testing.T) {
	// new bceClient, bosContext
	bosContext := &BosContext{}
	testHttpClient := &http.Client{}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	config := &bce.BceClientConfiguration{
		Endpoint:                  "endpoint",
		Region:                    "bce.DEFAULT_REGION",
		UserAgent:                 "bce.DEFAULT_USER_AGENT",
		Credentials:               &auth.BceCredentials{AccessKeyId: "ak", SecretAccessKey: "sk"},
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
		HTTPClient:                testHttpClient,
	}
	bceClient, err := bce.NewBceClientWithExclusiveHTTPClient(config, &auth.BceV1Signer{})
	ExpectEqual(t, nil, err)
	ExpectEqual(t, bceClient.HTTPClient, testHttpClient)
	ExpectEqual(t, false, bosContext.EnableCalcMd5)
	ExpectEqual(t, "", bosContext.ApiVersion)

	// HandleBosClientOptions
	mockHttpClient := util.NewMockHTTPClient()
	option := HTTPClient(mockHttpClient)
	options := []Option{option}
	option = EnableCalcMd5(true)
	options = append(options, option)
	option = ApiVersion(API_VERSION_V2)
	options = append(options, option)
	err = HandleBosClientOptions(bceClient, bosContext, options)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, bceClient.HTTPClient, mockHttpClient)
	ExpectEqual(t, true, bosContext.EnableCalcMd5)
	ExpectEqual(t, API_VERSION_V2, bosContext.ApiVersion)

	// handleBceClientOptions
	option = HTTPClient(testHttpClient)
	options1 := []Option{option}
	option = EnableCalcMd5(false)
	options1 = append(options1, option)
	option = ApiVersion(API_VERSION_V1)
	options1 = append(options1, option)

	err = handleBceClientOptions(bceClient, options1)
	ExpectEqual(t, nil, err)
	err = handleBosContextOptions(bosContext, options1)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, bceClient.HTTPClient, testHttpClient)
	ExpectEqual(t, false, bosContext.EnableCalcMd5)
	ExpectEqual(t, API_VERSION_V1, bosContext.ApiVersion)
	var client1 bce.Client = nil
	ExpectEqual(t, fmt.Errorf("unknown bceClient type"), handleBceClientOptions(client1, options1))
}

func TestOptions(t *testing.T) {
	ExpectEqual(t, nil, StorageClass("sdhfiuerg"))
	ExpectEqual(t, true, StorageClass(STORAGE_CLASS_ARCHIVE) != nil)
	ExpectEqual(t, nil, TrafficLimit(TRAFFIC_LIMIT_MAX+1))
	ExpectEqual(t, nil, TrafficLimit(TRAFFIC_LIMIT_MIN-1))
	ExpectEqual(t, true, TrafficLimit(TRAFFIC_LIMIT_MIN+1) != nil)
	ExpectEqual(t, nil, TaggingStr(""))
	ExpectEqual(t, nil, TaggingStr("dsafhierhg"))
	ExpectEqual(t, true, TaggingStr("key1=val1") != nil)
	ExpectEqual(t, nil, setBceClient("key string", ""))
	ExpectEqual(t, true, setBceClient("key string", "value") != nil)
	ExpectEqual(t, nil, setBosContext("key string", ""))
	ExpectEqual(t, true, setBosContext("key string", "value") != nil)
	ExpectEqual(t, nil, setSpecifiedTagParam(optionBceClient, "key", ""))
	ExpectEqual(t, true, setSpecifiedTagParam(optionBceClient, "key", "value") != nil)
	ExpectEqual(t, nil, IfMatch(""))
	ExpectEqual(t, nil, IfNoneMatch(""))
	ExpectEqual(t, nil, IfModifiedSince("dsnfisdf"))
	ExpectEqual(t, nil, IfUnModifiedSince("dsnfisdf"))
	ExpectEqual(t, true, IfModifiedSince(time.Now().Format(HTTPTimeFormat)) != nil)
	ExpectEqual(t, true, IfUnModifiedSince(time.Now().Format(HTTPTimeFormat)) != nil)
}

func TestHandleGetOptions(t *testing.T) {
	bosResp := &BosResponse{}
	myHttpResp := &my_http.Response{}
	httpResp := &http.Response{
		Header: http.Header{
			my_http.ETAG:                []string{"a44cc9bab11cbd156984767aad637851"},
			my_http.CACHE_CONTROL:       []string{"CACHE_CONTROL"},
			my_http.CONTENT_DISPOSITION: []string{"CONTENT_DISPOSITION"},
			my_http.CONTENT_LENGTH:      []string{"123456"},
			my_http.CONTENT_TYPE:        []string{"CONTENT_TYPE"},
			my_http.CONTENT_MD5:         []string{"CONTENT_MD5"},
			my_http.EXPIRES:             []string{"EXPIRES"},
			my_http.LAST_MODIFIED:       []string{"LAST_MODIFIED"},
			my_http.CONTENT_LANGUAGE:    []string{"CONTENT_LANGUAGE"},
			my_http.CONTENT_ENCODING:    []string{"CONTENT_ENCODING"},

			toHttpHeaderKey(my_http.BCE_CONTENT_SHA256):                 []string{"BCE_CONTENT_SHA256"},
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC32):                  []string{"BCE_CONTENT_CRC32"},
			toHttpHeaderKey(my_http.BCE_STORAGE_CLASS):                  []string{"BCE_STORAGE_CLASS"},
			toHttpHeaderKey(my_http.BCE_VERSION_ID):                     []string{"BCE_VERSION_ID"},
			toHttpHeaderKey(my_http.BCE_OBJECT_TYPE):                    []string{"BCE_OBJECT_TYPE"},
			toHttpHeaderKey(my_http.BCE_NEXT_APPEND_OFFSET):             []string{"123457"},
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC32C):                 []string{"BCE_CONTENT_CRC32C"},
			toHttpHeaderKey(my_http.BCE_EXPIRATION_DATE):                []string{"BCE_EXPIRATION_DATE"},
			toHttpHeaderKey(my_http.BCE_SERVER_SIDE_ENCRYPTION):         []string{"BCE_SERVER_SIDE_ENCRYPTION"},
			toHttpHeaderKey(my_http.BCE_SERVER_SIDE_ENCRYPTION_KEY):     []string{"BCE_SERVER_SIDE_ENCRYPTION_KEY"},
			toHttpHeaderKey(my_http.BCE_SERVER_SIDE_ENCRYPTION_KEY_MD5): []string{"BCE_SERVER_SIDE_ENCRYPTION_KEY_MD5"},
			toHttpHeaderKey(my_http.BCE_SERVER_SIDE_ENCRYPTION_KEY_ID):  []string{"BCE_SERVER_SIDE_ENCRYPTION_KEY_ID"},
			toHttpHeaderKey(my_http.BCE_OBJECT_RETENTION_DATE):          []string{"BCE_OBJECT_RETENTION_DATE"},
			toHttpHeaderKey(my_http.BCE_TAGGING_COUNT):                  []string{"3"},
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA):              []string{"BCE_CONTENT_CRC64ECMA"},
			toHttpHeaderKey(my_http.BCE_USER_METADATA_PREFIX + "key1"):  []string{"value1"},
			toHttpHeaderKey(my_http.BCE_USER_METADATA_PREFIX + "key2"):  []string{"value2"},
		},
	}
	myHttpResp.SetHttpResponse(httpResp)
	bosResp.SetHttpResponse(myHttpResp)
	objMeta := &ObjectMeta{}
	getOptions := getObjectMetaOptions(objMeta)
	ExpectEqual(t, nil, handleGetOptions(bosResp, getOptions))
	ExpectEqual(t, "a44cc9bab11cbd156984767aad637851", objMeta.ETag)
	ExpectEqual(t, "CACHE_CONTROL", objMeta.CacheControl)
	ExpectEqual(t, "CONTENT_DISPOSITION", objMeta.ContentDisposition)
	ExpectEqual(t, 123456, objMeta.ContentLength)
	ExpectEqual(t, "CONTENT_MD5", objMeta.ContentMD5)
	ExpectEqual(t, "EXPIRES", objMeta.Expires)
	ExpectEqual(t, "LAST_MODIFIED", objMeta.LastModified)
	ExpectEqual(t, "CONTENT_LANGUAGE", objMeta.ContentLanguage)
	ExpectEqual(t, "CONTENT_ENCODING", objMeta.ContentEncoding)
	ExpectEqual(t, "BCE_CONTENT_SHA256", objMeta.ContentSha256)
	ExpectEqual(t, "BCE_CONTENT_CRC32", objMeta.ContentCrc32)
	ExpectEqual(t, "CONTENT_DISPOSITION", objMeta.ContentDisposition)
	ExpectEqual(t, "BCE_STORAGE_CLASS", objMeta.StorageClass)
	ExpectEqual(t, "BCE_VERSION_ID", objMeta.VersionId)
	ExpectEqual(t, "BCE_OBJECT_TYPE", objMeta.ObjectType)
	ExpectEqual(t, "123457", objMeta.NextAppendOffset)
	ExpectEqual(t, "BCE_CONTENT_CRC32C", objMeta.ContentCrc32c)
	ExpectEqual(t, "BCE_EXPIRATION_DATE", objMeta.ExpirationDate)
	ExpectEqual(t, "BCE_SERVER_SIDE_ENCRYPTION", objMeta.Encryption.ServerSideEncryption)
	ExpectEqual(t, "BCE_SERVER_SIDE_ENCRYPTION_KEY", objMeta.Encryption.SSECKey)
	ExpectEqual(t, "BCE_SERVER_SIDE_ENCRYPTION_KEY_MD5", objMeta.Encryption.SSECKeyMD5)
	ExpectEqual(t, "BCE_SERVER_SIDE_ENCRYPTION_KEY_ID", objMeta.Encryption.SSEKmsKeyId)
	ExpectEqual(t, "BCE_OBJECT_RETENTION_DATE", objMeta.RetentionDate)
	ExpectEqual(t, 3, objMeta.objectTagCount)
	ExpectEqual(t, "BCE_CONTENT_CRC64ECMA", objMeta.ContentCrc64ECMA)
	ExpectEqual(t, map[string]string{"Key1": "value1", "Key2": "value2"}, objMeta.UserMeta)

	bosResp1 := &BosResponse{}
	myHttpResp1 := &my_http.Response{}
	httpResp1 := &http.Response{
		Header: http.Header{
			toHttpHeaderKey(my_http.BCE_TAG): []string{"key1=value1,key2=value2"},
		},
	}
	myHttpResp1.SetHttpResponse(httpResp1)
	bosResp1.SetHttpResponse(myHttpResp1)
	testStruct := struct {
		Tags []string
	}{
		Tags: []string{},
	}

	getOptions1 := []GetOption{
		getHeader(my_http.BCE_TAG, &(testStruct.Tags)),
	}
	ExpectEqual(t, nil, handleGetOptions(bosResp1, getOptions1))
	ExpectEqual(t, 2, len(testStruct.Tags))
	ExpectEqual(t, "key1=value1", testStruct.Tags[0])
	ExpectEqual(t, "key2=value2", testStruct.Tags[1])
}

// TestSetSpecifiedTagParam_NilValue 是用于测试 SetSpecifiedTagParam_NilValue
// generated by Comate
func TestSetSpecifiedTagParam_NilValue(t *testing.T) {
	// 测试场景：value 为 nil 的情况，应该覆盖第438-439行
	// 当 value 为 nil 时，函数应该返回一个 Option，该 Option 执行时不会设置参数
	option := setSpecifiedTagParam("testTag", "testKey", nil)

	if option == nil {
		t.Error("Expected non-nil option function for nil value")
		return
	}

	// 执行返回的 Option 函数
	params := make(map[string]optionValue)
	err := option(params)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// 验证参数没有被设置
	if _, exists := params["testKey"]; exists {
		t.Error("Params should not contain testKey when value is nil")
	}
}

// Test_handleBceClientOptions_OptionError 是用于测试 _handleBceClientOptions_OptionError
// generated by Comate
func Test_handleBceClientOptions_OptionError(t *testing.T) {
	// 创建一个模拟的BceClient
	mockClient := &bce.BceClient{}

	// 创建一个会返回错误的Option
	errorOption := func(params map[string]optionValue) error {
		return fmt.Errorf("mock option error")
	}

	options := []Option{errorOption}

	err := handleBceClientOptions(mockClient, options)

	if err == nil {
		t.Error("Expected error from option execution, but got nil")
	}

	expectedError := "mock option error"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s', but got '%s'", expectedError, err.Error())
	}
}

// TestHandleBosContextOptions_OptionReturnsError 是用于测试 HandleBosContextOptions_OptionReturnsError
// generated by Comate
func TestHandleBosContextOptions_OptionReturnsError(t *testing.T) {
	ctx := &BosContext{}
	errorOption := func(params map[string]optionValue) error {
		return &bce.BceServiceError{Code: "TestError", Message: "Test error from option"}
	}
	options := []Option{errorOption}
	err := handleBosContextOptions(ctx, options)
	if err == nil {
		t.Errorf("Expected error from option, but got nil")
	}
	bceErr, ok := err.(*bce.BceServiceError)
	if !ok {
		t.Errorf("Expected BceServiceError, got %T", err)
	}
	if bceErr.Code != "TestError" {
		t.Errorf("Expected error code 'TestError', got '%s'", bceErr.Code)
	}
}

// TestSetHeader_NilValue 是用于测试 SetHeader_NilValue
// generated by Comate
func TestSetHeader_NilValue(t *testing.T) {
	key := "Test-Header"
	var value interface{} = nil

	option := setHeader(key, value)

	// 当 value 为 nil 时，setHeader 应该返回一个 Option 函数
	if option == nil {
		t.Error("setHeader should return an Option function even when value is nil")
		return
	}

	// 执行返回的 Option 函数
	params := make(map[string]optionValue)
	err := option(params)

	// 验证没有错误返回
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// 验证参数 map 没有被修改（因为 value 为 nil 时应该直接返回）
	if len(params) != 0 {
		t.Errorf("Expected params to be empty when value is nil, but got %d entries", len(params))
	}
}

// TestHandleBosClientOptions_BosContextError 是用于测试 HandleBosClientOptions_BosContextError
// generated by Comate
func TestHandleBosClientOptions_BosContextError(t *testing.T) {
	client := &bce.BceClient{}
	ctx := &BosContext{}

	// 创建一个会返回错误的option
	errorOption := func(params map[string]optionValue) error {
		return fmt.Errorf("bos context option error")
	}
	options := []Option{errorOption}

	err := HandleBosClientOptions(client, ctx, options)
	if err == nil {
		t.Error("Expected error from handleBosContextOptions, but got nil")
	}

	expectedErrorMsg := "BosContext Options: bos context option error"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message '%s', but got '%s'", expectedErrorMsg, err.Error())
	}
}

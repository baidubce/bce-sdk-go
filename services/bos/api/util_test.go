package api

import (
	"hash/crc64"
	"net/http"
	"runtime"
	"strconv"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
	my_http "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

var (
	bceServiceErro503 = bce.NewBceServiceError("Unavailable", "Service Unavailable", "", http.StatusServiceUnavailable)
	bceServiceErro408 = bce.NewBceServiceError("Timeout", "Request Timeout", "", http.StatusRequestTimeout)
)

func ExpectEqual(t *testing.T, exp interface{}, act interface{}) bool {
	if !util.Equal(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d: missmatch, expect %v but %v", file, line, exp, act)
		return false
	}
	return true
}

func TestUtil(t *testing.T) {
	_, ok := VALID_RESTORE_TIER[RESTORE_TIER_STANDARD]
	ExpectEqual(t, true, ok)
	_, ok = VALID_RESTORE_TIER[RESTORE_TIER_EXPEDITED]
	ExpectEqual(t, true, ok)
	_, ok = VALID_RESTORE_TIER[RESTORE_TIER_LOWCOST]
	ExpectEqual(t, true, ok)
	_, ok = VALID_RESTORE_TIER["restore_tier_unknown"]
	ExpectEqual(t, false, ok)

	bucket := "test-bucket"
	object := "test-object"
	ExpectEqual(t, "/"+bucket, getBucketUri(bucket))
	ExpectEqual(t, "/"+bucket+"/"+object, getObjectUri(bucket, object))
	ExpectEqual(t, "", getCnameUri(""))
	ExpectEqual(t, "/", getCnameUri("/"))
	ExpectEqual(t, "/", getCnameUri("/path"))
	ExpectEqual(t, "/dir", getCnameUri("/path/dir"))

	ExpectEqual(t, true, validMetadataDirective(METADATA_DIRECTIVE_COPY))
	ExpectEqual(t, true, validMetadataDirective(METADATA_DIRECTIVE_COPY))
	ExpectEqual(t, false, validMetadataDirective("unknown_metadata_directive"))

	tooLengthTagging := make([]byte, 4096)
	for i := 0; i < len(tooLengthTagging); i++ {
		tooLengthTagging[i] = 't'
	}
	tooLengthKey := make([]byte, 150)
	for i := 0; i < len(tooLengthKey); i++ {
		tooLengthKey[i] = 'k'
	}
	tooLengthVal := make([]byte, 300)
	for i := 0; i < len(tooLengthVal); i++ {
		tooLengthVal[i] = 'v'
	}

	testTagging := []struct {
		tag string
		ok  bool
		res string
	}{
		{"", false, ""},
		{"k=v=vv", false, ""},
		{"testtagging", false, ""},
		{string(tooLengthTagging), false, ""},
		{string(tooLengthKey) + "=val", false, ""},
		{"key=" + string(tooLengthVal), false, ""},
		{"key1=val1&key2=val2", true, "key1=val1&key2=val2"},
	}

	for _, v := range testTagging {
		ok, res := validObjectTagging(v.tag)
		ExpectEqual(t, v.ok, ok)
		ExpectEqual(t, v.res, res)
	}

	testTagsToStr := []struct {
		tags map[string]string
		res  []string
	}{
		{make(map[string]string), []string{""}},
		{map[string]string{string(tooLengthKey): "val"}, []string{""}},
		{map[string]string{"key": string(tooLengthVal)}, []string{""}},
		{map[string]string{"key1": "val1", "key2": "val2"}, []string{"key1=val1&key2=val2", "key2=val2&key1=val1"}},
	}
	for _, v := range testTagsToStr {
		match := false
		for _, res := range v.res {
			match = match || util.Equal(res, taggingMapToStr(v.tags))
		}
		ExpectEqual(t, true, match)
	}

	httpKey := []struct {
		input  string
		output string
	}{
		{"x-bce-header1", "X-Bce-Header1"},
		{"", ""},
		{"1-2-3", "1-2-3"},
	}
	for _, v := range httpKey {
		ExpectEqual(t, v.output, toHttpHeaderKey(v.input))
	}
}

func TestSendRequestWithBodyV2(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	//prepare request / response
	req := &BosRequest{}
	resp := &BosResponse{}
	bosContext := newDefaultBosContext()
	bucket := "test-bucket"
	object := "prefix/test-object"
	test_url := getObjectUri(bucket, object)
	reqBody := "this is a request body string for testing SendRequest."
	// calc crc64 of reqBody
	crc64Hasher := crc64.New(crc64.MakeTable(crc64.ECMA))
	n, err := crc64Hasher.Write([]byte(reqBody))
	ExpectEqual(t, len(reqBody), n)
	ExpectEqual(t, nil, err)
	crc64Str := strconv.FormatUint(crc64Hasher.Sum64(), 10)
	req.SetUri(test_url)
	req.SetMethod(http.MethodPut)
	req.SetBucket(bucket)

	//case1: body is TeeReadNopCloser, return ok
	body1, err := bce.NewBodyFromStringV2(reqBody, true)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, body1 != nil)
	req.SetBody(body1)
	AddCrc64Check(req, resp)
	for _, tracker := range req.Tracker {
		ExpectEqual(t, nil, tracker(req))
	}
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str,
		}),
	}
	mockHttpClient := util.NewMockHTTPClient(options...)
	client.HTTPClient = mockHttpClient
	err = SendRequest(client, req, resp, bosContext)
	ExpectEqual(t, nil, err)
	//case1.1: send request v2
	bosContext.ApiVersion = API_VERSION_V2
	//reset body
	body11, err := bce.NewBodyFromStringV2(reqBody, true)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, body11 != nil)
	req.SetBody(body11)
	//reset tracker and handler
	req.Tracker = []RequestTracker{}
	resp.Handler = []ResponseHandler{}
	AddCrc64Check(req, resp)
	for _, tracker := range req.Tracker {
		ExpectEqual(t, nil, tracker(req))
	}
	err = SendRequest(client, req, resp, bosContext)
	ExpectEqual(t, nil, err)

	//case2: body is TeeReadNopCloser, return 500, retry with BackupEndpoint
	bosContext.ApiVersion = API_VERSION_V1
	options2 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusServiceUnavailable),
		util.SetStatusMsg(http.StatusText(http.StatusServiceUnavailable)),
		util.AddHeaders(map[string]string{
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str,
		}),
	}
	mockHttpClient2 := util.NewMockHTTPClient(options2...)
	client.HTTPClient = mockHttpClient2
	client.Config.BackupEndpoint = "bj-bk.bcebos.com"
	//reset body
	body21, err := bce.NewBodyFromStringV2(reqBody, true)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, body21 != nil)
	req.SetBody(body21)
	//reset tracker and handler
	req.Tracker = []RequestTracker{}
	resp.Handler = []ResponseHandler{}
	AddCrc64Check(req, resp)
	for _, tracker := range req.Tracker {
		ExpectEqual(t, nil, tracker(req))
	}
	err = SendRequest(client, req, resp, bosContext)
	ExpectEqual(t, bceServiceErro503, err)
	// case2.1: send request v2
	bosContext.ApiVersion = API_VERSION_V2
	//reset body
	body22, err := bce.NewBodyFromStringV2(reqBody, true)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, body22 != nil)
	req.SetBody(body22)
	//reset tracker and handler
	req.Tracker = []RequestTracker{}
	resp.Handler = []ResponseHandler{}
	AddCrc64Check(req, resp)
	for _, tracker := range req.Tracker {
		ExpectEqual(t, nil, tracker(req))
	}
	err = SendRequest(client, req, resp, bosContext)
	ExpectEqual(t, bceServiceErro503, err)
}

func TestSendRequestWithBodyV1(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	//prepare request / response
	req := &BosRequest{}
	resp := &BosResponse{}
	bosContext := newDefaultBosContext()
	bucket := "test-bucket"
	object := "prefix/test-object"
	test_url := getObjectUri(bucket, object)
	reqBody := "this is a request body string for testing SendRequest."
	// calc crc64 of reqBody
	crc64Hasher := crc64.New(crc64.MakeTable(crc64.ECMA))
	n, err := crc64Hasher.Write([]byte(reqBody))
	ExpectEqual(t, len(reqBody), n)
	ExpectEqual(t, nil, err)
	crc64Str := strconv.FormatUint(crc64Hasher.Sum64(), 10)
	req.SetUri(test_url)
	req.SetMethod(http.MethodPut)
	req.SetBucket(bucket)

	//case1: old body, send request v1, return ok
	body1, err := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, body1 != nil)
	req.SetBody(body1)
	AddCrc64Check(req, resp)
	for _, tracker := range req.Tracker {
		ExpectEqual(t, nil, tracker(req))
	}
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str,
		}),
	}
	mockHttpClient := util.NewMockHTTPClient(options...)
	client.HTTPClient = mockHttpClient
	err = SendRequest(client, req, resp, bosContext)
	ExpectEqual(t, nil, err)
	//case1.1: old body, send request v2, return ok
	bosContext.ApiVersion = API_VERSION_V2
	//reset body
	body11, err := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, body11 != nil)
	req.SetBody(body11)
	//reset tracker and handler
	req.Tracker = []RequestTracker{}
	resp.Handler = []ResponseHandler{}
	AddCrc64Check(req, resp)
	for _, tracker := range req.Tracker {
		ExpectEqual(t, nil, tracker(req))
	}
	err = SendRequest(client, req, resp, bosContext)
	ExpectEqual(t, nil, err)

	//case2: old body, return 500, send request v1, retry
	bosContext.ApiVersion = API_VERSION_V1
	options2 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusServiceUnavailable),
		util.SetStatusMsg(http.StatusText(http.StatusServiceUnavailable)),
		util.AddHeaders(map[string]string{
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str,
		}),
	}
	mockHttpClient2 := util.NewMockHTTPClient(options2...)
	client.HTTPClient = mockHttpClient2
	client.Config.BackupEndpoint = "bj-bk.bcebos.com"
	//reset body
	body21, err := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, body21 != nil)
	req.SetBody(body21)
	//reset tracker and handler
	req.Tracker = []RequestTracker{}
	resp.Handler = []ResponseHandler{}
	AddCrc64Check(req, resp)
	for _, tracker := range req.Tracker {
		ExpectEqual(t, nil, tracker(req))
	}
	err = SendRequest(client, req, resp, bosContext)
	ExpectEqual(t, bceServiceErro503, err)

	// case2.1: old body, return 500, send request v2, not retry
	bosContext.ApiVersion = API_VERSION_V2
	//reset body
	body22, err := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, body22 != nil)
	req.SetBody(body22)
	//reset tracker and handler
	req.Tracker = []RequestTracker{}
	resp.Handler = []ResponseHandler{}
	AddCrc64Check(req, resp)
	for _, tracker := range req.Tracker {
		ExpectEqual(t, nil, tracker(req))
	}
	err = SendRequest(client, req, resp, bosContext)
	ExpectEqual(t, bceServiceErro503, err)
}

func TestGetObjectMetaOptions(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	//prepare request / response
	req1 := &BosRequest{}
	resp1 := &BosResponse{}
	req1.SetUri("test_url")
	req1.SetMethod(http.MethodPut)
	req1.SetBucket("bucket")
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			my_http.ETAG:                               "a44cc9bab11cbd156984767aad637851",
			my_http.CACHE_CONTROL:                      "CACHE_CONTROL",
			my_http.CONTENT_DISPOSITION:                "CONTENT_DISPOSITION",
			my_http.CONTENT_LENGTH:                     "123456",
			my_http.CONTENT_TYPE:                       "CONTENT_TYPE",
			my_http.CONTENT_MD5:                        "CONTENT_MD5",
			my_http.EXPIRES:                            "EXPIRES",
			my_http.LAST_MODIFIED:                      "LAST_MODIFIED",
			my_http.CONTENT_LANGUAGE:                   "CONTENT_LANGUAGE",
			my_http.CONTENT_ENCODING:                   "CONTENT_ENCODING",
			my_http.BCE_CONTENT_SHA256:                 "BCE_CONTENT_SHA256",
			my_http.BCE_CONTENT_CRC32:                  "BCE_CONTENT_CRC32",
			my_http.BCE_STORAGE_CLASS:                  "BCE_STORAGE_CLASS",
			my_http.BCE_VERSION_ID:                     "BCE_VERSION_ID",
			my_http.BCE_OBJECT_TYPE:                    "BCE_OBJECT_TYPE",
			my_http.BCE_NEXT_APPEND_OFFSET:             "123457",
			my_http.BCE_CONTENT_CRC32C:                 "BCE_CONTENT_CRC32C",
			my_http.BCE_EXPIRATION_DATE:                "BCE_EXPIRATION_DATE",
			my_http.BCE_SERVER_SIDE_ENCRYPTION:         "BCE_SERVER_SIDE_ENCRYPTION",
			my_http.BCE_SERVER_SIDE_ENCRYPTION_KEY:     "BCE_SERVER_SIDE_ENCRYPTION_KEY",
			my_http.BCE_SERVER_SIDE_ENCRYPTION_KEY_MD5: "BCE_SERVER_SIDE_ENCRYPTION_KEY_MD5",
			my_http.BCE_SERVER_SIDE_ENCRYPTION_KEY_ID:  "BCE_SERVER_SIDE_ENCRYPTION_KEY_ID",
			my_http.BCE_OBJECT_RETENTION_DATE:          "BCE_OBJECT_RETENTION_DATE",
			my_http.BCE_TAGGING_COUNT:                  "3",
			my_http.BCE_CONTENT_CRC64ECMA:              "BCE_CONTENT_CRC64ECMA",
			my_http.BCE_USER_METADATA_PREFIX + "key1":  "value1",
			my_http.BCE_USER_METADATA_PREFIX + "key2":  "value2",
		}),
	}
	mockHttpClient := util.NewMockHTTPClient(options...)
	client.HTTPClient = mockHttpClient
	err = SendRequest(client, req1, resp1, newDefaultBosContext())
	ExpectEqual(t, nil, err)
	objMeta1 := &ObjectMeta{}
	getOptions1 := getObjectMetaOptions(objMeta1)
	ExpectEqual(t, nil, handleGetOptions(resp1, getOptions1))
	ExpectEqual(t, "a44cc9bab11cbd156984767aad637851", objMeta1.ETag)
	ExpectEqual(t, "CACHE_CONTROL", objMeta1.CacheControl)
	ExpectEqual(t, "CONTENT_DISPOSITION", objMeta1.ContentDisposition)
	ExpectEqual(t, 123456, objMeta1.ContentLength)
	ExpectEqual(t, "CONTENT_MD5", objMeta1.ContentMD5)
	ExpectEqual(t, "EXPIRES", objMeta1.Expires)
	ExpectEqual(t, "LAST_MODIFIED", objMeta1.LastModified)
	ExpectEqual(t, "CONTENT_LANGUAGE", objMeta1.ContentLanguage)
	ExpectEqual(t, "CONTENT_ENCODING", objMeta1.ContentEncoding)
	ExpectEqual(t, "BCE_CONTENT_SHA256", objMeta1.ContentSha256)
	ExpectEqual(t, "BCE_CONTENT_CRC32", objMeta1.ContentCrc32)
	ExpectEqual(t, "CONTENT_DISPOSITION", objMeta1.ContentDisposition)
	ExpectEqual(t, "BCE_STORAGE_CLASS", objMeta1.StorageClass)
	ExpectEqual(t, "BCE_VERSION_ID", objMeta1.VersionId)
	ExpectEqual(t, "BCE_OBJECT_TYPE", objMeta1.ObjectType)
	ExpectEqual(t, "123457", objMeta1.NextAppendOffset)
	ExpectEqual(t, "BCE_CONTENT_CRC32C", objMeta1.ContentCrc32c)
	ExpectEqual(t, "BCE_EXPIRATION_DATE", objMeta1.ExpirationDate)

	ExpectEqual(t, "BCE_SERVER_SIDE_ENCRYPTION", objMeta1.Encryption.ServerSideEncryption)
	ExpectEqual(t, "BCE_SERVER_SIDE_ENCRYPTION_KEY", objMeta1.Encryption.SSECKey)
	ExpectEqual(t, "BCE_SERVER_SIDE_ENCRYPTION_KEY_MD5", objMeta1.Encryption.SSECKeyMD5)
	ExpectEqual(t, "BCE_SERVER_SIDE_ENCRYPTION_KEY_ID", objMeta1.Encryption.SSEKmsKeyId)
	ExpectEqual(t, "BCE_OBJECT_RETENTION_DATE", objMeta1.RetentionDate)
	ExpectEqual(t, 3, objMeta1.objectTagCount)
	ExpectEqual(t, "BCE_CONTENT_CRC64ECMA", objMeta1.ContentCrc64ECMA)
	ExpectEqual(t, map[string]string{"Key1": "value1", "Key2": "value2"}, objMeta1.UserMeta)
}

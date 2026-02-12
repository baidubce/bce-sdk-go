package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	my_http "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

func TestPutObject(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	reqBody := "this is a request body string for testing put object."
	body, err := bce.NewBodyFromStringV2(reqBody, false)
	ExpectEqual(t, nil, err)
	args := &PutObjectArgs{}
	// calc crc64
	crc64Hash := crc64.New(crc64.MakeTable(crc64.ECMA))
	n, err := crc64Hash.Write([]byte(reqBody))
	ExpectEqual(t, n, len(reqBody))
	ExpectEqual(t, err, nil)
	crc64Str := strconv.FormatUint(crc64Hash.Sum64(), 10)
	// calc crc32c
	crc32cHash := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	n, err = crc32cHash.Write([]byte(reqBody))
	ExpectEqual(t, n, len(reqBody))
	ExpectEqual(t, err, nil)
	crc32cStr := strconv.FormatUint(uint64(crc32cHash.Sum32()), 10)
	// case1: body is nil
	etag, res, err := PutObject(client, bucket, object, nil, nil, nil)
	ExpectEqual(t, "", etag)
	ExpectEqual(t, nil, res)
	ExpectEqual(t, bce.NewBceClientError("PutObject body should not be emtpy"), err)

	// case2: args.ContentLength > body.Size()
	args.ContentLength = 1024 * 1024
	etag, res, err = PutObject(client, bucket, object, body, args, nil)
	ExpectEqual(t, "", etag)
	ExpectEqual(t, nil, res)
	ExpectEqual(t, bce.NewBceClientError(fmt.Sprintf("ContentLength %d is bigger than body size %d",
		args.ContentLength, body.Size())), err)
	args.ContentLength = int64(len(reqBody))

	// case3: args.TrafficLimit
	args.TrafficLimit = TRAFFIC_LIMIT_MIN - 1
	etag, res, err = PutObject(client, bucket, object, body, args, nil)
	ExpectEqual(t, "", etag)
	ExpectEqual(t, nil, res)
	ExpectEqual(t, fmt.Errorf("TrafficLimit must between %d ~ %d, current value:%d",
		TRAFFIC_LIMIT_MIN, TRAFFIC_LIMIT_MAX, args.TrafficLimit), err)
	args.TrafficLimit = TRAFFIC_LIMIT_MIN + 1

	// case4: invalid storage class
	args.StorageClass = "dkfneirgerg"
	etag, res, err = PutObject(client, bucket, object, body, args, nil)
	ExpectEqual(t, "", etag)
	ExpectEqual(t, nil, res)
	ExpectEqual(t, fmt.Errorf("invalid storage class value: %s", args.StorageClass), err)
	args.StorageClass = STORAGE_CLASS_STANDARD

	// case5: args.ContentLength < body.Size()
	args.ContentLength = int64(len(reqBody) / 2)
	// calc crc64
	crc64Hash1 := crc64.New(crc64.MakeTable(crc64.ECMA))
	n, err = crc64Hash1.Write([]byte(reqBody)[:args.ContentLength])
	ExpectEqual(t, n, args.ContentLength)
	ExpectEqual(t, err, nil)
	crc64Str1 := strconv.FormatUint(crc64Hash1.Sum64(), 10)
	// calc crc32c
	crc32cHash1 := crc32.New(crc32.MakeTable(crc32.Castagnoli))
	n, err = crc32cHash.Write([]byte(reqBody)[:args.ContentLength])
	ExpectEqual(t, n, args.ContentLength)
	ExpectEqual(t, err, nil)
	crc32cStr1 := strconv.FormatUint(uint64(crc32cHash1.Sum32()), 10)
	// mock http client
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			toHttpHeaderKey(my_http.BCE_VERSION_ID):        "AKyQ9DRhhoY=",
			toHttpHeaderKey(my_http.BCE_STORAGE_CLASS):     STORAGE_CLASS_STANDARD,
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC32C):    crc32cStr1,
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str1,
			my_http.ETAG: "9b2cf535f27731c974343645a3985328",
		}),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	etag, res, err = PutObject(client, bucket, object, body, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "9b2cf535f27731c974343645a3985328", etag)
	ExpectEqual(t, "AKyQ9DRhhoY=", res.VersionId)
	ExpectEqual(t, STORAGE_CLASS_STANDARD, res.StorageClass)
	ExpectEqual(t, crc32cStr1, res.ContentCrc32c)
	ExpectEqual(t, crc64Str1, res.ContentCrc64ECMA)
	args.ContentLength = int64(len(reqBody))

	// case6 : all is ok
	// prepare parameters
	args.ContentLength = int64(len(reqBody))
	args.CacheControl = "testcachecontrol"
	args.CannedAcl = "private"
	args.ContentCrc64ECMA = crc64Str
	args.StorageClass = STORAGE_CLASS_STANDARD
	args.TrafficLimit = TRAFFIC_LIMIT_MIN + 1
	// mock http client
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			toHttpHeaderKey(my_http.BCE_VERSION_ID):        "AKyQ9DRhhoY=",
			toHttpHeaderKey(my_http.BCE_STORAGE_CLASS):     STORAGE_CLASS_STANDARD,
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC32C):    crc32cStr,
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str,
			my_http.ETAG: "9b2cf535f27731c974343645a3985328",
		}),
	}
	mockHttpClient := util.NewMockHTTPClient(options...)
	ExpectEqual(t, true, mockHttpClient != nil)
	client.HTTPClient = mockHttpClient

	// put object
	body1, err := bce.NewBodyFromStringV2(reqBody, false)
	ExpectEqual(t, nil, err)
	etag, res, err = PutObject(client, bucket, object, body1, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "9b2cf535f27731c974343645a3985328", etag)
	ExpectEqual(t, "AKyQ9DRhhoY=", res.VersionId)
	ExpectEqual(t, STORAGE_CLASS_STANDARD, res.StorageClass)
	ExpectEqual(t, crc32cStr, res.ContentCrc32c)
	ExpectEqual(t, crc64Str, res.ContentCrc64ECMA)

	// case7: args.Process = "callback-ksdfigreger"
	args.Process = "callback-ksdfigreger"
	respBody7 := `{
		"callback": {
			"result": "${callback_result}"
		}
	}`
	// mock http client
	options7 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(respBody7),
		util.AddHeaders(map[string]string{
			toHttpHeaderKey(my_http.BCE_VERSION_ID):        "AKyQ9DRhhoY=",
			toHttpHeaderKey(my_http.BCE_STORAGE_CLASS):     STORAGE_CLASS_STANDARD,
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC32C):    crc32cStr,
			toHttpHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str,
			my_http.ETAG: "9b2cf535f27731c974343645a3985328",
		}),
	}
	mockHttpClient7 := util.NewMockHTTPClient(options7...)
	ExpectEqual(t, true, mockHttpClient7 != nil)
	client.HTTPClient = mockHttpClient7

	// put object
	body7, err := bce.NewBodyFromStringV2(reqBody, false)
	ExpectEqual(t, nil, err)
	etag, res, err = PutObject(client, bucket, object, body7, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "9b2cf535f27731c974343645a3985328", etag)
	ExpectEqual(t, "AKyQ9DRhhoY=", res.VersionId)
	ExpectEqual(t, STORAGE_CLASS_STANDARD, res.StorageClass)
	ExpectEqual(t, crc32cStr, res.ContentCrc32c)
	ExpectEqual(t, crc64Str, res.ContentCrc64ECMA)
	ExpectEqual(t, "${callback_result}", res.Callback.Result)
	args.Process = ""

	// case8: handle options error
	body8, err := bce.NewBodyFromStringV2(reqBody, false)
	ExpectEqual(t, nil, err)
	etag, res, err = PutObject(client, bucket, object, body8, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)
	ExpectEqual(t, "", etag)

	// case9: send request error
	body9, err := bce.NewBodyFromStringV2(reqBody, false)
	ExpectEqual(t, nil, err)
	err9 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err9)
	etag, res, err = PutObject(client, bucket, object, body9, args, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)
	ExpectEqual(t, "", etag)

	// case10: resp is fail
	body10, err := bce.NewBodyFromStringV2(reqBody, false)
	ExpectEqual(t, nil, err)
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	etag, res, err = PutObject(client, bucket, object, body10, args, nil)
	ExpectEqual(t, bceServiceErro404.Error(), err.Error())
	ExpectEqual(t, nil, res)
	ExpectEqual(t, "", etag)
}

func TestPostObject(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test-bucket"
	object := "test-object"
	content := "this is a test-string for testing PostObject"

	//case1: args is nil
	res, err := PostObject(client, bucket, object, bytes.NewBufferString(content), nil, nil)
	ExpectEqual(t, bce.NewBceClientError("post object argument is nil."), err)
	ExpectEqual(t, nil, res)
	//case2: handlePostOptions error
	args := &PostObjectArgs{
		Expiration:         180 * time.Second,
		ContentLengthLower: 0,
		ContentLengthUpper: 4096,
	}
	res, err = PostObject(client, bucket, object, bytes.NewBufferString(content), args, nil, ErrorOption)
	ExpectEqual(t, postOptionError, err)
	ExpectEqual(t, nil, res)
	//case3: send request fail
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = PostObject(client, bucket, object, bytes.NewBufferString(content), args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket, "", "Post", err3, err)
	ExpectEqual(t, nil, res)
	//case4: response is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = PostObject(client, bucket, object, bytes.NewBufferString(content), args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
	//case5: all is ok
	options5 := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.CONTENT_MD5):       "Zh+ACfqOVqnQ6UoKZEOX1w==",
			http.CanonicalHeaderKey(my_http.ETAG):              "827ccb0eea8a706c4c34a16891f84e7b",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32): "1922069637",
		}),
	}
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
	res, err = PostObject(client, bucket, object, bytes.NewBufferString(content), args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "1922069637", res.ContentCrc32)
	ExpectEqual(t, "827ccb0eea8a706c4c34a16891f84e7b", res.ETag)
	ExpectEqual(t, "Zh+ACfqOVqnQ6UoKZEOX1w==", res.ContentMD5)
}
func TestOptionObject(t *testing.T) {
	// mock bos client
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	bucket := "test-bucket"
	object := "test-object"
	resource := fmt.Sprintf("/%s/%s", bucket, object)
	args := &OptionsObjectArgs{
		Origin:         "string",
		RequestMethod:  "Post",
		RequestHeaders: []string{"headr1", "header2"},
	}
	//case1: all is ok
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.ACCESS_CONTROL_ALLOW_CREDENTIALS): "true",
			http.CanonicalHeaderKey(my_http.ACCESS_CONTROL_ALLOW_HEADERS):     "header1,header2",
			http.CanonicalHeaderKey(my_http.ACCESS_CONTROL_ALLOW_METHODS):     "PUT,GET",
			http.CanonicalHeaderKey(my_http.ACCESS_CONTROL_ALLOW_ORIGIN):      "origin",
			http.CanonicalHeaderKey(my_http.ACCESS_CONTROL_EXPOSE_HEADERS):    "POST",
			http.CanonicalHeaderKey(my_http.ACCESS_CONTROL_MAX_AGE):           "10",
		}),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	res, err := OptionsObject(client, bucket, object, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, res.AllowCredentials)
	ExpectEqual(t, 2, len(res.AllowHeaders))
	ExpectEqual(t, "header1", res.AllowHeaders[0])
	ExpectEqual(t, "header2", res.AllowHeaders[1])
	ExpectEqual(t, 2, len(res.AllowMethods))
	ExpectEqual(t, "PUT", res.AllowMethods[0])
	ExpectEqual(t, "GET", res.AllowMethods[1])
	ExpectEqual(t, "origin", res.AllowOrigin)
	ExpectEqual(t, 1, len(res.ExposeHeaders))
	ExpectEqual(t, "POST", res.ExposeHeaders[0])
	ExpectEqual(t, 10, res.MaxAge)
	//case2: handle options error
	res, err = OptionsObject(client, bucket, object, args, nil, ErrorOption)
	ExpectEqual(t, optionError1, err)
	ExpectEqual(t, nil, res)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = OptionsObject(client, bucket, object, args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, resource, "", "Options", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	AttachMockHttpClientRespFail(t, client, util.RoundTripperOpts404)
	res, err = OptionsObject(client, bucket, object, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)
}
func TestCopyObject(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	source := "src-bucket/src-object"
	args := &CopyObjectArgs{}

	// case1: empty source
	res, err := CopyObject(client, bucket, object, "", args, nil)
	ExpectEqual(t, bce.NewBceClientError("copy source should not be null"), err)
	ExpectEqual(t, nil, res)

	// case2: invalid meta redirective
	args.MetadataDirective = "meta-redirective"
	res, err = CopyObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, bce.NewBceClientError("invalid metadata directive value: "+args.MetadataDirective), err)
	ExpectEqual(t, nil, res)
	args.MetadataDirective = METADATA_DIRECTIVE_COPY

	// case3: invalid storage class
	args.StorageClass = "storage-class"
	res, err = CopyObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, bce.NewBceClientError("invalid storage class value: "+args.StorageClass), err)
	ExpectEqual(t, nil, res)
	args.StorageClass = STORAGE_CLASS_STANDARD

	// case4: invalid traffic limit
	args.TrafficLimit = TRAFFIC_LIMIT_MAX + 1
	res, err = CopyObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, trafficLimitInvalidError(args.TrafficLimit), err)
	ExpectEqual(t, nil, res)
	args.TrafficLimit = TRAFFIC_LIMIT_MAX - 1

	// case5: handle options error
	res, err = CopyObject(client, bucket, object, source, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case6: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	res, err = CopyObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	//case7: resp is fail
	options7 := util.RoundTripperOpts404
	mockHttpClient7 := util.NewMockHTTPClient(options7...)
	ExpectEqual(t, true, mockHttpClient7 != nil)
	client.HTTPClient = mockHttpClient7
	res, err = CopyObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)

	//case8: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = CopyObject(client, bucket, object, source, args, nil)
	result5 := &CopyObjectResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)

	// case7: resp body is fail content
	respBody9 := `{
		"code":"InternalError",
		"message":"We encountered an internal error. Please try again.",
		"requestId":"52454655-5345-4420-4259-204e47494e58"
	}`
	options9 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			my_http.BCE_REQUEST_ID: "52454655-5345-4420-4259-204e47494e58",
			my_http.BCE_VERSION_ID: "version-id",
		}),
		util.SetRespBody(respBody9),
	}
	mockHttpClient9 := util.NewMockHTTPClient(options9...)
	ExpectEqual(t, true, mockHttpClient9 != nil)
	client.HTTPClient = mockHttpClient9
	res, err = CopyObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, bce.NewBceServiceError("InternalError", "We encountered an internal error. Please try again.",
		"52454655-5345-4420-4259-204e47494e58", 500), err)
	ExpectEqual(t, nil, res)

	// case8: all is ok
	args.ObjectExpires = 3
	args.TaggingDirective = METADATA_DIRECTIVE_COPY
	args.ObjectTagging = "key1=value2&key2=value2"
	args.CannedAcl = CANNED_ACL_PRIVATE
	args.UserMeta = map[string]string{"key1": "value1", "key2": "val2"}
	respBody8 := `{ "lastModified":"2016-05-12T09:14:32Z", "eTag":"67b92a7c2a9b9c1809a6ae3295dcc127" }`
	AttachMockHttpClientOk(t, client, &respBody8)
	res, err = CopyObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "2016-05-12T09:14:32Z", res.LastModified)
	ExpectEqual(t, "67b92a7c2a9b9c1809a6ae3295dcc127", res.ETag)
}
func TestGetObject(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	respBody := "this is a test-string for testing GetObject"
	resource := fmt.Sprintf("/%s/%s", bucket, object)

	// case1: all is ok
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.SetRespBody(respBody),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.BCE_VERSION_ID):                    "AKyQ9DRhhoY=",
			http.CanonicalHeaderKey(my_http.CACHE_CONTROL):                     "private",
			http.CanonicalHeaderKey(my_http.CONTENT_DISPOSITION):               "attachment; filename=\"download.txt\"",
			http.CanonicalHeaderKey(my_http.CONTENT_LENGTH):                    strconv.FormatInt(int64(len(respBody)), 10),
			http.CanonicalHeaderKey(my_http.CONTENT_TYPE):                      "text/plain",
			http.CanonicalHeaderKey(my_http.BCE_USER_METADATA_PREFIX) + "Key1": "Value1",
			http.CanonicalHeaderKey(my_http.BCE_USER_METADATA_PREFIX) + "Key2": "Value2",
			http.CanonicalHeaderKey(my_http.BCE_STORAGE_CLASS):                 STORAGE_CLASS_ARCHIVE,
		}),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	args1 := map[string]string{
		my_http.CONTENT_DISPOSITION: "CONTENT_DISPOSITION",
		my_http.CONTENT_TYPE:        "CONTENT_TYPE",
	}
	res, err := GetObject(client, bucket, object, nil, args1, int64(0), int64(len(respBody)))
	ExpectEqual(t, err, nil)
	ExpectEqual(t, res.VersionId, "AKyQ9DRhhoY=")
	ExpectEqual(t, res.CacheControl, "private")
	ExpectEqual(t, res.ContentDisposition, "attachment; filename=\"download.txt\"")
	ExpectEqual(t, res.ContentLength, int64(len(respBody)))
	ExpectEqual(t, res.ContentType, "text/plain")
	ExpectEqual(t, res.UserMeta["Key1"], "Value1")
	ExpectEqual(t, res.UserMeta["Key2"], "Value2")
	ExpectEqual(t, res.StorageClass, STORAGE_CLASS_ARCHIVE)
	buf := make([]byte, len(respBody))
	n, err := res.Body.Read(buf)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, len(respBody), n)
	ExpectEqual(t, respBody, string(buf))
	//case1.1: GetObjectWithArgs
	args11 := &GetObjectArgs{
		Params:            args1,
		Ranges:            []int64{0, int64(len(respBody))},
		IfMatch:           "string",
		IfNoneMatch:       "*",
		IfModifiedSince:   time.Now().Format(HTTPTimeFormat),
		IfUnModifiedSince: time.Now().Format(HTTPTimeFormat),
	}
	res, err = GetObjectWithArgs(client, bucket, object, nil, args11)
	ExpectEqual(t, err, nil)
	ExpectEqual(t, res.VersionId, "AKyQ9DRhhoY=")
	ExpectEqual(t, res.CacheControl, "private")
	ExpectEqual(t, res.ContentDisposition, "attachment; filename=\"download.txt\"")
	ExpectEqual(t, res.ContentLength, int64(len(respBody)))
	ExpectEqual(t, res.ContentType, "text/plain")
	ExpectEqual(t, res.UserMeta["Key1"], "Value1")
	ExpectEqual(t, res.UserMeta["Key2"], "Value2")
	ExpectEqual(t, res.StorageClass, STORAGE_CLASS_ARCHIVE)

	//case2: object is empty
	args2 := args1
	res, err = GetObject(client, bucket, "", nil, args2)
	ExpectEqual(t, err, fmt.Errorf("get object don't accept \"\" as a parameter"))
	ExpectEqual(t, res, nil)
	//case2.1: GetObjectWithArgs
	args21 := args11
	res, err = GetObjectWithArgs(client, bucket, "", nil, args21)
	ExpectEqual(t, err, fmt.Errorf("get object don't accept \"\" as a parameter"))
	ExpectEqual(t, res, nil)
	//case2.2: GetObjectWithArgs with error option
	res, err = GetObjectWithArgs(client, bucket, object, nil, args21, ErrorOption)
	ExpectEqual(t, err, optionError)
	ExpectEqual(t, res, nil)

	//case3: send request error
	args3 := args1
	err3 := fmt.Errorf("IO Error")
	options3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err3),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = GetObject(client, bucket, object, nil, args3, 0)
	CheckMockHttpClientError(t, client.Config.Endpoint, resource, "", "Get", err3, err)
	ExpectEqual(t, nil, res)
	//case3.1: GetObjectWithArgs
	args31 := args11
	args31.Ranges = []int64{0}
	res, err = GetObjectWithArgs(client, bucket, object, nil, args31)
	CheckMockHttpClientError(t, client.Config.Endpoint, resource, "", "Get", err3, err)
	ExpectEqual(t, nil, res)

	//case4: resp is fail
	args4 := args1
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	res, err = GetObject(client, bucket, object, nil, args4, 0)
	ExpectEqual(t, bceServiceErro404.Error(), err.Error())
	ExpectEqual(t, nil, res)
	//case41: GetObjectWithArgs
	args41 := args11
	res, err = GetObjectWithArgs(client, bucket, object, nil, args41)
	ExpectEqual(t, bceServiceErro404.Error(), err.Error())
	ExpectEqual(t, nil, res)
}

func TestGetObjectMeta(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	resource := "/" + bucket + "/" + object

	//case1: all is ok
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.BCE_VERSION_ID):                    "AKyQ9DRhhoY=",
			http.CanonicalHeaderKey(my_http.CACHE_CONTROL):                     "private",
			http.CanonicalHeaderKey(my_http.CONTENT_DISPOSITION):               "attachment; filename=\"download.txt\"",
			http.CanonicalHeaderKey(my_http.CONTENT_LENGTH):                    "1234567",
			http.CanonicalHeaderKey(my_http.CONTENT_TYPE):                      "application/octet-stream",
			http.CanonicalHeaderKey(my_http.BCE_USER_METADATA_PREFIX) + "Key1": "Value1",
			http.CanonicalHeaderKey(my_http.BCE_USER_METADATA_PREFIX) + "Key2": "Value2",
			http.CanonicalHeaderKey(my_http.BCE_STORAGE_CLASS):                 STORAGE_CLASS_ARCHIVE,
			http.CanonicalHeaderKey(my_http.CONTENT_MD5):                       "Zh+ACfqOVqnQ6UoKZEOX1w==",
			http.CanonicalHeaderKey(my_http.LAST_MODIFIED):                     "Wed, 17 Dec 2025 06:25:34 GMT",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32):                 "1922069637",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC64ECMA):             "12759301844125077625",
		}),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	res, err := GetObjectMeta(client, bucket, object, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, res.VersionId, "AKyQ9DRhhoY=")
	ExpectEqual(t, res.CacheControl, "private")
	ExpectEqual(t, res.ContentDisposition, "attachment; filename=\"download.txt\"")
	ExpectEqual(t, res.ContentLength, 1234567)
	ExpectEqual(t, res.ContentType, "application/octet-stream")
	ExpectEqual(t, res.UserMeta["Key1"], "Value1")
	ExpectEqual(t, res.UserMeta["Key2"], "Value2")
	ExpectEqual(t, res.StorageClass, STORAGE_CLASS_ARCHIVE)
	ExpectEqual(t, res.ContentMD5, "Zh+ACfqOVqnQ6UoKZEOX1w==")
	ExpectEqual(t, res.LastModified, "Wed, 17 Dec 2025 06:25:34 GMT")
	ExpectEqual(t, res.ContentCrc32, "1922069637")
	ExpectEqual(t, res.ContentCrc64ECMA, "12759301844125077625")
	//case2: handle option error
	res, err = GetObjectMeta(client, bucket, object, nil, ErrorOption)
	ExpectEqual(t, err, optionError)
	ExpectEqual(t, res, nil)
	//case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	res, err = GetObjectMeta(client, bucket, object, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, resource, "", "Head", err3, err)
	ExpectEqual(t, nil, res)
	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	res, err = GetObjectMeta(client, bucket, object, nil)
	ExpectEqual(t, bceServiceErro404.Error(), err.Error())
	ExpectEqual(t, nil, res)
}

func TestSelectObject(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	args := &SelectObjectArgs{}
	reqBody := `{
		"selectRequest": {
			"expression": "c2VsZWN0IGNvdW50KCopIGZyb20gbxkl2JqZWN0IHdoZXJlIF80ID4gNDU=",
			"expressionType": "SQL",
			"inputSerialization": {
				"compressionType": "NONE",
				"json": { "type": "DOCUMENT" }
			},
			"outputSerialization": {
				"json": { "recordDelimiter": "Cg==" }
			},
			"requestProgress": { "enabled": false }
		}
	}`
	ExpectEqual(t, nil, json.Unmarshal([]byte(reqBody), args))
	// case1: handle options error
	res, err := SelectObject(client, bucket, object, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case6: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	res, err = SelectObject(client, bucket, object, args, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	//case7: resp is fail
	options7 := util.RoundTripperOpts404
	mockHttpClient7 := util.NewMockHTTPClient(options7...)
	ExpectEqual(t, true, mockHttpClient7 != nil)
	client.HTTPClient = mockHttpClient7
	res, err = SelectObject(client, bucket, object, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)

	//case8: all is ok
	respBody8 := "response body"
	AttachMockHttpClientOk(t, client, &respBody8)
	res, err = SelectObject(client, bucket, object, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, true, res != nil)
}

func TestFetchObject(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	source := "fetch-resource"
	args := &FetchObjectArgs{}

	//case1: empty source
	res, err := FetchObject(client, bucket, object, "", args, nil)
	ExpectEqual(t, bce.NewBceClientError("invalid fetch source value: "), err)
	ExpectEqual(t, nil, res)

	//case2: invalid fetch mode
	args.FetchMode = "fetch-mode"
	res, err = FetchObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, bce.NewBceClientError("invalid fetch mode value: "+args.FetchMode), err)
	ExpectEqual(t, nil, res)
	args.FetchMode = FETCH_MODE_ASYNC

	//case3: invalid storage class
	args.StorageClass = "storage-class"
	res, err = FetchObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, bce.NewBceClientError("invalid storage class value: "+args.StorageClass), err)
	ExpectEqual(t, nil, res)
	args.StorageClass = STORAGE_CLASS_ARCHIVE

	// case4: handle options error
	res, err = FetchObject(client, bucket, object, source, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case6: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	res, err = FetchObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	//case7: resp is fail
	options7 := util.RoundTripperOpts404
	mockHttpClient7 := util.NewMockHTTPClient(options7...)
	ExpectEqual(t, true, mockHttpClient7 != nil)
	client.HTTPClient = mockHttpClient7
	res, err = FetchObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)

	//case8: parse json fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = FetchObject(client, bucket, object, source, args, nil)
	result4 := &ListMultipartUploadsResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result4), err)
	ExpectEqual(t, nil, res)

	// case9: all is ok
	args.ContentEncoding = "content_encoding"
	args.ObjectExpires = 3
	args.FetchCallBackAddress = "callback_address"
	respBody9 := `{
		"code": "success",
		"message": "success",
		"requestId": "4db2b34d-654d-4d8a-b49b-3049ca786409"
	}`
	AttachMockHttpClientOk(t, client, &respBody9)
	res, err = FetchObject(client, bucket, object, source, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "success", res.Code)
	ExpectEqual(t, "4db2b34d-654d-4d8a-b49b-3049ca786409", res.RequestId)
}

func TestAppendObject(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	data := "data-string"
	content, err := bce.NewBodyFromString(data)
	ExpectEqual(t, nil, err)
	args := &AppendObjectArgs{}

	// case1: content is nil
	res, err := AppendObject(client, bucket, object, nil, args, nil)
	ExpectEqual(t, bce.NewBceClientError("AppendObject body should not be emtpy"), err)
	ExpectEqual(t, nil, res)

	// case2: offset < 0
	args.Offset = -1
	res, err = AppendObject(client, bucket, object, content, args, nil)
	ExpectEqual(t, bce.NewBceClientError(fmt.Sprintf("invalid append offset value: %d", args.Offset)), err)
	ExpectEqual(t, nil, res)
	args.Offset = 10

	// case3: invalid storage class
	args.StorageClass = "storage-class"
	res, err = AppendObject(client, bucket, object, content, args, nil)
	ExpectEqual(t, bce.NewBceClientError("invalid storage class value: "+args.StorageClass), err)
	ExpectEqual(t, nil, res)
	args.StorageClass = STORAGE_CLASS_ARCHIVE

	// case4: invalid traffic limit
	args.TrafficLimit = TRAFFIC_LIMIT_MAX + 1
	res, err = AppendObject(client, bucket, object, content, args, nil)
	ExpectEqual(t, trafficLimitInvalidError(args.TrafficLimit), err)
	ExpectEqual(t, nil, res)
	args.TrafficLimit = TRAFFIC_LIMIT_MAX - 1

	// case5: handle options error
	res, err = AppendObject(client, bucket, object, content, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case6: send request error
	err6 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err6)
	content6, err := bce.NewBodyFromString(data)
	ExpectEqual(t, nil, err)
	res, err = AppendObject(client, bucket, object, content6, args, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	//case7: resp is fail
	options7 := util.RoundTripperOpts404
	mockHttpClient7 := util.NewMockHTTPClient(options7...)
	ExpectEqual(t, true, mockHttpClient7 != nil)
	client.HTTPClient = mockHttpClient7
	content7, err := bce.NewBodyFromString(data)
	ExpectEqual(t, nil, err)
	res, err = AppendObject(client, bucket, object, content7, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)

	// case8: all is ok
	args.ObjectExpires = 3
	options8 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			my_http.BCE_NEXT_APPEND_OFFSET: "100",
			my_http.ETAG:                   "67b92a7c2a9b9c1809a6ae3295dcc127",
		}),
	}
	mockHttpClient8 := util.NewMockHTTPClient(options8...)
	ExpectEqual(t, true, mockHttpClient8 != nil)
	client.HTTPClient = mockHttpClient8
	content8, err := bce.NewBodyFromString(data)
	ExpectEqual(t, nil, err)
	res, err = AppendObject(client, bucket, object, content8, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "67b92a7c2a9b9c1809a6ae3295dcc127", res.ETag)
	ExpectEqual(t, 100, res.NextAppendOffset)
}

func TestDeleteObject(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	versionId := "version_id"

	// case1: handle options error
	err = DeleteObject(client, bucket, object, versionId, nil, ErrorOption)
	ExpectEqual(t, optionError, err)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	err = DeleteObject(client, bucket, object, versionId, nil)
	ExpectEqual(t, true, err != nil)

	//case3: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	err = DeleteObject(client, bucket, object, versionId, nil)
	ExpectEqual(t, bceServiceErro404, err)

	// case4: all is ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteObject(client, bucket, object, versionId, nil)
	ExpectEqual(t, nil, err)
}

func TestDeleteMultipleObjects(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	objectList := `  {
		"objects": [
			{ "key": "my-object1" },
			{ "key": "my-object2" }
		]
	}`

	// case1: object list stream is nil
	res, err := DeleteMultipleObjects(client, bucket, nil, nil)
	ExpectEqual(t, bce.NewBceClientError("DeleteMultipleObjects body should not be emtpy"), err)
	ExpectEqual(t, nil, res)

	// case2: handle options error
	objectListStream2, err := bce.NewBodyFromStringV2(objectList, false)
	ExpectEqual(t, nil, err)
	res, err = DeleteMultipleObjects(client, bucket, objectListStream2, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	objectListStream3, err := bce.NewBodyFromStringV2(objectList, false)
	ExpectEqual(t, nil, err)
	res, err = DeleteMultipleObjects(client, bucket, objectListStream3, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	objectListStream4, err := bce.NewBodyFromStringV2(objectList, false)
	ExpectEqual(t, nil, err)
	res, err = DeleteMultipleObjects(client, bucket, objectListStream4, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)

	// case5: response content length is 0
	options5 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			my_http.CONTENT_LENGTH: "0",
		}),
	}
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
	objectListStream5, err := bce.NewBodyFromStringV2(objectList, false)
	ExpectEqual(t, nil, err)
	res, err = DeleteMultipleObjects(client, bucket, objectListStream5, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 0, len(res.Errors))

	// case6: parse json fail
	AttachMockHttpClientJsonBodyError(t, client)
	objectListStream6, err := bce.NewBodyFromStringV2(objectList, false)
	ExpectEqual(t, nil, err)
	res, err = DeleteMultipleObjects(client, bucket, objectListStream6, nil)
	jsonBody := &DeleteMultipleObjectsResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(jsonBody), err)
	ExpectEqual(t, nil, res)

	// case7: ok
	respBody7 := `  {
		"errors": [
			{
				"key": "my-object1",
				"code": "NoSuchKey",
				"message": "The specified key does not exist."
			},
			{
				"key": "my-object2",
				"code": "InvalidArgument",
				"message": "Invalid Argument."
			}
		]
	}`
	AttachMockHttpClientOk(t, client, &respBody7)
	objectListStream7, err := bce.NewBodyFromStringV2(objectList, false)
	ExpectEqual(t, nil, err)
	res, err = DeleteMultipleObjects(client, bucket, objectListStream7, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.Errors))
}

func TestGeneratePresignedUrlInternal(t *testing.T) {
	conf := &bce.BceClientConfiguration{
		Endpoint: "localhost:8080",
		Region:   "bj",
		Credentials: &auth.BceCredentials{
			AccessKeyId:     "ak",
			SecretAccessKey: "sk",
			SessionToken:    "sts",
		},
		SignOption: &auth.SignOptions{
			Timestamp: time.Now().Unix(),
		},
	}
	bucket := "test-bucket"
	object := "test-object"
	expire := 1800
	headers := map[string]string{"headr1": "value1", "header2": "value2"}
	params := map[string]string{"param1": "value1", "param2": "value2"}
	signer := &auth.BceV1Signer{}

	// case1: object is empty
	res := GeneratePresignedUrlInternal(conf, signer, bucket, "", expire, "", headers, params, true)
	ExpectEqual(t, "", res)

	// case2: path style, cname enable
	conf.CnameEnabled = true
	res = GeneratePresignedUrlInternal(conf, signer, bucket, "", expire, object, headers, params, true)
	log.Warnf("pre-signed url: %s", res)

	// case3: virtual host, endpoint is not ip
	res = GeneratePresignedUrlInternal(conf, signer, bucket, "", expire, object, headers, params, false)
	log.Warnf("pre-signed url: %s", res)

	// case3: virtual host, endpoint is ip
	conf.Endpoint = "192.168.1.1:8080"
	res = GeneratePresignedUrlInternal(conf, signer, bucket, "", expire, object, headers, params, false)
	log.Warnf("pre-signed url: %s", res)

	// case:
	res = GeneratePresignedUrl(conf, signer, bucket, "", expire, object, headers, params)
	log.Warnf("pre-signed url: %s", res)
	res = GeneratePresignedUrlPathStyle(conf, signer, bucket, "", expire, object, headers, params)
	log.Warnf("pre-signed url: %s", res)
}

func TestPutObjectAcl(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	cannedAcl := CANNED_ACL_PRIVATE
	grantRead := []string{"id1", "id2"}
	grantFullControl := []string{"id3", "id4"}
	aclString := ` {
		"accessControlList":[
			{
				"grantee":[{ "id":"e13b12d0131b4c8bae959df4969387b8" }],
				"permission":["READ"]
			}
		]
	}`
	aclBody, err := bce.NewBodyFromStringV2(aclString, false)
	ExpectEqual(t, nil, err)

	// case1:
	err = PutObjectAcl(client, bucket, object, cannedAcl, grantRead, grantFullControl, aclBody, nil)
	ExpectEqual(t, bce.NewBceClientError("BOS only support one acl setting method at the same time"), err)

	// case2: handle options error
	err = PutObjectAcl(client, bucket, object, cannedAcl, nil, nil, nil, nil, ErrorOption)
	ExpectEqual(t, optionError, err)

	// case3: send request error
	err3 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err3)
	err = PutObjectAcl(client, bucket, object, cannedAcl, nil, nil, nil, nil, ErrorOption)
	ExpectEqual(t, true, err != nil)

	//case4: resp is fail
	options4 := util.RoundTripperOpts404
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	err = PutObjectAcl(client, bucket, object, cannedAcl, nil, nil, nil, nil)
	ExpectEqual(t, bceServiceErro404, err)

	// case5: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutObjectAcl(client, bucket, object, cannedAcl, nil, nil, nil, nil)
	ExpectEqual(t, nil, err)
}

func TestGetObjectAcl(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"

	// case1: handle option error
	res, err := GetObjectAcl(client, bucket, object, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	res, err = GetObjectAcl(client, bucket, object, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	//case3: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = GetObjectAcl(client, bucket, object, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)

	// case4: parse json fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetObjectAcl(client, bucket, object, nil)
	result4 := &GetObjectAclResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result4), err)
	ExpectEqual(t, nil, res)

	// case5: ok
	respBody5 := `{
		"accessControlList":[
			{
				"grantee":[{ "id":"e13b12d0131b4c8bae959df4969387b8" }],
				"permission":["FULL_CONTROL"]
			},
			{
				"grantee":[{ "id":"8c47a952db4444c5a097b41be3f24c94" }],
				"permission":["READ"]
			}
		]
	}`
	AttachMockHttpClientOk(t, client, &respBody5)
	res, err = GetObjectAcl(client, bucket, object, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.AccessControlList))
}

func TestDeleteObjectAcl(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"

	// case1: handle option error
	err = DeleteObjectAcl(client, bucket, object, nil, ErrorOption)
	ExpectEqual(t, optionError, err)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	err = DeleteObjectAcl(client, bucket, object, nil)
	ExpectEqual(t, true, err != nil)

	//case4: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	err = DeleteObjectAcl(client, bucket, object, nil)
	ExpectEqual(t, bceServiceErro404, err)

	//case5: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteObjectAcl(client, bucket, object, nil)
	ExpectEqual(t, nil, err)
}

func TestRestoreObject(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	args := ArchiveRestoreArgs{
		RestoreTier: RESTORE_TIER_EXPEDITED,
		RestoreDays: 3,
	}

	// case1: handle option error
	err = RestoreObject(client, bucket, object, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	err = RestoreObject(client, bucket, object, args, nil)
	ExpectEqual(t, true, err != nil)

	//case3: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	err = RestoreObject(client, bucket, object, args, nil)
	ExpectEqual(t, bceServiceErro404, err)

	//case4: ok
	AttachMockHttpClientOk(t, client, nil)
	err = RestoreObject(client, bucket, object, args, nil)
	ExpectEqual(t, nil, err)
}

func TestPutObjectSymlink(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	symlinkKey := "symlink-key"
	symlinkArgs := &PutSymlinkArgs{}

	// case1 : invalid ForbidOverwrite
	symlinkArgs.ForbidOverwrite = "forbiden"
	err = PutObjectSymlink(client, bucket, object, symlinkKey, symlinkArgs, nil)
	ExpectEqual(t, bce.NewBceClientError("invalid forbid overwrite val,"+symlinkArgs.ForbidOverwrite), err)
	symlinkArgs.ForbidOverwrite = "true"

	//case2: invalid storage class
	symlinkArgs.StorageClass = "storage-class"
	err = PutObjectSymlink(client, bucket, object, symlinkKey, symlinkArgs, nil)
	ExpectEqual(t, bce.NewBceClientError("invalid storage class val,"+symlinkArgs.StorageClass), err)
	symlinkArgs.StorageClass = STORAGE_CLASS_ARCHIVE
	err = PutObjectSymlink(client, bucket, object, symlinkKey, symlinkArgs, nil)
	ExpectEqual(t, bce.NewBceClientError("archive storage class not support"), err)
	symlinkArgs.StorageClass = STORAGE_CLASS_COLD

	symlinkArgs.UserMeta = map[string]string{"key1": "value1", "key2": "value2"}
	symlinkArgs.SymlinkBucket = "symlink-bucket"
	// case3: handle option error
	err = PutObjectSymlink(client, bucket, object, symlinkKey, symlinkArgs, nil, ErrorOption)
	ExpectEqual(t, optionError, err)

	// case4: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	err = PutObjectSymlink(client, bucket, object, symlinkKey, symlinkArgs, nil)
	ExpectEqual(t, true, err != nil)

	//case5: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	err = PutObjectSymlink(client, bucket, object, symlinkKey, symlinkArgs, nil)
	ExpectEqual(t, bceServiceErro404, err)

	//case6: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutObjectSymlink(client, bucket, object, symlinkKey, symlinkArgs, nil)
	ExpectEqual(t, nil, err)
}

func TestGetObjectSymlink(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	symlinkKey := "test-object"

	// case1: handle option error
	res, err := GetObjectSymlink(client, bucket, symlinkKey, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, "", res)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	res, err = GetObjectSymlink(client, bucket, symlinkKey, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, "", res)

	//case3: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = GetObjectSymlink(client, bucket, symlinkKey, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, "", res)

	// case4: ok
	options4 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			my_http.BCE_SYMLINK_TARGET: "target-object",
		}),
	}
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	res, err = GetObjectSymlink(client, bucket, symlinkKey, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "target-object", res)

	// case5: ok
	options5 := append(options4, util.AddHeaders(map[string]string{
		my_http.BCE_SYMLINK_BUCKET: "target-bucket",
	}))
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
	res, err = GetObjectSymlink(client, bucket, symlinkKey, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, BOS_CONFIG_PREFIX+"target-bucket/target-object", res)
}

func TestPutObjectTag(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	putObjectTagArgs := &PutObjectTagArgs{
		ObjectTags: []ObjectTags{
			{
				TagInfo: []ObjectTag{
					{Key: "key1", Value: "value1"},
					{Key: "key2", Value: "value2"},
				},
			},
		},
	}

	// case1: handle option error
	err = PutObjectTag(client, bucket, object, putObjectTagArgs, nil, ErrorOption)
	ExpectEqual(t, optionError, err)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	err = PutObjectTag(client, bucket, object, putObjectTagArgs, nil)
	ExpectEqual(t, true, err != nil)

	//case3: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	err = PutObjectTag(client, bucket, object, putObjectTagArgs, nil)
	ExpectEqual(t, bceServiceErro404, err)

	// case4: ok
	AttachMockHttpClientOk(t, client, nil)
	err = PutObjectTag(client, bucket, object, putObjectTagArgs, nil)
	ExpectEqual(t, nil, err)
}

func TestGetObjectTag(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"

	// case1: handle option error
	res, err := GetObjectTag(client, bucket, object, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	res, err = GetObjectTag(client, bucket, object, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	//case3: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = GetObjectTag(client, bucket, object, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)

	// case5: parser json error
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = GetObjectTag(client, bucket, object, nil)
	var data map[string]interface{}
	ExpectEqual(t, json.Unmarshal([]byte(errorJsonBody), &data), err)
	ExpectEqual(t, nil, res)

	// case6: ok
	respBody6 := `{
		"tagSet": [{
			"tagInfo": { 
				"key9": "value9", "key8": "value8", "key10": "value10", "key3": "value3", "key2": "value2",
				"key1": "value1", "key0": "value0", "key6": "value6", "key5": "value5", "key4": "value4"
			}
		}]
	}`
	AttachMockHttpClientOk(t, client, &respBody6)
	res, err = GetObjectTag(client, bucket, object, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 10, len(res))
}

func TestDeleteObjectTag(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"

	// case1: handle option error
	err = DeleteObjectTag(client, bucket, object, nil, ErrorOption)
	ExpectEqual(t, optionError, err)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	err = DeleteObjectTag(client, bucket, object, nil)
	ExpectEqual(t, true, err != nil)

	//case4: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	err = DeleteObjectTag(client, bucket, object, nil)
	ExpectEqual(t, bceServiceErro404, err)

	//case5: ok
	AttachMockHttpClientOk(t, client, nil)
	err = DeleteObjectTag(client, bucket, object, nil)
	ExpectEqual(t, nil, err)
}

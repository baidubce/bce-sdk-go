package api

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	my_http "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
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
	options3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err3),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
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

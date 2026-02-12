package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	my_http "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	optionError       *bce.BceClientError  = bce.NewBceClientError("Handle options occur error: error option")
	optionError1      *bce.BceClientError  = bce.NewBceClientError("Handle options error: error option")
	postOptionError   *bce.BceClientError  = bce.NewBceClientError("Handle post options error: error option")
	bceServiceErro404 *bce.BceServiceError = bce.NewBceServiceError("NOTFound", "404 NOT Found", "", http.StatusNotFound)
)

func init() {
	log.SetLogLevel(log.WARN)
	log.SetLogHandler(log.STDOUT)
}

func trafficLimitInvalidError(trafficLimit int64) *bce.BceClientError {
	return bce.NewBceClientError(fmt.Sprintf("TrafficLimit must between %d ~ %d, current value:%d",
		TRAFFIC_LIMIT_MIN, TRAFFIC_LIMIT_MAX, trafficLimit))
}

func httpClientDoError(retry int, uErr *url.Error) *bce.BceClientError {
	return &bce.BceClientError{
		Message: fmt.Sprintf("execute http request failed! Retried %d times, error: %v", retry, uErr),
	}
}

func crc32cMisMatchError(cliCrc32c, svrCrc32c string) *bce.BceClientError {
	return &bce.BceClientError{
		Message: fmt.Sprintf("crc32 is inconsistence, client: %s, server: %s", cliCrc32c, svrCrc32c),
	}
}

func NewMockBosClient() (*bce.BceClient, error) {
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	config := &bce.BceClientConfiguration{
		Endpoint:                  "192.168.1.1",
		Region:                    "bce.DEFAULT_REGION",
		UserAgent:                 "bce.DEFAULT_USER_AGENT",
		Credentials:               &auth.BceCredentials{AccessKeyId: "ak", SecretAccessKey: "sk"},
		SignOption:                defaultSignOptions,
		Retry:                     bce.NewBackOffRetryPolicy(3, 100, 20000),
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
	}
	return bce.NewBceClientWithExclusiveHTTPClient(config, &auth.BceV1Signer{})
}

func TestInitiateMultipartUpload(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameter
	bucket := "test-bucket"
	object := "test-object"
	respBody := `{ "bucket": "BucketName", "key":"ObjectName", "uploadId": "a44cc9bab11cbd156984767aad637851" }`
	args := &InitiateMultipartUploadArgs{}
	AttachMockHttpClientOk(t, client, &respBody)

	// case1: invalid storage class
	args.StorageClass = "storage-class"
	res, err := InitiateMultipartUpload(client, bucket, object, "", args, nil)
	expectErr1 := bce.NewBceClientError("invalid storage class value: " + args.StorageClass)
	ExpectEqual(t, expectErr1, err)
	ExpectEqual(t, nil, res)
	args.StorageClass = STORAGE_CLASS_STANDARD

	// case2: invalid acl
	args.CannedAcl = "canned-acl"
	res, err = InitiateMultipartUpload(client, bucket, object, "", args, nil)
	expectErr2 := bce.NewBceClientError("invalid canned acl value: " + args.CannedAcl)
	ExpectEqual(t, expectErr2, err)
	ExpectEqual(t, nil, res)
	args.CannedAcl = CANNED_ACL_PRIVATE

	// case3: handle options error
	res, err = InitiateMultipartUpload(client, bucket, object, "", args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case4: send request error
	err4 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err4)
	res, err = InitiateMultipartUpload(client, bucket, object, "", args, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, "/"+bucket+"/"+object, "uploads", "Post", err4, err)
	ExpectEqual(t, nil, res)

	//case5: resp is fail
	options5 := util.RoundTripperOpts404
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
	res, err = InitiateMultipartUpload(client, bucket, object, "", args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)

	//case5: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = InitiateMultipartUpload(client, bucket, object, "", args, nil)
	result5 := &InitiateMultipartUploadResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)

	// case6: all is ok
	args.ObjectExpires = 3
	args.CopySource = "copy-source"
	args.GrantRead = []string{"id1", "id2"}
	args.GrantFullControl = []string{"id3", "id3"}
	args.ObjectTagging = "key1=value1&key2=value2"
	AttachMockHttpClientOk(t, client, &respBody)
	res, err = InitiateMultipartUpload(client, bucket, object, "", args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "a44cc9bab11cbd156984767aad637851", res.UploadId)
	ExpectEqual(t, "ObjectName", res.Key)
	ExpectEqual(t, "BucketName", res.Bucket)
}

func TestUploadPart(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameter
	bucket := "test-bucket"
	object := "test-object"
	uploadId := "a44cc9bab11bdc157676984aad851637"
	reqBody := "this is a body string for testing upload part."
	partNumber := 1
	content, err := bce.NewBodyFromStringV2(reqBody, true)
	ExpectEqual(t, nil, err)
	crc64Hash := crc64.New(crc64.MakeTable(crc64.ECMA))
	n, err := crc64Hash.Write([]byte(reqBody))
	ExpectEqual(t, n, len(reqBody))
	ExpectEqual(t, nil, err)
	crc64Str := strconv.FormatUint(crc64Hash.Sum64(), 10)
	// case1: mock http client, crc64 match
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str,
			my_http.ETAG: "9b2cf535f27731c974343645a3985328",
		}),
	}
	mockHttpClient := util.NewMockHTTPClient(options...)
	ExpectEqual(t, true, mockHttpClient != nil)
	client.HTTPClient = mockHttpClient
	args := &UploadPartArgs{
		ContentMD5:       content.ContentMD5(),
		TrafficLimit:     TRAFFIC_LIMIT_MIN + 1,
		ContentCrc32:     string(crc32.NewIEEE().Sum([]byte(reqBody))),
		ContentCrc64ECMA: crc64Str,
	}
	bosContext := &BosContext{
		PathStyleEnable: false,
		Ctx:             context.Background(),
		ApiVersion:      API_VERSION_V1,
	}
	etag, err := UploadPart(client, bucket, object, uploadId, partNumber, content, args, bosContext)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "9b2cf535f27731c974343645a3985328", etag)
	//case2: mock http client, crc64 mismatch
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): "crc64Str",
			my_http.ETAG: "9b2cf535f27731c974343645a3985328",
		}),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	bosContext.ApiVersion = API_VERSION_V2
	content1, err := bce.NewBodyFromStringV2(reqBody, true)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, len(reqBody), content1.Size())
	bosContext1 := &BosContext{
		PathStyleEnable: false,
		Ctx:             context.Background(),
		ApiVersion:      API_VERSION_V1,
	}
	etag, err = UploadPart(client, bucket, object, uploadId, partNumber, content1, args, bosContext1)
	ExpectEqual(t, fmt.Sprintf("crc64 is not consistent, client: %s, server: crc64Str", crc64Str), err.Error())
	ExpectEqual(t, "", etag)

	//case3: bos context is nil
	client.HTTPClient = mockHttpClient
	args3 := args
	content3, err := bce.NewBodyFromStringV2(reqBody, true)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, len(reqBody), content3.Size())
	etag, err = UploadPart(client, bucket, object, uploadId, partNumber, content3, args3, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "9b2cf535f27731c974343645a3985328", etag)

	//case4 : traffic limit is invalid
	args4 := &UploadPartArgs{
		TrafficLimit: TRAFFIC_LIMIT_MIN - 1,
	}
	etag, err = UploadPart(client, bucket, object, uploadId, partNumber, content, args4, nil)
	ExpectEqual(t, trafficLimitInvalidError(args4.TrafficLimit), err)
	ExpectEqual(t, "", etag)

	//case5: handle options error
	args5 := args
	etag, err = UploadPart(client, bucket, object, uploadId, partNumber, content, args5, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, "", etag)

	//case6: send request error
	args6 := args
	err6 := fmt.Errorf("IO Error")
	options6 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err6),
	}
	mockHttpClient6 := util.NewMockHTTPClient(options6...)
	ExpectEqual(t, true, mockHttpClient6 != nil)
	client.HTTPClient = mockHttpClient6
	etag, err = UploadPart(client, bucket, object, uploadId, partNumber, content, args6, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, "", etag)

	// case7: resp is fail
	args7 := args
	options7 := util.RoundTripperOpts404
	mockHttpClient7 := util.NewMockHTTPClient(options7...)
	ExpectEqual(t, true, mockHttpClient7 != nil)
	client.HTTPClient = mockHttpClient7
	content7, err := bce.NewBodyFromStringV2(reqBody, true)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, len(reqBody), content7.Size())
	etag, err = UploadPart(client, bucket, object, uploadId, partNumber, content7, args7, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, "", etag)
}

func TestUploadPartFromBytes(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	uploadId := "a44cc9bab11bdc157676984aad851637"
	reqBody := "this is a body string for testing upload part."
	partNumber := 1
	// calculate crc64 of reqBody string
	crc64Hash := crc64.New(crc64.MakeTable(crc64.ECMA))
	n, err := crc64Hash.Write([]byte(reqBody))
	ExpectEqual(t, n, len(reqBody))
	ExpectEqual(t, nil, err)
	crc64Str := strconv.FormatUint(crc64Hash.Sum64(), 10)
	//mock http client, crc64 match
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC64ECMA): crc64Str,
			my_http.ETAG: "9b2cf535f27731c974343645a3985328",
		}),
	}
	mockHttpClient := util.NewMockHTTPClient(options...)
	ExpectEqual(t, true, mockHttpClient != nil)
	client.HTTPClient = mockHttpClient
	//construct upload part args
	args := &UploadPartArgs{
		TrafficLimit:     TRAFFIC_LIMIT_MIN + 1,
		ContentCrc32:     string(crc32.NewIEEE().Sum([]byte(reqBody))),
		ContentCrc64ECMA: crc64Str,
		ContentMD5:       "string",
	}
	bosContext := &BosContext{
		PathStyleEnable: false,
		Ctx:             context.Background(),
		ApiVersion:      API_VERSION_V1,
	}
	// case1: all is ok
	etag, err := UploadPartFromBytes(client, bucket, object, uploadId, partNumber, []byte(reqBody), args, bosContext)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "9b2cf535f27731c974343645a3985328", etag)

	// case2: content is nil
	etag, err = UploadPartFromBytes(client, bucket, object, uploadId, partNumber, nil, args, bosContext)
	ExpectEqual(t, bce.NewBceClientError("upload part content should not be empty"), err)
	ExpectEqual(t, "", etag)

	// case3: trafficLimit is invalid
	args3 := &UploadPartArgs{
		TrafficLimit: TRAFFIC_LIMIT_MIN - 1,
	}
	etag, err = UploadPartFromBytes(client, bucket, object, uploadId, partNumber, []byte(reqBody), args3, bosContext)
	ExpectEqual(t, trafficLimitInvalidError(args3.TrafficLimit), err)
	ExpectEqual(t, "", etag)

	//case4: handle options error
	args4 := args
	etag, err = UploadPartFromBytes(client, bucket, object, uploadId, partNumber, []byte(reqBody), args4, bosContext, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, "", etag)

	//case5: send request error
	args5 := args
	err5 := fmt.Errorf("IO Error")
	options5 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err5),
	}
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
	etag, err = UploadPartFromBytes(client, bucket, object, uploadId, partNumber, []byte(reqBody), args5, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, "", etag)

	// case6: resp is fail
	args6 := args
	options6 := util.RoundTripperOpts404
	mockHttpClient6 := util.NewMockHTTPClient(options6...)
	ExpectEqual(t, true, mockHttpClient6 != nil)
	client.HTTPClient = mockHttpClient6
	content7, err := bce.NewBodyFromStringV2(reqBody, true)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, len(reqBody), content7.Size())
	etag, err = UploadPartFromBytes(client, bucket, object, uploadId, partNumber, []byte(reqBody), args6, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, "", etag)
}

func TestUploadPartCopy(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	uploadId := "a44cc9bab11bdc157676984aad851637"
	source := "copy-source"
	partNumber := 1
	args := &UploadPartCopyArgs{}

	// case1: source is empty
	res, err := UploadPartCopy(client, bucket, object, "", uploadId, partNumber, args, nil)
	ExpectEqual(t, bce.NewBceClientError("upload part copy source should not be empty"), err)
	ExpectEqual(t, nil, res)

	// case2: invalid TRAFFIC_LIMIT
	args.TrafficLimit = TRAFFIC_LIMIT_MAX + 1
	res, err = UploadPartCopy(client, bucket, object, source, uploadId, partNumber, args, nil)
	ExpectEqual(t, trafficLimitInvalidError(args.TrafficLimit), err)
	ExpectEqual(t, nil, res)
	args.TrafficLimit = TRAFFIC_LIMIT_MAX - 1

	// case3: handle options error
	res, err = UploadPartCopy(client, bucket, object, source, uploadId, partNumber, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	//case4: send request error
	err4 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err4)
	res, err = UploadPartCopy(client, bucket, object, source, uploadId, partNumber, args, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	// case5: resp is fail
	options5 := util.RoundTripperOpts404
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
	res, err = UploadPartCopy(client, bucket, object, source, uploadId, partNumber, args, nil)
	ExpectEqual(t, bceServiceErro404, err)
	ExpectEqual(t, nil, res)

	//case6: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = UploadPartCopy(client, bucket, object, source, uploadId, partNumber, args, nil)
	result5 := &CopyObjectResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)

	// case7: resp body is fail content
	respBody7 := `{
		"code":"InternalError",
		"message":"We encountered an internal error. Please try again.",
		"requestId":"52454655-5345-4420-4259-204e47494e58"
	}`
	options7 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			my_http.BCE_REQUEST_ID: "9b2cf535f27731c974343645a3985328",
		}),
		util.SetRespBody(respBody7),
	}
	mockHttpClient7 := util.NewMockHTTPClient(options7...)
	ExpectEqual(t, true, mockHttpClient7 != nil)
	client.HTTPClient = mockHttpClient7
	res, err = UploadPartCopy(client, bucket, object, source, uploadId, partNumber, args, nil)
	ExpectEqual(t, bce.NewBceServiceError("InternalError", "We encountered an internal error. Please try again.",
		"9b2cf535f27731c974343645a3985328", 500), err)
	ExpectEqual(t, nil, res)

	// case8: all is ok
	respBody8 := `{ "lastModified":"2016-05-12T09:14:32Z", "eTag":"67b92a7c2a9b9c1809a6ae3295dcc127" }`
	AttachMockHttpClientOk(t, client, &respBody8)
	res, err = UploadPartCopy(client, bucket, object, source, uploadId, partNumber, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "2016-05-12T09:14:32Z", res.LastModified)
	ExpectEqual(t, "67b92a7c2a9b9c1809a6ae3295dcc127", res.ETag)
}

func TestCompleteMultipartUpload(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	uploadId := "a44cc9bab11bdc157676984aad851637"
	resource := fmt.Sprintf("/%s/%s", bucket, object)
	query := fmt.Sprintf("uploadId=%s", uploadId)
	reqBody := `{
		"parts":[
			{ "partNumber":1, "eTag":"a54357aff0632cce46d942af68356b38" },
			{ "partNumber":2, "eTag":"0c78aef83f66abc1fa1e8477f296d394" },
			{ "partNumber":3, "eTag":"acbd18db4cc2f85cedef654fccc4a4d8" }
		]
	}`
	respBody := `{
		"location":"http://bj.bcebos.com/BucketName/ObjectName",
		"bucket":"BucketName",
		"key":"object",
		"eTag":"3858f62230ac3c915f300c664312c11f"
	}`

	//case1: all is ok
	//mock http client, crc64 match
	options1 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(respBody),
		util.AddHeaders(map[string]string{
			my_http.BCE_CONTENT_CRC32C: "ContentCrc32c",
			my_http.ETAG:               "3858f62230ac3c915f300c664312c11f",
		}),
	}
	mockHttpClient1 := util.NewMockHTTPClient(options1...)
	ExpectEqual(t, true, mockHttpClient1 != nil)
	client.HTTPClient = mockHttpClient1
	args1 := &CompleteMultipartUploadArgs{
		UserMeta: map[string]string{
			"Key1": "Value1", "Key2": "Value2",
		},
		Process:           "process",
		ContentCrc32:      "ContentCrc32",
		ContentCrc32c:     "ContentCrc32c",
		ContentCrc32cFlag: true,
		ObjectExpires:     3,
		ContentCrc64ECMA:  "ContentCrc64ECMA",
		IfMatch:           "a54357aff0632cce46d942af68356b38",
		IfNoneMatch:       "*",
	}
	body1, _ := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, true, body1 != nil)
	res, err := CompleteMultipartUpload(client, bucket, object, uploadId, body1, args1, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, "http://bj.bcebos.com/BucketName/ObjectName", res.Location)
	ExpectEqual(t, "3858f62230ac3c915f300c664312c11f", res.ETag)

	//case2: body is nil
	res, err = CompleteMultipartUpload(client, bucket, object, uploadId, nil, args1, nil)
	ExpectEqual(t, bce.NewBceClientError("upload body info should not be emtpy"), err)
	ExpectEqual(t, nil, res)

	//case3: handle option error
	body3, _ := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, true, body3 != nil)
	res, err = CompleteMultipartUpload(client, bucket, object, uploadId, body3, args1, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	//case4: send request error
	args4 := args1
	err4 := fmt.Errorf("IO Error")
	options4 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetHTTPClientDoError(err4),
	}
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, true, mockHttpClient4 != nil)
	client.HTTPClient = mockHttpClient4
	body4, _ := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, true, body4 != nil)
	res, err = CompleteMultipartUpload(client, bucket, object, uploadId, body4, args4, nil)
	CheckMockHttpClientError(t, client.Config.Endpoint, resource, query, "Post", err4, err)
	ExpectEqual(t, nil, res)

	//case5:resp is fail
	args5 := args1
	options5 := util.RoundTripperOpts404
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, true, mockHttpClient5 != nil)
	client.HTTPClient = mockHttpClient5
	body5, _ := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, true, body5 != nil)
	res, err = CompleteMultipartUpload(client, bucket, object, uploadId, body5, args5, nil)
	ExpectEqual(t, bceServiceErro404.Error(), err.Error())
	ExpectEqual(t, nil, res)

	//case6: parse resp body fail
	args6 := args1
	errRespBody := "{adsf,dsf:fgrkjh}"
	options6 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(errRespBody),
		util.AddHeaders(map[string]string{
			my_http.BCE_CONTENT_CRC32C: "ContentCrc32c",
			my_http.ETAG:               "3858f62230ac3c915f300c664312c11f",
		}),
	}
	mockHttpClient6 := util.NewMockHTTPClient(options6...)
	ExpectEqual(t, true, mockHttpClient6 != nil)
	client.HTTPClient = mockHttpClient6
	body6, _ := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, true, body6 != nil)
	res, err = CompleteMultipartUpload(client, bucket, object, uploadId, body6, args6, nil)
	ExpectEqual(t, false, err == nil)
	ExpectEqual(t, nil, res)

	// case7:crc32c not match
	args7 := args1
	args7.ContentCrc32c = "sdfiuerhgfiuerg"
	options7 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(respBody),
		util.AddHeaders(map[string]string{
			my_http.BCE_CONTENT_CRC32C: "ContentCrc32c",
			my_http.ETAG:               "3858f62230ac3c915f300c664312c11f",
		}),
	}
	mockHttpClient7 := util.NewMockHTTPClient(options7...)
	ExpectEqual(t, true, mockHttpClient7 != nil)
	client.HTTPClient = mockHttpClient7
	body7, _ := bce.NewBodyFromString(reqBody)
	ExpectEqual(t, true, body7 != nil)
	res, err = CompleteMultipartUpload(client, bucket, object, uploadId, body7, args7, nil)
	ExpectEqual(t, crc32cMisMatchError(args7.ContentCrc32c, "ContentCrc32c"), err)
	ExpectEqual(t, "http://bj.bcebos.com/BucketName/ObjectName", res.Location)
	ExpectEqual(t, "3858f62230ac3c915f300c664312c11f", res.ETag)
}

func TestAbortMultipartUpload(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	uploadId := "a44cc9bab11bdc157676984aad851637"

	// case1: handle options error
	err = AbortMultipartUpload(client, bucket, object, uploadId, nil, ErrorOption)
	ExpectEqual(t, optionError, err)

	// case2: send request fail
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	err = AbortMultipartUpload(client, bucket, object, uploadId, nil, ErrorOption)
	ExpectEqual(t, true, err != nil)

	// case3: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	err = AbortMultipartUpload(client, bucket, object, uploadId, nil)
	ExpectEqual(t, bceServiceErro404, err)

	// case4: all is ok
	AttachMockHttpClientOk(t, client, nil)
	err = AbortMultipartUpload(client, bucket, object, uploadId, nil)
	ExpectEqual(t, nil, err)
}

func TestListParts(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	object := "test-object"
	uploadId := "a44cc9bab11bdc157676984aad851637"
	args := &ListPartsArgs{
		MaxParts:         10,
		PartNumberMarker: "part-number-marker",
	}

	// case1: handle options error
	res, err := ListParts(client, bucket, object, uploadId, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	res, err = ListParts(client, bucket, object, uploadId, args, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	// case3: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = ListParts(client, bucket, object, uploadId, args, nil)
	ExpectEqual(t, bceServiceErro404.Error(), err.Error())
	ExpectEqual(t, nil, res)

	//case4: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = ListParts(client, bucket, object, uploadId, args, nil)
	result5 := &ListPartsResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result5), err)
	ExpectEqual(t, nil, res)

	// case5: all is ok
	respBody5 := `{
		"bucket":"BucketName",
		"key":"object",
		"uploadId":"a44cc9bab11cbd156984767aad637851",
		"initiated":"2010-11-10T20:48:33Z",
		"owner":{ "id":"75aa570f8e7faeebf76c078efc7c6caea54ba06a", "displayName":"someName" },
		"storageClass":"STANDARD",
		"partNumberMarker":1,
		"nextPartNumberMarker":3,
		"maxParts":2,
		"isTruncated":true,
		"parts":[
			{
				"partNumber":2,
				"lastModified":"2010-11-10T20:48:34Z",
				"ETag":"7778aef83f66abc1fa1e8477f296d394",
				"size":10485760
			},
			{
				"partNumber":3,
				"lastModified":"2010-11-10T20:48:33Z",
				"ETag":"aaaa18db4cc2f85cedef654fccc4a4x8",
				"size":10485760
			}
		]
	}`
	AttachMockHttpClientOk(t, client, &respBody5)
	res, err = ListParts(client, bucket, object, uploadId, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.Parts))
}

func TestListMultipartUploads(t *testing.T) {
	//mock bce client
	client, err := NewMockBosClient()
	ExpectEqual(t, nil, err)
	// prepare parameters
	bucket := "test-bucket"
	args := &ListMultipartUploadsArgs{
		Delimiter:  "/",
		KeyMarker:  "key-marker",
		MaxUploads: 10,
		Prefix:     "prefix",
	}

	// case1: handle options error
	res, err := ListMultipartUploads(client, bucket, args, nil, ErrorOption)
	ExpectEqual(t, optionError, err)
	ExpectEqual(t, nil, res)

	// case2: send request error
	err2 := fmt.Errorf("IO Error")
	AttachMockHttpClientError(t, client, err2)
	res, err = ListMultipartUploads(client, bucket, args, nil)
	ExpectEqual(t, true, err != nil)
	ExpectEqual(t, nil, res)

	// case3: resp is fail
	options3 := util.RoundTripperOpts404
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, true, mockHttpClient3 != nil)
	client.HTTPClient = mockHttpClient3
	res, err = ListMultipartUploads(client, bucket, args, nil)
	ExpectEqual(t, bceServiceErro404.Error(), err.Error())
	ExpectEqual(t, nil, res)

	//case4: parse json body fail
	AttachMockHttpClientJsonBodyError(t, client)
	res, err = ListMultipartUploads(client, bucket, args, nil)
	result4 := &ListMultipartUploadsResult{}
	jsonDecoder := json.NewDecoder(bytes.NewBufferString(errorJsonBody))
	ExpectEqual(t, jsonDecoder.Decode(result4), err)
	ExpectEqual(t, nil, res)

	// case5: all is ok
	respBody5 := `{
		"bucket": "bucket",
		"keyMarker": "",
		"nextKeyMarker": "my-movie.m2ts",
		"nextUploadMarker": "c41cc9aad11cbd637851767bab156984",
		"maxUploads": 3,
		"isTruncated": true,
		"uploads": [
			{
				"key": "my-divisor",
				"uploadId": "a44cc9bab11bdc157676984aad851637",
				"owner": {
					"id": "75aa57f09aa0c8caeab4aeebf76c078efc7c6caea54ba06a",
					"displayName": "OwnerDisplayName"
				},
				"initiated": "2010-11-10T20:48:33Z",
				"storageClass": "STANDARD_IA"
			},
			{
				"key": "my-movie",
				"uploadId": "b44cc9bab11cbd156984767aad637851",
				"owner": {
					"id": "b1d16700c70b0b05597d7acd6a3f92be",
					"displayName": "OwnerDisplayName"
				},
				"initiated": "2010-11-10T20:48:33Z",
				"storageClass": "STANDARD"
			}
		]
	}`
	AttachMockHttpClientOk(t, client, &respBody5)
	res, err = ListMultipartUploads(client, bucket, args, nil)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, 2, len(res.Uploads))
}

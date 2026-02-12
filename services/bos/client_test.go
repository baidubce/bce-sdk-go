// nolint
package bos

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	my_http "github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/services/bos/api"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	EXISTS_BUCKET = "gosdk-unittest-bucket"
	EXISTS_OBJECT = "gosdk-unittest-object"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	ENDPOINT string
}

// init client with your ak/sk written in config.json
func init() {
	log.SetLogLevel(log.WARN)
	log.SetLogHandler(log.STDOUT)
	//log.SetLogHandler(log.STDERR | log.FILE)
	//log.SetRotateType(log.ROTATE_SIZE)
}

func currentFileName() string {
	_, file, _, _ := runtime.Caller(0)
	return file
}

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

type ErrorTypeTransport struct {
}

func (ett *ErrorTypeTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

var (
	bceServiceErro403 *bce.BceServiceError = bce.NewBceServiceError("AccessDenied", "403 Forbidden", "", http.StatusForbidden)
	bceServiceErro404 *bce.BceServiceError = bce.NewBceServiceError("NOTFound", "404 NOT Found", "", http.StatusNotFound)
	bceServiceErro408 *bce.BceServiceError = bce.NewBceServiceError("Timeout", "Request Timeout", "", http.StatusRequestTimeout)
	optionError       *bce.BceClientError  = bce.NewBceClientError("Handle bos client options failed: BosContext Options: error option")
)

func NewMockBosClient(ak, sk, endpoint, respBody string, options ...util.MockRoundTripperOption) (*Client, error) {
	config := NewBosClientConfig(ak, sk, endpoint)
	if len(options) == 0 {
		options = []util.MockRoundTripperOption{
			util.SetStatusCode(200),
			util.SetStatusMsg("200 OK"),
		}
	}
	if len(respBody) > 0 {
		options = append(options, util.SetRespBody(respBody))
	}

	mockHttpClient := util.NewMockHTTPClient(options...)
	if mockHttpClient == nil {
		return nil, fmt.Errorf("util.NewMockHTTPClient fail")
	}
	config = config.WithHttpClient(*mockHttpClient)
	return NewClientWithConfig(config)
}

func TestNewBosClient(t *testing.T) {
	//case1: NewClient(ak, sk, endpoint)
	ak, sk, endpoint := "test-ak", "test-sk", "test-endpoint"
	_, err := NewClient(ak, sk, endpoint)
	ExpectEqual(t.Errorf, nil, err)
	// case2: NewClientWithConfig
	config := &BosClientConfiguration{
		Ak:       ak,
		Sk:       sk,
		Endpoint: endpoint,
	}
	_, err = NewClientWithConfig(config)
	ExpectEqual(t.Errorf, nil, err)
	// case3: NewConfigWithXX
	http_client := &http.Client{}
	retry_policy := &bce.BackOffRetryPolicy{}
	config1 := config.WithAk("test-ak1").
		WithSk("test-sk1").
		WithEndpoint("test-endpoint1").
		WithDialTimeout(100).
		WithDisableKeepAlives(true).
		WithDownloadRateLimit(4096000).
		WithExclusiveHTTPClient(true).
		WithHttpClient(*http_client).
		WithHttpClientTimeout(100).
		WithIdleConnectionTimeout(100).
		WithKeepAlive(100).
		WithNoVerifySSL(true).
		WithPathStyleEnable(true).
		WithReadTimeout(100).
		WithRedirectDisabled(true).
		WithResponseHeaderTimeout(100).
		WithRetryPolicy(retry_policy).
		WithTLSHandshakeTimeout(100).
		WithUploadRateLimit(8192000).
		WithWriteTimeout(100).
		WithApiVersion("")
	ExpectEqual(t.Errorf, "test-ak1", config.Ak)
	ExpectEqual(t.Errorf, "test-sk1", config.Sk)
	ExpectEqual(t.Errorf, "test-endpoint1", config.Endpoint)
	ExpectEqual(t.Errorf, 100, *(config.DialTimeout))
	ExpectEqual(t.Errorf, true, config.DisableKeepAlives)
	ExpectEqual(t.Errorf, 4096000, *(config.DownloadRatelimit))
	ExpectEqual(t.Errorf, true, config.ExclusiveHTTPClient)
	ExpectEqual(t.Errorf, http_client, config.HTTPClient)
	ExpectEqual(t.Errorf, 100, *(config.HTTPClientTimeout))
	ExpectEqual(t.Errorf, 100, *(config.IdleConnectionTimeout))
	ExpectEqual(t.Errorf, 100, *(config.KeepAlive))
	ExpectEqual(t.Errorf, true, config.NoVerifySSL)
	ExpectEqual(t.Errorf, true, config.PathStyleEnable)
	ExpectEqual(t.Errorf, 100, *(config.ReadTimeout))
	ExpectEqual(t.Errorf, true, config.RedirectDisabled)
	ExpectEqual(t.Errorf, 100, *(config.ResponseHeaderTimeout))
	ExpectEqual(t.Errorf, retry_policy, config.retryPolicy)
	ExpectEqual(t.Errorf, 100, *(config.TLSHandshakeTimeout))
	ExpectEqual(t.Errorf, 8192000, *(config.UploadRatelimit))
	ExpectEqual(t.Errorf, 100, *(config.WriteTimeOut))
	ExpectEqual(t.Errorf, "", config.ApiVersion)
	_, err = NewClientWithConfig(config1)
	ExpectEqual(t.Errorf, nil, err)

	// error http transport
	config2 := config.WithHttpClient(http.Client{
		Timeout:   1000,
		Transport: &ErrorTypeTransport{},
	})
	_, err = NewClientWithConfig(config2)
	ExpectEqual(t.Errorf, nil, err)

	// empty parameter
	config3 := &BosClientConfiguration{}
	_, err = NewClientWithConfig(config3)
	ExpectEqual(t.Errorf, nil, err)

	config3.Ak = "AK"
	_, err = NewClientWithConfig(config3)
	ExpectEqual(t.Errorf, errors.New("secretKey should not be empty"), err)

	// NewStsClient
	_, err = NewStsClient("", "sk", "endpoint", 180)
	ExpectEqual(t.Errorf, errors.New("accessKeyId should not be empty"), err)

}

func TestListBuckets(t *testing.T) {
	respBody := `{
		"owner":{
			"id":"10eb6f5ff6ff4605bf044313e8f3ffa5",
			"displayName":"BosUser"
		},
		"buckets":[
			{
				"name":"bucket1",
				"location":"bj",
				"creationDate":"2016-04-05T10:20:35Z",
				"enableMultiAz":true
			},
			{
				"name":"bucket2",
				"location":"bj",
				"creationDate":"2016-04-05T16:41:58Z"
			}
		]
	}`

	// new mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	bosContext := client.NewBosContext(nil)
	ExpectEqual(t.Errorf, bosContext.ApiVersion, client.BosContext.ApiVersion)
	ExpectEqual(t.Errorf, bosContext.PathStyleEnable, client.BosContext.PathStyleEnable)

	// case1: ok, api version v1
	res, err := client.ListBuckets()
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Buckets))
	ExpectEqual(t.Errorf, "bucket1", res.Buckets[0].Name)
	ExpectEqual(t.Errorf, "bucket2", res.Buckets[1].Name)
	res, err = client.ListBucketsWithContext(context.Background())
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Buckets))
	ExpectEqual(t.Errorf, "bucket1", res.Buckets[0].Name)
	ExpectEqual(t.Errorf, "bucket2", res.Buckets[1].Name)

	// case2: ok, api version v2
	clientOptions := []api.Option{
		api.ApiVersion(api.API_VERSION_V2),
	}
	res, err = client.ListBuckets(clientOptions...)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Buckets))
	ExpectEqual(t.Errorf, "bucket1", res.Buckets[0].Name)
	ExpectEqual(t.Errorf, "bucket2", res.Buckets[1].Name)
	res, err = client.ListBucketsWithContext(context.Background(), clientOptions...)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Buckets))
	ExpectEqual(t.Errorf, "bucket1", res.Buckets[0].Name)
	ExpectEqual(t.Errorf, "bucket2", res.Buckets[1].Name)

	// case3: http client return 404, api version v1
	options := util.RoundTripperOpts404
	errorHttpClient := util.NewMockHTTPClient(options...)
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
	}
	var nilRes *api.ListBucketsResult = nil
	res, err = client.ListBuckets(clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro404, err)
	ExpectEqual(t.Errorf, nilRes, res)
	res, err = client.ListBucketsWithContext(context.Background(), clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro404, err)
	ExpectEqual(t.Errorf, nilRes, res)

	// case4: http client return 404, api version v2
	clientOptions = []api.Option{
		api.ApiVersion(api.API_VERSION_V2),
	}
	res, err = client.ListBuckets(clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro404, err)
	ExpectEqual(t.Errorf, nilRes, res)
	res, err = client.ListBucketsWithContext(context.Background(), clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro404, err)
	ExpectEqual(t.Errorf, nilRes, res)

	// case5: http client do error, api version v1
	netError := fmt.Errorf("net error")
	options = []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(netError),
	}
	urlError := url.Error{
		Op:  "Get",
		URL: "http://" + endpoint,
		Err: netError,
	}
	expectError := &bce.BceClientError{
		Message: fmt.Sprintf("execute http request failed! Retried 3 times, error: %s", urlError.Error()),
	}
	errorHttpClient = util.NewMockHTTPClient(options...)
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
	}
	res, err = client.ListBuckets(clientOptions...)
	ExpectEqual(t.Errorf, expectError, err)
	ExpectEqual(t.Errorf, nilRes, res)
	res, err = client.ListBucketsWithContext(context.Background(), clientOptions...)
	ExpectEqual(t.Errorf, expectError, err)
	ExpectEqual(t.Errorf, nilRes, res)

	// case6: http client do error, api version v2
	clientOptions = []api.Option{
		api.ApiVersion(api.API_VERSION_V2),
	}
	res, err = client.ListBuckets(clientOptions...)
	ExpectEqual(t.Errorf, expectError, err)
	ExpectEqual(t.Errorf, nilRes, res)
	res, err = client.ListBucketsWithContext(context.Background(), clientOptions...)
	ExpectEqual(t.Errorf, expectError, err)
	ExpectEqual(t.Errorf, nilRes, res)

	// option return error
	_, err = client.ListBuckets(api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	_, err = client.ListBucketsWithContext(context.Background(), api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
}

func TestListObjects(t *testing.T) {
	respBody := `{
		"name":"bucket",
		"prefix":"",
		"delimiter":"/",
		"marker":"",
		"maxKeys":1000,
		"isTruncated":false,
		"contents":[
			{
			    "key":"my-image.jpg",
			    "lastModified":"2009-10-12T17:50:30Z",
			    "eTag":"fba9dede5f27731c9771645a39863328",
			    "size":434234,
			    "storageClass":"STANDARD",
			    "owner":{
				    "id":"168bf6fd8fa74d9789f35a283a1f15e2",
				    "displayName":"mtd"
			    }
			},
			{
			    "key":"my-image1.jpg",
			    "lastModified":"2009-10-12T17:51:30Z",
			    "eTag":"0cce7caecc8309864f663d78d1293f98",
			    "size":124231,
			    "storageClass":"COLD",
			    "owner":{
				   "id":"168bf6fd8fa74d9789f35a283a1f15e2",
				   "displayName":"mtd"
				}
			}
		],
		"commonPrefixes":[
			{"prefix":"photos/"},
			{"prefix":"mtd/"}
	    ]
    }`

	// new mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1:8080"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1: ok, api version v1
	args := &api.ListObjectsArgs{Prefix: "test", MaxKeys: 10}
	res, err := client.ListObjects(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Contents))
	ExpectEqual(t.Errorf, 2, len(res.CommonPrefixes))
	ExpectEqual(t.Errorf, "my-image.jpg", res.Contents[0].Key)
	ExpectEqual(t.Errorf, "2009-10-12T17:51:30Z", res.Contents[1].LastModified)
	ExpectEqual(t.Errorf, "photos/", res.CommonPrefixes[0].Prefix)
	ExpectEqual(t.Errorf, "mtd/", res.CommonPrefixes[1].Prefix)
	res, err = client.ListObjectsWithContext(context.TODO(), EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Contents))
	ExpectEqual(t.Errorf, 2, len(res.CommonPrefixes))
	ExpectEqual(t.Errorf, "STANDARD", res.Contents[0].StorageClass)
	ExpectEqual(t.Errorf, 124231, res.Contents[1].Size)

	// case2: ok, api version v2
	clientOptions := []api.Option{
		api.ApiVersion(api.API_VERSION_V2),
	}
	res, err = client.ListObjects(EXISTS_BUCKET, args, clientOptions...)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Contents))
	ExpectEqual(t.Errorf, 2, len(res.CommonPrefixes))
	ExpectEqual(t.Errorf, "168bf6fd8fa74d9789f35a283a1f15e2", res.Contents[0].Owner.Id)
	ExpectEqual(t.Errorf, "COLD", res.Contents[1].StorageClass)
	ExpectEqual(t.Errorf, "photos/", res.CommonPrefixes[0].Prefix)
	ExpectEqual(t.Errorf, "mtd/", res.CommonPrefixes[1].Prefix)
	res, err = client.ListObjectsWithContext(context.TODO(), EXISTS_BUCKET, args, clientOptions...)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Contents))
	ExpectEqual(t.Errorf, 2, len(res.CommonPrefixes))
	ExpectEqual(t.Errorf, 434234, res.Contents[0].Size)
	ExpectEqual(t.Errorf, "my-image1.jpg", res.Contents[1].Key)

	// mock 403 error
	var nilRes *api.ListObjectsResult = nil
	roundTripperOptions := util.RoundTripperOpts403
	errorHttpClient := util.NewMockHTTPClient(roundTripperOptions...)
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
	}
	// case3: 403, api version v1
	res, err = client.ListObjects(EXISTS_BUCKET, args, clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro403, err)
	ExpectEqual(t.Errorf, nilRes, res)

	// case3: 403, api version v2
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
		api.ApiVersion(api.API_VERSION_V2),
	}
	res, err = client.ListObjects(EXISTS_BUCKET, args, clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro403, err)
	ExpectEqual(t.Errorf, nilRes, res)

	// option return error
	_, err = client.ListObjects(EXISTS_BUCKET, args, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	_, err = client.ListObjectsWithContext(context.TODO(), EXISTS_BUCKET, args, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
}

func TestSimpleListObjects(t *testing.T) {
	respBody := `{
		"name":"bucket",
		"prefix":"" ,
		"delimiter":"/",
		"marker":"",
		"maxKeys":1000,
		"isTruncated":false,
		"contents":[
			{
				"key":"my-image.jpg",
				"lastModified":"2009-10-12T17:50:30Z",
				"eTag":"fba9dede5f27731c9771645a39863328",
				"size":434234,
				"storageClass":"STANDARD",
				"owner":{
					"id":"168bf6fd8fa74d9789f35a283a1f15e2",
					"displayName":"mtd"
				}
			}
		],
		"commonPrefixes":[
			 {"prefix":"photos/"},
			 {"prefix":"mtd/"}
		]
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	// case1.1: ok, api version v1
	res, err := client.SimpleListObjects(EXISTS_BUCKET, "test", 10, "", "")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(res.Contents))
	ExpectEqual(t.Errorf, 2, len(res.CommonPrefixes))
	ExpectEqual(t.Errorf, "my-image.jpg", res.Contents[0].Key)
	ExpectEqual(t.Errorf, "photos/", res.CommonPrefixes[0].Prefix)

	// case1.2: ok, api version v2
	clientOptions := []api.Option{
		api.ApiVersion(api.API_VERSION_V2),
	}
	res, err = client.SimpleListObjects(EXISTS_BUCKET, "test", 10, "", "", clientOptions...)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(res.Contents))
	ExpectEqual(t.Errorf, 2, len(res.CommonPrefixes))
	ExpectEqual(t.Errorf, "my-image.jpg", res.Contents[0].Key)
	ExpectEqual(t.Errorf, "photos/", res.CommonPrefixes[0].Prefix)

	// mock server error: 408
	roundTripperOptions := util.RoundTripperOpts408
	errorHttpClient := util.NewMockHTTPClient(roundTripperOptions...)
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
	}
	var nilRes *api.ListObjectsResult = nil

	// case2.1: response 408 error, api version v1
	res, err = client.SimpleListObjects(EXISTS_BUCKET, "test", 10, "", "", clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro408, err)
	ExpectEqual(t.Errorf, nilRes, res)

	// case2.2: response 408 error, api version v2
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
		api.ApiVersion(api.API_VERSION_V2),
	}
	res, err = client.SimpleListObjects(EXISTS_BUCKET, "test", 10, "", "", clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro408, err)
	ExpectEqual(t.Errorf, nilRes, res)

	// option return error
	_, err = client.SimpleListObjects(EXISTS_BUCKET, "test", 10, "", "", api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
}
func TestListObjectVersions(t *testing.T) {
	respBody := `{
		"name":"bucket",
		"prefix":"",
		"delimiter":"",
		"marker":"",
		"maxKeys":1,
		"isTruncated":false,
		"nextVersionIdMarker":"AKyQMlG4ihY=",
		"contents":[
			{
			    "key":"my-image.jpg",
			    "lastModified":"2009-10-12T17:50:30Z",
			    "eTag":"fba9dede5f27731c9771645a39863328",
			    "size":434234,
			    "storageClass":"STANDARD",
			    "isLatest":1,
			    "versionId":"AKyQMlG4ihY=",
			    "owner":{
				   "id":"168bf6fd8fa74d9789f35a283a1f15e2",
				   "displayName":"mtd"
			    }
			}
		]
    }`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	res, err := client.ListObjectVersions(EXISTS_BUCKET, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "bucket", res.Name)
	ExpectEqual(t.Errorf, 1, len(res.Contents))

	// option return error
	_, err = client.ListObjectVersions(EXISTS_BUCKET, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
}
func TestHeadBucket(t *testing.T) {
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	// case1.1: ok, api version v1
	err = client.HeadBucket(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
	err = client.HeadBucketWithContext(context.TODO(), EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)

	// case1.2: ok, api version v2
	clientOptions := []api.Option{
		api.ApiVersion(api.API_VERSION_V2),
	}
	err = client.HeadBucket(EXISTS_BUCKET, clientOptions...)
	ExpectEqual(t.Errorf, nil, err)
	err = client.HeadBucketWithContext(context.TODO(), EXISTS_BUCKET, clientOptions...)
	ExpectEqual(t.Errorf, nil, err)

	// mock server error: 404
	roundTripperOptions := util.RoundTripperOpts404
	errorHttpClient := util.NewMockHTTPClient(roundTripperOptions...)
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
	}

	// case 2.1: response error: 404, api version v1
	err = client.HeadBucket(EXISTS_BUCKET, clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro404, err)
	err = client.HeadBucketWithContext(context.TODO(), EXISTS_BUCKET, clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro404, err)

	//case 2.2: response error: 404, api version v2
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
		api.ApiVersion(api.API_VERSION_V2),
	}
	err = client.HeadBucket(EXISTS_BUCKET, clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro404, err)
	err = client.HeadBucketWithContext(context.TODO(), EXISTS_BUCKET, clientOptions...)
	ExpectEqual(t.Errorf, bceServiceErro404, err)

	// option return error
	err = client.HeadBucket(EXISTS_BUCKET, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	err = client.HeadBucketWithContext(context.TODO(), EXISTS_BUCKET, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
}

func TestDoesBucketExist(t *testing.T) {
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	// case 1.1: ok, api version v1
	exist, err := client.DoesBucketExist(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, true, exist)
	ExpectEqual(t.Errorf, nil, err)

	// case 1.2: ok, api version v2
	clientOptions := []api.Option{
		api.ApiVersion(api.API_VERSION_V2),
	}
	exist, _ = client.DoesBucketExist("xxx", clientOptions...)
	ExpectEqual(t.Errorf, true, exist)
	ExpectEqual(t.Errorf, nil, err)

	// mock server error: 404
	roundTripperOptions := util.RoundTripperOpts404
	errorHttpClient := util.NewMockHTTPClient(roundTripperOptions...)
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
	}

	// case 2.1: response 404 error, api version v1
	exist, err = client.DoesBucketExist(EXISTS_BUCKET, clientOptions...)
	ExpectEqual(t.Errorf, false, exist)
	ExpectEqual(t.Errorf, nil, err)

	// case 2.2: response 404 error, api version v2
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
		api.ApiVersion(api.API_VERSION_V2),
	}
	exist, _ = client.DoesBucketExist("xxx", clientOptions...)
	ExpectEqual(t.Errorf, false, exist)
	ExpectEqual(t.Errorf, nil, err)

	// mock server error: 403
	roundTripperOptions = util.RoundTripperOpts403
	errorHttpClient = util.NewMockHTTPClient(roundTripperOptions...)
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
	}

	// case 2.1: response 403 error, api version v1
	exist, err = client.DoesBucketExist(EXISTS_BUCKET, clientOptions...)
	ExpectEqual(t.Errorf, true, exist)
	ExpectEqual(t.Errorf, nil, err)

	// case 2.2: response 403 error, api version v2
	clientOptions = []api.Option{
		api.HTTPClient(errorHttpClient),
		api.ApiVersion(api.API_VERSION_V2),
	}
	exist, _ = client.DoesBucketExist("xxx", clientOptions...)
	ExpectEqual(t.Errorf, true, exist)
	ExpectEqual(t.Errorf, nil, err)

	// option return error
	_, err = client.DoesBucketExist("xxx", api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
}

func TestIsNsBucket(t *testing.T) {
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	// case1: 200 ok, return true
	roundTripperOptions := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			my_http.BCE_BUCKET_TYPE: api.NAMESPACE_BUCKET,
		}),
	}
	httpClient1 := util.NewMockHTTPClient(roundTripperOptions...)
	clientOptions := []api.Option{
		api.HTTPClient(httpClient1),
	}

	ExpectEqual(t.Errorf, true, client.IsNsBucket(EXISTS_BUCKET, clientOptions...))

	// case2: 403 ok, return true
	roundTripperOptions2 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusForbidden),
		util.SetStatusMsg(http.StatusText(http.StatusForbidden)),
		util.AddHeaders(map[string]string{
			my_http.BCE_BUCKET_TYPE: api.NAMESPACE_BUCKET,
		}),
	}
	httpClient2 := util.NewMockHTTPClient(roundTripperOptions2...)
	clientOptions2 := []api.Option{
		api.HTTPClient(httpClient2),
	}

	ExpectEqual(t.Errorf, true, client.IsNsBucket(EXISTS_BUCKET, clientOptions2...))

	// case3: retunrn false
	roundTripperOptions3 := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
	}
	httpClient3 := util.NewMockHTTPClient(roundTripperOptions3...)
	clientOptions3 := []api.Option{
		api.HTTPClient(httpClient3),
	}

	ExpectEqual(t.Errorf, false, client.IsNsBucket(EXISTS_BUCKET, clientOptions3...))

}

func TestPutBucket(t *testing.T) {
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.LOCATION: "beijing",
		}),
	}
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	res, err := client.PutBucket("test-put-bucket")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, "beijing", res)

	res, err = client.PutBucketWithArgs("test-put-bucket", nil)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, "beijing", res)

	res, err = client.PutBucket("test-put-bucket", api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, "", res)

	res, err = client.PutBucketWithArgs("test-put-bucket", nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, "", res)
}
func TestDeleteBucket(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucket("test-put-bucket")
	ExpectEqual(t.Errorf, err, nil)
	err = client.DeleteBucket("test-put-bucket", api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
}
func TestGetBucketLocation(t *testing.T) {
	respBody := `{
		"locationConstraint": "bj"
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketLocation(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, "bj", res)
	res, err = client.GetBucketLocation(EXISTS_BUCKET, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, "", res)
}
func TestPutBucketAcl(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	// NewBodyFromString
	body, _ := bce.NewBodyFromString("acl")
	err = client.PutBucketAcl(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	err = client.PutBucketAcl(EXISTS_BUCKET, body, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	// PutBucketAclFromCanned
	err = client.PutBucketAclFromCanned(EXISTS_BUCKET, api.CANNED_ACL_PUBLIC_READ)
	ExpectEqual(t.Errorf, err, nil)
	err = client.PutBucketAclFromCanned(EXISTS_BUCKET, api.CANNED_ACL_PUBLIC_READ, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	// PutBucketAclFromString
	err = client.PutBucketAclFromString(EXISTS_BUCKET, "acl")
	ExpectEqual(t.Errorf, err, nil)
	err = client.PutBucketAclFromString(EXISTS_BUCKET, "acl", api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	// PutBucketAclFromFile
	fname := "/tmp/test-put-bucket-acl-by-file"
	f, _ := os.Create(fname)
	f.WriteString("acl")
	f.Close()
	err = client.PutBucketAclFromFile(EXISTS_BUCKET, fname)
	ExpectEqual(t.Errorf, err, nil)
	err = client.PutBucketAclFromFile(EXISTS_BUCKET, fname, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	os.Remove(fname)
	// PutBucketAclFromStruct
	args := &api.PutBucketAclArgs{
		AccessControlList: []api.GrantType{
			{
				Grantee: []api.GranteeType{
					{Id: "e13b12d0131b4c8bae959df4969387b8"},
				},
				Permission: []string{
					"FULL_CONTROL",
				},
			},
		},
	}
	err = client.PutBucketAclFromStruct(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
	err = client.PutBucketAclFromStruct(EXISTS_BUCKET, args, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
}
func TestGetBucketAcl(t *testing.T) {
	acl := `{
		"accessControlList":[
			{
				"grantee":[{
					"id":"e13b12d0131b4c8bae959df4969387b8"
				}],
				"permission":["FULL_CONTROL"]
			}
		]
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, acl)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketAcl(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, "e13b12d0131b4c8bae959df4969387b8", res.AccessControlList[0].Grantee[0].Id)
	ExpectEqual(t.Errorf, "FULL_CONTROL", res.AccessControlList[0].Permission[0])

	acl1 := `{
		"accessControlList":[
			{
				"grantee":[
					{"id":"e13b12d0131b4c8bae959df4969387b8"},
					{"id":"a13b12d0131b4c8bae959df4969387b8"}
				],
				"permission":["FULL_CONTROL"]
			}
		]
	}`
	// mock bos client
	client1, err := NewMockBosClient(ak, sk, endpoint, acl1)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client1.GetBucketAcl(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.AccessControlList[0].Grantee[0].Id, "e13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, res.AccessControlList[0].Grantee[1].Id, "a13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, res.AccessControlList[0].Permission[0], "FULL_CONTROL")
	res, err = client1.GetBucketAcl(EXISTS_BUCKET, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, res, nil)
}
func TestPutBucketLogging(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	//PutBucketLogging
	body, _ := bce.NewBodyFromString(
		`{"targetBucket": "gosdk-unittest-bucket", "targetPrefix": "my-log/"}`)
	err = client.PutBucketLogging(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	//PutBucketLoggingFromString
	err = client.PutBucketLoggingFromString(EXISTS_BUCKET, "logging")
	ExpectEqual(t.Errorf, err, nil)
	//PutBucketLoggingFromStruct
	obj := &api.PutBucketLoggingArgs{
		TargetBucket: "gosdk-unittest-bucket",
		TargetPrefix: "my-log3/",
	}
	err = client.PutBucketLoggingFromStruct(EXISTS_BUCKET, obj)
	ExpectEqual(t.Errorf, err, nil)
}
func TestGetBucketLogging(t *testing.T) {
	logging := `{"status":"enabled", "targetBucket": "gosdk-unittest-bucket", "targetPrefix": "my-log/"}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, logging)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketLogging(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.TargetBucket, "gosdk-unittest-bucket")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.TargetPrefix, "my-log/")

	logging1 := `{"status":"enabled", "targetBucket": "gosdk-unittest-bucket", "targetPrefix": "my-log2/"}`
	client, err = NewMockBosClient(ak, sk, endpoint, logging1)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client.GetBucketLogging(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.TargetBucket, "gosdk-unittest-bucket")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.TargetPrefix, "my-log2/")

	logging2 := `{"status":"disabled"}`
	client, err = NewMockBosClient(ak, sk, endpoint, logging2)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client.GetBucketLogging(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Status, "disabled")
}
func TestDeleteBucketLogging(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketLogging(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutBucketLifecycle(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	//PutBucketLifecycle
	body, _ := bce.NewBodyFromString("lifecyle json string")
	err = client.PutBucketLifecycle(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	//PutBucketLifecycleFromString
	err = client.PutBucketLifecycleFromString(EXISTS_BUCKET, "lifecycle json string")
	ExpectEqual(t.Errorf, err, nil)
}
func TestGetBucketLifecycle(t *testing.T) {
	str := `{
		"rule": [
			{
				"id": "transition-to-cold",
				"status": "enabled",
				"resource": ["gosdk-unittest-bucket/test*"],
				"condition": { "time": { "dateGreaterThan": "2018-09-07T00:00:00Z" } },
				"action": { "name": "DeleteObject" }
			}
		]
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, str)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketLifecycle(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Rule[0].Id, "transition-to-cold")
	ExpectEqual(t.Errorf, res.Rule[0].Status, "enabled")
	ExpectEqual(t.Errorf, res.Rule[0].Resource[0], "gosdk-unittest-bucket/test*")
	ExpectEqual(t.Errorf, res.Rule[0].Condition.Time.DateGreaterThan, "2018-09-07T00:00:00Z")
	ExpectEqual(t.Errorf, res.Rule[0].Action.Name, "DeleteObject")

	obj := `{
		"rule": [
			{
				"id": "transition-to-cold",
				"status": "enabled",
				"resource": ["gosdk-unittest-bucket/test*"],
				"condition": {"time": {"dateGreaterThan": "2018-09-07T00:00:00Z"}},
				"action": {"name": "DeleteObject"}
			}
		]
	}`
	client, err = NewMockBosClient(ak, sk, endpoint, obj)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client.GetBucketLifecycle(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Rule[0].Id, "transition-to-cold")
	ExpectEqual(t.Errorf, res.Rule[0].Status, "enabled")
	ExpectEqual(t.Errorf, res.Rule[0].Resource[0], "gosdk-unittest-bucket/test*")
	ExpectEqual(t.Errorf, res.Rule[0].Condition.Time.DateGreaterThan, "2018-09-07T00:00:00Z")
	ExpectEqual(t.Errorf, res.Rule[0].Action.Name, "DeleteObject")
}
func TestDeleteBucketLifecycle(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketLifecycle(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutBucketStorageClass(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.PutBucketStorageclass(EXISTS_BUCKET, api.STORAGE_CLASS_STANDARD_IA)
	ExpectEqual(t.Errorf, err, nil)
}
func TestGetBucketStorageClass(t *testing.T) {
	respBody := ` {
		"storageClass": "COLD"
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	res, err := client.GetBucketStorageclass(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, "COLD", res)
}
func TestPutBucketReplication(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	mockReplicationJosnStr := "replication json stirng"
	//PutBucketReplication
	body, _ := bce.NewBodyFromString(mockReplicationJosnStr)
	err = client.PutBucketReplication(EXISTS_BUCKET, body, "")
	ExpectEqual(t.Errorf, err, nil)

	//PutBucketReplicationFromFile
	fname := "/tmp/test-put-bucket-replication-by-file"
	f, _ := os.Create(fname)
	f.WriteString(mockReplicationJosnStr)
	f.Close()
	err = client.PutBucketReplicationFromFile(EXISTS_BUCKET, fname, "")
	os.Remove(fname)
	ExpectEqual(t.Errorf, err, nil)

	//PutBucketReplicationFromString
	err = client.PutBucketReplicationFromString(EXISTS_BUCKET, mockReplicationJosnStr, "")
	ExpectEqual(t.Errorf, err, nil)

	//PutBucketReplicationFromStruct
	args := &api.PutBucketReplicationArgs{
		Id:               "abc",
		Status:           "enabled",
		Resource:         []string{"gosdk-unittest-bucket/films"},
		Destination:      &api.BucketReplicationDescriptor{Bucket: "bos-rd-su-test", StorageClass: "COLD"},
		ReplicateDeletes: "disabled",
	}
	err = client.PutBucketReplicationFromStruct(EXISTS_BUCKET, args, "")
	ExpectEqual(t.Errorf, err, nil)
}
func TestGetBucketReplication(t *testing.T) {
	str := `{
		"id": "abc",
		"status":"enabled",
		"resource": ["gosdk-unittest-bucket/films"],
		"destination": {
			"bucket": "bos-rd-su-test",
			"storageClass": "COLD"
		},
		"replicateDeletes": "disabled"
    }`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, str)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketReplication(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Id, "abc")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.Resource[0], "gosdk-unittest-bucket/films")
	ExpectEqual(t.Errorf, res.Destination.Bucket, "bos-rd-su-test")
	ExpectEqual(t.Errorf, res.ReplicateDeletes, "disabled")

	str1 := `{
		"id": "abc",
		"status":"enabled",
		"resource": ["gosdk-unittest-bucket/films"],
		"destination": {
			"bucket": "bos-rd-su-test",
			"storageClass": "COLD"
		},
		"replicateDeletes": "disabled"
    }`
	client1, err := NewMockBosClient(ak, sk, endpoint, str1)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client1.GetBucketReplication(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Id, "abc")
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.Resource[0], "gosdk-unittest-bucket/films")
	ExpectEqual(t.Errorf, res.Destination.Bucket, "bos-rd-su-test")
	ExpectEqual(t.Errorf, res.ReplicateDeletes, "disabled")

	listResStr := `{
	"rules": [
		{
			"status": "enabled",
			"resource": ["src-bucket-name/abc","src-bucket-name/cd*"],
			"destination": {"bucket": "dst-bucket-name","storageClass": "COLD"},
			"replicateHistory": {"storageClass": "COLD"},
			"replicateDeletes": "enabled",
			"id": "sample-bucket-replication-config",
			"createTime": 1583060606,
			"destRegion": "bj"
		}
	]}`
	client2, err := NewMockBosClient(ak, sk, endpoint, listResStr)
	ExpectEqual(t.Errorf, nil, err)
	listRes, err := client2.ListBucketReplication(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, true, listRes != nil)
}
func TestGetBucketReplicationProcess(t *testing.T) {
	progress := `{
		"status":"enabled",
		"historyReplicationPercent":5,
		"latestReplicationTime":"1504448315"
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, progress)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketReplicationProgress(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Status, "enabled")
	ExpectEqual(t.Errorf, res.LatestReplicationTime, "1504448315")
	ExpectEqual(t.Errorf, strconv.FormatFloat(res.HistoryReplicationPercent, 'f', 5, 64), "5.00000")
}
func TestDeleteBucketReplication(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketReplication(EXISTS_BUCKET, "")
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutBucketEncryption(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.PutBucketEncryption(EXISTS_BUCKET, api.ENCRYPTION_AES256)
	ExpectEqual(t.Errorf, err, nil)
}
func TestGetBucketEncryption(t *testing.T) {
	encryption := ` {
		"encryptionAlgorithm":"AES256"
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, encryption)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketEncryption(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, api.ENCRYPTION_AES256)
}
func TestDeleteBucketEncryption(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketEncryption(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutBucketStaticWebsite(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	conf := `{"index": "index.html", "notFound":"blank.html"}`

	//PutBucketStaticWebsite
	body, _ := bce.NewBodyFromString(conf)
	err = client.PutBucketStaticWebsite(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)

	//PutBucketStaticWebsiteFromString
	err = client.PutBucketStaticWebsiteFromString(EXISTS_BUCKET, conf)
	ExpectEqual(t.Errorf, err, nil)

	//PutBucketStaticWebsiteFromStruct
	obj := &api.PutBucketStaticWebsiteArgs{Index: "index.html", NotFound: "blank.html"}
	err = client.PutBucketStaticWebsiteFromStruct(EXISTS_BUCKET, obj)
	ExpectEqual(t.Errorf, err, nil)

	//SimplePutBucketStaticWebsite
	err = client.SimplePutBucketStaticWebsite(EXISTS_BUCKET, "index.html", "blank.html")
	ExpectEqual(t.Errorf, err, nil)

}
func TestGetBucketStaticWebsite(t *testing.T) {
	conf := `{"index": "index.html", "notFound":"blank.html"}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, conf)
	ExpectEqual(t.Errorf, nil, err)

	res, err := client.GetBucketStaticWebsite(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Index, "index.html")
	ExpectEqual(t.Errorf, res.NotFound, "blank.html")
}
func TestDeleteBucketStaticWebsite(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketStaticWebsite(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutBucketCors(t *testing.T) {
	corsConf := `{
		"corsConfiguration": [
			{
				"allowedOrigins": ["https://www.baidu.com"],
				"allowedMethods": ["GET"],
				"maxAgeSeconds": 1800
			}
		]
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, corsConf)
	ExpectEqual(t.Errorf, nil, err)
	//PutBucketCors
	body, _ := bce.NewBodyFromString(corsConf)
	err = client.PutBucketCors(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	//PutBucketCorsFromFile
	fname := "/tmp/test-put-bucket-cors-by-file"
	f, _ := os.Create(fname)
	f.WriteString(corsConf)
	f.Close()
	err = client.PutBucketCorsFromFile(EXISTS_BUCKET, fname)
	os.Remove(fname)
	ExpectEqual(t.Errorf, err, nil)
	//PutBucketCorsFromString
	err = client.PutBucketCorsFromString(EXISTS_BUCKET, corsConf)
	ExpectEqual(t.Errorf, err, nil)
	//PutBucketCorsFromStruct
	obj := &api.PutBucketCorsArgs{
		CorsConfiguration: []api.BucketCORSType{
			{
				AllowedOrigins: []string{"https://www.baidu.com"},
				AllowedMethods: []string{"GET"},
				MaxAgeSeconds:  1200,
			},
		},
	}
	err = client.PutBucketCorsFromStruct(EXISTS_BUCKET, obj)
	ExpectEqual(t.Errorf, err, nil)
}
func TestGetBucketCors(t *testing.T) {
	corsConf := `{
		"corsConfiguration": [
			{
				"allowedOrigins": ["https://www.baidu.com"],
				"allowedMethods": ["GET"],
				"maxAgeSeconds": 1800
			}
		]
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, corsConf)
	ExpectEqual(t.Errorf, nil, err)

	res, err := client.GetBucketCors(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedOrigins[0], "https://www.baidu.com")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].AllowedMethods[0], "GET")
	ExpectEqual(t.Errorf, res.CorsConfiguration[0].MaxAgeSeconds, 1800)
}
func TestDeleteBucketCors(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketCors(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutBucketCopyrightProtection(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.PutBucketCopyrightProtection(EXISTS_BUCKET,
		"gosdk-unittest-bucket/test-put-object", "gosdk-unittest-bucket/films/*")
	ExpectEqual(t.Errorf, err, nil)
}
func TestGetBucketCopyrightProtection(t *testing.T) {
	copyRight := `{	
		"resource": [
			"bucket/prefix/*",
			"bucket/*/suffix"
		]
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, copyRight)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketCopyrightProtection(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res[0], "bucket/prefix/*")
	ExpectEqual(t.Errorf, res[1], "bucket/*/suffix")
}
func TestDeleteBucketCopyrightProtection(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketCopyrightProtection(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}

func TestPutObject(t *testing.T) {
	// mock http transport options
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "827ccb0eea8a706c4c34a16891f84e7b",
		}),
	}
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	args := &api.PutObjectArgs{StorageClass: api.STORAGE_CLASS_COLD}
	body, _ := bce.NewBodyFromString("12345")
	//PutObject
	etag, err := client.PutObject(EXISTS_BUCKET, "test-put-object", body, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObject(EXISTS_BUCKET, "test-put-object", body, args, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)

	//PutObjectWithContext
	body1, _ := bce.NewBodyFromString("12345")
	etag, err = client.PutObjectWithContext(context.Background(), EXISTS_BUCKET, "test-put-object", body1, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObjectWithContext(context.Background(), EXISTS_BUCKET, "test-put-object", body1, args, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)
	//PutObjectWithCallback
	body2, _ := bce.NewBodyFromString("12345")
	etag, _, err = client.PutObjectWithCallback(EXISTS_BUCKET, "test-put-object", body2, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, _, err = client.PutObjectWithCallback(EXISTS_BUCKET, "test-put-object", body2, args, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)
}

func TestBasicPutObject(t *testing.T) {
	// mock http transport options
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "827ccb0eea8a706c4c34a16891f84e7b",
		}),
	}
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	body, _ := bce.NewBodyFromString("12345")
	etag, err := client.BasicPutObject(EXISTS_BUCKET, "test-put-object", body)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.BasicPutObject(EXISTS_BUCKET, "test-put-object", body, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, "", etag)
}

func TestPutObjectFromBytes(t *testing.T) {
	// mock http transport options
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "827ccb0eea8a706c4c34a16891f84e7b",
		}),
	}
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	arr := []byte("12345")
	//PutObjectFromBytes
	etag, err := client.PutObjectFromBytes(EXISTS_BUCKET, "test-put-object", arr, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObjectFromBytes(EXISTS_BUCKET, "test-put-object", arr, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)
	//PutObjectFromBytesWithContext
	etag, err = client.PutObjectFromBytesWithContext(context.Background(), EXISTS_BUCKET, "test-put-object", arr, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObjectFromBytesWithContext(context.Background(),
		EXISTS_BUCKET, "test-put-object", arr, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)
}

func TestPutObjectFromString(t *testing.T) {
	// mock http transport options
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "827ccb0eea8a706c4c34a16891f84e7b",
		}),
	}
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	etag, err := client.PutObjectFromString(EXISTS_BUCKET, "test-put-object", "12345", nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObjectFromString(EXISTS_BUCKET, "test-put-object", "12345", nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, "", etag)

	etag, err = client.PutObjectFromStringWithContext(context.Background(), EXISTS_BUCKET, "test-put-object", "12345", nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObjectFromStringWithContext(context.Background(),
		EXISTS_BUCKET, "test-put-object", "12345", nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, "", etag)
}

func TestPutObjectFromFile(t *testing.T) {
	// mock http transport options
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "827ccb0eea8a706c4c34a16891f84e7b",
		}),
	}
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	fname := "/tmp/test-put-file"
	f, _ := os.Create(fname)
	f.WriteString("12345")
	f.Close()
	etag, err := client.PutObjectFromFile(EXISTS_BUCKET, "test-put-object", fname, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObjectFromFile(EXISTS_BUCKET, "test-put-object", fname, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)
	etag, err = client.PutObjectFromFileWithContext(context.Background(), EXISTS_BUCKET, "test-put-object", fname, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObjectFromFileWithContext(context.Background(),
		EXISTS_BUCKET, "test-put-object", fname, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)
	etag, _, err = client.PutObjectFromFileWithCallback(EXISTS_BUCKET, "test-put-object", fname, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, _, err = client.PutObjectFromFileWithCallback(EXISTS_BUCKET, "test-put-object", fname, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)
	args := &api.PutObjectArgs{ContentLength: 6}
	etag, err = client.PutObjectFromFile(EXISTS_BUCKET, "test-put-object", fname, args)
	ExpectEqual(t.Errorf, true, err != nil)
	ExpectEqual(t.Errorf, "", etag)
	os.Remove(fname)
}

func TestPutObjectFromStream(t *testing.T) {
	// mock http transport options
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "827ccb0eea8a706c4c34a16891f84e7b",
		}),
	}
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	fname := "/tmp/test-put-file"
	fw, _ := os.Create(fname)
	defer os.Remove(fname)
	fw.WriteString("12345")
	fw.Close()
	fr, _ := os.Open(fname)
	defer fr.Close()
	etag, err := client.PutObjectFromStream(EXISTS_BUCKET, "test-put-object", fr, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObjectFromStream(EXISTS_BUCKET, "test-put-object", fr, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)
	buffer := bytes.NewBufferString("s string")
	etag, err = client.PutObjectFromStreamWithContext(context.Background(), EXISTS_BUCKET, "test-put-object", buffer, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", etag)
	etag, err = client.PutObjectFromStreamWithContext(context.Background(),
		EXISTS_BUCKET, "test-put-object", buffer, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, "", etag)
}

func TestPostObject(t *testing.T) {
	// mock bos client
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.CONTENT_MD5):       "Zh+ACfqOVqnQ6UoKZEOX1w==",
			http.CanonicalHeaderKey(my_http.ETAG):              "827ccb0eea8a706c4c34a16891f84e7b",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32): "1922069637",
		}),
	}
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)
	content := "this is a test string for post object"
	args := &api.PostObjectArgs{
		Expiration:         180 * time.Second,
		ContentLengthLower: 0,
		ContentLengthUpper: 4096,
	}
	res, err := client.PostObjectFromBytes("test-bucket", "test-object", []byte(content), args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "1922069637", res.ContentCrc32)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", res.ETag)
	ExpectEqual(t.Errorf, "Zh+ACfqOVqnQ6UoKZEOX1w==", res.ContentMD5)
	res, err = client.PostObjectFromString("test-bucket", "test-object", content, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "1922069637", res.ContentCrc32)
	ExpectEqual(t.Errorf, "827ccb0eea8a706c4c34a16891f84e7b", res.ETag)
	ExpectEqual(t.Errorf, "Zh+ACfqOVqnQ6UoKZEOX1w==", res.ContentMD5)
	res, err = client.PostObjectFromStream("test-bucket", "test-object", bytes.NewBufferString(content), args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "1922069637", res.ContentCrc32)
	res, err = client.PostObjectFromFile("test-bucket", "test-object", currentFileName(), args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "1922069637", res.ContentCrc32)

	res, err = client.PostObjectFromBytes("test-bucket", "test-object", []byte(content), args, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, nil, res)
	res, err = client.PostObjectFromString("test-bucket", "test-object", content, args, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, nil, res)
	res, err = client.PostObjectFromStream("test-bucket", "test-object", bytes.NewBufferString(content), args, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, nil, res)
	res, err = client.PostObjectFromFile("test-bucket", "test-object", currentFileName(), args, api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, nil, res)
}
func TestOptionObject(t *testing.T) {
	// mock bos client
	roundTripperOpts := []util.MockRoundTripperOption{
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
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	args := &api.OptionsObjectArgs{
		Origin:         "string",
		RequestMethod:  "Post",
		RequestHeaders: []string{"headr1", "header2"},
	}

	res, err := client.OptionsObject("test-bucket", "test-object", args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.AllowCredentials)
	ExpectEqual(t.Errorf, 2, len(res.AllowHeaders))
	ExpectEqual(t.Errorf, "header1", res.AllowHeaders[0])
	ExpectEqual(t.Errorf, "header2", res.AllowHeaders[1])
	ExpectEqual(t.Errorf, 2, len(res.AllowMethods))
	ExpectEqual(t.Errorf, "PUT", res.AllowMethods[0])
	ExpectEqual(t.Errorf, "GET", res.AllowMethods[1])
	ExpectEqual(t.Errorf, "origin", res.AllowOrigin)
	ExpectEqual(t.Errorf, 1, len(res.ExposeHeaders))
	ExpectEqual(t.Errorf, "POST", res.ExposeHeaders[0])
	ExpectEqual(t.Errorf, 10, res.MaxAge)
}
func TestCopyObject(t *testing.T) {
	// mock bos client
	respBody := `{
		"lastModified":"2009-10-28T22:32:00Z",
		"ETag":"9b2cf535f27731c974343645a3985328"
	}`
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.BCE_VERSION_ID): "AKyQ9DRhhoY=",
		}),
	}
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody, roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	// copy object ok
	args := new(api.CopyObjectArgs)
	args.StorageClass = api.STORAGE_CLASS_COLD
	res, err := client.CopyObject(EXISTS_BUCKET, "test-copy-object", EXISTS_BUCKET, "test-put-object", args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.LastModified, "2009-10-28T22:32:00Z")
	ExpectEqual(t.Errorf, res.ETag, "9b2cf535f27731c974343645a3985328")
	t.Logf("copy result: %+v", res)
	res, err = client.CopyObjectWithContext(context.TODO(), EXISTS_BUCKET,
		"test-copy-object", EXISTS_BUCKET, "test-put-object", args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.LastModified, "2009-10-28T22:32:00Z")
	ExpectEqual(t.Errorf, res.ETag, "9b2cf535f27731c974343645a3985328")

	// copy object service error
	respBody1 := `{
		"code":"InternalError",
		"message":"We encountered an internal error. Please try again.",
		"requestId":"52454655-5345-4420-4259-204e47494e58"
	}`
	client1, err := NewMockBosClient(ak, sk, endpoint, respBody1)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client1.CopyObject(EXISTS_BUCKET, "test-copy-object", EXISTS_BUCKET, "test-put-object", args)
	ExpectEqual(t.Errorf, err, bce.NewBceServiceError("InternalError", "We encountered an internal error. Please try again.",
		"52454655-5345-4420-4259-204e47494e58", 500))
	ExpectEqual(t.Errorf, res, nil)

	//BasicCopyObject
	client2, err := NewMockBosClient(ak, sk, endpoint, respBody, roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client2.BasicCopyObject(EXISTS_BUCKET, "test-copy-object", EXISTS_BUCKET, "test-put-object")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.LastModified, "2009-10-28T22:32:00Z")
	ExpectEqual(t.Errorf, res.ETag, "9b2cf535f27731c974343645a3985328")
}
func TestGetObject(t *testing.T) {
	//mock bos client
	respBody := "this is a test-string for testing GetObject"
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.BCE_VERSION_ID):                    "AKyQ9DRhhoY=",
			http.CanonicalHeaderKey(my_http.CACHE_CONTROL):                     "private",
			http.CanonicalHeaderKey(my_http.CONTENT_DISPOSITION):               "attachment; filename=\"download.txt\"",
			http.CanonicalHeaderKey(my_http.CONTENT_LENGTH):                    strconv.FormatInt(int64(len(respBody)), 10),
			http.CanonicalHeaderKey(my_http.CONTENT_TYPE):                      "text/plain",
			http.CanonicalHeaderKey(my_http.BCE_USER_METADATA_PREFIX) + "Key1": "Value1",
			http.CanonicalHeaderKey(my_http.BCE_USER_METADATA_PREFIX) + "Key2": "Value2",
			http.CanonicalHeaderKey(my_http.BCE_STORAGE_CLASS):                 api.STORAGE_CLASS_ARCHIVE,
		}),
	}
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody, roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)

	// GetObject
	args := map[string]string{"ContentType": "text/html"}
	res, err := client.GetObject(EXISTS_BUCKET, "test-put-object", args, int64(0), int64(len(respBody)))
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.VersionId, "AKyQ9DRhhoY=")
	ExpectEqual(t.Errorf, res.CacheControl, "private")
	ExpectEqual(t.Errorf, res.ContentDisposition, "attachment; filename=\"download.txt\"")
	ExpectEqual(t.Errorf, res.ContentLength, int64(len(respBody)))
	ExpectEqual(t.Errorf, res.ContentType, "text/plain")
	ExpectEqual(t.Errorf, res.UserMeta["Key1"], "Value1")
	ExpectEqual(t.Errorf, res.UserMeta["Key2"], "Value2")
	ExpectEqual(t.Errorf, res.StorageClass, api.STORAGE_CLASS_ARCHIVE)
	buf := make([]byte, len(respBody))
	n, err := res.Body.Read(buf)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, len(respBody), n)
	ExpectEqual(t.Errorf, respBody, string(buf))

	// GetObjectWithContext
	res, err = client.GetObjectWithContext(context.Background(), EXISTS_BUCKET,
		"test-put-object", args, int64(0), int64(len(respBody)))
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.VersionId, "AKyQ9DRhhoY=")

	// GetObjectWithArgs
	res, err = client.GetObjectWithArgs(context.Background(), EXISTS_BUCKET, "test-put-object", nil)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.VersionId, "AKyQ9DRhhoY=")
	// GetObjectWithArgs with error option
	res, err = client.GetObjectWithArgs(context.Background(), EXISTS_BUCKET, "test-put-object", nil, api.ErrorOption)
	ExpectEqual(t.Errorf, nil, res)
	ExpectEqual(t.Errorf, bce.NewBceClientError("Handle bos client options failed: BosContext Options: error option"), err)

	//BasicGetObject
	res, err = client.BasicGetObject(EXISTS_BUCKET, "test-put-object")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, res.VersionId, "AKyQ9DRhhoY=")
	ExpectEqual(t.Errorf, res.CacheControl, "private")
	ExpectEqual(t.Errorf, res.ContentDisposition, "attachment; filename=\"download.txt\"")
	ExpectEqual(t.Errorf, res.ContentLength, int64(len(respBody)))
	ExpectEqual(t.Errorf, res.ContentType, "text/plain")
	ExpectEqual(t.Errorf, res.UserMeta["Key1"], "Value1")
	ExpectEqual(t.Errorf, res.UserMeta["Key2"], "Value2")
	ExpectEqual(t.Errorf, res.StorageClass, api.STORAGE_CLASS_ARCHIVE)
	n, err = res.Body.Read(buf)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, len(respBody), n)
	ExpectEqual(t.Errorf, respBody, string(buf))

	//BasicGetObjectToFile
	fname := "/tmp/test-get-object"
	err = client.BasicGetObjectToFile(EXISTS_BUCKET, "test-put-object", fname)
	ExpectEqual(t.Errorf, err, nil)
	err = client.GetObjectToFileWithContext(context.TODO(), EXISTS_BUCKET, "test-put-object", fname)
	ExpectEqual(t.Errorf, err, nil)
	file, err := os.Open(fname)
	ExpectEqual(t.Errorf, err, nil)
	n, err = file.Read(buf)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, len(respBody), n)
	ExpectEqual(t.Errorf, respBody, string(buf))
	os.Remove(fname)
}
func TestGetObjectMeta(t *testing.T) {
	roundTripperOpts := []util.MockRoundTripperOption{
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
			http.CanonicalHeaderKey(my_http.BCE_STORAGE_CLASS):                 api.STORAGE_CLASS_ARCHIVE,
			http.CanonicalHeaderKey(my_http.CONTENT_MD5):                       "Zh+ACfqOVqnQ6UoKZEOX1w==",
			http.CanonicalHeaderKey(my_http.LAST_MODIFIED):                     "Wed, 17 Dec 2025 06:25:34 GMT",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32):                 "1922069637",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC64ECMA):             "12759301844125077625",
		}),
	}
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)
	// case1: all is ok
	res, err := client.GetObjectMeta(EXISTS_BUCKET, "test-put-object")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, res.VersionId, "AKyQ9DRhhoY=")
	ExpectEqual(t.Errorf, res.CacheControl, "private")
	ExpectEqual(t.Errorf, res.ContentDisposition, "attachment; filename=\"download.txt\"")
	ExpectEqual(t.Errorf, res.ContentLength, 1234567)
	ExpectEqual(t.Errorf, res.ContentType, "application/octet-stream")
	ExpectEqual(t.Errorf, res.UserMeta["Key1"], "Value1")
	ExpectEqual(t.Errorf, res.UserMeta["Key2"], "Value2")
	ExpectEqual(t.Errorf, res.StorageClass, api.STORAGE_CLASS_ARCHIVE)
	ExpectEqual(t.Errorf, res.ContentMD5, "Zh+ACfqOVqnQ6UoKZEOX1w==")
	ExpectEqual(t.Errorf, res.LastModified, "Wed, 17 Dec 2025 06:25:34 GMT")
	ExpectEqual(t.Errorf, res.ContentCrc32, "1922069637")
	ExpectEqual(t.Errorf, res.ContentCrc64ECMA, "12759301844125077625")
	t.Logf("get object meta result: %+v", res)
	res, err = client.GetObjectMetaWithContext(context.TODO(), EXISTS_BUCKET, "test-put-object")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, res.VersionId, "AKyQ9DRhhoY=")
	ExpectEqual(t.Errorf, res.CacheControl, "private")
	ExpectEqual(t.Errorf, res.ContentDisposition, "attachment; filename=\"download.txt\"")
	ExpectEqual(t.Errorf, res.ContentLength, 1234567)
	ExpectEqual(t.Errorf, res.ContentType, "application/octet-stream")
	ExpectEqual(t.Errorf, res.UserMeta["Key1"], "Value1")
	ExpectEqual(t.Errorf, res.UserMeta["Key2"], "Value2")
	ExpectEqual(t.Errorf, res.StorageClass, api.STORAGE_CLASS_ARCHIVE)
	ExpectEqual(t.Errorf, res.ContentMD5, "Zh+ACfqOVqnQ6UoKZEOX1w==")
	ExpectEqual(t.Errorf, res.LastModified, "Wed, 17 Dec 2025 06:25:34 GMT")
	ExpectEqual(t.Errorf, res.ContentCrc32, "1922069637")
	ExpectEqual(t.Errorf, res.ContentCrc64ECMA, "12759301844125077625")
	// case2: handle options err
	res, err = client.GetObjectMeta(EXISTS_BUCKET, "test-put-object", api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, nil, res)
	res, err = client.GetObjectMetaWithContext(context.TODO(), EXISTS_BUCKET, "test-put-object", api.ErrorOption)
	ExpectEqual(t.Errorf, optionError, err)
	ExpectEqual(t.Errorf, nil, res)
}
func TestSelectObject(t *testing.T) {
	// mock bos client
	respBody := `<Records message>
	……
	<Continuation Message>
	……
	<Records message>
	<Continuation Message>
	<End message>  `
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	args := &api.SelectObjectArgs{}
	reqBody := `{
		"selectRequest": {
			"expression": "c2VsZWN0IGNvdW50KCopIGZyb20gbxkl2JqZWN0IHdoZXJlIF80ID4gNDU=",
			"expressionType": "SQL",
			"inputSerialization": {
				"compressionType": "NONE",
				"json": {
					"type": "DOCUMENT"
				}
			},
			"outputSerialization": {
				"json": {
					"recordDelimiter": "Cg=="
				}
			},
			"requestProgress": {
				"enabled": false
			}
		}
	}`
	ExpectEqual(t.Errorf, nil, json.Unmarshal([]byte(reqBody), args))
	res, err := client.SelectObject(EXISTS_BUCKET, "test-object", args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Body != nil)
	res, err = client.SelectObjectWithContext(context.Background(), EXISTS_BUCKET, "test-object", args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, res.Body != nil)
}
func TestFetchObject(t *testing.T) {
	asyncResp := `{
		"code": "success",
		"message": "success",
		"requestId": "4db2b34d-654d-4d8a-b49b-3049ca786409",
		"jobId": "fo-9a98f238d56a33b4eb6664fede20b747"
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, asyncResp)
	ExpectEqual(t.Errorf, nil, err)
	// async fetch object
	args := &api.FetchObjectArgs{
		FetchMode:    api.FETCH_MODE_ASYNC,
		StorageClass: api.STORAGE_CLASS_COLD,
	}
	res, err := client.FetchObject(EXISTS_BUCKET, "test-fetch-object", "https://cloud.baidu.com/doc/BOS/API.html", args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Code, "success")
	ExpectEqual(t.Errorf, res.Message, "success")
	ExpectEqual(t.Errorf, res.RequestId, "4db2b34d-654d-4d8a-b49b-3049ca786409")
	ExpectEqual(t.Errorf, res.JobId, "fo-9a98f238d56a33b4eb6664fede20b747")
	t.Logf("result: %+v", res)
	res, err = client.FetchObjectWithContext(context.Background(), EXISTS_BUCKET, "test-fetch-object", "https://cloud.baidu.com/doc/BOS/API.html", args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Code, "success")
	ExpectEqual(t.Errorf, res.Message, "success")
	ExpectEqual(t.Errorf, res.RequestId, "4db2b34d-654d-4d8a-b49b-3049ca786409")
	ExpectEqual(t.Errorf, res.JobId, "fo-9a98f238d56a33b4eb6664fede20b747")
	t.Logf("result: %+v", res)

	//BasicFetchObject
	res, err = client.BasicFetchObject(EXISTS_BUCKET, "test-fetch-object", "https://bj.bcebos.com/gosdk-test-bucket/testsumlink")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Code, "success")
	ExpectEqual(t.Errorf, res.Message, "success")
	ExpectEqual(t.Errorf, res.RequestId, "4db2b34d-654d-4d8a-b49b-3049ca786409")
	ExpectEqual(t.Errorf, res.JobId, "fo-9a98f238d56a33b4eb6664fede20b747")

	//SimpleFetchObject
	res, err = client.SimpleFetchObject(EXISTS_BUCKET, "test-fetch-object", "https://bj.bcebos.com/gosdk-test-bucket/testsumlink",
		api.FETCH_MODE_ASYNC, api.STORAGE_CLASS_COLD)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Code, "success")
	ExpectEqual(t.Errorf, res.Message, "success")
	ExpectEqual(t.Errorf, res.RequestId, "4db2b34d-654d-4d8a-b49b-3049ca786409")
	ExpectEqual(t.Errorf, res.JobId, "fo-9a98f238d56a33b4eb6664fede20b747")
}
func TestAppendObject(t *testing.T) {
	// mock bos client
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.BCE_NEXT_APPEND_OFFSET): "12345",
			http.CanonicalHeaderKey(my_http.CONTENT_MD5):            "Zh+ACfqOVqnQ6UoKZEOX1w==",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32):      "1922069637",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32C):     "43574823532456",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC64ECMA):  "12759301844125077625",
		}),
	}
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", options...)
	ExpectEqual(t.Errorf, nil, err)

	//AppendObject
	args := &api.AppendObjectArgs{}
	body, _ := bce.NewBodyFromString("aaaaaaaaaaa")
	res, err := client.AppendObject(EXISTS_BUCKET, "test-append-object", body, args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, int64(12345), res.NextAppendOffset)
	ExpectEqual(t.Errorf, res.ContentMD5, "Zh+ACfqOVqnQ6UoKZEOX1w==")
	ExpectEqual(t.Errorf, res.ContentCrc32, "1922069637")
	ExpectEqual(t.Errorf, res.ContentCrc32c, "43574823532456")
	ExpectEqual(t.Errorf, res.ContentCrc64ECMA, "12759301844125077625")
	t.Logf("%+v", res)
	body01, _ := bce.NewBodyFromString("aaaaaaaaaaa")
	res, err = client.AppendObjectWithContext(context.Background(), EXISTS_BUCKET, "test-append-object", body01, args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, int64(12345), res.NextAppendOffset)

	//SimpleAppendObject
	body1, _ := bce.NewBodyFromString("bbbbbbbbbbb")
	res, err = client.SimpleAppendObject(EXISTS_BUCKET, "test-append-object", body1, 11)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, int64(12345), res.NextAppendOffset)

	//SimpleAppendObjectFromString
	res, err = client.SimpleAppendObjectFromString(EXISTS_BUCKET, "test-append-object", "123", 22)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, int64(12345), res.NextAppendOffset)

	//SimpleAppendObjectFromFile
	fname := "/tmp/test-append-file"
	f, _ := os.Create(fname)
	f.WriteString("12345")
	f.Close()
	res, err = client.SimpleAppendObjectFromFile(EXISTS_BUCKET, "test-append-object", fname, 25)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, int64(12345), res.NextAppendOffset)
	os.Remove(fname)
	_, err = client.SimpleAppendObjectFromFile(EXISTS_BUCKET, "test-append-object", fname, 25)
	ExpectEqual(t.Errorf, err.Error(), fmt.Errorf("open %s: no such file or directory", fname).Error())
}
func TestDeleteObject(t *testing.T) {
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteObject(EXISTS_BUCKET, "test-put-object")
	ExpectEqual(t.Errorf, err, nil)
	err = client.DeleteObjectVersion(EXISTS_BUCKET, "test-put-object", "version_id")
	ExpectEqual(t.Errorf, err, nil)
}
func TestDeleteMultipleObjects(t *testing.T) {
	respBody := `  {
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
	} `
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	//DeleteMultipleObjectsFromString
	multiDeleteStr := `{
		"objects":[
			{"key": "aaaa"},
			{"key": "test-copy-object"},
			{"key": "test-append-object"},
			{"key": "cccc"}
		]
	}`
	res, err := client.DeleteMultipleObjectsFromString(EXISTS_BUCKET, multiDeleteStr)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 2, len(res.Errors))
	ExpectEqual(t.Errorf, "my-object1", res.Errors[0].Key)
	ExpectEqual(t.Errorf, "NoSuchKey", res.Errors[0].Code)
	ExpectEqual(t.Errorf, "The specified key does not exist.", res.Errors[0].Message)
	ExpectEqual(t.Errorf, "my-object2", res.Errors[1].Key)
	ExpectEqual(t.Errorf, "InvalidArgument", res.Errors[1].Code)
	ExpectEqual(t.Errorf, "Invalid Argument.", res.Errors[1].Message)
	t.Logf("%+v", res)

	//DeleteMultipleObjectsFromStruct
	multiDeleteObj := &api.DeleteMultipleObjectsArgs{
		Objects: []api.DeleteObjectArgs{
			{Key: "1"}, {Key: "test-fetch-object"},
		},
	}
	res, err = client.DeleteMultipleObjectsFromStruct(EXISTS_BUCKET, multiDeleteObj)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 2, len(res.Errors))

	//DeleteMultipleObjectsFromKeyList
	keyList := []string{"aaaa", "test-copy-object", "test-append-object", "cccc"}
	res, err = client.DeleteMultipleObjectsFromKeyList(EXISTS_BUCKET, keyList)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 2, len(res.Errors))

	//DeleteMultipleObjects
	body, err := bce.NewBodyFromString(multiDeleteStr)
	ExpectEqual(t.Errorf, err, nil)
	res, err = client.DeleteMultipleObjects(EXISTS_BUCKET, body)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 2, len(res.Errors))
}
func TestInitiateMultipartUpload(t *testing.T) {
	// mock bos client
	respBody := `{
		"bucket": "BucketName",
		"key":"ObjectName",
		"uploadId": "a44cc9bab11cbd156984767aad637851"
	}`
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	args := &api.InitiateMultipartUploadArgs{Expires: "aaaaaaa"}
	res, err := client.InitiateMultipartUpload(EXISTS_BUCKET, "test-multipart-upload", "", args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Bucket, "BucketName")
	ExpectEqual(t.Errorf, res.Key, "ObjectName")
	ExpectEqual(t.Errorf, res.UploadId, "a44cc9bab11cbd156984767aad637851")
	t.Logf("%+v", res)

	//BasicInitiateMultipartUpload
	res, err = client.BasicInitiateMultipartUpload(EXISTS_BUCKET, "test-multipart-upload")
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Bucket, "BucketName")
	ExpectEqual(t.Errorf, res.Key, "ObjectName")
	ExpectEqual(t.Errorf, res.UploadId, "a44cc9bab11cbd156984767aad637851")
}
func TestUploadPart(t *testing.T) {
	//mock bos clientF
	roundTripperOpts := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "827ccb0eea8a706c4c34a16891f84e7b",
		}),
	}
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "", roundTripperOpts...)
	ExpectEqual(t.Errorf, nil, err)
	// body is nil
	res, err := client.UploadPart(EXISTS_BUCKET, "a", "b", 1, nil, nil)
	ExpectEqual(t.Errorf, err, bce.NewBceClientError("upload part content should not be empty"))
	ExpectEqual(t.Errorf, res, "")
	res, err = client.UploadPart(EXISTS_BUCKET, "a", "b", 1, nil, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, res, "")
	// body is valid
	content := "this is a body string for testing upload part"
	body, err := bce.NewBodyFromString(content)
	ExpectEqual(t.Errorf, err, nil)
	res, err = client.UploadPart(EXISTS_BUCKET, "a", "b", 1, body, nil)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, "827ccb0eea8a706c4c34a16891f84e7b")
	t.Logf("%+v, %+v", res, err)
	body1, err := bce.NewBodyFromString(content)
	ExpectEqual(t.Errorf, err, nil)
	res, err = client.UploadPartWithContext(context.TODO(), EXISTS_BUCKET, "a", "b", 1, body1, nil)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, "827ccb0eea8a706c4c34a16891f84e7b")
	res, err = client.UploadPartWithContext(context.TODO(), EXISTS_BUCKET, "a", "b", 1, nil, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, res, "")

	//BasicUploadPart
	body2, err := bce.NewBodyFromString(content)
	ExpectEqual(t.Errorf, err, nil)
	res, err = client.BasicUploadPart(EXISTS_BUCKET, "a", "b", 1, body2, nil)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, "827ccb0eea8a706c4c34a16891f84e7b")
	res, err = client.BasicUploadPart(EXISTS_BUCKET, "a", "b", 1, body2, nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, res, "")

	//UploadPartFromSectionFile
	fname := "/tmp/test-put-object-acl-by-file"
	f, _ := os.Create(fname)
	f.WriteString(content)
	f.Close()
	f, err = os.Open(fname)
	ExpectEqual(t.Errorf, err, nil)
	res, err = client.UploadPartFromSectionFile(EXISTS_BUCKET, "a", "b", 1, f, 0, int64(len(content)), nil)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, "827ccb0eea8a706c4c34a16891f84e7b")
	f.Seek(0, io.SeekStart)
	res, err = client.UploadPartFromSectionFileWithContext(context.TODO(),
		EXISTS_BUCKET, "a", "b", 1, f, 0, int64(len(content)), nil)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, "827ccb0eea8a706c4c34a16891f84e7b")
	res, err = client.UploadPartFromSectionFile(EXISTS_BUCKET, "a", "b", 1, f, 0, int64(len(content)), nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, res, "")
	res, err = client.UploadPartFromSectionFileWithContext(context.TODO(), EXISTS_BUCKET,
		"a", "b", 1, f, 0, int64(len(content)), nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, res, "")
	f.Close()
	os.Remove(fname)

	//UploadPartFromBytes
	res, err = client.UploadPartFromBytes(EXISTS_BUCKET, "a", "b", 1, []byte(content), nil)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, "827ccb0eea8a706c4c34a16891f84e7b")
	res, err = client.UploadPartFromBytesWithContext(context.TODO(), EXISTS_BUCKET, "a", "b", 1, []byte(content), nil)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, "827ccb0eea8a706c4c34a16891f84e7b")
	res, err = client.UploadPartFromBytes(EXISTS_BUCKET, "a", "b", 1, []byte(content), nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, res, "")
	res, err = client.UploadPartFromBytesWithContext(context.TODO(), EXISTS_BUCKET,
		"a", "b", 1, []byte(content), nil, api.ErrorOption)
	ExpectEqual(t.Errorf, err, optionError)
	ExpectEqual(t.Errorf, res, "")
}
func TestUploadPartCopy(t *testing.T) {
	respBody := `{   
		"lastModified":"2016-05-12T09:14:32Z",
		"eTag":"67b92a7c2a9b9c1809a6ae3295dcc127"
	}`
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	// response ok
	res, err := client.UploadPartCopy(EXISTS_BUCKET, "test-multipart-upload",
		EXISTS_BUCKET, "test-multipart-copy", "12345", 1, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, res.ETag, "67b92a7c2a9b9c1809a6ae3295dcc127")
	ExpectEqual(t.Errorf, res.LastModified, "2016-05-12T09:14:32Z")
	t.Logf("%+v, %+v", res, err)
	res, err = client.UploadPartCopyWithContext(context.TODO(), EXISTS_BUCKET, "test-multipart-upload",
		EXISTS_BUCKET, "test-multipart-copy", "12345", 1, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, res.ETag, "67b92a7c2a9b9c1809a6ae3295dcc127")
	ExpectEqual(t.Errorf, res.LastModified, "2016-05-12T09:14:32Z")
	//BasicUploadPartCopy
	res, err = client.BasicUploadPartCopy(EXISTS_BUCKET, "test-multipart-upload",
		EXISTS_BUCKET, "test-multipart-copy", "12345", 1)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, res.ETag, "67b92a7c2a9b9c1809a6ae3295dcc127")
	ExpectEqual(t.Errorf, res.LastModified, "2016-05-12T09:14:32Z")
	// response error
	respBody1 := `{
		"code":"InternalError",
		"message":"We encountered an internal error. Please try again.",
		"requestId":"52454655-5345-4420-4259-204e47494e58"
	}`
	client1, err := NewMockBosClient(ak, sk, endpoint, respBody1)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client1.UploadPartCopy(EXISTS_BUCKET, "test-multipart-upload",
		EXISTS_BUCKET, "test-multipart-copy", "12345", 1, nil)
	ExpectEqual(t.Errorf, bce.NewBceServiceError("InternalError", "We encountered an internal error. Please try again.",
		"52454655-5345-4420-4259-204e47494e58", 500), err)
	ExpectEqual(t.Errorf, res, nil)
}
func TestCompleteMultipartUpload(t *testing.T) {
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
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	body, err := bce.NewBodyFromStringV2(reqBody, false)
	ExpectEqual(t.Errorf, nil, err)
	uploadId := "3858f62230ac3c915f300c664312c11f"
	res, err := client.CompleteMultipartUpload(EXISTS_BUCKET, "test-object", uploadId, body, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "BucketName", res.Bucket)
	args := &api.CompleteMultipartUploadArgs{}
	ExpectEqual(t.Errorf, nil, json.Unmarshal([]byte(reqBody), args))
	res, err = client.CompleteMultipartUploadFromStruct(EXISTS_BUCKET, "test-object", uploadId, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "BucketName", res.Bucket)
}
func TestAbortMultipartUpload(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	uploadId := "3858f62230ac3c915f300c664312c11f"
	ExpectEqual(t.Errorf, nil, client.AbortMultipartUpload(EXISTS_BUCKET, EXISTS_OBJECT, uploadId))
}
func TestListParts(t *testing.T) {
	respBody := `{
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
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	uploadId := "a44cc9bab11bdc157676984aad851637"
	args := &api.ListPartsArgs{
		MaxParts:         10,
		PartNumberMarker: "part-number-marker",
	}
	res, err := client.ListParts(EXISTS_BUCKET, EXISTS_OBJECT, uploadId, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Parts))
	res, err = client.ListPartsWithContext(context.Background(), EXISTS_BUCKET, EXISTS_OBJECT, uploadId, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Parts))
	res, err = client.BasicListParts(EXISTS_BUCKET, EXISTS_OBJECT, uploadId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res.Parts))
}
func TestListMultipartUploads(t *testing.T) {
	respBody := `{
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
			},
			{
				"key": "my-movie.m2ts",
				"uploadId": "b41cc9aad11cbd637851767bab156984",
				"owner": {
					"id": "b1d16700c70b0b05597d7acd6a3f92be",
					"displayName": "OwnerDisplayName"
				},
				"initiated": "2010-11-10T10:49:33Z",
				"storageClass": "STANDARD_IA"
			}
		]
	}`
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	//ListMultipartUploads
	args := &api.ListMultipartUploadsArgs{MaxUploads: 10}
	res, err := client.ListMultipartUploads(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Bucket, "bucket")
	ExpectEqual(t.Errorf, res.NextKeyMarker, "my-movie.m2ts")
	ExpectEqual(t.Errorf, res.MaxUploads, 3)
	ExpectEqual(t.Errorf, res.IsTruncated, true)
	ExpectEqual(t.Errorf, len(res.Uploads), 3)
	t.Logf("%+v", res)

	res, err = client.ListMultipartUploadsWithContext(context.Background(), EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res.Bucket, "bucket")
	ExpectEqual(t.Errorf, res.NextKeyMarker, "my-movie.m2ts")

	//BasicListMultipartUploads
	res, err = client.BasicListMultipartUploads(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, len(res.Uploads), 3)
	ExpectEqual(t.Errorf, res.Uploads[0].Key, "my-divisor")
	ExpectEqual(t.Errorf, res.Uploads[0].UploadId, "a44cc9bab11bdc157676984aad851637")
	ExpectEqual(t.Errorf, res.Uploads[0].Initiated, "2010-11-10T20:48:33Z")
	ExpectEqual(t.Errorf, res.Uploads[0].StorageClass, "STANDARD_IA")
	ExpectEqual(t.Errorf, res.Uploads[0].Owner.Id, "75aa57f09aa0c8caeab4aeebf76c078efc7c6caea54ba06a")
	ExpectEqual(t.Errorf, res.Uploads[0].Owner.DisplayName, "OwnerDisplayName")
	t.Logf("%+v", res)
}
func TestUploadSuperFile(t *testing.T) {
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "b54357faf0632cce46e942fa68356b38",
		}),
	}
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	// case1: not exist
	err = client.UploadSuperFile(EXISTS_BUCKET, "super-object", "test-object", "")
	ExpectEqual(t.Errorf, err.Error(), "open test-object: no such file or directory")
	// case2: small file
	fileName := "/tmp/test-upload-super-file-" + strconv.FormatInt(time.Now().UnixMicro(), 10)
	file, err := os.Create(fileName)
	ExpectEqual(t.Errorf, nil, err)
	file.WriteString("test content")
	file.Close()
	err = client.UploadSuperFile(EXISTS_BUCKET, "super-object", fileName, "")
	ExpectEqual(t.Errorf, err.Error(), "multipart size should not be less than 1MB")
	err = client.UploadSuperFile(EXISTS_BUCKET, "super-object", currentFileName(), "")
	ExpectEqual(t.Errorf, err.Error(), "EOF")
	t.Logf("%+v", err)
	os.Remove(fileName)
	// case3: ok
	respBody1 := `{ "bucket": "BucketName", "key":"ObjectName", "uploadId": "a44cc9bab11cbd156984767aad637851" }`
	respBody2 := `{
		"location":"http://bj.bcebos.com/BucketName/ObjectName",
		"bucket":"BucketName",
		"key":"object",
		"eTag":"3858f62230ac3c915f300c664312c11f"
	}`
	options3 := append(options, util.AppendRespBody([]string{respBody1, respBody2}))
	client3, err := NewMockBosClient(ak, sk, endpoint, "", options3...)
	ExpectEqual(t.Errorf, nil, err)
	err = client3.UploadSuperFile(EXISTS_BUCKET, "super-object", currentFileName(), "")
	ExpectEqual(t.Errorf, nil, err)

	options4 := append(options, util.AppendRespBody([]string{respBody1, ""}))
	client4, err := NewMockBosClient(ak, sk, endpoint, "", options4...)
	ExpectEqual(t.Errorf, nil, err)
	err = client4.UploadSuperFile(EXISTS_BUCKET, "super-object", currentFileName(), "")
	ExpectEqual(t.Errorf, "EOF", err.Error())
}
func TestDownloadSuperFile(t *testing.T) {
	//mock bos client
	options := []util.MockRoundTripperOption{
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
			http.CanonicalHeaderKey(my_http.BCE_STORAGE_CLASS):                 api.STORAGE_CLASS_ARCHIVE,
			http.CanonicalHeaderKey(my_http.CONTENT_MD5):                       "Zh+ACfqOVqnQ6UoKZEOX1w==",
			http.CanonicalHeaderKey(my_http.LAST_MODIFIED):                     "Wed, 17 Dec 2025 06:25:34 GMT",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32):                 "1922069637",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC64ECMA):             "12759301844125077625",
		}),
	}
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	options1 := append(options, util.AppendRespBody([]string{"test data"}))
	client1, err := NewMockBosClient(ak, sk, endpoint, "", options1...)
	ExpectEqual(t.Errorf, nil, err)
	err = client1.DownloadSuperFile(EXISTS_BUCKET, "super-object", "/dev/null")
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", err)
}

func TestGeneratePresignedUrl(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	//BasicGeneratePresignedUrl
	url1 := client.BasicGeneratePresignedUrl(EXISTS_BUCKET, EXISTS_OBJECT, 100)
	rawUrl1, err := url.Parse(url1)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, EXISTS_BUCKET+"."+client.Config.Endpoint, rawUrl1.Host)
	//params
	params := map[string]string{"responseContentType": "text"}
	url2 := client.GeneratePresignedUrl(EXISTS_BUCKET, EXISTS_OBJECT, 1000, "HEAD", nil, params)
	rawUrl2, err := url.Parse(url2)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, "text", rawUrl2.Query().Get("responseContentType"))
	//change endpoint
	client.Config.Endpoint = "10.180.112.31"
	url3 := client.GeneratePresignedUrl(EXISTS_BUCKET, EXISTS_OBJECT, 100, "HEAD", nil, params)
	rawUrl3, err := url.Parse(url3)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, client.Config.Endpoint, rawUrl3.Host)
	//change endpoint
	client.Config.Endpoint = "10.180.112.31:80"
	url4 := client.GeneratePresignedUrl(EXISTS_BUCKET, EXISTS_OBJECT, 100, "HEAD", nil, params)
	rawUrl4, err := url.Parse(url4)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, client.Config.Endpoint, rawUrl4.Host)

	url5 := client.GeneratePresignedUrlPathStyle(EXISTS_BUCKET, EXISTS_OBJECT, 1000, "", nil, nil)
	rawUrl5, err := url.Parse(url5)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, client.Config.Endpoint, rawUrl5.Host)
}

func TestPutObjectAcl(t *testing.T) {
	acl := `{
		"accessControlList":[
			{
				"grantee":[{
					"id":"e13b12d0131b4c8bae959df4969387b8"
				}],
				"permission":["READ"]
			}
		]
	}`

	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	body, _ := bce.NewBodyFromString(acl)
	err = client.PutObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT, body)
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutObjectAclFromCanned(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	err = client.PutObjectAclFromCanned(EXISTS_BUCKET, EXISTS_OBJECT, api.CANNED_ACL_PUBLIC_READ)
	ExpectEqual(t.Errorf, err, nil)
}

func TestPutObjectAclGrantRead(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	err = client.PutObjectAclGrantRead(EXISTS_BUCKET,
		EXISTS_OBJECT, "e13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutObjectAclGrantFullControl(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	err = client.PutObjectAclGrantFullControl(EXISTS_BUCKET,
		EXISTS_OBJECT, "e13b12d0131b4c8bae959df4969387b8")
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutObjectAclFromFile(t *testing.T) {
	acl := `{
        "accessControlList":[
            {
                "grantee":[{
                    "id":"e13b12d0131b4c8bae959df4969387b8"
                }],
                "permission":["READ"]
            }
        ]
    }`
	fname := "/tmp/test-put-object-acl-by-file"
	f, _ := os.Create(fname)
	f.WriteString(acl)
	f.Close()
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	err = client.PutObjectAclFromFile(EXISTS_BUCKET, EXISTS_OBJECT, fname)
	os.Remove(fname)
	ExpectEqual(t.Errorf, err, nil)

	err = client.PutObjectAclFromFile(EXISTS_BUCKET, EXISTS_OBJECT, fname)
	ExpectEqual(t.Errorf, err.Error(), fmt.Sprintf("open %s: no such file or directory", fname))
}
func TestPutObjectAclFromString(t *testing.T) {
	acl := `{
    "accessControlList":[
        {
            "grantee":[{
                "id":"e13b12d0131b4c8bae959df4969387b8"
            }],
            "permission":["FULL_CONTROL"]
        }
    ]}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	err = client.PutObjectAclFromString(EXISTS_BUCKET, EXISTS_OBJECT, acl)
	ExpectEqual(t.Errorf, err, nil)
}
func TestPutObjectAclFromStruct(t *testing.T) {
	aclObj := &api.PutObjectAclArgs{
		AccessControlList: []api.GrantType{
			{
				Grantee: []api.GranteeType{
					{Id: "e13b12d0131b4c8bae959df4969387b8"},
				},
				Permission: []string{
					"READ",
				},
			},
		},
	}
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "192.168.1.1"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	err = client.PutObjectAclFromStruct(EXISTS_BUCKET, EXISTS_OBJECT, aclObj)
	ExpectEqual(t.Errorf, err, nil)
}
func TestGetObjectAcl(t *testing.T) {
	respBody := `{
		"accessControlList":[
			{
				"grantee":[{
					"id":"e13b12d0131b4c8bae959df4969387b8"
				}],
				"permission":["FULL_CONTROL"]
			},
			{
				"grantee":[{
					"id":"8c47a952db4444c5a097b41be3f24c94"
				}],
				"permission":["READ"]
			}
		]
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 2, len(res.AccessControlList))
}
func TestDeleteObjectAcl(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteObjectAcl(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
}
func TestRestoreObject(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	// case1: invalid restore tier
	err = client.RestoreObject(EXISTS_BUCKET, EXISTS_OBJECT, 3, "restore_tier")
	ExpectEqual(t.Errorf, errors.New("invalid restore tier"), err)
	//case2: invalid restore days
	err = client.RestoreObject(EXISTS_BUCKET, EXISTS_OBJECT, 0, api.RESTORE_TIER_EXPEDITED)
	ExpectEqual(t.Errorf, errors.New("invalid restore days"), err)
	// case3: ok
	err = client.RestoreObject(EXISTS_BUCKET, EXISTS_OBJECT, 3, api.RESTORE_TIER_EXPEDITED)
	ExpectEqual(t.Errorf, err, nil)
}
func TestBucketTrashPut(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	args := api.PutBucketTrashReq{
		TrashDir: ".trash",
	}
	err = client.PutBucketTrash(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
}
func TestBucketTrashGet(t *testing.T) {
	respBody := `{
		"trashDir": ".trash/"
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketTrash(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, res.TrashDir, ".trash/")
	t.Logf("%v, %v", res, err)
}
func TestBucketTrashDelete(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketTrash(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}
func TestBucketNotificationPut(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	//PutBucketNotification
	notificationReq := api.PutBucketNotificationReq{
		Notifications: []api.PutBucketNotificationSt{
			{
				Id:        "water-test-1",
				Name:      "water-rule-1",
				AppId:     "water-app-id-1",
				Status:    "enabled",
				Resources: []string{EXISTS_BUCKET + "/path1", "/path2"},
				Events:    []string{"PutObject"},
				Apps: []api.PutBucketNotificationAppsSt{
					{
						Id:       "app-id-1",
						EventUrl: "http://xxx.com/event",
						XVars:    "",
					},
				},
			},
		},
	}
	err = client.PutBucketNotification(EXISTS_BUCKET, notificationReq)
	ExpectEqual(t.Errorf, err, nil)
}
func TestBucketNotificationGet(t *testing.T) {
	respBody := `{
		"notifications": [
			{
				"id": "notify-id-1",
				"name": "rule-name",
				"appId": "app-id-1",
				"status": "enabled",
				"resources": [
					"bucket-a/path1", "/path2", "/path3/*.jpg", "/path4/*"
				],
				"events": [
					"PutObject"
				],
				"apps": [
					{
						"id": "app-id-1",
						"eventUrl": "http://xxx.com/event",
						"xVars": ""
					},
					{
						"id": "app-id-2",
						"eventUrl": "brn:bce:cfc:bj:1f1c3e383c31e6467c4c44523f0d5b22:function:hello_bos:$LATEST"
					},
					{
						"id": "app-id-3",
						"eventUrl": "app:ImageOcr",
						"xVars": "{\"saveUrl\": \"http://xxx.com/ocr\"}"
					}
				]
			}
		]
	}`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketNotification(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, len(res.Notifications), 1)
	ExpectEqual(t.Errorf, res.Notifications[0].AppId, "app-id-1")
	ExpectEqual(t.Errorf, len(res.Notifications[0].Apps), 3)
	ExpectEqual(t.Errorf, res.Notifications[0].Id, "notify-id-1")
	ExpectEqual(t.Errorf, res.Notifications[0].Name, "rule-name")
	ExpectEqual(t.Errorf, res.Notifications[0].Status, "enabled")
	ExpectEqual(t.Errorf, len(res.Notifications[0].Resources), 4)
	ExpectEqual(t.Errorf, len(res.Notifications[0].Events), 1)
}
func TestBucketNotificationDelete(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketNotification(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}

func TestParallelUpload(t *testing.T) {
	// mock bos client
	respBody1 := `{ "bucket": "BucketName", "key":"ObjectName", "uploadId": "a44cc9bab11cbd156984767aad637851" }`
	respBody3 := `{
		"location":"http://bj.bcebos.com/BucketName/ObjectName",
		"bucket":"BucketName",
		"key":"object",
		"eTag":"3858f62230ac3c915f300c664312c11f"
	}`

	options := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "b54357faf0632cce46e942fa68356b38",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32C): "723497213897532",
		}),
	}
	args := &api.InitiateMultipartUploadArgs{
		ObjectExpires: 3,
		IfMatch:       "if-match",
		IfNoneMatch:   "IfNoneMatch",
	}

	// case1: ok
	options1 := append(options, util.AppendRespBody([]string{respBody1, respBody3}))
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client1, err := NewMockBosClient(ak, sk, endpoint, "", options1...)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client1.ParallelUpload(EXISTS_BUCKET, "test_multiupload", currentFileName(), "", args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "BucketName", res.Bucket)

	// case2: init error
	options2 := append(options, util.AppendRespBody([]string{""}))
	client2, err := NewMockBosClient(ak, sk, endpoint, "", options2...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client2.ParallelUpload(EXISTS_BUCKET, "test_multiupload", currentFileName(), "", args)
	ExpectEqual(t.Errorf, "EOF", err.Error())
	ExpectEqual(t.Errorf, nil, res)

	// case3: ParallelUpload error
	options3 := append(options, util.AppendRespBody([]string{respBody1, ""}))
	client3, err := NewMockBosClient(ak, sk, endpoint, "", options3...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client3.ParallelUpload(EXISTS_BUCKET, "test_multiupload", "fileName", "", args)
	ExpectEqual(t.Errorf, nil, res)
	ExpectEqual(t.Errorf, fmt.Sprintf("open %s: no such file or directory", "fileName"), err.Error())

	// case4: complete error
	options4 := append(options, util.AppendRespBody([]string{respBody1, ""}))
	client4, err := NewMockBosClient(ak, sk, endpoint, "", options4...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client4.ParallelUpload(EXISTS_BUCKET, "test_multiupload", currentFileName(), "", args)
	ExpectEqual(t.Errorf, "EOF", err.Error())
	ExpectEqual(t.Errorf, nil, res)

	// case5: parallelPartUpload
	options5 := append(options, util.AppendRespBody([]string{""}))
	client5, err := NewMockBosClient(ak, sk, endpoint, "", options5...)
	ExpectEqual(t.Errorf, nil, err)
	uploadId := "a44cc9bab11cbd156984767aad637851"
	uploadInfo, err := client5.parallelPartUpload(EXISTS_BUCKET, EXISTS_OBJECT, currentFileName(), uploadId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(uploadInfo))

	// case6:
	options6 := append(options, util.AppendRespBody([]string{""}))
	client6, err := NewMockBosClient(ak, sk, endpoint, "", options6...)
	ExpectEqual(t.Errorf, nil, err)
	content6, err := bce.NewBodyFromStringV2("str string", false)
	ExpectEqual(t.Errorf, nil, err)
	parallelChan := make(chan int, 1)
	errChan := make(chan error, 1)
	resultChan := make(chan api.UploadInfoType, 1)
	parallelChan <- 1
	client6.singlePartUpload(EXISTS_BUCKET, EXISTS_OBJECT, uploadId, 1, content6, parallelChan, errChan, resultChan)
	result7 := <-resultChan
	ExpectEqual(t.Errorf, "b54357faf0632cce46e942fa68356b38", result7.ETag)
}

func TestParallelUpload_CompleteOk(t *testing.T) {
	// mock bos client
	respBody := []string{`{
		"bucket": "BucketName",
		"key":"ObjectName",
		"uploadId": "a44cc9bab11cbd156984767aad637851"
	    }`,
		`{
			"location":"http://bj.bcebos.com/BucketName/ObjectName",
			"bucket":"BucketName",
			"key":"object",
			"eTag":"3858f62230ac3c915f300c664312c11f"
		}`,
	}
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AppendRespBody(respBody),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "b54357faf0632cce46e942fa68356b38",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32C): "723497213897532",
		}),
	}
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "", options...)
	ExpectEqual(t.Errorf, nil, err)

	// CompleteMultipartUploadFromStruct ok
	args := &api.InitiateMultipartUploadArgs{
		StorageClass: api.STORAGE_CLASS_ARCHIVE,
		IfMatch:      "string",
		IfNoneMatch:  "*",
	}
	args.ObjectExpires = 3
	args.IfMatch = "dskfhsad"
	args.IfNoneMatch = "asfiewhui"
	res, err := client.ParallelUpload(EXISTS_BUCKET, "test_multiupload", currentFileName(), "", args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "3858f62230ac3c915f300c664312c11f", res.ETag)
}

func TestParallelUpload_CompleteError(t *testing.T) {
	// mock bos client
	respBody := []string{`{
		"bucket": "BucketName",
		"key":"ObjectName",
		"uploadId": "a44cc9bab11cbd156984767aad637851"
	    }`,
		`{"loc}`,
	}
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
		util.AppendRespBody(respBody),
		util.AddHeaders(map[string]string{
			my_http.ETAG: "b54357faf0632cce46e942fa68356b38",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32C): "723497213897532",
		}),
	}
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "", options...)
	ExpectEqual(t.Errorf, nil, err)

	// CompleteMultipartUploadFromStruct ok
	args := &api.InitiateMultipartUploadArgs{
		StorageClass: api.STORAGE_CLASS_ARCHIVE,
	}
	args.ObjectExpires = 3
	args.IfMatch = "dskfhsad"
	args.IfNoneMatch = "asfiewhui"
	res, err := client.ParallelUpload(EXISTS_BUCKET, "test_multiupload", currentFileName(), "", args)
	ExpectEqual(t.Errorf, nil, res)
	ExpectEqual(t.Errorf, "unexpected EOF", err.Error())
}

func TestParallelCopy(t *testing.T) {
	respBody1 := `{ "bucket": "BucketName", "key":"ObjectName", "uploadId": "a44cc9bab11cbd156984767aad637851" }`
	respBody2 := `{ "lastModified":"2016-05-12T09:14:32Z", "eTag":"67b92a7c2a9b9c1809a6ae3295dcc127" }`
	respBody3 := `{
		"location":"http://bj.bcebos.com/BucketName/ObjectName",
		"bucket":"BucketName",
		"key":"object",
		"eTag":"3858f62230ac3c915f300c664312c11f"
	}`
	// mock bos client
	options := []util.MockRoundTripperOption{
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
			http.CanonicalHeaderKey(my_http.BCE_STORAGE_CLASS):                 api.STORAGE_CLASS_ARCHIVE,
			http.CanonicalHeaderKey(my_http.CONTENT_MD5):                       "Zh+ACfqOVqnQ6UoKZEOX1w==",
			http.CanonicalHeaderKey(my_http.LAST_MODIFIED):                     "Wed, 17 Dec 2025 06:25:34 GMT",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC32):                 "1922069637",
			http.CanonicalHeaderKey(my_http.BCE_CONTENT_CRC64ECMA):             "12759301844125077625",
		}),
	}
	// case1: ok
	options1 := append(options, util.AppendRespBody([]string{"", respBody1, respBody2, respBody3}))
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "", options1...)
	ExpectEqual(t.Errorf, nil, err)
	args := api.MultiCopyObjectArgs{
		StorageClass:     api.STORAGE_CLASS_COLD,
		ObjectTagging:    "k1=v1&k2=v2",
		TaggingDirective: api.METADATA_DIRECTIVE_COPY,
		CannedAcl:        api.CANNED_ACL_PRIVATE,
		GrantRead:        []string{"id1"},
		GrantFullControl: []string{"id2"},
	}
	res, err := client.ParallelCopy(EXISTS_BUCKET, "test_multiupload", EXISTS_BUCKET, "test_multiupload_copy", &args, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "BucketName", res.Bucket)

	// case2: get object meta fail
	options2 := util.RoundTripperOpts404
	client2, err := NewMockBosClient(ak, sk, endpoint, "", options2...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client2.ParallelCopy(EXISTS_BUCKET, "test_multiupload", EXISTS_BUCKET, "test_multiupload_copy", &args, nil)
	ExpectEqual(t.Errorf, bceServiceErro404.Error(), err.Error())
	ExpectEqual(t.Errorf, nil, res)

	// case3: init fail
	options3 := options
	client3, err := NewMockBosClient(ak, sk, endpoint, "", options3...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client3.ParallelCopy(EXISTS_BUCKET, "test_multiupload", EXISTS_BUCKET, "test_multiupload_copy", &args, nil)
	ExpectEqual(t.Errorf, "EOF", err.Error())
	ExpectEqual(t.Errorf, nil, res)

	// case4: upload part copy fail
	options4 := append(options, util.AppendRespBody([]string{"", respBody1, ""}))
	client4, err := NewMockBosClient(ak, sk, endpoint, "", options4...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client4.ParallelCopy(EXISTS_BUCKET, "test_multiupload", EXISTS_BUCKET, "test_multiupload_copy", &args, nil)
	ExpectEqual(t.Errorf, "EOF", err.Error())
	ExpectEqual(t.Errorf, nil, res)

	//case5: complete fail
	options5 := append(options, util.AppendRespBody([]string{"", respBody1, respBody2, ""}))
	client5, err := NewMockBosClient(ak, sk, endpoint, "", options5...)
	ExpectEqual(t.Errorf, nil, err)
	res, err = client5.ParallelCopy(EXISTS_BUCKET, "test_multiupload", EXISTS_BUCKET, "test_multiupload_copy", &args, nil)
	ExpectEqual(t.Errorf, "EOF", err.Error())
	ExpectEqual(t.Errorf, nil, res)

	// case6: parallelPartCopy
	options6 := append(options, util.AppendRespBody([]string{"", respBody2}))
	client6, err := NewMockBosClient(ak, sk, endpoint, "", options6...)
	ExpectEqual(t.Errorf, nil, err)
	metaRes, err := client6.GetObjectMeta(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, nil, err)
	uploadId := "a44cc9bab11cbd156984767aad637851"
	uploadInfo, err := client6.parallelPartCopy(*metaRes, "source", EXISTS_BUCKET, EXISTS_OBJECT, uploadId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(uploadInfo))

	// case7: singlePartCopy
	parallelChan := make(chan int, 1)
	errChan := make(chan error, 1)
	resultChan := make(chan api.UploadInfoType, 1)
	parallelChan <- 1
	client6.singlePartCopy("source", EXISTS_BUCKET, EXISTS_OBJECT, uploadId, 1, nil, parallelChan, errChan, resultChan)
	result7 := <-resultChan
	ExpectEqual(t.Errorf, "67b92a7c2a9b9c1809a6ae3295dcc127", result7.ETag)
}

func TestBucketTagPut(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	args := &api.PutBucketTagArgs{
		Tags: []api.Tag{
			{
				TagKey:   "key1",
				TagValue: "value1",
			},
		},
	}
	err = client.PutBucketTag(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
}
func TestBucketTagGet(t *testing.T) {
	respBody := `{
        "tag":[
            {
                "tag_key":"key1",
                "tag_value":"value123"
            },
            {
                "tag_key":"ttt2",
                "tag_value":"6863gerg"
            }
        ]
    }`
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetBucketTag(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, len(res.Tags), 2)
	ExpectEqual(t.Errorf, res.Tags[0].TagKey, "key1")
	ExpectEqual(t.Errorf, res.Tags[1].TagKey, "ttt2")
	ExpectEqual(t.Errorf, res.Tags[0].TagValue, "value123")
	ExpectEqual(t.Errorf, res.Tags[1].TagValue, "6863gerg")
}
func TestBucketTagDelete(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketTag(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}

func TestSymLinkPut(t *testing.T) {
	// mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	args := &api.PutSymlinkArgs{}
	err = client.PutSymlink(EXISTS_BUCKET, "test-symlink", EXISTS_OBJECT, args)
	ExpectEqual(t.Errorf, err, nil)
}
func TestSymLinkGet(t *testing.T) {
	// mock bos client
	options := []util.MockRoundTripperOption{
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.AddHeaders(map[string]string{
			http.CanonicalHeaderKey(my_http.BCE_SYMLINK_TARGET): "test-symlink",
		}),
	}
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "", options...)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetSymlink(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, res, "test-symlink")
	t.Logf("%v, %v", res, err)
}

func TestBucketMirrorPut(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	//PutBucketMirror
	args := &api.PutBucketMirrorArgs{
		BucketMirroringConfiguration: []api.MirrorConfigurationRule{
			{
				SourceUrl:    "http://gosdk-unittest-bucket.bj.bcebos.com",
				Prefix:       "testprefix",
				StorageClass: api.STORAGE_CLASS_STANDARD,
				Mode:         "fetch",
			},
		},
	}
	err = client.PutBucketMirror(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, err, nil)
}
func TestBucketMirrorGet(t *testing.T) {
	respBody := `{
		"bucketMirroringConfiguration":[
			{
			    "mode":"fetch",
			    "sourceUrl":"http://www.baidu.com",
			    "backSourceUrl":"bos://bj.bcebos.com/bucket",
			    "resource" : "folder1/folder2*.jpeg",
			    "prefix": "testprefix",
			    "suffix": ".jpeg",
			    "fixedKey":"folder1/404.jpeg", 
			    "version":"v2",                          
			    "prefixReplace" : "a/b/c",           
			    "passQuerystring":true,
			    "storageClass":"STANDARD",
			    "allHeader":"custom",
			    "customHeaders": [
				    {
					    "headerName":"testheader1",
					    "headerValue":"name1"
				    },
				    {
					    "headerName":"TestHeaderName",
					    "headerValue":"TestHeaderValue"
				    }
			    ],
			    "ignoreHeaders": ["BanHeader1","BanHeader2"],
			    "passHeaders":["AllowHeader1","AllowHeader2"]
			}
		]
	}`
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	//GetBucketMirror
	MirroConfig, err := client.GetBucketMirror(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, len(MirroConfig.BucketMirroringConfiguration), 1)
	ExpectEqual(t.Errorf, len(MirroConfig.BucketMirroringConfiguration[0].CustomHeaders), 2)
	ExpectEqual(t.Errorf, len(MirroConfig.BucketMirroringConfiguration[0].IgnoreHeaders), 2)
	ExpectEqual(t.Errorf, len(MirroConfig.BucketMirroringConfiguration[0].PassHeaders), 2)
	ExpectEqual(t.Errorf, MirroConfig.BucketMirroringConfiguration[0].SourceUrl, "http://www.baidu.com")
	ExpectEqual(t.Errorf, MirroConfig.BucketMirroringConfiguration[0].BackSourceUrl, "bos://bj.bcebos.com/bucket")
	ExpectEqual(t.Errorf, MirroConfig.BucketMirroringConfiguration[0].Prefix, "testprefix")
}
func TestBucketMirrorDelete(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteBucketMirror(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, err, nil)
}

// TestNewClientWithConfig_EmptyAkSk 是用于测试 NewClientWithConfig_EmptyAkSk
// generated by Comate
func TestNewClientWithConfig_EmptyAkSk(t *testing.T) {
	config := &BosClientConfiguration{
		Ak:       "",
		Sk:       "",
		Endpoint: "test-endpoint",
	}

	client, err := NewClientWithConfig(config)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if client == nil {
		t.Fatal("Expected client to be created")
	}
}

// TestNewClientWithConfig_EmptyEndpoint 是用于测试 NewClientWithConfig_EmptyEndpoint
// generated by Comate
func TestNewClientWithConfig_EmptyEndpoint(t *testing.T) {
	config := &BosClientConfiguration{
		Ak: "test-ak",
		Sk: "test-sk",
		// Endpoint为空，应该使用默认值
	}

	client, err := NewClientWithConfig(config)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if client == nil {
		t.Fatal("Expected client to be created")
	}
}

func TestPutObjectTag(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	putObjectTagArgs := &api.PutObjectTagArgs{
		ObjectTags: []api.ObjectTags{
			{
				TagInfo: []api.ObjectTag{
					{Key: "key1", Value: "value1"},
					{Key: "key2", Value: "value2"},
				},
			},
		},
	}
	err = client.PutObjectTag(EXISTS_BUCKET, EXISTS_OBJECT, putObjectTagArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetObjectTag(t *testing.T) {
	respBody := `{
		"tagSet": [{
			"tagInfo": { 
				"key9": "value9", "key8": "value8", "key10": "value10", "key3": "value3", "key2": "value2",
				"key1": "value1", "key0": "value0", "key6": "value6", "key5": "value5", "key4": "value4"
			}
		}]
	}`
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client.GetObjectTag(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, err, nil)
	ExpectEqual(t.Errorf, 10, len(res))
}

func TestDeleteObjectTag(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client.DeleteObjectTag(EXISTS_BUCKET, EXISTS_OBJECT)
	ExpectEqual(t.Errorf, nil, err)
}

func TestBosShareLinkGet(t *testing.T) {
	respBody := `{"shareUrl":"url","linkExpireTime":180,"shareCode":"111111"}`
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	res, err := client.BosShareLinkGet(EXISTS_BUCKET, "prefix", "111111", 180)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, respBody, res)
}

func TestPutBucketVersioning(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	args := &api.BucketVersioningArgs{Status: "enabled"}
	err = client.PutBucketVersioning(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetBucketVersioning(t *testing.T) {
	respBody := `{"status": "enabled"}`
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, respBody)
	ExpectEqual(t.Errorf, nil, err)

	res, err := client.GetBucketVersioning(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "enabled", res.Status)
}

func TestBucketInventory(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)

	argsJsonStr := `{
		"inventoryRuleList":[
			{
				"id": "inventory-configuration-ID1", 
				"status": "enabled", 
				"resource": [ "bucket/prefix/*" ], 
				"schedule": "Weekly", 
				"destination": { "targetBucket": "destBucketName", "targetPrefix": "destination-prefix", "format": "CSV" }
			}, 
			{
				"id": "inventory-configuration-ID2", 
				"status": "enabled", 
				"resource": [ "bucket/prefix2/*" ], 
				"schedule": "Monthly", 
				"monthlyDate": 15,
				"destination": { "targetBucket": "destBucketName", "targetPrefix": "destination-prefix-another", "format": "CSV" }
			}
		]
	}`

	listRes := &api.ListBucketInventoryResult{}
	ExpectEqual(t.Errorf, nil, json.Unmarshal([]byte(argsJsonStr), listRes))
	ExpectEqual(t.Errorf, 2, len(listRes.RuleList))
	args := &api.PutBucketInventoryArgs{Rule: listRes.RuleList[0]}

	// put bucket inventory
	err = client.PutBucketInventory(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, nil, err)

	// get bucket inventory
	jsonStr, err := json.Marshal(args.Rule)
	ExpectEqual(t.Errorf, nil, err)
	singleRuleStr := string(jsonStr)
	client1, err := NewMockBosClient(ak, sk, endpoint, singleRuleStr)
	ExpectEqual(t.Errorf, nil, err)

	getRes, err := client1.GetBucketInventory(EXISTS_BUCKET, "id")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "inventory-configuration-ID1", getRes.Rule.Id)

	// list bucket inventory
	client2, err := NewMockBosClient(ak, sk, endpoint, argsJsonStr)
	ExpectEqual(t.Errorf, nil, err)
	res1, err := client2.ListBucketInventory(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 2, len(res1.RuleList))

	// delete bucket inventory
	client3, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	err = client3.DeleteBucketInventory(EXISTS_BUCKET, "id")
	ExpectEqual(t.Errorf, nil, err)
}

func TestBucketRequestPayment(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	args := &api.RequestPaymentArgs{RequestPayment: "Requester"}
	//PutBucketRequestPayment
	err = client.PutBucketRequestPayment(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, nil, err)
	//GetBucketRequestPayment
	respBody1 := `{"requestPayment": "Requester"}`
	client1, err := NewMockBosClient(ak, sk, endpoint, respBody1)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client1.GetBucketRequestPayment(EXISTS_BUCKET, nil)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Requester", res.RequestPayment)
}

func TestBucketObjectLock(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	// InitBucketObjectLock
	args := &api.InitBucketObjectLockArgs{
		RetentionDays: 10,
	}
	err = client.InitBucketObjectLock(EXISTS_BUCKET, args, nil)
	ExpectEqual(t.Errorf, nil, err)
	//GetBucketObjectLock
	respBody1 := `{
		"lockStatus": "IN_PROGRESS",
		"createDate": 1569317168,
		"expirationDate": 1569403568,
		"retentionDays": 3
	}`
	client1, err := NewMockBosClient(ak, sk, endpoint, respBody1)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client1.GetBucketObjectLock(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1569317168, res.CreateDate)
	ExpectEqual(t.Errorf, 1569403568, res.ExpirationDate)
	ExpectEqual(t.Errorf, "IN_PROGRESS", res.LockStatus)
	ExpectEqual(t.Errorf, 3, res.RetentionDays)
	//DeleteBucketObjectLock
	err = client.DeleteBucketObjectLock(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
	//CompleteBucketObjectLock
	err = client.CompleteBucketObjectLock(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
	//ExtendBucketObjectLock
	extArgs := &api.ExtendBucketObjectLockArgs{ExtendRetentionDays: 1}
	err = client.ExtendBucketObjectLock(EXISTS_BUCKET, extArgs)
	ExpectEqual(t.Errorf, nil, err)
}
func TestBucketQuota(t *testing.T) {
	//mock bos client
	ak, sk, endpoint := "ak", "sk", "bj.bcebos.com"
	client, err := NewMockBosClient(ak, sk, endpoint, "")
	ExpectEqual(t.Errorf, nil, err)
	args := &api.BucketQuotaArgs{
		MaxObjectCount:       100,
		MaxCapacityMegaBytes: 99999999,
	}
	//put bucket quota
	err = client.PutBucketQuota(EXISTS_BUCKET, args)
	ExpectEqual(t.Errorf, nil, err)

	//get bucket quota
	respBody1 := `{ "maxObjectCount": 50,  "maxCapacityMegaBytes"  : 12334424 }`
	client1, err := NewMockBosClient(ak, sk, endpoint, respBody1)
	ExpectEqual(t.Errorf, nil, err)
	res, err := client1.GetBucketQuota(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 50, res.MaxObjectCount)
	ExpectEqual(t.Errorf, 12334424, res.MaxCapacityMegaBytes)

	// delete bucket quota
	err = client.DeleteBucketQuota(EXISTS_BUCKET)
	ExpectEqual(t.Errorf, nil, err)
}

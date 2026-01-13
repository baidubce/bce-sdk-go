package bce

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

func init() {
	log.SetLogLevel(log.WARN)
	log.SetLogHandler(log.STDOUT)
}

var (
	bceServiceErro408 *BceServiceError = NewBceServiceError("Timeout", "Request Timeout", "", http.StatusRequestTimeout)
	bceServiceErro500 *BceServiceError = NewBceServiceError("InternalError", "Internal Server Error", "", http.StatusInternalServerError)
)

type ErrorTypeTransport struct {
}

func (ett *ErrorTypeTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func ExpectEqual(t *testing.T, exp interface{}, act interface{}) bool {
	if !util.Equal(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d: missmatch, expect %v but %v", file, line, exp, act)
		return false
	}
	return true
}

func TestNewBceClientWithExclusiveHTTPClient(t *testing.T) {
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	// case1: error type http transport
	config := &BceClientConfiguration{
		Endpoint:                  "endpoint",
		Region:                    "bce.DEFAULT_REGION",
		UserAgent:                 "bce.DEFAULT_USER_AGENT",
		Credentials:               &auth.BceCredentials{AccessKeyId: "ak", SecretAccessKey: "sk"},
		SignOption:                defaultSignOptions,
		Retry:                     NewBackOffRetryPolicy(3, 100, 20000),
		ConnectionTimeoutInMillis: DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
		HTTPClient: &http.Client{
			Timeout:   10000,
			Transport: &ErrorTypeTransport{},
		},
	}
	_, err := NewBceClientWithExclusiveHTTPClient(config, &auth.BceV1Signer{})
	ExpectEqual(t, nil, err)
	// case2: rate limit parameters
	rateLimit := int64(8192000)
	config1 := &BceClientConfiguration{
		Endpoint:                  "endpoint",
		Region:                    "bce.DEFAULT_REGION",
		UserAgent:                 "bce.DEFAULT_USER_AGENT",
		Credentials:               &auth.BceCredentials{AccessKeyId: "ak", SecretAccessKey: "sk"},
		SignOption:                defaultSignOptions,
		Retry:                     NewBackOffRetryPolicy(3, 100, 20000),
		ConnectionTimeoutInMillis: DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
		HTTPClient:                &http.Client{},
		UploadRatelimit:           &rateLimit,
		DownloadRatelimit:         &rateLimit,
	}
	_, err = NewBceClientWithExclusiveHTTPClient(config1, &auth.BceV1Signer{})
	ExpectEqual(t, nil, err)
}

func TestSendRequestV2(t *testing.T) {
	// new bce client
	maxRetry := 3
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	config := &BceClientConfiguration{
		Endpoint:                  "endpoint",
		Region:                    "bce.DEFAULT_REGION",
		UserAgent:                 "bce.DEFAULT_USER_AGENT",
		Credentials:               &auth.BceCredentials{AccessKeyId: "ak", SecretAccessKey: "sk"},
		SignOption:                defaultSignOptions,
		Retry:                     NewBackOffRetryPolicy(maxRetry, 100, 20000),
		ConnectionTimeoutInMillis: DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS,
	}
	client, err := NewBceClientWithExclusiveHTTPClient(config, &auth.BceV1Signer{})
	ExpectEqual(t, nil, err)
	// case1: client error not nil
	req1 := &BceRequest{}
	resp1 := &BceResponse{}
	clientErr1 := NewBceClientError("client error")
	req1.SetClientError(clientErr1)
	err = client.SendRequestV2(req1, resp1)
	ExpectEqual(t, clientErr1, err)

	// case2: has body, body is not TeeReadNopCloser, ok
	req2 := &BceRequest{}
	resp2 := &BceResponse{}
	body2, err := NewBodyFromString("body string")
	ExpectEqual(t, nil, err)
	req2.SetBody(body2)
	mockHttpClient2 := util.NewMockHTTPClient()
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient2))
	client.HTTPClient = mockHttpClient2
	err = client.SendRequestV2(req2, resp2)
	ExpectEqual(t, nil, err)

	// case3: has not body, http client do error and retry
	respErr3 := fmt.Errorf("IO Error")
	options3 := []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(respErr3),
	}
	mockHttpClient3 := util.NewMockHTTPClient(options3...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient3))
	client.HTTPClient = mockHttpClient3

	host := "hostname:8080"
	uri := "bucket/object"
	test_url := "http://" + host + "/" + uri
	req3 := &BceRequest{}
	resp3 := &BceResponse{}
	req3.SetMethod(http.MethodGet)
	req3.SetUri(uri)
	req3.SetHost(host)
	req3.SetProtocol("http")
	err = client.SendRequestV2(req3, resp3)
	urlErr3 := url.Error{Op: "Get", URL: test_url, Err: respErr3}
	expectErr3 := &BceClientError{fmt.Sprintf("execute http request failed! Retried %d times, error: %v", maxRetry, urlErr3.Error())}
	ExpectEqual(t, expectErr3, err)

	// case4: has not body, http client do error and not retry
	respErr4 := fmt.Errorf("context deadline exceeded")
	options4 := []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(respErr4),
	}
	mockHttpClient4 := util.NewMockHTTPClient(options4...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient4))
	client.HTTPClient = mockHttpClient4

	req4 := &BceRequest{}
	resp4 := &BceResponse{}
	req4.SetMethod(http.MethodGet)
	req4.SetUri(uri)
	req4.SetHost(host)
	req4.SetProtocol("http")
	err = client.SendRequestV2(req4, resp4)
	urlErr4 := url.Error{Op: "Get", URL: test_url, Err: respErr4}
	expectErr4 := &BceClientError{fmt.Sprintf("execute http request failed! Retried 0 times, error: %v", urlErr4.Error())}
	ExpectEqual(t, expectErr4, err)

	// case5: has not body, handle time exceed 5s
	options5 := []util.MockRoundTripperOption{
		util.SetRequestTime(5010 * time.Millisecond),
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
	}
	mockHttpClient5 := util.NewMockHTTPClient(options5...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient5))
	client.HTTPClient = mockHttpClient5
	err = client.SendRequestV2(req4, resp4)
	ExpectEqual(t, nil, err)

	// case6: has not body, response error 500, retry
	options6 := util.RoundTripperOpts500
	mockHttpClient6 := util.NewMockHTTPClient(options6...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient6))
	client.HTTPClient = mockHttpClient6
	err = client.SendRequestV2(req4, resp4)
	ExpectEqual(t, bceServiceErro500, err)

	// case7: has not body response error 408, not retry
	options7 := util.RoundTripperOpts408
	mockHttpClient7 := util.NewMockHTTPClient(options7...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient7))
	client.HTTPClient = mockHttpClient7
	err = client.SendRequestV2(req4, resp4)
	ExpectEqual(t, bceServiceErro408, err)

	// case8: has body, body is not TeeReadNopCloser, http client do error, not retry
	req8 := &BceRequest{}
	resp8 := &BceResponse{}
	req8.SetMethod(http.MethodGet)
	req8.SetUri(uri)
	req8.SetHost(host)
	req8.SetProtocol("http")
	body8, err := NewBodyFromString("body string")
	ExpectEqual(t, nil, err)
	req8.SetBody(body8)
	err8 := fmt.Errorf("error 8")
	options8 := []util.MockRoundTripperOption{util.SetHTTPClientDoError(err8)}
	urlErr8 := url.Error{Op: "Get", URL: test_url, Err: err8}
	mockHttpClient8 := util.NewMockHTTPClient(options8...)
	ExpectEqual(t, true, mockHttpClient8 != nil)
	client.HTTPClient = mockHttpClient8
	err = client.SendRequestV2(req8, resp8)
	ExpectEqual(t, urlErr8.Error(), err.Error())

	// case9: has body, body is not TeeReadNopCloser, response error 500, not retry
	req9 := &BceRequest{}
	resp9 := &BceResponse{}
	req9.SetMethod(http.MethodGet)
	req9.SetUri(uri)
	req9.SetHost(host)
	req9.SetProtocol("http")
	body9, err := NewBodyFromString("body string")
	ExpectEqual(t, nil, err)
	req9.SetBody(body9)
	options9 := util.RoundTripperOpts500
	mockHttpClient9 := util.NewMockHTTPClient(options9...)
	ExpectEqual(t, true, mockHttpClient9 != nil)
	client.HTTPClient = mockHttpClient9
	err = client.SendRequestV2(req9, resp9)
	ExpectEqual(t, bceServiceErro500, err)
	// case10: has body, body is not TeeReadNopCloser, response error 408, not retry
	req10 := &BceRequest{}
	resp10 := &BceResponse{}
	req10.SetMethod(http.MethodGet)
	req10.SetUri(uri)
	req10.SetHost(host)
	req10.SetProtocol("http")
	body10, err := NewBodyFromString("body string")
	ExpectEqual(t, nil, err)
	req10.SetBody(body10)
	options10 := util.RoundTripperOpts408
	mockHttpClient10 := util.NewMockHTTPClient(options10...)
	ExpectEqual(t, true, mockHttpClient10 != nil)
	client.HTTPClient = mockHttpClient10
	err = client.SendRequestV2(req10, resp10)
	ExpectEqual(t, bceServiceErro408, err)
	// case11: has body, body is TeeReadNopCloser, http client do error, retry
	req11 := &BceRequest{}
	resp11 := &BceResponse{}
	req11.SetMethod(http.MethodGet)
	req11.SetUri(uri)
	req11.SetHost(host)
	req11.SetProtocol("http")
	body11, err := NewBodyFromStringV2("body string", false)
	ExpectEqual(t, nil, err)
	req11.SetBody(body11)
	err11 := fmt.Errorf("error 11")
	options11 := []util.MockRoundTripperOption{util.SetHTTPClientDoError(err11)}
	urlErr11 := url.Error{Op: "Get", URL: test_url, Err: err11}
	expectErr11 := &BceClientError{fmt.Sprintf("execute http request failed! Retried 3 times, error: %v", urlErr11.Error())}
	mockHttpClient11 := util.NewMockHTTPClient(options11...)
	ExpectEqual(t, true, mockHttpClient11 != nil)
	client.HTTPClient = mockHttpClient11
	err = client.SendRequestV2(req11, resp11)
	ExpectEqual(t, expectErr11, err)
	// case12: has body, body is TeeReadNopCloser, response error 500, retry
	req12 := &BceRequest{}
	resp12 := &BceResponse{}
	req12.SetMethod(http.MethodGet)
	req12.SetUri(uri)
	req12.SetHost(host)
	req12.SetProtocol("http")
	body12, err := NewBodyFromStringV2("body string", false)
	ExpectEqual(t, nil, err)
	req12.SetBody(body12)
	options12 := util.RoundTripperOpts500
	mockHttpClient12 := util.NewMockHTTPClient(options12...)
	ExpectEqual(t, true, mockHttpClient12 != nil)
	client.HTTPClient = mockHttpClient12
	err = client.SendRequestV2(req12, resp12)
	ExpectEqual(t, bceServiceErro500, err)
	// case13: has body, body is TeeReadNopCloser, response error 408, not retry
	req13 := &BceRequest{}
	resp13 := &BceResponse{}
	req13.SetMethod(http.MethodGet)
	req13.SetUri(uri)
	req13.SetHost(host)
	req13.SetProtocol("http")
	body13, err := NewBodyFromStringV2("body string", false)
	ExpectEqual(t, nil, err)
	req13.SetBody(body13)
	options13 := util.RoundTripperOpts408
	mockHttpClient13 := util.NewMockHTTPClient(options13...)
	ExpectEqual(t, true, mockHttpClient13 != nil)
	client.HTTPClient = mockHttpClient13
	err = client.SendRequestV2(req13, resp13)
	ExpectEqual(t, bceServiceErro408, err)
}

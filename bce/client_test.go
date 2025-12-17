package bce

import (
	"fmt"
	net_http "net/http"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/baidubce/bce-sdk-go/util"
)

type ErrorTypeTransport struct {
}

func (ett *ErrorTypeTransport) RoundTrip(*net_http.Request) (*net_http.Response, error) {
	return &net_http.Response{}, nil
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
		HTTPClient: &net_http.Client{
			Timeout:   10000,
			Transport: &ErrorTypeTransport{},
		},
	}
	_, err := NewBceClientWithExclusiveHTTPClient(config, &auth.BceV1Signer{})
	ExpectEqual(t, http.UnknownHTTPTransport, fmt.Errorf("%s", err))
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
		HTTPClient:                &net_http.Client{},
		UploadRatelimit:           &rateLimit,
		DownloadRatelimit:         &rateLimit,
	}
	_, err = NewBceClientWithExclusiveHTTPClient(config1, &auth.BceV1Signer{})
	ExpectEqual(t, nil, err)
}

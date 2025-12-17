package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util"
)

var (
	roundTripperOpts408 = []MockRoundTripperOption{setStatusCode(http.StatusRequestTimeout), setStatusMsg(http.StatusText(http.StatusRequestTimeout))}
)

func ExpectEqual(t *testing.T, exp interface{}, act interface{}) bool {
	if !util.Equal(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d: missmatch, expect %v but %v", file, line, exp, act)
		return false
	}
	return true
}

type MockRoundTripperOption func(*MockRoundTripper)

type MockRoundTripper struct {
	http.Transport
	Err        error
	StatusCode int
	StatusMsg  string
	RespBody   string
}

func setHTTPClientDoError(err error) MockRoundTripperOption {
	return func(m *MockRoundTripper) {
		m.Err = err
	}
}

func setStatusCode(statusCode int) MockRoundTripperOption {
	return func(m *MockRoundTripper) {
		m.StatusCode = statusCode
	}
}

func setStatusMsg(statusMsg string) MockRoundTripperOption {
	return func(m *MockRoundTripper) {
		m.StatusMsg = statusMsg
	}
}

func setRespBody(respBody string) MockRoundTripperOption {
	return func(m *MockRoundTripper) {
		m.RespBody = respBody
	}
}

func (m *MockRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if m.Err != nil {
		return nil, m.Err
	}

	resp := &http.Response{
		StatusCode: m.StatusCode,
		Status:     m.StatusMsg,
		Body:       io.NopCloser(bytes.NewBufferString(m.RespBody)),
		Header:     make(http.Header),
	}
	return resp, nil
}

func NewMockHTTPClient(options ...MockRoundTripperOption) *http.Client {
	mockRoundTripper := &MockRoundTripper{}
	for _, option := range options {
		option(mockRoundTripper)
	}
	return &http.Client{
		Transport: mockRoundTripper,
	}
}

type ErrorTypeTransport struct {
}

func (ett *ErrorTypeTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{}, nil
}

func TestInitWithSpecifiedClient(t *testing.T) {
	var nilTranport *http.Transport = nil
	transport = nil
	// case1: error type http transport
	testHTTPClient := &http.Client{
		Timeout:   10 * 1000,
		Transport: &ErrorTypeTransport{},
	}
	err := InitWithSpecifiedClient(testHTTPClient)
	ExpectEqual(t, UnknownHTTPTransport, err)
	ExpectEqual(t, nilTranport, transport)
	// case2:default http client
	transport = nil
	testHTTPClient1 := &http.Client{}
	err = InitWithSpecifiedClient(testHTTPClient1)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, nilTranport, transport)
	// case3:customized http client
	transport = nil
	testHTTPClient2 := &http.Client{
		Timeout:   10 * 1000,
		Transport: NewTransportCustom(&DefaultClientConfig),
	}
	err = InitWithSpecifiedClient(testHTTPClient2)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, false, util.Equal(nilTranport, transport))
	// case4:nil http client
	transport = nil
	err = InitWithSpecifiedClient(nil)
	ExpectEqual(t, NilHTTPClient, err)
	// case5: mock http transport
	testHTTPClient3 := &http.Client{
		Timeout:   10 * 1000,
		Transport: &MockRoundTripper{},
	}
	err = InitWithSpecifiedClient(testHTTPClient3)
	ExpectEqual(t, nilTranport, transport)
	ExpectEqual(t, UnknownHTTPTransport, err)
}

func TestExecute(t *testing.T) {
	var nilClient *http.Client = nil
	transport = nil
	// mock http client
	options := []MockRoundTripperOption{
		setRespBody("respBody"),
		setStatusCode(200),
		setStatusMsg("200 OK"),
	}
	mockHttpClient := NewMockHTTPClient(options...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient))

	// case1: response is ok
	host := "host"
	request := &Request{}
	request.SetMethod(http.MethodGet)
	request.SetHost("host")
	request.SetProtocol("http")
	request.SetHTTPClient(mockHttpClient)
	_, err := Execute(request)
	ExpectEqual(t, nil, err)

	// case2: http client do error
	ioErr := fmt.Errorf("IO error.")
	roundTripperOptions := []MockRoundTripperOption{
		setHTTPClientDoError(ioErr),
	}
	mockHttpClient1 := NewMockHTTPClient(roundTripperOptions...)
	ExpectEqual(t, false, util.Equal(nilClient, mockHttpClient1))

	request.SetHTTPClient(mockHttpClient1)
	_, err = Execute(request)
	urlError := url.Error{
		Op:  "Get",
		URL: "http://" + host,
		Err: ioErr,
	}
	ExpectEqual(t, urlError.Error(), err.Error())

	// case2: reaponse 408
	roundTripperOptions = roundTripperOpts408
	mockHttpClient2 := NewMockHTTPClient(roundTripperOptions...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient2))

	request.SetMethod(http.MethodPut)
	request.SetHTTPClient(mockHttpClient2)
	_, err = Execute(request)
	ExpectEqual(t, nil, err)
}

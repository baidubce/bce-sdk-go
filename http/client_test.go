package http

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util"
)

func ExpectEqual(t *testing.T, exp interface{}, act interface{}) bool {
	if !util.Equal(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d: missmatch, expect %v but %v", file, line, exp, act)
		return false
	}
	return true
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
	ExpectEqual(t, nil, err)
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
		Transport: &util.MockRoundTripper{},
	}
	err = InitWithSpecifiedClient(testHTTPClient3)
	ExpectEqual(t, nilTranport, transport)
	ExpectEqual(t, nil, err)
}

func TestExecute(t *testing.T) {
	var nilClient *http.Client = nil
	transport = nil
	// mock http client
	options := []util.MockRoundTripperOption{
		util.SetRespBody("respBody"),
		util.SetStatusCode(200),
		util.SetStatusMsg("200 OK"),
	}
	mockHttpClient := util.NewMockHTTPClient(options...)
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
	roundTripperOptions := []util.MockRoundTripperOption{
		util.SetHTTPClientDoError(ioErr),
	}
	mockHttpClient1 := util.NewMockHTTPClient(roundTripperOptions...)
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
	roundTripperOptions = util.RoundTripperOpts408
	mockHttpClient2 := util.NewMockHTTPClient(roundTripperOptions...)
	ExpectEqual(t, false, util.Equal(nil, mockHttpClient2))

	request.SetMethod(http.MethodPut)
	request.SetHTTPClient(mockHttpClient2)
	_, err = Execute(request)
	ExpectEqual(t, nil, err)
}

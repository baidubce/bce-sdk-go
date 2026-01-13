// nolint
package util

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"testing"
	"time"
)

func ExpectEqual(t *testing.T, exp interface{}, act interface{}) bool {
	if !Equal(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d: missmatch, expect %v but %v", file, line, exp, act)
		return false
	}
	return true
}

type testStruct struct {
	Member string
}

func TestExpectEqual(t *testing.T) {
	// case1: both nil
	Equal(true, Equal(nil, nil))
	var nilStruct *testStruct = nil
	Equal(true, Equal(nil, nilStruct))
	Equal(true, Equal(nilStruct, nil))
	ExpectEqual(t, nil, nilStruct)
	// case2: both not nil and equal
	struct1 := &testStruct{}
	struct2 := &testStruct{}
	Equal(true, Equal(struct1, struct2))
	ExpectEqual(t, struct1, struct2)
	// case3: one nil, one not nil
	Equal(false, Equal(struct1, nilStruct))
	Equal(false, Equal(nilStruct, struct1))
	// case4: both not nil, and unequal
	struct3 := &testStruct{
		Member: "member",
	}
	Equal(false, Equal(struct1, struct3))
}

func TestMockHTTPClient(t *testing.T) {
	// case1: mock http client do return 200 OK
	respBodyStr := "respBodyStr"
	options1 := []MockRoundTripperOption{
		SetStatusCode(http.StatusOK),
		SetStatusMsg(http.StatusText(http.StatusOK)),
		SetRespBody(respBodyStr),
	}

	mockHttpClient := NewMockHTTPClient(options1...)
	ExpectEqual(t, false, Equal(nil, mockHttpClient))

	test_url := "http://hostname:8888/bucket/object"
	req1, err := http.NewRequest(http.MethodGet, test_url, nil)
	ExpectEqual(t, nil, err)
	resp1, err := mockHttpClient.Do(req1)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, http.StatusOK, resp1.StatusCode)
	ExpectEqual(t, http.StatusText(http.StatusOK), resp1.Status)
	resp1Body, err := io.ReadAll(resp1.Body)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, respBodyStr, string(resp1Body))

	// case2: http client do return error
	ioErr := fmt.Errorf("IO Error")
	options2 := []MockRoundTripperOption{
		SetHTTPClientDoError(ioErr),
	}
	mockHttpClient2 := NewMockHTTPClient(options2...)
	ExpectEqual(t, false, Equal(nil, mockHttpClient))
	_, err = mockHttpClient2.Do(req1)
	expectErr := url.Error{
		Op:  "Get",
		URL: req1.URL.String(),
		Err: ioErr,
	}
	ExpectEqual(t, expectErr.Error(), err.Error())

	// case3: test single func
	time1 := 2 * time.Second
	respBody := "body-string"
	testMockRoundTripper := &MockRoundTripper{
		Err:         fmt.Errorf("erro1"),
		StatusCode:  http.StatusOK,
		StatusMsg:   http.StatusText(http.StatusOK),
		RespBody:    []string{respBody},
		RequestTime: &time1,
		Headers:     make(map[string]string),
	}
	ExpectEqual(t, fmt.Errorf("erro1"), testMockRoundTripper.Err)
	ExpectEqual(t, http.StatusOK, testMockRoundTripper.StatusCode)
	ExpectEqual(t, http.StatusText(http.StatusOK), testMockRoundTripper.StatusMsg)
	ExpectEqual(t, respBody, testMockRoundTripper.RespBody[0])
	ExpectEqual(t, time1, *testMockRoundTripper.RequestTime)
	ExpectEqual(t, make(map[string]string), testMockRoundTripper.Headers)
	option := SetHTTPClientDoError(fmt.Errorf("error2"))
	option(testMockRoundTripper)
	ExpectEqual(t, fmt.Errorf("error2"), testMockRoundTripper.Err)
	option = SetStatusCode(http.StatusAccepted)
	option(testMockRoundTripper)
	ExpectEqual(t, http.StatusAccepted, testMockRoundTripper.StatusCode)
	option = SetStatusMsg(http.StatusText(http.StatusAccepted))
	option(testMockRoundTripper)
	ExpectEqual(t, http.StatusText(http.StatusAccepted), testMockRoundTripper.StatusMsg)
	respBody1 := "body--string--1"
	option = SetRespBody(respBody1)
	option(testMockRoundTripper)
	ExpectEqual(t, respBody1, testMockRoundTripper.RespBody[0])
	option = SetRequestTime(1 * time.Second)
	option(testMockRoundTripper)
	ExpectEqual(t, 1*time.Second, *testMockRoundTripper.RequestTime)
	testMockRoundTripper.Headers = nil
	option = AddHeaders(map[string]string{"key1": "value1"})
	option(testMockRoundTripper)
	ExpectEqual(t, map[string]string{"key1": "value1"}, testMockRoundTripper.Headers)
	reqBody := "request-body"
	req, err := http.NewRequest(http.MethodGet, "http://endpoint", bytes.NewBufferString(reqBody))
	ExpectEqual(t, nil, err)
	ExpectEqual(t, len(reqBody), req.ContentLength)
	resp, err := testMockRoundTripper.RoundTrip(req)
	ExpectEqual(t, nil, resp)
	ExpectEqual(t, testMockRoundTripper.Err, err)
	testMockRoundTripper.Err = nil
	resp, err = testMockRoundTripper.RoundTrip(req)
	ExpectEqual(t, nil, err)
	ExpectEqual(t, http.StatusAccepted, resp.StatusCode)
	ExpectEqual(t, http.StatusText(http.StatusAccepted), resp.Status)
	buf := make([]byte, len(respBody1))
	resp.Body.Read(buf)
	ExpectEqual(t, respBody1, string(buf))
	ExpectEqual(t, map[string][]string{"Key1": {"value1"}}, resp.Header)
	//multi resp body
	respBody11 := "resp-body-001"
	respBody22 := "resp-body-002"
	roundTripper1 := &MockRoundTripper{
		RespBody: []string{
			respBody11,
			respBody22,
		},
	}
	httpReq1 := &http.Request{}
	resp11, err := roundTripper1.RoundTrip(httpReq1)
	ExpectEqual(t, nil, err)
	buf11 := make([]byte, len(respBody11))
	resp11.Body.Read(buf11)
	ExpectEqual(t, respBody11, string(buf11))
	resp22, err := roundTripper1.RoundTrip(httpReq1)
	ExpectEqual(t, nil, err)
	buf22 := make([]byte, len(respBody22))
	resp22.Body.Read(buf22)
	ExpectEqual(t, respBody22, string(buf22))
	resp33, err := roundTripper1.RoundTrip(httpReq1)
	ExpectEqual(t, nil, err)
	buf33 := make([]byte, len(respBody22))
	resp33.Body.Read(buf33)
	ExpectEqual(t, respBody22, string(buf33))
}

// TestEqual_ComplexStructsWithDifferentTypes 是用于测试 Equal_ComplexStructsWithDifferentTypes
// generated by Comate
func TestEqual_ComplexStructsWithDifferentTypes(t *testing.T) {
	// 另一个测试用例，覆盖第26行 - 复杂结构体类型不可转换
	type Person struct {
		Name string
		Age  int
	}

	type Animal struct {
		Species string
		Age     int
	}

	expected := Person{Name: "John", Age: 30}
	actual := Animal{Species: "Dog", Age: 5}

	result := Equal(expected, actual)
	expectedResult := false

	if result != expectedResult {
		t.Errorf("Equal(%v, %v) = %v, expected %v", expected, actual, result, expectedResult)
	}
}

package api

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
)

func Equal(expected, actual interface{}) bool {
	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}
	return false
}

func ExpectEqual(t *testing.T, exp interface{}, act interface{}) bool {
	if !Equal(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d: missmatch, expect %v but %v", file, line, exp, act)
		return false
	}
	return true
}

func TestBosRetryPolicyShouldRetry(t *testing.T) {
	retry_policy := NewBosRetryPolicy(3, 20000, 300)
	// attempt exceed the maxErrorRetry
	var err bce.BceError
	ExpectEqual(t, false, retry_policy.ShouldRetry(err, 4))
	// nil err
	ExpectEqual(t, true, retry_policy.ShouldRetry(err, 1))
	// net error, context deadline
	err = &net.OpError{
		Err: fmt.Errorf("context deadline exceeded"),
	}
	ExpectEqual(t, false, retry_policy.ShouldRetry(err, 1))
	err = &net.OpError{
		Err: fmt.Errorf("io timeout"),
	}
	ExpectEqual(t, true, retry_policy.ShouldRetry(err, 1))
	// bce service error, StatusInternalServerError
	err = &bce.BceServiceError{
		StatusCode: http.StatusInternalServerError,
	}
	ExpectEqual(t, true, retry_policy.ShouldRetry(err, 1))
	// bce service error, StatusBadGateway
	err = &bce.BceServiceError{
		StatusCode: http.StatusBadGateway,
	}
	ExpectEqual(t, true, retry_policy.ShouldRetry(err, 1))
	// bce service error, StatusServiceUnavailable
	err = &bce.BceServiceError{
		StatusCode: http.StatusServiceUnavailable,
	}
	ExpectEqual(t, true, retry_policy.ShouldRetry(err, 1))
	// bce service error, StatusForbidden
	err = &bce.BceServiceError{
		StatusCode: http.StatusForbidden,
	}
	ExpectEqual(t, true, retry_policy.ShouldRetry(err, 1))
	// bce service error, StatusBadRequest
	err = &bce.BceServiceError{
		StatusCode: http.StatusBadGateway,
	}
	ExpectEqual(t, true, retry_policy.ShouldRetry(err, 1))
	// bce service error, EREQUEST_EXPIRED
	err = &bce.BceServiceError{
		StatusCode: http.StatusNotFound,
		Code:       bce.EREQUEST_EXPIRED,
	}
	ExpectEqual(t, true, retry_policy.ShouldRetry(err, 1))
	// bce service error, ok
	err = &bce.BceServiceError{
		StatusCode: http.StatusOK,
		Code:       "ok",
	}
	ExpectEqual(t, false, retry_policy.ShouldRetry(err, 1))
}

func TestGetDelayBeforeNextRetryInMillis(t *testing.T) {
	maxRetry := 3
	maxDelayMs := int64(20000)
	baseIntervalMs := int64(300)
	retry_policy := NewBosRetryPolicy(maxRetry, maxDelayMs, baseIntervalMs)
	// attempt < 0
	var err bce.BceError
	ExpectEqual(t, 0, retry_policy.GetDelayBeforeNextRetryInMillis(err, -1))
	// attempt = 0, return 300
	ExpectEqual(t, baseIntervalMs*int64(time.Millisecond), retry_policy.GetDelayBeforeNextRetryInMillis(err, 0))
	// attempt = 1, return 600
	ExpectEqual(t, 2*baseIntervalMs*int64(time.Millisecond), retry_policy.GetDelayBeforeNextRetryInMillis(err, 1))
	// attempt = 2, return 1200
	ExpectEqual(t, 4*baseIntervalMs*int64(time.Millisecond), retry_policy.GetDelayBeforeNextRetryInMillis(err, 2))
	// attempt = 3, return 2400
	ExpectEqual(t, 8*baseIntervalMs*int64(time.Millisecond), retry_policy.GetDelayBeforeNextRetryInMillis(err, 3))
	// attempt = 10, return maxDelayMs
	ExpectEqual(t, maxDelayMs*int64(time.Millisecond), retry_policy.GetDelayBeforeNextRetryInMillis(err, 10))
}

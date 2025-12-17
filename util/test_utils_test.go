package util

import (
	"fmt"
	"runtime"
	"testing"
)

func ExpectEqual(t *testing.T, exp interface{}, act interface{}) bool {
	if !Equal(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		t.Errorf("%s:%d: missmatch, expect %v but %v", file, line, exp, act)
		return false
	}
	return true
}

func TestExpectEqual(t *testing.T) {
	// case1: both nil
	var nilErr error = nil
	Equal(true, Equal(nil, nilErr))
	ExpectEqual(t, nil, nilErr)
	// case2: both not nil and equal
	var err1 error = fmt.Errorf("error 1")
	var err2 error = fmt.Errorf("error 1")
	Equal(true, Equal(err1, err2))
	ExpectEqual(t, err1, err2)
	// case3: one nil, one not nil
	Equal(false, Equal(err1, nilErr))
	// case4: both not nil, and unequal
	var err3 error = fmt.Errorf("error 3")
	Equal(false, Equal(err1, err3))
}

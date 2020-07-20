package dts

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	DTS_CLIENT *Client
	DTS_ID     string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	DTS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
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

func TestClient_CreateDts(t *testing.T) {
	args := &CreateDtsArgs{
		ProductType:        "postpay",
		Type:               "migration",
		Standard:           "Large",
		SourceInstanceType: "public",
		TargetInstanceType: "public",
		CrossRegionTag:     0,
	}
	result, err := DTS_CLIENT.CreateDts(args)
	ExpectEqual(t.Errorf, nil, err)

	DTS_ID = result.DtsTasks[0].DtsId
}

func TestClient_GetDetail(t *testing.T) {
	result, err := DTS_CLIENT.GetDetail(DTS_ID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "unConfig", result.Status)
}

func TestClient_ListDts(t *testing.T) {
	args := &ListDtsArgs{
		Type: "migration",
	}
	result, err := DTS_CLIENT.ListDts(args)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Task {
		if e.DtsId == DTS_ID {
			ExpectEqual(t.Errorf, "unConfig", e.Status)
		}
	}
}

func TestClient_DeleteDts(t *testing.T) {
	err := DTS_CLIENT.DeleteDts(DTS_ID)
	ExpectEqual(t.Errorf, nil, err)
}
package eccr

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	CCR_CLIENT      *Client
	CCR_INSTANCE_ID string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fmt.Println(conf)
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	if err := decoder.Decode(confObj); err != nil {
		log.Fatal("decode config obj err:", err)
		os.Exit(1)
	}

	log.SetLogLevel(log.WARN)

	CCR_CLIENT, err = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Setup Complete")
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

func TestClient_ListInstances(t *testing.T) {
	args := &ListInstancesArgs{
		KeywordType: "clusterName",
		Keyword:     "",
		PageNo:      1,
		PageSize:    10,
	}
	resp, err := CCR_CLIENT.ListInstances(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_GetInstanceDetail(t *testing.T) {

	resp, err := CCR_CLIENT.GetInstanceDetail(CCR_INSTANCE_ID)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_ListPrivateNetworks(t *testing.T) {

	resp, err := CCR_CLIENT.ListPrivateNetworks(CCR_INSTANCE_ID)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

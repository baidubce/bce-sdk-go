package as

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	asClient *Client
	asConf   *Conf
)

type Conf struct {
	AK         string `json:"AK"`
	SK         string `json:"SK"`
	Endpoint   string `json:"Endpoint"`
	InstanceId string `json:"InstanceId"`
	UserId     string `json:"UserId"`
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	asConf = &Conf{}
	_ = decoder.Decode(asConf)

	asClient, _ = NewClient(asConf.AK, asConf.SK, asConf.Endpoint)
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

func TestClient_GetAsGroupList(t *testing.T) {
	req := &ListAsGroupRequest{
		GroupName: "test-name",
	}
	resp, err := asClient.ListAsGroup(req)
	fmt.Println(resp)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetAsGroupDetail(t *testing.T) {
	req := &GetAsGroupRequest{
		GroupId: "asg-wqksXo95",
	}
	resp, err := asClient.GetAsGroup(req)
	fmt.Println(resp)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetAsNodeList(t *testing.T) {
	req := &ListAsNodeRequest{
		GroupId: "asg-wqksXo95",
	}
	resp, err := asClient.ListAsNode(req)
	fmt.Println(resp)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_IncreaseAs(t *testing.T) {
	req := &IncreaseAsGroupRequest{
		GroupId:   "asg-Hhm2ucIK",
		Zone:      []string{"zoneB"},
		NodeCount: 1,
	}
	err := asClient.IncreaseAsGroup(req)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DecreaseAs(t *testing.T) {
	req := &DecreaseAsGroupRequest{
		GroupId: "asg-Hhm2ucIK",
		Nodes:   []string{"i-z0PXqFD3"},
	}
	err := asClient.DecreaseAsGroup(req)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_AdjustAs(t *testing.T) {
	req := &AdjustAsGroupRequest{
		GroupId:   "asg-Hhm2ucIK",
		AdjustNum: 1,
	}
	err := asClient.AdjustAsGroup(req)
	ExpectEqual(t.Errorf, nil, err)
}

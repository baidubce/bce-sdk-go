package hpas

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/hpas/api"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	HPAS_CLIENT *Client
	Hpas_id     string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	log.Error("config json file of ak/sk:", conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	HPAS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestCreateInstance(t *testing.T) {
	createInstanceArgs := &api.CreateHpasReq{
		AppType:             "llama2_7B_train_ic5",
		AppPerformanceLevel: "10k",
		Name:                "create_hpas_test",
		PurchaseNum:         1,
		ZoneName:            "cn-bj-a",
		ImageId:             "m-Xz6svNFM",
		SubnetId:            "sbn-s2haxxwvw8yi",
		SecurityGroupIds:    []string{"g-9vjwstn24c2v"},
		Password:            "71fa62c0059fa8624a4fbe110e236ab31ceede74cc7349df2f75f7ed2a279665",
		Tags:                []api.TagModel{{TagKey: "test1", TagValue: "test1"}},
	}
	createResult, err := HPAS_CLIENT.CreateHpas(createInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(createResult)
	Hpas_id = createResult.HpasIds[0]
}

func TestDeleteHpas(t *testing.T) {
	deleteInstanceArgs := &api.DeleteHpasReq{
		HpasIds: []string{"hpas-yQtvbIDe"},
	}
	err := HPAS_CLIENT.DeleteHpas(deleteInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestStopHpas(t *testing.T) {
	stopHpasArgs := &api.StopHpasReq{
		HpasIds:   []string{"hpas-0eGs9gbM"},
		ForceStop: true,
	}
	err := HPAS_CLIENT.StopHpas(stopHpasArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestStartHpas(t *testing.T) {
	startInstanceArgs := &api.StartHpasReq{
		HpasIds: []string{"hpas-0eGs9gbM"},
	}
	err := HPAS_CLIENT.StartHpas(startInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRebootHpas(t *testing.T) {
	rebootInstanceArgs := &api.RebootHpasReq{
		HpasIds:   []string{"hpas-0eGs9gbM"},
		ForceStop: true,
	}
	err := HPAS_CLIENT.RebootHpas(rebootInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestResetHpas(t *testing.T) {
	resetHpasArgs := &api.ResetHpasReq{
		HpasIds:  []string{"hpas-YuZJjrZ1"},
		ImageId:  "m-Xz6svNFM",
		Password: "71fa62c0059fa8624a4fbe110e236ab31ceede74cc7349df2f75f7ed2a279665",
	}
	err := HPAS_CLIENT.ResetHpas(resetHpasArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyPasswordHpas(t *testing.T) {
	modifyPasswordHpasArgs := &api.ModifyPasswordHpasReq{
		HpasId:   "hpas-0eGs9gbM",
		Password: "71fa62c0059fa8624a4fbe110e236ab31ceede74cc7349df2f75f7ed2a279665",
	}
	err := HPAS_CLIENT.ModifyPasswordHpas(modifyPasswordHpasArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateReservedHpas(t *testing.T) {
	createReservedHpasReq := &api.CreateReservedHpasReq{
		AppType:             "llama2_7B_train",
		AppPerformanceLevel: "10k",
		Name:                "create_hpas_test",
		PurchaseNum:         1,
		ZoneName:            "cn-bj-a",
		Tags:                []api.TagModel{{TagKey: "test1", TagValue: "test1"}},
	}
	createResult, err := HPAS_CLIENT.CreateReservedHpas(createReservedHpasReq)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(createResult)
	Hpas_id = createResult.ReservedHpasIds[0]
}

func TestDescribeReservedHpas(t *testing.T) {
	listReservedHpasPageReq := &api.ListReservedHpasPageReq{
		ReservedHpasIds: []string{"k-eZgiHDOY"},
		PageNo:          1,
		PageSize:        10,
	}
	describeResult, err := HPAS_CLIENT.DescribeReservedHpas(listReservedHpasPageReq)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(describeResult)
}

func TestListHpas(t *testing.T) {
	listHpasPageReq := &api.ListHpasPageReq{
		HpasIds:    []string{"hpas-K1ZPbS07"},
		Name:       "create_hpas_test",
		ZoneName:   "cn-bj-a",
		HpasStatus: "Active",
		AppType:    "llama2_7B_train_ic5",
		PageNo:     1,
		PageSize:   10,
	}
	listResult, err := HPAS_CLIENT.ListHpas(listHpasPageReq)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(listResult)
}

func TestImageList(t *testing.T) {
	body := &api.BaseMarkerV3Req{
		Marker:  "m-11111",
		MaxKeys: 2,
	}
	resp, err := HPAS_CLIENT.ImageList(body)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(resp)
}

func TestAttachTags(t *testing.T) {
	attachTagsArgs := &api.TagsOperationRequest{
		ResourceType: "hpasri",
		ResourceIds:  []string{"k-NEbO71Mz"},
		Tags: []api.TagModel{
			{
				TagKey:   "test1",
				TagValue: "test1",
			},
		},
	}
	err := HPAS_CLIENT.AttachTags(attachTagsArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDetachTags(t *testing.T) {
	attachTagsArgs := &api.TagsOperationRequest{
		ResourceType: "hpasri",
		ResourceIds:  []string{"k-NEbO71Mz"},
		Tags: []api.TagModel{
			{
				TagKey:   "test1",
				TagValue: "test1",
			},
		},
	}
	err := HPAS_CLIENT.DetachTags(attachTagsArgs)
	ExpectEqual(t.Errorf, err, nil)
}

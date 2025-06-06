package dev

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	AIHC_DEV_CLIENT *Client
)

const (
	AK       = "your ak"
	SK       = "your sk"
	ENDPOINT = "aihc.bj.baidubce.com"
)

// For security reason, ak/sk should not hard write here.

func init() {
	AIHC_DEV_CLIENT, _ = NewClient(AK, SK, ENDPOINT)
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

func ReadJson(fileName string) (*CreateDevInstanceArgs, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	var args CreateDevInstanceArgs
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&args)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &args, nil
}

func ToJSON(o interface{}) string {
	bs, _ := json.MarshalIndent(o, "", "\t")
	return string(bs)
}

func TestCreateDevInstance(t *testing.T) {
	args, err := ReadJson("you_json_file")
	if err != nil {
		t.Fatal(err)
		return
	}
	//fmt.Printf("create devinstance :%+v\n", args)
	res, err := AIHC_DEV_CLIENT.CreateDevInstance(args)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ToJSON(res))
	ExpectEqual(t.Errorf, err, nil)
}

func TestUpdateDevInstance(t *testing.T) {
	args, err := ReadJson("you_json_file")
	if err != nil {
		t.Fatal(err)
		return
	}
	//fmt.Printf("create devinstance :%+v\n", args)
	res, err := AIHC_DEV_CLIENT.UpdateDevInstance(args)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ToJSON(res))
	ExpectEqual(t.Errorf, err, nil)
}

func TestQueryDevInstaceDetail(t *testing.T) {
	args := &QueryDevInstanceDetailArgs{
		DevInstanceId: "your devInstanceId",
	}
	res, err := AIHC_DEV_CLIENT.QueryDevInstanceDetail(args)
	if err != nil {
		fmt.Println(err)
	}
	ExpectEqual(t.Errorf, err, nil)
	ret, _ := json.Marshal(res)
	fmt.Println(string(ret))
}

func TestListDevInstance(t *testing.T) {
	args := &ListDevInstanceArgs{
		QueryKey: "devInstanceId",
		QueryVal: "your devInstanceId",
		// ResourcePoolId: "cce-a7aphlyu",
		// PageNumber: 1,
		// PageSize:   10,
	}
	res, err := AIHC_DEV_CLIENT.ListDevInstance(args)
	if err != nil {
		fmt.Println(err)
	}
	ExpectEqual(t.Errorf, err, nil)
	ret, _ := json.Marshal(res)
	fmt.Println(string(ret))
}

func TestStopDevInstance(t *testing.T) {
	stopDevInstanceArgs := &StopDevInstanceArgs{
		DevInstanceId: "your devInstanceId",
	}
	res, err := AIHC_DEV_CLIENT.StopDevInstance(stopDevInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestStartDevInstance(t *testing.T) {
	startDevInstanceArgs := &StartDevInstanceArgs{
		DevInstanceId: "your devInstanceId",
	}
	res, err := AIHC_DEV_CLIENT.StartDevInstance(startDevInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestDeleteDevInstance(t *testing.T) {
	deleteDevInstanceArgs := &DeleteDevInstanceArgs{
		DevInstanceId: "your devInstanceId",
	}
	res, err := AIHC_DEV_CLIENT.DeleteDevInstance(deleteDevInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestTimedStopDevInstance(t *testing.T) {
	timedStopDevInstanceArgs := &TimedStopDevInstanceArgs{
		DevInstanceId: "",
		DelaySec:      3600,
		Enable:        false,
	}
	res, err := AIHC_DEV_CLIENT.TimedStopDevInstance(timedStopDevInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestCreateDevInstanceImagePackJob(t *testing.T) {
	args := &CreateDevInstanceImagePackJobArgs{
		DevInstanceID: "your devInstanceId",
		ImageName:     "image name",
		ImageTag:      "image tag",
		Namespace:     "registry namespace",
		Password:      "password",
		Registry:      "registry address",
		Username:      "username",
	}
	res, err := AIHC_DEV_CLIENT.CreateDevInstanceImagePackJob(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestDevInstanceImagePackJobDetail(t *testing.T) {
	args := &DevInstanceImagePackJobDetailArgs{
		ImagePackJobId: "your imagePackJobId",
		DevInstanceId:  "your devInstanceId",
	}
	res, err := AIHC_DEV_CLIENT.DevInstanceImagePackJobDetail(args)
	if err != nil {
		fmt.Println(err)
	}
	ExpectEqual(t.Errorf, err, nil)
	ret, _ := json.Marshal(res)
	fmt.Println(string(ret))
}

func TestListDevInstanceEvent(t *testing.T) {
	args := &ListDevInstanceEventArgs{
		DevInstanceId: "your devInstanceId",
		StartTime:     "2025-05-18T17:12:20.761Z",
		EndTime:       "2025-06-4T05:30:23.337Z",
	}
	res, err := AIHC_DEV_CLIENT.ListDevInstanceEvent(args)
	if err != nil {
		fmt.Println(err)
	}
	ExpectEqual(t.Errorf, err, nil)
	ret, _ := json.Marshal(res)
	fmt.Println(string(ret))
}

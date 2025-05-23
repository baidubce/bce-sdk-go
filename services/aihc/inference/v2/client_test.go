package api

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
	AIHC_INFERENCE_CLIENT *Client
)

const (
	AK       = "your ak"
	SK       = "your sk"
	ENDPOINT = "aihc.bj.baidubce.com"
)

// For security reason, ak/sk should not hard write here.

func init() {
	AIHC_INFERENCE_CLIENT, _ = NewClient(AK, SK, ENDPOINT)
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

func ToJSON(o interface{}) string {
	bs, _ := json.MarshalIndent(o, "", "\t")
	return string(bs)
}

func ReadJson(fileName string) (*ServiceConf, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	// 读取文件内容
	var serviceConf ServiceConf
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&serviceConf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &serviceConf, nil
}

func TestCreateService(t *testing.T) {
	createServiceArgs, err := ReadJson("you_json_file")
	if err != nil {
		t.Fatal(err)
		return
	}
	//fmt.Printf("create serviceConf :%+v\n", createServiceArgs)
	res, err := AIHC_INFERENCE_CLIENT.CreateService(createServiceArgs, "")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ToJSON(res))
	ExpectEqual(t.Errorf, err, nil)
}

func TestListService(t *testing.T) {
	listServiceArgs := &ListServiceArgs{
		PageSize:   10,
		PageNumber: 1,
		OrderBy:    "createdAt",
		Order:      "desc",
	}
	res, err := AIHC_INFERENCE_CLIENT.ListService(listServiceArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestListServiceStats(t *testing.T) {
	listServiceStatsArgs := &ListServiceStatsArgs{
		ServiceId: "you serviceId",
	}
	res, err := AIHC_INFERENCE_CLIENT.ListServiceStats(listServiceStatsArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestServiceDetails(t *testing.T) {
	serviceDetailsArgs := &ServiceDetailsArgs{
		ServiceId: "you serviceId",
	}
	res, err := AIHC_INFERENCE_CLIENT.ServiceDetails(serviceDetailsArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestUpdateService(t *testing.T) {
	serviceConf, err := ReadJson("you_json_file")
	if err != nil {
		t.Fatal(err)
		return
	}
	updateServiceArgs := &UpdateServiceArgs{
		ServiceId:   "you serviceId",
		ServiceConf: *serviceConf,
		Description: "you description",
	}
	res, err := AIHC_INFERENCE_CLIENT.UpdateService(updateServiceArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestScaleService(t *testing.T) {
	scaleServiceArgs := &ScaleServiceArgs{
		ServiceId:     "you serviceId",
		InstanceCount: 1,
	}
	res, err := AIHC_INFERENCE_CLIENT.ScaleService(scaleServiceArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestPubAccess(t *testing.T) {
	pubAccessArgs := &PubAccessArgs{
		ServiceId:    "you serviceId",
		PublicAccess: false,
		Eip:          "",
	}
	res, err := AIHC_INFERENCE_CLIENT.PubAccess(pubAccessArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestListChange(t *testing.T) {
	listChangeArgs := &ListChangeArgs{
		ServiceId: "you serviceId",
	}
	res, err := AIHC_INFERENCE_CLIENT.ListChange(listChangeArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestChangeDetail(t *testing.T) {
	changeDetailArgs := &ChangeDetailArgs{
		ChangeId: "you changeId",
	}
	res, err := AIHC_INFERENCE_CLIENT.ChangeDetail(changeDetailArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestDeleteService(t *testing.T) {
	deleteServiceArgs := &DeleteServiceArgs{
		ServiceId: "you serviceId",
	}
	res, err := AIHC_INFERENCE_CLIENT.DeleteService(deleteServiceArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestListPod(t *testing.T) {
	listPodArgs := &ListPodArgs{
		ServiceId: "you serviceId",
	}
	res, err := AIHC_INFERENCE_CLIENT.ListPod(listPodArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestBlockPod(t *testing.T) {
	blockPodArgs := &BlockPodArgs{
		ServiceId:  "you serviceId",
		InstanceId: "you instanceId",
		Block:      false,
	}
	res, err := AIHC_INFERENCE_CLIENT.BlockPod(blockPodArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestDeletePod(t *testing.T) {
	deletePodArgs := &DeletePodArgs{
		ServiceId:  "you serviceId",
		InstanceId: "you instanceId",
	}
	res, err := AIHC_INFERENCE_CLIENT.DeletePod(deletePodArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

func TestListPodGroups(t *testing.T) {
	listPodGroupArgs := &ListPodGroupsArgs{
		ServiceId: "you serviceId",
	}
	res, err := AIHC_INFERENCE_CLIENT.ListPodGroups(listPodGroupArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(ToJSON(res))
}

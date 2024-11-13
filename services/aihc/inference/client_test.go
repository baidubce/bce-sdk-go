package inference

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	AIHC_INFERENCE_CLIENT   *Client
	REGION                  string
	AIHC_INFERENCE_APPID    string
	AIHC_INFERENCE_CHANGEID string
	AIHC_INFERENCE_INSID    string
	AIHC_INFERENCE_RESPOOL  string
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
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	AIHC_INFERENCE_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	REGION = "bj"
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

func TestCreateApp(t *testing.T) {
	createAppArgs := &api.CreateAppArgs{
		AppName:  "testApp",
		ChipType: "baidu.com/test",
	}
	result, err := AIHC_INFERENCE_CLIENT.CreateApp(createAppArgs, REGION, nil)
	ExpectEqual(t.Errorf, err, nil)
	AIHC_INFERENCE_APPID = result.Data.AppId
}

func TestListApp(t *testing.T) {
	listAppArgs := &api.ListAppArgs{
		PageSize: 10,
		PageNo:   1,
	}
	_, err := AIHC_INFERENCE_CLIENT.ListApp(listAppArgs, REGION, nil)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListAppStats(t *testing.T) {
	listAppStatsArgs := &api.ListAppStatsArgs{
		AppIds: []string{AIHC_INFERENCE_APPID},
	}
	_, err := AIHC_INFERENCE_CLIENT.ListAppStats(listAppStatsArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

func TestAppDetails(t *testing.T) {
	appDetailsArgs := &api.AppDetailsArgs{
		AppId: AIHC_INFERENCE_APPID,
	}
	_, err := AIHC_INFERENCE_CLIENT.AppDetails(appDetailsArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUpdateApp(t *testing.T) {
	updateAppArgs := &api.UpdateAppArgs{
		AppId: AIHC_INFERENCE_APPID,
	}
	_, err := AIHC_INFERENCE_CLIENT.UpdateApp(updateAppArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

func TestScaleApp(t *testing.T) {
	scaleAppArgs := &api.ScaleAppArgs{
		AppId:    AIHC_INFERENCE_APPID,
		InsCount: 2,
	}
	_, err := AIHC_INFERENCE_CLIENT.ScaleApp(scaleAppArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

func TestPubAccess(t *testing.T) {
	pubAccessArgs := &api.PubAccessArgs{
		AppId:        AIHC_INFERENCE_APPID,
		PublicAccess: false,
	}
	_, err := AIHC_INFERENCE_CLIENT.PubAccess(pubAccessArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListChange(t *testing.T) {
	listChangeArgs := &api.ListChangeArgs{
		AppId: AIHC_INFERENCE_APPID,
	}
	result, err := AIHC_INFERENCE_CLIENT.ListChange(listChangeArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
	AIHC_INFERENCE_CHANGEID = result.Data.List[0].ChangeId
}

func TestChangeDetail(t *testing.T) {
	changeDetailArgs := &api.ChangeDetailArgs{
		ChangeId: AIHC_INFERENCE_CHANGEID,
	}
	_, err := AIHC_INFERENCE_CLIENT.ChangeDetail(changeDetailArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteApp(t *testing.T) {
	deleteAppArgs := &api.DeleteAppArgs{
		AppId: AIHC_INFERENCE_APPID,
	}
	_, err := AIHC_INFERENCE_CLIENT.DeleteApp(deleteAppArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListPod(t *testing.T) {
	listPodArgs := &api.ListPodArgs{
		AppId: AIHC_INFERENCE_APPID,
	}
	result, err := AIHC_INFERENCE_CLIENT.ListPod(listPodArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
	AIHC_INFERENCE_INSID = result.Data.List[0].InsID
}

func TestBlockPod(t *testing.T) {
	blockPodArgs := &api.BlockPodArgs{
		AppId: AIHC_INFERENCE_APPID,
		InsID: AIHC_INFERENCE_INSID,
		Block: true,
	}
	_, err := AIHC_INFERENCE_CLIENT.BlockPod(blockPodArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeletePod(t *testing.T) {
	deletePodArgs := &api.DeletePodArgs{
		AppId: AIHC_INFERENCE_APPID,
		InsID: AIHC_INFERENCE_INSID,
	}
	_, err := AIHC_INFERENCE_CLIENT.DeletePod(deletePodArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListBriefResPool(t *testing.T) {
	listBreifResPoolArgs := &api.ListBriefResPoolArgs{
		PageNo:   1,
		PageSize: 10,
	}
	result, err := AIHC_INFERENCE_CLIENT.ListBriefResPool(listBreifResPoolArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
	AIHC_INFERENCE_RESPOOL = result.Data.ResPools[0].ResPoolId
}

func TestResPoolDetail(t *testing.T) {
	resPoolDetailArgs := &api.ResPoolDetailArgs{
		ResPoolId: AIHC_INFERENCE_RESPOOL,
	}
	_, err := AIHC_INFERENCE_CLIENT.ResPoolDetail(resPoolDetailArgs, REGION)
	ExpectEqual(t.Errorf, err, nil)
}

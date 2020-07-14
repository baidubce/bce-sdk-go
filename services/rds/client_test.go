package rds

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/util"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	RDS_CLIENT *Client
	RDS_ID     string

	// set this value before start test
	ACCOUNT_NAME = "baidu"
	PASSWORD     = "baidu@123"
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

	RDS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateRds(t *testing.T) {
	args := &CreateRdsArgs{
		Engine:         "mysql",
		EngineVersion:  "5.6",
		CpuCount:       1,
		MemoryCapacity: 1,
		VolumeCapacity: 5,
		Billing: Billing{
			PaymentTiming: "Postpaid",
		},
		ClientToken: getClientToken(),
	}
	result, err := RDS_CLIENT.CreateRds(args)
	ExpectEqual(t.Errorf, nil, err)

	RDS_ID = result.InstanceIds[0]
	isAvailable(RDS_ID)
}

func TestClient_ResizeRds(t *testing.T) {
	args := &ResizeRdsArgs{
		CpuCount:       1,
		MemoryCapacity: 2,
		VolumeCapacity: 10,
	}
	err := RDS_CLIENT.ResizeRds(RDS_ID, args)
	ExpectEqual(t.Errorf, nil, err)
	time.Sleep(30*time.Second)
	isAvailable(RDS_ID)
}

func TestClient_ListRds(t *testing.T) {
	args := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(args)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if e.InstanceId == RDS_ID {
			ExpectEqual(t.Errorf, "MySQL", e.Engine)
			ExpectEqual(t.Errorf, "5.6", e.EngineVersion)
		}
	}
}

func TestClient_GetDetail(t *testing.T) {
	result, err := RDS_CLIENT.GetDetail(RDS_ID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "MySQL", result.Engine)
	ExpectEqual(t.Errorf, "5.6", result.EngineVersion)
}

func TestClient_CreateAccount(t *testing.T) {

	args := &CreateAccountArgs{
		AccountName: ACCOUNT_NAME,
		Password:    PASSWORD,
		ClientToken: getClientToken(),
	}

	isAvailable(RDS_ID)
	err := RDS_CLIENT.CreateAccount(RDS_ID,args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListAccount(t *testing.T) {
	isAvailable(RDS_ID)
	result, err := RDS_CLIENT.ListAccount(RDS_ID)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Accounts {
		if e.AccountName == ACCOUNT_NAME {
			ExpectEqual(t.Errorf, "Available", e.Status)
		}
	}
}

func TestClient_GetAccount(t *testing.T) {
	result, err := RDS_CLIENT.GetAccount(RDS_ID,ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Available", result.Status)
}

func TestClient_DeleteAccount(t *testing.T) {
	err := RDS_CLIENT.DeleteAccount(RDS_ID,ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateReadReplica(t *testing.T) {
	args := &CreateReadReplicaArgs{
		SourceInstanceId: RDS_ID,
		CpuCount:         1,
		MemoryCapacity:   2,
		VolumeCapacity:   10,
		Billing: Billing{
			PaymentTiming: "Postpaid",
		},
		ClientToken: getClientToken(),
	}
	time.Sleep(30*time.Second)
	isAvailable(RDS_ID)
	_, err := RDS_CLIENT.CreateReadReplica(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateRdsProxy(t *testing.T) {
	args := &CreateRdsProxyArgs{
		SourceInstanceId: RDS_ID,
		NodeAmount:       2,
		Billing: Billing{
			PaymentTiming: "Postpaid",
		},
		ClientToken: getClientToken(),
	}
	isAvailable(RDS_ID)
	_, err := RDS_CLIENT.CreateRdsProxy(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteRds(t *testing.T) {
	time.Sleep(30*time.Second)
	isAvailable(RDS_ID)
	err := RDS_CLIENT.DeleteRds(RDS_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	return util.NewUUID()
}

func isAvailable(instanceId string) {
	for {
		result, err := RDS_CLIENT.GetDetail(instanceId)
		if err == nil && result.InstanceStatus == "Available" {
			break
		}
	}
}
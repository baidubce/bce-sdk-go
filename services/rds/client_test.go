package rds

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/util"
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

const (
	SDK_NAME_PREFIX = "sdk_rds_"
)

func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 1; i++ {
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
	id := strconv.FormatInt(time.Now().Unix(), 10)
	args := &CreateRdsArgs{
		Engine:         "mysql",
		EngineVersion:  "5.6",
		Category:       "Standard",
		InstanceName:   SDK_NAME_PREFIX + id,
		CpuCount:       1,
		DiskIoType:     "normal_io",
		MemoryCapacity: 1,
		VolumeCapacity: 5,
		VpcId:          "vpc-it3vbqt3jhv",
		ZoneNames:      []string{"cn-bj-d"},
		Billing: Billing{
			PaymentTiming: "Prepaid",
			Reservation: Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
		ClientToken: getClientToken(),
		IsDirectPay: true,
	}
	result, err := RDS_CLIENT.CreateRds(args)
	ExpectEqual(t.Errorf, nil, err)

	RDS_ID = result.InstanceIds[0]
	fmt.Println("RDS: ", RDS_ID)
	// isAvailable(RDS_ID)
}

func TestClient_ResizeRds(t *testing.T) {
	args := &ResizeRdsArgs{
		CpuCount:       1,
		MemoryCapacity: 2,
		VolumeCapacity: 10,
	}
	err := RDS_CLIENT.ResizeRds(RDS_ID, args)
	ExpectEqual(t.Errorf, nil, err)
	time.Sleep(30 * time.Second)
	isAvailable(RDS_ID)
}

func TestClient_ListRds(t *testing.T) {
	args := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(args)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		fmt.Println(e.InstanceId)
		if e.InstanceId == RDS_ID {
			ExpectEqual(t.Errorf, "MySQL", e.Engine)
			ExpectEqual(t.Errorf, "5.6", e.EngineVersion)
		}
	}
}

func TestClient_GetDetail(t *testing.T) {
	result, err := RDS_CLIENT.GetDetail("rds-XNd7nidO")
	re, error := json.Marshal(result)
	fmt.Print(error)
	fmt.Println(string(re))
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
	err := RDS_CLIENT.CreateAccount(RDS_ID, args)
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
	result, err := RDS_CLIENT.GetAccount(RDS_ID, ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Available", result.Status)
}

func TestClient_DeleteAccount(t *testing.T) {
	err := RDS_CLIENT.DeleteAccount(RDS_ID, ACCOUNT_NAME)
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
	time.Sleep(30 * time.Second)
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

func TestClient_RebootInstance(t *testing.T) {
	isAvailable(RDS_ID)
	err := RDS_CLIENT.RebootInstance(RDS_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateInstanceName(t *testing.T) {
	isAvailable(RDS_ID)
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			args := &UpdateInstanceNameArgs{
				InstanceName: e.InstanceName + "_new",
			}
			err := RDS_CLIENT.UpdateInstanceName(e.InstanceId, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_ModifySyncMode(t *testing.T) {
	isAvailable(RDS_ID)
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			args := &ModifySyncModeArgs{
				SyncMode: "Async",
			}
			err := RDS_CLIENT.ModifySyncMode(e.InstanceId, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_ModifyEndpoint(t *testing.T) {
	isAvailable(RDS_ID)
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			args := &ModifyEndpointArgs{
				Address: "newsdkrds",
			}
			err := RDS_CLIENT.ModifyEndpoint(e.InstanceId, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_ModifyPublicAccess(t *testing.T) {
	isAvailable(RDS_ID)
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			args := &ModifyPublicAccessArgs{
				PublicAccess: false,
			}
			err := RDS_CLIENT.ModifyPublicAccess(e.InstanceId, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_GetBackupList(t *testing.T) {
	isAvailable("rds-ZLlMF0c3")
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			args := &GetBackupListArgs{}
			_, err := RDS_CLIENT.GetBackupList(e.InstanceId, args)
			ExpectEqual(t.Errorf, nil, err)
		}
		fmt.Println(e)
	}
}

func TestClient_GetZoneList(t *testing.T) {
	isAvailable(RDS_ID)
	_, err := RDS_CLIENT.GetZoneList()
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListSubnets(t *testing.T) {
	isAvailable(RDS_ID)
	args := &ListSubnetsArgs{}
	_, err := RDS_CLIENT.ListSubnets(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSecurityIps(t *testing.T) {
	//isAvailable(RDS_ID)
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			res, err := RDS_CLIENT.GetSecurityIps(e.InstanceId)
			fmt.Println(res.SecurityIps)
			fmt.Println(res.Etag)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_SecurityIps(t *testing.T) {
	//isAvailable(RDS_ID)
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			res, err := RDS_CLIENT.GetSecurityIps(e.InstanceId)
			ExpectEqual(t.Errorf, nil, err)
			args := &UpdateSecurityIpsArgs{
				SecurityIps: []string{
					"%",
					"192.0.0.1",
					"192.0.0.2",
				},
			}
			er := RDS_CLIENT.UpdateSecurityIps(e.InstanceId, res.Etag, args)
			ExpectEqual(t.Errorf, nil, er)
		}
	}
}

func TestClient_ListParameters(t *testing.T) {
	//isAvailable(RDS_ID)
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			res, err := RDS_CLIENT.ListParameters(e.InstanceId)
			data, _ := json.Marshal(res)
			fmt.Println(string(data))
			fmt.Println(res.Etag)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_UpdateParameter(t *testing.T) {
	//isAvailable(RDS_ID)
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			res, err := RDS_CLIENT.ListParameters(e.InstanceId)
			ExpectEqual(t.Errorf, nil, err)
			args := &UpdateParameterArgs{
				Parameters: []KVParameter{
					{
						Name:  "connect_timeout",
						Value: "15",
					},
				},
			}
			er := RDS_CLIENT.UpdateParameter(e.InstanceId, res.Etag, args)
			ExpectEqual(t.Errorf, nil, er)
		}
	}
}

func TestClient_DeleteRds(t *testing.T) {
	time.Sleep(30 * time.Second)
	isAvailable(RDS_ID)
	listRdsArgs := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Available" == e.InstanceStatus {
			err := RDS_CLIENT.DeleteRds(e.InstanceId)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
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

func TestClient_AutoRenew(t *testing.T) {
	err := RDS_CLIENT.AutoRenew(&AutoRenewArgs{
		AutoRenewTimeUnit: "month",
		AutoRenewTime:     1,
		InstanceIds: []string{
			"rds-rbmh6gJl",
		},
	})
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowLogDownloadTaskList(t *testing.T) {
	res, err := RDS_CLIENT.GetSlowLogDownloadTaskList("rdsmv5aumcrpynd", "2022-11-14T16:00:00Z")
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowLogDownloadDetail(t *testing.T) {
	res, err := RDS_CLIENT.GetSlowLogDownloadDetail("rds-qJG8sHPY", "slowlog.202211141158", "60")
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_newApi(t *testing.T) {
	result, err := RDS_CLIENT.Request("GET", "/v1/instance/rds-TSIlv3Sd/performance/processlist", nil)
	ExpectEqual(t.Errorf, nil, err)
	if result != nil {
		fmt.Println(result)
	}
}

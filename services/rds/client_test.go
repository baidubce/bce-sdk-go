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
	RDS_ID     = "rds-W3NYp4m9"
	ORDERID    string
	// set this value before start test
	ACCOUNT_NAME = "baidu"
	PASSWORD     = "xxxxxx"
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
		VpcId:          "vpc-it3v6qt3jhvj",
		ZoneNames:      []string{"cn-bj-d"},
		Subnets: []SubnetMap{
			{
				ZoneName: "cn-bj-d",
				SubnetId: "sbn-na4tmg4v11hs",
			},
		},
		Billing: Billing{
			PaymentTiming: "Postpaid",
		},
		ClientToken: getClientToken(),
		IsDirectPay: true,
	}
	result, err := RDS_CLIENT.CreateRds(args)

	ExpectEqual(t.Errorf, nil, err)

	RDS_ID = result.InstanceIds[0]
	ORDERID = result.OrderId
	fmt.Println("RDS: ", RDS_ID)
	fmt.Println("ORDERID: ", ORDERID)
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
	jsonData, err := json.Marshal(result)
	fmt.Println(string(jsonData))
	for _, e := range result.Instances {
		if e.InstanceId == RDS_ID {
			ExpectEqual(t.Errorf, "MySQL", e.Engine)
			ExpectEqual(t.Errorf, "5.6", e.EngineVersion)
		}
	}
}

func TestClient_GetDetail(t *testing.T) {
	result, err := RDS_CLIENT.GetDetail("rds-yZ4qXER9")
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

func TestClient_ModifyAccountDesc(t *testing.T) {
	args := &ModifyAccountDesc{
		Remark: "test",
	}
	err := RDS_CLIENT.ModifyAccountDesc(RDS_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAccount(t *testing.T) {
	err := RDS_CLIENT.DeleteAccount(RDS_ID, ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAccountPassword(t *testing.T) {
	args := &UpdatePasswordArgs{
		Password: "test",
	}
	err := RDS_CLIENT.UpdateAccountPassword(RDS_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_UpdateAccountPrivileges(t *testing.T) {
	args := &UpdateAccountPrivileges{
		DatabasePrivileges: []DatabasePrivilege{{
			DbName:   "test_db",
			AuthType: "ReadOnly",
		}},
	}
	err := RDS_CLIENT.UpdateAccountPrivileges(RDS_ID, ACCOUNT_NAME, args)
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

func TestClient_ModifyBackupPolicy(t *testing.T) {
	isAvailable(RDS_ID)
	modifyBackupPolicyArgs := &ModifyBackupPolicyArgs{
		ExpireInDays: 10,
		BackupDays:   "0,1,2,3,4,5,6",
		BackupTime:   "17:00:00Z",
		Persistent:   true,
	}
	err := RDS_CLIENT.ModifyBackupPolicy(RDS_ID, modifyBackupPolicyArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetBackupList(t *testing.T) {
	isAvailable(RDS_ID)
	getBackupListArgs := &GetBackupListArgs{
		Marker:  RDS_ID,
		MaxKeys: 100,
	}
	result, err := RDS_CLIENT.GetBackupList(RDS_ID, getBackupListArgs)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetBackupDetail(t *testing.T) {
	isAvailable(RDS_ID)
	result, err := RDS_CLIENT.GetBackupDetail(RDS_ID, "1691679661534780403")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteBackup(t *testing.T) {
	isAvailable(RDS_ID)
	err := RDS_CLIENT.DeleteBackup(RDS_ID, "1691734023130272802")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetBinlogList(t *testing.T) {
	result, err := RDS_CLIENT.GetBinlogList(RDS_ID, "2022-07-12T23:59:59Z")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetBinlogInfo(t *testing.T) {
	result, err := RDS_CLIENT.GetBinlogInfo(RDS_ID, "1691734023130272802", "1800")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RecoveryToSourceInstanceByDatetime(t *testing.T) {
	isAvailable(RDS_ID)
	recoveryByDatetimeArgs := &RecoveryByDatetimeArgs{
		Datetime: "2022-01-11T16:05:52Z",
		Data: []RecoveryData{
			{
				DbName:      "test_db",
				NewDbname:   "new_test_db",
				RestoreMode: "database",
				Tables: []TableData{
					{
						TableName:    "table_name",
						NewTablename: "new_table_name",
					},
				},
			},
		},
	}
	err := RDS_CLIENT.RecoveryToSourceInstanceByDatetime(RDS_ID, recoveryByDatetimeArgs)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RecoveryToSourceInstanceBySnapshot(t *testing.T) {
	isAvailable(RDS_ID)
	recoveryBySnapshotArgs := &RecoveryBySnapshotArgs{
		SnapshotId: "1691734023130272802",
		Data: []RecoveryData{
			{
				DbName:      "test_db",
				NewDbname:   "new_test_db",
				RestoreMode: "database",
				Tables: []TableData{
					{
						TableName:    "table_name",
						NewTablename: "new_table_name",
					},
				},
			},
		},
	}
	err := RDS_CLIENT.RecoveryToSourceInstanceBySnapshot(RDS_ID, recoveryBySnapshotArgs)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetZoneList(t *testing.T) {
	isAvailable(RDS_ID)
	_, err := RDS_CLIENT.GetZoneList()
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListSubnets(t *testing.T) {
	isAvailable(RDS_ID)
	args := &ListSubnetsArgs{
		VpcId:    "vpc-it3v6qt3jhvj",
		ZoneName: "cn-bj-d",
	}
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
	listRdsArgs := &ListRdsArgs{
		Marker:  "0",
		MaxKeys: 100,
	}
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

func TestClient_ParameterHistory(t *testing.T) {
	result, err := RDS_CLIENT.ParameterHistory(RDS_ID)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
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

func TestClient_UpdateMaintainTime(t *testing.T) {
	err := RDS_CLIENT.UpdateMaintainTime(RDS_ID, &MaintainTimeArgs{
		MaintainStartTime: "14:00:00",
		MaintainDuration:  2,
	})
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ConfigDiskAutoResize(t *testing.T) {
	err := RDS_CLIENT.ConfigDiskAutoResize(RDS_ID, "open", &DiskAutoResizeArgs{
		FreeSpaceThreshold: 10,
		DiskMaxLimit:       2000,
	})
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetAutoResizeConfig(t *testing.T) {
	res, err := RDS_CLIENT.GetAutoResizeConfig(RDS_ID)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_EnableAutoExpansion(t *testing.T) {
	res, err := RDS_CLIENT.EnableAutoExpansion(RDS_ID)
	fmt.Print(res)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_AzoneMigration(t *testing.T) {
	err := RDS_CLIENT.AzoneMigration(RDS_ID, &AzoneMigration{
		MasterAzone: "cn-bj-d",
		BackupAzone: "cn-bj-e",
		ZoneNames:   []string{"cn-bj-d", "cn-bj-e"},
		Subnets: []SubnetMap{
			{
				ZoneName: "cn-bj-d",
				SubnetId: "sbn-nedt51qre6r2",
			},
			{
				ZoneName: "cn-bj-e",
				SubnetId: "sbn-hc20wss3idai",
			},
		},
		EffectiveTime: "timewindow",
	})
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateDatabasePort(t *testing.T) {
	err := RDS_CLIENT.UpdateDatabasePort(RDS_ID, &UpdateDatabasePortArgs{
		EntryPort: 3309,
	})
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListDatabases(t *testing.T) {
	result, err := RDS_CLIENT.ListDatabases(RDS_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ModifyDatabaseDesc(t *testing.T) {
	args := &ModifyDatabaseDesc{
		Remark: "test",
	}
	err := RDS_CLIENT.ModifyDatabaseDesc(RDS_ID, "test_db", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteDatabase(t *testing.T) {
	err := RDS_CLIENT.DeleteDatabase(RDS_ID, "test_db")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateDatabase(t *testing.T) {

	args := &CreateDatabaseArgs{
		CharacterSetName: "utf8",
		DbName:           "test_db",
		Remark:           "test_db",
		AccountPrivileges: []AccountPrivilege{
			{
				AccountName: "baidu",
				AuthType:    "ReadOnly",
			},
		},
	}

	isAvailable(RDS_ID)
	err := RDS_CLIENT.CreateDatabase(RDS_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_TaskList(t *testing.T) {
	args := &TaskListArgs{
		InstanceId: RDS_ID,
	}
	result, err := RDS_CLIENT.TaskList(args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListRecyclerInstance(t *testing.T) {
	args := &ListRdsArgs{}
	result, err := RDS_CLIENT.ListRecyclerInstance(args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RecyclerRecover(t *testing.T) {
	args := &RecyclerRecoverArgs{
		InstanceIds: []string{RDS_ID},
	}
	err := RDS_CLIENT.RecyclerRecover(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteRecyclerInstance(t *testing.T) {
	err := RDS_CLIENT.DeleteRecyclerInstance(RDS_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateInstanceGroup(t *testing.T) {
	args := &InstanceGroupArgs{
		Name:     "test_group",
		LeaderId: RDS_ID,
	}
	result, err := RDS_CLIENT.CreateInstanceGroup(args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListInstanceGroup(t *testing.T) {
	args := &ListInstanceGroupArgs{
		Manner: "page",
	}
	result, err := RDS_CLIENT.ListInstanceGroup(args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupDetail(t *testing.T) {
	result, err := RDS_CLIENT.InstanceGroupDetail("rdcg6034psv")
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupCheckGtid(t *testing.T) {
	args := &CheckGtidArgs{
		InstanceId: RDS_ID,
	}
	result, err := RDS_CLIENT.InstanceGroupCheckGtid(args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupCheckPing(t *testing.T) {
	args := &CheckPingArgs{
		SourceId: RDS_ID,
		TargetId: RDS_ID,
	}
	result, err := RDS_CLIENT.InstanceGroupCheckPing(args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupCheckData(t *testing.T) {
	args := &CheckDataArgs{
		InstanceId: RDS_ID,
	}
	result, err := RDS_CLIENT.InstanceGroupCheckData(args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupCheckVersion(t *testing.T) {
	args := &CheckVersionArgs{
		LeaderId:   RDS_ID,
		FollowerId: RDS_ID,
	}
	result, err := RDS_CLIENT.InstanceGroupCheckVersion(args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateInstanceGroupName(t *testing.T) {
	args := &InstanceGroupNameArgs{
		Name: "test_group_name",
	}
	err := RDS_CLIENT.UpdateInstanceGroupName("rdcg6034psv", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupAdd(t *testing.T) {
	args := &InstanceGroupAddArgs{
		FollowerId: RDS_ID,
	}
	err := RDS_CLIENT.InstanceGroupAdd("rdcg6034psv", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupBatchAdd(t *testing.T) {
	args := &InstanceGroupBatchAddArgs{
		FollowerIds: []string{RDS_ID},
		Name:        "test_group_name",
		LeaderId:    RDS_ID,
	}
	err := RDS_CLIENT.InstanceGroupBatchAdd(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupForceChange(t *testing.T) {
	args := &ForceChangeArgs{
		LeaderId: RDS_ID,
		Force:    0,
	}
	result, err := RDS_CLIENT.InstanceGroupForceChange("rdcg6034psv", args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupLeaderChange(t *testing.T) {
	args := &GroupLeaderChangeArgs{
		LeaderId: RDS_ID,
	}
	err := RDS_CLIENT.InstanceGroupLeaderChange("rdcg6034psv", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceGroupRemove(t *testing.T) {
	err := RDS_CLIENT.InstanceGroupRemove("rdcg6034psv", RDS_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteInstanceGroup(t *testing.T) {
	err := RDS_CLIENT.DeleteInstanceGroup("rdcg6034psv")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceMinorVersionList(t *testing.T) {
	result, err := RDS_CLIENT.InstanceMinorVersionList(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InstanceUpgradeMinorVersion(t *testing.T) {
	args := &UpgradeMinorVersionArgs{
		TargetMinorVersion: "5.7.38",
		EffectiveTime:      "immediate",
	}
	err := RDS_CLIENT.InstanceUpgradeMinorVersion(RDS_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_SlowSqlFlowStatus(t *testing.T) {
	result, err := RDS_CLIENT.SlowSqlFlowStatus(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_EnableSlowSqlFlow(t *testing.T) {
	err := RDS_CLIENT.EnableSlowSqlFlow(RDS_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DisableSlowSqlFlow(t *testing.T) {
	err := RDS_CLIENT.DisableSlowSqlFlow(RDS_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlList(t *testing.T) {
	args := &GetSlowSqlArgs{}
	result, err := RDS_CLIENT.GetSlowSqlList(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlBySqlId(t *testing.T) {
	result, err := RDS_CLIENT.GetSlowSqlBySqlId(RDS_ID, "sqlidxxx")
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlExplain(t *testing.T) {
	result, err := RDS_CLIENT.GetSlowSqlExplain(RDS_ID, "sqlidxxx", "db1")
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlStatsDigest(t *testing.T) {
	args := &GetSlowSqlArgs{}
	result, err := RDS_CLIENT.GetSlowSqlStatsDigest(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlDuration(t *testing.T) {
	args := &GetSlowSqlDurationArgs{}
	result, err := RDS_CLIENT.GetSlowSqlDuration(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlSource(t *testing.T) {
	args := &GetSlowSqlSourceArgs{}
	result, err := RDS_CLIENT.GetSlowSqlSource(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlSchema(t *testing.T) {
	result, err := RDS_CLIENT.GetSlowSqlSchema(RDS_ID, "sqlidxxx", "db1")
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlTable(t *testing.T) {
	result, err := RDS_CLIENT.GetSlowSqlTable(RDS_ID, "sqlidxxx", "db1", "table1")
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlIndex(t *testing.T) {
	args := &GetSlowSqlIndexArgs{
		SqlId:  "e9fa9802-0d0e-41b4-b3ba-6496466b6cad",
		Schema: "db1",
		Table:  "table1",
	}
	result, err := RDS_CLIENT.GetSlowSqlIndex(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlTrend(t *testing.T) {
	args := &GetSlowSqlTrendArgs{
		Start: "2023-05-05T05:30:13.000Z",
		End:   "2023-05-06T05:30:13.000Z",
	}
	result, err := RDS_CLIENT.GetSlowSqlTrend(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlAdvice(t *testing.T) {
	result, err := RDS_CLIENT.GetSlowSqlAdvice(RDS_ID, "sqlidxxx", "db1")
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetDiskInfo(t *testing.T) {
	result, err := RDS_CLIENT.GetDiskInfo(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetDbListSize(t *testing.T) {
	result, err := RDS_CLIENT.GetDbListSize(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetTableListInfo(t *testing.T) {
	args := &GetTableListArgs{
		DbName: "db1",
	}
	result, err := RDS_CLIENT.GetTableListInfo(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetKillSessionTypes(t *testing.T) {
	result, err := RDS_CLIENT.GetKillSessionTypes(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSessionSummary(t *testing.T) {
	result, err := RDS_CLIENT.GetSessionSummary(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSessionDetail(t *testing.T) {
	args := &SessionDetailArgs{}
	result, err := RDS_CLIENT.GetSessionDetail(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CheckKillSessionAuth(t *testing.T) {
	args := &KillSessionAuthArgs{}
	result, err := RDS_CLIENT.CheckKillSessionAuth(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetKillSessionHistory(t *testing.T) {
	args := &KillSessionHistory{}
	result, err := RDS_CLIENT.GetKillSessionHistory(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_KillSession(t *testing.T) {
	args := &KillSessionArgs{}
	result, err := RDS_CLIENT.KillSession(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSessionStatistics(t *testing.T) {
	result, err := RDS_CLIENT.GetSessionStatistics(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetErrorLogStatus(t *testing.T) {
	result, err := RDS_CLIENT.GetErrorLogStatus(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_EnableErrorLog(t *testing.T) {
	result, err := RDS_CLIENT.EnableErrorLog(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DisableErrorLog(t *testing.T) {
	result, err := RDS_CLIENT.DisableErrorLog(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetErrorLogList(t *testing.T) {
	args := &ErrorLogListArgs{}
	result, err := RDS_CLIENT.GetErrorLogList(RDS_ID, args)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSqlFilterList(t *testing.T) {
	result, err := RDS_CLIENT.GetSqlFilterList(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSqlFilterDetail(t *testing.T) {
	result, err := RDS_CLIENT.GetSqlFilterDetail(RDS_ID, "83")
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_AddSqlFilter(t *testing.T) {
	args := &SqlFilterArgs{
		FilterType:  "SELECT",
		FilterKey:   "123",
		FilterLimit: 0,
	}
	err := RDS_CLIENT.AddSqlFilter(RDS_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateSqlFilter(t *testing.T) {
	args := &SqlFilterArgs{
		FilterType:  "SELECT",
		FilterKey:   "1234",
		FilterLimit: 0,
	}
	err := RDS_CLIENT.UpdateSqlFilter(RDS_ID, "83", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_StartOrStopSqlFilter(t *testing.T) {
	args := &StartOrStopSqlFilterArgs{
		Action: "OFF",
	}
	err := RDS_CLIENT.StartOrStopSqlFilter(RDS_ID, "83", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteSqlFilter(t *testing.T) {
	err := RDS_CLIENT.DeleteSqlFilter(RDS_ID, "83")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_IsAllowedSqlFilter(t *testing.T) {
	result, err := RDS_CLIENT.IsAllowedSqlFilter(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ProcessKill(t *testing.T) {
	args := &ProcessArgs{
		Ids: []int64{123},
	}
	err := RDS_CLIENT.ProcessKill(RDS_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_InnodbStatus(t *testing.T) {
	result, err := RDS_CLIENT.InnodbStatus(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ProcessList(t *testing.T) {
	result, err := RDS_CLIENT.ProcessList(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_TransactionList(t *testing.T) {
	result, err := RDS_CLIENT.TransactionList(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ConnectionList(t *testing.T) {
	result, err := RDS_CLIENT.ConnectionList(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_FailInjectWhiteList(t *testing.T) {
	result, err := RDS_CLIENT.FailInjectWhiteList()
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_AddToFailInjectWhiteList(t *testing.T) {
	args := &FailInjectArgs{
		AppList: []string{RDS_ID},
	}
	err := RDS_CLIENT.AddToFailInjectWhiteList(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RemoveFailInjectWhiteList(t *testing.T) {
	args := &FailInjectArgs{
		AppList: []string{RDS_ID},
	}
	err := RDS_CLIENT.RemoveFailInjectWhiteList(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_FailInjectStart(t *testing.T) {
	result, err := RDS_CLIENT.FailInjectStart(RDS_ID)
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetOrderStatus(t *testing.T) {
	result, err := RDS_CLIENT.GetOrderStatus("xxx")
	jsonData, _ := json.Marshal(result)
	fmt.Println(string(jsonData))
	ExpectEqual(t.Errorf, nil, err)
}

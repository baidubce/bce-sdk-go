package gaiadb

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	GAIADB_CLIENT *Client
	GAIADB_ID     = "gaiadbwjeq27"
	ORDERID       string
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
	SDK_NAME_PREFIX = "sdk_gaiadb_"
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

	GAIADB_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func getClientToken() string {
	return util.NewUUID()
}
func TestClient_CreateGaiadb(t *testing.T) {

	args := &CreateClusterArgs{
		ClientToken: getClientToken(),
		Number:      1,
		ProductType: "postpay",
		InstanceParam: InstanceParam{
			ReleaseVersion:       "8.0",
			SubnetId:             "sbn-na4tmg4v11hs",
			AllocatedCpuInCore:   2,
			AllocatedMemoryInMB:  8192,
			AllocatedStorageInGB: 5120,
			VpcId:                "vpc-it3v6qt3jhvj",
			InstanceAmount:       2,
			ProxyAmount:          2,
		},
	}
	result, err := GAIADB_CLIENT.CreateCluster(args)

	ExpectEqual(t.Errorf, nil, err)

	GAIADB_ID = result.ClusterIds[0]
	ORDERID = result.OrderId
	fmt.Println("GAIADB: ", GAIADB_ID)
	fmt.Println("ORDERID: ", ORDERID)
}

func TestClient_DeleteCluster(t *testing.T) {
	err := GAIADB_CLIENT.DeleteCluster(GAIADB_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RenameCluster(t *testing.T) {
	args := &ClusterName{
		ClusterName: "cluster_test",
	}
	err := GAIADB_CLIENT.RenameCluster(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ResizeCluster(t *testing.T) {
	args := &ResizeClusterArgs{
		ResizeType:          "resizeSlave",
		AllocatedCpuInCore:  4,
		AllocatedMemoryInMB: 8192,
	}
	result, err := GAIADB_CLIENT.ResizeCluster(GAIADB_ID, args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetClusterList(t *testing.T) {
	args := &Marker{
		Marker:  "-1",
		MaxKeys: 1000,
	}
	result, err := GAIADB_CLIENT.GetClusterList(args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetClusterDetail(t *testing.T) {
	result, err := GAIADB_CLIENT.GetClusterDetail(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetClusterCapacity(t *testing.T) {
	result, err := GAIADB_CLIENT.GetClusterCapacity(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_QueryClusterPrice(t *testing.T) {
	args := &QueryPriceArgs{
		Number: 1,
		InstanceParam: InstanceInfo{
			ReleaseVersion:       "8.0",
			AllocatedCpuInCore:   2,
			AllocatedMemoryInMB:  8192,
			AllocatedStorageInGB: 5120,
			InstanceAmount:       2,
			ProxyAmount:          2,
		},
		ProductType: "postpay",
	}
	result, err := GAIADB_CLIENT.QueryClusterPrice(args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_QueryResizeClusterPrice(t *testing.T) {
	args := &QueryResizePriceArgs{
		ClusterId:           GAIADB_ID,
		ResizeType:          "resizeSlave",
		AllocatedCpuInCore:  2,
		AllocatedMemoryInMB: 8192,
	}
	result, err := GAIADB_CLIENT.QueryResizeClusterPrice(args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RebootInstance(t *testing.T) {
	args := &RebootInstanceArgs{
		ExecuteAction: "executeNow",
	}
	err := GAIADB_CLIENT.RebootInstance(GAIADB_ID, "gaiadbm5h6ys-secondary-129aafc0", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_BindTags(t *testing.T) {
	args := &BindTagsArgs{
		Resources: []Resource{
			{
				ResourceId: GAIADB_ID,
				Tags: []Tag{
					{
						TagKey:   "testTagKey",
						TagValue: "testTagValue",
					},
				},
			},
		},
	}
	err := GAIADB_CLIENT.BindTags(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ClusterSwitch(t *testing.T) {
	args := &ClusterSwitchArgs{
		ExecuteAction:       "executeNow",
		SecondaryInstanceId: GAIADB_ID,
	}
	result, err := GAIADB_CLIENT.ClusterSwitch(GAIADB_ID, args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetInterfaceList(t *testing.T) {
	result, err := GAIADB_CLIENT.GetInterfaceList(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateDnsName(t *testing.T) {
	args := &UpdateDnsNameArgs{
		InterfaceId: "gaiadbm5h6ys_interface0000",
		DnsName:     "my.gaiadb.bj.baidubce.com",
	}
	err := GAIADB_CLIENT.UpdateDnsName(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateInterface(t *testing.T) {
	args := &UpdateInterfaceArgs{
		InterfaceId: "gaiadbm5h6ys_interface0000",
		Interface: InterfaceInfo{
			MasterReadable: 1,
			AddressName:    "addressname",
			InstanceBinding: []string{
				"gaiadbymbrc8-primary-6f1cc3a2",
				"gaiadbymbrc8-secondary-ec909467",
			},
		},
	}
	err := GAIADB_CLIENT.UpdateInterface(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_NewInstanceAutoJoin(t *testing.T) {
	args := &NewInstanceAutoJoinArgs{
		AutoJoinRequestItems: []AutoJoinRequestItem{
			{
				NewInstanceAutoJoin: "off",
				InterfaceId:         "gaiadbymbrc8-primary-6f1cc3a2",
			},
		},
	}
	err := GAIADB_CLIENT.NewInstanceAutoJoin(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateAccount(t *testing.T) {
	args := &CreateAccountArgs{
		AccountName: "testaccount",
		Password:    "xxxxxx",
		AccountType: "common",
		Remark:      "testRemark",
	}
	err := GAIADB_CLIENT.CreateAccount(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAccount(t *testing.T) {
	err := GAIADB_CLIENT.DeleteAccount(GAIADB_ID, "testaccount")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetAccountDetail(t *testing.T) {
	result, err := GAIADB_CLIENT.GetAccountDetail(GAIADB_ID, "testaccount")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetAccountList(t *testing.T) {
	result, err := GAIADB_CLIENT.GetAccountList(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAccountRemark(t *testing.T) {
	args := &RemarkArgs{
		Remark: "remark",
		Etag:   "v0",
	}
	err := GAIADB_CLIENT.UpdateAccountRemark(GAIADB_ID, "testaccount", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAccountAuthIp(t *testing.T) {
	args := &AuthIpArgs{
		Action: "ipAdd",
		Value: AuthIp{
			Authip:  []string{"10.10.10.10"},
			Authbns: []string{},
		},
	}
	err := GAIADB_CLIENT.UpdateAccountAuthIp(GAIADB_ID, "testaccount", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAccountPrivileges(t *testing.T) {
	args := &PrivilegesArgs{
		DatabasePrivileges: []DatabasePrivilege{
			{
				DbName:     "testdb",
				AuthType:   "definePrivilege",
				Privileges: []string{"UPDATE"},
			},
		},
		Etag: "v0",
	}
	err := GAIADB_CLIENT.UpdateAccountPrivileges(GAIADB_ID, "testaccount", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAccountPassword(t *testing.T) {
	args := &PasswordArgs{
		Password: "testpassword",
		Etag:     "v0",
	}
	err := GAIADB_CLIENT.UpdateAccountPassword(GAIADB_ID, "testaccount", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateDatabase(t *testing.T) {
	args := &CreateDatabaseArgs{
		DbName:           "test_db",
		CharacterSetName: "utf8",
		Remark:           "sdk test",
	}
	err := GAIADB_CLIENT.CreateDatabase(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteDatabase(t *testing.T) {
	err := GAIADB_CLIENT.DeleteDatabase(GAIADB_ID, "test_db")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListDatabase(t *testing.T) {
	result, err := GAIADB_CLIENT.ListDatabase(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_CreateSnapshot(t *testing.T) {
	err := GAIADB_CLIENT.CreateSnapshot(GAIADB_ID)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_ListSnapshot(t *testing.T) {
	result, err := GAIADB_CLIENT.ListSnapshot(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateSnapshotPolicy(t *testing.T) {
	args := &UpdateSnapshotPolicyArgs{
		DataBackupWeekDay: []string{"Monday"},
		DataBackupRetainStrategys: []DataBackupRetainStrategy{{
			StartSeconds: 0,
			RetainCount:  "8",
			Precision:    86400,
			EndSeconds:   -691200,
		}},
		DataBackupTime: "02:00:00Z",
	}
	err := GAIADB_CLIENT.UpdateSnapshotPolicy(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetSnapshotPolicy(t *testing.T) {
	result, err := GAIADB_CLIENT.GetSnapshotPolicy(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateWhiteList(t *testing.T) {
	args := &WhiteList{
		AuthIps: []string{"192.168.1.2"},
		Etag:    "v0",
	}
	err := GAIADB_CLIENT.UpdateWhiteList(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetWhiteList(t *testing.T) {
	result, err := GAIADB_CLIENT.GetWhiteList(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateMultiactiveGroup(t *testing.T) {
	args := &CreateMultiactiveGroupArgs{
		LeaderClusterId:      GAIADB_ID,
		MultiActiveGroupName: "test_multiactive_group",
	}
	result, err := GAIADB_CLIENT.CreateMultiactiveGroup(args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteMultiactiveGroup(t *testing.T) {
	err := GAIADB_CLIENT.DeleteMultiactiveGroup("gaiagroup-fru5yw")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RenameMultiactiveGroup(t *testing.T) {
	args := &RenameMultiactiveGroupArgs{
		MultiActiveGroupName: "test_multiactive_group",
	}
	err := GAIADB_CLIENT.RenameMultiactiveGroup("gaiagroup-5r5aur", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_MultiactiveGroupList(t *testing.T) {
	result, err := GAIADB_CLIENT.MultiactiveGroupList()
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_MultiactiveGroupDetail(t *testing.T) {
	result, err := GAIADB_CLIENT.MultiactiveGroupDetail("gaiagroup-0luzwo")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSyncStatus(t *testing.T) {
	result, err := GAIADB_CLIENT.GetSyncStatus("gaiagroup-0luzwo", GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GroupExchange(t *testing.T) {
	args := &ExchangeArgs{
		ExecuteAction:      "executeNow",
		NewLeaderClusterId: GAIADB_ID,
	}
	err := GAIADB_CLIENT.GroupExchange("gaiagroup-0luzwo", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetParamsList(t *testing.T) {
	result, err := GAIADB_CLIENT.GetParamsList(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetParamsHistory(t *testing.T) {
	result, err := GAIADB_CLIENT.GetParamsHistory(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateParams(t *testing.T) {
	args := &UpdateParamsArgs{
		Params: map[string]interface{}{
			"auto_increment_increment": "5",
		},
		Timing: "now",
	}
	err := GAIADB_CLIENT.UpdateParams(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListParamTemplate(t *testing.T) {
	args := &ListParamTempArgs{
		Detail:   0,
		Type:     "mysql",
		PageNo:   1,
		PageSize: 10,
	}
	result, err := GAIADB_CLIENT.ListParamTemplate(args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_SaveAsParamTemplate(t *testing.T) {
	args := &ParamTempArgs{
		Type:        "mysql",
		Version:     "8.0",
		Description: "create by sdk",
		Name:        "sdk_test",
		Source:      GAIADB_ID,
	}
	err := GAIADB_CLIENT.SaveAsParamTemplate(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetTemplateApplyRecords(t *testing.T) {
	result, err := GAIADB_CLIENT.GetTemplateApplyRecords("ce8a8245-1d9b-c844-bff8-818933461baa")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteParamsFromTemp(t *testing.T) {
	args := &Params{
		Params: []string{"long_query_time"},
	}
	err := GAIADB_CLIENT.DeleteParamsFromTemp("ce8a8245-1d9b-c844-bff8-818933461baa", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateParamTemplate(t *testing.T) {
	args := &UpdateParamTplArgs{
		Name:        "test_template",
		Description: "test_template_description",
	}
	err := GAIADB_CLIENT.UpdateParamTemplate("ce8a8245-1d9b-c844-bff8-818933461baa", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ModifyParams(t *testing.T) {
	args := &ModifyParamsArgs{
		Params: map[string]interface{}{
			"auto_increment_increment": "5",
			"long_query_time":          "6.6",
		},
	}
	err := GAIADB_CLIENT.ModifyParams("ce8a8245-1d9b-c844-bff8-818933461baa", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteParamTemplate(t *testing.T) {
	err := GAIADB_CLIENT.DeleteParamTemplate("ce8a8245-1d9b-c844-bff8-818933461baa")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateParamTemplate(t *testing.T) {
	args := &CreateParamTemplateArgs{
		Name:        "test_template",
		Type:        "mysql",
		Version:     "8.0",
		Description: "test_template_description",
	}
	err := GAIADB_CLIENT.CreateParamTemplate(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetParamTemplateDetail(t *testing.T) {
	result, err := GAIADB_CLIENT.GetParamTemplateDetail("ce8a8245-1d9b-c844-bff8-818933461baa", "0")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetParamTemplateHistory(t *testing.T) {
	result, err := GAIADB_CLIENT.GetParamTemplateHistory("ce8a8245-1d9b-c844-bff8-818933461baa", "addParam")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ApplyParamTemplate(t *testing.T) {
	args := &ApplyParamTemplateArgs{
		Timing: "now",
		Clusters: map[string]interface{}{
			"gaiadbk3pyxv": []interface{}{},
		},
	}
	err := GAIADB_CLIENT.ApplyParamTemplate("ce8a8245-1d9b-c844-bff8-818933461baa", args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateMaintenTime(t *testing.T) {
	args := &UpdateMaintenTimeArgs{
		Period:    "1,2,3",
		StartTime: "03:00",
		Duration:  1,
	}
	err := GAIADB_CLIENT.UpdateMaintenTime(GAIADB_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetMaintenTime(t *testing.T) {
	result, err := GAIADB_CLIENT.GetMaintenTime(GAIADB_ID)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetSlowSqlDetail(t *testing.T) {
	args := &GetSlowSqlArgs{
		Page:     "1",
		PageSize: "10",
	}
	result, err := GAIADB_CLIENT.GetSlowSqlDetail(GAIADB_ID, args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_SlowSqlAdvice(t *testing.T) {
	result, err := GAIADB_CLIENT.SlowSqlAdvice(GAIADB_ID, "a41ff8fe-b0c6-407c-91f6-354717c3fbaa")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetBinlogDetail(t *testing.T) {
	args := &GetBinlogArgs{
		AppID:         GAIADB_ID,
		LogBackupType: "logical",
	}
	result, err := GAIADB_CLIENT.GetBinlogDetail("1694660508228961814", args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetBinlogList(t *testing.T) {
	args := &GetBinlogListArgs{
		AppID:         GAIADB_ID,
		LogBackupType: "logical",
		PageNo:        1,
		PageSize:      10,
	}
	result, err := GAIADB_CLIENT.GetBinlogList(args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ExecuteTaskNow(t *testing.T) {
	taskId := "3773297"
	err := GAIADB_CLIENT.ExecuteTaskNow(taskId)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CancelTask(t *testing.T) {
	taskId := "3773297"
	err := GAIADB_CLIENT.CancelTask(taskId)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetTaskList(t *testing.T) {
	args := &TaskListArgs{
		Region:    "bj",
		StartTime: "2023-09-11 16:00:00",
	}
	result, err := GAIADB_CLIENT.GetTaskList(args)
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetClusterByVpcId(t *testing.T) {
	result, err := GAIADB_CLIENT.GetClusterByVpcId("9556bf45-5867-4495-83c5-bd945b782503")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetClusterByLbId(t *testing.T) {
	result, err := GAIADB_CLIENT.GetClusterByLbId("496d6b552f316b313863786a4b32457a2f4732416e773d3d")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetOrderInfo(t *testing.T) {
	result, err := GAIADB_CLIENT.GetOrderInfo("8a3c9bb4313e489f859938b7d199487f")
	re, _ := json.Marshal(result)
	fmt.Println(string(re))
	ExpectEqual(t.Errorf, nil, err)
}

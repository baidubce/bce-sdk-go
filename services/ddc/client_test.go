package ddc

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	DDC_CLIENT            *Client
	DDC_ID                string = "rdsm2h07068gw39"
	ACCOUNT_NAME          string = "go_sdk_account_1"
	ACCOUNT_PASSWORD      string = "go_sdk_password_1"
	ACCOUNT_REMARK        string = "go-sdk-remark-1"
	DB_NAME               string = "go_sdk_db_1"
	DB_CHARACTER_SET_NAME string = "utf8"
	DB_REMARK             string = "go_sdk_db_remark"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

const (
	SDK_NAME_PREFIX = "sdk_rds_"
	POOL            = "xdb_gaiabase_pool"
	PNETIP          = "100.88.65.121"
	DEPLOY_ID       = "ab89d829-9068-d88e-75bc-64bb6367d036"
	INSTANCE_ID     = "rdsmcb38t7b8atx"
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

	DDC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_CreateInstance(t *testing.T) {
	//id := strconv.FormatInt(time.Now().Unix(),10)
	args := &CreateInstanceArgs{

		InstanceType: "RDS",
		Number:       1,
		Instance: Instance{
			Engine:               "mysql",
			EngineVersion:        "5.7",
			CpuCount:             1,
			AllocatedMemoryInGB:  8,
			AllocatedStorageInGB: 10,
			AZone:                "zoneA",
			SubnetId:             "zoneA:11c4f322-3a0e-4f26-8883-285cf64d0f03",
			DiskIoType:           "normal_io",
			DeployId:             "",
			PoolId:               "",
		},
	}
	DDC_CLIENT.CreateInstance(args)

}

func TestClient_ListDeploySets(t *testing.T) {
	deploys, err := DDC_CLIENT.ListDeploySets(POOL, &Marker{MaxKeys: 1})
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, deploys.MaxKeys, 1)
	fmt.Println(deploys.Result[0].DeployName)
}

func TestClient_GetDeploySet(t *testing.T) {
	deploy, err := DDC_CLIENT.GetDeploySet(POOL, DEPLOY_ID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, DEPLOY_ID, deploy.DeployID)
}

func TestClient_DeleteDeploySet(t *testing.T) {
	err := DDC_CLIENT.DeleteDeploySet(POOL, "b444142c-69ed-87a6-396b-fc4a76a9754f")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateDeploySet(t *testing.T) {
	err := DDC_CLIENT.CreateDeploySet(POOL, &CreateDeployRequest{
		DeployName: "api-from-go",
		Strategy:   "distribute",
	})
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListParameters(t *testing.T) {
	parameters, err := DDC_CLIENT.ListParameters("rdsmcb38t7b8atx")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, parameters.Items, parameters.Items)
	fmt.Println(parameters.Items)
}

func TestClient_UpdateParameter(t *testing.T) {
	DDC_CLIENT.UpdateParameter(INSTANCE_ID, &UpdateParameterArgs{Parameters: []KVParameter{
		{
			Name:  "auto_increment_increment",
			Value: "2",
		},
	}})
}

func TestClient_GetSecurityIps(t *testing.T) {
	ips, err := DDC_CLIENT.GetSecurityIps(INSTANCE_ID)
	ExpectEqual(t.Errorf, ips, ips)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(ips.SecurityIps)
}
func TestClient_ListRoGroup(t *testing.T) {
	ips, err := DDC_CLIENT.ListRoGroup("rdsm2bpbozmj59z")
	ExpectEqual(t.Errorf, ips, ips)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(ips)
}

func TestClient_ListVpc(t *testing.T) {
	vpc, err := DDC_CLIENT.ListVpc()
	ExpectEqual(t.Errorf, vpc, vpc)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(vpc)
}

func TestClient_GetDetail(t *testing.T) {
	detail, err := DDC_CLIENT.GetDetail("rdsm2bpbozmj59z")
	ExpectEqual(t.Errorf, detail, detail)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(detail)
}

func TestClient_UpdateSecurityIps(t *testing.T) {
	DDC_CLIENT.UpdateSecurityIps(INSTANCE_ID, &UpdateSecurityIpsArgs{
		SecurityIps: []string{
			"10.10.0.0/16",
		},
	})
}




func TestClient_GetBackupList(t *testing.T) {
	list, err := DDC_CLIENT.GetBackupList(INSTANCE_ID)
	ExpectEqual(t.Errorf, list.Snapshots, list.Snapshots)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(list.Snapshots)
}

func TestClient_GetBackupDetail(t *testing.T) {
	list, err := DDC_CLIENT.GetBackupDetail(INSTANCE_ID, "1611156600150859846")
	ExpectEqual(t.Errorf, list.Snapshot, list.Snapshot)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(list.Snapshot)
}

func TestClient_CreateBackup(t *testing.T) {
	DDC_CLIENT.CreateBackup(INSTANCE_ID)
}

func TestClient_ModifyBackupPolicy(t *testing.T) {
	DDC_CLIENT.ModifyBackupPolicy(INSTANCE_ID, &BackupPolicy{
		BackupDays:   "0,1,2,3,5",
		ExpireInDays: 100,
		BackupTime:   "17:00:00Z",
	})
}

func TestClient_GetZoneList(t *testing.T) {
	list, err := DDC_CLIENT.GetZoneList()
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, list, list)
	fmt.Println(list)
}

func TestClient_ListSubnets(t *testing.T) {
	subnets, err := DDC_CLIENT.ListSubnets()
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, subnets, subnets)
	fmt.Println(subnets)
}

func TestClient_GetBinlogList(t *testing.T) {
	list, err := DDC_CLIENT.GetBinlogList(INSTANCE_ID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, list, list)
	fmt.Println(list.Binlogs)
}

func TestClient_GetBinlogDetail(t *testing.T) {
	detail, err := DDC_CLIENT.GetBinlogDetail(INSTANCE_ID, "1611072328934047748")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, detail.Binlog.BinlogId, "1611072328934047748")
	fmt.Println(detail)
}

func TestClient_DeleteDdcInstance(t *testing.T) {
	instanceId := "rdsmvn9mmomzaw"
	DDC_CLIENT.DeleteDdcInstance(instanceId)
}

func TestClient_SwitchInstance(t *testing.T) {
	DDC_CLIENT.SwitchInstance("rdsma7fcyu2anvi")
}

// Database
func TestClient_CreateDatabase(t *testing.T) {
	args := &CreateDatabaseArgs{
		ClientToken:      getClientToken(),
		DbName:           DB_NAME,
		CharacterSetName: DB_CHARACTER_SET_NAME,
		Remark:           DB_REMARK,
	}

	err := DDC_CLIENT.CreateDatabase(DDC_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetDatabase(t *testing.T) {
	result, err := DDC_CLIENT.GetDatabase(DDC_ID, DB_NAME)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "available", result.DbStatus)
}

func TestClient_ListDatabase(t *testing.T) {
	result, err := DDC_CLIENT.ListDatabase(DDC_ID)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Databases {
		if e.DbName == DB_NAME {
			ExpectEqual(t.Errorf, "available", e.DbStatus)
		}
	}
}

func TestClient_UpdateDatabaseRemark(t *testing.T) {
	args := &UpdateDatabaseRemarkArgs{
		Remark: DB_REMARK + "_update",
	}
	err := DDC_CLIENT.UpdateDatabaseRemark(DDC_ID, DB_NAME, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteDatabase(t *testing.T) {
	err := DDC_CLIENT.DeleteDatabase(DDC_ID, DB_NAME)
	ExpectEqual(t.Errorf, nil, err)
}

// Account
func TestClient_CreateAccount(t *testing.T) {
	args := &CreateAccountArgs{
		ClientToken: getClientToken(),
		AccountName: ACCOUNT_NAME,
		Password:    ACCOUNT_PASSWORD,
		Type:        AccountType_Common,
		Remark:      ACCOUNT_REMARK,
	}

	err := DDC_CLIENT.CreateAccount(DDC_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetAccount(t *testing.T) {
	result, err := DDC_CLIENT.GetAccount(DDC_ID, ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "available", result.AccountStatus)
}

func TestClient_ListAccount(t *testing.T) {
	result, err := DDC_CLIENT.ListAccount(DDC_ID)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Accounts {
		if e.AccountName == ACCOUNT_NAME {
			ExpectEqual(t.Errorf, "available", e.AccountStatus)
		}
	}
}

func TestClient_UpdateAccountPassword(t *testing.T) {
	args := &UpdateAccountPasswordArgs{
		Password: ACCOUNT_PASSWORD + "_update",
	}
	err := DDC_CLIENT.UpdateAccountPassword(DDC_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAccountRemark(t *testing.T) {
	args := &UpdateAccountRemarkArgs{
		Remark: ACCOUNT_REMARK + "_update",
	}
	err := DDC_CLIENT.UpdateAccountRemark(DDC_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAccountPrivileges(t *testing.T) {
	databasePrivileges := []DatabasePrivilege{
		{
			DbName:   "hello",
			AuthType: AuthType_ReadWrite,
		},
	}
	args := &UpdateAccountPrivilegesArgs{
		DatabasePrivileges: databasePrivileges,
	}
	err := DDC_CLIENT.UpdateAccountPrivileges(DDC_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteAccount(t *testing.T) {
	err := DDC_CLIENT.DeleteAccount(DDC_ID, ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
}

// util
func getClientToken() string {
	return util.NewUUID()
}

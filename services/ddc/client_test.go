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
	DDC_ID                string = "ddc-m8rs4yjz"
	ACCOUNT_NAME          string = "go_sdk_account_2"
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
		Instance: CreateInstance{
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
		Strategy:   "distributed",
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
	detail, err := DDC_CLIENT.GetDetail("ddc-mmqptugx")
	ExpectEqual(t.Errorf, detail, detail)
	ExpectEqual(t.Errorf, err, nil)
	res, _ := json.Marshal(detail)
	fmt.Println(string(res))
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
	args := &ListSubnetsArgs{ZoneName: "zoneA"}
	subnets, err := DDC_CLIENT.ListSubnets(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, subnets, subnets)
	result, _ := json.Marshal(subnets.Subnets)
	fmt.Println(string(result))
}

func TestClient_GetBinlogList(t *testing.T) {
	list, err := DDC_CLIENT.GetBinlogList(INSTANCE_ID, "")
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
	instanceId := "ddc-m8rs4yjz"
	DDC_CLIENT.DeleteRds(instanceId)
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
		AccountType: "Common",
		Desc:        ACCOUNT_REMARK,
		DatabasePrivileges: []DatabasePrivilege{
			{
				DbName:   "hello",
				AuthType: "ReadOnly",
			},
		},
	}

	err := DDC_CLIENT.CreateAccount(DDC_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetAccount(t *testing.T) {
	result, err := DDC_CLIENT.GetAccount(DDC_ID, ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Available", result.Status)
}

func TestClient_ListAccount(t *testing.T) {
	result, err := DDC_CLIENT.ListAccount(DDC_ID)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Accounts {
		if e.AccountName == ACCOUNT_NAME {
			ExpectEqual(t.Errorf, "Available", e.Status)
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

func TestClient_UpdateAccountDesc(t *testing.T) {
	args := &UpdateAccountDescArgs{
		Desc: ACCOUNT_REMARK + "_update",
	}
	err := DDC_CLIENT.UpdateAccountDesc(DDC_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateAccountPrivileges(t *testing.T) {
	databasePrivileges := []DatabasePrivilege{
		{
			DbName:   "hello",
			AuthType: "ReadWrite",
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

func TestClient_CreateRds(t *testing.T) {
	args := &CreateRdsArgs{
		PurchaseCount: 1,
		InstanceName: "mysql_5.7",
		//SourceInstanceId: "ddc-mmqptugx",
		Engine: "mysql",
		EngineVersion: "5.7",
		CpuCount: 1,
		MemoryCapacity: 1,
		VolumeCapacity: 5,
		Billing: Billing{
			PaymentTiming: "Postpaid",
			Reservation: Reservation{ReservationLength: 1, ReservationTimeUnit: "Month"},
		},
		VpcId: "vpc-80m2ksi6sv0f",
		ZoneNames: []string{
			"cn-su-c",
		},
		Subnets: []SubnetMap{
			{
				ZoneName: "cn-su-c",
				SubnetId: "sbn-8v3p33vhyhq5",
			},

		},
	}
	rds, err := DDC_CLIENT.CreateRds(args)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(rds)
}

func TestClient_ListDdcInstance(t *testing.T) {
	args := &ListRdsArgs{
		// 批量获取列表的查询的起始位置，实例列表中Marker需要指定实例Id，可选
		// Marker: "marker",
		// 指定每页包含的最大数量(主实例)，最大数量不超过1000，缺省值为1000，可选
		MaxKeys: 100,
	}
	resp, err := DDC_CLIENT.ListRds(args)

	if err != nil {
		fmt.Printf("get instance error: %+v\n", err)
		return
	}

	// 返回标记查询的起始位置
	fmt.Println("ddc list marker: ", resp.Marker)
	// true表示后面还有数据，false表示已经是最后一页
	fmt.Println("ddc list isTruncated: ", resp.IsTruncated)
	// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	fmt.Println("ddc list nextMarker: ", resp.NextMarker)
	// 每页包含的最大数量
	fmt.Println("ddc list maxKeys: ", resp.MaxKeys)

	// 获取instance的列表信息
	for _, e := range resp.Result {
		fmt.Println("ddc instanceId: ", e.InstanceId)
		fmt.Println("ddc instanceName: ", e.InstanceName)
		fmt.Println("ddc engine: ", e.Engine)
		fmt.Println("ddc engineVersion: ", e.EngineVersion)
		fmt.Println("ddc instanceStatus: ", e.InstanceStatus)
		fmt.Println("ddc cpuCount: ", e.CpuCount)
		fmt.Println("ddc memoryCapacity: ", e.MemoryCapacity)
		fmt.Println("ddc volumeCapacity: ", e.VolumeCapacity)
		fmt.Println("ddc usedStorage: ", e.UsedStorage)
		fmt.Println("ddc paymentTiming: ", e.PaymentTiming)
		fmt.Println("ddc instanceType: ", e.InstanceType)
		fmt.Println("ddc instanceCreateTime: ", e.InstanceCreateTime)
		fmt.Println("ddc instanceExpireTime: ", e.InstanceExpireTime)
		fmt.Println("ddc publicAccessStatus: ", e.PublicAccessStatus)
		fmt.Println("ddc vpcId: ", e.VpcId)
	}

}
func TestClient_GetDetail2(t *testing.T) {
	result, err := DDC_CLIENT.GetDetail("ddc-m67du0mh")
	if err != nil {
		fmt.Printf("get instance error: %+v\n", err)
		return
	}
	// 获取实例详情信息
	fmt.Println("ddc instanceId: ", result.InstanceId)
	fmt.Println("ddc instanceName: ", result.InstanceName)
	fmt.Println("ddc engine: ", result.Engine)
	fmt.Println("ddc engineVersion: ", result.EngineVersion)
	fmt.Println("ddc instanceStatus: ", result.InstanceStatus)
	fmt.Println("ddc cpuCount: ", result.CpuCount)
	fmt.Println("ddc memoryCapacity: ", result.MemoryCapacity)
	fmt.Println("ddc volumeCapacity: ", result.VolumeCapacity)
	fmt.Println("ddc usedStorage: ", result.UsedStorage)
	fmt.Println("ddc paymentTiming: ", result.PaymentTiming)
	fmt.Println("ddc instanceType: ", result.InstanceType)
	fmt.Println("ddc instanceCreateTime: ", result.InstanceCreateTime)
	fmt.Println("ddc instanceExpireTime: ", result.InstanceExpireTime)
	fmt.Println("ddc publicAccessStatus: ", result.PublicAccessStatus)
	fmt.Println("ddc vpcId: ", result.VpcId)
	fmt.Println("ddc Subnets: ", result.Subnets)
	fmt.Println("ddc BackupPolicy: ", result.BackupPolicy)
	fmt.Println("ddc RoGroupList: ", result.RoGroupList)
	fmt.Println("ddc NodeMaster: ", result.NodeMaster)
	fmt.Println("ddc NodeSlave: ", result.NodeSlave)
	fmt.Println("ddc NodeReadReplica: ", result.NodeReadReplica)
	fmt.Println("ddc DeployId: ", result.DeployId)

}

func TestClient_UpdateInstanceName(t *testing.T) {
	args := &UpdateInstanceNameArgs{
		// DDC实例名称，允许小写字母、数字，中文，长度限制为1~64
		InstanceName: "ssss",
	}
	err := DDC_CLIENT.UpdateInstanceName("ddc-m67du0mh", args)
	if err != nil {
		fmt.Printf("update instance name error: %+v\n", err)
		return
	}
	fmt.Printf("update instance name success\n")
}
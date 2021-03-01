package ddcrds

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
	"time"
)

var (
	DDCRDS_CLIENT         *Client
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
	DDC_INSTANCE_ID = "ddc-me4mtqdi"
	RDS_INSTANCE_ID = "rds-OtTkC1OD"
	ETAG            = "v0"
)

var instanceId = "ddc-mnvw691i"
var client = DDCRDS_CLIENT

func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 2; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		panic("config json file of ak/sk not given:")
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	DDCRDS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	client = DDCRDS_CLIENT
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

func Json(v interface{}) string {
	jsonStr, err := json.Marshal(v)
	if err != nil {
		panic("convert to json faild")
	}
	return string(jsonStr)
}

func assertAvailable(instanceId string, t *testing.T) {
	result, err := DDCRDS_CLIENT.GetDetail(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Available", result.InstanceStatus)
}

func TestClient_CreateInstance(t *testing.T) {
	args := &CreateRdsArgs{
		PurchaseCount: 1,
		InstanceName:  "mysql_5.7",
		//SourceInstanceId: "ddc-mmqptugx",
		Engine:         "mysql",
		EngineVersion:  "5.7",
		CpuCount:       1,
		MemoryCapacity: 1,
		VolumeCapacity: 5,
		Billing: Billing{
			PaymentTiming: "Postpaid",
			Reservation:   Reservation{ReservationLength: 1, ReservationTimeUnit: "Month"},
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
		DeployId: "",
		PoolId:   "xdb_gaiabase_pool",
	}
	rds, err := DDCRDS_CLIENT.CreateRds(args, "ddc")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(rds)
}

func TestClient_ListDeploySets(t *testing.T) {
	result, err := client.ListDeploySets(POOL, nil)
	if err != nil {
		fmt.Printf("list deploy set error: %+v\n", err)
		return
	}

	for i := range result.Result {
		deploy := result.Result[i]
		fmt.Println("ddc deploy id: ", deploy.DeployID)
		fmt.Println("ddc deploy name: ", deploy.DeployName)
		fmt.Println("ddc deploy strategy: ", deploy.Strategy)
		fmt.Println("ddc deploy create time: ", deploy.CreateTime)
		fmt.Println("ddc deploy centralizeThreshold: ", deploy.CentralizeThreshold)
		fmt.Println("ddc instance ids: ", deploy.Instances)
	}
}

func TestClient_GetDeploySet(t *testing.T) {
	deploy, err := DDCRDS_CLIENT.GetDeploySet(POOL, DEPLOY_ID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, DEPLOY_ID, deploy.DeployID)
}

func TestClient_DeleteDeploySet(t *testing.T) {
	err := DDCRDS_CLIENT.DeleteDeploySet(POOL, "b444142c-69ed-87a6-396b-fc4a76a9754f")
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateDeploySet(t *testing.T) {
	result, err := DDCRDS_CLIENT.CreateDeploySet(POOL, &CreateDeployRequest{
		DeployName:          "api-from-go3",
		Strategy:            "distributed",
		CentralizeThreshold: 10,
	})
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(result.DeployID)
}

func TestClient_UpdateDeploySet(t *testing.T) {
	deployId := "1424c210-f998-1072-9608-9def4d93dec7"
	//err := DDCRDS_CLIENT.UpdateDeploySet(POOL, deployId, &UpdateDeployRequest{
	//	Strategy:   "centralized",
	//	CentralizeThreshold: 23,
	//})
	//ExpectEqual(t.Errorf, nil, err)

	client := DDCRDS_CLIENT
	args := &UpdateDeployRequest{
		// 幂等 Token
		ClientToken: "xxxyyyzzz",
		// 部署策略 支持集中部署(centralized)/完全打散(distributed)
		Strategy: "distributed",
		// 亲和度阈值 取值范围【0-96】，必须大于原亲和度阈值
		CentralizeThreshold: 30,
	}
	err := client.UpdateDeploySet(POOL, deployId, args)
	if err != nil {
		fmt.Printf("update deploy set error: %+v\n", err)
		return
	}

	fmt.Println("update deploy set success.")
}

func TestClient_ListParameters(t *testing.T) {
	parameters, err := DDCRDS_CLIENT.ListParameters(DDC_INSTANCE_ID)
	ExpectEqual(t.Errorf, nil, err)
	res, err := json.Marshal(parameters.Parameters)
	fmt.Println(len(parameters.Parameters))

	parameters, err = DDCRDS_CLIENT.ListParameters(RDS_INSTANCE_ID)
	ExpectEqual(t.Errorf, nil, err)
	res, err = json.Marshal(parameters.Parameters[:3])
	fmt.Println(string(res))
	fmt.Println(string(parameters.Etag))
}

func TestClient_UpdateParameter(t *testing.T) {
	//DDCRDS_CLIENT.UpdateParameter(DDC_INSTANCE_ID, "", &UpdateParameterArgs{Parameters: []KVParameter{
	//	{
	//		Name:  "auto_increment_increment",
	//		Value: "3",
	//	},
	//}}, "ddc")

	instances := getRdsList(t)
	for _, e := range *instances {
		if "Available" == e.InstanceStatus {
			res, err := DDCRDS_CLIENT.ListParameters(e.InstanceId)
			ExpectEqual(t.Errorf, nil, err)
			args := &UpdateParameterArgs{
				Parameters: []KVParameter{
					{
						Name:  "auto_increment_increment",
						Value: "2",
					},
				},
			}
			er := DDCRDS_CLIENT.UpdateParameter(e.InstanceId, res.Etag, args)
			ExpectEqual(t.Errorf, nil, er)
		}
	}
}

func getRdsList(t *testing.T) *[]Instance {
	listRdsArgs := &ListRdsArgs{}
	result, err := DDCRDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	return &result.Instances
}

func TestClient_GetSecurityIps(t *testing.T) {
	ips, err := DDCRDS_CLIENT.GetSecurityIps(DDC_INSTANCE_ID)
	ExpectEqual(t.Errorf, ips, ips)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(ips.SecurityIps)

	ips, err = DDCRDS_CLIENT.GetSecurityIps(RDS_INSTANCE_ID)
	ExpectEqual(t.Errorf, ips, ips)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(ips.SecurityIps)
}

// Only DDC
func TestClient_ListRoGroup(t *testing.T) {
	ips, err := DDCRDS_CLIENT.ListRoGroup(DDC_INSTANCE_ID)
	ExpectEqual(t.Errorf, ips, ips)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(ips)
}

// Only DDC
func TestClient_UpdateRoGroup(t *testing.T) {
	detail, err := client.GetDetail(instanceId)
	if err != nil {
		fmt.Printf("get ddc detail error: %+v\n", err)
		return
	}
	roGroupList := detail.RoGroupList
	if len(roGroupList) < 1 {
		fmt.Printf("the ddc instance %s do not have roGroup \n", instanceId)
	}
	roGroupId := roGroupList[0].RoGroupID
	fmt.Println(roGroupId)
	args := &UpdateRoGroupArgs{
		RoGroupName:         "testRo",
		IsBalanceRoLoad:     1,
		EnableDelayOff:      1,
		DelayThreshold:      0,
		LeastInstanceAmount: 0,
	}
	err = client.UpdateRoGroup(roGroupId, args, "ddc")
	if err != nil {
		fmt.Printf("update ddc roGroup error: %+v\n", err)
		return
	}
	fmt.Println("update ddc roGroup success")
}

// Only DDC
func TestClient_UpdateRoGroupReplicaWeight(t *testing.T) {
	detail, err := client.GetDetail(instanceId)
	if err != nil {
		fmt.Printf("get ddc detail error: %+v\n", err)
		return
	}
	roGroupList := detail.RoGroupList
	if len(roGroupList) < 1 {
		fmt.Printf("the ddc instance %s do not have roGroup \n", instanceId)
	}
	roGroupId := roGroupList[0].RoGroupID
	fmt.Println(roGroupId)
	replicaId := roGroupList[0].ReplicaList[0].InstanceId
	replicaWeight := ReplicaWeight{
		InstanceId: replicaId,
		Weight:     20,
	}
	args := &UpdateRoGroupWeightArgs{
		IsBalanceRoLoad: 0,
		ReplicaList:     []ReplicaWeight{replicaWeight},
	}
	err = client.UpdateRoGroupReplicaWeight(roGroupId, args, "ddc")
	if err != nil {
		fmt.Printf("update ddc roGroup replica weight error: %+v\n", err)
		return
	}
	fmt.Println("update ddc roGroup replica weight success")
}

// Only DDC
func TestClient_ReBalanceRoGroup(t *testing.T) {
	detail, err := client.GetDetail(instanceId)
	if err != nil {
		fmt.Printf("get ddc detail error: %+v\n", err)
		return
	}
	roGroupList := detail.RoGroupList
	if len(roGroupList) < 1 {
		fmt.Printf("the ddc instance %s do not have roGroup \n", instanceId)
	}
	roGroupId := roGroupList[0].RoGroupID
	fmt.Println(roGroupId)
	err = client.ReBalanceRoGroup(roGroupId, "ddc")
	if err != nil {
		fmt.Printf("reBalance ddc roGroup error: %+v\n", err)
		return
	}
	fmt.Println("reBalance ddc roGroup success")
}

// Only DDC
func TestClient_ListVpc(t *testing.T) {
	vpc, err := DDCRDS_CLIENT.ListVpc("ddc")
	ExpectEqual(t.Errorf, vpc, vpc)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(vpc)
}

func TestClient_GetDetail(t *testing.T) {
	result, err := DDCRDS_CLIENT.GetDetail(DDC_INSTANCE_ID)
	ExpectEqual(t.Errorf, err, nil)
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
	fmt.Println("ddc SyncMode: ", result.SyncMode)
	fmt.Println("ddc Category: ", result.Category)
	fmt.Println("ddc ZoneNames: ", result.ZoneNames)
	fmt.Println("ddc Endpoint: ", result.Endpoint)

	result, err = DDCRDS_CLIENT.GetDetail(RDS_INSTANCE_ID)
	ExpectEqual(t.Errorf, err, nil)
	res, _ := json.Marshal(result)
	fmt.Println(string(res))
}

func TestClient_UpdateSecurityIps(t *testing.T) {
	err := DDCRDS_CLIENT.UpdateSecurityIps(DDC_INSTANCE_ID, "", &UpdateSecurityIpsArgs{
		SecurityIps: []string{
			"10.10.0.0/16",
		},
	})
	ExpectEqual(t.Errorf, nil, err)

	instances := getRdsList(t)
	for _, e := range *instances {
		if "Available" == e.InstanceStatus {
			res, err := DDCRDS_CLIENT.GetSecurityIps(e.InstanceId)
			ExpectEqual(t.Errorf, nil, err)
			args := &UpdateSecurityIpsArgs{
				SecurityIps: []string{
					"%",
					"192.0.0.1",
					"192.0.0.2",
				},
			}
			er := DDCRDS_CLIENT.UpdateSecurityIps(e.InstanceId, res.Etag, args)
			ExpectEqual(t.Errorf, nil, er)
		}
	}
}

func TestClient_GetBackupList(t *testing.T) {
	args := &GetBackupListArgs{
		MaxKeys: 3,
	}
	list, err := DDCRDS_CLIENT.GetBackupList(DDC_INSTANCE_ID, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 3, list.MaxKeys)
	res, err := json.Marshal(list.Backups)
	fmt.Println(string(res))

	list, err = DDCRDS_CLIENT.GetBackupList(RDS_INSTANCE_ID, args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 3, list.MaxKeys)
	res, err = json.Marshal(list.Backups)
	fmt.Println(string(res))
}

// Only DDC
func TestClient_GetBackupDetail(t *testing.T) {
	list, err := DDCRDS_CLIENT.GetBackupDetail(DDC_INSTANCE_ID, "1612625400590890206")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(Json(list.Snapshot))

	list, err = DDCRDS_CLIENT.GetBackupDetail(RDS_INSTANCE_ID, "snap-rdsmi35jqm07uyg-2021_02_07T17_00_19Z")
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

// Only DDC
func TestClient_CreateBackup(t *testing.T) {
	err := DDCRDS_CLIENT.CreateBackup(DDC_INSTANCE_ID)
	ExpectEqual(t.Errorf, nil, err)
	err = DDCRDS_CLIENT.CreateBackup(RDS_INSTANCE_ID)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

// Only DDC
func TestClient_ModifyBackupPolicy(t *testing.T) {
	err := DDCRDS_CLIENT.ModifyBackupPolicy(DDC_INSTANCE_ID, &BackupPolicy{
		BackupDays:      "0,1,2,3,5",
		ExpireInDaysInt: 100,
		BackupTime:      "17:00:00Z",
	})
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.ModifyBackupPolicy(RDS_INSTANCE_ID, &BackupPolicy{
		BackupDays:      "0,1,2,3,5",
		ExpireInDaysInt: 100,
		BackupTime:      "17:00:00Z",
	})
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

// ddc 有多可用区
func TestClient_GetZoneList(t *testing.T) {
	list, err := DDCRDS_CLIENT.GetZoneList("ddc")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, list, list)
	fmt.Println(Json(list))

	list, err = DDCRDS_CLIENT.GetZoneList("rds")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(Json(list))
}

// rds使用apiZoneNames,ddc使用zoneNames逻辑可用区
func TestClient_ListSubnets(t *testing.T) {
	args := &ListSubnetsArgs{ZoneName: "zoneC"}
	subnets, err := DDCRDS_CLIENT.ListSubnets(args, "ddc")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, subnets, subnets)
	fmt.Println(Json(subnets))

	args = &ListSubnetsArgs{ZoneName: "cn-su-a"}
	subnets, err = DDCRDS_CLIENT.ListSubnets(args, "rds")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(Json(subnets))
}

// Only DDC
func TestClient_GetBinlogList(t *testing.T) {
	// 获取两天前的日志备份
	yesterday := time.Now().
		AddDate(0, 0, -2).
		Format("2006-01-02T15:04:05Z")
	list, err := DDCRDS_CLIENT.GetBinlogList(DDC_INSTANCE_ID, yesterday)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(Json(list.Binlogs))

	list, err = DDCRDS_CLIENT.GetBinlogList(RDS_INSTANCE_ID, "")
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

// Only DDC
func TestClient_GetBinlogDetail(t *testing.T) {
	detail, err := DDCRDS_CLIENT.GetBinlogDetail(DDC_INSTANCE_ID, "1612713949797105170")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, detail.Binlog.BinlogId, "1612713949797105170")
	fmt.Println(Json(detail))

	detail, err = DDCRDS_CLIENT.GetBinlogDetail(RDS_INSTANCE_ID, "1612355516300875570")
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

func TestClient_DeleteDdcInstance(t *testing.T) {
	instanceId := "rds-EOWuPqrI,rds-OtTkC1OD,ddc-m1e78f5f"
	err := DDCRDS_CLIENT.DeleteRds(instanceId)
	ExpectEqual(t.Errorf, nil, err)
}

// Only DDC
func TestClient_SwitchInstance(t *testing.T) {
	err := DDCRDS_CLIENT.SwitchInstance(DDC_INSTANCE_ID)
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.SwitchInstance(RDS_INSTANCE_ID)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

// Database
func TestClient_CreateDatabase(t *testing.T) {
	args := &CreateDatabaseArgs{
		ClientToken:      getClientToken(),
		DbName:           DB_NAME,
		CharacterSetName: DB_CHARACTER_SET_NAME,
		Remark:           DB_REMARK,
	}

	err := DDCRDS_CLIENT.CreateDatabase(DDC_INSTANCE_ID, args)
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.CreateDatabase(RDS_INSTANCE_ID, args)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

func TestClient_GetDatabase(t *testing.T) {
	result, err := DDCRDS_CLIENT.GetDatabase(DDC_INSTANCE_ID, DB_NAME)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Available", result.DbStatus)

	result, err = DDCRDS_CLIENT.GetDatabase(RDS_INSTANCE_ID, DB_NAME)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

func TestClient_ListDatabase(t *testing.T) {
	result, err := DDCRDS_CLIENT.ListDatabase(DDC_INSTANCE_ID)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Databases {
		if e.DbName == DB_NAME {
			ExpectEqual(t.Errorf, "available", e.DbStatus)
		}
		fmt.Println("ddc dbStatus: ", e.DbStatus)
		fmt.Println("ddc remark: ", e.Remark)
		fmt.Println("ddc accountPrivileges: ", e.AccountPrivileges)
	}

	result, err = DDCRDS_CLIENT.ListDatabase(RDS_INSTANCE_ID)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

func TestClient_UpdateDatabaseRemark(t *testing.T) {
	args := &UpdateDatabaseRemarkArgs{
		Remark: DB_REMARK + "_update",
	}
	err := DDCRDS_CLIENT.UpdateDatabaseRemark(DDC_INSTANCE_ID, DB_NAME, args)
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.UpdateDatabaseRemark(RDS_INSTANCE_ID, DB_NAME, args)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

func TestClient_DeleteDatabase(t *testing.T) {
	err := DDCRDS_CLIENT.DeleteDatabase(DDC_INSTANCE_ID, DB_NAME)
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.DeleteDatabase(RDS_INSTANCE_ID, DB_NAME)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
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
				DbName:   DB_NAME,
				AuthType: "ReadOnly",
			},
		},
	}

	err := DDCRDS_CLIENT.CreateAccount(DDC_INSTANCE_ID, args)
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.CreateAccount(RDS_INSTANCE_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetAccount(t *testing.T) {
	result, err := DDCRDS_CLIENT.GetAccount(DDC_INSTANCE_ID, ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Available", result.Status)

	result, err = DDCRDS_CLIENT.GetAccount(RDS_INSTANCE_ID, ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(Json(result))
	ExpectEqual(t.Errorf, "Available", result.Status)
}

func TestClient_ListAccount(t *testing.T) {
	result, err := DDCRDS_CLIENT.ListAccount(DDC_INSTANCE_ID)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Accounts {
		if e.AccountName == ACCOUNT_NAME {
			ExpectEqual(t.Errorf, "Available", e.Status)
		}
	}

	result, err = DDCRDS_CLIENT.ListAccount(RDS_INSTANCE_ID)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Accounts {
		if e.AccountName == ACCOUNT_NAME {
			ExpectEqual(t.Errorf, "Available", e.Status)
		}
	}
}

// Only DDC
func TestClient_UpdateAccountPassword(t *testing.T) {
	args := &UpdateAccountPasswordArgs{
		Password: ACCOUNT_PASSWORD + "_update",
	}
	err := DDCRDS_CLIENT.UpdateAccountPassword(DDC_INSTANCE_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.UpdateAccountPassword(RDS_INSTANCE_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

// Only DDC
func TestClient_UpdateAccountDesc(t *testing.T) {
	args := &UpdateAccountDescArgs{
		Desc: ACCOUNT_REMARK + "_update",
	}
	err := DDCRDS_CLIENT.UpdateAccountDesc(DDC_INSTANCE_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.UpdateAccountDesc(RDS_INSTANCE_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

// Only DDC
func TestClient_UpdateAccountPrivileges(t *testing.T) {
	databasePrivileges := []DatabasePrivilege{
		{
			DbName:   DB_NAME,
			AuthType: "ReadWrite",
		},
	}
	args := &UpdateAccountPrivilegesArgs{
		DatabasePrivileges: databasePrivileges,
	}
	err := DDCRDS_CLIENT.UpdateAccountPrivileges(DDC_INSTANCE_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.UpdateAccountPrivileges(RDS_INSTANCE_ID, ACCOUNT_NAME, args)
	ExpectEqual(t.Errorf, RDSNotSupportError(), err)
}

func TestClient_DeleteAccount(t *testing.T) {
	err := DDCRDS_CLIENT.DeleteAccount(DDC_INSTANCE_ID, ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)

	err = DDCRDS_CLIENT.DeleteAccount(RDS_INSTANCE_ID, ACCOUNT_NAME)
	ExpectEqual(t.Errorf, nil, err)
}

// util
func getClientToken() string {
	return util.NewUUID()
}

func TestClient_CreateRds(t *testing.T) {
	args := &CreateRdsArgs{
		PurchaseCount: 1,
		InstanceName:  "mysql57fromgo",
		//SourceInstanceId: "ddc-mmqptugx",
		Engine:         "mysql",
		EngineVersion:  "5.7",
		CpuCount:       1,
		MemoryCapacity: 1,
		VolumeCapacity: 5,
		Billing: Billing{
			PaymentTiming: "Postpaid",
			Reservation:   Reservation{ReservationLength: 1, ReservationTimeUnit: "Month"},
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
		//DeployId: "86be443c-a40d-be6a-58d5-e3aedc966cc1",
		PoolId: "xdb_5cf97afb-ee06-4b80-9146-4a840e5d0288_pool",
	}
	subnetArgs := &ListSubnetsArgs{ZoneName: "zoneC"}
	resp, err := DDCRDS_CLIENT.ListSubnets(subnetArgs, "ddc")
	ExpectEqual(t.Errorf, nil, err)
	if len(resp.Subnets) > 0 {
		subnet := resp.Subnets[0]
		fmt.Printf("ddc use subnet: %s\n", Json(subnet))
		// DDC使用ShortId
		args.Subnets[0].SubnetId = subnet.ShortId
	}
	ddc, err := DDCRDS_CLIENT.CreateRds(args, "ddc")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(Json(ddc))

	// 修改VPCId和子网Id
	subnetArgs = &ListSubnetsArgs{ZoneName: "cn-su-c"}
	resp, err = DDCRDS_CLIENT.ListSubnets(subnetArgs, "rds")
	ExpectEqual(t.Errorf, nil, err)
	if len(resp.Subnets) > 0 {
		subnet := resp.Subnets[0]
		fmt.Printf("rds use subnet: %s\n", Json(subnet))
		args.VpcId = subnet.VpcId
		args.Subnets[0].SubnetId = subnet.ShortId
	}
	rds, err := DDCRDS_CLIENT.CreateRds(args, "rds")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(Json(rds))
}

func TestClient_CreateReadReplica(t *testing.T) {
	client := DDCRDS_CLIENT
	instanceId := "ddc-mpsb5qre"
	args := &CreateReadReplicaArgs{
		//主实例ID，必选
		SourceInstanceId: instanceId,
		// 计费相关参数，只读实例只支持后付费Postpaid，必选
		Billing: Billing{
			PaymentTiming: "Postpaid",
		},
		// CPU核数，必选
		CpuCount: 2,
		//套餐内存大小，单位GB，必选
		MemoryCapacity: 8,
		//套餐磁盘大小，单位GB，每5G递增，必选
		VolumeCapacity: 20,
		//批量创建云数据库 ddc 只读实例个数, 目前只支持一次创建一个,可选
		//PurchaseCount: 2,
		//实例名称，允许小写字母、数字，长度限制为1~32，默认命名规则:{engine} + {engineVersion}，可选
		//指定zone信息，默认为空，由系统自动选择，可选
		//zoneName命名规范是小写的“国家-region-可用区序列"，例如北京可用区A为"cn-bj-a"。
		//ZoneNames: []string{"cn-su-c"},
		//与主实例 vpcId 相同，可选
		//VpcId: "vpc-IyrqYIQ7",
		//是否进行直接支付，默认false，设置为直接支付的变配订单会直接扣款，不需要再走支付逻辑，可选
		IsDirectPay: false,
		//vpc内，每个可用区的subnetId；如果不是默认vpc则必须指定 subnetId，可选
		//Subnets: []SubnetMap{
		//	{
		//		ZoneName: "cn-su-c",
		//		SubnetId: "sbn-8v3p33vhyhq5",
		//	},
		//},
		// 资源池id 必选与主实例保持一致
		//PoolId:"xdb_5cf97afb-ee06-4b80-9146-4a840e5d0288_pool",
		// RO组ID。(创建只读实例时) 可选
		// 如果不传，默认会创建一个RO组，并将该只读加入RO组中
		RoGroupId: "yyzzcc2",
		// RO组是否启用延迟剔除，默认不启动。（创建只读实例时）可选
		EnableDelayOff: 0,
		// 延迟阈值。（创建只读实例时）可选
		DelayThreshold: 1,
		// RO组最少保留实例数目。默认为1. （创建只读实例时）可选
		LeastInstanceAmount: 1,
		// 只读实例在RO组中的读流量权重。默认为1（创建只读实例时）可选
		RoGroupWeight: 1,
	}
	result, err := client.CreateReadReplica(args)
	if err != nil {
		fmt.Printf("create ddc readReplica error: %+v\n", err)
		return
	}

	for _, e := range result.InstanceIds {
		fmt.Println("create ddc readReplica success, instanceId: ", e)
	}
}

func TestClient_ListRds(t *testing.T) {
	args := &ListRdsArgs{
		// 批量获取列表的查询的起始位置，实例列表中Marker需要指定实例Id，可选
		// Marker: "marker",
		// 指定每页包含的最大数量(主实例)，最大数量不超过1000，缺省值为1000，可选
		MaxKeys: 20,
	}
	resp, err := DDCRDS_CLIENT.ListRds(args)

	if err != nil {
		fmt.Printf("get instance error: %+v\n", err)
		return
	}

	// 返回标记查询的起始位置
	fmt.Println("list marker: ", resp.Marker)
	// true表示后面还有数据，false表示已经是最后一页
	fmt.Println("list isTruncated: ", resp.IsTruncated)
	// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	fmt.Println("list nextMarker: ", resp.NextMarker)
	// 每页包含的最大数量
	fmt.Println("list maxKeys: ", resp.MaxKeys)

	// 获取instance的列表信息
	for _, e := range resp.Instances {
		if len(e.RoGroupList) > 0 {
			fmt.Println("=====================================>")
			fmt.Println("instance productType: ", e.ProductType())
			fmt.Println("instanceId: ", e.InstanceId)
			fmt.Println("instanceName: ", e.InstanceName)
			fmt.Println("engine: ", e.Engine)
			fmt.Println("engineVersion: ", e.EngineVersion)
			fmt.Println("instanceStatus: ", e.InstanceStatus)
			fmt.Println("cpuCount: ", e.CpuCount)
			fmt.Println("memoryCapacity: ", e.MemoryCapacity)
			fmt.Println("volumeCapacity: ", e.VolumeCapacity)
			fmt.Println("usedStorage: ", e.UsedStorage)
			fmt.Println("paymentTiming: ", e.PaymentTiming)
			fmt.Println("instanceType: ", e.InstanceType)
			fmt.Println("instanceCreateTime: ", e.InstanceCreateTime)
			fmt.Println("instanceExpireTime: ", e.InstanceExpireTime)
			fmt.Println("publiclyAccessible: ", e.PubliclyAccessible)
			fmt.Println("backup expireInDays: ", e.BackupPolicy.ExpireInDaysInt)
			fmt.Println("backup expireInDays: ", e.BackupPolicy.ExpireInDays)
			fmt.Println("vpcId: ", e.VpcId)
		}
	}
}

func TestClient_UpdateInstanceName(t *testing.T) {
	args := &UpdateInstanceNameArgs{
		// DDC实例名称，允许小写字母、数字，中文，长度限制为1~64
		InstanceName: "备份恢复测试-勿删-02",
	}
	err := DDCRDS_CLIENT.UpdateInstanceName(DDC_INSTANCE_ID, args)
	if err != nil {
		fmt.Printf("update instance name error: %+v\n", err)
		return
	}
	fmt.Printf("update instance name success\n")

	args = &UpdateInstanceNameArgs{
		// DDC实例名称，允许小写字母、数字，中文，长度限制为1~64
		InstanceName: "mysql56_rds",
	}
	err = DDCRDS_CLIENT.UpdateInstanceName(RDS_INSTANCE_ID, args)
	if err != nil {
		fmt.Printf("update instance name error: %+v\n", err)
		return
	}
	fmt.Printf("update instance name success\n")
}

// 以下操作仅支持RDS
// Only RDS
func TestClient_CreateRdsProxy(t *testing.T) {
	args := &CreateRdsProxyArgs{
		SourceInstanceId: RDS_INSTANCE_ID,
		NodeAmount:       2,
		Billing: Billing{
			PaymentTiming: "Postpaid",
		},
		ClientToken: getClientToken(),
	}
	assertAvailable(RDS_INSTANCE_ID, t)
	_, err := DDCRDS_CLIENT.CreateRdsProxy(args)
	ExpectEqual(t.Errorf, nil, err)
}

// Only RDS
func TestClient_ResizeRds(t *testing.T) {
	args := &ResizeRdsArgs{
		CpuCount:       1,
		MemoryCapacity: 2,
		VolumeCapacity: 10,
	}
	err := DDCRDS_CLIENT.ResizeRds(RDS_INSTANCE_ID, args)
	ExpectEqual(t.Errorf, nil, err)
	time.Sleep(30 * time.Second)
	assertAvailable(RDS_INSTANCE_ID, t)
}

func TestClient_RebootInstance(t *testing.T) {
	client := DDCRDS_CLIENT
	err := client.RebootInstance("rds-deaaDuV9")
	if err != nil {
		fmt.Printf("reboot error: %+v\n", err)
		return
	}

	// 延迟重启(仅支持DDC)
	args := &RebootArgs{
		IsRebootNow: true,
	}
	err = client.RebootInstanceWithArgs(instanceId, args)
	if err != nil {
		fmt.Printf("reboot ddc error: %+v\n", err)
		return
	}
}

// Only RDS
func TestClient_ModifySyncMode(t *testing.T) {
	assertAvailable(RDS_INSTANCE_ID, t)
	listRdsArgs := &ListRdsArgs{}
	result, err := DDCRDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if "Available" == e.InstanceStatus {
			args := &ModifySyncModeArgs{
				SyncMode: "Semi_sync", // Semi_sync
			}
			err := DDCRDS_CLIENT.ModifySyncMode(e.InstanceId, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

// Only RDS
func TestClient_ModifyEndpoint(t *testing.T) {
	assertAvailable(RDS_INSTANCE_ID, t)
	listRdsArgs := &ListRdsArgs{}
	result, err := DDCRDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if "Available" == e.InstanceStatus {
			args := &ModifyEndpointArgs{
				Address: "gosdk",
			}
			err := DDCRDS_CLIENT.ModifyEndpoint(e.InstanceId, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

// Only RDS
func TestClient_ModifyPublicAccess(t *testing.T) {
	assertAvailable(RDS_INSTANCE_ID, t)
	listRdsArgs := &ListRdsArgs{}
	result, err := DDCRDS_CLIENT.ListRds(listRdsArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if "Available" == e.InstanceStatus {
			args := &ModifyPublicAccessArgs{
				PublicAccess: true,
			}
			err := DDCRDS_CLIENT.ModifyPublicAccess(e.InstanceId, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_AutoRenew(t *testing.T) {
	args := &AutoRenewArgs{
		// 自动续费时长（续费单位为year 不大于3，续费单位为month 不大于9）必选
		AutoRenewTime: 1,
		// 自动续费单位（"year";"month"）必选
		AutoRenewTimeUnit: "month",
		// 实例id集合 必选
		InstanceIds: []string{
			"rds-OtTkC1OD",
			"rds-rbmh6gJl",
		},
	}
	err := client.AutoRenew(args, "rds")
	if err != nil {
		fmt.Printf("create auto renew error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetMaintainTime(t *testing.T) {
	maintainTime, err := client.GetMaintainTime(instanceId)
	if err != nil {
		fmt.Printf("get maintain time error: %+v\n", err)
		return
	}
	fmt.Println("maintainTime duration", maintainTime.Duration)
	fmt.Println("maintainTime period", maintainTime.Period)
	fmt.Println("maintainTime startTime", maintainTime.StartTime)
}

func TestClient_UpdateMaintainTime(t *testing.T) {
	client := DDCRDS_CLIENT
	args := &MaintainTime{
		// 时长间隔
		Duration: 2,
		// 1-7分别代表周一到周日
		Period: "1,2,3,4,5,0",
		// 所有涉及的时间皆为北京时间24小时制
		StartTime: "14:07",
	}
	err := client.UpdateMaintainTime(instanceId, args)
	if err != nil {
		fmt.Printf("update maintain time error: %+v\n", err)
	}
}

package ddcrds

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
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
	TASK_ID               string = "1173906"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

const (
	SDK_NAME_PREFIX = "sdk_rds_"
	POOL            = "xdb_005a2d79-a4f4-4bfb-8284-0ffe9ddaa307_pool"
	PNETIP          = "100.88.65.121"
	DEPLOY_ID       = "ab89d829-9068-d88e-75bc-64bb6367d036"
	DDC_INSTANCE_ID = "ddc-mqv3e72u"
	RDS_INSTANCE_ID = "rds-OtTkC1OD"
	ETAG            = "v0"
)

var instanceId = DDC_INSTANCE_ID
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

func assertAvailable(instanceId string, t *testing.T) {
	result, err := DDCRDS_CLIENT.GetDetail(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, "Available", result.InstanceStatus)
}

func TestClient_CreateInstance(t *testing.T) {
	args := &CreateRdsArgs{
		ClientToken: "81adc02cf0221a753d1ef969eb6c6360",
		Billing: Billing{
			PaymentTiming: "Prepaid",
			Reservation:   Reservation{ReservationLength: 1, ReservationTimeUnit: "Month"},
		},
		PurchaseCount:     1,
		InstanceName:      "go_sdk_tester",
		Engine:            "mysql",
		EngineVersion:     "5.7",
		Category:          "Standard",
		CpuCount:          2,
		MemoryCapacity:    4,
		VolumeCapacity:    50,
		IsDirectPay:       true,
		AutoRenewTime:     1,
		AutoRenewTimeUnit: "month",
		PoolId:            "xdb_9c72b2ea-a24c-41ba-b6c7-fc4eb7e8f538_pool",
		VpcId:             "vpc-4mcfvqcitav5",
		ZoneNames:         []string{"cn-bj-a"},
		Subnets: []SubnetMap{
			{
				ZoneName: "cn-bj-a",
				SubnetId: "sbn-izx7eq3wy87e",
			},
		},
	}
	result, err := DDCRDS_CLIENT.CreateRds(args, "ddc")
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println("create ddc success, orderId: ", result.OrderId)
	for _, e := range result.InstanceIds {
		fmt.Println("create ddc success, instanceId: ", e)
	}
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
		DeployName:          "api-from-go4",
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
				WaitSwitch: 0,
			}
			result, er := DDCRDS_CLIENT.UpdateParameter(e.InstanceId, res.Etag, args)
			ExpectEqual(t.Errorf, nil, er)
			if result != nil {
				fmt.Println("update parameter task success: ", result.Result.TaskID)
				TASK_ID = result.Result.TaskID
				TestClient_GetMaintainTaskDetail(t)
			} else {
				fmt.Println("update parameter task success.")
			}
			break
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
		IsBalanceRoLoad:     "1",
		EnableDelayOff:      "1",
		DelayThreshold:      "0",
		LeastInstanceAmount: "1",
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
		RoGroupName:         "testRo",
		EnableDelayOff:      "1",
		DelayThreshold:      "",
		LeastInstanceAmount: "0",
		IsBalanceRoLoad:     "0",
		ReplicaList:         []ReplicaWeight{replicaWeight},
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
func TestClient_ListPool(t *testing.T) {
	//pools, err := DDCRDS_CLIENT.ListPool(nil, "ddc")
	result, err := client.ListPool(nil, "ddc")
	if err != nil {
		fmt.Printf("list pool error: %+v\n", err)
		return
	}

	for i := range result.Result {
		pool := result.Result[i]
		fmt.Println("ddc pool id: ", pool.PoolID)
		fmt.Println("ddc pool vpc id: ", pool.VpcID)
		fmt.Println("ddc pool engine: ", pool.Engine)
		fmt.Println("ddc pool create time: ", pool.CreateTime)
		fmt.Println("ddc pool hosts: ", pool.Hosts)
	}
}

// Only DDC
func TestClient_ListVpc(t *testing.T) {
	vpc, err := DDCRDS_CLIENT.ListVpc("ddc")
	ExpectEqual(t.Errorf, vpc, vpc)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println(vpc)
}

func TestClient_GetDetail(t *testing.T) {
	instanceId = "ddc-mqqint6z"
	result, err := DDCRDS_CLIENT.GetDetail(instanceId)
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
	fmt.Println("master BBC hostname: ", result.NodeMaster.HostName)
	fmt.Println("ddc NodeSlave: ", result.NodeSlave)
	fmt.Println("slave BBC hostname: ", result.NodeSlave.HostName)
	fmt.Println("ddc NodeReadReplica: ", result.NodeReadReplica)
	fmt.Println("ddc DeployId: ", result.DeployId)
	fmt.Println("ddc SyncMode: ", result.SyncMode)
	fmt.Println("ddc Category: ", result.Category)
	fmt.Println("ddc ZoneNames: ", result.ZoneNames)
	fmt.Println("ddc Endpoint: ", result.Endpoint)
	fmt.Println("ddc Endpoint vnetIpBackup: ", result.Endpoint.VnetIpBackup)
	fmt.Println("ddc long BBC Id: ", result.LongBBCId)
	fmt.Println("ddc topo: ", result.InstanceTopoForReadonly)
	// 自动续费规则
	if result.AutoRenewRule != nil {
		fmt.Println("ddc renewTime: ", result.AutoRenewRule.RenewTime)
		fmt.Println("ddc renewTimeUnit: ", result.AutoRenewRule.RenewTimeUnit)
	}
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
	//args := &ListSubnetsArgs{VpcId: "vpc-fhsajv3w2j7h"}
	subnets, err := DDCRDS_CLIENT.ListSubnets(nil, "ddc")
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, subnets, subnets)
	fmt.Println(Json(subnets))

	//args = &ListSubnetsArgs{ZoneName: "cn-su-a"}
	//subnets, err = DDCRDS_CLIENT.ListSubnets(args, "rds")
	//ExpectEqual(t.Errorf, nil, err)
	//fmt.Println(Json(subnets))
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
	instanceId = "ddc-m8xc5hmz"
	args := &SwitchArgs{
		IsSwitchNow: false,
	}
	result, err := client.SwitchInstance(instanceId, args)
	if err != nil {
		fmt.Printf(" main standby switching of the instance error: %+v\n", err)
		return
	}
	if result != nil {
		fmt.Printf(" main standby switching of the instance success, taskId: %v\n", result.Result.TaskID)
	} else {
		fmt.Printf(" main standby switching of the instance success\n")
	}
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

func TestClient_GetTableAmount(t *testing.T) {
	args := &GetTableAmountArgs{
		InstanceId: instanceId,
		DbName:     "test1",
		Pattern:    "0",
	}
	result, err := DDCRDS_CLIENT.GetTableAmount(args)
	if err != nil {
		fmt.Printf("get table amount error: %+v\n", err)
		return
	}
	fmt.Printf("get table amount success.\n")
	fmt.Println("ddc return amount ", result.ReturnAmount)
	fmt.Println("ddc total amount ", result.TotalAmount)
	for k, v := range result.Tables {
		fmt.Println("ddc table ", k, " size: ", v)
	}
}

func TestClient_GetDatabaseDiskUsage(t *testing.T) {
	result, err := DDCRDS_CLIENT.GetDatabaseDiskUsage(instanceId, "")
	if err != nil {
		fmt.Printf("get database disk usage error: %+v\n", err)
		return
	}
	fmt.Printf("get database disk usage success.\n")
	fmt.Println("ddc rest disk size(byte) ", result.RestDisk)
	fmt.Println("ddc used disk size(byte) ", result.UsedDisk)
	for k, v := range result.Databases {
		fmt.Println("ddc table ", k, " size: ", v)
	}
}

func TestClient_GetRecoverableDateTime(t *testing.T) {
	result, err := DDCRDS_CLIENT.GetRecoverableDateTime(instanceId)
	if err != nil {
		fmt.Printf("get recoverable datetimes error: %+v\n", err)
		return
	}
	fmt.Printf("get recoverable datetimes success.\n")
	for _, e := range result.RecoverableDateTimes {
		fmt.Println("recover startTime: ", e.StartDateTime, " endTime: ", e.EndDateTime)
	}
}

func TestClient_RecoverToSourceInstanceByDatetime(t *testing.T) {
	dbName := "test2"
	tableName := "app_1"
	args := &RecoverInstanceArgs{
		Datetime: "2021-11-03T11:38:04Z",
		RecoverData: []RecoverData{
			{
				DbName:      dbName,
				NewDbName:   dbName + "_new",
				RestoreMode: "database",
			},
			{
				DbName:      dbName,
				NewDbName:   dbName + "_new",
				RestoreMode: "table",
				RecoverTables: []RecoverTable{
					{
						TableName:    tableName,
						NewTableName: tableName + "_new",
					},
				},
			},
		},
	}
	taskResult, err := DDCRDS_CLIENT.RecoverToSourceInstanceByDatetime(instanceId, args)
	if err != nil {
		fmt.Printf("recover instance database error: %+v\n", err)
		return
	}
	fmt.Printf("recover instance database success. taskId:%s\n", taskResult.TaskID)
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
		VolumeCapacity: 50,
		Billing: Billing{
			PaymentTiming: "Postpaid",
			Reservation:   Reservation{ReservationLength: 1, ReservationTimeUnit: "Month"},
		},
		VpcId: "vpc-fhsajv3w2j7h",
		ZoneNames: []string{
			"cn-bj-a",
		},
		Subnets: []SubnetMap{
			{
				ZoneName: "cn-bj-a",
				SubnetId: "sbn-7zdak3vr8jv2",
			},
		},
		//DeployId: "86be443c-a40d-be6a-58d5-e3aedc966cc1",
		PoolId:   "xdb_005a2d79-a4f4-4bfb-8284-0ffe9ddaa307_pool",
		Category: STANDARD,
		SyncMode: "Semi_sync",
	}
	subnetArgs := &ListSubnetsArgs{VpcId: args.VpcId}
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
	//subnetArgs = &ListSubnetsArgs{ZoneName: "cn-su-c"}
	//resp, err = DDCRDS_CLIENT.ListSubnets(subnetArgs, "rds")
	//ExpectEqual(t.Errorf, nil, err)
	//if len(resp.Subnets) > 0 {
	//	subnet := resp.Subnets[0]
	//	fmt.Printf("rds use subnet: %s\n", Json(subnet))
	//	args.VpcId = subnet.VpcId
	//	args.Subnets[0].SubnetId = subnet.ShortId
	//}
	//rds, err := DDCRDS_CLIENT.CreateRds(args, "rds")
	//ExpectEqual(t.Errorf, nil, err)
	//fmt.Println(Json(rds))
}

func TestClient_CreateReadReplica(t *testing.T) {
	client := DDCRDS_CLIENT
	instanceId := "ddc-m1b5gjr5"
	args := &CreateReadReplicaArgs{
		ClientToken: "320cfd8dceaf98529bd9f7c1d43a52c5",
		// 计费相关参数，DDC 只读实例只支持预付费，RDS 只读实例只支持后付费Postpaid，必选
		Billing: Billing{
			PaymentTiming: "Prepaid",
			Reservation:   Reservation{ReservationLength: 5, ReservationTimeUnit: "Month"},
		},
		PurchaseCount: 3,
		//主实例ID，必选
		SourceInstanceId: instanceId,
		InstanceName:     "go_tester_read",
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
		EnableDelayOff: "0",
		// 延迟阈值。（创建只读实例时）可选
		DelayThreshold: "1",
		// RO组最少保留实例数目。默认为1. （创建只读实例时）可选
		LeastInstanceAmount: "1",
		// 只读实例在RO组中的读流量权重。默认为1（创建只读实例时）可选
		RoGroupWeight: "1",
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
		Marker: "-1",
		// 指定每页包含的最大数量(主实例)，最大数量不超过1000，缺省值为1000，可选
		MaxKeys: 10,
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
		fmt.Println("vpcId: ", e.VpcId)
		fmt.Println("endpoint: ", e.Endpoint)
		fmt.Println("vnetIp: ", e.Endpoint.VnetIp)
		fmt.Println("vnetIpBackup: ", e.Endpoint.VnetIpBackup)
		fmt.Println("long BBC Id: ", e.LongBBCId)
		fmt.Println("bbc hostname: ", e.HostName)
		if e.AutoRenewRule != nil {
			fmt.Println("renewTime: ", e.AutoRenewRule.RenewTime)
			fmt.Println("renewTimeUnit: ", e.AutoRenewRule.RenewTimeUnit)
		}
	}
}

func TestClient_ListPage(t *testing.T) {
	args := &ListPageArgs{
		// 页码
		PageNo: 1,
		// 页大小
		PageSize: 10,
		// 筛选条件
		// 筛选字段类型,各筛选条件只能单独筛选，当取值为type、status、dbType、zone时可在集合中增加筛选项
		// all：匹配全部
		// instanceName：匹配实例名称
		// instanceId：匹配实例id；
		// vnetIpBackup：备库ip；
		// vnetIp：主库ip
		//Filters: []Filter{
		//	{KeywordType: "all", Keyword: "mysql"},
		//	{KeywordType: "zone", Keyword: "cn-bj-a"},
		//},
	}
	resp, err := DDCRDS_CLIENT.ListPage(args)

	if err != nil {
		fmt.Printf("get instance error: %+v\n", err)
		return
	}

	// 返回分页页码
	fmt.Println("list pageNo: ", resp.Page.PageNo)
	// 返回页大小
	fmt.Println("list pageSize: ", resp.Page.PageSize)
	// 返回总数量
	fmt.Println("list totalCount: ", resp.Page.TotalCount)

	// 获取instance的列表信息
	for _, e := range resp.Page.Result {
		fmt.Println("=====================================>")
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
		fmt.Println("vpcId: ", e.VpcId)
		fmt.Println("endpoint: ", e.Endpoint)
		fmt.Println("vnetIp: ", e.Endpoint.VnetIp)
		fmt.Println("vnetIpBackup: ", e.Endpoint.VnetIpBackup)
		fmt.Println("long BBC Id: ", e.LongBBCId)
		fmt.Println("bbc hostname: ", e.HostName)
		if e.AutoRenewRule != nil {
			fmt.Println("renewTime: ", e.AutoRenewRule.RenewTime)
			fmt.Println("renewTimeUnit: ", e.AutoRenewRule.RenewTimeUnit)
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
		IsResizeNow:    true,
		IsDirectPay:    true,
	}
	orderIdResponse, err := DDCRDS_CLIENT.ResizeRds(DDC_INSTANCE_ID, args)
	ExpectEqual(t.Errorf, nil, err)
	fmt.Println("resize ddc success, orderId: ", orderIdResponse.OrderId)
	time.Sleep(30 * time.Second)
	assertAvailable(DDC_INSTANCE_ID, t)
}

func TestClient_RebootInstance(t *testing.T) {
	// client := DDCRDS_CLIENT
	// err := client.RebootInstance("rds-deaaDuV9")
	// if err != nil {
	// 	fmt.Printf("reboot error: %+v\n", err)
	// 	return
	// }

	// 延迟重启(仅支持DDC)
	args := &RebootArgs{
		IsRebootNow: false,
	}
	result, err := client.RebootInstanceWithArgs(instanceId, args)
	if err != nil {
		fmt.Printf("reboot ddc error: %+v\n", err)
		return
	}
	if result != nil {
		fmt.Printf("reboot ddc success, taskId: %+v\n", result.TaskID)
	}
}

func TestClient_ModifySyncMode(t *testing.T) {
	assertAvailable(DDC_INSTANCE_ID, t)
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
func TestClient_ListRecycleInstances(t *testing.T) {
	marker := &Marker{MaxKeys: 10}
	instances, err := client.ListRecycleInstances(marker, "ddc")
	if err != nil {
		fmt.Printf("list recycler instances error: %+v\n", err)
		return
	}
	fmt.Println(Json(instances))
	for _, instance := range instances.Result {
		fmt.Println("+-------------------------------------------+")
		fmt.Println("instanceId: ", instance.InstanceId)
		fmt.Println("instanceName: ", instance.InstanceName)
		fmt.Println("engine: ", instance.Engine)
		fmt.Println("engineVersion: ", instance.EngineVersion)
		fmt.Println("instanceStatus: ", instance.InstanceStatus)
		fmt.Println("cpuCount: ", instance.CpuCount)
		fmt.Println("memoryCapacity: ", instance.MemoryCapacity)
		fmt.Println("volumeCapacity: ", instance.VolumeCapacity)
		fmt.Println("usedStorage: ", instance.UsedStorage)
		fmt.Println("instanceType: ", instance.InstanceType)
		fmt.Println("instanceCreateTime: ", instance.InstanceCreateTime)
		fmt.Println("instanceExpireTime: ", instance.InstanceExpireTime)
		fmt.Println("publicAccessStatus: ", instance.PublicAccessStatus)
		fmt.Println("vpcId: ", instance.VpcId)
	}
}

func TestClient_RecoverRecyclerInstances(t *testing.T) {
	instanceIds := []string{
		"ddc-mv8zcy6u",
		"ddc-mof1m3hb",
	}
	orderIdResponse, err := client.RecoverRecyclerInstances(instanceIds)
	if err != nil {
		fmt.Printf("recover recycler instances error: %+v\n", err)
		return
	}
	fmt.Println("recover recycler instances success, orderId: ", orderIdResponse.OrderId)
}

func TestClient_DeleteRecyclerInstances(t *testing.T) {
	instanceIds := []string{
		"ddc-mof1m3hb",
		"ddc-moxgq5dm",
	}
	err := client.DeleteRecyclerInstances(instanceIds)
	if err != nil {
		fmt.Printf("delete recycler instances error: %+v\n", err)
		return
	}
	fmt.Println("delete recycler instances success.")
}

func TestClient_ListSecurityGroupByVpcId(t *testing.T) {
	vpcId := "vpc-j1vaxw1cx2mw"
	securityGroups, err := client.ListSecurityGroupByVpcId(vpcId)
	if err != nil {
		fmt.Printf("list security group by vpcId error: %+v\n", err)
		return
	}
	for _, group := range *securityGroups {
		fmt.Println("+-------------------------------------------+")
		fmt.Println("id: ", group.SecurityGroupID)
		fmt.Println("name: ", group.Name)
		fmt.Println("description: ", group.Description)
		fmt.Println("associateNum: ", group.AssociateNum)
		fmt.Println("createdTime: ", group.CreatedTime)
		fmt.Println("version: ", group.Version)
		fmt.Println("defaultSecurityGroup: ", group.DefaultSecurityGroup)
		fmt.Println("vpc name: ", group.VpcName)
		fmt.Println("vpc id: ", group.VpcShortID)
		fmt.Println("tenantId: ", group.TenantID)
	}
	fmt.Println("list security group by vpcId success.")
}

func TestClient_ListSecurityGroupByInstanceId(t *testing.T) {
	instanceId := "ddc-m1h4mma5"
	result, err := client.ListSecurityGroupByInstanceId(instanceId)
	if err != nil {
		fmt.Printf("list security group by instanceId error: %+v\n", err)
		return
	}
	for _, group := range result.Groups {
		fmt.Println("+-------------------------------------------+")
		fmt.Println("securityGroupId: ", group.SecurityGroupID)
		fmt.Println("securityGroupName: ", group.SecurityGroupName)
		fmt.Println("securityGroupRemark: ", group.SecurityGroupRemark)
		fmt.Println("projectId: ", group.ProjectID)
		fmt.Println("vpcId: ", group.VpcID)
		fmt.Println("vpcName: ", group.VpcName)
		fmt.Println("inbound: ", group.Inbound)
		fmt.Println("outbound: ", group.Outbound)
	}
	fmt.Println("list security group by instanceId success.")
}

func TestClient_BindSecurityGroups(t *testing.T) {
	instanceIds := []string{
		"ddc-mf4c901b",
	}
	securityGroupIds := []string{
		"g-mi74p78rtq07",
	}
	args := &SecurityGroupArgs{
		InstanceIds:      instanceIds,
		SecurityGroupIds: securityGroupIds,
	}

	err := client.BindSecurityGroups(args)
	if err != nil {
		fmt.Printf("bind security groups to instances error: %+v\n", err)
		return
	}
	fmt.Println("bind security groups to instances success.")
}

func TestClient_UnBindSecurityGroups(t *testing.T) {
	instanceIds := []string{
		"ddc-mjafcdu0",
	}
	securityGroupIds := []string{
		"g-iutg5rtcydsk",
	}
	args := &SecurityGroupArgs{
		InstanceIds:      instanceIds,
		SecurityGroupIds: securityGroupIds,
	}

	err := client.UnBindSecurityGroups(args)
	if err != nil {
		fmt.Printf("unbind security groups to instances error: %+v\n", err)
		return
	}
	fmt.Println("unbind security groups to instances success.")
}

func TestClient_ReplaceSecurityGroups(t *testing.T) {
	instanceIds := []string{
		"ddc-mjafcdu0",
	}
	securityGroupIds := []string{
		"g-iutg5rtcydsk",
	}
	args := &SecurityGroupArgs{
		InstanceIds:      instanceIds,
		SecurityGroupIds: securityGroupIds,
	}

	err := client.ReplaceSecurityGroups(args)
	if err != nil {
		fmt.Printf("replace security groups to instances error: %+v\n", err)
		return
	}
	fmt.Println("replace security groups to instances success.")
}

func TestClient_ListLogByInstanceId(t *testing.T) {
	// 两天前
	date := time.Now().
		AddDate(0, 0, -2).
		Format("2006-01-02")
	fmt.Println(date)
	args := &ListLogArgs{
		LogType:  "error",
		Datetime: date,
	}
	logs, err := client.ListLogByInstanceId(instanceId, args)
	if err != nil {
		fmt.Printf("list logs of instance error: %+v\n", err)
		return
	}
	fmt.Println("list logs of instance success.")
	for _, log := range *logs {
		fmt.Println("+-------------------------------------------+")
		fmt.Println("id: ", log.LogID)
		fmt.Println("size: ", log.LogSizeInBytes)
		fmt.Println("start time: ", log.LogStartTime)
		fmt.Println("end time: ", log.LogEndTime)
	}
}

func TestClient_GetLogById(t *testing.T) {
	args := &GetLogArgs{
		ValidSeconds: 20,
	}
	logId := "errlog.202103091300"
	log, err := client.GetLogById(instanceId, logId, args)
	if err != nil {
		fmt.Printf("get log detail of instance error: %+v\n", err)
		return
	}
	fmt.Println("list logs of instances success.")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("id: ", log.LogID)
	fmt.Println("size: ", log.LogSizeInBytes)
	fmt.Println("start time: ", log.LogStartTime)
	fmt.Println("end time: ", log.LogEndTime)
	fmt.Println("download url: ", log.DownloadURL)
	fmt.Println("download url expires: ", log.DownloadExpires)
}

func TestClient_LazyDropCreateHardLink(t *testing.T) {
	dbName := "test2"
	tableName := "app_3"
	err := client.LazyDropCreateHardLink(instanceId, dbName, tableName)
	if err != nil {
		fmt.Printf("[lazy drop] create hard link error: %+v\n", err)
		return
	}
	fmt.Println("[lazy drop] create hard link success.")
}

func TestClient_LazyDropDeleteHardLink(t *testing.T) {
	dbName := "test2"
	tableName := "app_1"
	result, err := client.LazyDropDeleteHardLink(instanceId, dbName, tableName)
	if err != nil {
		fmt.Printf("[lazy drop] delete hard link error: %+v\n", err)
		return
	}
	fmt.Println("[lazy drop] delete hard link success.taskId:", result.TaskID)
}

func TestClient_GetDisk(t *testing.T) {
	disk, err := client.GetDisk(instanceId)
	if err != nil {
		fmt.Printf("get disk of instance error: %+v\n", err)
		return
	}
	fmt.Println("get disk of instance success.")
	for _, diskItem := range disk.Response.Items {
		fmt.Println("instance id: ", diskItem.InstanceID)
		fmt.Println("instance role: ", diskItem.InstanceRole)
		fmt.Println("disk disk partition: ", diskItem.DiskPartition)
		fmt.Println("disk totle size in bytes: ", diskItem.TotalSize)
		fmt.Println("disk used size in bytes: ", diskItem.UsedSize)
		fmt.Println("disk report time: ", diskItem.ReportTime)
	}
}

func TestClient_GetMachineInfo(t *testing.T) {
	machine, err := client.GetMachineInfo(instanceId)
	if err != nil {
		fmt.Printf("get machine info error: %+v\n", err)
		return
	}
	fmt.Println("get machine info success.")
	for _, machine := range machine.Response.Items {
		fmt.Println("instance id: ", machine.InstanceID)
		fmt.Println("instance role: ", machine.Role)
		fmt.Println("cpu(core): ", machine.CPUInCore)
		fmt.Println("cpu(core) free: ", machine.FreeCPUInCore)
		fmt.Println("memory(MB): ", machine.MemSizeInMB)
		fmt.Println("memory(MB) free: ", machine.FreeMemSizeInMB)
		fmt.Println("disk info: ", machine.SizeInGB)
		fmt.Println("----------------------")
	}
}

func TestClient_GetResidual(t *testing.T) {
	residual, err := client.GetResidual(POOL)
	if err != nil {
		fmt.Printf("get residual of pool error: %+v\n", err)
		return
	}
	fmt.Println("get residual of pool success.")
	for zoneName, residualByZone := range residual.Residual {
		fmt.Println("zone name: ", zoneName)
		fmt.Printf("Single residual: disk %v GB, memory %v GB, cpu cores %d\n",
			residualByZone.Single.DiskInGb, residualByZone.Single.MemoryInGb, residualByZone.Single.CPUInCore)
		fmt.Printf("Slave residual: disk %v GB, memory %v GB, cpu cores %d\n",
			residualByZone.Slave.DiskInGb, residualByZone.Slave.MemoryInGb, residualByZone.Slave.CPUInCore)
		fmt.Printf("HA residual: disk %v GB, memory %v GB, cpu cores %d\n",
			residualByZone.HA.DiskInGb, residualByZone.HA.MemoryInGb, residualByZone.HA.CPUInCore)
	}
}

func TestClient_GetFlavorCapacity(t *testing.T) {
	args := &GetFlavorCapacityArgs{
		CpuInCore:  2,
		MemoryInGb: 4,
		DiskInGb:   50,
		Affinity:   2,
	}

	args = NewDefaultGetFlavorCapacityArgs(2, 4, 50)
	capacityResult, err := client.GetFlavorCapacity(POOL, args)
	if err != nil {
		fmt.Printf("get flavor capacity of pool error: %+v\n", err)
		return
	}
	fmt.Println("get flavor capacity of pool success.")
	for zoneName, residualByZone := range capacityResult.Capacity {
		fmt.Println("zone name: ", zoneName)
		fmt.Printf("HA capacity: %d\n", residualByZone.HA)
		fmt.Printf("Single capacity: %d\n", residualByZone.Single)
		fmt.Printf("Slave capacity: %d\n", residualByZone.Slave)
	}
}

func TestClient_KillSession(t *testing.T) {
	args := &KillSessionArgs{
		Role:       "master",
		SessionIds: []int{8661, 8662},
	}
	result, err := client.KillSession(instanceId, args)
	if err != nil {
		fmt.Printf("start kill session task error: %+v\n", err)
		return
	}
	fmt.Println("start kill session task success. TaskID:", result.TaskID)
}

func TestClient_GetKillSessionTaskResult(t *testing.T) {
	taskId := 285647
	result, err := client.GetKillSessionTask(instanceId, taskId)
	if err != nil {
		fmt.Printf("get kill session task error: %+v\n", err)
		return
	}
	fmt.Println("get kill session task success.")
	for _, task := range result.Tasks {
		fmt.Println("sessionId: ", task.SessionID)
		fmt.Println("task status: ", task.Status)
	}
}

func TestClient_GetMaintainTaskList(t *testing.T) {
	args := &GetMaintainTaskListArgs{
		Marker: Marker{
			MaxKeys: 10,
		},
		// 任务起始时间 必选
		StartTime: "2021-08-10 00:00:00",
	}
	result, err := client.GetMaintainTaskList(args)
	if err != nil {
		fmt.Printf("get tasks error: %+v\n", err)
		return
	}
	fmt.Println("get tasks success.")
	// 返回标记查询的起始位置
	fmt.Println("list marker: ", result.Marker)
	// true表示后面还有数据，false表示已经是最后一页
	fmt.Println("list isTruncated: ", result.IsTruncated)
	// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	fmt.Println("list nextMarker: ", result.NextMarker)
	// 每页包含的最大数量
	fmt.Println("list maxKeys: ", result.MaxKeys)
	for _, task := range result.Result {
		fmt.Println("task id: ", task.TaskID)
		fmt.Println("task name: ", task.TaskName)
		fmt.Println("task status: ", task.TaskStatus)
		fmt.Println("instance id: ", task.InstanceID)
		fmt.Println("instance name: ", task.InstanceName)
		fmt.Println("instance region: ", task.Region)
		fmt.Println("start time: ", task.StartTime)
		fmt.Println("end time: ", task.EndTime)
		fmt.Println("--------------------------")
	}
}

func TestClient_GetMaintainTaskDetail(t *testing.T) {
	result, err := client.GetMaintainTaskDetail(TASK_ID)
	if err != nil {
		fmt.Printf("get task detail error: %+v\n", err)
		return
	}
	fmt.Println("get task detail success.")
	for _, task := range result.Tasks {
		fmt.Println("task id: ", task.TaskID)
		fmt.Println("task name: ", task.TaskName)
		fmt.Println("instance instanceId: ", task.InstanceID)
		fmt.Println("instance instanceName: ", task.InstanceName)
		fmt.Println("instance instanceName: ", task.AppID)
		fmt.Println("instance task type: ", task.TaskType)
		fmt.Println("task status: ", task.TaskStatus)
		fmt.Println("instance create time: ", task.CreateTime)
		fmt.Println("instance create time: ", task.TaskSpecialAction)
		fmt.Println("instance create time: ", task.TaskSpecialActionTime)
		fmt.Println("--------------------------")
	}
}

func TestClient_ExecuteMaintainTaskImmediately(t *testing.T) {
	taskId := "880337"
	err := client.ExecuteMaintainTaskImmediately(taskId)
	if err != nil {
		fmt.Printf("execute task invoke error: %+v\n", err)
		return
	}
	fmt.Println("execute task invoke  success.")
}

func TestClient_CancelMaintainTask(t *testing.T) {
	taskId := "880337"
	err := client.CancelMaintainTask(taskId)
	if err != nil {
		fmt.Printf("cancel task invoke error: %+v\n", err)
		return
	}
	fmt.Println("cancel task invoke  success.")
}

func TestClient_GetAccessLog(t *testing.T) {
	date := "20210810"
	downloadInfo, err := client.GetAccessLog(date)
	if err != nil {
		fmt.Printf("get access logs error: %+v\n", err)
		return
	}
	fmt.Println("get access logs success.")
	fmt.Println("mysql access logs link: ", downloadInfo.Downloadurl.Mysql)
	fmt.Println("bbc access logs link: ", downloadInfo.Downloadurl.Bbc)
	fmt.Println("bos access logs link: ", downloadInfo.Downloadurl.Bos)
}

func TestClient_GetErrorLogs(t *testing.T) {
	args := &GetErrorLogsArgs{
		InstanceId: "ddc-mp8lme9w",
		StartTime:  "2021-08-16T02:28:51Z",
		EndTime:    "2021-08-17T02:28:51Z",
		PageNo:     1,
		PageSize:   10,
		Role:       "master",
		KeyWord:    "Aborted",
	}
	errorLogsResponse, err := client.GetErrorLogs(args)
	if err != nil {
		fmt.Printf("get error logs error: %+v\n", err)
		return
	}
	fmt.Println("get error logs success.")
	fmt.Println("error logs count: ", errorLogsResponse.Count)
	for _, errorLog := range errorLogsResponse.ErrorLogs {
		fmt.Println("=================================================")
		fmt.Println("error log instanceId: ", errorLog.InstanceId)
		fmt.Println("error log executeTime: ", errorLog.ExecuteTime)
		fmt.Println("error log logLevel: ", errorLog.LogLevel)
		fmt.Println("error log logText: ", errorLog.LogText)
	}
}

func TestClient_GetSlowLogs(t *testing.T) {
	args := &GetSlowLogsArgs{
		InstanceId: "ddc-mp8lme9w",
		StartTime:  "2021-08-16T02:28:51Z",
		EndTime:    "2021-08-17T02:28:51Z",
		PageNo:     1,
		PageSize:   10,
		Role:       "master",
		DbName:     []string{"baidu_dba"},
		UserName:   []string{"_root"},
		HostIp:     []string{"localhost"},
		Sql:        "update heartbeat set id=?, value=?",
	}
	slowLogsResponse, err := client.GetSlowLogs(args)
	if err != nil {
		fmt.Printf("get slow logs error: %+v\n", err)
		return
	}
	fmt.Println("get slow logs success.")
	fmt.Println("slow logs count: ", slowLogsResponse.Count)
	for _, slowLog := range slowLogsResponse.SlowLogs {
		fmt.Println("=================================================")
		fmt.Println("slow log instanceId: ", slowLog.InstanceId)
		fmt.Println("slow log userName: ", slowLog.UserName)
		fmt.Println("slow log dbName: ", slowLog.DbName)
		fmt.Println("slow log hostIp: ", slowLog.HostIp)
		fmt.Println("slow log queryTime: ", slowLog.QueryTime)
		fmt.Println("slow log lockTime: ", slowLog.LockTime)
		fmt.Println("slow log rowsExamined: ", slowLog.RowsExamined)
		fmt.Println("slow log rowsSent: ", slowLog.RowsSent)
		fmt.Println("slow log sql: ", slowLog.Sql)
		fmt.Println("slow log executeTime: ", slowLog.ExecuteTime)
	}
}

func TestClient_GetInstanceBackupStatus(t *testing.T) {
	instanceId = "ddc-mvxhc1fq"
	backupStatusResult, err := client.GetInstanceBackupStatus(instanceId)
	if err != nil {
		fmt.Printf("get backup status error: %+v\n", err)
		return
	}
	fmt.Println("get backup status success.")
	fmt.Println("instance is backuping: ", backupStatusResult.IsBackuping)
	if backupStatusResult.IsBackuping {
		fmt.Println("instance backup start time: ", backupStatusResult.SnapshotStartTime)
	}
}

func TestClient_InstanceVersionRollBack(t *testing.T) {
	instanceId = "ddc-mvxhc1fq"
	args := &InstanceVersionRollBackArg{
		// 是否维护时间执行
		WaitSwitch: true,
	}
	result, err := client.InstanceVersionRollBack(instanceId, args)
	if err != nil {
		fmt.Printf("rollback instance version faild: %+v\n", err)
		return
	}
	fmt.Printf("rollback instance version success. taskId:%s\n", result.TaskID)
}

func TestClient_InstanceVersionUpgrade(t *testing.T) {
	instanceId = "ddc-mvxhc1fq"
	args := &InstanceVersionUpgradeArg{
		// 是否立即执行
		IsUpgradeNow: true,
	}
	result, err := client.InstanceVersionUpgrade(instanceId, args)
	if err != nil {
		fmt.Printf("upgrade instance version faild: %+v\n", err)
		return
	}
	fmt.Printf("upgrade instance version success. taskId:%s\n", result.TaskID)
}

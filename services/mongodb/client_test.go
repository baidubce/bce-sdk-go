package mongodb

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	CLIENT     *Client
	INSTANCEID = "m-oHGYu8"
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

func TestListMongodb(t *testing.T) {
	dbInstanceType := "replica"
	args := &ListMongodbArgs{
		DbInstanceType: dbInstanceType,
		MaxKeys:        10,
	}
	result, err := CLIENT.ListMongodb(args)
	ExpectEqual(t.Errorf, nil, err)
	for _, instance := range result.DbInstances {
		ExpectEqual(t.Errorf, dbInstanceType, instance.DbInstanceType)
	}
}

func TestGetInstanceDetail(t *testing.T) {
	detail, err := CLIENT.GetInstanceDetail(INSTANCEID)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, INSTANCEID, detail.DbInstanceId)
}

func TestCreateReplica(t *testing.T) {
	vpcId := "vpc-e5kk1bvt1t11"
	subnetId := "sbn-urdqd0y5sf11"
	zoneName := "cn-gz-f"
	resGroupId := "RESG-XiqksQtNE11"
	args := &CreateReplicaArgs{
		StorageEngine:            "WiredTiger",
		DbInstanceType:           "replica",
		DbInstanceCpuCount:       1,
		DbInstanceMemoryCapacity: 2,
		ReadonlyNodeNum:          0,
		DbInstanceStorage:        5,
		VotingMemberNum:          1,
		EngineVersion:            "3.6",
		DbInstanceName:           "test_name",
		ResGroupId:               resGroupId,
		VpcId:                    vpcId,
		Subnets: []SubnetMap{
			{
				ZoneName: zoneName,
				SubnetId: subnetId,
			},
		},
		Billing: BillingModel{
			Reservation: Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "Month",
			},
			PaymentTiming: "Postpaid",
		},
		Tags: []TagModel{
			{
				TagKey:   "123",
				TagValue: "13",
			},
			{
				TagKey:   "test",
				TagValue: "test",
			},
		},
	}
	result, err := CLIENT.CreateReplica(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(result.DbInstanceSimpleModels) > 0)
}

func TestCreateReplicaPrepaid(t *testing.T) {
	vpcId := "vpc-e5kk1bvt1t11"
	subnetId := "sbn-urdqd0y5sf11"
	zoneName := "cn-gz-f"
	resGroupId := "RESG-XiqksQtNE11"
	args := &CreateReplicaArgs{
		StorageEngine:            "WiredTiger",
		DbInstanceType:           "replica",
		DbInstanceCpuCount:       1,
		DbInstanceMemoryCapacity: 2,
		ReadonlyNodeNum:          0,
		DbInstanceStorage:        5,
		VotingMemberNum:          1,
		EngineVersion:            "3.6",
		DbInstanceName:           "test_name",
		ResGroupId:               resGroupId,
		VpcId:                    vpcId,
		Subnets: []SubnetMap{
			{
				ZoneName: zoneName,
				SubnetId: subnetId,
			},
		},
		Billing: BillingModel{
			Reservation: Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "Month",
			},
			AutoRenew: AutoRenewModel{
				AutoRenewLength:   2,
				AutoRenewTimeUnit: "Month",
			},
			PaymentTiming: "Prepaid",
		},
		Tags: []TagModel{
			{
				TagKey:   "123",
				TagValue: "13",
			},
			{
				TagKey:   "test",
				TagValue: "test",
			},
		},
	}
	result, err := CLIENT.CreateReplica(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(result.DbInstanceSimpleModels) > 0)
}

func TestReplicaResize(t *testing.T) {
	instanceId := "m-cvbCWv"
	args := ReplicaResizeArgs{
		DbInstanceStorage:        7,
		DbInstanceCpuCount:       2,
		DbInstanceMemoryCapacity: 4,
	}
	err := CLIENT.ReplicaResize(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestRestartMongodb(t *testing.T) {
	instanceId := "m-cvbCWv"
	err := CLIENT.RestartMongodb(instanceId)
	ExpectEqual(t.Errorf, nil, err)
}

func TestRestartMongodbs(t *testing.T) {
	instanceIds := []string{"m-ARidmu", "m-oHGYu8"}
	err := CLIENT.RestartMongodbs(instanceIds)
	ExpectEqual(t.Errorf, nil, err)
}

func TestUpdateInstanceName(t *testing.T) {
	args := UpdateInstanceNameArgs{
		DbInstanceName: "test",
	}
	instanceId := "m-oHGYu8"
	err := CLIENT.UpdateInstanceName(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestUpdateAccountPassword(t *testing.T) {
	args := UpdatePasswordArgs{
		AccountPassword: "LegalPassword", // LegalPassword
	}
	instanceId := "m-oHGYu8"
	err := CLIENT.UpdateAccountPassword(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)

	args = UpdatePasswordArgs{
		AccountPassword: "IllegalPassword",
	}
	err = CLIENT.UpdateAccountPassword(instanceId, &args)
	ExpectEqual(t.Errorf, true, err != nil)
}

func TestReplicaSwitch(t *testing.T) {
	instanceId := "m-ARidmu"
	err := CLIENT.ReplicaSwitch(instanceId)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetReadonlyNodes(t *testing.T) {
	instanceId := "m-ARidmu"
	result, err := CLIENT.GetReadonlyNodes(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(result.ReadOnlyList))
	ExpectEqual(t.Errorf, 0, len(result.ReadOnlyList[0].NodeIds))
}

func TestReplicaAddReadonlyNodes(t *testing.T) {
	increment := 1
	instanceId := "m-ARidmu"
	subnetId := "sbn-urdqd0y5sf11"
	zoneName := "cn-gz-f"

	result, err := CLIENT.GetReadonlyNodes(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(result.ReadOnlyList))
	existing_cnt := len(result.ReadOnlyList[0].NodeIds)

	args := ReplicaAddReadonlyNodesArgs{
		ReadonlyNodeNum: increment,
		Subnet: SubnetMap{
			ZoneName: zoneName,
			SubnetId: subnetId,
		},
	}
	add_result, err := CLIENT.ReplicaAddReadonlyNodes(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, increment, len(add_result.ReadonlyMemberIds))

	time.Sleep(10 * time.Second)
	result, err = CLIENT.GetReadonlyNodes(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(result.ReadOnlyList))
	ExpectEqual(t.Errorf, existing_cnt+increment, len(result.ReadOnlyList[0].NodeIds))
}

func TestCreateSharding(t *testing.T) {
	vpcId := "vpc-e5kk1bvt1t11"
	subnetId := "sbn-urdqd0y5sf11"
	zoneName := "cn-gz-f"
	resGroupId := "RESG-XiqksQtNE11"
	args := &CreateShardingArgs{
		StorageEngine:        "WiredTiger",
		DbInstanceType:       "sharding",
		MongosCpuCount:       1,
		MongosMemoryCapacity: 2,
		ShardCpuCount:        1,
		ShardMemoryCapacity:  2,
		ShardStorage:         5,
		EngineVersion:        "3.6",
		DbInstanceName:       "shard_test_name",
		ResGroupId:           resGroupId,
		VpcId:                vpcId,
		Subnets: []SubnetMap{
			{
				ZoneName: zoneName,
				SubnetId: subnetId,
			},
		},
		Billing: BillingModel{
			Reservation: Reservation{
				ReservationLength:   0,
				ReservationTimeUnit: "Month",
			},
			PaymentTiming: "Postpaid",
		},
		Tags: []TagModel{
			{
				TagKey:   "123",
				TagValue: "13",
			},
			{
				TagKey:   "test",
				TagValue: "test",
			},
		},
	}
	result, err := CLIENT.CreateSharding(args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(result.DbInstanceSimpleModels) > 0)
}

func TestShardingComponentResize(t *testing.T) {
	instanceId := "m-2ke5iF"
	nodeId := "shd-dc2N9I"
	args := ShardingComponentResizeArgs{
		NodeCpuCount:       1,
		NodeMemoryCapacity: 2,
		NodeStorage:        6,
	}
	err := CLIENT.ShardingComponentResize(instanceId, nodeId, &args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestRestartShardingComponent(t *testing.T) {
	instanceId := "m-2ke5iF"
	nodeId := "shd-dc2N9I"
	err := CLIENT.RestartShardingComponent(instanceId, nodeId)
	ExpectEqual(t.Errorf, nil, err)
}

func TestUpdateShardingComponentName(t *testing.T) {
	instanceId := "m-2ke5iF"
	nodeId := "shd-dc2N9I"
	args := UpdateComponentNameArgs{
		NodeName: "test_name",
	}
	err := CLIENT.UpdateShardingComponentName(instanceId, nodeId, &args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestShardingComponentSwitch(t *testing.T) {
	instanceId := "m-2ke5iF"
	nodeId := "shd-dc2N9I"
	err := CLIENT.ShardingComponentSwitch(instanceId, nodeId)
	ExpectEqual(t.Errorf, nil, err)
}

func TestShardingAddComponent(t *testing.T) {
	instanceId := "m-2ke5iF"
	args := ShardingAddComponentArgs{
		NodeCpuCount:       1,
		NodeMemoryCapacity: 2,
		NodeStorage:        10,
		NodeType:           "mongos",
	}
	result, err := CLIENT.ShardingAddComponent(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 1, len(result.NodeIds))
	ExpectEqual(t.Errorf, 10, len(result.NodeIds[0]))
}

func TestMigrateAzone(t *testing.T) {
	instanceId := "m-2ke5iF"
	subnetId := "sbn-jxm38c1hjk11"
	zoneName := "cn-gz-a"
	args := MigrateAzoneArgs{
		Subnets: []SubnetMap{
			{
				subnetId,
				zoneName,
			},
		},
		Members: []MemberRoleModel{
			{
				subnetId,
				"primary",
			},
			{
				subnetId,
				"secondary",
			},
			{
				subnetId,
				"hidden",
			},
		},
	}
	err := CLIENT.MigrateAzone(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestInstanceAssignTags(t *testing.T) {
	instanceId := "m-oHGYu8"
	tags := []TagModel{
		// {
		// 	TagKey:   "123",
		// 	TagValue: "13",
		// },
		{
			TagKey:   "test",
			TagValue: "test",
		},
	}
	err := CLIENT.InstanceAssignTags(instanceId, tags)
	ExpectEqual(t.Errorf, nil, err)
}

func TestInstanceBindTags(t *testing.T) {
	instanceId := "m-oHGYu8"
	tags := []TagModel{
		{
			TagKey:   "123",
			TagValue: "13",
		},
		// {
		// 	TagKey:   "test",
		// 	TagValue: "test",
		// },
	}
	err := CLIENT.InstanceBindTags(instanceId, tags)
	ExpectEqual(t.Errorf, nil, err)
}

func TestInstanceUnbindTags(t *testing.T) {
	instanceId := "m-oHGYu8"
	tags := []TagModel{
		// {
		// 	TagKey:   "123",
		// 	TagValue: "13",
		// },
		{
			TagKey:   "test",
			TagValue: "test",
		},
	}
	err := CLIENT.InstanceUnbindTags(instanceId, tags)
	ExpectEqual(t.Errorf, nil, err)
}

func TestCreateBackup(t *testing.T) {
	instanceId := "m-ARidmu"
	backupMethod := "Physical"
	backupDescription := "test"
	result, err := CLIENT.CreateBackup(instanceId, backupMethod, backupDescription)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 13, len(result.BackupId))
}

func TestListBackup(t *testing.T) {
	instanceId := "m-ARidmu"
	args := ListBackupArgs{}
	result, err := CLIENT.ListBackup(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(result.Backups) > 0)
}

func TestGetBackupDetail(t *testing.T) {
	instanceId := "m-ARidmu"
	backupId := "backup-BQTR77"
	result, err := CLIENT.GetBackupDetail(instanceId, backupId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, backupId, result.BackupId)
}

func TestModifyBackupDescription(t *testing.T) {
	instanceId := "m-ARidmu"
	backupId := "backup-BQTR77"
	description1 := "Description1"
	description2 := "Description2"
	args1 := ModifyBackupDescriptionArgs{
		BackupDescription: description1,
	}
	args2 := ModifyBackupDescriptionArgs{
		BackupDescription: description2,
	}

	err := CLIENT.ModifyBackupDescription(instanceId, backupId, &args1)
	ExpectEqual(t.Errorf, nil, err)
	detail, err := CLIENT.GetBackupDetail(instanceId, backupId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, description1, detail.BackupDescription)

	err = CLIENT.ModifyBackupDescription(instanceId, backupId, &args2)
	ExpectEqual(t.Errorf, nil, err)
	detail, err = CLIENT.GetBackupDetail(instanceId, backupId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, description2, detail.BackupDescription)
}

func TestBackupPolicy(t *testing.T) {
	instanceId := "m-ARidmu"
	autoBackupEnableOFF := "OFF"
	autoBackupEnableON := "ON"

	policy, err := CLIENT.GetBackupPolicy(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, autoBackupEnableOFF, policy.AutoBackupEnable)

	policy.AutoBackupEnable = autoBackupEnableON
	policy.PreferredBackupPeriod = "Monday"
	err = CLIENT.ModifyBackupPolicy(instanceId, policy)
	ExpectEqual(t.Errorf, nil, err)

	policy, err = CLIENT.GetBackupPolicy(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, autoBackupEnableON, policy.AutoBackupEnable)

	policy.AutoBackupEnable = autoBackupEnableOFF
	err = CLIENT.ModifyBackupPolicy(instanceId, policy)
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeleteBackup(t *testing.T) {
	instanceId := "m-ARidmu"
	backupId := "backup-XE1CWC"
	err := CLIENT.DeleteBackup(instanceId, backupId)
	ExpectEqual(t.Errorf, nil, err)
}

func TestReleaseMongodb(t *testing.T) {
	instanceId := "m-nsC80W"
	err := CLIENT.ReleaseMongodb(instanceId)
	ExpectEqual(t.Errorf, nil, err)
}

func TestRecoverMongodbs(t *testing.T) {
	instanceIds := []string{"m-nsC80W"}
	err := CLIENT.RecoverMongodbs(instanceIds)
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeleteMongodb(t *testing.T) {
	instanceId := "m-nsC80W"
	err := CLIENT.DeleteMongodb(instanceId)
	ExpectEqual(t.Errorf, nil, err)
}

func TestStartLogging(t *testing.T) {
	instanceId := "m-ARidmu"
	args := StartLoggingArgs{
		// Type: "error",
		Type: "slow",
	}
	err := CLIENT.StartLogging(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestListLogFilesArgs(t *testing.T) {
	instanceId := "m-ARidmu"
	args := ListLogFilesArgs{
		Type:     "running",
		MemberId: "node-gBvCGc",
	}
	result, err := CLIENT.ListLogFiles(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(result.Logs) > 0)
}

func TestGetSecurityIps(t *testing.T) {
	instanceId := "m-ARidmu"
	result, err := CLIENT.GetSecurityIps(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(result.SecurityIps) > 0)
}

func TestAddSecurityIps(t *testing.T) {
	instanceId := "m-ARidmu"
	args := SecurityIpModel{
		SecurityIps: []string{
			"192.168.0.1",
		},
	}
	err := CLIENT.AddSecurityIps(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestDeleteSecurityIps(t *testing.T) {
	instanceId := "m-ARidmu"
	securityIp := "192.168.0.1"
	result, err := CLIENT.GetSecurityIps(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	cnt := 0
	for _, ip := range result.SecurityIps {
		if ip == securityIp {
			cnt++
		}
	}
	ExpectEqual(t.Errorf, 1, cnt)
	args := SecurityIpModel{
		SecurityIps: []string{
			securityIp,
		},
	}
	err = CLIENT.DeleteSecurityIps(instanceId, &args)
	ExpectEqual(t.Errorf, nil, err)

	time.Sleep(10 * time.Second)
	result, err = CLIENT.GetSecurityIps(instanceId)
	ExpectEqual(t.Errorf, nil, err)
	cnt = 0
	for _, ip := range result.SecurityIps {
		if ip == securityIp {
			cnt++
		}
	}
	ExpectEqual(t.Errorf, 0, cnt)
}

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
	err = decoder.Decode(confObj)
	if err != nil {
		log.Fatal("error in Decode:", err)
		os.Exit(1)
	}

	CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

package scs

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

	"github.com/baidubce/bce-sdk-go/model"

	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	SCS_CLIENT  *Client
	SCS_TEST_ID string
	client      *Client
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

const (
	SDK_NAME_PREFIX = "sdk_scs_"
)

var instanceId = SCS_TEST_ID

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

	SCS_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
	client = SCS_CLIENT
	SCS_TEST_ID = "scs-bj-rlmhqtbehihj"
	instanceId = SCS_TEST_ID
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
	id := strconv.FormatInt(time.Now().Unix(), 10)
	args := &CreateInstanceArgs{
		Billing: Billing{
			PaymentTiming: "Postpaid",
			Reservation: &Reservation{
				ReservationLength: 1,
			},
		},
		ClientToken:   getClientToken(),
		PurchaseCount: 1,
		InstanceName:  SDK_NAME_PREFIX + id,
		Port:          6379,
		Engine:        3,
		// EngineVersion:  "5.0",
		NodeType:       "pega.g4s1.micro",
		ClusterType:    "cluster",
		ReplicationNum: 2,
		ShardNum:       1,
		ProxyNum:       2,
		DiskFlavor:     60,
		DiskType:       "essd",
		BgwGroupId:     "",
		StoreType:      3,
		EnableReadOnly: 1,
		ClientAuth:     "ABlockIs16Bytes!",
	}
	result, err := SCS_CLIENT.CreateInstance(args)
	ExpectEqual(t.Errorf, nil, err)

	if len(result.InstanceIds) > 0 {
		SCS_TEST_ID = result.InstanceIds[0]
	}
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}

func TestClient_ListInstances(t *testing.T) {
	args := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(args)
	ExpectEqual(t.Errorf, nil, err)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	for _, e := range result.Instances {
		if e.InstanceID == SCS_TEST_ID {
			ExpectEqual(t.Errorf, "Postpaid", e.PaymentTiming)
		}
	}
}

func TestClient_GetInstanceDetail(t *testing.T) {
	result, err := SCS_CLIENT.GetInstanceDetail(SCS_TEST_ID)
	ExpectEqual(t.Errorf, nil, err)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}

func TestClient_UpdateInstanceName(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &UpdateInstanceNameArgs{
				InstanceName: e.InstanceName + "_new",
				ClientToken:  getClientToken(),
			}
			err := SCS_CLIENT.UpdateInstanceName(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_ResizeInstance(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	args := &ResizeInstanceArgs{
		NodeType:    "cache.n1.mirco",
		ShardNum:    4,
		ClientToken: getClientToken(),
		IsDefer:     true,
	}
	result, err := SCS_CLIENT.GetInstanceDetail(SCS_TEST_ID)
	ExpectEqual(t.Errorf, nil, err)
	if result.InstanceStatus == "Running" {
		err := SCS_CLIENT.ResizeInstance(SCS_TEST_ID, args)
		ExpectEqual(t.Errorf, nil, err)
	}
}

func TestClient_GetCreatePrice(t *testing.T) {
	args := &CreatePriceArgs{
		Engine:         2,
		Period:         1,
		ChargeType:     "prepay",
		NodeType:       "cache.n1.small",
		ReplicationNum: 2,
		ClusterType:    "cluster",
	}
	result, err := SCS_CLIENT.GetCreatePrice(args)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetResizePrice(t *testing.T) {
	args := &ResizePriceArgs{
		ChangeType:     "nodeModify",
		ShardNum:       2,
		ReplicationNum: 1,
		NodeType:       "cache.n1.small",
		Period:         1,
	}
	result, err := SCS_CLIENT.GetResizePrice(instanceId, args)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)

}

func TestClient_AddReplication(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	args := &ReplicationArgs{
		ResizeType: "add",
		ReplicationInfo: []Replication{
			{AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 1},
			{AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 0},
			{AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 0},
			{AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 0},
		},
		ClientToken: getClientToken(),
	}
	result, err := SCS_CLIENT.GetInstanceDetail(SCS_TEST_ID)
	ExpectEqual(t.Errorf, nil, err)
	if result.InstanceStatus == "Running" {
		err := SCS_CLIENT.AddReplication(SCS_TEST_ID, args)
		ExpectEqual(t.Errorf, nil, err)
	}
}

func TestClient_Deletelication(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	args := &ReplicationArgs{
		ResizeType: "delete",
		ReplicationInfo: []Replication{
			{AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 1},
			{AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 0},
			{AvailabilityZone: "cn-bj-a", SubnetId: "sbn-fh56wbtv1ycw", IsMaster: 0},
		},
		ClientToken: getClientToken(),
	}
	result, err := SCS_CLIENT.GetInstanceDetail(SCS_TEST_ID)
	ExpectEqual(t.Errorf, nil, err)
	if result.InstanceStatus == "Running" {
		err := SCS_CLIENT.DeleteReplication(SCS_TEST_ID, args)
		ExpectEqual(t.Errorf, nil, err)
	}
}

func TestClient_RestartInstance(t *testing.T) {
	args := &RestartInstanceArgs{
		IsDefer: true,
	}
	err := SCS_CLIENT.RestartInstance(SCS_TEST_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetNodeTypeList(t *testing.T) {
	_, err := SCS_CLIENT.GetNodeTypeList()
	ExpectEqual(t.Errorf, nil, err)
}

func getClientToken() string {
	return util.NewUUID()
}

func TestClient_ListSubnets(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	args := &ListSubnetsArgs{}
	_, err := SCS_CLIENT.ListSubnets(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateInstanceDomainName(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &UpdateInstanceDomainNameArgs{
				Domain:      "new" + e.Domain,
				ClientToken: getClientToken(),
			}
			err := SCS_CLIENT.UpdateInstanceDomainName(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_GetZoneList(t *testing.T) {
	//isAvailable(SCS_TEST_ID)
	_, err := SCS_CLIENT.GetZoneList()
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ModifyPassword(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &ModifyPasswordArgs{
				Password:    "1234qweR",
				ClientToken: getClientToken(),
			}
			err := SCS_CLIENT.ModifyPassword(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_RenameDomain(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	args := &RenameDomainArgs{
		Domain:      "newDomain",
		ClientToken: getClientToken(),
	}
	err := SCS_CLIENT.RenameDomain(SCS_TEST_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_SwapDomain(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	args := &SwapDomainArgs{
		SourceInstanceId: SCS_TEST_ID,
		TargetInstanceId: "scs-bj-xeelkkdsx",
		ClientToken:      getClientToken(),
	}
	err := SCS_CLIENT.SwapDomain(SCS_TEST_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_FlushInstance(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	time.Sleep(30 * time.Second)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &FlushInstanceArgs{
				Password:    "1234qweR",
				ClientToken: getClientToken(),
			}
			err := SCS_CLIENT.FlushInstance(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_BindingTag(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &BindingTagArgs{
				ChangeTags: []model.TagModel{
					{
						TagKey:   "tag1",
						TagValue: "var1",
					},
				},
			}
			err := SCS_CLIENT.BindingTag(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_UnBindingTag(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	time.Sleep(30 * time.Second)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &BindingTagArgs{
				ChangeTags: []model.TagModel{
					{
						TagKey:   "tag1",
						TagValue: "var1",
					},
				},
			}
			err := SCS_CLIENT.UnBindingTag(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_SetAsMaster(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	err := SCS_CLIENT.SetAsMaster(SCS_TEST_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_SetAsSlave(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	args := &SetAsSlaveArgs{
		MasterDomain: "masterDomain",
		MasterPort:   6379,
	}
	err := SCS_CLIENT.SetAsSlave(SCS_TEST_ID, args)
	ExpectEqual(t.Errorf, nil, err)
}
func TestClient_GetSecurityIp(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			_, err := SCS_CLIENT.GetSecurityIp(e.InstanceID)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_AddSecurityIp(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &SecurityIpArgs{
				SecurityIps: []string{
					"192.0.0.1",
				},
				ClientToken: getClientToken(),
			}
			err := SCS_CLIENT.AddSecurityIp(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_DeleteSecurityIp(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	time.Sleep(30 * time.Second)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &SecurityIpArgs{
				SecurityIps: []string{
					"192.0.0.1",
				},
				ClientToken: getClientToken(),
			}
			err := SCS_CLIENT.DeleteSecurityIp(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_GetParameters(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			_, err := SCS_CLIENT.GetParameters(e.InstanceID)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_ModifyParameters(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &ModifyParametersArgs{
				Parameter: InstanceParam{
					Name:  "timeout",
					Value: "0",
				},
				ClientToken: getClientToken(),
			}
			err := SCS_CLIENT.ModifyParameters(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_GetBackupList(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			_, err := SCS_CLIENT.GetBackupList(e.InstanceID)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_GetBackupDetail(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	result, err := SCS_CLIENT.GetBackupDetail(SCS_TEST_ID, "2587532")
	ExpectEqual(t.Errorf, nil, err)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}
func TestClient_ModifyBackupPolicy(t *testing.T) {
	isAvailable(SCS_TEST_ID)
	listInstancesArgs := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(listInstancesArgs)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Running" == e.InstanceStatus {
			args := &ModifyBackupPolicyArgs{
				BackupDays:  "Sun,Mon,Tue,Wed,Thu,Fri,Sta",
				BackupTime:  "01:05:00",
				ExpireDay:   7,
				ClientToken: getClientToken(),
			}
			err := SCS_CLIENT.ModifyBackupPolicy(e.InstanceID, args)
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func TestClient_DeleteInstance(t *testing.T) {
	//isAvailable(SCS_TEST_ID)
	//time.Sleep(50*time.Second)
	args := &ListInstancesArgs{}
	result, err := SCS_CLIENT.ListInstances(args)
	ExpectEqual(t.Errorf, nil, err)
	for _, e := range result.Instances {
		if strings.HasPrefix(e.InstanceName, SDK_NAME_PREFIX) && "Run"+
			"ning" == e.InstanceStatus && "Postpaid" == e.PaymentTiming {
			err := SCS_CLIENT.DeleteInstance(e.InstanceID, getClientToken())
			ExpectEqual(t.Errorf, nil, err)
		}
	}
}

func isAvailable(instanceId string) {
	for {
		result, err := SCS_CLIENT.GetInstanceDetail(instanceId)
		fmt.Println(instanceId, " => ", result.InstanceStatus)
		if err == nil && result.InstanceStatus == "Running" {
			break
		}
	}
}

func TestClient_ListSecurityGroupByVpcId(t *testing.T) {
	vpcId := "vpc-t7yi6xyrapjz"
	securityGroups, err := client.ListSecurityGroupByVpcId(vpcId)
	if err != nil {
		fmt.Printf("list security group by vpcId error: %+v\n", err)
		return
	}
	for _, group := range securityGroups.Groups {
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
	instanceId := "scs-su-bbjhgxyqyddd"
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
		"scs-su-bbjhgxyqyddd",
	}
	securityGroupIds := []string{
		"g-eun39daa38qf",
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
	securityGroupIds := []string{
		"g-gtj7wknuw3h9",
	}
	args := &UnbindSecurityGroupArgs{
		InstanceId:       "scs-su-bbjhgxyqyddd",
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
		"scs-mjafcdu0",
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

func TestClient_ListRecycleInstances(t *testing.T) {
	marker := &Marker{MaxKeys: 10}
	instances, err := client.ListRecycleInstances(marker)
	if err != nil {
		fmt.Printf("list recycler instances error: %+v\n", err)
		return
	}
	for _, instance := range instances.Result {
		fmt.Println("+-------------------------------------------+")
		fmt.Println("instanceId: ", instance.InstanceID)
		fmt.Println("instanceName: ", instance.InstanceName)
		fmt.Println("engine: ", instance.Engine)
		fmt.Println("engineVersion: ", instance.EngineVersion)
		fmt.Println("instanceStatus: ", instance.InstanceStatus)
		fmt.Println("isolatedStatus: ", instance.IsolatedStatus)
		fmt.Println("PaymentTiming: ", instance.PaymentTiming)
		fmt.Println("ClusterType: ", instance.ClusterType)
		fmt.Println("Domain: ", instance.Domain)
		fmt.Println("Port: ", instance.Port)
		fmt.Println("VnetIP: ", instance.VnetIP)
		fmt.Println("InstanceCreateTime: ", instance.InstanceCreateTime)
		fmt.Println("UsedCapacity: ", instance.UsedCapacity)
		fmt.Println("ZoneNames: ", instance.ZoneNames)
		fmt.Println("tags: ", instance.Tags)
	}
}

func TestClient_RecoverRecyclerInstances(t *testing.T) {
	instanceIds := []string{
		"scs-bj-xjgriqupoftn",
	}
	err := client.RecoverRecyclerInstances(instanceIds)
	if err != nil {
		fmt.Printf("recover recycler instances error: %+v\n", err)
		return
	}
	fmt.Println("recover recycler instances success.")
}

func TestClient_DeleteRecyclerInstances(t *testing.T) {
	instanceIds := []string{
		"scs-bj-xuuasbccatzr",
	}
	err := client.DeleteRecyclerInstances(instanceIds)
	if err != nil {
		fmt.Printf("delete recycler instances error: %+v\n", err)
		return
	}
	fmt.Println("delete recycler instances success.")
}

func TestClient_RenewInstances(t *testing.T) {
	instanceIds := []string{
		"scs-bj-xuuasbccatzr",
	}
	args := &RenewInstanceArgs{
		// 实例Id列表
		InstanceIds: instanceIds,
		// 续费周期，单位为月
		Duration: 1,
	}
	result, err := client.RenewInstances(args)
	if err != nil {
		fmt.Printf("renew instances error: %+v\n", err)
		return
	}
	fmt.Println("renew instances success. orderId:" + result.OrderId)
}
func TestClient_ListLogByInstanceId(t *testing.T) {
	// 一天前
	date := time.Now().
		AddDate(0, 0, -1).
		Format("2006-01-02 03:04:05")
	fmt.Println(date)
	args := &ListLogArgs{
		// 运行日志 runlog 慢日志 slowlog
		FileType: "runlog",
		// 开始时间格式 "yyyy-MM-dd hh:mm:ss"
		StartTime: date,
		// 结束时间,可选,默认返回开始时间+24小时内的日志
		// EndTime: date,
	}
	listLogResult, err := client.ListLogByInstanceId(instanceId, args)
	if err != nil {
		fmt.Printf("list logs of instance error: %+v\n", err)
		return
	}
	fmt.Println("list logs of instance success.")
	for _, shardLog := range listLogResult.LogList {
		fmt.Println("+-------------------------------------------+")
		fmt.Println("shard id: ", shardLog.ShardID)
		fmt.Println("logs size: ", shardLog.TotalNum)
		for _, log := range shardLog.LogItem {
			fmt.Println("log id: ", log.LogID)
			fmt.Println("size: ", log.LogSizeInBytes)
			fmt.Println("start time: ", log.LogStartTime)
			fmt.Println("end time: ", log.LogEndTime)
			fmt.Println("download url: ", log.DownloadURL)
			fmt.Println("download url expires: ", log.DownloadExpires)
		}
	}
}

func TestClient_GetLogById(t *testing.T) {
	args := &GetLogArgs{
		// 下载链接有效时间，单位为秒，可选，默认为1800秒
		ValidSeconds: 60,
	}
	logId := "scs-bj-mktaypucksot_8742_slowlog_202104160330"
	log, err := client.GetLogById(instanceId, logId, args)
	if err != nil {
		fmt.Printf("get log detail of instance error: %+v\n", err)
		return
	}
	fmt.Println("get log detail success.")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("id: ", log.LogID)
	fmt.Println("download url: ", log.DownloadURL)
	fmt.Println("download url expires: ", log.DownloadExpires)
}

func TestClient_GetMaintainTime(t *testing.T) {
	resp, err := client.GetMaintainTime(instanceId)
	if err != nil {
		fmt.Printf("get maintainTime of instance error: %+v\n", err)
		return
	}
	fmt.Println("get maintainTime success.")
	fmt.Println("+-------------------------------------------+")
	fmt.Println("start time: ", resp.MaintainTime.StartTime)
	fmt.Println("dutation: ", resp.MaintainTime.Duration)
	fmt.Println("period: ", resp.MaintainTime.Period)
}

func TestClient_ModifyMaintainTime(t *testing.T) {
	newMaintainTime := &MaintainTime{
		StartTime: "16:00",
		Duration:  1,
		Period:    []int{1, 2, 3},
	}
	err := client.ModifyMaintainTime(instanceId, newMaintainTime)
	if err != nil {
		fmt.Printf("modify maintainTime of instance error: %+v\n", err)
		return
	}
	fmt.Println("modify maintainTime success.")
}

func TestClient_GroupPreCheck(t *testing.T) {
	args := &GroupPreCheckArgs{
		Leader: GroupLeader{
			LeaderRegion: "bj",
			LeaderId:     SCS_TEST_ID,
		},
		Followers: []GroupFollower{
			{
				FollowerId:     "scs-bdbl-dzkqigawuhzy",
				FollowerRegion: "bd",
			},
		},
	}
	result, err := client.GroupPreCheck(args)
	if err != nil {
		fmt.Printf("group pre check error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}

func TestClient_CreateGroup(t *testing.T) {
	args := &CreateGroupArgs{
		Leader: CreateGroupLeader{
			GroupName:    "test_create",
			LeaderRegion: "bj",
			LeaderId:     SCS_TEST_ID,
		},
	}
	result, err := client.CreateGroup(args)
	if err != nil {
		fmt.Printf("group create error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}

func TestClient_GetGroupList(t *testing.T) {
	args := &GetGroupListArgs{
		PageNo:   1,
		PageSize: 10,
	}
	result, err := client.GetGroupList(args)
	if err != nil {
		fmt.Printf("get group list error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}

func TestClient_GetGroupDetail(t *testing.T) {
	result, err := client.GetGroupDetail("scs-group-vobpzinoqadm")
	if err != nil {
		fmt.Printf("get group detail error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}

func TestClient_DeleteGroup(t *testing.T) {
	err := client.DeleteGroup("scs-group-ekeveqhmekvd")
	if err != nil {
		fmt.Printf("delete group error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GroupAddFollower(t *testing.T) {
	args := &FollowerInfo{
		FollowerId:     "scs-bdbl-dzkqigawuhzy",
		FollowerRegion: "bd",
		SyncMaster:     "sync",
	}
	err := client.GroupAddFollower("scs-group-ekeveqhmekvd", args)
	if err != nil {
		fmt.Printf("join group error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GroupRemoveFollower(t *testing.T) {
	err := client.GroupRemoveFollower("scs-group-ekeveqhmekvd", SCS_TEST_ID)
	if err != nil {
		fmt.Printf("quit group error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateGroupName(t *testing.T) {
	args := &GroupNameArgs{
		GroupName: "test_group",
	}
	err := client.UpdateGroupName("scs-group-nqkkmbdjlacx", args)
	if err != nil {
		fmt.Printf("update group name error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_SetAsLeader(t *testing.T) {
	err := client.SetAsLeader("scs-group-nqkkmbdjlacx", SCS_TEST_ID)
	if err != nil {
		fmt.Printf("set as leader error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GroupForbidWrite(t *testing.T) {
	args := &ForbidWriteArgs{
		ForbidWriteFlag: true,
	}
	err := client.GroupForbidWrite("scs-group-nqkkmbdjlacx", args)
	if err != nil {
		fmt.Printf("forbid write error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GroupSetQps(t *testing.T) {
	args := &GroupSetQpsArgs{
		ClusterShowId: "scs-bj-bftgjzjxbmex",
		QpsWrite:      10,
		QpsRead:       20,
	}
	err := client.GroupSetQps("scs-group-nqkkmbdjlacx", args)
	if err != nil {
		fmt.Printf("group set qps error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GroupSyncStatus(t *testing.T) {
	result, err := client.GroupSyncStatus("scs-group-szqbjupjjhpl")
	if err != nil {
		fmt.Printf("group sync status error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}

func TestClient_GroupWhiteList(t *testing.T) {
	result, err := client.GroupWhiteList("scs-group-szqbjupjjhpl")
	if err != nil {
		fmt.Printf("get group white list error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
}

func TestClient_GroupWhiteListAdd(t *testing.T) {
	args := &GroupWhiteList{
		WhiteLists: []string{"127.0.0.1"},
	}
	err := client.GroupWhiteListAdd("scs-group-szqbjupjjhpl", args)
	if err != nil {
		fmt.Printf("group white list add error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GroupWhiteListDelete(t *testing.T) {
	args := &GroupWhiteList{
		WhiteLists: []string{"127.0.0.1"},
	}
	err := client.GroupWhiteListDelete("scs-group-szqbjupjjhpl", args)
	if err != nil {
		fmt.Printf("group white list delete error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GroupStaleReadable(t *testing.T) {
	args := &StaleReadableArgs{
		FollowerId:    SCS_TEST_ID,
		StaleReadable: true,
	}
	err := client.GroupStaleReadable("scs-group-szqbjupjjhpl", args)
	if err != nil {
		fmt.Printf("group stale readable error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_CreateParamsTemplate(t *testing.T) {
	args := &CreateTemplateArgs{
		EngineVersion: "5.0",
		TemplateType:  1,
		ClusterType:   "master_slave",
		Engine:        "redis",
		Name:          "test_template",
		Comment:       "test template",
		Parameters: []ParameterItem{
			{
				ConfName:   "disable_commands",
				ConfModule: 1,
				ConfValue:  "flushall,flushdb",
				ConfType:   3,
			},
		},
	}
	result, err := client.CreateParamsTemplate(args)
	if err != nil {
		fmt.Printf("create params template error: %+v\n", err)
		return
	}
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetParamsTemplateList(t *testing.T) {
	args := &Marker{
		Marker:  "-1",
		MaxKeys: 1000,
	}
	result, err := client.GetParamsTemplateList(args)
	if err != nil {
		fmt.Printf("get params template error: %+v\n", err)
		return
	}
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetParamsTemplateDetail(t *testing.T) {

	result, err := client.GetParamsTemplateDetail("scs-tmpl-kctbndsfdhya")
	if err != nil {
		fmt.Printf("get params template detail error: %+v\n", err)
		return
	}
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteParamsTemplate(t *testing.T) {
	err := client.DeleteParamsTemplate("scs-tmpl-vxslemqppzuz")
	if err != nil {
		fmt.Printf("delete params template error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_RenameParamsTemplate(t *testing.T) {
	args := &RenameTemplateArgs{
		Name: "scs-test-template",
	}
	err := client.RenameParamsTemplate("scs-tmpl-kctbndsfdhya", args)
	if err != nil {
		fmt.Printf("rename params template error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ApplyParamsTemplate(t *testing.T) {
	args := &ApplyTemplateArgs{
		RebootType: 0,
		Extra:      "0",
		CacheClusterShowIdItem: []CacheClusterShowId{
			{
				CacheClusterShowId: SCS_TEST_ID,
				Region:             "bj",
			},
		},
		Parameters: []ParameterItem{
			{
				ConfName:   "disable_commands",
				ConfModule: 1,
				ConfValue:  "flushall,flushdb",
				ConfType:   3,
			},
		},
	}
	err := client.ApplyParamsTemplate("scs-tmpl-kctbndsfdhya", args)
	if err != nil {
		fmt.Printf("apply params template error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_TemplateAddParams(t *testing.T) {
	args := &AddParamsArgs{
		Parameters: []ParameterItem{
			{
				ConfName:   "disable_commands",
				ConfModule: 1,
				ConfValue:  "flushall,flushdb",
				ConfType:   3,
			},
		},
	}
	err := client.TemplateAddParams("scs-tmpl-kctbndsfdhya", args)
	if err != nil {
		fmt.Printf("add params template error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_TemplateModifyParams(t *testing.T) {
	args := &ModifyParamsArgs{
		Parameters: []ParameterItem{
			{
				ConfName:   "disable_commands",
				ConfModule: 1,
				ConfValue:  "flushall,flushdb",
				ConfType:   3,
			},
		},
	}
	err := client.TemplateModifyParams("scs-tmpl-kctbndsfdhya", args)
	if err != nil {
		fmt.Printf("modify params template error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_TemplateDeleteParams(t *testing.T) {
	args := &DeleteParamsArgs{
		Parameters: []string{"appendonly"},
	}
	err := client.TemplateDeleteParams("scs-tmpl-kctbndsfdhya", args)
	if err != nil {
		fmt.Printf("delete params template error: %+v\n", err)
		return
	}
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetapplyTemplate(t *testing.T) {
	args := &GetSystemTemplateArgs{
		Engine:        "redis",
		EngineVersion: "5.0",
		ClusterType:   "master_slave",
	}
	result, err := client.GetSystemTemplate(args)
	if err != nil {
		fmt.Printf("get system template error: %+v\n", err)
		return
	}
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_GetApplyRecords(t *testing.T) {
	args := &Marker{
		Marker:  "-1",
		MaxKeys: 100,
	}
	result, err := client.GetApplyRecords("scs-tmpl-kctbndsfdhya", args)
	if err != nil {
		fmt.Printf("get apply record error: %+v\n", err)
		return
	}
	data, _ := json.Marshal(result)
	fmt.Println(string(data))
	ExpectEqual(t.Errorf, nil, err)
}

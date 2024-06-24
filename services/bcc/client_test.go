package bcc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	BCC_CLIENT              *Client
	BCC_TestCdsId           string
	BCC_TestBccId           string
	BCC_TestSecurityGroupId string
	BCC_TestImageId         string
	BCC_TestSnapshotId      string
	BCC_TestAspId           string
	BCC_TestDeploySetId     string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	BCC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
	// log.SetLogLevel(log.DEBUG)
	BCC_TestBccId = "i-nOKDVvyq"
	BCC_TestCdsId = "cds_id"
	BCC_TestImageId = "m-Q0ezqMIa"
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

func TestCreateInstance(t *testing.T) {
	InternalIps := []string{"ip"}
	DeploySetIds := []string{"DeploySetId1", "DeploySetId2"}
	createInstanceArgs := &api.CreateInstanceArgs{
		ImageId: "m-aCVG7Jxt",
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
		InstanceType:        api.InstanceTypeN5,
		CpuCount:            1,
		MemoryCapacityInGB:  1,
		RootDiskSizeInGb:    40,
		RootDiskStorageType: api.StorageTypeEnhancedPl1,
		ZoneName:            "ZoneName",
		SubnetId:            "SubnetId",
		SecurityGroupId:     "SecurityGroupId",
		RelationTag:         true,
		PurchaseCount:       1,
		Name:                "sdkTest",
		KeypairId:           "KeypairId",
		InternalIps:         InternalIps,
		DeployIdList:        DeploySetIds,
	}
	createResult, err := BCC_CLIENT.CreateInstance(createInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestBccId = createResult.InstanceIds[0]
}

func TestCreateInstanceV2(t *testing.T) {
	InternalIps := []string{"ip"}
	DeploySetIds := []string{"DeploySetId1", "DeploySetId2"}
	RelationTag := true
	IsOpenHostEye := true
	argsV2 := &api.CreateInstanceArgsV2{
		ClientToken:  "clientToken",
		RequestToken: "requestToken",
		ImageId:      "m-aCVG7Jxt",
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
		InstanceType:        api.InstanceTypeN5,
		CpuCount:            1,
		MemoryCapacityInGB:  1,
		RootDiskSizeInGb:    40,
		RootDiskStorageType: api.StorageTypeEnhancedPl1,
		ZoneName:            "ZoneName",
		SubnetId:            "SubnetId",
		SecurityGroupId:     "SecurityGroupId",
		RelationTag:         &RelationTag,
		PurchaseCount:       1,
		Name:                "sdkTest",
		KeypairId:           "KeypairId",
		InternalIps:         InternalIps,
		DeployIdList:        DeploySetIds,
		IsOpenHostEye:       &IsOpenHostEye,
	}

	createResult, err := BCC_CLIENT.CreateInstanceV2(argsV2)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestBccId = createResult.InstanceIds[0]
}

func TestCreateSpecialInstanceBySpec(t *testing.T) {
	createInstanceBySpecArgs := &api.CreateSpecialInstanceBySpecArgs{
		ImageId:  "ImageId",
		Spec:     "bcc.g5.c1m4",
		ZoneName: "cn-bj-a",
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},

		LabelConstraints: []api.LabelConstraint{{
			Key:      "feaA",
			Operator: api.LabelOperatorExist,
		}, {
			Key:      "feaB",
			Value:    "typeB",
			Operator: api.LabelOperatorNotEqual,
		}},
	}
	// 将使用『有 feaA 这个 label』且『feaB 这个 label 的值不是 typeB』的测试机创建实例
	createResult, err := BCC_CLIENT.CreateInstanceByLabel(createInstanceBySpecArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(createResult)
	BCC_TestBccId = createResult.InstanceIds[0]
}

func TestCreateInstanceBySpec(t *testing.T) {
	DeploySetIds := []string{"DeploySetId"}
	createInstanceBySpecArgs := &api.CreateInstanceBySpecArgs{
		ImageId:   "ImageId",
		Spec:      "bcc.l3.c16m64",
		Name:      "sdkTest2",
		AdminPass: "123qaz!@#",
		ZoneName:  "cn-bj-a",
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
		DeployIdList: DeploySetIds,
		EhcClusterId: "ehc-bk4hM1N3",
	}
	createResult, err := BCC_CLIENT.CreateInstanceBySpec(createInstanceBySpecArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(createResult)
	BCC_TestBccId = createResult.InstanceIds[0]
}

func TestCreateInstanceBySpecV2(t *testing.T) {
	DeploySetIds := []string{"DeploySetId"}
	EnableHt := true
	createInstanceBySpecArgs := &api.CreateInstanceBySpecArgsV2{
		ImageId:   "ImageId",
		Spec:      "bcc.l3.c16m64",
		Name:      "sdkTest2",
		AdminPass: "123qaz!@#",
		ZoneName:  "cn-bj-a",
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
		DeployIdList: DeploySetIds,
		EnableHt:     &EnableHt,
	}
	createResult, err := BCC_CLIENT.CreateInstanceBySpecV2(createInstanceBySpecArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(createResult)
	BCC_TestBccId = createResult.InstanceIds[0]
}

func TestCreateInstanceV3(t *testing.T) {
	createInstanceV3Args := &api.CreateInstanceV3Args{
		InstanceSpec: "bcc.l1.c1m1",
		SystemVolume: api.SystemVolume{
			StorageType: api.StorageTypeV3CloudSSDGeneral,
			VolumeSize:  20,
		},
		DataVolumes: []api.DataVolume{
			{
				StorageType: api.StorageTypeV3LocalSSD,
				VolumeSize:  460,
			},
		},
		PurchaseCount: 1,
		InstanceName:  "sdkTest8",
		Password:      "123qaz!@#",
		ZoneName:      "cn-bj-b",
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
			Reservation: &api.Reservation{
				ReservationLength: 1,
			},
		},
		AssociatedResourceTag: true,
		Tags: []model.TagModel{
			{
				TagKey:   "v3",
				TagValue: "1",
			},
		},
		AutoRenewTime: 12,
		CdsAutoRenew:  true,
		// 私有网络子网 IP 数组，当前仅支持批量创建多台实例时支持传入相同子网的多个 IP。
		PrivateIpAddresses: []string{
			"ip",
		},
		DeployIdList: []string{
			// "dset-PAAeNoJt",
			"DeployId",
		},
		ImageId: "ImageId",
		InternetAccessible: api.InternetAccessible{
			InternetMaxBandwidthOut: 5,
			InternetChargeType:      api.TrafficPostpaidByHour,
		},
		InstanceMarketOptions: api.InstanceMarketOptions{
			SpotOption: "custom",
			SpotPrice:  "10",
		},
	}
	createResult, err := BCC_CLIENT.CreateInstanceV3(createInstanceV3Args)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestBccId = createResult.InstanceIds[0]
}

func TestCreateSecurityGroup(t *testing.T) {
	args := &api.CreateSecurityGroupArgs{
		Name:  "testSecurityGroup",
		VpcId: "vpc-uiudcexceb7y",
		Desc:  "vpc1 sdk test create security group",
		Rules: []api.SecurityGroupRuleModel{
			{
				Remark:        "备注",
				Protocol:      "tcp",
				PortRange:     "1-65535",
				Direction:     "ingress",
				SourceIp:      "",
				SourceGroupId: "",
			},
		},
		Tags: []model.TagModel{
			{
				TagKey:   "tagKey",
				TagValue: "tagValue",
			},
		},
	}
	result, err := BCC_CLIENT.CreateSecurityGroup(args)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestSecurityGroupId = result.SecurityGroupId
}

func TestCreateImage(t *testing.T) {
	args := &api.CreateImageArgs{
		ImageName:  "testImageName",
		InstanceId: BCC_TestBccId,
	}
	result, err := BCC_CLIENT.CreateImage(args)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestImageId = result.ImageId
}

func TestListInstances(t *testing.T) {
	listArgs := &api.ListInstanceArgs{
		ZoneName:     "cn-bj-a",
		EhcClusterId: "ehc-bk4hM1N3",
	}
	res, err := BCC_CLIENT.ListInstances(listArgs)
	ExpectEqual(t.Errorf, err, nil)
	// fmt.Println(res.Instances[0].NetEthQueueCount)
	fmt.Println(res)
}

func TestListInstancesByVpcAndIp(t *testing.T) {
	listArgs := &api.ListInstanceArgs{
		VpcId:         "vpc-41avheyaawqc",
		Ipv6Addresses: "240c:4082:2:5d08::2",
	}
	res, err := BCC_CLIENT.ListInstances(listArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res)
}

func TestListRecycleInstances(t *testing.T) {
	listArgs := &api.ListRecycleInstanceArgs{}
	res, err := BCC_CLIENT.ListRecycleInstances(listArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res)
}

func TestGetInstanceDetail(t *testing.T) {
	res, err := BCC_CLIENT.GetInstanceDetail("i-JVXcfQ6M")
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res.Instance.NetEthQueueCount)
	fmt.Println(res)
}

func TestGetInstanceDetailWithDeploySetAndFailed(t *testing.T) {
	res, err := BCC_CLIENT.GetInstanceDetailWithDeploySetAndFailed("InstanceId", false, true)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res)
}

func TestResizeInstance(t *testing.T) {
	resizeArgs := &api.ResizeInstanceArgs{
		CpuCount:           2,
		MemoryCapacityInGB: 4,
	}
	err := BCC_CLIENT.ResizeInstance(BCC_TestBccId, resizeArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestLiveResizeInstance(t *testing.T) {
	resizeArgs := &api.ResizeInstanceArgs{
		CpuCount:           2,
		MemoryCapacityInGB: 4,
		LiveResize:         true,
	}
	err := BCC_CLIENT.ResizeInstance(BCC_TestBccId, resizeArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestStopInstanceWithNoCharge(t *testing.T) {
	err := BCC_CLIENT.StopInstanceWithNoCharge(BCC_TestBccId, true, true)
	ExpectEqual(t.Errorf, err, nil)
}

func TestStopInstance(t *testing.T) {
	err := BCC_CLIENT.StopInstance(BCC_TestBccId, true)
	ExpectEqual(t.Errorf, err, nil)
}

func TestStartInstance(t *testing.T) {
	err := BCC_CLIENT.StartInstance(BCC_TestBccId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRebootInstance(t *testing.T) {
	err := BCC_CLIENT.RebootInstance(BCC_TestBccId, true)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRebuildInstance(t *testing.T) {
	rebuildArgs := &api.RebuildInstanceArgs{
		ImageId:   "ImageId",
		AdminPass: "123qaz!@#",
	}
	err := BCC_CLIENT.RebuildInstance(BCC_TestBccId, rebuildArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRebuildInstanceV2(t *testing.T) {
	rebuildArgs := &api.RebuildInstanceArgsV2{
		ImageId:   "ImageId",
		AdminPass: "123qaz!@#",
	}
	err := BCC_CLIENT.RebuildInstanceV2(BCC_TestBccId, rebuildArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestChangeInstancePass(t *testing.T) {
	changeArgs := &api.ChangeInstancePassArgs{
		AdminPass: "321zaq#@!",
	}
	err := BCC_CLIENT.ChangeInstancePass(BCC_TestBccId, changeArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyInstanceAttribute(t *testing.T) {
	modifyArgs := &api.ModifyInstanceAttributeArgs{
		Name:             "test-modify",
		NetEthQueueCount: "3",
	}
	err := BCC_CLIENT.ModifyInstanceAttribute(BCC_TestBccId, modifyArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyInstanceDesc(t *testing.T) {
	modifyArgs := &api.ModifyInstanceDescArgs{
		Description: "test modify",
	}
	err := BCC_CLIENT.ModifyInstanceDesc(BCC_TestBccId, modifyArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyInstanceHostname(t *testing.T) {
	modifyArgs := &api.ModifyInstanceHostnameArgs{
		Hostname:             "test-modify-domain",
		IsOpenHostnameDomain: false,
		Reboot:               true,
	}
	err := BCC_CLIENT.ModifyInstanceHostname(BCC_TestBccId, modifyArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetInstanceVNC(t *testing.T) {
	_, err := BCC_CLIENT.GetInstanceVNC(BCC_TestBccId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBatchAddIp(t *testing.T) {
	privateIps := []string{"privateIp"}
	batchAddIpArgs := &api.BatchAddIpArgs{
		InstanceId:                     BCC_TestBccId,
		PrivateIps:                     privateIps,
		SecondaryPrivateIpAddressCount: 1,
	}
	result, err := BCC_CLIENT.BatchAddIP(batchAddIpArgs)
	fmt.Println(result)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBatchDelIp(t *testing.T) {
	privateIps := []string{"privateIp"}
	batchDelIpArgs := &api.BatchDelIpArgs{
		InstanceId: BCC_TestBccId,
		PrivateIps: privateIps,
	}
	err := BCC_CLIENT.BatchDelIP(batchDelIpArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBindSecurityGroup(t *testing.T) {
	err := BCC_CLIENT.BindSecurityGroup(BCC_TestBccId, BCC_TestSecurityGroupId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUnBindSecurityGroup(t *testing.T) {
	err := BCC_CLIENT.UnBindSecurityGroup(BCC_TestBccId, BCC_TestSecurityGroupId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateCDSVolume(t *testing.T) {
	args := &api.CreateCDSVolumeArgs{
		PurchaseCount: 1,
		CdsSizeInGB:   40,
		Billing: &api.Billing{
			PaymentTiming: api.PaymentTimingPrePaid,
			Reservation: &api.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "MONTH",
			},
		},
		AutoSnapshotPolicy: []api.AutoSnapshotPolicy{
			{
				AutoSnapshotPolicyId: "Test_AutoSnapshotPolicyId",
			},
		},
		RenewTimeUnit: "month",
		RenewTime:     2,
	}

	result, _ := BCC_CLIENT.CreateCDSVolume(args)
	BCC_TestCdsId = result.VolumeIds[0]
	fmt.Print(BCC_TestCdsId)
}

func TestCreateCDSVolumeV3(t *testing.T) {
	args := &api.CreateCDSVolumeV3Args{
		ZoneName:      "cn-bj-a",
		VolumeName:    "volumeV3Test",
		Description:   "v3 test",
		PurchaseCount: 1,
		VolumeSize:    50,
		StorageType:   api.StorageTypeV3CloudSSDGeneral,
		Billing: &api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
	}

	result, err := BCC_CLIENT.CreateCDSVolumeV3(args)
	fmt.Println(result)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestCdsId = result.VolumeIds[0]
}

func TestGetBidInstancePrice(t *testing.T) {
	args := &api.GetBidInstancePriceArgs{
		InstanceType:        api.InstanceTypeN1,
		CpuCount:            1,
		MemoryCapacityInGB:  2,
		RootDiskSizeInGb:    45,
		RootDiskStorageType: api.StorageTypeCloudHP1,
		PurchaseCount:       1,
	}
	result, err := BCC_CLIENT.GetBidInstancePrice(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestListBidFlavor(t *testing.T) {
	result, err := BCC_CLIENT.ListBidFlavor()
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestCreateSnapshot(t *testing.T) {
	args := &api.CreateSnapshotArgs{
		VolumeId:     BCC_TestCdsId,
		SnapshotName: "testSnapshotName",
		Tags: []model.TagModel{
			{
				TagKey:   "test",
				TagValue: "val",
			},
		},
	}
	result, err := BCC_CLIENT.CreateSnapshot(args)
	ExpectEqual(t.Errorf, err, nil)

	BCC_TestSnapshotId = result.SnapshotId
}

func TestListSnapshot(t *testing.T) {
	args := &api.ListSnapshotArgs{}
	_, err := BCC_CLIENT.ListSnapshot(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListSnapshotChain(t *testing.T) {
	args := &api.ListSnapshotChainArgs{
		OrderBy:  "chainId",
		Order:    "asc",
		PageSize: 2,
		PageNo:   2,
	}
	res, err := BCC_CLIENT.ListSnapshotChain(args)
	fmt.Println(res)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetSnapshotDetail(t *testing.T) {
	_, err := BCC_CLIENT.GetSnapshotDetail(BCC_TestSnapshotId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteSnapshot(t *testing.T) {
	err := BCC_CLIENT.DeleteSnapshot(BCC_TestSnapshotId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDescribeRegions(t *testing.T) {
	queryArgs := &api.DescribeRegionsArgs{
		Region: "",
	}
	if res, err := BCC_CLIENT.DescribeRegions(queryArgs); err != nil {
		fmt.Println("list all region's endpoint information failed: ", err)
	} else {
		fmt.Println("list all region's endpoint information: ", res)
	}
}

func TestCreateAutoSnapshotPolicy(t *testing.T) {
	args := &api.CreateASPArgs{
		Name:           "testAspName",
		TimePoints:     []string{"20"},
		RepeatWeekdays: []string{"1", "5"},
		RetentionDays:  "7",
	}
	result, err := BCC_CLIENT.CreateAutoSnapshotPolicy(args)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestAspId = result.AspId
}

func TestAttachAutoSnapshotPolicy(t *testing.T) {
	args := &api.AttachASPArgs{
		VolumeIds: []string{BCC_TestCdsId},
	}
	err := BCC_CLIENT.AttachAutoSnapshotPolicy(BCC_TestAspId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDetachAutoSnapshotPolicy(t *testing.T) {
	args := &api.DetachASPArgs{
		VolumeIds: []string{BCC_TestCdsId},
	}
	err := BCC_CLIENT.DetachAutoSnapshotPolicy(BCC_TestAspId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUpdateAutoSnapshotPolicy(t *testing.T) {
	args := &api.UpdateASPArgs{
		AspId:          "AspId",
		Name:           "Name",
		TimePoints:     []string{"21"},
		RepeatWeekdays: []string{"2"},
	}
	err := BCC_CLIENT.UpdateAutoSnapshotPolicy(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListAutoSnapshotPolicy(t *testing.T) {
	args := &api.ListASPArgs{}
	_, err := BCC_CLIENT.ListAutoSnapshotPolicy(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetAutoSnapshotPolicy(t *testing.T) {
	_, err := BCC_CLIENT.GetAutoSnapshotPolicy(BCC_TestAspId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteAutoSnapshotPolicy(t *testing.T) {
	err := BCC_CLIENT.DeleteAutoSnapshotPolicy(BCC_TestAspId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListCDSVolume(t *testing.T) {
	queryArgs := &api.ListCDSVolumeArgs{
		InstanceId: BCC_TestBccId,
	}
	res, err := BCC_CLIENT.ListCDSVolume(queryArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res)
}

func TestListCDSVolumeV3(t *testing.T) {
	queryArgs := &api.ListCDSVolumeArgs{
		InstanceId: BCC_TestBccId,
	}
	res, err := BCC_CLIENT.ListCDSVolumeV3(queryArgs)
	fmt.Println(res)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetCDSVolumeDetail(t *testing.T) {
	res, err := BCC_CLIENT.GetCDSVolumeDetail(BCC_TestCdsId)
	fmt.Println(res.Volume)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetCDSVolumeDetailV3(t *testing.T) {
	res, err := BCC_CLIENT.GetCDSVolumeDetailV3(BCC_TestCdsId)
	fmt.Println(res.Volume)
	ExpectEqual(t.Errorf, err, nil)
}

func TestAttachCDSVolume(t *testing.T) {
	args := &api.AttachVolumeArgs{
		InstanceId: BCC_TestBccId,
	}

	_, err := BCC_CLIENT.AttachCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDetachCDSVolume(t *testing.T) {
	args := &api.DetachVolumeArgs{
		InstanceId: BCC_TestBccId,
	}

	err := BCC_CLIENT.DetachCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestResizeCDSVolume(t *testing.T) {
	args := &api.ResizeCSDVolumeArgs{
		NewCdsSizeInGB: 461,
		NewVolumeType:  "enhanced_ssd_pl2",
	}

	res, err := BCC_CLIENT.ResizeCDSVolume(BCC_TestCdsId, args)

	fmt.Println(res)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRollbackCDSVolume(t *testing.T) {
	args := &api.RollbackCSDVolumeArgs{
		SnapshotId: "SnapshotId",
	}

	err := BCC_CLIENT.RollbackCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestPurchaseReservedCDSVolume(t *testing.T) {
	args := &api.PurchaseReservedCSDVolumeArgs{
		Billing: &api.Billing{
			Reservation: &api.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "Month",
			},
		},
	}

	err := BCC_CLIENT.PurchaseReservedCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRenameCDSVolume(t *testing.T) {
	args := &api.RenameCSDVolumeArgs{
		Name: "testRenamedName",
	}

	err := BCC_CLIENT.RenameCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyCDSVolume(t *testing.T) {
	args := &api.ModifyCSDVolumeArgs{
		CdsName: "modifiedName",
		Desc:    "modifiedDesc",
	}

	err := BCC_CLIENT.ModifyCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyChargeTypeCDSVolume(t *testing.T) {
	args := &api.ModifyChargeTypeCSDVolumeArgs{
		Billing: &api.Billing{
			Reservation: &api.Reservation{
				ReservationLength: 1,
			},
		},
	}

	err := BCC_CLIENT.ModifyChargeTypeCDSVolume(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteCDSVolumeNew(t *testing.T) {
	args := &api.DeleteCDSVolumeArgs{
		AutoSnapshot: "on",
		Recycle:      "on",
	}

	err := BCC_CLIENT.DeleteCDSVolumeNew(BCC_TestCdsId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteCDSVolume(t *testing.T) {
	err := BCC_CLIENT.DeleteCDSVolume(BCC_TestCdsId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestAutoRenewCDSVolume(t *testing.T) {
	args := &api.AutoRenewCDSVolumeArgs{
		VolumeId:      "VolumeId",
		RenewTimeUnit: "month",
		RenewTime:     2,
	}

	err := BCC_CLIENT.AutoRenewCDSVolume(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCancelAutoRenewCDSVolume(t *testing.T) {
	args := &api.CancelAutoRenewCDSVolumeArgs{
		VolumeId: "VolumeId",
	}

	err := BCC_CLIENT.CancelAutoRenewCDSVolume(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListSecurityGroup(t *testing.T) {
	queryArgs := &api.ListSecurityGroupArgs{
		VpcId: "vpc-uiudcexceb7y",
	}
	result, err := BCC_CLIENT.ListSecurityGroup(queryArgs)
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
	ExpectEqual(t.Errorf, err, nil)
}

func TestInstanceChangeVpc(t *testing.T) {
	args := &api.InstanceChangeVpcArgs{
		InstanceId:                 "InstanceId",
		SubnetId:                   "SubnetId",
		Reboot:                     false,
		SecurityGroupIds:           []string{"SecurityGroupId"},
		EnterpriseSecurityGroupIds: []string{"EnterpriseSecurityGroupIds"},
	}

	err := BCC_CLIENT.InstanceChangeVpc(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestInstanceChangeSubnet(t *testing.T) {
	args := &api.InstanceChangeSubnetArgs{
		InstanceId:       "i-YRMFRB6Z",
		SubnetId:         "sbn-z1y9tcedqnh6",
		InternalIp:       "192.168.5.12",
		Reboot:           false,
		SecurityGroupIds: []string{"g-i24fkh******"},
	}

	err := BCC_CLIENT.InstanceChangeSubnet(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestAuthorizeSecurityGroupRule(t *testing.T) {
	args := &api.AuthorizeSecurityGroupArgs{
		Rule: &api.SecurityGroupRuleModel{
			Remark:        "备注",
			Protocol:      "udp",
			PortRange:     "1-65535",
			Direction:     "ingress",
			SourceIp:      "",
			SourceGroupId: "",
		},
	}
	err := BCC_CLIENT.AuthorizeSecurityGroupRule(BCC_TestSecurityGroupId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRevokeSecurityGroupRule(t *testing.T) {
	args := &api.RevokeSecurityGroupArgs{
		Rule: &api.SecurityGroupRuleModel{
			Remark:        "备注",
			Protocol:      "udp",
			PortRange:     "1-65535",
			Direction:     "ingress",
			SourceIp:      "",
			SourceGroupId: "",
		},
	}
	err := BCC_CLIENT.RevokeSecurityGroupRule(BCC_TestSecurityGroupId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteSecurityGroupRule(t *testing.T) {
	err := BCC_CLIENT.DeleteSecurityGroup(BCC_TestSecurityGroupId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListImage(t *testing.T) {
	queryArgs := &api.ListImageArgs{}
	_, err := BCC_CLIENT.ListImage(queryArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetImageDetail(t *testing.T) {
	_, err := BCC_CLIENT.GetImageDetail(BCC_TestImageId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRemoteCopyImage(t *testing.T) {
	args := &api.RemoteCopyImageArgs{
		Name:       "testRemoteCopy",
		DestRegion: []string{"hkg"},
	}
	err := BCC_CLIENT.RemoteCopyImage(BCC_TestImageId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRemoteCopyImageReturnImageIds(t *testing.T) {
	args := &api.RemoteCopyImageArgs{
		Name:       "Copy",
		DestRegion: []string{"hkg"},
	}
	result, err := BCC_CLIENT.RemoteCopyImageReturnImageIds(BCC_TestImageId, args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestCancelRemoteCopyImage(t *testing.T) {
	err := BCC_CLIENT.CancelRemoteCopyImage(BCC_TestImageId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestShareImage(t *testing.T) {
	args := &api.SharedUser{
		AccountId: "id",
	}
	err := BCC_CLIENT.ShareImage(BCC_TestImageId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUnShareImage(t *testing.T) {
	args := &api.SharedUser{
		AccountId: "id",
	}
	err := BCC_CLIENT.UnShareImage(BCC_TestImageId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetImageSharedUser(t *testing.T) {
	_, err := BCC_CLIENT.GetImageSharedUser(BCC_TestImageId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetImageOS(t *testing.T) {
	args := &api.GetImageOsArgs{}
	_, err := BCC_CLIENT.GetImageOS(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteImage(t *testing.T) {
	err := BCC_CLIENT.DeleteImage(BCC_TestImageId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteInstance(t *testing.T) {
	err := BCC_CLIENT.DeleteInstance(BCC_TestBccId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteInstanceWithRelateResource(t *testing.T) {
	args := &api.DeleteInstanceWithRelateResourceArgs{
		BccRecycleFlag:  true,
		DeleteImmediate: true,
	}

	err := BCC_CLIENT.DeleteInstanceWithRelateResource(BCC_TestBccId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListSpec(t *testing.T) {
	_, err := BCC_CLIENT.ListSpec()
	ExpectEqual(t.Errorf, err, nil)
}

func TestListZone(t *testing.T) {
	_, err := BCC_CLIENT.ListZone()
	ExpectEqual(t.Errorf, err, nil)
}

func TestListFlavorSpec(t *testing.T) {
	args := &api.ListFlavorSpecArgs{}
	res, err := BCC_CLIENT.ListFlavorSpec(args)
	ExpectEqual(t.Errorf, err, nil)
	// fmt.Println(res.ZoneResources[0].BccResources.FlavorGroups[0].Flavors[0].NetEthQueueCount)
	// fmt.Println(res.ZoneResources[0].BccResources.FlavorGroups[0].Flavors[0].NetEthMaxQueueCount)
	fmt.Println(res)
}

func TestGetPriceBySpec(t *testing.T) {
	args := &api.GetPriceBySpecArgs{
		SpecId:         "g1",
		Spec:           "",
		PaymentTiming:  "Postpaid",
		ZoneName:       "cn-gz-b",
		PurchaseCount:  1,
		PurchaseLength: 1,
	}
	res, err := BCC_CLIENT.GetPriceBySpec(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res)
}

func TestDeletePrepaidInstanceWithRelateResource(t *testing.T) {
	args := &api.DeletePrepaidInstanceWithRelateResourceArgs{
		InstanceId:            BCC_TestBccId,
		RelatedReleaseFlag:    true,
		DeleteCdsSnapshotFlag: true,
		DeleteRelatedEnisFlag: true,
	}
	result, err := BCC_CLIENT.DeletePrepaidInstanceWithRelateResource(args)
	fmt.Println(result)
	fmt.Println(err)
}

func TestCreateDeploySet(t *testing.T) {
	testDeploySetName := "testName"
	testDeployDesc := "testDesc"
	testStrategy := "HOST_HA"
	queryArgs := &api.CreateDeploySetArgs{
		Strategy:    testStrategy,
		Name:        testDeploySetName,
		Desc:        testDeployDesc,
		Concurrency: 5,
	}
	rep, err := BCC_CLIENT.CreateDeploySet(queryArgs)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListDeploySets(t *testing.T) {
	rep, err := BCC_CLIENT.ListDeploySets()
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyDeploySet(t *testing.T) {
	testDeploySetName := "testName"
	testDeployDesc := "goDesc"
	queryArgs := &api.ModifyDeploySetArgs{
		Name: testDeploySetName,
		Desc: testDeployDesc,
	}
	BCC_TestDeploySetId = "DeploySetId"
	rep, err := BCC_CLIENT.ModifyDeploySet(BCC_TestDeploySetId, queryArgs)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteDeploySet(t *testing.T) {
	testDeleteDeploySetId := "DeploySetId"
	err := BCC_CLIENT.DeleteDeploySet(testDeleteDeploySetId)
	fmt.Println(err)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetDeploySet(t *testing.T) {
	testDeploySetID := "DeploySetId"
	rep, err := BCC_CLIENT.GetDeploySet(testDeploySetID)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUpdateInstanceDeploySet(t *testing.T) {
	queryArgs := &api.UpdateInstanceDeployArgs{
		InstanceId:   "InstanceId",
		DeploySetIds: []string{"DeploySetId"},
	}
	rep, err := BCC_CLIENT.UpdateInstanceDeploySet(queryArgs)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDelInstanceDeploySet(t *testing.T) {
	queryArgs := &api.DelInstanceDeployArgs{
		DeploySetId: "DeploySetId",
		InstanceIds: []string{"InstanceId"},
	}
	rep, err := BCC_CLIENT.DelInstanceDeploySet(queryArgs)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestResizeInstanceBySpec(t *testing.T) {
	resizeArgs := &api.ResizeInstanceArgs{
		Spec: "Spec",
	}
	err := BCC_CLIENT.ResizeInstanceBySpec(BCC_TestBccId, resizeArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBatchRebuildInstances(t *testing.T) {
	rebuildArgs := &api.RebuildBatchInstanceArgs{
		ImageId:     "ImageId",
		AdminPass:   "123qaz!@#",
		InstanceIds: []string{BCC_TestBccId},
	}
	err := BCC_CLIENT.BatchRebuildInstances(rebuildArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBatchRebuildInstancesV2(t *testing.T) {
	rebuildArgs := &api.RebuildBatchInstanceArgsV2{
		ImageId:     "ImageId",
		AdminPass:   "123qaz!@#",
		InstanceIds: []string{BCC_TestBccId},
	}
	err := BCC_CLIENT.BatchRebuildInstancesV2(rebuildArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestChangeToPrepaid(t *testing.T) {
	args := &api.ChangeToPrepaidRequest{
		Duration:    1,
		RelationCds: true,
	}
	_, err := BCC_CLIENT.ChangeToPrepaid(BCC_TestBccId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBindInstanceToTags(t *testing.T) {
	args := &api.BindTagsRequest{
		ChangeTags: []model.TagModel{
			{
				TagKey:   "TagKey",
				TagValue: "TagValue",
			},
		},
	}
	err := BCC_CLIENT.BindInstanceToTags(BCC_TestBccId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUnBindInstanceToTags(t *testing.T) {
	args := &api.UnBindTagsRequest{
		ChangeTags: []model.TagModel{
			{
				TagKey:   "TagKey",
				TagValue: "TagValue",
			},
		},
	}
	err := BCC_CLIENT.UnBindInstanceToTags(BCC_TestBccId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBindReservedInstanceToTags(t *testing.T) {
	args := &api.ReservedTagsRequest{
		ChangeTags: []model.TagModel{
			{
				TagKey:   "TagKey-go",
				TagValue: "TagValue",
			},
		},
		ReservedInstanceIds: []string{
			"r-Qyycx1SX",
		},
	}
	err := BCC_CLIENT.BindReservedInstanceToTags(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUnbindReservedInstanceToTags(t *testing.T) {
	args := &api.ReservedTagsRequest{
		ChangeTags: []model.TagModel{
			{
				TagKey:   "TagKey-go",
				TagValue: "TagValue",
			},
		},
		ReservedInstanceIds: []string{
			"r-Qyycx1SX",
		},
	}
	err := BCC_CLIENT.UnbindReservedInstanceFromTags(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBindInstanceToTagsByResourceType(t *testing.T) {
	args := &api.TagsOperationRequest{
		ResourceType: "bccri",
		ResourceIds: []string{
			"r-oFpMXKhv", "r-HrztSVk0",
		},
		Tags: []model.TagModel{
			{
				TagKey:   "TagKey-go",
				TagValue: "TagValue",
			},
		},
		IsRelationTag: false,
	}
	err := BCC_CLIENT.BindInstanceToTagsByResourceType(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUnbindInstanceToTagsByResourceType(t *testing.T) {
	args := &api.TagsOperationRequest{
		ResourceType: "bccri",
		ResourceIds: []string{
			"r-oFpMXKhv", "r-HrztSVk0",
		},
		Tags: []model.TagModel{
			{
				TagKey:   "TagKey-go",
				TagValue: "TagValue",
			},
		},
		IsRelationTag: false,
	}
	err := BCC_CLIENT.UnbindInstanceToTagsByResourceType(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetInstanceNoChargeList(t *testing.T) {
	listArgs := &api.ListInstanceArgs{}
	_, err := BCC_CLIENT.GetInstanceNoChargeList(listArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateBidInstance(t *testing.T) {
	createInstanceArgs := &api.CreateInstanceArgs{
		ImageId: "ImageId",
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingBidding,
		},
		InstanceType:        api.InstanceTypeN3,
		CpuCount:            1,
		MemoryCapacityInGB:  4,
		RootDiskSizeInGb:    40,
		RootDiskStorageType: api.StorageTypeHP1,
		ZoneName:            "zoneName",
		SubnetId:            "SubnetId",
		SecurityGroupId:     "SecurityGroupId",
		RelationTag:         true,
		PurchaseCount:       1,
		Name:                "sdkTest",
		BidModel:            "BidModel",
		BidPrice:            "BidPrice",
	}
	createResult, err := BCC_CLIENT.CreateBidInstance(createInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
	BCC_TestBccId = createResult.InstanceIds[0]
}

func TestCancelBidOrder(t *testing.T) {
	createInstanceArgs := &api.CancelBidOrderRequest{
		OrderId: "OrderId",
	}
	_, err := BCC_CLIENT.CancelBidOrder(createInstanceArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestInstancePurchaseReserved(t *testing.T) {
	purchaseReservedArgs := &api.PurchaseReservedArgs{
		Billing: api.Billing{
			PaymentTiming: api.PaymentTimingPrePaid,
			Reservation: &api.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "MONTH",
			},
		},
		RelatedRenewFlag: "",
	}
	_, err := BCC_CLIENT.InstancePurchaseReserved(BCC_TestBccId, purchaseReservedArgs)
	// fmt.Print(err)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetAvailableDiskInfo(t *testing.T) {
	zoneName := "cn-bj-a"
	if res, err := BCC_CLIENT.GetAvailableDiskInfo(zoneName); err != nil {
		fmt.Println("Get the specific zone flavor failed: ", err)
	} else {
		fmt.Println("Get the specific zone flavor success, result: ", res)
	}
}

func TestListTypeZones(t *testing.T) {
	args := &api.ListTypeZonesArgs{
		InstanceType: "",
		ProductType:  "",
		Spec:         "bcc.g3.c2m12",
		SpecId:       "",
	}
	if res, err := BCC_CLIENT.ListTypeZones(args); err != nil {
		fmt.Println("Get the specific zone flavor failed: ", err)
	} else {
		fmt.Println("Get the specific zone flavor success, result: ", res)
	}
}

func TestListInstanceEnis(t *testing.T) {
	instanceId := "InstanceId"
	if res, err := BCC_CLIENT.ListInstanceEnis(instanceId); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success, result: ", res)
	}
}

func TestCreateKeypair(t *testing.T) {
	args := &api.CreateKeypairArgs{
		Name:        "gosdk",
		Description: "go sdk test",
	}
	if res, err := BCC_CLIENT.CreateKeypair(args); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success, result: ", res)
	}
}

func TestImportKeypair(t *testing.T) {
	args := &api.ImportKeypairArgs{
		ClientToken: "",
		Name:        "goImport",
		Description: "go sdk test",
		PublicKey:   "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCNItVsPPOYbMH4W5fyFqoYZwfL2A1G9IWgofhrrNYVmUr22qx42FPcyR6Fj1frHGNUIZ0NN3CzS8wXg/aKWJkYMiZGjlwmppdrNGWUjmPZD9GbHw/w8sVGCBEyyCEVlTZHQe+AgfzOr/yzqpUmCareBIlQDlR1PzX39wDf7ohpzmJy2e+B+amNy2pgsxG9OI9a4RacGLAeD/OTE/nvj027pEwbWbxM1BsJjrMeH51gWGqv8zANJFL2MGqdBaUGH0r4iXTWGZ+TkA1L7np8qWNCwquve2iy8dlHw7OnzA+hsFVZJROjJimzMY+yNNiy3CqzdO+WaBXG9MWUxtLf3ZjF",
	}
	if res, err := BCC_CLIENT.ImportKeypair(args); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success, result: ", res)
	}
}

func TestListKeypairs(t *testing.T) {
	args := &api.ListKeypairArgs{
		Marker:  "",
		MaxKeys: 0,
		Name:    "ac",
	}
	if res, err := BCC_CLIENT.ListKeypairs(args); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success, result: ", res)
	}
}

func TestRenameKeypair(t *testing.T) {
	args := &api.RenameKeypairArgs{
		Name:      "renameKeypair",
		KeypairId: "KeypairId",
	}
	if err := BCC_CLIENT.RenameKeypair(args); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success")
	}
}

func TestUpdateKeypairDescription(t *testing.T) {
	args := &api.KeypairUpdateDescArgs{
		KeypairId:   "KeypairId",
		Description: "UpdateKeypairDescription test",
	}
	if err := BCC_CLIENT.UpdateKeypairDescription(args); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success")
	}
}

func TestGetKeypairDetail(t *testing.T) {
	keypairId := "KeypairId"
	if resp, err := BCC_CLIENT.GetKeypairDetail(keypairId); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success resp:", resp.Keypair.InstanceCount)
	}
}

func TestAttachKeypair(t *testing.T) {
	args := &api.AttackKeypairArgs{
		KeypairId:   "KeypairId",
		InstanceIds: []string{"InstanceId"},
	}
	if err := BCC_CLIENT.AttachKeypair(args); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success")
	}
}

func TestDetachKeypair(t *testing.T) {
	args := &api.DetachKeypairArgs{
		KeypairId:   "KeypairId",
		InstanceIds: []string{"InstanceId"},
	}
	if err := BCC_CLIENT.DetachKeypair(args); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success")
	}
}

func TestDeleteKeypair(t *testing.T) {
	args := &api.DeleteKeypairArgs{
		KeypairId: "KeypairId",
	}
	if err := BCC_CLIENT.DeleteKeypair(args); err != nil {
		fmt.Println("Get specific instance eni failed: ", err)
	} else {
		fmt.Println("Get specific instance eni success")
	}
}

func TestBatchCreateAutoRenewRules(t *testing.T) {
	bccAutoRenewArgs := &api.BccCreateAutoRenewArgs{
		InstanceId:    BCC_TestBccId,
		RenewTimeUnit: "month",
		RenewTime:     1,
	}
	err := BCC_CLIENT.BatchCreateAutoRenewRules(bccAutoRenewArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBatchDeleteAutoRenewRules(t *testing.T) {
	bccAutoRenewArgs := &api.BccDeleteAutoRenewArgs{
		InstanceId: BCC_TestBccId,
	}
	err := BCC_CLIENT.BatchDeleteAutoRenewRules(bccAutoRenewArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteInstanceIngorePayment(t *testing.T) {
	args := &api.DeleteInstanceIngorePaymentArgs{
		InstanceId:            "InstanceId",
		RelatedReleaseFlag:    true,
		DeleteRelatedEnisFlag: true,
		DeleteCdsSnapshotFlag: true,
	}
	if res, err := BCC_CLIENT.DeleteInstanceIngorePayment(args); err != nil {
		fmt.Println("delete  instance failed: ", err)
	} else {
		fmt.Println("delete  instance success, result: ", res)
	}
}

func TestRecoveryInstance(t *testing.T) {
	args := &api.RecoveryInstanceArgs{
		InstanceIds: []api.RecoveryInstanceModel{
			{
				InstanceId: BCC_TestBccId,
			},
		},
	}
	if err := BCC_CLIENT.RecoveryInstance(args); err != nil {
		fmt.Println("recovery instance failed: ", err)
	} else {
		fmt.Println("recovery instance success")
	}
}

func TestGetAllStocks(t *testing.T) {
	if res, err := BCC_CLIENT.GetAllStocks(); err != nil {
		fmt.Println("get all stocks failed: ", err)
	} else {
		fmt.Println("get all stocks success, result: ", res)
	}
}

func TestGetStockWithDeploySet(t *testing.T) {
	args := &api.GetStockWithDeploySetArgs{
		Spec:         "ehc.lgn5.c128m1024.8a100.8re.4d",
		DeploySetIds: []string{"dset-Z3aEKdeY"},
		EhcClusterId: "ehc-bk4hM1N3",
	}
	if res, err := BCC_CLIENT.GetStockWithDeploySet(args); err != nil {
		fmt.Println("get stock with deploySet failed: ", err)
	} else {
		fmt.Println("get stock with deploySet, result: ", res)
	}
}

func TestGetStockWithSpec(t *testing.T) {
	args := &api.GetStockWithSpecArgs{
		Spec:         "bcc.g3.c2m8",
		DeploySetIds: []string{"dset-RekVqK7V"},
	}
	if res, err := BCC_CLIENT.GetStockWithSpec(args); err != nil {
		fmt.Println("get stock with spec failed: ", err)
	} else {
		fmt.Println("get stock with spec, result: ", res)
	}
}

func TestListInstanceByInstanceIds(t *testing.T) {
	args := &api.ListInstanceByInstanceIdArgs{
		InstanceIds: []string{"i-gRYyYyjr", "i-GGc7Buqs"},
		Marker:      "",
		MaxKeys:     3,
	}
	result, err := BCC_CLIENT.ListInstanceByInstanceIds(args)
	if err != nil {
		fmt.Println("list instance failed: ", err)
	} else {
		fmt.Println("list instance  success")
		data, e := json.Marshal(result)
		if e != nil {
			fmt.Println("json marshal failed!")
			return
		}
		fmt.Printf("list instance : %s", data)
	}
}

func TestListServersByMarkerV3(t *testing.T) {
	args := &api.ListServerRequestV3Args{
		Marker:  "",
		MaxKeys: 3,
	}
	result, err := BCC_CLIENT.ListServersByMarkerV3(args)
	if err != nil {
		fmt.Println("list instance failed: ", err)
	} else {
		fmt.Println("list instance  success")
		data, e := json.Marshal(result)
		if e != nil {
			fmt.Println("json marshal failed!")
			return
		}
		fmt.Printf("list instance : %s", data)
	}
}

func TestDeletePrepayVolume(t *testing.T) {
	args := &api.VolumePrepayDeleteRequestArgs{
		VolumeId:           "v-tVDW1NkK",
		RelatedReleaseFlag: false,
	}
	result, err := BCC_CLIENT.DeletePrepayVolume(args)
	if err != nil {
		fmt.Println("delete volume failed: ", err)
	} else {
		fmt.Println("delete volume  success")
		data, e := json.Marshal(result)
		if e != nil {
			fmt.Println("json marshal failed!")
			return
		}
		fmt.Printf("delete volume : %s", data)
	}
}

func TestBatchDeleteInstanceWithRelateResource(t *testing.T) {
	args := &api.BatchDeleteInstanceWithRelateResourceArgs{
		RelatedReleaseFlag: true,
		BccRecycleFlag:     true,
		InstanceIds:        []string{"i-gRYyYyjx", "i-GGc7Buqd"},
	}

	err := BCC_CLIENT.BatchDeleteInstanceWithRelateResource(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBatchStartInstance(t *testing.T) {
	args := &api.BatchStartInstanceArgs{
		InstanceIds: []string{"i-gRYyYyjx", "i-GGc7Buqd"},
	}
	err := BCC_CLIENT.BatchStartInstance(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBatchStopInstance(t *testing.T) {
	args := &api.BatchStopInstanceArgs{
		ForceStop:        true,
		StopWithNoCharge: false,
		InstanceIds:      []string{"i-gRYyYyjx", "i-GGc7Buqd"},
	}
	err := BCC_CLIENT.BatchStopInstance(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListInstanceTypes(t *testing.T) {
	listArgs := &api.ListInstanceTypeArgs{
		ZoneName: "cn-bj-a",
	}
	res, err := BCC_CLIENT.ListInstanceTypes(listArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res)
}

func TestListIdMappings(t *testing.T) {
	listArgs := &api.ListIdMappingArgs{
		IdType:     "shot",
		ObjectType: "vm",
		InstanceIds: []string{
			"i-wQzV1qYZ",
			"i-b1jcrdt5",
		},
	}
	res, err := BCC_CLIENT.ListIdMappings(listArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res)
}

func TestBatchResizeInstance(t *testing.T) {
	listArgs := &api.BatchResizeInstanceArgs{
		Spec: "spec",
		InstanceIdList: []string{
			"i-wQzV1qYZ",
			"i-b1jcrdt5",
		},
	}
	res, err := BCC_CLIENT.BatchResizeInstance(listArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res)
}

func TestClient_DeleteSecurityGroupRule(t *testing.T) {
	args := &api.DeleteSecurityGroupRuleArgs{
		SecurityGroupRuleId: "r-zkcrsnesy13b",
	}
	err := BCC_CLIENT.DeleteSecurityGroupRule(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateSecurityGroupRule(t *testing.T) {
	remark := ""
	args := &api.UpdateSecurityGroupRuleArgs{
		SecurityGroupRuleId: "r-sdxzpzxe2igh",
		Remark:              &remark,
	}
	err := BCC_CLIENT.UpdateSecurityGroupRule(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestGetInstanceDeleteProgress(t *testing.T) {
	args := &api.GetInstanceDeleteProgressArgs{
		InstanceIds: []string{
			BCC_TestBccId,
		},
	}

	res, err := BCC_CLIENT.GetInstanceDeleteProgress(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(res)
}

func TestTagVolume(t *testing.T) {
	tagArgs := &api.TagVolumeArgs{
		ChangeTags: []api.Tag{
			{
				TagKey:   "go-SDK-Tag-Key3",
				TagValue: "go_SDK-Tag-Value2",
			},
		},
	}
	err := BCC_CLIENT.TagVolume(BCC_TestCdsId, tagArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUntagVolume(t *testing.T) {
	tagArgs := &api.TagVolumeArgs{
		ChangeTags: []api.Tag{
			{
				TagKey:   "go-SDK-Tag-Key3",
				TagValue: "go_SDK-Tag-Value2",
			},
		},
	}
	err := BCC_CLIENT.UntagVolume(BCC_TestCdsId, tagArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestTagSnapshotChain(t *testing.T) {
	tagArgs := &api.TagVolumeArgs{
		ChangeTags: []api.Tag{
			{
				TagKey:   "go-k",
				TagValue: "go-v",
			},
		},
	}
	err := BCC_CLIENT.TagSnapshotChain("sl-PdPu6Oel", tagArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUntagSnapshotChain(t *testing.T) {
	tagArgs := &api.TagVolumeArgs{
		ChangeTags: []api.Tag{
			{
				TagKey:   "go-k",
				TagValue: "go-v",
			},
		},
	}
	err := BCC_CLIENT.UntagSnapshotChain("sl-PdPu6Oel", tagArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListAvailableResizeSpecs(t *testing.T) {
	listAvailableResizeSpecsArgs := &api.ListAvailableResizeSpecsArgs{
		Spec: "bcc.ic5.c1m1",
		Zone: "cn-bj-a",
	}
	createResult, err := BCC_CLIENT.ListAvailableResizeSpecs(listAvailableResizeSpecsArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(createResult)
}

func TestBatchChangeInstanceToPrepay(t *testing.T) {
	batchChangeInstanceToPrepayArgs := &api.BatchChangeInstanceToPrepayArgs{
		Config: []api.PrepayConfig{
			{
				InstanceId: BCC_TestBccId,
				Duration:   1,
				CdsList: []string{
					BCC_TestCdsId,
				},
			},
		},
	}
	result, err := BCC_CLIENT.BatchChangeInstanceToPrepay(batchChangeInstanceToPrepayArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestBatchChangeInstanceToPostpay(t *testing.T) {
	batchChangeInstanceToPostArgs := &api.BatchChangeInstanceToPostpayArgs{
		Config: []api.PostpayConfig{
			{
				InstanceId: BCC_TestBccId,
				CdsList: []string{
					BCC_TestCdsId,
				},
			},
		},
	}
	result, err := BCC_CLIENT.BatchChangeInstanceToPostpay(batchChangeInstanceToPostArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestListInstanceRoles(t *testing.T) {
	result, err := BCC_CLIENT.ListInstanceRoles()
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestBindInstanceRole(t *testing.T) {
	bindInstanceRoleArgs := &api.BindInstanceRoleArgs{
		RoleName: "Test_BCC",
		Instances: []api.Instances{
			{
				InstanceId: BCC_TestBccId,
			},
		},
	}

	result, err := BCC_CLIENT.BindInstanceRole(bindInstanceRoleArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestUnBindInstanceRole(t *testing.T) {
	unbindInstanceRoleArgs := &api.UnBindInstanceRoleArgs{
		RoleName: "Test_BCC",
		Instances: []api.Instances{
			{
				InstanceId: BCC_TestBccId,
			},
		},
	}

	result, err := BCC_CLIENT.UnBindInstanceRole(unbindInstanceRoleArgs)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestDeleteIpv6(t *testing.T) {
	deleteIpv6Args := &api.DeleteIpv6Args{
		InstanceId: BCC_TestBccId,
		Reboot:     true,
	}

	err := BCC_CLIENT.DeleteIpv6(deleteIpv6Args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteIpv6WithIpv6Address(t *testing.T) {
	deleteIpv6Args := &api.DeleteIpv6Args{
		InstanceId:  "i-0nPl9WFJ",
		Ipv6Address: "2400:da00:e003:0:41c:4100:0:5",
		Reboot:      true,
	}

	err := BCC_CLIENT.DeleteIpv6(deleteIpv6Args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestAddIpv6(t *testing.T) {
	addIpv6Args := &api.AddIpv6Args{
		InstanceId:  BCC_TestBccId,
		Reboot:      true,
		Ipv6Address: "2400:da00:e003:0:41c:4100:0:2",
	}

	result, err := BCC_CLIENT.AddIpv6(addIpv6Args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestBindImageToTags(t *testing.T) {
	args := &api.BindTagsRequest{
		ChangeTags: []model.TagModel{
			{
				TagKey:   "TagKey",
				TagValue: "TagValue",
			},
		},
	}
	err := BCC_CLIENT.BindImageToTags(BCC_TestImageId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUnBindImageToTags(t *testing.T) {
	args := &api.UnBindTagsRequest{
		ChangeTags: []model.TagModel{
			{
				TagKey:   "TagKey",
				TagValue: "TagValue",
			},
		},
	}
	err := BCC_CLIENT.UnBindImageToTags(BCC_TestImageId, args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateRemoteCopySnapshot(t *testing.T) {
	args := &api.RemoteCopySnapshotArgs{
		ClientToken: "ClientTokenForTest",
		DestRegionInfos: []api.DestRegionInfo{
			{
				Name:       "Test",
				DestRegion: "bj",
			},
		},
	}
	result, err := BCC_CLIENT.CreateRemoteCopySnapshot("s-S9HdTie0", args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestImportCustomImage(t *testing.T) {
	args := &api.ImportCustomImageArgs{
		OsName:    "Centos",
		OsArch:    "32",
		OsType:    "linux",
		OsVersion: "6.5",
		Name:      "import_image_test",
		BosURL:    "http://cloud.baidu.com/testurl",
	}

	result, err := BCC_CLIENT.ImportCustomImage(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestGetAvailableImagesBySpec(t *testing.T) {
	args := &api.GetAvailableImagesBySpecArg{
		OsName:  "Centos",
		Spec:    "bcc.ic4.c1m1",
		MaxKeys: 1,
		Marker:  "m-21bmeYvH",
	}

	result, err := BCC_CLIENT.GetAvailableImagesBySpec(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestListCDSVolumeV3New(t *testing.T) {
	args := &api.ListCDSVolumeArgs{
		ChargeFilter: "postpay",
		Name:         "cdsTest",
		UsageFilter:  "data",
	}

	result, err := BCC_CLIENT.ListCDSVolumeV3(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestBatchRefundResource(t *testing.T) {
	arg := &api.BatchRefundResourceArg{
		InstanceIds: []string{
			"i-",
		},
		DeleteRelatedEnisFlag: true,
	}

	result, err := BCC_CLIENT.BatchRefundResource(arg)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestListCDSVolumeNew(t *testing.T) {
	args := &api.ListCDSVolumeArgs{
		ChargeFilter: "postpay",
		Name:         "test-ebcc-api-gc_datadiskvCSM",
		UsageFilter:  "data",
	}

	result, err := BCC_CLIENT.ListCDSVolume(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestGetAvailableStockWithSpec(t *testing.T) {
	rootOnLocal := true
	args := &api.GetAvailableStockWithSpecArgs{
		SpecList:     []string{"ehc.lgn5.c128m1024.8a100.8re.4d"},
		DeploySetIds: []string{"dset-Z3aEKdeY"},
		RootOnLocal:  &rootOnLocal,
		EhcClusterId: "ehc-bk4hM1N3",
	}
	result, err := BCC_CLIENT.GetAvailableStockWithSpec(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestModifyRelatedDeletePolicy(t *testing.T) {
	args := &api.RelatedDeletePolicy{
		IsEipAutoRelatedDelete: true,
	}
	err := BCC_CLIENT.ModifyRelatedDeletePolicy("i-ZMRzyU8f", args)
	ExpectEqual(t.Errorf, err, nil)
	instance, _ := BCC_CLIENT.GetInstanceDetail("i-ZMRzyU8f")
	ExpectEqual(t.Errorf, instance.Instance.IsEipAutoRelatedDelete, true)
	args = &api.RelatedDeletePolicy{
		IsEipAutoRelatedDelete: false,
	}
	_ = BCC_CLIENT.ModifyRelatedDeletePolicy("i-ZMRzyU8f", args)
	ExpectEqual(t.Errorf, err, nil)
	instance, _ = BCC_CLIENT.GetInstanceDetail("i-ZMRzyU8f")
	ExpectEqual(t.Errorf, instance.Instance.IsEipAutoRelatedDelete, false)
}

func TestModifyInstanceAttributeForJumboFrame(t *testing.T) {
	var openJumboFrame = new(bool)
	*openJumboFrame = true
	modifyArgs := &api.ModifyInstanceAttributeArgs{
		EnableJumboFrame: openJumboFrame,
	}
	err := BCC_CLIENT.ModifyInstanceAttribute(BCC_TestBccId, modifyArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestTransferReservedInstanceOrder(t *testing.T) {
	args := &api.TransferReservedInstanceRequest{
		ReservedInstanceIds: []string{
			"r-3p89YnJf",
		},
		RecipientAccountId: "",
	}
	result, err := BCC_CLIENT.TransferReservedInstanceOrder(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestRevokeTransferReservedInstanceOrder(t *testing.T) {
	args := &api.TransferReservedInstanceOperateRequest{
		TransferRecordIds: []string{
			"t-tgQYk4Rx",
		},
	}
	err := BCC_CLIENT.RevokeTransferReservedInstanceOrder(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRefuseTransferReservedInstanceOrder(t *testing.T) {
	args := &api.TransferReservedInstanceOperateRequest{
		TransferRecordIds: []string{
			"t-tgQYk4Rx",
		},
	}
	err := BCC_CLIENT.RefuseTransferReservedInstanceOrder(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestAcceptTransferReservedInstanceOrder(t *testing.T) {
	args := &api.AcceptTransferReservedInstanceRequest{
		TransferRecordId: "t-uNwDdZO9",
		EhcClusterId:     "ehc-bk4hM1N3",
	}
	err := BCC_CLIENT.AcceptTransferReservedInstanceOrder(args)
	ExpectEqual(t.Errorf, err, nil)
}

func TestTransferInReservedInstanceOrders(t *testing.T) {
	args := &api.DescribeTransferReservedInstancesRequest{
		ReservedInstanceIds: []string{
			// "r-I8rLAfcM",
		},
		TransferRecordIds: []string{
			// "t-FoM4l1xI",
		},
		Spec:   "bcc.g3.c1m1",
		Status: "timeout",
	}
	result, err := BCC_CLIENT.TransferInReservedInstanceOrders(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestTransferOutReservedInstanceOrders(t *testing.T) {
	args := &api.DescribeTransferReservedInstancesRequest{
		ReservedInstanceIds: []string{
			"r-ObFTPNIp",
		},
		TransferRecordIds: []string{
			"t-PKnSYeWh",
		},
		Spec:   "bcc.ic4.c2m2",
		Status: "fail",
	}
	result, err := BCC_CLIENT.TransferOutReservedInstanceOrders(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestCreateVolumeWithResGroup(t *testing.T) {
	args := &api.CreateCDSVolumeArgs{
		PurchaseCount: 1,
		CdsSizeInGB:   40,
		Billing: &api.Billing{
			PaymentTiming: api.PaymentTimingPostPaid,
		},
		ResGroupId: "RESG-4xiymzjDzqX",
	}

	result, _ := BCC_CLIENT.CreateCDSVolume(args)
	BCC_TestCdsId = result.VolumeIds[0]
	fmt.Print(BCC_TestCdsId)
	res, _ := BCC_CLIENT.GetCDSVolumeDetail(BCC_TestCdsId)
	fmt.Println(res.Volume)
}

func TestGetCDSPrice(t *testing.T) {
	args := &api.VolumePriceRequestArgs{
		PurchaseLength: 1,
		PaymentTiming:  "Prepaid",
		StorageType:    "cloud_hp1",
		CdsSizeInGB:    1000,
		PurchaseCount:  1,
		ZoneName:       "cn-bj-a",
	}

	result, err := BCC_CLIENT.getCdsPrice(args)
	ExpectEqual(t.Errorf, err, nil)
	fmt.Println(result)
}

func TestCreateEhcCluster(t *testing.T) {
	args := &api.CreateEhcClusterArg{
		Name:        "test-ehcCluster",
		ZoneName:    "cn-bj-a",
		Description: "test description",
	}
	result, _ := BCC_CLIENT.CreateEhcCluster(args)
	fmt.Println(result)
}

func TestEhcClusterList(t *testing.T) {
	args := &api.DescribeEhcClusterListArg{
		EhcClusterIdList: []string{
			"ehc-bk4hM1N3",
		}, NameList: []string{
			"test-modify",
		},
		ZoneName: "cn-bj-a",
		SortKey:  "name",
		SortDir:  "asc",
	}
	result, err := BCC_CLIENT.ListEhcCluster(args)
	fmt.Println(result)
	fmt.Println(err)
}

func TestCreateReservedInstance(t *testing.T) {

	args := &api.CreateReservedInstanceArgs{
		ClientToken:              "myClientToken",
		ReservedInstanceName:     "myReservedInstance",
		Scope:                    "AZ",
		ZoneName:                 "cn-bj-a",
		Spec:                     "bcc.g5.c2m8",
		OfferingType:             "FullyPrepay",
		InstanceCount:            1,
		ReservedInstanceCount:    1,
		ReservedInstanceTime:     1,
		ReservedInstanceTimeUnit: "month",
		AutoRenewTimeUnit:        "month",
		AutoRenewTime:            1,
		AutoRenew:                true,
		Tags: []api.Tag{
			{
				TagKey:   "Env",
				TagValue: "Production",
			},
		},
		TicketId: "ticket456",
	}

	result, err := BCC_CLIENT.CreateReservedInstance(args)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestModifyReservedInstances(t *testing.T) {
	args := &api.ModifyReservedInstancesArgs{
		ReservedInstances: []api.ModifyReservedInstance{
			{
				ReservedInstanceId:   "r-UBVQFB5b",
				ReservedInstanceName: "update-reserved-instance",
			},
		},
	}

	result, err := BCC_CLIENT.ModifyReservedInstances(args)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestEhcClusterModify(t *testing.T) {
	descriptions := ""
	args := &api.ModifyEhcClusterArg{
		EhcClusterId: "ehc-bk4hM1N3",
		Name:         "test-modify",
		Description:  &descriptions,
	}
	err := BCC_CLIENT.ModifyEhcCluster(args)
	fmt.Println(err)
}

func TestEhcClusterDelete(t *testing.T) {
	args := &api.DeleteEhcClusterArg{
		EhcClusterIdList: []string{
			"ehc-tBmphmZE",
		},
	}
	err := BCC_CLIENT.DeleteEhcCluster(args)
	fmt.Println(err)
}

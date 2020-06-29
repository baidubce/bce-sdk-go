package bbc

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	BBC_CLIENT              *Client
	BBC_TestBbcId           string
	BBC_TestImageId         string
	BBC_TestFlavorId        string
	BBC_TestRaidId          string
	BBC_TestZoneName        string
	BBC_TestSubnetId        string
	BBC_TestName            string
	BBC_TestAdminPass       string
	BBC_TestDeploySetId     string
	BBC_TestClientToken     string
	BBC_TestSecurityGroupId string
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 6; i++ {
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

	BBC_TestFlavorId = "flavor-id"
	BBC_TestImageId = "image-id"
	BBC_TestRaidId = "raid-id"
	BBC_TestZoneName = "zone-name"
	BBC_TestSubnetId = "subnet-id"
	BBC_TestName = "sdkTest"
	BBC_TestAdminPass = "123@adminPass"
	BBC_TestDeploySetId = "deployset-id"
	BBC_TestBbcId = "bbc_id"
	BBC_TestSecurityGroupId = "bbc-security-group-id"
	BBC_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.WARN)
	//log.SetLogLevel(log.DEBUG)
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
	createInstanceArgs := &CreateInstanceArgs{
		FlavorId:         BBC_TestFlavorId,
		ImageId:          BBC_TestImageId,
		RaidId:           BBC_TestRaidId,
		RootDiskSizeInGb: 40,
		PurchaseCount:    1,
		ZoneName:         BBC_TestZoneName,
		SubnetId:         BBC_TestSubnetId,
		ClientToken:      BBC_TestClientToken,
		Billing: Billing{
			PaymentTiming: PaymentTimingPostPaid,
		},
		DeploySetId: BBC_TestDeploySetId,
		Name:        BBC_TestName,
	}
	res, err := BBC_CLIENT.CreateInstance(createInstanceArgs)
	fmt.Println(res)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListInstances(t *testing.T) {
	listArgs := &ListInstancesArgs{
		MaxKeys: 100,
	}
	res, err := BBC_CLIENT.ListInstances(listArgs)
	fmt.Println(res)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetInstanceDetail(t *testing.T) {
	res, err := BBC_CLIENT.GetInstanceDetail(BBC_TestBbcId)
	fmt.Println(res)
	ExpectEqual(t.Errorf, err, nil)
}

func TestStopInstance(t *testing.T) {
	err := BBC_CLIENT.StopInstance(BBC_TestBbcId, false)
	ExpectEqual(t.Errorf, err, nil)
}

func TestStartInstance(t *testing.T) {
	err := BBC_CLIENT.StartInstance(BBC_TestBbcId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRebootInstance(t *testing.T) {
	err := BBC_CLIENT.RebootInstance(BBC_TestBbcId, true)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyInstanceName(t *testing.T) {
	modifyInstanceNameArgs := &ModifyInstanceNameArgs{
		Name: "new_bbc_name",
	}
	err := BBC_CLIENT.ModifyInstanceName(BBC_TestBbcId, modifyInstanceNameArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyInstanceDesc(t *testing.T) {
	modifyInstanceDescArgs := &ModifyInstanceDescArgs{
		Description: "new_bbc_description",
	}
	err := BBC_CLIENT.ModifyInstanceDesc(BBC_TestBbcId, modifyInstanceDescArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestRebuildInstance(t *testing.T) {
	rebuildArgs := &RebuildInstanceArgs{
		ImageId:        BBC_TestImageId,
		AdminPass:      BBC_TestAdminPass,
		IsPreserveData: true,
		RaidId:         BBC_TestRaidId,
		SysRootSize:    20,
	}
	err := BBC_CLIENT.RebuildInstance(BBC_TestBbcId, true, rebuildArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestReleaseInstance(t *testing.T) {
	err := BBC_CLIENT.DeleteInstance(BBC_TestBbcId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestModifyInstancePassword(t *testing.T) {
	modifyInstancePasswordArgs := &ModifyInstancePasswordArgs{
		AdminPass: BBC_TestAdminPass,
	}
	err := BBC_CLIENT.ModifyInstancePassword(BBC_TestBbcId, modifyInstancePasswordArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetVpcSubnet(t *testing.T) {
	getVpcSubnetArgs := &GetVpcSubnetArgs{
		BbcIds: []string{BBC_TestBbcId},
	}
	result, err := BBC_CLIENT.GetVpcSubnet(getVpcSubnetArgs)
	fmt.Println(result)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBatchAddIp(t *testing.T) {
	privateIps := []string{"192.168.1.25"}
	batchAddIpArgs := &BatchAddIpArgs{
		InstanceId: BBC_TestBbcId,
		PrivateIps: privateIps,
	}
	err := BBC_CLIENT.BatchAddIP(batchAddIpArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestBatchDelIp(t *testing.T) {
	privateIps := []string{"192.168.1.25"}
	batchDelIpArgs := &BatchDelIpArgs{
		InstanceId: BBC_TestBbcId,
		PrivateIps: privateIps,
	}
	err := BBC_CLIENT.BatchDelIP(batchDelIpArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestUnbindTags(t *testing.T) {
	unbindTagsArgs := &UnbindTagsArgs{
		ChangeTags: []model.TagModel{
			{
				TagKey:   "BCC",
				TagValue: "aaa",
			},
		},
	}
	err := BBC_CLIENT.UnbindTags(BBC_TestBbcId, unbindTagsArgs)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListFlavors(t *testing.T) {
	res, err := BBC_CLIENT.ListFlavors()
	fmt.Println(res)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetFlavorDetail(t *testing.T) {
	testFlavorId := BBC_TestFlavorId
	rep, err := BBC_CLIENT.GetFlavorDetail(testFlavorId)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetFlavorRaid(t *testing.T) {
	testFlavorId := BBC_TestFlavorId
	rep, err := BBC_CLIENT.GetFlavorRaid(testFlavorId)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateImageFromInstanceId(t *testing.T) {
	testInstanceId := BBC_TestBbcId
	testImageName := "testCreateImage"
	queryArgs := &CreateImageArgs{
		ImageName:  testImageName,
		InstanceId: testInstanceId,
	}
	rep, err := BBC_CLIENT.CreateImageFromInstanceId(queryArgs)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListImage(t *testing.T) {
	queryArgs := &ListImageArgs{}
	rep, err := BBC_CLIENT.ListImage(queryArgs)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetImageDetail(t *testing.T) {
	testImageId := BBC_TestImageId
	rep, err := BBC_CLIENT.GetImageDetail(testImageId)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteImage(t *testing.T) {
	testImageId := BBC_TestImageId
	err := BBC_CLIENT.DeleteImage(testImageId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetOperationLog(t *testing.T) {
	queryArgs := &GetOperationLogArgs{}
	rep, err := BBC_CLIENT.GetOperationLog(queryArgs)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCreateDeploySet(t *testing.T) {
	testDeploySetName := "testName"
	testDeployDesc := "testDesc"
	testConcurrency := 1
	testStrategy := "tor_ha"
	queryArgs := &CreateDeploySetArgs{
		Strategy:    testStrategy,
		Concurrency: testConcurrency,
		Name:        testDeploySetName,
		Desc:        testDeployDesc,
	}
	rep, err := BBC_CLIENT.CreateDeploySet(queryArgs)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestListDeploySets(t *testing.T) {
	rep, err := BBC_CLIENT.ListDeploySets()
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestGetDeploySet(t *testing.T) {
	testDeploySetID := BBC_TestDeploySetId
	rep, err := BBC_CLIENT.GetDeploySet(testDeploySetID)
	fmt.Println(rep)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteDeploySet(t *testing.T) {
	testDeleteDeploySetId := BBC_TestDeploySetId
	err := BBC_CLIENT.DeleteDeploySet(testDeleteDeploySetId)
	fmt.Println(err)
	ExpectEqual(t.Errorf, err, nil)
}

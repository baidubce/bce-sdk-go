package cce

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	CCE_CLIENT             *Client
	CCE_CLUSTER_ID         string
	CCE_SHIFT_IN_NODE_ID   string
	CCE_SHIFT_OUT_NODE_ID  string
	CCE_CONTAINER_NET      string
	CCE_KUBERNETES_VERSION string

	CCE_NODE_ADMIN_PASSWD = "123qwe!"
	VPC_TEST_ID           = ""
	SUBNET_TEST_ID        = ""
	SECURITY_GROUP_ID     = ""
	ZONE_TEST_ID          = "zoneA"
	IMAGE_TEST_ID         = "m-Nlv9C0tF"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 7; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fmt.Println(conf)
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	CCE_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
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

func TestClient_ListVersion(t *testing.T) {
	result, err := CCE_CLIENT.ListVersions()
	ExpectEqual(t.Errorf, nil, err)

	CCE_KUBERNETES_VERSION = result.Data[0]
}

func TestClient_GetContainerNet(t *testing.T) {
	args := &GetContainerNetArgs{
		VpcShortId: VPC_TEST_ID,
		VpcCidr:    "192.168.0.0/24",
		Size:       1000,
	}
	result, err := CCE_CLIENT.GetContainerNet(args)
	ExpectEqual(t.Errorf, nil, err)

	CCE_CONTAINER_NET = result.ContainerNet
}

func TestClient_CreateCluster(t *testing.T) {
	args := &CreateClusterArgs{
		ClusterName:       "sdk-test",
		Version:           CCE_KUBERNETES_VERSION,
		MainAvailableZone: "zoneA",
		ContainerNet:      CCE_CONTAINER_NET,
		DeployMode:        DeployModeBcc,
		OrderContent: &BaseCreateOrderRequestVo{Items: []Item{
			{
				Config: BccConfig{
					ProductType:     ProductTypePostpay,
					InstanceType:    InstanceTypeG3,
					Cpu:             1,
					Memory:          2,
					ImageType:       ImageTypeCommon,
					SubnetUuid:      SUBNET_TEST_ID,
					SecurityGroupId: SECURITY_GROUP_ID,
					AdminPass:       CCE_NODE_ADMIN_PASSWD,
					PurchaseNum:     1,
					ImageId:         IMAGE_TEST_ID,
					ServiceType:     ServiceTypeBCC,
				},
			},
		}},
	}
	result, err := CCE_CLIENT.CreateCluster(args)
	ExpectEqual(t.Errorf, nil, err)

	CCE_CLUSTER_ID = result.ClusterUuid
}

func TestClient_GetCluster(t *testing.T) {
	_, err := CCE_CLIENT.GetCluster(CCE_CLUSTER_ID)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListClusters(t *testing.T) {
	args := &ListClusterArgs{}
	_, err := CCE_CLIENT.ListClusters(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListNodes(t *testing.T) {
	args := &ListNodeArgs{
		ClusterUuid: CCE_CLUSTER_ID,
	}
	result, err := CCE_CLIENT.ListNodes(args)
	ExpectEqual(t.Errorf, nil, err)

	CCE_SHIFT_OUT_NODE_ID = result.Nodes[0].InstanceShortId
}

func TestClient_GetKubeConfig(t *testing.T) {
	args := &GetKubeConfigArgs{
		ClusterUuid: CCE_CLUSTER_ID,
	}
	_, err := CCE_CLIENT.GetKubeConfig(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ScalingUp(t *testing.T) {
	args := &ScalingUpArgs{
		ClusterUuid: CCE_CLUSTER_ID,
		OrderContent: &BaseCreateOrderRequestVo{Items: []Item{
			{
				Config: BccConfig{
					ProductType:     ProductTypePostpay,
					InstanceType:    InstanceTypeG3,
					Cpu:             1,
					Memory:          2,
					ImageType:       ImageTypeCommon,
					SubnetUuid:      SUBNET_TEST_ID,
					SecurityGroupId: SECURITY_GROUP_ID,
					AdminPass:       CCE_NODE_ADMIN_PASSWD,
					PurchaseNum:     1,
					ImageId:         IMAGE_TEST_ID,
					ServiceType:     ServiceTypeBCC,
				},
			},
		}},
	}
	_, err := CCE_CLIENT.ScalingUp(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ShiftOutNode(t *testing.T) {
	args := &ShiftOutNodeArgs{
		ClusterUuid: CCE_CLUSTER_ID,
		NodeInfoList: []CceNodeInfo{
			{InstanceId: CCE_SHIFT_OUT_NODE_ID},
		},
	}
	err := CCE_CLIENT.ShiftOutNode(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListExistedBccNode(t *testing.T) {
	args := &ListExistedNodeArgs{
		ClusterUuid: CCE_CLUSTER_ID,
	}
	result, err := CCE_CLIENT.ListExistedBccNode(args)
	ExpectEqual(t.Errorf, nil, err)

	CCE_SHIFT_IN_NODE_ID = result.NodeList[0].InstanceId
}

func TestClient_ShiftInNode(t *testing.T) {
	args := &ShiftInNodeArgs{
		ClusterUuid:  CCE_CLUSTER_ID,
		NeedRebuild:  false,
		AdminPass:    CCE_NODE_ADMIN_PASSWD,
		InstanceType: ShiftInstanceTypeBcc,
		NodeInfoList: []CceNodeInfo{
			{
				InstanceId: CCE_SHIFT_IN_NODE_ID,
				AdminPass:  CCE_NODE_ADMIN_PASSWD,
			},
		},
	}
	err := CCE_CLIENT.ShiftInNode(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ScalingDown(t *testing.T) {
	args := &ScalingDownArgs{
		ClusterUuid:  CCE_CLUSTER_ID,
		DeleteEipCds: true,
		DeleteSnap:   true,
		NodeInfo: []NodeInfo{
			{
				InstanceId: CCE_SHIFT_IN_NODE_ID,
			},
		},
	}
	err := CCE_CLIENT.ScalingDown(args)
	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteCluster(t *testing.T) {
	deleteClusterArgs := &DeleteClusterArgs{
		ClusterUuid:  CCE_CLUSTER_ID,
		DeleteEipCds: true,
		DeleteSnap:   true,
	}
	err := CCE_CLIENT.DeleteCluster(deleteClusterArgs)
	ExpectEqual(t.Errorf, nil, err)
}

package v2

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	bccapi "github.com/baidubce/bce-sdk-go/services/bcc/api"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/types"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	CCE_CLIENT             *Client
	CCE_CLUSTER_ID         string
	CCE_INSTANCE_ID        string
	VPC_TEST_ID            = "vpc-mwbgygrjb72w"
	IMAGE_TEST_ID          = "m-gTpZ1k6n"
	SECURITY_GROUP_TESI_ID = "g-xh04bcdkq5n6"
	ADMIN_PASSWORD_TEST    = "test123!YT"
	SSH_KEY_TEST_ID        = "k-3uvrdvVq"
	VPC_SUBNET_TEST_ID     = "sbn-mnbvhnuupv1u"
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
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

//获取ak sk与endpoint并构建ccev2Client
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

	CCE_CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
}

func TestClient_CheckClusterIPCIDR(t *testing.T) {
	args := &CheckClusterIPCIDRArgs{
		VPCID:         VPC_TEST_ID,
		VPCCIDR:       "192.168.0.0/16",
		ClusterIPCIDR: "172.31.0.0/16",
		IPVersion:     "ipv4",
	}
	resp, err := CCE_CLIENT.CheckClusterIPCIDR(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_CheckContainerNetworkCIDR(t *testing.T) {
	args := &CheckContainerNetworkCIDRArgs{
		VPCID:          VPC_TEST_ID,
		VPCCIDR:        "192.168.0.0/16",
		ContainerCIDR:  "172.28.0.0/16",
		ClusterIPCIDR:  "172.31.0.0/16",
		MaxPodsPerNode: 256,
	}
	resp, err := CCE_CLIENT.CheckContainerNetworkCIDR(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_RecommendClusterIPCIDR(t *testing.T) {
	args := &RecommendClusterIPCIDRArgs{
		ClusterMaxServiceNum: 2,
		ContainerCIDR:        "172.28.0.0/16",
		IPVersion:            "ipv4",
		PrivateNetCIDRs:      []PrivateNetString{PrivateIPv4Net172},
		VPCCIDR:              "192.168.0.0/16",
	}
	resp, err := CCE_CLIENT.RecommendClusterIPCIDR(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_RecommendContainerCIDR(t *testing.T) {
	args := &RecommendContainerCIDRArgs{
		ClusterMaxNodeNum: 2,
		IPVersion:         "ipv4",
		K8SVersion:        types.K8S_1_16_8,
		MaxPodsPerNode:    32,
		PrivateNetCIDRs:   []PrivateNetString{PrivateIPv4Net172},
		VPCCIDR:           "192.168.0.0/16",
		VPCID:             VPC_TEST_ID,
	}

	resp, err := CCE_CLIENT.RecommendContainerCIDR(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Request ID:" + resp.RequestID)
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_CreateCluster(t *testing.T) {
	args := &CreateClusterArgs{
		CreateClusterRequest: &CreateClusterRequest{
			ClusterSpec: &types.ClusterSpec{
				ClusterName: "sdk-ccev2-test",
				K8SVersion:  types.K8S_1_16_8,
				RuntimeType: types.RuntimeTypeDocker,
				VPCID:       VPC_TEST_ID,
				MasterConfig: types.MasterConfig{
					MasterType:            types.MasterTypeManaged,
					ClusterHA:             1,
					ExposedPublic:         false,
					ClusterBLBVPCSubnetID: VPC_SUBNET_TEST_ID,
					ManagedClusterMasterOption: types.ManagedClusterMasterOption{
						MasterVPCSubnetZone: types.AvailableZoneA,
					},
				},
				ContainerNetworkConfig: types.ContainerNetworkConfig{
					Mode:                 types.ContainerNetworkModeKubenet,
					LBServiceVPCSubnetID: VPC_SUBNET_TEST_ID,
					ClusterPodCIDR:       "172.28.0.0/16",
					ClusterIPServiceCIDR: "172.31.0.0/16",
				},
			},
			NodeSpecs: []*InstanceSet{
				{
					Count: 1,
					InstanceSpec: types.InstanceSpec{
						InstanceName: "",
						ClusterRole:  types.ClusterRoleNode,
						Existed:      false,
						MachineType:  types.MachineTypeBCC,
						InstanceType: bccapi.InstanceTypeN3,
						VPCConfig: types.VPCConfig{
							VPCID:           VPC_TEST_ID,
							VPCSubnetID:     VPC_SUBNET_TEST_ID,
							SecurityGroupID: SECURITY_GROUP_TESI_ID,
							AvailableZone:   types.AvailableZoneA,
						},
						InstanceResource: types.InstanceResource{
							CPU:           4,
							MEM:           8,
							RootDiskSize:  40,
							LocalDiskSize: 0,
							CDSList:       []types.CDSConfig{},
						},
						ImageID: IMAGE_TEST_ID,
						InstanceOS: types.InstanceOS{
							ImageType: bccapi.ImageTypeSystem,
						},
						NeedEIP:              false,
						AdminPassword:        ADMIN_PASSWORD_TEST,
						SSHKeyID:             "",
						InstanceChargingType: bccapi.PaymentTimingPostPaid,
						RuntimeType:          types.RuntimeTypeDocker,
					},
				},
			},
		},
	}
	resp, err := CCE_CLIENT.CreateCluster(args)

	ExpectEqual(t.Errorf, nil, err)
	CCE_CLUSTER_ID = resp.ClusterID

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

	//等集群创建完成
	time.Sleep(time.Duration(240) * time.Second)
}

func TestClient_GetCluster(t *testing.T) {
	resp, err := CCE_CLIENT.GetCluster(CCE_CLUSTER_ID)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_ListClusters(t *testing.T) {
	args := &ListClustersArgs{
		KeywordType: "clusterName",
		Keyword:     "",
		OrderBy:     "clusterID",
		Order:       OrderASC,
		PageSize:    10,
		PageNum:     1,
	}
	resp, err := CCE_CLIENT.ListClusters(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_CreateInstances(t *testing.T) {
	args := &CreateInstancesArgs{
		ClusterID: CCE_CLUSTER_ID,
		Instances: []*InstanceSet{
			{
				Count: 1,
				InstanceSpec: types.InstanceSpec{
					ClusterRole:  types.ClusterRoleNode,
					Existed:      false,
					MachineType:  types.MachineTypeBCC,
					InstanceType: bccapi.InstanceTypeN3,
					VPCConfig: types.VPCConfig{
						VPCID:           VPC_TEST_ID,
						VPCSubnetID:     VPC_SUBNET_TEST_ID,
						SecurityGroupID: SECURITY_GROUP_TESI_ID,
						AvailableZone:   types.AvailableZoneA,
					},
					InstanceResource: types.InstanceResource{
						CPU:           1,
						MEM:           4,
						RootDiskSize:  40,
						LocalDiskSize: 0,
					},
					ImageID: IMAGE_TEST_ID,
					InstanceOS: types.InstanceOS{
						ImageType: bccapi.ImageTypeSystem,
					},
					NeedEIP:              false,
					AdminPassword:        ADMIN_PASSWORD_TEST,
					SSHKeyID:             SSH_KEY_TEST_ID,
					InstanceChargingType: bccapi.PaymentTimingPostPaid,
					RuntimeType:          types.RuntimeTypeDocker,
				},
			},
		},
	}
	resp, err := CCE_CLIENT.CreateInstances(args)

	ExpectEqual(t.Errorf, nil, err)
	CCE_INSTANCE_ID = resp.CCEInstanceIDs[0]

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_ListInstancesByPage(t *testing.T) {
	args := &ListInstancesByPageArgs{
		ClusterID: CCE_CLUSTER_ID,
		Params: &ListInstancesByPageParams{
			KeywordType: InstanceKeywordTypeInstanceName,
			Keyword:     "",
			OrderBy:     "createdAt",
			Order:       OrderASC,
			PageNo:      1,
			PageSize:    10,
		},
	}
	resp, err := CCE_CLIENT.ListInstancesByPage(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_GetInstance(t *testing.T) {
	args := &GetInstanceArgs{
		ClusterID:  CCE_CLUSTER_ID,
		InstanceID: CCE_INSTANCE_ID,
	}
	resp, err := CCE_CLIENT.GetInstance(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_GetClusterQuota(t *testing.T) {
	resp, err := CCE_CLIENT.GetClusterQuota()

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_GetClusterNodeQuota(t *testing.T) {
	resp, err := CCE_CLIENT.GetClusterNodeQuota(CCE_CLUSTER_ID)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_DeleteInstances(t *testing.T) {
	args := &DeleteInstancesArgs{
		ClusterID: CCE_CLUSTER_ID,
		DeleteInstancesRequest: &DeleteInstancesRequest{
			InstanceIDs: []string{CCE_INSTANCE_ID},
			DeleteOption: &types.DeleteOption{
				MoveOut:           false,
				DeleteCDSSnapshot: true,
				DeleteResource:    true,
			},
		},
	}
	resp, err := CCE_CLIENT.DeleteInstances(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_DeleteCluster(t *testing.T) {
	args := &DeleteClusterArgs{
		ClusterID:         CCE_CLUSTER_ID,
		DeleteResource:    true,
		DeleteCDSSnapshot: true,
	}
	resp, err := CCE_CLIENT.DeleteCluster(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Request ID:" + resp.RequestID)
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

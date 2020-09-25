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
	CCE_INSTANCE_GROUP_ID  string
	CCE_INSTANCE_ID        string
	VPC_TEST_ID            = ""
	IMAGE_TEST_ID          = ""
	SECURITY_GROUP_TEST_ID = ""
	VPC_SUBNET_TEST_ID     = ""
)

// For security reason, ak/sk should not hard write here.
type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
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

	log.SetLogLevel(log.WARN)

	CCE_CLIENT, err = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Setup Complete")
}

//Try to clean environment
func teardown() {
	if CCE_INSTANCE_ID != "" && CCE_CLIENT != nil {
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
		_, err := CCE_CLIENT.DeleteInstances(args)
		if err != nil {
			log.Error(err.Error())
		}
	}

	if CCE_INSTANCE_GROUP_ID != "" && CCE_CLIENT != nil {
		args := &DeleteInstanceGroupArgs{
			ClusterID:       CCE_CLUSTER_ID,
			InstanceGroupID: CCE_INSTANCE_GROUP_ID,
			DeleteInstances: true,
		}
		_, err := CCE_CLIENT.DeleteInstanceGroup(args)
		if err != nil {
			log.Error(err.Error())
		}
	}

	if CCE_CLUSTER_ID != "" && CCE_CLIENT != nil {
		args := &DeleteClusterArgs{
			ClusterID:         CCE_CLUSTER_ID,
			DeleteResource:    true,
			DeleteCDSSnapshot: true,
		}
		_, err := CCE_CLIENT.DeleteCluster(args)
		if err != nil {
			log.Error(err.Error())
		}
	}
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
		ClusterMaxServiceNum: 8,
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
				K8SCustomConfig: types.K8SCustomConfig{
					KubeAPIQPS:   1000,
					KubeAPIBurst: 2000,
				},
			},
		},
	}
	resp, err := CCE_CLIENT.CreateCluster(args)

	ExpectEqual(t.Errorf, nil, err)
	if resp.ClusterID == "" {
		t.Fatalf("Request Fail. Cluster ID is empty.")
	}

	CCE_CLUSTER_ID = resp.ClusterID

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

	//等集群创建完成
	time.Sleep(time.Duration(180) * time.Second)
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

func TestClient_CreateInstanceGroup(t *testing.T) {
	args := &CreateInstanceGroupArgs{
		ClusterID: CCE_CLUSTER_ID,
		Request: &CreateInstanceGroupRequest{
			types.InstanceGroupSpec{
				InstanceGroupName: "jichao-sdk-testcase",
				CleanPolicy:       types.DeleteCleanPolicy,
				Replicas:          3,
				InstanceTemplate: types.InstanceTemplate{
					InstanceSpec: types.InstanceSpec{
						ClusterRole:  types.ClusterRoleNode,
						Existed:      false,
						MachineType:  types.MachineTypeBCC,
						InstanceType: bccapi.InstanceTypeN3,
						VPCConfig: types.VPCConfig{
							VPCID:           VPC_TEST_ID,
							VPCSubnetID:     VPC_SUBNET_TEST_ID,
							SecurityGroupID: SECURITY_GROUP_TEST_ID,
							AvailableZone:   types.AvailableZoneA,
						},
						DeployCustomConfig: types.DeployCustomConfig{
							PreUserScript:  "ls",
							PostUserScript: "ls",
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
						InstanceChargingType: bccapi.PaymentTimingPostPaid,
						RuntimeType:          types.RuntimeTypeDocker,
					},
				},
			},
		},
	}

	resp, err := CCE_CLIENT.CreateInstanceGroup(args)

	ExpectEqual(t.Errorf, nil, err)
	if resp.InstanceGroupID == "" {
		t.Fatalf("Request Fail. Instance Group ID is empty.")
	}
	CCE_INSTANCE_GROUP_ID = resp.InstanceGroupID

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

	time.Sleep(time.Duration(180) * time.Second)
}

func TestClient_ListInstanceGroups(t *testing.T) {
	args := &ListInstanceGroupsArgs{
		ClusterID: CCE_CLUSTER_ID,
		ListOption: &InstanceGroupListOption{
			PageNo:   0,
			PageSize: 0,
		},
	}
	resp, err := CCE_CLIENT.ListInstanceGroups(args)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_ListInstancesByInstanceGroupID(t *testing.T) {
	args := &ListInstanceByInstanceGroupIDArgs{
		ClusterID:       CCE_CLUSTER_ID,
		InstanceGroupID: CCE_INSTANCE_GROUP_ID,
		PageSize:        0,
		PageNo:          0,
	}

	resp, err := CCE_CLIENT.ListInstancesByInstanceGroupID(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_GetInstanceGroup(t *testing.T) {
	args := &GetInstanceGroupArgs{
		ClusterID:       CCE_CLUSTER_ID,
		InstanceGroupID: CCE_INSTANCE_GROUP_ID,
	}

	resp, err := CCE_CLIENT.GetInstanceGroup(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_UpdateInstanceGroupReplicas(t *testing.T) {
	args := &UpdateInstanceGroupReplicasArgs{
		ClusterID:       CCE_CLUSTER_ID,
		InstanceGroupID: CCE_INSTANCE_GROUP_ID,
		Request: &UpdateInstanceGroupReplicasRequest{
			Replicas:       1,
			DeleteInstance: true,
			DeleteOption: &types.DeleteOption{
				MoveOut:           false,
				DeleteCDSSnapshot: true,
				DeleteResource:    true,
			},
		},
	}

	resp, err := CCE_CLIENT.UpdateInstanceGroupReplicas(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

	time.Sleep(time.Duration(120) * time.Second)
}

func TestClient_CreateAutoscaler(t *testing.T) {
	args := &CreateAutoscalerArgs{
		ClusterID: CCE_CLUSTER_ID,
	}

	resp, err := CCE_CLIENT.CreateAutoscaler(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_GetAutoscaler(t *testing.T) {
	args := &GetAutoscalerArgs{
		ClusterID: CCE_CLUSTER_ID,
	}

	resp, err := CCE_CLIENT.GetAutoscaler(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_UpdateAutoscaler(t *testing.T) {
	args := &UpdateAutoscalerArgs{
		ClusterID: CCE_CLUSTER_ID,
		AutoscalerConfig: ClusterAutoscalerConfig{
			ReplicaCount:     5,
			ScaleDownEnabled: true,
			Expander:         "random",
		},
	}

	resp, err := CCE_CLIENT.UpdateAutoscaler(args)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_UpdateInstanceGroupClusterAutoscalerSpec(t *testing.T) {
	args := &UpdateInstanceGroupClusterAutoscalerSpecArgs{
		ClusterID:       CCE_CLUSTER_ID,
		InstanceGroupID: CCE_INSTANCE_GROUP_ID,
		Request: &ClusterAutoscalerSpec{
			Enabled:              true,
			MinReplicas:          2,
			MaxReplicas:          5,
			ScalingGroupPriority: 1,
		},
	}

	resp, err := CCE_CLIENT.UpdateInstanceGroupClusterAutoscalerSpec(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func TestClient_GetKubeConfig(t *testing.T) {
	args := &GetKubeConfigArgs{
		ClusterID:      CCE_CLUSTER_ID,
		KubeConfigType: KubeConfigTypeVPC,
	}

	resp, err := CCE_CLIENT.GetKubeConfig(args)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

	ExpectEqual(t.Errorf, nil, err)
}

func TestClient_DeleteInstanceGroup(t *testing.T) {
	args := &DeleteInstanceGroupArgs{
		ClusterID:       CCE_CLUSTER_ID,
		InstanceGroupID: CCE_INSTANCE_GROUP_ID,
		DeleteInstances: true,
	}

	resp, err := CCE_CLIENT.DeleteInstanceGroup(args)

	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

	time.Sleep(time.Duration(180) * time.Second)
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
						SecurityGroupID: SECURITY_GROUP_TEST_ID,
						AvailableZone:   types.AvailableZoneA,
					},
					DeployCustomConfig: types.DeployCustomConfig{
						PreUserScript:  "ls",
						PostUserScript: "time",
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
						OSType: types.OSTypeLinux,
						OSName: types.OSNameCentOS,
						OSVersion: "7.5",
						OSArch: "x86_64 (64bit)",
					},
					NeedEIP:              false,
					InstanceChargingType: bccapi.PaymentTimingPostPaid,
					RuntimeType:          types.RuntimeTypeDocker,
				},
			},
		},
	}
	resp, err := CCE_CLIENT.CreateInstances(args)

	ExpectEqual(t.Errorf, nil, err)
	if resp.CCEInstanceIDs == nil || len(resp.CCEInstanceIDs) == 0 {
		t.Fatalf("Request Fail. Instance ID is empty.")
	}
	CCE_INSTANCE_ID = resp.CCEInstanceIDs[0]

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

	time.Sleep(time.Duration(180) * time.Second)
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

func TestClient_UpdateInstance(t *testing.T) {
	args := &GetInstanceArgs{
		ClusterID:  CCE_CLUSTER_ID,
		InstanceID: CCE_INSTANCE_ID,
	}
	respGet, err := CCE_CLIENT.GetInstance(args)

	oldInstanceSpec := respGet.Instance.Spec

	oldInstanceSpec.CCEInstancePriority = 1

	argsUpdate := &UpdateInstanceArgs{
		ClusterID:  CCE_CLUSTER_ID,
		InstanceID: CCE_INSTANCE_ID,
		InstanceSpec: oldInstanceSpec,
	}

	respUpdate, err := CCE_CLIENT.UpdateInstance(argsUpdate)
	ExpectEqual(t.Errorf, nil, err)

	s, _ := json.MarshalIndent(respUpdate, "", "\t")
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

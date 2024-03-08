package cluster

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bcc/api"
	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/types"
)

func CreateCluster() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	ccev2Client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}
	args := &v2.CreateClusterArgs{
		CreateClusterRequest: &v2.CreateClusterRequest{
			ClusterSpec: &types.ClusterSpec{
				ClusterName: "your-cluster-name",
				K8SVersion:  "1.21.14",
				RuntimeType: types.RuntimeTypeDocker,
				VPCID:       "vpc-id",
				MasterConfig: types.MasterConfig{
					MasterType:            types.MasterTypeManaged,
					ClusterHA:             1,
					ExposedPublic:         false,
					ClusterBLBVPCSubnetID: "cluster-blb-vpc-subnet-id",
					ManagedClusterMasterOption: types.ManagedClusterMasterOption{
						MasterVPCSubnetZone: types.AvailableZoneA,
					},
				},
				ContainerNetworkConfig: types.ContainerNetworkConfig{
					Mode:                 types.ContainerNetworkModeKubenet,
					LBServiceVPCSubnetID: "lb-service-vpc-subnet-id",
					ClusterPodCIDR:       "172.28.0.0/16",
					ClusterIPServiceCIDR: "172.31.0.0/16",
				},
			},
			NodeSpecs: []*v2.InstanceSet{
				{
					Count: 1,
					InstanceSpec: types.InstanceSpec{
						InstanceName: "",
						ClusterRole:  types.ClusterRoleNode,
						Existed:      false,
						MachineType:  types.MachineTypeBCC,
						InstanceType: api.InstanceTypeN3,
						VPCConfig: types.VPCConfig{
							VPCID:           "",
							VPCSubnetID:     "vpc-subnet-id",
							SecurityGroupID: "security-group-id",
							AvailableZone:   types.AvailableZoneA,
						},
						InstanceResource: types.InstanceResource{
							CPU:           4,
							MEM:           8,
							RootDiskSize:  40,
							LocalDiskSize: 0,
							CDSList:       []types.CDSConfig{},
						},
						ImageID: "image-id",
						InstanceOS: types.InstanceOS{
							ImageType: api.ImageTypeSystem,
						},
						NeedEIP:              false,
						AdminPassword:        "admin-password",
						SSHKeyID:             "ssh-key-id",
						InstanceChargingType: api.PaymentTimingPostPaid,
						RuntimeType:          types.RuntimeTypeDocker,
					},
				},
			},
		},
	}

	resp, err := ccev2Client.CreateCluster(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

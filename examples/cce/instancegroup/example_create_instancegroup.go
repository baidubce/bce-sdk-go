package instancegroup

import (
	"encoding/json"
	"fmt"

	ccev2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/types"
)

func CreateHPASInstanceGroup() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	ccev2Client, err := ccev2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	// 参数配置
	clusterID := ""
	instanceGroupName := ""
	securityGroupType := types.SecurityGroupTypeEnterprise
	securityGroupID := ""
	securityGroupName := ""
	machineType := types.MachineTypeHPAS
	instanceType := types.InstanceTypeHPAS
	subnetID := ""
	hpasAppType := ""
	hpasAppPerformanceLevel := ""
	imageID := ""
	adminPassword := ""

	args := &ccev2.CreateInstanceGroupArgs{
		ClusterID: clusterID,
		Request: &ccev2.CreateInstanceGroupRequest{
			InstanceGroupSpec: types.InstanceGroupSpec{
				Replicas:          0,
				InstanceGroupName: instanceGroupName,
				SecurityGroups: []types.SecurityGroupV2{
					{
						Type: securityGroupType,
						ID:   securityGroupID,
						Name: securityGroupName,
					},
				},
				InstanceTemplates: []types.InstanceTemplate{
					{
						InstanceSpec: types.InstanceSpec{
							MachineType:  machineType,
							InstanceType: instanceType,
							VPCConfig: types.VPCConfig{
								VPCSubnetID: subnetID,
							},
							HPASOption: &types.HPASOption{
								AppType:             hpasAppType,
								AppPerformanceLevel: hpasAppPerformanceLevel,
							},
							ImageID:       imageID,
							AdminPassword: adminPassword,
						},
					},
				},
			},
		},
	}
	if args == nil {
		fmt.Println("args is nil")
		return
	}
	resp, err := ccev2Client.CreateInstanceGroup(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))

}

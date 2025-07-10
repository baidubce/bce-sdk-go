package instancegroup

import (
	"encoding/json"
	"fmt"

	bccapi "github.com/baidubce/bce-sdk-go/services/bcc/api"
	ccev2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/types"
)

func AttachHPASInstance() {
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
	instanceGroupID := ""
	adminPassword := ""
	hpasID := ""
	rebuild := false
	securityGroupType := types.SecurityGroupTypeEnterprise
	securityGroupID := ""

	args := &ccev2.AttachInstancesToInstanceGroupArgs{
		ClusterID:       clusterID,
		InstanceGroupID: instanceGroupID,
		Request: &ccev2.AttachInstancesToInstanceGroupRequest{
			Incluster:                          false,
			UseInstanceGroupConfig:             false,
			UseInstanceGroupConfigWithDiskInfo: false,
			ExistedInstances: []*ccev2.InstanceSet{
				{
					InstanceSpec: types.InstanceSpec{
						AdminPassword: adminPassword,
						Existed:       true,
						ExistedOption: types.ExistedOption{
							ExistedInstanceID: hpasID,
							Rebuild:           &rebuild,
						},
						VPCConfig: types.VPCConfig{
							SecurityGroups: []types.SecurityGroupV2{
								{
									Type: securityGroupType,
									ID:   securityGroupID,
								},
							},
						},
						MachineType: types.MachineTypeHPAS,
						ClusterRole: types.ClusterRoleNode,
						InstanceOS: types.InstanceOS{
							ImageType: bccapi.ImageTypeSystem,
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
	resp, err := ccev2Client.AttachInstancesToInstanceGroup(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

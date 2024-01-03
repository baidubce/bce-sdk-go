package bccsgexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/bcc"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
)

// securityGroup sdk
// CreateSecurityGroup - create a security group
//
// PARAMS:
//   - args: the arguments to create security group
//
// RETURNS:
//   - *api.CreateSecurityGroupResult: the result of create security group
//   - error: nil if success otherwise the specific error
//
// CreateSecurityGroup 创建安全组
func CreateSecurityGroup() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "endpoint"
	client, _ := bcc.NewClient(ak, sk, endpoint) // 创建BCC Client
	args := &api.CreateSecurityGroupArgs{
		Name:  "testSecurityGroup",                   // 创建的SG名称
		VpcId: "Your VPCID",                          // 创建的SG所属VPCID
		Desc:  "vpc1 sdk test create security group", // 创建的SG描述
		Rules: []api.SecurityGroupRuleModel{
			{
				Remark:        "备注",      // 规则备注
				Protocol:      "tcp",     // 协议
				PortRange:     "1-65535", // 端口范围
				Direction:     "ingress", // 方向
				SourceIp:      "",        // 源IP
				SourceGroupId: "",        // 源SGID
			},
		},
		Tags: []model.TagModel{
			{
				TagKey:   "tagKey",   // 标签key
				TagValue: "tagValue", // 标签value
			},
		},
	}

	response, err := client.CreateSecurityGroup(args) // 创建SG

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

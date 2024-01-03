package bccsgexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bcc"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
)

// DeleteSecurityGroupRule - delete a security group rule
//
// PARAMS:
//   - securityGroupRuleId: the id of the specific security group rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DeleteSecurityGroupRule() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := bcc.NewClient(ak, sk, endpoint) // 创建bcc client
	args := &api.DeleteSecurityGroupRuleArgs{    // 创建参数
		SecurityGroupRuleId: "Your SecurityGroupRuleID", // 指定要删除的规则id
	}
	err := client.DeleteSecurityGroupRule(args) // 删除安全组规则
	if err != nil {
		panic(err)
	}
	fmt.Printf("Delete SecurityGroupRule %s success\n", args.SecurityGroupRuleId)
}

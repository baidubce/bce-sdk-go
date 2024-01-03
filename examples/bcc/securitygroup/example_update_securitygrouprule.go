package bccsgexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bcc"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
)

// UpdateSecurityGroupRule - update security group rule with the specific parameters
//
// PARAMS:
//   - args: the arguments to update the specific security group rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func UpdateSecurityGroupRule() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := bcc.NewClient(ak, sk, endpoint) // 创建bcc client
	remark := "test"                             // 描述备注
	args := &api.UpdateSecurityGroupRuleArgs{
		SecurityGroupRuleId: "Your SecurityGroupRuleId", // 安全组规则ID
		Remark:              &remark,
	}
	err := client.UpdateSecurityGroupRule(args) // 更新安全组规则
	if err != nil {
		panic(err)
	}
	fmt.Print("The security group rule has been updated successfully.\n")
}

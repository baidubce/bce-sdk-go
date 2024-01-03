package bccsgexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bcc"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
)

// RevokeSecurityGroupRule - revoke a security group rule
//
// PARAMS:
//   - securityGroupId: the specific securityGroup ID
//   - args: the arguments to revoke security group rule
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func RevokeSecurityGroupRule() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := bcc.NewClient(ak, sk, endpoint) // 创建BCC Client
	args := &api.RevokeSecurityGroupArgs{
		Rule: &api.SecurityGroupRuleModel{
			Remark:        "备注",      // 规则备注
			Protocol:      "tcp",     // 协议
			PortRange:     "1-65535", // 端口范围
			Direction:     "ingress", // 方向
			SourceIp:      "",        // 源IP
			SourceGroupId: "",        // 源SGID
		},
	}
	err := client.RevokeSecurityGroupRule("Your SecurityGroupID", args)
	if err != nil {
		panic(err)
	}
	fmt.Print("RevokeSecurityGroupRule Success!\n")
}

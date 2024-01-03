package esgexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/esg"
)

// AuthorizeEsgRule 函数用于授权企业安全组规则。
func AuthorizeEsgRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := esg.NewClient(ak, sk, endpoint) // 创建esg client
	args := &esg.CreateEsgRuleArgs{
		Rules: []esg.EnterpriseSecurityGroupRule{
			{
				Action:    "deny",        // 允许为allow, 不允许为deny
				Direction: "ingress",     // 方向
				Ethertype: "IPv4",        // IP协议类型
				PortRange: "1-65535",     // 端口范围
				Priority:  1000,          // 优先级
				Protocol:  "udp",         // 传输协议
				Remark:    "go sdk test", // 备注
				SourceIp:  "all",         // 源ip
			},
		},
		EnterpriseSecurityGroupId: "Your EnterpriseSecurityGroupId",
		ClientToken:               "ClientToken",
	}
	err := client.CreateEsgRules(args)
	if err != nil {
		panic(err)
	}
	fmt.Print("AuthorizeEsgRule success!\n")
}

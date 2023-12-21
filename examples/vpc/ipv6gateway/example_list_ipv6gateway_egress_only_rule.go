package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func ListIPv6GatewayEgressOnlyRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	ipv6gatewayID := "Your ipv6gateway's id"

	args := &vpc.ListIPv6GatewayEgressOnlyRuleArgs{
		// 设置每页包含的最大数量
		MaxKeys: 10,
		// 设置查询起始位置
		Marker: "Your EgressOnlyRule's id",
	}
	result, err := ipv6gatewayClient.ListIPv6GatewayEgressOnlyRule(ipv6gatewayID, args)
	if err != nil {
		fmt.Println("list ipv6 gateway egress only rules error: ", err)
		return
	}
	// 获取ipv6gatewayEgressOnlyRule列表
	for _, egressOnlyRule := range result.EgressOnlyRules {
		fmt.Println("ipv6 gateway egress only rule id: ", egressOnlyRule.EgressOnlyRuleId)
		fmt.Println("ipv6 gateway egress only rule cidr: ", egressOnlyRule.Cidr)
	}
}

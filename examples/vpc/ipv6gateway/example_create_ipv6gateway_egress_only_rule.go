package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func CreateIPv6GatewayEgressOnlyRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	ipv6gatewayID := "Your ipv6gateway's id"

	args := &vpc.CreateIPv6GatewayEgressOnlyRuleArgs{
		// 设置只出不进策略的CIDR
		Cidr: "2400:da00:e003:d01::/64",
	}
	result, err := ipv6gatewayClient.CreateIPv6GatewayEgressOnlyRule(ipv6gatewayID, args)
	if err != nil {
		fmt.Println("create ipv6 gateway egress only rule error: ", err)
		return
	}
	fmt.Println("ipv6 gateway egress only rule id: ", result.EgressOnlyRuleId)
}

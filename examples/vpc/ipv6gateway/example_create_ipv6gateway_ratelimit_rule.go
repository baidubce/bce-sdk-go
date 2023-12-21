package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func CreateIPv6GatewayRateLimitRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	ipv6gatewayID := "Your ipv6gateway's id"

	args := &vpc.CreateIPv6GatewayRateLimitRuleArgs{
		// 设置要限速IPv6的地址
		IPv6Address: "240c:4082:0:100::",
		// 设置限速的入向带宽
		IngressBandwidthInMbps: 10,
		// 设置限速的出向带宽
		EgressBandwidthInMbps: 10,
	}
	result, err := ipv6gatewayClient.CreateIPv6GatewayRateLimitRule(ipv6gatewayID, args)
	if err != nil {
		fmt.Println("create ipv6 gateway rate limit rule error: ", err)
		return
	}
	fmt.Println("ipv6 gateway rate limit rule id: ", result.RateLimitRuleId)
}

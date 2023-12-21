package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func UpdateIPv6GatewayRateLimitRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	ipv6gatewayID := "Your ipv6gateway's id"

	rateLimitRuleID := "Your rateLimitRule's id"

	args := &vpc.UpdateIPv6GatewayRateLimitRuleArgs{
		// 客户端Token，用于请求的幂等性
		ClientToken: "-",
		// 设置限速的入向带宽
		IngressBandwidthInMbps: 10,
		// 设置限速的出向带宽
		EgressBandwidthInMbps: 10,
	}
	if err := ipv6gatewayClient.UpdateIPv6GatewayRateLimitRule(ipv6gatewayID, rateLimitRuleID, args); err != nil {
		fmt.Println("update ipv6 gateway rate limit rule error: ", err)
		return
	}

	fmt.Printf("update ipv6 gateway %s rate limit rule %s success.", ipv6gatewayID, rateLimitRuleID)
}

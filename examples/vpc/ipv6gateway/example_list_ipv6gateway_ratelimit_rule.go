package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func ListIPv6GatewayRateLimitRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	ipv6gatewayID := "Your ipv6gateway's id"

	args := &vpc.ListIPv6GatewayRateLimitRuleArgs{
		// 设置每页包含的最大数量
		MaxKeys: 10,
		// 设置查询起始位置
		Marker: "Your RateLimitRule's id",
	}
	result, err := ipv6gatewayClient.ListIPv6GatewayRateLimitRule(ipv6gatewayID, args)
	if err != nil {
		fmt.Println("list ipv6 gateway rate limit rules error: ", err)
		return
	}
	// 获取ipv6gatewayRateLimitRule列表
	for _, rateLimitRule := range result.RateLimitRules {
		fmt.Println("ipv6 gateway rate limit rule id: ", rateLimitRule.RateLimitRuleId)
		fmt.Println("ipv6 gateway rate limit rule ingress bandwidthInMbps: ", rateLimitRule.IngressBandwidthInMbps)
		fmt.Println("ipv6 gateway rate limit rule egress bandwidthInMbps: ", rateLimitRule.EgressBandwidthInMbps)
	}
}

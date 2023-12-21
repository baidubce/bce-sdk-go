package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func ListIPv6Gateway() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	vpcID := "Your vpc's id"

	args := &vpc.ListIPv6GatewayArgs{
		// 设置ipv6网关所属的vpc id，必选
		VpcId: vpcID,
	}
	result, err := ipv6gatewayClient.ListIPv6Gateway(args)
	if err != nil {
		fmt.Println("list ipv6 gateway error: ", err)
		return
	}
	// 获取ipv6gateway信息
	fmt.Println("ipv6 gateway id: ", result.GatewayId)
	fmt.Println("ipv6 gateway name: ", result.Name)
	fmt.Println("ipv6 gateway vpcId: ", result.VpcId)
	fmt.Println("ipv6 gateway bandwidthInMbps: ", result.BandwidthInMbps)
	fmt.Printf("ipv6 gateway egress only rules: %v", result.EgressOnlyRules)
	fmt.Printf("ipv6 gateway rate limit rules: %v", result.RateLimitRules)
}

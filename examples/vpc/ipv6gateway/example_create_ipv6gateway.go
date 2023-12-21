package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func CreateIPv6Gateway() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpc.CreateIPv6GatewayArgs{
		// 设置ipv6网关的名称
		Name: "ipv6gateway-sdk-go",
		// 设置ipv6网关所属的vpc id
		VpcId: "vpc-id",
		// 设置ipv6网关的带宽上限
		BandwidthInMbps: 10,
		// 设置ipv6网关的计费信息
		Billing: &vpc.Billing{
			PaymentTiming: vpc.PAYMENT_TIMING_POSTPAID,
		},
	}
	result, err := ipv6gatewayClient.CreateIPv6Gateway(args)
	if err != nil {
		fmt.Println("create ipv6 gateway error: ", err)
		return
	}
	fmt.Println("ipv6 gateway id: ", result.GatewayId)
}

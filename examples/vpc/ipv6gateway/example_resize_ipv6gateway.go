package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func ResizeIPv6Gateway() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	ipv6gatewayID := "Your ipv6gateway's id"

	args := &vpc.ResizeIPv6GatewayArgs{
		// 客户端Token，用于请求的幂等性
		ClientToken: "-",
		// 指定IPv6网关的带宽
		BandwidthInMbps: 10,
	}
	if err := ipv6gatewayClient.ResizeIPv6Gateway(ipv6gatewayID, args); err != nil {
		fmt.Println("resize ipv6 gateway error: ", err)
		return
	}

	fmt.Printf("resize ipv6 gateway %s success.", ipv6gatewayID)
}

package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func DeleteIPv6Gateway() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	ipv6gatewayID := "Your ipv6gateway's id"

	args := &vpc.DeleteIPv6GatewayArgs{
		// 客户端Token，用于请求的幂等性
		ClientToken: "-",
	}
	if err := ipv6gatewayClient.DeleteIPv6Gateway(ipv6gatewayID, args); err != nil {
		fmt.Println("delete ipv6 gateway error: ", err)
		return
	}

	fmt.Printf("delete ipv6 gateway %s success.", ipv6gatewayID)
}

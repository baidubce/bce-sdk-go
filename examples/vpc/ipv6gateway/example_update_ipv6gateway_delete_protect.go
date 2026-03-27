package ipv6gateway

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func UpdateIPv6GatewayDeleteProtect() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	ipv6gatewayClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	ipv6gatewayID := "Your ipv6gateway's id"

	args := &vpc.UpdateIPv6GatewayDeleteProtectArgs{
		// 客户端Token，用于请求的幂等性
		ClientToken: "-",
		// 设置释放保护开关，true表示开启，false表示关闭
		DeleteProtect: true,
	}
	if err := ipv6gatewayClient.UpdateIPv6GatewayDeleteProtect(ipv6gatewayID, args); err != nil {
		fmt.Println("update ipv6 gateway delete protect error: ", err)
		return
	}

	fmt.Printf("update ipv6 gateway %s delete protect success.", ipv6gatewayID)
}

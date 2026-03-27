package etgatewayexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/etGateway"
)

// updateEtGateway 更新etGateway的函数
func UpdateEtGateway() {
	client, err := etGateway.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et gateway client, err: %v.\n", err)
		return
	}

	// 更新etGateway
	description := "test update et gateway"
	enableIpv6 := 1
	args := &etGateway.UpdateEtGatewayArgs{
		ClientToken:    getClientToken(),
		EtGatewayId:    "dcgw-iiyc0ers2qx4",
		Name:           "test-et-gateway",
		Description:    &description, // 使用指针类型，支持更新为空字符串
		Speed:          10,
		LocalCidrs:     []string{"10.240.0.0/16", "192.168.3.0/24"},
		EnableIpv6:     &enableIpv6,                           // IPv6功能开关，1开启0关闭
		Ipv6LocalCidrs: []string{"2400:da00:e003:0:15f::/87"}, // IPv6云端网络
	}
	if err = client.UpdateEtGateway(args); err != nil {
		fmt.Printf("Failed to update et gateway, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully update et gateway.")
}

package etgatewayexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/etGateway"
	"github.com/baidubce/bce-sdk-go/util"
)

// getClientToken 生成一个长度为32位的随机字符串作为客户端token。
func getClientToken() string {
	return util.NewUUID()
}

// updateEtGateway 更新etGateway的函数
func updateEtGateway() {
	client, err := etGateway.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et gateway client, err: %v.\n", err)
		return
	}

	// 更新etGateway
	args := &etGateway.UpdateEtGatewayArgs{
		ClientToken: getClientToken(),
		EtGatewayId: "dcgw-iiyc0ers2qx4",
		Name:        "test-et-gateway",
		Description: "test update et gateway",
		Speed:       10,
		LocalCidrs:  []string{"10.240.0.0/16", "192.168.3.0/24"},
	}
	if err = client.UpdateEtGateway(args); err != nil {
		fmt.Printf("Failed to update et gateway, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully update et gateway.")
}

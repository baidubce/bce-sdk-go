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

// BindEt 函数用于绑定ET网关
func BindEt() {
	client, err := etGateway.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et gateway client, err: %v.\n", err)
		return
	}

	args := etGateway.BindEtArgs{
		EtId:        "et-aaccd",
		ClientToken: getClientToken(),
		EtGatewayId: "dcgw-iiyc0ers2qx4",
		ChannelId:   "sdxs",
		LocalCidrs:  []string{"10.240.0.0/16", "192.168.3.0/24"},
	}
	err = client.BindEt(&args)
	if err != nil {
		fmt.Printf("Failed to bind et, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully bind et.")
}

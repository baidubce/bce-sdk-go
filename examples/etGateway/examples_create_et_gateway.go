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

// CreateEtGateway 函数用于创建ET网关
func CreateEtGateway() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := etGateway.NewClient(ak, sk, endpoint)      // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &etGateway.CreateEtGatewayArgs{
		Name:        "test-et-gateway",
		VpcId:       "vpc-2pa2x0bjt26i",
		Description: "test create et gateway",
		Speed:       100,
		EtId:        "et-aaccd",
		ChannelId:   "sdxs",
		LocalCidrs:  []string{"10.240.0.0/16", "192.168.3.0/24"},
		ClientToken: getClientToken(),
	}
	result, err := client.CreateEtGateway(args)
	if err != nil {
		fmt.Println("create et gateway error: ", err)
		return
	}

	fmt.Println("create et gateway success, et gateway id: ", result.EtGatewayId)
}

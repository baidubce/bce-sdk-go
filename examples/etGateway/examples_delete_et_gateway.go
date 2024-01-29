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

// DeleteEtGateway 删除指定网关
func DeleteEtGateway() {
	client, err := etGateway.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et gateway client, err: %v.\n", err)
		return
	}

	if err = client.DeleteEtGateway("dcgw-iiyc0ers2qx4", getClientToken()); err != nil {
		fmt.Printf("Failed to delete et gateway, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully delete et gateway.")
}

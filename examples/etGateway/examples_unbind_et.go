package etgatewayexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/etGateway"
)

// UnBindEt 函数用于解绑ET。
func UnBindEt() {
	client, err := etGateway.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et gateway client, err: %v.\n", err)
		return
	}

	err = client.UnBindEt("dcgw-iiyc0ers2qx4", getClientToken())
	if err != nil {
		fmt.Printf("Failed to unbind et, err: %v.\n", err)
		return
	}
	fmt.Printf("Unbind et success.\n")
}

package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func UnBindEip() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	err := client.UnBindEip(
		"vpnId",            // vpn id
		"Your ClientToken", // client token
	)
	if err != nil {
		fmt.Printf("unbind eip error: %+v\n", err)
		return
	}

	fmt.Println("unbind eip success")
}

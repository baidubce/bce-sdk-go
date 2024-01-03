package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func DeleteVpnGateway() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	err := client.DeleteVpnGateway(
		"vpnId",            // vpn id
		"Your ClientToken", // client token
	)
	if err != nil {
		fmt.Printf("delete vpn error: %+v\n", err)
		return
	}

	fmt.Println("delete vpn success")
}

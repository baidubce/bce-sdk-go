package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func DeleteSslVpnServer() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	err := client.DeleteSslVpnServer(
		"Your Vpn Id",          // vpn id
		"Your SslVpnServer Id", // ssl vpn server id
		"Your ClientToken",     // client token
	)
	if err != nil {
		fmt.Printf("delete ssl vpn server error: %+v\n", err)
		return
	}

	fmt.Println("delete ssl vpn server success")
}

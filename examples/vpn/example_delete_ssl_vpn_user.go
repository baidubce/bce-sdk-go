package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func DeleteSslVpnUser() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	err := client.DeleteSslVpnUser(
		"Your Vpn Id",      // vpn id
		"Your User Id",     // user id
		"Your ClientToken", // client token
	)
	if err != nil {
		fmt.Printf("delete ssl vpn user error: %+v\n", err)
		return
	}

	fmt.Println("delete ssl vpn user success")
}

package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func GetSslVpnServer() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	result, err := client.GetSslVpnServer(
		"Your Vpn Id",      // vpn id
		"Your ClientToken", // client token
	)
	if err != nil {
		fmt.Printf("get ssl vpn server error: %+v\n", err)
		return
	}

	fmt.Println("get ssl vpn server success, result: ", result)
}

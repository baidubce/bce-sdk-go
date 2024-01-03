package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func DeleteVpnConn() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	err := client.DeleteVpnConn(
		"Your VpnConnId",   // vpn conn id
		"Your ClientToken", // client token
	)
	if err != nil {
		fmt.Printf("delete vpn conn error: %+v\n", err)
		return
	}

	fmt.Println("delete vpn conn success")
}

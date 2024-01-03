package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func GetVpnGatewayDetail() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	result, err := client.GetVpnGatewayDetail("Your VPN ID")
	if err != nil {
		fmt.Printf("get vpn error: %+v\n", err)
		return
	}

	fmt.Println("get vpn success, result: ", result)
}

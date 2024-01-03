package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func UpdateVpnGateway() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpn.UpdateVpnGatewayArgs{
		ClientToken: "Your ClientToken", // client token
		Name:        "Your VPN Name",    // vpn name
	}

	err := client.UpdateVpnGateway(
		"vpnId", // vpn id
		args,
	)
	if err != nil {
		fmt.Printf("update vpn error: %+v\n", err)
		return
	}

	fmt.Println("update vpn success")
}

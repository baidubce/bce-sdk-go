package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func UpdateVpnDeleteProtect() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpn.UpdateVpnDeleteProtectArgs{
		ClientToken:   "Your ClientToken", // client token
		VpnId:         "Your Vpn Id",      // vpn id
		DeleteProtect: true,
	}
	err := client.UpdateVpnDeleteProtect(args)
	if err != nil {
		fmt.Printf("update delete protect error: %+v\n", err)
		return
	}

	fmt.Println("update delete protect success")
}

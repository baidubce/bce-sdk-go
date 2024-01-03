package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func BindEip() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpn.BindEipArgs{
		ClientToken: "Your ClientToken", // client token
		Eip:         "Your Eip Address", // eip address
	}
	err := client.BindEip(
		"vpnId", // vpn id
		args,
	)
	if err != nil {
		fmt.Printf("bind eip error: %+v\n", err)
		return
	}

	fmt.Println("bind eip success")
}

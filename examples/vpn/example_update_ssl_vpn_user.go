package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func UpdateSslVpnUser() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	desc := "333"
	args := &vpn.UpdateSslVpnUserArgs{
		ClientToken: "Your ClientToken", // client token
		VpnId:       "Your Vpn Id",      // vpn id
		UserId:      "Your User Id",     // user id
		SslVpnUser: &vpn.UpdateSslVpnUser{
			Password:    "psd123456!", // 密码
			Description: &desc,        // 描述
		},
	}
	err := client.UpdateSslVpnUser(args)
	if err != nil {
		fmt.Printf("update ssl vpn user error: %+v\n", err)
		return
	}

	fmt.Println("update ssl vpn user success")
}

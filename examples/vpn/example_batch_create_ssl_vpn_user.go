package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func BatchCreateSslVpnUser() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	desc1 := "user1 description"
	args := &vpn.BatchCreateSslVpnUserArgs{
		ClientToken: "Your ClientToken", // client token
		VpnId:       "Your Vpn Id",      // vpn id
		SslVpnUsers: []vpn.SslVpnUser{
			{
				UserName:    "user1test",  // 用户名
				Password:    "psd123456!", // 密码
				Description: &desc1,       // 描述
			},
			{
				UserName: "user2test",  // 用户名
				Password: "psd123456!", // 密码
			},
		},
	}
	result, err := client.BatchCreateSslVpnUser(args)
	if err != nil {
		fmt.Printf("batch create ssl vpn user error: %+v\n", err)
		return
	}

	fmt.Println("batch create ssl vpn user success, result: ", result)
}

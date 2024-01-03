package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func ListSslVpnUser() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpn.ListSslVpnUserArgs{
		VpnId: "Your Vpn Id", // vpn id
	}
	result, err := client.ListSslVpnUser(args)
	if err != nil {
		fmt.Printf("list ssl vpn user error: %+v\n", err)
		return
	}

	fmt.Println("list ssl vpn user success, result: ", result)
}

package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func ListVpnConn() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	result, err := client.ListVpnConn("Your VpnConnId")
	if err != nil {
		fmt.Printf("list vpn conn error: %+v\n", err)
		return
	}

	fmt.Println("list vpn conn success, result: ", result)
}

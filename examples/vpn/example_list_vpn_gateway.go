package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func ListVpnGateway() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpn.ListVpnGatewayArgs{
		VpcId: "vpcId", // vpc id
	}
	results, err := client.ListVpnGateway(args)
	if err != nil {
		fmt.Printf("list vpn error: %+v\n", err)
		return
	}

	fmt.Println("list vpn success, results: ", results)
}

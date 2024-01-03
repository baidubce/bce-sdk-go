package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func CreateSslVpnServer() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	interfaceTypeStr := "tun"
	clientDnsStr := ""
	args := &vpn.CreateSslVpnServerArgs{
		ClientToken:      "Your ClientToken", // client token
		VpnId:            "Your Vpn Id",      // vpn id
		SslVpnServerName: "server_1",
		InterfaceType:    &interfaceTypeStr,
		LocalSubnets:     []string{"192.168.0.0/20"},
		RemoteSubnet:     "192.168.100.0/24",
		ClientDns:        &clientDnsStr,
	}
	result, err := client.CreateSslVpnServer(args)
	if err != nil {
		fmt.Printf("create ssl vpn server error: %+v\n", err)
		return
	}

	fmt.Println("create ssl vpn server success, result: ", result)
}

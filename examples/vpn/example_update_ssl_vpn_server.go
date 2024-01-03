package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func UpdateSslVpnServer() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	clientDnsStr := ""
	args := &vpn.UpdateSslVpnServerArgs{
		ClientToken:    "Your ClientToken",     // client token
		VpnId:          "Your Vpn Id",          // vpn id
		SslVpnServerId: "Your SslVpnServer Id", // ssl vpn server id
		UpdateSslVpnServer: &vpn.UpdateSslVpnServer{
			SslVpnServerName: "Your SslVpnServer Name",   // ssl vpn server name
			LocalSubnets:     []string{"192.168.0.0/20"}, // local subnets
			RemoteSubnet:     "192.168.100.0/24",         // remote subnet
			ClientDns:        &clientDnsStr,              // client dns
		},
	}
	err := client.UpdateSslVpnServer(args)
	if err != nil {
		fmt.Printf("update ssl vpn server error: %+v\n", err)
		return
	}

	fmt.Println("update ssl vpn server success")
}

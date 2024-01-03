package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func UpdateVpnConn() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpn.UpdateVpnConnArgs{
		VpnConnId: "",
		UpdateVpnconn: &vpn.CreateVpnConnArgs{
			VpnId:         "Your VPN ID",                //	vpn id
			VpnConnName:   "Your VPN Conn Name",         // vpn conn name
			LocalIp:       "0.1.2.3",                    // local ip
			SecretKey:     "!sdse154d",                  // secret key
			LocalSubnets:  []string{"192.168.0.0/20"},   // local subnets
			RemoteIp:      "3.4.5.6",                    // remote ip
			RemoteSubnets: []string{"192.168.100.0/24"}, // remote subnets
			CreateIkeConfig: &vpn.CreateIkeConfig{ // ike config
				IkeVersion:  "v1",
				IkeMode:     "main",
				IkeEncAlg:   "aes",
				IkeAuthAlg:  "sha1",
				IkePfs:      "group2",
				IkeLifeTime: 25500,
			},
			CreateIpsecConfig: &vpn.CreateIpsecConfig{ // ipsec config
				IpsecEncAlg:   "aes",
				IpsecAuthAlg:  "sha1",
				IpsecPfs:      "group2",
				IpsecLifetime: 25500,
			},
		},
	}
	err := client.UpdateVpnConn(args)
	if err != nil {
		fmt.Printf("update vpn conn error: %+v\n", err)
		return
	}

	fmt.Println("update vpn conn success")
}

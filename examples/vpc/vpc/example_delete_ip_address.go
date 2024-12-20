package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func DeleteIpAddress() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "c587aab8-cc6d-4e36-a7a6-b78339b1469f"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	deleteIpAddressArgs := &vpc.DeleteIpAddressArgs{
		ClientToken: clientToken,
		IpAddressInfo: []string{ // 删除的IP地址信息，单次最多指定10个
			"192.168.11.0/24",
			"192.168.12.0/24",
		},
	}
	ipSetID := "ips-2etsti1g24hv"
	err := client.DeleteIpAddress(ipSetID, deleteIpAddressArgs) // IP地址组删除IP地址

	if err != nil {
		fmt.Println(err)
	}
}

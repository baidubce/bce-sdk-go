package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func DelIpAddress() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "c587aab8-cc6d-4e36-a7a6-b78339b1469f"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	delIpAddressArgs := &vpc.DelIpAddressArgs{
		ClientToken: clientToken,
		IpAddressInfo: []string{
			"192.168.11.0/24",
			"192.168.12.0/24",
		},
	}
	IpSetID := "ips-2etsti1g24hv"
	err := client.DelIpAddress(IpSetID, delIpAddressArgs) // 创建vpc

	if err != nil {
		fmt.Println(err)
	}
}

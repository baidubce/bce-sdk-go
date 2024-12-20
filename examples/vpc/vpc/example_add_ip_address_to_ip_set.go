package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func AddIpAddressToIpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "c587aab8-cc6d-4e36-a7a6-b78339b1469f"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	addIpAddress2IpSetArgs := &vpc.AddIpAddress2IpSetArgs{
		ClientToken: clientToken,
		IpAddressInfo: []vpc.TemplateIpAddressInfo{ // 添加的IP地址信息，其ipVersion需与指定的IP地址组保持一致，单次最多指定10个
			{IpAddress: "192.168.11.0/24", Description: "test1"},
			{IpAddress: "192.168.12.0/24", Description: "test2"},
		},
	}
	ipSetID := "ips-2etsti1g24hv"
	err := client.AddIpAddress2IpSet(ipSetID, addIpAddress2IpSetArgs) // IP地址组添加IP地址

	if err != nil {
		fmt.Println(err)
	}
}

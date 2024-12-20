package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func CreateIpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "c587aab8-cc6d-4e36-a7a6-b78339b1469f"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	createIpSetArgs := &vpc.CreateIpSetArgs{
		ClientToken: clientToken,
		Name:        "test_create_ip_set", // IP地址组的名称
		IpVersion:   "IPv4",               // ipVersion，取值IPv4或IPv6
		IpAddressInfo: []vpc.TemplateIpAddressInfo{ // 参数模板IP地址信息，单次最多指定10个
			{IpAddress: "192.168.11.0/24", Description: "test1"},
			{IpAddress: "192.168.12.0/24", Description: "test2"},
		},
		Description: "this is a test", // IP地址组描述
	}

	response, err := client.CreateIpSet(createIpSetArgs) // 创建IP地址组

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

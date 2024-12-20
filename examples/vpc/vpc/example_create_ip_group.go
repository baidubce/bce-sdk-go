package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func CreateIpGroup() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "c587aab8-cc6d-4e36-a7a6-b78339b1469f"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	createIpGroupArgs := &vpc.CreateIpGroupArgs{
		ClientToken: clientToken,
		Name:        "test_create_ip_set", // IP地址族的名称
		IpVersion:   "IPv4",               // ipVersion，取值IPv4或IPv6
		IpSetIds: []string{ // 关联的IP地址组ID，其ipVersion需与本次创建的IP地址族一致，单次最多指定5个
			"ips-z2a8uk9qnkc1",
			"ips-hms1n8fu184f",
		},
		Description: "this is a test", //IP地址族描述
	}

	response, err := client.CreateIpGroup(createIpGroupArgs)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

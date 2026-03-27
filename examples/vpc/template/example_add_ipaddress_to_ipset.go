package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func AddIpAddress2IpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	ipSetId := "ips-xxxxx"
	args := &vpc.AddIpAddress2IpSetArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 添加的IP地址信息列表，单次最多10个
		IpAddressInfo: []vpc.TemplateIpAddressInfo{
			{IpAddress: "192.168.2.1", Description: "new ip1"},
		},
	}
	err = client.AddIpAddress2IpSet(ipSetId, args)
	if err != nil {
		fmt.Println("add ip address to ip set error: ", err)
		return
	}

	fmt.Println("add ip address to ip set success")
}

package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func CreateIpGroup() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.CreateIpGroupArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 设置IP地址族名称
		Name: "test_ip_group",
		// 设置IP版本，取值IPv4或IPv6
		IpVersion: "IPv4",
		// 关联的IP地址组ID列表，单次最多5个
		IpSetIds: []string{
			"ips-xxxxx",
		},
		// 设置IP地址族描述（可选）
		Description: "this is a test ip group",
	}
	result, err := client.CreateIpGroup(args)
	if err != nil {
		fmt.Println("create ip group error: ", err)
		return
	}

	fmt.Println("create ip group success, ip group id: ", result.IpGroupId)
}

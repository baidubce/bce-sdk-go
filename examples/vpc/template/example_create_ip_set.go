package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
	"github.com/baidubce/bce-sdk-go/util"
)

func getClientToken() string {
	return util.NewUUID()
}

func CreateIpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.CreateIpSetArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 设置IP地址组名称
		Name: "test_ip_set",
		// 设置IP版本，取值IPv4或IPv6
		IpVersion: "IPv4",
		// 设置IP地址信息列表
		IpAddressInfo: []vpc.TemplateIpAddressInfo{
			{IpAddress: "192.168.1.0/24", Description: "test cidr1"},
			{IpAddress: "10.0.0.1", Description: "test ip1"},
		},
		// 设置IP地址组描述（可选）
		Description: "this is a test ip set",
	}
	result, err := client.CreateIpSet(args)
	if err != nil {
		fmt.Println("create ip set error: ", err)
		return
	}

	fmt.Println("create ip set success, ip set id: ", result.IpSetId)
}

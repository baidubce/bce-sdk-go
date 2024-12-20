package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ListIpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	ListIpSetArgs := &vpc.ListIpSetArgs{
		IpVersion: "IPv4", // ipVersion，取值IPv4或IPv6
		Marker:    "",     // 批量获取列表的查询的起始位置，是一个由系统生成的字符串
		MaxKeys:   2,      // 每页包含的最大数量，最大数量不超过1000。缺省值为1000
	}
	response, err := client.ListIpSet(ListIpSetArgs) // 查询IP地址组列表

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

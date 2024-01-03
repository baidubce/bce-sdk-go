package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ListVpc() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	args := &vpc.ListVPCArgs{
		MaxKeys: 100, // 设置每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
		Marker:  "",  // 设置批量获取列表的查询的起始位置，是一个由系统生成的字符串
	}

	response, err := client.ListVPC(args) // 获取vpc列表

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

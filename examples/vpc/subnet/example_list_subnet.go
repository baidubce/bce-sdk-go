package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ListSubnet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	args := &vpc.ListSubnetArgs{
		MaxKeys: 100,                // 设置每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
		Marker:  "",                 // 设置批量获取列表的查询的起始位置，是一个由系统生成的字符串
		VpcId:   "vpc-p1eawhw5rx4n", // 子网所属vpc id
	}

	response, err := client.ListSubnets(args) // 获取子网列表

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

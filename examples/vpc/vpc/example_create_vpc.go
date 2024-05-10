package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func CreateVpc() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	createVpcArgs := &vpc.CreateVPCArgs{
		Name:        "TestSDK-VPC",    // vpc名称
		Description: "vpc test",       // vpc描述
		Cidr:        "192.168.0.0/16", // vpc网段
		EnableIpv6:  true,             // 是否分配Ipv6网段
		Tags: []model.TagModel{ // vpc标签
			{
				TagKey:   "tagK",
				TagValue: "tagV",
			},
		},
	}
	response, err := client.CreateVPC(createVpcArgs) // 创建vpc

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

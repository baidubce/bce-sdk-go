package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func CreateSubnet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	createSubnetArgs := &vpc.CreateSubnetArgs{
		Name:        "TestSDK-Subnet",   // 子网名称
		Description: "vpc test",         // 子网描述
		Cidr:        "192.168.96.0/20",  // 子网网段
		VpcId:       "vpc-p1eawhw5rx4n", // vpc id
		ZoneName:    "cn-bj-a",          // 子网可用区
	}
	response, err := client.CreateSubnet(createSubnetArgs) // 创建子网

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

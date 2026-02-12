package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func GetResourceIp() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	getResourceIpArgs := &vpc.GetResourceIpArgs{
		VpcId:    "vpc-xxxxx", // VPC ID
		PageNo:   1,           // 页码，从1开始，默认值1
		PageSize: 100,         // 每页数量，取值范围[1,1000]，默认值100
	}
	response, err := client.GetResourceIp(getResourceIpArgs) // 查询VPC内产品占用IP

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

func GetResourceIpWithFilter() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	getResourceIpArgs := &vpc.GetResourceIpArgs{
		VpcId:        "vpc-xxxxx", // VPC ID
		SubnetId:     "sbn-xxxxx", // 子网ID（可选）
		ResourceType: "enic",      // 产品类型（可选），如bcc、enic、blb等
		PageNo:       1,           // 页码
		PageSize:     50,          // 每页数量
	}
	response, err := client.GetResourceIp(getResourceIpArgs) // 查询VPC内产品占用IP（带过滤条件）

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

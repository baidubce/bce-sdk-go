package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ShowSubnet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	subnetID := "sbn-u166vdnqqubi" // 子网id

	response, err := client.GetSubnetDetail(subnetID) // 获取子网详情

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

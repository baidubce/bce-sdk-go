package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ShowVpc() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	vpcID := "vpc-p1eawhw5rx4n" // vpc id

	response, err := client.GetVPCDetail(vpcID) // 获取vpc详情

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

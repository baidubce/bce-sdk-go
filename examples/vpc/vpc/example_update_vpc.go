package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// UpdateVpc 更新VPC基本信息
func UpdateVpc() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	clientToken := "be31b98c-5e41-4838-9830-9be700de5a20" // 幂等性Token，是一个长度不超过64位的ASCII字符串
	vpcID := "vpc-p1eawhw5rx4n"                           // vpc id

	enableIpv6 := true
	updateVpcArgs := &vpc.UpdateVPCArgs{
		ClientToken: clientToken,
		Name:        "test_vpc",    // 更新vpc名称
		Description: "vpc updated", // 更新vpc描述
		EnableIpv6:  &enableIpv6,   // 更新vpc是否开启ipv6功能（使用指针类型）
		SecondaryCidr: []string{ // 更新vpc辅助网段
			"172.16.0.0/16",
		},
	}
	err := client.UpdateVPC(vpcID, updateVpcArgs) // 更新vpc

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("update vpc success.")
}

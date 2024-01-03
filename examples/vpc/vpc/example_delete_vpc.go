package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func DeleteVpc() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	vpcID := "vpc-is0j69zp08cx"                           // 要删除的VPC的ID
	clientToken := "be31b98c-5141-4838-9830-9be700de5a20" // 幂等性Token，是一个长度不超过64位的ASCII字符串，见 https://cloud.baidu.com/doc/VPC/s/gjwvyu77i#%E5%B9%82%E7%AD%89%E6%80%A7

	err := client.DeleteVPC(vpcID, clientToken) // 删除 VPC

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("delete vpc success.")
}

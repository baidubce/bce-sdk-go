package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UpdateSubnet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	clientToken := "be31b98c-5e41-4838-9830-9be700de5a20" // 幂等性Token，是一个长度不超过64位的ASCII字符串，见 https://cloud.baidu.com/doc/VPC/s/gjwvyu77i#%E5%B9%82%E7%AD%89%E6%80%A7
	subnetID := "sbn-u166vdnqqubi"                        // 子网ID

	updateSubnetArgs := &vpc.UpdateSubnetArgs{
		ClientToken: clientToken,
		Name:        "test_subnet", // 更新子网名称
		Description: "",            // 更新子网描述
	}
	err := client.UpdateSubnet(subnetID, updateSubnetArgs) // 更新子网

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("update subnet success.")
}

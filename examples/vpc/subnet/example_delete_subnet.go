package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func DeleteSubnet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)

	subnetID := "sbn-5uj2t9f1jazj"                        // 子网ID
	clientToken := "be31b98c-5141-4838-9830-9be700de5a20" // 幂等性Token，是一个长度不超过64位的ASCII字符串，见 https://cloud.baidu.com/doc/VPC/s/gjwvyu77i#%E5%B9%82%E7%AD%89%E6%80%A7

	err := client.DeleteSubnet(subnetID, clientToken) // 删除子网

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("delete subnet success.")
}

package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func DeleteIpGroup() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "c587aab8-cc6d-4e36-a7a6-b78339b1469f"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	deleteIpGroupArgs := &vpc.DeleteIpGroupArgs{
		ClientToken: clientToken,
	}
	ipGroupID := "ipg-riivgeymsiwe"                           // IP地址族的ID
	err := client.DeleteIpGroup(ipGroupID, deleteIpGroupArgs) // 删除IP地址族

	if err != nil {
		fmt.Println(err)
	}
}

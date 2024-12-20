package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UpdateIpGroup() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "c587aab8-cc6d-4e36-a7a6-b78339b1469f"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	updateIpGroupArgs := &vpc.UpdateIpGroupArgs{
		ClientToken: clientToken,
		Name:        "test_update_ip_group", // IP地址族的名称
		Description: "this is a test",       // IP地址族的描述
	}

	ipGroupID := "ipg-9vd6xtyjz0in"                           // IP地址族的ID
	err := client.UpdateIpGroup(ipGroupID, updateIpGroupArgs) // 更新IP地址族

	if err != nil {
		fmt.Println(err)
	}
}

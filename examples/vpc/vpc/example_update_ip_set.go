package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UpdateIpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "c587aab8-cc6d-4e36-a7a6-b78339b1469f"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	updateIpSetArgs := &vpc.UpdateIpSetArgs{
		Name:        "test_update_ip_set", // IP地址组的名称
		Description: "this is a test",     // IP地址组的描述
		ClientToken: clientToken,
	}
	ipSetID := "ips-2etsti1g24hv"                       // IP地址组的ID
	err := client.UpdateIpSet(ipSetID, updateIpSetArgs) // 更新IP地址组

	if err != nil {
		fmt.Println(err)
	}
}

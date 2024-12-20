package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UnbindIpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "c587aab8-cc6d-4e36-a7a6-b78339b1469f"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	unbindIpsetArgs := &vpc.UnbindIpSetArgs{
		ClientToken: clientToken,
		IpSetIds: []string{ // 	移除的IP地址组ID，单次最多指定5个
			"ips-5eekehr75vbv",
			"ips-vn4nfjau2t2u",
		},
	}
	ipGroupID := "ipg-9vd6xtyjz0in"                       // IP地址族的ID
	err := client.UnbindIpSet(ipGroupID, unbindIpsetArgs) // IP地址族移除IP地址组

	if err != nil {
		fmt.Println(err)
	}
}

package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func UpdateDnatRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	DnatRuleID := "Your dnat rule's id"

	// replace with your data
	args := &vpc.UpdateNatGatewayDnatRuleArgs{
		RuleName:         "dnat_go_update",
		PublicIpAddress:  "100.88.14.90",
		PrivateIpAddress: "192.168.1.2",
		Protocol:         "TCP",
		PublicPort:       "1212",
		PrivatePort:      "1212",
	}
	err := natClient.UpdateNatGatewayDnatRule(NatID, DnatRuleID, args)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("dnat rule update successfully")
}

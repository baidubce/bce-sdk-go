package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func BatchCreateDnatRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"
	// replace by your data
	args := &vpc.BatchCreateNatGatewayDnatRuleArgs{
		Rules: []vpc.DnatRuleArgs{
			{
				RuleName:         "dnat_go_1",
				PublicIpAddress:  "100.88.14.90",
				PrivateIpAddress: "192.168.1.1",
				Protocol:         "TCP",
				PublicPort:       "1212",
				PrivatePort:      "1212",
			},
			{
				RuleName:         "dnat_go_2",
				PublicIpAddress:  "100.88.14.52",
				PrivateIpAddress: "192.168.1.2",
				Protocol:         "UDP",
				PublicPort:       "65535",
				PrivatePort:      "65535",
			},
		},
	}
	result, err := natClient.BatchCreateNatGatewayDnatRule(NatID, args)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("dnat rules id: ", result.RuleIds)
}

package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func CreateSnatRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	args := &vpc.CreateNatGatewaySnatRuleArgs{
		RuleName:          "snat_go",
		SourceCIDR:        "192.168.1.0/24",         // 源网段，eg. 192.168.1.0/24
		PublicIpAddresses: []string{"100.88.14.90"}, // 替换为需要绑定的 EIP 列表
	}
	result, err := natClient.CreateNatGatewaySnatRule(NatID, args)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("snat rule id: ", result.RuleId)
}

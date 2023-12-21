package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func UpdateSnatRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	SnatRuleID := "Your snat rule's id"

	args := &vpc.UpdateNatGatewaySnatRuleArgs{
		RuleName:          "snat_go_modify",
		SourceCIDR:        "192.168.1.0/24",    // 源网段，eg. 192.168.1.0/24
		PublicIpAddresses: []string{"x.x.x.x"}, // 替换为需要绑定的 EIP 列表
	}
	if err := natClient.UpdateNatGatewaySnatRule(NatID, SnatRuleID, args); err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("ok.")
}

package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func BatchCreateSnatRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	args := &vpc.BatchCreateNatGatewaySnatRuleArgs{
		NatId: NatID,
		SnatRules: []vpc.SnatRuleArgs{
			{RuleName: "batchTest1", SourceCIDR: "192.168.1.0/24", PublicIpAddresses: []string{"100.88.14.90"}},
			{RuleName: "batchTest2", SourceCIDR: "192.168.16.0/20", PublicIpAddresses: []string{"100.88.14.90"}},
		},
	}
	results, err := natClient.BatchCreateNatGatewaySnatRule(args)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Println("snat rules id: ", results.SnatRuleIds)
}

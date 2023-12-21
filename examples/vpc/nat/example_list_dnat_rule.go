package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func ListDnatRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	args := &vpc.ListNatGatewaDnatRuleArgs{
		MaxKeys: 10,
	}
	result, err := natClient.ListNatGatewayDnatRules(NatID, args)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("Get %d dnat rules: \n", len(result.Rules))
}

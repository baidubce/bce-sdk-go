package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func DeleteDnatRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	DnatRuleID := "Your dnat rule's id"

	clientToken := "-" // 客户端Token，用于请求的幂等性

	if err := natClient.DeleteNatGatewayDnatRule(NatID, DnatRuleID, clientToken); err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("ok.")
}

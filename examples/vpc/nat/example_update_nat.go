package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func UpdateNat() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	args := &vpc.UpdateNatGatewayArgs{
		// 设置nat网关的最新名称
		Name: "GO-SDK-TestNatUpdate",
	}

	if err := natClient.UpdateNatGateway(NatID, args); err != nil {
		fmt.Println("update nat gateway error: ", err)
		return
	}

	fmt.Printf("update nat gateway %s success.", NatID)
}

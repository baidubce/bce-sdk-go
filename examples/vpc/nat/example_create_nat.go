package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func CreateNat() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpc.CreateNatGatewayArgs{
		// 设置nat网关的名称
		Name: "nat-sdk-go",
		// 设置nat网关所属的vpc id
		VpcId: "vpc-id",
		// 设置nat网关的规格
		Spec: vpc.NAT_GATEWAY_SPEC_SMALL,
		// 设置nat网关的snat eip列表
		Eips: []string{},
		// 设置nat网关的dnat eip列表
		DnatEips: []string{},
		// 设置nat绑定的资源组ID，此字段选传，传则表示绑定资源组
		ResourceGroupId: "ResourceGroupId",
		// 设置nat网关的计费信息
		Billing: &vpc.Billing{
			PaymentTiming: vpc.PAYMENT_TIMING_PREPAID,
		},
	}
	result, err := natClient.CreateNatGateway(args)
	if err != nil {
		fmt.Println("create nat gateway error: ", err)
		return
	}
	fmt.Println("create nat gateway success, nat gateway id: ", result.NatId)
}

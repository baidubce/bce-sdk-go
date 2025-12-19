package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func ListNat() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	VPCID := "Your vpc's id"

	args := &vpc.ListNatGatewayArgs{
		// 设置nat网关所属的vpc id，必选
		VpcId: VPCID,
		// 指定查询的NAT的Id
	}
	result, err := natClient.ListNatGateway(args)
	if err != nil {
		fmt.Println("list nat gateway error: ", err)
		return
	}
	// 获取nat的列表信息
	for _, nat := range result.Nats {
		fmt.Println("nat id: ", nat.Id)
		fmt.Println("nat name: ", nat.Name)
		fmt.Println("nat vpcId: ", nat.VpcId)
		fmt.Println("nat type: ", nat.NatType)
		fmt.Println("nat spec: ", nat.Spec)
		fmt.Println("nat snat eips: ", nat.Eips)
		fmt.Println("nat dnat eips: ", nat.DnatEips)
		fmt.Println("nat bind eips: ", nat.BindEips)
		fmt.Println("nat status: ", nat.Status)
		fmt.Println("nat paymentTiming: ", nat.PaymentTiming)
		fmt.Println("nat expireTime: ", nat.ExpiredTime)
		fmt.Println("nat ipVerion: ", *nat.IpVersion)
	}
}

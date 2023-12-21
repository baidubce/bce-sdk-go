package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func GetNatDetail() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	result, err := natClient.GetNatGatewayDetail(NatID)
	if err != nil {
		fmt.Println("get nat gateway details error: ", err)
		return
	}

	// 查询得到nat网关的id
	fmt.Println("nat id: ", result.Id)
	// 查询得到nat网关的名称
	fmt.Println("nat name: ", result.Name)
	// 查询得到nat网关所属的vpc id
	fmt.Println("nat vpcId: ", result.VpcId)
	// 查询得到nat网关的大小
	fmt.Println("nat spec: ", result.Spec)
	// 查询得到nat网关绑定的snat EIP的IP地址列表
	fmt.Println("nat snat eips: ", result.Eips)
	// 查询得到nat网关绑定的dnat EIP的IP地址列表
	fmt.Println("nat dnat eips: ", result.DnatEips)
	// 查询得到nat网关的状态
	fmt.Println("nat status: ", result.Status)
	// 查询得到nat网关的付费方式
	fmt.Println("nat paymentTiming: ", result.PaymentTiming)
	// 查询得到nat网关的过期时间
	fmt.Println("nat expireTime: ", result.ExpiredTime)
}

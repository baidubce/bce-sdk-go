package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func RenewNat() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	args := &vpc.RenewNatGatewayArgs{
		// 设置nat网关续费的订单信息
		Billing: &vpc.Billing{
			Reservation: &vpc.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
	}
	if err := natClient.RenewNatGateway(NatID, args); err != nil {
		fmt.Println("renew nat gateway error: ", err)
		return
	}

	fmt.Printf("renew nat gateway %s success.", NatID)
}

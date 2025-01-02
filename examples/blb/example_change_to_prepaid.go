package blbexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/blb"
)

func ChangeToPrePaid() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	blbClient, _ := blb.NewClient(ak, sk, endpoint) // 初始化client

	blbID := "Your BlbID" // 后付费BLB实例ID

	orderIdResult, err := blbClient.ChangeToPrepaid(blbID, &blb.ChangeToPrepaidArgs{
		BillingMethod:     "BySpec", // 计费方式:仅支持"BySpec"(预付费-按规格)，默认"BySpec"
		PerformanceLevel:  "small2", // 性能规格:取值"small1","small2","medium1","medium2","large1","large2","large3"
		ReservationLength: 2,        // 购买月份，必填。
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("orderId=%s\n", orderIdResult.OrderId)
}

package blbexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/blb"
)

func ChangeToPostPaid() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	blbClient, _ := blb.NewClient(ak, sk, endpoint) // 初始化client

	blbID := "Your BlbID" // 预付费BLB实例ID

	// billingMethod 非必填
	// "BySpec"表示转后付费-按固定规格计费。"ByCapacityUnit"表示转后付费-按使用量计费
	billingMethod := "ByCapacityUnit"

	// performanceLevel: 非必填
	// 取值"small1","small2","medium1","medium2","large1","large2","large3"
	// 若billingMethod="ByCapacityUnit"还可选择"unlimited"
	performanceLevel := "small2"

	orderIdResult, err := blbClient.ChangeToPostpaid(blbID, &blb.ChangeToPostpaidArgs{
		BillingMethod:    billingMethod,
		PerformanceLevel: performanceLevel,
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("orderId=%s\n", orderIdResult.OrderId)
}

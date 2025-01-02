package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func GetCsnBpPrice() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	request := &csn.GetCsnBpPriceRequest{
		Name:        "test",  // 带宽包的名称，不能为空
		Bandwidth:   2,       // 带宽包的带宽值，后付费按流量不需要该值
		GeographicA: "China", // 网络实例所属的区域。取值 [ China | Asia-Pacific ]，分别表示中国大陆、亚太区域
		GeographicB: "China", // 另一个网络实例所属的区域。取值 [ China | Asia-Pacific ]，分别表示中国大陆、亚太区域
		Billing: csn.Billing{ // 计费信息
			PaymentTiming: "Prepaid",     // 付款时间，预支付（Prepaid）和后支付（Postpaid）
			BillingMethod: "ByBandwidth", // 后付费需要传该参数。ByTraffic按量计费 ByBandwidth按带宽计费 PeakBandwidth_Percent_95 95计费 Enhanced_Percent_95增强型95计费
			Reservation: &csn.Reservation{ // 保留信息，支付方式为后支付时不需要设置，预支付时必须设置
				ReservationLength:   1,       // 时长，[1,2,3,4,5,6,7,8,9,12,24,36]
				ReservationTimeUnit: "month", // 时间单位，当前仅支持按月，取值month
			},
		},
	}
	response, err := client.GetCsnBpPrice(request)
	if err != nil {
		fmt.Printf("Failed to get csn Bp price, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func CreateCsnBp() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	interworkType := "center"
	request := &csn.CreateCsnBpRequest{
		Name:          "csn_bp_test",  // 带宽包的名称，不能为空
		InterworkType: &interworkType, // 带宽包的互通类型，取值 [ center | center-edge | edge-edge ]，分别表示云间互通、云边互通、边边互通，默认为云间互通
		Bandwidth:     10,             // 带宽包的带宽值，最大值为10000
		GeographicA:   "China",        // 网络实例所属的区域。取值 [ China | Asia-Pacific ]，分别表示中国大陆、亚太区域
		GeographicB:   "China",        // 另一个网络实例所属的区域。取值 [ China | Asia-Pacific ]，分别表示中国大陆、亚太区域
		Billing: csn.Billing{ // 计费信息
			PaymentTiming: "Prepaid", // 付款时间，预支付（Prepaid）和后支付（Postpaid）
			Reservation: &csn.Reservation{ // 保留信息，支付方式为后支付时不需要设置，预支付时必须设置
				ReservationLength:   1,       // 时长，[1,2,3,4,5,6,7,8,9,12,24,36]
				ReservationTimeUnit: "month", // 时间单位，当前仅支持按月，取值month
			},
		},
		Tags: []model.TagModel{ // 带宽包的标签列表
			{
				TagKey:   "tagKey1",
				TagValue: "tagValue1",
			},
		},
	}
	response, err := client.CreateCsnBp(request, util.NewUUID())
	if err != nil {
		fmt.Printf("Failed to create csn bp, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

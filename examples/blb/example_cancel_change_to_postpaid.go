package blbexamples

import (
	"github.com/baidubce/bce-sdk-go/services/blb"
)

func CancelChangeToPostPaid() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	blbClient, _ := blb.NewClient(ak, sk, endpoint) // 初始化client

	blbID := "Your BlbID" // 正在进行预付费转后付费的BLB实例ID

	if err := blbClient.CancelChangeToPostpaid(blbID); err != nil {
		panic(err.Error())
	}
}

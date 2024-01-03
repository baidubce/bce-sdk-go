package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
)

func GetEniStatus() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	EniId := "eni-efeq0vm3pu6e"                     // 弹性网卡ID
	response, err := ENI_CLIENT.GetEniStatus(EniId) // 获取弹性网卡状态
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

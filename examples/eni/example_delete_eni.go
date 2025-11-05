package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
)

func DeleteEni() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.DeleteEniArgs{
		EniId:       "eni-efeq0vm3pu6e", // 弹性网卡ID
		ClientToken: getClientToken(),   // 客户端Token
	}
	err := ENI_CLIENT.DeleteEni(args) // 删除弹性网卡
	if err != nil {
		panic(err)
	}
	fmt.Printf("Delete eni %s success\n", args.EniId)
}

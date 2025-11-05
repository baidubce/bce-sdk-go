package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
)

func BatchDeletePrivateIp() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.EniBatchPrivateIpArgs{
		EniId:       "eni-477g9akswgjv", // 弹性网卡ID
		ClientToken: getClientToken(),   // 客户端Token
		PrivateIpAddresses: []string{ // 弹性网卡私有IP地址列表
			"10.0.1.201",
			"10.0.1.202",
		},
	}
	err := ENI_CLIENT.BatchDeletePrivateIp(args) // 批量删除弹性网卡私有IP地址
	if err != nil {
		panic(err)
	}
	fmt.Println("BatchDeletePrivateIp success!")
}

package eniexamples

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
)

func BatchAddPrivateIp() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.EniBatchPrivateIpArgs{
		EniId:       "eni-477g9akswgjv", // 弹性网卡ID
		ClientToken: getClientToken(),   // 客户端Token
		PrivateIpAddresses: []string{ // 弹性网卡私有IP列表
			"10.0.1.201",
			"10.0.1.202",
		},
	}
	response, err := ENI_CLIENT.BatchAddPrivateIp(args) //批量添加弹性网卡私有IP
	if err != nil {
		panic(err)
	}
	r, _ := json.Marshal(response)
	fmt.Println(string(r))
}

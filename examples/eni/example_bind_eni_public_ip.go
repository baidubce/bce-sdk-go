package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
)

func BindEniPublicIp() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.BindEniPublicIpArgs{
		EniId:            "eni-477g9akswgjv", // 弹性网卡ID
		ClientToken:      getClientToken(),   // 客户端Token
		PrivateIpAddress: "10.0.1.100",       // 弹性网卡内网IP地址
		PublicIpAddress:  "120.48.142.121",   // 弹性网卡绑定的EIP地址
	}
	err := ENI_CLIENT.BindEniPublicIp(args) // 弹性网卡绑定EIP
	if err != nil {
		panic(err)
	}
	fmt.Println("Bind eni public ip success")
}

package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
	"github.com/baidubce/bce-sdk-go/util"
)

func getClientToken() string {
	return util.NewUUID()
}
func UnBindEniPublicIp() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.UnBindEniPublicIpArgs{
		EniId:           "eni-477g9akswgjv", // 弹性网卡ID
		ClientToken:     getClientToken(),   // 客户端Token
		PublicIpAddress: "120.48.142.121",   // 弹性网卡绑定的EIP
	}
	err := ENI_CLIENT.UnBindEniPublicIp(args) // 弹性网卡解绑EIP
	if err != nil {
		panic(err)
	}
	fmt.Println("UnBind eni public ip success")
}

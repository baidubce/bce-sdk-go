package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
	"github.com/baidubce/bce-sdk-go/util"
)

func getClientToken() string {
	return util.NewUUID()
}
func DeletePrivateIp() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.EniPrivateIpArgs{
		EniId:            "eni-477g9akswgjv", // 弹性网卡ID
		ClientToken:      getClientToken(),   // 客户端Token
		PrivateIpAddress: "10.0.1.108",       // 私有IP地址
	}
	err := ENI_CLIENT.DeletePrivateIp(args) // 删除弹性网卡私有IP地址
	if err != nil {
		panic(err)
	}
	fmt.Println("DeletePrivateIp success")
}

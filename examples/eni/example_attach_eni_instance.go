package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
	"github.com/baidubce/bce-sdk-go/util"
)

func getClientToken() string {
	return util.NewUUID()
}
func AttachEniInstance() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.EniInstance{
		EniId:       "eni-477g9akswgjv", // 弹性网卡ID
		ClientToken: getClientToken(),   // 客户端Token
		InstanceId:  "i-Dqf1k9ul",       // 云主机ID
	}
	err := ENI_CLIENT.AttachEniInstance(args) // 弹性网卡挂载云主机
	if err != nil {
		panic(err)
	}
	fmt.Println("AttachEniInstance success!")
}

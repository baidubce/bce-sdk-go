package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
	"github.com/baidubce/bce-sdk-go/util"
)

func AttachEniInstance() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	// 示例1: 挂载到云主机（默认）
	args := &eni.EniInstance{
		EniId:        "eni-477g9akswgjv",           // 弹性网卡ID
		ClientToken:  getClientToken(),             // 客户端Token
		InstanceId:   util.PtrString("i-Dqf1k9ul"), // 云主机ID
		InstanceType: util.PtrString("server"),     // 实例类型：server（云主机，默认）
	}
	err := ENI_CLIENT.AttachEniInstance(args) // 弹性网卡挂载云主机
	if err != nil {
		panic(err)
	}
	fmt.Println("AttachEniInstance to server success!")

	// 示例2: 挂载到HPAS实例
	args2 := &eni.EniInstance{
		EniId:        "eni-477g9akswgjv",           // 弹性网卡ID
		ClientToken:  getClientToken(),             // 客户端Token
		InstanceId:   util.PtrString("hpas-xxxxx"), // HPAS实例ID
		InstanceType: util.PtrString("hpas"),       // 实例类型：hpas（HPAS实例）
	}
	err = ENI_CLIENT.AttachEniInstance(args2) // 弹性网卡挂载HPAS实例
	if err != nil {
		panic(err)
	}
	fmt.Println("AttachEniInstance to HPAS success!")
}

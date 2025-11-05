package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
)

func UpdateEniSecurityGroup() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.UpdateEniSecurityGroupArgs{
		EniId:       "eni-477g9akswgjv", // 待更新安全组的弹性网卡ID
		ClientToken: getClientToken(),   // 客户端Token
		SecurityGroupIds: []string{ // 待更新的安全组ID列表
			"g-jpppuref4vbh",
			"g-f8u628jzeq84",
		},
	}
	err := ENI_CLIENT.UpdateEniSecurityGroup(args) // 更新弹性网卡关联的安全组
	if err != nil {
		panic(err)
	}
	fmt.Println("UpdateEniSecurityGroup success")
}

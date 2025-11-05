package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
)

func UpdateEniEnterpriseSecurityGroup() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.UpdateEniEnterpriseSecurityGroupArgs{
		EniId:       "eni-477g9akswgjv", // 待更新的弹性网卡ID
		ClientToken: getClientToken(),   // 客户端Token
		EnterpriseSecurityGroupIds: []string{ // 待更新的企业安全组列表
			"esg-1atxb1iqd1e2",
		},
	}
	err := ENI_CLIENT.UpdateEniEnterpriseSecurityGroup(args) // 更新弹性网卡关联的企业安全组
	if err != nil {
		panic(err)
	}
	fmt.Println("UpdateEniEnterpriseSecurityGroup success")
}

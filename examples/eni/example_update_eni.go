package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
	"github.com/baidubce/bce-sdk-go/util"
)

func UpdateEni() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	// 示例1: 同时更新名称和描述
	args := &eni.UpdateEniArgs{
		EniId:       "eni-477g9akswgjv",                        // 待更新的eni id
		ClientToken: getClientToken(),                          // 客户端Token
		Name:        util.PtrString("GO_SDK_TEST_UPDATE"),      // 更新后的名称（使用指针）
		Description: util.PtrString("go sdk test: update eni"), // 更新后的描述（使用指针）
	}
	err := ENI_CLIENT.UpdateEni(args) // 更新eni
	if err != nil {
		panic(err)
	}
	fmt.Printf("update eni %s success\n", args.EniId)

	// 示例2: 只更新名称，不修改描述
	args2 := &eni.UpdateEniArgs{
		EniId:       "eni-477g9akswgjv",                             // 待更新的eni id
		ClientToken: getClientToken(),                               // 客户端Token
		Name:        util.PtrString("GO_SDK_TEST_UPDATE_NAME_ONLY"), // 只更新名称
		// Description 不设置，保持原值不变
	}
	err = ENI_CLIENT.UpdateEni(args2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("update eni name only success\n")
}

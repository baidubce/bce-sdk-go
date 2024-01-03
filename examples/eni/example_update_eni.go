package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
	"github.com/baidubce/bce-sdk-go/util"
)

func getClientToken() string {
	return util.NewUUID()
}
func UpdateEni() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.UpdateEniArgs{
		EniId:       "eni-477g9akswgjv",        // 待更新的eni id
		ClientToken: getClientToken(),          // 客户端Token
		Name:        "GO_SDK_TEST_UPDATE",      // 更新后的名称
		Description: "go sdk test: update eni", // 更新后的描述
	}
	err := ENI_CLIENT.UpdateEni(args) // 更新eni
	if err != nil {
		panic(err)
	}
	fmt.Printf("update eni %s success\n", args.EniId)
}

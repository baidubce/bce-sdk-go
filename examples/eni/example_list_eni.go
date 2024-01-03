package eniexamples

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
)

func ListEni() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.ListEniArgs{
		VpcId: "vpc-jm7h2j497ut7",   // VPC ID
		Name:  "GO_SDK_TEST_CREATE", // 弹性网卡实例名称
	}
	response, err := ENI_CLIENT.ListEni(args) // 查询弹性网卡列表
	if err != nil {
		panic(err)
	}
	r, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))
}

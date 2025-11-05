package eniexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/model"

	"github.com/baidubce/bce-sdk-go/services/eni"
)

func CreateEni() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	createEniArgs := &eni.CreateEniArgs{
		ClientToken: getClientToken(),
		Name:        "GO_SDK_TEST_CREATE", // 弹性网卡名称
		SubnetId:    "sbn-d63m7t0bbwt5",   // 子网ID
		SecurityGroupIds: []string{ // 安全组ID列表
			"g-92600fd1grhr", // 安全组ID
		},
		PrivateIpSet: []eni.PrivateIp{ // 弹性网卡内网IP列表
			{
				Primary:          true,         // 弹性网卡内网IP是否为主IP
				PrivateIpAddress: "10.0.1.108", // 弹性网卡内网IP
			},
		},
		Ipv6PrivateIpSet: []eni.PrivateIp{ // 弹性网卡内网IPv6列表
			{
				Primary:          false, // 弹性网卡内网IPv6是否为主IP
				PrivateIpAddress: "",    // 弹性网卡内网IPv6
			},
		},
		Tags: []model.TagModel{
			{
				TagKey:   "tagKey",
				TagValue: "tagValue",
			},
		},
		Description:                 "go sdk test: create eni", // 弹性网卡描述
		NetworkInterfaceTrafficMode: "standard",                // 区分创建弹性RDMA网卡（ERI）和普通弹性网卡（ENI）
	}

	response, err := ENI_CLIENT.CreateEni(createEniArgs) // 创建eni
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

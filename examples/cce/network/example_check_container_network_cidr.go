package main

import (
	"encoding/json"
	"fmt"

	ccev2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
)

func CheckContainerNetworkCIDR() {
	// 设置您的ak、sk、要访问地域对应的 endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"
	ccev2Client, err := ccev2.NewClient(ak, sk, endpoint) // 初始化client
	if err != nil {
		panic(err)
	}

	args := &ccev2.CheckContainerNetworkCIDRArgs{
		VPCID:          "Your VPC ID",
		VPCCIDR:        "Your VPC CIDR",
		ContainerCIDR:  "172.28.0.0/16",
		ClusterIPCIDR:  "172.31.0.0/16",
		MaxPodsPerNode: 256,
	}

	resp, err := ccev2Client.CheckContainerNetworkCIDR(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

package main

import (
	"encoding/json"
	"fmt"

	ccev2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/types"
)

func RecommendContainerCIDR() {
	// 设置您的ak、sk、要访问地域对应的 endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"
	ccev2Client, err := ccev2.NewClient(ak, sk, endpoint) // 初始化client
	if err != nil {
		panic(err)
	}

	args := &ccev2.RecommendContainerCIDRArgs{
		ClusterMaxNodeNum: 2,
		IPVersion:         "ipv4",
		K8SVersion:        types.K8S_1_16_8,
		MaxPodsPerNode:    32,
		PrivateNetCIDRs:   []ccev2.PrivateNetString{ccev2.PrivateIPv4Net172},
		VPCCIDR:           "Your VPC CIDR",
		VPCID:             "Your VPC ID",
	}

	resp, err := ccev2Client.RecommendContainerCIDR(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

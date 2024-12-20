package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func GetIpGroupDetail() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	ipGroupID := "ipg-kzcc0bfteds6"                           // IP地址族的ID
	response, err := client.GetIpGroupDetail(ipGroupID)       // 查询指定的IP地址族
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

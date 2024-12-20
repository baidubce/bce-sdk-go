package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func GetIpSetDetail() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	ipSetID := "ips-hms1n8fu184f"                             // IP地址组的ID
	response, err := client.GetIpSetDetail(ipSetID)           // 查询指定的IP地址组
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

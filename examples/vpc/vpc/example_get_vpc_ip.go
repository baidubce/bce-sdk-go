package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ShowVpcIP() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	vpcID := "vpc-p1eawhw5rx4n"                   // vpc id
	privateIPAddresses := []string{"192.168.0.1"} // 私有ip地址

	args := &vpc.GetVpcPrivateIpArgs{
		VpcId:              vpcID,
		PrivateIpAddresses: privateIPAddresses,
	}

	response, err := client.GetPrivateIpAddressesInfo(args) // 获取vpc私有ip信息

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

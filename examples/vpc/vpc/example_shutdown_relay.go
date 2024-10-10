package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ShutdownRelay() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client
	clientToken := "ca9ab08f-55e1-4675-a55d-6939a8efe3dd"     //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性支持。
	vpcID := "vpc-p1eawhw5rx4n"                               // vpc id
	args := &vpc.UpdateVpcRelayArgs{
		ClientToken: clientToken,
		VpcId:       vpcID, // vpc id
	}
	err := client.ShutdownRelay(args)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("shutdown relay success.")
}

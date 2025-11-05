package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// main函数定义了程序入口
func ListIpreserve() {
	// 设置AK、SK和Endpoint
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // Initialize ak, sk, and endpoint

	// 创建VPC客户端
	VPC_CLIENT, _ := vpc.NewClient(ak, sk, endpoint)

	args := &vpc.ListIpeserveArgs{
		SubnetId: "sbn-4fxx51yxxxx",
		Marker:   "", // 查询的起始位置，为空则从第一条开始查询
		MaxKeys:  10,
	}

	// 添加查询保留IP范围的代码
	result, err := VPC_CLIENT.ListIpreserve(args)
	if err != nil {
		fmt.Printf("List reserved IP ranges failed with %s\n", err)
	}

	// 输出子网ID和保留IP范围信息
	for _, IpReserve := range result.IpReserves {
		fmt.Printf("IP Range: %s, Description: %s\n", IpReserve.IpCidr, IpReserve.SubnetId)
		fmt.Println("isTruncated:", result.IsTruncated)
	}
}

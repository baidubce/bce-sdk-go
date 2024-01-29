package etgatewayexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/etGateway"
)

// ListEtGatewayDetail 函数用于获取指定VPC下的ET网关列表
func ListEtGatewayDetail() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := etGateway.NewClient(ak, sk, endpoint)      // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &etGateway.ListEtGatewayArgs{
		VpcId:       "vpc-2pa2x0bjt26i",  // vpc id
		EtGatewayId: "dcgw-iiyc0ers2qx4", // et gateway id
		Name:        "test-et-gateway",   // et gateway name
		Status:      "running",           // et gateway status
		Marker:      "20",                // 起始偏移量
		MaxKeys:     1000,                // 最大返回数量
	}
	result, err := client.ListEtGateway(args)
	if err != nil {
		fmt.Println("list et gateway error: ", err)
		return
	}

	// 返回标记查询的起始位置
	fmt.Println("et gateway list marker: ", result.Marker)
	// true表示后面还有数据，false表示已经是最后一页
	fmt.Println("et gateway list isTruncated: ", result.IsTruncated)
	// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	fmt.Println("et gateway list nextMarker: ", result.NextMarker)
	// 每页包含的最大数量
	fmt.Println("et gateway list maxKeys: ", result.MaxKeys)
	// 获取et gateway列表
	for _, etGateway := range result.EtGateways {
		fmt.Println("et gateway id: ", etGateway.EtGatewayId)
		fmt.Println("et gateway name: ", etGateway.Name)
		fmt.Println("et gateway status: ", etGateway.Status)
		fmt.Println("et gateway vpc id: ", etGateway.VpcId)
		fmt.Println("et gateway create time: ", etGateway.CreateTime)
		fmt.Println("et gateway etGateway description: ", etGateway.Description)
		fmt.Println("et gateway etGateway et id: ", etGateway.EtId)
		fmt.Println("et gateway etGateway speed: ", etGateway.Speed)
		fmt.Println("et gateway etGateway channel id: ", etGateway.ChannelId)
		fmt.Println("et gateway etGateway channel local cidrs: ", etGateway.LocalCidrs)
	}
}

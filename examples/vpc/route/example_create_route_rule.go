package routeexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// main函数用于启动程序
func CreateRouteRule() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	VPC_CLIENT, _ := vpc.NewClient(ak, sk, endpoint)

	args := &vpc.CreateRouteRuleArgs{
		RouteTableId:       "rt-hf1ezaxxxxxx",
		SourceAddress:      "192.168.0.0/20",
		DestinationAddress: "0.0.0.0/0",
		// NexthopId:           "yourNexthopId",     //optional
		// NexthopType:         "nat",               //optional
		// Description:         "Your Description",  //optional
		// ClientToken:         "Your Client Token", //optional
		// NextHopList:         []NextHop, // if you want to create multiple nexthop, you can use this field.
	}

	// 创建路由规则
	result, err := VPC_CLIENT.CreateRouteRule(args)
	if err != nil {
		fmt.Println("create route rule error: ", err)
	}

	// 输出结果
	fmt.Println(result.RouteRuleId)
}

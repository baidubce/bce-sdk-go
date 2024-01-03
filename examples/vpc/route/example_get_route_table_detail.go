package routeexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// main 函数是 Go 程序的入口函数
func GetRouteTableDetail() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	VPC_CLIENT, _ := vpc.NewClient(ak, sk, endpoint)

	routeTableId := "rt-hf1ezardxxxx"
	vpcId := "vpc-nx6bs5xxxxxx"

	// routeTableId and vpcId should not be empty at the same time
	result, err := VPC_CLIENT.GetRouteTableDetail(routeTableId, vpcId)
	if err != nil {
		fmt.Println("get route table error: ", err)
	}

	// print result
	fmt.Println("result of route table id: ", result.RouteTableId)
	fmt.Println("result of vpc id: ", result.VpcId)

	for _, route := range result.RouteRules {
		fmt.Println("route rule id: ", route.RouteRuleId)
		fmt.Println("route rule routeTableId: ", route.RouteTableId)
		fmt.Println("route rule sourceAddress: ", route.SourceAddress)
		fmt.Println("route rule destinationAddress: ", route.DestinationAddress)
		fmt.Println("route rule nexthopId: ", route.NexthopId)
		fmt.Println("route rule nexthopType: ", route.NexthopType)
		fmt.Println("route rule description: ", route.Description)
	}
}

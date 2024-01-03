package routeexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// main函数是主函数，用于执行程序
func GetRouteRuleDetail() {
	// 设置AK、SK和endpoint
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"

	VPC_CLIENT, _ := vpc.NewClient(ak, sk, endpoint)

	args := &vpc.GetRouteRuleArgs{
		RouteTableId: "rt-hf1ezaxxxxxf",
		VpcId:        "vpc-nx6xxxxxxxd2",
		Marker:       "rr-j4snfssssssm", // 该参数为路由表内的某个路由规则ID，表示从这个路由规则开始取列表
		MaxKeys:      10,
	}

	// 获取路由表详情
	result, err := VPC_CLIENT.GetRouteRuleDetail(args)
	if err != nil {
		panic(err)
	}

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

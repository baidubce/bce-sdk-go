package routeexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// main 函数作为主函数
func UpdateRouteRule() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	VPC_CLIENT, _ := vpc.NewClient(ak, sk, endpoint)

	args := &vpc.UpdateRouteRuleArgs{
		RouteRuleId: "rr-1zcxxxxxxyyy",
		// SourceAddress:      "Your SourceAddress", // optional
		// DestinationAddress: "Your DestinationAddress", // optional
		// NexthopId:           "your NewNexthopId", // optional
		// Description:         "Your New Description", // optional
	}

	err := VPC_CLIENT.UpdateRouteRule(args)
	if err != nil {
		fmt.Println("Route rule updated fail")
	}

	fmt.Println("Route rule %s updated successfully", args.RouteRuleId)
}

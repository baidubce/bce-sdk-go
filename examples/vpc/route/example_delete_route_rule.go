package routeexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
	"github.com/baidubce/bce-sdk-go/util"
)

// main函数，删除路由规则
func DeleteRouteRule() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	VPC_CLIENT, _ := vpc.NewClient(ak, sk, endpoint)

	routeRuleId := "rr-1zctgrxxxxxx"
	clientToken := util.NewUUID() // "yourClientToken"

	err := VPC_CLIENT.DeleteRouteRule(routeRuleId, clientToken)
	if err != nil {
		fmt.Println("delete route rule error: ", err)
	}

	fmt.Printf("delete route rule %s success.", routeRuleId)
}

package routeexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// switchRoute
//
// PARAMS:
//   - routeRuleId: 主路由规则ID
//   - clientToken: the idempotent token, 可选
//   - action: 切换路由规则的动作，可选值：switchRouteHA
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SwitchRoute() {
	ak, sk, endpoint := "Your AK", "Your SK", "endpoint"
	client, _ := vpc.NewClient(ak, sk, endpoint) // 创建vpc client
	err := client.SwitchRoute("Your routeRuleId", "ClientToken")
	if err != nil {
		panic(err)
	}
	fmt.Print("SwitchRoute success!\n")
}

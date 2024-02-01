package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

// DeleteRouteRule - 删除云智能网路由表的指定路由条目。
//
// PARAMS:
//   - csnRtId: 路由表的ID
//   - csnRtRuleId: 路由条目的ID
//   - clientToken: 幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
//
// RETURNS:
//   - error: the return error if any occurs

func DeleteRouteRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "csn.baidubce.com"
	client, err := csn.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	if err = client.DeleteRouteRule("csnRtId", // 路由表的ID
		"csnRtRuleId", "clientToken"); err != nil {
		fmt.Printf("Failed to delete RouteRule, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully deleted RouteRule.")
}

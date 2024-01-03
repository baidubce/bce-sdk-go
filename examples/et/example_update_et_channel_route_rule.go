package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
	"github.com/baidubce/bce-sdk-go/util"
)

// getClientToken 生成一个长度为32位的随机字符串作为客户端token。
func getClientToken() string {
	return util.NewUUID()
}

func UpdateEtChannelRouteRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	CLIENT, _ := et.NewClient(ak, sk, endpoint) // 初始化client

	updateEtChannelRouteRuleArgs := &et.UpdateEtChannelRouteRuleArgs{
		ClientToken: getClientToken(),             // 幂等性Token
		EtId:        "dcphy-gq65bz9ip712",         // 专线ID
		EtChannelId: "dedicatedconn-zy9t7n91k0iq", // 专线通道ID
		RouteRuleId: "dcrr-5afcf643-94e",          // 专线通道路由规则ID
		Description: "route_2",                    // 描述信息

	}
	err := CLIENT.UpdateEtChannelRouteRule(updateEtChannelRouteRuleArgs) // 修改专线通道路由规则

	if err != nil {
		fmt.Printf("Failed to update et channel route rule, err: %+v.\n", err)
		return
	}
	fmt.Println("Successfully update et channel route rule.")
}

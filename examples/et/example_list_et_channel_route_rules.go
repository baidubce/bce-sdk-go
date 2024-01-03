package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

func ListEtChannelRouteRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	CLIENT, _ := et.NewClient(ak, sk, endpoint) // 初始化client

	listEtChannelRouteRuleArgs := &et.ListEtChannelRouteRuleArgs{
		EtId:        "dcphy-gq65bz9ip712",         // 专线ID
		EtChannelId: "dedicatedconn-zy9t7n91k0iq", // 专线通道ID
	}
	response, err := CLIENT.ListEtChannelRouteRule(listEtChannelRouteRuleArgs) // 查询专线通道路由规则

	if err != nil {
		fmt.Printf("Failed to list et channel route rule, err: %v.\n", err)
		return
	}
	fmt.Printf("Successfully list et channel route rule, response: %+v.\n", *response)
}

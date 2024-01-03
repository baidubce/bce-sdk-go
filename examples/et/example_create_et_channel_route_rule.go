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

func CreateEtChannelRouteRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	CLIENT, _ := et.NewClient(ak, sk, endpoint) // 初始化client

	createEtChannelRouteRuleArgs := &et.CreateEtChannelRouteRuleArgs{
		EtId:        "dcphy-gq65bz9ip712",         // 专线ID
		EtChannelId: "dedicatedconn-zy9t7n91k0iq", // 专线通道ID
		ClientToken: getClientToken(),             // 幂等性Token
		IpVersion:   4,                            // IP协议类型
		DestAddress: "192.168.0.7/32",             // 目标网段
		NextHopType: "etGateWay",                  // 下一跳类型
		NextHopId:   "etGateWayId",                // 下一跳实例ID
		Description: "Your description",           // 描述
	}
	response, err := CLIENT.CreateEtChannelRouteRule(createEtChannelRouteRuleArgs) // 创建专线通道路由规则

	if err != nil {
		fmt.Printf("Failed to create et channel route rule, err: %v.\n", err)
		return
	}
	fmt.Printf("Successfully create et channel route rule, response: %+v\n", *response)
}

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

func DisableEtChannelIPv6() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	CLIENT, _ := et.NewClient(ak, sk, endpoint) // 初始化client

	disableEtChannelIPv6Args := &et.DisableEtChannelIPv6Args{
		ClientToken: getClientToken(),             // 幂等性Token
		EtId:        "dcphy-gq65bz9ip712",         // 专线ID
		EtChannelId: "dedicatedconn-zy9t7n91k0iq", // 专线通道ID
	}
	err := CLIENT.DisableEtChannelIPv6(disableEtChannelIPv6Args) // 关闭专线通道IPv6功能

	if err != nil {
		fmt.Printf("Failed to enable et channel IPv6, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully disable et channel IPv6.")
}

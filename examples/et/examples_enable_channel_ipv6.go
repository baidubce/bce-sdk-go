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

// EnableEtChannelIPv6
func EnableEtChannelIPv6() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &et.EnableEtChannelIPv6Args{
		ClientToken			:	getClientToken(),       			// client token
		EtId       			:	"Your EtId",						// et id
		EtChannelId			:	"Your EtChannelId",					// et channel id
		Result				: 	et.EnableEtChannelIPv6Result{		// enable et channel ipv6 result
			BaiduIpv6Address	:   "Your BaiduIpv6Address", 		// baidu ipv6 address
			Ipv6Networks 		: 	[]string{"Your Ipv6Networks"},	// ipv6 networks
			CustomerIpv6Address	:  	"Your CustomerIpv6Address", 	// customer ipv6 address
		},
	}
	
	if err := client.EnableEtChannelIPv6(args); err != nil {
		fmt.Printf("Failed to enable et channel IPv6, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully enable et channel IPv6.")
}
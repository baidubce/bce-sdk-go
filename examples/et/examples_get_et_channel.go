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

// GetEtChannel
func GetEtChannel() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}
	args := &et.GetEtChannelArgs{
		ClientToken	: getClientToken(),    	// client token
		EtId       	: "Your Et ID",			// et id
	}
	response, err := client.GetEtChannel(args)
	if err != nil {
		fmt.Printf("Failed to get et channel, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}
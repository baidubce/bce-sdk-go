package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// GetEtChannel
func GetEtChannel() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}
	args := &et.GetEtChannelArgs{
		ClientToken: getClientToken(), // client token
		EtId:        "Your Et ID",     // et id
	}
	response, err := client.GetEtChannel(args)
	if err != nil {
		fmt.Printf("Failed to get et channel, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

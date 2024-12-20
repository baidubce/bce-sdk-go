package etexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/et"
	"github.com/baidubce/bce-sdk-go/util"
)

func getClientToken() string {
	return util.NewUUID()
}

// AssociateEtChannel
func AssociateEtChannel() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &et.AssociateEtChannelArgs{
		ClientToken:    getClientToken(),      // client token
		EtId:           "Your EtId",           // et id
		EtChannelId:    "Your EtChannelId",    // et channel id
		ExtraChannelId: "Your ExtraChannelId", // extra channel id
	}

	if err := client.AssociateEtChannel(args); err != nil {
		fmt.Printf("Failed to associate et channel, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully associate et channel.")
}

package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// DisAssociateEtChannel
func DisAssociateEtChannel() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &et.DisAssociateEtChannelArgs{
		ClientToken:    getClientToken(),      // client token
		EtId:           "Your EtId",           // et id
		EtChannelId:    "Your EtChannelId",    // et channel id
		ExtraChannelId: "Your ExtraChannelId", // extra channel id
	}

	if err := client.DisAssociateEtChannel(args); err != nil {
		fmt.Printf("Failed to disassociate et channel, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully disassociate et channel.")
}

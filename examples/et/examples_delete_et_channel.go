package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// DeleteEtChannel
func DeleteEtChannel() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}
	args := &et.DeleteEtChannelArgs{
		ClientToken: getClientToken(),   // client token
		EtId:        "Your EtId",        // et id
		EtChannelId: "Your EtChannelId", // et channel id
	}

	if err = client.DeleteEtChannel(args); err != nil {
		fmt.Printf("Failed to delete et channel, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully delete et channel.")
}

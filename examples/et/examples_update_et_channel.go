package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// UpdateEtChannel
func UpdateEtChannel() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}
	args := &et.UpdateEtChannelArgs{
		ClientToken: getClientToken(),   // client token
		EtId:        "Your EtId",        // et id
		EtChannelId: "Your EtChannelId", // et channel id
		Result: et.UpdateEtChannelResult{ // update et channel result
			Name:        "Your Name",        // et channel name
			Description: "Your Description", // et channel description
		},
	}

	if err = client.UpdateEtChannel(args); err != nil {
		fmt.Printf("Failed to update et channel, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully update et channel.")
}

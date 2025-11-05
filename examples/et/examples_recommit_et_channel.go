package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// RecommitEtChannel
func RecommitEtChannel() {
	client, err := et.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &et.RecommitEtChannelArgs{
		ClientToken: getClientToken(),   // client token
		EtId:        "Your EtId",        // et id
		EtChannelId: "Your EtChannelId", // et channel id
		Result: et.RecommitEtChannelResult{ // recommit et channel result
			AuthorizedUsers:     []string{"Your AuthorizedUsers"}, // authorized users
			Description:         "Your Description",               // description
			BaiduAddress:        "Your BaiduAddress",              // baidu address
			Name:                "Your Name",                      // name
			Networks:            []string{"Your Networks"},        // networks
			CustomerAddress:     "Your CustomerAddress",           // customer address
			RouteType:           "Your RouteType",                 // route type
			VlanId:              "Your Vlan ID",                   // vlan id
			Status:              "Your Status",                    // status
			EnableIpv6:          0,                                // enable ipv6
			BaiduIpv6Address:    "Your BaiduIpv6Address",          // baidu ipv6 address
			Ipv6Networks:        []string{"Your Ipv6Networks"},    // ipv6 networks
			CustomerIpv6Address: "Your CustomerIpv6Address",       // customer ipv6 address
		},
	}

	if err := client.RecommitEtChannel(args); err != nil {
		fmt.Printf("Failed to recommit et channel, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully recommit et channel.")
}

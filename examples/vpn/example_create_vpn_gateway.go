package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func CreateVpnGateway() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpn.CreateVpnGatewayArgs{
		VpnName:     "Your VPN Name", // vpn name
		Description: "vpn test",      // vpn description
		VpcId:       "vpcId",         // vpc id
		Billing: &vpn.Billing{ // billing info
			PaymentTiming: vpn.PAYMENT_TIMING_PREPAID,
			Reservation: &vpn.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
		ClientToken: "Your ClientToken", // client token
	}
	result, err := client.CreateVpnGateway(args)
	if err != nil {
		fmt.Printf("create vpn error: %+v\n", err)
		return
	}

	fmt.Println("create vpn success, vpn: ", result.VpnId)
}

package vpnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpn"
)

func RenewVpnGateway() {
	ak, sk, endpoint := "Your AK", "Your SK", "vpn.bj.baidubce.com"

	client, _ := vpn.NewClient(ak, sk, endpoint) // 初始化client

	args := &vpn.RenewVpnGatewayArgs{
		Billing: &vpn.Billing{ // billing info
			Reservation: &vpn.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
		ClientToken: "Your ClientToken", // client token
	}
	err := client.RenewVpnGateway(
		"vpnId", // vpn id
		args,
	)
	if err != nil {
		fmt.Printf("renew vpn error: %+v\n", err)
		return
	}

	fmt.Println("renew vpn success")
}

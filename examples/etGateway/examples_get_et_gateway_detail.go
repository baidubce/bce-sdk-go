package etgatewayexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/etGateway"
)

// GetEtGatewayDetail get et gateway detail.
func GetEtGatewayDetail() {
	client, err := etGateway.NewClient("Your AK", "Your SK", "bcc.bj.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}
	etGatewayId := "dcgw-iiyc0ers2qx4"
	response, err := client.GetEtGatewayDetail(etGatewayId)
	if err != nil {
		fmt.Printf("Failed to get et channel, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

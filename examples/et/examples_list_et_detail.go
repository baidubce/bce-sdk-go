package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// ListEtDcphyDetail
func ListEtDcphyDetail() {
	client, err := et.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	dcphyId := "Your Et Id" // 专线ID

	response, err := client.ListEtDcphyDetail(dcphyId)
	if err != nil {
		fmt.Printf("Failed to get et's detail, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

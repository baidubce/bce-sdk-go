package csnexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func ListCsnBpLimit() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	response, err := client.ListCsnBpLimit("csnBpId")
	if err != nil {
		fmt.Printf("Failed to list csn bp limit, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

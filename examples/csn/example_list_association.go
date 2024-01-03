package csnexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func ListAssociation() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	response, err := client.ListAssociation("csnRtId")
	if err != nil {
		fmt.Printf("Failed to list csn route association, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

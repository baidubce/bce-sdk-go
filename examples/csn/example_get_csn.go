package main

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func main() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	response, err := client.GetCsn("Your csnId")
	if err != nil {
		fmt.Printf("Failed to get csn, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

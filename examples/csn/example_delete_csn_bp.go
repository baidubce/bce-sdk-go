package main

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func main() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	if err = client.DeleteCsnBp("Your csnBpId", util.NewUUID()); err != nil {
		fmt.Printf("Failed to delete csn bp, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully deleted csn bp.")
}

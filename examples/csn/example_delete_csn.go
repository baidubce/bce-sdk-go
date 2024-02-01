package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func DeleteCsn() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	if err = client.DeleteCsn("Your csnId", util.NewUUID()); err != nil {
		fmt.Printf("Failed to delete csn, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully deleted csn.")
}

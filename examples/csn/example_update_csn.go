package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func UpdateCsn() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	name := "csn_test_update"
	description := "csn_test_update description"
	request := &csn.UpdateCsnRequest{
		Name:        &name,        // 云智能网的名称
		Description: &description, // 云智能网的描述
	}
	if err = client.UpdateCsn("Your csnId", request, util.NewUUID()); err != nil {
		fmt.Printf("Failed to update csn, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully updated csn.")
}

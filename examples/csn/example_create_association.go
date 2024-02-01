package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func CreateAssociation() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	description := "csn_test description"
	request := &csn.CreateAssociationRequest{
		AttachId:    "attachId",
		Description: &description,
	}
	if err = client.CreateAssociation("csnRtId", request, util.NewUUID()); err != nil {
		fmt.Printf("Failed to create csn route association, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully creat csn route association.")
}

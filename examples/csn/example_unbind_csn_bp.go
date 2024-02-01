package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func UnbindCsnBp() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	request := &csn.UnbindCsnBpRequest{
		CsnId: "csnId",
	}
	if err = client.UnbindCsnBp("csnBpId", request, util.NewUUID()); err != nil {
		fmt.Printf("Failed to unbind csn bp, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully unbind csn bp.")
}

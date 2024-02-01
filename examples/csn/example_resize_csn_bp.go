package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func ResizeCsnBp() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	request := &csn.ResizeCsnBpRequest{
		Bandwidth: 1000, // 带宽包带宽
	}
	if err = client.ResizeCsnBp("csnBpId", request, util.NewUUID()); err != nil {
		fmt.Printf("Failed to resize csn bp, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully resize csn bp.")
}

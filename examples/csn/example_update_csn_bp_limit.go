package csnexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func UpdateCsnBpLimit() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	request := &csn.CreateCsnBpLimitRequest{
		LocalRegion: "bj",
		PeerRegion:  "cn-hangzhou-cm",
		Bandwidth:   10,
	}
	if err = client.UpdateCsnBpLimit("csnBpId", request, util.NewUUID()); err != nil {
		fmt.Printf("Failed to update csn bp limit, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully update csn bp limit.")
}

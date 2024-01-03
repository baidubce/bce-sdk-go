package csnexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func CreateCsnBpLimit() {
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
	if err = client.CreateCsnBpLimit("csnBpId", request, util.NewUUID()); err != nil {
		fmt.Printf("Failed to create csn bp limit, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully create csn bp limit.")
}

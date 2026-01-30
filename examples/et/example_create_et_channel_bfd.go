package etexamples

import (
	"fmt"

	etPackage "github.com/baidubce/bce-sdk-go/services/et"
	"github.com/baidubce/bce-sdk-go/util"
)

func CreateEtChannelBfd() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := etPackage.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createEtChannelBfdRequest := &etPackage.CreateEtChannelBfdRequest{
		EtId:             util.PtrString(""),
		EtChannelId:      util.PtrString(""),
		ClientToken:      util.PtrString(""),
		SendInterval:     util.PtrInt32(int32(0)),
		ReceivInterval:   util.PtrInt32(int32(0)),
		DetectMultiplier: util.PtrInt32(int32(0)),
	}
	err = client.CreateEtChannelBfd(createEtChannelBfdRequest)
	if err != nil {
		fmt.Println("request failed:", err)
	}
}

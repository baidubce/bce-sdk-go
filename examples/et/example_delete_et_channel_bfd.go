package etexamples

import (
	"fmt"

	etPackage "github.com/baidubce/bce-sdk-go/services/et"
	"github.com/baidubce/bce-sdk-go/util"
)

func DeleteEtChannelBfd() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := etPackage.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	deleteEtChannelBfdRequest := &etPackage.DeleteEtChannelBfdRequest{
		EtId:        util.PtrString(""),
		EtChannelId: util.PtrString(""),
		ClientToken: util.PtrString(""),
	}
	err = client.DeleteEtChannelBfd(deleteEtChannelBfdRequest)
	if err != nil {
		fmt.Println("request failed:", err)
	}
}

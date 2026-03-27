package eipexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eip"
	"github.com/baidubce/bce-sdk-go/util"
)

func AcceptEipTransfer() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := eip.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	acceptEipTransferRequest := &eip.AcceptEipTransferRequest{
		ClientToken:    util.PtrString(""),
		TransferIdList: []*string{util.PtrString("tf-6unuzqor")},
	}
	err = client.AcceptEipTransfer(acceptEipTransferRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
	}
}

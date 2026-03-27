package eipexamples

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eip"
	"github.com/baidubce/bce-sdk-go/util"
)

func CreateEipTransfer() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := eip.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	createEipTransferRequest := &eip.CreateEipTransferRequest{
		ClientToken:          util.PtrString(""),
		TransferType:         util.PtrString("eip"),
		TransferResourceList: []*string{util.PtrString("ip-e98yixqq")},
		ToUserId:             util.PtrString("07782bbe2d114c4dbbbc8d584cc6aa3f"),
	}
	result := &eip.CreateEipTransferResponse{}
	result, err = client.CreateEipTransfer(createEipTransferRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		fmt.Println("json marshalIndent failed:", err)
		return
	}
	fmt.Println(string(data))
}

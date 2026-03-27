package eipexamples

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eip"
	"github.com/baidubce/bce-sdk-go/util"
)

func ListEipTransfer() {
	// 设置Client的Access Key ID和Secret Access Key，获取AKSK详见:https://cloud.baidu.com/doc/Reference/s/9jwvz2egb
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint"
	client, err := eip.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	listEipTransferRequest := &eip.ListEipTransferRequest{
		MaxKeys:           util.PtrInt32(10),
		Marker:            util.PtrString("tf-1l4m5etb"),
		Direction:         util.PtrString("sent"),
		TransferId:        util.PtrString("tf-1l4m5etb"),
		Status:            util.PtrString("success"),
		FuzzyTransferId:   util.PtrString("tf-1l4m5etb"),
		FuzzyInstanceId:   util.PtrString("ip-3pblsyay"),
		FuzzyInstanceName: util.PtrString("EIP1768387961009"),
		FuzzyInstanceIp:   util.PtrString("100.88.8.139"),
	}
	result := &eip.ListEipTransferResponse{}
	result, err = client.ListEipTransfer(listEipTransferRequest)
	if err != nil {
		// 此处仅做打印展示，请谨慎对待异常处理，在工程项目中切勿直接忽略异常。
		fmt.Println("request failed:", err)
		return
	}
	data, err := json.Marshal(result)
	if err != nil {
		fmt.Println("json marshal failed:", err)
		return
	}
	var out bytes.Buffer
	err = json.Indent(&out, data, "", "  ")
	if err != nil {
		fmt.Println("json indent failed:", err)
		return
	}
	fmt.Println(out.String())
}

package probeexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
	"github.com/baidubce/bce-sdk-go/util"
)

func DeleteProbe() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	probeId := "Your Probe ID"
	clientToken := util.NewUUID()
	err := client.DeleteProbe(probeId, clientToken) // 删除探测

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("delete probe %s success.", probeId)
}

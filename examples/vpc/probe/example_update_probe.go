package probeexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
	"github.com/baidubce/bce-sdk-go/util"
)

func UpdateProbe() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	probeId := "Your Probe ID"
	updateProbeArgs := &vpc.UpdateProbeArgs{
		Name:        util.PtrString("TestSDK1"),          // 探测名称
		Frequency:   util.PtrInt(20),                     // 探测频率
		DestIp:      util.PtrString("192.168.0.8"),       // 探测目标ip
		DestPort:    util.PtrInt(88),                     // 探测目标端口
		Payload:     util.PtrString("test udp update"),   // 探测包内容
		Description: util.PtrString("test update probe"), // 子网描述
		ClientToken: util.NewUUID(),
	}
	err := client.UpdateProbe(probeId, updateProbeArgs) // 更新探测

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("update probe %s success.", probeId)
}

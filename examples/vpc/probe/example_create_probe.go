package probeexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
	"github.com/baidubce/bce-sdk-go/util"
)

func CreateProbe() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	createProbeArgs := &vpc.CreateProbeArgs{
		Name:        "TestSDK",                // 探测名称
		VpcId:       "Your Vpc ID",            // 探测所在vpc id
		SubnetId:    "Your Subnet ID",         // 探测所在子网id
		Protocol:    "UDP",                    // 探测协议
		Frequency:   10,                       // 探测频率
		DestIp:      "192.168.0.4",            // 探测目标ip
		DestPort:    "53",                     // 探测目标端口
		SourceIps:   []string{"192.168.0.19"}, // 探测源ip列表
		SourceIpNum: 0,                        // 探测源ip数量
		Payload:     "test udp",               // 探测包内容
		Description: "test create probe",      // 探测描述
		ClientToken: util.NewUUID(),
	}
	response, err := client.CreateProbe(createProbeArgs) // 创建探测

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)
}

package probeexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func GetProbeDetail() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	probeId := "Your Probe ID"
	result, err := client.GetProbeDetail(probeId) // 查询探测详情
	if err != nil {
		fmt.Println("get probe detail error: ", err)
		return
	}

	fmt.Println("probe id: ", result.ProbeId)
	fmt.Println("probe name: ", result.Name)
	fmt.Println("probe vpcId: ", result.VpcId)
	fmt.Println("probe subnetId: ", result.SubnetId)
	fmt.Println("probe description: ", result.Description)
	fmt.Println("probe protocol: ", result.Protocol)
	fmt.Println("probe frequency: ", result.Frequency)
	fmt.Println("probe destIp: ", result.DestIp)
	fmt.Println("probe destPort: ", result.DestPort)
	fmt.Println("probe sourceIps: ", result.SourceIps)
	fmt.Println("probe sourceIpNum: ", result.SourceIpNum)
	fmt.Println("probe payload: ", result.Payload)
	fmt.Println("probe status: ", result.Status)
}

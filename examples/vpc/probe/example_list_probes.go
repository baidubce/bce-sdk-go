package probeexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ListProbes() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := vpc.NewClient(ak, sk, endpoint)              // 初始化client

	args := &vpc.ListProbesArgs{
		Marker:  "20", // 指定批量获取列表的查询的起始位置
		MaxKeys: 1000, // 指定每页包含的最大数量，最大数量不超过1000。缺省值为1000
	}
	result, err := client.ListProbes(args) // 查询探测列表
	if err != nil {
		fmt.Println("list probes error: ", err)
		return
	}

	// 返回标记查询的起始位置
	fmt.Println("probe list marker: ", result.Marker)
	// true表示后面还有数据，false表示已经是最后一页
	fmt.Println("probe list isTruncated: ", result.IsTruncated)
	// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	fmt.Println("probe list nextMarker: ", result.NextMarker)
	// 每页包含的最大数量
	fmt.Println("probe list maxKeys: ", result.MaxKeys)
	// 获取探测列表信息
	for _, probe := range result.Probes {
		fmt.Println("probe id: ", probe.ProbeId)
		fmt.Println("probe name: ", probe.Name)
		fmt.Println("probe vpcId: ", probe.VpcId)
		fmt.Println("probe subnetId: ", probe.SubnetId)
		fmt.Println("probe description: ", probe.Description)
		fmt.Println("probe protocol: ", probe.Protocol)
		fmt.Println("probe frequency: ", probe.Frequency)
		fmt.Println("probe destIp: ", probe.DestIp)
		fmt.Println("probe destPort: ", probe.DestPort)
		fmt.Println("probe sourceIps: ", probe.SourceIps)
		fmt.Println("probe sourceIpNum: ", probe.SourceIpNum)
		fmt.Println("probe payload: ", probe.Payload)
	}
}

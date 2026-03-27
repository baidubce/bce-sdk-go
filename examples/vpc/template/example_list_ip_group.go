package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ListIpGroup() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.ListIpGroupArgs{
		// 按IP版本过滤（可选），取值IPv4或IPv6
		IpVersion: "IPv4",
		// 分页起始标记（可选）
		Marker: "",
		// 每页最大数量（可选），最大1000，默认1000
		MaxKeys: 100,
	}
	result, err := client.ListIpGroup(args)
	if err != nil {
		fmt.Println("list ip group error: ", err)
		return
	}

	fmt.Println("list ip group success, total: ", len(result.IpGroups))
	for _, ipGroup := range result.IpGroups {
		fmt.Println("ip group id: ", ipGroup.IpGroupId, ", name: ", ipGroup.Name)
	}
	fmt.Println("isTruncated: ", result.IsTruncated)
	fmt.Println("nextMarker: ", result.NextMarker)
}

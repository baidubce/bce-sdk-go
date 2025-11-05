package ldexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/localDns"
)

func ListPrivateZone() {

	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := localDns.NewClient(ak, sk, endpoint)       // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}

	args := &localDns.ListPrivateZoneRequest{
		Marker:  "Get from last response", // 分页标示，从上次请求的返回值中获取，第一次请求不需要设置
		MaxKeys: 10,                       // 每页包含的最大数量
	}
	result, err := client.ListPrivateZone(args)
	if err != nil {
		fmt.Println("list private zone err:", err)
		return
	}

	fmt.Println("private zone list marker: ", result.Marker)           // 返回标记查询的起始位置
	fmt.Println("private zone list isTruncated: ", result.IsTruncated) // true表示后面还有数据，false表示已经是最后一页
	fmt.Println("private zone list nextMarker: ", result.NextMarker)   // 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	fmt.Println("private zone list maxKeys: ", result.MaxKeys)         // 每页包含的最大数量
	for _, zone := range result.Zones {                                // 获取私有域的列表信息
		fmt.Println("zoneId: ", zone.ZoneId)
		fmt.Println("zoneName: ", zone.ZoneName)
		fmt.Println("recordCount: ", zone.RecordCount)
		fmt.Println("createTime: ", zone.CreateTime)
		fmt.Println("updateTime: ", zone.UpdateTime)
	}
	fmt.Println("list private zone success")
}

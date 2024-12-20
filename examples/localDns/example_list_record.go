package ldexample

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/localDns"
	"github.com/baidubce/bce-sdk-go/util"
)

// getClientToken 生成一个长度为32位的随机字符串作为客户端token
func getClientToken() string {
	return util.NewUUID()
}

func ListRecord() {

	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := localDns.NewClient(ak, sk, endpoint)       // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}

	zoneId := "Your private zone id"
	sourceType := "" // 默认为空字符串, 可以指定类型
	marker := ""     // 首次查询传递空字符串，后续查询时需要传递上次接口返回的marker值
	maxKeys := 10    // 每次查询的最大数量
	args := &localDns.ListRecordRequest{
		SourceType: sourceType,
		Marker:     marker,
		MaxKeys:    maxKeys,
	}
	result, err := client.ListRecord(zoneId, args)
	if err != nil {
		fmt.Println("list record err:", err)
		return
	}

	fmt.Println("record list marker: ", result.Marker)           // 返回标记查询的起始位置
	fmt.Println("record list isTruncated: ", result.IsTruncated) // true表示后面还有数据，false表示已经是最后一页
	fmt.Println("record list nextMarker: ", result.NextMarker)   // 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	fmt.Println("record list maxKeys: ", result.MaxKeys)         // 每页包含的最大数量
	for _, record := range result.Records {                      // 获取对等连接的列表信息
		fmt.Println("recordId: ", record.RecordId)
		fmt.Println("rr: ", record.Rr)
		fmt.Println("value: ", record.Value)
		fmt.Println("status: ", record.Status)
		fmt.Println("type: ", record.Type)
		fmt.Println("ttl: ", record.Ttl)
		fmt.Println("priority: ", record.Priority)
		fmt.Println("description: ", record.Description)
	}
	fmt.Println("list record success")
}

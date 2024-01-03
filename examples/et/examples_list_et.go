package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
	"github.com/baidubce/bce-sdk-go/util"
)

// getClientToken 生成一个长度为32位的随机字符串作为客户端token。
func getClientToken() string {
	return util.NewUUID()
}

// ListEtDcphy
func ListEtDcphy() {
	client, err := et.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &ListEtDcphyArgs {
		Marker:     "",          // 批量获取列表的查询的起始位置
		MaxKeys:    1000,        // 每页包含的最大数量，最大数量不超过1000，缺省值为1000
		Status:     "",          // 专线状态
	}

	response, err := client.ListEtDcphy(args)
	if err != nil {
		fmt.Printf("Failed to get ets, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}
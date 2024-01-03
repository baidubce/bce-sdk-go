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

// ListEtDcphyDetail
func ListEtDcphyDetail() {
	client, err := et.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &ListEtDcphyDetailArgs {
		DcphyId:     "Your Et Id",          // 专线ID
	}

	response, err := client.ListEtDcphyDetail(args)
	if err != nil {
		fmt.Printf("Failed to get et's detail, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}
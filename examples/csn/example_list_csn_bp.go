package csnexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func ListCsnBp() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	request := &csn.ListCsnBpArgs{
		Marker:  "",   // 批量获取列表的查询的起始位置
		MaxKeys: 1000, // 每页包含的最大数量，最大数量不超过1000，缺省值为1000
	}
	response, err := client.ListCsnBp(request)
	if err != nil {
		fmt.Printf("Failed to list csn bp, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

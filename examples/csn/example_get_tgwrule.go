package csnexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

// ListTgwRule - 查询指定TGW的路由条目。
//
// PARAMS:
//     - csnId: 云智能网的ID
//     - tgwId: TGW的ID
//     - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//     - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
// RETURNS:
//     - *ListTgwRuleResponse:
//     - error: the return error if any occurs

func ListTgwRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "csn.baidubce.com"
	client, err := csn.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	args := &csn.ListTgwRuleArgs{
		MaxKeys: 1000,
	}
	// csnId: 云智能网的ID tgwId: TGW的ID
	response, err := client.ListTgwRule("csnid", "tgwid", args)
	if err != nil {
		panic(err)
	}
	if response != nil {
		fmt.Printf("%+v\n", *response)
	}
	fmt.Println("Successfully get TGW Rules.")
}

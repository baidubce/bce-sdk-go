package csnexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

// ListTgw - 查询云智能网TGW列表。
//
// PARAMS:
//     - csnId: 云智能网的ID
//     - marker: 批量获取列表的查询的起始位置，是一个由系统生成的字符串
//     - maxKeys: 每页包含的最大数量，最大数量不超过1000，缺省值为1000
// RETURNS:
//     - *ListTgwResponse:
//     - error: the return error if any occurs

func ListTgw() {
	ak, sk, endpoint := "Your AK", "Your SK", "csn.baidubce.com"
	client, err := csn.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	args := &csn.ListTgwArgs{
		Marker:  "",
		MaxKeys: 1000,
	}
	response, err := client.ListTgw("csnid", args) // 云智能网csn的ID
	if err != nil {
		fmt.Printf("Failed to list tgw, err: %v.\n", err)
		panic(err)
	}
	if response != nil {
		fmt.Printf("%+v\n", *response)
	}
}

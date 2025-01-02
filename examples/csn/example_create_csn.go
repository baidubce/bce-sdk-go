package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func CreateCsn() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	description := "csn_test description"
	request := &csn.CreateCsnRequest{
		Name:        "csn_test",   // 云智能网的名称
		Description: &description, // 云智能网的描述
		Tags: []model.TagModel{ // 云智能网的标签信息
			{
				TagKey:   "tagKey1",
				TagValue: "tagValue1",
			},
		},
	}
	response, err := client.CreateCsn(request, util.NewUUID())
	if err != nil {
		fmt.Printf("Failed to create csn, err: %v.\n", err)
		return
	}
	fmt.Printf("%+v\n", *response)
}

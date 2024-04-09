package bbc

import (
	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/bbc"
)

func UnbindReservedTagsDemo() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "ak", "sk", "bbc.bj.baidubce.com"
	client, _ := bbc.NewClient(ak, sk, endpoint) // 创建BBC Client
	args := &bbc.ReservedTagsRequest{
		ChangeTags: []model.TagModel{
			{
				TagKey:   "TagKey-go",
				TagValue: "TagValue",
			},
		},
		ReservedInstanceIds: []string{
			"r-oFpMXKhv", "r-HrztSVk0",
		},
	}
	err := client.UnbindReservedInstanceFromTags(args)
	if err != nil {
		panic(err)
	}
}

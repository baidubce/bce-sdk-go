package bcc

import (
	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"

	"github.com/baidubce/bce-sdk-go/services/bcc"
)

func BindBccReservedTagsDemo() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "ak", "sk", "bcc.bj.baidubce.com"
	client, _ := bcc.NewClient(ak, sk, endpoint) // 创建BCC Client
	args := &api.ReservedTagsRequest{
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
	err := client.BindReservedInstanceToTags(args)
	if err != nil {
		panic(err)
	}
}

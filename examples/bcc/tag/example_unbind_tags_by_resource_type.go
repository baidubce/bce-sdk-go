package tag

import (
	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"

	"github.com/baidubce/bce-sdk-go/services/bcc"
)

func UnbindTagsByResourceTypeDemo() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "ak", "sk", "bcc.bj.baidubce.com"
	client, _ := bcc.NewClient(ak, sk, endpoint) // 创建BCC Client
	args := &api.TagsOperationRequest{
		ResourceType: "bccri",
		ResourceIds: []string{
			"r-oFpMXKhv", "r-HrztSVk0",
		},
		Tags: []model.TagModel{
			{
				TagKey:   "TagKey-go",
				TagValue: "TagValue",
			},
		},
		IsRelationTag: false,
	}
	err := client.UnbindInstanceToTagsByResourceType(args)
	if err != nil {
		panic(err)
	}
}

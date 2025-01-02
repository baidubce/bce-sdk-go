package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// CreateBackupRepositorys 创建备份仓库
func CreateBackupRepositorys() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.CreateBackupRequest{
		&model.BackupRepository{
			Name:          "仓库名",
			BucketName:    "bucket名",
			BucketSubPath: "子路径",
			Region:        "地域",
		},
	}

	resp, err := client.CreateBackupRepositorys(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// DeleteBackupRepositorys 删除备份仓库
func DeleteBackupRepositorys() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.DeleteBackupRepositoryReq{
		// 备份仓库ID，如 repo-xxxxxxxx
		RepositoryID: "repo_ID",
	}

	resp, err := client.DeleteBackupRepositorys(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

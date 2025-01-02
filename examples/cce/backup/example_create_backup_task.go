package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// CreateBackupTasks 创建备份任务
func CreateBackupTasks() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.CreateBackupTaskRequest{
		ClusterID: "集群ID",
		BackupTaskArgs: &model.BackupTaskArgs{
			TaskName:             "备份名称",
			RepositoryID:         "仓库id",
			BackupScope:          "Specified",
			IncludedNamespaces:   []string{"namespace"},
			IncludedResources:    []string{"ConfigMap"},
			BackupExpirationDays: 30,
		},
	}

	resp, err := client.CreateBackupTasks(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

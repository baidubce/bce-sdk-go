package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// CreateRestoreTasks 创建恢复任务
func CreateRestoreTasks() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.CreateRestoreTaskRequest{
		ClusterID: "集群名",
		RestoreTaskArgs: &model.RestoreTaskArgs{
			TaskName:               "任务名称",
			BackupScope:            "specified",
			BackupTaskID:           "备份任务名称",
			BackupTaskName:         "备份任务名称",
			ExistingResourcePolicy: "none",
			IncludedNamespaces:     []string{"namespace"},
			RepositoryID:           "仓库名称",
		},
	}

	resp, err := client.CreateRestoreTasks(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

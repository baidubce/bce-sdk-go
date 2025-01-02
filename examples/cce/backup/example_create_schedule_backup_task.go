package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// CreateBackupScheduleRules 创建定时备份任务
func CreateBackupScheduleRules() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.CreateScheduleRulesRequest{
		ClusterID: "集群ID",
		BackupTaskArgs: &model.BackupTaskArgs{

			TaskName:             "备份名称",
			RepositoryID:         "仓库id",
			BackupScope:          "Specified",
			IncludedNamespaces:   []string{"namespace"},
			IncludedResources:    []string{"ConfigMap"},
			Schedule:             "",
			BackupExpirationDays: 1,
		},
	}

	resp, err := client.CreateBackupScheduleRules(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// ListBackupScheduleRules 获取定时备份任务
func ListBackupScheduleRules() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.ListTasksRequest{
		ClusterID: "集群ID",
		PageNo:    1,
		PageSize:  10,

		// 获取指定定时策略
		KeywordType: "taskName",
		Keyword:     "任务名称",
	}

	resp, err := client.ListBackupScheduleRules(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

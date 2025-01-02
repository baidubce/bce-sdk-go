package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// ListInstanceGroupRemedyTasks 获取节点组自愈任务列表
func ListInstanceGroupRemedyTasks() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	remedyClient, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.ListRemedyTaskOptions{
		ClusterID:       "cluster_id",
		InstanceGroupID: "instancegroup_id",
		PageSize:        10,
		PageNo:          1,

		// 时间范围，可选参数，如2024-10-21T16:00:00Z
		StartTime: "",
		StopTime:  "",
	}

	resp, err := remedyClient.ListInstanceGroupRemedyTasks(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

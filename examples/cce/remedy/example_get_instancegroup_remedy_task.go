package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// GetInstanceGroupRemedyTask 获取节点组自愈任务
func GetInstanceGroupRemedyTask() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	remedyClient, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.GetRemedyTaskOptions{
		ClusterID:       "cluster_id",
		InstanceGroupID: "instancegroup_id",
		RemedyTaskID:    "remedytask_id",
	}

	resp, err := remedyClient.GetInstanceGroupRemedyTask(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

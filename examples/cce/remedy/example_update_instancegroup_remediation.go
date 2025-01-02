package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// UpdateInstanceGroupRemediation 更新节点组自愈规则
func UpdateInstanceGroupRemediation() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	remedyClient, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.BindingOrUnBindingRequest{
		ClusterID:       "cluster_id",
		InstanceGroupID: "instancegroup_id",
		RemedyRulesBinding: &model.RemedyRulesBinding{
			RemedyRuleID:         "remedyrule_id",
			EnableCheckANDRemedy: true,
		},
	}

	resp, err := remedyClient.UpdateInstanceGroupRemediation(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

package main

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
	nrv1 "github.com/baidubce/bce-sdk-go/services/cce/v2/model/remedy/api/v1"
)

// CreateRemedyRule 创建自愈规则
func CreateRemedyRule() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	remedyClient, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.RemedyRule{
		//ClusterID: "cluster_id",

		ObjectMeta: model.ObjectMeta{
			Name: "remedy_rule_name",
		},
		Spec: &model.RemedyRuleSpec{
			Conditions: []model.RemedyCondition{
				model.RemedyCondition{
					Type:        model.ConditionType("DiskHardwareFail"),
					ConfigType:  nrv1.ConfigTypeSystemPresetConfig,
					EnableCheck: true,
				},
			},
		},
	}

	resp, err := remedyClient.CreateRemedyRule(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

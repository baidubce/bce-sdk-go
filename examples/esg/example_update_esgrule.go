package esgexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/esg"
)

// UpdateESGRule 函数用于更新ESG规则
func UpdateESGRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := esg.NewClient(ak, sk, endpoint) // 创建esg client
	args := &esg.UpdateEsgRuleArgs{
		Priority:                      900,                                  // 优先级
		Remark:                        "go sdk test update",                 // 备注
		ClientToken:                   "ClientToken",                        // token
		EnterpriseSecurityGroupRuleId: "Your EnterpriseSecurityGroupRuleId", // 企业安全组规则ID
	}
	err := client.UpdateEsgRule(args)
	if err != nil {
		panic(err)
	}
	fmt.Print("UpdateESGRule success\n")
}

package esgexample

import "github.com/baidubce/bce-sdk-go/services/esg"

// DeleteESGRule 删除指定企业安全组规则
func DeleteESGRule() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := esg.NewClient(ak, sk, endpoint) // 创建esg client
	args := &esg.DeleteEsgRuleArgs{
		EnterpriseSecurityGroupRuleId: "Your EnterpriseSecurityGroupRuleId", // 企业安全组规则ID
	}
	err := client.DeleteEsgRule(args)
	if err != nil {
		panic(err)
	}
	println("DeleteESGRule success")
}

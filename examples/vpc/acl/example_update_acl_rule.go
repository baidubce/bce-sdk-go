package aclexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UpdateAclRule() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak、sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	aclRuleId := "ar-gkf7pdiy0cuj"
	args := &vpc.UpdateAclRuleArgs{
		// 设置acl的最新协议
		Protocol: vpc.ACL_RULE_PROTOCOL_TCP,
		// 设置acl的源ip
		SourceIpAddress: "0.0.0.0",
		// 设置acl的目的ip
		DestinationIpAddress: "192.168.0.0/24",
		// 设置acl的源端口
		SourcePort: "3333",
		// 设置acl的目的端口
		DestinationPort: "4444",
		// 设置acl的优先级
		Position: 12,
		// 设置acl的策略
		Action: vpc.ACL_RULE_ACTION_ALLOW,
		// 设置acl最新的描述信息
		Description: "test",
	}
	if err := client.UpdateAclRule(aclRuleId, args); err != nil {
		fmt.Println("update acl rule err:", err)
		return
	}
	fmt.Println("update acl rule success")
}

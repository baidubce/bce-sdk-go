package main

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func CreateAclRule() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	requests := []vpc.AclRuleRequest{
		{
			// 设置acl规则所属的子网id
			SubnetId: "sbn-zuabnf2w6qtn",
			// 设置acl规则的协议
			Protocol: vpc.ACL_RULE_PROTOCOL_TCP,
			// 设置acl规则的源ip
			SourceIpAddress: "192.168.2.0",
			// 设置acl规则的目的ip
			DestinationIpAddress: "192.168.0.0/24",
			// 设置acl规则的源端口
			SourcePort: "1-65535",
			// 设置acl规则的目的端口
			DestinationPort: "443",
			// 设置acl规则的优先级
			Position: 2,
			// 设置acl规则的方向
			Direction: vpc.ACL_RULE_DIRECTION_INGRESS,
			// 设置acl规则的策略
			Action: vpc.ACL_RULE_ACTION_ALLOW,
			// 设置acl规则的描述信息
			Description: "test",
		},
	}
	args := &vpc.CreateAclRuleArgs{
		AclRules: requests,
	}
	if err := client.CreateAclRule(args); err != nil {
		fmt.Println("create acl rule err:", err)
		return
	}
	fmt.Println("create acl rule success")
}

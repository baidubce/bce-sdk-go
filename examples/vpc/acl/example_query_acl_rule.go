package main

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func QueryAclRule() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak、sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.ListAclRulesArgs{
		// 设置acl所属子网的id
		SubnetId: "sbn-zuabnf2w6qtn",
		// 设置批量获取列表的查询的起始位置
		Marker: "",
		// 设置每页包含的最大数量
		MaxKeys: 100,
	}
	result, err := client.ListAclRules(args)
	if err != nil {
		fmt.Println("query acl rule err:", err)
		return
	}
	// 返回标记查询的起始位置
	fmt.Println("acl list marker: ", result.Marker)
	// true表示后面还有数据，false表示已经是最后一页
	fmt.Println("acl list isTruncated: ", result.IsTruncated)
	// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	fmt.Println("acl list nextMarker: ", result.NextMarker)
	// 每页包含的最大数量
	fmt.Println("acl list maxKeys: ", result.MaxKeys)
	// 获取acl的列表信息
	for _, acl := range result.AclRules {
		fmt.Println("acl rule id: ", acl.Id)
		fmt.Println("acl rule subnetId: ", acl.SubnetId)
		fmt.Println("acl rule description: ", acl.Description)
		fmt.Println("acl rule protocol: ", acl.Protocol)
		fmt.Println("acl rule sourceIpAddress: ", acl.SourceIpAddress)
		fmt.Println("acl rule destinationIpAddress: ", acl.DestinationIpAddress)
		fmt.Println("acl rule sourcePort: ", acl.SourcePort)
		fmt.Println("acl rule destinationPort: ", acl.DestinationPort)
		fmt.Println("acl rule position: ", acl.Position)
		fmt.Println("acl rule direction: ", acl.Direction)
		fmt.Println("acl rule action: ", acl.Action)
	}
	fmt.Println("query acl rule success")
}

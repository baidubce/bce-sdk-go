package aclexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func QueryAcl() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak、sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	vpcId := "vpc-12345678"
	result, err := client.ListAclEntrys(vpcId)
	if err != nil {
		fmt.Println("query acl err:", err)
		return
	}
	// 查询得到acl所属的vpc id
	fmt.Println("acl entrys of vpcId: ", result.VpcId)
	// 查询得到acl所属的vpc名称
	fmt.Println("acl entrys of vpcName: ", result.VpcName)
	// 查询得到acl所属的vpc网段
	fmt.Println("acl entrys of vpcCidr: ", result.VpcCidr)
	// 查询得到acl的详细信息
	for _, acl := range result.AclEntrys {
		fmt.Println("subnetId: ", acl.SubnetId)
		fmt.Println("subnetName: ", acl.SubnetName)
		fmt.Println("subnetCidr: ", acl.SubnetCidr)
		fmt.Println("aclRules: ", acl.AclRules)
	}
	fmt.Println("query acl success")
}

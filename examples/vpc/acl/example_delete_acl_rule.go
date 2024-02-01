package aclexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
	"github.com/baidubce/bce-sdk-go/util"
)

// getClientToken 生成一个长度为32位的随机字符串作为客户端token。
func getClientToken() string {
	return util.NewUUID()
}

func DeleteAclRule() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	aclRuleId := "acl-rule-id"
	clientToken := getClientToken()
	if err := client.DeleteAclRule(aclRuleId, clientToken); err != nil {
		fmt.Println("delete acl rule err:", err)
		return
	}
	fmt.Println("delete acl rule success")
}

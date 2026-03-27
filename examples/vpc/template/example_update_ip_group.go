package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UpdateIpGroup() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	ipGroupId := "ipg-xxxxx"
	name := "new_ip_group_name"
	description := "new description"
	args := &vpc.UpdateIpGroupArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 更新IP地址族名称（可选，不传则不更新）
		Name: &name,
		// 更新IP地址族描述（可选，不传则不更新）
		Description: &description,
	}
	err = client.UpdateIpGroup(ipGroupId, args)
	if err != nil {
		fmt.Println("update ip group error: ", err)
		return
	}

	fmt.Println("update ip group success")
}

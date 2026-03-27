package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UpdateIpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	ipSetId := "ips-xxxxx"
	name := "new_ip_set_name"
	description := "new description"
	args := &vpc.UpdateIpSetArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 更新IP地址组名称（可选，不传则不更新）
		Name: &name,
		// 更新IP地址组描述（可选，不传则不更新）
		Description: &description,
	}
	err = client.UpdateIpSet(ipSetId, args)
	if err != nil {
		fmt.Println("update ip set error: ", err)
		return
	}

	fmt.Println("update ip set success")
}

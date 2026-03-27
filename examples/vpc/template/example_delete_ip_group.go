package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func DeleteIpGroup() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	ipGroupId := "ipg-xxxxx"
	args := &vpc.DeleteIpGroupArgs{
		// 请求标识
		ClientToken: getClientToken(),
	}
	err = client.DeleteIpGroup(ipGroupId, args)
	if err != nil {
		fmt.Println("delete ip group error: ", err)
		return
	}

	fmt.Println("delete ip group success")
}

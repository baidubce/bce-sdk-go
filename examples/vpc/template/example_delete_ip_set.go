package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func DeleteIpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	ipSetId := "ips-xxxxx"
	args := &vpc.DeleteIpSetArgs{
		// 请求标识
		ClientToken: getClientToken(),
	}
	err = client.DeleteIpSet(ipSetId, args)
	if err != nil {
		fmt.Println("delete ip set error: ", err)
		return
	}

	fmt.Println("delete ip set success")
}

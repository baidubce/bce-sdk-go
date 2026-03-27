package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UnbindIpSet() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	ipGroupId := "ipg-xxxxx"
	args := &vpc.UnbindIpSetArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 移除的IP地址组ID列表，单次最多5个
		IpSetIds: []string{
			"ips-xxxxx",
		},
	}
	err = client.UnbindIpSet(ipGroupId, args)
	if err != nil {
		fmt.Println("unbind ip set from ip group error: ", err)
		return
	}

	fmt.Println("unbind ip set from ip group success")
}

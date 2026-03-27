package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func DeleteIpAddress() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	ipSetId := "ips-xxxxx"
	args := &vpc.DeleteIpAddressArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 删除的IP地址列表，单次最多10个
		IpAddressInfo: []string{
			"192.168.1.0/24",
			"10.0.0.1",
		},
	}
	err = client.DeleteIpAddress(ipSetId, args)
	if err != nil {
		fmt.Println("delete ip address error: ", err)
		return
	}

	fmt.Println("delete ip address success")
}

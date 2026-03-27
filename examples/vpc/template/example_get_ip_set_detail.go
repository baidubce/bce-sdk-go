package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func GetIpSetDetail() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	ipSetId := "ips-xxxxx"
	result, err := client.GetIpSetDetail(ipSetId)
	if err != nil {
		fmt.Println("get ip set detail error: ", err)
		return
	}

	fmt.Println("ip set id: ", result.IpSetId)
	fmt.Println("ip set name: ", result.Name)
	fmt.Println("ip set description: ", result.Description)
	fmt.Println("ip set ipVersion: ", result.IpVersion)
	fmt.Println("ip set ipAddressInfo: ", result.IpAddressInfo)
	fmt.Println("ip set bindedInstances: ", result.BindedInstances)
}

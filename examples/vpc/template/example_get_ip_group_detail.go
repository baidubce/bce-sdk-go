package templateexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func GetIpGroupDetail() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	ipGroupId := "ipg-xxxxx"
	result, err := client.GetIpGroupDetail(ipGroupId)
	if err != nil {
		fmt.Println("get ip group detail error: ", err)
		return
	}

	fmt.Println("ip group id: ", result.IpGroupId)
	fmt.Println("ip group name: ", result.Name)
	fmt.Println("ip group description: ", result.Description)
	fmt.Println("ip group ipVersion: ", result.IpVersion)
	fmt.Println("ip group ipSetIds: ", result.IpSetIds)
	fmt.Println("ip group bindedInstances: ", result.BindedInstances)
}

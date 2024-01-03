package ldexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/localDns"
	"github.com/baidubce/bce-sdk-go/util"
)

// getClientToken 生成一个长度为32位的随机字符串作为客户端token
func getClientToken() string {
	return util.NewUUID()
}

func BindVpcs() {

	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := localDns.NewClient(ak, sk, endpoint)       // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}

	zoneId := "Your Zone Id"
	args := &localDns.BindVpcRequest{
		ClientToken: getClientToken(),                   // 客户端 token
		Region:      "Your vpc region",                  // region
		VpcIds:      []string{"Your vpc1", "Your vpc2"}, // 该 region 下需要绑定的 vpc 列表
	}
	if err := client.BindVpc(zoneId, args); err != nil {
		fmt.Println("bind vpcs err:", err)
		return
	}

	fmt.Println("bind vpcs success")
}

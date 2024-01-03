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

func CreatePrivateZone() {

	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := localDns.NewClient(ak, sk, endpoint)       // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}

	args := &localDns.CreatePrivateZoneRequest{
		ClientToken: getClientToken(), // 请求标识
		ZoneName:    "Your Zone Name", // 私有域名称
	}
	result, err := client.CreatePrivateZone(args)
	if err != nil {
		fmt.Println("create private zone err:", err)
		return
	}

	fmt.Println("private zone id: ", result.ZoneId) // 创建 private zone 的 id
	fmt.Println("create private zone success")
}

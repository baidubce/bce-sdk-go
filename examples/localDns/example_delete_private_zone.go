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

func DeletePrivateZone() {

	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := localDns.NewClient(ak, sk, endpoint)       // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}

	zoneId := "Your private zone id" // 创建 private zone 的 id
	clientToken := getClientToken()
	if err := client.DeletePrivateZone(zoneId, clientToken); err != nil {
		fmt.Println("delete private zone err:", err)
		return
	}

	fmt.Println("delete private zone success")
}

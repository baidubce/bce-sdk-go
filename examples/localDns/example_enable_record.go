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

func EnableRecord() {

	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := localDns.NewClient(ak, sk, endpoint)       // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}

	recordId := "Your record Id"
	clientToken := getClientToken()
	if err := client.EnableRecord(recordId, clientToken); err != nil {
		fmt.Println("enable record err:", err)
		return
	}

	fmt.Println("enable record success")
}

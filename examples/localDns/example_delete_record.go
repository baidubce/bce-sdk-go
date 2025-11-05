package ldexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/localDns"
)

func DeleteRecord() {

	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := localDns.NewClient(ak, sk, endpoint)       // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}

	recordId := "Your record Id"
	clientToken := getClientToken()
	if err := client.DeleteRecord(recordId, clientToken); err != nil {
		fmt.Println("delete record err:", err)
		return
	}

	fmt.Println("delete record success")
}

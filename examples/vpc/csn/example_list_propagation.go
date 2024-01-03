package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func ListPropagation() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := csn.NewClient(ak, sk, endpoint)              // 初始化client

	csnRtId := "xxxxx" //云智能网路由表的ID

	response, err := client.ListPropagation(csnRtId)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}

package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

func ListRouteRule() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := csn.NewClient(ak, sk, endpoint)              // 初始化client

	csnRtId := "xxxxx" //云智能网路由表的ID
	args := &csn.ListRouteRuleArgs{}

	response, err := client.ListRouteRule(csnRtId, args)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
}

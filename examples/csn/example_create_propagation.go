package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func CreatePropagation() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := csn.NewClient(ak, sk, endpoint)              // 初始化client

	csnRtId := "xxxxxx"           //云智能网路由表的ID
	clientToken := util.NewUUID() //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性
	desc := "desc"
	args := &csn.CreatePropagationRequest{
		AttachId:    "tgwAttach-uff0gvjkis95f6xg",
		Description: &desc,
	}
	err := client.CreatePropagation(csnRtId, args, clientToken) // 创建学习关系

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("create propagation success.")
}

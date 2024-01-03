package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func DeletePropagation() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, _ := csn.NewClient(ak, sk, endpoint)              // 初始化client

	csnRtId := "xxxx"             //云智能网路由表的ID
	attachId := "tgwAttach-xxxx"  //网络实例在云智能网中的身份ID
	clientToken := util.NewUUID() //幂等性Token，是一个长度不超过64位的ASCII字符串，详见ClientToken幂等性

	err := client.DeletePropagation(csnRtId, attachId, clientToken) // 删除学习关系

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("delete propagation success.")
}

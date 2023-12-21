package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func BindEip() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	args := &vpc.BindEipsArgs{
		// 设置要绑定的 EIP 列表
		Eips: []string{"180.76.186.174"}, // 替换为需要绑定的 EIP 列表
	}
	if err := natClient.BindEips(NatID, args); err != nil {
		fmt.Println("bind eips error: ", err)
		return
	}
}

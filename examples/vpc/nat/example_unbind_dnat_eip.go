package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// 以下为示例代码，实际开发中请根据需要进行修改和补充

func UnDnatBindEip() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"

	natClient, _ := vpc.NewClient(ak, sk, endpoint) // 初始化client

	NatID := "Your nat's id"

	args := &vpc.UnBindDnatEipsArgs{
		// 设置要解绑的 Dnat EIP 列表
		DnatEips: []string{"180.76.186.174"}, // 替换为需要解绑的 Dnat EIP 列表
	}
	if err := natClient.UnBindDnatEips(NatID, args); err != nil {
		fmt.Println("unbind eips error: ", err)
		return
	}
}

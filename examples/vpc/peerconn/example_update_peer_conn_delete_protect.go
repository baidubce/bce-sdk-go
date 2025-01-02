package peerconnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UpdatePeerConnDeleteProtect() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.UpdatePeerConnDeleteProtectArgs{
		// 设置对等连接删除保护，必选
		DeleteProtect: true,
		// 请求标识
		ClientToken: getClientToken(),
	}
	peerConnId := "peer-conn-id"
	if err := client.UpdatePeerConnDeleteProtect(peerConnId, args); err != nil {
		fmt.Println("update peer conn delete protect error: ", err)
		return
	}

	fmt.Printf("update peer conn delete protect for peer conn %s success", peerConnId)
}

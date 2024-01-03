package peerconnexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func UpdatePeerConn() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.UpdatePeerConnArgs{
		// 设置对等连接的接口ID 不可更改，必选
		LocalIfId:   "localIfId",
		// 设置对等连接的本端端口名称
		LocalIfName: "test-update",
		// 设置对等连接的本端端口描述
		Description: "test-description",
	}
	peerConnId := "peer-conn-id"
	if err := client.UpdatePeerConn(peerConnId, args); err != nil {
		fmt.Println("update peer conn error: ", err)
		return 
	}
	
	fmt.Printf("update peer conn %s success", peerConnId)
}

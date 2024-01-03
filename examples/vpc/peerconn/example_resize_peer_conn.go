package peerconnexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ResizePeerConn() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.ResizePeerConnArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 指定对等连接升降级的带宽
		NewBandwidthInMbps: 20,
	}
	peerConnId := "peer-conn-id"
	if err := client.ResizePeerConn(peerConnId, args); err != nil {
		fmt.Println("resize peer conn error: ", err)
		return 
	}
	
	fmt.Printf("resize peer conn %s success.", peerConnId)
}

package peerconnexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ClosePeerConnSyncDns() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.PeerConnSyncDNSArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 指定对等连接的角色，发起端"initiator" 接收端"acceptor"
		Role:        vpc.PEERCONN_ROLE_INITIATOR,
	}
	peerConnId := "peer-conn-id"
	if err := client.ClosePeerConnSyncDNS(peerConnId, args); err != nil {
		fmt.Println("close peer conn sync dns error: ", err)
		return 
	}
	
	fmt.Printf("close peer conn %s sync dns success.", peerConnId)
}

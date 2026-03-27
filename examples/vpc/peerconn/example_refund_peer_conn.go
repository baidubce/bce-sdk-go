package peerconnexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func RefundPeerConn() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	peerConnId := "peer-conn-id"
	clientToken := getClientToken()
	if err := client.RefundPeerConn(peerConnId, clientToken); err != nil {
		fmt.Println("refund peer conn error: ", err)
		return
	}
	fmt.Printf("refund peer conn %s success.", peerConnId)
}

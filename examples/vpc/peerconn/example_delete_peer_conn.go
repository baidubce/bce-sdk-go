package peerconnexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func DeletePeerConn() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	peerConnId := "peer-conn-id"
	clientToken := getClientToken()
	if err := client.DeletePeerConn(peerConnId, clientToken); err != nil {
		fmt.Println("delete peer conn error: ", err)
		return 
	}
	
	fmt.Printf("delete peer conn %s success", peerConnId)
}

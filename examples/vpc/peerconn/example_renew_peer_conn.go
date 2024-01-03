package peerconnexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func RenewPeerConn() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.RenewPeerConnArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 指定对等连接的续费信息
		Billing: &vpc.Billing{
			Reservation: &vpc.Reservation{
				ReservationLength:   1,
				ReservationTimeUnit: "month",
			},
		},
	}
	peerConnId := "peer-conn-id"
	if err := client.RenewPeerConn(peerConnId, args); err != nil {
		fmt.Println("renew peer conn error: ", err)
		return 
	}
	
	fmt.Printf("renew peer conn %s success.", peerConnId)
}

package peerconnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func GetPeerConnDetail() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	peerConnId := "peer-conn-id"
	result, err := client.GetPeerConnDetail(peerConnId, vpc.PEERCONN_ROLE_INITIATOR)
	if err != nil {
		fmt.Println("get peer conn detail error: ", err)
		return
	}

	// 查询得到对等连接的id
	fmt.Println("peerconn id: ", result.PeerConnId)
	// 查询得到对等连接的角色, "initiator"表示发起端"acceptor"表示接受端
	fmt.Println("peerconn role: ", result.Role)
	// 查询得到对等连接的状态
	fmt.Println("peerconn status: ", result.Status)
	// 查询得到对等连接的带宽
	fmt.Println("peerconn bandwithInMbp: ", result.BandwidthInMbps)
	// 查询得到对等连接的描述
	fmt.Println("peerconn description: ", result.Description)
	// 查询得到对等连接的本端接口ID
	fmt.Println("peerconn localIfId: ", result.LocalIfId)
	// 查询得到对等连接的本端接口名称
	fmt.Println("peerconn localIfName: ", result.LocalIfName)
	// 查询得到对等连接的本端VPC ID
	fmt.Println("peerconn localVpcId: ", result.LocalVpcId)
	// 查询得到对等连接的本端区域
	fmt.Println("peerconn localRegion: ", result.LocalRegion)
	// 查询得到对等连接的对端VPC ID
	fmt.Println("peerconn peerVpcId: ", result.PeerVpcId)
	// 查询得到对等连接的对端区域
	fmt.Println("peerconn peerRegion: ", result.PeerRegion)
	// 查询得到对等连接的对端账户ID
	fmt.Println("peerconn peerAccountId: ", result.PeerAccountId)
	// 查询得到对等连接的计费方式
	fmt.Println("peerconn paymentTiming: ", result.PaymentTiming)
	// 查询得到对等连接的dns状态
	fmt.Println("peerconn dnsStatus: ", result.DnsStatus)
	// 查询得到对等连接的创建时间
	fmt.Println("peerconn createdTime: ", result.CreatedTime)
	// 查询得到对等连接的过期时间
	fmt.Println("peerconn expiredTime: ", result.ExpiredTime)
	// 查询得到对等连接的标签
	fmt.Println("peerconn tags: ", result.Tags)
	// 查询得到对等连接是否开启删除保护
	fmt.Println("peerconn deleteProtect: ", result.DeleteProtect)
}

package peerconnexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

func ListPeerConn() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.ListPeerConnsArgs{
		// 指定对等连接所属的vpc id
		VpcId: "vpcId",
		// 指定批量获取列表的查询的起始位置
		Marker: "20",
		// 指定每页包含的最大数量，最大数量不超过1000。缺省值为1000
		MaxKeys: 1000,
	}
	result, err := client.ListPeerConn(args)
	if err != nil {
		fmt.Println("list peer conns error: ", err)
		return 
	}
	
	// 返回标记查询的起始位置
	fmt.Println("peerconn list marker: ", result.Marker)
	// true表示后面还有数据，false表示已经是最后一页
	fmt.Println("peerconn list isTruncated: ", result.IsTruncated)
	// 获取下一页所需要传递的marker值。当isTruncated为false时，该域不出现
	fmt.Println("peerconn list nextMarker: ", result.NextMarker)
	// 每页包含的最大数量
	fmt.Println("peerconn list maxKeys: ", result.MaxKeys)
	// 获取对等连接的列表信息
	for _, pc := range result.PeerConns {
		fmt.Println("peerconn id: ", pc.PeerConnId)
		fmt.Println("peerconn role: ", pc.Role)
		fmt.Println("peerconn status: ", pc.Status)
		fmt.Println("peerconn bandwithInMbp: ", pc.BandwidthInMbps)
		fmt.Println("peerconn description: ", pc.Description)
		fmt.Println("peerconn localIfId: ", pc.LocalIfId)
		fmt.Println("peerconn localIfName: ", pc.LocalIfName)
		fmt.Println("peerconn localVpcId: ", pc.LocalVpcId)
		fmt.Println("peerconn localRegion: ", pc.LocalRegion)
		fmt.Println("peerconn peerVpcId: ", pc.PeerVpcId)
		fmt.Println("peerconn peerRegion: ", pc.PeerRegion)
		fmt.Println("peerconn peerAccountId: ", pc.PeerAccountId)
		fmt.Println("peerconn paymentTiming: ", pc.PaymentTiming)
		fmt.Println("peerconn dnsStatus: ", pc.DnsStatus)
		fmt.Println("peerconn createdTime: ", pc.CreatedTime)
		fmt.Println("peerconn expiredTime: ", pc.ExpiredTime)
	}
}

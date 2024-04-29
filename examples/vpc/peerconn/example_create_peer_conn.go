package peerconnexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/vpc"
	"github.com/baidubce/bce-sdk-go/util"
)

func getClientToken() string {
	return util.NewUUID()
}

func CreatePeerConn() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := vpc.NewClient(ak, sk, endpoint)            // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}
	args := &vpc.CreatePeerConnArgs{
		// 请求标识
		ClientToken: getClientToken(),
		// 设置对等连接的带宽
		BandwidthInMbps: 10,
		// 设置对等连接的描述信息
		Description: "test peer conn",
		// 设置对等连接的本端端口名称
		LocalIfName: "local-interface",
		// 设置对等连接的本端vpc的id
		LocalVpcId: "local-vpc-id",
		// 设置对等连接的对端账户ID，只有在建立跨账号的对等连接时需要该字段
		PeerAccountId: "peer-account-id",
		// 设置对等连接的对端vpc的id
		PeerVpcId: "peer-vpc-id",
		// 设置对等连接的对端区域
		PeerRegion: "bj",
		// 设置对等连接的对端接口名称，只有本账号的对等连接才允许设置该字段
		PeerIfName: "peer-interface",
		// 设置对等连接绑定的资源组ID，此字段选传，传则表示绑定资源组
		ResourceGroupId: "ResourceGroupId",
		// 设置对等连接的计费信息
		Billing: &vpc.Billing{
			PaymentTiming: vpc.PAYMENT_TIMING_POSTPAID,
		},
	}
	result, err := client.CreatePeerConn(args)
	if err != nil {
		fmt.Println("create peerconn error: ", err)
		return
	}

	fmt.Println("create peerconn success, peerconn id: ", result.PeerConnId)
}

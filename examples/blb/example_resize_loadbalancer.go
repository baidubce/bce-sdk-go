package blbexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/blb"
)

func ResizeLoadBalancer() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"

	blbClient, _ := blb.NewClient(ak, sk, endpoint) // 初始化client

	blbID := "Your BlbID" // BLB实例ID

	// 性能规格：取值"small1","small2","medium1","medium2","large1","large2","large3"
	// 注意：若当前BLB实例为预付费,仅支持升级配置,不支持"unlimited"
	// 注意：若当前BLB实例为后付费-按使用量计费，还可取值"unlimited"
	performanceLevel := "large1"

	orderIdResult, err := blbClient.ResizeLoadBalancer(blbID, &blb.ResizeLoadBalancerArgs{
		PerformanceLevel: performanceLevel,
	})
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("orderId=%s\n", orderIdResult.OrderId)
}

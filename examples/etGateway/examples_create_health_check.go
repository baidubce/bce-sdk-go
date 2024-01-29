package etgatewayexamples

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/etGateway"
	"github.com/baidubce/bce-sdk-go/util"
)

// getClientToken 生成一个长度为32位的随机字符串作为客户端token。
func getClientToken() string {
	return util.NewUUID()
}

// CreateHealthCheck 函数用于创建健康检查
func CreateHealthCheck() {
	client, err := etGateway.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et gateway client, err: %v.\n", err)
		return
	}
	auto := true
	args := &etGateway.CreateHealthCheckArgs{
		ClientToken:           getClientToken(),
		EtGatewayId:           "dcgw-iiyc0ers2qx4",
		HealthCheckSourceIp:   "192.168.0.1",
		HealthCheckType:       etGateway.HEALTH_CHECK_ICMP,
		HealthCheckPort:       80,
		HealthCheckInterval:   3,
		HealthThreshold:       2,
		UnhealthThreshold:     2,
		AutoGenerateRouteRule: &auto,
	}

	err = client.CreateHealthCheck(args)
	if err != nil {
		fmt.Printf("Failed to create health check, err: %v.\n", err)
		return
	}
	fmt.Printf("Create health check success.\n")
}

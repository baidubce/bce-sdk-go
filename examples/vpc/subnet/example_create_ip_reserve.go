package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// main 函数作为 Go 语言的主函数，用于创建保留 IP。
func CreateIpreserve() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // Initialize ak, sk, and endpoint
	VPC_CLIENT, _ := vpc.NewClient(ak, sk, endpoint)          // Initialize VPC client

	args := &vpc.CreateIpreserveArgs{
		SubnetId:  "sbn-4fa15xxxxxxx", // ID of the subnet to create the reserved ip segment
		IpCidr:    "192.168.0.0/31",   // Reserved CIDR
		IpVersion: 4,                  // IP version (4 for IPv4, 6 for IPv6)
		// Description: "test",          // Description of the reserved CIDR, optional
		// ClientToken: "",              // Client token, optional
	}

	result, err := VPC_CLIENT.CreateIpreserve(args)

	if err != nil {
		fmt.Println("create reserved ip error: ", err)
		return
	}

	fmt.Println("create reserved ip success, reserved CIDR id: ", result.IpReserveId)
}

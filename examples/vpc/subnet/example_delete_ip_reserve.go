package vpcexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// main 函数用于启动程序
func DeleteIpreserve() {
	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // Initialize ak, sk, and endpoint
	VPC_CLIENT, _ := vpc.NewClient(ak, sk, endpoint)          // Initialize VPC client

	ipReserveId := "ipr-nc4xxxxx" // ID of the reserved CIDR to be deleted
	clientToken := ""             // optional yourclientToken

	err := VPC_CLIENT.DeleteIpreserve(ipReserveId, clientToken)

	if err != nil {
		fmt.Println("DeleteIpreserve error: ", err)
		return
	}

	fmt.Printf("delete reserved CIDR %s success.", ipReserveId)
}

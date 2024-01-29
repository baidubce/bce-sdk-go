package csnexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
	"github.com/baidubce/bce-sdk-go/util"
)

func DetachInstance() {
	client, err := csn.NewClient("Your AK", "Your SK", "csn.baidubce.com")
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	request := &csn.DetachInstanceRequest{
		InstanceType:      "vpc",        // 加载的实例类型，取值 [ vpc | channel | bec_vpc ]，分别表示私有网络、专线通道、边缘网络
		InstanceId:        "Your vpcId", // 加载的实例ID
		InstanceRegion:    "bj",         // 加载的实例所属的region
		InstanceAccountId: nil,          // 跨账号加载网络实例场景下，网络实例所属账号的ID
	}
	if err = client.DetachInstance("Your csnId", request, util.NewUUID()); err != nil {
		fmt.Printf("Failed to detach instance, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully detached instance.")
}

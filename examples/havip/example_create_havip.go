package havipexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/havip"
)

func CreateHaVip() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"
	HAVIP_CLIENT, _ := havip.NewClient(ak, sk, endpoint) // 初始化client

	createHaVipArgs := &havip.CreateHaVipArgs{
		PrivateIpAddress: "0.0.0.0",        // 指定的IP地址，为""表示自动分配IP地址
		Name:             "havip_sdk_test", // 高可用虚拟IP的名称
		SubnetId:         "sbn-id",         // 高可用虚拟IP所属的子网ID
	}
	response, err := HAVIP_CLIENT.CreateHaVip(createHaVipArgs) // 创建havip

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

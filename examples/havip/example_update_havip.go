package havipexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/havip"
)

func UpdateHaVip() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"
	HAVIP_CLIENT, _ := havip.NewClient(ak, sk, endpoint) // 初始化client

	updateHaVipArgs := &havip.UpdateHaVipArgs{
		HaVipId: "havip_id",   // 高可用虚拟IP的ID
		Name:    "havip_name", // 高可用虚拟IP的名称
	}
	response, err := HAVIP_CLIENT.UpdateHaVip(updateHaVipArgs) // 更新高可用虚拟IP

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

package havipexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/havip"
)

func DeleteHaVip() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"
	HAVIP_CLIENT, _ := havip.NewClient(ak, sk, endpoint) // 初始化client

	deleteHaVipArgs := &havip.DeleteHaVipArgs{
		HaVipId: "havip_id", // 高可用虚拟IP的ID
	}
	response, err := HAVIP_CLIENT.DeleteHaVip(deleteHaVipArgs) // 删除havip

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

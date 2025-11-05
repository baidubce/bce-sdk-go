package havipexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/havip"
)

func GetHaVipDetail() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"
	HAVIP_CLIENT, _ := havip.NewClient(ak, sk, endpoint) // 初始化client

	haVipId := "havip_id"                                 // 高可用虚拟IP的ID
	response, err := HAVIP_CLIENT.GetHaVipDetail(haVipId) // 查询指定的高可用虚拟IP

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

package havipexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/havip"
)

func HaVipUnbindPublicIp() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"
	HAVIP_CLIENT, _ := havip.NewClient(ak, sk, endpoint) // 初始化client

	haVipUnbindPublicIpArgs := &havip.HaVipUnbindPublicIpArgs{
		haVipId: "havip_id", // 高可用虚拟IP的ID
	}
	response, err := HAVIP_CLIENT.HaVipUnbindPublicIp(haVipUnbindPublicIpArgs) // 高可用虚拟IP解绑EIP

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

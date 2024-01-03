package havipexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/havip"
)

func HaVipBindPublicIp() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"
	HAVIP_CLIENT, _ := havip.NewClient(ak, sk, endpoint) // 初始化client

	haVipBindPublicIpArgs := &havip.HaVipBindPublicIpArgs{
		haVipId:         "havip_id", // 高可用虚拟IP的ID
		publicIpAddress: "0.0.0.0",  // 弹性公网IP的地址
	}
	response, err := HAVIP_CLIENT.HaVipBindPublicIp(haVipBindPublicIpArgs) // 高可用虚拟IP绑定EIP

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

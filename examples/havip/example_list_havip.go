package havipexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/havip"
)

func ListHaVip() {
	ak, sk, endpoint := "Your AK", "Your SK", "bcc.bj.baidubce.com"
	HAVIP_CLIENT, _ := havip.NewClient(ak, sk, endpoint) // 初始化client

	listHaVipArgs := &havip.ListHaVipArgs{
		VpcId: "vpc_id", // 高可用虚拟IP所属的VPC ID
	}
	response, err := HAVIP_CLIENT.ListHaVip(listHaVipArgs) // 查询高可用虚拟IP列表

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

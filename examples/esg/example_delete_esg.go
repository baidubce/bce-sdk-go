package esgexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/esg"
)

// DeleteEsg 函数用于删除企业安全组（ESG）。
func DeleteEsg() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := esg.NewClient(ak, sk, endpoint) // 创建esg client
	args := &esg.DeleteEsgArgs{
		EnterpriseSecurityGroupId: "Your EnterpriseSecurityGroupId", // 企业安全组的id
		ClientToken:               "ClientToken",
	}
	err := client.DeleteEsg(args) // 删除esg
	if err != nil {
		panic(err)
	}
	fmt.Print("Delete Esg Success!")
}

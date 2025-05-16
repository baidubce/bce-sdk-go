package becvpcexample

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bec"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// BEC vpc sdk
// GetVpcList - get vpc list
//
// PARAMS:
//   - args: the arguments to get vpc list
//
// RETURNS:
//   - *api.LogicPageVpcResult: the result of vpc list
//   - error: nil if success otherwise the specific error
func GetVpcList() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "endpoint"
	client, _ := bec.NewClient(ak, sk, endpoint) // 创建BEC Client
	req := &api.ListRequest{}

	response, err := client.GetVpcList(req) // 查询BEC VPC列表

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

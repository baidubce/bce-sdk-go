package becvpcexample

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bec"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// BEC vpc sdk
// UpdateVpc - update a vpc
//
// PARAMS:
//   - args: the arguments to update vpc
//
// RETURNS:
//   - *api.VpcCommonResult: the result of updated vpc
//   - error: nil if success otherwise the specific error
func UpdateVpc() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "endpoint"
	client, _ := bec.NewClient(ak, sk, endpoint) // 创建BEC Client
	req := &api.UpdateVpcRequest{
		Name:        "vpc-zyc-gosdk",
		Description: "vpc-zyc-gosdk-desc",
	}

	response, err := client.UpdateVpc("vpc-07szxxom4iu5", req) // 更新BEC VPC

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

package becvpcexample

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bec"
	"github.com/baidubce/bce-sdk-go/services/bec/api"
)

// BEC vpc sdk
// CreateVpc - create a vpc
//
// PARAMS:
//   - args: the arguments to create vpc
//
// RETURNS:
//   - *api.VpcCommonResult: the result of created vpc
//   - error: nil if success otherwise the specific error
func CreateVpc() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "endpoint"
	client, _ := bec.NewClient(ak, sk, endpoint) // 创建BEC Client
	req := &api.CreateVpcRequest{
		Name:        "vpc-gosdk-fin",
		RegionId:    "cn-nanning-cm",
		Cidr:        "192.168.33.0/24",
		Description: "gogogoexample",
		Tags: &[]api.Tag{
			api.Tag{
				TagKey:   "bec-zyc-key",
				TagValue: "bec-zyc-key-val",
			},
		},
	}

	response, err := client.CreateVpc(req) // 创建BEC VPC

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

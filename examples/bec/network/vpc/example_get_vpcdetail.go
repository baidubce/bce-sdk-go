package becvpcexample

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bec"
)

// BEC vpc sdk
// GetVpcDetail - get a vpc detail
//
// PARAMS:
//   - args: the arguments to get a vpc detail
//
// RETURNS:
//   - *api.VpcCommonResult: the result of vpc detail
//   - error: nil if success otherwise the specific error
func GetVpcDetail() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "endpoint"
	client, _ := bec.NewClient(ak, sk, endpoint) // 创建BEC Client

	response, err := client.GetVpcDetail("vpc-07szdxxm4iu5") // 获取BEC VPC 详情

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

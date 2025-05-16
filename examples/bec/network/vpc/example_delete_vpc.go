package becvpcexample

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bec"
)

// BEC vpc sdk
// DeleteVpc - delete a vpc
//
// PARAMS:
//   - args: the arguments to delete vpc
//
// RETURNS:
//   - *api.VpcCommonResult: the result of deleted vpc
//   - error: nil if success otherwise the specific error
func DeleteVpc() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "endpoint"
	client, _ := bec.NewClient(ak, sk, endpoint) // 创建BEC Client

	response, err := client.DeleteVpc("vpc-orznbxxwed71") // 删除BEC VPC

	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

package bccsgexamples

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bcc"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
)

// ListSecurityGroup - list all security group
//
// PARAMS:
//   - args: the arguments to list all security group
//
// RETURNS:
//   - *api.ListSecurityGroupResult: the result of create Instance, contains new Instance ID
//   - error: nil if success otherwise the specific error
func ListSecurityGroup() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := bcc.NewClient(ak, sk, endpoint) // 创建bcc client
	queryArgs := &api.ListSecurityGroupArgs{
		VpcId: "Your VPCID", // vpc id
	}
	result, err := client.ListSecurityGroup(queryArgs) // 查询所有的普通安全组
	if err != nil {
		panic(err)
	}
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
}

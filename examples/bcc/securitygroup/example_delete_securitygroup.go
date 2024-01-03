package bccsgexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bcc"
)

// DeleteSecurityGroup - delete a security group
//
// PARAMS:
//   - securityGroupId: the specific securityGroup ID
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DeleteSecurityGroup() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := bcc.NewClient(ak, sk, endpoint)              // 创建BCC Client
	err := client.DeleteSecurityGroup("Your SecurityGroupID") // 删除安全组
	if err != nil {
		panic(err)
	}
	fmt.Print("Delete SecurityGroup Success!")
}

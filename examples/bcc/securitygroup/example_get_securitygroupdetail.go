package bccsgexamples

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/bcc"
)

// GetSecurityGroupDetail 获取安全组详情
func GetSecurityGroupDetail() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := bcc.NewClient(ak, sk, endpoint)                   // 创建bcc client
	result, err := client.GetSecurityGroupDetail("g-7dyv27r6pnse") // 查询普通安全组详情
	if err != nil {
		panic(err)
	}
	r, _ := json.Marshal(result)
	fmt.Println(string(r))
}

package esgexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/esg"
)

// GetEsgList 函数用于查询企业安全组列表
func GetEsgList() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := esg.NewClient(ak, sk, endpoint) // 创建esg client
	args := &esg.ListEsgArgs{
		InstanceId: "instanceID", // 实例ID, 可查询实例关联的企业安全组列表。如需查询所有创建的企业安全组信息，不填写此参数
		Marker:     "",           // 批量获取列表的查询的起始位置，是一个由系统生成的字符串
		MaxKeys:    10,           // 每页包含的最大数量，最大数量通常不超过1000。缺省值为1000
	}
	response, err := client.ListEsg(args) // 查询企业安全组列表
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

package eniexamples

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/eni"
	"github.com/baidubce/bce-sdk-go/util"
)

func getClientToken() string {
	return util.NewUUID()
}
func AddPrivateIp() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	args := &eni.EniPrivateIpArgs{
		EniId:            "eni-477g9akswgjv", // 弹性网卡ID
		ClientToken:      getClientToken(),   // 客户端Token
		PrivateIpAddress: "10.0.1.108",       // 私有IP地址
	}
	response, err := ENI_CLIENT.AddPrivateIp(args) // 添加私有IP地址
	if err != nil {
		panic(err)
	}
	r, err := json.Marshal(response)
	fmt.Println(string(r))
}

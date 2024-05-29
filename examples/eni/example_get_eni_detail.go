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
func GetEniDetail() {
	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	ENI_CLIENT, _ := eni.NewClient(ak, sk, endpoint) // 初始化client

	response, err := ENI_CLIENT.GetEniDetail("eni-477g9akswgjv") // 查询指定的弹性网卡
	if err != nil {
		panic(err)
	}
	r, err := json.Marshal(response)
	fmt.Println(string(r))

}

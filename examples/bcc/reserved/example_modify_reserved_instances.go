package reserved

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bcc"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"
)

func ModifyReservedInstances() {

	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "ak", "sk", "bcc.bj.baidubce.com"
	// 创建BCC Client
	client, _ := bcc.NewClient(ak, sk, endpoint)

	args := &api.ModifyReservedInstancesArgs{
		ReservedInstances: []api.ModifyReservedInstance{
			{
				ReservedInstanceId:   "r-UBVQFB5b",
				ReservedInstanceName: "update-reserved-instance",
			},
		},
	}

	result, err := client.ModifyReservedInstances(args)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

package reserved

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/bcc/api"

	"github.com/baidubce/bce-sdk-go/services/bcc"
)

func CreateReservedInstance() {

	// 初始化AK/SK/Endpoint
	ak, sk, endpoint := "ak", "sk", "bcc.bj.baidubce.com"
	// 创建BCC Client
	client, _ := bcc.NewClient(ak, sk, endpoint)

	args := &api.CreateReservedInstanceArgs{
		ClientToken:              "myClientToken",
		ReservedInstanceName:     "myReservedInstance",
		Scope:                    "AZ",
		ZoneName:                 "cn-bj-a",
		Spec:                     "bcc.g5.c2m8",
		OfferingType:             "FullyPrepay",
		InstanceCount:            1,
		ReservedInstanceCount:    1,
		ReservedInstanceTime:     365,
		ReservedInstanceTimeUnit: "month",
		AutoRenewTimeUnit:        "month",
		AutoRenewTime:            1,
		AutoRenew:                true,
		Tags: []api.Tag{
			{
				TagKey:   "Env",
				TagValue: "Production",
			},
		},
		EhcClusterId: "ehcCluster123",
		TicketId:     "ticket456",
	}

	result, err := client.CreateReservedInstance(args)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

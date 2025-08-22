package v2

import (
	"encoding/json"
	"fmt"

	aihcv2 "github.com/baidubce/bce-sdk-go/services/aihc/v2"
	v2 "github.com/baidubce/bce-sdk-go/services/aihc/v2/api"
)

func DescribeResourcePools() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	client, _ := aihcv2.NewClient(ak, sk, endpoint)
	result, err := client.DescribeResourcePools(v2.DescribeResourcePoolsRequest{
		PageNumber:       2,
		PageSize:         1,
		KeywordType:      v2.ResourcePoolKeywordTypeResourcePoolName,
		OrderBy:          v2.ResorucePoolOrderByCreatedAt,
		ResourcePoolType: v2.ResourcePoolTypeCommon,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonBytes))
}

func DescribeResourcePool() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"
	resourcePoolID := "cce-xxxx"

	client, _ := aihcv2.NewClient(ak, sk, endpoint)
	result, err := client.DescribeResourcePool(resourcePoolID)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonBytes))
}

func DescribeResourcePoolOverview() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	client, _ := aihcv2.NewClient(ak, sk, endpoint)
	result, err := client.DescribeResourcePoolOverview()

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonBytes))
}

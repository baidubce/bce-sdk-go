package v2

import (
	"encoding/json"
	"fmt"

	aihcv2 "github.com/baidubce/bce-sdk-go/services/aihc/v2"
	v2 "github.com/baidubce/bce-sdk-go/services/aihc/v2/api"
)

func DescribeResourceQueues() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"

	client, _ := aihcv2.NewClient(ak, sk, endpoint)
	result, err := client.DescribeResourceQueues(&v2.DescribeResourceQueuesRequest{
		PageNumber:     2,
		PageSize:       1,
		ResourcePoolID: "cce-234afznf",
		Keyword:        "test",
		KeywordType:    "queueName",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonBytes))
}

func DescribeResourceQueue() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"
	resourceQueueID := "queue-xvhldim9m0yq"

	client, _ := aihcv2.NewClient(ak, sk, endpoint)
	result, err := client.DescribeResourceQueue(resourceQueueID)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonBytes))
}

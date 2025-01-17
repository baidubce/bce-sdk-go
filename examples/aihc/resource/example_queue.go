package resource

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/aihc"
	v1 "github.com/baidubce/bce-sdk-go/services/aihc/api/v1"
)

func ListQueue() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"
	resourcePoolID := ""

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.ListQueue(resourcePoolID, &v1.ListQueueRequest{
		PageNo:   1,
		PageSize: 1,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func GetQueue() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"
	resourcePoolID, queueName := "", ""

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.GetQueue(resourcePoolID, queueName)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

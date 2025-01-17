package resource

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/aihc"
	v1 "github.com/baidubce/bce-sdk-go/services/aihc/api/v1"
)

func ListResourcePool() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.ListResourcePool(&v1.ListResourcePoolRequest{
		PageNo:   1,
		PageSize: 1,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func GetResourcePool() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"
	resourcePoolID := ""

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.GetResourcePool(resourcePoolID)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListResourcePoolNode() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"
	resourcePoolID := ""

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.ListNodeByResourcePoolID(resourcePoolID, &v1.ListResourcePoolNodeRequest{
		PageNo:   1,
		PageSize: 2,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

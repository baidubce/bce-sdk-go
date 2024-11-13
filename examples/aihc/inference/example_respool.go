package inference

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/aihc/inference"
	"github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
)

func ListBriefResPool() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.ListBriefResPool(&api.ListBriefResPoolArgs{
		PageNo:   1,
		PageSize: 10,
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ResPoolDetail() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.ResPoolDetail(&api.ResPoolDetailArgs{
		ResPoolId: "",
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

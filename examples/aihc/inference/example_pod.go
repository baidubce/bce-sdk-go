package inference

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/aihc/inference"
	"github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
)

func ListPod() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.ListPod(&api.ListPodArgs{
		AppId: "",
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func BlockPod() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.BlockPod(&api.BlockPodArgs{
		AppId: "",
		InsID: "",
		Block: true,
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func DeletePod() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.DeletePod(&api.DeletePodArgs{
		AppId: "",
		InsID: "",
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

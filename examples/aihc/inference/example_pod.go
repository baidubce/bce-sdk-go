package inference

import (
	"encoding/json"
	"fmt"

	api "github.com/baidubce/bce-sdk-go/services/aihc/inference/v2"
)

func ListPod() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.ListPod(&api.ListPodArgs{
		ServiceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func BlockPod() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.BlockPod(&api.BlockPodArgs{
		ServiceId:  "",
		InstanceId: "",
		Block:      true,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func DeletePod() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.DeletePod(&api.DeletePodArgs{
		ServiceId:  "",
		InstanceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListPodGroups() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.ListPodGroups(&api.ListPodGroupsArgs{
		ServiceId: "",
	})
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

package inference

import (
	"encoding/json"
	"fmt"

	api "github.com/baidubce/bce-sdk-go/services/aihc/inference/v2"
)

func CreateService() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"
	client, _ := api.NewClient(ak, sk, endpoint)
	createServiceArgs, err := ReadJson("you_create_json_file")
	if err != nil {
		panic(err)
		return
	}
	result, err := client.CreateService(createServiceArgs, "your clientToken")

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListService() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"
	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.ListService(&api.ListServiceArgs{
		PageSize:   10,
		PageNumber: 1,
	})
	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListServiceStats() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.ListServiceStats(&api.ListServiceStatsArgs{
		ServiceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ServiceDetails() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.ServiceDetails(&api.ServiceDetailsArgs{
		ServiceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func UpdateService() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	updateServiceArgs, err := ReadJson("you_update_json_file")
	if err != nil {
		panic(err)
		return
	}
	result, err := client.UpdateService(&api.UpdateServiceArgs{
		ServiceId:   "",
		ServiceConf: *updateServiceArgs,
		Description: "your description",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ScaleService() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.ScaleService(&api.ScaleServiceArgs{
		ServiceId:     "",
		InstanceCount: 1,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func PubAccess() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.PubAccess(&api.PubAccessArgs{
		ServiceId:    "",
		PublicAccess: false,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListChange() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.ListChange(&api.ListChangeArgs{
		ServiceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ChangeDetail() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.ChangeDetail(&api.ChangeDetailArgs{
		ChangeId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func DeleteService() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.DeleteService(&api.DeleteServiceArgs{
		ServiceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

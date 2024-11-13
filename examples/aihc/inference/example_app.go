package inference

import (
	"encoding/json"
	"fmt"
	"github.com/baidubce/bce-sdk-go/services/aihc/inference"
	"github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
)

func CreateApp() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.CreateApp(&api.CreateAppArgs{
		AppName: "",
	}, region, nil)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListApp() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.ListApp(&api.ListAppArgs{
		PageSize: 10,
		PageNo:   1,
	}, region, nil)
	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListAppStats() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.ListAppStats(&api.ListAppStatsArgs{
		AppIds: []string{""},
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func AppDetails() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.AppDetails(&api.AppDetailsArgs{
		AppId: "",
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func UpdateApp() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.UpdateApp(&api.UpdateAppArgs{}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ScaleApp() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.ScaleApp(&api.ScaleAppArgs{
		AppId:    "",
		InsCount: 1,
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func PubAccess() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.PubAccess(&api.PubAccessArgs{
		AppId:        "",
		PublicAccess: false,
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListChange() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.ListChange(&api.ListChangeArgs{
		AppId: "",
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ChangeDetail() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.ChangeDetail(&api.ChangeDetailArgs{
		ChangeId: "",
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func DeleteApp() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.baidubce.com"
	region := "bj"

	client, _ := inference.NewClient(ak, sk, endpoint)
	result, err := client.DeleteApp(&api.DeleteAppArgs{
		AppId: "",
	}, region)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

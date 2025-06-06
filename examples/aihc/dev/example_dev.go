package dev

import (
	"encoding/json"
	"fmt"

	api "github.com/baidubce/bce-sdk-go/services/aihc/dev"
)

func CreateDevInstance() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"
	client, _ := api.NewClient(ak, sk, endpoint)
	args, err := ReadJson("you_create_json_file")
	if err != nil {
		panic(err)
	}
	result, err := client.CreateDevInstance(args)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func UpdateDevInstance() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"
	client, _ := api.NewClient(ak, sk, endpoint)
	args, err := ReadJson("you_update_json_file")
	if err != nil {
		panic(err)
	}
	result, err := client.UpdateDevInstance(args)

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func QueryDevInstanceDetail() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.QueryDevInstanceDetail(&api.QueryDevInstanceDetailArgs{
		DevInstanceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListDevInstance() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.ListDevInstance(&api.ListDevInstanceArgs{
		QueryKey: "devInstanceId",
		QueryVal: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func StopDevInstance() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.StopDevInstance(&api.StopDevInstanceArgs{
		DevInstanceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func TestStartDevInstance() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.StartDevInstance(&api.StartDevInstanceArgs{
		DevInstanceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func DeleteDevInstance() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.DeleteDevInstance(&api.DeleteDevInstanceArgs{
		DevInstanceId: "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func TimedStopDevInstance() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.TimedStopDevInstance(&api.TimedStopDevInstanceArgs{
		DevInstanceId: "",
		DelaySec:      3600,
		Enable:        true,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func CreateDevInstanceImagePackJob() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)

	result, err := client.CreateDevInstanceImagePackJob(&api.CreateDevInstanceImagePackJobArgs{
		DevInstanceID: "",
		ImageName:     "",
		ImageTag:      "",
		Namespace:     "",
		Password:      "",
		Registry:      "",
		Username:      "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func DevInstanceImagePackJobDetail() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)
	result, err := client.DevInstanceImagePackJobDetail(&api.DevInstanceImagePackJobDetailArgs{
		ImagePackJobId: "",
		DevInstanceId:  "",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListDevInstanceEvent() {
	ak, sk, endpoint := "Your ak", "Your sk", "aihc.bj.baidubce.com"

	client, _ := api.NewClient(ak, sk, endpoint)

	result, err := client.ListDevInstanceEvent(&api.ListDevInstanceEventArgs{
		DevInstanceId: "",
		StartTime:     "2025-05-18T17:12:20.761Z",
		EndTime:       "2025-06-4T05:30:23.337Z",
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

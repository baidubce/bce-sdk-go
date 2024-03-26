package main

import (
	"encoding/json"
	"fmt"

	ccev2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
)

func UpdateAddons() {
	// 设置您的ak、sk、要访问地域对应的 endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"
	ccev2Client, err := ccev2.NewClient(ak, sk, endpoint) // 初始化client
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	args := &ccev2.UpdateAddonArgs{
		Name:              "addon-name",
		ClusterID:         "cluster-id",
		AddOnInstanceName: "addon-instance-name",
		Params:            "EnableHook: true\\nEnableSGPU: true\\n\\n",
	}

	resp, err := ccev2Client.UpdateAddon(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

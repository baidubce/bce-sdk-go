package main

import (
	"encoding/json"
	"fmt"

	ccev2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
)

func UninstallAddons() {
	// 设置您的ak、sk、要访问地域对应的 endpoint
	ak, sk, endpoint := "Your AK", "Your SK", "Your endpoint"
	ccev2Client, err := ccev2.NewClient(ak, sk, endpoint) // 初始化client
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	args := &ccev2.UninstallAddonArgs{
		Name:         "addon-name",
		ClusterID:    "cluster-id",
		InstanceName: "instance-name",
	}

	resp, err := ccev2Client.UnInstallAddon(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

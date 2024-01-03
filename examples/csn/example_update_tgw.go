package csnexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/csn"
)

// UpdateTgw - 更新TGW的名称、描述。
//
// PARAMS:
//     - csnId: 云智能网的ID
//     - tgwId: TGW实例的ID
//     - body: body参数
// RETURNS:
//     - error: the return error if any occurs

func UpdateTgw() {
	ak, sk, endpoint := "Your AK", "Your SK", "csn.baidubce.com"
	client, err := csn.NewClient(ak, sk, endpoint)
	if err != nil {
		fmt.Printf("Failed to new csn client, err: %v.\n", err)
		return
	}
	name := "tgw_1"
	desc := "desc"
	args := &csn.UpdateTgwRequest{
		Name:        &name, // TGW名称
		Description: &desc, // TGW描述
	}
	// csnId: 云智能网的ID tgwId: TGW的ID
	if err := client.UpdateTgw("csnId", "tgwId", args, "clientToken"); err != nil {
		fmt.Printf("Failed to update TGW, err: %v.\n", err)
		panic(err)
	}
	fmt.Println("Successfully updated TGW.")
}

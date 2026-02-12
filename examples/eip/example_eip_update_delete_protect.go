package eipexamples

import (
	"fmt"
	EIP "github.com/baidubce/bce-sdk-go/services/eip"
)

func UpdateEipDeleteProtect() {
	Init()

	eip := "x.x.x.x"
	deleteProtect := true
	args := &EIP.UpdateEipDeleteProtectArgs{
		DeleteProtect: &deleteProtect,
		ClientToken:   "",
	}

	err := eipClient.UpdateEipDeleteProtect(eip, args)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update eip deleteProtect success!")
}

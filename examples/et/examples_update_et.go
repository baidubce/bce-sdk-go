package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
)

// UpdateEtDcphy
func UpdateEtDcphy() {
	client, err := et.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &et.UpdateEtDcphyArgs{
		Name:        "Your Et name",        // 专线名称
		Description: "Your Et description", // 描述
		UserName:    "Your name",           // 用户名称
		UserPhone:   "Your Phone",          // 用户手机号码
		UserEmail:   "Your Email",          // 用户邮箱
	}

	if err = client.UpdateEtDcphy("Your Et Id", args); err != nil {
		fmt.Printf("Failed to update et, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully update et.")
}

package etexamples

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/et"
	"github.com/baidubce/bce-sdk-go/util"
)

// getClientToken 生成一个长度为32位的随机字符串作为客户端token。
func getClientToken() string {
	return util.NewUUID()
}

// CreateEtDcphy
func CreateEtDcphy() {
	client, err := et.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &CreateEtDcphyArgs {
		ClientToken: getClientToken(),            // 请求标识
		Name:        "Your Et name",              // 专线名称
		Description: "Your Et description"        // 描述
		Isp:         "ISP_CMCC",                  // 运营商
		IntfType:    "1G",                        // 物理端口规格
		ApType:      "SINGLE",                    // 线路类型
		ApAddr:      "BJYZ",                      // 接入点
		UserName:    "Your name",                 // 用户名称
		UserPhone:   "Your Phone",                // 用户手机号码
		UserEmail:   "Your Email",                // 用户邮箱
		UserIdc:     "Your Idc",                  // 对端地址
		Tags:        []Tag{{"Your TagKey", "Your TagValue"}}, // 标签
	}

	if err = client.CreateEtDcphy(args); err != nil {
		fmt.Printf("Failed to create a new et, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully create a new et.")
}
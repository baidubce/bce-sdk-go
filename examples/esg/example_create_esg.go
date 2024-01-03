package esgexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/esg"
)

// CreateESG函数用于创建企业安全组
func CreateESG() {
	ak, sk, endpoint := "Your AK", "Your SK", "Your Endpoint"
	client, _ := esg.NewClient(ak, sk, endpoint) // 创建esg client
	args := &esg.CreateEsgArgs{
		Name: "esgGoSdkTest",
		Rules: []esg.EnterpriseSecurityGroupRule{ // 企业安全组规则
			{
				Action:    "deny",        // 	规则“允许”或者“拒绝”，取值‘allow’或者‘deny’,必填
				Direction: "ingress",     //	指定方向，“出向”或者“入向”，取值“egress”或者“ingress”，必填
				Ethertype: "IPv4",        //	指定IP协议类型v4/v6
				PortRange: "1-65535",     //	指定端口范围 1-65535
				Priority:  1000,          //	指定优先级 1-1000
				Protocol:  "udp",         //	指定协议，取值“udp”/"tcp"/"icmp"/"all"，必填
				Remark:    "go sdk test", //	指定备注
				SourceIp:  "all",         //	源（目的）IP地址，"all"代表全部地址，必填
			},
			{
				Action:    "allow",
				Direction: "ingress",
				Ethertype: "IPv4",
				PortRange: "1-65535",
				Priority:  1000,
				Protocol:  "icmp",
				Remark:    "go sdk test",
				SourceIp:  "all",
			},
		},
		Desc:        "go sdk test", // 描述
		ClientToken: "ClientToken",
		Tags: []model.TagModel{ // 标签
			{
				TagKey:   "test",
				TagValue: "",
			},
		},
	}
	response, err := client.CreateEsg(args) // 创建企业安全组
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}

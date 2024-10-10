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

// CreateEtChannel
func CreateEtChannel() {
	client, err := et.NewClient("Your AK", "Your SK", "Your endpoint") // 初始化ak、sk和endpoint
	if err != nil {
		fmt.Printf("Failed to new et client, err: %v.\n", err)
		return
	}

	args := &CreateEtChannelArgs{
		ClientToken:         getClientToken(),
		EtId:                "Your EtId",                             // 专线ID
		AuthorizedUsers:     []string{"Your AuthorizedUsers"},        // 分配对象
		Description:         "Your Description",                      // 描述
		BaiduAddress:        "Your BaiduAddress",                     // 云端网络互联IP
		Name:                "Your Channel Name",                     // 通道名称
		Networks:            []string{"Your Networks"},               // 路由参数
		CustomerAddress:     "Your CustomerAddress",                  // IDC互联IP
		RouteType:           "Your RouteType",                        // 路由协议
		VlanId:              "Your Vlan ID",                          // VLAN ID
		BgpAsn:              "Your Bgp Asn",                          // BGP ASN
		BgpKey:              "Your Bgp Key",                          // BGP 密钥
		EnableIpv6:          0,                                       // IPv6功能是否开启
		BaiduIpv6Address:    "Your BaiduIpv6Address",                 // 云端网络侧IPv6互联地址
		CustomerIpv6Address: "Your CustomerIpv6Address",              // IDC侧IPv6互联地址
		Ipv6Networks:        []string{"Your Ipv6Networks"},           // IPv6路由参数
		Tags:                []Tag{{"Your TagKey", "Your TagValue"}}, // 标签
	}

	if err = client.CreateEtChannel(args); err != nil {
		fmt.Printf("Failed to create et channel, err: %v.\n", err)
		return
	}
	fmt.Println("Successfully create et channel.")
}

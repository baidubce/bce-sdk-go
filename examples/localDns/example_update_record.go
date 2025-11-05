package ldexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/localDns"
)

func UpdateRecord() {

	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := localDns.NewClient(ak, sk, endpoint)       // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}

	recordId := "Your record Id"
	args := &localDns.UpdateRecordRequest{
		ClientToken: getClientToken(),                         // 幂等性 token
		Rr:          "Your record name, such as www.test.com", // 记录名称
		Value:       "Your record value, such as 192.168.0.1", // 记录值
		Type:        "Your record type, such as A",            // 解析记录类型，目前支持A, AAAA,CNAME, TXT, MX, PTR, SRV
		Ttl:         60,                                       // 生存时间，值为[5,24*3600]，默认为60
		Priority:    0,                                        // MX记录的优先级，取值范围：[0,100]。记录类型为MX记录时，此参数必选。
		Description: "You record description",                 // 描述
	}
	if err := client.UpdateRecord(recordId, args); err != nil {
		fmt.Println("update record err:", err)
		return
	}

	fmt.Println("udpate record success")
}

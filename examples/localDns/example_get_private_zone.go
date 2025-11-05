package ldexample

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/localDns"
)

func GetPrivateZone() {

	ak, sk, endpoint := "Your Ak", "Your Sk", "Your endpoint" // 初始化ak, sk和endpoint
	client, err := localDns.NewClient(ak, sk, endpoint)       // 初始化client
	if err != nil {
		fmt.Println("create client err:", err)
		return
	}

	zoneId := "Your private zone id"
	result, err := client.GetPrivateZone(zoneId)
	if err != nil {
		fmt.Println("get private zone err:", err)
		return
	}

	fmt.Println("private zone zoneId: ", result.ZoneId)
	fmt.Println("private zone zoneName: ", result.ZoneName)
	fmt.Println("private zone recordCount: ", result.RecordCount)
	fmt.Println("private zone createTime: ", result.CreateTime)
	fmt.Println("private zone updateTime: ", result.UpdateTime)
	for _, vpc := range result.BindVpcs {
		fmt.Println("vpcId: ", vpc.VpcId)
		fmt.Println("vpcName: ", vpc.VpcName)
		fmt.Println("vpcRegion: ", vpc.VpcRegion)
	}
	fmt.Println("get private zone success")
}

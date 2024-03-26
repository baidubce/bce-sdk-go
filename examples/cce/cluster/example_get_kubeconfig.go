package cluster

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
)

func GetKubeConfig() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""
			
	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	ccev2Client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}
	args := &v2.GetKubeConfigArgs{
		ClusterID: "your-cluster-id",
		KubeConfigType: v2.KubeConfigTypeVPC, // kubeconfig-type you need
	}
	resp, err := ccev2Client.GetKubeConfig(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

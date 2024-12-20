package cluster

import (
	"encoding/json"
	"fmt"

	ccev2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
)

func UpdateClusterForbidDelete() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	ccev2Client, err := ccev2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}
	clusterID := ""
	args := &ccev2.UpdateClusterForbidDeleteArgs{
		ClusterID: clusterID,
		UpdateClusterForbidDeleteRequest: ccev2.UpdateClusterForbidDeleteRequest{
			ForbidDelete: true,
		},
	}

	resp, err := ccev2Client.UpdateClusterForbidDelete(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

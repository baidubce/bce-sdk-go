package rbac

import (
	"encoding/json"
	"fmt"
	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

func RenewRBACForNamespace() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	ccev2Client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.RBACRequest{
		ClusterID: "集群 ID", // 必须指定具体的 clusterID
		UserID:    "用户 ID",
		Namespace: "已授权的 namespace",   // 该值需跟授权时保持一致，如果授权时是 all，则重置时必须也是 all。
		Role:      model.RoleReadonly, // 该值需跟授权时保持一致
	}

	resp, err := ccev2Client.RenewRBAC(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

func RenewRBACForAllNamespace() {
	// 用户的Access Key ID和Secret Access Key
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	ccev2Client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.RBACRequest{
		ClusterID: "集群 ID", // 必须指定具体的 clusterID
		UserID:    "用户 ID",
		Namespace: model.AllNamespace, // 该值需跟授权时保持一致，如果授权时是 all，则重置时必须也是 all。
		Role:      model.RoleReadonly, // 该值需跟授权时保持一致
	}

	resp, err := ccev2Client.RenewRBAC(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

package rbac

import (
	"encoding/json"
	"fmt"
	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// ListRBAC 根据用户 ID 查询该用户已有的 RBAC 权限
func ListRBAC() {
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
		UserID: "用户 ID", // 仅支持根据用户 ID 查询 RBAC 信息
	}

	resp, err := ccev2Client.ListRBAC(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

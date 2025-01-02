package rbac

import (
	"encoding/json"
	"fmt"

	v2 "github.com/baidubce/bce-sdk-go/services/cce/v2"
	"github.com/baidubce/bce-sdk-go/services/cce/v2/model"
)

// CreateRBACForAllCluster 为指定用户授权所有集群
func CreateRBACForAllCluster() {
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
		ClusterID: model.AllCluster, // 仅对当前已存在的集群生效，后续新增集群不生效。
		UserID:    "待授权的用户 ID",
		Role:      model.RoleReadonly,
	}

	resp, err := ccev2Client.CreateRBAC(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

// CreateRBACForAllNamespace 为指定用户授权指定集群的所有 namepsace
func CreateRBACForAllNamespace() {
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
		ClusterID: "待授权的集群 ID",
		UserID:    "待授权的用户 ID",
		Namespace: model.AllNamespace,
		Role:      model.RoleReadonly,
	}

	resp, err := ccev2Client.CreateRBAC(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

// CreateRBACForOneNamespace 为指定用户授权指定集群的指定 namepsace
func CreateRBACForOneNamespace() {
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
		ClusterID: "待授权的集群 ID",
		UserID:    "待授权的用户 ID",
		Namespace: "待授权的 namespace",
		Role:      model.RoleReadonly,
	}

	resp, err := ccev2Client.CreateRBAC(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Println("Response:" + string(s))
}

// CreateTempKubeConfig 创建临时 kubeconfig
func CreateTempKubeConfig() {
	AK, SK := "", ""

	// 用户指定的endpoint
	ENDPOINT := ""

	// 初始化一个CCEClient
	ccev2Client, err := v2.NewClient(AK, SK, ENDPOINT)
	if err != nil {
		panic(err)
	}

	args := &model.RBACRequest{
		ClusterID:      "集群 id",
		ExpireHours:    1, // 如果需要7天，则填写 168 小时
		Temp:           true,
		KubeConfigType: model.KubeConfigTypeVPC, // vpc 内访问，public 公网访问
	}

	resp, err := ccev2Client.CreateRBAC(args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("临时访问凭证:\r\n" + resp.TemporaryKubeConfig)
}

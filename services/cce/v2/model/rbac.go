package model

const (
	AllCluster   = "all"
	AllNamespace = "all"
)

// RBACRole - RBAC 角色
type RBACRole string

const (
	// RoleAdmin - 管理员权限
	RoleAdmin RBACRole = "cce:admin" // 管理员

	// RoleDevOps - 运维开发权限
	RoleDevOps RBACRole = "cce:devops" // 运维开发

	// RoleReadonly - 只读权限
	RoleReadonly RBACRole = "cce:readonly" // 只读
)

// KubeConfigType - kube config 类型
type KubeConfigType string

const (
	// KubeConfigTypeVPC 使用 BLB VPCIP
	KubeConfigTypeVPC KubeConfigType = "vpc"

	// KubeConfigTypePublic 使用 BLB EIP
	KubeConfigTypePublic KubeConfigType = "public"
)

// RBACResponse - 应答，仅包含 requestID
type RBACResponse struct {
	RequestID string `json:"requestID,omitempty"` // request id
}

// RBACRequest - 创建 RBAC 请求内容
type RBACRequest struct {
	ClusterID string   `json:"clusterID,omitempty"`
	UserID    string   `json:"userID,omitempty"`
	Namespace string   `json:"namespace,omitempty"`
	Role      RBACRole `json:"role,omitempty"`

	Temp           bool           `json:"temp,omitempty"`
	ExpireHours    int            `json:"expireHours,omitempty"`
	KubeConfigType KubeConfigType `json:"kubeConfigType,omitempty"`
}

// CreateRBACResponse - 创建 RBAC 权限接口的返回
type CreateRBACResponse struct {
	RBACResponse
	Data []*CreateRBACMessage `json:"data,omitempty"`

	// 临时访问凭证
	TemporaryKubeConfig string `json:"temporaryKubeConfig,omitempty"`
}

// CreateRBACMessage - 创建RBAC权限接口的主要信息
type CreateRBACMessage struct {
	Success   bool   `json:"success,omitempty"` // 是否创建成功
	ClusterID string `json:"clusterID"`
	Message   string `json:"message,omitempty"` // 创建结果的信息
}

// GetRBACResponse - 获取RBAC权限列表的返回
type GetRBACResponse struct {
	RequestID string            `json:"requestID,omitempty"` // request id
	Data      []*GetRBACMessage `json:"data"`
}

type GetRBACMessage struct {
	Role        RBACRole `json:"role,omitempty"`
	ClusterID   string   `json:"clusterID,omitempty"`
	Namespace   string   `json:"namespace,omitempty"`
	ClusterName string   `json:"clusterName,omitempty"`
}

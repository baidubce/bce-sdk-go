package v2

// PermissionEntry结构体
type PermissionEntry struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Permission string `json:"permission"` // 取值 "r" 或 "rw"
}

// DatasetVersionEntry结构体
type DatasetVersionEntry struct {
	ID          string `json:"id,omitempty"`          // 数据集版本ID，新建时可不填
	Version     string `json:"version,omitempty"`     // 版本号，新建时可不填
	Description string `json:"description,omitempty"` // 版本描述
	StoragePath string `json:"storagePath"`           // 存储路径，必填
	MountPath   string `json:"mountPath"`             // 默认挂载路径，必填
	CreateUser  string `json:"createUser,omitempty"`  // 创建用户
}

// Dataset结构体
type Dataset struct {
	ID              string            `json:"id,omitempty"`              // 数据集ID，新建时可不填
	Name            string            `json:"name"`                      // 数据集名称
	StorageType     string            `json:"storageType"`               // 存储类型
	StorageInstance string            `json:"storageInstance"`           // 存储实例
	ImportFormat    string            `json:"importFormat"`              // 导入格式
	Description     string            `json:"description,omitempty"`     // 描述
	Owner           string            `json:"owner"`                     // 所有者
	OwnerName       string            `json:"ownerName"`                 // 所有者名称
	VisibilityScope string            `json:"visibilityScope"`           // 可见范围
	VisibilityUser  []PermissionEntry `json:"visibilityUser,omitempty"`  // 用户权限列表
	VisibilityGroup []PermissionEntry `json:"visibilityGroup,omitempty"` // 用户组权限列表
	Permission      string            `json:"permission"`                // 当前用户拥有的读写权限
	LatestVersionID string            `json:"latestVersionId,omitempty"` // 最新版本ID
	LatestVersion   string            `json:"latestVersion,omitempty"`   // 最新版本号
	CreatedAt       string            `json:"createdAt"`                 // 创建时间
	UpdatedAt       string            `json:"updatedAt"`                 // 更新时间
}

// CreateDataset
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc091fap
type CreateDatasetRequest struct {
	Name             string              `json:"name"`                      // 数据集名称
	StorageType      string              `json:"storageType"`               // 存储类型
	StorageInstance  string              `json:"storageInstance"`           // 存储实例
	ImportFormat     string              `json:"importFormat"`              // 导入格式
	Description      string              `json:"description,omitempty"`     // 描述
	Owner            string              `json:"owner,omitempty"`           // 所有者
	VisibilityScope  string              `json:"visibilityScope"`           // 可见范围
	VisibilityUser   []PermissionEntry   `json:"visibilityUser,omitempty"`  // 用户权限列表
	VisibilityGroup  []PermissionEntry   `json:"visibilityGroup,omitempty"` // 用户组权限列表
	InitVersionEntry DatasetVersionEntry `json:"initVersionEntry"`          // 初始版本相关信息
}

// CreateDatasetResponse结构体
type CreateDatasetResponse struct {
	RequestID string `json:"requestId"`
	ID        string `json:"id"`
}

// DeleteDataset
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/wmc09407x
type DeleteDatasetResponse struct {
	RequestID string `json:"requestId"`
}

// ModifyDataset
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Imc095v8z
type ModifyDatasetRequest struct {
	Name            string            `json:"name,omitempty"`            // 数据集名称
	Description     string            `json:"description,omitempty"`     // 描述
	VisibilityScope string            `json:"visibilityScope,omitempty"` // 可见范围
	VisibilityUser  []PermissionEntry `json:"visibilityUser,omitempty"`  // 用户权限列表
	VisibilityGroup []PermissionEntry `json:"visibilityGroup,omitempty"` // 用户组权限列表
}

// ModifyDatasetResponse结构体
type ModifyDatasetResponse struct {
	RequestID string `json:"requestId"`
}

// DescribeDataset
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Umc0988jj
type DescribeDatasetResponse struct {
	RequestID          string                `json:"requestId"`
	ID                 string                `json:"id"`
	Name               string                `json:"name"`
	StorageType        string                `json:"storageType"`
	StorageInstance    string                `json:"storageInstance"`
	ImportFormat       string                `json:"importFormat"`
	Description        string                `json:"description"`
	Owner              string                `json:"owner"`
	OwnerName          string                `json:"ownerName"`
	VisibilityScope    string                `json:"visibilityScope"`
	VisibilityUser     []PermissionEntry     `json:"visibilityUser"`
	VisibilityGroup    []PermissionEntry     `json:"visibilityGroup"`
	Permission         string                `json:"permission"`
	LatestVersionID    string                `json:"latestVersionId"`
	LatestVersion      string                `json:"latestVersion"`
	LatestVersionEntry *DatasetVersionDetail `json:"latestVersionEntry"`
	CreatedAt          string                `json:"createdAt"`
	UpdatedAt          string                `json:"updatedAt"`
}

// DatasetVersionDetail结构体
type DatasetVersionDetail struct {
	RequestID      string `json:"requestId"`
	ID             string `json:"id"`
	Version        string `json:"version"`
	Description    string `json:"description"`
	StoragePath    string `json:"storagePath"`
	MountPath      string `json:"mountPath"`
	CreateUser     string `json:"createUser"`
	CreateUserName string `json:"createUserName"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
}

// DescribeDatasets
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Emc099va4
type DescribeDatasetsOptions struct {
	Keyword          string `json:"keyword,omitempty"`          // 名称关键字
	StorageType      string `json:"storageType,omitempty"`      // 存储类型
	StorageInstances string `json:"storageInstances,omitempty"` // 存储实例列表，英文逗号分隔
	ImportFormat     string `json:"importFormat,omitempty"`     // 导入格式
	PageNumber       int    `json:"pageNumber,omitempty"`       // 第几页，默认1
	PageSize         int    `json:"pageSize,omitempty"`         // 单页结果数
}

// DescribeDatasetsResponse结构体
type DescribeDatasetsResponse struct {
	RequestID  string    `json:"requestId"`
	TotalCount int       `json:"totalCount"`
	Datasets   []Dataset `json:"datasets"`
}

// DescribeDatasetVersion
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Dmc09bpj1
type DescribeDatasetVersionResponse struct {
	RequestID       string              `json:"requestId"`
	ID              string              `json:"id"`
	Name            string              `json:"name"`
	StorageType     string              `json:"storageType"`
	StorageInstance string              `json:"storageInstance"`
	ImportFormat    string              `json:"importFormat"`
	Description     string              `json:"description"`
	Owner           string              `json:"owner"`
	OwnerName       string              `json:"ownerName"`
	VisibilityScope string              `json:"visibilityScope"`
	VisibilityUser  []PermissionEntry   `json:"visibilityUser,omitempty"`
	VisibilityGroup []PermissionEntry   `json:"visibilityGroup,omitempty"`
	Permission      string              `json:"permission"`
	VersionEntry    DatasetVersionEntry `json:"versionEntry"`
	CreatedAt       string              `json:"createdAt"`
	UpdatedAt       string              `json:"updatedAt"`
}

// DescribeDatasetVersions
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Tmc09d4k0
type DescribeDatasetVersionsOptions struct {
	PageNumber int `json:"pageNumber,omitempty"`
	PageSize   int `json:"pageSize,omitempty"`
}

// DescribeDatasetVersionsResponse结构体
type DescribeDatasetVersionsResponse struct {
	RequestID  string                `json:"requestId"`
	TotalCount int                   `json:"totalCount"`
	Versions   []DatasetVersionEntry `json:"versions"`
}

// CreateDatasetVersion
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/hmc09en7q
type CreateDatasetVersionRequest struct {
	Description string `json:"description,omitempty"` // 版本描述
	StoragePath string `json:"storagePath"`           // 存储路径
	MountPath   string `json:"mountPath"`             // 默认挂载路径
}

// CreateDatasetVersionResponse结构体
type CreateDatasetVersionResponse struct {
	RequestID string `json:"requestId"`
	ID        string `json:"id"`
}

// DeleteDatasetVersion
// 参考百度智能云AIHC文档 https://cloud.baidu.com/doc/AIHC/s/Omc09gd0f
type DeleteDatasetVersionResponse struct {
	RequestID string `json:"requestId"`
}

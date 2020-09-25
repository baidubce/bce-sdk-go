// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/
/*
CCE V2 版本 GO SDK, Interface 定义
*/

package v2

import (
	"fmt"
	"time"

	"github.com/baidubce/bce-sdk-go/services/cce/v2/types"
	"github.com/baidubce/bce-sdk-go/services/vpc"
)

// Interface 定义 CCE V2 SDK
type Interface interface {
	CreateCluster(args *CreateClusterArgs) (*CreateClusterResponse, error)
	GetCluster(clusterID string) (*GetClusterResponse, error)
	DeleteCluster(args *DeleteClusterArgs) (*DeleteClusterResponse, error)
	ListClusters(args *ListClustersArgs) (*ListClustersResponse, error)

	CreateInstances(args *CreateInstancesArgs) (*CreateInstancesResponse, error)
	GetInstance(args *GetInstanceArgs) (*GetInstanceResponse, error)
	DeleteInstances(args *DeleteInstancesArgs) (*DeleteInstancesResponse, error)
	ListInstancesByPage(args *ListInstancesByPageArgs) (*ListInstancesResponse, error)

	GetClusterQuota() (*GetQuotaResponse, error)
	GetClusterNodeQuota(clusterID string) (*GetQuotaResponse, error)

	CheckContainerNetworkCIDR(args *CheckContainerNetworkCIDRArgs) (*CheckContainerNetworkCIDRResponse, error)
	CheckClusterIPCIDR(args *CheckClusterIPCIDRArgs) (*CheckClusterIPCIDRResponse, error)
	RecommendContainerCIDR(args *RecommendContainerCIDRArgs) (*RecommendContainerCIDRResponse, error)
	RecommendClusterIPCIDR(args *RecommendClusterIPCIDRArgs) (*RecommendClusterIPCIDRResponse, error)
}

//CreateCluterArgs为后续支持clientToken预留空间
type CreateClusterArgs struct {
	CreateClusterRequest *CreateClusterRequest
}

type DeleteClusterArgs struct {
	ClusterID         string
	DeleteResource    bool
	DeleteCDSSnapshot bool
}

type ListClustersArgs struct {
	KeywordType ClusterKeywordType
	Keyword     string
	OrderBy     ClusterOrderBy
	Order       Order
	PageNum     int
	PageSize    int
}

type CreateInstancesArgs struct {
	ClusterID string
	Instances []*InstanceSet
}

type GetInstanceArgs struct {
	ClusterID  string
	InstanceID string
}

type DeleteInstancesArgs struct {
	ClusterID              string
	DeleteInstancesRequest *DeleteInstancesRequest
}

type ListInstancesByPageArgs struct {
	ClusterID string
	Params    *ListInstancesByPageParams
}

// CreateClusterRequest - 创建 Cluster 参数
type CreateClusterRequest struct {
	ClusterSpec *types.ClusterSpec `json:"cluster"`
	MasterSpecs []*InstanceSet     `json:"masters,omitempty"`
	NodeSpecs   []*InstanceSet     `json:"nodes,omitempty"`
}

type InstanceSet struct {
	InstanceSpec types.InstanceSpec `json:"instanceSpec"`
	Count        int                `json:"count"`
}

// ListInstancesByPageParams - 分页查询集群实例列表参数
type ListInstancesByPageParams struct {
	KeywordType InstanceKeywordType `json:"keywordType"`
	Keyword     string              `json:"keyword"`
	OrderBy     InstanceOrderBy     `json:"orderBy"`
	Order       Order               `json:"order"`
	PageNo      int                 `json:"pageNo"`
	PageSize    int                 `json:"pageSize"`
}

// CreateClusterResponse - 创建 Cluster 返回
type CreateClusterResponse struct {
	ClusterID string `json:"clusterID"`
	RequestID string `json:"requestID"`
}

// UpdateClusterResponse - 更新 Cluster 返回
type UpdateClusterResponse struct {
	Cluster   *Cluster `json:"cluster"`
	RequestID string   `json:"requestID"`
}

// GetClusterResponse - 查询 Cluster 返回
type GetClusterResponse struct {
	Cluster   *Cluster `json:"cluster"`
	RequestID string   `json:"requestID"`
}

// DeleteClusterResponse - 删除 Cluster 返回
type DeleteClusterResponse struct {
	RequestID string `json:"requestID"`
}

// ListClustersResponse - List 用户 Cluster 返回
type ListClustersResponse struct {
	ClusterPage *ClusterPage `json:"clusterPage"`
	RequestID   string       `json:"requestID"`
}

// CreateInstancesResponse - 创建 Instances 返回
type CreateInstancesResponse struct {
	CCEInstanceIDs []string `json:"cceInstanceIDs"`
	RequestID      string   `json:"requestID"`
}

type UpdateInstanceArgs struct {
	ClusterID string
	InstanceID string
	InstanceSpec *types.InstanceSpec
}

// UpdateInstancesResponse - 更新 Instances 返回
type UpdateInstancesResponse struct {
	Instance  *Instance `json:"instance"`
	RequestID string    `json:"requestID"`
}

// ClusterPage - 集群分页查询返回
type ClusterPage struct {
	KeywordType ClusterKeywordType `json:"keywordType"`
	Keyword     string             `json:"keyword"`
	OrderBy     ClusterOrderBy     `json:"orderBy"`
	Order       Order              `json:"order"`
	PageNo      int                `json:"pageNo"`
	PageSize    int                `json:"pageSize"`
	TotalCount  int                `json:"totalCount"`
	ClusterList []*Cluster         `json:"clusterList"`
}

// ClusterKeywordType 集群模糊查询字段
type ClusterKeywordType string

const (
	// ClusterKeywordTypeClusterName 集群模糊查询字段: ClusterName
	ClusterKeywordTypeClusterName ClusterKeywordType = "clusterName"
	// ClusterKeywordTypeClusterID 集群模糊查询字段: ClusterID
	ClusterKeywordTypeClusterID ClusterKeywordType = "clusterID"
)

// ClusterOrderBy 集群查询排序字段
type ClusterOrderBy string

const (
	// ClusterOrderByClusterName 集群查询排序字段: ClusterName
	ClusterOrderByClusterName ClusterOrderBy = "clusterName"
	// ClusterOrderByClusterID 集群查询排序字段: ClusterID
	ClusterOrderByClusterID ClusterOrderBy = "clusterID"
	// ClusterOrderByCreatedAt 集群查询排序字段: CreatedAt
	ClusterOrderByCreatedAt ClusterOrderBy = "createdAt"
)

// Order 集群查询排序
type Order string

const (
	// OrderASC 集群查询排序: 升序
	OrderASC Order = "ASC"
	// OrderDESC 集群查询排序: 降序
	OrderDESC Order = "DESC"
)

const (
	// PageNoDefault 分页查询默认页码
	PageNoDefault int = 1
	// PageSizeDefault 分页查询默认页面元素数目
	PageSizeDefault int = 10
)

// GetInstanceResponse - 查询 Instances 返回
type GetInstanceResponse struct {
	Instance  *Instance `json:"instance"`
	RequestID string    `json:"requestID"`
}

// DeleteInstancesResponse - 删除 Instances 返回
type DeleteInstancesResponse struct {
	RequestID string `json:"requestID"`
}

// ListInstancesResponse - List Instances 返回
type ListInstancesResponse struct {
	InstancePage *InstancePage `json:"instancePage"`
	RequestID    string        `json:"requestID"`
}

// GetQuotaResponse - 查询 Quota 返回
type GetQuotaResponse struct {
	types.Quota
	RequestID string `json:"requestID"`
}

// Cluster - Cluster 返回
type Cluster struct {
	Spec   *ClusterSpec   `json:"spec"`
	Status *ClusterStatus `json:"status"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// 作为返回值的ClusterSpec
type ClusterSpec struct {
	ClusterID   string            `json:"clusterID"`
	ClusterName string            `json:"clusterName"`
	ClusterType types.ClusterType `json:"clusterType"`

	Description string `json:"description"`

	K8SVersion types.K8SVersion `json:"k8sVersion"`

	VPCID   string `json:"vpcID"`
	VPCCIDR string `json:"vpcCIDR"`

	Plugins []string `json:"plugins"`

	MasterConfig           types.MasterConfig           `json:"masterConfig"`
	ContainerNetworkConfig types.ContainerNetworkConfig `json:"containerNetworkConfig"`
}

// ClusterStatus - Cluster Status
type ClusterStatus struct {
	ClusterBLB BLB `json:"clusterBLB"`

	ClusterPhase types.ClusterPhase `json:"clusterPhase"`

	NodeNum int `json:"nodeNum"`
}

// BLB 定义 BLB 类型
type BLB struct {
	ID    string `json:"id"`
	VPCIP string `json:"vpcIP"`
	EIP   string `json:"eip"`
}

// InstancePage - 节点分页查询返回
type InstancePage struct {
	ClusterID    string              `json:"clusterID"`
	KeywordType  InstanceKeywordType `json:"keywordType"`
	Keyword      string              `json:"keyword"`
	OrderBy      InstanceOrderBy     `json:"orderBy"`
	Order        Order               `json:"order"`
	PageNo       int                 `json:"pageNo"`
	PageSize     int                 `json:"pageSize"`
	TotalCount   int                 `json:"totalCount"`
	InstanceList []*Instance         `json:"instanceList"`
}

// InstanceKeywordType 节点模糊查询字段
type InstanceKeywordType string

const (
	// InstanceKeywordTypeInstanceName 节点模糊查询字段: InstanceName
	InstanceKeywordTypeInstanceName InstanceKeywordType = "instanceName"
	// InstanceKeywordTypeInstanceID 节点模糊查询字段: InstanceID
	InstanceKeywordTypeInstanceID InstanceKeywordType = "instanceID"
)

// InstanceOrderBy 节点查询排序字段
type InstanceOrderBy string

const (
	// InstanceOrderByInstanceName 节点查询排序字段: InstanceName
	InstanceOrderByInstanceName InstanceOrderBy = "instanceName"
	// InstanceOrderByInstanceID 节点查询排序字段: InstanceID
	InstanceOrderByInstanceID InstanceOrderBy = "instanceID"
	// InstanceOrderByCreatedAt 节点查询排序字段: CreatedAt
	InstanceOrderByCreatedAt InstanceOrderBy = "createdAt"
)

// Instance - 节点详情
// 作为sdk返回结果的Instance
type Instance struct {
	Spec   *types.InstanceSpec   `json:"spec"`
	Status *InstanceStatus `json:"status"`

	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// InstanceStatus - Instance Status
type InstanceStatus struct {
	Machine Machine `json:"machine"`

	InstancePhase types.InstancePhase `json:"instancePhase"`
	MachineStatus types.ServerStatus  `json:"machineStatus"`
}

// Machine - 定义机器相关信息
type Machine struct {
	InstanceID string `json:"instanceID"`

	OrderID string `json:"orderID,omitempty"`

	MountList []types.MountConfig `json:"mountList,omitempty"`

	VPCIP     string `json:"vpcIP,omitempty"`
	VPCIPIPv6 string `json:"vpcIPIPv6,omitempty"`

	EIP string `json:"eip,omitempty"`
}

// DeleteInstancesRequest - 删除节点请求
type DeleteInstancesRequest struct {
	InstanceIDs  []string            `json:"instanceIDs,omitempty"`
	DeleteOption *types.DeleteOption `json:"deleteOption,omitempty"`
}

// InstanceKeyType - ListInstanceByPage 参数
type InstanceKeyType string

// NetworkConflictType 冲突类型
type NetworkConflictType string

const (
	// ContainerCIDRAndNodeCIDRConflict 容器网段和本集群的节点网段冲突
	ContainerCIDRAndNodeCIDRConflict NetworkConflictType = "ContainerCIDRAndNodeCIDR"
	// ContainerCIDRAndExistedClusterContainerCIDRConflict 容器网段和 VPC 内已有集群的容器网段冲突
	ContainerCIDRAndExistedClusterContainerCIDRConflict NetworkConflictType = "ContainerCIDRAndExistedClusterContainerCIDR"
	// ContainerCIDRAndVPCRouteConflict 容器网段与 VPC 路由冲突
	ContainerCIDRAndVPCRouteConflict NetworkConflictType = "ContainerCIDRAndVPCRoute"
	// ClusterIPCIDRAndNodeCIDRConflict ClusterIP 网段与本集群节点网段冲突
	ClusterIPCIDRAndNodeCIDRConflict NetworkConflictType = "ClusterIPCIDRAndNodeCIDR"
	// ClusterIPCIDRAndContainerCIDRConflict ClusterIP 网段与本集群容器网段冲突
	ClusterIPCIDRAndContainerCIDRConflict NetworkConflictType = "ClusterIPCIDRAndContainerCIDR"
)

// PrivateNetString IPv4/IPv6 私有网络地址类型
type PrivateNetString string

const (
	// PrivateIPv4Net10 - IPv4 10 段
	PrivateIPv4Net10 PrivateNetString = "10.0.0.0/8"

	// PrivateIPv4Net172 - IPv4 172 段
	PrivateIPv4Net172 PrivateNetString = "172.16.0.0/12"

	// PrivateIPv4Net192 - IPv4 192 段
	PrivateIPv4Net192 PrivateNetString = "192.168.0.0/16"

	// PrivateIPv6Net - IPv6 段
	PrivateIPv6Net PrivateNetString = "fc00::/7"
)

const (
	// MaxClusterIPServiceNum 集群最大的 ClusterIP Service 数量
	MaxClusterIPServiceNum = 65536
)

// CheckContainerNetworkCIDRRequest 包含检查容器网络网段冲突的请求参数
type CheckContainerNetworkCIDRArgs struct {
	VPCID             string                       `json:"vpcID"`
	VPCCIDR           string                       `json:"vpcCIDR"`
	VPCCIDRIPv6       string                       `json:"vpcCIDRIPv6"`
	ContainerCIDR     string                       `json:"containerCIDR"`
	ContainerCIDRIPv6 string                       `json:"containerCIDRIPv6"`
	ClusterIPCIDR     string                       `json:"clusterIPCIDR"`
	ClusterIPCIDRIPv6 string                       `json:"clusterIPCIDRIPv6"`
	MaxPodsPerNode    int                          `json:"maxPodsPerNode"`
	IPVersion         types.ContainerNetworkIPType `json:"ipVersion"` // if not set, set ipv4
}

// CheckClusterIPCIDRequest - 检查 ClusterIP CIDR 请求
type CheckClusterIPCIDRArgs struct {
	VPCID             string                       `json:"vpcID"`
	VPCCIDR           string                       `json:"vpcCIDR"`
	VPCCIDRIPv6       string                       `json:"vpcCIDRIPv6"`
	ClusterIPCIDR     string                       `json:"clusterIPCIDR"`
	ClusterIPCIDRIPv6 string                       `json:"clusterIPCIDRIPv6"`
	IPVersion         types.ContainerNetworkIPType `json:"ipVersion"` // if not set, set ipv4
}

// CheckContainerNetworkCIDRResponse 检查容器网络网段冲突的响应
type CheckContainerNetworkCIDRResponse struct {
	MaxNodeNum int `json:"maxNodeNum"`
	NetworkConflictInfo
	RequestID string `json:"requestID"`
}

// CheckClusterIPCIDRResponse - 检查 ClusterIP CIDR 返回
type CheckClusterIPCIDRResponse struct {
	IsConflict bool   `json:"isConflict"`
	ErrMsg     string `json:"errMsg"`
	RequestID  string `json:"requestID"`
}

// RecommendContainerCIDRRequest 推荐容器网段的请求参数
type RecommendContainerCIDRArgs struct {
	VPCID       string `json:"vpcID"`
	VPCCIDR     string `json:"vpcCIDR"`
	VPCCIDRIPv6 string `json:"vpcCIDRIPv6"`
	// ClusterMaxNodeNum 集群节点的最大规模
	ClusterMaxNodeNum int `json:"clusterMaxNodeNum"`
	MaxPodsPerNode    int `json:"maxPodsPerNode"`
	// PrivateNetCIDRs 候选的容器网段列表，只能从 [10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16] 里选择
	PrivateNetCIDRs     []PrivateNetString           `json:"privateNetCIDRs"`
	PrivateNetCIDRIPv6s []PrivateNetString           `json:"privateNetCIDRIPv6s"`
	K8SVersion          types.K8SVersion             `json:"k8sVersion"`
	IPVersion           types.ContainerNetworkIPType `json:"ipVersion"` // if not set, set ipv4
}

// RecommendContainerCIDRResponse 推荐容器网段的响应
type RecommendContainerCIDRResponse struct {
	RecommendedContainerCIDRs     []string `json:"recommendedContainerCIDRs"`
	RecommendedContainerCIDRIPv6s []string `json:"recommendedContainerCIDRIPv6s"`
	IsSuccess                     bool     `json:"isSuccess"`
	ErrMsg                        string   `json:"errMsg"`
	RequestID                     string   `json:"requestID"`
}

// RecommendClusterIPCIDRRequest 推荐 ClusterIP 网段的请求参数
type RecommendClusterIPCIDRArgs struct {
	VPCCIDR           string `json:"vpcCIDR"`
	VPCCIDRIPv6       string `json:"vpcCIDRIPv6"`
	ContainerCIDR     string `json:"containerCIDR"`
	ContainerCIDRIPv6 string `json:"containerCIDRIPv6"`
	// ClusterMaxServiceNum 集群 Service 最大规模
	ClusterMaxServiceNum int `json:"clusterMaxServiceNum"`
	// PrivateNetCIDRs 候选的 ClusterIP 网段列表，只能从 [10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16] 里选择
	PrivateNetCIDRs     []PrivateNetString           `json:"privateNetCIDRs"`
	PrivateNetCIDRIPv6s []PrivateNetString           `json:"privateNetCIDRIPv6s"`
	IPVersion           types.ContainerNetworkIPType `json:"ipVersion"` // if not set, set ipv4
}

// RecommendClusterIPCIDRResponse 推荐 ClusterIP 网段的响应
type RecommendClusterIPCIDRResponse struct {
	RecommendedClusterIPCIDRs     []string `json:"recommendedClusterIPCIDRs"`
	RecommendedClusterIPCIDRIPv6s []string `json:"recommendedClusterIPCIDRIPv6s"`
	IsSuccess                     bool     `json:"isSuccess"`
	ErrMsg                        string   `json:"errMsg"`
	RequestID                     string   `json:"requestID"`
}

// NetworkConflictInfo 容器网络整体配置冲突信息
type NetworkConflictInfo struct {
	IsConflict            bool                   `json:"isConflict"`
	ErrMsg                string                 `json:"errMsg"`
	ContainerCIDRConflict *ContainerCIDRConflict `json:"containerCIDRConflict"` // 容器网段冲突信息
	ClusterIPCIDRConflict *ClusterIPCIDRConflict `json:"clusterIPCIDRConflict"` // ClusterIP 网段冲突信息
}

// ContainerCIDRConflict 容器网段冲突信息
type ContainerCIDRConflict struct {
	// NetworkConflictType 冲突类型，可取的值： ContainerCIDRAndNodeCIDRConflict、ContainerCIDRAndExistedClusterContainerCIDRConflict、ContainerCIDRAndVPCRouteConflict
	ConflictType NetworkConflictType `json:"conflictType"`
	// ConflictNodeCIDR 与容器网段冲突的节点网段，当且仅当 NetworkConflictType 为 ContainerCIDRAndNodeCIDRConflict 不为 nil
	ConflictNodeCIDR *ConflictNodeCIDR `json:"conflictNodeCIDR"`
	// ConflictCluster 与容器网段冲突的VPC内集群，当且仅当 NetworkConflictType 为 ContainerCIDRAndExistedClusterContainerCIDRConflict 不为 nil
	ConflictCluster *ConflictCluster `json:"conflictCluster"`
	// ConflictVPCRoute 与容器网段冲突的VPC路由，当且仅当 NetworkConflictType 为 ContainerCIDRAndVPCRouteConflict 不为 nil
	ConflictVPCRoute *ConflictVPCRoute `json:"conflictVPCRoute"`
}

// ClusterIPCIDRConflict ClusterIP 网段冲突信息
type ClusterIPCIDRConflict struct {
	// NetworkConflictType 冲突类型，可取的值： ClusterIPCIDRAndNodeCIDRConflict、ClusterIPCIDRAndContainerCIDRConflict
	ConflictType NetworkConflictType `json:"conflictType"`
	// ConflictNodeCIDR 与 ClusterIP 网段冲突的节点网段，当且仅当 NetworkConflictType 为 ClusterIPCIDRAndNodeCIDRConflict 不为 nil
	ConflictNodeCIDR *ConflictNodeCIDR `json:"conflictNodeCIDR"`
	// ConflictContainerCIDR 与 ClusterIP 网段冲突的节点网段，当且仅当 NetworkConflictType 为 ClusterIPCIDRAndContainerCIDRConflict 不为 nil
	ConflictContainerCIDR *ConflictContainerCIDR `json:"conflictContainerCIDR"`
}

// ConflictNodeCIDR 节点网段冲突信息
type ConflictNodeCIDR struct {
	NodeCIDR string `json:"nodeCIDR"`
}

// ConflictContainerCIDR 容器网段冲突信息
type ConflictContainerCIDR struct {
	ContainerCIDR string `json:"containerCIDR"`
}

// ConflictCluster 同一 VPC 内容器网段冲突的集群信息
type ConflictCluster struct {
	ClusterID     string `json:"clusterID"`
	ContainerCIDR string `json:"containerCIDR"`
}

// ConflictVPCRoute 冲突的 VPC 路由
type ConflictVPCRoute struct {
	RouteRule vpc.RouteRule `json:"routeRule"`
}

type InstanceTemplate struct {
	types.InstanceSpec `json:",inline"`
}

type CommonResponse struct {
	RequestID string `json:"requestID"`
}

type InstanceGroup struct {
	Spec      *InstanceGroupSpec   `json:"spec"`
	Status    *InstanceGroupStatus `json:"status"`
	CreatedAt time.Time            `json:"createdAt"`
}

type InstanceGroupSpec struct {
	CCEInstanceGroupID string `json:"cceInstanceGroupID,omitempty"`
	InstanceGroupName  string `json:"instanceGroupName"`

	ClusterID    string            `json:"clusterID,omitempty"`
	ClusterRole  types.ClusterRole `json:"clusterRole,omitempty"`
	ShrinkPolicy ShrinkPolicy      `json:"shrinkPolicy,omitempty"`
	UpdatePolicy UpdatePolicy      `json:"updatePolicy,omitempty"`
	CleanPolicy  CleanPolicy       `json:"cleanPolicy,omitempty"`

	InstanceTemplate InstanceTemplate `json:"instanceTemplate"`
	Replicas         int              `json:"replicas"`

	ClusterAutoscalerSpec *ClusterAutoscalerSpec `json:"clusterAutoscalerSpec,omitempty"`
}

type ShrinkPolicy string
type UpdatePolicy string
type CleanPolicy string

type ClusterAutoscalerSpec struct {
	Enabled              bool `json:"enabled"`
	MinReplicas          int  `json:"minReplicas"`
	MaxReplicas          int  `json:"maxReplicas"`
	ScalingGroupPriority int  `json:"scalingGroupPriority"`
}

// InstanceGroupStatus -
type InstanceGroupStatus struct {
	ReadyReplicas int          `json:"readyReplicas"`
	Pause         *PauseDetail `json:"pause,omitempty"`
}

type PauseDetail struct {
	Paused bool   `json:"paused"`
	Reason string `json:"reason"`
}

// CreateInstanceGroupRequest - 创建InstanceGroup request
type CreateInstanceGroupRequest struct {
	types.InstanceGroupSpec
}

// CreateInstanceGroupResponse - 创建InstanceGroup response
type CreateInstanceGroupResponse struct {
	CommonResponse
	InstanceGroupID string `json:"instanceGroupID"`
}

type ListInstanceGroupResponse struct {
	CommonResponse
	Page ListInstanceGroupPage `json:"page"`
}

type ListInstanceGroupPage struct {
	PageNo     int              `json:"pageNo"`
	PageSize   int              `json:"pageSize"`
	TotalCount int              `json:"totalCount"`
	List       []*InstanceGroup `json:"list"`
}

type GetInstanceGroupResponse struct {
	CommonResponse
	InstanceGroup *InstanceGroup `json:"instanceGroup"`
}

type UpdateInstanceGroupReplicasRequest struct {
	Replicas       int                 `json:"replicas"`
	InstanceIDs    []string            `json:"instanceIDs"`
	DeleteInstance bool                `json:"deleteInstance"`
	DeleteOption   *types.DeleteOption `json:"deleteOption,omitempty"`
}

type UpdateInstanceGroupReplicasResponse struct {
	CommonResponse
}

type UpdateInstanceGroupClusterAutoscalerSpecResponse struct {
	CommonResponse
}

type DeleteInstanceGroupResponse struct {
	CommonResponse
}

type ListInstancesByInstanceGroupIDPage struct {
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	List       []*Instance `json:"list"`
}

type ListInstancesByInstanceGroupIDResponse struct {
	CommonResponse
	Page ListInstancesByInstanceGroupIDPage `json:"page"`
}

type GetAutoscalerArgs struct {
	ClusterID string
}

type GetAutoscalerResponse struct {
	Autoscaler *Autoscaler `json:"autoscaler"`
	RequestID  string      `json:"requestID"`
}

type UpdateAutoscalerArgs struct {
	ClusterID        string
	AutoscalerConfig ClusterAutoscalerConfig
}

type UpdateAutoscalerResponse struct {
	CommonResponse
}

type CreateAutoscalerArgs struct {
	ClusterID string
}

type CreateAutoscalerResponse struct {
	CommonResponse
}

type Autoscaler struct {
	ClusterID   string                  `json:"clusterID"`
	ClusterName string                  `json:"clusterName"`
	CAConfig    ClusterAutoscalerConfig `json:"caConfig,omitempty"`
}

type ClusterAutoscalerConfig struct {
	KubeVersion    string                           `json:"kubeVersion,omitempty"`
	ReplicaCount   int                              `json:"replicaCount"`
	InstanceGroups []ClusterAutoscalerInstanceGroup `json:"instanceGroups,omitempty"`
	// default: false
	ScaleDownEnabled bool `json:"scaleDownEnabled"`
	// 可选，缩容阈值百分比，范围(0, 100)
	ScaleDownUtilizationThreshold *int `json:"scaleDownUtilizationThreshold,omitempty"`
	// 可选，GPU缩容阈值百分比，范围(0, 100)
	ScaleDownGPUUtilizationThreshold *int `json:"scaleDownGPUUtilizationThreshold,omitempty"`
	// 可选，缩容触发时延，单位：m
	ScaleDownUnneededTime *int `json:"scaleDownUnneededTime,omitempty"`
	// 可选，扩容后缩容启动时延，单位：m
	ScaleDownDelayAfterAdd *int `json:"scaleDownDelayAfterAdd,omitempty"`
	// 可选，最大并发缩容数
	MaxEmptyBulkDelete *int `json:"maxEmptyBulkDelete,omitempty"`
	// 可选，
	SkipNodesWithLocalStorage *bool `json:"skipNodesWithLocalStorage,omitempty"`
	// 可选，
	SkipNodesWithSystemPods *bool `json:"skipNodesWithSystemPods,omitempty"`
	// supported: random, most-pods, least-waste, priority; default: random
	Expander string `json:"expander"`
}

type ClusterAutoscalerInstanceGroup struct {
	InstanceGroupID string
	MinReplicas     int
	MaxReplicas     int
	Priority        int
}

type InstanceGroupListOption struct {
	PageNo   int
	PageSize int
}

type CreateInstanceGroupArgs struct {
	ClusterID string
	Request   *CreateInstanceGroupRequest
}

type ListInstanceGroupsArgs struct {
	ClusterID  string
	ListOption *InstanceGroupListOption
}

type ListInstanceByInstanceGroupIDArgs struct {
	ClusterID       string
	InstanceGroupID string
	PageNo          int
	PageSize        int
}

type GetInstanceGroupArgs struct {
	ClusterID       string
	InstanceGroupID string
}

type UpdateInstanceGroupClusterAutoscalerSpecArgs struct {
	ClusterID       string
	InstanceGroupID string
	Request         *ClusterAutoscalerSpec
}

type UpdateInstanceGroupReplicasArgs struct {
	ClusterID       string
	InstanceGroupID string
	Request         *UpdateInstanceGroupReplicasRequest
}

type DeleteInstanceGroupArgs struct {
	ClusterID       string
	InstanceGroupID string
	DeleteInstances bool
}

// KubeConfigType - kube config 类型
type KubeConfigType string

const (
	// KubeConfigTypeInternal 使用 BLB FloatingIP
	KubeConfigTypeInternal KubeConfigType = "internal"

	// KubeConfigTypeVPC 使用 BLB VPCIP
	KubeConfigTypeVPC KubeConfigType = "vpc"

	// KubeConfigTypePublic 使用 BLB EIP
	KubeConfigTypePublic KubeConfigType = "public"
)

type GetKubeConfigArgs struct {
	ClusterID      string
	KubeConfigType KubeConfigType
}

// GetKubeConfigResponse - 查询 KubeConfig 返回
type GetKubeConfigResponse struct {
	KubeConfigType KubeConfigType `json:"kubeConfigType"`
	KubeConfig     string         `json:"kubeConfig"`
	RequestID      string         `json:"requestID"`
}

func CheckKubeConfigType(kubeConfigType string) error {
	if kubeConfigType != string(KubeConfigTypePublic) &&
		kubeConfigType != string(KubeConfigTypeInternal) &&
		kubeConfigType != string(KubeConfigTypeVPC) {
		return fmt.Errorf("KubeConfigType %s not valid", kubeConfigType)
	}
	return nil
}

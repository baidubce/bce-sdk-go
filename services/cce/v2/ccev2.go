package v2

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// 创建集群
func (c *Client) CreateCluster(args *CreateClusterArgs) (*CreateClusterResponse, error) {
	if args == nil || args.CreateClusterRequest == nil {
		return nil, fmt.Errorf("args is nil")
	}

	//给其中可能存在的user script用base64编码
	if err := encodeUserScriptInInstanceSet(args.CreateClusterRequest.MasterSpecs); err != nil {
		return nil, err
	}

	if err := encodeUserScriptInInstanceSet(args.CreateClusterRequest.NodeSpecs); err != nil {
		return nil, err
	}

	result := &CreateClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterURI()).
		WithBody(args.CreateClusterRequest).
		WithResult(result).
		Do()

	return result, err
}

//删除集群
func (c *Client) DeleteCluster(args *DeleteClusterArgs) (*DeleteClusterResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &DeleteClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getClusterUriWithIDURI(args.ClusterID)).
		WithQueryParamFilter("deleteResource", strconv.FormatBool(args.DeleteResource)).
		WithQueryParamFilter("deleteCDSSnapshot", strconv.FormatBool(args.DeleteCDSSnapshot)).
		WithResult(result).
		Do()

	return result, err
}

//获得集群详情
func (c *Client) GetCluster(clusterID string) (*GetClusterResponse, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("clusterID is empty")
	}

	result := &GetClusterResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterUriWithIDURI(clusterID)).
		WithResult(result).
		Do()

	return result, err
}

//集群列表
func (c *Client) ListClusters(args *ListClustersArgs) (*ListClustersResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.PageNum <= 0 || args.PageSize <= 0 {
		return nil, fmt.Errorf("invlaid pageNo or pageSize")
	}

	result := &ListClustersResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterListURI()).
		WithQueryParamFilter("keywordType", string(args.KeywordType)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("orderBy", string(args.OrderBy)).
		WithQueryParamFilter("order", string(args.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNum)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()

	return result, err
}

//创建节点（扩容）
func (c *Client) CreateInstances(args *CreateInstancesArgs) (*CreateInstancesResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	//给其中可能存在的user script用base64编码
	if err := encodeUserScriptInInstanceSet(args.Instances); err != nil {
		return nil, err
	}

	s, _ := json.MarshalIndent(args, "", "\t")
	fmt.Println("Args:" + string(s))

	result := &CreateInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterInstanceListURI(args.ClusterID)).
		WithBody(args.Instances).
		WithResult(result).
		Do()

	return result, err
}

//查询节点
func (c *Client) GetInstance(args *GetInstanceArgs) (*GetInstanceResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &GetInstanceResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterInstanceURI(args.ClusterID, args.InstanceID)).
		WithResult(result).
		Do()

	return result, err
}

//更新节点配置
func (c *Client) UpdateInstance(args *UpdateInstanceArgs) (*UpdateInstancesResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &UpdateInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getClusterInstanceURI(args.ClusterID, args.InstanceID)).
		WithBody(args.InstanceSpec).
		WithResult(result).
		Do()

	return result, err
}

//删除节点（缩容）
func (c *Client) DeleteInstances(args *DeleteInstancesArgs) (*DeleteInstancesResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &DeleteInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getClusterInstanceListURI(args.ClusterID)).
		WithBody(args.DeleteInstancesRequest).
		WithResult(result).
		Do()

	return result, err
}

//集群内节点列表
func (c *Client) ListInstancesByPage(args *ListInstancesByPageArgs) (*ListInstancesResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &ListInstancesResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterInstanceListURI(args.ClusterID)).
		WithQueryParamFilter("keywordType", string(args.Params.KeywordType)).
		WithQueryParamFilter("keyword", args.Params.Keyword).
		WithQueryParamFilter("orderBy", string(args.Params.OrderBy)).
		WithQueryParamFilter("order", string(args.Params.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.Params.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.Params.PageSize)).
		WithResult(result).
		Do()

	return result, err
}

//检查容器网络网段
func (c *Client) CheckContainerNetworkCIDR(args *CheckContainerNetworkCIDRArgs) (*CheckContainerNetworkCIDRResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("CheckContainerNetworkCIDRRequest is nil")
	}

	result := &CheckContainerNetworkCIDRResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNetCheckContainerNetworkCIDRURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

//检查集群网络网段
func (c *Client) CheckClusterIPCIDR(args *CheckClusterIPCIDRArgs) (*CheckClusterIPCIDRResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &CheckClusterIPCIDRResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNetCheckClusterIPCIDRURL()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

//推荐容器CIDR
func (c *Client) RecommendContainerCIDR(args *RecommendContainerCIDRArgs) (*RecommendContainerCIDRResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &RecommendContainerCIDRResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNetRecommendContainerCidrURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

//推荐集群CIDR
func (c *Client) RecommendClusterIPCIDR(args *RecommendClusterIPCIDRArgs) (*RecommendClusterIPCIDRResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &RecommendClusterIPCIDRResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getNetRecommendClusterIpCidrURI()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

//用户集群 Quota
func (c *Client) GetClusterQuota() (*GetQuotaResponse, error) {
	result := &GetQuotaResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuotaURI()).
		WithResult(result).
		Do()

	return result, err
}

//用户集群 Node Quota
func (c *Client) GetClusterNodeQuota(clusterID string) (*GetQuotaResponse, error) {
	if clusterID == "" {
		return nil, fmt.Errorf("clusterID is empty")
	}

	result := &GetQuotaResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getQuotaNodeURI(clusterID)).
		WithResult(result).
		Do()

	return result, err
}

//创建节点组
func (c *Client) CreateInstanceGroup(args *CreateInstanceGroupArgs) (*CreateInstanceGroupResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	encodeUserScript(&args.Request.InstanceTemplate.InstanceSpec)

	result := &CreateInstanceGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getInstanceGroupURI(args.ClusterID)).
		WithBody(args.Request).
		WithResult(result).
		Do()

	return result, err
}

//获取节点组列表
func (c *Client) ListInstanceGroups(args *ListInstanceGroupsArgs) (*ListInstanceGroupResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &ListInstanceGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.ListOption.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.ListOption.PageSize)).
		WithURL(getInstanceGroupListURI(args.ClusterID)).
		WithResult(result).
		Do()

	return result, err
}

//获取节点组的节点列表
func (c *Client) ListInstancesByInstanceGroupID(args *ListInstanceByInstanceGroupIDArgs) (*ListInstancesByInstanceGroupIDResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &ListInstancesByInstanceGroupIDResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithURL(getClusterInstanceListWithInstanceGroupIDURI(args.ClusterID, args.InstanceGroupID)).
		WithResult(result).
		Do()

	return result, err
}

//获取节点组详情
func (c *Client) GetInstanceGroup(args *GetInstanceGroupArgs) (*GetInstanceGroupResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &GetInstanceGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getInstanceGroupWithIDURI(args.ClusterID, args.InstanceGroupID)).
		WithResult(result).
		Do()

	return result, err
}

//更新节点组副本数
func (c *Client) UpdateInstanceGroupReplicas(args *UpdateInstanceGroupReplicasArgs) (*UpdateInstanceGroupReplicasResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &UpdateInstanceGroupReplicasResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceGroupReplicasURI(args.ClusterID, args.InstanceGroupID)).
		WithBody(args.Request).
		WithResult(result).
		Do()

	return result, err
}

//修改节点组节点Autoscaler配置
func (c *Client) UpdateInstanceGroupClusterAutoscalerSpec(args *UpdateInstanceGroupClusterAutoscalerSpecArgs) (*UpdateInstanceGroupClusterAutoscalerSpecResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.Request == nil {
		return nil, fmt.Errorf("nil UpdateInstanceGroupReplicasRequest")
	}
	if args.Request.Enabled {
		if args.Request.MinReplicas < 0 || args.Request.MaxReplicas < args.Request.MinReplicas {
			return nil, fmt.Errorf("invalid minReplicas or maxReplicas")
		}
		if args.Request.ScalingGroupPriority < 0 {
			return nil, fmt.Errorf("invalid scalingGroupPriority")
		}
	}

	result := &UpdateInstanceGroupClusterAutoscalerSpecResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getInstanceGroupAutoScalerURI(args.ClusterID, args.InstanceGroupID)).
		WithBody(args.Request).
		WithResult(result).
		Do()

	return result, err
}

//删除节点组
func (c *Client) DeleteInstanceGroup(args *DeleteInstanceGroupArgs) (*DeleteInstanceGroupResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &DeleteInstanceGroupResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getInstanceGroupWithIDURI(args.ClusterID, args.InstanceGroupID)).
		WithResult(result).
		Do()

	return result, err
}

//创建autoscaler配置
func (c *Client) CreateAutoscaler(args *CreateAutoscalerArgs) (*CreateAutoscalerResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &CreateAutoscalerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAutoscalerURI(args.ClusterID)).
		WithResult(result).
		Do()

	return result, err
}

//查询autoscaler配置
func (c *Client) GetAutoscaler(args *GetAutoscalerArgs) (*GetAutoscalerResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &GetAutoscalerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAutoscalerURI(args.ClusterID)).
		WithResult(result).
		Do()

	return result, err
}

//更新autoscaler配置
func (c *Client) UpdateAutoscaler(args *UpdateAutoscalerArgs) (*UpdateAutoscalerResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}

	result := &UpdateAutoscalerResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAutoscalerURI(args.ClusterID)).
		WithBody(args.AutoscalerConfig).
		WithResult(result).
		Do()

	return result, err
}

//获取kubeconfig
func (c *Client) GetKubeConfig(args *GetKubeConfigArgs) (*GetKubeConfigResponse, error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if err := CheckKubeConfigType(string(args.KubeConfigType)); err != nil {
		return nil, err
	}

	result := &GetKubeConfigResponse{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getKubeconfigURI(args.ClusterID, args.KubeConfigType)).
		WithResult(result).
		Do()

	return result, err
}
